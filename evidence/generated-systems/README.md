# Generated System Files

Each of the directories contains all the files compiled by Harbormaster for:

- The Spring 3.5 [blueprint](../blueprints/README.md) (others include: Golang, ASP.NET, Angular, React, etc...)
- An [industry domain model](../domain-models/README.md)( 1 of the 240+ included with HM, or any custom well-formed model)
- Dynamic features applied to the target blueprint (ex: AWS, Terraform, Docker) [see system-as-code YAML](../system-as-code/README.md)
- System specific options (database engine, docker/git/terraform/kubernetes/aws params, etc...) [see system-as-code YAML](../system-as-code/README.md

Upon system creation, Harbormaster commits all files to a target repository.  
Next, invokes a [verfication system](../certification/README.md) on GitHub Workflow.
