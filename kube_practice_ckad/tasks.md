## Core Concepts

<details>
<summary>
<b>Create a namespace called 'mynamespace' and a pod with image nginx called nginx on this namespace</b>
</summary>
      <code>
      
      kubectl create namespace mynamespace 
      kubectl run nginx --image=nginx --restart=Never -n mynamespace
</code>
</details>

<details>
<summary>
<b>Create the pod that was just described using YAML</b>
</summary>
      <code>
      
      kubectl run nginx --image=nginx --restart=Never --dry-run=client -n mynamespace -o yaml > pod.yaml

</code>
</details>

<details>
<summary>
    <b>Create a busybox pod (using kubectl command) that runs the command "env". Run it and see the output</b>
</summary>

<code>
      kubectl run nginx --image=nginx --restart=Never --dry-run=client -n mynamespace -o yaml > pod.yaml
</code>

</details>

<details>
<summary>
<b>Create a busybox pod (using kubectl command) that runs the command "env". Run it and see the output</b>
</summary>
      <code> 
     
      kubectl run busybox --image=busybox --command --restart=Never -it --rm -- env # -it will help in seeing the output, --rm will immediately delete the pod after it exits

      # or, just run it without -it

      kubectl run busybox --image=busybox --command --restart=Never -- env

     # and then, check its logs
        kubectl logs busybox

</code>

</details>

<details>
<summary>
<b>Create a busybox pod (using YAML) that runs the command "env". Run it and see the output
</b>
</summary>
<code>

    kubectl run busybox --image=busybox --restart=Never --dry-run=client -o yaml --command -- env > envpod.yaml

</code>

</details>
