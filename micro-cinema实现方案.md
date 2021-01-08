# micro-cinema实现方案

github原demo：

https://github.com/mmorejon/microservices-docker-go-mongodb.git

其有两套方案，一套是通过docker-compose单机上群起一堆服务，一套是通过helm部署到k8s上（但是其实现时的helm版本过老，现在版本的helm很多参数已经修改，其给的方案已经不可用了），所以方便学习k8s的话，可以通过手工部署到k8s上

## 手工部署到k8s方案

四个服务：bookings，movies，showtimes，users全部用deployment部署pod。

mongo用statefulset部署，pv用的是openstack cinder

traefik用daemonset部署

service的type最好使用load balancer，如果你使用的是GKE这样的平台会给你提供external ip，但是如果你是自己搭建的k8s集群，不会直接提供external ip，可以采用metallb，地址：https://metallb.universe.tf/

但是笔者的k8s集群是搭建在openstack上的，load balancer的方案不太行，所以service的type是node port，这样的话traefik就没用上了，直接通过node的指定port就能访问service



Yaml全部放在k8s-yaml文件夹中了，有一些ingress的yaml是直接apply的traefik官网上的，地址：https://doc.traefik.io/traefik/v1.7/user-guide/kubernetes/#submitting-an-ingress-to-the-cluster

