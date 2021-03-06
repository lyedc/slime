# Slime

![slime-logo](logo/slime-logo.png)

---

slime是针对istio的CRD控制器。旨在通过简单配置，自动更便捷的使用istio/envoy高阶功能。不同功能对应slime中的不同模块，目前slime包含了三个子模块：

**[配置懒加载](#配置懒加载):** 无须配置SidecarScope，自动按需加载配置/服务发现信息  

**[Http插件管理](#http插件管理):** 使用新的的CRD pluginmanager/envoyplugin包装了可读性，可维护性差的envoyfilter,使得插件扩展更为便捷  

**[自适应限流](#自适应限流):** 实现了本地限流，同时可以结合监控信息自动调整限流策略

后续我们会开放更多的功能模块~

## 安装slime-boot
在使用slime之前，需要安装slime-boot，通过slime-boot，可以方便的安装和卸载slime模块。 执行如下命令：
```
kubectl create ns mesh-operator
kubectl apply -f https://raw.githubusercontent.com/ydh926/slime/master/install/crds.yaml
kubectl apply -f https://raw.githubusercontent.com/ydh926/slime/master/install/slime-boot-install.yaml
```

## 配置懒加载
### 安装和使用

**请先按照[安装slime-boot](#安装slime-boot)小节的指引安装`slime-boot`**     

1. 使用Slime的配置懒加载功能需打开Fence模块，同时安装附加组件，如下：
```yaml
apiVersion: config.netease.com/v1alpha1
kind: SlimeBoot
metadata:
  name: lazyload
  namespace: mesh-operator
spec:
  # Default values copied from <project_dir>/helm-charts/slimeboot/values.yaml\
  module:
    - fence:
        enable: true
        wormholePort:
        - {{port1}} # 业务svc的端口
        - {{port2}}
        - ...
      name: slime-fence
 component:
   globalSidecar:
     enable: true
     namespace:
       - default # 替换为业务所在的namespace
       - {{you namespace}}
   pilot:
     enable: true
     image:
       repository: docker.io/bcxq/pilot
       tag: preview-1.3.7-v0.0.1
   reportServer:
     enable: true
     image:
       repository: docker.io/bcxq/mixer
       tag: preview-1.3.7-v0.0.1  
```
2. 确认所有组件已正常运行：
```
$ kubectl get po -n mesh-operator
NAME                                    READY     STATUS    RESTARTS   AGE
global-sidecar-pilot-796fb554d7-blbml   1/1       Running   0          27s
lazyload-fbcd5dbd9-jvp2s                1/1       Running   0          27s
report-server-855c8cf558-wdqjs          2/2       Running   0          27s
slime-boot-68b6f88b7b-wwqnd             1/1       Running   0          39s
```
```
$ kubectl get po -n {{your namespace}}
NAME                              READY     STATUS    RESTARTS   AGE
global-sidecar-785b58d4b4-fl8j4   1/1       Running   0          68s
```
3. 打开配置懒加载：
在namespace上打上`istio-dependency-servicefence=true`的标签。
```shell
kubectl label ns {{your namespace}} istio-dependency-servicefence=true
```
为需要开启懒加载的服务打上标签`istio.dependency.servicefence/status: "true"`。
```shell
kubectl annotate svc {{your svc}} istio.dependency.servicefence/status=true
```
4. 确认懒加载已开启
执行`kubectl get sidecar {{svc name}} -oyaml`，可以看到对应服务生成了一个sidecar，如下：
```yaml
apiVersion: networking.istio.io/v1beta1
kind: Sidecar
metadata:
  name: {{your svc}}
  namespace: {{your ns}}
  ownerReferences:
  - apiVersion: microservice.netease.com/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: ServiceFence
    name: {{your svc}}
spec:
  egress:
  - hosts:
    - istio-system/*
    - mesh-operator/*
    - '*/global-sidecar.{{your ns}}.svc.cluster.local'
  workloadSelector:
    labels:
      app: {{your svc}}
```

### 卸载
1. 删除slime-boot配置
2. 删除servicefence配置
```shell
for i in $(kubectl get ns);do kubectl delete servicefence -n $i --all;done
```
### 示例: 为bookinfo开启配置懒加载
1. 安装 istio ( > 1.8 )
2. 安装 slime 
```shell
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/ydh926/slime/master/install/easy_install_lazyload.sh)"
```
3. 确认所有组件已正常运行：
```
$ kubectl get po -n mesh-operator
NAME                                    READY     STATUS    RESTARTS   AGE
global-sidecar-pilot-796fb554d7-blbml   1/1       Running   0          27s
lazyload-fbcd5dbd9-jvp2s                1/1       Running   0          27s
report-server-855c8cf558-wdqjs          2/2       Running   0          27s
slime-boot-68b6f88b7b-wwqnd             1/1       Running   0          39s
```

```
$ kubectl get po 
NAME                              READY     STATUS    RESTARTS   AGE
global-sidecar-785b58d4b4-fl8j4   1/1       Running   0          68s
```
4. 在default namespace下安装bookinfo
5. 开启配置懒加载
```shell
kubectl label ns default istio-dependency-servicefence=true
```
```shell
kubectl annotate svc productpage istio.dependency.servicefence/status=true
kubectl annotate svc reviews istio.dependency.servicefence/status=true
kubectl annotate svc details istio.dependency.servicefence/status=true
kubectl annotate svc ratings istio.dependency.servicefence/status=true
```
6. 确认sidecarScope已经生成
```
$ kubectl get sidecar
NAME          AGE
details       12s
kubernetes    11s
productpage   11s
ratings       11s
reviews       11s
```
```
$ kubectl get sidecar productpage -oyaml
apiVersion: networking.istio.io/v1beta1
kind: Sidecar
metadata:
  name: productpage
  namespace: default
  ownerReferences:
  - apiVersion: microservice.netease.com/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: ServiceFence
    name: productpage
spec:
  egress:
  - hosts:
    - istio-system/*
    - mesh-operator/*
    - '*/global-sidecar.default.svc.cluster.local'
  workloadSelector:
    labels:
      app: productpage
```
7. 访问productpage并查看accesslog
```
[2021-01-04T07:12:24.101Z] "GET /details/0 HTTP/1.1" 200 - "-" 0 178 36 35 "-" "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0" "83793ccf-545c-4cc2-9a48-82bb70d81a2a" "details:9080" "10.244.3.83:9080" outbound|9080||global-sidecar.default.svc.cluster.local 10.244.1.206:42786 10.97.33.96:9080 10.244.1.206:40108 - -
[2021-01-04T07:12:24.171Z] "GET /reviews/0 HTTP/1.1" 200 - "-" 0 295 33 33 "-" "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0" "010bb2bc-54ab-4809-b3a0-288d60670ded" "reviews:9080" "10.244.3.83:9080" outbound|9080||global-sidecar.default.svc.cluster.local 10.244.1.206:42786 10.99.230.151:9080 10.244.1.206:51512 - -
```
成功访问, 访问日志显示后端服务是global-sidecar.

8. 查看productpage的sidecarScope
```
$ kubectl get sidecar productpage -oyaml
apiVersion: networking.istio.io/v1beta1
kind: Sidecar
metadata:
  name: productpage
  namespace: default
  ownerReferences:
  - apiVersion: microservice.netease.com/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: ServiceFence
    name: productpage
spec:
  egress:
  - hosts:
    - '*/details.default.svc.cluster.local'
    - '*/reviews.default.svc.cluster.local'
    - istio-system/*
    - mesh-operator/*
    - '*/global-sidecar.default.svc.cluster.local'
  workloadSelector:
    labels:
      app: productpage
```
reviews 和 details 被自动加入！

9. 再次访问productpage
```
[2021-01-04T07:35:57.622Z] "GET /details/0 HTTP/1.1" 200 - "-" 0 178 2 2 "-" "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0" "73a6de0b-aac9-422b-af7b-2094bd37094c" "details:9080" "10.244.7.30:9080" outbound|9080||details.default.svc.cluster.local 10.244.1.206:52626 10.97.33.96:9080 10.244.1.206:47396 - default
[2021-01-04T07:35:57.628Z] "GET /reviews/0 HTTP/1.1" 200 - "-" 0 379 134 134 "-" "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0" "edf8c7eb-9558-4d1e-834c-4f238b387fc5" "reviews:9080" "10.244.7.14:9080" outbound|9080||reviews.default.svc.cluster.local 10.244.1.206:42204 10.99.230.151:9080 10.244.1.206:58798 - default
```
访问成功, 后端服务是reviews和details.

## HTTP插件管理
// TODO
### 安装
// TODO
### 卸载
// TODO

## 自适应限流
### 安装和使用

**注意:** 自适应限流功能可以对接envoy社区支持的限流插件`envoy.filters.http.local_ratelimit`，也可以对接网易自研插件`com.netease.local_flow_control`。envoy社区的限流插件暂不支持HeaderMatch的配置，使用`com.netease.local_flow_control`插件前需确认envoy二进制中是否包含该插件。      

**请先按照[安装slime-boot](#安装slime-boot)小节的指引安装`slime-boot`**  

使用Slime的自适应限流功能需打开Limiter模块：
```yaml
apiVersion: config.netease.com/v1alpha1
kind: SlimeBoot
metadata:
  name: limiter
  namespace: mesh-operator
spec:
  # Default values copied from <project_dir>/helm-charts/slimeboot/values.yaml\
  module:
    - limiter:
        enable: true
        backend: 1
      name: slime-limiter
  //...      
```

根据限流规则为目标服务定义SmartLimite资源，如下所示：

```yaml
apiVersion: microservice.netease.com/v1alpha1
kind: SmartLimiter
metadata:
  name: test-svc
  namespace: default
spec:
  descriptors:
  - action:
      quota: "3"     # 配额数
      fill_interval:
        seconds: 1   # 统计配额的周期
    condition: "true"
```
上述配置为test-svc服务限制了每秒3次的请求。将配置提交之后，该服务下实例的状态信息以及限流信息会显示在`status`中，如下：

```yaml
apiVersion: microservice.netease.com/v1alpha1
kind: SmartLimiter
metadata:
  name: test-svc
  namespace: default
spec:
  descriptors:
  - action:
      quota: "3"
      fill_interval:
        seconds: 1
    condition: "true"
status:
  endPointStatus:
    cpu: "398293"        # 业务容器和sidecar容器占用CPU之和 
    cpu_max: "286793"    # CPU占用最大的容器
    memory: "68022"      # 业务容器和sidecar容器内存占用之和  
    memory_max: "55236"  # 内存占用最大的容器
    pod: "1"
  ratelimitStatus:
  - action:
      fill_interval:
        seconds: 1
      quota: "3"
```
#### 基于监控的自适应限流

可以将监控信息条目配置到`condition`中，例如希望cpu超过300m时触发限流，可以进行如下配置：

```yaml
apiVersion: microservice.netease.com/v1alpha1
kind: SmartLimiter
metadata:
  name: test-svc
  namespace: default
spec:
  descriptors:
  - action:
      quota: "3"
      fill_interval:
        seconds: 1
    condition: {cpu}>300000 # cpu的单位为ns，首先会根据endPointStatus中cpu的值将算式渲染为398293>300000
status:
  endPointStatus:
    cpu: "398293"        # 业务容器和sidecar容器占用CPU之和 
    cpu_max: "286793"    # CPU占用最大的容器
    memory: "68022"      # 业务容器和sidecar容器内存占用之和  
    memory_max: "55236"  # 内存占用最大的容器
    pod: "1"
  ratelimitStatus:
  - action:
      fill_interval:
        seconds: 1
      quota: "3"
```

condition中的算式会根据endPointStatus的条目进行渲染，渲染后的算式若计算结果为true，则会触发限流。

#### 服务限流
由于缺乏全局配额管理组件，我们无法做到精确的服务限流，但是假定负载均衡理想的情况下，实例限流数=服务限流数/实例个数。test-svc的服务限流数为3，那么可以将quota字段配置为3/{pod}以实现服务级别的限流。在服务发生扩容时，可以在限流状态栏中看到实例限流数的变化。
```yaml
apiVersion: microservice.netease.com/v1alpha1
kind: SmartLimiter
metadata:
  name: test-svc
  namespace: default
spec:
  descriptors:
  - action:
      quota: "3/{pod}" # 算式会根据endPointStatus中pod值渲染为3/3
      fill_interval:
        seconds: 1
    condition: "{cpu}>300000" 
    match:
    - exact_match: user
      invert_match: false
      name: Bob
status:
  endPointStatus:
    cpu: "xxxxx"        
    cpu_max: "xxxx"    
    memory: "xxx"       
    memory_max: "xx" 
    pod: "3" # test-svc的endpoint扩容成了3
  ratelimitStatus:
  - action:
      fill_interval:
        seconds: 1
      quota: "1" 显然，3/3=1
```
### 卸载
1. 删除slime-boot配置
2. 删除smartlimiter配置
```
for i in $(kubectl get ns);do kubectl delete smartlimiter -n $i --all;done
```
### 示例
以bookinfo为例，介绍配置限流如何使用。开始之前，确保已经安装了istio1.8版本以及slime-boot。示例中的bookinfo安装在default namespace。   

**安装bookinfo**
```
$ kubectl apply -f samples/bookinfo.yaml 
```

**清理slime-boot资源**
```
$ kubectl delete slimeboot -n mesh-operator --all
```

**安装限流模块**
```
$ kubectl apply -f samples/limiter-install.yaml
```

**为reviews服务设置限流规则**
```
$ kubectl apply -f samples/reviews-svc-limiter.yaml  
```

**确认配置已经创建**
```
$ kubectl get smartlimiter reviews -oyaml
apiVersion: microservice.netease.com/v1alpha1
kind: SmartLimiter
metadata:
  name: reviews
  namespace: default
spec:
  descriptors:
  - action:
      quota: "3/{pod}"
      fill_interval:
        seconds: 10
    condition: "true"
```
该配置表明review服务会被限制为10s访问三次

**确认对应的EnvoyFilter是否创建**
```
$ kubectl get envoyfilter  reviews.default.local-ratelimit -oyaml
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  creationTimestamp: "2021-01-05T07:28:01Z"
  generation: 1
  name: reviews.default.local-ratelimit
  namespace: default
  ownerReferences:
  - apiVersion: microservice.netease.com/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: SmartLimiter
    name: reviews
    uid: 5eed7271-e8a9-4eda-b5d8-6cd2dd6b3659
  resourceVersion: "59145684"
  selfLink: /apis/networking.istio.io/v1alpha3/namespaces/default/envoyfilters/reviews.default.local-ratelimit
  uid: 04549089-4bf5-4200-98ae-59dd993cda9d
spec:
  configPatches:
  - applyTo: HTTP_FILTER
    match:
      context: SIDECAR_INBOUND
      listener:
        filterChain:
          filter:
            name: envoy.http_connection_manager
            subFilter:
              name: envoy.router
    patch:
      operation: INSERT_BEFORE
      value:
        name: envoy.filters.http.local_ratelimit
        typed_config:
          '@type': type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit
          value:
            filter_enabled:
              default_value:
                numerator: 100
              runtime_key: local_rate_limit_enabled
            filter_enforced:
              default_value:
                numerator: 100
              runtime_key: local_rate_limit_enforced
            stat_prefix: http_local_rate_limiter
            token_bucket:
              fill_interval:
                seconds: "10"
              max_tokens: 1
  workloadSelector:
    labels:
      app: reviews
```
由于review服务有3个实例，因此每个实例的10s只能获得1个配额

**访问productpage页面**     
多次访问productpage页面将触发限流，查看productpage的accesslog可以更直观的看出限流效果：

```
$ kubectl logs {productpage pod} -c istio-proxy
[2021-01-05T07:29:03.986Z] "GET /reviews/0 HTTP/1.1" 429 - "-" 0 18 10 10 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36" "d59c781a-f62c-4e98-9efe-5ace68579654" "reviews:9080" "10.244.8.95:9080" outbound|9080||reviews.default.svc.cluster.local 10.244.1.206:35784 10.99.230.151:9080 10.244.1.206:39864 - default                              
```