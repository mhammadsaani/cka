Scaling:
 - Adding more resources [cpu, memory,  disk]

Horizontal [Adding more instances of application ] Pod Autoscaling vs Vertical [Adding infra in existing infra] Pod Scaling

Scaling Infra also have two categories.
- Horizontal adding more nodes in k8s cluster 
- Vertical  adding more resources in existing node

Scaling workloads
- Horizontal, adding more pods
- Vertical, increasing allocated resources.

Two ways to achive this,  
- Manual and Automatic



We can configure HPA for pods and define a threshold like the cpu-usage, The HPA will continuously 
monitor the pod and if the threshold is met, it will provision a new pod.
For this to happen, we must have metrics server.


there must be request, limits set.


Imperative Way ```kubectl autoscale deployment deploymentName --cpu-usage=50 --min=1 --max=5```
kubectl scale deployment flask-web-app --replicas=3
Also, we can do this via the manifest.


Inplace Pod Resizing 

If we change the resource limit, it will decrease the old pod and then create the new pod. [in case of vertical scaling] but this can be
dangerous in case of statefulset because they have data. To avoid this, we can enable inplace pod resizing. we enable InPlacePodVerticalScaling=true.
then, in manifest we can define certain attributues like cpu restartpolicy: notrequired which will cause the pod to not restart.

Init and ephemeral containers cannot be resized

Vertical Pod Autoscaler - It doesn't come by default, so need to manually deploy it.

When we deploy it, there are three pods
- vpa-admission-control
- vpa-recommender
- vpa-updater

VPA-Recommender continuously monitors the pod usage from the metrics api and tells to updater
Updater, idenitfy the pods with sub optimal resources, evicts the pod
Admission control takes the recomendation from the recommender and mutate the manifest


Then, the new pod starts with the updated values.


We make a manifest file, tell which deployment to monitor, give max and min limits.
One important thing is
updatePolicy:
    updateMode: values 

values can be
Off - only recommends, no change
Initial - only changes on pod creation not later.  updater doesn't evict pod,if pod is evicted for some other reason, it only then apply the recommended changes
Recreate - evicts pod if usage goes above.
Auto -  similar to recreate, only works if inplace update is enabled. otherwise, only evict pod and no change.

