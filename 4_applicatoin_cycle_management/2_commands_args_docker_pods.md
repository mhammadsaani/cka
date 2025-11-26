Containers are not means to host OS but a task/process and when the process is completed, the 
container dies.


Task/Process can be
- instance of an application
- database
- script


Container only exists as long the process inside it is alive. If a webservice crashes/stops, the container dies
Therefore, usually we have a CMD ['ngnix'] like instruction at the end of the Dockerfile, and if it is starting 
a process, the container will run, otherwise, it will die.


Dockerfile 1 


FROM ubuntu

CMD ['sleep' '10']


If Dockerfile 1, and docker build -t sleep_ten .
then, 
`docker run sleep_ten `will cause the container to live for 10s. Container will run and sleep 10 command will run inside

To pass an argument, we need to provide the full new command which will overwrite the 
sleep 10 command.

`docker run sleep_ten sleep 20` will overwrite the sleep 10 command.


Dockerfile 2


FROM ubuntu

Entrypoint ['sleep']

If Dockerfile 2, and docker build -t sleep_x .

then to run we must give a argument otherwise, there will be error because only sleep command will run 

docker run sleep_x 10 -> a container will be up, sleep 10 will run 10 will be appended in the command written in entrypoint

if nothing is passed, error will come.


Bettr Approach is to use both

FROM ubuntu

Entrypoint ['sleep']

CMD [10]

now, default will be 10, and we can override as well.

Entry point is the command that is run at startup and CMD is the parameter passed.

---

For Pods, 

spec:
    containers:
        - name: test
          image: sleep_10
          command: ['sleep2.0'] # similar to entry point
          args: ['20'] # similar to cmd in docker