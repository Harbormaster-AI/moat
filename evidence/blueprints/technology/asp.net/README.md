# ASP.NET Blueprint

## Overview

**ASP.NET** is a Harbormaster blueprint for generating the core UI and backend capabilities of an ASP.NET application from a model-independent foundation.

The blueprint uses the ASP.NET framework with an **MVC** design pattern, web presentation technologies including jQuery, HTML5, CSS3 and `.cshtml`, and nHibernate for persistence.

## Blueprint Information

| Property               | Value             |
| ---------------------- |-------------------|
| **Name**               | ASP.NET-10        |
| **Short Name**         | `ASP.NET10`       |
| **Version**            | 1.4               |
| **Type**               | Firstclass        |
| **Application Type**   | Web Application   |
| **Release Status**     | Production        |
| **Technology Source**  | Commercial        |
| **Published**          | Yes               |
| **Docker Support**     | yes               |
| **Kubernetes Support** | Yes               |
| **Major Technology**   | ASPDotNet         |
| **Derived From**       | common, |
| **Category**           | Backend Framework |
| **Primary Vendor**     | Microsoft         |
| **Design Pattern**     | MVC               |
| **Architecture Style** | Layered           |

## Capabilities

The blueprint is designed to generate core ASP.NET application capabilities, including:

* ASP.NET web applications
* MVC-based application structure
* Controllers
* Domain and business logic
* Persistence integration
* ORM-based data access
* REST APIs
* Enterprise application foundations

## Supported Languages

| Language   | Version |
| ---------- |---------|
| C#         | 14      |
| JavaScript | —       |
| HTML5      | —       |
| `.cshtml`  | —       |

## Technology Stack

| Technology    | Layer        | Purpose                                                 |
|---------------| ------------ |---------------------------------------------------------|
| **ASPDotNet** | Business     | Model and controller                                    |
| **Database**  | Data         | ORM-like features with multiple RDBMS and NoSQL options |


## Architecture

The blueprint uses a **Layered** architecture with the **MVC** design pattern.

```text
                    ASP.NET
                       │
                       ▼
                Layered Architecture
                       │
          ┌────────────┼────────────┐
          ▼            ▼            ▼
   View (Optional)   Model       Controller
          │            │            │
          │            └─────┬──────┘
          │                  ▼
          │        Entity Framework Core
          │                  │
          │                  ▼
          │               Database
          │
          ▼
       Presentation Blueprint
          │
          ▼
       Web UI
```

## Supported Use Cases

The blueprint identifies the following use cases:

* Portal
* REST API
* SaaS Platform
* Marketplace
* Workflow Engine
* Case Management
* Financial Services
* Healthcare
* Supply Chain Management

## Metadata

```text
ASP.NET Core
Clean Architecture
Domain-Driven Design
Microservices
CQRS
Event Sourcing
Entity Framework Core
SQL Server
Azure
Docker
Kubernetes
OAuth2
Azure AD
Multi-Tenant
Production Ready
Gold Certified
```

## Harbormaster Production Model

The ASP.NET blueprint is intended to capture reusable application production knowledge for generating ASP.NET systems through Harbormaster.

```text
Domain Model
      │
      ▼
ASP.NET Blueprint
      │
      ├── MVC Structure
      ├── Models
      ├── Services
      ├── Persistence
      └── Docker
      │
      ▼
Generated ASP.NET Application
```

The model-independent nature of the blueprint allows the same production foundation to be applied across different application domains.

## Blueprint Assets

**Blueprint Icon**

`blueprints/asp-dot-net-blueprint-header-image.png`

**Architecture Preview**

`blueprints/architecture.diagrams/asp.net-blueprint-system-diagram.png`

**Information Page**

`http://www.harbormaster.net/infopages/asp.net.blueprint`

---

[<<< return](../README.md)
