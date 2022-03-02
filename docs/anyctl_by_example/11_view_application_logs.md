# View Application logs

### 1. Show the logs of the application

```
anyctl runtimemanager deployments logs hello-mule-app
```

Expected output:

```
 2022-02-24T20:59:42   hello-mule-app-76f675fd97-2z2r4   INFO   No SessionScavenger set, using defaults
 2022-02-24T20:59:42   hello-mule-app-76f675fd97-2z2r4   INFO   NO JSP Support for /, did not find org.eclipse.jetty.jsp.JettyJspServlet
 2022-02-24T20:59:42   hello-mule-app-76f675fd97-2z2r4   INFO   jetty-9.4.43.v20210629; built: 2021-06-30T11:07:22.254Z; git: 526006ecfa3af7f1a27ef3a288e2bef7ea9dd7e8; jvm 1.8.0_312-8u312-b07-0ubuntu1~20.04-b07
 2022-02-24T20:59:41   hello-mule-app-76f675fd97-2z2r4   INFO   Logging initialized @4379ms to org.eclipse.jetty.util.log.Slf4jLog
 2022-02-24T20:59:41   hello-mule-app-76f675fd97-2z2r4   INFO   Starting plugin: mule-agent-plugin
 2022-02-24T20:59:41   hello-mule-app-76f675fd97-2z2r4   INFO   Starting plugin: object-store-plugin
 2022-02-24T20:59:41   hello-mule-app-76f675fd97-2z2r4   INFO   Loaded object store core extension
 2022-02-24T20:59:41   hello-mule-app-76f675fd97-2z2r4   INFO   Registering plugin: object-store-plugin
 2022-02-24T20:59:41   hello-mule-app-76f675fd97-2z2r4   INFO   Registering plugin: mule-agent-plugin
 2022-02-24T20:59:40   hello-mule-app-76f675fd97-2z2r4   INFO   Creating plugin from folder /opt/mule/server-plugins/cloud-object-store-plugin-4.1.13-mule-server-plugin
 2022-02-24T20:59:40   hello-mule-app-76f675fd97-2z2r4   INFO   Creating plugin from folder /opt/mule/server-plugins/mule-agent-plugin
 2022-02-24T20:59:40   hello-mule-app-76f675fd97-2z2r4   INFO   Unzipping plugin file: cloud-object-store-plugin-4.1.13-mule-server-plugin.jar
```