# AWS Cluster
data "aws_eks_cluster" "governance-on-angular-cluster" {
  name = "governance-on-angular-cluster"
}

output "endpoint" {
  value = "${data.aws_eks_cluster.governance-on-angular-cluster.endpoint}"
}

output "kubeconfig-certificate-authority-data" {
  value = "${data.aws_eks_cluster.governance-on-angular-cluster.certificate_authority.0.data}"
}

output "eks_cluster_endpoint" {
  description = "Endpoint for your Kubernetes API server"
  value       = "${data.aws_eks_cluster.governance-on-angular-cluster.endpoint
}


# Output for K8S
