import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {OrganizationService} from '../services/Organization.service';
import {DepartmentService} from '../services/Department.service';
import {LocationService} from '../services/Location.service';
import {CostCenterService} from '../services/CostCenter.service';
import {JobFamilyService} from '../services/JobFamily.service';
import {JobProfileService} from '../services/JobProfile.service';
import {CompetencyService} from '../services/Competency.service';
import {PositionService} from '../services/Position.service';
import {EmployeeService} from '../services/Employee.service';
import {EmploymentAssignmentService} from '../services/EmploymentAssignment.service';
import {EmploymentContractService} from '../services/EmploymentContract.service';
import {WorkScheduleService} from '../services/WorkSchedule.service';
import {WorkShiftService} from '../services/WorkShift.service';
import {ScheduleExceptionService} from '../services/ScheduleException.service';
import {CompensationPackageService} from '../services/CompensationPackage.service';
import {SalaryComponentService} from '../services/SalaryComponent.service';
import {BonusPlanService} from '../services/BonusPlan.service';
import {EquityGrantService} from '../services/EquityGrant.service';
import {BenefitPlanService} from '../services/BenefitPlan.service';
import {BenefitEnrollmentService} from '../services/BenefitEnrollment.service';
import {DependentService} from '../services/Dependent.service';
import {PayrollCalendarService} from '../services/PayrollCalendar.service';
import {PayrollRunService} from '../services/PayrollRun.service';
import {PayrollItemService} from '../services/PayrollItem.service';
import {TaxWithholdingService} from '../services/TaxWithholding.service';
import {PaymentMethodService} from '../services/PaymentMethod.service';
import {TimesheetService} from '../services/Timesheet.service';
import {TimeEntryService} from '../services/TimeEntry.service';
import {ApprovalService} from '../services/Approval.service';
import {LeavePolicyService} from '../services/LeavePolicy.service';
import {LeaveRequestService} from '../services/LeaveRequest.service';
import {PerformanceCycleService} from '../services/PerformanceCycle.service';
import {GoalService} from '../services/Goal.service';
import {PerformanceReviewService} from '../services/PerformanceReview.service';
import {CompetencyRatingService} from '../services/CompetencyRating.service';
import {TrainingCourseService} from '../services/TrainingCourse.service';
import {TrainingEnrollmentService} from '../services/TrainingEnrollment.service';
import {CertificationService} from '../services/Certification.service';
import {JobRequisitionService} from '../services/JobRequisition.service';
import {CandidateService} from '../services/Candidate.service';
import {JobApplicationService} from '../services/JobApplication.service';
import {InterviewService} from '../services/Interview.service';
import {ScreeningService} from '../services/Screening.service';
import {OfferService} from '../services/Offer.service';
import {OnboardingTaskService} from '../services/OnboardingTask.service';
import {BackgroundCheckService} from '../services/BackgroundCheck.service';
import {DocumentService} from '../services/Document.service';
import {PolicyService} from '../services/Policy.service';
import {PolicyAcknowledgementService} from '../services/PolicyAcknowledgement.service';
import {TerminationService} from '../services/Termination.service';
import {WorkAuthorizationService} from '../services/WorkAuthorization.service';
import {BankAccountService} from '../services/BankAccount.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    JobLevels = Object.keys(enumTypes.JobLevel);
    ExemptStatuss = Object.keys(enumTypes.ExemptStatus);
    PositionStatuss = Object.keys(enumTypes.PositionStatus);
    WorkLocationTypes = Object.keys(enumTypes.WorkLocationType);
    EmploymentStatuss = Object.keys(enumTypes.EmploymentStatus);
    AssignmentTypes = Object.keys(enumTypes.AssignmentType);
    AssignmentStatuss = Object.keys(enumTypes.AssignmentStatus);
    EmploymentTypes = Object.keys(enumTypes.EmploymentType);
    ContractStatuss = Object.keys(enumTypes.ContractStatus);
    PayFrequencys = Object.keys(enumTypes.PayFrequency);
    ScheduleTypes = Object.keys(enumTypes.ScheduleType);
    DayOfWeeks = Object.keys(enumTypes.DayOfWeek);
    SalaryComponentTypes = Object.keys(enumTypes.SalaryComponentType);
    EquityTypes = Object.keys(enumTypes.EquityType);
    BenefitTypes = Object.keys(enumTypes.BenefitType);
    BenefitEnrollmentStatuss = Object.keys(enumTypes.BenefitEnrollmentStatus);
    CoverageLevels = Object.keys(enumTypes.CoverageLevel);
    DependentRelationships = Object.keys(enumTypes.DependentRelationship);
    TimeEntryTypes = Object.keys(enumTypes.TimeEntryType);
    TimesheetStatuss = Object.keys(enumTypes.TimesheetStatus);
    ApprovalStatuss = Object.keys(enumTypes.ApprovalStatus);
    LeaveCategorys = Object.keys(enumTypes.LeaveCategory);
    AccrualUnits = Object.keys(enumTypes.AccrualUnit);
    LeaveStatuss = Object.keys(enumTypes.LeaveStatus);
    PayrollStatuss = Object.keys(enumTypes.PayrollStatus);
    PayrollItemTypes = Object.keys(enumTypes.PayrollItemType);
    FilingStatuss = Object.keys(enumTypes.FilingStatus);
    CycleStatuss = Object.keys(enumTypes.CycleStatus);
    GoalStatuss = Object.keys(enumTypes.GoalStatus);
    PerformanceRatings = Object.keys(enumTypes.PerformanceRating);
    ReviewStatuss = Object.keys(enumTypes.ReviewStatus);
    DeliveryMethods = Object.keys(enumTypes.DeliveryMethod);
    TrainingStatuss = Object.keys(enumTypes.TrainingStatus);
    DisciplinaryActionTypes = Object.keys(enumTypes.DisciplinaryActionType);
    RequisitionStatuss = Object.keys(enumTypes.RequisitionStatus);
    RequisitionPrioritys = Object.keys(enumTypes.RequisitionPriority);
    CandidateSources = Object.keys(enumTypes.CandidateSource);
    ApplicationStatuss = Object.keys(enumTypes.ApplicationStatus);
    InterviewStages = Object.keys(enumTypes.InterviewStage);
    InterviewResults = Object.keys(enumTypes.InterviewResult);
    OfferStatuss = Object.keys(enumTypes.OfferStatus);
    OnboardingTaskStatuss = Object.keys(enumTypes.OnboardingTaskStatus);
    BackgroundCheckStatuss = Object.keys(enumTypes.BackgroundCheckStatus);
    DocumentTypes = Object.keys(enumTypes.DocumentType);
    AcknowledgementStatuss = Object.keys(enumTypes.AcknowledgementStatus);
    TerminationReasons = Object.keys(enumTypes.TerminationReason);
    TerminationTypes = Object.keys(enumTypes.TerminationType);
    WorkAuthorizationStatuss = Object.keys(enumTypes.WorkAuthorizationStatus);
    PaymentMethodTypes = Object.keys(enumTypes.PaymentMethodType);

