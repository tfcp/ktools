## 持久卷相关
与volume区别 
对于有状态应用和数据持久化的应用 hostPath(仅支持单机),emptyDir不太合适了
需要容器重建之后依然可以使用之前的数据

### 基础说明
pv: 用于持久化存储的细节实现 类似于镜像(可能不太恰当)
pvc: pvc是pv的具体应用 类似pod 对于使用者不关系底层实现 直接用

### pv
文档: https://www.qikqiak.com/k8s-book/docs/33.PV.html
安装nfs使用
service nfs-server start
service nfs-server status

### pvc
文档: https://www.qikqiak.com/k8s-book/docs/34.PVC.html

### storageClass

### 常见问题
1. 1.20X版本nfs pending问题
https://www.cnblogs.com/Applogize/p/15161379.html
