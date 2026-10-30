# Harbormaster Engineering Expertise Corpus — Rough Draft v0.1

The Harbormaster Engineering Expertise Corpus is a structured body of reusable engineering knowledge that captures how experienced engineers design, build, verify, and operate software systems.

It contains:

- **Principles** — what good engineering should follow
- **Patterns** — proven ways of implementing things
- **Decisions** — choices made and why
- **Constraints** — what should or should not be done
- **Examples and counterexamples** — what works and what doesn't
- **Verification rules** — how to determine whether the engineering is correct
- **Evidence** — what has actually been demonstrated in generated systems  

The key distinction

- An expertise corpus is more than documentation.
- Documentation tells someone what a system does.
- An expertise corpus captures the knowledge required to make the right engineering decisions.
- That distinction is important to Harbormaster because:

``A blueprint is the executable form of a portion of the expertise corpus.``


## 1. Engineering Principles

These are the underlying engineering beliefs that have repeatedly guided our decisions.

### EP-001 — A Blueprint Is Executable Engineering Expertise

**Principle**

A blueprint is not a collection of code templates. It is a structured, executable representation of engineering expertise that can be applied repeatedly to create systems.

**Rationale**

The value of a blueprint comes from capturing decisions about architecture, technology, implementation, infrastructure, verification, and operational behavior—not simply generating files.

**Implication**

Blueprints should be treated as engineering IP and as first-class artifacts within Harbormaster.

---

### EP-002 — Harbormaster Is a System Compiler

**Principle**

Harbormaster should be understood as a **system compiler**, not primarily as a code generator.

**Rationale**

It accepts a declarative description of a system, applies engineering expertise encoded in blueprints, resolves dependencies and configuration, and produces a complete system.

**Implication**

The compiler remains the central execution mechanism regardless of how much AI participates in creating or selecting blueprints.

---

### EP-003 — Separate Intent From Implementation

**Principle**

The desired system should be expressed declaratively and independently of its generated implementation.

**Rationale**

System intent should survive changes in technology, framework, implementation strategy, and deployment environment.

**Pattern**

```text
System Intent
    ↓
System-as-Code
    ↓
Blueprint
    ↓
Implementation
```

---

### EP-004 — Engineering Expertise Should Be Reusable

**Principle**

Engineering knowledge should be captured once and applied repeatedly rather than recreated inside every project.

**Rationale**

Traditional software delivery repeatedly consumes architect and developer expertise. Harbormaster converts portions of that expertise into reusable production capability.

**Economic implication**

```text
People × Hours × Rate
        ↓
Expertise × Blueprint × Systems
```

---

### EP-005 — Generated Systems Should Be Complete

**Principle**

A blueprint should generate the system required to operate, deliver, and verify the application—not merely application source code.

**Scope includes**

* Application code
* Domain model
* Persistence
* APIs
* Tests
* Configuration
* Infrastructure
* Containers
* CI/CD
* Documentation
* Operational configuration

---

### EP-006 — Verification Is Part of System Generation

**Principle**

Generation without verification is incomplete.

**Rationale**

The output of a compiler must have evidence that it satisfies its expected engineering properties.

**Implication**

Verification is a first-class part of the Harbormaster lifecycle rather than an external QA activity.

---

### EP-007 — Verification Must Be Layered

**Principle**

Different classes of correctness require different verification stages.

```text
Pre-Build
    ↓
Build
    ↓
Runtime
    ↓
Delivery
    ↓
Infrastructure
```

Each stage should validate the concerns appropriate to that stage.

---

### EP-008 — Build Verification Owns Application Dependency Resolution

**Principle**

The build stage owns resolving application dependencies required to compile, test, and package the generated system.

**Rationale**

Pre-Build should validate and prepare source; it should not become a second application build system.

**Implication**

Generated source verification should avoid unnecessarily installing or resolving application dependencies before Build Verification.

---

### EP-009 — Generated Tool Artifacts Are Not Source

**Principle**

Artifacts created by development or verification tooling should not automatically be considered generated application source.

Examples include:

* `go.mod`
* `go.sum`
* `Gemfile.lock`
* temporary reports
* analyzer output
* build artifacts

**Implication**

The source boundary must be explicit.

---

### EP-010 — Formatting Can Change Source and Therefore Belongs Before the Source Commit

**Principle**

Code formatting is a source transformation.

**Pipeline**

```text
Generate
→ Format
→ Commit formatted source
→ Analyze
→ Security
→ Build
```

**Rationale**

