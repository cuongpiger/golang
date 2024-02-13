```bash=
kubebuilder create api --group example.com --version v1alpha1 --kind Memcached --image=memcached:1.4.36-alpine --image-container-command="memcached,-m=64,-o,modern,-v" --image-container-port="11211" --run-as-user="1001" --plugins="deploy-image/v1-alpha" --make=false
```