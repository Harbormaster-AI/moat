import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatInputModule} from '@angular/material/input';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatMomentDateModule} from "@angular/material-moment-adapter";
import {NgModule} from '@angular/core';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {RouterModule} from '@angular/router';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule} from '@angular/forms';
import {ReactiveFormsModule} from '@angular/forms';
import {AppComponent} from './app.component';
import {MatMenuModule} from '@angular/material/menu';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatSidenavModule} from '@angular/material/sidenav'

import {IndexOrganizationComponent} from './components/Organization/index/index.component';
import {CreateOrganizationComponent} from './components/Organization/create/create.component';
import {EditOrganizationComponent} from './components/Organization/edit/edit.component';
import {IndexGovernanceBodyComponent} from './components/GovernanceBody/index/index.component';
import {CreateGovernanceBodyComponent} from './components/GovernanceBody/create/create.component';
import {EditGovernanceBodyComponent} from './components/GovernanceBody/edit/edit.component';
import {IndexPersonComponent} from './components/Person/index/index.component';
import {CreatePersonComponent} from './components/Person/create/create.component';
import {EditPersonComponent} from './components/Person/edit/edit.component';
import {IndexRoleComponent} from './components/Role/index/index.component';
import {CreateRoleComponent} from './components/Role/create/create.component';
import {EditRoleComponent} from './components/Role/edit/edit.component';
import {IndexRoleAssignmentComponent} from './components/RoleAssignment/index/index.component';
import {CreateRoleAssignmentComponent} from './components/RoleAssignment/create/create.component';
import {EditRoleAssignmentComponent} from './components/RoleAssignment/edit/edit.component';
import {IndexPolicyComponent} from './components/Policy/index/index.component';
import {CreatePolicyComponent} from './components/Policy/create/create.component';
import {EditPolicyComponent} from './components/Policy/edit/edit.component';
import {IndexProcedureComponent} from './components/Procedure/index/index.component';
import {CreateProcedureComponent} from './components/Procedure/create/create.component';
import {EditProcedureComponent} from './components/Procedure/edit/edit.component';
import {IndexRegulationComponent} from './components/Regulation/index/index.component';
import {CreateRegulationComponent} from './components/Regulation/create/create.component';
import {EditRegulationComponent} from './components/Regulation/edit/edit.component';
import {IndexObligationComponent} from './components/Obligation/index/index.component';
import {CreateObligationComponent} from './components/Obligation/create/create.component';
import {EditObligationComponent} from './components/Obligation/edit/edit.component';
import {IndexControlComponent} from './components/Control/index/index.component';
import {CreateControlComponent} from './components/Control/create/create.component';
import {EditControlComponent} from './components/Control/edit/edit.component';
import {IndexControlTest_Component} from './components/ControlTest_/index/index.component';
import {CreateControlTest_Component} from './components/ControlTest_/create/create.component';
import {EditControlTest_Component} from './components/ControlTest_/edit/edit.component';
import {IndexEvidenceComponent} from './components/Evidence/index/index.component';
import {CreateEvidenceComponent} from './components/Evidence/create/create.component';
import {EditEvidenceComponent} from './components/Evidence/edit/edit.component';
import {IndexRiskComponent} from './components/Risk/index/index.component';
import {CreateRiskComponent} from './components/Risk/create/create.component';
import {EditRiskComponent} from './components/Risk/edit/edit.component';
import {IndexRiskAssessmentComponent} from './components/RiskAssessment/index/index.component';
import {CreateRiskAssessmentComponent} from './components/RiskAssessment/create/create.component';
import {EditRiskAssessmentComponent} from './components/RiskAssessment/edit/edit.component';
import {IndexComplianceProgramComponent} from './components/ComplianceProgram/index/index.component';
import {CreateComplianceProgramComponent} from './components/ComplianceProgram/create/create.component';
import {EditComplianceProgramComponent} from './components/ComplianceProgram/edit/edit.component';
import {IndexComplianceRequirementComponent} from './components/ComplianceRequirement/index/index.component';
import {CreateComplianceRequirementComponent} from './components/ComplianceRequirement/create/create.component';
import {EditComplianceRequirementComponent} from './components/ComplianceRequirement/edit/edit.component';
import {IndexAttestationComponent} from './components/Attestation/index/index.component';
import {CreateAttestationComponent} from './components/Attestation/create/create.component';
import {EditAttestationComponent} from './components/Attestation/edit/edit.component';
import {IndexAuditProgramComponent} from './components/AuditProgram/index/index.component';
import {CreateAuditProgramComponent} from './components/AuditProgram/create/create.component';
import {EditAuditProgramComponent} from './components/AuditProgram/edit/edit.component';
import {IndexAuditEngagementComponent} from './components/AuditEngagement/index/index.component';
import {CreateAuditEngagementComponent} from './components/AuditEngagement/create/create.component';
import {EditAuditEngagementComponent} from './components/AuditEngagement/edit/edit.component';
import {IndexAuditWorkpaperComponent} from './components/AuditWorkpaper/index/index.component';
import {CreateAuditWorkpaperComponent} from './components/AuditWorkpaper/create/create.component';
import {EditAuditWorkpaperComponent} from './components/AuditWorkpaper/edit/edit.component';
import {IndexAuditFindingComponent} from './components/AuditFinding/index/index.component';
import {CreateAuditFindingComponent} from './components/AuditFinding/create/create.component';
import {EditAuditFindingComponent} from './components/AuditFinding/edit/edit.component';
import {IndexCorrectiveActionComponent} from './components/CorrectiveAction/index/index.component';
import {CreateCorrectiveActionComponent} from './components/CorrectiveAction/create/create.component';
import {EditCorrectiveActionComponent} from './components/CorrectiveAction/edit/edit.component';
import {IndexIssueComponent} from './components/Issue/index/index.component';
import {CreateIssueComponent} from './components/Issue/create/create.component';
import {EditIssueComponent} from './components/Issue/edit/edit.component';
import {IndexBusinessUnitComponent} from './components/BusinessUnit/index/index.component';
import {CreateBusinessUnitComponent} from './components/BusinessUnit/create/create.component';
import {EditBusinessUnitComponent} from './components/BusinessUnit/edit/edit.component';
import {IndexDataProcessingActivityComponent} from './components/DataProcessingActivity/index/index.component';
import {CreateDataProcessingActivityComponent} from './components/DataProcessingActivity/create/create.component';
import {EditDataProcessingActivityComponent} from './components/DataProcessingActivity/edit/edit.component';
import {IndexDataCategoryComponent} from './components/DataCategory/index/index.component';
import {CreateDataCategoryComponent} from './components/DataCategory/create/create.component';
import {EditDataCategoryComponent} from './components/DataCategory/edit/edit.component';
import {IndexSystem_Component} from './components/System_/index/index.component';
import {CreateSystem_Component} from './components/System_/create/create.component';
import {EditSystem_Component} from './components/System_/edit/edit.component';
import {IndexPrivacyNoticeComponent} from './components/PrivacyNotice/index/index.component';
import {CreatePrivacyNoticeComponent} from './components/PrivacyNotice/create/create.component';
import {EditPrivacyNoticeComponent} from './components/PrivacyNotice/edit/edit.component';
import {IndexDataSubjectRequestComponent} from './components/DataSubjectRequest/index/index.component';
import {CreateDataSubjectRequestComponent} from './components/DataSubjectRequest/create/create.component';
import {EditDataSubjectRequestComponent} from './components/DataSubjectRequest/edit/edit.component';
import {IndexRecordsRepositoryComponent} from './components/RecordsRepository/index/index.component';
import {CreateRecordsRepositoryComponent} from './components/RecordsRepository/create/create.component';
import {EditRecordsRepositoryComponent} from './components/RecordsRepository/edit/edit.component';
import {IndexRecord_Component} from './components/Record_/index/index.component';
import {CreateRecord_Component} from './components/Record_/create/create.component';
import {EditRecord_Component} from './components/Record_/edit/edit.component';
import {IndexRetentionScheduleComponent} from './components/RetentionSchedule/index/index.component';
import {CreateRetentionScheduleComponent} from './components/RetentionSchedule/create/create.component';
import {EditRetentionScheduleComponent} from './components/RetentionSchedule/edit/edit.component';
import {IndexDispositionReviewComponent} from './components/DispositionReview/index/index.component';
import {CreateDispositionReviewComponent} from './components/DispositionReview/create/create.component';
import {EditDispositionReviewComponent} from './components/DispositionReview/edit/edit.component';
import {IndexLegalHoldComponent} from './components/LegalHold/index/index.component';
import {CreateLegalHoldComponent} from './components/LegalHold/create/create.component';
import {EditLegalHoldComponent} from './components/LegalHold/edit/edit.component';
import {IndexMatterComponent} from './components/Matter/index/index.component';
import {CreateMatterComponent} from './components/Matter/create/create.component';
import {EditMatterComponent} from './components/Matter/edit/edit.component';
import {IndexThirdPartyComponent} from './components/ThirdParty/index/index.component';
import {CreateThirdPartyComponent} from './components/ThirdParty/create/create.component';
import {EditThirdPartyComponent} from './components/ThirdParty/edit/edit.component';
import {IndexThirdPartyAssessmentComponent} from './components/ThirdPartyAssessment/index/index.component';
import {CreateThirdPartyAssessmentComponent} from './components/ThirdPartyAssessment/create/create.component';
import {EditThirdPartyAssessmentComponent} from './components/ThirdPartyAssessment/edit/edit.component';
import {IndexContractComponent} from './components/Contract/index/index.component';
import {CreateContractComponent} from './components/Contract/create/create.component';
import {EditContractComponent} from './components/Contract/edit/edit.component';
import {IndexException_Component} from './components/Exception_/index/index.component';
import {CreateException_Component} from './components/Exception_/create/create.component';
import {EditException_Component} from './components/Exception_/edit/edit.component';
import {IndexConsentComponent} from './components/Consent/index/index.component';
import {CreateConsentComponent} from './components/Consent/create/create.component';
import {EditConsentComponent} from './components/Consent/edit/edit.component';
import {IndexDataBreachComponent} from './components/DataBreach/index/index.component';
import {CreateDataBreachComponent} from './components/DataBreach/create/create.component';
import {EditDataBreachComponent} from './components/DataBreach/edit/edit.component';

