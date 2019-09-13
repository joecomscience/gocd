# Deploy script

Create service account
```
oc create secret generic github --from-literal=username=<username> --from-literal=password=<password> --type=kubernetes.io/basic-auth
```