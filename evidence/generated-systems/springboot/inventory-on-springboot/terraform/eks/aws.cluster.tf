# AWS Cluster
data "aws_eks_cluster" "inventoryonspringboot-cluster" {
  name = "inventoryonspringboot-cluster"
}

output "endpoint" {
  value = "${data.aws_eks_cluster.inventoryonspringboot-cluster.endpoint}"
}

output "kubeconfig-certificate-authority-data" {
  value = "${data.aws_eks_cluster.inventoryonspringboot-cluster.certificate_authority.0.data}"
}

output "eks_cluster_endpoint" {
  description = "Endpoint for your Kubernetes API server"
  value       = "${data.aws_eks_cluster.inventoryonspringboot-cluster.endpoint
}


# Output for K8S
