# Sign in to Anypoint Platform

### 1. Create an account in Anypoint Platform

URL: https://anypoint.mulesoft.com

### 2.  Use Anyctl to sign in to your Anypoint account

```
anyctl login --username <username> --password <password>
```

Expected output:

```
Configuration file '/Users/arubio/.anypoint/anyconfig' created.
```

### 3. Inspect the anyconfig file

```
less /Users/arubio/.anypoint/anyconfig
```

The configuration file have the following structure:

```
apiVersion: v1
kind: Config
credentials:
  username: <plain-text-username>
  password: <plain-text-password>
organizationId: 3f20b860-9181-44bb-bea3-35f3ec4c749c
environments:
- id: 696e1c12-92a9-4044-a4ea-d45ff2db0146
  name: development
  kind: design
- id: 76d9d729-dda7-4bbf-99b3-42d629cfb909
  name: quality
  kind: sandbox
- id: 3a90bc25-d573-46ae-9cad-9f0c175f93ad
  name: production
  kind: production
currentContext:
  environmentId: 696e1c12-92a9-4044-a4ea-d45ff2db0146
  authorizationToken: <anypoint-authorization-token>
```