// all collection instances
    organizations : any;
    departments : any;
    locations : any;
    costCenters : any;
    jobFamilys : any;
    jobProfiles : any;
    competencys : any;
    positions : any;
    employees : any;
    employmentAssignments : any;
    employmentContracts : any;
    workSchedules : any;
    workShifts : any;
    scheduleExceptions : any;
    compensationPackages : any;
    salaryComponents : any;
    bonusPlans : any;
    equityGrants : any;
    benefitPlans : any;
    benefitEnrollments : any;
    dependents : any;
    payrollCalendars : any;
    payrollRuns : any;
    payrollItems : any;
    taxWithholdings : any;
    paymentMethods : any;
    timesheets : any;
    timeEntrys : any;
    approvals : any;
    leavePolicys : any;
    leaveRequests : any;
    performanceCycles : any;
    goals : any;
    performanceReviews : any;
    competencyRatings : any;
    trainingCourses : any;
    trainingEnrollments : any;
    certifications : any;
    jobRequisitions : any;
    candidates : any;
    jobApplications : any;
    interviews : any;
    screenings : any;
    offers : any;
    onboardingTasks : any;
    backgroundChecks : any;
    documents : any;
    policys : any;
    policyAcknowledgements : any;
    terminations : any;
    workAuthorizations : any;
    bankAccounts : any;
  
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
    
    initDepartmentList() {
        if ( this.departments == null ) {
            new DepartmentService(this.http).getDepartments().subscribe(res => {
                this.departments = res;
            });
        }
    }
    
    initLocationList() {
        if ( this.locations == null ) {
            new LocationService(this.http).getLocations().subscribe(res => {
                this.locations = res;
            });
        }
    }
    
    initCostCenterList() {
        if ( this.costCenters == null ) {
            new CostCenterService(this.http).getCostCenters().subscribe(res => {
                this.costCenters = res;
            });
        }
    }
    
    initJobFamilyList() {
        if ( this.jobFamilys == null ) {
            new JobFamilyService(this.http).getJobFamilys().subscribe(res => {
                this.jobFamilys = res;
            });
        }
    }
    
    initJobProfileList() {
        if ( this.jobProfiles == null ) {
            new JobProfileService(this.http).getJobProfiles().subscribe(res => {
                this.jobProfiles = res;
            });
        }
    }
    
    initCompetencyList() {
        if ( this.competencys == null ) {
            new CompetencyService(this.http).getCompetencys().subscribe(res => {
                this.competencys = res;
            });
        }
    }
    
    initPositionList() {
        if ( this.positions == null ) {
            new PositionService(this.http).getPositions().subscribe(res => {
                this.positions = res;
            });
        }
    }
    
    initEmployeeList() {
        if ( this.employees == null ) {
            new EmployeeService(this.http).getEmployees().subscribe(res => {
                this.employees = res;
            });
        }
    }
    
    initEmploymentAssignmentList() {
        if ( this.employmentAssignments == null ) {
            new EmploymentAssignmentService(this.http).getEmploymentAssignments().subscribe(res => {
                this.employmentAssignments = res;
            });
        }
    }
    
    initEmploymentContractList() {
        if ( this.employmentContracts == null ) {
            new EmploymentContractService(this.http).getEmploymentContracts().subscribe(res => {
                this.employmentContracts = res;
            });
        }
    }
    
    initWorkScheduleList() {
        if ( this.workSchedules == null ) {
            new WorkScheduleService(this.http).getWorkSchedules().subscribe(res => {
                this.workSchedules = res;
            });
        }
    }
    
    initWorkShiftList() {
        if ( this.workShifts == null ) {
            new WorkShiftService(this.http).getWorkShifts().subscribe(res => {
                this.workShifts = res;
            });
        }
    }
    
    initScheduleExceptionList() {
        if ( this.scheduleExceptions == null ) {
            new ScheduleExceptionService(this.http).getScheduleExceptions().subscribe(res => {
                this.scheduleExceptions = res;
            });
        }
    }
    
    initCompensationPackageList() {
        if ( this.compensationPackages == null ) {
            new CompensationPackageService(this.http).getCompensationPackages().subscribe(res => {
                this.compensationPackages = res;
            });
        }
    }
    
    initSalaryComponentList() {
        if ( this.salaryComponents == null ) {
            new SalaryComponentService(this.http).getSalaryComponents().subscribe(res => {
                this.salaryComponents = res;
            });
        }
    }
    
    initBonusPlanList() {
        if ( this.bonusPlans == null ) {
            new BonusPlanService(this.http).getBonusPlans().subscribe(res => {
                this.bonusPlans = res;
            });
        }
    }
    
    initEquityGrantList() {
        if ( this.equityGrants == null ) {
            new EquityGrantService(this.http).getEquityGrants().subscribe(res => {
                this.equityGrants = res;
            });
        }
    }
    
    initBenefitPlanList() {
        if ( this.benefitPlans == null ) {
            new BenefitPlanService(this.http).getBenefitPlans().subscribe(res => {
                this.benefitPlans = res;
            });
        }
    }
    
    initBenefitEnrollmentList() {
        if ( this.benefitEnrollments == null ) {
            new BenefitEnrollmentService(this.http).getBenefitEnrollments().subscribe(res => {
                this.benefitEnrollments = res;
            });
        }
    }
    
    initDependentList() {
        if ( this.dependents == null ) {
            new DependentService(this.http).getDependents().subscribe(res => {
                this.dependents = res;
            });
        }
    }
    
    initPayrollCalendarList() {
        if ( this.payrollCalendars == null ) {
            new PayrollCalendarService(this.http).getPayrollCalendars().subscribe(res => {
                this.payrollCalendars = res;
            });
        }
    }
    
    initPayrollRunList() {
        if ( this.payrollRuns == null ) {
            new PayrollRunService(this.http).getPayrollRuns().subscribe(res => {
                this.payrollRuns = res;
            });
        }
    }
    
    initPayrollItemList() {
        if ( this.payrollItems == null ) {
            new PayrollItemService(this.http).getPayrollItems().subscribe(res => {
                this.payrollItems = res;
            });
        }
    }
    
    initTaxWithholdingList() {
        if ( this.taxWithholdings == null ) {
            new TaxWithholdingService(this.http).getTaxWithholdings().subscribe(res => {
                this.taxWithholdings = res;
            });
        }
    }
    
    initPaymentMethodList() {
        if ( this.paymentMethods == null ) {
            new PaymentMethodService(this.http).getPaymentMethods().subscribe(res => {
                this.paymentMethods = res;
            });
        }
    }
    
    initTimesheetList() {
        if ( this.timesheets == null ) {
            new TimesheetService(this.http).getTimesheets().subscribe(res => {
                this.timesheets = res;
            });
        }
    }
    
    initTimeEntryList() {
        if ( this.timeEntrys == null ) {
            new TimeEntryService(this.http).getTimeEntrys().subscribe(res => {
                this.timeEntrys = res;
            });
        }
    }
    
    initApprovalList() {
        if ( this.approvals == null ) {
            new ApprovalService(this.http).getApprovals().subscribe(res => {
                this.approvals = res;
            });
        }
    }
    
    initLeavePolicyList() {
        if ( this.leavePolicys == null ) {
            new LeavePolicyService(this.http).getLeavePolicys().subscribe(res => {
                this.leavePolicys = res;
            });
        }
    }
    
    initLeaveRequestList() {
        if ( this.leaveRequests == null ) {
            new LeaveRequestService(this.http).getLeaveRequests().subscribe(res => {
                this.leaveRequests = res;
            });
        }
    }
    
    initPerformanceCycleList() {
        if ( this.performanceCycles == null ) {
            new PerformanceCycleService(this.http).getPerformanceCycles().subscribe(res => {
                this.performanceCycles = res;
            });
        }
    }
    
    initGoalList() {
        if ( this.goals == null ) {
            new GoalService(this.http).getGoals().subscribe(res => {
                this.goals = res;
            });
        }
    }
    
    initPerformanceReviewList() {
        if ( this.performanceReviews == null ) {
            new PerformanceReviewService(this.http).getPerformanceReviews().subscribe(res => {
                this.performanceReviews = res;
            });
        }
    }
    
    initCompetencyRatingList() {
        if ( this.competencyRatings == null ) {
            new CompetencyRatingService(this.http).getCompetencyRatings().subscribe(res => {
                this.competencyRatings = res;
            });
        }
    }
    
    initTrainingCourseList() {
        if ( this.trainingCourses == null ) {
            new TrainingCourseService(this.http).getTrainingCourses().subscribe(res => {
                this.trainingCourses = res;
            });
        }
    }
    
    initTrainingEnrollmentList() {
        if ( this.trainingEnrollments == null ) {
            new TrainingEnrollmentService(this.http).getTrainingEnrollments().subscribe(res => {
                this.trainingEnrollments = res;
            });
        }
    }
    
    initCertificationList() {
        if ( this.certifications == null ) {
            new CertificationService(this.http).getCertifications().subscribe(res => {
                this.certifications = res;
            });
        }
    }
    
    initJobRequisitionList() {
        if ( this.jobRequisitions == null ) {
            new JobRequisitionService(this.http).getJobRequisitions().subscribe(res => {
                this.jobRequisitions = res;
            });
        }
    }
    
    initCandidateList() {
        if ( this.candidates == null ) {
            new CandidateService(this.http).getCandidates().subscribe(res => {
                this.candidates = res;
            });
        }
    }
    
    initJobApplicationList() {
        if ( this.jobApplications == null ) {
            new JobApplicationService(this.http).getJobApplications().subscribe(res => {
                this.jobApplications = res;
            });
        }
    }
    
    initInterviewList() {
        if ( this.interviews == null ) {
            new InterviewService(this.http).getInterviews().subscribe(res => {
                this.interviews = res;
            });
        }
    }
    
    initScreeningList() {
        if ( this.screenings == null ) {
            new ScreeningService(this.http).getScreenings().subscribe(res => {
                this.screenings = res;
            });
        }
    }
    
    initOfferList() {
        if ( this.offers == null ) {
            new OfferService(this.http).getOffers().subscribe(res => {
                this.offers = res;
            });
        }
    }
    
    initOnboardingTaskList() {
        if ( this.onboardingTasks == null ) {
            new OnboardingTaskService(this.http).getOnboardingTasks().subscribe(res => {
                this.onboardingTasks = res;
            });
        }
    }
    
    initBackgroundCheckList() {
        if ( this.backgroundChecks == null ) {
            new BackgroundCheckService(this.http).getBackgroundChecks().subscribe(res => {
                this.backgroundChecks = res;
            });
        }
    }
    
    initDocumentList() {
        if ( this.documents == null ) {
            new DocumentService(this.http).getDocuments().subscribe(res => {
                this.documents = res;
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
    
    initPolicyAcknowledgementList() {
        if ( this.policyAcknowledgements == null ) {
            new PolicyAcknowledgementService(this.http).getPolicyAcknowledgements().subscribe(res => {
                this.policyAcknowledgements = res;
            });
        }
    }
    
    initTerminationList() {
        if ( this.terminations == null ) {
            new TerminationService(this.http).getTerminations().subscribe(res => {
                this.terminations = res;
            });
        }
    }
    
    initWorkAuthorizationList() {
        if ( this.workAuthorizations == null ) {
            new WorkAuthorizationService(this.http).getWorkAuthorizations().subscribe(res => {
                this.workAuthorizations = res;
            });
        }
    }
    
    initBankAccountList() {
        if ( this.bankAccounts == null ) {
            new BankAccountService(this.http).getBankAccounts().subscribe(res => {
                this.bankAccounts = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
