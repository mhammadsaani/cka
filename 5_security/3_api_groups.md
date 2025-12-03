K8s API is the server, we interact with. MasterNodeIP:6443

Then, it has different applicatinos under it. think of it as microservices. 

/version -> gives us the versino

/api/v1 contains all the major k8s components like ns, pods, secrets, pvcs, pvs, configmaps, rcs, services

/apis it is more organized and future features will fall under it.

Under /apis, we have apps, extensions, certificates etc. [These sub groups are called api groups. ]

Under apps, we have v1/ -> deployment, replicasets, statefulsets [These are the resources inside specific groups]


we can perform actions on these resources. Like list, get, create, delete deployment, watch. [These are called verbs]

Rest on the following docs: https://docs.google.com/document/d/1K1x8oZplKg9wL24Hj4vAnRf7pkvuQmtvp8D8Rn3y3Vw/edit?tab=t.0

