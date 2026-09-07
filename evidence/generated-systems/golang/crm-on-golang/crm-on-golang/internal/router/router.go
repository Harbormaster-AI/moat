package router

import (

    OrganizationController "crm-on-golang/internal/controller"
    UserController "crm-on-golang/internal/controller"
    TeamController "crm-on-golang/internal/controller"
    TerritoryController "crm-on-golang/internal/controller"
    AccountController "crm-on-golang/internal/controller"
    ContactController "crm-on-golang/internal/controller"
    LeadController "crm-on-golang/internal/controller"
    OpportunityController "crm-on-golang/internal/controller"
    OpportunityLineItemController "crm-on-golang/internal/controller"
    OpportunityStageHistoryController "crm-on-golang/internal/controller"
    ProductController "crm-on-golang/internal/controller"
    PriceBookController "crm-on-golang/internal/controller"
    PriceBookEntryController "crm-on-golang/internal/controller"
    QuoteController "crm-on-golang/internal/controller"
    QuoteLineItemController "crm-on-golang/internal/controller"
    OrderController "crm-on-golang/internal/controller"
    OrderItemController "crm-on-golang/internal/controller"
    ContractController "crm-on-golang/internal/controller"
    Case_Controller "crm-on-golang/internal/controller"
    ActivityController "crm-on-golang/internal/controller"
    CampaignController "crm-on-golang/internal/controller"
    CampaignMemberController "crm-on-golang/internal/controller"
    NoteController "crm-on-golang/internal/controller"
    EmailMessageController "crm-on-golang/internal/controller"
    jsonResponseFormatter "crm-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "crm-on-golang/internal/controller"

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
    // Organization Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Organization/{id}", jsonResponseFormatter.FormatToJSON(OrganizationController.GetOrganization)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Organization", jsonResponseFormatter.FormatToJSON(OrganizationController.GetAllOrganization)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOrganization", jsonResponseFormatter.FormatToJSON(OrganizationController.CreateOrganization)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Organization/{id}", jsonResponseFormatter.FormatToJSON(OrganizationController.UpdateOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOrganization/{id}", jsonResponseFormatter.FormatToJSON(OrganizationController.DeleteOrganization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddUsersToOrganization/{parentId}/usersId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddUsersToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveUsersFromOrganization/{parentId}/usersIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveUsersFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAccountsToOrganization/{parentId}/accountsId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddAccountsToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAccountsFromOrganization/{parentId}/accountsIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveAccountsFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTeamsToOrganization/{parentId}/teamsId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddTeamsToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTeamsFromOrganization/{parentId}/teamsIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveTeamsFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTerritoriesToOrganization/{parentId}/territoriesId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddTerritoriesToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTerritoriesFromOrganization/{parentId}/territoriesIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveTerritoriesFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddProductsToOrganization/{parentId}/productsId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddProductsToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProductsFromOrganization/{parentId}/productsIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveProductsFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPriceBooksToOrganization/{parentId}/priceBooksId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddPriceBooksToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePriceBooksFromOrganization/{parentId}/priceBooksIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemovePriceBooksFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCampaignsToOrganization/{parentId}/campaignsId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddCampaignsToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCampaignsFromOrganization/{parentId}/campaignsIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveCampaignsFromOrganization)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AssignOrganizationToUser/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(UserController.AssignOrganizationToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromUser/{parentId}", jsonResponseFormatter.FormatToJSON(UserController.UnassignOrganizationFromUser)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTeamsToUser/{parentId}/teamsId", jsonResponseFormatter.FormatToJSON(UserController.AddTeamsToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTeamsFromUser/{parentId}/teamsIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveTeamsFromUser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddActivitiesToUser/{parentId}/activitiesId", jsonResponseFormatter.FormatToJSON(UserController.AddActivitiesToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveActivitiesFromUser/{parentId}/activitiesIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveActivitiesFromUser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOwnedAccountsToUser/{parentId}/ownedAccountsId", jsonResponseFormatter.FormatToJSON(UserController.AddOwnedAccountsToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOwnedAccountsFromUser/{parentId}/ownedAccountsIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveOwnedAccountsFromUser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOwnedLeadsToUser/{parentId}/ownedLeadsId", jsonResponseFormatter.FormatToJSON(UserController.AddOwnedLeadsToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOwnedLeadsFromUser/{parentId}/ownedLeadsIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveOwnedLeadsFromUser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOwnedOpportunitiesToUser/{parentId}/ownedOpportunitiesId", jsonResponseFormatter.FormatToJSON(UserController.AddOwnedOpportunitiesToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOwnedOpportunitiesFromUser/{parentId}/ownedOpportunitiesIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveOwnedOpportunitiesFromUser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOwnedCasesToUser/{parentId}/ownedCasesId", jsonResponseFormatter.FormatToJSON(UserController.AddOwnedCasesToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOwnedCasesFromUser/{parentId}/ownedCasesIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveOwnedCasesFromUser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddQuotesToUser/{parentId}/quotesId", jsonResponseFormatter.FormatToJSON(UserController.AddQuotesToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQuotesFromUser/{parentId}/quotesIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveQuotesFromUser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOrdersToUser/{parentId}/ordersId", jsonResponseFormatter.FormatToJSON(UserController.AddOrdersToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOrdersFromUser/{parentId}/ordersIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveOrdersFromUser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddContractsToUser/{parentId}/contractsId", jsonResponseFormatter.FormatToJSON(UserController.AddContractsToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContractsFromUser/{parentId}/contractsIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveContractsFromUser)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEmailMessagesToUser/{parentId}/emailMessagesId", jsonResponseFormatter.FormatToJSON(UserController.AddEmailMessagesToUser)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEmailMessagesFromUser/{parentId}/emailMessagesIds", jsonResponseFormatter.FormatToJSON(UserController.RemoveEmailMessagesFromUser)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AssignOrganizationToTeam/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(TeamController.AssignOrganizationToTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromTeam/{parentId}", jsonResponseFormatter.FormatToJSON(TeamController.UnassignOrganizationFromTeam)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddUsersToTeam/{parentId}/usersId", jsonResponseFormatter.FormatToJSON(TeamController.AddUsersToTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveUsersFromTeam/{parentId}/usersIds", jsonResponseFormatter.FormatToJSON(TeamController.RemoveUsersFromTeam)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAccountsToTeam/{parentId}/accountsId", jsonResponseFormatter.FormatToJSON(TeamController.AddAccountsToTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAccountsFromTeam/{parentId}/accountsIds", jsonResponseFormatter.FormatToJSON(TeamController.RemoveAccountsFromTeam)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOpportunitiesToTeam/{parentId}/opportunitiesId", jsonResponseFormatter.FormatToJSON(TeamController.AddOpportunitiesToTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOpportunitiesFromTeam/{parentId}/opportunitiesIds", jsonResponseFormatter.FormatToJSON(TeamController.RemoveOpportunitiesFromTeam)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCasesToTeam/{parentId}/casesId", jsonResponseFormatter.FormatToJSON(TeamController.AddCasesToTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCasesFromTeam/{parentId}/casesIds", jsonResponseFormatter.FormatToJSON(TeamController.RemoveCasesFromTeam)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCampaignsToTeam/{parentId}/campaignsId", jsonResponseFormatter.FormatToJSON(TeamController.AddCampaignsToTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCampaignsFromTeam/{parentId}/campaignsIds", jsonResponseFormatter.FormatToJSON(TeamController.RemoveCampaignsFromTeam)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Territory Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Territory/{id}", jsonResponseFormatter.FormatToJSON(TerritoryController.GetTerritory)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Territory", jsonResponseFormatter.FormatToJSON(TerritoryController.GetAllTerritory)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTerritory", jsonResponseFormatter.FormatToJSON(TerritoryController.CreateTerritory)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Territory/{id}", jsonResponseFormatter.FormatToJSON(TerritoryController.UpdateTerritory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTerritory/{id}", jsonResponseFormatter.FormatToJSON(TerritoryController.DeleteTerritory)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToTerritory/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(TerritoryController.AssignOrganizationToTerritory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromTerritory/{parentId}", jsonResponseFormatter.FormatToJSON(TerritoryController.UnassignOrganizationFromTerritory)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAccountsToTerritory/{parentId}/accountsId", jsonResponseFormatter.FormatToJSON(TerritoryController.AddAccountsToTerritory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAccountsFromTerritory/{parentId}/accountsIds", jsonResponseFormatter.FormatToJSON(TerritoryController.RemoveAccountsFromTerritory)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddUsersToTerritory/{parentId}/usersId", jsonResponseFormatter.FormatToJSON(TerritoryController.AddUsersToTerritory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveUsersFromTerritory/{parentId}/usersIds", jsonResponseFormatter.FormatToJSON(TerritoryController.RemoveUsersFromTerritory)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AssignOrganizationToAccount/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(AccountController.AssignOrganizationToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AccountController.UnassignOrganizationFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignParentAccountToAccount/{parentId}/parentAccountId", jsonResponseFormatter.FormatToJSON(AccountController.AssignParentAccountToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignParentAccountFromAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AccountController.UnassignParentAccountFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToAccount/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(AccountController.AssignOwnerToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AccountController.UnassignOwnerFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTerritoryToAccount/{parentId}/territoryId", jsonResponseFormatter.FormatToJSON(AccountController.AssignTerritoryToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTerritoryFromAccount/{parentId}", jsonResponseFormatter.FormatToJSON(AccountController.UnassignTerritoryFromAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddChildAccountsToAccount/{parentId}/childAccountsId", jsonResponseFormatter.FormatToJSON(AccountController.AddChildAccountsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveChildAccountsFromAccount/{parentId}/childAccountsIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveChildAccountsFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddContactsToAccount/{parentId}/contactsId", jsonResponseFormatter.FormatToJSON(AccountController.AddContactsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContactsFromAccount/{parentId}/contactsIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveContactsFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOpportunitiesToAccount/{parentId}/opportunitiesId", jsonResponseFormatter.FormatToJSON(AccountController.AddOpportunitiesToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOpportunitiesFromAccount/{parentId}/opportunitiesIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveOpportunitiesFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCasesToAccount/{parentId}/casesId", jsonResponseFormatter.FormatToJSON(AccountController.AddCasesToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCasesFromAccount/{parentId}/casesIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveCasesFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddActivitiesToAccount/{parentId}/activitiesId", jsonResponseFormatter.FormatToJSON(AccountController.AddActivitiesToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveActivitiesFromAccount/{parentId}/activitiesIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveActivitiesFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCampaignsToAccount/{parentId}/campaignsId", jsonResponseFormatter.FormatToJSON(AccountController.AddCampaignsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCampaignsFromAccount/{parentId}/campaignsIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveCampaignsFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddQuotesToAccount/{parentId}/quotesId", jsonResponseFormatter.FormatToJSON(AccountController.AddQuotesToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQuotesFromAccount/{parentId}/quotesIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveQuotesFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOrdersToAccount/{parentId}/ordersId", jsonResponseFormatter.FormatToJSON(AccountController.AddOrdersToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOrdersFromAccount/{parentId}/ordersIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveOrdersFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddContractsToAccount/{parentId}/contractsId", jsonResponseFormatter.FormatToJSON(AccountController.AddContractsToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContractsFromAccount/{parentId}/contractsIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveContractsFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddNotesToAccount/{parentId}/notesId", jsonResponseFormatter.FormatToJSON(AccountController.AddNotesToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveNotesFromAccount/{parentId}/notesIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveNotesFromAccount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEmailMessagesToAccount/{parentId}/emailMessagesId", jsonResponseFormatter.FormatToJSON(AccountController.AddEmailMessagesToAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEmailMessagesFromAccount/{parentId}/emailMessagesIds", jsonResponseFormatter.FormatToJSON(AccountController.RemoveEmailMessagesFromAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Contact Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Contact/{id}", jsonResponseFormatter.FormatToJSON(ContactController.GetContact)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Contact", jsonResponseFormatter.FormatToJSON(ContactController.GetAllContact)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewContact", jsonResponseFormatter.FormatToJSON(ContactController.CreateContact)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Contact/{id}", jsonResponseFormatter.FormatToJSON(ContactController.UpdateContact)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteContact/{id}", jsonResponseFormatter.FormatToJSON(ContactController.DeleteContact)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToContact/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(ContactController.AssignOrganizationToContact)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromContact/{parentId}", jsonResponseFormatter.FormatToJSON(ContactController.UnassignOrganizationFromContact)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToContact/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(ContactController.AssignAccountToContact)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromContact/{parentId}", jsonResponseFormatter.FormatToJSON(ContactController.UnassignAccountFromContact)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToContact/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(ContactController.AssignOwnerToContact)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromContact/{parentId}", jsonResponseFormatter.FormatToJSON(ContactController.UnassignOwnerFromContact)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddActivitiesToContact/{parentId}/activitiesId", jsonResponseFormatter.FormatToJSON(ContactController.AddActivitiesToContact)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveActivitiesFromContact/{parentId}/activitiesIds", jsonResponseFormatter.FormatToJSON(ContactController.RemoveActivitiesFromContact)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOpportunitiesToContact/{parentId}/opportunitiesId", jsonResponseFormatter.FormatToJSON(ContactController.AddOpportunitiesToContact)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOpportunitiesFromContact/{parentId}/opportunitiesIds", jsonResponseFormatter.FormatToJSON(ContactController.RemoveOpportunitiesFromContact)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCasesToContact/{parentId}/casesId", jsonResponseFormatter.FormatToJSON(ContactController.AddCasesToContact)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCasesFromContact/{parentId}/casesIds", jsonResponseFormatter.FormatToJSON(ContactController.RemoveCasesFromContact)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCampaignsToContact/{parentId}/campaignsId", jsonResponseFormatter.FormatToJSON(ContactController.AddCampaignsToContact)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCampaignsFromContact/{parentId}/campaignsIds", jsonResponseFormatter.FormatToJSON(ContactController.RemoveCampaignsFromContact)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddNotesToContact/{parentId}/notesId", jsonResponseFormatter.FormatToJSON(ContactController.AddNotesToContact)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveNotesFromContact/{parentId}/notesIds", jsonResponseFormatter.FormatToJSON(ContactController.RemoveNotesFromContact)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEmailMessagesToContact/{parentId}/emailMessagesId", jsonResponseFormatter.FormatToJSON(ContactController.AddEmailMessagesToContact)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEmailMessagesFromContact/{parentId}/emailMessagesIds", jsonResponseFormatter.FormatToJSON(ContactController.RemoveEmailMessagesFromContact)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Lead Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Lead/{id}", jsonResponseFormatter.FormatToJSON(LeadController.GetLead)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Lead", jsonResponseFormatter.FormatToJSON(LeadController.GetAllLead)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLead", jsonResponseFormatter.FormatToJSON(LeadController.CreateLead)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Lead/{id}", jsonResponseFormatter.FormatToJSON(LeadController.UpdateLead)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLead/{id}", jsonResponseFormatter.FormatToJSON(LeadController.DeleteLead)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToLead/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(LeadController.AssignOrganizationToLead)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromLead/{parentId}", jsonResponseFormatter.FormatToJSON(LeadController.UnassignOrganizationFromLead)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToLead/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(LeadController.AssignOwnerToLead)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromLead/{parentId}", jsonResponseFormatter.FormatToJSON(LeadController.UnassignOwnerFromLead)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignConvertedAccountToLead/{parentId}/convertedAccountId", jsonResponseFormatter.FormatToJSON(LeadController.AssignConvertedAccountToLead)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignConvertedAccountFromLead/{parentId}", jsonResponseFormatter.FormatToJSON(LeadController.UnassignConvertedAccountFromLead)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignConvertedContactToLead/{parentId}/convertedContactId", jsonResponseFormatter.FormatToJSON(LeadController.AssignConvertedContactToLead)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignConvertedContactFromLead/{parentId}", jsonResponseFormatter.FormatToJSON(LeadController.UnassignConvertedContactFromLead)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignConvertedOpportunityToLead/{parentId}/convertedOpportunityId", jsonResponseFormatter.FormatToJSON(LeadController.AssignConvertedOpportunityToLead)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignConvertedOpportunityFromLead/{parentId}", jsonResponseFormatter.FormatToJSON(LeadController.UnassignConvertedOpportunityFromLead)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddActivitiesToLead/{parentId}/activitiesId", jsonResponseFormatter.FormatToJSON(LeadController.AddActivitiesToLead)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveActivitiesFromLead/{parentId}/activitiesIds", jsonResponseFormatter.FormatToJSON(LeadController.RemoveActivitiesFromLead)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCampaignsToLead/{parentId}/campaignsId", jsonResponseFormatter.FormatToJSON(LeadController.AddCampaignsToLead)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCampaignsFromLead/{parentId}/campaignsIds", jsonResponseFormatter.FormatToJSON(LeadController.RemoveCampaignsFromLead)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddNotesToLead/{parentId}/notesId", jsonResponseFormatter.FormatToJSON(LeadController.AddNotesToLead)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveNotesFromLead/{parentId}/notesIds", jsonResponseFormatter.FormatToJSON(LeadController.RemoveNotesFromLead)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEmailMessagesToLead/{parentId}/emailMessagesId", jsonResponseFormatter.FormatToJSON(LeadController.AddEmailMessagesToLead)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEmailMessagesFromLead/{parentId}/emailMessagesIds", jsonResponseFormatter.FormatToJSON(LeadController.RemoveEmailMessagesFromLead)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Opportunity Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Opportunity/{id}", jsonResponseFormatter.FormatToJSON(OpportunityController.GetOpportunity)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Opportunity", jsonResponseFormatter.FormatToJSON(OpportunityController.GetAllOpportunity)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOpportunity", jsonResponseFormatter.FormatToJSON(OpportunityController.CreateOpportunity)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Opportunity/{id}", jsonResponseFormatter.FormatToJSON(OpportunityController.UpdateOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOpportunity/{id}", jsonResponseFormatter.FormatToJSON(OpportunityController.DeleteOpportunity)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToOpportunity/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(OpportunityController.AssignOrganizationToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromOpportunity/{parentId}", jsonResponseFormatter.FormatToJSON(OpportunityController.UnassignOrganizationFromOpportunity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToOpportunity/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(OpportunityController.AssignAccountToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromOpportunity/{parentId}", jsonResponseFormatter.FormatToJSON(OpportunityController.UnassignAccountFromOpportunity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToOpportunity/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(OpportunityController.AssignOwnerToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromOpportunity/{parentId}", jsonResponseFormatter.FormatToJSON(OpportunityController.UnassignOwnerFromOpportunity)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddContactsToOpportunity/{parentId}/contactsId", jsonResponseFormatter.FormatToJSON(OpportunityController.AddContactsToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContactsFromOpportunity/{parentId}/contactsIds", jsonResponseFormatter.FormatToJSON(OpportunityController.RemoveContactsFromOpportunity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLineItemsToOpportunity/{parentId}/lineItemsId", jsonResponseFormatter.FormatToJSON(OpportunityController.AddLineItemsToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLineItemsFromOpportunity/{parentId}/lineItemsIds", jsonResponseFormatter.FormatToJSON(OpportunityController.RemoveLineItemsFromOpportunity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddStageHistoryToOpportunity/{parentId}/stageHistoryId", jsonResponseFormatter.FormatToJSON(OpportunityController.AddStageHistoryToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveStageHistoryFromOpportunity/{parentId}/stageHistoryIds", jsonResponseFormatter.FormatToJSON(OpportunityController.RemoveStageHistoryFromOpportunity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddQuotesToOpportunity/{parentId}/quotesId", jsonResponseFormatter.FormatToJSON(OpportunityController.AddQuotesToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQuotesFromOpportunity/{parentId}/quotesIds", jsonResponseFormatter.FormatToJSON(OpportunityController.RemoveQuotesFromOpportunity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOrdersToOpportunity/{parentId}/ordersId", jsonResponseFormatter.FormatToJSON(OpportunityController.AddOrdersToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOrdersFromOpportunity/{parentId}/ordersIds", jsonResponseFormatter.FormatToJSON(OpportunityController.RemoveOrdersFromOpportunity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCampaignsToOpportunity/{parentId}/campaignsId", jsonResponseFormatter.FormatToJSON(OpportunityController.AddCampaignsToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCampaignsFromOpportunity/{parentId}/campaignsIds", jsonResponseFormatter.FormatToJSON(OpportunityController.RemoveCampaignsFromOpportunity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddActivitiesToOpportunity/{parentId}/activitiesId", jsonResponseFormatter.FormatToJSON(OpportunityController.AddActivitiesToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveActivitiesFromOpportunity/{parentId}/activitiesIds", jsonResponseFormatter.FormatToJSON(OpportunityController.RemoveActivitiesFromOpportunity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTeamsToOpportunity/{parentId}/teamsId", jsonResponseFormatter.FormatToJSON(OpportunityController.AddTeamsToOpportunity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTeamsFromOpportunity/{parentId}/teamsIds", jsonResponseFormatter.FormatToJSON(OpportunityController.RemoveTeamsFromOpportunity)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // OpportunityLineItem Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/OpportunityLineItem/{id}", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.GetOpportunityLineItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/OpportunityLineItem", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.GetAllOpportunityLineItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOpportunityLineItem", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.CreateOpportunityLineItem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/OpportunityLineItem/{id}", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.UpdateOpportunityLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOpportunityLineItem/{id}", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.DeleteOpportunityLineItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOpportunityToOpportunityLineItem/{parentId}/opportunityId", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.AssignOpportunityToOpportunityLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOpportunityFromOpportunityLineItem/{parentId}", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.UnassignOpportunityFromOpportunityLineItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProductToOpportunityLineItem/{parentId}/productId", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.AssignProductToOpportunityLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductFromOpportunityLineItem/{parentId}", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.UnassignProductFromOpportunityLineItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPriceBookEntryToOpportunityLineItem/{parentId}/priceBookEntryId", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.AssignPriceBookEntryToOpportunityLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPriceBookEntryFromOpportunityLineItem/{parentId}", jsonResponseFormatter.FormatToJSON(OpportunityLineItemController.UnassignPriceBookEntryFromOpportunityLineItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // OpportunityStageHistory Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/OpportunityStageHistory/{id}", jsonResponseFormatter.FormatToJSON(OpportunityStageHistoryController.GetOpportunityStageHistory)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/OpportunityStageHistory", jsonResponseFormatter.FormatToJSON(OpportunityStageHistoryController.GetAllOpportunityStageHistory)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOpportunityStageHistory", jsonResponseFormatter.FormatToJSON(OpportunityStageHistoryController.CreateOpportunityStageHistory)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/OpportunityStageHistory/{id}", jsonResponseFormatter.FormatToJSON(OpportunityStageHistoryController.UpdateOpportunityStageHistory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOpportunityStageHistory/{id}", jsonResponseFormatter.FormatToJSON(OpportunityStageHistoryController.DeleteOpportunityStageHistory)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOpportunityToOpportunityStageHistory/{parentId}/opportunityId", jsonResponseFormatter.FormatToJSON(OpportunityStageHistoryController.AssignOpportunityToOpportunityStageHistory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOpportunityFromOpportunityStageHistory/{parentId}", jsonResponseFormatter.FormatToJSON(OpportunityStageHistoryController.UnassignOpportunityFromOpportunityStageHistory)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignChangedByToOpportunityStageHistory/{parentId}/changedById", jsonResponseFormatter.FormatToJSON(OpportunityStageHistoryController.AssignChangedByToOpportunityStageHistory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignChangedByFromOpportunityStageHistory/{parentId}", jsonResponseFormatter.FormatToJSON(OpportunityStageHistoryController.UnassignChangedByFromOpportunityStageHistory)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Product Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Product/{id}", jsonResponseFormatter.FormatToJSON(ProductController.GetProduct)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Product", jsonResponseFormatter.FormatToJSON(ProductController.GetAllProduct)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewProduct", jsonResponseFormatter.FormatToJSON(ProductController.CreateProduct)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Product/{id}", jsonResponseFormatter.FormatToJSON(ProductController.UpdateProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteProduct/{id}", jsonResponseFormatter.FormatToJSON(ProductController.DeleteProduct)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToProduct/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(ProductController.AssignOrganizationToProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromProduct/{parentId}", jsonResponseFormatter.FormatToJSON(ProductController.UnassignOrganizationFromProduct)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPriceBookEntriesToProduct/{parentId}/priceBookEntriesId", jsonResponseFormatter.FormatToJSON(ProductController.AddPriceBookEntriesToProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePriceBookEntriesFromProduct/{parentId}/priceBookEntriesIds", jsonResponseFormatter.FormatToJSON(ProductController.RemovePriceBookEntriesFromProduct)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOpportunityLineItemsToProduct/{parentId}/opportunityLineItemsId", jsonResponseFormatter.FormatToJSON(ProductController.AddOpportunityLineItemsToProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOpportunityLineItemsFromProduct/{parentId}/opportunityLineItemsIds", jsonResponseFormatter.FormatToJSON(ProductController.RemoveOpportunityLineItemsFromProduct)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddQuoteLineItemsToProduct/{parentId}/quoteLineItemsId", jsonResponseFormatter.FormatToJSON(ProductController.AddQuoteLineItemsToProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQuoteLineItemsFromProduct/{parentId}/quoteLineItemsIds", jsonResponseFormatter.FormatToJSON(ProductController.RemoveQuoteLineItemsFromProduct)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOrderItemsToProduct/{parentId}/orderItemsId", jsonResponseFormatter.FormatToJSON(ProductController.AddOrderItemsToProduct)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOrderItemsFromProduct/{parentId}/orderItemsIds", jsonResponseFormatter.FormatToJSON(ProductController.RemoveOrderItemsFromProduct)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PriceBook Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PriceBook/{id}", jsonResponseFormatter.FormatToJSON(PriceBookController.GetPriceBook)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PriceBook", jsonResponseFormatter.FormatToJSON(PriceBookController.GetAllPriceBook)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPriceBook", jsonResponseFormatter.FormatToJSON(PriceBookController.CreatePriceBook)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PriceBook/{id}", jsonResponseFormatter.FormatToJSON(PriceBookController.UpdatePriceBook)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePriceBook/{id}", jsonResponseFormatter.FormatToJSON(PriceBookController.DeletePriceBook)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToPriceBook/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(PriceBookController.AssignOrganizationToPriceBook)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromPriceBook/{parentId}", jsonResponseFormatter.FormatToJSON(PriceBookController.UnassignOrganizationFromPriceBook)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddEntriesToPriceBook/{parentId}/entriesId", jsonResponseFormatter.FormatToJSON(PriceBookController.AddEntriesToPriceBook)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEntriesFromPriceBook/{parentId}/entriesIds", jsonResponseFormatter.FormatToJSON(PriceBookController.RemoveEntriesFromPriceBook)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddQuotesToPriceBook/{parentId}/quotesId", jsonResponseFormatter.FormatToJSON(PriceBookController.AddQuotesToPriceBook)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQuotesFromPriceBook/{parentId}/quotesIds", jsonResponseFormatter.FormatToJSON(PriceBookController.RemoveQuotesFromPriceBook)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOrdersToPriceBook/{parentId}/ordersId", jsonResponseFormatter.FormatToJSON(PriceBookController.AddOrdersToPriceBook)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOrdersFromPriceBook/{parentId}/ordersIds", jsonResponseFormatter.FormatToJSON(PriceBookController.RemoveOrdersFromPriceBook)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PriceBookEntry Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PriceBookEntry/{id}", jsonResponseFormatter.FormatToJSON(PriceBookEntryController.GetPriceBookEntry)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PriceBookEntry", jsonResponseFormatter.FormatToJSON(PriceBookEntryController.GetAllPriceBookEntry)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPriceBookEntry", jsonResponseFormatter.FormatToJSON(PriceBookEntryController.CreatePriceBookEntry)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PriceBookEntry/{id}", jsonResponseFormatter.FormatToJSON(PriceBookEntryController.UpdatePriceBookEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePriceBookEntry/{id}", jsonResponseFormatter.FormatToJSON(PriceBookEntryController.DeletePriceBookEntry)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPriceBookToPriceBookEntry/{parentId}/priceBookId", jsonResponseFormatter.FormatToJSON(PriceBookEntryController.AssignPriceBookToPriceBookEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPriceBookFromPriceBookEntry/{parentId}", jsonResponseFormatter.FormatToJSON(PriceBookEntryController.UnassignPriceBookFromPriceBookEntry)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProductToPriceBookEntry/{parentId}/productId", jsonResponseFormatter.FormatToJSON(PriceBookEntryController.AssignProductToPriceBookEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductFromPriceBookEntry/{parentId}", jsonResponseFormatter.FormatToJSON(PriceBookEntryController.UnassignProductFromPriceBookEntry)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

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
    router.HandleFunc("/api/AssignOrganizationToQuote/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(QuoteController.AssignOrganizationToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromQuote/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteController.UnassignOrganizationFromQuote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToQuote/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(QuoteController.AssignAccountToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromQuote/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteController.UnassignAccountFromQuote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOpportunityToQuote/{parentId}/opportunityId", jsonResponseFormatter.FormatToJSON(QuoteController.AssignOpportunityToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOpportunityFromQuote/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteController.UnassignOpportunityFromQuote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToQuote/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(QuoteController.AssignOwnerToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromQuote/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteController.UnassignOwnerFromQuote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPriceBookToQuote/{parentId}/priceBookId", jsonResponseFormatter.FormatToJSON(QuoteController.AssignPriceBookToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPriceBookFromQuote/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteController.UnassignPriceBookFromQuote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOrderToQuote/{parentId}/orderId", jsonResponseFormatter.FormatToJSON(QuoteController.AssignOrderToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrderFromQuote/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteController.UnassignOrderFromQuote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLineItemsToQuote/{parentId}/lineItemsId", jsonResponseFormatter.FormatToJSON(QuoteController.AddLineItemsToQuote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLineItemsFromQuote/{parentId}/lineItemsIds", jsonResponseFormatter.FormatToJSON(QuoteController.RemoveLineItemsFromQuote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // QuoteLineItem Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/QuoteLineItem/{id}", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.GetQuoteLineItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/QuoteLineItem", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.GetAllQuoteLineItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewQuoteLineItem", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.CreateQuoteLineItem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/QuoteLineItem/{id}", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.UpdateQuoteLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteQuoteLineItem/{id}", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.DeleteQuoteLineItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignQuoteToQuoteLineItem/{parentId}/quoteId", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.AssignQuoteToQuoteLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignQuoteFromQuoteLineItem/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.UnassignQuoteFromQuoteLineItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProductToQuoteLineItem/{parentId}/productId", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.AssignProductToQuoteLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductFromQuoteLineItem/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.UnassignProductFromQuoteLineItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPriceBookEntryToQuoteLineItem/{parentId}/priceBookEntryId", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.AssignPriceBookEntryToQuoteLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPriceBookEntryFromQuoteLineItem/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.UnassignPriceBookEntryFromQuoteLineItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOpportunityLineItemToQuoteLineItem/{parentId}/opportunityLineItemId", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.AssignOpportunityLineItemToQuoteLineItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOpportunityLineItemFromQuoteLineItem/{parentId}", jsonResponseFormatter.FormatToJSON(QuoteLineItemController.UnassignOpportunityLineItemFromQuoteLineItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Order Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Order/{id}", jsonResponseFormatter.FormatToJSON(OrderController.GetOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Order", jsonResponseFormatter.FormatToJSON(OrderController.GetAllOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOrder", jsonResponseFormatter.FormatToJSON(OrderController.CreateOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Order/{id}", jsonResponseFormatter.FormatToJSON(OrderController.UpdateOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOrder/{id}", jsonResponseFormatter.FormatToJSON(OrderController.DeleteOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToOrder/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(OrderController.AssignOrganizationToOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromOrder/{parentId}", jsonResponseFormatter.FormatToJSON(OrderController.UnassignOrganizationFromOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToOrder/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(OrderController.AssignAccountToOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromOrder/{parentId}", jsonResponseFormatter.FormatToJSON(OrderController.UnassignAccountFromOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOpportunityToOrder/{parentId}/opportunityId", jsonResponseFormatter.FormatToJSON(OrderController.AssignOpportunityToOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOpportunityFromOrder/{parentId}", jsonResponseFormatter.FormatToJSON(OrderController.UnassignOpportunityFromOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignQuoteToOrder/{parentId}/quoteId", jsonResponseFormatter.FormatToJSON(OrderController.AssignQuoteToOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignQuoteFromOrder/{parentId}", jsonResponseFormatter.FormatToJSON(OrderController.UnassignQuoteFromOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToOrder/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(OrderController.AssignOwnerToOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromOrder/{parentId}", jsonResponseFormatter.FormatToJSON(OrderController.UnassignOwnerFromOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignContractToOrder/{parentId}/contractId", jsonResponseFormatter.FormatToJSON(OrderController.AssignContractToOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignContractFromOrder/{parentId}", jsonResponseFormatter.FormatToJSON(OrderController.UnassignContractFromOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPriceBookToOrder/{parentId}/priceBookId", jsonResponseFormatter.FormatToJSON(OrderController.AssignPriceBookToOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPriceBookFromOrder/{parentId}", jsonResponseFormatter.FormatToJSON(OrderController.UnassignPriceBookFromOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddItemsToOrder/{parentId}/itemsId", jsonResponseFormatter.FormatToJSON(OrderController.AddItemsToOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveItemsFromOrder/{parentId}/itemsIds", jsonResponseFormatter.FormatToJSON(OrderController.RemoveItemsFromOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // OrderItem Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/OrderItem/{id}", jsonResponseFormatter.FormatToJSON(OrderItemController.GetOrderItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/OrderItem", jsonResponseFormatter.FormatToJSON(OrderItemController.GetAllOrderItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOrderItem", jsonResponseFormatter.FormatToJSON(OrderItemController.CreateOrderItem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/OrderItem/{id}", jsonResponseFormatter.FormatToJSON(OrderItemController.UpdateOrderItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOrderItem/{id}", jsonResponseFormatter.FormatToJSON(OrderItemController.DeleteOrderItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrderToOrderItem/{parentId}/orderId", jsonResponseFormatter.FormatToJSON(OrderItemController.AssignOrderToOrderItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrderFromOrderItem/{parentId}", jsonResponseFormatter.FormatToJSON(OrderItemController.UnassignOrderFromOrderItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProductToOrderItem/{parentId}/productId", jsonResponseFormatter.FormatToJSON(OrderItemController.AssignProductToOrderItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductFromOrderItem/{parentId}", jsonResponseFormatter.FormatToJSON(OrderItemController.UnassignProductFromOrderItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPriceBookEntryToOrderItem/{parentId}/priceBookEntryId", jsonResponseFormatter.FormatToJSON(OrderItemController.AssignPriceBookEntryToOrderItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPriceBookEntryFromOrderItem/{parentId}", jsonResponseFormatter.FormatToJSON(OrderItemController.UnassignPriceBookEntryFromOrderItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Contract Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Contract/{id}", jsonResponseFormatter.FormatToJSON(ContractController.GetContract)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Contract", jsonResponseFormatter.FormatToJSON(ContractController.GetAllContract)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewContract", jsonResponseFormatter.FormatToJSON(ContractController.CreateContract)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Contract/{id}", jsonResponseFormatter.FormatToJSON(ContractController.UpdateContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteContract/{id}", jsonResponseFormatter.FormatToJSON(ContractController.DeleteContract)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToContract/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(ContractController.AssignOrganizationToContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromContract/{parentId}", jsonResponseFormatter.FormatToJSON(ContractController.UnassignOrganizationFromContract)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToContract/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(ContractController.AssignAccountToContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromContract/{parentId}", jsonResponseFormatter.FormatToJSON(ContractController.UnassignAccountFromContract)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToContract/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(ContractController.AssignOwnerToContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromContract/{parentId}", jsonResponseFormatter.FormatToJSON(ContractController.UnassignOwnerFromContract)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddOrdersToContract/{parentId}/ordersId", jsonResponseFormatter.FormatToJSON(ContractController.AddOrdersToContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOrdersFromContract/{parentId}/ordersIds", jsonResponseFormatter.FormatToJSON(ContractController.RemoveOrdersFromContract)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCasesToContract/{parentId}/casesId", jsonResponseFormatter.FormatToJSON(ContractController.AddCasesToContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCasesFromContract/{parentId}/casesIds", jsonResponseFormatter.FormatToJSON(ContractController.RemoveCasesFromContract)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Case_ Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Case_/{id}", jsonResponseFormatter.FormatToJSON(Case_Controller.GetCase_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Case_", jsonResponseFormatter.FormatToJSON(Case_Controller.GetAllCase_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCase_", jsonResponseFormatter.FormatToJSON(Case_Controller.CreateCase_)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Case_/{id}", jsonResponseFormatter.FormatToJSON(Case_Controller.UpdateCase_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCase_/{id}", jsonResponseFormatter.FormatToJSON(Case_Controller.DeleteCase_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToCase_/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(Case_Controller.AssignOrganizationToCase_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromCase_/{parentId}", jsonResponseFormatter.FormatToJSON(Case_Controller.UnassignOrganizationFromCase_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToCase_/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(Case_Controller.AssignAccountToCase_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromCase_/{parentId}", jsonResponseFormatter.FormatToJSON(Case_Controller.UnassignAccountFromCase_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignContactToCase_/{parentId}/contactId", jsonResponseFormatter.FormatToJSON(Case_Controller.AssignContactToCase_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignContactFromCase_/{parentId}", jsonResponseFormatter.FormatToJSON(Case_Controller.UnassignContactFromCase_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToCase_/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(Case_Controller.AssignOwnerToCase_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromCase_/{parentId}", jsonResponseFormatter.FormatToJSON(Case_Controller.UnassignOwnerFromCase_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTeamToCase_/{parentId}/teamId", jsonResponseFormatter.FormatToJSON(Case_Controller.AssignTeamToCase_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTeamFromCase_/{parentId}", jsonResponseFormatter.FormatToJSON(Case_Controller.UnassignTeamFromCase_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddActivitiesToCase_/{parentId}/activitiesId", jsonResponseFormatter.FormatToJSON(Case_Controller.AddActivitiesToCase_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveActivitiesFromCase_/{parentId}/activitiesIds", jsonResponseFormatter.FormatToJSON(Case_Controller.RemoveActivitiesFromCase_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCaseCommentsToCase_/{parentId}/caseCommentsId", jsonResponseFormatter.FormatToJSON(Case_Controller.AddCaseCommentsToCase_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCaseCommentsFromCase_/{parentId}/caseCommentsIds", jsonResponseFormatter.FormatToJSON(Case_Controller.RemoveCaseCommentsFromCase_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEmailsToCase_/{parentId}/emailsId", jsonResponseFormatter.FormatToJSON(Case_Controller.AddEmailsToCase_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEmailsFromCase_/{parentId}/emailsIds", jsonResponseFormatter.FormatToJSON(Case_Controller.RemoveEmailsFromCase_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRelatedOpportunitiesToCase_/{parentId}/relatedOpportunitiesId", jsonResponseFormatter.FormatToJSON(Case_Controller.AddRelatedOpportunitiesToCase_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRelatedOpportunitiesFromCase_/{parentId}/relatedOpportunitiesIds", jsonResponseFormatter.FormatToJSON(Case_Controller.RemoveRelatedOpportunitiesFromCase_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Activity Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Activity/{id}", jsonResponseFormatter.FormatToJSON(ActivityController.GetActivity)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Activity", jsonResponseFormatter.FormatToJSON(ActivityController.GetAllActivity)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewActivity", jsonResponseFormatter.FormatToJSON(ActivityController.CreateActivity)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Activity/{id}", jsonResponseFormatter.FormatToJSON(ActivityController.UpdateActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteActivity/{id}", jsonResponseFormatter.FormatToJSON(ActivityController.DeleteActivity)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToActivity/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(ActivityController.AssignOrganizationToActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromActivity/{parentId}", jsonResponseFormatter.FormatToJSON(ActivityController.UnassignOrganizationFromActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToActivity/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(ActivityController.AssignOwnerToActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromActivity/{parentId}", jsonResponseFormatter.FormatToJSON(ActivityController.UnassignOwnerFromActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToActivity/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(ActivityController.AssignAccountToActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromActivity/{parentId}", jsonResponseFormatter.FormatToJSON(ActivityController.UnassignAccountFromActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignContactToActivity/{parentId}/contactId", jsonResponseFormatter.FormatToJSON(ActivityController.AssignContactToActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignContactFromActivity/{parentId}", jsonResponseFormatter.FormatToJSON(ActivityController.UnassignContactFromActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLeadToActivity/{parentId}/leadId", jsonResponseFormatter.FormatToJSON(ActivityController.AssignLeadToActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLeadFromActivity/{parentId}", jsonResponseFormatter.FormatToJSON(ActivityController.UnassignLeadFromActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOpportunityToActivity/{parentId}/opportunityId", jsonResponseFormatter.FormatToJSON(ActivityController.AssignOpportunityToActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOpportunityFromActivity/{parentId}", jsonResponseFormatter.FormatToJSON(ActivityController.UnassignOpportunityFromActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCaseToActivity/{parentId}/caseId", jsonResponseFormatter.FormatToJSON(ActivityController.AssignCaseToActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCaseFromActivity/{parentId}", jsonResponseFormatter.FormatToJSON(ActivityController.UnassignCaseFromActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCampaignToActivity/{parentId}/campaignId", jsonResponseFormatter.FormatToJSON(ActivityController.AssignCampaignToActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCampaignFromActivity/{parentId}", jsonResponseFormatter.FormatToJSON(ActivityController.UnassignCampaignFromActivity)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

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
    router.HandleFunc("/api/AssignOrganizationToCampaign/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(CampaignController.AssignOrganizationToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromCampaign/{parentId}", jsonResponseFormatter.FormatToJSON(CampaignController.UnassignOrganizationFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignParentCampaignToCampaign/{parentId}/parentCampaignId", jsonResponseFormatter.FormatToJSON(CampaignController.AssignParentCampaignToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignParentCampaignFromCampaign/{parentId}", jsonResponseFormatter.FormatToJSON(CampaignController.UnassignParentCampaignFromCampaign)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddChildCampaignsToCampaign/{parentId}/childCampaignsId", jsonResponseFormatter.FormatToJSON(CampaignController.AddChildCampaignsToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveChildCampaignsFromCampaign/{parentId}/childCampaignsIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveChildCampaignsFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMembersToCampaign/{parentId}/membersId", jsonResponseFormatter.FormatToJSON(CampaignController.AddMembersToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMembersFromCampaign/{parentId}/membersIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveMembersFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOpportunitiesToCampaign/{parentId}/opportunitiesId", jsonResponseFormatter.FormatToJSON(CampaignController.AddOpportunitiesToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOpportunitiesFromCampaign/{parentId}/opportunitiesIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveOpportunitiesFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAccountsToCampaign/{parentId}/accountsId", jsonResponseFormatter.FormatToJSON(CampaignController.AddAccountsToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAccountsFromCampaign/{parentId}/accountsIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveAccountsFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLeadsToCampaign/{parentId}/leadsId", jsonResponseFormatter.FormatToJSON(CampaignController.AddLeadsToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLeadsFromCampaign/{parentId}/leadsIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveLeadsFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddContactsToCampaign/{parentId}/contactsId", jsonResponseFormatter.FormatToJSON(CampaignController.AddContactsToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContactsFromCampaign/{parentId}/contactsIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveContactsFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTeamsToCampaign/{parentId}/teamsId", jsonResponseFormatter.FormatToJSON(CampaignController.AddTeamsToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTeamsFromCampaign/{parentId}/teamsIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveTeamsFromCampaign)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddActivitiesToCampaign/{parentId}/activitiesId", jsonResponseFormatter.FormatToJSON(CampaignController.AddActivitiesToCampaign)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveActivitiesFromCampaign/{parentId}/activitiesIds", jsonResponseFormatter.FormatToJSON(CampaignController.RemoveActivitiesFromCampaign)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CampaignMember Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CampaignMember/{id}", jsonResponseFormatter.FormatToJSON(CampaignMemberController.GetCampaignMember)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CampaignMember", jsonResponseFormatter.FormatToJSON(CampaignMemberController.GetAllCampaignMember)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCampaignMember", jsonResponseFormatter.FormatToJSON(CampaignMemberController.CreateCampaignMember)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CampaignMember/{id}", jsonResponseFormatter.FormatToJSON(CampaignMemberController.UpdateCampaignMember)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCampaignMember/{id}", jsonResponseFormatter.FormatToJSON(CampaignMemberController.DeleteCampaignMember)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCampaignToCampaignMember/{parentId}/campaignId", jsonResponseFormatter.FormatToJSON(CampaignMemberController.AssignCampaignToCampaignMember)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCampaignFromCampaignMember/{parentId}", jsonResponseFormatter.FormatToJSON(CampaignMemberController.UnassignCampaignFromCampaignMember)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLeadToCampaignMember/{parentId}/leadId", jsonResponseFormatter.FormatToJSON(CampaignMemberController.AssignLeadToCampaignMember)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLeadFromCampaignMember/{parentId}", jsonResponseFormatter.FormatToJSON(CampaignMemberController.UnassignLeadFromCampaignMember)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignContactToCampaignMember/{parentId}/contactId", jsonResponseFormatter.FormatToJSON(CampaignMemberController.AssignContactToCampaignMember)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignContactFromCampaignMember/{parentId}", jsonResponseFormatter.FormatToJSON(CampaignMemberController.UnassignContactFromCampaignMember)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Note Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Note/{id}", jsonResponseFormatter.FormatToJSON(NoteController.GetNote)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Note", jsonResponseFormatter.FormatToJSON(NoteController.GetAllNote)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewNote", jsonResponseFormatter.FormatToJSON(NoteController.CreateNote)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Note/{id}", jsonResponseFormatter.FormatToJSON(NoteController.UpdateNote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteNote/{id}", jsonResponseFormatter.FormatToJSON(NoteController.DeleteNote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToNote/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(NoteController.AssignOrganizationToNote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromNote/{parentId}", jsonResponseFormatter.FormatToJSON(NoteController.UnassignOrganizationFromNote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToNote/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(NoteController.AssignOwnerToNote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromNote/{parentId}", jsonResponseFormatter.FormatToJSON(NoteController.UnassignOwnerFromNote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToNote/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(NoteController.AssignAccountToNote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromNote/{parentId}", jsonResponseFormatter.FormatToJSON(NoteController.UnassignAccountFromNote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignContactToNote/{parentId}/contactId", jsonResponseFormatter.FormatToJSON(NoteController.AssignContactToNote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignContactFromNote/{parentId}", jsonResponseFormatter.FormatToJSON(NoteController.UnassignContactFromNote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOpportunityToNote/{parentId}/opportunityId", jsonResponseFormatter.FormatToJSON(NoteController.AssignOpportunityToNote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOpportunityFromNote/{parentId}", jsonResponseFormatter.FormatToJSON(NoteController.UnassignOpportunityFromNote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCaseToNote/{parentId}/caseId", jsonResponseFormatter.FormatToJSON(NoteController.AssignCaseToNote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCaseFromNote/{parentId}", jsonResponseFormatter.FormatToJSON(NoteController.UnassignCaseFromNote)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLeadToNote/{parentId}/leadId", jsonResponseFormatter.FormatToJSON(NoteController.AssignLeadToNote)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLeadFromNote/{parentId}", jsonResponseFormatter.FormatToJSON(NoteController.UnassignLeadFromNote)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // EmailMessage Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/EmailMessage/{id}", jsonResponseFormatter.FormatToJSON(EmailMessageController.GetEmailMessage)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/EmailMessage", jsonResponseFormatter.FormatToJSON(EmailMessageController.GetAllEmailMessage)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEmailMessage", jsonResponseFormatter.FormatToJSON(EmailMessageController.CreateEmailMessage)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/EmailMessage/{id}", jsonResponseFormatter.FormatToJSON(EmailMessageController.UpdateEmailMessage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEmailMessage/{id}", jsonResponseFormatter.FormatToJSON(EmailMessageController.DeleteEmailMessage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToEmailMessage/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(EmailMessageController.AssignOrganizationToEmailMessage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromEmailMessage/{parentId}", jsonResponseFormatter.FormatToJSON(EmailMessageController.UnassignOrganizationFromEmailMessage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToEmailMessage/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(EmailMessageController.AssignOwnerToEmailMessage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromEmailMessage/{parentId}", jsonResponseFormatter.FormatToJSON(EmailMessageController.UnassignOwnerFromEmailMessage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAccountToEmailMessage/{parentId}/accountId", jsonResponseFormatter.FormatToJSON(EmailMessageController.AssignAccountToEmailMessage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAccountFromEmailMessage/{parentId}", jsonResponseFormatter.FormatToJSON(EmailMessageController.UnassignAccountFromEmailMessage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignContactToEmailMessage/{parentId}/contactId", jsonResponseFormatter.FormatToJSON(EmailMessageController.AssignContactToEmailMessage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignContactFromEmailMessage/{parentId}", jsonResponseFormatter.FormatToJSON(EmailMessageController.UnassignContactFromEmailMessage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLeadToEmailMessage/{parentId}/leadId", jsonResponseFormatter.FormatToJSON(EmailMessageController.AssignLeadToEmailMessage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLeadFromEmailMessage/{parentId}", jsonResponseFormatter.FormatToJSON(EmailMessageController.UnassignLeadFromEmailMessage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCaseToEmailMessage/{parentId}/caseId", jsonResponseFormatter.FormatToJSON(EmailMessageController.AssignCaseToEmailMessage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCaseFromEmailMessage/{parentId}", jsonResponseFormatter.FormatToJSON(EmailMessageController.UnassignCaseFromEmailMessage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOpportunityToEmailMessage/{parentId}/opportunityId", jsonResponseFormatter.FormatToJSON(EmailMessageController.AssignOpportunityToEmailMessage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOpportunityFromEmailMessage/{parentId}", jsonResponseFormatter.FormatToJSON(EmailMessageController.UnassignOpportunityFromEmailMessage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCampaignToEmailMessage/{parentId}/campaignId", jsonResponseFormatter.FormatToJSON(EmailMessageController.AssignCampaignToEmailMessage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCampaignFromEmailMessage/{parentId}", jsonResponseFormatter.FormatToJSON(EmailMessageController.UnassignCampaignFromEmailMessage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    return router
}
