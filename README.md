What is Docker?
Docker is a lighter and simpler alternative to a virtual machine.
Why? Because Docker uses the same operating system and the same kernel as the host device, unlike a virtual machine, which runs its own separate operating system with its own kernel.
Docker connects hardware resources (CPU, RAM, storage) with the software (applications, files) through the host’s kernel and OS. The operating system manages these resources using file descriptors, small numerical identifiers that allow programs to access files, sockets, or other input/output resources without directly dealing with the hardware. For example, when an application opens a file, the OS assigns it a file descriptor to keep track of that file and place it in the correct location in memory or storage.
In contrast, a virtual machine uses a hypervisor to virtualize hardware and run an entirely separate OS, including its own kernel. This adds more overhead and makes VMs heavier and slower to start compared to Docker containers.
Docker is more practical in many cases, such as when you need to run a project on a machine that has a different environment from the one it was developed on.
For example:
    A developer builds a project in Linux using Golang v1.20.
    Later, they need to work on it from a Windows machine without reinstalling or changing the environment.
    In this situation, they can create a Docker image that contains the exact Linux environment and Go version they need. That image can then be run in a container on the Windows machine, working exactly the same way as on Linux.
This is one of Docker’s main advantages it lets you package an application with its dependencies into an image and run it anywhere, without worrying about OS differences, and without the heavy setup of a virtual machine.

// small explication for vituelle machine: 

Software de la VM (ex: VirtualBox) → Hypervisor → Kernel de l’OS hôte → Hardware réel

Et à l’intérieur de la VM :

Software invité → Kernel invité (de la VM) → Matériel virtuel → Hypervisor → Kernel hôte → Hardware réel

// but the docker :
Docker = juste ce dont ton programme a besoin → léger et rapide.