Once formatted source is committed, subsequent verification tools may create artifacts without contaminating the source commit.

---

### EP-011 — Security Should Prefer Detection and Correction Over Suppression

**Principle**

Security findings should normally result in actual engineering remediation rather than suppressions designed only to make verification pass.

**Implication**

The verification system should expose security problems as engineering feedback.

---

### EP-012 — Deterministic Generation Matters

**Principle**

Given the same model, blueprint, configuration, and blueprint version, Harbormaster should produce a predictable result.

**Rationale**

Deterministic generation makes systems reproducible, testable, auditable, and comparable.

---

### EP-013 — Engineering Knowledge Should Be Explicit

**Principle**

Important engineering decisions should be represented explicitly rather than hidden inside templates or conversational context.

This includes:

* principles
* constraints
* architectural decisions
* patterns
* rationale
* verification rules
* exceptions
* counterexamples

This is foundational to making the substrate usable by AI.

---

### EP-014 — Verification Evidence Becomes Engineering Knowledge

**Principle**

The result of compiling and verifying systems should be retained as evidence that can improve future engineering decisions.

```text
Blueprint
→ Compile
→ Verify
→ Evidence
→ Knowledge
→ Improve Blueprint
```

---

### EP-015 — AI Maturation Is an Intent, Not a Product Feature

**Principle**

Harbormaster's AI maturation is not simply another feature roadmap. It is the long-term purpose of capturing engineering expertise in a structured, executable substrate.

**Destination**

AI progressively moves from:

```text
Assist
→ Co-Author
→ Author
→ Refine
→ Validate
→ Discover
→ Evolve / Repair
```

while Harbormaster remains the compiler.

---

### EP-016 — Framework Version Isolation

**Principle**

When generated systems succeed on a prior framework version but fail after a framework upgrade, first isolate the framework version by testing the latest maintenance release and the prior known-good minor line before modifying the generated domain model, architecture, or framework-specific implementation.

---

# 2. Architecture Decisions

These describe architectural structures and boundaries we've established or repeatedly converged upon.

### AD-001 — Harbormaster Uses a Compiler Architecture

```text
System-as-Code
      ↓
HM Compiler
      ↓
Blueprint Engine
      ↓
Generated System
      ↓
Verification / Delivery / Infrastructure
```

---

### AD-002 — System-as-Code Is the Declarative System Input

System-as-Code describes:

* blueprint selection
* domain model
* system configuration
* optional features
* technology choices
* deployment/configuration choices

The generated implementation should not be the primary system specification.

---

### AD-003 — HCIDL Is the Domain Modeling Substrate

HCIDL provides a technology-independent representation of domain concepts.

The model can express:

* entities
* attributes
* enums
* associations
* multiplicity
* value objects
* services/APIs
* events/messaging
* constraints/rules

The technology blueprint maps this model into implementation.

---

### AD-004 — Blueprint and Compiler Are Separate Concerns

The blueprint defines engineering expertise.

The compiler interprets and executes that expertise.

```text
Blueprint
    ≠
Compiler
```

This separation allows the compiler to support multiple technologies and allows blueprints to evolve independently.

---

### AD-005 — Technology Is a Blueprint Concern, Not a Domain Concern

The domain model should not need to know whether its implementation is:

* Spring Boot
* ASP.NET
* Go
* Rails
* Django
* etc.

The blueprint provides the technology realization.

---

### AD-006 — Blueprint Architecture Has Multiple Layers

The evolving taxonomy separates:

```text
Architecture Style
      ↓
Blueprint
      ↓
Architecture Patterns
      ↓
Technologies
      ↓
Implementation
```

For example, an architecture style can be **Serverless**, while a blueprint may specify a particular realization using Azure Functions and associated patterns.

This taxonomy was developed as a proposed classification and should be treated as **evolving**, not immutable.

---

### AD-007 — Verification Is an Independent Pipeline

Verification is not buried inside individual technology blueprints.

```text
Generated System
       ↓
Verification Orchestrator
       ↓
Pre-Build
Build
Runtime
Delivery
Infrastructure
```

This provides consistency across different blueprints.

---

### AD-008 — Pre-Build Is Source Correctness

Pre-Build focuses on whether the generated source is structurally and statically correct before application build resolution.

Typical responsibilities:

* files/directories
* formatting
* static analysis
* source-level security
* source commit
* report

---

### AD-009 — Runtime Verification Tests the Actual Generated System

Runtime verification should move beyond compilation.

It can include:

