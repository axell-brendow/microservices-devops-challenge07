# microservices-devops-challenge07

Course "[Architecting and Developing modern and scalable applications based on Microservices (pt-br)](https://drive.google.com/file/d/1JXXmYhfi-Sk0zwiEbBIEswDC6AFeyBlD/view?usp=sharing)".

- Mount CI/CD pipeline for a Go application using Docker, Google Cloud Build and Kubernetes

To create this Deployment in your cluster:

```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

To test Go application:

```
GOPATH=$(pwd) go test challenge07 -v -cover
```

To run Go application:

```
# You can choose any port you want if you will not use k8s/deployment.yaml
# to run the application in a Kubernetes Cluster
GOPATH=$(pwd) go run challenge07 --port=9000
```

## Other challenges

- [Nginx + PHP-FPM + Redis + MySQL with docker-compose](https://github.com/axell-brendow/microservices-devops-challenge01)
- [Create Hello World in Go with a Docker image less than 2MB](https://github.com/axell-brendow/microservices-devops-challenge02)
- [Apply CI (Continuous Integration) to test and create a docker image of a simple Go application](https://github.com/axell-brendow/microservices-devops-challenge04)
- [3 Nginx replicas with a Load Balancer; MySQL with Secret Object and PersistentVolume; Http server in Go with CI and tests in GCP](https://github.com/axell-brendow/microservices-devops-challenge05)
- [Apply a Horizontal Pod Autoscaler in a Go application under high usage](https://github.com/axell-brendow/microservices-devops-challenge06)
- [Mount CI/CD pipeline for a Go application using Docker, Google Cloud Build and Kubernetes](https://github.com/axell-brendow/microservices-devops-challenge07)
