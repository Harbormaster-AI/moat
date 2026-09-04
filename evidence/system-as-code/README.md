# System-as-Code

## Introduction

A declarative system simplifies the process of creating a system. Hidden are the complexities of a technology blueprint, an industry domain model, and targets in the stack of technologies

## Format

 YAML; single file

## Content

Application blueprints by name, domain models by name with system and technology configurations.

## Technology

A command-line-interface using NodeJS, available on [npmjs.com](https://www.npmjs.com/package/@system-as-code/cli)

## Usage

### CLI
Like docker-compose in form, simple commands issued from an OS prompt give full control over a system creation session. Commands include:

Model Commands:
> list [options] [hint] [category] [industry]  List available models. Use hint, category, and/or industry as filters.
> profile [options] <id>                       Display details about a specific domain model.  
> industries [options]                         List all industries for the supported domain models.  
> categories [options]                         List all categories for supported domain models.   

Blueprint Commands:
> list [options] [hint]   List available blueprints.  User [hint] as a filter.  
> profile [options] <id>  Display details about a specific blueprint.  
> inputs <id>             Available user input options, to include in a System-as-Code file to allow customization of a created system.  

System Commands:
> list [options]                  List previously created systems. For authenticated users only.  
> generate [options] <yaml_file>  Generates a system using the directives of a System-as-Code YAML file.    
> certification [options] <id>    Checks the status of a system certification.    
> delete [options] <id>           Delete a previously created system.  

### Programmatic 
The CLI uses a separate NodeJS package which serves as an SDK to communicate with the system creation platform. With this it is possible to programmatically integrate system creation as part of a larger software workflow.