* container startup
* provisioning dependencies
* application startup
* health checks
* API behavior
* integration tests
* end-to-end behavior
* observability

---

### AD-010 — Delivery Verification Tests the Delivery System

Delivery Verification validates that the generated system can move through its delivery mechanism.

Examples:

* source repository
* CI/CD
* artifact management
* container image
* deployment
* deployed environment
* smoke tests

---

### AD-011 — Infrastructure Verification Is a First-Class Concern

Infrastructure should be generated and verified rather than treated as an external deployment responsibility.

Examples:

* Terraform validation
* security scanning
* provisioning
* Kubernetes/OpenShift
* Helm/manifests
* runtime infrastructure validation

---

### AD-012 — Observability Is Part of the Generated System Contract

Generated systems should have a baseline observability contract including:

* system identity
* generation identity
* blueprint identity/version
* structured logging
* correlation/trace context
* health/readiness
* metrics
* startup/shutdown behavior

OpenTelemetry is the preferred standards-based instrumentation approach.

---

### AD-013 — Instrumentation and Destination Are Separate

The generated system should define **what is instrumented**, while deployment configuration defines **where telemetry goes**.

For example:

```yaml
observability:
  instrumentation:
    tracing: true
    metrics: true
    logs: true
  export:
    protocol: otlp

deployment:
  observability:
    destination: sumologic
```

---

### AD-014 — Access Channels Should Share the Same Capabilities

Harbormaster should expose the same underlying capabilities through:

* Web UI
* CLI
* SDK/programmatic APIs

The access mechanism should not define the compiler's capabilities.

---

### AD-015 — The Current Blueprints Demonstrate the Model, Not Its Ultimate Boundary

The current focus is application/system generation, but the architecture is intentionally broader than any individual technology blueprint.

This allows the blueprint concept to expand into application, architecture, infrastructure, cloud, security, governance, operational, and domain expertise.

---

# 3. Technology Patterns

These are recurring implementation patterns we've established while building and debugging actual blueprints.

## 3.1 General Blueprint Pattern

### TP-001 — Domain-to-Technology Mapping

The canonical mapping we've been using is:

```text
HCIDL Entity
   ↓
Technology Domain Class
   ↓
Persistence Entity
   ↓
Repository
   ↓
Application Service
   ↓
API Handler
```

This is the core contract demonstrated in the ASP.NET blueprint.

---

### TP-002 — Technology Blueprint Owns Implementation Conventions

The blueprint determines conventions such as:

* file locations
* naming
* dependency injection
* persistence mapping
* API structure
* framework configuration
* tests
* infrastructure

The HCIDL model should remain technology-independent.

---

### TP-003 — Generated IDs Are Platform-Level Concerns

Harbormaster automatically introduces an identifier even when the domain model does not explicitly declare one.

The technology blueprint determines how that identifier is represented.

---

## 3.2 ASP.NET

### TP-004 — ASP.NET Layering

Current ASP.NET backend blueprint contract:

```text
HCIDL Entity
→ ASP.NET Domain Class
→ EF Core Entity
→ Repository
→ Application Service
→ Minimal API Handler
```

---

### TP-005 — EF Core Is the ORM Boundary

EF Core provides the persistence mapping between domain/application structures and the relational database.

---

### TP-006 — ASP.NET Service Registration Precedes Build

All `builder.Services.Add...` registrations must occur before:

```csharp
var app = builder.Build();
```

After `Build()`, the service collection is immutable.

This became an explicit verification concern after generated systems exposed the structural error.

---

### TP-007 — OpenTelemetry Uses Native .NET Metrics + OpenTelemetry Export

The generated application can define native meters/counters and connect them to OpenTelemetry.

Example architecture:

```text
System Metrics
    ↓
System.Diagnostics.Metrics
    ↓
OpenTelemetry
    ↓
OTLP / Prometheus / Collector
```

---

## 3.3 Go

### TP-008 — Go Favors Composition

Go blueprints should generally favor:

* composition
* interfaces
* explicit dependency injection
* adapters
* factories
* strategies

rather than inheritance-oriented designs.

---

### TP-009 — Go Layering

A recurring Go architecture is:

```text
Router
→ Controller / Handler
→ Service
→ Repository
→ Database
```

with Docker/Kubernetes as common deployment targets.

---

### TP-010 — Go Repository Pattern

Repositories provide the persistence boundary.

The service layer should not become tightly coupled to database implementation details.

---

### TP-011 — Go Concurrency Uses Native Concurrency Primitives

