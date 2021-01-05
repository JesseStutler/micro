# k8s手工改造方案

四个服务：bookings，movies，showtimes，users全部用deployment部署pod，然后通过NodePort service暴露，因为无法实现ingress分配ip，所以只能通过nodeport的形式访问了，有点可惜（尽管部署了traefik但实际没用上）



mongo用statefulset部署，pv用的是cinder



traefik（实际没用上，只有个ui界面用来查看有哪些前后端服务）用daemonset部署