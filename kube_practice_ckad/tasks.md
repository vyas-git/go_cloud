## Core Concepts

#### 1 . Create a namespace called 'mynamespace' and a pod with image nginx called nginx on this namespace</b>

<details>
<summary>Solution</summary>
      
      kubectl create namespace mynamespace 
      kubectl run nginx --image=nginx --restart=Never -n mynamespace
</details>

#### 2. Create the pod that was just described using YAML</b>

<details><summary>Solution</summary> 
        
      kubectl run nginx --image=nginx --restart=Never --dry-run=client -n mynamespace -o yaml > pod.yaml

</details>

#### 3. Create a busybox pod (using kubectl command) that runs the command "env". Run it and see the output</b>

<details><summary>Solution</summary>
        kubectl run nginx --image=nginx --restart=Never --dry-run=client -n mynamespace -o yaml > pod.yaml
</details>

#### 4. Create a busybox pod (using kubectl command) that runs the command "env". Run it and see the output

<details><summary>Solution</summary>
     
      kubectl run busybox --image=busybox --command --restart=Never -it --rm -- env # -it will help in seeing the output, --rm will immediately delete the pod after it exits

      # or, just run it without -it

      kubectl run busybox --image=busybox --command --restart=Never -- env

     # and then, check its logs
        kubectl logs busybox

</details>

#### 5. Create a busybox pod (using YAML) that runs the command "env". Run it and see the output

<details><summary>Solution</summary>

    kubectl run busybox --image=busybox --restart=Never --dry-run=client -o yaml --command -- env > envpod.yaml

</details>

#### 6. Create the YAML for a new ResourceQuota called 'myrq' with hard limits of 1 CPU, 1G memory and 2 pods without creating it

<details><summary>Solution</summary>
        k create quota myrq --hard=cpu=1,memory=1G,pods=2 --dry-run=client -o yaml
</details>

#### 7. Get pods on all namespaces

<details><summary>Solution</summary>

        k get pods --all-namespaces

        k get pods -A

</details>

#### 8. Create a pod with image nginx called nginx and expose traffic on port 80

<details><summary>Solution</summary>
        k run nginx --image=nginx --port=80 --restart=Never
</details>

#### 9.Change pod's image to nginx:1.7.1. Observe that the container will be restarted as soon as the image gets pulled

<details><summary>Solution</summary>
      
        k set image pod/nginx nginx=nginx:1.7.1
        k describe po nginx # you will see an event 'Container will be killed and recreated'
        k get po nginx -w # watch it
</details>
