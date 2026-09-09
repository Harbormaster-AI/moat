import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListOrganizationComponent from './components/ListOrganizationComponent';
import CreateOrganizationComponent from './components/CreateOrganizationComponent';
import ViewOrganizationComponent from './components/ViewOrganizationComponent';
import ListGovernanceBodyComponent from './components/ListGovernanceBodyComponent';
import CreateGovernanceBodyComponent from './components/CreateGovernanceBodyComponent';
import ViewGovernanceBodyComponent from './components/ViewGovernanceBodyComponent';
import ListPersonComponent from './components/ListPersonComponent';
import CreatePersonComponent from './components/CreatePersonComponent';
import ViewPersonComponent from './components/ViewPersonComponent';
import ListRoleComponent from './components/ListRoleComponent';
import CreateRoleComponent from './components/CreateRoleComponent';
import ViewRoleComponent from './components/ViewRoleComponent';
import ListRoleAssignmentComponent from './components/ListRoleAssignmentComponent';
import CreateRoleAssignmentComponent from './components/CreateRoleAssignmentComponent';
import ViewRoleAssignmentComponent from './components/ViewRoleAssignmentComponent';
import ListPolicyComponent from './components/ListPolicyComponent';
import CreatePolicyComponent from './components/CreatePolicyComponent';
import ViewPolicyComponent from './components/ViewPolicyComponent';
import ListProcedureComponent from './components/ListProcedureComponent';
import CreateProcedureComponent from './components/CreateProcedureComponent';
import ViewProcedureComponent from './components/ViewProcedureComponent';
import ListRegulationComponent from './components/ListRegulationComponent';
import CreateRegulationComponent from './components/CreateRegulationComponent';
import ViewRegulationComponent from './components/ViewRegulationComponent';
import ListObligationComponent from './components/ListObligationComponent';
import CreateObligationComponent from './components/CreateObligationComponent';
import ViewObligationComponent from './components/ViewObligationComponent';
import ListControlComponent from './components/ListControlComponent';
import CreateControlComponent from './components/CreateControlComponent';
import ViewControlComponent from './components/ViewControlComponent';
import ListControlTest_Component from './components/ListControlTest_Component';
import CreateControlTest_Component from './components/CreateControlTest_Component';
import ViewControlTest_Component from './components/ViewControlTest_Component';
import ListEvidenceComponent from './components/ListEvidenceComponent';
import CreateEvidenceComponent from './components/CreateEvidenceComponent';
import ViewEvidenceComponent from './components/ViewEvidenceComponent';
import ListRiskComponent from './components/ListRiskComponent';
import CreateRiskComponent from './components/CreateRiskComponent';
import ViewRiskComponent from './components/ViewRiskComponent';
import ListRiskAssessmentComponent from './components/ListRiskAssessmentComponent';
import CreateRiskAssessmentComponent from './components/CreateRiskAssessmentComponent';
import ViewRiskAssessmentComponent from './components/ViewRiskAssessmentComponent';
import ListComplianceProgramComponent from './components/ListComplianceProgramComponent';
import CreateComplianceProgramComponent from './components/CreateComplianceProgramComponent';
import ViewComplianceProgramComponent from './components/ViewComplianceProgramComponent';
import ListComplianceRequirementComponent from './components/ListComplianceRequirementComponent';
import CreateComplianceRequirementComponent from './components/CreateComplianceRequirementComponent';
import ViewComplianceRequirementComponent from './components/ViewComplianceRequirementComponent';
import ListAttestationComponent from './components/ListAttestationComponent';
import CreateAttestationComponent from './components/CreateAttestationComponent';
import ViewAttestationComponent from './components/ViewAttestationComponent';
import ListAuditProgramComponent from './components/ListAuditProgramComponent';
import CreateAuditProgramComponent from './components/CreateAuditProgramComponent';
import ViewAuditProgramComponent from './components/ViewAuditProgramComponent';
import ListAuditEngagementComponent from './components/ListAuditEngagementComponent';
import CreateAuditEngagementComponent from './components/CreateAuditEngagementComponent';
import ViewAuditEngagementComponent from './components/ViewAuditEngagementComponent';
import ListAuditWorkpaperComponent from './components/ListAuditWorkpaperComponent';
import CreateAuditWorkpaperComponent from './components/CreateAuditWorkpaperComponent';
import ViewAuditWorkpaperComponent from './components/ViewAuditWorkpaperComponent';
import ListAuditFindingComponent from './components/ListAuditFindingComponent';
import CreateAuditFindingComponent from './components/CreateAuditFindingComponent';
import ViewAuditFindingComponent from './components/ViewAuditFindingComponent';
import ListCorrectiveActionComponent from './components/ListCorrectiveActionComponent';
import CreateCorrectiveActionComponent from './components/CreateCorrectiveActionComponent';
import ViewCorrectiveActionComponent from './components/ViewCorrectiveActionComponent';
import ListIssueComponent from './components/ListIssueComponent';
import CreateIssueComponent from './components/CreateIssueComponent';
import ViewIssueComponent from './components/ViewIssueComponent';
import ListBusinessUnitComponent from './components/ListBusinessUnitComponent';
import CreateBusinessUnitComponent from './components/CreateBusinessUnitComponent';
import ViewBusinessUnitComponent from './components/ViewBusinessUnitComponent';
import ListDataProcessingActivityComponent from './components/ListDataProcessingActivityComponent';
import CreateDataProcessingActivityComponent from './components/CreateDataProcessingActivityComponent';
import ViewDataProcessingActivityComponent from './components/ViewDataProcessingActivityComponent';
import ListDataCategoryComponent from './components/ListDataCategoryComponent';
import CreateDataCategoryComponent from './components/CreateDataCategoryComponent';
import ViewDataCategoryComponent from './components/ViewDataCategoryComponent';
import ListSystem_Component from './components/ListSystem_Component';
import CreateSystem_Component from './components/CreateSystem_Component';
import ViewSystem_Component from './components/ViewSystem_Component';
import ListPrivacyNoticeComponent from './components/ListPrivacyNoticeComponent';
import CreatePrivacyNoticeComponent from './components/CreatePrivacyNoticeComponent';
import ViewPrivacyNoticeComponent from './components/ViewPrivacyNoticeComponent';
import ListDataSubjectRequestComponent from './components/ListDataSubjectRequestComponent';
import CreateDataSubjectRequestComponent from './components/CreateDataSubjectRequestComponent';
import ViewDataSubjectRequestComponent from './components/ViewDataSubjectRequestComponent';
import ListRecordsRepositoryComponent from './components/ListRecordsRepositoryComponent';
import CreateRecordsRepositoryComponent from './components/CreateRecordsRepositoryComponent';
import ViewRecordsRepositoryComponent from './components/ViewRecordsRepositoryComponent';
import ListRecord_Component from './components/ListRecord_Component';
import CreateRecord_Component from './components/CreateRecord_Component';
import ViewRecord_Component from './components/ViewRecord_Component';
import ListRetentionScheduleComponent from './components/ListRetentionScheduleComponent';
import CreateRetentionScheduleComponent from './components/CreateRetentionScheduleComponent';
import ViewRetentionScheduleComponent from './components/ViewRetentionScheduleComponent';
import ListDispositionReviewComponent from './components/ListDispositionReviewComponent';
import CreateDispositionReviewComponent from './components/CreateDispositionReviewComponent';
import ViewDispositionReviewComponent from './components/ViewDispositionReviewComponent';
import ListLegalHoldComponent from './components/ListLegalHoldComponent';
import CreateLegalHoldComponent from './components/CreateLegalHoldComponent';
import ViewLegalHoldComponent from './components/ViewLegalHoldComponent';
import ListMatterComponent from './components/ListMatterComponent';
import CreateMatterComponent from './components/CreateMatterComponent';
import ViewMatterComponent from './components/ViewMatterComponent';
import ListThirdPartyComponent from './components/ListThirdPartyComponent';
import CreateThirdPartyComponent from './components/CreateThirdPartyComponent';
import ViewThirdPartyComponent from './components/ViewThirdPartyComponent';
import ListThirdPartyAssessmentComponent from './components/ListThirdPartyAssessmentComponent';
import CreateThirdPartyAssessmentComponent from './components/CreateThirdPartyAssessmentComponent';
import ViewThirdPartyAssessmentComponent from './components/ViewThirdPartyAssessmentComponent';
import ListContractComponent from './components/ListContractComponent';
import CreateContractComponent from './components/CreateContractComponent';
import ViewContractComponent from './components/ViewContractComponent';
import ListException_Component from './components/ListException_Component';
import CreateException_Component from './components/CreateException_Component';
import ViewException_Component from './components/ViewException_Component';
import ListConsentComponent from './components/ListConsentComponent';
import CreateConsentComponent from './components/CreateConsentComponent';
import ViewConsentComponent from './components/ViewConsentComponent';
import ListDataBreachComponent from './components/ListDataBreachComponent';
import CreateDataBreachComponent from './components/CreateDataBreachComponent';
import ViewDataBreachComponent from './components/ViewDataBreachComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/organizations" component = {ListOrganizationComponent}></Route>
                            <Route path = "/add-organization/:id" component = {CreateOrganizationComponent}></Route>
                            <Route path = "/view-organization/:id" component = {ViewOrganizationComponent}></Route>
                          {/* <Route path = "/update-organization/:id" component = {UpdateOrganizationComponent}></Route> */}
                            <Route path = "/governanceBodys" component = {ListGovernanceBodyComponent}></Route>
                            <Route path = "/add-governanceBody/:id" component = {CreateGovernanceBodyComponent}></Route>
                            <Route path = "/view-governanceBody/:id" component = {ViewGovernanceBodyComponent}></Route>
                          {/* <Route path = "/update-governanceBody/:id" component = {UpdateGovernanceBodyComponent}></Route> */}
                            <Route path = "/persons" component = {ListPersonComponent}></Route>
                            <Route path = "/add-person/:id" component = {CreatePersonComponent}></Route>
                            <Route path = "/view-person/:id" component = {ViewPersonComponent}></Route>
                          {/* <Route path = "/update-person/:id" component = {UpdatePersonComponent}></Route> */}
                            <Route path = "/roles" component = {ListRoleComponent}></Route>
                            <Route path = "/add-role/:id" component = {CreateRoleComponent}></Route>
                            <Route path = "/view-role/:id" component = {ViewRoleComponent}></Route>
                          {/* <Route path = "/update-role/:id" component = {UpdateRoleComponent}></Route> */}
                            <Route path = "/roleAssignments" component = {ListRoleAssignmentComponent}></Route>
                            <Route path = "/add-roleAssignment/:id" component = {CreateRoleAssignmentComponent}></Route>
                            <Route path = "/view-roleAssignment/:id" component = {ViewRoleAssignmentComponent}></Route>
                          {/* <Route path = "/update-roleAssignment/:id" component = {UpdateRoleAssignmentComponent}></Route> */}
                            <Route path = "/policys" component = {ListPolicyComponent}></Route>
                            <Route path = "/add-policy/:id" component = {CreatePolicyComponent}></Route>
                            <Route path = "/view-policy/:id" component = {ViewPolicyComponent}></Route>
                          {/* <Route path = "/update-policy/:id" component = {UpdatePolicyComponent}></Route> */}
                            <Route path = "/procedures" component = {ListProcedureComponent}></Route>
                            <Route path = "/add-procedure/:id" component = {CreateProcedureComponent}></Route>
                            <Route path = "/view-procedure/:id" component = {ViewProcedureComponent}></Route>
                          {/* <Route path = "/update-procedure/:id" component = {UpdateProcedureComponent}></Route> */}
                            <Route path = "/regulations" component = {ListRegulationComponent}></Route>
                            <Route path = "/add-regulation/:id" component = {CreateRegulationComponent}></Route>
                            <Route path = "/view-regulation/:id" component = {ViewRegulationComponent}></Route>
                          {/* <Route path = "/update-regulation/:id" component = {UpdateRegulationComponent}></Route> */}
                            <Route path = "/obligations" component = {ListObligationComponent}></Route>
                            <Route path = "/add-obligation/:id" component = {CreateObligationComponent}></Route>
                            <Route path = "/view-obligation/:id" component = {ViewObligationComponent}></Route>
                          {/* <Route path = "/update-obligation/:id" component = {UpdateObligationComponent}></Route> */}
                            <Route path = "/controls" component = {ListControlComponent}></Route>
                            <Route path = "/add-control/:id" component = {CreateControlComponent}></Route>
                            <Route path = "/view-control/:id" component = {ViewControlComponent}></Route>
                          {/* <Route path = "/update-control/:id" component = {UpdateControlComponent}></Route> */}
                            <Route path = "/controlTest_s" component = {ListControlTest_Component}></Route>
                            <Route path = "/add-controlTest_/:id" component = {CreateControlTest_Component}></Route>
                            <Route path = "/view-controlTest_/:id" component = {ViewControlTest_Component}></Route>
                          {/* <Route path = "/update-controlTest_/:id" component = {UpdateControlTest_Component}></Route> */}
                            <Route path = "/evidences" component = {ListEvidenceComponent}></Route>
                            <Route path = "/add-evidence/:id" component = {CreateEvidenceComponent}></Route>
                            <Route path = "/view-evidence/:id" component = {ViewEvidenceComponent}></Route>
                          {/* <Route path = "/update-evidence/:id" component = {UpdateEvidenceComponent}></Route> */}
                            <Route path = "/risks" component = {ListRiskComponent}></Route>
                            <Route path = "/add-risk/:id" component = {CreateRiskComponent}></Route>
                            <Route path = "/view-risk/:id" component = {ViewRiskComponent}></Route>
                          {/* <Route path = "/update-risk/:id" component = {UpdateRiskComponent}></Route> */}
                            <Route path = "/riskAssessments" component = {ListRiskAssessmentComponent}></Route>
                            <Route path = "/add-riskAssessment/:id" component = {CreateRiskAssessmentComponent}></Route>
                            <Route path = "/view-riskAssessment/:id" component = {ViewRiskAssessmentComponent}></Route>
                          {/* <Route path = "/update-riskAssessment/:id" component = {UpdateRiskAssessmentComponent}></Route> */}
                            <Route path = "/compliancePrograms" component = {ListComplianceProgramComponent}></Route>
                            <Route path = "/add-complianceProgram/:id" component = {CreateComplianceProgramComponent}></Route>
                            <Route path = "/view-complianceProgram/:id" component = {ViewComplianceProgramComponent}></Route>
                          {/* <Route path = "/update-complianceProgram/:id" component = {UpdateComplianceProgramComponent}></Route> */}
                            <Route path = "/complianceRequirements" component = {ListComplianceRequirementComponent}></Route>
                            <Route path = "/add-complianceRequirement/:id" component = {CreateComplianceRequirementComponent}></Route>
                            <Route path = "/view-complianceRequirement/:id" component = {ViewComplianceRequirementComponent}></Route>
                          {/* <Route path = "/update-complianceRequirement/:id" component = {UpdateComplianceRequirementComponent}></Route> */}
                            <Route path = "/attestations" component = {ListAttestationComponent}></Route>
                            <Route path = "/add-attestation/:id" component = {CreateAttestationComponent}></Route>
                            <Route path = "/view-attestation/:id" component = {ViewAttestationComponent}></Route>
                          {/* <Route path = "/update-attestation/:id" component = {UpdateAttestationComponent}></Route> */}
                            <Route path = "/auditPrograms" component = {ListAuditProgramComponent}></Route>
                            <Route path = "/add-auditProgram/:id" component = {CreateAuditProgramComponent}></Route>
                            <Route path = "/view-auditProgram/:id" component = {ViewAuditProgramComponent}></Route>
                          {/* <Route path = "/update-auditProgram/:id" component = {UpdateAuditProgramComponent}></Route> */}
                            <Route path = "/auditEngagements" component = {ListAuditEngagementComponent}></Route>
                            <Route path = "/add-auditEngagement/:id" component = {CreateAuditEngagementComponent}></Route>
                            <Route path = "/view-auditEngagement/:id" component = {ViewAuditEngagementComponent}></Route>
                          {/* <Route path = "/update-auditEngagement/:id" component = {UpdateAuditEngagementComponent}></Route> */}
                            <Route path = "/auditWorkpapers" component = {ListAuditWorkpaperComponent}></Route>
                            <Route path = "/add-auditWorkpaper/:id" component = {CreateAuditWorkpaperComponent}></Route>
                            <Route path = "/view-auditWorkpaper/:id" component = {ViewAuditWorkpaperComponent}></Route>
                          {/* <Route path = "/update-auditWorkpaper/:id" component = {UpdateAuditWorkpaperComponent}></Route> */}
                            <Route path = "/auditFindings" component = {ListAuditFindingComponent}></Route>
                            <Route path = "/add-auditFinding/:id" component = {CreateAuditFindingComponent}></Route>
                            <Route path = "/view-auditFinding/:id" component = {ViewAuditFindingComponent}></Route>
                          {/* <Route path = "/update-auditFinding/:id" component = {UpdateAuditFindingComponent}></Route> */}
                            <Route path = "/correctiveActions" component = {ListCorrectiveActionComponent}></Route>
                            <Route path = "/add-correctiveAction/:id" component = {CreateCorrectiveActionComponent}></Route>
                            <Route path = "/view-correctiveAction/:id" component = {ViewCorrectiveActionComponent}></Route>
                          {/* <Route path = "/update-correctiveAction/:id" component = {UpdateCorrectiveActionComponent}></Route> */}
                            <Route path = "/issues" component = {ListIssueComponent}></Route>
                            <Route path = "/add-issue/:id" component = {CreateIssueComponent}></Route>
                            <Route path = "/view-issue/:id" component = {ViewIssueComponent}></Route>
                          {/* <Route path = "/update-issue/:id" component = {UpdateIssueComponent}></Route> */}
                            <Route path = "/businessUnits" component = {ListBusinessUnitComponent}></Route>
                            <Route path = "/add-businessUnit/:id" component = {CreateBusinessUnitComponent}></Route>
                            <Route path = "/view-businessUnit/:id" component = {ViewBusinessUnitComponent}></Route>
                          {/* <Route path = "/update-businessUnit/:id" component = {UpdateBusinessUnitComponent}></Route> */}
                            <Route path = "/dataProcessingActivitys" component = {ListDataProcessingActivityComponent}></Route>
                            <Route path = "/add-dataProcessingActivity/:id" component = {CreateDataProcessingActivityComponent}></Route>
                            <Route path = "/view-dataProcessingActivity/:id" component = {ViewDataProcessingActivityComponent}></Route>
                          {/* <Route path = "/update-dataProcessingActivity/:id" component = {UpdateDataProcessingActivityComponent}></Route> */}
                            <Route path = "/dataCategorys" component = {ListDataCategoryComponent}></Route>
                            <Route path = "/add-dataCategory/:id" component = {CreateDataCategoryComponent}></Route>
                            <Route path = "/view-dataCategory/:id" component = {ViewDataCategoryComponent}></Route>
                          {/* <Route path = "/update-dataCategory/:id" component = {UpdateDataCategoryComponent}></Route> */}
                            <Route path = "/system_s" component = {ListSystem_Component}></Route>
                            <Route path = "/add-system_/:id" component = {CreateSystem_Component}></Route>
                            <Route path = "/view-system_/:id" component = {ViewSystem_Component}></Route>
                          {/* <Route path = "/update-system_/:id" component = {UpdateSystem_Component}></Route> */}
                            <Route path = "/privacyNotices" component = {ListPrivacyNoticeComponent}></Route>
                            <Route path = "/add-privacyNotice/:id" component = {CreatePrivacyNoticeComponent}></Route>
                            <Route path = "/view-privacyNotice/:id" component = {ViewPrivacyNoticeComponent}></Route>
                          {/* <Route path = "/update-privacyNotice/:id" component = {UpdatePrivacyNoticeComponent}></Route> */}
                            <Route path = "/dataSubjectRequests" component = {ListDataSubjectRequestComponent}></Route>
                            <Route path = "/add-dataSubjectRequest/:id" component = {CreateDataSubjectRequestComponent}></Route>
                            <Route path = "/view-dataSubjectRequest/:id" component = {ViewDataSubjectRequestComponent}></Route>
                          {/* <Route path = "/update-dataSubjectRequest/:id" component = {UpdateDataSubjectRequestComponent}></Route> */}
                            <Route path = "/recordsRepositorys" component = {ListRecordsRepositoryComponent}></Route>
                            <Route path = "/add-recordsRepository/:id" component = {CreateRecordsRepositoryComponent}></Route>
                            <Route path = "/view-recordsRepository/:id" component = {ViewRecordsRepositoryComponent}></Route>
                          {/* <Route path = "/update-recordsRepository/:id" component = {UpdateRecordsRepositoryComponent}></Route> */}
                            <Route path = "/record_s" component = {ListRecord_Component}></Route>
                            <Route path = "/add-record_/:id" component = {CreateRecord_Component}></Route>
                            <Route path = "/view-record_/:id" component = {ViewRecord_Component}></Route>
                          {/* <Route path = "/update-record_/:id" component = {UpdateRecord_Component}></Route> */}
                            <Route path = "/retentionSchedules" component = {ListRetentionScheduleComponent}></Route>
                            <Route path = "/add-retentionSchedule/:id" component = {CreateRetentionScheduleComponent}></Route>
                            <Route path = "/view-retentionSchedule/:id" component = {ViewRetentionScheduleComponent}></Route>
                          {/* <Route path = "/update-retentionSchedule/:id" component = {UpdateRetentionScheduleComponent}></Route> */}
                            <Route path = "/dispositionReviews" component = {ListDispositionReviewComponent}></Route>
                            <Route path = "/add-dispositionReview/:id" component = {CreateDispositionReviewComponent}></Route>
                            <Route path = "/view-dispositionReview/:id" component = {ViewDispositionReviewComponent}></Route>
                          {/* <Route path = "/update-dispositionReview/:id" component = {UpdateDispositionReviewComponent}></Route> */}
                            <Route path = "/legalHolds" component = {ListLegalHoldComponent}></Route>
                            <Route path = "/add-legalHold/:id" component = {CreateLegalHoldComponent}></Route>
                            <Route path = "/view-legalHold/:id" component = {ViewLegalHoldComponent}></Route>
                          {/* <Route path = "/update-legalHold/:id" component = {UpdateLegalHoldComponent}></Route> */}
                            <Route path = "/matters" component = {ListMatterComponent}></Route>
                            <Route path = "/add-matter/:id" component = {CreateMatterComponent}></Route>
                            <Route path = "/view-matter/:id" component = {ViewMatterComponent}></Route>
                          {/* <Route path = "/update-matter/:id" component = {UpdateMatterComponent}></Route> */}
                            <Route path = "/thirdPartys" component = {ListThirdPartyComponent}></Route>
                            <Route path = "/add-thirdParty/:id" component = {CreateThirdPartyComponent}></Route>
                            <Route path = "/view-thirdParty/:id" component = {ViewThirdPartyComponent}></Route>
                          {/* <Route path = "/update-thirdParty/:id" component = {UpdateThirdPartyComponent}></Route> */}
                            <Route path = "/thirdPartyAssessments" component = {ListThirdPartyAssessmentComponent}></Route>
                            <Route path = "/add-thirdPartyAssessment/:id" component = {CreateThirdPartyAssessmentComponent}></Route>
                            <Route path = "/view-thirdPartyAssessment/:id" component = {ViewThirdPartyAssessmentComponent}></Route>
                          {/* <Route path = "/update-thirdPartyAssessment/:id" component = {UpdateThirdPartyAssessmentComponent}></Route> */}
                            <Route path = "/contracts" component = {ListContractComponent}></Route>
                            <Route path = "/add-contract/:id" component = {CreateContractComponent}></Route>
                            <Route path = "/view-contract/:id" component = {ViewContractComponent}></Route>
                          {/* <Route path = "/update-contract/:id" component = {UpdateContractComponent}></Route> */}
                            <Route path = "/exception_s" component = {ListException_Component}></Route>
                            <Route path = "/add-exception_/:id" component = {CreateException_Component}></Route>
                            <Route path = "/view-exception_/:id" component = {ViewException_Component}></Route>
                          {/* <Route path = "/update-exception_/:id" component = {UpdateException_Component}></Route> */}
                            <Route path = "/consents" component = {ListConsentComponent}></Route>
                            <Route path = "/add-consent/:id" component = {CreateConsentComponent}></Route>
                            <Route path = "/view-consent/:id" component = {ViewConsentComponent}></Route>
                          {/* <Route path = "/update-consent/:id" component = {UpdateConsentComponent}></Route> */}
                            <Route path = "/dataBreachs" component = {ListDataBreachComponent}></Route>
                            <Route path = "/add-dataBreach/:id" component = {CreateDataBreachComponent}></Route>
                            <Route path = "/view-dataBreach/:id" component = {ViewDataBreachComponent}></Route>
                          {/* <Route path = "/update-dataBreach/:id" component = {UpdateDataBreachComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
