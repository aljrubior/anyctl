# Scale your deployment

### 1. List the change history of a deployment 

```
anyctl runtimemanager deployments history get hello-mule-app
```

Expected output:

```
 DATE                  CHANGES
 2022-02-24T03:40:05   a3d2af (Last successful)
```

### 2. Show the current state of a deployment

```
anyctl runtimemanager deployments history get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   1/1        APPLIED   runtime-mgr-private-space   4.4.0:20220124-4   getting-started-hello-mule:1.0.0
```

### 3. Scale the deployment

```
anyctl runtimemanager deployments scale hello-mule-app --replicas 2
```

Expected output:

```
Deployment 'hello-mule-app' scaled to '2' replicas.
```

### 4. Show the current status of the deployment

```
anyctl runtimemanager deployments history get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   2/2        APPLIED   runtime-mgr-private-space   4.4.0:20211227-2   getting-started-hello-mule:1.0.0
```


### 5. List the change history of a deployment

```
anyctl runtimemanager deployments history get hello-mule-app
```

Expected output:

```
 DATE                  CHANGES
 2022-02-24T03:41:15   88c839 (Last successful)
 2022-02-24T03:40:05   a3d2af
```

### 5. Show the current status of the deployment

```
anyctl runtimemanager deployments history get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   2/2        APPLIED   runtime-mgr-private-space   4.4.0:20211227-2   getting-started-hello-mule:1.0.0
```

### 6. Show differences between current deployment version and the previous version

```
anyctl runtimemanager deployments history diff hello-mule-app a3d2af 
```

Expected output:

```
   ~ target: 
     ~ replicas: 2 --> 1
   ~ version: 88c8397b-137f-4860-ac44-8046ac4809ff --> a3d2af14-fd42-48f7-bb12-bd1050a06c56
   ~ createdAt: 1645684875657 --> 1645684805014

Result: 0 to add, 3 to change, 0 to delete.
```