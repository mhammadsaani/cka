To store secrets information as envs [database passwords, etc], we use secrets


apiVersion: v1
kind: Secret
metadata:
    name: secret-data
data:
    password: alksfjd # data must be in encoded form


Inside Pod:

Then, inside Pod Template, 

    spec:
      containers: 
        - name: nginx-container
          image: nginx
          envFrom:
            - secretRef: 
                name: secret-data

    
---

Secrets are not encrypted, so it is not safer in that sense. However, some best practices around using secrets make it safer. 
As in best practices like:

Not checking in secret object definition files to source code repositories.
Enabling Encryption at Rest for Secrets so they are stored encrypted in ETCD.
