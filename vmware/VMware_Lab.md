# Setup
We're going to setup a VMware lab using VMware Workstation Pro and ESXi. This will allow us to create multiple VMs in a single server.

## VMware Workstation Pro
1. Download VMware Workstation Pro from [VMware](https://softwareupdate.vmware.com/cds/vmw-desktop/ws)
2. Install VMware Workstation Pro
   
## ESXi (hypervisor)
1. Download ESXi
2. Create a new VM using the ESXi ISO file
   1. Give aprox. 500GB of disk space (this isn't physical space, it's virtual space)
   2. Two options:
      - Store virtual disk as a single file (this will consume 500GB of disk space)
      - Split virtual disk into multiple files (whatever disk utilisations is used, that will be consumed, not the whole 500GB)
   3. Customize Harware: give the resources you want to allocate to the VM
3. The VM will boot from the ESXi ISO file (continue with the installation)
   1. We may have to disable Hyper-V if we're using Windows (bcdedit /set hypervisorlaunchtype off). That is because Hyper-V and VMware Workstation can't run at the same time, they are both hypervisors.
   2. We have to map the network adapter to the physical network adapter (Settings -> Network Adapter -> Bridged)
   3. Now we get the DHCP IP address of the ESXi server and make it static (F2 -> Network Configuration)

## Windows Server 2016 (VM)
1. 