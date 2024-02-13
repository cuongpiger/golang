# Setup
- Init the project
  ```bash=
  go mod init github.com/cuongpiger/golang
  ```

- Create the project
  ```bash=
  kubebuilder init --domain cuongpiger.io github.com/cuongpiger/golang
  ```
  
- Create the API
  ```bash=
  kubebuilder create api --group webapp --version v1 --kind Guestbook
  >> INFO Create Resource [y/n]                        
  y
  >> INFO Create Controller [y/n]                      
  y
  ```
  
- Update the CRD
  ```bash=
  make manifests
  ```

- Install the CRDs into the cluster
  ```bash=
  # export KUBECONFIG=<path_to_kubeconfig_file>
  
  make install
  ```
  
- Run the controller _(this will run in the foreground, so switch to a new terminal if you want to leave it running)_
  ```bash=
  make run
  ```