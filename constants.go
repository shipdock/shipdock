package shipdock

const (
  HTTPHEADERPREFIX     = "com.navercorp."
  CLUSTER              = "com.navercorp.cluster.name"
  RACK                 = "com.navercorp.rack.name"
  OWNER                = "com.docker.swarm.owner"
  SERVICE_IP           = "com.navercorp.shipdock.lb.service_ip"
  SERVICE_PORTS        = "com.navercorp.shipdock.lb.service_ports"
  
  ENV_ENABLE_ALLOC_VIP = "ENABLE_ALLOC_VIP"
  ENV_KVSTORE          = "KVSTORE"
  ENV_CLUSTER          = "CLUSTER"
  ENV_RACK             = "RACK"
  ENV_VIP_PATH         = "VIP_PATH"
)
