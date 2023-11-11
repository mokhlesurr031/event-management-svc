# How to Run


### Using Docker...

Execute
```'sh
docker-compose up
```

The application shoul start running on **8081** port. 



### Using Kubernetes Local Cluster...

Run the application service (MySQL)

```sh
helmfile -f helmfile-services.yaml sync
```

Run the application

```sh
helmfile sync
```