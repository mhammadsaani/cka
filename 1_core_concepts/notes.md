4 Basic Things, every manifest should have
- apiVersion: The version of the Object, you are trying to create (Pod, Deployment, DaemonSet)
- kind: The kind of Object
- metadata: more information about the object (name of the object, labels it will have) name is string, labels is a dictionary with many key: value pairs
- spec: specification of the object we are about to create, will be different for different objects. In case of pod, it will require docker images. It will be a list, as Pods can have many containers. 
