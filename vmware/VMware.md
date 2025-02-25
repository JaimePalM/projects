# Virtualization
In computing, virtualization refers to the act of creating a virtual (rather than actual) version of something, including virtual computer hardware platforms, storage devices, and computer network resources. 

### Vitualization Types
- **Storage Virtualization**
- **Network Virtualization**
- **Desktop Virtualization**
- **Server Virtualization**: this is which we are going to discuss in this document.

Datacenters usually uses a 1:1 ratio of applications to servers.

Example: we want to install some aplications in a server. Above this server we will have any OS (e.g. Windows Server 2022), on top of this OS we will have the applications. However, this has some disadvantages:
- If the server fails, all the applications will fail.
- It will slowdown the server.
- We can't install services like SAP, SCOM or Oracle on the Active Directory server. Whenever we install Active Directory on any server, it will not allow the other services to be installed.

Advantages of following the 1:1 ratio. We will have a dedicated server for each application (ADDS, DNS, Web, SQL...) on top of the OS. This will provide:
- If the server fails, only the application will fail, users will be able to access the other applications.
- You only have to work on the server where the application is installed.
- Better performance.
- You can install any service on any server. For example, you can install SQL server and Active Directory on other server, which wouldn't be possible in the same server.

However, this has some disadvantages:
- Requiere more hardware.
- More power consumption.
- Extra cooling.
- More space.
- Low utilization of resources.
- Too expensive.

Solution: run multiple OS on a single server. This is called **Server Virtualization**.

Between the physical hardware and the multiple OS, we have a software called **Hypervisor**. The hypervisor is the software that runs on the physical hardware and will replicate the hardware for the OS. The OS will think that it is running on a physical hardware, but it is actually running on a virtual hardware.

# VMware
VMWare was the first company to invent virtualization for the x86 architecture in the 90s. More than 80% of virtualized workloads and a large percentage of business apps are running on VMware technology. 

## VMware Products

- vCloud Director: Cloud Delivery Service
- VMware SD-WAN: WAN connections virtualization
- Horizon View: desktop
- ThinApp: application 
- NSX: network
- vSan: storage
- VSphere: server (goal of this document)

# vSphere
It is a VMware product used to virtualize the x86 servers.

## Hypervisor
It's a virtualization software which is installed on top of the physical server. It'll help us to create multiple VMs with their OS in the same server. 

There are 2 types:
- Type 1: it runs directly in top of the hardware. It's highly efficient and increase the security. E.g. OracleVM, VMware ESXi (core component of vSphere).
- Type 2: it runs on top of an OS. They are suitable for individual PC users needing to run multiple OS without formatting their current OS. It's less secure because adding a mediator adds extra vulnerabilities. E.g. Virtual Box or VMware Workstation. 

# ESXi
It's the VMware's Type 1 Hypervisor. ESXi shares the resources of the physical server, so it virtualize each one to give the VM the illusion of having the hardware exclusively for them. 

DCUI (Direct Console User Interface) is the console of ESXi. To access it you need to be physically connected to the server. It's used for initial configuration only, once created we cannot change it. To create advanced thins like VM, switches, etc. we need to use the VMware Host Client.

VMware Host Client. It'll allow use to create VMs, switches and manage network. It's used to manage only one ESXi host per session (if we have different ESXi host with different IPs we'll have to create one session per each). If you want to manage all the ESXi hosts with one session you need to use vCenter Server.

For the labs, we'll install ESXi over our OS, so we are going to use a type 2 hypervisor.

# vCenter
It's an advanced server management software. It's used to manage multiple ESXi hosts with one session centrally. There it'll be added all the IPs from the ESXi hosts. 

VCenter Server is mandatory to use the advanced features like:
- vMotion and Enhance vMotion
- HA, FT, DRS, VM cloning and much more

# vMotion
The migration of a live VM from one host to another host is called vMotion. During the migration, services won't be affected, servers will be working.

Storage vMotion: migrate the live VM files from one storage to another storage during maintenance or upgrade. For instance, if you have one HDD and you want to upgrade it, you cannot unmount it because all the VMs will stop working. So you connect an spare HDD and migrate the content with Storage vMotion and then change the old HDD.

## vMotion Enhanced
Basically it's the combination of vMotion and Storage vMotion. It migrates live VMs from one host to another, at the same time it also migrates VM files from one storage to another one. It's like a full migration.

# vSphere High Availability
It provides high availability for VMs, if any host fails, the VMs of the failed host are restarted on alternate hosts. Therefore we'll have an inevitable downtime while the VMs are started in the new host.

# vSphere Fault Tolerance
It also provides HA for VMs, but without downtime. That's possible because there'll be a live shadow instance running on another host.

# vSphere Distributed Resource Scheduler
It migrates the VMs from over-utilized ESXi hosts to under-utilized ESXi hosts. It's used to balance the resources between the hosts.

DRS uses vMotion to migrate the VMs. 

# Single Sign-On (SSO)
SSO is an authentication process that allows a user to access multiple applications with one set of login credentials. These applications can be, for example, vCenter Server, vCloud Director.

It supports multiple identity sources like Active Directory (Microsoft) and OpenLDAP (Lightweight Directory Access Protocol) (Linux).

# Virtual Machine
A VM is a software computer that, like a physical computer, runs an operating system and applications. The VM is composed of a set of specification and configuration files and is backed by the physical resources of a host.

Without VM, a single OS owns all the hardware resources. With VM, multiple OS can share the hardware resources. Each VM will see the hardware as if it were the only one using it. The hypervisor will manage the hardware resources and allocate them to the VMs.

The VM key files are: configuration file (.vmx), virtual disk (.vmdk), NVRAM file (.nvram), log file (.log).
- Configuration file (.vmx): it contains information about the VM like how many CPUs, NICs or RAM are available.
- NVRAM file (.nvram): it contains the BIOS information. Changes in the BIOS will be saved in this file.
- Log file (.log): it contains the logs of the VM.
- Virtual disk (.vmdk): it contains the OS and applications.

## Characteristics
VMs have four primary characteristics:
- **Compatibility**: VMs are compatible with all x86 computers, regardless of the operating system or hardware below it. 
- **Isolation**: VMs are isolated from each other, so if one VM fails, the others will continue working. One VM can't see the other VMs.
- **Encapsulation**: VMs are encapsulated into files, so they can be "easily" moved from one host to another host.
- **Hardware Independence**: VMs are independent of the hardware, so they can be moved from one host to another host without any problem. For example, you can move a VM from a Dell server to an HP server or from Intel to AMD.

## Advantages
Using VMs has several advantages:
- **Cloning**: the VM can be cloned to create a new VM, both exactly the same. Clones are used for testing, development, deployment, etc. You only have to create one VM and then clone it.
- **Flexibility**: you can move the VM from one host to another host, from one storage to another storage, etc.
- **Security**: malware can't affect the host because it's running on an isolated VM. If the VM is infected, you can delete it and create a new one.
- **Backup**: VMs can be backed up periodically easily and them can be restored at any time.
- **Restore**: if the VM fails, you can restore it from the backup, in the same host or in another host.
- **Resource Utilization**: you can run multiple VMs on a single host, so you can utilize the resources of the host efficiently. That is, if you run one application on one server, you will have a lot of resources wasted. If you run multiple applications on the same server, you will utilize the resources efficiently.