## 常用命令

### 文档地址
http://docs.kubernetes.org.cn/541.html

### 查看应用升级历史
 kubectl rollout history deployment nginx-deploy
### 查看具体版本详情
 kubectl rollout history deployment nginx-deploy --revision=3
### 回退到当前版本的前一个版本
 kubectl rollout undo deployment nginx-deploy
### 回退到具体的版本
 kubectl rollout undo deployment nginx-deploy --to-revision=2
### 封闭节点
 kubectl cordon node
### 解锁节点
 kubectl uncordon node
### 驱逐节点
 kubectl drain node
### 标签
 kubectl label node label-name=value    # 打标签
 kubectl label node label-name-         # 删除
 kubectl label node label-name=value --overwrite # 覆盖标签
### 扩容
 kubectl scale --current-replicas=2 --replicas=3 deployment/deploy # 扩容2个pod为3个
### depoly重启
 kubectl patch deployment deploy-name -p '{"spec":{"template":{"metadata":{"labels":{"date":"1603769318"}}}}}' -n default --record=true



