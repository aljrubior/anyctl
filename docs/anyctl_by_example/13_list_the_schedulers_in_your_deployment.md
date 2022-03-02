## List the schedulers in your deployments

### 1. Publish an application with schedulers to Anypoint Exchange

```
anyctl assets upload --file /tmp/getting-started-hello-brooklyn99.jar --name getting-started-hello-brooklyn99 --version 1.0.0
```

Expected output:

```
Asset 'getting-started-hello-brooklyn99' v1.0.0 created.
```

### 2. Deploy the Mule Application getting-started-hello-brooklyn99

```
anyctl runtimemanager deployments run hello-brooklyn99-app \
   --asset getting-started-hello-brooklyn99 \
   --target-name private-space-finance-qa
```

Expected output:

```
Deployment 'hello-brooklyn99-app' created.
```

### 3. Show the current state of the deployment

```
anyctl runtimemanager deployments get hello-brooklyn99-app
```

Expected output:

```
 NAME                   REPLICAS   STATUS    TARGET                      RUNTIME            ASSET
 hello-brooklyn99-app   1/1        APPLIED   runtime-mgr-private-space   4.4.0:20220124-4   getting-started-hello-brooklyn99:1.0.0
```

### 4. List the schedulers available in the application

```
anyctl runtimemanager deployments scheduler  get hello-brooklyn99-app
```

Expected output:

```
 FLOW NAME       TYPE                      ENABLED   SCHEDULE
 TerryJeffords   FixedFrequencyScheduler   true      Every 2000 milliseconds
 RosaDiaz        FixedFrequencyScheduler   true      Every 2000 milliseconds
 JakePeralta     FixedFrequencyScheduler   true      Every 5 seconds
 AmySantiago     CronScheduler             true      0 0/1 * 1/1 * ? *
```