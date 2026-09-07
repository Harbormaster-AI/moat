import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {OrganizationService} from '../services/Organization.service';
import {GovernanceBodyService} from '../services/GovernanceBody.service';
import {PersonService} from '../services/Person.service';
import {RoleService} from '../services/Role.service';
import {RoleAssignmentService} from '../services/RoleAssignment.service';
import {PolicyService} from '../services/Policy.service';
import {ProcedureService} from '../services/Procedure.service';
import {RegulationService} from '../services/Regulation.service';
import {ObligationService} from '../services/Obligation.service';
import {ControlService} from '../services/Control.service';
import {ControlTest_Service} from '../services/ControlTest_.service';
import {EvidenceService} from '../services/Evidence.service';
import {RiskService} from '../services/Risk.service';
import {RiskAssessmentService} from '../services/RiskAssessment.service';
import {ComplianceProgramService} from '../services/ComplianceProgram.service';
import {ComplianceRequirementService} from '../services/ComplianceRequirement.service';
import {AttestationService} from '../services/Attestation.service';
import {AuditProgramService} from '../services/AuditProgram.service';
import {AuditEngagementService} from '../services/AuditEngagement.service';
import {AuditWorkpaperService} from '../services/AuditWorkpaper.service';
import {AuditFindingService} from '../services/AuditFinding.service';
import {CorrectiveActionService} from '../services/CorrectiveAction.service';
import {IssueService} from '../services/Issue.service';
import {BusinessUnitService} from '../services/BusinessUnit.service';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {DataCategoryService} from '../services/DataCategory.service';
import {System_Service} from '../services/System_.service';
import {PrivacyNoticeService} from '../services/PrivacyNotice.service';
import {DataSubjectRequestService} from '../services/DataSubjectRequest.service';
import {RecordsRepositoryService} from '../services/RecordsRepository.service';
import {Record_Service} from '../services/Record_.service';
import {RetentionScheduleService} from '../services/RetentionSchedule.service';
import {DispositionReviewService} from '../services/DispositionReview.service';
import {LegalHoldService} from '../services/LegalHold.service';
import {MatterService} from '../services/Matter.service';
import {ThirdPartyService} from '../services/ThirdParty.service';
import {ThirdPartyAssessmentService} from '../services/ThirdPartyAssessment.service';
import {ContractService} from '../services/Contract.service';
import {Exception_Service} from '../services/Exception_.service';
import {ConsentService} from '../services/Consent.service';
import {DataBreachService} from '../services/DataBreach.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    GovernanceBodyTypes = Object.keys(enumTypes.GovernanceBodyType);
    PolicyTypes = Object.keys(enumTypes.PolicyType);
    DocumentStatuss = Object.keys(enumTypes.DocumentStatus);
    ControlTypes = Object.keys(enumTypes.ControlType);
    ControlFrequencys = Object.keys(enumTypes.ControlFrequency);
    ControlStatuss = Object.keys(enumTypes.ControlStatus);
    RiskCategorys = Object.keys(enumTypes.RiskCategory);
    RiskImpacts = Object.keys(enumTypes.RiskImpact);
    RiskLikelihoods = Object.keys(enumTypes.RiskLikelihood);
    RiskStatuss = Object.keys(enumTypes.RiskStatus);
    AssessmentTypes = Object.keys(enumTypes.AssessmentType);
    TestTypes = Object.keys(enumTypes.TestType);
    ControlEffectivenesss = Object.keys(enumTypes.ControlEffectiveness);
    TestStatuss = Object.keys(enumTypes.TestStatus);
    EvidenceTypes = Object.keys(enumTypes.EvidenceType);
    AuditCycles = Object.keys(enumTypes.AuditCycle);
    AuditStatuss = Object.keys(enumTypes.AuditStatus);
    FindingSeveritys = Object.keys(enumTypes.FindingSeverity);
    FindingStatuss = Object.keys(enumTypes.FindingStatus);
    ActionStatuss = Object.keys(enumTypes.ActionStatus);
    IssueTypes = Object.keys(enumTypes.IssueType);
    Prioritys = Object.keys(enumTypes.Priority);
    IssueStatuss = Object.keys(enumTypes.IssueStatus);
    ComplianceStatuss = Object.keys(enumTypes.ComplianceStatus);
    Applicabilitys = Object.keys(enumTypes.Applicability);
    AttestationResults = Object.keys(enumTypes.AttestationResult);
    LawfulBasiss = Object.keys(enumTypes.LawfulBasis);
    DataClassificationLevels = Object.keys(enumTypes.DataClassificationLevel);
    SystemTypes = Object.keys(enumTypes.SystemType);
    DataSubjectRequestTypes = Object.keys(enumTypes.DataSubjectRequestType);
    RequestStatuss = Object.keys(enumTypes.RequestStatus);
    RepositoryTypes = Object.keys(enumTypes.RepositoryType);
    RecordTypes = Object.keys(enumTypes.RecordType);
    RecordStatuss = Object.keys(enumTypes.RecordStatus);
    RetentionTriggers = Object.keys(enumTypes.RetentionTrigger);
    DispositionActions = Object.keys(enumTypes.DispositionAction);
    RetentionStatuss = Object.keys(enumTypes.RetentionStatus);
    DispositionOutcomes = Object.keys(enumTypes.DispositionOutcome);
    LegalHoldStatuss = Object.keys(enumTypes.LegalHoldStatus);
    MatterTypes = Object.keys(enumTypes.MatterType);
    MatterStatuss = Object.keys(enumTypes.MatterStatus);
    ThirdPartyTypes = Object.keys(enumTypes.ThirdPartyType);
    VendorCriticalitys = Object.keys(enumTypes.VendorCriticality);
    AssessmentResults = Object.keys(enumTypes.AssessmentResult);
    ContractStatuss = Object.keys(enumTypes.ContractStatus);
    ObligationTypes = Object.keys(enumTypes.ObligationType);
    ExceptionTypes = Object.keys(enumTypes.ExceptionType);
    ExceptionStatuss = Object.keys(enumTypes.ExceptionStatus);
    ConsentTypes = Object.keys(enumTypes.ConsentType);
    ConsentStatuss = Object.keys(enumTypes.ConsentStatus);
    BreachSeveritys = Object.keys(enumTypes.BreachSeverity);
    IncidentStatuss = Object.keys(enumTypes.IncidentStatus);

