# Update the deployment configuration

### 1. Save into a file the current deployment state

```
anyctl runtimemanager deployments describe hello-mule-app
```

### 2. Declare your deployment configurations in the file

#### 2.1 Add database configurations into the application properties

```
  db.host: mysql.localhost
  db.port: 3306
  db.user: root
  db.pass: qwerty
```

#### 2.2 Add Java Virtual Machine arguments

```
-XX:MaxRAMPercentage=66.0
```

#### 2.3 Increase Deployment resources

```
CPU Reserved     from 20m   to 50m
Memory Limit     from 700Mi to 750Mi
Memory Reserved  from 700Mi to 750Mi
```

#### 2.4 Increase the Deployment replicas

```
Replicas         from 2 to 1
```

#### 2.5 Change the asset version

```
Asset version    from 1.0.2 to 1.0.1
```


Expected File: /tmp/hello-mule-app.yaml

```
apiVersion: v1
kind: Deployment
metadata:
  name: hello-mule-app
spec:
  id: 50c765f0-06dc-4bc4-8f7b-0b506280996e
  name: hello-mule-app
  target:
    deploymentSettings:
      anypointMonitoringScope: cluster
      http:
        inbound:
          publicUrl: ""
      jvm:
        args: -XX:MaxRAMPercentage=66.0
      resources:
        cpu:
          limit: 3500m
          reserved: 50m
        memory:
          limit: 750Mi
          reserved: 750Mi
      runtimeVersion: 4.4.0:20211227-2
      sidecars:
        anypoint-monitoring:
          image: auto
          resources:
            cpu:
              limit: 50m
              reserved: 0m
            memory:
              limit: 50Mi
              reserved: 50Mi
      updateStrategy: rolling
    provider: MC
    replicas: 1
    targetId: a8949688-f7b6-4302-8ca5-c9f6dea5f9aa
  application:
    configuration:
      mule.agent.application.properties.service:
        applicationName: hello-mule-app
        properties:
          db.host: mysql.localhost
          db.port: 3306
          db.user: root
          db.pass: qwerty
    desiredState: STARTED
    ref:
      artifactId: getting-started-hello-mule
      groupId: a9ee3922-f172-452f-8eef-5db2e685eaa1
      packaging: jar
      version: 1.0.1
```

### 3. Show the deployment plan

```
anyctl runtimemanager deployments apply -f /tmp/hello-mule-app.yaml --plan
```

Expected output:

```
   ~ target: 
     ~ deploymentSettings: 
       ~ resources: 
         ~ cpu: 
           ~ reserved: 20m --> 50m
         ~ memory: 
           ~ reserved: 700Mi --> 750Mi
           ~ limit: 700Mi --> 750Mi
       ~ jvm: 
         + args: -XX:MaxRAMPercentage=66.0
     ~ replicas: 2 --> 1
   ~ application: 
     ~ ref: 
       ~ version: 1.0.2 --> 1.0.1
     ~ configuration: 
       ~ mule.agent.application.properties.service: 
         ~ properties: 
           + db.host: mysql.localhost
           + db.pass: qwerty
           + db.port: 3306
           + db.user: root

Result: 5 to add, 5 to change, 0 to delete.
```

### 4. Apply the deployment file

```
anyctl runtimemanager deployments apply -f /tmp/hello-mule-app.yaml 
```

Expected output:

```
Deployment 'hello-mule-app' updated.
```

### 5. Show the current deployment state

```
anyctl runtimemanager deployments get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
  hello-mule-app   1/1        APPLIED   runtime-mgr-private-space   4.4.0:20211227-2   getting-started-hello-mule:1.0.1
```

### 6. List the change history of the deployment

```
anyctl runtimemanager deployments history get hello-mule-app
```

Expected output:

```
 DATE                  CHANGES
 2022-02-24T05:02:08   5c394d (Last successful)
 2022-02-24T04:10:04   87eb15
 2022-02-24T03:41:15   88c839
 2022-02-24T03:40:05   a3d2af
```

### 7. Show differences between current deployment version and the previous version

```
anyctl runtimemanager deployments history diff hello-mule-app 87eb15 
```

Expected output:

```
   ~ version: 5c394d4d-d753-4511-adb3-f64ad0a0d0fd --> 87eb15b1-b44f-4e9f-a7a8-df24c95a67a5
   ~ createdAt: 1645689728888 --> 1645686604926
   ~ target: 
     ~ deploymentSettings: 
       ~ jvm: 
         - args: -XX:MaxRAMPercentage=66.0
       ~ resources: 
         ~ cpu: 
           ~ reserved: 50m --> 20m
         ~ memory: 
           ~ limit: 750Mi --> 700Mi
           ~ reserved: 750Mi --> 700Mi
     ~ replicas: 1 --> 2
   ~ application: 
     ~ configuration: 
       ~ mule.agent.application.properties.service: 
         ~ properties: 
           - db.host: mysql.localhost
           - db.pass: qwerty
           - db.port: "3306"
           - db.user: root
     ~ ref: 
       ~ version: 1.0.1 --> 1.0.2

Result: 0 to add, 7 to change, 5 to delete.
```
