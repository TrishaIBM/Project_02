Node : working machine
Replicaset: replica of each pod running in a node
Deployment : manages replicaset
Daemonset: Nodes run copy of specifin pod
Service: abstract way of seeing running pods
etcd: The cluster’s "memory bank"
kube-scheduler: The "matchmaker". It looks for new Pods and decides which Worker Node has enough CPU and RAM to run them.
kubelet: The "on-site manager". It takes orders from the Control Plane and ensures the containers in a Pod are running and healthy.
kubeproxy: It manages network rules on each node so Pods can talk to each other and the internet.
Container Runtime: starts and stops the container



List all pods: kubectl get pods
List all services: kubectl get svc
Get detailed info on a node: kubectl describe node <name>
Show all resources in all namespaces: kubectl get all -A
Delete a resource: kubectl delete -f <filename>.yaml
View container logs: kubectl logs <pod-name>
Check current context: kubectl config current-context
