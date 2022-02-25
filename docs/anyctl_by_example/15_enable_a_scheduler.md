## Enable a scheduler

### 1. List the schedulers available in the application

```
anyctl runtimemanager deployments scheduler get hello-brooklyn99-app
```

Expected output:

```
 FLOW NAME       TYPE                      ENABLED   SCHEDULE
 TerryJeffords   FixedFrequencyScheduler   false      Every 2000 milliseconds
 RosaDiaz        FixedFrequencyScheduler   true      Every 2000 milliseconds
 JakePeralta     FixedFrequencyScheduler   true      Every 5 seconds
 AmySantiago     CronScheduler             true      0 0/1 * 1/1 * ? *
```

### 2. List the scheduler to enable

```
anyctl runtimemanager deployments scheduler get hello-brooklyn99-app TerryJeffords
```

Expected output:

```
 FLOW NAME       TYPE                      ENABLED   DELAY   FREQUENCY   TIME UNIT
 TerryJeffords   FixedFrequencyScheduler   false     0       2000        MILLISECONDS 
```

### 3. Enable a scheduler

```
anyctl  runtimemanager deployments scheduler set hello-brooklyn99-app TerryJeffords --enabled true
```

Expected output:

```
 FLOW NAME       TYPE                      ENABLED   SCHEDULE
 TerryJeffords   FixedFrequencyScheduler   true     Every 2000 milliseconds
```

### 4. List the schedulers in the application

```
anyctl runtimemanager deployments scheduler  get hello-brooklyn99-app
```

Expected output:

```
 FLOW NAME       TYPE                      ENABLED   SCHEDULE
 RosaDiaz        FixedFrequencyScheduler   true      Every 2000 milliseconds
 JakePeralta     FixedFrequencyScheduler   true      Every 5 seconds
 AmySantiago     CronScheduler             true      0 0/1 * 1/1 * ? *
 TerryJeffords   FixedFrequencyScheduler   true     Every 2000 milliseconds
```