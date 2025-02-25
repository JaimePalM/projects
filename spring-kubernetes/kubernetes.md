# KUBERNETES with Spring Boot

## Get information
```bash
kubectl get all
```

## Create deployment
```bash
kubectl create deployment my-deployment --image <DcokerHub image> --dry-run client -o=yaml > deployment.yaml
```
This will create a deployment.yaml file with the deployment configuration. There you can change the configuration as needed. 

## Create service
```bash
kubectl create service clusterip my-service --tcp=80:80 --dry-run=client -o=yaml > service.yaml
```

## Apply deployment and service
```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```
Now if you run `kubectl get all` you should see the deployment and service running.

## Delete deployment and service
```bash 
kubectl delete -f deployment.yaml
kubectl delete -f service.yaml
```

## Expose service
To expose the service to the internet you can edit the service.yaml file and change the type from ClusterIP to LoadBalancer or NodePort. Then run `kubectl apply -f service.yaml` again.
Now you can access the service using the external IP of the service.

## Setting environment variables
One very common use cases to use environment variables is to override log levels. 

For example, if you have a Spring Boot application you can set the log level to DEBUG by setting the environment variable `LOG_LEVEL` to `DEBUG`. 

In the deployment.yaml file you can add the following configuration to set the environment variable:
```yaml
containers:
- image: springframeworkguru/kbe-rest-brewery
  name: kbe-rest-brewery
  resources: {}
  env: 
  - name: LOGGING_LEVEL_GURU_SPRINGFRAMEWORK_SFGRESTBREWERY
    value: info
```
The name of the environment variable comes from the `application.properties` file. So it's the path to the property in the application.properties file.

To apply the changes run `kubectl apply -f deployment.yaml`. We will see another pod created with the new configuration. If we do `kubectl logs <pod-name>` we will see there is no DEBUG logs anymore.

## Readiness and Liveness probes
Readiness and Liveness probes are used to check if the application is ready to receive traffic and if the application is still alive. Spring Boot Actuator provides endpoints to check if the application is ready and alive.
To add the probes to the deployment.yaml file you can add the following configuration:
```yaml
    - name: MANAGEMENT_ENDPOINT_HEALTH_PROBES_ENABLED
        value: "true"
    - name: MANAGEMENT_HEALTH_READINESSTATE_ENABLED
        value: "true"
    readinessProbe:
        httpGet:
        port: 8080
        path: /actuator/health/readiness
```
The `readinessProbe` will check if the application is ready to receive traffic, sending a GET request to the `/actuator/health/readiness` endpoint. If the application is not ready, the pod will be restarted.
Note: Kubernetes expects the values as strings, so you need to put the values in quotes.

Now if we go to Postman and sent the following URL:
```
http://localhost:<NodePort>/actuator/health/readiness
```
We should see the response `{"status":"UP"}`.

Now we can add the liveness probe to the deployment.yaml file:
```yaml
    - name: MANAGEMENT_HEALTH_LIVENESSTATE_ENABLED
        value: "true"
    livenessProbe:
        httpGet:
        port: 8080
        path: /actuator/health/liveness
```
Restart the pod and check in Postman if the liveness probe is working:
```
http://localhost:<NodePort>/actuator/health/liveness
```
We should see the response `{"status":"UP"}`.

Now Kubernetes will check periodically if the application is ready to receive traffic and if the application is still alive. If the application is not ready or alive, the pod will be restarted.

## Graceful Shutdown
When a pod is terminated, Kubernetes sends a SIGTERM signal to the pod. The pod has 30 seconds to shutdown gracefully. If the pod doesn't shutdown in 30 seconds, Kubernetes will send a SIGKILL signal to the pod, which will kill the pod immediately.

We have to set and environment variable in the deployment.yaml file to enable the graceful shutdown in Spring Boot:
```yaml
    - name: SERVER_SHUTDOWN
        value: "graceful"
```