Goroutines and channels are the foundation for patterns such as:

* worker pools
* pipelines
* asynchronous processing

---

### TP-012 — GORM Value Objects Use Embedded Mapping

For example:

```go
SwiftBic BIC `gorm:"embedded;embeddedPrefix:swift_bic_"`
Balance Money `gorm:"embedded;embeddedPrefix:balance_"`
```

This preserves value-object structure while mapping to relational columns.

---

## 3.4 Rails

### TP-013 — Rails Naming Follows Ruby/Rails Conventions

Generated Rails artifacts use:

* classes → PascalCase
* variables → snake_case
* attributes → snake_case
* controllers/models aligned with Rails naming conventions

For example:

```text
KycProfile
@kyc_profile
kyc_profile_params
:kyc_profile
```

---

### TP-014 — Rails Static Analysis Is Independent of Application Dependency Resolution

RuboCop is installed/invoked independently in the verification stage rather than requiring `bundle exec`.

This preserves the boundary that Build Verification owns application dependencies.

---

### TP-015 — Formatting and Static Analysis Are Separate

RuboCop autocorrection can modify source, after which the remaining offenses are analyzed.

```text
Generate
→ RuboCop Auto-Correct
→ Commit
→ RuboCop Analysis
```

---

## 3.5 Frontend

### TP-016 — Frontend Technology Is Explicit

Angular, React, and Apollo GraphQL are treated as technology choices rather than loosely inferred capabilities.

```text
technology:
  ANGULAR
  REACT
  APOLLO
```

---

### TP-017 — Frontend Verification Uses Technology-Specific Structural Requirements

Examples:

**Angular**

* `angular.json`
* TypeScript configuration
* build/test tooling

**React**

* TypeScript configuration
* application configuration

**Apollo GraphQL**

* TypeScript configuration
* GraphQL/Apollo-specific implementation

Pre-Build should validate the expected source structure before build resolution.

---

## 3.6 Verification Patterns

### TP-018 — Pre-Build Verification Pattern

```text
VERIFY_FILES
→ VERIFY_DIRECTORIES
→ FORMAT_CODE
→ COMMIT_FORMATTED_SOURCE
→ STATIC_ANALYSIS
→ SECURITY_CHECK
→ BUILD
```

The final `BUILD` represents transition into Build Verification.

---

### TP-019 — Trivy Source Security Pattern

A common source-security pattern is:

```bash
trivy fs . \
  --scanners vuln \
  --severity HIGH,CRITICAL \
  --ignore-unfixed \
  --format table \
  --exit-code 1
```

The exact Trivy version/configuration should remain blueprint/configuration-specific rather than being treated as an immutable platform constant.

---

### TP-020 — Docker Delivery Pattern

The Dockerfile may exist at the repository root or application level.

The verification process should discover the Dockerfile and explicitly pass its path to the build rather than assuming a fixed location.

---

### TP-021 — Verification Must Produce Evidence

Every verification stage should produce a machine-readable or reportable result.

```text
Verification
→ Evidence
→ Harbormaster Report
```

This evidence becomes useful both operationally and as future engineering knowledge.

---

# 4. Decisions and Rationale

This category is slightly different from the first three.

It records **why we made particular decisions**, including decisions that emerged because generated systems exposed problems.

That makes this category particularly valuable for AI because it captures the reasoning behind the engineering substrate.

---

### DR-001 — Why Blueprints Are Technology-Specific

**Decision**

Blueprints encode technology-specific engineering expertise.

**Rationale**

HCIDL provides the technology-independent domain model. The blueprint is where technology-specific realization decisions belong.

This prevents the domain model from becoming polluted with framework assumptions.

---

### DR-002 — Why the Compiler Should Not Become Technology-Specific

**Decision**

The compiler orchestrates generation but should not contain large amounts of technology-specific implementation logic.

**Rationale**

Otherwise every new technology increases compiler complexity and reduces blueprint independence.

---

### DR-003 — Why Build Dependencies Belong to Build Verification

**Decision**

Do not make Pre-Build responsible for resolving application dependencies simply to perform static analysis.

**Rationale**

It creates duplicated responsibility and makes the pipeline harder to reason about.

The clean boundary is:

```text
Pre-Build
    = source correctness

Build
    = dependency resolution + compilation + tests + packaging
```

---

### DR-004 — Why Formatting Is Before the Source Commit

**Decision**

Formatting occurs before the formatted source is committed.

**Rationale**

