import { Routes } from '@angular/router';

import { CreateOrganizationComponent } from './components/Organization/create/create.component';
import { EditOrganizationComponent } from './components/Organization/edit/edit.component';
import { IndexOrganizationComponent } from './components/Organization/index/index.component';
import { CreateDepartmentComponent } from './components/Department/create/create.component';
import { EditDepartmentComponent } from './components/Department/edit/edit.component';
import { IndexDepartmentComponent } from './components/Department/index/index.component';
import { CreateLocationComponent } from './components/Location/create/create.component';
import { EditLocationComponent } from './components/Location/edit/edit.component';
import { IndexLocationComponent } from './components/Location/index/index.component';
import { CreateCostCenterComponent } from './components/CostCenter/create/create.component';
import { EditCostCenterComponent } from './components/CostCenter/edit/edit.component';
import { IndexCostCenterComponent } from './components/CostCenter/index/index.component';
import { CreateJobFamilyComponent } from './components/JobFamily/create/create.component';
import { EditJobFamilyComponent } from './components/JobFamily/edit/edit.component';
import { IndexJobFamilyComponent } from './components/JobFamily/index/index.component';
import { CreateJobProfileComponent } from './components/JobProfile/create/create.component';
import { EditJobProfileComponent } from './components/JobProfile/edit/edit.component';
import { IndexJobProfileComponent } from './components/JobProfile/index/index.component';
import { CreateCompetencyComponent } from './components/Competency/create/create.component';
import { EditCompetencyComponent } from './components/Competency/edit/edit.component';
import { IndexCompetencyComponent } from './components/Competency/index/index.component';
import { CreatePositionComponent } from './components/Position/create/create.component';
import { EditPositionComponent } from './components/Position/edit/edit.component';
import { IndexPositionComponent } from './components/Position/index/index.component';
import { CreateEmployeeComponent } from './components/Employee/create/create.component';
import { EditEmployeeComponent } from './components/Employee/edit/edit.component';
import { IndexEmployeeComponent } from './components/Employee/index/index.component';
import { CreateEmploymentAssignmentComponent } from './components/EmploymentAssignment/create/create.component';
import { EditEmploymentAssignmentComponent } from './components/EmploymentAssignment/edit/edit.component';
import { IndexEmploymentAssignmentComponent } from './components/EmploymentAssignment/index/index.component';
import { CreateEmploymentContractComponent } from './components/EmploymentContract/create/create.component';
import { EditEmploymentContractComponent } from './components/EmploymentContract/edit/edit.component';
import { IndexEmploymentContractComponent } from './components/EmploymentContract/index/index.component';
import { CreateWorkScheduleComponent } from './components/WorkSchedule/create/create.component';
import { EditWorkScheduleComponent } from './components/WorkSchedule/edit/edit.component';
import { IndexWorkScheduleComponent } from './components/WorkSchedule/index/index.component';
import { CreateWorkShiftComponent } from './components/WorkShift/create/create.component';
import { EditWorkShiftComponent } from './components/WorkShift/edit/edit.component';
import { IndexWorkShiftComponent } from './components/WorkShift/index/index.component';
import { CreateScheduleExceptionComponent } from './components/ScheduleException/create/create.component';
import { EditScheduleExceptionComponent } from './components/ScheduleException/edit/edit.component';
import { IndexScheduleExceptionComponent } from './components/ScheduleException/index/index.component';
import { CreateCompensationPackageComponent } from './components/CompensationPackage/create/create.component';
import { EditCompensationPackageComponent } from './components/CompensationPackage/edit/edit.component';
import { IndexCompensationPackageComponent } from './components/CompensationPackage/index/index.component';
import { CreateSalaryComponentComponent } from './components/SalaryComponent/create/create.component';
import { EditSalaryComponentComponent } from './components/SalaryComponent/edit/edit.component';
import { IndexSalaryComponentComponent } from './components/SalaryComponent/index/index.component';
import { CreateBonusPlanComponent } from './components/BonusPlan/create/create.component';
import { EditBonusPlanComponent } from './components/BonusPlan/edit/edit.component';
import { IndexBonusPlanComponent } from './components/BonusPlan/index/index.component';
import { CreateEquityGrantComponent } from './components/EquityGrant/create/create.component';
import { EditEquityGrantComponent } from './components/EquityGrant/edit/edit.component';
import { IndexEquityGrantComponent } from './components/EquityGrant/index/index.component';
import { CreateBenefitPlanComponent } from './components/BenefitPlan/create/create.component';
import { EditBenefitPlanComponent } from './components/BenefitPlan/edit/edit.component';
import { IndexBenefitPlanComponent } from './components/BenefitPlan/index/index.component';
import { CreateBenefitEnrollmentComponent } from './components/BenefitEnrollment/create/create.component';
import { EditBenefitEnrollmentComponent } from './components/BenefitEnrollment/edit/edit.component';
import { IndexBenefitEnrollmentComponent } from './components/BenefitEnrollment/index/index.component';
import { CreateDependentComponent } from './components/Dependent/create/create.component';
import { EditDependentComponent } from './components/Dependent/edit/edit.component';
import { IndexDependentComponent } from './components/Dependent/index/index.component';
import { CreatePayrollCalendarComponent } from './components/PayrollCalendar/create/create.component';
import { EditPayrollCalendarComponent } from './components/PayrollCalendar/edit/edit.component';
import { IndexPayrollCalendarComponent } from './components/PayrollCalendar/index/index.component';
import { CreatePayrollRunComponent } from './components/PayrollRun/create/create.component';
import { EditPayrollRunComponent } from './components/PayrollRun/edit/edit.component';
import { IndexPayrollRunComponent } from './components/PayrollRun/index/index.component';
import { CreatePayrollItemComponent } from './components/PayrollItem/create/create.component';
import { EditPayrollItemComponent } from './components/PayrollItem/edit/edit.component';
import { IndexPayrollItemComponent } from './components/PayrollItem/index/index.component';
import { CreateTaxWithholdingComponent } from './components/TaxWithholding/create/create.component';
import { EditTaxWithholdingComponent } from './components/TaxWithholding/edit/edit.component';
import { IndexTaxWithholdingComponent } from './components/TaxWithholding/index/index.component';
import { CreatePaymentMethodComponent } from './components/PaymentMethod/create/create.component';
import { EditPaymentMethodComponent } from './components/PaymentMethod/edit/edit.component';
import { IndexPaymentMethodComponent } from './components/PaymentMethod/index/index.component';
import { CreateTimesheetComponent } from './components/Timesheet/create/create.component';
import { EditTimesheetComponent } from './components/Timesheet/edit/edit.component';
import { IndexTimesheetComponent } from './components/Timesheet/index/index.component';
import { CreateTimeEntryComponent } from './components/TimeEntry/create/create.component';
import { EditTimeEntryComponent } from './components/TimeEntry/edit/edit.component';
import { IndexTimeEntryComponent } from './components/TimeEntry/index/index.component';
import { CreateApprovalComponent } from './components/Approval/create/create.component';
import { EditApprovalComponent } from './components/Approval/edit/edit.component';
import { IndexApprovalComponent } from './components/Approval/index/index.component';
import { CreateLeavePolicyComponent } from './components/LeavePolicy/create/create.component';
import { EditLeavePolicyComponent } from './components/LeavePolicy/edit/edit.component';
import { IndexLeavePolicyComponent } from './components/LeavePolicy/index/index.component';
import { CreateLeaveRequestComponent } from './components/LeaveRequest/create/create.component';
import { EditLeaveRequestComponent } from './components/LeaveRequest/edit/edit.component';
import { IndexLeaveRequestComponent } from './components/LeaveRequest/index/index.component';
import { CreatePerformanceCycleComponent } from './components/PerformanceCycle/create/create.component';
import { EditPerformanceCycleComponent } from './components/PerformanceCycle/edit/edit.component';
import { IndexPerformanceCycleComponent } from './components/PerformanceCycle/index/index.component';
import { CreateGoalComponent } from './components/Goal/create/create.component';
import { EditGoalComponent } from './components/Goal/edit/edit.component';
import { IndexGoalComponent } from './components/Goal/index/index.component';
import { CreatePerformanceReviewComponent } from './components/PerformanceReview/create/create.component';
import { EditPerformanceReviewComponent } from './components/PerformanceReview/edit/edit.component';
import { IndexPerformanceReviewComponent } from './components/PerformanceReview/index/index.component';
import { CreateCompetencyRatingComponent } from './components/CompetencyRating/create/create.component';
import { EditCompetencyRatingComponent } from './components/CompetencyRating/edit/edit.component';
import { IndexCompetencyRatingComponent } from './components/CompetencyRating/index/index.component';
import { CreateTrainingCourseComponent } from './components/TrainingCourse/create/create.component';
import { EditTrainingCourseComponent } from './components/TrainingCourse/edit/edit.component';
import { IndexTrainingCourseComponent } from './components/TrainingCourse/index/index.component';
import { CreateTrainingEnrollmentComponent } from './components/TrainingEnrollment/create/create.component';
import { EditTrainingEnrollmentComponent } from './components/TrainingEnrollment/edit/edit.component';
import { IndexTrainingEnrollmentComponent } from './components/TrainingEnrollment/index/index.component';
import { CreateCertificationComponent } from './components/Certification/create/create.component';
import { EditCertificationComponent } from './components/Certification/edit/edit.component';
import { IndexCertificationComponent } from './components/Certification/index/index.component';
import { CreateJobRequisitionComponent } from './components/JobRequisition/create/create.component';
import { EditJobRequisitionComponent } from './components/JobRequisition/edit/edit.component';
import { IndexJobRequisitionComponent } from './components/JobRequisition/index/index.component';
import { CreateCandidateComponent } from './components/Candidate/create/create.component';
import { EditCandidateComponent } from './components/Candidate/edit/edit.component';
import { IndexCandidateComponent } from './components/Candidate/index/index.component';
import { CreateJobApplicationComponent } from './components/JobApplication/create/create.component';
import { EditJobApplicationComponent } from './components/JobApplication/edit/edit.component';
import { IndexJobApplicationComponent } from './components/JobApplication/index/index.component';
import { CreateInterviewComponent } from './components/Interview/create/create.component';
import { EditInterviewComponent } from './components/Interview/edit/edit.component';
import { IndexInterviewComponent } from './components/Interview/index/index.component';
import { CreateScreeningComponent } from './components/Screening/create/create.component';
import { EditScreeningComponent } from './components/Screening/edit/edit.component';
import { IndexScreeningComponent } from './components/Screening/index/index.component';
import { CreateOfferComponent } from './components/Offer/create/create.component';
import { EditOfferComponent } from './components/Offer/edit/edit.component';
import { IndexOfferComponent } from './components/Offer/index/index.component';
import { CreateOnboardingTaskComponent } from './components/OnboardingTask/create/create.component';
import { EditOnboardingTaskComponent } from './components/OnboardingTask/edit/edit.component';
import { IndexOnboardingTaskComponent } from './components/OnboardingTask/index/index.component';
import { CreateBackgroundCheckComponent } from './components/BackgroundCheck/create/create.component';
import { EditBackgroundCheckComponent } from './components/BackgroundCheck/edit/edit.component';
import { IndexBackgroundCheckComponent } from './components/BackgroundCheck/index/index.component';
import { CreateDocumentComponent } from './components/Document/create/create.component';
import { EditDocumentComponent } from './components/Document/edit/edit.component';
import { IndexDocumentComponent } from './components/Document/index/index.component';
import { CreatePolicyComponent } from './components/Policy/create/create.component';
import { EditPolicyComponent } from './components/Policy/edit/edit.component';
import { IndexPolicyComponent } from './components/Policy/index/index.component';
import { CreatePolicyAcknowledgementComponent } from './components/PolicyAcknowledgement/create/create.component';
import { EditPolicyAcknowledgementComponent } from './components/PolicyAcknowledgement/edit/edit.component';
import { IndexPolicyAcknowledgementComponent } from './components/PolicyAcknowledgement/index/index.component';
import { CreateTerminationComponent } from './components/Termination/create/create.component';
import { EditTerminationComponent } from './components/Termination/edit/edit.component';
import { IndexTerminationComponent } from './components/Termination/index/index.component';
import { CreateWorkAuthorizationComponent } from './components/WorkAuthorization/create/create.component';
import { EditWorkAuthorizationComponent } from './components/WorkAuthorization/edit/edit.component';
import { IndexWorkAuthorizationComponent } from './components/WorkAuthorization/index/index.component';
import { CreateBankAccountComponent } from './components/BankAccount/create/create.component';
import { EditBankAccountComponent } from './components/BankAccount/edit/edit.component';
import { IndexBankAccountComponent } from './components/BankAccount/index/index.component';

