package router

import (

    FinancialInstitutionController "fintech-on-golang/internal/controller"
    BranchController "fintech-on-golang/internal/controller"
    ProductOfferingController "fintech-on-golang/internal/controller"
    PricingPlanController "fintech-on-golang/internal/controller"
    FeeScheduleController "fintech-on-golang/internal/controller"
    UsageLimitController "fintech-on-golang/internal/controller"
    CustomerController "fintech-on-golang/internal/controller"
    KYCProfileController "fintech-on-golang/internal/controller"
    KYCDocumentController "fintech-on-golang/internal/controller"
    ScreeningController "fintech-on-golang/internal/controller"
    VerifiedAddressController "fintech-on-golang/internal/controller"
    CompliancePolicyController "fintech-on-golang/internal/controller"
    ComplianceAlertController "fintech-on-golang/internal/controller"
    ConsentController "fintech-on-golang/internal/controller"
    APIClientController "fintech-on-golang/internal/controller"
    AgreementController "fintech-on-golang/internal/controller"
    AccountController "fintech-on-golang/internal/controller"
    WalletController "fintech-on-golang/internal/controller"
    PaymentCardController "fintech-on-golang/internal/controller"
    CardTokenizationController "fintech-on-golang/internal/controller"
    MerchantController "fintech-on-golang/internal/controller"
    TerminalController "fintech-on-golang/internal/controller"
    PaymentContractController "fintech-on-golang/internal/controller"
    PaymentProcessorController "fintech-on-golang/internal/controller"
    TransactionController "fintech-on-golang/internal/controller"
    PaymentOrderController "fintech-on-golang/internal/controller"
    BeneficiaryController "fintech-on-golang/internal/controller"
    AppliedFeeController "fintech-on-golang/internal/controller"
    FXQuoteController "fintech-on-golang/internal/controller"
    FXDealController "fintech-on-golang/internal/controller"
    SettlementBatchController "fintech-on-golang/internal/controller"
    PayoutController "fintech-on-golang/internal/controller"
    DisputeController "fintech-on-golang/internal/controller"
    ChargebackController "fintech-on-golang/internal/controller"
    InvoiceController "fintech-on-golang/internal/controller"
    AccountStatementController "fintech-on-golang/internal/controller"
    DirectDebitMandateController "fintech-on-golang/internal/controller"
    CreditorController "fintech-on-golang/internal/controller"
    LoanApplicationController "fintech-on-golang/internal/controller"
    RiskAssessmentController "fintech-on-golang/internal/controller"
    LoanController "fintech-on-golang/internal/controller"
    RepaymentScheduleController "fintech-on-golang/internal/controller"
    CollateralController "fintech-on-golang/internal/controller"
    LoanTransactionController "fintech-on-golang/internal/controller"
    InvestmentPortfolioController "fintech-on-golang/internal/controller"
    InvestmentAccountController "fintech-on-golang/internal/controller"
    SecurityController "fintech-on-golang/internal/controller"
    PositionController "fintech-on-golang/internal/controller"
    TradeOrderController "fintech-on-golang/internal/controller"
    TradeController "fintech-on-golang/internal/controller"
    ExchangeRateController "fintech-on-golang/internal/controller"
    jsonResponseFormatter "fintech-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "fintech-on-golang/internal/controller"

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
    // FinancialInstitution Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FinancialInstitution/{id}", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.GetFinancialInstitution)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FinancialInstitution", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.GetAllFinancialInstitution)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFinancialInstitution", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.CreateFinancialInstitution)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FinancialInstitution/{id}", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.UpdateFinancialInstitution)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFinancialInstitution/{id}", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.DeleteFinancialInstitution)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddBranchesToFinancialInstitution/{parentId}/branchesId", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.AddBranchesToFinancialInstitution)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBranchesFromFinancialInstitution/{parentId}/branchesIds", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.RemoveBranchesFromFinancialInstitution)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCustomersToFinancialInstitution/{parentId}/customersId", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.AddCustomersToFinancialInstitution)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCustomersFromFinancialInstitution/{parentId}/customersIds", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.RemoveCustomersFromFinancialInstitution)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddProductOfferingsToFinancialInstitution/{parentId}/productOfferingsId", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.AddProductOfferingsToFinancialInstitution)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProductOfferingsFromFinancialInstitution/{parentId}/productOfferingsIds", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.RemoveProductOfferingsFromFinancialInstitution)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPaymentProcessorsToFinancialInstitution/{parentId}/paymentProcessorsId", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.AddPaymentProcessorsToFinancialInstitution)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentProcessorsFromFinancialInstitution/{parentId}/paymentProcessorsIds", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.RemovePaymentProcessorsFromFinancialInstitution)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCompliancePoliciesToFinancialInstitution/{parentId}/compliancePoliciesId", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.AddCompliancePoliciesToFinancialInstitution)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCompliancePoliciesFromFinancialInstitution/{parentId}/compliancePoliciesIds", jsonResponseFormatter.FormatToJSON(FinancialInstitutionController.RemoveCompliancePoliciesFromFinancialInstitution)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Branch Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Branch/{id}", jsonResponseFormatter.FormatToJSON(BranchController.GetBranch)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Branch", jsonResponseFormatter.FormatToJSON(BranchController.GetAllBranch)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBranch", jsonResponseFormatter.FormatToJSON(BranchController.CreateBranch)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Branch/{id}", jsonResponseFormatter.FormatToJSON(BranchController.UpdateBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBranch/{id}", jsonResponseFormatter.FormatToJSON(BranchController.DeleteBranch)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInstitutionToBranch/{parentId}/institutionId", jsonResponseFormatter.FormatToJSON(BranchController.AssignInstitutionToBranch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInstitutionFromBranch/{parentId}", jsonResponseFormatter.FormatToJSON(BranchController.UnassignInstitutionFromBranch)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ProductOffering Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ProductOffering/{id}", jsonResponseFormatter.FormatToJSON(ProductOfferingController.GetProductOffering)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ProductOffering", jsonResponseFormatter.FormatToJSON(ProductOfferingController.GetAllProductOffering)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewProductOffering", jsonResponseFormatter.FormatToJSON(ProductOfferingController.CreateProductOffering)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ProductOffering/{id}", jsonResponseFormatter.FormatToJSON(ProductOfferingController.UpdateProductOffering)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteProductOffering/{id}", jsonResponseFormatter.FormatToJSON(ProductOfferingController.DeleteProductOffering)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInstitutionToProductOffering/{parentId}/institutionId", jsonResponseFormatter.FormatToJSON(ProductOfferingController.AssignInstitutionToProductOffering)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInstitutionFromProductOffering/{parentId}", jsonResponseFormatter.FormatToJSON(ProductOfferingController.UnassignInstitutionFromProductOffering)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPricingPlansToProductOffering/{parentId}/pricingPlansId", jsonResponseFormatter.FormatToJSON(ProductOfferingController.AddPricingPlansToProductOffering)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePricingPlansFromProductOffering/{parentId}/pricingPlansIds", jsonResponseFormatter.FormatToJSON(ProductOfferingController.RemovePricingPlansFromProductOffering)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PricingPlan Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PricingPlan/{id}", jsonResponseFormatter.FormatToJSON(PricingPlanController.GetPricingPlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PricingPlan", jsonResponseFormatter.FormatToJSON(PricingPlanController.GetAllPricingPlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPricingPlan", jsonResponseFormatter.FormatToJSON(PricingPlanController.CreatePricingPlan)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PricingPlan/{id}", jsonResponseFormatter.FormatToJSON(PricingPlanController.UpdatePricingPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePricingPlan/{id}", jsonResponseFormatter.FormatToJSON(PricingPlanController.DeletePricingPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignProductOfferingToPricingPlan/{parentId}/productOfferingId", jsonResponseFormatter.FormatToJSON(PricingPlanController.AssignProductOfferingToPricingPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductOfferingFromPricingPlan/{parentId}", jsonResponseFormatter.FormatToJSON(PricingPlanController.UnassignProductOfferingFromPricingPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddFeeSchedulesToPricingPlan/{parentId}/feeSchedulesId", jsonResponseFormatter.FormatToJSON(PricingPlanController.AddFeeSchedulesToPricingPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeeSchedulesFromPricingPlan/{parentId}/feeSchedulesIds", jsonResponseFormatter.FormatToJSON(PricingPlanController.RemoveFeeSchedulesFromPricingPlan)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLimitsToPricingPlan/{parentId}/limitsId", jsonResponseFormatter.FormatToJSON(PricingPlanController.AddLimitsToPricingPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLimitsFromPricingPlan/{parentId}/limitsIds", jsonResponseFormatter.FormatToJSON(PricingPlanController.RemoveLimitsFromPricingPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // FeeSchedule Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FeeSchedule/{id}", jsonResponseFormatter.FormatToJSON(FeeScheduleController.GetFeeSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FeeSchedule", jsonResponseFormatter.FormatToJSON(FeeScheduleController.GetAllFeeSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFeeSchedule", jsonResponseFormatter.FormatToJSON(FeeScheduleController.CreateFeeSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FeeSchedule/{id}", jsonResponseFormatter.FormatToJSON(FeeScheduleController.UpdateFeeSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFeeSchedule/{id}", jsonResponseFormatter.FormatToJSON(FeeScheduleController.DeleteFeeSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPricingPlanToFeeSchedule/{parentId}/pricingPlanId", jsonResponseFormatter.FormatToJSON(FeeScheduleController.AssignPricingPlanToFeeSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPricingPlanFromFeeSchedule/{parentId}", jsonResponseFormatter.FormatToJSON(FeeScheduleController.UnassignPricingPlanFromFeeSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // UsageLimit Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/UsageLimit/{id}", jsonResponseFormatter.FormatToJSON(UsageLimitController.GetUsageLimit)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/UsageLimit", jsonResponseFormatter.FormatToJSON(UsageLimitController.GetAllUsageLimit)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewUsageLimit", jsonResponseFormatter.FormatToJSON(UsageLimitController.CreateUsageLimit)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/UsageLimit/{id}", jsonResponseFormatter.FormatToJSON(UsageLimitController.UpdateUsageLimit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteUsageLimit/{id}", jsonResponseFormatter.FormatToJSON(UsageLimitController.DeleteUsageLimit)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPricingPlanToUsageLimit/{parentId}/pricingPlanId", jsonResponseFormatter.FormatToJSON(UsageLimitController.AssignPricingPlanToUsageLimit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPricingPlanFromUsageLimit/{parentId}", jsonResponseFormatter.FormatToJSON(UsageLimitController.UnassignPricingPlanFromUsageLimit)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

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
    router.HandleFunc("/api/AssignInstitutionToCustomer/{parentId}/institutionId", jsonResponseFormatter.FormatToJSON(CustomerController.AssignInstitutionToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInstitutionFromCustomer/{parentId}", jsonResponseFormatter.FormatToJSON(CustomerController.UnassignInstitutionFromCustomer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAccountsToCustomer/{parentId}/accountsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddAccountsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAccountsFromCustomer/{parentId}/accountsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveAccountsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddWalletsToCustomer/{parentId}/walletsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddWalletsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWalletsFromCustomer/{parentId}/walletsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveWalletsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCardsToCustomer/{parentId}/cardsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddCardsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCardsFromCustomer/{parentId}/cardsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveCardsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddKycProfilesToCustomer/{parentId}/kycProfilesId", jsonResponseFormatter.FormatToJSON(CustomerController.AddKycProfilesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveKycProfilesFromCustomer/{parentId}/kycProfilesIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveKycProfilesFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddConsentsToCustomer/{parentId}/consentsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddConsentsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveConsentsFromCustomer/{parentId}/consentsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveConsentsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAgreementsToCustomer/{parentId}/agreementsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddAgreementsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAgreementsFromCustomer/{parentId}/agreementsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveAgreementsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLoanApplicationsToCustomer/{parentId}/loanApplicationsId", jsonResponseFormatter.FormatToJSON(CustomerController.AddLoanApplicationsToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLoanApplicationsFromCustomer/{parentId}/loanApplicationsIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveLoanApplicationsFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLoansToCustomer/{parentId}/loansId", jsonResponseFormatter.FormatToJSON(CustomerController.AddLoansToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLoansFromCustomer/{parentId}/loansIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveLoansFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPortfoliosToCustomer/{parentId}/portfoliosId", jsonResponseFormatter.FormatToJSON(CustomerController.AddPortfoliosToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePortfoliosFromCustomer/{parentId}/portfoliosIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemovePortfoliosFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDisputesToCustomer/{parentId}/disputesId", jsonResponseFormatter.FormatToJSON(CustomerController.AddDisputesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDisputesFromCustomer/{parentId}/disputesIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveDisputesFromCustomer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // KYCProfile Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/KYCProfile/{id}", jsonResponseFormatter.FormatToJSON(KYCProfileController.GetKYCProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/KYCProfile", jsonResponseFormatter.FormatToJSON(KYCProfileController.GetAllKYCProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewKYCProfile", jsonResponseFormatter.FormatToJSON(KYCProfileController.CreateKYCProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/KYCProfile/{id}", jsonResponseFormatter.FormatToJSON(KYCProfileController.UpdateKYCProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteKYCProfile/{id}", jsonResponseFormatter.FormatToJSON(KYCProfileController.DeleteKYCProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToKYCProfile/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(KYCProfileController.AssignCustomerToKYCProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromKYCProfile/{parentId}", jsonResponseFormatter.FormatToJSON(KYCProfileController.UnassignCustomerFromKYCProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDocumentsToKYCProfile/{parentId}/documentsId", jsonResponseFormatter.FormatToJSON(KYCProfileController.AddDocumentsToKYCProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDocumentsFromKYCProfile/{parentId}/documentsIds", jsonResponseFormatter.FormatToJSON(KYCProfileController.RemoveDocumentsFromKYCProfile)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddScreeningsToKYCProfile/{parentId}/screeningsId", jsonResponseFormatter.FormatToJSON(KYCProfileController.AddScreeningsToKYCProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveScreeningsFromKYCProfile/{parentId}/screeningsIds", jsonResponseFormatter.FormatToJSON(KYCProfileController.RemoveScreeningsFromKYCProfile)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAddressesToKYCProfile/{parentId}/addressesId", jsonResponseFormatter.FormatToJSON(KYCProfileController.AddAddressesToKYCProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAddressesFromKYCProfile/{parentId}/addressesIds", jsonResponseFormatter.FormatToJSON(KYCProfileController.RemoveAddressesFromKYCProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // KYCDocument Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/KYCDocument/{id}", jsonResponseFormatter.FormatToJSON(KYCDocumentController.GetKYCDocument)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/KYCDocument", jsonResponseFormatter.FormatToJSON(KYCDocumentController.GetAllKYCDocument)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewKYCDocument", jsonResponseFormatter.FormatToJSON(KYCDocumentController.CreateKYCDocument)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/KYCDocument/{id}", jsonResponseFormatter.FormatToJSON(KYCDocumentController.UpdateKYCDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteKYCDocument/{id}", jsonResponseFormatter.FormatToJSON(KYCDocumentController.DeleteKYCDocument)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignKycProfileToKYCDocument/{parentId}/kycProfileId", jsonResponseFormatter.FormatToJSON(KYCDocumentController.AssignKycProfileToKYCDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignKycProfileFromKYCDocument/{parentId}", jsonResponseFormatter.FormatToJSON(KYCDocumentController.UnassignKycProfileFromKYCDocument)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Screening Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Screening/{id}", jsonResponseFormatter.FormatToJSON(ScreeningController.GetScreening)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Screening", jsonResponseFormatter.FormatToJSON(ScreeningController.GetAllScreening)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewScreening", jsonResponseFormatter.FormatToJSON(ScreeningController.CreateScreening)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Screening/{id}", jsonResponseFormatter.FormatToJSON(ScreeningController.UpdateScreening)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteScreening/{id}", jsonResponseFormatter.FormatToJSON(ScreeningController.DeleteScreening)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignKycProfileToScreening/{parentId}/kycProfileId", jsonResponseFormatter.FormatToJSON(ScreeningController.AssignKycProfileToScreening)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignKycProfileFromScreening/{parentId}", jsonResponseFormatter.FormatToJSON(ScreeningController.UnassignKycProfileFromScreening)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAlertsToScreening/{parentId}/alertsId", jsonResponseFormatter.FormatToJSON(ScreeningController.AddAlertsToScreening)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAlertsFromScreening/{parentId}/alertsIds", jsonResponseFormatter.FormatToJSON(ScreeningController.RemoveAlertsFromScreening)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // VerifiedAddress Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/VerifiedAddress/{id}", jsonResponseFormatter.FormatToJSON(VerifiedAddressController.GetVerifiedAddress)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/VerifiedAddress", jsonResponseFormatter.FormatToJSON(VerifiedAddressController.GetAllVerifiedAddress)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewVerifiedAddress", jsonResponseFormatter.FormatToJSON(VerifiedAddressController.CreateVerifiedAddress)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/VerifiedAddress/{id}", jsonResponseFormatter.FormatToJSON(VerifiedAddressController.UpdateVerifiedAddress)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteVerifiedAddress/{id}", jsonResponseFormatter.FormatToJSON(VerifiedAddressController.DeleteVerifiedAddress)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignKycProfileToVerifiedAddress/{parentId}/kycProfileId", jsonResponseFormatter.FormatToJSON(VerifiedAddressController.AssignKycProfileToVerifiedAddress)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignKycProfileFromVerifiedAddress/{parentId}", jsonResponseFormatter.FormatToJSON(VerifiedAddressController.UnassignKycProfileFromVerifiedAddress)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // CompliancePolicy Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CompliancePolicy/{id}", jsonResponseFormatter.FormatToJSON(CompliancePolicyController.GetCompliancePolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CompliancePolicy", jsonResponseFormatter.FormatToJSON(CompliancePolicyController.GetAllCompliancePolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCompliancePolicy", jsonResponseFormatter.FormatToJSON(CompliancePolicyController.CreateCompliancePolicy)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CompliancePolicy/{id}", jsonResponseFormatter.FormatToJSON(CompliancePolicyController.UpdateCompliancePolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCompliancePolicy/{id}", jsonResponseFormatter.FormatToJSON(CompliancePolicyController.DeleteCompliancePolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInstitutionToCompliancePolicy/{parentId}/institutionId", jsonResponseFormatter.FormatToJSON(CompliancePolicyController.AssignInstitutionToCompliancePolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInstitutionFromCompliancePolicy/{parentId}", jsonResponseFormatter.FormatToJSON(CompliancePolicyController.UnassignInstitutionFromCompliancePolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ComplianceAlert Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ComplianceAlert/{id}", jsonResponseFormatter.FormatToJSON(ComplianceAlertController.GetComplianceAlert)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ComplianceAlert", jsonResponseFormatter.FormatToJSON(ComplianceAlertController.GetAllComplianceAlert)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewComplianceAlert", jsonResponseFormatter.FormatToJSON(ComplianceAlertController.CreateComplianceAlert)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ComplianceAlert/{id}", jsonResponseFormatter.FormatToJSON(ComplianceAlertController.UpdateComplianceAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteComplianceAlert/{id}", jsonResponseFormatter.FormatToJSON(ComplianceAlertController.DeleteComplianceAlert)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignScreeningToComplianceAlert/{parentId}/screeningId", jsonResponseFormatter.FormatToJSON(ComplianceAlertController.AssignScreeningToComplianceAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignScreeningFromComplianceAlert/{parentId}", jsonResponseFormatter.FormatToJSON(ComplianceAlertController.UnassignScreeningFromComplianceAlert)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTransactionToComplianceAlert/{parentId}/transactionId", jsonResponseFormatter.FormatToJSON(ComplianceAlertController.AssignTransactionToComplianceAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTransactionFromComplianceAlert/{parentId}", jsonResponseFormatter.FormatToJSON(ComplianceAlertController.UnassignTransactionFromComplianceAlert)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Consent Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Consent/{id}", jsonResponseFormatter.FormatToJSON(ConsentController.GetConsent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Consent", jsonResponseFormatter.FormatToJSON(ConsentController.GetAllConsent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewConsent", jsonResponseFormatter.FormatToJSON(ConsentController.CreateConsent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Consent/{id}", jsonResponseFormatter.FormatToJSON(ConsentController.UpdateConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteConsent/{id}", jsonResponseFormatter.FormatToJSON(ConsentController.DeleteConsent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToConsent/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(ConsentController.AssignCustomerToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromConsent/{parentId}", jsonResponseFormatter.FormatToJSON(ConsentController.UnassignCustomerFromConsent)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignApiClientToConsent/{parentId}/apiClientId", jsonResponseFormatter.FormatToJSON(ConsentController.AssignApiClientToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignApiClientFromConsent/{parentId}", jsonResponseFormatter.FormatToJSON(ConsentController.UnassignApiClientFromConsent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // APIClient Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/APIClient/{id}", jsonResponseFormatter.FormatToJSON(APIClientController.GetAPIClient)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/APIClient", jsonResponseFormatter.FormatToJSON(APIClientController.GetAllAPIClient)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAPIClient", jsonResponseFormatter.FormatToJSON(APIClientController.CreateAPIClient)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/APIClient/{id}", jsonResponseFormatter.FormatToJSON(APIClientController.UpdateAPIClient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAPIClient/{id}", jsonResponseFormatter.FormatToJSON(APIClientController.DeleteAPIClient)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddConsentsToAPIClient/{parentId}/consentsId", jsonResponseFormatter.FormatToJSON(APIClientController.AddConsentsToAPIClient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveConsentsFromAPIClient/{parentId}/consentsIds", jsonResponseFormatter.FormatToJSON(APIClientController.RemoveConsentsFromAPIClient)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Agreement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Agreement/{id}", jsonResponseFormatter.FormatToJSON(AgreementController.GetAgreement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Agreement", jsonResponseFormatter.FormatToJSON(AgreementController.GetAllAgreement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAgreement", jsonResponseFormatter.FormatToJSON(AgreementController.CreateAgreement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Agreement/{id}", jsonResponseFormatter.FormatToJSON(AgreementController.UpdateAgreement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAgreement/{id}", jsonResponseFormatter.FormatToJSON(AgreementController.DeleteAgreement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToAgreement/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(AgreementController.AssignCustomerToAgreement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromAgreement/{parentId}", jsonResponseFormatter.FormatToJSON(AgreementController.UnassignCustomerFromAgreement)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProductOfferingToAgreement/{parentId}/productOfferingId", jsonResponseFormatter.FormatToJSON(AgreementController.AssignProductOfferingToAgreement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductOfferingFromAgreement/{parentId}", jsonResponseFormatter.FormatToJSON(AgreementController.UnassignProductOfferingFromAgreement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Account Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Account/{id}", jsonResponseFormatter.FormatToJSON(AccountController.GetAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Account", jsonResponseFormatter.FormatToJSON(AccountController.GetAllAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAccount", jsonResponseFormatter.FormatToJSON(AccountController.CreateAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Account/{id}", jsonResponseFormatter.FormatToJSON(AccountController.UpdateAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAccount/{id}", jsonResponseFormatter.FormatToJSON(AccountController.DeleteAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToAccount/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(AccountController.AssignCustomerToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AccountController.UnassignCustomerFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInstitutionToAccount/{parentId}/institutionId", jsonResponseFormatter.FormatToJSON(AccountController.AssignInstitutionToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInstitutionFromAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AccountController.UnassignInstitutionFromAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTransactionsToAccount/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(AccountController.AddTransactionsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromAccount/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveTransactionsFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCardsToAccount/{parentId}/cardsId", jsonResponseFormatter.FormatToJSON(AccountController.AddCardsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCardsFromAccount/{parentId}/cardsIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveCardsFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddStatementsToAccount/{parentId}/statementsId", jsonResponseFormatter.FormatToJSON(AccountController.AddStatementsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveStatementsFromAccount/{parentId}/statementsIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveStatementsFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMandatesToAccount/{parentId}/mandatesId", jsonResponseFormatter.FormatToJSON(AccountController.AddMandatesToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMandatesFromAccount/{parentId}/mandatesIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveMandatesFromAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Wallet Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Wallet/{id}", jsonResponseFormatter.FormatToJSON(WalletController.GetWallet)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Wallet", jsonResponseFormatter.FormatToJSON(WalletController.GetAllWallet)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewWallet", jsonResponseFormatter.FormatToJSON(WalletController.CreateWallet)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Wallet/{id}", jsonResponseFormatter.FormatToJSON(WalletController.UpdateWallet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteWallet/{id}", jsonResponseFormatter.FormatToJSON(WalletController.DeleteWallet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToWallet/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(WalletController.AssignCustomerToWallet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromWallet/{parentId}", jsonResponseFormatter.FormatToJSON(WalletController.UnassignCustomerFromWallet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTransactionsToWallet/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(WalletController.AddTransactionsToWallet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromWallet/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(WalletController.RemoveTransactionsFromWallet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PaymentCard Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PaymentCard/{id}", jsonResponseFormatter.FormatToJSON(PaymentCardController.GetPaymentCard)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PaymentCard", jsonResponseFormatter.FormatToJSON(PaymentCardController.GetAllPaymentCard)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPaymentCard", jsonResponseFormatter.FormatToJSON(PaymentCardController.CreatePaymentCard)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentCard/{id}", jsonResponseFormatter.FormatToJSON(PaymentCardController.UpdatePaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePaymentCard/{id}", jsonResponseFormatter.FormatToJSON(PaymentCardController.DeletePaymentCard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToPaymentCard/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(PaymentCardController.AssignCustomerToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromPaymentCard/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentCardController.UnassignCustomerFromPaymentCard)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToPaymentCard/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(PaymentCardController.AssignAccountToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromPaymentCard/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentCardController.UnassignAccountFromPaymentCard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTokenizationsToPaymentCard/{parentId}/tokenizationsId", jsonResponseFormatter.FormatToJSON(PaymentCardController.AddTokenizationsToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTokenizationsFromPaymentCard/{parentId}/tokenizationsIds", jsonResponseFormatter.FormatToJSON(PaymentCardController.RemoveTokenizationsFromPaymentCard)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDisputesToPaymentCard/{parentId}/disputesId", jsonResponseFormatter.FormatToJSON(PaymentCardController.AddDisputesToPaymentCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDisputesFromPaymentCard/{parentId}/disputesIds", jsonResponseFormatter.FormatToJSON(PaymentCardController.RemoveDisputesFromPaymentCard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CardTokenization Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CardTokenization/{id}", jsonResponseFormatter.FormatToJSON(CardTokenizationController.GetCardTokenization)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CardTokenization", jsonResponseFormatter.FormatToJSON(CardTokenizationController.GetAllCardTokenization)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCardTokenization", jsonResponseFormatter.FormatToJSON(CardTokenizationController.CreateCardTokenization)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CardTokenization/{id}", jsonResponseFormatter.FormatToJSON(CardTokenizationController.UpdateCardTokenization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCardTokenization/{id}", jsonResponseFormatter.FormatToJSON(CardTokenizationController.DeleteCardTokenization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCardToCardTokenization/{parentId}/cardId", jsonResponseFormatter.FormatToJSON(CardTokenizationController.AssignCardToCardTokenization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCardFromCardTokenization/{parentId}", jsonResponseFormatter.FormatToJSON(CardTokenizationController.UnassignCardFromCardTokenization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Merchant Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Merchant/{id}", jsonResponseFormatter.FormatToJSON(MerchantController.GetMerchant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Merchant", jsonResponseFormatter.FormatToJSON(MerchantController.GetAllMerchant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMerchant", jsonResponseFormatter.FormatToJSON(MerchantController.CreateMerchant)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Merchant/{id}", jsonResponseFormatter.FormatToJSON(MerchantController.UpdateMerchant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMerchant/{id}", jsonResponseFormatter.FormatToJSON(MerchantController.DeleteMerchant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTerminalsToMerchant/{parentId}/terminalsId", jsonResponseFormatter.FormatToJSON(MerchantController.AddTerminalsToMerchant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTerminalsFromMerchant/{parentId}/terminalsIds", jsonResponseFormatter.FormatToJSON(MerchantController.RemoveTerminalsFromMerchant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPaymentContractsToMerchant/{parentId}/paymentContractsId", jsonResponseFormatter.FormatToJSON(MerchantController.AddPaymentContractsToMerchant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentContractsFromMerchant/{parentId}/paymentContractsIds", jsonResponseFormatter.FormatToJSON(MerchantController.RemovePaymentContractsFromMerchant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPayoutsToMerchant/{parentId}/payoutsId", jsonResponseFormatter.FormatToJSON(MerchantController.AddPayoutsToMerchant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePayoutsFromMerchant/{parentId}/payoutsIds", jsonResponseFormatter.FormatToJSON(MerchantController.RemovePayoutsFromMerchant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSettlementsToMerchant/{parentId}/settlementsId", jsonResponseFormatter.FormatToJSON(MerchantController.AddSettlementsToMerchant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSettlementsFromMerchant/{parentId}/settlementsIds", jsonResponseFormatter.FormatToJSON(MerchantController.RemoveSettlementsFromMerchant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDisputesToMerchant/{parentId}/disputesId", jsonResponseFormatter.FormatToJSON(MerchantController.AddDisputesToMerchant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDisputesFromMerchant/{parentId}/disputesIds", jsonResponseFormatter.FormatToJSON(MerchantController.RemoveDisputesFromMerchant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInvoicesToMerchant/{parentId}/invoicesId", jsonResponseFormatter.FormatToJSON(MerchantController.AddInvoicesToMerchant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInvoicesFromMerchant/{parentId}/invoicesIds", jsonResponseFormatter.FormatToJSON(MerchantController.RemoveInvoicesFromMerchant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Terminal Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Terminal/{id}", jsonResponseFormatter.FormatToJSON(TerminalController.GetTerminal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Terminal", jsonResponseFormatter.FormatToJSON(TerminalController.GetAllTerminal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTerminal", jsonResponseFormatter.FormatToJSON(TerminalController.CreateTerminal)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Terminal/{id}", jsonResponseFormatter.FormatToJSON(TerminalController.UpdateTerminal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTerminal/{id}", jsonResponseFormatter.FormatToJSON(TerminalController.DeleteTerminal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignMerchantToTerminal/{parentId}/merchantId", jsonResponseFormatter.FormatToJSON(TerminalController.AssignMerchantToTerminal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMerchantFromTerminal/{parentId}", jsonResponseFormatter.FormatToJSON(TerminalController.UnassignMerchantFromTerminal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // PaymentContract Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PaymentContract/{id}", jsonResponseFormatter.FormatToJSON(PaymentContractController.GetPaymentContract)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PaymentContract", jsonResponseFormatter.FormatToJSON(PaymentContractController.GetAllPaymentContract)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPaymentContract", jsonResponseFormatter.FormatToJSON(PaymentContractController.CreatePaymentContract)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentContract/{id}", jsonResponseFormatter.FormatToJSON(PaymentContractController.UpdatePaymentContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePaymentContract/{id}", jsonResponseFormatter.FormatToJSON(PaymentContractController.DeletePaymentContract)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignMerchantToPaymentContract/{parentId}/merchantId", jsonResponseFormatter.FormatToJSON(PaymentContractController.AssignMerchantToPaymentContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMerchantFromPaymentContract/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentContractController.UnassignMerchantFromPaymentContract)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAcquirerToPaymentContract/{parentId}/acquirerId", jsonResponseFormatter.FormatToJSON(PaymentContractController.AssignAcquirerToPaymentContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAcquirerFromPaymentContract/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentContractController.UnassignAcquirerFromPaymentContract)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // PaymentProcessor Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PaymentProcessor/{id}", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.GetPaymentProcessor)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PaymentProcessor", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.GetAllPaymentProcessor)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPaymentProcessor", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.CreatePaymentProcessor)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentProcessor/{id}", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.UpdatePaymentProcessor)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePaymentProcessor/{id}", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.DeletePaymentProcessor)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInstitutionsToPaymentProcessor/{parentId}/institutionsId", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.AddInstitutionsToPaymentProcessor)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInstitutionsFromPaymentProcessor/{parentId}/institutionsIds", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.RemoveInstitutionsFromPaymentProcessor)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddContractsToPaymentProcessor/{parentId}/contractsId", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.AddContractsToPaymentProcessor)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContractsFromPaymentProcessor/{parentId}/contractsIds", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.RemoveContractsFromPaymentProcessor)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSettlementsToPaymentProcessor/{parentId}/settlementsId", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.AddSettlementsToPaymentProcessor)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSettlementsFromPaymentProcessor/{parentId}/settlementsIds", jsonResponseFormatter.FormatToJSON(PaymentProcessorController.RemoveSettlementsFromPaymentProcessor)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Transaction Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Transaction/{id}", jsonResponseFormatter.FormatToJSON(TransactionController.GetTransaction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Transaction", jsonResponseFormatter.FormatToJSON(TransactionController.GetAllTransaction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTransaction", jsonResponseFormatter.FormatToJSON(TransactionController.CreateTransaction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Transaction/{id}", jsonResponseFormatter.FormatToJSON(TransactionController.UpdateTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTransaction/{id}", jsonResponseFormatter.FormatToJSON(TransactionController.DeleteTransaction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAccountToTransaction/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignAccountToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignAccountFromTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWalletToTransaction/{parentId}/walletId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignWalletToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWalletFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignWalletFromTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPaymentOrderToTransaction/{parentId}/paymentOrderId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignPaymentOrderToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPaymentOrderFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignPaymentOrderFromTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignMerchantToTransaction/{parentId}/merchantId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignMerchantToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMerchantFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignMerchantFromTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCardToTransaction/{parentId}/cardId", jsonResponseFormatter.FormatToJSON(TransactionController.AssignCardToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCardFromTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(TransactionController.UnassignCardFromTransaction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddRelatedTransactionsToTransaction/{parentId}/relatedTransactionsId", jsonResponseFormatter.FormatToJSON(TransactionController.AddRelatedTransactionsToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRelatedTransactionsFromTransaction/{parentId}/relatedTransactionsIds", jsonResponseFormatter.FormatToJSON(TransactionController.RemoveRelatedTransactionsFromTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAlertsToTransaction/{parentId}/alertsId", jsonResponseFormatter.FormatToJSON(TransactionController.AddAlertsToTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAlertsFromTransaction/{parentId}/alertsIds", jsonResponseFormatter.FormatToJSON(TransactionController.RemoveAlertsFromTransaction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PaymentOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PaymentOrder/{id}", jsonResponseFormatter.FormatToJSON(PaymentOrderController.GetPaymentOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PaymentOrder", jsonResponseFormatter.FormatToJSON(PaymentOrderController.GetAllPaymentOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPaymentOrder", jsonResponseFormatter.FormatToJSON(PaymentOrderController.CreatePaymentOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentOrder/{id}", jsonResponseFormatter.FormatToJSON(PaymentOrderController.UpdatePaymentOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePaymentOrder/{id}", jsonResponseFormatter.FormatToJSON(PaymentOrderController.DeletePaymentOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSourceAccountToPaymentOrder/{parentId}/sourceAccountId", jsonResponseFormatter.FormatToJSON(PaymentOrderController.AssignSourceAccountToPaymentOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSourceAccountFromPaymentOrder/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentOrderController.UnassignSourceAccountFromPaymentOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDestinationAccountToPaymentOrder/{parentId}/destinationAccountId", jsonResponseFormatter.FormatToJSON(PaymentOrderController.AssignDestinationAccountToPaymentOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDestinationAccountFromPaymentOrder/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentOrderController.UnassignDestinationAccountFromPaymentOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBeneficiaryToPaymentOrder/{parentId}/beneficiaryId", jsonResponseFormatter.FormatToJSON(PaymentOrderController.AssignBeneficiaryToPaymentOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBeneficiaryFromPaymentOrder/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentOrderController.UnassignBeneficiaryFromPaymentOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignFxDealToPaymentOrder/{parentId}/fxDealId", jsonResponseFormatter.FormatToJSON(PaymentOrderController.AssignFxDealToPaymentOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFxDealFromPaymentOrder/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentOrderController.UnassignFxDealFromPaymentOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTransactionsToPaymentOrder/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(PaymentOrderController.AddTransactionsToPaymentOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromPaymentOrder/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(PaymentOrderController.RemoveTransactionsFromPaymentOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFeesToPaymentOrder/{parentId}/feesId", jsonResponseFormatter.FormatToJSON(PaymentOrderController.AddFeesToPaymentOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeesFromPaymentOrder/{parentId}/feesIds", jsonResponseFormatter.FormatToJSON(PaymentOrderController.RemoveFeesFromPaymentOrder)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AssignCustomerToBeneficiary/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(BeneficiaryController.AssignCustomerToBeneficiary)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromBeneficiary/{parentId}", jsonResponseFormatter.FormatToJSON(BeneficiaryController.UnassignCustomerFromBeneficiary)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // AppliedFee Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AppliedFee/{id}", jsonResponseFormatter.FormatToJSON(AppliedFeeController.GetAppliedFee)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AppliedFee", jsonResponseFormatter.FormatToJSON(AppliedFeeController.GetAllAppliedFee)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAppliedFee", jsonResponseFormatter.FormatToJSON(AppliedFeeController.CreateAppliedFee)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AppliedFee/{id}", jsonResponseFormatter.FormatToJSON(AppliedFeeController.UpdateAppliedFee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAppliedFee/{id}", jsonResponseFormatter.FormatToJSON(AppliedFeeController.DeleteAppliedFee)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPaymentOrderToAppliedFee/{parentId}/paymentOrderId", jsonResponseFormatter.FormatToJSON(AppliedFeeController.AssignPaymentOrderToAppliedFee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPaymentOrderFromAppliedFee/{parentId}", jsonResponseFormatter.FormatToJSON(AppliedFeeController.UnassignPaymentOrderFromAppliedFee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTransactionToAppliedFee/{parentId}/transactionId", jsonResponseFormatter.FormatToJSON(AppliedFeeController.AssignTransactionToAppliedFee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTransactionFromAppliedFee/{parentId}", jsonResponseFormatter.FormatToJSON(AppliedFeeController.UnassignTransactionFromAppliedFee)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // FXQuote Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FXQuote/{id}", jsonResponseFormatter.FormatToJSON(FXQuoteController.GetFXQuote)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FXQuote", jsonResponseFormatter.FormatToJSON(FXQuoteController.GetAllFXQuote)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFXQuote", jsonResponseFormatter.FormatToJSON(FXQuoteController.CreateFXQuote)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FXQuote/{id}", jsonResponseFormatter.FormatToJSON(FXQuoteController.UpdateFXQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFXQuote/{id}", jsonResponseFormatter.FormatToJSON(FXQuoteController.DeleteFXQuote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRequestedByToFXQuote/{parentId}/requestedById", jsonResponseFormatter.FormatToJSON(FXQuoteController.AssignRequestedByToFXQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRequestedByFromFXQuote/{parentId}", jsonResponseFormatter.FormatToJSON(FXQuoteController.UnassignRequestedByFromFXQuote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // FXDeal Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FXDeal/{id}", jsonResponseFormatter.FormatToJSON(FXDealController.GetFXDeal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FXDeal", jsonResponseFormatter.FormatToJSON(FXDealController.GetAllFXDeal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFXDeal", jsonResponseFormatter.FormatToJSON(FXDealController.CreateFXDeal)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FXDeal/{id}", jsonResponseFormatter.FormatToJSON(FXDealController.UpdateFXDeal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFXDeal/{id}", jsonResponseFormatter.FormatToJSON(FXDealController.DeleteFXDeal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignQuoteToFXDeal/{parentId}/quoteId", jsonResponseFormatter.FormatToJSON(FXDealController.AssignQuoteToFXDeal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignQuoteFromFXDeal/{parentId}", jsonResponseFormatter.FormatToJSON(FXDealController.UnassignQuoteFromFXDeal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPaymentOrdersToFXDeal/{parentId}/paymentOrdersId", jsonResponseFormatter.FormatToJSON(FXDealController.AddPaymentOrdersToFXDeal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentOrdersFromFXDeal/{parentId}/paymentOrdersIds", jsonResponseFormatter.FormatToJSON(FXDealController.RemovePaymentOrdersFromFXDeal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // SettlementBatch Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SettlementBatch/{id}", jsonResponseFormatter.FormatToJSON(SettlementBatchController.GetSettlementBatch)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SettlementBatch", jsonResponseFormatter.FormatToJSON(SettlementBatchController.GetAllSettlementBatch)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSettlementBatch", jsonResponseFormatter.FormatToJSON(SettlementBatchController.CreateSettlementBatch)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SettlementBatch/{id}", jsonResponseFormatter.FormatToJSON(SettlementBatchController.UpdateSettlementBatch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSettlementBatch/{id}", jsonResponseFormatter.FormatToJSON(SettlementBatchController.DeleteSettlementBatch)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignProcessorToSettlementBatch/{parentId}/processorId", jsonResponseFormatter.FormatToJSON(SettlementBatchController.AssignProcessorToSettlementBatch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProcessorFromSettlementBatch/{parentId}", jsonResponseFormatter.FormatToJSON(SettlementBatchController.UnassignProcessorFromSettlementBatch)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignMerchantToSettlementBatch/{parentId}/merchantId", jsonResponseFormatter.FormatToJSON(SettlementBatchController.AssignMerchantToSettlementBatch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMerchantFromSettlementBatch/{parentId}", jsonResponseFormatter.FormatToJSON(SettlementBatchController.UnassignMerchantFromSettlementBatch)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPayoutsToSettlementBatch/{parentId}/payoutsId", jsonResponseFormatter.FormatToJSON(SettlementBatchController.AddPayoutsToSettlementBatch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePayoutsFromSettlementBatch/{parentId}/payoutsIds", jsonResponseFormatter.FormatToJSON(SettlementBatchController.RemovePayoutsFromSettlementBatch)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTransactionsToSettlementBatch/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(SettlementBatchController.AddTransactionsToSettlementBatch)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromSettlementBatch/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(SettlementBatchController.RemoveTransactionsFromSettlementBatch)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Payout Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Payout/{id}", jsonResponseFormatter.FormatToJSON(PayoutController.GetPayout)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Payout", jsonResponseFormatter.FormatToJSON(PayoutController.GetAllPayout)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPayout", jsonResponseFormatter.FormatToJSON(PayoutController.CreatePayout)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Payout/{id}", jsonResponseFormatter.FormatToJSON(PayoutController.UpdatePayout)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePayout/{id}", jsonResponseFormatter.FormatToJSON(PayoutController.DeletePayout)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignMerchantToPayout/{parentId}/merchantId", jsonResponseFormatter.FormatToJSON(PayoutController.AssignMerchantToPayout)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMerchantFromPayout/{parentId}", jsonResponseFormatter.FormatToJSON(PayoutController.UnassignMerchantFromPayout)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSettlementBatchToPayout/{parentId}/settlementBatchId", jsonResponseFormatter.FormatToJSON(PayoutController.AssignSettlementBatchToPayout)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSettlementBatchFromPayout/{parentId}", jsonResponseFormatter.FormatToJSON(PayoutController.UnassignSettlementBatchFromPayout)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDestinationAccountToPayout/{parentId}/destinationAccountId", jsonResponseFormatter.FormatToJSON(PayoutController.AssignDestinationAccountToPayout)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDestinationAccountFromPayout/{parentId}", jsonResponseFormatter.FormatToJSON(PayoutController.UnassignDestinationAccountFromPayout)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Dispute Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Dispute/{id}", jsonResponseFormatter.FormatToJSON(DisputeController.GetDispute)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Dispute", jsonResponseFormatter.FormatToJSON(DisputeController.GetAllDispute)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDispute", jsonResponseFormatter.FormatToJSON(DisputeController.CreateDispute)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Dispute/{id}", jsonResponseFormatter.FormatToJSON(DisputeController.UpdateDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDispute/{id}", jsonResponseFormatter.FormatToJSON(DisputeController.DeleteDispute)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignTransactionToDispute/{parentId}/transactionId", jsonResponseFormatter.FormatToJSON(DisputeController.AssignTransactionToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTransactionFromDispute/{parentId}", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignTransactionFromDispute)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCardToDispute/{parentId}/cardId", jsonResponseFormatter.FormatToJSON(DisputeController.AssignCardToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCardFromDispute/{parentId}", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignCardFromDispute)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignMerchantToDispute/{parentId}/merchantId", jsonResponseFormatter.FormatToJSON(DisputeController.AssignMerchantToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMerchantFromDispute/{parentId}", jsonResponseFormatter.FormatToJSON(DisputeController.UnassignMerchantFromDispute)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddChargebacksToDispute/{parentId}/chargebacksId", jsonResponseFormatter.FormatToJSON(DisputeController.AddChargebacksToDispute)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveChargebacksFromDispute/{parentId}/chargebacksIds", jsonResponseFormatter.FormatToJSON(DisputeController.RemoveChargebacksFromDispute)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Chargeback Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Chargeback/{id}", jsonResponseFormatter.FormatToJSON(ChargebackController.GetChargeback)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Chargeback", jsonResponseFormatter.FormatToJSON(ChargebackController.GetAllChargeback)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewChargeback", jsonResponseFormatter.FormatToJSON(ChargebackController.CreateChargeback)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Chargeback/{id}", jsonResponseFormatter.FormatToJSON(ChargebackController.UpdateChargeback)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteChargeback/{id}", jsonResponseFormatter.FormatToJSON(ChargebackController.DeleteChargeback)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignDisputeToChargeback/{parentId}/disputeId", jsonResponseFormatter.FormatToJSON(ChargebackController.AssignDisputeToChargeback)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDisputeFromChargeback/{parentId}", jsonResponseFormatter.FormatToJSON(ChargebackController.UnassignDisputeFromChargeback)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTransactionToChargeback/{parentId}/transactionId", jsonResponseFormatter.FormatToJSON(ChargebackController.AssignTransactionToChargeback)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTransactionFromChargeback/{parentId}", jsonResponseFormatter.FormatToJSON(ChargebackController.UnassignTransactionFromChargeback)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

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
    router.HandleFunc("/api/AssignMerchantToInvoice/{parentId}/merchantId", jsonResponseFormatter.FormatToJSON(InvoiceController.AssignMerchantToInvoice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMerchantFromInvoice/{parentId}", jsonResponseFormatter.FormatToJSON(InvoiceController.UnassignMerchantFromInvoice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPaymentsToInvoice/{parentId}/paymentsId", jsonResponseFormatter.FormatToJSON(InvoiceController.AddPaymentsToInvoice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentsFromInvoice/{parentId}/paymentsIds", jsonResponseFormatter.FormatToJSON(InvoiceController.RemovePaymentsFromInvoice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AccountStatement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AccountStatement/{id}", jsonResponseFormatter.FormatToJSON(AccountStatementController.GetAccountStatement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AccountStatement", jsonResponseFormatter.FormatToJSON(AccountStatementController.GetAllAccountStatement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAccountStatement", jsonResponseFormatter.FormatToJSON(AccountStatementController.CreateAccountStatement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AccountStatement/{id}", jsonResponseFormatter.FormatToJSON(AccountStatementController.UpdateAccountStatement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAccountStatement/{id}", jsonResponseFormatter.FormatToJSON(AccountStatementController.DeleteAccountStatement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAccountToAccountStatement/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(AccountStatementController.AssignAccountToAccountStatement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromAccountStatement/{parentId}", jsonResponseFormatter.FormatToJSON(AccountStatementController.UnassignAccountFromAccountStatement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // DirectDebitMandate Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DirectDebitMandate/{id}", jsonResponseFormatter.FormatToJSON(DirectDebitMandateController.GetDirectDebitMandate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DirectDebitMandate", jsonResponseFormatter.FormatToJSON(DirectDebitMandateController.GetAllDirectDebitMandate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDirectDebitMandate", jsonResponseFormatter.FormatToJSON(DirectDebitMandateController.CreateDirectDebitMandate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DirectDebitMandate/{id}", jsonResponseFormatter.FormatToJSON(DirectDebitMandateController.UpdateDirectDebitMandate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDirectDebitMandate/{id}", jsonResponseFormatter.FormatToJSON(DirectDebitMandateController.DeleteDirectDebitMandate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAccountToDirectDebitMandate/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(DirectDebitMandateController.AssignAccountToDirectDebitMandate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromDirectDebitMandate/{parentId}", jsonResponseFormatter.FormatToJSON(DirectDebitMandateController.UnassignAccountFromDirectDebitMandate)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCreditorToDirectDebitMandate/{parentId}/creditorId", jsonResponseFormatter.FormatToJSON(DirectDebitMandateController.AssignCreditorToDirectDebitMandate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCreditorFromDirectDebitMandate/{parentId}", jsonResponseFormatter.FormatToJSON(DirectDebitMandateController.UnassignCreditorFromDirectDebitMandate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Creditor Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Creditor/{id}", jsonResponseFormatter.FormatToJSON(CreditorController.GetCreditor)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Creditor", jsonResponseFormatter.FormatToJSON(CreditorController.GetAllCreditor)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCreditor", jsonResponseFormatter.FormatToJSON(CreditorController.CreateCreditor)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Creditor/{id}", jsonResponseFormatter.FormatToJSON(CreditorController.UpdateCreditor)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCreditor/{id}", jsonResponseFormatter.FormatToJSON(CreditorController.DeleteCreditor)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddMandatesToCreditor/{parentId}/mandatesId", jsonResponseFormatter.FormatToJSON(CreditorController.AddMandatesToCreditor)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMandatesFromCreditor/{parentId}/mandatesIds", jsonResponseFormatter.FormatToJSON(CreditorController.RemoveMandatesFromCreditor)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // LoanApplication Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanApplication/{id}", jsonResponseFormatter.FormatToJSON(LoanApplicationController.GetLoanApplication)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LoanApplication", jsonResponseFormatter.FormatToJSON(LoanApplicationController.GetAllLoanApplication)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLoanApplication", jsonResponseFormatter.FormatToJSON(LoanApplicationController.CreateLoanApplication)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanApplication/{id}", jsonResponseFormatter.FormatToJSON(LoanApplicationController.UpdateLoanApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLoanApplication/{id}", jsonResponseFormatter.FormatToJSON(LoanApplicationController.DeleteLoanApplication)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToLoanApplication/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(LoanApplicationController.AssignCustomerToLoanApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromLoanApplication/{parentId}", jsonResponseFormatter.FormatToJSON(LoanApplicationController.UnassignCustomerFromLoanApplication)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRiskAssessmentToLoanApplication/{parentId}/riskAssessmentId", jsonResponseFormatter.FormatToJSON(LoanApplicationController.AssignRiskAssessmentToLoanApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRiskAssessmentFromLoanApplication/{parentId}", jsonResponseFormatter.FormatToJSON(LoanApplicationController.UnassignRiskAssessmentFromLoanApplication)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLoanToLoanApplication/{parentId}/loanId", jsonResponseFormatter.FormatToJSON(LoanApplicationController.AssignLoanToLoanApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLoanFromLoanApplication/{parentId}", jsonResponseFormatter.FormatToJSON(LoanApplicationController.UnassignLoanFromLoanApplication)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // RiskAssessment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RiskAssessment/{id}", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.GetRiskAssessment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RiskAssessment", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.GetAllRiskAssessment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRiskAssessment", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.CreateRiskAssessment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RiskAssessment/{id}", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.UpdateRiskAssessment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRiskAssessment/{id}", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.DeleteRiskAssessment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignApplicationToRiskAssessment/{parentId}/applicationId", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.AssignApplicationToRiskAssessment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignApplicationFromRiskAssessment/{parentId}", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.UnassignApplicationFromRiskAssessment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Loan Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Loan/{id}", jsonResponseFormatter.FormatToJSON(LoanController.GetLoan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Loan", jsonResponseFormatter.FormatToJSON(LoanController.GetAllLoan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLoan", jsonResponseFormatter.FormatToJSON(LoanController.CreateLoan)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Loan/{id}", jsonResponseFormatter.FormatToJSON(LoanController.UpdateLoan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLoan/{id}", jsonResponseFormatter.FormatToJSON(LoanController.DeleteLoan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToLoan/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(LoanController.AssignCustomerToLoan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromLoan/{parentId}", jsonResponseFormatter.FormatToJSON(LoanController.UnassignCustomerFromLoan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddScheduleToLoan/{parentId}/scheduleId", jsonResponseFormatter.FormatToJSON(LoanController.AddScheduleToLoan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveScheduleFromLoan/{parentId}/scheduleIds", jsonResponseFormatter.FormatToJSON(LoanController.RemoveScheduleFromLoan)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCollateralToLoan/{parentId}/collateralId", jsonResponseFormatter.FormatToJSON(LoanController.AddCollateralToLoan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCollateralFromLoan/{parentId}/collateralIds", jsonResponseFormatter.FormatToJSON(LoanController.RemoveCollateralFromLoan)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTransactionsToLoan/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(LoanController.AddTransactionsToLoan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromLoan/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(LoanController.RemoveTransactionsFromLoan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // RepaymentSchedule Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RepaymentSchedule/{id}", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.GetRepaymentSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.GetAllRepaymentSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRepaymentSchedule", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.CreateRepaymentSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RepaymentSchedule/{id}", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.UpdateRepaymentSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRepaymentSchedule/{id}", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.DeleteRepaymentSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignLoanToRepaymentSchedule/{parentId}/loanId", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.AssignLoanToRepaymentSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLoanFromRepaymentSchedule/{parentId}", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.UnassignLoanFromRepaymentSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPaymentsToRepaymentSchedule/{parentId}/paymentsId", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.AddPaymentsToRepaymentSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentsFromRepaymentSchedule/{parentId}/paymentsIds", jsonResponseFormatter.FormatToJSON(RepaymentScheduleController.RemovePaymentsFromRepaymentSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Collateral Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Collateral/{id}", jsonResponseFormatter.FormatToJSON(CollateralController.GetCollateral)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Collateral", jsonResponseFormatter.FormatToJSON(CollateralController.GetAllCollateral)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCollateral", jsonResponseFormatter.FormatToJSON(CollateralController.CreateCollateral)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Collateral/{id}", jsonResponseFormatter.FormatToJSON(CollateralController.UpdateCollateral)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCollateral/{id}", jsonResponseFormatter.FormatToJSON(CollateralController.DeleteCollateral)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignLoanToCollateral/{parentId}/loanId", jsonResponseFormatter.FormatToJSON(CollateralController.AssignLoanToCollateral)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLoanFromCollateral/{parentId}", jsonResponseFormatter.FormatToJSON(CollateralController.UnassignLoanFromCollateral)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // LoanTransaction Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LoanTransaction/{id}", jsonResponseFormatter.FormatToJSON(LoanTransactionController.GetLoanTransaction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LoanTransaction", jsonResponseFormatter.FormatToJSON(LoanTransactionController.GetAllLoanTransaction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLoanTransaction", jsonResponseFormatter.FormatToJSON(LoanTransactionController.CreateLoanTransaction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LoanTransaction/{id}", jsonResponseFormatter.FormatToJSON(LoanTransactionController.UpdateLoanTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLoanTransaction/{id}", jsonResponseFormatter.FormatToJSON(LoanTransactionController.DeleteLoanTransaction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignLoanToLoanTransaction/{parentId}/loanId", jsonResponseFormatter.FormatToJSON(LoanTransactionController.AssignLoanToLoanTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLoanFromLoanTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(LoanTransactionController.UnassignLoanFromLoanTransaction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // InvestmentPortfolio Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InvestmentPortfolio/{id}", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.GetInvestmentPortfolio)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InvestmentPortfolio", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.GetAllInvestmentPortfolio)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInvestmentPortfolio", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.CreateInvestmentPortfolio)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InvestmentPortfolio/{id}", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.UpdateInvestmentPortfolio)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInvestmentPortfolio/{id}", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.DeleteInvestmentPortfolio)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToInvestmentPortfolio/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.AssignCustomerToInvestmentPortfolio)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromInvestmentPortfolio/{parentId}", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.UnassignCustomerFromInvestmentPortfolio)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAccountsToInvestmentPortfolio/{parentId}/accountsId", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.AddAccountsToInvestmentPortfolio)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAccountsFromInvestmentPortfolio/{parentId}/accountsIds", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.RemoveAccountsFromInvestmentPortfolio)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOrdersToInvestmentPortfolio/{parentId}/ordersId", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.AddOrdersToInvestmentPortfolio)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOrdersFromInvestmentPortfolio/{parentId}/ordersIds", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.RemoveOrdersFromInvestmentPortfolio)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddHoldingsToInvestmentPortfolio/{parentId}/holdingsId", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.AddHoldingsToInvestmentPortfolio)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveHoldingsFromInvestmentPortfolio/{parentId}/holdingsIds", jsonResponseFormatter.FormatToJSON(InvestmentPortfolioController.RemoveHoldingsFromInvestmentPortfolio)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InvestmentAccount Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InvestmentAccount/{id}", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.GetInvestmentAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InvestmentAccount", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.GetAllInvestmentAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInvestmentAccount", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.CreateInvestmentAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InvestmentAccount/{id}", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.UpdateInvestmentAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInvestmentAccount/{id}", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.DeleteInvestmentAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPortfolioToInvestmentAccount/{parentId}/portfolioId", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.AssignPortfolioToInvestmentAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPortfolioFromInvestmentAccount/{parentId}", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.UnassignPortfolioFromInvestmentAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTradesToInvestmentAccount/{parentId}/tradesId", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.AddTradesToInvestmentAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTradesFromInvestmentAccount/{parentId}/tradesIds", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.RemoveTradesFromInvestmentAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOrdersToInvestmentAccount/{parentId}/ordersId", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.AddOrdersToInvestmentAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOrdersFromInvestmentAccount/{parentId}/ordersIds", jsonResponseFormatter.FormatToJSON(InvestmentAccountController.RemoveOrdersFromInvestmentAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Security Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Security/{id}", jsonResponseFormatter.FormatToJSON(SecurityController.GetSecurity)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Security", jsonResponseFormatter.FormatToJSON(SecurityController.GetAllSecurity)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSecurity", jsonResponseFormatter.FormatToJSON(SecurityController.CreateSecurity)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Security/{id}", jsonResponseFormatter.FormatToJSON(SecurityController.UpdateSecurity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSecurity/{id}", jsonResponseFormatter.FormatToJSON(SecurityController.DeleteSecurity)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPositionsToSecurity/{parentId}/positionsId", jsonResponseFormatter.FormatToJSON(SecurityController.AddPositionsToSecurity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePositionsFromSecurity/{parentId}/positionsIds", jsonResponseFormatter.FormatToJSON(SecurityController.RemovePositionsFromSecurity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTradesToSecurity/{parentId}/tradesId", jsonResponseFormatter.FormatToJSON(SecurityController.AddTradesToSecurity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTradesFromSecurity/{parentId}/tradesIds", jsonResponseFormatter.FormatToJSON(SecurityController.RemoveTradesFromSecurity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOrdersToSecurity/{parentId}/ordersId", jsonResponseFormatter.FormatToJSON(SecurityController.AddOrdersToSecurity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOrdersFromSecurity/{parentId}/ordersIds", jsonResponseFormatter.FormatToJSON(SecurityController.RemoveOrdersFromSecurity)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Position Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Position/{id}", jsonResponseFormatter.FormatToJSON(PositionController.GetPosition)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Position", jsonResponseFormatter.FormatToJSON(PositionController.GetAllPosition)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPosition", jsonResponseFormatter.FormatToJSON(PositionController.CreatePosition)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Position/{id}", jsonResponseFormatter.FormatToJSON(PositionController.UpdatePosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePosition/{id}", jsonResponseFormatter.FormatToJSON(PositionController.DeletePosition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPortfolioToPosition/{parentId}/portfolioId", jsonResponseFormatter.FormatToJSON(PositionController.AssignPortfolioToPosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPortfolioFromPosition/{parentId}", jsonResponseFormatter.FormatToJSON(PositionController.UnassignPortfolioFromPosition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSecurityToPosition/{parentId}/securityId", jsonResponseFormatter.FormatToJSON(PositionController.AssignSecurityToPosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSecurityFromPosition/{parentId}", jsonResponseFormatter.FormatToJSON(PositionController.UnassignSecurityFromPosition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // TradeOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TradeOrder/{id}", jsonResponseFormatter.FormatToJSON(TradeOrderController.GetTradeOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TradeOrder", jsonResponseFormatter.FormatToJSON(TradeOrderController.GetAllTradeOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTradeOrder", jsonResponseFormatter.FormatToJSON(TradeOrderController.CreateTradeOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TradeOrder/{id}", jsonResponseFormatter.FormatToJSON(TradeOrderController.UpdateTradeOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTradeOrder/{id}", jsonResponseFormatter.FormatToJSON(TradeOrderController.DeleteTradeOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPortfolioToTradeOrder/{parentId}/portfolioId", jsonResponseFormatter.FormatToJSON(TradeOrderController.AssignPortfolioToTradeOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPortfolioFromTradeOrder/{parentId}", jsonResponseFormatter.FormatToJSON(TradeOrderController.UnassignPortfolioFromTradeOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSecurityToTradeOrder/{parentId}/securityId", jsonResponseFormatter.FormatToJSON(TradeOrderController.AssignSecurityToTradeOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSecurityFromTradeOrder/{parentId}", jsonResponseFormatter.FormatToJSON(TradeOrderController.UnassignSecurityFromTradeOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTradesToTradeOrder/{parentId}/tradesId", jsonResponseFormatter.FormatToJSON(TradeOrderController.AddTradesToTradeOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTradesFromTradeOrder/{parentId}/tradesIds", jsonResponseFormatter.FormatToJSON(TradeOrderController.RemoveTradesFromTradeOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Trade Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Trade/{id}", jsonResponseFormatter.FormatToJSON(TradeController.GetTrade)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Trade", jsonResponseFormatter.FormatToJSON(TradeController.GetAllTrade)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTrade", jsonResponseFormatter.FormatToJSON(TradeController.CreateTrade)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Trade/{id}", jsonResponseFormatter.FormatToJSON(TradeController.UpdateTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTrade/{id}", jsonResponseFormatter.FormatToJSON(TradeController.DeleteTrade)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrderToTrade/{parentId}/orderId", jsonResponseFormatter.FormatToJSON(TradeController.AssignOrderToTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrderFromTrade/{parentId}", jsonResponseFormatter.FormatToJSON(TradeController.UnassignOrderFromTrade)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSecurityToTrade/{parentId}/securityId", jsonResponseFormatter.FormatToJSON(TradeController.AssignSecurityToTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSecurityFromTrade/{parentId}", jsonResponseFormatter.FormatToJSON(TradeController.UnassignSecurityFromTrade)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInvestmentAccountToTrade/{parentId}/investmentAccountId", jsonResponseFormatter.FormatToJSON(TradeController.AssignInvestmentAccountToTrade)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInvestmentAccountFromTrade/{parentId}", jsonResponseFormatter.FormatToJSON(TradeController.UnassignInvestmentAccountFromTrade)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ExchangeRate Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExchangeRate/{id}", jsonResponseFormatter.FormatToJSON(ExchangeRateController.GetExchangeRate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ExchangeRate", jsonResponseFormatter.FormatToJSON(ExchangeRateController.GetAllExchangeRate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewExchangeRate", jsonResponseFormatter.FormatToJSON(ExchangeRateController.CreateExchangeRate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExchangeRate/{id}", jsonResponseFormatter.FormatToJSON(ExchangeRateController.UpdateExchangeRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteExchangeRate/{id}", jsonResponseFormatter.FormatToJSON(ExchangeRateController.DeleteExchangeRate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddUsedByQuotesToExchangeRate/{parentId}/usedByQuotesId", jsonResponseFormatter.FormatToJSON(ExchangeRateController.AddUsedByQuotesToExchangeRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveUsedByQuotesFromExchangeRate/{parentId}/usedByQuotesIds", jsonResponseFormatter.FormatToJSON(ExchangeRateController.RemoveUsedByQuotesFromExchangeRate)).Methods("DELETE", "OPTIONS")

    return router
}