import * as appRoutes from './routerConfig';

import {OrganizationService} from './services/Organization.service';
import {GovernanceBodyService} from './services/GovernanceBody.service';
import {PersonService} from './services/Person.service';
import {RoleService} from './services/Role.service';
import {RoleAssignmentService} from './services/RoleAssignment.service';
import {PolicyService} from './services/Policy.service';
import {ProcedureService} from './services/Procedure.service';
import {RegulationService} from './services/Regulation.service';
import {ObligationService} from './services/Obligation.service';
import {ControlService} from './services/Control.service';
import {ControlTest_Service} from './services/ControlTest_.service';
import {EvidenceService} from './services/Evidence.service';
import {RiskService} from './services/Risk.service';
import {RiskAssessmentService} from './services/RiskAssessment.service';
import {ComplianceProgramService} from './services/ComplianceProgram.service';
import {ComplianceRequirementService} from './services/ComplianceRequirement.service';
import {AttestationService} from './services/Attestation.service';
import {AuditProgramService} from './services/AuditProgram.service';
import {AuditEngagementService} from './services/AuditEngagement.service';
import {AuditWorkpaperService} from './services/AuditWorkpaper.service';
import {AuditFindingService} from './services/AuditFinding.service';
import {CorrectiveActionService} from './services/CorrectiveAction.service';
import {IssueService} from './services/Issue.service';
import {BusinessUnitService} from './services/BusinessUnit.service';
import {DataProcessingActivityService} from './services/DataProcessingActivity.service';
import {DataCategoryService} from './services/DataCategory.service';
import {System_Service} from './services/System_.service';
import {PrivacyNoticeService} from './services/PrivacyNotice.service';
import {DataSubjectRequestService} from './services/DataSubjectRequest.service';
import {RecordsRepositoryService} from './services/RecordsRepository.service';
import {Record_Service} from './services/Record_.service';
import {RetentionScheduleService} from './services/RetentionSchedule.service';
import {DispositionReviewService} from './services/DispositionReview.service';
import {LegalHoldService} from './services/LegalHold.service';
import {MatterService} from './services/Matter.service';
import {ThirdPartyService} from './services/ThirdParty.service';
import {ThirdPartyAssessmentService} from './services/ThirdPartyAssessment.service';
import {ContractService} from './services/Contract.service';
import {Exception_Service} from './services/Exception_.service';
import {ConsentService} from './services/Consent.service';
import {DataBreachService} from './services/DataBreach.service';

