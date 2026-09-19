# Generated System Files

Each of the directories contains all the files compiled by Harbormaster for:

Spring 3.5, ASP.NET-10, Golang, Angular 22, Django, RubyOnRails [blueprints](../blueprints/README.md)  
- An [industry domain model](../domain-models/README.md)( 1 of the 240+ included with HM, or any custom well-formed model)
- Dynamic features applied to the target blueprint (ex: AWS, Terraform, Docker) [see system-as-code YAML](../system-as-code/README.md)
- System specific options (database engine, docker/git/terraform/kubernetes/aws params, etc...) [see system-as-code YAML](../system-as-code/README.md

Upon system creation, Harbormaster commits all files to a target repository.  
Next, to validate the system, invokes a separate [verfication system](../certification/README.md) on GitHub Workflow.
