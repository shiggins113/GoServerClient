# GoServerClient

Sample Go Application that runs a client and server application, writes decoded certificate data to a MySQL database.

To install on a K8s cluster, run kubectl create -f server/manifests/

Once running, you can view the UI using:

kubectl port-forward server-<xxx> 8080:8080
kubectl port-forward server-<xxx> 8081:8081

Opening http://localhost:8080/cert will display the upload page.

Docker images:

https://hub.docker.com/repository/docker/higgins113/my-go-client
https://hub.docker.com/repository/docker/higgins113/my-go-server/general
