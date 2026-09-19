# Generated System Files

Each of the directories contains all the files compiled by Harbormaster as evidence for the thesis:

- [Spring Boot 3.5](./springboot)
- [Golang](./golang)  
- [Angular 22](./angular)
- [Django](./django)  
- [RubyOnRails](./rubyonrails) 
- [Axon Framework](./axon)

Each generated system is the result of the following:

- A [blueprints](../blueprints/README.md) used to capture the knowledge of a technical SME
- An [industry domain model](../domain-models/README.md)( 1 of the 240+ included with HM, or any custom well-formed model) to capture the business context- Dynamic features applied to the target blueprint (ex: AWS, Terraform, Docker) [see system-as-code YAML](../system-as-code/README.md)
- System specific features (database engine, docker/git/terraform/kubernetes/aws params, etc...) [see system-as-code YAML](../system-as-code/README.md) to customize the resulting system.

Upon system creation, Harbormaster commits all files to a target repository.  
Next, to validate the system, invokes a separate [verfication system](../certification/README.md) on GitHub Workflow.
