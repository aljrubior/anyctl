# Publish a Mule Application to Anypoint Exchange

### 1. Build a Hello World Mule Application

Resources:
* https://docs.mulesoft.com/mule-runtime/4.4/mule-app-dev-hellomule
* https://developer.mulesoft.com/tutorials-and-howtos/getting-started/hello-mule/
* https://github.com/mulesoft-developers/getting-started-hello-mule

Download File:
* https://mulesoft.s3.us-west-1.amazonaws.com/tutorials-jars/getting-started-hello-mule.jar

### 2. Upload the Mule Application file to Anypoint Exchange

```
anyctl assets upload --file /tmp/getting-started-hello-mule.jar --name getting-started-hello-mule --version 1.0.0
```

Expected output

```
Asset 'getting-started-hello-mule' v1.0.0 created.
```

### 3. List the Mule Application from Anypoint Exchange

```
anyctl assets get getting-started-hello-mule                                                                                                                                                                                       develop  ✭ ✱
```

Expected output

```
 NAME                         VERSION   STATUS      RUNTIME   UPDATED AT                 CREATED BY
 getting-started-hello-mule   1.0.0     published   4.3.0     2022-02-23T02:19:05.277Z   foo-bar
```