Formatting changes the source itself. Therefore the repository should contain the canonical formatted version before later verification stages begin generating additional artifacts.

---

### DR-005 — Why Tool Artifacts Should Not Be Required Source Files

**Decision**

Do not require generated dependency/tool artifacts simply because a later verification stage may create them.

**Example**

A Rails `Gemfile.lock` does not need to be a required Pre-Build source file simply because dependency installation later creates it.

**Rationale**

Otherwise stages become unnecessarily coupled.

---

### DR-006 — Why Generated Code Should Sometimes Be Fixed Through the Blueprint

**Decision**

When a generated system repeatedly exposes the same structural defect, the preferred solution is to correct the generation rule rather than manually patching the generated system.

**Rationale**

The defect represents a failure of reusable engineering expertise.

```text
Generated Defect
→ Identify Pattern
→ Correct Blueprint
→ Regenerate
→ Verify
```

---

### DR-007 — Why Generated Code Should Not Always Be Changed to Satisfy Generic Linters

**Decision**

When generated code intentionally violates a generic style preference that is not an actual engineering defect, adjust the verification rule rather than distorting the generated architecture.

**Rationale**

The blueprint defines the intended engineering system. Generic tooling should verify it, not redefine it.

---

### DR-008 — Why Security Findings Should Improve the Blueprint

**Decision**

Repeated security findings should become candidates for blueprint-level rules.

**Rationale**

A security correction made once in a blueprint can prevent the same issue across every subsequently generated system.

---

### DR-009 — Why Observability Should Be Generated

**Decision**

Observability should be part of the generated system rather than something added manually after generation.

**Rationale**

A system that cannot expose its health, identity, telemetry, and operational state is incomplete from an operational perspective.

---

### DR-010 — Why Evidence Belongs in the Knowledge Loop

**Decision**

Verification evidence should become an input to future engineering improvement.

**Rationale**

It provides empirical information about whether encoded expertise actually works.

This creates:

```text
Expertise
→ Blueprint
→ System
→ Verification
→ Evidence
→ Improved Expertise
```

---

### DR-011 — Why AI Knowledge Should Remain External and Structured

**Decision**

The initial Harbormaster AI strategy should not depend on embedding all engineering knowledge into model weights.

**Rationale**

Blueprints, decisions, patterns, constraints, verification evidence, and other engineering knowledge need to remain:

* versioned
* inspectable
* updateable
* attributable
* technology-aware
* independently governed

AI should reason over the substrate rather than replace the substrate.

---

### DR-012 — Why Harbormaster Should Be Model-Agnostic

**Decision**

Harbormaster should use frontier AI models as reasoning engines without making the company dependent on one foundation model.

**Rationale**

The scarce strategic asset is the engineering substrate—not ownership of a foundation model.

The long-term architecture is:

```text
Foundation Models
        ↓
Harbormaster AI Layer
        ↓
Engineering Knowledge
        ↓
Blueprints
        ↓
Compiler
        ↓
Systems
```

---

### DR-013 — Why SME Corrections Are Valuable Knowledge

**Decision**

A correction made during engineering collaboration should be captured as more than a conversation outcome.

The useful structure is:

```text
Initial Proposal
→ SME Correction
→ Final Decision
→ Rationale
→ Generalized Rule
```

**Rationale**

The correction reveals the boundary between what the AI initially believed and what an expert knows.

That boundary is particularly valuable training/evaluation material for future Harbormaster AI.

---

### DR-014 — Why the Engineering Corpus Is More Valuable Than the Conversations

**Decision**

The conversations themselves should not be considered the primary knowledge asset.

The asset is the engineering knowledge extracted from them.

That knowledge includes:

* principles
* decisions
* patterns
* constraints
* rationale
* corrections
* examples
* counterexamples
* verification rules
* evidence

---

### DR-015 — Why Harbormaster's AI Strategy Starts With Existing Intelligence

**Decision**

Initially use high-capability external/private AI models rather than attempting to build a foundation model.

**Rationale**

This allows Harbormaster to spend its scarce resources building the differentiated asset:

> **the engineering substrate that gives AI the context required to reason about real software systems.**

Specialized models can later be introduced for tasks where the Harbormaster corpus provides enough proprietary training/evaluation data.

---

### DR-016 — Prefer Version Isolation Over Third-Party Debugging

**Decision**

If a failure occurs inside a third-party framework during runtime initialization, and the same generated system succeeds on a prior framework version, first test an alternative supported framework version before investigating or compensating for third-party implementation behavior.