package router

import (

    InsurerController "insurance-on-golang/internal/controller"
    InsuranceProductController "insurance-on-golang/internal/controller"
    CoverageDefinitionController "insurance-on-golang/internal/controller"
    DistributorController "insurance-on-golang/internal/controller"
    AgentController "insurance-on-golang/internal/controller"
    CustomerController "insurance-on-golang/internal/controller"
    ApplicationController "insurance-on-golang/internal/controller"
    QuoteController "insurance-on-golang/internal/controller"
    UnderwritingDecisionController "insurance-on-golang/internal/controller"
    UnderwriterController "insurance-on-golang/internal/controller"
    PolicyController "insurance-on-golang/internal/controller"
    EndorsementController "insurance-on-golang/internal/controller"
    PolicyCoverageController "insurance-on-golang/internal/controller"
    InsuredObjectController "insurance-on-golang/internal/controller"
    BeneficiaryController "insurance-on-golang/internal/controller"
    BillingAccountController "insurance-on-golang/internal/controller"
    InvoiceController "insurance-on-golang/internal/controller"
    PaymentController "insurance-on-golang/internal/controller"
    ClaimController "insurance-on-golang/internal/controller"
    IncidentController "insurance-on-golang/internal/controller"
    ExposureController "insurance-on-golang/internal/controller"
    AdjusterController "insurance-on-golang/internal/controller"
    ClaimReserveController "insurance-on-golang/internal/controller"
    ClaimPaymentController "insurance-on-golang/internal/controller"
    ServiceProviderController "insurance-on-golang/internal/controller"
    ReinsuranceAgreementController "insurance-on-golang/internal/controller"
    SubrogationRecoveryController "insurance-on-golang/internal/controller"
    ThirdPartyController "insurance-on-golang/internal/controller"
    DocumentController "insurance-on-golang/internal/controller"
    jsonResponseFormatter "insurance-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "insurance-on-golang/internal/controller"

)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    //----------------------------------------------------------------------------
    // default controllers for health and availability checking
    //----------------------------------------------------------------------------

    router.HandleFunc("/", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Default__)).Methods("GET", "OPTIONS")
    router.HandleFunc("/health", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Health__)).Methods("GET", "OPTIONS")


    //----------------------------------------------------------------------------
    // Insurer Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Insurer/{id}", jsonResponseFormatter.FormatToJSON(InsurerController.GetInsurer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Insurer", jsonResponseFormatter.FormatToJSON(InsurerController.GetAllInsurer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInsurer", jsonResponseFormatter.FormatToJSON(InsurerController.CreateInsurer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Insurer/{id}", jsonResponseFormatter.FormatToJSON(InsurerController.UpdateInsurer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInsurer/{id}", jsonResponseFormatter.FormatToJSON(InsurerController.DeleteInsurer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProductsToInsurer/{parentId}/productsId", jsonResponseFormatter.FormatToJSON(InsurerController.AddProductsToInsurer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProductsFromInsurer/{parentId}/productsIds", jsonResponseFormatter.FormatToJSON(InsurerController.RemoveProductsFromInsurer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDistributionPartnersToInsurer/{parentId}/distributionPartnersId", jsonResponseFormatter.FormatToJSON(InsurerController.AddDistributionPartnersToInsurer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDistributionPartnersFromInsurer/{parentId}/distributionPartnersIds", jsonResponseFormatter.FormatToJSON(InsurerController.RemoveDistributionPartnersFromInsurer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPoliciesToInsurer/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(InsurerController.AddPoliciesToInsurer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromInsurer/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(InsurerController.RemovePoliciesFromInsurer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddClaimsToInsurer/{parentId}/claimsId", jsonResponseFormatter.FormatToJSON(InsurerController.AddClaimsToInsurer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveClaimsFromInsurer/{parentId}/claimsIds", jsonResponseFormatter.FormatToJSON(InsurerController.RemoveClaimsFromInsurer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReinsuranceAgreementsToInsurer/{parentId}/reinsuranceAgreementsId", jsonResponseFormatter.FormatToJSON(InsurerController.AddReinsuranceAgreementsToInsurer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReinsuranceAgreementsFromInsurer/{parentId}/reinsuranceAgreementsIds", jsonResponseFormatter.FormatToJSON(InsurerController.RemoveReinsuranceAgreementsFromInsurer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InsuranceProduct Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InsuranceProduct/{id}", jsonResponseFormatter.FormatToJSON(InsuranceProductController.GetInsuranceProduct)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InsuranceProduct", jsonResponseFormatter.FormatToJSON(InsuranceProductController.GetAllInsuranceProduct)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInsuranceProduct", jsonResponseFormatter.FormatToJSON(InsuranceProductController.CreateInsuranceProduct)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InsuranceProduct/{id}", jsonResponseFormatter.FormatToJSON(InsuranceProductController.UpdateInsuranceProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInsuranceProduct/{id}", jsonResponseFormatter.FormatToJSON(InsuranceProductController.DeleteInsuranceProduct)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInsurerToInsuranceProduct/{parentId}/insurerId", jsonResponseFormatter.FormatToJSON(InsuranceProductController.AssignInsurerToInsuranceProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInsurerFromInsuranceProduct/{parentId}", jsonResponseFormatter.FormatToJSON(InsuranceProductController.UnassignInsurerFromInsuranceProduct)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCoverageDefinitionsToInsuranceProduct/{parentId}/coverageDefinitionsId", jsonResponseFormatter.FormatToJSON(InsuranceProductController.AddCoverageDefinitionsToInsuranceProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCoverageDefinitionsFromInsuranceProduct/{parentId}/coverageDefinitionsIds", jsonResponseFormatter.FormatToJSON(InsuranceProductController.RemoveCoverageDefinitionsFromInsuranceProduct)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CoverageDefinition Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CoverageDefinition/{id}", jsonResponseFormatter.FormatToJSON(CoverageDefinitionController.GetCoverageDefinition)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CoverageDefinition", jsonResponseFormatter.FormatToJSON(CoverageDefinitionController.GetAllCoverageDefinition)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCoverageDefinition", jsonResponseFormatter.FormatToJSON(CoverageDefinitionController.CreateCoverageDefinition)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CoverageDefinition/{id}", jsonResponseFormatter.FormatToJSON(CoverageDefinitionController.UpdateCoverageDefinition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCoverageDefinition/{id}", jsonResponseFormatter.FormatToJSON(CoverageDefinitionController.DeleteCoverageDefinition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignProductToCoverageDefinition/{parentId}/productId", jsonResponseFormatter.FormatToJSON(CoverageDefinitionController.AssignProductToCoverageDefinition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductFromCoverageDefinition/{parentId}", jsonResponseFormatter.FormatToJSON(CoverageDefinitionController.UnassignProductFromCoverageDefinition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Distributor Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Distributor/{id}", jsonResponseFormatter.FormatToJSON(DistributorController.GetDistributor)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Distributor", jsonResponseFormatter.FormatToJSON(DistributorController.GetAllDistributor)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDistributor", jsonResponseFormatter.FormatToJSON(DistributorController.CreateDistributor)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Distributor/{id}", jsonResponseFormatter.FormatToJSON(DistributorController.UpdateDistributor)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDistributor/{id}", jsonResponseFormatter.FormatToJSON(DistributorController.DeleteDistributor)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInsurersToDistributor/{parentId}/insurersId", jsonResponseFormatter.FormatToJSON(DistributorController.AddInsurersToDistributor)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInsurersFromDistributor/{parentId}/insurersIds", jsonResponseFormatter.FormatToJSON(DistributorController.RemoveInsurersFromDistributor)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAgentsToDistributor/{parentId}/agentsId", jsonResponseFormatter.FormatToJSON(DistributorController.AddAgentsToDistributor)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAgentsFromDistributor/{parentId}/agentsIds", jsonResponseFormatter.FormatToJSON(DistributorController.RemoveAgentsFromDistributor)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPoliciesToDistributor/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(DistributorController.AddPoliciesToDistributor)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromDistributor/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(DistributorController.RemovePoliciesFromDistributor)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Agent Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Agent/{id}", jsonResponseFormatter.FormatToJSON(AgentController.GetAgent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Agent", jsonResponseFormatter.FormatToJSON(AgentController.GetAllAgent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAgent", jsonResponseFormatter.FormatToJSON(AgentController.CreateAgent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Agent/{id}", jsonResponseFormatter.FormatToJSON(AgentController.UpdateAgent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAgent/{id}", jsonResponseFormatter.FormatToJSON(AgentController.DeleteAgent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignDistributorToAgent/{parentId}/distributorId", jsonResponseFormatter.FormatToJSON(AgentController.AssignDistributorToAgent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDistributorFromAgent/{parentId}", jsonResponseFormatter.FormatToJSON(AgentController.UnassignDistributorFromAgent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPoliciesToAgent/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(AgentController.AddPoliciesToAgent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromAgent/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(AgentController.RemovePoliciesFromAgent)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCustomersToAgent/{parentId}/customersId", jsonResponseFormatter.FormatToJSON(AgentController.AddCustomersToAgent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCustomersFromAgent/{parentId}/customersIds", jsonResponseFormatter.FormatToJSON(AgentController.RemoveCustomersFromAgent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Customer Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Customer/{id}", jsonResponseFormatter.FormatToJSON(CustomerController.GetCustomer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Customer", jsonResponseFormatter.FormatToJSON(CustomerController.GetAllCustomer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCustomer", jsonResponseFormatter.FormatToJSON(CustomerController.CreateCustomer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Customer/{id}", jsonResponseFormatter.FormatToJSON(CustomerController.UpdateCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCustomer/{id}", jsonResponseFormatter.FormatToJSON(CustomerController.DeleteCustomer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddApplicationsToCustomer/{parentId}/applicationsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddApplicationsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveApplicationsFromCustomer/{parentId}/applicationsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveApplicationsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPoliciesToCustomer/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(CustomerController.AddPoliciesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromCustomer/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemovePoliciesFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddClaimsToCustomer/{parentId}/claimsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddClaimsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveClaimsFromCustomer/{parentId}/claimsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveClaimsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAgentsToCustomer/{parentId}/agentsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddAgentsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAgentsFromCustomer/{parentId}/agentsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveAgentsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddBeneficiariesToCustomer/{parentId}/beneficiariesId", jsonResponseFormatter.FormatToJSON(CustomerController.AddBeneficiariesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBeneficiariesFromCustomer/{parentId}/beneficiariesIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveBeneficiariesFromCustomer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Application Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Application/{id}", jsonResponseFormatter.FormatToJSON(ApplicationController.GetApplication)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Application", jsonResponseFormatter.FormatToJSON(ApplicationController.GetAllApplication)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewApplication", jsonResponseFormatter.FormatToJSON(ApplicationController.CreateApplication)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Application/{id}", jsonResponseFormatter.FormatToJSON(ApplicationController.UpdateApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteApplication/{id}", jsonResponseFormatter.FormatToJSON(ApplicationController.DeleteApplication)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToApplication/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(ApplicationController.AssignCustomerToApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromApplication/{parentId}", jsonResponseFormatter.FormatToJSON(ApplicationController.UnassignCustomerFromApplication)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProductToApplication/{parentId}/productId", jsonResponseFormatter.FormatToJSON(ApplicationController.AssignProductToApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductFromApplication/{parentId}", jsonResponseFormatter.FormatToJSON(ApplicationController.UnassignProductFromApplication)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDistributorToApplication/{parentId}/distributorId", jsonResponseFormatter.FormatToJSON(ApplicationController.AssignDistributorToApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDistributorFromApplication/{parentId}", jsonResponseFormatter.FormatToJSON(ApplicationController.UnassignDistributorFromApplication)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSelectedQuoteToApplication/{parentId}/selectedQuoteId", jsonResponseFormatter.FormatToJSON(ApplicationController.AssignSelectedQuoteToApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSelectedQuoteFromApplication/{parentId}", jsonResponseFormatter.FormatToJSON(ApplicationController.UnassignSelectedQuoteFromApplication)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddQuotesToApplication/{parentId}/quotesId", jsonResponseFormatter.FormatToJSON(ApplicationController.AddQuotesToApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQuotesFromApplication/{parentId}/quotesIds", jsonResponseFormatter.FormatToJSON(ApplicationController.RemoveQuotesFromApplication)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Quote Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Quote/{id}", jsonResponseFormatter.FormatToJSON(QuoteController.GetQuote)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Quote", jsonResponseFormatter.FormatToJSON(QuoteController.GetAllQuote)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewQuote", jsonResponseFormatter.FormatToJSON(QuoteController.CreateQuote)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Quote/{id}", jsonResponseFormatter.FormatToJSON(QuoteController.UpdateQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteQuote/{id}", jsonResponseFormatter.FormatToJSON(QuoteController.DeleteQuote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignApplicationToQuote/{parentId}/applicationId", jsonResponseFormatter.FormatToJSON(QuoteController.AssignApplicationToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignApplicationFromQuote/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteController.UnassignApplicationFromQuote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPolicyToQuote/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(QuoteController.AssignPolicyToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromQuote/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteController.UnassignPolicyFromQuote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddUnderwritingDecisionsToQuote/{parentId}/underwritingDecisionsId", jsonResponseFormatter.FormatToJSON(QuoteController.AddUnderwritingDecisionsToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveUnderwritingDecisionsFromQuote/{parentId}/underwritingDecisionsIds", jsonResponseFormatter.FormatToJSON(QuoteController.RemoveUnderwritingDecisionsFromQuote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // UnderwritingDecision Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/UnderwritingDecision/{id}", jsonResponseFormatter.FormatToJSON(UnderwritingDecisionController.GetUnderwritingDecision)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/UnderwritingDecision", jsonResponseFormatter.FormatToJSON(UnderwritingDecisionController.GetAllUnderwritingDecision)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewUnderwritingDecision", jsonResponseFormatter.FormatToJSON(UnderwritingDecisionController.CreateUnderwritingDecision)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/UnderwritingDecision/{id}", jsonResponseFormatter.FormatToJSON(UnderwritingDecisionController.UpdateUnderwritingDecision)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteUnderwritingDecision/{id}", jsonResponseFormatter.FormatToJSON(UnderwritingDecisionController.DeleteUnderwritingDecision)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignQuoteToUnderwritingDecision/{parentId}/quoteId", jsonResponseFormatter.FormatToJSON(UnderwritingDecisionController.AssignQuoteToUnderwritingDecision)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignQuoteFromUnderwritingDecision/{parentId}", jsonResponseFormatter.FormatToJSON(UnderwritingDecisionController.UnassignQuoteFromUnderwritingDecision)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignUnderwriterToUnderwritingDecision/{parentId}/underwriterId", jsonResponseFormatter.FormatToJSON(UnderwritingDecisionController.AssignUnderwriterToUnderwritingDecision)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignUnderwriterFromUnderwritingDecision/{parentId}", jsonResponseFormatter.FormatToJSON(UnderwritingDecisionController.UnassignUnderwriterFromUnderwritingDecision)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Underwriter Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Underwriter/{id}", jsonResponseFormatter.FormatToJSON(UnderwriterController.GetUnderwriter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Underwriter", jsonResponseFormatter.FormatToJSON(UnderwriterController.GetAllUnderwriter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewUnderwriter", jsonResponseFormatter.FormatToJSON(UnderwriterController.CreateUnderwriter)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Underwriter/{id}", jsonResponseFormatter.FormatToJSON(UnderwriterController.UpdateUnderwriter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteUnderwriter/{id}", jsonResponseFormatter.FormatToJSON(UnderwriterController.DeleteUnderwriter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInsurerToUnderwriter/{parentId}/insurerId", jsonResponseFormatter.FormatToJSON(UnderwriterController.AssignInsurerToUnderwriter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInsurerFromUnderwriter/{parentId}", jsonResponseFormatter.FormatToJSON(UnderwriterController.UnassignInsurerFromUnderwriter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDecisionsToUnderwriter/{parentId}/decisionsId", jsonResponseFormatter.FormatToJSON(UnderwriterController.AddDecisionsToUnderwriter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDecisionsFromUnderwriter/{parentId}/decisionsIds", jsonResponseFormatter.FormatToJSON(UnderwriterController.RemoveDecisionsFromUnderwriter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Policy Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Policy/{id}", jsonResponseFormatter.FormatToJSON(PolicyController.GetPolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Policy", jsonResponseFormatter.FormatToJSON(PolicyController.GetAllPolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPolicy", jsonResponseFormatter.FormatToJSON(PolicyController.CreatePolicy)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Policy/{id}", jsonResponseFormatter.FormatToJSON(PolicyController.UpdatePolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePolicy/{id}", jsonResponseFormatter.FormatToJSON(PolicyController.DeletePolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInsurerToPolicy/{parentId}/insurerId", jsonResponseFormatter.FormatToJSON(PolicyController.AssignInsurerToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInsurerFromPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(PolicyController.UnassignInsurerFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCustomerToPolicy/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(PolicyController.AssignCustomerToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(PolicyController.UnassignCustomerFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProductToPolicy/{parentId}/productId", jsonResponseFormatter.FormatToJSON(PolicyController.AssignProductToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductFromPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(PolicyController.UnassignProductFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAgentToPolicy/{parentId}/agentId", jsonResponseFormatter.FormatToJSON(PolicyController.AssignAgentToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAgentFromPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(PolicyController.UnassignAgentFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBillingAccountToPolicy/{parentId}/billingAccountId", jsonResponseFormatter.FormatToJSON(PolicyController.AssignBillingAccountToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBillingAccountFromPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(PolicyController.UnassignBillingAccountFromPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCoveragesToPolicy/{parentId}/coveragesId", jsonResponseFormatter.FormatToJSON(PolicyController.AddCoveragesToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCoveragesFromPolicy/{parentId}/coveragesIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveCoveragesFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInsuredObjectsToPolicy/{parentId}/insuredObjectsId", jsonResponseFormatter.FormatToJSON(PolicyController.AddInsuredObjectsToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInsuredObjectsFromPolicy/{parentId}/insuredObjectsIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveInsuredObjectsFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEndorsementsToPolicy/{parentId}/endorsementsId", jsonResponseFormatter.FormatToJSON(PolicyController.AddEndorsementsToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEndorsementsFromPolicy/{parentId}/endorsementsIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveEndorsementsFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddBeneficiariesToPolicy/{parentId}/beneficiariesId", jsonResponseFormatter.FormatToJSON(PolicyController.AddBeneficiariesToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBeneficiariesFromPolicy/{parentId}/beneficiariesIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveBeneficiariesFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddClaimsToPolicy/{parentId}/claimsId", jsonResponseFormatter.FormatToJSON(PolicyController.AddClaimsToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveClaimsFromPolicy/{parentId}/claimsIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveClaimsFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReinsuranceAgreementsToPolicy/{parentId}/reinsuranceAgreementsId", jsonResponseFormatter.FormatToJSON(PolicyController.AddReinsuranceAgreementsToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReinsuranceAgreementsFromPolicy/{parentId}/reinsuranceAgreementsIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveReinsuranceAgreementsFromPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Endorsement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Endorsement/{id}", jsonResponseFormatter.FormatToJSON(EndorsementController.GetEndorsement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Endorsement", jsonResponseFormatter.FormatToJSON(EndorsementController.GetAllEndorsement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEndorsement", jsonResponseFormatter.FormatToJSON(EndorsementController.CreateEndorsement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Endorsement/{id}", jsonResponseFormatter.FormatToJSON(EndorsementController.UpdateEndorsement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEndorsement/{id}", jsonResponseFormatter.FormatToJSON(EndorsementController.DeleteEndorsement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPolicyToEndorsement/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(EndorsementController.AssignPolicyToEndorsement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromEndorsement/{parentId}", jsonResponseFormatter.FormatToJSON(EndorsementController.UnassignPolicyFromEndorsement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // PolicyCoverage Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PolicyCoverage/{id}", jsonResponseFormatter.FormatToJSON(PolicyCoverageController.GetPolicyCoverage)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PolicyCoverage", jsonResponseFormatter.FormatToJSON(PolicyCoverageController.GetAllPolicyCoverage)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPolicyCoverage", jsonResponseFormatter.FormatToJSON(PolicyCoverageController.CreatePolicyCoverage)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PolicyCoverage/{id}", jsonResponseFormatter.FormatToJSON(PolicyCoverageController.UpdatePolicyCoverage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePolicyCoverage/{id}", jsonResponseFormatter.FormatToJSON(PolicyCoverageController.DeletePolicyCoverage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPolicyToPolicyCoverage/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(PolicyCoverageController.AssignPolicyToPolicyCoverage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromPolicyCoverage/{parentId}", jsonResponseFormatter.FormatToJSON(PolicyCoverageController.UnassignPolicyFromPolicyCoverage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInsuredObjectsToPolicyCoverage/{parentId}/insuredObjectsId", jsonResponseFormatter.FormatToJSON(PolicyCoverageController.AddInsuredObjectsToPolicyCoverage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInsuredObjectsFromPolicyCoverage/{parentId}/insuredObjectsIds", jsonResponseFormatter.FormatToJSON(PolicyCoverageController.RemoveInsuredObjectsFromPolicyCoverage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InsuredObject Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InsuredObject/{id}", jsonResponseFormatter.FormatToJSON(InsuredObjectController.GetInsuredObject)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InsuredObject", jsonResponseFormatter.FormatToJSON(InsuredObjectController.GetAllInsuredObject)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInsuredObject", jsonResponseFormatter.FormatToJSON(InsuredObjectController.CreateInsuredObject)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InsuredObject/{id}", jsonResponseFormatter.FormatToJSON(InsuredObjectController.UpdateInsuredObject)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInsuredObject/{id}", jsonResponseFormatter.FormatToJSON(InsuredObjectController.DeleteInsuredObject)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPolicyToInsuredObject/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(InsuredObjectController.AssignPolicyToInsuredObject)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromInsuredObject/{parentId}", jsonResponseFormatter.FormatToJSON(InsuredObjectController.UnassignPolicyFromInsuredObject)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCoveragesToInsuredObject/{parentId}/coveragesId", jsonResponseFormatter.FormatToJSON(InsuredObjectController.AddCoveragesToInsuredObject)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCoveragesFromInsuredObject/{parentId}/coveragesIds", jsonResponseFormatter.FormatToJSON(InsuredObjectController.RemoveCoveragesFromInsuredObject)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Beneficiary Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Beneficiary/{id}", jsonResponseFormatter.FormatToJSON(BeneficiaryController.GetBeneficiary)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Beneficiary", jsonResponseFormatter.FormatToJSON(BeneficiaryController.GetAllBeneficiary)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBeneficiary", jsonResponseFormatter.FormatToJSON(BeneficiaryController.CreateBeneficiary)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Beneficiary/{id}", jsonResponseFormatter.FormatToJSON(BeneficiaryController.UpdateBeneficiary)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBeneficiary/{id}", jsonResponseFormatter.FormatToJSON(BeneficiaryController.DeleteBeneficiary)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPolicyToBeneficiary/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(BeneficiaryController.AssignPolicyToBeneficiary)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromBeneficiary/{parentId}", jsonResponseFormatter.FormatToJSON(BeneficiaryController.UnassignPolicyFromBeneficiary)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCustomerToBeneficiary/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(BeneficiaryController.AssignCustomerToBeneficiary)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromBeneficiary/{parentId}", jsonResponseFormatter.FormatToJSON(BeneficiaryController.UnassignCustomerFromBeneficiary)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // BillingAccount Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BillingAccount/{id}", jsonResponseFormatter.FormatToJSON(BillingAccountController.GetBillingAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BillingAccount", jsonResponseFormatter.FormatToJSON(BillingAccountController.GetAllBillingAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBillingAccount", jsonResponseFormatter.FormatToJSON(BillingAccountController.CreateBillingAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BillingAccount/{id}", jsonResponseFormatter.FormatToJSON(BillingAccountController.UpdateBillingAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBillingAccount/{id}", jsonResponseFormatter.FormatToJSON(BillingAccountController.DeleteBillingAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToBillingAccount/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(BillingAccountController.AssignCustomerToBillingAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromBillingAccount/{parentId}", jsonResponseFormatter.FormatToJSON(BillingAccountController.UnassignCustomerFromBillingAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPoliciesToBillingAccount/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(BillingAccountController.AddPoliciesToBillingAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromBillingAccount/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(BillingAccountController.RemovePoliciesFromBillingAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInvoicesToBillingAccount/{parentId}/invoicesId", jsonResponseFormatter.FormatToJSON(BillingAccountController.AddInvoicesToBillingAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInvoicesFromBillingAccount/{parentId}/invoicesIds", jsonResponseFormatter.FormatToJSON(BillingAccountController.RemoveInvoicesFromBillingAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPaymentsToBillingAccount/{parentId}/paymentsId", jsonResponseFormatter.FormatToJSON(BillingAccountController.AddPaymentsToBillingAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentsFromBillingAccount/{parentId}/paymentsIds", jsonResponseFormatter.FormatToJSON(BillingAccountController.RemovePaymentsFromBillingAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Invoice Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Invoice/{id}", jsonResponseFormatter.FormatToJSON(InvoiceController.GetInvoice)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Invoice", jsonResponseFormatter.FormatToJSON(InvoiceController.GetAllInvoice)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInvoice", jsonResponseFormatter.FormatToJSON(InvoiceController.CreateInvoice)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Invoice/{id}", jsonResponseFormatter.FormatToJSON(InvoiceController.UpdateInvoice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInvoice/{id}", jsonResponseFormatter.FormatToJSON(InvoiceController.DeleteInvoice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBillingAccountToInvoice/{parentId}/billingAccountId", jsonResponseFormatter.FormatToJSON(InvoiceController.AssignBillingAccountToInvoice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBillingAccountFromInvoice/{parentId}", jsonResponseFormatter.FormatToJSON(InvoiceController.UnassignBillingAccountFromInvoice)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPolicyToInvoice/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(InvoiceController.AssignPolicyToInvoice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromInvoice/{parentId}", jsonResponseFormatter.FormatToJSON(InvoiceController.UnassignPolicyFromInvoice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPaymentsToInvoice/{parentId}/paymentsId", jsonResponseFormatter.FormatToJSON(InvoiceController.AddPaymentsToInvoice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentsFromInvoice/{parentId}/paymentsIds", jsonResponseFormatter.FormatToJSON(InvoiceController.RemovePaymentsFromInvoice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Payment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Payment/{id}", jsonResponseFormatter.FormatToJSON(PaymentController.GetPayment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Payment", jsonResponseFormatter.FormatToJSON(PaymentController.GetAllPayment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPayment", jsonResponseFormatter.FormatToJSON(PaymentController.CreatePayment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Payment/{id}", jsonResponseFormatter.FormatToJSON(PaymentController.UpdatePayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePayment/{id}", jsonResponseFormatter.FormatToJSON(PaymentController.DeletePayment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInvoiceToPayment/{parentId}/invoiceId", jsonResponseFormatter.FormatToJSON(PaymentController.AssignInvoiceToPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInvoiceFromPayment/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentController.UnassignInvoiceFromPayment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBillingAccountToPayment/{parentId}/billingAccountId", jsonResponseFormatter.FormatToJSON(PaymentController.AssignBillingAccountToPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBillingAccountFromPayment/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentController.UnassignBillingAccountFromPayment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPolicyToPayment/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(PaymentController.AssignPolicyToPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromPayment/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentController.UnassignPolicyFromPayment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Claim Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Claim/{id}", jsonResponseFormatter.FormatToJSON(ClaimController.GetClaim)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Claim", jsonResponseFormatter.FormatToJSON(ClaimController.GetAllClaim)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewClaim", jsonResponseFormatter.FormatToJSON(ClaimController.CreateClaim)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Claim/{id}", jsonResponseFormatter.FormatToJSON(ClaimController.UpdateClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteClaim/{id}", jsonResponseFormatter.FormatToJSON(ClaimController.DeleteClaim)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPolicyToClaim/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(ClaimController.AssignPolicyToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromClaim/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimController.UnassignPolicyFromClaim)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCustomerToClaim/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(ClaimController.AssignCustomerToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromClaim/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimController.UnassignCustomerFromClaim)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAdjusterToClaim/{parentId}/adjusterId", jsonResponseFormatter.FormatToJSON(ClaimController.AssignAdjusterToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdjusterFromClaim/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimController.UnassignAdjusterFromClaim)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignIncidentToClaim/{parentId}/incidentId", jsonResponseFormatter.FormatToJSON(ClaimController.AssignIncidentToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignIncidentFromClaim/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimController.UnassignIncidentFromClaim)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddExposuresToClaim/{parentId}/exposuresId", jsonResponseFormatter.FormatToJSON(ClaimController.AddExposuresToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveExposuresFromClaim/{parentId}/exposuresIds", jsonResponseFormatter.FormatToJSON(ClaimController.RemoveExposuresFromClaim)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReservesToClaim/{parentId}/reservesId", jsonResponseFormatter.FormatToJSON(ClaimController.AddReservesToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReservesFromClaim/{parentId}/reservesIds", jsonResponseFormatter.FormatToJSON(ClaimController.RemoveReservesFromClaim)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddClaimPaymentsToClaim/{parentId}/claimPaymentsId", jsonResponseFormatter.FormatToJSON(ClaimController.AddClaimPaymentsToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveClaimPaymentsFromClaim/{parentId}/claimPaymentsIds", jsonResponseFormatter.FormatToJSON(ClaimController.RemoveClaimPaymentsFromClaim)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddServiceProvidersToClaim/{parentId}/serviceProvidersId", jsonResponseFormatter.FormatToJSON(ClaimController.AddServiceProvidersToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveServiceProvidersFromClaim/{parentId}/serviceProvidersIds", jsonResponseFormatter.FormatToJSON(ClaimController.RemoveServiceProvidersFromClaim)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSubrogationsToClaim/{parentId}/subrogationsId", jsonResponseFormatter.FormatToJSON(ClaimController.AddSubrogationsToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSubrogationsFromClaim/{parentId}/subrogationsIds", jsonResponseFormatter.FormatToJSON(ClaimController.RemoveSubrogationsFromClaim)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Incident Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Incident/{id}", jsonResponseFormatter.FormatToJSON(IncidentController.GetIncident)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Incident", jsonResponseFormatter.FormatToJSON(IncidentController.GetAllIncident)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewIncident", jsonResponseFormatter.FormatToJSON(IncidentController.CreateIncident)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Incident/{id}", jsonResponseFormatter.FormatToJSON(IncidentController.UpdateIncident)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteIncident/{id}", jsonResponseFormatter.FormatToJSON(IncidentController.DeleteIncident)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignClaimToIncident/{parentId}/claimId", jsonResponseFormatter.FormatToJSON(IncidentController.AssignClaimToIncident)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignClaimFromIncident/{parentId}", jsonResponseFormatter.FormatToJSON(IncidentController.UnassignClaimFromIncident)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInsuredObjectsToIncident/{parentId}/insuredObjectsId", jsonResponseFormatter.FormatToJSON(IncidentController.AddInsuredObjectsToIncident)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInsuredObjectsFromIncident/{parentId}/insuredObjectsIds", jsonResponseFormatter.FormatToJSON(IncidentController.RemoveInsuredObjectsFromIncident)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Exposure Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Exposure/{id}", jsonResponseFormatter.FormatToJSON(ExposureController.GetExposure)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Exposure", jsonResponseFormatter.FormatToJSON(ExposureController.GetAllExposure)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewExposure", jsonResponseFormatter.FormatToJSON(ExposureController.CreateExposure)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Exposure/{id}", jsonResponseFormatter.FormatToJSON(ExposureController.UpdateExposure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteExposure/{id}", jsonResponseFormatter.FormatToJSON(ExposureController.DeleteExposure)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignClaimToExposure/{parentId}/claimId", jsonResponseFormatter.FormatToJSON(ExposureController.AssignClaimToExposure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignClaimFromExposure/{parentId}", jsonResponseFormatter.FormatToJSON(ExposureController.UnassignClaimFromExposure)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPolicyCoverageToExposure/{parentId}/policyCoverageId", jsonResponseFormatter.FormatToJSON(ExposureController.AssignPolicyCoverageToExposure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyCoverageFromExposure/{parentId}", jsonResponseFormatter.FormatToJSON(ExposureController.UnassignPolicyCoverageFromExposure)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInsuredObjectToExposure/{parentId}/insuredObjectId", jsonResponseFormatter.FormatToJSON(ExposureController.AssignInsuredObjectToExposure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInsuredObjectFromExposure/{parentId}", jsonResponseFormatter.FormatToJSON(ExposureController.UnassignInsuredObjectFromExposure)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddReservesToExposure/{parentId}/reservesId", jsonResponseFormatter.FormatToJSON(ExposureController.AddReservesToExposure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReservesFromExposure/{parentId}/reservesIds", jsonResponseFormatter.FormatToJSON(ExposureController.RemoveReservesFromExposure)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPaymentsToExposure/{parentId}/paymentsId", jsonResponseFormatter.FormatToJSON(ExposureController.AddPaymentsToExposure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentsFromExposure/{parentId}/paymentsIds", jsonResponseFormatter.FormatToJSON(ExposureController.RemovePaymentsFromExposure)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Adjuster Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Adjuster/{id}", jsonResponseFormatter.FormatToJSON(AdjusterController.GetAdjuster)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Adjuster", jsonResponseFormatter.FormatToJSON(AdjusterController.GetAllAdjuster)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAdjuster", jsonResponseFormatter.FormatToJSON(AdjusterController.CreateAdjuster)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Adjuster/{id}", jsonResponseFormatter.FormatToJSON(AdjusterController.UpdateAdjuster)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAdjuster/{id}", jsonResponseFormatter.FormatToJSON(AdjusterController.DeleteAdjuster)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddClaimsToAdjuster/{parentId}/claimsId", jsonResponseFormatter.FormatToJSON(AdjusterController.AddClaimsToAdjuster)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveClaimsFromAdjuster/{parentId}/claimsIds", jsonResponseFormatter.FormatToJSON(AdjusterController.RemoveClaimsFromAdjuster)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddServiceProvidersToAdjuster/{parentId}/serviceProvidersId", jsonResponseFormatter.FormatToJSON(AdjusterController.AddServiceProvidersToAdjuster)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveServiceProvidersFromAdjuster/{parentId}/serviceProvidersIds", jsonResponseFormatter.FormatToJSON(AdjusterController.RemoveServiceProvidersFromAdjuster)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ClaimReserve Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ClaimReserve/{id}", jsonResponseFormatter.FormatToJSON(ClaimReserveController.GetClaimReserve)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ClaimReserve", jsonResponseFormatter.FormatToJSON(ClaimReserveController.GetAllClaimReserve)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewClaimReserve", jsonResponseFormatter.FormatToJSON(ClaimReserveController.CreateClaimReserve)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ClaimReserve/{id}", jsonResponseFormatter.FormatToJSON(ClaimReserveController.UpdateClaimReserve)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteClaimReserve/{id}", jsonResponseFormatter.FormatToJSON(ClaimReserveController.DeleteClaimReserve)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignClaimToClaimReserve/{parentId}/claimId", jsonResponseFormatter.FormatToJSON(ClaimReserveController.AssignClaimToClaimReserve)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignClaimFromClaimReserve/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimReserveController.UnassignClaimFromClaimReserve)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignExposureToClaimReserve/{parentId}/exposureId", jsonResponseFormatter.FormatToJSON(ClaimReserveController.AssignExposureToClaimReserve)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignExposureFromClaimReserve/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimReserveController.UnassignExposureFromClaimReserve)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ClaimPayment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ClaimPayment/{id}", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.GetClaimPayment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ClaimPayment", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.GetAllClaimPayment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewClaimPayment", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.CreateClaimPayment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ClaimPayment/{id}", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.UpdateClaimPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteClaimPayment/{id}", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.DeleteClaimPayment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignClaimToClaimPayment/{parentId}/claimId", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.AssignClaimToClaimPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignClaimFromClaimPayment/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.UnassignClaimFromClaimPayment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignExposureToClaimPayment/{parentId}/exposureId", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.AssignExposureToClaimPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignExposureFromClaimPayment/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.UnassignExposureFromClaimPayment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBeneficiaryToClaimPayment/{parentId}/beneficiaryId", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.AssignBeneficiaryToClaimPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBeneficiaryFromClaimPayment/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.UnassignBeneficiaryFromClaimPayment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignServiceProviderToClaimPayment/{parentId}/serviceProviderId", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.AssignServiceProviderToClaimPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignServiceProviderFromClaimPayment/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.UnassignServiceProviderFromClaimPayment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCustomerToClaimPayment/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.AssignCustomerToClaimPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromClaimPayment/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimPaymentController.UnassignCustomerFromClaimPayment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ServiceProvider Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ServiceProvider/{id}", jsonResponseFormatter.FormatToJSON(ServiceProviderController.GetServiceProvider)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ServiceProvider", jsonResponseFormatter.FormatToJSON(ServiceProviderController.GetAllServiceProvider)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewServiceProvider", jsonResponseFormatter.FormatToJSON(ServiceProviderController.CreateServiceProvider)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ServiceProvider/{id}", jsonResponseFormatter.FormatToJSON(ServiceProviderController.UpdateServiceProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteServiceProvider/{id}", jsonResponseFormatter.FormatToJSON(ServiceProviderController.DeleteServiceProvider)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddClaimsToServiceProvider/{parentId}/claimsId", jsonResponseFormatter.FormatToJSON(ServiceProviderController.AddClaimsToServiceProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveClaimsFromServiceProvider/{parentId}/claimsIds", jsonResponseFormatter.FormatToJSON(ServiceProviderController.RemoveClaimsFromServiceProvider)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ReinsuranceAgreement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ReinsuranceAgreement/{id}", jsonResponseFormatter.FormatToJSON(ReinsuranceAgreementController.GetReinsuranceAgreement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ReinsuranceAgreement", jsonResponseFormatter.FormatToJSON(ReinsuranceAgreementController.GetAllReinsuranceAgreement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewReinsuranceAgreement", jsonResponseFormatter.FormatToJSON(ReinsuranceAgreementController.CreateReinsuranceAgreement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ReinsuranceAgreement/{id}", jsonResponseFormatter.FormatToJSON(ReinsuranceAgreementController.UpdateReinsuranceAgreement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteReinsuranceAgreement/{id}", jsonResponseFormatter.FormatToJSON(ReinsuranceAgreementController.DeleteReinsuranceAgreement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInsurerToReinsuranceAgreement/{parentId}/insurerId", jsonResponseFormatter.FormatToJSON(ReinsuranceAgreementController.AssignInsurerToReinsuranceAgreement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInsurerFromReinsuranceAgreement/{parentId}", jsonResponseFormatter.FormatToJSON(ReinsuranceAgreementController.UnassignInsurerFromReinsuranceAgreement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPoliciesToReinsuranceAgreement/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(ReinsuranceAgreementController.AddPoliciesToReinsuranceAgreement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromReinsuranceAgreement/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(ReinsuranceAgreementController.RemovePoliciesFromReinsuranceAgreement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // SubrogationRecovery Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SubrogationRecovery/{id}", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.GetSubrogationRecovery)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SubrogationRecovery", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.GetAllSubrogationRecovery)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSubrogationRecovery", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.CreateSubrogationRecovery)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SubrogationRecovery/{id}", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.UpdateSubrogationRecovery)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSubrogationRecovery/{id}", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.DeleteSubrogationRecovery)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignClaimToSubrogationRecovery/{parentId}/claimId", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.AssignClaimToSubrogationRecovery)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignClaimFromSubrogationRecovery/{parentId}", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.UnassignClaimFromSubrogationRecovery)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignExposureToSubrogationRecovery/{parentId}/exposureId", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.AssignExposureToSubrogationRecovery)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignExposureFromSubrogationRecovery/{parentId}", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.UnassignExposureFromSubrogationRecovery)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCounterpartyToSubrogationRecovery/{parentId}/counterpartyId", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.AssignCounterpartyToSubrogationRecovery)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCounterpartyFromSubrogationRecovery/{parentId}", jsonResponseFormatter.FormatToJSON(SubrogationRecoveryController.UnassignCounterpartyFromSubrogationRecovery)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ThirdParty Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ThirdParty/{id}", jsonResponseFormatter.FormatToJSON(ThirdPartyController.GetThirdParty)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ThirdParty", jsonResponseFormatter.FormatToJSON(ThirdPartyController.GetAllThirdParty)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewThirdParty", jsonResponseFormatter.FormatToJSON(ThirdPartyController.CreateThirdParty)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ThirdParty/{id}", jsonResponseFormatter.FormatToJSON(ThirdPartyController.UpdateThirdParty)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteThirdParty/{id}", jsonResponseFormatter.FormatToJSON(ThirdPartyController.DeleteThirdParty)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSubrogationsToThirdParty/{parentId}/subrogationsId", jsonResponseFormatter.FormatToJSON(ThirdPartyController.AddSubrogationsToThirdParty)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSubrogationsFromThirdParty/{parentId}/subrogationsIds", jsonResponseFormatter.FormatToJSON(ThirdPartyController.RemoveSubrogationsFromThirdParty)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Document Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Document/{id}", jsonResponseFormatter.FormatToJSON(DocumentController.GetDocument)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Document", jsonResponseFormatter.FormatToJSON(DocumentController.GetAllDocument)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDocument", jsonResponseFormatter.FormatToJSON(DocumentController.CreateDocument)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Document/{id}", jsonResponseFormatter.FormatToJSON(DocumentController.UpdateDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDocument/{id}", jsonResponseFormatter.FormatToJSON(DocumentController.DeleteDocument)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPolicyToDocument/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(DocumentController.AssignPolicyToDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromDocument/{parentId}", jsonResponseFormatter.FormatToJSON(DocumentController.UnassignPolicyFromDocument)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignClaimToDocument/{parentId}/claimId", jsonResponseFormatter.FormatToJSON(DocumentController.AssignClaimToDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignClaimFromDocument/{parentId}", jsonResponseFormatter.FormatToJSON(DocumentController.UnassignClaimFromDocument)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignApplicationToDocument/{parentId}/applicationId", jsonResponseFormatter.FormatToJSON(DocumentController.AssignApplicationToDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignApplicationFromDocument/{parentId}", jsonResponseFormatter.FormatToJSON(DocumentController.UnassignApplicationFromDocument)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCustomerToDocument/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(DocumentController.AssignCustomerToDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromDocument/{parentId}", jsonResponseFormatter.FormatToJSON(DocumentController.UnassignCustomerFromDocument)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    return router
}
