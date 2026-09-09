# Proof of Hypotheses
The contents of this directory exists to serve as a proof to the hypotheses.

## Overcoming Common Friction

The following table are 5 points of friction to help demonstrate how each is overcome

| #     | Friction                                       | What is demonstrated                                                                                                     | Evidence of momentum                                                                                         |
| ----- | ---------------------------------------------- |--------------------------------------------------------------------------------------------------------------------------| ------------------------------------------------------------------------------------------------------------ |
| **1** | **Recreating common application architecture** | A model + blueprint can produce a working system rather than starting from an empty repository.                          | Multiple systems generated from the same blueprint with materially less effort.                              |
| **2** | **Recreating enterprise/domain knowledge**     | Domain models and blueprints capture knowledge that would otherwise live in SMEs, documents and previous projects.       | The same knowledge is reused across multiple generated systems.                                              |
| **3** | **Inconsistent implementation of standards**   | Configuration, policies and blueprints turn architectural/engineering standards into executable production rules.        | Generated systems consistently conform to defined standards.                                                 |
| **4** | **The cost of modernization**                  | Existing application requirements can be expressed through models/configuration and regenerated onto a modern blueprint. | A credible modernization example showing reduced production effort.                                          |
| **5** | **Verification of generated systems**          | Generation isn't enough—the resulting system can be tested and verified as part of production.                           | Runtime tests, infrastructure validation and certification demonstrate that generated systems actually work. |

<!--
## Generated System Highlights

| Blueprint       | Domain     | Generated | Builds | Tests | Runs | Deploys | Repeatable |
| --------------- | ---------- | --------: | -----: | ----: | ---: | ------: | ---------: |
| Spring Boot 3.5 | Banking    |         ✓ |      ✓ |     ✓ |    ✓ |       ✓ |          ✓ |
| Spring Boot 3.5 | Healthcare |         ✓ |      ✓ |     ✓ |    ✓ |       ✓ |          ✓ |
| Go              | Banking    |         ✓ |      ✓ |     ✓ |    ✓ |       ✓ |          ✓ |
| Rails           | Retail     |         ✓ |      ✓ |     ✓ |    ✓ |       ✓ |          ✓ |
| Django          | Insurance  |         ✓ |      ✓ |     ✓ |    ✓ |       ✓ |          ✓ |
| Angular         | Healthcare |         ✓ |      ✓ |     ✓ |    ✓ |       ✓ |          ✓ |
-->

## Points of Value


#1 — Blueprint certification

#2 — Runtime verification

#3 — Enterprise system completeness

Not code snippets.

Not boilerplate.

A complete deployable system.

#4 — Productivity measurement

Demonstrate:

Traditional delivery: X people / Y months

versus

Harbormaster: X' people / Y' months.


#5 — Cloud/platform deployment

Prove:

Generates workloads that actually run on OpenShift, AWS, Azure, etc.

#6 — Industry applicability

## Proof Points
Offer real world examples where the platform and approach offered real measurable value.

### Wells Fargo Microservice Enterprise Strategy

#### Problem: Microservice enterprise adoption stalled

#### Before HM:

#### With HM:

### AxonIQ Flagship Product Migration

Problem: Version migration too difficult 

Before HM:  
AxonIQ introduces a new version (v5) of its framework and server product.  The difference between it and the previous version (v4) proved challenging for developers.  AxonIQ created AI skills to assist, but migration was still unsuccessful.

The AxonIQ team then used their v4 GitHub demo as a prompt for Claude, but Claude didn't have enough detail to create a migration path to v5.

With Harbormaster:
Harbormaster created a blueprint for the AxonIQ Framework v4, chose an industry domain model from our library, and generated a system.

Once generated, the system was validated for build, runtime, and deployment integrity.  Next, the AxonIQ team used the system's source code as a prompt for Claude, along with other prompts. With over 75,000 LOCs, the generated system was a rich representation of v4, allowing Claude to create a solution to migrate the system to v5.

## System Component Details

### [Blueprints](./blueprints/READ.md) 
These serve as an example of the blueprints available on Harbormaster.  Use them to verify their contents and observe their form.

### [Domain Models](./domain-models/READ.md)
These serve as an example of the over 200 AI procured industry domain models.  The generated systems are a result of Harbormaster compile a system against a blueprint (Spring Boot) and a selected domain model.

### [System-as-Code](./system-as-code)
The YAML that serves to declare the intent of a system.  Declare a blueprint, blueprint features, a domain model, Docker/GitHub params and custom input args to control what system results after compilations.

### [Docker](./docker/README.md)
The docker images of the systems generated as evidence.

### Certification
Offers an introduction to Harbormaster's certification system used to verify the completeness of a generated system.

### [Measurements](./measurements/READ.md)
Explains the Harbormaster Measurement Framework to demonstrate the dozens of data points that are possible when generation takes place at the system level.  

Many of the measurement categories are now possible because of the capture and reuse of software engineering knowledge within a blueprint.

