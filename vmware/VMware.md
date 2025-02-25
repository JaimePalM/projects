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

VMWare was the first company to invent virtualization for the x86 architecture in the 90s.

# VMware

# vSphere

# ESXi

# vCenter

# vMotion

# vSphere High Availability

# vSphere Fault Tolerance

# vSphere Distributed Resource Scheduler

# Single Sign-On (SSO)

# Virtual Machine