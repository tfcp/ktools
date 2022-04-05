## 服务发现知识点

### sessionAffinity作用
会话亲和性 默认为 none, ClientIP
当需要保持会话状态或者状态保存要求 需要使用ClientIP

    使用示例(使用base-java.yaml):
    curl <cluster-ip>/hostname;echo
    clientIP可以看到同一个client请求的pod永远是同一个

### 类型

    Cluster-ip: 默认类型 内部访问 也是其余类型的基础
    Nodeport: 会在每个Node上暴露一个对外端口(30000-32767 该端口范围可以修改) 通过该端口对外提供服务
    LoadBalancer: 云厂商提供的负载均衡器 对外提供服务 用于自动负载均衡
    
### ingress

    为了节省lb 能够一个ip暴露更多的服务 可以使用ingress, 可以理解为service的service
    使用方式: 1. nginx-ingress绑定一个slb 2. nginx中和普通服务一样正常代理到service中 3. 新的域名指向nginx的slb
    好处: 节省了宝贵的静态 IP 和 LoadBalancers 资源。
    

