# System Certification


![](./assets/certification.flow.png)

## Introduction
Each software system automatically undergoes a 5-step verification process.  

Once complete, the system is assigned a certification level (Platinum, Gold, Silver or Bronze) based on its overall score. Certification is meant to inform and instill confidence of the software system consumer, normally a developer.

## Process 

Each software system undergoes the following certification steps; 

### (1)-Blueprint Integrity 
Ensure using blueprint results in a software system free of generation errors. 

### (2)-Build Verification
Based on the supported language of the blueprint, tests the following: 
- directory structure
- minimum required files
- validate external dependencies
- build system into an executable artifact
  
### (3)-Runtime Verification
Using the executable artifact from the build verification phase, tests the following: 
- the executable artifact can in fact be executed
- a running executable can have it's health checked
- a running executable can be communicated with

### (4)-Delivery Verification 
Using the executable artifact from the build verification phase, tests the following for Docker: 
- login credentials
- Docker file structure
- build image 
- push image to designated docker repo

### (5)-Deploy Verification
Using Hashicorp Terraform, create the following on AWS: 

- A VPC
- Security group
- Subnet
- Identity
- A load balanced kubernetes cluster on EKS or Redhat Open shift using:
  - system defined database (MySQL, Postgres, or MongoDb)
  - the Docker image built during delivery verification 
- An EC2 instance with a default AMI
