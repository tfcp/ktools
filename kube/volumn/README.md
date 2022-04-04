# Volume

## emptyDir
最基础的类型 pod重启 迁移 数据都会丢失

## hostPath
允许把pod的数据挂载到node的本地磁盘中去 如果pod需要使用Node上的文件 可以使用hostPath


## PV

### pv 回收策略
Retain，不清理, 保留 Volume（需要手动清理）
Recycle，删除数据，即 rm -rf /thevolume/*（只有 NFS 和 HostPath 支持）
Delete，删除存储资源，比如删除 AWS EBS 卷（只有 AWS EBS, GCE PD, Azure Disk 和 Cinder 支持

### 访问模式
ReadWriteOnce（RWO）：是最基本的方式，可读可写，但只支持被单个节点挂载。
ReadOnlyMany（ROX）：可以以只读的方式被多个节点挂载。
ReadWriteMany（RWX）：这种存储可以以读写的方式被多个节点共享。不是每一种存储都支持这三种方式，像共享方式，目前支持的还比较少，比较常用的是 NFS。在 PVC 绑定 PV 时通常根据两个条件来绑定，一个是存储的大小，另一个就是访问模式。

