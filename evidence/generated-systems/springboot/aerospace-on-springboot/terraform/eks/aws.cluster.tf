# AWS Cluster
data "aws_eks_cluster" "aerospaceonspringboot-cluster" {
  name = "aerospaceonspringboot-cluster"
}

output "endpoint" {
  value = "${data.aws_eks_cluster.aerospaceonspringboot-cluster.endpoint}"
}

output "kubeconfig-certificate-authority-data" {
  value = "${data.aws_eks_cluster.aerospaceonspringboot-cluster.certificate_authority.0.data}"
}

output "eks_cluster_endpoint" {
  description = "Endpoint for your Kubernetes API server"
  value       = "${data.aws_eks_cluster.aerospaceonspringboot-cluster.endpoint
}


# Output for K8S
