kubectl create secret generic cloudsql-instances-credentials --from-file=credentials.json=D:\Users\jmeixner\Code\goCode\src\TacIt-go\config\tacit-db67b0097365.json
kubectl create secret generic cloudsql-db-credentials --from-literal=username=gorm --from-literal=password=<redacted>
