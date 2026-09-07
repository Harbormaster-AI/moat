package router

import (

    OrganizationController "governance-on-golang/internal/controller"
    GovernanceBodyController "governance-on-golang/internal/controller"
    PersonController "governance-on-golang/internal/controller"
    RoleController "governance-on-golang/internal/controller"
    RoleAssignmentController "governance-on-golang/internal/controller"
    PolicyController "governance-on-golang/internal/controller"
    ProcedureController "governance-on-golang/internal/controller"
    RegulationController "governance-on-golang/internal/controller"
    ObligationController "governance-on-golang/internal/controller"
    ControlController "governance-on-golang/internal/controller"
    ControlTest_Controller "governance-on-golang/internal/controller"
    EvidenceController "governance-on-golang/internal/controller"
    RiskController "governance-on-golang/internal/controller"
    RiskAssessmentController "governance-on-golang/internal/controller"
    ComplianceProgramController "governance-on-golang/internal/controller"
    ComplianceRequirementController "governance-on-golang/internal/controller"
    AttestationController "governance-on-golang/internal/controller"
    AuditProgramController "governance-on-golang/internal/controller"
    AuditEngagementController "governance-on-golang/internal/controller"
    AuditWorkpaperController "governance-on-golang/internal/controller"
    AuditFindingController "governance-on-golang/internal/controller"
    CorrectiveActionController "governance-on-golang/internal/controller"
    IssueController "governance-on-golang/internal/controller"
    BusinessUnitController "governance-on-golang/internal/controller"
    DataProcessingActivityController "governance-on-golang/internal/controller"
    DataCategoryController "governance-on-golang/internal/controller"
    System_Controller "governance-on-golang/internal/controller"
    PrivacyNoticeController "governance-on-golang/internal/controller"
    DataSubjectRequestController "governance-on-golang/internal/controller"
    RecordsRepositoryController "governance-on-golang/internal/controller"
    Record_Controller "governance-on-golang/internal/controller"
    RetentionScheduleController "governance-on-golang/internal/controller"
    DispositionReviewController "governance-on-golang/internal/controller"
    LegalHoldController "governance-on-golang/internal/controller"
    MatterController "governance-on-golang/internal/controller"
    ThirdPartyController "governance-on-golang/internal/controller"
    ThirdPartyAssessmentController "governance-on-golang/internal/controller"
    ContractController "governance-on-golang/internal/controller"
    Exception_Controller "governance-on-golang/internal/controller"
    ConsentController "governance-on-golang/internal/controller"
    DataBreachController "governance-on-golang/internal/controller"
    jsonResponseFormatter "governance-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "governance-on-golang/internal/controller"

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
    router.HandleFunc("/api/AddGovernanceBodiesToOrganization/{parentId}/governanceBodiesId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddGovernanceBodiesToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveGovernanceBodiesFromOrganization/{parentId}/governanceBodiesIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveGovernanceBodiesFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPoliciesToOrganization/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddPoliciesToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromOrganization/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemovePoliciesFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRisksToOrganization/{parentId}/risksId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddRisksToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRisksFromOrganization/{parentId}/risksIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveRisksFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddThirdPartiesToOrganization/{parentId}/thirdPartiesId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddThirdPartiesToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveThirdPartiesFromOrganization/{parentId}/thirdPartiesIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveThirdPartiesFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRecordsRepositoriesToOrganization/{parentId}/recordsRepositoriesId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddRecordsRepositoriesToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRecordsRepositoriesFromOrganization/{parentId}/recordsRepositoriesIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveRecordsRepositoriesFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataProcessingActivitiesToOrganization/{parentId}/dataProcessingActivitiesId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddDataProcessingActivitiesToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataProcessingActivitiesFromOrganization/{parentId}/dataProcessingActivitiesIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveDataProcessingActivitiesFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddComplianceProgramsToOrganization/{parentId}/complianceProgramsId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddComplianceProgramsToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveComplianceProgramsFromOrganization/{parentId}/complianceProgramsIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveComplianceProgramsFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAuditProgramsToOrganization/{parentId}/auditProgramsId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddAuditProgramsToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAuditProgramsFromOrganization/{parentId}/auditProgramsIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveAuditProgramsFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddBusinessUnitsToOrganization/{parentId}/businessUnitsId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddBusinessUnitsToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBusinessUnitsFromOrganization/{parentId}/businessUnitsIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveBusinessUnitsFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMattersToOrganization/{parentId}/mattersId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddMattersToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMattersFromOrganization/{parentId}/mattersIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveMattersFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataBreachesToOrganization/{parentId}/dataBreachesId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddDataBreachesToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataBreachesFromOrganization/{parentId}/dataBreachesIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveDataBreachesFromOrganization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // GovernanceBody Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/GovernanceBody/{id}", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.GetGovernanceBody)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/GovernanceBody", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.GetAllGovernanceBody)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewGovernanceBody", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.CreateGovernanceBody)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/GovernanceBody/{id}", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.UpdateGovernanceBody)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteGovernanceBody/{id}", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.DeleteGovernanceBody)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToGovernanceBody/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.AssignOrganizationToGovernanceBody)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromGovernanceBody/{parentId}", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.UnassignOrganizationFromGovernanceBody)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddRoleAssignmentsToGovernanceBody/{parentId}/roleAssignmentsId", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.AddRoleAssignmentsToGovernanceBody)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRoleAssignmentsFromGovernanceBody/{parentId}/roleAssignmentsIds", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.RemoveRoleAssignmentsFromGovernanceBody)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPoliciesToGovernanceBody/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.AddPoliciesToGovernanceBody)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromGovernanceBody/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(GovernanceBodyController.RemovePoliciesFromGovernanceBody)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Person Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Person/{id}", jsonResponseFormatter.FormatToJSON(PersonController.GetPerson)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Person", jsonResponseFormatter.FormatToJSON(PersonController.GetAllPerson)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPerson", jsonResponseFormatter.FormatToJSON(PersonController.CreatePerson)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Person/{id}", jsonResponseFormatter.FormatToJSON(PersonController.UpdatePerson)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePerson/{id}", jsonResponseFormatter.FormatToJSON(PersonController.DeletePerson)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddRoleAssignmentsToPerson/{parentId}/roleAssignmentsId", jsonResponseFormatter.FormatToJSON(PersonController.AddRoleAssignmentsToPerson)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRoleAssignmentsFromPerson/{parentId}/roleAssignmentsIds", jsonResponseFormatter.FormatToJSON(PersonController.RemoveRoleAssignmentsFromPerson)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOwnedPoliciesToPerson/{parentId}/ownedPoliciesId", jsonResponseFormatter.FormatToJSON(PersonController.AddOwnedPoliciesToPerson)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOwnedPoliciesFromPerson/{parentId}/ownedPoliciesIds", jsonResponseFormatter.FormatToJSON(PersonController.RemoveOwnedPoliciesFromPerson)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCorrectiveActionsToPerson/{parentId}/correctiveActionsId", jsonResponseFormatter.FormatToJSON(PersonController.AddCorrectiveActionsToPerson)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCorrectiveActionsFromPerson/{parentId}/correctiveActionsIds", jsonResponseFormatter.FormatToJSON(PersonController.RemoveCorrectiveActionsFromPerson)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Role Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Role/{id}", jsonResponseFormatter.FormatToJSON(RoleController.GetRole)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Role", jsonResponseFormatter.FormatToJSON(RoleController.GetAllRole)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRole", jsonResponseFormatter.FormatToJSON(RoleController.CreateRole)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Role/{id}", jsonResponseFormatter.FormatToJSON(RoleController.UpdateRole)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRole/{id}", jsonResponseFormatter.FormatToJSON(RoleController.DeleteRole)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAssignmentsToRole/{parentId}/assignmentsId", jsonResponseFormatter.FormatToJSON(RoleController.AddAssignmentsToRole)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAssignmentsFromRole/{parentId}/assignmentsIds", jsonResponseFormatter.FormatToJSON(RoleController.RemoveAssignmentsFromRole)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // RoleAssignment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RoleAssignment/{id}", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.GetRoleAssignment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RoleAssignment", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.GetAllRoleAssignment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRoleAssignment", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.CreateRoleAssignment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RoleAssignment/{id}", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.UpdateRoleAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRoleAssignment/{id}", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.DeleteRoleAssignment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPersonToRoleAssignment/{parentId}/personId", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.AssignPersonToRoleAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPersonFromRoleAssignment/{parentId}", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.UnassignPersonFromRoleAssignment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRoleToRoleAssignment/{parentId}/roleId", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.AssignRoleToRoleAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRoleFromRoleAssignment/{parentId}", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.UnassignRoleFromRoleAssignment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignGovernanceBodyToRoleAssignment/{parentId}/governanceBodyId", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.AssignGovernanceBodyToRoleAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignGovernanceBodyFromRoleAssignment/{parentId}", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.UnassignGovernanceBodyFromRoleAssignment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOrganizationToRoleAssignment/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.AssignOrganizationToRoleAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromRoleAssignment/{parentId}", jsonResponseFormatter.FormatToJSON(RoleAssignmentController.UnassignOrganizationFromRoleAssignment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

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
    router.HandleFunc("/api/AssignOrganizationToPolicy/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(PolicyController.AssignOrganizationToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(PolicyController.UnassignOrganizationFromPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddOwnersToPolicy/{parentId}/ownersId", jsonResponseFormatter.FormatToJSON(PolicyController.AddOwnersToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOwnersFromPolicy/{parentId}/ownersIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveOwnersFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRelatedRequirementsToPolicy/{parentId}/relatedRequirementsId", jsonResponseFormatter.FormatToJSON(PolicyController.AddRelatedRequirementsToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRelatedRequirementsFromPolicy/{parentId}/relatedRequirementsIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveRelatedRequirementsFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddControlsToPolicy/{parentId}/controlsId", jsonResponseFormatter.FormatToJSON(PolicyController.AddControlsToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveControlsFromPolicy/{parentId}/controlsIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveControlsFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddProceduresToPolicy/{parentId}/proceduresId", jsonResponseFormatter.FormatToJSON(PolicyController.AddProceduresToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProceduresFromPolicy/{parentId}/proceduresIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveProceduresFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddExceptionsToPolicy/{parentId}/exceptionsId", jsonResponseFormatter.FormatToJSON(PolicyController.AddExceptionsToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveExceptionsFromPolicy/{parentId}/exceptionsIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveExceptionsFromPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAttestationsToPolicy/{parentId}/attestationsId", jsonResponseFormatter.FormatToJSON(PolicyController.AddAttestationsToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAttestationsFromPolicy/{parentId}/attestationsIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveAttestationsFromPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Procedure Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Procedure/{id}", jsonResponseFormatter.FormatToJSON(ProcedureController.GetProcedure)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Procedure", jsonResponseFormatter.FormatToJSON(ProcedureController.GetAllProcedure)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewProcedure", jsonResponseFormatter.FormatToJSON(ProcedureController.CreateProcedure)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Procedure/{id}", jsonResponseFormatter.FormatToJSON(ProcedureController.UpdateProcedure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteProcedure/{id}", jsonResponseFormatter.FormatToJSON(ProcedureController.DeleteProcedure)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPolicyToProcedure/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(ProcedureController.AssignPolicyToProcedure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromProcedure/{parentId}", jsonResponseFormatter.FormatToJSON(ProcedureController.UnassignPolicyFromProcedure)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddControlsToProcedure/{parentId}/controlsId", jsonResponseFormatter.FormatToJSON(ProcedureController.AddControlsToProcedure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveControlsFromProcedure/{parentId}/controlsIds", jsonResponseFormatter.FormatToJSON(ProcedureController.RemoveControlsFromProcedure)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Regulation Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Regulation/{id}", jsonResponseFormatter.FormatToJSON(RegulationController.GetRegulation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Regulation", jsonResponseFormatter.FormatToJSON(RegulationController.GetAllRegulation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRegulation", jsonResponseFormatter.FormatToJSON(RegulationController.CreateRegulation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Regulation/{id}", jsonResponseFormatter.FormatToJSON(RegulationController.UpdateRegulation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRegulation/{id}", jsonResponseFormatter.FormatToJSON(RegulationController.DeleteRegulation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddObligationsToRegulation/{parentId}/obligationsId", jsonResponseFormatter.FormatToJSON(RegulationController.AddObligationsToRegulation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveObligationsFromRegulation/{parentId}/obligationsIds", jsonResponseFormatter.FormatToJSON(RegulationController.RemoveObligationsFromRegulation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddComplianceProgramsToRegulation/{parentId}/complianceProgramsId", jsonResponseFormatter.FormatToJSON(RegulationController.AddComplianceProgramsToRegulation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveComplianceProgramsFromRegulation/{parentId}/complianceProgramsIds", jsonResponseFormatter.FormatToJSON(RegulationController.RemoveComplianceProgramsFromRegulation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Obligation Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Obligation/{id}", jsonResponseFormatter.FormatToJSON(ObligationController.GetObligation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Obligation", jsonResponseFormatter.FormatToJSON(ObligationController.GetAllObligation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewObligation", jsonResponseFormatter.FormatToJSON(ObligationController.CreateObligation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Obligation/{id}", jsonResponseFormatter.FormatToJSON(ObligationController.UpdateObligation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteObligation/{id}", jsonResponseFormatter.FormatToJSON(ObligationController.DeleteObligation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRegulationToObligation/{parentId}/regulationId", jsonResponseFormatter.FormatToJSON(ObligationController.AssignRegulationToObligation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRegulationFromObligation/{parentId}", jsonResponseFormatter.FormatToJSON(ObligationController.UnassignRegulationFromObligation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddControlsToObligation/{parentId}/controlsId", jsonResponseFormatter.FormatToJSON(ObligationController.AddControlsToObligation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveControlsFromObligation/{parentId}/controlsIds", jsonResponseFormatter.FormatToJSON(ObligationController.RemoveControlsFromObligation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPoliciesToObligation/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(ObligationController.AddPoliciesToObligation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromObligation/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(ObligationController.RemovePoliciesFromObligation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddContractsToObligation/{parentId}/contractsId", jsonResponseFormatter.FormatToJSON(ObligationController.AddContractsToObligation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContractsFromObligation/{parentId}/contractsIds", jsonResponseFormatter.FormatToJSON(ObligationController.RemoveContractsFromObligation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Control Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Control/{id}", jsonResponseFormatter.FormatToJSON(ControlController.GetControl)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Control", jsonResponseFormatter.FormatToJSON(ControlController.GetAllControl)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewControl", jsonResponseFormatter.FormatToJSON(ControlController.CreateControl)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Control/{id}", jsonResponseFormatter.FormatToJSON(ControlController.UpdateControl)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteControl/{id}", jsonResponseFormatter.FormatToJSON(ControlController.DeleteControl)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPolicyToControl/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(ControlController.AssignPolicyToControl)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromControl/{parentId}", jsonResponseFormatter.FormatToJSON(ControlController.UnassignPolicyFromControl)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddControlTestsToControl/{parentId}/controlTestsId", jsonResponseFormatter.FormatToJSON(ControlController.AddControlTestsToControl)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveControlTestsFromControl/{parentId}/controlTestsIds", jsonResponseFormatter.FormatToJSON(ControlController.RemoveControlTestsFromControl)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEvidenceToControl/{parentId}/evidenceId", jsonResponseFormatter.FormatToJSON(ControlController.AddEvidenceToControl)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEvidenceFromControl/{parentId}/evidenceIds", jsonResponseFormatter.FormatToJSON(ControlController.RemoveEvidenceFromControl)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRisksToControl/{parentId}/risksId", jsonResponseFormatter.FormatToJSON(ControlController.AddRisksToControl)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRisksFromControl/{parentId}/risksIds", jsonResponseFormatter.FormatToJSON(ControlController.RemoveRisksFromControl)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddObligationsToControl/{parentId}/obligationsId", jsonResponseFormatter.FormatToJSON(ControlController.AddObligationsToControl)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveObligationsFromControl/{parentId}/obligationsIds", jsonResponseFormatter.FormatToJSON(ControlController.RemoveObligationsFromControl)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddProceduresToControl/{parentId}/proceduresId", jsonResponseFormatter.FormatToJSON(ControlController.AddProceduresToControl)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProceduresFromControl/{parentId}/proceduresIds", jsonResponseFormatter.FormatToJSON(ControlController.RemoveProceduresFromControl)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddIssuesToControl/{parentId}/issuesId", jsonResponseFormatter.FormatToJSON(ControlController.AddIssuesToControl)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveIssuesFromControl/{parentId}/issuesIds", jsonResponseFormatter.FormatToJSON(ControlController.RemoveIssuesFromControl)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ControlTest_ Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ControlTest_/{id}", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.GetControlTest_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ControlTest_", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.GetAllControlTest_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewControlTest_", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.CreateControlTest_)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ControlTest_/{id}", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.UpdateControlTest_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteControlTest_/{id}", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.DeleteControlTest_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignControlToControlTest_/{parentId}/controlId", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.AssignControlToControlTest_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignControlFromControlTest_/{parentId}", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.UnassignControlFromControlTest_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEngagementToControlTest_/{parentId}/engagementId", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.AssignEngagementToControlTest_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEngagementFromControlTest_/{parentId}", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.UnassignEngagementFromControlTest_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddEvidenceToControlTest_/{parentId}/evidenceId", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.AddEvidenceToControlTest_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEvidenceFromControlTest_/{parentId}/evidenceIds", jsonResponseFormatter.FormatToJSON(ControlTest_Controller.RemoveEvidenceFromControlTest_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Evidence Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Evidence/{id}", jsonResponseFormatter.FormatToJSON(EvidenceController.GetEvidence)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Evidence", jsonResponseFormatter.FormatToJSON(EvidenceController.GetAllEvidence)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEvidence", jsonResponseFormatter.FormatToJSON(EvidenceController.CreateEvidence)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Evidence/{id}", jsonResponseFormatter.FormatToJSON(EvidenceController.UpdateEvidence)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEvidence/{id}", jsonResponseFormatter.FormatToJSON(EvidenceController.DeleteEvidence)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignControlTestToEvidence/{parentId}/controlTestId", jsonResponseFormatter.FormatToJSON(EvidenceController.AssignControlTestToEvidence)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignControlTestFromEvidence/{parentId}", jsonResponseFormatter.FormatToJSON(EvidenceController.UnassignControlTestFromEvidence)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignControlToEvidence/{parentId}/controlId", jsonResponseFormatter.FormatToJSON(EvidenceController.AssignControlToEvidence)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignControlFromEvidence/{parentId}", jsonResponseFormatter.FormatToJSON(EvidenceController.UnassignControlFromEvidence)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignObligationToEvidence/{parentId}/obligationId", jsonResponseFormatter.FormatToJSON(EvidenceController.AssignObligationToEvidence)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignObligationFromEvidence/{parentId}", jsonResponseFormatter.FormatToJSON(EvidenceController.UnassignObligationFromEvidence)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkpaperToEvidence/{parentId}/workpaperId", jsonResponseFormatter.FormatToJSON(EvidenceController.AssignWorkpaperToEvidence)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkpaperFromEvidence/{parentId}", jsonResponseFormatter.FormatToJSON(EvidenceController.UnassignWorkpaperFromEvidence)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Risk Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Risk/{id}", jsonResponseFormatter.FormatToJSON(RiskController.GetRisk)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Risk", jsonResponseFormatter.FormatToJSON(RiskController.GetAllRisk)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRisk", jsonResponseFormatter.FormatToJSON(RiskController.CreateRisk)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Risk/{id}", jsonResponseFormatter.FormatToJSON(RiskController.UpdateRisk)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRisk/{id}", jsonResponseFormatter.FormatToJSON(RiskController.DeleteRisk)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToRisk/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(RiskController.AssignOrganizationToRisk)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromRisk/{parentId}", jsonResponseFormatter.FormatToJSON(RiskController.UnassignOrganizationFromRisk)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddControlsToRisk/{parentId}/controlsId", jsonResponseFormatter.FormatToJSON(RiskController.AddControlsToRisk)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveControlsFromRisk/{parentId}/controlsIds", jsonResponseFormatter.FormatToJSON(RiskController.RemoveControlsFromRisk)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAssessmentsToRisk/{parentId}/assessmentsId", jsonResponseFormatter.FormatToJSON(RiskController.AddAssessmentsToRisk)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAssessmentsFromRisk/{parentId}/assessmentsIds", jsonResponseFormatter.FormatToJSON(RiskController.RemoveAssessmentsFromRisk)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddIssuesToRisk/{parentId}/issuesId", jsonResponseFormatter.FormatToJSON(RiskController.AddIssuesToRisk)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveIssuesFromRisk/{parentId}/issuesIds", jsonResponseFormatter.FormatToJSON(RiskController.RemoveIssuesFromRisk)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFindingsToRisk/{parentId}/findingsId", jsonResponseFormatter.FormatToJSON(RiskController.AddFindingsToRisk)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFindingsFromRisk/{parentId}/findingsIds", jsonResponseFormatter.FormatToJSON(RiskController.RemoveFindingsFromRisk)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AssignRiskToRiskAssessment/{parentId}/riskId", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.AssignRiskToRiskAssessment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRiskFromRiskAssessment/{parentId}", jsonResponseFormatter.FormatToJSON(RiskAssessmentController.UnassignRiskFromRiskAssessment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ComplianceProgram Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ComplianceProgram/{id}", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.GetComplianceProgram)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ComplianceProgram", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.GetAllComplianceProgram)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewComplianceProgram", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.CreateComplianceProgram)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ComplianceProgram/{id}", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.UpdateComplianceProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteComplianceProgram/{id}", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.DeleteComplianceProgram)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToComplianceProgram/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.AssignOrganizationToComplianceProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromComplianceProgram/{parentId}", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.UnassignOrganizationFromComplianceProgram)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddRequirementsToComplianceProgram/{parentId}/requirementsId", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.AddRequirementsToComplianceProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRequirementsFromComplianceProgram/{parentId}/requirementsIds", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.RemoveRequirementsFromComplianceProgram)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddControlsToComplianceProgram/{parentId}/controlsId", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.AddControlsToComplianceProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveControlsFromComplianceProgram/{parentId}/controlsIds", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.RemoveControlsFromComplianceProgram)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAttestationsToComplianceProgram/{parentId}/attestationsId", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.AddAttestationsToComplianceProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAttestationsFromComplianceProgram/{parentId}/attestationsIds", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.RemoveAttestationsFromComplianceProgram)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRegulationsToComplianceProgram/{parentId}/regulationsId", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.AddRegulationsToComplianceProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRegulationsFromComplianceProgram/{parentId}/regulationsIds", jsonResponseFormatter.FormatToJSON(ComplianceProgramController.RemoveRegulationsFromComplianceProgram)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ComplianceRequirement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ComplianceRequirement/{id}", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.GetComplianceRequirement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ComplianceRequirement", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.GetAllComplianceRequirement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewComplianceRequirement", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.CreateComplianceRequirement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ComplianceRequirement/{id}", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.UpdateComplianceRequirement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteComplianceRequirement/{id}", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.DeleteComplianceRequirement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignComplianceProgramToComplianceRequirement/{parentId}/complianceProgramId", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.AssignComplianceProgramToComplianceRequirement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignComplianceProgramFromComplianceRequirement/{parentId}", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.UnassignComplianceProgramFromComplianceRequirement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPoliciesToComplianceRequirement/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.AddPoliciesToComplianceRequirement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromComplianceRequirement/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.RemovePoliciesFromComplianceRequirement)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddControlsToComplianceRequirement/{parentId}/controlsId", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.AddControlsToComplianceRequirement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveControlsFromComplianceRequirement/{parentId}/controlsIds", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.RemoveControlsFromComplianceRequirement)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddObligationsToComplianceRequirement/{parentId}/obligationsId", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.AddObligationsToComplianceRequirement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveObligationsFromComplianceRequirement/{parentId}/obligationsIds", jsonResponseFormatter.FormatToJSON(ComplianceRequirementController.RemoveObligationsFromComplianceRequirement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Attestation Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Attestation/{id}", jsonResponseFormatter.FormatToJSON(AttestationController.GetAttestation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Attestation", jsonResponseFormatter.FormatToJSON(AttestationController.GetAllAttestation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAttestation", jsonResponseFormatter.FormatToJSON(AttestationController.CreateAttestation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Attestation/{id}", jsonResponseFormatter.FormatToJSON(AttestationController.UpdateAttestation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAttestation/{id}", jsonResponseFormatter.FormatToJSON(AttestationController.DeleteAttestation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignControlToAttestation/{parentId}/controlId", jsonResponseFormatter.FormatToJSON(AttestationController.AssignControlToAttestation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignControlFromAttestation/{parentId}", jsonResponseFormatter.FormatToJSON(AttestationController.UnassignControlFromAttestation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPolicyToAttestation/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(AttestationController.AssignPolicyToAttestation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromAttestation/{parentId}", jsonResponseFormatter.FormatToJSON(AttestationController.UnassignPolicyFromAttestation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignComplianceProgramToAttestation/{parentId}/complianceProgramId", jsonResponseFormatter.FormatToJSON(AttestationController.AssignComplianceProgramToAttestation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignComplianceProgramFromAttestation/{parentId}", jsonResponseFormatter.FormatToJSON(AttestationController.UnassignComplianceProgramFromAttestation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // AuditProgram Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AuditProgram/{id}", jsonResponseFormatter.FormatToJSON(AuditProgramController.GetAuditProgram)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AuditProgram", jsonResponseFormatter.FormatToJSON(AuditProgramController.GetAllAuditProgram)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAuditProgram", jsonResponseFormatter.FormatToJSON(AuditProgramController.CreateAuditProgram)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AuditProgram/{id}", jsonResponseFormatter.FormatToJSON(AuditProgramController.UpdateAuditProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAuditProgram/{id}", jsonResponseFormatter.FormatToJSON(AuditProgramController.DeleteAuditProgram)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToAuditProgram/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(AuditProgramController.AssignOrganizationToAuditProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromAuditProgram/{parentId}", jsonResponseFormatter.FormatToJSON(AuditProgramController.UnassignOrganizationFromAuditProgram)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddEngagementsToAuditProgram/{parentId}/engagementsId", jsonResponseFormatter.FormatToJSON(AuditProgramController.AddEngagementsToAuditProgram)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEngagementsFromAuditProgram/{parentId}/engagementsIds", jsonResponseFormatter.FormatToJSON(AuditProgramController.RemoveEngagementsFromAuditProgram)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AuditEngagement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AuditEngagement/{id}", jsonResponseFormatter.FormatToJSON(AuditEngagementController.GetAuditEngagement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AuditEngagement", jsonResponseFormatter.FormatToJSON(AuditEngagementController.GetAllAuditEngagement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAuditEngagement", jsonResponseFormatter.FormatToJSON(AuditEngagementController.CreateAuditEngagement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AuditEngagement/{id}", jsonResponseFormatter.FormatToJSON(AuditEngagementController.UpdateAuditEngagement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAuditEngagement/{id}", jsonResponseFormatter.FormatToJSON(AuditEngagementController.DeleteAuditEngagement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAuditProgramToAuditEngagement/{parentId}/auditProgramId", jsonResponseFormatter.FormatToJSON(AuditEngagementController.AssignAuditProgramToAuditEngagement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAuditProgramFromAuditEngagement/{parentId}", jsonResponseFormatter.FormatToJSON(AuditEngagementController.UnassignAuditProgramFromAuditEngagement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddBusinessUnitsToAuditEngagement/{parentId}/businessUnitsId", jsonResponseFormatter.FormatToJSON(AuditEngagementController.AddBusinessUnitsToAuditEngagement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBusinessUnitsFromAuditEngagement/{parentId}/businessUnitsIds", jsonResponseFormatter.FormatToJSON(AuditEngagementController.RemoveBusinessUnitsFromAuditEngagement)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddControlTestsToAuditEngagement/{parentId}/controlTestsId", jsonResponseFormatter.FormatToJSON(AuditEngagementController.AddControlTestsToAuditEngagement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveControlTestsFromAuditEngagement/{parentId}/controlTestsIds", jsonResponseFormatter.FormatToJSON(AuditEngagementController.RemoveControlTestsFromAuditEngagement)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddWorkpapersToAuditEngagement/{parentId}/workpapersId", jsonResponseFormatter.FormatToJSON(AuditEngagementController.AddWorkpapersToAuditEngagement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWorkpapersFromAuditEngagement/{parentId}/workpapersIds", jsonResponseFormatter.FormatToJSON(AuditEngagementController.RemoveWorkpapersFromAuditEngagement)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFindingsToAuditEngagement/{parentId}/findingsId", jsonResponseFormatter.FormatToJSON(AuditEngagementController.AddFindingsToAuditEngagement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFindingsFromAuditEngagement/{parentId}/findingsIds", jsonResponseFormatter.FormatToJSON(AuditEngagementController.RemoveFindingsFromAuditEngagement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AuditWorkpaper Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AuditWorkpaper/{id}", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.GetAuditWorkpaper)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AuditWorkpaper", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.GetAllAuditWorkpaper)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAuditWorkpaper", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.CreateAuditWorkpaper)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AuditWorkpaper/{id}", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.UpdateAuditWorkpaper)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAuditWorkpaper/{id}", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.DeleteAuditWorkpaper)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEngagementToAuditWorkpaper/{parentId}/engagementId", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.AssignEngagementToAuditWorkpaper)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEngagementFromAuditWorkpaper/{parentId}", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.UnassignEngagementFromAuditWorkpaper)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddEvidenceToAuditWorkpaper/{parentId}/evidenceId", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.AddEvidenceToAuditWorkpaper)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEvidenceFromAuditWorkpaper/{parentId}/evidenceIds", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.RemoveEvidenceFromAuditWorkpaper)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFindingsToAuditWorkpaper/{parentId}/findingsId", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.AddFindingsToAuditWorkpaper)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFindingsFromAuditWorkpaper/{parentId}/findingsIds", jsonResponseFormatter.FormatToJSON(AuditWorkpaperController.RemoveFindingsFromAuditWorkpaper)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AuditFinding Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AuditFinding/{id}", jsonResponseFormatter.FormatToJSON(AuditFindingController.GetAuditFinding)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AuditFinding", jsonResponseFormatter.FormatToJSON(AuditFindingController.GetAllAuditFinding)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAuditFinding", jsonResponseFormatter.FormatToJSON(AuditFindingController.CreateAuditFinding)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AuditFinding/{id}", jsonResponseFormatter.FormatToJSON(AuditFindingController.UpdateAuditFinding)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAuditFinding/{id}", jsonResponseFormatter.FormatToJSON(AuditFindingController.DeleteAuditFinding)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEngagementToAuditFinding/{parentId}/engagementId", jsonResponseFormatter.FormatToJSON(AuditFindingController.AssignEngagementToAuditFinding)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEngagementFromAuditFinding/{parentId}", jsonResponseFormatter.FormatToJSON(AuditFindingController.UnassignEngagementFromAuditFinding)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkpaperToAuditFinding/{parentId}/workpaperId", jsonResponseFormatter.FormatToJSON(AuditFindingController.AssignWorkpaperToAuditFinding)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkpaperFromAuditFinding/{parentId}", jsonResponseFormatter.FormatToJSON(AuditFindingController.UnassignWorkpaperFromAuditFinding)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCorrectiveActionsToAuditFinding/{parentId}/correctiveActionsId", jsonResponseFormatter.FormatToJSON(AuditFindingController.AddCorrectiveActionsToAuditFinding)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCorrectiveActionsFromAuditFinding/{parentId}/correctiveActionsIds", jsonResponseFormatter.FormatToJSON(AuditFindingController.RemoveCorrectiveActionsFromAuditFinding)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRelatedRisksToAuditFinding/{parentId}/relatedRisksId", jsonResponseFormatter.FormatToJSON(AuditFindingController.AddRelatedRisksToAuditFinding)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRelatedRisksFromAuditFinding/{parentId}/relatedRisksIds", jsonResponseFormatter.FormatToJSON(AuditFindingController.RemoveRelatedRisksFromAuditFinding)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRelatedControlsToAuditFinding/{parentId}/relatedControlsId", jsonResponseFormatter.FormatToJSON(AuditFindingController.AddRelatedControlsToAuditFinding)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRelatedControlsFromAuditFinding/{parentId}/relatedControlsIds", jsonResponseFormatter.FormatToJSON(AuditFindingController.RemoveRelatedControlsFromAuditFinding)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddIssuesToAuditFinding/{parentId}/issuesId", jsonResponseFormatter.FormatToJSON(AuditFindingController.AddIssuesToAuditFinding)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveIssuesFromAuditFinding/{parentId}/issuesIds", jsonResponseFormatter.FormatToJSON(AuditFindingController.RemoveIssuesFromAuditFinding)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CorrectiveAction Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CorrectiveAction/{id}", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.GetCorrectiveAction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CorrectiveAction", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.GetAllCorrectiveAction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCorrectiveAction", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.CreateCorrectiveAction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CorrectiveAction/{id}", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.UpdateCorrectiveAction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCorrectiveAction/{id}", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.DeleteCorrectiveAction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignFindingToCorrectiveAction/{parentId}/findingId", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.AssignFindingToCorrectiveAction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFindingFromCorrectiveAction/{parentId}", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.UnassignFindingFromCorrectiveAction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignIssueToCorrectiveAction/{parentId}/issueId", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.AssignIssueToCorrectiveAction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignIssueFromCorrectiveAction/{parentId}", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.UnassignIssueFromCorrectiveAction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Issue Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Issue/{id}", jsonResponseFormatter.FormatToJSON(IssueController.GetIssue)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Issue", jsonResponseFormatter.FormatToJSON(IssueController.GetAllIssue)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewIssue", jsonResponseFormatter.FormatToJSON(IssueController.CreateIssue)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Issue/{id}", jsonResponseFormatter.FormatToJSON(IssueController.UpdateIssue)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteIssue/{id}", jsonResponseFormatter.FormatToJSON(IssueController.DeleteIssue)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRiskToIssue/{parentId}/riskId", jsonResponseFormatter.FormatToJSON(IssueController.AssignRiskToIssue)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRiskFromIssue/{parentId}", jsonResponseFormatter.FormatToJSON(IssueController.UnassignRiskFromIssue)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignFindingToIssue/{parentId}/findingId", jsonResponseFormatter.FormatToJSON(IssueController.AssignFindingToIssue)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFindingFromIssue/{parentId}", jsonResponseFormatter.FormatToJSON(IssueController.UnassignFindingFromIssue)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignControlToIssue/{parentId}/controlId", jsonResponseFormatter.FormatToJSON(IssueController.AssignControlToIssue)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignControlFromIssue/{parentId}", jsonResponseFormatter.FormatToJSON(IssueController.UnassignControlFromIssue)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCorrectiveActionsToIssue/{parentId}/correctiveActionsId", jsonResponseFormatter.FormatToJSON(IssueController.AddCorrectiveActionsToIssue)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCorrectiveActionsFromIssue/{parentId}/correctiveActionsIds", jsonResponseFormatter.FormatToJSON(IssueController.RemoveCorrectiveActionsFromIssue)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // BusinessUnit Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BusinessUnit/{id}", jsonResponseFormatter.FormatToJSON(BusinessUnitController.GetBusinessUnit)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BusinessUnit", jsonResponseFormatter.FormatToJSON(BusinessUnitController.GetAllBusinessUnit)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBusinessUnit", jsonResponseFormatter.FormatToJSON(BusinessUnitController.CreateBusinessUnit)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BusinessUnit/{id}", jsonResponseFormatter.FormatToJSON(BusinessUnitController.UpdateBusinessUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBusinessUnit/{id}", jsonResponseFormatter.FormatToJSON(BusinessUnitController.DeleteBusinessUnit)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToBusinessUnit/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(BusinessUnitController.AssignOrganizationToBusinessUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromBusinessUnit/{parentId}", jsonResponseFormatter.FormatToJSON(BusinessUnitController.UnassignOrganizationFromBusinessUnit)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAuditsToBusinessUnit/{parentId}/auditsId", jsonResponseFormatter.FormatToJSON(BusinessUnitController.AddAuditsToBusinessUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAuditsFromBusinessUnit/{parentId}/auditsIds", jsonResponseFormatter.FormatToJSON(BusinessUnitController.RemoveAuditsFromBusinessUnit)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DataProcessingActivity Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DataProcessingActivity/{id}", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.GetDataProcessingActivity)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DataProcessingActivity", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.GetAllDataProcessingActivity)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDataProcessingActivity", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.CreateDataProcessingActivity)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DataProcessingActivity/{id}", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.UpdateDataProcessingActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDataProcessingActivity/{id}", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.DeleteDataProcessingActivity)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToDataProcessingActivity/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.AssignOrganizationToDataProcessingActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromDataProcessingActivity/{parentId}", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.UnassignOrganizationFromDataProcessingActivity)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDataCategoriesToDataProcessingActivity/{parentId}/dataCategoriesId", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.AddDataCategoriesToDataProcessingActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataCategoriesFromDataProcessingActivity/{parentId}/dataCategoriesIds", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.RemoveDataCategoriesFromDataProcessingActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSystemsToDataProcessingActivity/{parentId}/systemsId", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.AddSystemsToDataProcessingActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSystemsFromDataProcessingActivity/{parentId}/systemsIds", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.RemoveSystemsFromDataProcessingActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRecordsToDataProcessingActivity/{parentId}/recordsId", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.AddRecordsToDataProcessingActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRecordsFromDataProcessingActivity/{parentId}/recordsIds", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.RemoveRecordsFromDataProcessingActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPrivacyNoticesToDataProcessingActivity/{parentId}/privacyNoticesId", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.AddPrivacyNoticesToDataProcessingActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePrivacyNoticesFromDataProcessingActivity/{parentId}/privacyNoticesIds", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.RemovePrivacyNoticesFromDataProcessingActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddThirdPartiesToDataProcessingActivity/{parentId}/thirdPartiesId", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.AddThirdPartiesToDataProcessingActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveThirdPartiesFromDataProcessingActivity/{parentId}/thirdPartiesIds", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.RemoveThirdPartiesFromDataProcessingActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddConsentsToDataProcessingActivity/{parentId}/consentsId", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.AddConsentsToDataProcessingActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveConsentsFromDataProcessingActivity/{parentId}/consentsIds", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.RemoveConsentsFromDataProcessingActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataBreachesToDataProcessingActivity/{parentId}/dataBreachesId", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.AddDataBreachesToDataProcessingActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataBreachesFromDataProcessingActivity/{parentId}/dataBreachesIds", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.RemoveDataBreachesFromDataProcessingActivity)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataSubjectRequestsToDataProcessingActivity/{parentId}/dataSubjectRequestsId", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.AddDataSubjectRequestsToDataProcessingActivity)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataSubjectRequestsFromDataProcessingActivity/{parentId}/dataSubjectRequestsIds", jsonResponseFormatter.FormatToJSON(DataProcessingActivityController.RemoveDataSubjectRequestsFromDataProcessingActivity)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DataCategory Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DataCategory/{id}", jsonResponseFormatter.FormatToJSON(DataCategoryController.GetDataCategory)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DataCategory", jsonResponseFormatter.FormatToJSON(DataCategoryController.GetAllDataCategory)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDataCategory", jsonResponseFormatter.FormatToJSON(DataCategoryController.CreateDataCategory)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DataCategory/{id}", jsonResponseFormatter.FormatToJSON(DataCategoryController.UpdateDataCategory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDataCategory/{id}", jsonResponseFormatter.FormatToJSON(DataCategoryController.DeleteDataCategory)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProcessingActivitiesToDataCategory/{parentId}/processingActivitiesId", jsonResponseFormatter.FormatToJSON(DataCategoryController.AddProcessingActivitiesToDataCategory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProcessingActivitiesFromDataCategory/{parentId}/processingActivitiesIds", jsonResponseFormatter.FormatToJSON(DataCategoryController.RemoveProcessingActivitiesFromDataCategory)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRecordsToDataCategory/{parentId}/recordsId", jsonResponseFormatter.FormatToJSON(DataCategoryController.AddRecordsToDataCategory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRecordsFromDataCategory/{parentId}/recordsIds", jsonResponseFormatter.FormatToJSON(DataCategoryController.RemoveRecordsFromDataCategory)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataBreachesToDataCategory/{parentId}/dataBreachesId", jsonResponseFormatter.FormatToJSON(DataCategoryController.AddDataBreachesToDataCategory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataBreachesFromDataCategory/{parentId}/dataBreachesIds", jsonResponseFormatter.FormatToJSON(DataCategoryController.RemoveDataBreachesFromDataCategory)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // System_ Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/System_/{id}", jsonResponseFormatter.FormatToJSON(System_Controller.GetSystem_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/System_", jsonResponseFormatter.FormatToJSON(System_Controller.GetAllSystem_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSystem_", jsonResponseFormatter.FormatToJSON(System_Controller.CreateSystem_)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/System_/{id}", jsonResponseFormatter.FormatToJSON(System_Controller.UpdateSystem_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSystem_/{id}", jsonResponseFormatter.FormatToJSON(System_Controller.DeleteSystem_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProcessingActivitiesToSystem_/{parentId}/processingActivitiesId", jsonResponseFormatter.FormatToJSON(System_Controller.AddProcessingActivitiesToSystem_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProcessingActivitiesFromSystem_/{parentId}/processingActivitiesIds", jsonResponseFormatter.FormatToJSON(System_Controller.RemoveProcessingActivitiesFromSystem_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRecordsRepositoriesToSystem_/{parentId}/recordsRepositoriesId", jsonResponseFormatter.FormatToJSON(System_Controller.AddRecordsRepositoriesToSystem_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRecordsRepositoriesFromSystem_/{parentId}/recordsRepositoriesIds", jsonResponseFormatter.FormatToJSON(System_Controller.RemoveRecordsRepositoriesFromSystem_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PrivacyNotice Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PrivacyNotice/{id}", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.GetPrivacyNotice)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PrivacyNotice", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.GetAllPrivacyNotice)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPrivacyNotice", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.CreatePrivacyNotice)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PrivacyNotice/{id}", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.UpdatePrivacyNotice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePrivacyNotice/{id}", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.DeletePrivacyNotice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToPrivacyNotice/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.AssignOrganizationToPrivacyNotice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromPrivacyNotice/{parentId}", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.UnassignOrganizationFromPrivacyNotice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProcessingActivitiesToPrivacyNotice/{parentId}/processingActivitiesId", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.AddProcessingActivitiesToPrivacyNotice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProcessingActivitiesFromPrivacyNotice/{parentId}/processingActivitiesIds", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.RemoveProcessingActivitiesFromPrivacyNotice)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddConsentsToPrivacyNotice/{parentId}/consentsId", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.AddConsentsToPrivacyNotice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveConsentsFromPrivacyNotice/{parentId}/consentsIds", jsonResponseFormatter.FormatToJSON(PrivacyNoticeController.RemoveConsentsFromPrivacyNotice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DataSubjectRequest Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DataSubjectRequest/{id}", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.GetDataSubjectRequest)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DataSubjectRequest", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.GetAllDataSubjectRequest)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDataSubjectRequest", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.CreateDataSubjectRequest)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DataSubjectRequest/{id}", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.UpdateDataSubjectRequest)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDataSubjectRequest/{id}", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.DeleteDataSubjectRequest)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToDataSubjectRequest/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.AssignOrganizationToDataSubjectRequest)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromDataSubjectRequest/{parentId}", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.UnassignOrganizationFromDataSubjectRequest)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProcessingActivitiesToDataSubjectRequest/{parentId}/processingActivitiesId", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.AddProcessingActivitiesToDataSubjectRequest)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProcessingActivitiesFromDataSubjectRequest/{parentId}/processingActivitiesIds", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.RemoveProcessingActivitiesFromDataSubjectRequest)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRecordsToDataSubjectRequest/{parentId}/recordsId", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.AddRecordsToDataSubjectRequest)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRecordsFromDataSubjectRequest/{parentId}/recordsIds", jsonResponseFormatter.FormatToJSON(DataSubjectRequestController.RemoveRecordsFromDataSubjectRequest)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // RecordsRepository Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RecordsRepository/{id}", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.GetRecordsRepository)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RecordsRepository", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.GetAllRecordsRepository)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRecordsRepository", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.CreateRecordsRepository)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RecordsRepository/{id}", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.UpdateRecordsRepository)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRecordsRepository/{id}", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.DeleteRecordsRepository)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToRecordsRepository/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.AssignOrganizationToRecordsRepository)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromRecordsRepository/{parentId}", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.UnassignOrganizationFromRecordsRepository)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddRecordsToRecordsRepository/{parentId}/recordsId", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.AddRecordsToRecordsRepository)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRecordsFromRecordsRepository/{parentId}/recordsIds", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.RemoveRecordsFromRecordsRepository)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSystemsToRecordsRepository/{parentId}/systemsId", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.AddSystemsToRecordsRepository)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSystemsFromRecordsRepository/{parentId}/systemsIds", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.RemoveSystemsFromRecordsRepository)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRetentionSchedulesToRecordsRepository/{parentId}/retentionSchedulesId", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.AddRetentionSchedulesToRecordsRepository)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRetentionSchedulesFromRecordsRepository/{parentId}/retentionSchedulesIds", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.RemoveRetentionSchedulesFromRecordsRepository)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLegalHoldsToRecordsRepository/{parentId}/legalHoldsId", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.AddLegalHoldsToRecordsRepository)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLegalHoldsFromRecordsRepository/{parentId}/legalHoldsIds", jsonResponseFormatter.FormatToJSON(RecordsRepositoryController.RemoveLegalHoldsFromRecordsRepository)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Record_ Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Record_/{id}", jsonResponseFormatter.FormatToJSON(Record_Controller.GetRecord_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Record_", jsonResponseFormatter.FormatToJSON(Record_Controller.GetAllRecord_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRecord_", jsonResponseFormatter.FormatToJSON(Record_Controller.CreateRecord_)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Record_/{id}", jsonResponseFormatter.FormatToJSON(Record_Controller.UpdateRecord_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRecord_/{id}", jsonResponseFormatter.FormatToJSON(Record_Controller.DeleteRecord_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRepositoryToRecord_/{parentId}/repositoryId", jsonResponseFormatter.FormatToJSON(Record_Controller.AssignRepositoryToRecord_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRepositoryFromRecord_/{parentId}", jsonResponseFormatter.FormatToJSON(Record_Controller.UnassignRepositoryFromRecord_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRetentionScheduleToRecord_/{parentId}/retentionScheduleId", jsonResponseFormatter.FormatToJSON(Record_Controller.AssignRetentionScheduleToRecord_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRetentionScheduleFromRecord_/{parentId}", jsonResponseFormatter.FormatToJSON(Record_Controller.UnassignRetentionScheduleFromRecord_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProcessingActivitiesToRecord_/{parentId}/processingActivitiesId", jsonResponseFormatter.FormatToJSON(Record_Controller.AddProcessingActivitiesToRecord_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProcessingActivitiesFromRecord_/{parentId}/processingActivitiesIds", jsonResponseFormatter.FormatToJSON(Record_Controller.RemoveProcessingActivitiesFromRecord_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataCategoriesToRecord_/{parentId}/dataCategoriesId", jsonResponseFormatter.FormatToJSON(Record_Controller.AddDataCategoriesToRecord_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataCategoriesFromRecord_/{parentId}/dataCategoriesIds", jsonResponseFormatter.FormatToJSON(Record_Controller.RemoveDataCategoriesFromRecord_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLegalHoldsToRecord_/{parentId}/legalHoldsId", jsonResponseFormatter.FormatToJSON(Record_Controller.AddLegalHoldsToRecord_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLegalHoldsFromRecord_/{parentId}/legalHoldsIds", jsonResponseFormatter.FormatToJSON(Record_Controller.RemoveLegalHoldsFromRecord_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataSubjectRequestsToRecord_/{parentId}/dataSubjectRequestsId", jsonResponseFormatter.FormatToJSON(Record_Controller.AddDataSubjectRequestsToRecord_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataSubjectRequestsFromRecord_/{parentId}/dataSubjectRequestsIds", jsonResponseFormatter.FormatToJSON(Record_Controller.RemoveDataSubjectRequestsFromRecord_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // RetentionSchedule Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RetentionSchedule/{id}", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.GetRetentionSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RetentionSchedule", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.GetAllRetentionSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRetentionSchedule", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.CreateRetentionSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RetentionSchedule/{id}", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.UpdateRetentionSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRetentionSchedule/{id}", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.DeleteRetentionSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddRepositoriesToRetentionSchedule/{parentId}/repositoriesId", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.AddRepositoriesToRetentionSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRepositoriesFromRetentionSchedule/{parentId}/repositoriesIds", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.RemoveRepositoriesFromRetentionSchedule)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRecordsToRetentionSchedule/{parentId}/recordsId", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.AddRecordsToRetentionSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRecordsFromRetentionSchedule/{parentId}/recordsIds", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.RemoveRecordsFromRetentionSchedule)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddExceptionsToRetentionSchedule/{parentId}/exceptionsId", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.AddExceptionsToRetentionSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveExceptionsFromRetentionSchedule/{parentId}/exceptionsIds", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.RemoveExceptionsFromRetentionSchedule)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDispositionReviewsToRetentionSchedule/{parentId}/dispositionReviewsId", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.AddDispositionReviewsToRetentionSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDispositionReviewsFromRetentionSchedule/{parentId}/dispositionReviewsIds", jsonResponseFormatter.FormatToJSON(RetentionScheduleController.RemoveDispositionReviewsFromRetentionSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DispositionReview Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DispositionReview/{id}", jsonResponseFormatter.FormatToJSON(DispositionReviewController.GetDispositionReview)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DispositionReview", jsonResponseFormatter.FormatToJSON(DispositionReviewController.GetAllDispositionReview)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDispositionReview", jsonResponseFormatter.FormatToJSON(DispositionReviewController.CreateDispositionReview)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DispositionReview/{id}", jsonResponseFormatter.FormatToJSON(DispositionReviewController.UpdateDispositionReview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDispositionReview/{id}", jsonResponseFormatter.FormatToJSON(DispositionReviewController.DeleteDispositionReview)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRecordToDispositionReview/{parentId}/recordId", jsonResponseFormatter.FormatToJSON(DispositionReviewController.AssignRecordToDispositionReview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRecordFromDispositionReview/{parentId}", jsonResponseFormatter.FormatToJSON(DispositionReviewController.UnassignRecordFromDispositionReview)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRetentionScheduleToDispositionReview/{parentId}/retentionScheduleId", jsonResponseFormatter.FormatToJSON(DispositionReviewController.AssignRetentionScheduleToDispositionReview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRetentionScheduleFromDispositionReview/{parentId}", jsonResponseFormatter.FormatToJSON(DispositionReviewController.UnassignRetentionScheduleFromDispositionReview)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // LegalHold Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LegalHold/{id}", jsonResponseFormatter.FormatToJSON(LegalHoldController.GetLegalHold)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LegalHold", jsonResponseFormatter.FormatToJSON(LegalHoldController.GetAllLegalHold)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLegalHold", jsonResponseFormatter.FormatToJSON(LegalHoldController.CreateLegalHold)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LegalHold/{id}", jsonResponseFormatter.FormatToJSON(LegalHoldController.UpdateLegalHold)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLegalHold/{id}", jsonResponseFormatter.FormatToJSON(LegalHoldController.DeleteLegalHold)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignMatterToLegalHold/{parentId}/matterId", jsonResponseFormatter.FormatToJSON(LegalHoldController.AssignMatterToLegalHold)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMatterFromLegalHold/{parentId}", jsonResponseFormatter.FormatToJSON(LegalHoldController.UnassignMatterFromLegalHold)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddRepositoriesToLegalHold/{parentId}/repositoriesId", jsonResponseFormatter.FormatToJSON(LegalHoldController.AddRepositoriesToLegalHold)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRepositoriesFromLegalHold/{parentId}/repositoriesIds", jsonResponseFormatter.FormatToJSON(LegalHoldController.RemoveRepositoriesFromLegalHold)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRecordsToLegalHold/{parentId}/recordsId", jsonResponseFormatter.FormatToJSON(LegalHoldController.AddRecordsToLegalHold)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRecordsFromLegalHold/{parentId}/recordsIds", jsonResponseFormatter.FormatToJSON(LegalHoldController.RemoveRecordsFromLegalHold)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Matter Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Matter/{id}", jsonResponseFormatter.FormatToJSON(MatterController.GetMatter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Matter", jsonResponseFormatter.FormatToJSON(MatterController.GetAllMatter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMatter", jsonResponseFormatter.FormatToJSON(MatterController.CreateMatter)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Matter/{id}", jsonResponseFormatter.FormatToJSON(MatterController.UpdateMatter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMatter/{id}", jsonResponseFormatter.FormatToJSON(MatterController.DeleteMatter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToMatter/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(MatterController.AssignOrganizationToMatter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromMatter/{parentId}", jsonResponseFormatter.FormatToJSON(MatterController.UnassignOrganizationFromMatter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLegalHoldsToMatter/{parentId}/legalHoldsId", jsonResponseFormatter.FormatToJSON(MatterController.AddLegalHoldsToMatter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLegalHoldsFromMatter/{parentId}/legalHoldsIds", jsonResponseFormatter.FormatToJSON(MatterController.RemoveLegalHoldsFromMatter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataBreachesToMatter/{parentId}/dataBreachesId", jsonResponseFormatter.FormatToJSON(MatterController.AddDataBreachesToMatter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataBreachesFromMatter/{parentId}/dataBreachesIds", jsonResponseFormatter.FormatToJSON(MatterController.RemoveDataBreachesFromMatter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddContractsToMatter/{parentId}/contractsId", jsonResponseFormatter.FormatToJSON(MatterController.AddContractsToMatter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContractsFromMatter/{parentId}/contractsIds", jsonResponseFormatter.FormatToJSON(MatterController.RemoveContractsFromMatter)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AssignOrganizationToThirdParty/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(ThirdPartyController.AssignOrganizationToThirdParty)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromThirdParty/{parentId}", jsonResponseFormatter.FormatToJSON(ThirdPartyController.UnassignOrganizationFromThirdParty)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProcessingActivitiesToThirdParty/{parentId}/processingActivitiesId", jsonResponseFormatter.FormatToJSON(ThirdPartyController.AddProcessingActivitiesToThirdParty)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProcessingActivitiesFromThirdParty/{parentId}/processingActivitiesIds", jsonResponseFormatter.FormatToJSON(ThirdPartyController.RemoveProcessingActivitiesFromThirdParty)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAssessmentsToThirdParty/{parentId}/assessmentsId", jsonResponseFormatter.FormatToJSON(ThirdPartyController.AddAssessmentsToThirdParty)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAssessmentsFromThirdParty/{parentId}/assessmentsIds", jsonResponseFormatter.FormatToJSON(ThirdPartyController.RemoveAssessmentsFromThirdParty)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddContractsToThirdParty/{parentId}/contractsId", jsonResponseFormatter.FormatToJSON(ThirdPartyController.AddContractsToThirdParty)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContractsFromThirdParty/{parentId}/contractsIds", jsonResponseFormatter.FormatToJSON(ThirdPartyController.RemoveContractsFromThirdParty)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddObligationsToThirdParty/{parentId}/obligationsId", jsonResponseFormatter.FormatToJSON(ThirdPartyController.AddObligationsToThirdParty)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveObligationsFromThirdParty/{parentId}/obligationsIds", jsonResponseFormatter.FormatToJSON(ThirdPartyController.RemoveObligationsFromThirdParty)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataBreachesToThirdParty/{parentId}/dataBreachesId", jsonResponseFormatter.FormatToJSON(ThirdPartyController.AddDataBreachesToThirdParty)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataBreachesFromThirdParty/{parentId}/dataBreachesIds", jsonResponseFormatter.FormatToJSON(ThirdPartyController.RemoveDataBreachesFromThirdParty)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ThirdPartyAssessment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ThirdPartyAssessment/{id}", jsonResponseFormatter.FormatToJSON(ThirdPartyAssessmentController.GetThirdPartyAssessment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyAssessment", jsonResponseFormatter.FormatToJSON(ThirdPartyAssessmentController.GetAllThirdPartyAssessment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewThirdPartyAssessment", jsonResponseFormatter.FormatToJSON(ThirdPartyAssessmentController.CreateThirdPartyAssessment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ThirdPartyAssessment/{id}", jsonResponseFormatter.FormatToJSON(ThirdPartyAssessmentController.UpdateThirdPartyAssessment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteThirdPartyAssessment/{id}", jsonResponseFormatter.FormatToJSON(ThirdPartyAssessmentController.DeleteThirdPartyAssessment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignThirdPartyToThirdPartyAssessment/{parentId}/thirdPartyId", jsonResponseFormatter.FormatToJSON(ThirdPartyAssessmentController.AssignThirdPartyToThirdPartyAssessment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignThirdPartyFromThirdPartyAssessment/{parentId}", jsonResponseFormatter.FormatToJSON(ThirdPartyAssessmentController.UnassignThirdPartyFromThirdPartyAssessment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddIssuesToThirdPartyAssessment/{parentId}/issuesId", jsonResponseFormatter.FormatToJSON(ThirdPartyAssessmentController.AddIssuesToThirdPartyAssessment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveIssuesFromThirdPartyAssessment/{parentId}/issuesIds", jsonResponseFormatter.FormatToJSON(ThirdPartyAssessmentController.RemoveIssuesFromThirdPartyAssessment)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AssignThirdPartyToContract/{parentId}/thirdPartyId", jsonResponseFormatter.FormatToJSON(ContractController.AssignThirdPartyToContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignThirdPartyFromContract/{parentId}", jsonResponseFormatter.FormatToJSON(ContractController.UnassignThirdPartyFromContract)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignMatterToContract/{parentId}/matterId", jsonResponseFormatter.FormatToJSON(ContractController.AssignMatterToContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMatterFromContract/{parentId}", jsonResponseFormatter.FormatToJSON(ContractController.UnassignMatterFromContract)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddObligationsToContract/{parentId}/obligationsId", jsonResponseFormatter.FormatToJSON(ContractController.AddObligationsToContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveObligationsFromContract/{parentId}/obligationsIds", jsonResponseFormatter.FormatToJSON(ContractController.RemoveObligationsFromContract)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataProcessingActivitiesToContract/{parentId}/dataProcessingActivitiesId", jsonResponseFormatter.FormatToJSON(ContractController.AddDataProcessingActivitiesToContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataProcessingActivitiesFromContract/{parentId}/dataProcessingActivitiesIds", jsonResponseFormatter.FormatToJSON(ContractController.RemoveDataProcessingActivitiesFromContract)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Exception_ Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Exception_/{id}", jsonResponseFormatter.FormatToJSON(Exception_Controller.GetException_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Exception_", jsonResponseFormatter.FormatToJSON(Exception_Controller.GetAllException_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewException_", jsonResponseFormatter.FormatToJSON(Exception_Controller.CreateException_)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Exception_/{id}", jsonResponseFormatter.FormatToJSON(Exception_Controller.UpdateException_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteException_/{id}", jsonResponseFormatter.FormatToJSON(Exception_Controller.DeleteException_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRetentionScheduleToException_/{parentId}/retentionScheduleId", jsonResponseFormatter.FormatToJSON(Exception_Controller.AssignRetentionScheduleToException_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRetentionScheduleFromException_/{parentId}", jsonResponseFormatter.FormatToJSON(Exception_Controller.UnassignRetentionScheduleFromException_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPolicyToException_/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(Exception_Controller.AssignPolicyToException_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromException_/{parentId}", jsonResponseFormatter.FormatToJSON(Exception_Controller.UnassignPolicyFromException_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignControlToException_/{parentId}/controlId", jsonResponseFormatter.FormatToJSON(Exception_Controller.AssignControlToException_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignControlFromException_/{parentId}", jsonResponseFormatter.FormatToJSON(Exception_Controller.UnassignControlFromException_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRiskToException_/{parentId}/riskId", jsonResponseFormatter.FormatToJSON(Exception_Controller.AssignRiskToException_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRiskFromException_/{parentId}", jsonResponseFormatter.FormatToJSON(Exception_Controller.UnassignRiskFromException_)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AssignPrivacyNoticeToConsent/{parentId}/privacyNoticeId", jsonResponseFormatter.FormatToJSON(ConsentController.AssignPrivacyNoticeToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPrivacyNoticeFromConsent/{parentId}", jsonResponseFormatter.FormatToJSON(ConsentController.UnassignPrivacyNoticeFromConsent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProcessingActivitiesToConsent/{parentId}/processingActivitiesId", jsonResponseFormatter.FormatToJSON(ConsentController.AddProcessingActivitiesToConsent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProcessingActivitiesFromConsent/{parentId}/processingActivitiesIds", jsonResponseFormatter.FormatToJSON(ConsentController.RemoveProcessingActivitiesFromConsent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DataBreach Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DataBreach/{id}", jsonResponseFormatter.FormatToJSON(DataBreachController.GetDataBreach)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DataBreach", jsonResponseFormatter.FormatToJSON(DataBreachController.GetAllDataBreach)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDataBreach", jsonResponseFormatter.FormatToJSON(DataBreachController.CreateDataBreach)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DataBreach/{id}", jsonResponseFormatter.FormatToJSON(DataBreachController.UpdateDataBreach)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDataBreach/{id}", jsonResponseFormatter.FormatToJSON(DataBreachController.DeleteDataBreach)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToDataBreach/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(DataBreachController.AssignOrganizationToDataBreach)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromDataBreach/{parentId}", jsonResponseFormatter.FormatToJSON(DataBreachController.UnassignOrganizationFromDataBreach)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignMatterToDataBreach/{parentId}/matterId", jsonResponseFormatter.FormatToJSON(DataBreachController.AssignMatterToDataBreach)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMatterFromDataBreach/{parentId}", jsonResponseFormatter.FormatToJSON(DataBreachController.UnassignMatterFromDataBreach)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProcessingActivitiesToDataBreach/{parentId}/processingActivitiesId", jsonResponseFormatter.FormatToJSON(DataBreachController.AddProcessingActivitiesToDataBreach)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProcessingActivitiesFromDataBreach/{parentId}/processingActivitiesIds", jsonResponseFormatter.FormatToJSON(DataBreachController.RemoveProcessingActivitiesFromDataBreach)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataCategoriesToDataBreach/{parentId}/dataCategoriesId", jsonResponseFormatter.FormatToJSON(DataBreachController.AddDataCategoriesToDataBreach)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataCategoriesFromDataBreach/{parentId}/dataCategoriesIds", jsonResponseFormatter.FormatToJSON(DataBreachController.RemoveDataCategoriesFromDataBreach)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddThirdPartiesToDataBreach/{parentId}/thirdPartiesId", jsonResponseFormatter.FormatToJSON(DataBreachController.AddThirdPartiesToDataBreach)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveThirdPartiesFromDataBreach/{parentId}/thirdPartiesIds", jsonResponseFormatter.FormatToJSON(DataBreachController.RemoveThirdPartiesFromDataBreach)).Methods("DELETE", "OPTIONS")

    return router
}
