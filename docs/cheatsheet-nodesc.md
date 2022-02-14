# anyctl Cheat Sheet

### Config command

```
anyctl config set-environment <environment-name>
anyctl config current-environment
```

### Assets command

```
anyctl assets upload --file <file-path> --name <asset-name> --version <asset-version>
anyctl assets get <asset-name>
```

### Shared Spaces command

```
anyctl admin sharedspaces get
anyctl admin sharedspaces get <shared-space-name|name-pattern>
anyctl admin sharedspaces describe <shared-space-name>
```

### Private Spaces command

```
anyctl admin privatespaces get <private-space-name|private-space-id|name-pattern>
anyctl admin privatespaces get-firewall-rules <private-space-name|private-space-id>
anyctl admin privatespaces get-fabrics <private-space-name|private-space-id>
anyctl admin privatespaces describe <private-space-name|private-space-id>
```

### Fabrics command

```
anyctl admin fabrics get <fabric-name|fabric-id|name-pattern>
anyctl admin fabrics get-versions <fabric-name|fabric-id>
anyctl admin fabrics describe <fabric-name|fabric-id>
```

### Targets command

```
anyctl runtimemanager targets get [target-name|name-pattern]
anyctl runtimemanager targets get-supported-versions <target-name>
anyctl runtimemanager targets get-details <target-name>
anyctl runtimemanager targets get-addresses <target-name>
anyctl runtimemanager targets describe target-name
```

### Runtime Fabrics command

```
anyctl runtimemanager runtimefabrics get [runtime-fabric-name|name-pattern]
anyctl runtimemanager runtimefabrics get-nodes runtime-fabric-name
anyctl targets runtimefabrics describe runtime-fabric-name
```

### Private Spaces command

```
anyctl runtimemanager privatespaces get [private-space-name|name-pattern]
anyctl runtimemanager privatespaces get-firewall-rules <private-space-name>
anyctl runtimemanager privatespaces get-network <private-space-name>
anyctl runtimemanager privatespaces get-fabrics <private-space-name>
anyctl runtimemanager privatespaces describe <private-space-name>
```

### Deployments command

```
anyctl runtimemanager deployments get [deployment-name|name-pattern]
anyctl runtimemanager deployment run <deployment-name> --asset [group-id:]<asset-name>[:asset-version] --target-name <target-name> [--runtime-version runtime-version]
anyctl runtimemanager deployments stop <deployment-name>
anyctl runtimemanager deployments start <deployment-name>
anyctl runtimemanager deployments delete <deployment-name>
anyctl runtimemanager deployments scale <deployment-name> --replicas <number-of-replicas>
anyctl runtimemanager deployments set <deployment-name> --asset <asset-name>[:asset-version]
anyctl runtimemanager deployments clone <deployment-name> --with-name <target-name> [--to-target-name <target-name>] [--to-environment-name <environment-name>]
anyctl runtimemanager deployments migrate <deployment-name> --to-target-name <target-name> [--to-environment-name <environment-name>] [--to-environment-name <environment-name>]
anyctl runtimemanager deployments describe <deployment-name>
```

### Schedulers command

```
anyctl runtimemanager deployments schedulers get <deployment-name> [scheduler-name]
anyctl runtimemanager deployments schedulers set <deployment-name> <scheduler-name> —enabled <false|true>
anyctl runtimemanager deployments schedulers run <deployment-name> <scheduler-name>
anyctl runtimemanager deployments schedulers unmanage <deployment-name> <scheduler-name>
```

### Organization command

```
anyctl admin organization get-usage [--all]
anyctl organizations get-quotas
```