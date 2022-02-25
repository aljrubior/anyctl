# Update the asset in your deployment

### 1. Publish two more versions of the asset

```
anyctl assets upload --file /tmp/getting-started-hello-mule.jar --name getting-started-hello-mule --version 1.0.1
anyctl assets upload --file /tmp/getting-started-hello-mule.jar --name getting-started-hello-mule --version 1.0.2
```

Expected output:

```
Asset 'getting-started-hello-mule' v1.0.1 created.
Asset 'getting-started-hello-mule' v1.0.2 created.
```

### 2. Show the latest version of the asset

```
anyctl assets get getting-started-hello-mule
```

Expected output:

```
 NAME                         VERSION   STATUS      RUNTIME   UPDATED AT                 CREATED BY
 getting-started-hello-mule   1.0.2     published   4.3.0     2022-02-24T07:04:11.383Z   cloudhub-rest
```

### 3. Show the current deployment state

```
anyctl runtimemanager deployments get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   2/2        APPLIED   private-space-regions-qa    4.4.0:20211227-2   getting-started-hello-mule:1.0.0
```

### 4. Update the deployment to the latest version of the same asset 

```
anyctl runtimemanager deployments set hello-mule-app --asset getting-started-hello-mule 
```

Expected output:

```
Deployment 'hello-mule-app' updated with asset 'getting-started-hello-mule:1.0.2'.
```

### 4. Show the current deployment state

```
anyctl runtimemanager deployments get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   2/2        APPLIED   private-space-regions-qa    4.4.0:20211227-2   getting-started-hello-mule:1.0.2
```

### 5. List the change history of a deployment

```
anyctl runtimemanager deployments history get hello-mule-app
```

Expected output:

```
 DATE                  CHANGES
 2022-02-24T04:10:04   87eb15 (Last successful)
 2022-02-24T03:41:15   88c839
 2022-02-24T03:40:05   a3d2af
```

### 6. Show differences between current deployment version and the previous version

```
anyctl runtimemanager deployments history diff hello-mule-app 88c839 
```

Expected output:

```
anyctl runtimemanager deployments history diff hello-mule-app 88c839                                                                                                                                                               develop  ✭ ✱
   ~ createdAt: 1645686604926 --> 1645684875657
   ~ application: 
     ~ ref: 
       ~ version: 1.0.2 --> 1.0.0
   ~ version: 87eb15b1-b44f-4e9f-a7a8-df24c95a67a5 --> 88c8397b-137f-4860-ac44-8046ac4809ff

Result: 0 to add, 3 to change, 0 to delete.
```