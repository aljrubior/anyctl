## Compare your deployment versions

### 1. List the change history of the deployment

```
anyctl runtimemanager deployments history get hello-mule-app
```

Expected output:

```
 DATE                  CHANGES
 2022-02-24T06:14:46   063ad9 (Last successful) 
 2022-02-24T06:03:49   88f6eb 
 2022-02-24T05:02:08   5c394d
 2022-02-24T04:10:04   87eb15
 2022-02-24T03:41:15   88c839
 2022-02-24T03:40:05   a3d2af
```

### 2. Compare the current version with the first deployment

```
anyctl runtimemanager deployments history diff hello-mule-app 
```

If you follow all the previous examples you will a result like this:

```
   ~ createdAt: 1645694086372 --> 1645684805014
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
     ~ replicas: 2 --> 1
   ~ application: 
     ~ configuration: 
       ~ mule.agent.application.properties.service: 
         ~ properties: 
           - db.host: mysql.localhost
           - db.pass: qwerty
           - db.port: 3306
           - db.user: root
     ~ ref: 
       ~ version: 1.0.1 --> 1.0.0
   ~ version: 063ad966-2f35-4ef5-8c86-ac8730f08b96 --> a3d2af14-fd42-48f7-bb12-bd1050a06c56

Result: 0 to add, 7 to change, 5 to delete.
```