Now set the property for Kubernetes:
```yaml
    lifecycle:
      preStop:
        exec:
          command: ["sh", "-c", "sleep 10"]
```
This will make the pod wait 10 seconds before shutting down. This is useful if you have a database connection that needs to be closed before the pod is terminated.

There is no an easy way to test the graceful shutdown. You can try to simulate a shutdown by running `kubectl delete pod <pod-name>`. The pod should wait 10 seconds before shutting down.

# MICROSERVICES with Kubernetes
When you have multiple microservices running in Kubernetes, you can use Kubernetes Services to communicate between the microservices.

Guru Brewery example:
- 4 microservices: beer, inventory-failover, inventory-service, and order-service.
  - The order-service communicates with the inventory-service to check if there is enough inventory to place an order.
  - The inventory-service communicates with the inventory-failover to check if the inventory-service is down.
  - The inventory-failover is a dummy service that always returns true.
  - The beer service is a dummy service that returns a list of beers.
- Each service is a Spring Boot application (like the brewery example).
- Brewery gateway: a Spring Boot application that acts as a gateway to the microservices.
- Each microservice has its own deployment and service in Kubernetes so each one will be a container in a pod.

## Docker Compose
Define and run multi-container Docker applications. With Docker Compose you can define the services, networks, and volumes in a YAML file. Then you can run the application with a single command.

The configuration file is called `compose-local.yml`. Here there is some services defined:
- mysql: a MySQL database, give persistence to the application (user and password)
- elasticsearch, kibana and filebeat: for loggin (ports to redirect the logs and volumes to store the logs)
- jms: a JMS (Java Message Service) server (ports to redirect the messages)
- And the services that are part of the application (defined before).

To run the application you can run `docker-compose -f compose-local.yml up`. This will start all the services defined in the configuration file. We can check the services running with `docker ps`.

Now if we go to Postman and send a GET request to `http://localhost:9090/api/v1/beer` we should see the response with the list of beers. Because we are using 9090 port, we are using the gateway service. If we send the request to `http://localhost:8080/api/v1/beer` we should see the same response, but we are using the beer service directly.

Finally, we can visit the kiabana dashboard to see the logs of the application. To do this we have to go to `http://localhost:5601` and create an index pattern to see the logs (the option may be in a footer). Not working for me.

For shutting down the application you can run `docker compose -f compose-local.yml down`.

## Infrastructure Services
Infrastructure services are services that are needed for the application to run. For example, a database, a message broker, a logging system, etc. These services are not part of the application, but the application needs them to run.

### MySQL
If we take a look to the Docker Compose file we have MySQL defined with two environment variables: `MYSQL_ROOT_PASSWORD` and `MYSQL_DATABASE`. The first one is the password for the root user and the second one is the name of the database. 
Now, we have to replicate this in a Kubernetes context. We creaate a deployment.yaml file for MySQL:
```bash
kubectl create deployment mysql --image=mysql:5.7 --dry-run=client -o=yaml > mysql-deployment.yml
```
In this file we have to add the environment variables of the Docker Compose file:
```yaml
  env:
  - name: MYSQL_ROOT_PASSWORD
    value: dbpassword
  - name: MYSQL_DATABASE
    value: beerservice
```
Now we can apply the deployment:
```bash
kubectl apply -f mysql-deployment.yml
```
Finally, we have to create a service for MySQL. We can create a service.yaml file:
```bash
kubectl create service clusterip mysql --tcp=3306:3306 --dry-run=client -o=yaml > mysql-service.yml
```
We use tcp port to expose the MySQL service. Now we can apply the service:
```bash
kubectl apply -f mysql-service.yml
```
If we check the services running with `kubectl get all` we should see the MySQL service running.
```
NAME                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
service/kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP    24h
service/mysql        ClusterIP   10.101.45.110   <none>        3306/TCP   4s

NAME                    READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/mysql   1/1     1            1           99s
````
Note: we can declare the deployment and service in the same file. For that we have to add the service configuration to the deployment.yaml file.

### JMS (Java Message Service)
JMS is a messaging standard that allows Java EE applications to create, send, receive, and read messages. It is a messaging system that allows communication between different components of a distributed application.

The JMS server is defined in the Docker Compose file with only the port mapping, so it's quite simple to replicate in Kubernetes. We can create a deployment.yaml file for the JMS server:
```bash
kubectl create deployment jms --image=vromero/activemq-artemis --dry-run=client -o=yaml > jms-deployment.yml
kubectl apply -f jms-deployment.yml
```
Now we can create a service for the JMS server. Because of the JMS only has two port mappings, we can use the `--tcp` flag to expose the ports and then replicate in the service.yaml file:
```bash
kubectl create service clusterip jms --tcp=8161:8161 --tcp=61616:61616 --dry-run=client -o=yaml > jms-service.yml
kubectl apply -f jms-service.yml
```
Now we can check the services running with `kubectl get all` and we should see the JMS service running.
```
NAME                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)              AGE
service/jms          ClusterIP   10.97.118.113   <none>        8161/TCP,61616/TCP   8s
service/kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP              24h
service/mysql        ClusterIP   10.101.45.110   <none>        3306/TCP             3m44s

