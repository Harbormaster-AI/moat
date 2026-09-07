// routerConfig.ts

import { Routes } from '@angular/router';
import { CreateOrganizationComponent } from './components/Organization/create/create.component';
import { EditOrganizationComponent } from './components/Organization/edit/edit.component';
import { IndexOrganizationComponent } from './components/Organization/index/index.component';
import { CreateGovernanceBodyComponent } from './components/GovernanceBody/create/create.component';
import { EditGovernanceBodyComponent } from './components/GovernanceBody/edit/edit.component';
import { IndexGovernanceBodyComponent } from './components/GovernanceBody/index/index.component';
import { CreatePersonComponent } from './components/Person/create/create.component';
import { EditPersonComponent } from './components/Person/edit/edit.component';
import { IndexPersonComponent } from './components/Person/index/index.component';
import { CreateRoleComponent } from './components/Role/create/create.component';
import { EditRoleComponent } from './components/Role/edit/edit.component';
import { IndexRoleComponent } from './components/Role/index/index.component';
import { CreateRoleAssignmentComponent } from './components/RoleAssignment/create/create.component';
import { EditRoleAssignmentComponent } from './components/RoleAssignment/edit/edit.component';
import { IndexRoleAssignmentComponent } from './components/RoleAssignment/index/index.component';
import { CreatePolicyComponent } from './components/Policy/create/create.component';
import { EditPolicyComponent } from './components/Policy/edit/edit.component';
import { IndexPolicyComponent } from './components/Policy/index/index.component';
import { CreateProcedureComponent } from './components/Procedure/create/create.component';
import { EditProcedureComponent } from './components/Procedure/edit/edit.component';
import { IndexProcedureComponent } from './components/Procedure/index/index.component';
import { CreateRegulationComponent } from './components/Regulation/create/create.component';
import { EditRegulationComponent } from './components/Regulation/edit/edit.component';
import { IndexRegulationComponent } from './components/Regulation/index/index.component';
import { CreateObligationComponent } from './components/Obligation/create/create.component';
import { EditObligationComponent } from './components/Obligation/edit/edit.component';
import { IndexObligationComponent } from './components/Obligation/index/index.component';
import { CreateControlComponent } from './components/Control/create/create.component';
import { EditControlComponent } from './components/Control/edit/edit.component';
import { IndexControlComponent } from './components/Control/index/index.component';
import { CreateControlTest_Component } from './components/ControlTest_/create/create.component';
import { EditControlTest_Component } from './components/ControlTest_/edit/edit.component';
import { IndexControlTest_Component } from './components/ControlTest_/index/index.component';
import { CreateEvidenceComponent } from './components/Evidence/create/create.component';
import { EditEvidenceComponent } from './components/Evidence/edit/edit.component';
import { IndexEvidenceComponent } from './components/Evidence/index/index.component';
import { CreateRiskComponent } from './components/Risk/create/create.component';
import { EditRiskComponent } from './components/Risk/edit/edit.component';
import { IndexRiskComponent } from './components/Risk/index/index.component';
import { CreateRiskAssessmentComponent } from './components/RiskAssessment/create/create.component';
import { EditRiskAssessmentComponent } from './components/RiskAssessment/edit/edit.component';
import { IndexRiskAssessmentComponent } from './components/RiskAssessment/index/index.component';
import { CreateComplianceProgramComponent } from './components/ComplianceProgram/create/create.component';
import { EditComplianceProgramComponent } from './components/ComplianceProgram/edit/edit.component';
import { IndexComplianceProgramComponent } from './components/ComplianceProgram/index/index.component';
import { CreateComplianceRequirementComponent } from './components/ComplianceRequirement/create/create.component';
import { EditComplianceRequirementComponent } from './components/ComplianceRequirement/edit/edit.component';
import { IndexComplianceRequirementComponent } from './components/ComplianceRequirement/index/index.component';
import { CreateAttestationComponent } from './components/Attestation/create/create.component';
import { EditAttestationComponent } from './components/Attestation/edit/edit.component';
import { IndexAttestationComponent } from './components/Attestation/index/index.component';
import { CreateAuditProgramComponent } from './components/AuditProgram/create/create.component';
import { EditAuditProgramComponent } from './components/AuditProgram/edit/edit.component';
import { IndexAuditProgramComponent } from './components/AuditProgram/index/index.component';
import { CreateAuditEngagementComponent } from './components/AuditEngagement/create/create.component';
import { EditAuditEngagementComponent } from './components/AuditEngagement/edit/edit.component';
import { IndexAuditEngagementComponent } from './components/AuditEngagement/index/index.component';
import { CreateAuditWorkpaperComponent } from './components/AuditWorkpaper/create/create.component';
import { EditAuditWorkpaperComponent } from './components/AuditWorkpaper/edit/edit.component';
import { IndexAuditWorkpaperComponent } from './components/AuditWorkpaper/index/index.component';
import { CreateAuditFindingComponent } from './components/AuditFinding/create/create.component';
import { EditAuditFindingComponent } from './components/AuditFinding/edit/edit.component';
import { IndexAuditFindingComponent } from './components/AuditFinding/index/index.component';
import { CreateCorrectiveActionComponent } from './components/CorrectiveAction/create/create.component';
import { EditCorrectiveActionComponent } from './components/CorrectiveAction/edit/edit.component';
import { IndexCorrectiveActionComponent } from './components/CorrectiveAction/index/index.component';
import { CreateIssueComponent } from './components/Issue/create/create.component';
import { EditIssueComponent } from './components/Issue/edit/edit.component';
import { IndexIssueComponent } from './components/Issue/index/index.component';
import { CreateBusinessUnitComponent } from './components/BusinessUnit/create/create.component';
import { EditBusinessUnitComponent } from './components/BusinessUnit/edit/edit.component';
import { IndexBusinessUnitComponent } from './components/BusinessUnit/index/index.component';
import { CreateDataProcessingActivityComponent } from './components/DataProcessingActivity/create/create.component';
import { EditDataProcessingActivityComponent } from './components/DataProcessingActivity/edit/edit.component';
import { IndexDataProcessingActivityComponent } from './components/DataProcessingActivity/index/index.component';
import { CreateDataCategoryComponent } from './components/DataCategory/create/create.component';
import { EditDataCategoryComponent } from './components/DataCategory/edit/edit.component';
import { IndexDataCategoryComponent } from './components/DataCategory/index/index.component';
import { CreateSystem_Component } from './components/System_/create/create.component';
import { EditSystem_Component } from './components/System_/edit/edit.component';
import { IndexSystem_Component } from './components/System_/index/index.component';
import { CreatePrivacyNoticeComponent } from './components/PrivacyNotice/create/create.component';
import { EditPrivacyNoticeComponent } from './components/PrivacyNotice/edit/edit.component';
import { IndexPrivacyNoticeComponent } from './components/PrivacyNotice/index/index.component';
import { CreateDataSubjectRequestComponent } from './components/DataSubjectRequest/create/create.component';
import { EditDataSubjectRequestComponent } from './components/DataSubjectRequest/edit/edit.component';
import { IndexDataSubjectRequestComponent } from './components/DataSubjectRequest/index/index.component';
import { CreateRecordsRepositoryComponent } from './components/RecordsRepository/create/create.component';
import { EditRecordsRepositoryComponent } from './components/RecordsRepository/edit/edit.component';
import { IndexRecordsRepositoryComponent } from './components/RecordsRepository/index/index.component';
import { CreateRecord_Component } from './components/Record_/create/create.component';
import { EditRecord_Component } from './components/Record_/edit/edit.component';
import { IndexRecord_Component } from './components/Record_/index/index.component';
import { CreateRetentionScheduleComponent } from './components/RetentionSchedule/create/create.component';
import { EditRetentionScheduleComponent } from './components/RetentionSchedule/edit/edit.component';
import { IndexRetentionScheduleComponent } from './components/RetentionSchedule/index/index.component';
import { CreateDispositionReviewComponent } from './components/DispositionReview/create/create.component';
import { EditDispositionReviewComponent } from './components/DispositionReview/edit/edit.component';
import { IndexDispositionReviewComponent } from './components/DispositionReview/index/index.component';
import { CreateLegalHoldComponent } from './components/LegalHold/create/create.component';
import { EditLegalHoldComponent } from './components/LegalHold/edit/edit.component';
import { IndexLegalHoldComponent } from './components/LegalHold/index/index.component';
import { CreateMatterComponent } from './components/Matter/create/create.component';
import { EditMatterComponent } from './components/Matter/edit/edit.component';
import { IndexMatterComponent } from './components/Matter/index/index.component';
import { CreateThirdPartyComponent } from './components/ThirdParty/create/create.component';
import { EditThirdPartyComponent } from './components/ThirdParty/edit/edit.component';
import { IndexThirdPartyComponent } from './components/ThirdParty/index/index.component';
import { CreateThirdPartyAssessmentComponent } from './components/ThirdPartyAssessment/create/create.component';
import { EditThirdPartyAssessmentComponent } from './components/ThirdPartyAssessment/edit/edit.component';
import { IndexThirdPartyAssessmentComponent } from './components/ThirdPartyAssessment/index/index.component';
import { CreateContractComponent } from './components/Contract/create/create.component';
import { EditContractComponent } from './components/Contract/edit/edit.component';
import { IndexContractComponent } from './components/Contract/index/index.component';
import { CreateException_Component } from './components/Exception_/create/create.component';
import { EditException_Component } from './components/Exception_/edit/edit.component';
import { IndexException_Component } from './components/Exception_/index/index.component';
import { CreateConsentComponent } from './components/Consent/create/create.component';
import { EditConsentComponent } from './components/Consent/edit/edit.component';
import { IndexConsentComponent } from './components/Consent/index/index.component';
import { CreateDataBreachComponent } from './components/DataBreach/create/create.component';
import { EditDataBreachComponent } from './components/DataBreach/edit/edit.component';
import { IndexDataBreachComponent } from './components/DataBreach/index/index.component';

