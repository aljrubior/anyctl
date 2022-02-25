# Stop your deployment configuration

### 1. Show the current state of the deployment

```
anyctl runtimemanager deployments get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   1/1        APPLIED   runtime-mgr-private-space   4.4.0:20211227-2   getting-started-hello-mule:1.0.1
```

### 2. Stop the deployment

```
anyctl runtimemanager deployments stop hello-mule-app
```

Expected output:

```
Deployment 'hello-mule-app' stopped.
```

### 3. Show the current state of the deployment

```
anyctl runtimemanager deployments get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   0/1        APPLIED   runtime-mgr-private-space   4.4.0:20211227-2   getting-started-hello-mule:1.0.1
```

### 4. List the change history of the deployment

```
anyctl runtimemanager deployments history get hello-mule-app
```

Expected output:

```
 DATE                  CHANGES
 2022-02-24T06:03:49   88f6eb (Last successful) 
 2022-02-24T05:02:08   5c394d
 2022-02-24T04:10:04   87eb15
 2022-02-24T03:41:15   88c839
 2022-02-24T03:40:05   a3d2af
```

### 5. Show differences between current deployment version and the previous version

```
anyctl runtimemanager deployments history diff hello-mule-app 5c394d 
```

Expected output:

```
   ~ createdAt: 1645693429274 --> 1645693227013
   ~ application: 
     ~ desiredState: STOPPED --> STARTED
   ~ version: 88f6eb3c-d0c3-45b3-8621-50bb6e06a508 --> de15a1bb-52bc-4ec6-a046-6094bb106b03

Result: 0 to add, 3 to change, 0 to delete.
```
