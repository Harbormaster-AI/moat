package router

import (

    AgencyController "advertising-on-golang/internal/controller"
    TeamController "advertising-on-golang/internal/controller"
    UserController "advertising-on-golang/internal/controller"
    AdvertiserController "advertising-on-golang/internal/controller"
    BillingProfileController "advertising-on-golang/internal/controller"
    PaymentMethodController "advertising-on-golang/internal/controller"
    AdAccountController "advertising-on-golang/internal/controller"
    DSPController "advertising-on-golang/internal/controller"
    CampaignController "advertising-on-golang/internal/controller"
    KPIController "advertising-on-golang/internal/controller"
    AudienceSegmentController "advertising-on-golang/internal/controller"
    DataProviderController "advertising-on-golang/internal/controller"
    LineItemController "advertising-on-golang/internal/controller"
    TargetingProfileController "advertising-on-golang/internal/controller"
    DeviceCriterionController "advertising-on-golang/internal/controller"
    BrandSafetyPolicyController "advertising-on-golang/internal/controller"
    ContentCategoryController "advertising-on-golang/internal/controller"
    PublisherController "advertising-on-golang/internal/controller"
    InventorySourceController "advertising-on-golang/internal/controller"
    AdSlotController "advertising-on-golang/internal/controller"
    DealController "advertising-on-golang/internal/controller"
    PlacementController "advertising-on-golang/internal/controller"
    CreativeAssetController "advertising-on-golang/internal/controller"
    CreativeFileController "advertising-on-golang/internal/controller"
    CreativeVariationController "advertising-on-golang/internal/controller"
    CreativeApprovalController "advertising-on-golang/internal/controller"
    TrackingPixelController "advertising-on-golang/internal/controller"
    ConversionEventController "advertising-on-golang/internal/controller"
    PerformanceMetricController "advertising-on-golang/internal/controller"
    ReportController "advertising-on-golang/internal/controller"
    InsertionOrderController "advertising-on-golang/internal/controller"
    RateCardController "advertising-on-golang/internal/controller"
    RateController "advertising-on-golang/internal/controller"
    ExperimentController "advertising-on-golang/internal/controller"
    ExperimentVariantController "advertising-on-golang/internal/controller"
    GeoRegionController "advertising-on-golang/internal/controller"
    jsonResponseFormatter "advertising-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "advertising-on-golang/internal/controller"

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
    // Agency Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Agency/{id}", jsonResponseFormatter.FormatToJSON(AgencyController.GetAgency)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Agency", jsonResponseFormatter.FormatToJSON(AgencyController.GetAllAgency)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAgency", jsonResponseFormatter.FormatToJSON(AgencyController.CreateAgency)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Agency/{id}", jsonResponseFormatter.FormatToJSON(AgencyController.UpdateAgency)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAgency/{id}", jsonResponseFormatter.FormatToJSON(AgencyController.DeleteAgency)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAdvertisersToAgency/{parentId}/advertisersId", jsonResponseFormatter.FormatToJSON(AgencyController.AddAdvertisersToAgency)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAdvertisersFromAgency/{parentId}/advertisersIds", jsonResponseFormatter.FormatToJSON(AgencyController.RemoveAdvertisersFromAgency)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTeamsToAgency/{parentId}/teamsId", jsonResponseFormatter.FormatToJSON(AgencyController.AddTeamsToAgency)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTeamsFromAgency/{parentId}/teamsIds", jsonResponseFormatter.FormatToJSON(AgencyController.RemoveTeamsFromAgency)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddUsersToAgency/{parentId}/usersId", jsonResponseFormatter.FormatToJSON(AgencyController.AddUsersToAgency)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveUsersFromAgency/{parentId}/usersIds", jsonResponseFormatter.FormatToJSON(AgencyController.RemoveUsersFromAgency)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInsertionOrdersToAgency/{parentId}/insertionOrdersId", jsonResponseFormatter.FormatToJSON(AgencyController.AddInsertionOrdersToAgency)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInsertionOrdersFromAgency/{parentId}/insertionOrdersIds", jsonResponseFormatter.FormatToJSON(AgencyController.RemoveInsertionOrdersFromAgency)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Team Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Team/{id}", jsonResponseFormatter.FormatToJSON(TeamController.GetTeam)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Team", jsonResponseFormatter.FormatToJSON(TeamController.GetAllTeam)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTeam", jsonResponseFormatter.FormatToJSON(TeamController.CreateTeam)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Team/{id}", jsonResponseFormatter.FormatToJSON(TeamController.UpdateTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTeam/{id}", jsonResponseFormatter.FormatToJSON(TeamController.DeleteTeam)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAgencyToTeam/{parentId}/agencyId", jsonResponseFormatter.FormatToJSON(TeamController.AssignAgencyToTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAgencyFromTeam/{parentId}", jsonResponseFormatter.FormatToJSON(TeamController.UnassignAgencyFromTeam)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddUsersToTeam/{parentId}/usersId", jsonResponseFormatter.FormatToJSON(TeamController.AddUsersToTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveUsersFromTeam/{parentId}/usersIds", jsonResponseFormatter.FormatToJSON(TeamController.RemoveUsersFromTeam)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAdAccountsToTeam/{parentId}/adAccountsId", jsonResponseFormatter.FormatToJSON(TeamController.AddAdAccountsToTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAdAccountsFromTeam/{parentId}/adAccountsIds", jsonResponseFormatter.FormatToJSON(TeamController.RemoveAdAccountsFromTeam)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // User Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/User/{id}", jsonResponseFormatter.FormatToJSON(UserController.GetUser)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/User", jsonResponseFormatter.FormatToJSON(UserController.GetAllUser)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewUser", jsonResponseFormatter.FormatToJSON(UserController.CreateUser)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/User/{id}", jsonResponseFormatter.FormatToJSON(UserController.UpdateUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteUser/{id}", jsonResponseFormatter.FormatToJSON(UserController.DeleteUser)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAgencyToUser/{parentId}/agencyId", jsonResponseFormatter.FormatToJSON(UserController.AssignAgencyToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAgencyFromUser/{parentId}", jsonResponseFormatter.FormatToJSON(UserController.UnassignAgencyFromUser)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTeamsToUser/{parentId}/teamsId", jsonResponseFormatter.FormatToJSON(UserController.AddTeamsToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTeamsFromUser/{parentId}/teamsIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveTeamsFromUser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAdAccountsToUser/{parentId}/adAccountsId", jsonResponseFormatter.FormatToJSON(UserController.AddAdAccountsToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAdAccountsFromUser/{parentId}/adAccountsIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveAdAccountsFromUser)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Advertiser Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Advertiser/{id}", jsonResponseFormatter.FormatToJSON(AdvertiserController.GetAdvertiser)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Advertiser", jsonResponseFormatter.FormatToJSON(AdvertiserController.GetAllAdvertiser)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAdvertiser", jsonResponseFormatter.FormatToJSON(AdvertiserController.CreateAdvertiser)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Advertiser/{id}", jsonResponseFormatter.FormatToJSON(AdvertiserController.UpdateAdvertiser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAdvertiser/{id}", jsonResponseFormatter.FormatToJSON(AdvertiserController.DeleteAdvertiser)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAgencyToAdvertiser/{parentId}/agencyId", jsonResponseFormatter.FormatToJSON(AdvertiserController.AssignAgencyToAdvertiser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAgencyFromAdvertiser/{parentId}", jsonResponseFormatter.FormatToJSON(AdvertiserController.UnassignAgencyFromAdvertiser)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAdAccountsToAdvertiser/{parentId}/adAccountsId", jsonResponseFormatter.FormatToJSON(AdvertiserController.AddAdAccountsToAdvertiser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAdAccountsFromAdvertiser/{parentId}/adAccountsIds", jsonResponseFormatter.FormatToJSON(AdvertiserController.RemoveAdAccountsFromAdvertiser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddBillingProfilesToAdvertiser/{parentId}/billingProfilesId", jsonResponseFormatter.FormatToJSON(AdvertiserController.AddBillingProfilesToAdvertiser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBillingProfilesFromAdvertiser/{parentId}/billingProfilesIds", jsonResponseFormatter.FormatToJSON(AdvertiserController.RemoveBillingProfilesFromAdvertiser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCampaignsToAdvertiser/{parentId}/campaignsId", jsonResponseFormatter.FormatToJSON(AdvertiserController.AddCampaignsToAdvertiser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCampaignsFromAdvertiser/{parentId}/campaignsIds", jsonResponseFormatter.FormatToJSON(AdvertiserController.RemoveCampaignsFromAdvertiser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTrackingPixelsToAdvertiser/{parentId}/trackingPixelsId", jsonResponseFormatter.FormatToJSON(AdvertiserController.AddTrackingPixelsToAdvertiser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTrackingPixelsFromAdvertiser/{parentId}/trackingPixelsIds", jsonResponseFormatter.FormatToJSON(AdvertiserController.RemoveTrackingPixelsFromAdvertiser)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // BillingProfile Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BillingProfile/{id}", jsonResponseFormatter.FormatToJSON(BillingProfileController.GetBillingProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BillingProfile", jsonResponseFormatter.FormatToJSON(BillingProfileController.GetAllBillingProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBillingProfile", jsonResponseFormatter.FormatToJSON(BillingProfileController.CreateBillingProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BillingProfile/{id}", jsonResponseFormatter.FormatToJSON(BillingProfileController.UpdateBillingProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBillingProfile/{id}", jsonResponseFormatter.FormatToJSON(BillingProfileController.DeleteBillingProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAdvertiserToBillingProfile/{parentId}/advertiserId", jsonResponseFormatter.FormatToJSON(BillingProfileController.AssignAdvertiserToBillingProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdvertiserFromBillingProfile/{parentId}", jsonResponseFormatter.FormatToJSON(BillingProfileController.UnassignAdvertiserFromBillingProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPaymentMethodsToBillingProfile/{parentId}/paymentMethodsId", jsonResponseFormatter.FormatToJSON(BillingProfileController.AddPaymentMethodsToBillingProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentMethodsFromBillingProfile/{parentId}/paymentMethodsIds", jsonResponseFormatter.FormatToJSON(BillingProfileController.RemovePaymentMethodsFromBillingProfile)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAdAccountsToBillingProfile/{parentId}/adAccountsId", jsonResponseFormatter.FormatToJSON(BillingProfileController.AddAdAccountsToBillingProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAdAccountsFromBillingProfile/{parentId}/adAccountsIds", jsonResponseFormatter.FormatToJSON(BillingProfileController.RemoveAdAccountsFromBillingProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PaymentMethod Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PaymentMethod/{id}", jsonResponseFormatter.FormatToJSON(PaymentMethodController.GetPaymentMethod)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PaymentMethod", jsonResponseFormatter.FormatToJSON(PaymentMethodController.GetAllPaymentMethod)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPaymentMethod", jsonResponseFormatter.FormatToJSON(PaymentMethodController.CreatePaymentMethod)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentMethod/{id}", jsonResponseFormatter.FormatToJSON(PaymentMethodController.UpdatePaymentMethod)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePaymentMethod/{id}", jsonResponseFormatter.FormatToJSON(PaymentMethodController.DeletePaymentMethod)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBillingProfileToPaymentMethod/{parentId}/billingProfileId", jsonResponseFormatter.FormatToJSON(PaymentMethodController.AssignBillingProfileToPaymentMethod)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBillingProfileFromPaymentMethod/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentMethodController.UnassignBillingProfileFromPaymentMethod)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // AdAccount Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AdAccount/{id}", jsonResponseFormatter.FormatToJSON(AdAccountController.GetAdAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AdAccount", jsonResponseFormatter.FormatToJSON(AdAccountController.GetAllAdAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAdAccount", jsonResponseFormatter.FormatToJSON(AdAccountController.CreateAdAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AdAccount/{id}", jsonResponseFormatter.FormatToJSON(AdAccountController.UpdateAdAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAdAccount/{id}", jsonResponseFormatter.FormatToJSON(AdAccountController.DeleteAdAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAdvertiserToAdAccount/{parentId}/advertiserId", jsonResponseFormatter.FormatToJSON(AdAccountController.AssignAdvertiserToAdAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdvertiserFromAdAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AdAccountController.UnassignAdvertiserFromAdAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBillingProfileToAdAccount/{parentId}/billingProfileId", jsonResponseFormatter.FormatToJSON(AdAccountController.AssignBillingProfileToAdAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBillingProfileFromAdAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AdAccountController.UnassignBillingProfileFromAdAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDspToAdAccount/{parentId}/dspId", jsonResponseFormatter.FormatToJSON(AdAccountController.AssignDspToAdAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDspFromAdAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AdAccountController.UnassignDspFromAdAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddUsersToAdAccount/{parentId}/usersId", jsonResponseFormatter.FormatToJSON(AdAccountController.AddUsersToAdAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveUsersFromAdAccount/{parentId}/usersIds", jsonResponseFormatter.FormatToJSON(AdAccountController.RemoveUsersFromAdAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCampaignsToAdAccount/{parentId}/campaignsId", jsonResponseFormatter.FormatToJSON(AdAccountController.AddCampaignsToAdAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCampaignsFromAdAccount/{parentId}/campaignsIds", jsonResponseFormatter.FormatToJSON(AdAccountController.RemoveCampaignsFromAdAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPerformanceMetricsToAdAccount/{parentId}/performanceMetricsId", jsonResponseFormatter.FormatToJSON(AdAccountController.AddPerformanceMetricsToAdAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePerformanceMetricsFromAdAccount/{parentId}/performanceMetricsIds", jsonResponseFormatter.FormatToJSON(AdAccountController.RemovePerformanceMetricsFromAdAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DSP Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DSP/{id}", jsonResponseFormatter.FormatToJSON(DSPController.GetDSP)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DSP", jsonResponseFormatter.FormatToJSON(DSPController.GetAllDSP)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDSP", jsonResponseFormatter.FormatToJSON(DSPController.CreateDSP)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DSP/{id}", jsonResponseFormatter.FormatToJSON(DSPController.UpdateDSP)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDSP/{id}", jsonResponseFormatter.FormatToJSON(DSPController.DeleteDSP)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAdAccountsToDSP/{parentId}/adAccountsId", jsonResponseFormatter.FormatToJSON(DSPController.AddAdAccountsToDSP)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAdAccountsFromDSP/{parentId}/adAccountsIds", jsonResponseFormatter.FormatToJSON(DSPController.RemoveAdAccountsFromDSP)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Campaign Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Campaign/{id}", jsonResponseFormatter.FormatToJSON(CampaignController.GetCampaign)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Campaign", jsonResponseFormatter.FormatToJSON(CampaignController.GetAllCampaign)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCampaign", jsonResponseFormatter.FormatToJSON(CampaignController.CreateCampaign)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Campaign/{id}", jsonResponseFormatter.FormatToJSON(CampaignController.UpdateCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCampaign/{id}", jsonResponseFormatter.FormatToJSON(CampaignController.DeleteCampaign)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAdAccountToCampaign/{parentId}/adAccountId", jsonResponseFormatter.FormatToJSON(CampaignController.AssignAdAccountToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdAccountFromCampaign/{parentId}", jsonResponseFormatter.FormatToJSON(CampaignController.UnassignAdAccountFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInsertionOrderToCampaign/{parentId}/insertionOrderId", jsonResponseFormatter.FormatToJSON(CampaignController.AssignInsertionOrderToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInsertionOrderFromCampaign/{parentId}", jsonResponseFormatter.FormatToJSON(CampaignController.UnassignInsertionOrderFromCampaign)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLineItemsToCampaign/{parentId}/lineItemsId", jsonResponseFormatter.FormatToJSON(CampaignController.AddLineItemsToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLineItemsFromCampaign/{parentId}/lineItemsIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveLineItemsFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddKpisToCampaign/{parentId}/kpisId", jsonResponseFormatter.FormatToJSON(CampaignController.AddKpisToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveKpisFromCampaign/{parentId}/kpisIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveKpisFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTrackingPixelsToCampaign/{parentId}/trackingPixelsId", jsonResponseFormatter.FormatToJSON(CampaignController.AddTrackingPixelsToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTrackingPixelsFromCampaign/{parentId}/trackingPixelsIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveTrackingPixelsFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAudiencesToCampaign/{parentId}/audiencesId", jsonResponseFormatter.FormatToJSON(CampaignController.AddAudiencesToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAudiencesFromCampaign/{parentId}/audiencesIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveAudiencesFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReportsToCampaign/{parentId}/reportsId", jsonResponseFormatter.FormatToJSON(CampaignController.AddReportsToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReportsFromCampaign/{parentId}/reportsIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveReportsFromCampaign)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // KPI Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/KPI/{id}", jsonResponseFormatter.FormatToJSON(KPIController.GetKPI)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/KPI", jsonResponseFormatter.FormatToJSON(KPIController.GetAllKPI)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewKPI", jsonResponseFormatter.FormatToJSON(KPIController.CreateKPI)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/KPI/{id}", jsonResponseFormatter.FormatToJSON(KPIController.UpdateKPI)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteKPI/{id}", jsonResponseFormatter.FormatToJSON(KPIController.DeleteKPI)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCampaignToKPI/{parentId}/campaignId", jsonResponseFormatter.FormatToJSON(KPIController.AssignCampaignToKPI)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCampaignFromKPI/{parentId}", jsonResponseFormatter.FormatToJSON(KPIController.UnassignCampaignFromKPI)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // AudienceSegment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AudienceSegment/{id}", jsonResponseFormatter.FormatToJSON(AudienceSegmentController.GetAudienceSegment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AudienceSegment", jsonResponseFormatter.FormatToJSON(AudienceSegmentController.GetAllAudienceSegment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAudienceSegment", jsonResponseFormatter.FormatToJSON(AudienceSegmentController.CreateAudienceSegment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AudienceSegment/{id}", jsonResponseFormatter.FormatToJSON(AudienceSegmentController.UpdateAudienceSegment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAudienceSegment/{id}", jsonResponseFormatter.FormatToJSON(AudienceSegmentController.DeleteAudienceSegment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignProviderToAudienceSegment/{parentId}/providerId", jsonResponseFormatter.FormatToJSON(AudienceSegmentController.AssignProviderToAudienceSegment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProviderFromAudienceSegment/{parentId}", jsonResponseFormatter.FormatToJSON(AudienceSegmentController.UnassignProviderFromAudienceSegment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCampaignsToAudienceSegment/{parentId}/campaignsId", jsonResponseFormatter.FormatToJSON(AudienceSegmentController.AddCampaignsToAudienceSegment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCampaignsFromAudienceSegment/{parentId}/campaignsIds", jsonResponseFormatter.FormatToJSON(AudienceSegmentController.RemoveCampaignsFromAudienceSegment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DataProvider Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DataProvider/{id}", jsonResponseFormatter.FormatToJSON(DataProviderController.GetDataProvider)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DataProvider", jsonResponseFormatter.FormatToJSON(DataProviderController.GetAllDataProvider)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDataProvider", jsonResponseFormatter.FormatToJSON(DataProviderController.CreateDataProvider)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DataProvider/{id}", jsonResponseFormatter.FormatToJSON(DataProviderController.UpdateDataProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDataProvider/{id}", jsonResponseFormatter.FormatToJSON(DataProviderController.DeleteDataProvider)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAudienceSegmentsToDataProvider/{parentId}/audienceSegmentsId", jsonResponseFormatter.FormatToJSON(DataProviderController.AddAudienceSegmentsToDataProvider)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAudienceSegmentsFromDataProvider/{parentId}/audienceSegmentsIds", jsonResponseFormatter.FormatToJSON(DataProviderController.RemoveAudienceSegmentsFromDataProvider)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // LineItem Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LineItem/{id}", jsonResponseFormatter.FormatToJSON(LineItemController.GetLineItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LineItem", jsonResponseFormatter.FormatToJSON(LineItemController.GetAllLineItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLineItem", jsonResponseFormatter.FormatToJSON(LineItemController.CreateLineItem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LineItem/{id}", jsonResponseFormatter.FormatToJSON(LineItemController.UpdateLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLineItem/{id}", jsonResponseFormatter.FormatToJSON(LineItemController.DeleteLineItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCampaignToLineItem/{parentId}/campaignId", jsonResponseFormatter.FormatToJSON(LineItemController.AssignCampaignToLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCampaignFromLineItem/{parentId}", jsonResponseFormatter.FormatToJSON(LineItemController.UnassignCampaignFromLineItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTargetingProfileToLineItem/{parentId}/targetingProfileId", jsonResponseFormatter.FormatToJSON(LineItemController.AssignTargetingProfileToLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTargetingProfileFromLineItem/{parentId}", jsonResponseFormatter.FormatToJSON(LineItemController.UnassignTargetingProfileFromLineItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDealToLineItem/{parentId}/dealId", jsonResponseFormatter.FormatToJSON(LineItemController.AssignDealToLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDealFromLineItem/{parentId}", jsonResponseFormatter.FormatToJSON(LineItemController.UnassignDealFromLineItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPlacementsToLineItem/{parentId}/placementsId", jsonResponseFormatter.FormatToJSON(LineItemController.AddPlacementsToLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePlacementsFromLineItem/{parentId}/placementsIds", jsonResponseFormatter.FormatToJSON(LineItemController.RemovePlacementsFromLineItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCreativesToLineItem/{parentId}/creativesId", jsonResponseFormatter.FormatToJSON(LineItemController.AddCreativesToLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCreativesFromLineItem/{parentId}/creativesIds", jsonResponseFormatter.FormatToJSON(LineItemController.RemoveCreativesFromLineItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPerformanceMetricsToLineItem/{parentId}/performanceMetricsId", jsonResponseFormatter.FormatToJSON(LineItemController.AddPerformanceMetricsToLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePerformanceMetricsFromLineItem/{parentId}/performanceMetricsIds", jsonResponseFormatter.FormatToJSON(LineItemController.RemovePerformanceMetricsFromLineItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // TargetingProfile Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TargetingProfile/{id}", jsonResponseFormatter.FormatToJSON(TargetingProfileController.GetTargetingProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TargetingProfile", jsonResponseFormatter.FormatToJSON(TargetingProfileController.GetAllTargetingProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTargetingProfile", jsonResponseFormatter.FormatToJSON(TargetingProfileController.CreateTargetingProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TargetingProfile/{id}", jsonResponseFormatter.FormatToJSON(TargetingProfileController.UpdateTargetingProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTargetingProfile/{id}", jsonResponseFormatter.FormatToJSON(TargetingProfileController.DeleteTargetingProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBrandSafetyPolicyToTargetingProfile/{parentId}/brandSafetyPolicyId", jsonResponseFormatter.FormatToJSON(TargetingProfileController.AssignBrandSafetyPolicyToTargetingProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBrandSafetyPolicyFromTargetingProfile/{parentId}", jsonResponseFormatter.FormatToJSON(TargetingProfileController.UnassignBrandSafetyPolicyFromTargetingProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAudienceSegmentsToTargetingProfile/{parentId}/audienceSegmentsId", jsonResponseFormatter.FormatToJSON(TargetingProfileController.AddAudienceSegmentsToTargetingProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAudienceSegmentsFromTargetingProfile/{parentId}/audienceSegmentsIds", jsonResponseFormatter.FormatToJSON(TargetingProfileController.RemoveAudienceSegmentsFromTargetingProfile)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddGeoRegionsToTargetingProfile/{parentId}/geoRegionsId", jsonResponseFormatter.FormatToJSON(TargetingProfileController.AddGeoRegionsToTargetingProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveGeoRegionsFromTargetingProfile/{parentId}/geoRegionsIds", jsonResponseFormatter.FormatToJSON(TargetingProfileController.RemoveGeoRegionsFromTargetingProfile)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddContentCategoriesToTargetingProfile/{parentId}/contentCategoriesId", jsonResponseFormatter.FormatToJSON(TargetingProfileController.AddContentCategoriesToTargetingProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContentCategoriesFromTargetingProfile/{parentId}/contentCategoriesIds", jsonResponseFormatter.FormatToJSON(TargetingProfileController.RemoveContentCategoriesFromTargetingProfile)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDeviceCriteriaToTargetingProfile/{parentId}/deviceCriteriaId", jsonResponseFormatter.FormatToJSON(TargetingProfileController.AddDeviceCriteriaToTargetingProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDeviceCriteriaFromTargetingProfile/{parentId}/deviceCriteriaIds", jsonResponseFormatter.FormatToJSON(TargetingProfileController.RemoveDeviceCriteriaFromTargetingProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DeviceCriterion Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DeviceCriterion/{id}", jsonResponseFormatter.FormatToJSON(DeviceCriterionController.GetDeviceCriterion)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DeviceCriterion", jsonResponseFormatter.FormatToJSON(DeviceCriterionController.GetAllDeviceCriterion)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDeviceCriterion", jsonResponseFormatter.FormatToJSON(DeviceCriterionController.CreateDeviceCriterion)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DeviceCriterion/{id}", jsonResponseFormatter.FormatToJSON(DeviceCriterionController.UpdateDeviceCriterion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDeviceCriterion/{id}", jsonResponseFormatter.FormatToJSON(DeviceCriterionController.DeleteDeviceCriterion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignTargetingProfileToDeviceCriterion/{parentId}/targetingProfileId", jsonResponseFormatter.FormatToJSON(DeviceCriterionController.AssignTargetingProfileToDeviceCriterion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTargetingProfileFromDeviceCriterion/{parentId}", jsonResponseFormatter.FormatToJSON(DeviceCriterionController.UnassignTargetingProfileFromDeviceCriterion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // BrandSafetyPolicy Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BrandSafetyPolicy/{id}", jsonResponseFormatter.FormatToJSON(BrandSafetyPolicyController.GetBrandSafetyPolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BrandSafetyPolicy", jsonResponseFormatter.FormatToJSON(BrandSafetyPolicyController.GetAllBrandSafetyPolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBrandSafetyPolicy", jsonResponseFormatter.FormatToJSON(BrandSafetyPolicyController.CreateBrandSafetyPolicy)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BrandSafetyPolicy/{id}", jsonResponseFormatter.FormatToJSON(BrandSafetyPolicyController.UpdateBrandSafetyPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBrandSafetyPolicy/{id}", jsonResponseFormatter.FormatToJSON(BrandSafetyPolicyController.DeleteBrandSafetyPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTargetingProfilesToBrandSafetyPolicy/{parentId}/targetingProfilesId", jsonResponseFormatter.FormatToJSON(BrandSafetyPolicyController.AddTargetingProfilesToBrandSafetyPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTargetingProfilesFromBrandSafetyPolicy/{parentId}/targetingProfilesIds", jsonResponseFormatter.FormatToJSON(BrandSafetyPolicyController.RemoveTargetingProfilesFromBrandSafetyPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ContentCategory Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ContentCategory/{id}", jsonResponseFormatter.FormatToJSON(ContentCategoryController.GetContentCategory)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ContentCategory", jsonResponseFormatter.FormatToJSON(ContentCategoryController.GetAllContentCategory)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewContentCategory", jsonResponseFormatter.FormatToJSON(ContentCategoryController.CreateContentCategory)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ContentCategory/{id}", jsonResponseFormatter.FormatToJSON(ContentCategoryController.UpdateContentCategory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteContentCategory/{id}", jsonResponseFormatter.FormatToJSON(ContentCategoryController.DeleteContentCategory)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Publisher Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Publisher/{id}", jsonResponseFormatter.FormatToJSON(PublisherController.GetPublisher)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Publisher", jsonResponseFormatter.FormatToJSON(PublisherController.GetAllPublisher)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPublisher", jsonResponseFormatter.FormatToJSON(PublisherController.CreatePublisher)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Publisher/{id}", jsonResponseFormatter.FormatToJSON(PublisherController.UpdatePublisher)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePublisher/{id}", jsonResponseFormatter.FormatToJSON(PublisherController.DeletePublisher)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInventorySourcesToPublisher/{parentId}/inventorySourcesId", jsonResponseFormatter.FormatToJSON(PublisherController.AddInventorySourcesToPublisher)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventorySourcesFromPublisher/{parentId}/inventorySourcesIds", jsonResponseFormatter.FormatToJSON(PublisherController.RemoveInventorySourcesFromPublisher)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDealsToPublisher/{parentId}/dealsId", jsonResponseFormatter.FormatToJSON(PublisherController.AddDealsToPublisher)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDealsFromPublisher/{parentId}/dealsIds", jsonResponseFormatter.FormatToJSON(PublisherController.RemoveDealsFromPublisher)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCreativeApprovalsToPublisher/{parentId}/creativeApprovalsId", jsonResponseFormatter.FormatToJSON(PublisherController.AddCreativeApprovalsToPublisher)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCreativeApprovalsFromPublisher/{parentId}/creativeApprovalsIds", jsonResponseFormatter.FormatToJSON(PublisherController.RemoveCreativeApprovalsFromPublisher)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInsertionOrdersToPublisher/{parentId}/insertionOrdersId", jsonResponseFormatter.FormatToJSON(PublisherController.AddInsertionOrdersToPublisher)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInsertionOrdersFromPublisher/{parentId}/insertionOrdersIds", jsonResponseFormatter.FormatToJSON(PublisherController.RemoveInsertionOrdersFromPublisher)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRateCardsToPublisher/{parentId}/rateCardsId", jsonResponseFormatter.FormatToJSON(PublisherController.AddRateCardsToPublisher)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRateCardsFromPublisher/{parentId}/rateCardsIds", jsonResponseFormatter.FormatToJSON(PublisherController.RemoveRateCardsFromPublisher)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InventorySource Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InventorySource/{id}", jsonResponseFormatter.FormatToJSON(InventorySourceController.GetInventorySource)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InventorySource", jsonResponseFormatter.FormatToJSON(InventorySourceController.GetAllInventorySource)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInventorySource", jsonResponseFormatter.FormatToJSON(InventorySourceController.CreateInventorySource)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InventorySource/{id}", jsonResponseFormatter.FormatToJSON(InventorySourceController.UpdateInventorySource)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInventorySource/{id}", jsonResponseFormatter.FormatToJSON(InventorySourceController.DeleteInventorySource)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPublisherToInventorySource/{parentId}/publisherId", jsonResponseFormatter.FormatToJSON(InventorySourceController.AssignPublisherToInventorySource)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPublisherFromInventorySource/{parentId}", jsonResponseFormatter.FormatToJSON(InventorySourceController.UnassignPublisherFromInventorySource)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAdSlotsToInventorySource/{parentId}/adSlotsId", jsonResponseFormatter.FormatToJSON(InventorySourceController.AddAdSlotsToInventorySource)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAdSlotsFromInventorySource/{parentId}/adSlotsIds", jsonResponseFormatter.FormatToJSON(InventorySourceController.RemoveAdSlotsFromInventorySource)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDealsToInventorySource/{parentId}/dealsId", jsonResponseFormatter.FormatToJSON(InventorySourceController.AddDealsToInventorySource)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDealsFromInventorySource/{parentId}/dealsIds", jsonResponseFormatter.FormatToJSON(InventorySourceController.RemoveDealsFromInventorySource)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AdSlot Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AdSlot/{id}", jsonResponseFormatter.FormatToJSON(AdSlotController.GetAdSlot)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AdSlot", jsonResponseFormatter.FormatToJSON(AdSlotController.GetAllAdSlot)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAdSlot", jsonResponseFormatter.FormatToJSON(AdSlotController.CreateAdSlot)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AdSlot/{id}", jsonResponseFormatter.FormatToJSON(AdSlotController.UpdateAdSlot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAdSlot/{id}", jsonResponseFormatter.FormatToJSON(AdSlotController.DeleteAdSlot)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInventorySourceToAdSlot/{parentId}/inventorySourceId", jsonResponseFormatter.FormatToJSON(AdSlotController.AssignInventorySourceToAdSlot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInventorySourceFromAdSlot/{parentId}", jsonResponseFormatter.FormatToJSON(AdSlotController.UnassignInventorySourceFromAdSlot)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPlacementsToAdSlot/{parentId}/placementsId", jsonResponseFormatter.FormatToJSON(AdSlotController.AddPlacementsToAdSlot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePlacementsFromAdSlot/{parentId}/placementsIds", jsonResponseFormatter.FormatToJSON(AdSlotController.RemovePlacementsFromAdSlot)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRatesToAdSlot/{parentId}/ratesId", jsonResponseFormatter.FormatToJSON(AdSlotController.AddRatesToAdSlot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRatesFromAdSlot/{parentId}/ratesIds", jsonResponseFormatter.FormatToJSON(AdSlotController.RemoveRatesFromAdSlot)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Deal Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Deal/{id}", jsonResponseFormatter.FormatToJSON(DealController.GetDeal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Deal", jsonResponseFormatter.FormatToJSON(DealController.GetAllDeal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDeal", jsonResponseFormatter.FormatToJSON(DealController.CreateDeal)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Deal/{id}", jsonResponseFormatter.FormatToJSON(DealController.UpdateDeal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDeal/{id}", jsonResponseFormatter.FormatToJSON(DealController.DeleteDeal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPublisherToDeal/{parentId}/publisherId", jsonResponseFormatter.FormatToJSON(DealController.AssignPublisherToDeal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPublisherFromDeal/{parentId}", jsonResponseFormatter.FormatToJSON(DealController.UnassignPublisherFromDeal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInventorySourcesToDeal/{parentId}/inventorySourcesId", jsonResponseFormatter.FormatToJSON(DealController.AddInventorySourcesToDeal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventorySourcesFromDeal/{parentId}/inventorySourcesIds", jsonResponseFormatter.FormatToJSON(DealController.RemoveInventorySourcesFromDeal)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPlacementsToDeal/{parentId}/placementsId", jsonResponseFormatter.FormatToJSON(DealController.AddPlacementsToDeal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePlacementsFromDeal/{parentId}/placementsIds", jsonResponseFormatter.FormatToJSON(DealController.RemovePlacementsFromDeal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Placement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Placement/{id}", jsonResponseFormatter.FormatToJSON(PlacementController.GetPlacement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Placement", jsonResponseFormatter.FormatToJSON(PlacementController.GetAllPlacement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPlacement", jsonResponseFormatter.FormatToJSON(PlacementController.CreatePlacement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Placement/{id}", jsonResponseFormatter.FormatToJSON(PlacementController.UpdatePlacement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePlacement/{id}", jsonResponseFormatter.FormatToJSON(PlacementController.DeletePlacement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignLineItemToPlacement/{parentId}/lineItemId", jsonResponseFormatter.FormatToJSON(PlacementController.AssignLineItemToPlacement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLineItemFromPlacement/{parentId}", jsonResponseFormatter.FormatToJSON(PlacementController.UnassignLineItemFromPlacement)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAdSlotToPlacement/{parentId}/adSlotId", jsonResponseFormatter.FormatToJSON(PlacementController.AssignAdSlotToPlacement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdSlotFromPlacement/{parentId}", jsonResponseFormatter.FormatToJSON(PlacementController.UnassignAdSlotFromPlacement)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDealToPlacement/{parentId}/dealId", jsonResponseFormatter.FormatToJSON(PlacementController.AssignDealToPlacement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDealFromPlacement/{parentId}", jsonResponseFormatter.FormatToJSON(PlacementController.UnassignDealFromPlacement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // CreativeAsset Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CreativeAsset/{id}", jsonResponseFormatter.FormatToJSON(CreativeAssetController.GetCreativeAsset)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CreativeAsset", jsonResponseFormatter.FormatToJSON(CreativeAssetController.GetAllCreativeAsset)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCreativeAsset", jsonResponseFormatter.FormatToJSON(CreativeAssetController.CreateCreativeAsset)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CreativeAsset/{id}", jsonResponseFormatter.FormatToJSON(CreativeAssetController.UpdateCreativeAsset)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCreativeAsset/{id}", jsonResponseFormatter.FormatToJSON(CreativeAssetController.DeleteCreativeAsset)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddFilesToCreativeAsset/{parentId}/filesId", jsonResponseFormatter.FormatToJSON(CreativeAssetController.AddFilesToCreativeAsset)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFilesFromCreativeAsset/{parentId}/filesIds", jsonResponseFormatter.FormatToJSON(CreativeAssetController.RemoveFilesFromCreativeAsset)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddApprovalsToCreativeAsset/{parentId}/approvalsId", jsonResponseFormatter.FormatToJSON(CreativeAssetController.AddApprovalsToCreativeAsset)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveApprovalsFromCreativeAsset/{parentId}/approvalsIds", jsonResponseFormatter.FormatToJSON(CreativeAssetController.RemoveApprovalsFromCreativeAsset)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddVariationsToCreativeAsset/{parentId}/variationsId", jsonResponseFormatter.FormatToJSON(CreativeAssetController.AddVariationsToCreativeAsset)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVariationsFromCreativeAsset/{parentId}/variationsIds", jsonResponseFormatter.FormatToJSON(CreativeAssetController.RemoveVariationsFromCreativeAsset)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLineItemsToCreativeAsset/{parentId}/lineItemsId", jsonResponseFormatter.FormatToJSON(CreativeAssetController.AddLineItemsToCreativeAsset)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLineItemsFromCreativeAsset/{parentId}/lineItemsIds", jsonResponseFormatter.FormatToJSON(CreativeAssetController.RemoveLineItemsFromCreativeAsset)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CreativeFile Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CreativeFile/{id}", jsonResponseFormatter.FormatToJSON(CreativeFileController.GetCreativeFile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CreativeFile", jsonResponseFormatter.FormatToJSON(CreativeFileController.GetAllCreativeFile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCreativeFile", jsonResponseFormatter.FormatToJSON(CreativeFileController.CreateCreativeFile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CreativeFile/{id}", jsonResponseFormatter.FormatToJSON(CreativeFileController.UpdateCreativeFile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCreativeFile/{id}", jsonResponseFormatter.FormatToJSON(CreativeFileController.DeleteCreativeFile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCreativeAssetToCreativeFile/{parentId}/creativeAssetId", jsonResponseFormatter.FormatToJSON(CreativeFileController.AssignCreativeAssetToCreativeFile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCreativeAssetFromCreativeFile/{parentId}", jsonResponseFormatter.FormatToJSON(CreativeFileController.UnassignCreativeAssetFromCreativeFile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // CreativeVariation Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CreativeVariation/{id}", jsonResponseFormatter.FormatToJSON(CreativeVariationController.GetCreativeVariation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CreativeVariation", jsonResponseFormatter.FormatToJSON(CreativeVariationController.GetAllCreativeVariation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCreativeVariation", jsonResponseFormatter.FormatToJSON(CreativeVariationController.CreateCreativeVariation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CreativeVariation/{id}", jsonResponseFormatter.FormatToJSON(CreativeVariationController.UpdateCreativeVariation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCreativeVariation/{id}", jsonResponseFormatter.FormatToJSON(CreativeVariationController.DeleteCreativeVariation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCreativeAssetToCreativeVariation/{parentId}/creativeAssetId", jsonResponseFormatter.FormatToJSON(CreativeVariationController.AssignCreativeAssetToCreativeVariation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCreativeAssetFromCreativeVariation/{parentId}", jsonResponseFormatter.FormatToJSON(CreativeVariationController.UnassignCreativeAssetFromCreativeVariation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // CreativeApproval Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CreativeApproval/{id}", jsonResponseFormatter.FormatToJSON(CreativeApprovalController.GetCreativeApproval)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CreativeApproval", jsonResponseFormatter.FormatToJSON(CreativeApprovalController.GetAllCreativeApproval)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCreativeApproval", jsonResponseFormatter.FormatToJSON(CreativeApprovalController.CreateCreativeApproval)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CreativeApproval/{id}", jsonResponseFormatter.FormatToJSON(CreativeApprovalController.UpdateCreativeApproval)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCreativeApproval/{id}", jsonResponseFormatter.FormatToJSON(CreativeApprovalController.DeleteCreativeApproval)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCreativeAssetToCreativeApproval/{parentId}/creativeAssetId", jsonResponseFormatter.FormatToJSON(CreativeApprovalController.AssignCreativeAssetToCreativeApproval)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCreativeAssetFromCreativeApproval/{parentId}", jsonResponseFormatter.FormatToJSON(CreativeApprovalController.UnassignCreativeAssetFromCreativeApproval)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPublisherToCreativeApproval/{parentId}/publisherId", jsonResponseFormatter.FormatToJSON(CreativeApprovalController.AssignPublisherToCreativeApproval)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPublisherFromCreativeApproval/{parentId}", jsonResponseFormatter.FormatToJSON(CreativeApprovalController.UnassignPublisherFromCreativeApproval)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // TrackingPixel Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TrackingPixel/{id}", jsonResponseFormatter.FormatToJSON(TrackingPixelController.GetTrackingPixel)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TrackingPixel", jsonResponseFormatter.FormatToJSON(TrackingPixelController.GetAllTrackingPixel)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTrackingPixel", jsonResponseFormatter.FormatToJSON(TrackingPixelController.CreateTrackingPixel)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TrackingPixel/{id}", jsonResponseFormatter.FormatToJSON(TrackingPixelController.UpdateTrackingPixel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTrackingPixel/{id}", jsonResponseFormatter.FormatToJSON(TrackingPixelController.DeleteTrackingPixel)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCampaignToTrackingPixel/{parentId}/campaignId", jsonResponseFormatter.FormatToJSON(TrackingPixelController.AssignCampaignToTrackingPixel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCampaignFromTrackingPixel/{parentId}", jsonResponseFormatter.FormatToJSON(TrackingPixelController.UnassignCampaignFromTrackingPixel)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAdvertiserToTrackingPixel/{parentId}/advertiserId", jsonResponseFormatter.FormatToJSON(TrackingPixelController.AssignAdvertiserToTrackingPixel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdvertiserFromTrackingPixel/{parentId}", jsonResponseFormatter.FormatToJSON(TrackingPixelController.UnassignAdvertiserFromTrackingPixel)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddConversionEventsToTrackingPixel/{parentId}/conversionEventsId", jsonResponseFormatter.FormatToJSON(TrackingPixelController.AddConversionEventsToTrackingPixel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveConversionEventsFromTrackingPixel/{parentId}/conversionEventsIds", jsonResponseFormatter.FormatToJSON(TrackingPixelController.RemoveConversionEventsFromTrackingPixel)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ConversionEvent Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ConversionEvent/{id}", jsonResponseFormatter.FormatToJSON(ConversionEventController.GetConversionEvent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ConversionEvent", jsonResponseFormatter.FormatToJSON(ConversionEventController.GetAllConversionEvent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewConversionEvent", jsonResponseFormatter.FormatToJSON(ConversionEventController.CreateConversionEvent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ConversionEvent/{id}", jsonResponseFormatter.FormatToJSON(ConversionEventController.UpdateConversionEvent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteConversionEvent/{id}", jsonResponseFormatter.FormatToJSON(ConversionEventController.DeleteConversionEvent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCampaignToConversionEvent/{parentId}/campaignId", jsonResponseFormatter.FormatToJSON(ConversionEventController.AssignCampaignToConversionEvent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCampaignFromConversionEvent/{parentId}", jsonResponseFormatter.FormatToJSON(ConversionEventController.UnassignCampaignFromConversionEvent)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLineItemToConversionEvent/{parentId}/lineItemId", jsonResponseFormatter.FormatToJSON(ConversionEventController.AssignLineItemToConversionEvent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLineItemFromConversionEvent/{parentId}", jsonResponseFormatter.FormatToJSON(ConversionEventController.UnassignLineItemFromConversionEvent)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTrackingPixelToConversionEvent/{parentId}/trackingPixelId", jsonResponseFormatter.FormatToJSON(ConversionEventController.AssignTrackingPixelToConversionEvent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTrackingPixelFromConversionEvent/{parentId}", jsonResponseFormatter.FormatToJSON(ConversionEventController.UnassignTrackingPixelFromConversionEvent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // PerformanceMetric Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PerformanceMetric/{id}", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.GetPerformanceMetric)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PerformanceMetric", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.GetAllPerformanceMetric)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPerformanceMetric", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.CreatePerformanceMetric)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PerformanceMetric/{id}", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.UpdatePerformanceMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePerformanceMetric/{id}", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.DeletePerformanceMetric)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAdAccountToPerformanceMetric/{parentId}/adAccountId", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.AssignAdAccountToPerformanceMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdAccountFromPerformanceMetric/{parentId}", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.UnassignAdAccountFromPerformanceMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCampaignToPerformanceMetric/{parentId}/campaignId", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.AssignCampaignToPerformanceMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCampaignFromPerformanceMetric/{parentId}", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.UnassignCampaignFromPerformanceMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLineItemToPerformanceMetric/{parentId}/lineItemId", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.AssignLineItemToPerformanceMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLineItemFromPerformanceMetric/{parentId}", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.UnassignLineItemFromPerformanceMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPlacementToPerformanceMetric/{parentId}/placementId", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.AssignPlacementToPerformanceMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlacementFromPerformanceMetric/{parentId}", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.UnassignPlacementFromPerformanceMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCreativeAssetToPerformanceMetric/{parentId}/creativeAssetId", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.AssignCreativeAssetToPerformanceMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCreativeAssetFromPerformanceMetric/{parentId}", jsonResponseFormatter.FormatToJSON(PerformanceMetricController.UnassignCreativeAssetFromPerformanceMetric)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Report Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Report/{id}", jsonResponseFormatter.FormatToJSON(ReportController.GetReport)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Report", jsonResponseFormatter.FormatToJSON(ReportController.GetAllReport)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewReport", jsonResponseFormatter.FormatToJSON(ReportController.CreateReport)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Report/{id}", jsonResponseFormatter.FormatToJSON(ReportController.UpdateReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteReport/{id}", jsonResponseFormatter.FormatToJSON(ReportController.DeleteReport)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAdAccountToReport/{parentId}/adAccountId", jsonResponseFormatter.FormatToJSON(ReportController.AssignAdAccountToReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdAccountFromReport/{parentId}", jsonResponseFormatter.FormatToJSON(ReportController.UnassignAdAccountFromReport)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCampaignToReport/{parentId}/campaignId", jsonResponseFormatter.FormatToJSON(ReportController.AssignCampaignToReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCampaignFromReport/{parentId}", jsonResponseFormatter.FormatToJSON(ReportController.UnassignCampaignFromReport)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLineItemToReport/{parentId}/lineItemId", jsonResponseFormatter.FormatToJSON(ReportController.AssignLineItemToReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLineItemFromReport/{parentId}", jsonResponseFormatter.FormatToJSON(ReportController.UnassignLineItemFromReport)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // InsertionOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InsertionOrder/{id}", jsonResponseFormatter.FormatToJSON(InsertionOrderController.GetInsertionOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InsertionOrder", jsonResponseFormatter.FormatToJSON(InsertionOrderController.GetAllInsertionOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInsertionOrder", jsonResponseFormatter.FormatToJSON(InsertionOrderController.CreateInsertionOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InsertionOrder/{id}", jsonResponseFormatter.FormatToJSON(InsertionOrderController.UpdateInsertionOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInsertionOrder/{id}", jsonResponseFormatter.FormatToJSON(InsertionOrderController.DeleteInsertionOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAdvertiserToInsertionOrder/{parentId}/advertiserId", jsonResponseFormatter.FormatToJSON(InsertionOrderController.AssignAdvertiserToInsertionOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdvertiserFromInsertionOrder/{parentId}", jsonResponseFormatter.FormatToJSON(InsertionOrderController.UnassignAdvertiserFromInsertionOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAgencyToInsertionOrder/{parentId}/agencyId", jsonResponseFormatter.FormatToJSON(InsertionOrderController.AssignAgencyToInsertionOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAgencyFromInsertionOrder/{parentId}", jsonResponseFormatter.FormatToJSON(InsertionOrderController.UnassignAgencyFromInsertionOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPublisherToInsertionOrder/{parentId}/publisherId", jsonResponseFormatter.FormatToJSON(InsertionOrderController.AssignPublisherToInsertionOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPublisherFromInsertionOrder/{parentId}", jsonResponseFormatter.FormatToJSON(InsertionOrderController.UnassignPublisherFromInsertionOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCampaignsToInsertionOrder/{parentId}/campaignsId", jsonResponseFormatter.FormatToJSON(InsertionOrderController.AddCampaignsToInsertionOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCampaignsFromInsertionOrder/{parentId}/campaignsIds", jsonResponseFormatter.FormatToJSON(InsertionOrderController.RemoveCampaignsFromInsertionOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // RateCard Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RateCard/{id}", jsonResponseFormatter.FormatToJSON(RateCardController.GetRateCard)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RateCard", jsonResponseFormatter.FormatToJSON(RateCardController.GetAllRateCard)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRateCard", jsonResponseFormatter.FormatToJSON(RateCardController.CreateRateCard)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RateCard/{id}", jsonResponseFormatter.FormatToJSON(RateCardController.UpdateRateCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRateCard/{id}", jsonResponseFormatter.FormatToJSON(RateCardController.DeleteRateCard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPublisherToRateCard/{parentId}/publisherId", jsonResponseFormatter.FormatToJSON(RateCardController.AssignPublisherToRateCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPublisherFromRateCard/{parentId}", jsonResponseFormatter.FormatToJSON(RateCardController.UnassignPublisherFromRateCard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddRatesToRateCard/{parentId}/ratesId", jsonResponseFormatter.FormatToJSON(RateCardController.AddRatesToRateCard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRatesFromRateCard/{parentId}/ratesIds", jsonResponseFormatter.FormatToJSON(RateCardController.RemoveRatesFromRateCard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Rate Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Rate/{id}", jsonResponseFormatter.FormatToJSON(RateController.GetRate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Rate", jsonResponseFormatter.FormatToJSON(RateController.GetAllRate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRate", jsonResponseFormatter.FormatToJSON(RateController.CreateRate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Rate/{id}", jsonResponseFormatter.FormatToJSON(RateController.UpdateRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRate/{id}", jsonResponseFormatter.FormatToJSON(RateController.DeleteRate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRateCardToRate/{parentId}/rateCardId", jsonResponseFormatter.FormatToJSON(RateController.AssignRateCardToRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRateCardFromRate/{parentId}", jsonResponseFormatter.FormatToJSON(RateController.UnassignRateCardFromRate)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAdSlotToRate/{parentId}/adSlotId", jsonResponseFormatter.FormatToJSON(RateController.AssignAdSlotToRate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdSlotFromRate/{parentId}", jsonResponseFormatter.FormatToJSON(RateController.UnassignAdSlotFromRate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Experiment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Experiment/{id}", jsonResponseFormatter.FormatToJSON(ExperimentController.GetExperiment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Experiment", jsonResponseFormatter.FormatToJSON(ExperimentController.GetAllExperiment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewExperiment", jsonResponseFormatter.FormatToJSON(ExperimentController.CreateExperiment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Experiment/{id}", jsonResponseFormatter.FormatToJSON(ExperimentController.UpdateExperiment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteExperiment/{id}", jsonResponseFormatter.FormatToJSON(ExperimentController.DeleteExperiment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCampaignToExperiment/{parentId}/campaignId", jsonResponseFormatter.FormatToJSON(ExperimentController.AssignCampaignToExperiment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCampaignFromExperiment/{parentId}", jsonResponseFormatter.FormatToJSON(ExperimentController.UnassignCampaignFromExperiment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddVariantsToExperiment/{parentId}/variantsId", jsonResponseFormatter.FormatToJSON(ExperimentController.AddVariantsToExperiment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVariantsFromExperiment/{parentId}/variantsIds", jsonResponseFormatter.FormatToJSON(ExperimentController.RemoveVariantsFromExperiment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ExperimentVariant Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExperimentVariant/{id}", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.GetExperimentVariant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ExperimentVariant", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.GetAllExperimentVariant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewExperimentVariant", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.CreateExperimentVariant)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExperimentVariant/{id}", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.UpdateExperimentVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteExperimentVariant/{id}", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.DeleteExperimentVariant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignExperimentToExperimentVariant/{parentId}/experimentId", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.AssignExperimentToExperimentVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignExperimentFromExperimentVariant/{parentId}", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.UnassignExperimentFromExperimentVariant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCreativeVariationToExperimentVariant/{parentId}/creativeVariationId", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.AssignCreativeVariationToExperimentVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCreativeVariationFromExperimentVariant/{parentId}", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.UnassignCreativeVariationFromExperimentVariant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLineItemToExperimentVariant/{parentId}/lineItemId", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.AssignLineItemToExperimentVariant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLineItemFromExperimentVariant/{parentId}", jsonResponseFormatter.FormatToJSON(ExperimentVariantController.UnassignLineItemFromExperimentVariant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // GeoRegion Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/GeoRegion/{id}", jsonResponseFormatter.FormatToJSON(GeoRegionController.GetGeoRegion)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/GeoRegion", jsonResponseFormatter.FormatToJSON(GeoRegionController.GetAllGeoRegion)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewGeoRegion", jsonResponseFormatter.FormatToJSON(GeoRegionController.CreateGeoRegion)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/GeoRegion/{id}", jsonResponseFormatter.FormatToJSON(GeoRegionController.UpdateGeoRegion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteGeoRegion/{id}", jsonResponseFormatter.FormatToJSON(GeoRegionController.DeleteGeoRegion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignParentToGeoRegion/{parentId}/parentId", jsonResponseFormatter.FormatToJSON(GeoRegionController.AssignParentToGeoRegion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignParentFromGeoRegion/{parentId}", jsonResponseFormatter.FormatToJSON(GeoRegionController.UnassignParentFromGeoRegion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddChildrenToGeoRegion/{parentId}/childrenId", jsonResponseFormatter.FormatToJSON(GeoRegionController.AddChildrenToGeoRegion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveChildrenFromGeoRegion/{parentId}/childrenIds", jsonResponseFormatter.FormatToJSON(GeoRegionController.RemoveChildrenFromGeoRegion)).Methods("DELETE", "OPTIONS")

    return router
}