// all collection instances
    organizations : any;
    governanceBodys : any;
    persons : any;
    roles : any;
    roleAssignments : any;
    policys : any;
    procedures : any;
    regulations : any;
    obligations : any;
    controls : any;
    controlTest_s : any;
    evidences : any;
    risks : any;
    riskAssessments : any;
    compliancePrograms : any;
    complianceRequirements : any;
    attestations : any;
    auditPrograms : any;
    auditEngagements : any;
    auditWorkpapers : any;
    auditFindings : any;
    correctiveActions : any;
    issues : any;
    businessUnits : any;
    dataProcessingActivitys : any;
    dataCategorys : any;
    system_s : any;
    privacyNotices : any;
    dataSubjectRequests : any;
    recordsRepositorys : any;
    record_s : any;
    retentionSchedules : any;
    dispositionReviews : any;
    legalHolds : any;
    matters : any;
    thirdPartys : any;
    thirdPartyAssessments : any;
    contracts : any;
    exception_s : any;
    consents : any;
    dataBreachs : any;
  
// initialization  
    ngOnInit() {
    }

    initOrganizationList() {
        if ( this.organizations == null ) {
            new OrganizationService(this.http).getOrganizations().subscribe(res => {
                this.organizations = res;
            });
        }
    }
    
    initGovernanceBodyList() {
        if ( this.governanceBodys == null ) {
            new GovernanceBodyService(this.http).getGovernanceBodys().subscribe(res => {
                this.governanceBodys = res;
            });
        }
    }
    
    initPersonList() {
        if ( this.persons == null ) {
            new PersonService(this.http).getPersons().subscribe(res => {
                this.persons = res;
            });
        }
    }
    
    initRoleList() {
        if ( this.roles == null ) {
            new RoleService(this.http).getRoles().subscribe(res => {
                this.roles = res;
            });
        }
    }
    
    initRoleAssignmentList() {
        if ( this.roleAssignments == null ) {
            new RoleAssignmentService(this.http).getRoleAssignments().subscribe(res => {
                this.roleAssignments = res;
            });
        }
    }
    
    initPolicyList() {
        if ( this.policys == null ) {
            new PolicyService(this.http).getPolicys().subscribe(res => {
                this.policys = res;
            });
        }
    }
    
    initProcedureList() {
        if ( this.procedures == null ) {
            new ProcedureService(this.http).getProcedures().subscribe(res => {
                this.procedures = res;
            });
        }
    }
    
    initRegulationList() {
        if ( this.regulations == null ) {
            new RegulationService(this.http).getRegulations().subscribe(res => {
                this.regulations = res;
            });
        }
    }
    
    initObligationList() {
        if ( this.obligations == null ) {
            new ObligationService(this.http).getObligations().subscribe(res => {
                this.obligations = res;
            });
        }
    }
    
    initControlList() {
        if ( this.controls == null ) {
            new ControlService(this.http).getControls().subscribe(res => {
                this.controls = res;
            });
        }
    }
    
    initControlTest_List() {
        if ( this.controlTest_s == null ) {
            new ControlTest_Service(this.http).getControlTest_s().subscribe(res => {
                this.controlTest_s = res;
            });
        }
    }
    
    initEvidenceList() {
        if ( this.evidences == null ) {
            new EvidenceService(this.http).getEvidences().subscribe(res => {
                this.evidences = res;
            });
        }
    }
    
    initRiskList() {
        if ( this.risks == null ) {
            new RiskService(this.http).getRisks().subscribe(res => {
                this.risks = res;
            });
        }
    }
    
    initRiskAssessmentList() {
        if ( this.riskAssessments == null ) {
            new RiskAssessmentService(this.http).getRiskAssessments().subscribe(res => {
                this.riskAssessments = res;
            });
        }
    }
    
    initComplianceProgramList() {
        if ( this.compliancePrograms == null ) {
            new ComplianceProgramService(this.http).getCompliancePrograms().subscribe(res => {
                this.compliancePrograms = res;
            });
        }
    }
    
    initComplianceRequirementList() {
        if ( this.complianceRequirements == null ) {
            new ComplianceRequirementService(this.http).getComplianceRequirements().subscribe(res => {
                this.complianceRequirements = res;
            });
        }
    }
    
    initAttestationList() {
        if ( this.attestations == null ) {
            new AttestationService(this.http).getAttestations().subscribe(res => {
                this.attestations = res;
            });
        }
    }
    
    initAuditProgramList() {
        if ( this.auditPrograms == null ) {
            new AuditProgramService(this.http).getAuditPrograms().subscribe(res => {
                this.auditPrograms = res;
            });
        }
    }
    
    initAuditEngagementList() {
        if ( this.auditEngagements == null ) {
            new AuditEngagementService(this.http).getAuditEngagements().subscribe(res => {
                this.auditEngagements = res;
            });
        }
    }
    
    initAuditWorkpaperList() {
        if ( this.auditWorkpapers == null ) {
            new AuditWorkpaperService(this.http).getAuditWorkpapers().subscribe(res => {
                this.auditWorkpapers = res;
            });
        }
    }
    
    initAuditFindingList() {
        if ( this.auditFindings == null ) {
            new AuditFindingService(this.http).getAuditFindings().subscribe(res => {
                this.auditFindings = res;
            });
        }
    }
    
    initCorrectiveActionList() {
        if ( this.correctiveActions == null ) {
            new CorrectiveActionService(this.http).getCorrectiveActions().subscribe(res => {
                this.correctiveActions = res;
            });
        }
    }
    
    initIssueList() {
        if ( this.issues == null ) {
            new IssueService(this.http).getIssues().subscribe(res => {
                this.issues = res;
            });
        }
    }
    
    initBusinessUnitList() {
        if ( this.businessUnits == null ) {
            new BusinessUnitService(this.http).getBusinessUnits().subscribe(res => {
                this.businessUnits = res;
            });
        }
    }
    
    initDataProcessingActivityList() {
        if ( this.dataProcessingActivitys == null ) {
            new DataProcessingActivityService(this.http).getDataProcessingActivitys().subscribe(res => {
                this.dataProcessingActivitys = res;
            });
        }
    }
    
    initDataCategoryList() {
        if ( this.dataCategorys == null ) {
            new DataCategoryService(this.http).getDataCategorys().subscribe(res => {
                this.dataCategorys = res;
            });
        }
    }
    
    initSystem_List() {
        if ( this.system_s == null ) {
            new System_Service(this.http).getSystem_s().subscribe(res => {
                this.system_s = res;
            });
        }
    }
    
    initPrivacyNoticeList() {
        if ( this.privacyNotices == null ) {
            new PrivacyNoticeService(this.http).getPrivacyNotices().subscribe(res => {
                this.privacyNotices = res;
            });
        }
    }
    
    initDataSubjectRequestList() {
        if ( this.dataSubjectRequests == null ) {
            new DataSubjectRequestService(this.http).getDataSubjectRequests().subscribe(res => {
                this.dataSubjectRequests = res;
            });
        }
    }
    
    initRecordsRepositoryList() {
        if ( this.recordsRepositorys == null ) {
            new RecordsRepositoryService(this.http).getRecordsRepositorys().subscribe(res => {
                this.recordsRepositorys = res;
            });
        }
    }
    
    initRecord_List() {
        if ( this.record_s == null ) {
            new Record_Service(this.http).getRecord_s().subscribe(res => {
                this.record_s = res;
            });
        }
    }
    
    initRetentionScheduleList() {
        if ( this.retentionSchedules == null ) {
            new RetentionScheduleService(this.http).getRetentionSchedules().subscribe(res => {
                this.retentionSchedules = res;
            });
        }
    }
    
    initDispositionReviewList() {
        if ( this.dispositionReviews == null ) {
            new DispositionReviewService(this.http).getDispositionReviews().subscribe(res => {
                this.dispositionReviews = res;
            });
        }
    }
    
    initLegalHoldList() {
        if ( this.legalHolds == null ) {
            new LegalHoldService(this.http).getLegalHolds().subscribe(res => {
                this.legalHolds = res;
            });
        }
    }
    
    initMatterList() {
        if ( this.matters == null ) {
            new MatterService(this.http).getMatters().subscribe(res => {
                this.matters = res;
            });
        }
    }
    
    initThirdPartyList() {
        if ( this.thirdPartys == null ) {
            new ThirdPartyService(this.http).getThirdPartys().subscribe(res => {
                this.thirdPartys = res;
            });
        }
    }
    
    initThirdPartyAssessmentList() {
        if ( this.thirdPartyAssessments == null ) {
            new ThirdPartyAssessmentService(this.http).getThirdPartyAssessments().subscribe(res => {
                this.thirdPartyAssessments = res;
            });
        }
    }
    
    initContractList() {
        if ( this.contracts == null ) {
            new ContractService(this.http).getContracts().subscribe(res => {
                this.contracts = res;
            });
        }
    }
    
    initException_List() {
        if ( this.exception_s == null ) {
            new Exception_Service(this.http).getException_s().subscribe(res => {
                this.exception_s = res;
            });
        }
    }
    
    initConsentList() {
        if ( this.consents == null ) {
            new ConsentService(this.http).getConsents().subscribe(res => {
                this.consents = res;
            });
        }
    }
    
    initDataBreachList() {
        if ( this.dataBreachs == null ) {
            new DataBreachService(this.http).getDataBreachs().subscribe(res => {
                this.dataBreachs = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
