# anyctl Cheat Sheet

### Config command

```
# Change the current Anypoint environment
anyctl config set-environment environment-name

# Display the current Anypoint environment
anyctl config current-environment
```

### Assets command

```
# Publish an asset to Anypoint Exchange
anyctl assets upload --file file-path --name asset-name --version asset-version

# List one or more versions of an asset
anyctl assets get asset-name
```

### Shared Spaces command

```
# List all Shared Spaces instances
anyctl admin sharedspaces get

# List Shared Spaces filtered by name pattern 
anyctl admin sharedspaces get shared-space-name-pattern

# List a single Shared Space
anyctl admin sharedspaces get shared-space-name

# Describe a Shared Space
anyctl admin sharedspaces describe shared-space-name
```

### Private Spaces command

```
# List Private Spaces filtered by name pattern
anyctl admin privatespaces get private-space-name-pattern

# List a Private Space by a given name
anyctl admin privatespaces get private-space-name

# List a Private Space by a given id
anyctl admin privatespaces get private-space-id

# List Managed Firewall Rules for a Private Space
anyctl admin privatespaces get-firewall-rules private-space-id/private-space-name

# List Fabrics for a Private Space
anyctl admin privatespaces get-fabrics private-space-id/private-space-name

# Describe a Private Space
anyctl admin privatespaces describe private-space-name/private-space-id
```

### Fabrics command

```
# List Fabrics filtered by name pattern
anyctl admin fabrics get fabric-name-pattern

# List a Fabric by a given name
anyctl admin fabrics get fabric-name

# List a Fabric by a given id
anyctl admin fabrics get fabric-id

# List versions information for a Fabric
anyctl admin fabrics get-versions fabric-name/fabric-id

# Describe a Fabric
anyctl admin fabrics describe fabric-name/fabric-id
```

### Targets command

```
# List targets available to deploy applications in the current environment
anyctl runtimemanager targets get

# List targets filtered by name pattern
anyctl runtimemanager targets get target-name-pattern

# List Supported Runtime Version of a Target (only for RTF)
anyctl runtimemanager targets get-supported-versions target-name

# Display the details of a Target (only for Standalone)
anyctl runtimemanager targets get-details target-name

# List addresses of Target (only for Standalone)
anyctl runtimemanager targets get-addresses target-name

# Describe a Target
anyctl runtimemanager targets describe target-name
```

### Runtime Fabrics command

```
# List all Runtime Fabrics in the current environment
anyctl runtimemanager runtimefabrics get

# List a Runtime Fabric
anyctl runtimemanager runtimefabrics get runtime-fabric-name

# List Runtime Fabrics filtered by name pattern
anyctl runtimemanager runtimefabrics get runtime-fabric-name-pattern

# List nodes of a Runtime Fabric
anyctl runtimemanager runtimefabrics get-nodes runtime-fabric-name

# Describe a Runtime Fabric
anyctl targets runtimefabrics describe runtime-fabric-name
```

### Private Spaces command

```
# List all Private Spaces in the current environment
anyctl runtimemanager privatespaces get

# List a Private Space
anyctl runtimemanager privatespaces get private-space-name

# List Private Spaces filtered by name pattern
anyctl runtimemanager privatespaces get private-space-name-pattern

# List firewall rules of a Private Space
anyctl runtimemanager privatespaces get-firewall-rules get-firewall-rules private-space-name

# Display the network configuration of a Private Space
anyctl runtimemanager privatespaces get-network private-space-name

# List Fabrics of a Private Space
anyctl runtimemanager privatespaces get-fabrics private-space-name

# Describe a Private Space
anyctl runtimemanager privatespaces describe private-space-name
```

### Deployments command

```
# List all deployments in the current environment
anyctl runtimemanager deployments get

# List Deployments by name pattern
anyctl deployments get deployment-name-pattern

# List a single deployment
anyctl deployments get deployment-name

# Deploying an application given a group id, asset name and asset version
anyctl runtimemanager deployment run deployment-name --asset group-id:asset-name:asset-version --target-name target-name --runtime-version runtime-version

# Deploying an application with the latest version of an asset given the group id and the name of the asset
anyctl deploy run deployment-name --asset group-id:asset-name --target-name target-name --runtime-version runtime-version

# Deploying an application with the latest version of an asset given the name of the asset
anyctl deployments run deployment-name --asset asset-name --target-name target-name --runtime-version runtime-version

# Deploying an application with a specific version of the asset
anyctl deployments run app-01 --asset asset-name:asset-version --target-name target-name --runtime-version runtime-version

# Deploying an application with a latest version of the asset on the latest runtime version in the target
anyctl runtimemanager deployments run deployment-name --asset asset-name --target-name target-name

# Stop a deployment
anyctl runtimemanager deployments stop deployment-name

# Start a deployment
anyctl runtimemanager deployments start deployment-name

# Delete a deployment
anyctl runtimemanager deployments delete deployment-name

# Scale a deployment
anyctl runtimemanager deployments scale deployment-name --replicas number-of-replicas

# Update the deployment asset to a specific version
anyctl runtimemanager deployments set deployment-name --asset asset-name:asset-version

# Update the deployment asset to the latest version
anyctl runtimemanager deployments set deployment-name --asset asset-name

# Copy a deployment on the same target
anyctl runtimemanager deployments clone deployment-name --with-name deployment-name

# Copy a deployment to another target in the same environment
anyctl runtimemanager deployments clone deployment-name --with-name target-name --to-target-name target-name

# Copy a deployment to another target and environment
anyctl runtimemanager deployments clone deployment-name --with-name deployment-name --to-target-name target-name --to-environment-name environment-name

# Migrate a deployment with the same name to another target in the same environment
anyctl runtimemanager deployments migrate deployment-name --to-target-name target-name

# Migrate a deployment with the same name to another target and environment
anyctl runtimemanager deployments migrate deployment-name --to-target-name target-name --to-environment-name environment-name

# Migrate a deployment with another name to another target and environment
anyctl runtimemanager deployments migrate deployment-name –with-name new-name --to-target-name target-name --to-environment-name environment-name

# Describe a deployment
anyctl runtimemanager deployments describe deployment-name
```

### Schedulers command

```
# List deployment schedulers
anyctl runtimemanager deployments schedulers get deployment-name

# List the details of a scheduler
anyctl runtimemanager deployments schedulers get deployment-name scheduler-name

# Disable a scheduler
anyctl runtimemanager deployments schedulers set deployment-name scheduler-name —enabled false

# Enable a scheduler
anyctl runtimemanager deployments schedulers set deployment-name scheduler-name —enabled true

# Run a scheduler
anyctl runtimemanager deployments schedulers run deployment-name scheduler-name

# Unmanage a scheduler
anyctl runtimemanager deployments schedulers unmanage deployment-name scheduler-name
```

### Organization command

```
# Display organization resource usage
anyctl admin organization get-usage

# Display resource usage for all master or sub organizations
anyctl organizations get-usage –all

# Display organization quotas
anyctl organizations get-quotas
```