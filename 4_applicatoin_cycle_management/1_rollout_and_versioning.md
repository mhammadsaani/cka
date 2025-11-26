Rollout means launching something new.
When a deployment is created, its create a rollout which create a new replicaset which is considered a revision lets say 
revision 1. When the same deployment is updated, again a new rollout is created -> replicaset -> revisoin 2


- A new version is only made when something in the spec.template section changes. 

`kubectl rollout history deployment/deployment-name`
`kubectl rollout status deployment/deployment-name`
`kubectl rollout undo deployment/deployment-name`

Types of Rolling Strategies
- RollingUpdateStrategy
- Recreate

which is defined in 

spec.strategy.type = RollingUpdate | Recreate