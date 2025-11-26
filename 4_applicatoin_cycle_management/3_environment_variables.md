docker run -e key=value imageName 

To pass it in the pod, we have three ways
- Direct inject 
- using configmaps
- using secrets

Direct Way

    spec:
      containers: 
        - name: nginx-container
          image: nginx
          env:
          - name: key
            vaule: "value"

ConfigMaps:

Define a k8s object type ConfigMap 
kind: ConfigMap
metadata:
    name: configMapsForApp1
data: 
    key1: value1
    key2: value2


Then, inside Pod Template, 

    spec:
      containers: 
        - name: nginx-container
          image: nginx
          envFrom:
            - configMapRef: 
                name: configMapsForApp1
          