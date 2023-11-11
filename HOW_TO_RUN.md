# How to Run


### Using Docker...

Prerequisite:

1. Need to have docker installed. 

#### Execute
```'sh
docker-compose up
```

The application shoul start running on local port **8081**.

### Dummy db table data_
The db is port-forwarded to 3306 on local machine. Execute [This Dummy Datas](zzz-dummy-data) to the respected tables. 




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

Then you need to port-forward the application service with local machine so that you can access the application.


### Dummy db table data_

Port-forward the database with local machine. Then execute [This Dummy Datas](zzz-dummy-data) to the respected tables. 
