Docker Image naming convension.

docker tag containerRegistery/PathOfFolder:tag

docker login containerRegistery -u "username" -p "password"

docker push tagName



kubectl create secret docker-registry <name> \
  --docker-server=<docker-registry-server> \
  --docker-username=<docker-user> \
  --docker-password=<docker-password> \
  --docker-email=<docker-email>



spec:
  containers:
    - name: foo
      image: janedoe/awesomeapp:v1
  imagePullSecrets:
    - name: myregistrykey




---


Docker Security

Container and Host share the same kernal. Container has its own namespace and cannot see any other processes. Host has the main namesapce 
and can view all the processees including the Docker processes.

Containers are isolated via namespaces in kernal. Each container is its own namespace and can only see what is in its namespace.
But host can see all the processes. 