NAME                    READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/jms     1/1     1            1           28s
deployment.apps/mysql   1/1     1            1           5m19s
```

## Spring Boot Microservices
Recap of the microservices: beer, inventory-failover, inventory-service, and order-service. Each microservice has its own deployment and service in Kubernetes. Microservices main goal is to deploy some functionality in a separate service. This way we can scale the services independently, and we can deploy new features without affecting the other services.

### Services
We start creating the deployment.yaml file for each microservice. And for those which have environment variables we have to add them to the deployment.yaml file manually.
```bash
kubectl create deployment inventory-service --image=springframeworkguru/kbe-brewery-inventory-service --dry-run=client -o=yaml > inventory-deployment.yml
kubectl create deployment inventory-failover --image=springframeworkguru/kbe-brewery-inventory-failover --dry-run=client -o=yaml > inventory-failover-deployment.yml
kubectl create deployment beer-service --image=springframeworkguru/kbe-brewery-beer-service --dry-run=client -o=yaml > beer-deployment.yml
kubectl create deployment order-service --image=springframeworkguru/kbe-brewery-order-service --dry-run=client -o=yaml > order-deployment.yml
```

Now we create the services for each microservice, keeping an eye on the ports that the services are using.
```bash
kubectl create service clusterip inventory-service --tcp=8082:8082 --dry-run=client -o=yaml > inventory-service.yml
kubectl create service clusterip inventory-failover --tcp=8083:8083 --dry-run=client -o=yaml > inventory-failover-service.yml
kubectl create service clusterip beer-service --tcp=8080:8080 --dry-run=client -o=yaml > beer-service.yml
kubectl create service clusterip order-service --tcp=8081:8081 --dry-run=client -o=yaml > order-service.yml
```

At this point, with every environment variable and port mapping set, we can apply the deployment and service files.
```bash
kubectl apply -f inventory-deployment.yml
kubectl apply -f inventory-failover-deployment.yml
kubectl apply -f beer-deployment.yml
kubectl apply -f order-deployment.yml
kubectl apply -f inventory-service.yml
kubectl apply -f inventory-failover-service.yml
kubectl apply -f beer-service.yml
kubectl apply -f order-service.yml
```

Now we can check the services running with `kubectl get all` and we should see all the microservices up and running.

### Readiness and Liveness Probes
We can add the readiness and liveness probes to the deployment.yaml file of each microservice. As we did before in the service exercise, now we define the same variables but in all the microservices (except the inventory-failover service).
```yaml
    - name: MANAGEMENT_ENDPOINT_HEALTH_PROBES_ENABLED
      value: "true"
    - name: MANAGEMENT_HEALTH_READINESSTATE_ENABLED
      value: "true"
    - name: MANAGEMENT_HEALTH_LIVENESSSTATE_ENABLED
      value: "true"
    readinessProbe:
        httpGet:
        port: 8080
    livenessProbe:
        httpGet:
        port: 8080
