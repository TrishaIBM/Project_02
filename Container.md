Container : a lighweight package that has everthing to run an application
Container Image : read only template use to create container
Containerfile: all the commands to assemble container
Docker Containers allow us to control​
What resources a process can see​
What resources a process can control​
What filesystem a process uses
Processes are running programs 
containers are launched by runiing image
docker ps wll show al the rinning images
docker build is use to create an image from containerfile
we can reuse image layers
images are stacked on top of each other to form base of a container file


Kubernetes
Sceduling
where the conatiner runs
Scaling
Load balancing
master notes , working notes and they communicate through API
Node : where containers are deployed
pod: smallest deployable units
Namespace: divide cluster resources btw projects
deployment :manage life cycle
Replicaset: no. of copies of pod running rn
Daemonset : ensures copy of. specific node runs on every node of the cluster


Volume: directory accesssible to the container in a pod
PV : Storage
PVC: request for storage
Podman runs individual containers
Kubernetes manage thousand of containers
 

 podman commands :
 podman build -t my-app
 podman images : list immages
 podman rmi <image_id>
 podman run -d --name my-container -p 8080:8080 my-app
 podman ps : list the running conatiners
 Stop a container: podman stop <container_id>
 Remove a container: podman rm <container_id>
 View logs: podman logs -f <container_id>
Create the file: touch container.md

