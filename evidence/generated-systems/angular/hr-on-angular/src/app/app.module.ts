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
import {IndexDepartmentComponent} from './components/Department/index/index.component';
import {CreateDepartmentComponent} from './components/Department/create/create.component';
import {EditDepartmentComponent} from './components/Department/edit/edit.component';
import {IndexLocationComponent} from './components/Location/index/index.component';
import {CreateLocationComponent} from './components/Location/create/create.component';
import {EditLocationComponent} from './components/Location/edit/edit.component';
import {IndexCostCenterComponent} from './components/CostCenter/index/index.component';
import {CreateCostCenterComponent} from './components/CostCenter/create/create.component';
import {EditCostCenterComponent} from './components/CostCenter/edit/edit.component';
import {IndexJobFamilyComponent} from './components/JobFamily/index/index.component';
import {CreateJobFamilyComponent} from './components/JobFamily/create/create.component';
import {EditJobFamilyComponent} from './components/JobFamily/edit/edit.component';
import {IndexJobProfileComponent} from './components/JobProfile/index/index.component';
import {CreateJobProfileComponent} from './components/JobProfile/create/create.component';
import {EditJobProfileComponent} from './components/JobProfile/edit/edit.component';
import {IndexCompetencyComponent} from './components/Competency/index/index.component';
import {CreateCompetencyComponent} from './components/Competency/create/create.component';
import {EditCompetencyComponent} from './components/Competency/edit/edit.component';
import {IndexPositionComponent} from './components/Position/index/index.component';
import {CreatePositionComponent} from './components/Position/create/create.component';
import {EditPositionComponent} from './components/Position/edit/edit.component';
import {IndexEmployeeComponent} from './components/Employee/index/index.component';
import {CreateEmployeeComponent} from './components/Employee/create/create.component';
import {EditEmployeeComponent} from './components/Employee/edit/edit.component';
import {IndexEmploymentAssignmentComponent} from './components/EmploymentAssignment/index/index.component';
import {CreateEmploymentAssignmentComponent} from './components/EmploymentAssignment/create/create.component';
import {EditEmploymentAssignmentComponent} from './components/EmploymentAssignment/edit/edit.component';
import {IndexEmploymentContractComponent} from './components/EmploymentContract/index/index.component';
import {CreateEmploymentContractComponent} from './components/EmploymentContract/create/create.component';
import {EditEmploymentContractComponent} from './components/EmploymentContract/edit/edit.component';
import {IndexWorkScheduleComponent} from './components/WorkSchedule/index/index.component';
import {CreateWorkScheduleComponent} from './components/WorkSchedule/create/create.component';
import {EditWorkScheduleComponent} from './components/WorkSchedule/edit/edit.component';
import {IndexWorkShiftComponent} from './components/WorkShift/index/index.component';
import {CreateWorkShiftComponent} from './components/WorkShift/create/create.component';
import {EditWorkShiftComponent} from './components/WorkShift/edit/edit.component';
import {IndexScheduleExceptionComponent} from './components/ScheduleException/index/index.component';
import {CreateScheduleExceptionComponent} from './components/ScheduleException/create/create.component';
import {EditScheduleExceptionComponent} from './components/ScheduleException/edit/edit.component';
import {IndexCompensationPackageComponent} from './components/CompensationPackage/index/index.component';
import {CreateCompensationPackageComponent} from './components/CompensationPackage/create/create.component';
import {EditCompensationPackageComponent} from './components/CompensationPackage/edit/edit.component';
import {IndexSalaryComponentComponent} from './components/SalaryComponent/index/index.component';
import {CreateSalaryComponentComponent} from './components/SalaryComponent/create/create.component';
import {EditSalaryComponentComponent} from './components/SalaryComponent/edit/edit.component';
import {IndexBonusPlanComponent} from './components/BonusPlan/index/index.component';
import {CreateBonusPlanComponent} from './components/BonusPlan/create/create.component';
import {EditBonusPlanComponent} from './components/BonusPlan/edit/edit.component';
import {IndexEquityGrantComponent} from './components/EquityGrant/index/index.component';
import {CreateEquityGrantComponent} from './components/EquityGrant/create/create.component';
import {EditEquityGrantComponent} from './components/EquityGrant/edit/edit.component';
import {IndexBenefitPlanComponent} from './components/BenefitPlan/index/index.component';
import {CreateBenefitPlanComponent} from './components/BenefitPlan/create/create.component';
import {EditBenefitPlanComponent} from './components/BenefitPlan/edit/edit.component';
import {IndexBenefitEnrollmentComponent} from './components/BenefitEnrollment/index/index.component';
import {CreateBenefitEnrollmentComponent} from './components/BenefitEnrollment/create/create.component';
import {EditBenefitEnrollmentComponent} from './components/BenefitEnrollment/edit/edit.component';
import {IndexDependentComponent} from './components/Dependent/index/index.component';
import {CreateDependentComponent} from './components/Dependent/create/create.component';
import {EditDependentComponent} from './components/Dependent/edit/edit.component';
import {IndexPayrollCalendarComponent} from './components/PayrollCalendar/index/index.component';
import {CreatePayrollCalendarComponent} from './components/PayrollCalendar/create/create.component';
import {EditPayrollCalendarComponent} from './components/PayrollCalendar/edit/edit.component';
import {IndexPayrollRunComponent} from './components/PayrollRun/index/index.component';
import {CreatePayrollRunComponent} from './components/PayrollRun/create/create.component';
import {EditPayrollRunComponent} from './components/PayrollRun/edit/edit.component';
import {IndexPayrollItemComponent} from './components/PayrollItem/index/index.component';
import {CreatePayrollItemComponent} from './components/PayrollItem/create/create.component';
import {EditPayrollItemComponent} from './components/PayrollItem/edit/edit.component';
import {IndexTaxWithholdingComponent} from './components/TaxWithholding/index/index.component';
import {CreateTaxWithholdingComponent} from './components/TaxWithholding/create/create.component';
import {EditTaxWithholdingComponent} from './components/TaxWithholding/edit/edit.component';
import {IndexPaymentMethodComponent} from './components/PaymentMethod/index/index.component';
import {CreatePaymentMethodComponent} from './components/PaymentMethod/create/create.component';
import {EditPaymentMethodComponent} from './components/PaymentMethod/edit/edit.component';
import {IndexTimesheetComponent} from './components/Timesheet/index/index.component';
import {CreateTimesheetComponent} from './components/Timesheet/create/create.component';
import {EditTimesheetComponent} from './components/Timesheet/edit/edit.component';
import {IndexTimeEntryComponent} from './components/TimeEntry/index/index.component';
import {CreateTimeEntryComponent} from './components/TimeEntry/create/create.component';
import {EditTimeEntryComponent} from './components/TimeEntry/edit/edit.component';
import {IndexApprovalComponent} from './components/Approval/index/index.component';
import {CreateApprovalComponent} from './components/Approval/create/create.component';
import {EditApprovalComponent} from './components/Approval/edit/edit.component';
import {IndexLeavePolicyComponent} from './components/LeavePolicy/index/index.component';
import {CreateLeavePolicyComponent} from './components/LeavePolicy/create/create.component';
import {EditLeavePolicyComponent} from './components/LeavePolicy/edit/edit.component';
import {IndexLeaveRequestComponent} from './components/LeaveRequest/index/index.component';
import {CreateLeaveRequestComponent} from './components/LeaveRequest/create/create.component';
import {EditLeaveRequestComponent} from './components/LeaveRequest/edit/edit.component';
import {IndexPerformanceCycleComponent} from './components/PerformanceCycle/index/index.component';
import {CreatePerformanceCycleComponent} from './components/PerformanceCycle/create/create.component';
import {EditPerformanceCycleComponent} from './components/PerformanceCycle/edit/edit.component';
import {IndexGoalComponent} from './components/Goal/index/index.component';
import {CreateGoalComponent} from './components/Goal/create/create.component';
import {EditGoalComponent} from './components/Goal/edit/edit.component';
import {IndexPerformanceReviewComponent} from './components/PerformanceReview/index/index.component';
import {CreatePerformanceReviewComponent} from './components/PerformanceReview/create/create.component';
import {EditPerformanceReviewComponent} from './components/PerformanceReview/edit/edit.component';
import {IndexCompetencyRatingComponent} from './components/CompetencyRating/index/index.component';
import {CreateCompetencyRatingComponent} from './components/CompetencyRating/create/create.component';
import {EditCompetencyRatingComponent} from './components/CompetencyRating/edit/edit.component';
import {IndexTrainingCourseComponent} from './components/TrainingCourse/index/index.component';
import {CreateTrainingCourseComponent} from './components/TrainingCourse/create/create.component';
import {EditTrainingCourseComponent} from './components/TrainingCourse/edit/edit.component';
import {IndexTrainingEnrollmentComponent} from './components/TrainingEnrollment/index/index.component';
import {CreateTrainingEnrollmentComponent} from './components/TrainingEnrollment/create/create.component';
import {EditTrainingEnrollmentComponent} from './components/TrainingEnrollment/edit/edit.component';
import {IndexCertificationComponent} from './components/Certification/index/index.component';
import {CreateCertificationComponent} from './components/Certification/create/create.component';
import {EditCertificationComponent} from './components/Certification/edit/edit.component';
import {IndexJobRequisitionComponent} from './components/JobRequisition/index/index.component';
import {CreateJobRequisitionComponent} from './components/JobRequisition/create/create.component';
import {EditJobRequisitionComponent} from './components/JobRequisition/edit/edit.component';
import {IndexCandidateComponent} from './components/Candidate/index/index.component';
import {CreateCandidateComponent} from './components/Candidate/create/create.component';
import {EditCandidateComponent} from './components/Candidate/edit/edit.component';
import {IndexJobApplicationComponent} from './components/JobApplication/index/index.component';
import {CreateJobApplicationComponent} from './components/JobApplication/create/create.component';
import {EditJobApplicationComponent} from './components/JobApplication/edit/edit.component';
import {IndexInterviewComponent} from './components/Interview/index/index.component';
import {CreateInterviewComponent} from './components/Interview/create/create.component';
import {EditInterviewComponent} from './components/Interview/edit/edit.component';
import {IndexScreeningComponent} from './components/Screening/index/index.component';
import {CreateScreeningComponent} from './components/Screening/create/create.component';
import {EditScreeningComponent} from './components/Screening/edit/edit.component';
import {IndexOfferComponent} from './components/Offer/index/index.component';
import {CreateOfferComponent} from './components/Offer/create/create.component';
import {EditOfferComponent} from './components/Offer/edit/edit.component';
import {IndexOnboardingTaskComponent} from './components/OnboardingTask/index/index.component';
import {CreateOnboardingTaskComponent} from './components/OnboardingTask/create/create.component';
import {EditOnboardingTaskComponent} from './components/OnboardingTask/edit/edit.component';
import {IndexBackgroundCheckComponent} from './components/BackgroundCheck/index/index.component';
import {CreateBackgroundCheckComponent} from './components/BackgroundCheck/create/create.component';
import {EditBackgroundCheckComponent} from './components/BackgroundCheck/edit/edit.component';
import {IndexDocumentComponent} from './components/Document/index/index.component';
import {CreateDocumentComponent} from './components/Document/create/create.component';
import {EditDocumentComponent} from './components/Document/edit/edit.component';
import {IndexPolicyComponent} from './components/Policy/index/index.component';
import {CreatePolicyComponent} from './components/Policy/create/create.component';
import {EditPolicyComponent} from './components/Policy/edit/edit.component';
import {IndexPolicyAcknowledgementComponent} from './components/PolicyAcknowledgement/index/index.component';
import {CreatePolicyAcknowledgementComponent} from './components/PolicyAcknowledgement/create/create.component';
import {EditPolicyAcknowledgementComponent} from './components/PolicyAcknowledgement/edit/edit.component';
import {IndexTerminationComponent} from './components/Termination/index/index.component';
import {CreateTerminationComponent} from './components/Termination/create/create.component';
import {EditTerminationComponent} from './components/Termination/edit/edit.component';
import {IndexWorkAuthorizationComponent} from './components/WorkAuthorization/index/index.component';
import {CreateWorkAuthorizationComponent} from './components/WorkAuthorization/create/create.component';
import {EditWorkAuthorizationComponent} from './components/WorkAuthorization/edit/edit.component';
import {IndexBankAccountComponent} from './components/BankAccount/index/index.component';
import {CreateBankAccountComponent} from './components/BankAccount/create/create.component';
import {EditBankAccountComponent} from './components/BankAccount/edit/edit.component';

