# anyctl Cheat Sheet (nodesc)

### Login command

```
anyctl lo -U <username> -P <password>
```

### Config command

```
anyctl config get-environments
anyctl config set-environment <environment-name>
anyctl config current-environment
```

### Assets command

```
anyctl assets upload --file <file-path> --name <asset-name> --version <asset-version>
anyctl assets g <asset-name>
```

### Shared Spaces command

```
anyctl adm ss g
anyctl adm ss g    <shared-space-name|name-pattern>
anyctl adm s desc  <shared-space-name>
```

### Private Spaces command

```
anyctl adm ps g                  <private-space-name|private-space-id|name-pattern>
anyctl adm ps get-firewall-rules <private-space-name|private-space-id>
anyctl adm ps get-fabrics        <private-space-name|private-space-id>
anyctl adm ps desc               <private-space-name|private-space-id>
```

### Fabrics command

```
anyctl adm fa g            <fabric-name|fabric-id|name-pattern>
anyctl adm fa get-versions <fabric-name|fabric-id>
anyctl adm fa desc         <fabric-name|fabric-id>
```

### Targets command

```
anyctl arm ta g                      [target-name|name-pattern]
anyctl arm ta get-supported-versions <target-name>
anyctl arm ta get-details            <target-name>
anyctl arm ta get-addresses          <target-name>
anyctl arm ta desc                   <target-name>
```

### Runtime Fabrics command

```
anyctl arm rtf g          [runtime-fabric-name|name-pattern]
anyctl arm rtf get-nodes  <runtime-fabric-name>
anyctl arm rtf desc       <runtime-fabric-name>
```

### Private Spaces command

```
anyctl arm ps g                  [private-space-name|name-pattern]
anyctl arm ps get-firewall-rules <private-space-name>
anyctl arm ps get-network        <private-space-name>
anyctl arm ps get-fabrics        <private-space-name>
anyctl arm ps desc               <private-space-name>
```

### Deployments command

```
anyctl arm deploy g        [deployment-name|name-pattern]
anyctl arm deploy run      <deployment-name> --asset [group-id:]<asset-name>[:asset-version] --target-name <target-name> [--runtime-version runtime-version]
anyctl arm deploy stop     <deployment-name>
anyctl arm deploy start    <deployment-name>
anyctl arm deploy delete   <deployment-name>
anyctl arm deploy scale    <deployment-name> --replicas <number-of-replicas>
anyctl arm deploy set      <deployment-name> --asset <asset-name>[:asset-version]
anyctl arm deploy clone    <deployment-name> --with-name <target-name> [--to-target-name <target-name>] [--to-environment-name <environment-name>]
anyctl arm deploy migrate  <deployment-name> --to-target-name <target-name> [--to-environment-name <environment-name>] [--to-environment-name <environment-name>]
anyctl arm deploy describe <deployment-name>
anyctl arm deploy logs     <deployment-name>

anyctl arm deploy hist get  <deployment-name>
anyctl arm deploy hist diff <deployment-name> <change-version>

anyctl arm deploy apply -f <deployment-file.yaml> --plan
anyctl arm deploy apply -f <deployment-file.yaml>

```

### Schedulers command

```
anyctl arm deploy sch g         <deployment-name> [scheduler-name]
anyctl arm deploy sch set       <deployment-name> <scheduler-name> —enabled <false|true> [--all]
anyctl arm deploy sch run       <deployment-name> <scheduler-name>
anyctl arm deploy sch unmanage  <deployment-name> <scheduler-name>
```

### Organization command

```
anyctl adm org get-usage [--all]
anyctl adm org get-quotas
```