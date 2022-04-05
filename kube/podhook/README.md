## pod hook相关

### 基础文档
    pod:
    https://kubernetes.io/zh/docs/concepts/workloads/pods/pod-lifecycle/
    pod webhook:
    https://kubernetes.io/zh/docs/tasks/configure-pod-container/attach-handler-lifecycle-event/

### pod执行流程

init container  => post start hook => liveness probe/readiness probe => pre stop hook
                =>                      main container
    
    PostStart和PreStop包括liveness和readiness是属于主容器的生命周期范围内
    而Init Container是独立于主容器之外的