```
Now if we apply the changes to the deployment files and check the services running with `kubectl get all` we could see services pods duplicated, because they are restarting to apply the changes (in a few seconds they should be running again).

### Graceful Shutdown
We can add the graceful shutdown to the deployment.yaml file of each microservice. We have to add the environment variable `SERVER_SHUTDOWN` with the value `graceful` and the lifecycle configuration to wait 10 seconds before shutting down.
```yaml
    - name: SERVER_SHUTDOWN
      value: "graceful"
    lifecycle:
      preStop:
        exec:
          command: ["sh", "-c", "sleep 10"]
```

And apply the changes to the deployment files with `kubectl apply -f <deployment-file>`.
Now if we delete a pod with `kubectl delete pod <pod-name>` we should see the pod waiting 10 seconds before shutting down.

### Ingess Controllers
Ingress controllers are used to expose the services to the internet. In Kubernetes, the services are only accessible inside the cluster. To expose the services to the internet we can use an Ingress controller.

![Ingress Diagram](https://kubernetes.io/docs/images/ingress.svg)

To make the Ingress Controller we are going to use Spring Cloud Gateway. We will be exposing it on a port on our NodePort. Sometimes you will be standing up services and paths behind the gateway and the actual ingress going to be handled by the DevOps team. 

### Spring Cloud Gateway
Spring Cloud Gateway is a Spring Boot application that acts as a gateway to the microservices. It is a reverse proxy that routes the requests to the correct microservice. It is a good practice to have a gateway in front of the microservices to handle the requests.

If we take a look to the `SfgBreweryGatewayApplication.java` file we can see the configuration of the gateway
```java
    @Bean
    public RouteLocator gatewayRoutes(RouteLocatorBuilder builder) {

        return builder.routes()
                .route("beer-service", r -> r.path("/api/v1/beer/*", "/api/v1/beerUpc/*")
                        .uri("http://beer-service:8080"))
                .route("inventory-service", r -> r.path("/api/v1/beer/*/inventory")
                        .uri("http://inventory-service:8082"))
                .route("order-service", r -> r.path(("/api/v1/customers/**"))
                        .uri("http://order-service:8081"))
                .build();
    }
```
This configuration is routing the requests to the correct microservice. For example, if the request is `/api/v1/beer` it will be routed to the beer-service microservice. If the request is `/api/v1/beer/1/inventory` it will be routed to the inventory-service microservice, so on.

Now we can see in the Docker Compose file the definition of the gateway service, it exposes the 9090 port. We can replicate this in Kubernetes creating a deployment.yaml and service.yaml files for the gateway service.
```bash
kubectl create deployment gateway --image=springframeworkguru/kbe-brewery-gateway --dry-run=client -o=yaml > gateway-deployment.yml
kubectl create service nodeport gateway --tcp=9090:9090 --dry-run=client -o=yaml > gateway-service.yml
```
Now we can apply both files with `kubectl apply -f <file>` and check the services running with `kubectl get all`. We should see the gateway service running. Something like this:
```
service/gateway              NodePort    10.109.204.165   <none>        9090:31248/TCP       2m47s
```
Because of we use NodePort, we can access the gateway service using the external IP of the service and it's listening on port 31248. If we go to Postman and send a GET request to `http://localhost:31248/api/v1/beer/` we should see the response with the list of beers.

What is happening?
1. The request is sent to the gateway service (NodePort 31248).
2. The gateway service routes the request to the correct microservice (ClusterIP).
3. The microservice processes the request and sends the response to the gateway service.
4. The gateway service sends the response to the client.
5. The client receives the response.

### Deleting Services and Deployments
There is a couple different ways to delete services and deployments. You can delete them one by one or you can delete them all at once.
One by one:
```bash
kubectl delete service <service-name>
kubectl delete deployment <deployment-name>
```
All at once:
```bash
kubectl delete all --all
kubectl delete deployment --all
kubectl delete service --all
```