import * as appRoutes from './routerConfig';

import {OrganizationService} from './services/Organization.service';
import {DepartmentService} from './services/Department.service';
import {LocationService} from './services/Location.service';
import {CostCenterService} from './services/CostCenter.service';
import {JobFamilyService} from './services/JobFamily.service';
import {JobProfileService} from './services/JobProfile.service';
import {CompetencyService} from './services/Competency.service';
import {PositionService} from './services/Position.service';
import {EmployeeService} from './services/Employee.service';
import {EmploymentAssignmentService} from './services/EmploymentAssignment.service';
import {EmploymentContractService} from './services/EmploymentContract.service';
import {WorkScheduleService} from './services/WorkSchedule.service';
import {WorkShiftService} from './services/WorkShift.service';
import {ScheduleExceptionService} from './services/ScheduleException.service';
import {CompensationPackageService} from './services/CompensationPackage.service';
import {SalaryComponentService} from './services/SalaryComponent.service';
import {BonusPlanService} from './services/BonusPlan.service';
import {EquityGrantService} from './services/EquityGrant.service';
import {BenefitPlanService} from './services/BenefitPlan.service';
import {BenefitEnrollmentService} from './services/BenefitEnrollment.service';
import {DependentService} from './services/Dependent.service';
import {PayrollCalendarService} from './services/PayrollCalendar.service';
import {PayrollRunService} from './services/PayrollRun.service';
import {PayrollItemService} from './services/PayrollItem.service';
import {TaxWithholdingService} from './services/TaxWithholding.service';
import {PaymentMethodService} from './services/PaymentMethod.service';
import {TimesheetService} from './services/Timesheet.service';
import {TimeEntryService} from './services/TimeEntry.service';
import {ApprovalService} from './services/Approval.service';
import {LeavePolicyService} from './services/LeavePolicy.service';
import {LeaveRequestService} from './services/LeaveRequest.service';
import {PerformanceCycleService} from './services/PerformanceCycle.service';
import {GoalService} from './services/Goal.service';
import {PerformanceReviewService} from './services/PerformanceReview.service';
import {CompetencyRatingService} from './services/CompetencyRating.service';
import {TrainingCourseService} from './services/TrainingCourse.service';
import {TrainingEnrollmentService} from './services/TrainingEnrollment.service';
import {CertificationService} from './services/Certification.service';
import {JobRequisitionService} from './services/JobRequisition.service';
import {CandidateService} from './services/Candidate.service';
import {JobApplicationService} from './services/JobApplication.service';
import {InterviewService} from './services/Interview.service';
import {ScreeningService} from './services/Screening.service';
import {OfferService} from './services/Offer.service';
import {OnboardingTaskService} from './services/OnboardingTask.service';
import {BackgroundCheckService} from './services/BackgroundCheck.service';
import {DocumentService} from './services/Document.service';
import {PolicyService} from './services/Policy.service';
import {PolicyAcknowledgementService} from './services/PolicyAcknowledgement.service';
import {TerminationService} from './services/Termination.service';
import {WorkAuthorizationService} from './services/WorkAuthorization.service';
import {BankAccountService} from './services/BankAccount.service';

