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
