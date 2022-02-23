# View your current Anypoint environment

### 1. Show your current Anypoint environment

```
anyctl config current-environment
```

Expected output:

```
Current environment: development
```

### 2. List all Anypoint environments available in your account

```
anyctl config get-environments                                                                                                                                                                                                     develop  ✭ ✱
```

Expected output
```
 ID                                     NAME         TYPE
 696e1c12-92a9-4044-a4ea-d45ff2db0146   development  design
 76d9d729-dda7-4bbf-99b3-42d629cfb909   quality      sandbox
 3a90bc25-d573-46ae-9cad-9f0c175f93ad   production   production
```

### 2. Change your current Anypoint environment

```
anyctl config set-environment quality
```

Expected output:

```
Updated current environment 'quality' in '/Users/arubio/.anypoint/anyconfig'.
```