export const OrganizationRoutes: Routes = [
  { path: 'createOrganization',
    component: CreateOrganizationComponent
  },
  {
    path: 'editOrganization/:id',
    component: EditOrganizationComponent
  },
  { path: 'indexOrganization',
    component: IndexOrganizationComponent
  }
];
export const GovernanceBodyRoutes: Routes = [
  { path: 'createGovernanceBody',
    component: CreateGovernanceBodyComponent
  },
  {
    path: 'editGovernanceBody/:id',
    component: EditGovernanceBodyComponent
  },
  { path: 'indexGovernanceBody',
    component: IndexGovernanceBodyComponent
  }
];
export const PersonRoutes: Routes = [
  { path: 'createPerson',
    component: CreatePersonComponent
  },
  {
    path: 'editPerson/:id',
    component: EditPersonComponent
  },
  { path: 'indexPerson',
    component: IndexPersonComponent
  }
];
export const RoleRoutes: Routes = [
  { path: 'createRole',
    component: CreateRoleComponent
  },
  {
    path: 'editRole/:id',
    component: EditRoleComponent
  },
  { path: 'indexRole',
    component: IndexRoleComponent
  }
];
export const RoleAssignmentRoutes: Routes = [
  { path: 'createRoleAssignment',
    component: CreateRoleAssignmentComponent
  },
  {
    path: 'editRoleAssignment/:id',
    component: EditRoleAssignmentComponent
  },
  { path: 'indexRoleAssignment',
    component: IndexRoleAssignmentComponent
  }
];
export const PolicyRoutes: Routes = [
  { path: 'createPolicy',
    component: CreatePolicyComponent
  },
  {
    path: 'editPolicy/:id',
    component: EditPolicyComponent
  },
  { path: 'indexPolicy',
    component: IndexPolicyComponent
  }
];
export const ProcedureRoutes: Routes = [
  { path: 'createProcedure',
    component: CreateProcedureComponent
  },
  {
    path: 'editProcedure/:id',
    component: EditProcedureComponent
  },
  { path: 'indexProcedure',
    component: IndexProcedureComponent
  }
];
export const RegulationRoutes: Routes = [
  { path: 'createRegulation',
    component: CreateRegulationComponent
  },
  {
    path: 'editRegulation/:id',
    component: EditRegulationComponent
  },
  { path: 'indexRegulation',
    component: IndexRegulationComponent
  }
];
export const ObligationRoutes: Routes = [
  { path: 'createObligation',
    component: CreateObligationComponent
  },
  {
    path: 'editObligation/:id',
    component: EditObligationComponent
  },
  { path: 'indexObligation',
    component: IndexObligationComponent
  }
];
export const ControlRoutes: Routes = [
  { path: 'createControl',
    component: CreateControlComponent
  },
  {
    path: 'editControl/:id',
    component: EditControlComponent
  },
  { path: 'indexControl',
    component: IndexControlComponent
  }
];
export const ControlTest_Routes: Routes = [
  { path: 'createControlTest_',
    component: CreateControlTest_Component
  },
  {
    path: 'editControlTest_/:id',
    component: EditControlTest_Component
  },
  { path: 'indexControlTest_',
    component: IndexControlTest_Component
  }
];
export const EvidenceRoutes: Routes = [
  { path: 'createEvidence',
    component: CreateEvidenceComponent
  },
  {
    path: 'editEvidence/:id',
    component: EditEvidenceComponent
  },
  { path: 'indexEvidence',
    component: IndexEvidenceComponent
  }
];
export const RiskRoutes: Routes = [
  { path: 'createRisk',
    component: CreateRiskComponent
  },
  {
    path: 'editRisk/:id',
    component: EditRiskComponent
  },
  { path: 'indexRisk',
    component: IndexRiskComponent
  }
];
export const RiskAssessmentRoutes: Routes = [
  { path: 'createRiskAssessment',
    component: CreateRiskAssessmentComponent
  },
  {
    path: 'editRiskAssessment/:id',
    component: EditRiskAssessmentComponent
  },
  { path: 'indexRiskAssessment',
    component: IndexRiskAssessmentComponent
  }
];
export const ComplianceProgramRoutes: Routes = [
  { path: 'createComplianceProgram',
    component: CreateComplianceProgramComponent
  },
  {
    path: 'editComplianceProgram/:id',
    component: EditComplianceProgramComponent
  },
  { path: 'indexComplianceProgram',
    component: IndexComplianceProgramComponent
  }
];
export const ComplianceRequirementRoutes: Routes = [
  { path: 'createComplianceRequirement',
    component: CreateComplianceRequirementComponent
  },
  {
    path: 'editComplianceRequirement/:id',
    component: EditComplianceRequirementComponent
  },
  { path: 'indexComplianceRequirement',
    component: IndexComplianceRequirementComponent
  }
];
export const AttestationRoutes: Routes = [
  { path: 'createAttestation',
    component: CreateAttestationComponent
  },
  {
    path: 'editAttestation/:id',
    component: EditAttestationComponent
  },
  { path: 'indexAttestation',
    component: IndexAttestationComponent
  }
];
export const AuditProgramRoutes: Routes = [
  { path: 'createAuditProgram',
    component: CreateAuditProgramComponent
  },
  {
    path: 'editAuditProgram/:id',
    component: EditAuditProgramComponent
  },
  { path: 'indexAuditProgram',
    component: IndexAuditProgramComponent
  }
];
export const AuditEngagementRoutes: Routes = [
  { path: 'createAuditEngagement',
    component: CreateAuditEngagementComponent
  },
  {
    path: 'editAuditEngagement/:id',
    component: EditAuditEngagementComponent
  },
  { path: 'indexAuditEngagement',
    component: IndexAuditEngagementComponent
  }
];
export const AuditWorkpaperRoutes: Routes = [
  { path: 'createAuditWorkpaper',
    component: CreateAuditWorkpaperComponent
  },
  {
    path: 'editAuditWorkpaper/:id',
    component: EditAuditWorkpaperComponent
  },
  { path: 'indexAuditWorkpaper',
    component: IndexAuditWorkpaperComponent
  }
];
export const AuditFindingRoutes: Routes = [
  { path: 'createAuditFinding',
    component: CreateAuditFindingComponent
  },
  {
    path: 'editAuditFinding/:id',
    component: EditAuditFindingComponent
  },
  { path: 'indexAuditFinding',
    component: IndexAuditFindingComponent
  }
];
export const CorrectiveActionRoutes: Routes = [
  { path: 'createCorrectiveAction',
    component: CreateCorrectiveActionComponent
  },
  {
    path: 'editCorrectiveAction/:id',
    component: EditCorrectiveActionComponent
  },
  { path: 'indexCorrectiveAction',
    component: IndexCorrectiveActionComponent
  }
];
export const IssueRoutes: Routes = [
  { path: 'createIssue',
    component: CreateIssueComponent
  },
  {
    path: 'editIssue/:id',
    component: EditIssueComponent
  },
  { path: 'indexIssue',
    component: IndexIssueComponent
  }
];
export const BusinessUnitRoutes: Routes = [
  { path: 'createBusinessUnit',
    component: CreateBusinessUnitComponent
  },
  {
    path: 'editBusinessUnit/:id',
    component: EditBusinessUnitComponent
  },
  { path: 'indexBusinessUnit',
    component: IndexBusinessUnitComponent
  }
];
export const DataProcessingActivityRoutes: Routes = [
  { path: 'createDataProcessingActivity',
    component: CreateDataProcessingActivityComponent
  },
  {
    path: 'editDataProcessingActivity/:id',
    component: EditDataProcessingActivityComponent
  },
  { path: 'indexDataProcessingActivity',
    component: IndexDataProcessingActivityComponent
  }
];
export const DataCategoryRoutes: Routes = [
  { path: 'createDataCategory',
    component: CreateDataCategoryComponent
  },
  {
    path: 'editDataCategory/:id',
    component: EditDataCategoryComponent
  },
  { path: 'indexDataCategory',
    component: IndexDataCategoryComponent
  }
];
export const System_Routes: Routes = [
  { path: 'createSystem_',
    component: CreateSystem_Component
  },
  {
    path: 'editSystem_/:id',
    component: EditSystem_Component
  },
  { path: 'indexSystem_',
    component: IndexSystem_Component
  }
];
export const PrivacyNoticeRoutes: Routes = [
  { path: 'createPrivacyNotice',
    component: CreatePrivacyNoticeComponent
  },
  {
    path: 'editPrivacyNotice/:id',
    component: EditPrivacyNoticeComponent
  },
  { path: 'indexPrivacyNotice',
    component: IndexPrivacyNoticeComponent
  }
];
export const DataSubjectRequestRoutes: Routes = [
  { path: 'createDataSubjectRequest',
    component: CreateDataSubjectRequestComponent
  },
  {
    path: 'editDataSubjectRequest/:id',
    component: EditDataSubjectRequestComponent
  },
  { path: 'indexDataSubjectRequest',
    component: IndexDataSubjectRequestComponent
  }
];
export const RecordsRepositoryRoutes: Routes = [
  { path: 'createRecordsRepository',
    component: CreateRecordsRepositoryComponent
  },
  {
    path: 'editRecordsRepository/:id',
    component: EditRecordsRepositoryComponent
  },
  { path: 'indexRecordsRepository',
    component: IndexRecordsRepositoryComponent
  }
];
export const Record_Routes: Routes = [
  { path: 'createRecord_',
    component: CreateRecord_Component
  },
  {
    path: 'editRecord_/:id',
    component: EditRecord_Component
  },
  { path: 'indexRecord_',
    component: IndexRecord_Component
  }
];
export const RetentionScheduleRoutes: Routes = [
  { path: 'createRetentionSchedule',
    component: CreateRetentionScheduleComponent
  },
  {
    path: 'editRetentionSchedule/:id',
    component: EditRetentionScheduleComponent
  },
  { path: 'indexRetentionSchedule',
    component: IndexRetentionScheduleComponent
  }
];
export const DispositionReviewRoutes: Routes = [
  { path: 'createDispositionReview',
    component: CreateDispositionReviewComponent
  },
  {
    path: 'editDispositionReview/:id',
    component: EditDispositionReviewComponent
  },
  { path: 'indexDispositionReview',
    component: IndexDispositionReviewComponent
  }
];
export const LegalHoldRoutes: Routes = [
  { path: 'createLegalHold',
    component: CreateLegalHoldComponent
  },
  {
    path: 'editLegalHold/:id',
    component: EditLegalHoldComponent
  },
  { path: 'indexLegalHold',
    component: IndexLegalHoldComponent
  }
];
export const MatterRoutes: Routes = [
  { path: 'createMatter',
    component: CreateMatterComponent
  },
  {
    path: 'editMatter/:id',
    component: EditMatterComponent
  },
  { path: 'indexMatter',
    component: IndexMatterComponent
  }
];
export const ThirdPartyRoutes: Routes = [
  { path: 'createThirdParty',
    component: CreateThirdPartyComponent
  },
  {
    path: 'editThirdParty/:id',
    component: EditThirdPartyComponent
  },
  { path: 'indexThirdParty',
    component: IndexThirdPartyComponent
  }
];
export const ThirdPartyAssessmentRoutes: Routes = [
  { path: 'createThirdPartyAssessment',
    component: CreateThirdPartyAssessmentComponent
  },
  {
    path: 'editThirdPartyAssessment/:id',
    component: EditThirdPartyAssessmentComponent
  },
  { path: 'indexThirdPartyAssessment',
    component: IndexThirdPartyAssessmentComponent
  }
];
export const ContractRoutes: Routes = [
  { path: 'createContract',
    component: CreateContractComponent
  },
  {
    path: 'editContract/:id',
    component: EditContractComponent
  },
  { path: 'indexContract',
    component: IndexContractComponent
  }
];
export const Exception_Routes: Routes = [
  { path: 'createException_',
    component: CreateException_Component
  },
  {
    path: 'editException_/:id',
    component: EditException_Component
  },
  { path: 'indexException_',
    component: IndexException_Component
  }
];
export const ConsentRoutes: Routes = [
  { path: 'createConsent',
    component: CreateConsentComponent
  },
  {
    path: 'editConsent/:id',
    component: EditConsentComponent
  },
  { path: 'indexConsent',
    component: IndexConsentComponent
  }
];
export const DataBreachRoutes: Routes = [
  { path: 'createDataBreach',
    component: CreateDataBreachComponent
  },
  {
    path: 'editDataBreach/:id',
    component: EditDataBreachComponent
  },
  { path: 'indexDataBreach',
    component: IndexDataBreachComponent
  }
];
