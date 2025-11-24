Usually, one pod has one container. But sometimes, we require that two containers should share the same lifecycle, die and live 
togather. To overcome these cases, we can have multicontainer pods. 


Design Patterns
- Colacated Containers [when two services are dependent on each other]
- Init Container [one container perform some inital checks and when they are passed, then the other container will run. First container will also be destroyed after it performs its job]
- Sidecar Container [Sidecar is like init container pattern, but first container will continue to perform its job and will not be deleted/stopped]

Colocated and Sidecar seems similar, but one difference is that in colacated, we cannot define in which order the container will start but in sidecar we can define.

Colocated:

spec:
      containers: 
        - name: nginx-container
          image: nginx
          ports:
            - containerPort: 8080
        - name: parallel-container
          image: nginx


Init Containers:
spec:
    containers: 
        - name: nginx-container
          image: nginx
        
    initContainers:
        - name: db-checker
          image: busybox
          command: wait-db.sh

        - name: api-checker
          image: busybox
          command: wait-api.sh

    # multiple init containers can be defined. First, db-checker will run. If complete, db-checker stops and then the api-checker. If complete,api checker stops and then the nginx-container will start which is the main container in this case.


SideCar Container:

spec:
    containers: 
        - name: nginx-container
          image: nginx
        
    initContainers:
        - name: api-checker
          image: busybox
          command: wait-api.sh
          restartPolicy: always # kind of sidecar now due to this line, as container will start again once stopped. Even if the main container stops, this will still start. 

    