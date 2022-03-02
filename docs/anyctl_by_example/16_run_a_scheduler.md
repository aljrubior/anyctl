## Enable a scheduler

### 1. List the schedulers available in the application

```
anyctl runtimemanager deployments scheduler get hello-brooklyn99-app
```

Expected output:

```
 FLOW NAME       TYPE                      ENABLED   SCHEDULE
 TerryJeffords   FixedFrequencyScheduler   false     Every 2000 milliseconds
 RosaDiaz        FixedFrequencyScheduler   true      Every 2000 milliseconds
 JakePeralta     FixedFrequencyScheduler   true      Every 5 seconds
 AmySantiago     CronScheduler             true      0 0/1 * 1/1 * ? *
```

### 2. List the scheduler to run

```
anyctl runtimemanager deployments scheduler get hello-brooklyn99-app JakePeralta
```

Expected output:

```
 FLOW NAME     TYPE                      ENABLED   DELAY   FREQUENCY   TIME UNIT
 JakePeralta   FixedFrequencyScheduler   true      15      5           SECONDS  
```

### 3. Run the scheduler

```
anyctl runtimemanager deployments scheduler run hello-brooklyn99-app JakePeralta
```

Expected output:

```
Scheduler 'JakePeralta' triggered.
```
