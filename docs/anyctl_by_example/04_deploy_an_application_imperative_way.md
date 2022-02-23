# Deploy an Application (Imperative)

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

### 2. Deploy the Mule Application getting-started-hello-mule

```
anyctl runtimemanager deployments run hello-mule-app \
   --asset getting-started-hello-mule \
   --target-name private-space-finance-qa
```

Expected output:

```
Deployment 'hello-mule-app' created.
```

### 3. List the application deployed in the previous step

```
anyctl runtimemanager deployments get hello-mule-app
```

Expected output:

```
 NAME             REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-mule-app   1/1        APPLIED   private-space-finance-qa    4.5.0:20220125-1   getting-started-hello-mule:1.0.0
```

### 4. Inspect in detail your deployment

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
  id: f66691e0-35cf-48e5-8857-6c39265087f0
  name: hello-mule-app
  desiredVersion: be855c64-c3db-4043-be87-77452fd98ad1
  lastSuccessfulVersion: be855c64-c3db-4043-be87-77452fd98ad1
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
          limit: 16000m
          reserved: 20m
        memory:
          limit: 700Mi
          reserved: 700Mi
      runtimeVersion: 4.5.0:20220125-1
      updateStrategy: rolling
    provider: MC
    replicas: 1
    targetId: 345c1246-4d7a-4b4f-b2c1-fbef442d2e2d
    type: ""  
  application:
    ref:
      artifactId: getting-started-hello-mule
      groupId: 21cfa350-cb56-4895-ac8c-825f4eabc928
      packaging: jar
      version: 1.0.0
    configuration:
      mule.agent.application.properties.service:
        applicationName: hello-mule-app
    desiredState: STARTED
    status: RUNNING
  replicas:
  - currentDeploymentVersion: be855c64-c3db-4043-be87-77452fd98ad1
    deploymentLocation: 345c1246-4d7a-4b4f-b2c1-fbef442d2e2d
    id: hello-mule-app-66746bbbcf-mvlfm
    reason: ""
    state: STARTED
  status: APPLIED
```