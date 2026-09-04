# Harbormaster Blueprints

The blueprints contained in this repository are partial set of a library, to illustrate the composition of a blueprint.  Each specification is used to define the purpose, structure, and intent of the blueprint.  The complete blueprint contents are in a private repository.

## Technology Blueprints
These capture core technologies which are usually in the form of a language or framework.  They can be used directly or derived from to extend their capabilities.

For example:
> Spring Boot 3.5 >>> derives from >>> Spring Boot 2.5 >>> derives from >>> spring.core
## Solution Blueprints
These capture the features and core functionality of a specific vendor solution.  Since the vendor typically leverages other technologies, a vendor blueprint normally derives from one or more other technology blueprints.

For example:
> Event Sourcing ===> Axon Framework 4 >>> derives from >>> Spring Boot 2.5 and React
> Block Chain ===> Corda >>> derives from Axon Framework 4
>
### Other Solution Blueprints

Imagine integrating APIs from a 3rd party vendor.  If those APIs are common across a set of custom applications, capture those APIs in a blueprint to more easily take advantage of vendor capabilities.  Importantly, any vendor feature could be integrated with a domain model to provide even more integration into your set of software systems.

For example:
> Salesforce >>> *encapsulates* >>> "How we deliver Salesforce at enterprise scale."
> ServiceNow >>> *encapsulates* >>> "How we implement enterprise workflow and service management."   
> SAP Hannah >>> *encapsulates* >>> "How we implement core enterprise operations." 
> AWS Cloud Native >>> *encapsulates* >>> "How we build and operate cloud-native systems."

## Feature Set

A blueprint can dynamically be bound with a feature at system generation time. This loose coupling maximizes flexibility in system definition options while simplifying feature enablement. 

### Derivation
Since many technologies and solutions are based on other technologies and/or solutions, a blueprint can "derive" from one or more other blueprints.  The library of blueprints have foundational blueprints that technology and solution blueprints leverage.  This simplifies the creation of new blueprints or versions of a blueprint category..

### Agnostic
An application blueprint can be comprised of any language and any number of languages.  When based on a supported language, the blueprint can derive from a foundational blueprint to support common language specific features.

### Non-Software Systems
Not every blueprint needs to result in a software system.  In fact, system compilation is devoid of any assumptions or dependencies of the resulting system.  If a blueprint captures knowledge, is well-formed, and can (optionally) consume a domain model and user inputs, the resulting system could be (for example):

#### Physical system
Manufacturing line, aircraft, warehouse, power plant

#### Business system
Lending operation, claims operation, procurement

#### Organizational system
Engineering methodology, governance process, operating model

#### Regulatory system
Compliance framework, safety system, quality system