@NgModule({
  declarations: [
    IndexOrganizationComponent,
    CreateOrganizationComponent,
    EditOrganizationComponent,
    IndexDepartmentComponent,
    CreateDepartmentComponent,
    EditDepartmentComponent,
    IndexLocationComponent,
    CreateLocationComponent,
    EditLocationComponent,
    IndexCostCenterComponent,
    CreateCostCenterComponent,
    EditCostCenterComponent,
    IndexJobFamilyComponent,
    CreateJobFamilyComponent,
    EditJobFamilyComponent,
    IndexJobProfileComponent,
    CreateJobProfileComponent,
    EditJobProfileComponent,
    IndexCompetencyComponent,
    CreateCompetencyComponent,
    EditCompetencyComponent,
    IndexPositionComponent,
    CreatePositionComponent,
    EditPositionComponent,
    IndexEmployeeComponent,
    CreateEmployeeComponent,
    EditEmployeeComponent,
    IndexEmploymentAssignmentComponent,
    CreateEmploymentAssignmentComponent,
    EditEmploymentAssignmentComponent,
    IndexEmploymentContractComponent,
    CreateEmploymentContractComponent,
    EditEmploymentContractComponent,
    IndexWorkScheduleComponent,
    CreateWorkScheduleComponent,
    EditWorkScheduleComponent,
    IndexWorkShiftComponent,
    CreateWorkShiftComponent,
    EditWorkShiftComponent,
    IndexScheduleExceptionComponent,
    CreateScheduleExceptionComponent,
    EditScheduleExceptionComponent,
    IndexCompensationPackageComponent,
    CreateCompensationPackageComponent,
    EditCompensationPackageComponent,
    IndexSalaryComponentComponent,
    CreateSalaryComponentComponent,
    EditSalaryComponentComponent,
    IndexBonusPlanComponent,
    CreateBonusPlanComponent,
    EditBonusPlanComponent,
    IndexEquityGrantComponent,
    CreateEquityGrantComponent,
    EditEquityGrantComponent,
    IndexBenefitPlanComponent,
    CreateBenefitPlanComponent,
    EditBenefitPlanComponent,
    IndexBenefitEnrollmentComponent,
    CreateBenefitEnrollmentComponent,
    EditBenefitEnrollmentComponent,
    IndexDependentComponent,
    CreateDependentComponent,
    EditDependentComponent,
    IndexPayrollCalendarComponent,
    CreatePayrollCalendarComponent,
    EditPayrollCalendarComponent,
    IndexPayrollRunComponent,
    CreatePayrollRunComponent,
    EditPayrollRunComponent,
    IndexPayrollItemComponent,
    CreatePayrollItemComponent,
    EditPayrollItemComponent,
    IndexTaxWithholdingComponent,
    CreateTaxWithholdingComponent,
    EditTaxWithholdingComponent,
    IndexPaymentMethodComponent,
    CreatePaymentMethodComponent,
    EditPaymentMethodComponent,
    IndexTimesheetComponent,
    CreateTimesheetComponent,
    EditTimesheetComponent,
    IndexTimeEntryComponent,
    CreateTimeEntryComponent,
    EditTimeEntryComponent,
    IndexApprovalComponent,
    CreateApprovalComponent,
    EditApprovalComponent,
    IndexLeavePolicyComponent,
    CreateLeavePolicyComponent,
    EditLeavePolicyComponent,
    IndexLeaveRequestComponent,
    CreateLeaveRequestComponent,
    EditLeaveRequestComponent,
    IndexPerformanceCycleComponent,
    CreatePerformanceCycleComponent,
    EditPerformanceCycleComponent,
    IndexGoalComponent,
    CreateGoalComponent,
    EditGoalComponent,
    IndexPerformanceReviewComponent,
    CreatePerformanceReviewComponent,
    EditPerformanceReviewComponent,
    IndexCompetencyRatingComponent,
    CreateCompetencyRatingComponent,
    EditCompetencyRatingComponent,
    IndexTrainingCourseComponent,
    CreateTrainingCourseComponent,
    EditTrainingCourseComponent,
    IndexTrainingEnrollmentComponent,
    CreateTrainingEnrollmentComponent,
    EditTrainingEnrollmentComponent,
    IndexCertificationComponent,
    CreateCertificationComponent,
    EditCertificationComponent,
    IndexJobRequisitionComponent,
    CreateJobRequisitionComponent,
    EditJobRequisitionComponent,
    IndexCandidateComponent,
    CreateCandidateComponent,
    EditCandidateComponent,
    IndexJobApplicationComponent,
    CreateJobApplicationComponent,
    EditJobApplicationComponent,
    IndexInterviewComponent,
    CreateInterviewComponent,
    EditInterviewComponent,
    IndexScreeningComponent,
    CreateScreeningComponent,
    EditScreeningComponent,
    IndexOfferComponent,
    CreateOfferComponent,
    EditOfferComponent,
    IndexOnboardingTaskComponent,
    CreateOnboardingTaskComponent,
    EditOnboardingTaskComponent,
    IndexBackgroundCheckComponent,
    CreateBackgroundCheckComponent,
    EditBackgroundCheckComponent,
    IndexDocumentComponent,
    CreateDocumentComponent,
    EditDocumentComponent,
    IndexPolicyComponent,
    CreatePolicyComponent,
    EditPolicyComponent,
    IndexPolicyAcknowledgementComponent,
    CreatePolicyAcknowledgementComponent,
    EditPolicyAcknowledgementComponent,
    IndexTerminationComponent,
    CreateTerminationComponent,
    EditTerminationComponent,
    IndexWorkAuthorizationComponent,
    CreateWorkAuthorizationComponent,
    EditWorkAuthorizationComponent,
    IndexBankAccountComponent,
    CreateBankAccountComponent,
    EditBankAccountComponent,
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
    RouterModule.forRoot(appRoutes.DepartmentRoutes), 
    RouterModule.forRoot(appRoutes.LocationRoutes), 
    RouterModule.forRoot(appRoutes.CostCenterRoutes), 
    RouterModule.forRoot(appRoutes.JobFamilyRoutes), 
    RouterModule.forRoot(appRoutes.JobProfileRoutes), 
    RouterModule.forRoot(appRoutes.CompetencyRoutes), 
    RouterModule.forRoot(appRoutes.PositionRoutes), 
    RouterModule.forRoot(appRoutes.EmployeeRoutes), 
    RouterModule.forRoot(appRoutes.EmploymentAssignmentRoutes), 
    RouterModule.forRoot(appRoutes.EmploymentContractRoutes), 
    RouterModule.forRoot(appRoutes.WorkScheduleRoutes), 
    RouterModule.forRoot(appRoutes.WorkShiftRoutes), 
    RouterModule.forRoot(appRoutes.ScheduleExceptionRoutes), 
    RouterModule.forRoot(appRoutes.CompensationPackageRoutes), 
    RouterModule.forRoot(appRoutes.SalaryComponentRoutes), 
    RouterModule.forRoot(appRoutes.BonusPlanRoutes), 
    RouterModule.forRoot(appRoutes.EquityGrantRoutes), 
    RouterModule.forRoot(appRoutes.BenefitPlanRoutes), 
    RouterModule.forRoot(appRoutes.BenefitEnrollmentRoutes), 
    RouterModule.forRoot(appRoutes.DependentRoutes), 
    RouterModule.forRoot(appRoutes.PayrollCalendarRoutes), 
    RouterModule.forRoot(appRoutes.PayrollRunRoutes), 
    RouterModule.forRoot(appRoutes.PayrollItemRoutes), 
    RouterModule.forRoot(appRoutes.TaxWithholdingRoutes), 
    RouterModule.forRoot(appRoutes.PaymentMethodRoutes), 
    RouterModule.forRoot(appRoutes.TimesheetRoutes), 
    RouterModule.forRoot(appRoutes.TimeEntryRoutes), 
    RouterModule.forRoot(appRoutes.ApprovalRoutes), 
    RouterModule.forRoot(appRoutes.LeavePolicyRoutes), 
    RouterModule.forRoot(appRoutes.LeaveRequestRoutes), 
    RouterModule.forRoot(appRoutes.PerformanceCycleRoutes), 
    RouterModule.forRoot(appRoutes.GoalRoutes), 
    RouterModule.forRoot(appRoutes.PerformanceReviewRoutes), 
    RouterModule.forRoot(appRoutes.CompetencyRatingRoutes), 
    RouterModule.forRoot(appRoutes.TrainingCourseRoutes), 
    RouterModule.forRoot(appRoutes.TrainingEnrollmentRoutes), 
    RouterModule.forRoot(appRoutes.CertificationRoutes), 
    RouterModule.forRoot(appRoutes.JobRequisitionRoutes), 
    RouterModule.forRoot(appRoutes.CandidateRoutes), 
    RouterModule.forRoot(appRoutes.JobApplicationRoutes), 
    RouterModule.forRoot(appRoutes.InterviewRoutes), 
    RouterModule.forRoot(appRoutes.ScreeningRoutes), 
    RouterModule.forRoot(appRoutes.OfferRoutes), 
    RouterModule.forRoot(appRoutes.OnboardingTaskRoutes), 
    RouterModule.forRoot(appRoutes.BackgroundCheckRoutes), 
    RouterModule.forRoot(appRoutes.DocumentRoutes), 
    RouterModule.forRoot(appRoutes.PolicyRoutes), 
    RouterModule.forRoot(appRoutes.PolicyAcknowledgementRoutes), 
    RouterModule.forRoot(appRoutes.TerminationRoutes), 
    RouterModule.forRoot(appRoutes.WorkAuthorizationRoutes), 
    RouterModule.forRoot(appRoutes.BankAccountRoutes), 
  ],
  providers: [OrganizationService,DepartmentService,LocationService,CostCenterService,JobFamilyService,JobProfileService,CompetencyService,PositionService,EmployeeService,EmploymentAssignmentService,EmploymentContractService,WorkScheduleService,WorkShiftService,ScheduleExceptionService,CompensationPackageService,SalaryComponentService,BonusPlanService,EquityGrantService,BenefitPlanService,BenefitEnrollmentService,DependentService,PayrollCalendarService,PayrollRunService,PayrollItemService,TaxWithholdingService,PaymentMethodService,TimesheetService,TimeEntryService,ApprovalService,LeavePolicyService,LeaveRequestService,PerformanceCycleService,GoalService,PerformanceReviewService,CompetencyRatingService,TrainingCourseService,TrainingEnrollmentService,CertificationService,JobRequisitionService,CandidateService,JobApplicationService,InterviewService,ScreeningService,OfferService,OnboardingTaskService,BackgroundCheckService,DocumentService,PolicyService,PolicyAcknowledgementService,TerminationService,WorkAuthorizationService,BankAccountService],
  bootstrap: [AppComponent]
})
export class AppModule { }
