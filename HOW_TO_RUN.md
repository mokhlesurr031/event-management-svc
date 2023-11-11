# How to Run


### Using Docker...

Prerequisite:

1. Need to have docker installed. 

#### Execute
```'sh
docker-compose up
```

The application shoul start running on **8081** port. 



### Using Kubernetes Local Cluster...


Prerequisite:

1. Need to have docker installed. 
2. Need to have kubernetes installed. 
3. Need to have helm installed. 
4. Need to have helmfile installed.


Execute

```sh
cd deployment
```

#### Run the application service (MySQL)

```sh
helmfile -f helmfile-services.yaml sync
```

Run the application

```sh
helmfile sync
```