## Logging with ELK Stack
ELK Stack is a combination of three open-source tools: Elasticsearch, Logstash, and Kibana. It is used to collect, store, and visualize logs. We can see the services in the Docker Compose file. Now as we did before, we can replicate this in Kubernetes.

### Elasticsearch
It is a distributed search and analytics engine designed for horizontal scalability, reliability, and real-time search capabilities. It is commonly used for log and event data analysis, full-text search, and data visualization.

We can create a deployment.yaml file for Elasticsearch:
```bash
kubectl create deployment elasticsearch --image=docker.elastic.co/elasticsearch/elasticsearch:7.12.1 --dry-run=client -o=yaml > elasticsearch-deployment.yml
```
And add the environment variables to the deployment.yml file:
```yaml
  env:
  - name: discovery.type
    value: single-node
```
Now we can apply the deployment file:
```bash
kubectl apply -f elasticsearch-deployment.yml
```
Finally, we can create a service for Elasticsearch. Attention to the port mapping, Elasticsearch uses the 9200 port.
```bash
kubectl create service clusterip elasticsearch --tcp=9200:9200 --dry-run=client -o=yaml > elasticsearch-service.yml
kubectl apply -f elasticsearch-service.yml
```

### Kibana
It is an open-source data visualization dashboard for Elasticsearch. It provides visualization capabilities on top of the content indexed on an Elasticsearch cluster.

We can create a deployment.yaml file for Kibana:
```bash
kubectl create deployment kibana --image=docker.elastic.co/kibana/kibana:7.12.1 --dry-run=client -o=yaml > kibana-deployment.yml
```
There is nothing to add to the deployment file, so we can create the service file and apply both files. This time we are going to use a NodePort service to expose the Kibana service to the internet and then allow us to access the Kibana dashboard.
```bash
kubectl create service nodeport kibana --tcp=5601:5601 --dry-run=client -o=yaml > kibana-service.yml
kubectl apply -f kibana-deployment.yml
kubectl apply -f kibana-service.yml
```

### Filebeat
It is a lightweight shipper for forwarding and centralizing log data. Installed as an agent on the servers, Filebeat monitors the log files or locations that you specify, collects log events, and forwards them either to Elasticsearch or Logstash for indexing. All the information is in the Filebeat manifest file (`filebeat-kubernetes.yaml`).

If we take a look to the Docker Compose file we can see the configuration of Filebeat:
```yaml
  filebeat:
    image: docker.elastic.co/beats/filebeat:7.12.1
    volumes:
      - ./filebeat/filebeat.yml:/usr/share/filebeat/filebeat.yml
      - /var/lib/docker/containers:/var/lib/docker/containers:ro
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - elk
```
The volumes are used to tell Filebeat where to find the configuration file and the log files. The first volume is the configuration file, the second volume is the log files, and the third volume is the Docker socket (to get the logs from the Docker containers).

The default Filebeat configuration is with 'false' in some variables. So we override them in the deployment files with annotations.
```yaml
      annotations:
        co.elastic.logs/enabled: "true"
        co.elastic.logs.json-logging/json.keys_under_root: "true"
        co.elastic.logs.json-logging/json.add_error_key: "true"
        co.elastic.logs.json-logging/json.message_key: "message"
```
Here we are enabling the logs, setting the keys under root, adding an error key, and setting the message key. I don't know what this means, but it's what the tutorial says.

### Final Steps
We can get the Kibana port looking with `kubectl get services` and then access the Kibana dashboard with `http://localhost:<NodePort>`. We can create an index pattern to see the logs. Not working for me (no data found).

# Troubleshooting

## Pods are pending
If the pods are pending, it means that the pod is not running. 

- Check the Kubernetes nodes with `kubectl get nodes`. If the nodes are not ready, the pods will be pending.
- Get more information about the pod with `kubectl describe pod <pod-name>`. 