@NgModule({
  declarations: [
    IndexOrganizationComponent,
    CreateOrganizationComponent,
    EditOrganizationComponent,
    IndexGovernanceBodyComponent,
    CreateGovernanceBodyComponent,
    EditGovernanceBodyComponent,
    IndexPersonComponent,
    CreatePersonComponent,
    EditPersonComponent,
    IndexRoleComponent,
    CreateRoleComponent,
    EditRoleComponent,
    IndexRoleAssignmentComponent,
    CreateRoleAssignmentComponent,
    EditRoleAssignmentComponent,
    IndexPolicyComponent,
    CreatePolicyComponent,
    EditPolicyComponent,
    IndexProcedureComponent,
    CreateProcedureComponent,
    EditProcedureComponent,
    IndexRegulationComponent,
    CreateRegulationComponent,
    EditRegulationComponent,
    IndexObligationComponent,
    CreateObligationComponent,
    EditObligationComponent,
    IndexControlComponent,
    CreateControlComponent,
    EditControlComponent,
    IndexControlTest_Component,
    CreateControlTest_Component,
    EditControlTest_Component,
    IndexEvidenceComponent,
    CreateEvidenceComponent,
    EditEvidenceComponent,
    IndexRiskComponent,
    CreateRiskComponent,
    EditRiskComponent,
    IndexRiskAssessmentComponent,
    CreateRiskAssessmentComponent,
    EditRiskAssessmentComponent,
    IndexComplianceProgramComponent,
    CreateComplianceProgramComponent,
    EditComplianceProgramComponent,
    IndexComplianceRequirementComponent,
    CreateComplianceRequirementComponent,
    EditComplianceRequirementComponent,
    IndexAttestationComponent,
    CreateAttestationComponent,
    EditAttestationComponent,
    IndexAuditProgramComponent,
    CreateAuditProgramComponent,
    EditAuditProgramComponent,
    IndexAuditEngagementComponent,
    CreateAuditEngagementComponent,
    EditAuditEngagementComponent,
    IndexAuditWorkpaperComponent,
    CreateAuditWorkpaperComponent,
    EditAuditWorkpaperComponent,
    IndexAuditFindingComponent,
    CreateAuditFindingComponent,
    EditAuditFindingComponent,
    IndexCorrectiveActionComponent,
    CreateCorrectiveActionComponent,
    EditCorrectiveActionComponent,
    IndexIssueComponent,
    CreateIssueComponent,
    EditIssueComponent,
    IndexBusinessUnitComponent,
    CreateBusinessUnitComponent,
    EditBusinessUnitComponent,
    IndexDataProcessingActivityComponent,
    CreateDataProcessingActivityComponent,
    EditDataProcessingActivityComponent,
    IndexDataCategoryComponent,
    CreateDataCategoryComponent,
    EditDataCategoryComponent,
    IndexSystem_Component,
    CreateSystem_Component,
    EditSystem_Component,
    IndexPrivacyNoticeComponent,
    CreatePrivacyNoticeComponent,
    EditPrivacyNoticeComponent,
    IndexDataSubjectRequestComponent,
    CreateDataSubjectRequestComponent,
    EditDataSubjectRequestComponent,
    IndexRecordsRepositoryComponent,
    CreateRecordsRepositoryComponent,
    EditRecordsRepositoryComponent,
    IndexRecord_Component,
    CreateRecord_Component,
    EditRecord_Component,
    IndexRetentionScheduleComponent,
    CreateRetentionScheduleComponent,
    EditRetentionScheduleComponent,
    IndexDispositionReviewComponent,
    CreateDispositionReviewComponent,
    EditDispositionReviewComponent,
    IndexLegalHoldComponent,
    CreateLegalHoldComponent,
    EditLegalHoldComponent,
    IndexMatterComponent,
    CreateMatterComponent,
    EditMatterComponent,
    IndexThirdPartyComponent,
    CreateThirdPartyComponent,
    EditThirdPartyComponent,
    IndexThirdPartyAssessmentComponent,
    CreateThirdPartyAssessmentComponent,
    EditThirdPartyAssessmentComponent,
    IndexContractComponent,
    CreateContractComponent,
    EditContractComponent,
    IndexException_Component,
    CreateException_Component,
    EditException_Component,
    IndexConsentComponent,
    CreateConsentComponent,
    EditConsentComponent,
    IndexDataBreachComponent,
    CreateDataBreachComponent,
    EditDataBreachComponent,
    AppComponent
  ],
  imports: [

    BrowserModule, 
    NgbModule,
    MatMenuModule,
    MatToolbarModule,
    MatCheckboxModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatDatepickerModule,
	MatMomentDateModule,
    BrowserAnimationsModule,
	HttpClientModule, 
    ReactiveFormsModule,
    FormsModule,
    MatSidenavModule,    
    RouterModule.forRoot(appRoutes.OrganizationRoutes), 
    RouterModule.forRoot(appRoutes.GovernanceBodyRoutes), 
    RouterModule.forRoot(appRoutes.PersonRoutes), 
    RouterModule.forRoot(appRoutes.RoleRoutes), 
    RouterModule.forRoot(appRoutes.RoleAssignmentRoutes), 
    RouterModule.forRoot(appRoutes.PolicyRoutes), 
    RouterModule.forRoot(appRoutes.ProcedureRoutes), 
    RouterModule.forRoot(appRoutes.RegulationRoutes), 
    RouterModule.forRoot(appRoutes.ObligationRoutes), 
    RouterModule.forRoot(appRoutes.ControlRoutes), 
    RouterModule.forRoot(appRoutes.ControlTest_Routes), 
    RouterModule.forRoot(appRoutes.EvidenceRoutes), 
    RouterModule.forRoot(appRoutes.RiskRoutes), 
    RouterModule.forRoot(appRoutes.RiskAssessmentRoutes), 
    RouterModule.forRoot(appRoutes.ComplianceProgramRoutes), 
    RouterModule.forRoot(appRoutes.ComplianceRequirementRoutes), 
    RouterModule.forRoot(appRoutes.AttestationRoutes), 
    RouterModule.forRoot(appRoutes.AuditProgramRoutes), 
    RouterModule.forRoot(appRoutes.AuditEngagementRoutes), 
    RouterModule.forRoot(appRoutes.AuditWorkpaperRoutes), 
    RouterModule.forRoot(appRoutes.AuditFindingRoutes), 
    RouterModule.forRoot(appRoutes.CorrectiveActionRoutes), 
    RouterModule.forRoot(appRoutes.IssueRoutes), 
    RouterModule.forRoot(appRoutes.BusinessUnitRoutes), 
    RouterModule.forRoot(appRoutes.DataProcessingActivityRoutes), 
    RouterModule.forRoot(appRoutes.DataCategoryRoutes), 
    RouterModule.forRoot(appRoutes.System_Routes), 
    RouterModule.forRoot(appRoutes.PrivacyNoticeRoutes), 
    RouterModule.forRoot(appRoutes.DataSubjectRequestRoutes), 
    RouterModule.forRoot(appRoutes.RecordsRepositoryRoutes), 
    RouterModule.forRoot(appRoutes.Record_Routes), 
    RouterModule.forRoot(appRoutes.RetentionScheduleRoutes), 
    RouterModule.forRoot(appRoutes.DispositionReviewRoutes), 
    RouterModule.forRoot(appRoutes.LegalHoldRoutes), 
    RouterModule.forRoot(appRoutes.MatterRoutes), 
    RouterModule.forRoot(appRoutes.ThirdPartyRoutes), 
    RouterModule.forRoot(appRoutes.ThirdPartyAssessmentRoutes), 
    RouterModule.forRoot(appRoutes.ContractRoutes), 
    RouterModule.forRoot(appRoutes.Exception_Routes), 
    RouterModule.forRoot(appRoutes.ConsentRoutes), 
    RouterModule.forRoot(appRoutes.DataBreachRoutes), 
  ],
  providers: [OrganizationService,GovernanceBodyService,PersonService,RoleService,RoleAssignmentService,PolicyService,ProcedureService,RegulationService,ObligationService,ControlService,ControlTest_Service,EvidenceService,RiskService,RiskAssessmentService,ComplianceProgramService,ComplianceRequirementService,AttestationService,AuditProgramService,AuditEngagementService,AuditWorkpaperService,AuditFindingService,CorrectiveActionService,IssueService,BusinessUnitService,DataProcessingActivityService,DataCategoryService,System_Service,PrivacyNoticeService,DataSubjectRequestService,RecordsRepositoryService,Record_Service,RetentionScheduleService,DispositionReviewService,LegalHoldService,MatterService,ThirdPartyService,ThirdPartyAssessmentService,ContractService,Exception_Service,ConsentService,DataBreachService],
  bootstrap: [AppComponent]
})
export class AppModule { }
