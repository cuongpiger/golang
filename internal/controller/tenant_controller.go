/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	multitenancyv1 "github.com/cuongpiger/golang/api/v1"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	rbacv1 "k8s.io/api/rbac/v1"
)

const (
	finalizerName = "tenant.codereliant.io/finalizer"
)

// TenantReconciler reconciles a Tenant object
type TenantReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=multitenancy.codereliant.io,resources=*,verbs=*
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=*
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=*,verbs=*

func (r *TenantReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	tenant := &multitenancyv1.Tenant{}

	log.Info("Reconciling tenant")
	if err := r.Get(ctx, req.NamespacedName, tenant); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if tenant.ObjectMeta.DeletionTimestamp.IsZero() {
		// Add a finalizer if not present
		if !controllerutil.ContainsFinalizer(tenant, finalizerName) {
			tenant.ObjectMeta.Finalizers = append(tenant.ObjectMeta.Finalizers, finalizerName)
			if err := r.Update(ctx, tenant); err != nil {
				log.Error(err, "unable to update Tenant")
				return ctrl.Result{}, err
			}
		}

		// Reconciliation logic for creating and managing namespaces
		for _, ns := range tenant.Spec.Namespaces {
			log.Info("Ensuring Namespace", "namespace", ns)
			if err := r.ensureNamespace(ctx, tenant, ns); err != nil {
				log.Error(err, "unable to ensure Namespace", "namespace", ns)
				return ctrl.Result{}, err
			}

			log.Info("Ensuring Admin RoleBinding", "namespace", ns)
			if err := r.ensureRoleBinding(ctx, ns, tenant.Spec.AdminGroups, "admin"); err != nil {
				log.Error(err, "unable to ensure Admin RoleBinding", "namespace", ns)
				return ctrl.Result{}, err
			}

			if err := r.ensureRoleBinding(ctx, ns, tenant.Spec.UserGroups, "edit"); err != nil {
				log.Error(err, "unable to ensure User RoleBinding", "namespace", ns)
				return ctrl.Result{}, err
			}
		}

		tenant.Status.NamespaceCount = len(tenant.Spec.Namespaces)
		tenant.Status.AdminEmail = tenant.Spec.AdminEmail
		if err := r.Update(ctx, tenant); err != nil {
			log.Error(err, "unable to update Tenant status")
			return ctrl.Result{}, err
		}
	} else {
		// Check if the finalizer is present
		if controllerutil.ContainsFinalizer(tenant, finalizerName) {
			log.Info("Finalizer found, cleaning up resources")

			// Cleanup Resources
			if err := r.deleteExternalResources(ctx, tenant); err != nil {
				// retry if failed
				log.Error(err, "Failed to cleanup resources")
				return ctrl.Result{}, err
			}
			log.Info("Resource cleanup succeeded")

			// Remove the finalizer from the Tenant object once the cleanup succeded
			// This will free up tenant resource to be deleted
			controllerutil.RemoveFinalizer(tenant, finalizerName)
			if err := r.Update(ctx, tenant); err != nil {
				log.Error(err, "Unable to remove finalizer and update Tenant")
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TenantReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&multitenancyv1.Tenant{}).
		Complete(r)
}

const (
	tenantOperatorAnnotation = "tenant-operator"
)

func (r *TenantReconciler) ensureNamespace(ctx context.Context, tenant *multitenancyv1.Tenant, namespaceName string) error {
	log := log.FromContext(ctx)

	// Define a namespace object
	namespace := &corev1.Namespace{}

	// Attempt to get the namespace with the provided name
	err := r.Get(ctx, client.ObjectKey{Name: namespaceName}, namespace)
	if err != nil {
		// If the namespace doesn't exist, create it
		if apierrors.IsNotFound(err) {
			log.Info("Creating Namespace", "namespace", namespaceName)
			namespace := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: namespaceName,
					Annotations: map[string]string{
						"adminEmail": tenant.Spec.AdminEmail,
						"managed-by": tenantOperatorAnnotation,
					},
				},
			}

			// Attempt to create the namespace
			if err = r.Create(ctx, namespace); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// If the namespace already exists, check for required annotations
		log.Info("Namespace already exists", "namespace", namespaceName)

		// If the namespace does not have any annotations, initialize the annotations map
		if namespace.Annotations == nil {
			namespace.Annotations = map[string]string{}
		}

		// Define required annotations and their desired values
		requiredAnnotations := map[string]string{
			"adminEmail": tenant.Spec.AdminEmail,
			"managed-by": tenantOperatorAnnotation,
		}

		// Iterate over the required annotations and update if necessary
		for annotationKey, desiredValue := range requiredAnnotations {
			existingValue, ok := namespace.Annotations[annotationKey]
			if !ok || existingValue != desiredValue {
				log.Info("Updating namespace annotation", "namespace", namespaceName, "annotation", annotationKey)
				namespace.Annotations[annotationKey] = desiredValue
				if err = r.Update(ctx, namespace); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *TenantReconciler) ensureRoleBinding(ctx context.Context, namespaceName string, groups []string, clusterRoleName string) error {
	log := log.FromContext(ctx)

	roleBindingName := fmt.Sprintf("%s-%s-rb", namespaceName, clusterRoleName)

	clusterRole := &rbacv1.ClusterRole{}
	err := r.Get(ctx, client.ObjectKey{Name: clusterRoleName}, clusterRole)
	if err != nil {
		log.Error(err, "Failed to get ClusterRole", "clusterRole", clusterRoleName)
	}

	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleBindingName,
			Namespace: namespaceName,
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "ClusterRole",
			Name:     clusterRoleName,
			APIGroup: rbacv1.GroupName,
		},
		Subjects: make([]rbacv1.Subject, len(groups)),
	}

	for i, group := range groups {
		roleBinding.Subjects[i] = rbacv1.Subject{
			Kind:     "Group",
			Name:     group,
			APIGroup: rbacv1.GroupName,
		}
	}

	err = r.Get(ctx, client.ObjectKey{Name: roleBindingName, Namespace: namespaceName}, roleBinding)
	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("Creating RoleBinding", "roleBinding", roleBindingName, "namespace", namespaceName)

			err = r.Create(ctx, roleBinding)
			if err != nil {
				log.Error(err, "Failed to create RoleBinding", "roleBinding", roleBindingName, "namespace", namespaceName)
				return err
			}
		} else {
			log.Error(err, "Failed to get RoleBinding", "roleBinding", roleBindingName, "namespace", namespaceName)
		}
	} else {
		// Compare current and desired roleBinding
		groupsChanged := false

		existingGroups := make(map[string]bool)
		newGroups := make(map[string]bool)

		for _, subject := range roleBinding.Subjects {
			if subject.Kind == "Group" {
				existingGroups[subject.Name] = true
			}
		}

		for _, group := range groups {
			newGroups[group] = true
			if _, exists := existingGroups[group]; !exists {
				groupsChanged = true
				break
			}
		}

		if len(existingGroups) != len(newGroups) {
			groupsChanged = true
		}

		if groupsChanged {
			log.Info("Updating RoleBinding", "roleBinding", roleBindingName, "namespace", namespaceName)

			roleBinding.Subjects = make([]rbacv1.Subject, len(groups))
			for i, group := range groups {
				roleBinding.Subjects[i] = rbacv1.Subject{
					Kind:     "Group",
					Name:     group,
					APIGroup: rbacv1.GroupName,
				}
			}

			err = r.Update(ctx, roleBinding)
			if err != nil {
				log.Error(err, "Failed to update RoleBinding", "roleBinding", roleBindingName, "namespace", namespaceName)
				return err
			}
		} else {
			log.Info("RoleBinding already exists", "roleBinding", roleBindingName, "namespace", namespaceName)
		}
	}

	return nil
}

func (r *TenantReconciler) deleteExternalResources(ctx context.Context, tenant *multitenancyv1.Tenant) error {
	// Delete any external resources created for this tenant
	log := log.FromContext(ctx)
	for _, ns := range tenant.Spec.Namespaces {
		// Delete Namespace
		namespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: ns,
			},
		}
		if err := r.Delete(ctx, namespace); client.IgnoreNotFound(err) != nil {
			log.Error(err, "unable to delete Namespace", "namespace", ns)
			return err
		}
		log.Info("Namespace deleted", "namespace", ns)
	}
	log.Info("All resources deleted for tenant", "tenant", tenant.Name)
	return nil
}
