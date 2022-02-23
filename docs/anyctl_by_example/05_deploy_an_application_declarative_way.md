# Deploy an Application (Declarative)

### 1. List the targets available to deploy applications in your current environment

```
anyctl runtimemanager targets get
```

Expected output:

```
 ID                                     NAME                           TYPE
 821f9ad2-4730-4916-9e98-bb837b2f4d42   private-space-products-qa      MC
 7be2ca41-ac41-4b81-9c37-883f2b8152ea   private-space-regions-qa       MC
 8751d9cf-4f38-40d7-9bb5-f4c93ebc89a2   private-space-services-qa      MC
 04992a72-ad0f-4c28-80cc-9abf862c7ef6   private-space-finance-qa       MC
```

### 2. Declare your deployment in a yaml file

File: /tmp/hello-mule-app.yaml

```
apiVersion: v1
kind: Deployment
metadata:
  name: hello-mule-app
spec:
  name: hello-mule-app
  labels:
  - beta
  target:
    provider: MC
    targetId: 7be2ca41-ac41-4b81-9c37-883f2b8152ea
    deploymentSettings:
      resources:
        cpu:
          reserved: 20m
          limit: 3500m
        memory:
          reserved: 700Mi
          limit: 700Mi
      clustered: false
      enforceDeployingReplicasAcrossNodes: false
      http:
        inbound:
          publicUrl: null
      jvm: {}
      runtimeVersion: 4.4.0:20210226-1
      lastMileSecurity: false
      forwardSslSession: false
      updateStrategy: rolling
    replicas: 1
  application:
    ref:
      groupId: 21cfa350-cb56-4895-ac8c-825f4eabc928
      artifactId: getting-started-hello-mule
      version: 1.0.0
      packaging: jar
    assets: []
    desiredState: STARTED
    configuration:
      mule.agent.application.properties.service:
        applicationName: hello-mule-app
        properties: {}
        secureproperties: {}
```

### 3. Show the deployment plan

```
anyctl runtimemanager deployments apply -f /tmp/hello-mule-app.yaml --plan
```

Expected output

```
Deployment 'hello-mule-app' will be created.
```

### 4. Apply the deployment file

```
anyctl runtimemanager deployments apply -f /tmp/hello-mule-app.yaml 
```

Expected output:

```
Deployment 'hello-mule-app' created.
```

### 5. List the application deployed in the previous step

```
anyctl runtimemanager deployments get hello-mule-app
```

Expected output: 

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   1/1        APPLIED   private-space-regions-qa    4.4.0:20220124-4   getting-started-hello-mule:1.0.0
```

### 5. Inspect in detail your deployment

```
anyctl runtimemanager deployments describe hello-mule-app 
```

Expected output:

```
apiVersion: v1
kind: Deployment
metadata:
  name: hello-mule-app
spec:
  id: 3a130eef-602d-49f7-a6c0-4edfbe77f342
  name: hello-mule-app  
  desiredVersion: c371f6d1-f970-4249-bbea-28d4a7cba685
  lastSuccessfulVersion: c371f6d1-f970-4249-bbea-28d4a7cba685
  target:
    deploymentSettings:
      sidecars:
        anypoint-monitoring:
          Image: auto
          Resources:
            cpu:
              limit: 50m
              reserved: 0m
            memory:
              limit: 50Mi
              reserved: 50Mi    
      anypointMonitoringScope: app
      http:
        inbound:
          decoratedProperties:
            defaultEndpoint: ""
      jvm: {}
      resources:
        cpu:
          limit: 3500m
          reserved: 20m
        memory:
          limit: 700Mi
          reserved: 700Mi
      runtimeVersion: 4.4.0:20220124-4
      updateStrategy: rolling
    provider: MC
    replicas: 1
    targetId: 345c1246-4d7a-4b4f-b2c1-fbef442d2e2d
    type: ""
  application:
    configuration:
      mule.agent.application.properties.service:
        applicationName: hello-mule-app
    desiredState: STARTED
    ref:
      artifactId: getting-started-hello-mule
      groupId: 66310c16-bce5-43c4-b978-5945ed2f99c5
      packaging: jar
      version: 1.0.0
    status: RUNNING
  replicas:
  - currentDeploymentVersion: c371f6d1-f970-4249-bbea-28d4a7cba685
    deploymentLocation: 345c1246-4d7a-4b4f-b2c1-fbef442d2e2d
    id: hello-mule-app-867655556-7d5rt
    reason: ""
    state: STARTED
  status: APPLIED
```