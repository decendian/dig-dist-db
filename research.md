##### Object storage: manages data as discrete units called objects
- manages data as discrete units called objects.
- Uses a UUID to locate the object. 
- An object has metadata.
- Accessed using RESTful APIs or S3-compatible interfaces.

##### Block storage
- breaks data into fixed-size chunks called blocks, 
each with a unique address, and stores them without metadata.
- iSCSI, Fibre Channel, or NVMe-oF.
- LVM or paritions
- Amazon EBS (Elastic Block Store)
- Requires OS-level file system management


##### File storage
- manages data in a hierarchical structure of directories and files.
- NFS, SMB/CIFS, or AFP
- built-in file metadata (e.g., permissions, timestamps).

##### RADOS (Reliable Autonomic Distributed Object Store)
- unstructured data over cluster of storage nodes.
- Supports atomic operations
- Each object consists of unique name (key), binary data payload (content), set of key-value metadata attributes.

##### OSD (Object Storage Daemon)

MON (Monitor)

MGR (Manager)

MDS (Metadata Server)

CephFS

RGW (RADOS Gateway)

CRUSH algorithm

Ceph Dashboard

Ceph + Kubernetes (Rook)