export const routes: Routes = [

        {
        path: 'createOrganization',
        component: CreateOrganizationComponent
},
{
    path: 'editOrganization/:id',
        component: EditOrganizationComponent
},
{
    path: 'indexOrganization',
        component: IndexOrganizationComponent
},
    {
        path: 'createDepartment',
        component: CreateDepartmentComponent
},
{
    path: 'editDepartment/:id',
        component: EditDepartmentComponent
},
{
    path: 'indexDepartment',
        component: IndexDepartmentComponent
},
    {
        path: 'createLocation',
        component: CreateLocationComponent
},
{
    path: 'editLocation/:id',
        component: EditLocationComponent
},
{
    path: 'indexLocation',
        component: IndexLocationComponent
},
    {
        path: 'createCostCenter',
        component: CreateCostCenterComponent
},
{
    path: 'editCostCenter/:id',
        component: EditCostCenterComponent
},
{
    path: 'indexCostCenter',
        component: IndexCostCenterComponent
},
    {
        path: 'createJobFamily',
        component: CreateJobFamilyComponent
},
{
    path: 'editJobFamily/:id',
        component: EditJobFamilyComponent
},
{
    path: 'indexJobFamily',
        component: IndexJobFamilyComponent
},
    {
        path: 'createJobProfile',
        component: CreateJobProfileComponent
},
{
    path: 'editJobProfile/:id',
        component: EditJobProfileComponent
},
{
    path: 'indexJobProfile',
        component: IndexJobProfileComponent
},
    {
        path: 'createCompetency',
        component: CreateCompetencyComponent
},
{
    path: 'editCompetency/:id',
        component: EditCompetencyComponent
},
{
    path: 'indexCompetency',
        component: IndexCompetencyComponent
},
    {
        path: 'createPosition',
        component: CreatePositionComponent
},
{
    path: 'editPosition/:id',
        component: EditPositionComponent
},
{
    path: 'indexPosition',
        component: IndexPositionComponent
},
    {
        path: 'createEmployee',
        component: CreateEmployeeComponent
},
{
    path: 'editEmployee/:id',
        component: EditEmployeeComponent
},
{
    path: 'indexEmployee',
        component: IndexEmployeeComponent
},
    {
        path: 'createEmploymentAssignment',
        component: CreateEmploymentAssignmentComponent
},
{
    path: 'editEmploymentAssignment/:id',
        component: EditEmploymentAssignmentComponent
},
{
    path: 'indexEmploymentAssignment',
        component: IndexEmploymentAssignmentComponent
},
    {
        path: 'createEmploymentContract',
        component: CreateEmploymentContractComponent
},
{
    path: 'editEmploymentContract/:id',
        component: EditEmploymentContractComponent
},
{
    path: 'indexEmploymentContract',
        component: IndexEmploymentContractComponent
},
    {
        path: 'createWorkSchedule',
        component: CreateWorkScheduleComponent
},
{
    path: 'editWorkSchedule/:id',
        component: EditWorkScheduleComponent
},
{
    path: 'indexWorkSchedule',
        component: IndexWorkScheduleComponent
},
    {
        path: 'createWorkShift',
        component: CreateWorkShiftComponent
},
{
    path: 'editWorkShift/:id',
        component: EditWorkShiftComponent
},
{
    path: 'indexWorkShift',
        component: IndexWorkShiftComponent
},
    {
        path: 'createScheduleException',
        component: CreateScheduleExceptionComponent
},
{
    path: 'editScheduleException/:id',
        component: EditScheduleExceptionComponent
},
{
    path: 'indexScheduleException',
        component: IndexScheduleExceptionComponent
},
    {
        path: 'createCompensationPackage',
        component: CreateCompensationPackageComponent
},
{
    path: 'editCompensationPackage/:id',
        component: EditCompensationPackageComponent
},
{
    path: 'indexCompensationPackage',
        component: IndexCompensationPackageComponent
},
    {
        path: 'createSalaryComponent',
        component: CreateSalaryComponentComponent
},
{
    path: 'editSalaryComponent/:id',
        component: EditSalaryComponentComponent
},
{
    path: 'indexSalaryComponent',
        component: IndexSalaryComponentComponent
},
    {
        path: 'createBonusPlan',
        component: CreateBonusPlanComponent
},
{
    path: 'editBonusPlan/:id',
        component: EditBonusPlanComponent
},
{
    path: 'indexBonusPlan',
        component: IndexBonusPlanComponent
},
    {
        path: 'createEquityGrant',
        component: CreateEquityGrantComponent
},
{
    path: 'editEquityGrant/:id',
        component: EditEquityGrantComponent
},
{
    path: 'indexEquityGrant',
        component: IndexEquityGrantComponent
},
    {
        path: 'createBenefitPlan',
        component: CreateBenefitPlanComponent
},
{
    path: 'editBenefitPlan/:id',
        component: EditBenefitPlanComponent
},
{
    path: 'indexBenefitPlan',
        component: IndexBenefitPlanComponent
},
    {
        path: 'createBenefitEnrollment',
        component: CreateBenefitEnrollmentComponent
},
{
    path: 'editBenefitEnrollment/:id',
        component: EditBenefitEnrollmentComponent
},
{
    path: 'indexBenefitEnrollment',
        component: IndexBenefitEnrollmentComponent
},
    {
        path: 'createDependent',
        component: CreateDependentComponent
},
{
    path: 'editDependent/:id',
        component: EditDependentComponent
},
{
    path: 'indexDependent',
        component: IndexDependentComponent
},
    {
        path: 'createPayrollCalendar',
        component: CreatePayrollCalendarComponent
},
{
    path: 'editPayrollCalendar/:id',
        component: EditPayrollCalendarComponent
},
{
    path: 'indexPayrollCalendar',
        component: IndexPayrollCalendarComponent
},
    {
        path: 'createPayrollRun',
        component: CreatePayrollRunComponent
},
{
    path: 'editPayrollRun/:id',
        component: EditPayrollRunComponent
},
{
    path: 'indexPayrollRun',
        component: IndexPayrollRunComponent
},
    {
        path: 'createPayrollItem',
        component: CreatePayrollItemComponent
},
{
    path: 'editPayrollItem/:id',
        component: EditPayrollItemComponent
},
{
    path: 'indexPayrollItem',
        component: IndexPayrollItemComponent
},
    {
        path: 'createTaxWithholding',
        component: CreateTaxWithholdingComponent
},
{
    path: 'editTaxWithholding/:id',
        component: EditTaxWithholdingComponent
},
{
    path: 'indexTaxWithholding',
        component: IndexTaxWithholdingComponent
},
    {
        path: 'createPaymentMethod',
        component: CreatePaymentMethodComponent
},
{
    path: 'editPaymentMethod/:id',
        component: EditPaymentMethodComponent
},
{
    path: 'indexPaymentMethod',
        component: IndexPaymentMethodComponent
},
    {
        path: 'createTimesheet',
        component: CreateTimesheetComponent
},
{
    path: 'editTimesheet/:id',
        component: EditTimesheetComponent
},
{
    path: 'indexTimesheet',
        component: IndexTimesheetComponent
},
    {
        path: 'createTimeEntry',
        component: CreateTimeEntryComponent
},
{
    path: 'editTimeEntry/:id',
        component: EditTimeEntryComponent
},
{
    path: 'indexTimeEntry',
        component: IndexTimeEntryComponent
},
    {
        path: 'createApproval',
        component: CreateApprovalComponent
},
{
    path: 'editApproval/:id',
        component: EditApprovalComponent
},
{
    path: 'indexApproval',
        component: IndexApprovalComponent
},
    {
        path: 'createLeavePolicy',
        component: CreateLeavePolicyComponent
},
{
    path: 'editLeavePolicy/:id',
        component: EditLeavePolicyComponent
},
{
    path: 'indexLeavePolicy',
        component: IndexLeavePolicyComponent
},
    {
        path: 'createLeaveRequest',
        component: CreateLeaveRequestComponent
},
{
    path: 'editLeaveRequest/:id',
        component: EditLeaveRequestComponent
},
{
    path: 'indexLeaveRequest',
        component: IndexLeaveRequestComponent
},
    {
        path: 'createPerformanceCycle',
        component: CreatePerformanceCycleComponent
},
{
    path: 'editPerformanceCycle/:id',
        component: EditPerformanceCycleComponent
},
{
    path: 'indexPerformanceCycle',
        component: IndexPerformanceCycleComponent
},
    {
        path: 'createGoal',
        component: CreateGoalComponent
},
{
    path: 'editGoal/:id',
        component: EditGoalComponent
},
{
    path: 'indexGoal',
        component: IndexGoalComponent
},
    {
        path: 'createPerformanceReview',
        component: CreatePerformanceReviewComponent
},
{
    path: 'editPerformanceReview/:id',
        component: EditPerformanceReviewComponent
},
{
    path: 'indexPerformanceReview',
        component: IndexPerformanceReviewComponent
},
    {
        path: 'createCompetencyRating',
        component: CreateCompetencyRatingComponent
},
{
    path: 'editCompetencyRating/:id',
        component: EditCompetencyRatingComponent
},
{
    path: 'indexCompetencyRating',
        component: IndexCompetencyRatingComponent
},
    {
        path: 'createTrainingCourse',
        component: CreateTrainingCourseComponent
},
{
    path: 'editTrainingCourse/:id',
        component: EditTrainingCourseComponent
},
{
    path: 'indexTrainingCourse',
        component: IndexTrainingCourseComponent
},
    {
        path: 'createTrainingEnrollment',
        component: CreateTrainingEnrollmentComponent
},
{
    path: 'editTrainingEnrollment/:id',
        component: EditTrainingEnrollmentComponent
},
{
    path: 'indexTrainingEnrollment',
        component: IndexTrainingEnrollmentComponent
},
    {
        path: 'createCertification',
        component: CreateCertificationComponent
},
{
    path: 'editCertification/:id',
        component: EditCertificationComponent
},
{
    path: 'indexCertification',
        component: IndexCertificationComponent
},
    {
        path: 'createJobRequisition',
        component: CreateJobRequisitionComponent
},
{
    path: 'editJobRequisition/:id',
        component: EditJobRequisitionComponent
},
{
    path: 'indexJobRequisition',
        component: IndexJobRequisitionComponent
},
    {
        path: 'createCandidate',
        component: CreateCandidateComponent
},
{
    path: 'editCandidate/:id',
        component: EditCandidateComponent
},
{
    path: 'indexCandidate',
        component: IndexCandidateComponent
},
    {
        path: 'createJobApplication',
        component: CreateJobApplicationComponent
},
{
    path: 'editJobApplication/:id',
        component: EditJobApplicationComponent
},
{
    path: 'indexJobApplication',
        component: IndexJobApplicationComponent
},
    {
        path: 'createInterview',
        component: CreateInterviewComponent
},
{
    path: 'editInterview/:id',
        component: EditInterviewComponent
},
{
    path: 'indexInterview',
        component: IndexInterviewComponent
},
    {
        path: 'createScreening',
        component: CreateScreeningComponent
},
{
    path: 'editScreening/:id',
        component: EditScreeningComponent
},
{
    path: 'indexScreening',
        component: IndexScreeningComponent
},
    {
        path: 'createOffer',
        component: CreateOfferComponent
},
{
    path: 'editOffer/:id',
        component: EditOfferComponent
},
{
    path: 'indexOffer',
        component: IndexOfferComponent
},
    {
        path: 'createOnboardingTask',
        component: CreateOnboardingTaskComponent
},
{
    path: 'editOnboardingTask/:id',
        component: EditOnboardingTaskComponent
},
{
    path: 'indexOnboardingTask',
        component: IndexOnboardingTaskComponent
},
    {
        path: 'createBackgroundCheck',
        component: CreateBackgroundCheckComponent
},
{
    path: 'editBackgroundCheck/:id',
        component: EditBackgroundCheckComponent
},
{
    path: 'indexBackgroundCheck',
        component: IndexBackgroundCheckComponent
},
    {
        path: 'createDocument',
        component: CreateDocumentComponent
},
{
    path: 'editDocument/:id',
        component: EditDocumentComponent
},
{
    path: 'indexDocument',
        component: IndexDocumentComponent
},
    {
        path: 'createPolicy',
        component: CreatePolicyComponent
},
{
    path: 'editPolicy/:id',
        component: EditPolicyComponent
},
{
    path: 'indexPolicy',
        component: IndexPolicyComponent
},
    {
        path: 'createPolicyAcknowledgement',
        component: CreatePolicyAcknowledgementComponent
},
{
    path: 'editPolicyAcknowledgement/:id',
        component: EditPolicyAcknowledgementComponent
},
{
    path: 'indexPolicyAcknowledgement',
        component: IndexPolicyAcknowledgementComponent
},
    {
        path: 'createTermination',
        component: CreateTerminationComponent
},
{
    path: 'editTermination/:id',
        component: EditTerminationComponent
},
{
    path: 'indexTermination',
        component: IndexTerminationComponent
},
    {
        path: 'createWorkAuthorization',
        component: CreateWorkAuthorizationComponent
},
{
    path: 'editWorkAuthorization/:id',
        component: EditWorkAuthorizationComponent
},
{
    path: 'indexWorkAuthorization',
        component: IndexWorkAuthorizationComponent
},
    {
        path: 'createBankAccount',
        component: CreateBankAccountComponent
},
{
    path: 'editBankAccount/:id',
        component: EditBankAccountComponent
},
{
    path: 'indexBankAccount',
        component: IndexBankAccountComponent
}
];