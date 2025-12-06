Controller is a process which runs in the backgroud and its job is to monitor the status of resources and make sure what is in the etcd [desired state]
is also reflected in the cluster. Deployment controller keeps the deployment up to date, Replica Controller keeps the number of pods replica 
up to date.

From Docs:

A controller tracks at least one Kubernetes resource type. These objects have a spec field that represents the desired state. The controller(s) for that resource are responsible for making the current state come closer to that desired state.

![Controller](./images/controller.png)

In robotics and automation, a control loop is a non-terminating loop that regulates the state of a system.

Here is one example of a control loop: a thermostat in a room.

When you set the temperature, that's telling the thermostat about your desired state. The actual room temperature is the current state. The thermostat acts to bring the current state closer to the desired state, by turning equipment on or off.

In Kubernetes, controllers are control loops that watch the state of your cluster, then make or request changes where needed. Each controller tries to move the current cluster state closer to the desired state.

