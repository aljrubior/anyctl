# Start your deployment configuration

### 1. Show the current state of the deployment

```
anyctl runtimemanager deployments get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   0/1        APPLIED   private-space-regions-qa    4.4.0:20211227-2   getting-started-hello-mule:1.0.1
```

### 2. Start the deployment

```
anyctl runtimemanager deployments start hello-mule-app
```

Expected output:

```
Deployment 'hello-mule-app' started.
```

### 3. Show the current state of the deployment

```
anyctl runtimemanager deployments get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   1/1        APPLIED   private-space-regions-qa    4.4.0:20211227-2   getting-started-hello-mule:1.0.1
```

### 4. List the change history of the deployment

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

### 5. Show differences between current deployment version and the previous version

```
anyctl runtimemanager deployments history diff hello-mule-app 88f6eb 
```

Expected output:

```
   ~ createdAt: 1645694086372 --> 1645693429274
   ~ application: 
     ~ desiredState: STARTED --> STOPPED
   ~ version: 063ad966-2f35-4ef5-8c86-ac8730f08b96 --> 88f6eb3c-d0c3-45b3-8621-50bb6e06a508

Result: 0 to add, 3 to change, 0 to delete.
```
