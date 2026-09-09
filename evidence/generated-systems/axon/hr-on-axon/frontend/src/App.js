import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListOrganizationComponent from './components/ListOrganizationComponent';
import CreateOrganizationComponent from './components/CreateOrganizationComponent';
import ViewOrganizationComponent from './components/ViewOrganizationComponent';
import ListDepartmentComponent from './components/ListDepartmentComponent';
import CreateDepartmentComponent from './components/CreateDepartmentComponent';
import ViewDepartmentComponent from './components/ViewDepartmentComponent';
import ListLocationComponent from './components/ListLocationComponent';
import CreateLocationComponent from './components/CreateLocationComponent';
import ViewLocationComponent from './components/ViewLocationComponent';
import ListCostCenterComponent from './components/ListCostCenterComponent';
import CreateCostCenterComponent from './components/CreateCostCenterComponent';
import ViewCostCenterComponent from './components/ViewCostCenterComponent';
import ListJobFamilyComponent from './components/ListJobFamilyComponent';
import CreateJobFamilyComponent from './components/CreateJobFamilyComponent';
import ViewJobFamilyComponent from './components/ViewJobFamilyComponent';
import ListJobProfileComponent from './components/ListJobProfileComponent';
import CreateJobProfileComponent from './components/CreateJobProfileComponent';
import ViewJobProfileComponent from './components/ViewJobProfileComponent';
import ListCompetencyComponent from './components/ListCompetencyComponent';
import CreateCompetencyComponent from './components/CreateCompetencyComponent';
import ViewCompetencyComponent from './components/ViewCompetencyComponent';
import ListPositionComponent from './components/ListPositionComponent';
import CreatePositionComponent from './components/CreatePositionComponent';
import ViewPositionComponent from './components/ViewPositionComponent';
import ListEmployeeComponent from './components/ListEmployeeComponent';
import CreateEmployeeComponent from './components/CreateEmployeeComponent';
import ViewEmployeeComponent from './components/ViewEmployeeComponent';
import ListEmploymentAssignmentComponent from './components/ListEmploymentAssignmentComponent';
import CreateEmploymentAssignmentComponent from './components/CreateEmploymentAssignmentComponent';
import ViewEmploymentAssignmentComponent from './components/ViewEmploymentAssignmentComponent';
import ListEmploymentContractComponent from './components/ListEmploymentContractComponent';
import CreateEmploymentContractComponent from './components/CreateEmploymentContractComponent';
import ViewEmploymentContractComponent from './components/ViewEmploymentContractComponent';
import ListWorkScheduleComponent from './components/ListWorkScheduleComponent';
import CreateWorkScheduleComponent from './components/CreateWorkScheduleComponent';
import ViewWorkScheduleComponent from './components/ViewWorkScheduleComponent';
import ListWorkShiftComponent from './components/ListWorkShiftComponent';
import CreateWorkShiftComponent from './components/CreateWorkShiftComponent';
import ViewWorkShiftComponent from './components/ViewWorkShiftComponent';
import ListScheduleExceptionComponent from './components/ListScheduleExceptionComponent';
import CreateScheduleExceptionComponent from './components/CreateScheduleExceptionComponent';
import ViewScheduleExceptionComponent from './components/ViewScheduleExceptionComponent';
import ListCompensationPackageComponent from './components/ListCompensationPackageComponent';
import CreateCompensationPackageComponent from './components/CreateCompensationPackageComponent';
import ViewCompensationPackageComponent from './components/ViewCompensationPackageComponent';
import ListSalaryComponentComponent from './components/ListSalaryComponentComponent';
import CreateSalaryComponentComponent from './components/CreateSalaryComponentComponent';
import ViewSalaryComponentComponent from './components/ViewSalaryComponentComponent';
import ListBonusPlanComponent from './components/ListBonusPlanComponent';
import CreateBonusPlanComponent from './components/CreateBonusPlanComponent';
import ViewBonusPlanComponent from './components/ViewBonusPlanComponent';
import ListEquityGrantComponent from './components/ListEquityGrantComponent';
import CreateEquityGrantComponent from './components/CreateEquityGrantComponent';
import ViewEquityGrantComponent from './components/ViewEquityGrantComponent';
import ListBenefitPlanComponent from './components/ListBenefitPlanComponent';
import CreateBenefitPlanComponent from './components/CreateBenefitPlanComponent';
import ViewBenefitPlanComponent from './components/ViewBenefitPlanComponent';
import ListBenefitEnrollmentComponent from './components/ListBenefitEnrollmentComponent';
import CreateBenefitEnrollmentComponent from './components/CreateBenefitEnrollmentComponent';
import ViewBenefitEnrollmentComponent from './components/ViewBenefitEnrollmentComponent';
import ListDependentComponent from './components/ListDependentComponent';
import CreateDependentComponent from './components/CreateDependentComponent';
import ViewDependentComponent from './components/ViewDependentComponent';
import ListPayrollCalendarComponent from './components/ListPayrollCalendarComponent';
import CreatePayrollCalendarComponent from './components/CreatePayrollCalendarComponent';
import ViewPayrollCalendarComponent from './components/ViewPayrollCalendarComponent';
import ListPayrollRunComponent from './components/ListPayrollRunComponent';
import CreatePayrollRunComponent from './components/CreatePayrollRunComponent';
import ViewPayrollRunComponent from './components/ViewPayrollRunComponent';
import ListPayrollItemComponent from './components/ListPayrollItemComponent';
import CreatePayrollItemComponent from './components/CreatePayrollItemComponent';
import ViewPayrollItemComponent from './components/ViewPayrollItemComponent';
import ListTaxWithholdingComponent from './components/ListTaxWithholdingComponent';
import CreateTaxWithholdingComponent from './components/CreateTaxWithholdingComponent';
import ViewTaxWithholdingComponent from './components/ViewTaxWithholdingComponent';
import ListPaymentMethodComponent from './components/ListPaymentMethodComponent';
import CreatePaymentMethodComponent from './components/CreatePaymentMethodComponent';
import ViewPaymentMethodComponent from './components/ViewPaymentMethodComponent';
import ListTimesheetComponent from './components/ListTimesheetComponent';
import CreateTimesheetComponent from './components/CreateTimesheetComponent';
import ViewTimesheetComponent from './components/ViewTimesheetComponent';
import ListTimeEntryComponent from './components/ListTimeEntryComponent';
import CreateTimeEntryComponent from './components/CreateTimeEntryComponent';
import ViewTimeEntryComponent from './components/ViewTimeEntryComponent';
import ListApprovalComponent from './components/ListApprovalComponent';
import CreateApprovalComponent from './components/CreateApprovalComponent';
import ViewApprovalComponent from './components/ViewApprovalComponent';
import ListLeavePolicyComponent from './components/ListLeavePolicyComponent';
import CreateLeavePolicyComponent from './components/CreateLeavePolicyComponent';
import ViewLeavePolicyComponent from './components/ViewLeavePolicyComponent';
import ListLeaveRequestComponent from './components/ListLeaveRequestComponent';
import CreateLeaveRequestComponent from './components/CreateLeaveRequestComponent';
import ViewLeaveRequestComponent from './components/ViewLeaveRequestComponent';
import ListPerformanceCycleComponent from './components/ListPerformanceCycleComponent';
import CreatePerformanceCycleComponent from './components/CreatePerformanceCycleComponent';
import ViewPerformanceCycleComponent from './components/ViewPerformanceCycleComponent';
import ListGoalComponent from './components/ListGoalComponent';
import CreateGoalComponent from './components/CreateGoalComponent';
import ViewGoalComponent from './components/ViewGoalComponent';
import ListPerformanceReviewComponent from './components/ListPerformanceReviewComponent';
import CreatePerformanceReviewComponent from './components/CreatePerformanceReviewComponent';
import ViewPerformanceReviewComponent from './components/ViewPerformanceReviewComponent';
import ListCompetencyRatingComponent from './components/ListCompetencyRatingComponent';
import CreateCompetencyRatingComponent from './components/CreateCompetencyRatingComponent';
import ViewCompetencyRatingComponent from './components/ViewCompetencyRatingComponent';
import ListTrainingCourseComponent from './components/ListTrainingCourseComponent';
import CreateTrainingCourseComponent from './components/CreateTrainingCourseComponent';
import ViewTrainingCourseComponent from './components/ViewTrainingCourseComponent';
import ListTrainingEnrollmentComponent from './components/ListTrainingEnrollmentComponent';
import CreateTrainingEnrollmentComponent from './components/CreateTrainingEnrollmentComponent';
import ViewTrainingEnrollmentComponent from './components/ViewTrainingEnrollmentComponent';
import ListCertificationComponent from './components/ListCertificationComponent';
import CreateCertificationComponent from './components/CreateCertificationComponent';
import ViewCertificationComponent from './components/ViewCertificationComponent';
import ListJobRequisitionComponent from './components/ListJobRequisitionComponent';
import CreateJobRequisitionComponent from './components/CreateJobRequisitionComponent';
import ViewJobRequisitionComponent from './components/ViewJobRequisitionComponent';
import ListCandidateComponent from './components/ListCandidateComponent';
import CreateCandidateComponent from './components/CreateCandidateComponent';
import ViewCandidateComponent from './components/ViewCandidateComponent';
import ListJobApplicationComponent from './components/ListJobApplicationComponent';
import CreateJobApplicationComponent from './components/CreateJobApplicationComponent';
import ViewJobApplicationComponent from './components/ViewJobApplicationComponent';
import ListInterviewComponent from './components/ListInterviewComponent';
import CreateInterviewComponent from './components/CreateInterviewComponent';
import ViewInterviewComponent from './components/ViewInterviewComponent';
import ListScreeningComponent from './components/ListScreeningComponent';
import CreateScreeningComponent from './components/CreateScreeningComponent';
import ViewScreeningComponent from './components/ViewScreeningComponent';
import ListOfferComponent from './components/ListOfferComponent';
import CreateOfferComponent from './components/CreateOfferComponent';
import ViewOfferComponent from './components/ViewOfferComponent';
import ListOnboardingTaskComponent from './components/ListOnboardingTaskComponent';
import CreateOnboardingTaskComponent from './components/CreateOnboardingTaskComponent';
import ViewOnboardingTaskComponent from './components/ViewOnboardingTaskComponent';
import ListBackgroundCheckComponent from './components/ListBackgroundCheckComponent';
import CreateBackgroundCheckComponent from './components/CreateBackgroundCheckComponent';
import ViewBackgroundCheckComponent from './components/ViewBackgroundCheckComponent';
import ListDocumentComponent from './components/ListDocumentComponent';
import CreateDocumentComponent from './components/CreateDocumentComponent';
import ViewDocumentComponent from './components/ViewDocumentComponent';
import ListPolicyComponent from './components/ListPolicyComponent';
import CreatePolicyComponent from './components/CreatePolicyComponent';
import ViewPolicyComponent from './components/ViewPolicyComponent';
import ListPolicyAcknowledgementComponent from './components/ListPolicyAcknowledgementComponent';
import CreatePolicyAcknowledgementComponent from './components/CreatePolicyAcknowledgementComponent';
import ViewPolicyAcknowledgementComponent from './components/ViewPolicyAcknowledgementComponent';
import ListTerminationComponent from './components/ListTerminationComponent';
import CreateTerminationComponent from './components/CreateTerminationComponent';
import ViewTerminationComponent from './components/ViewTerminationComponent';
import ListWorkAuthorizationComponent from './components/ListWorkAuthorizationComponent';
import CreateWorkAuthorizationComponent from './components/CreateWorkAuthorizationComponent';
import ViewWorkAuthorizationComponent from './components/ViewWorkAuthorizationComponent';
import ListBankAccountComponent from './components/ListBankAccountComponent';
import CreateBankAccountComponent from './components/CreateBankAccountComponent';
import ViewBankAccountComponent from './components/ViewBankAccountComponent';
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
                            <Route path = "/departments" component = {ListDepartmentComponent}></Route>
                            <Route path = "/add-department/:id" component = {CreateDepartmentComponent}></Route>
                            <Route path = "/view-department/:id" component = {ViewDepartmentComponent}></Route>
                          {/* <Route path = "/update-department/:id" component = {UpdateDepartmentComponent}></Route> */}
                            <Route path = "/locations" component = {ListLocationComponent}></Route>
                            <Route path = "/add-location/:id" component = {CreateLocationComponent}></Route>
                            <Route path = "/view-location/:id" component = {ViewLocationComponent}></Route>
                          {/* <Route path = "/update-location/:id" component = {UpdateLocationComponent}></Route> */}
                            <Route path = "/costCenters" component = {ListCostCenterComponent}></Route>
                            <Route path = "/add-costCenter/:id" component = {CreateCostCenterComponent}></Route>
                            <Route path = "/view-costCenter/:id" component = {ViewCostCenterComponent}></Route>
                          {/* <Route path = "/update-costCenter/:id" component = {UpdateCostCenterComponent}></Route> */}
                            <Route path = "/jobFamilys" component = {ListJobFamilyComponent}></Route>
                            <Route path = "/add-jobFamily/:id" component = {CreateJobFamilyComponent}></Route>
                            <Route path = "/view-jobFamily/:id" component = {ViewJobFamilyComponent}></Route>
                          {/* <Route path = "/update-jobFamily/:id" component = {UpdateJobFamilyComponent}></Route> */}
                            <Route path = "/jobProfiles" component = {ListJobProfileComponent}></Route>
                            <Route path = "/add-jobProfile/:id" component = {CreateJobProfileComponent}></Route>
                            <Route path = "/view-jobProfile/:id" component = {ViewJobProfileComponent}></Route>
                          {/* <Route path = "/update-jobProfile/:id" component = {UpdateJobProfileComponent}></Route> */}
                            <Route path = "/competencys" component = {ListCompetencyComponent}></Route>
                            <Route path = "/add-competency/:id" component = {CreateCompetencyComponent}></Route>
                            <Route path = "/view-competency/:id" component = {ViewCompetencyComponent}></Route>
                          {/* <Route path = "/update-competency/:id" component = {UpdateCompetencyComponent}></Route> */}
                            <Route path = "/positions" component = {ListPositionComponent}></Route>
                            <Route path = "/add-position/:id" component = {CreatePositionComponent}></Route>
                            <Route path = "/view-position/:id" component = {ViewPositionComponent}></Route>
                          {/* <Route path = "/update-position/:id" component = {UpdatePositionComponent}></Route> */}
                            <Route path = "/employees" component = {ListEmployeeComponent}></Route>
                            <Route path = "/add-employee/:id" component = {CreateEmployeeComponent}></Route>
                            <Route path = "/view-employee/:id" component = {ViewEmployeeComponent}></Route>
                          {/* <Route path = "/update-employee/:id" component = {UpdateEmployeeComponent}></Route> */}
                            <Route path = "/employmentAssignments" component = {ListEmploymentAssignmentComponent}></Route>
                            <Route path = "/add-employmentAssignment/:id" component = {CreateEmploymentAssignmentComponent}></Route>
                            <Route path = "/view-employmentAssignment/:id" component = {ViewEmploymentAssignmentComponent}></Route>
                          {/* <Route path = "/update-employmentAssignment/:id" component = {UpdateEmploymentAssignmentComponent}></Route> */}
                            <Route path = "/employmentContracts" component = {ListEmploymentContractComponent}></Route>
                            <Route path = "/add-employmentContract/:id" component = {CreateEmploymentContractComponent}></Route>
                            <Route path = "/view-employmentContract/:id" component = {ViewEmploymentContractComponent}></Route>
                          {/* <Route path = "/update-employmentContract/:id" component = {UpdateEmploymentContractComponent}></Route> */}
                            <Route path = "/workSchedules" component = {ListWorkScheduleComponent}></Route>
                            <Route path = "/add-workSchedule/:id" component = {CreateWorkScheduleComponent}></Route>
                            <Route path = "/view-workSchedule/:id" component = {ViewWorkScheduleComponent}></Route>
                          {/* <Route path = "/update-workSchedule/:id" component = {UpdateWorkScheduleComponent}></Route> */}
                            <Route path = "/workShifts" component = {ListWorkShiftComponent}></Route>
                            <Route path = "/add-workShift/:id" component = {CreateWorkShiftComponent}></Route>
                            <Route path = "/view-workShift/:id" component = {ViewWorkShiftComponent}></Route>
                          {/* <Route path = "/update-workShift/:id" component = {UpdateWorkShiftComponent}></Route> */}
                            <Route path = "/scheduleExceptions" component = {ListScheduleExceptionComponent}></Route>
                            <Route path = "/add-scheduleException/:id" component = {CreateScheduleExceptionComponent}></Route>
                            <Route path = "/view-scheduleException/:id" component = {ViewScheduleExceptionComponent}></Route>
                          {/* <Route path = "/update-scheduleException/:id" component = {UpdateScheduleExceptionComponent}></Route> */}
                            <Route path = "/compensationPackages" component = {ListCompensationPackageComponent}></Route>
                            <Route path = "/add-compensationPackage/:id" component = {CreateCompensationPackageComponent}></Route>
                            <Route path = "/view-compensationPackage/:id" component = {ViewCompensationPackageComponent}></Route>
                          {/* <Route path = "/update-compensationPackage/:id" component = {UpdateCompensationPackageComponent}></Route> */}
                            <Route path = "/salaryComponents" component = {ListSalaryComponentComponent}></Route>
                            <Route path = "/add-salaryComponent/:id" component = {CreateSalaryComponentComponent}></Route>
                            <Route path = "/view-salaryComponent/:id" component = {ViewSalaryComponentComponent}></Route>
                          {/* <Route path = "/update-salaryComponent/:id" component = {UpdateSalaryComponentComponent}></Route> */}
                            <Route path = "/bonusPlans" component = {ListBonusPlanComponent}></Route>
                            <Route path = "/add-bonusPlan/:id" component = {CreateBonusPlanComponent}></Route>
                            <Route path = "/view-bonusPlan/:id" component = {ViewBonusPlanComponent}></Route>
                          {/* <Route path = "/update-bonusPlan/:id" component = {UpdateBonusPlanComponent}></Route> */}
                            <Route path = "/equityGrants" component = {ListEquityGrantComponent}></Route>
                            <Route path = "/add-equityGrant/:id" component = {CreateEquityGrantComponent}></Route>
                            <Route path = "/view-equityGrant/:id" component = {ViewEquityGrantComponent}></Route>
                          {/* <Route path = "/update-equityGrant/:id" component = {UpdateEquityGrantComponent}></Route> */}
                            <Route path = "/benefitPlans" component = {ListBenefitPlanComponent}></Route>
                            <Route path = "/add-benefitPlan/:id" component = {CreateBenefitPlanComponent}></Route>
                            <Route path = "/view-benefitPlan/:id" component = {ViewBenefitPlanComponent}></Route>
                          {/* <Route path = "/update-benefitPlan/:id" component = {UpdateBenefitPlanComponent}></Route> */}
                            <Route path = "/benefitEnrollments" component = {ListBenefitEnrollmentComponent}></Route>
                            <Route path = "/add-benefitEnrollment/:id" component = {CreateBenefitEnrollmentComponent}></Route>
                            <Route path = "/view-benefitEnrollment/:id" component = {ViewBenefitEnrollmentComponent}></Route>
                          {/* <Route path = "/update-benefitEnrollment/:id" component = {UpdateBenefitEnrollmentComponent}></Route> */}
                            <Route path = "/dependents" component = {ListDependentComponent}></Route>
                            <Route path = "/add-dependent/:id" component = {CreateDependentComponent}></Route>
                            <Route path = "/view-dependent/:id" component = {ViewDependentComponent}></Route>
                          {/* <Route path = "/update-dependent/:id" component = {UpdateDependentComponent}></Route> */}
                            <Route path = "/payrollCalendars" component = {ListPayrollCalendarComponent}></Route>
                            <Route path = "/add-payrollCalendar/:id" component = {CreatePayrollCalendarComponent}></Route>
                            <Route path = "/view-payrollCalendar/:id" component = {ViewPayrollCalendarComponent}></Route>
                          {/* <Route path = "/update-payrollCalendar/:id" component = {UpdatePayrollCalendarComponent}></Route> */}
                            <Route path = "/payrollRuns" component = {ListPayrollRunComponent}></Route>
                            <Route path = "/add-payrollRun/:id" component = {CreatePayrollRunComponent}></Route>
                            <Route path = "/view-payrollRun/:id" component = {ViewPayrollRunComponent}></Route>
                          {/* <Route path = "/update-payrollRun/:id" component = {UpdatePayrollRunComponent}></Route> */}
                            <Route path = "/payrollItems" component = {ListPayrollItemComponent}></Route>
                            <Route path = "/add-payrollItem/:id" component = {CreatePayrollItemComponent}></Route>
                            <Route path = "/view-payrollItem/:id" component = {ViewPayrollItemComponent}></Route>
                          {/* <Route path = "/update-payrollItem/:id" component = {UpdatePayrollItemComponent}></Route> */}
                            <Route path = "/taxWithholdings" component = {ListTaxWithholdingComponent}></Route>
                            <Route path = "/add-taxWithholding/:id" component = {CreateTaxWithholdingComponent}></Route>
                            <Route path = "/view-taxWithholding/:id" component = {ViewTaxWithholdingComponent}></Route>
                          {/* <Route path = "/update-taxWithholding/:id" component = {UpdateTaxWithholdingComponent}></Route> */}
                            <Route path = "/paymentMethods" component = {ListPaymentMethodComponent}></Route>
                            <Route path = "/add-paymentMethod/:id" component = {CreatePaymentMethodComponent}></Route>
                            <Route path = "/view-paymentMethod/:id" component = {ViewPaymentMethodComponent}></Route>
                          {/* <Route path = "/update-paymentMethod/:id" component = {UpdatePaymentMethodComponent}></Route> */}
                            <Route path = "/timesheets" component = {ListTimesheetComponent}></Route>
                            <Route path = "/add-timesheet/:id" component = {CreateTimesheetComponent}></Route>
                            <Route path = "/view-timesheet/:id" component = {ViewTimesheetComponent}></Route>
                          {/* <Route path = "/update-timesheet/:id" component = {UpdateTimesheetComponent}></Route> */}
                            <Route path = "/timeEntrys" component = {ListTimeEntryComponent}></Route>
                            <Route path = "/add-timeEntry/:id" component = {CreateTimeEntryComponent}></Route>
                            <Route path = "/view-timeEntry/:id" component = {ViewTimeEntryComponent}></Route>
                          {/* <Route path = "/update-timeEntry/:id" component = {UpdateTimeEntryComponent}></Route> */}
                            <Route path = "/approvals" component = {ListApprovalComponent}></Route>
                            <Route path = "/add-approval/:id" component = {CreateApprovalComponent}></Route>
                            <Route path = "/view-approval/:id" component = {ViewApprovalComponent}></Route>
                          {/* <Route path = "/update-approval/:id" component = {UpdateApprovalComponent}></Route> */}
                            <Route path = "/leavePolicys" component = {ListLeavePolicyComponent}></Route>
                            <Route path = "/add-leavePolicy/:id" component = {CreateLeavePolicyComponent}></Route>
                            <Route path = "/view-leavePolicy/:id" component = {ViewLeavePolicyComponent}></Route>
                          {/* <Route path = "/update-leavePolicy/:id" component = {UpdateLeavePolicyComponent}></Route> */}
                            <Route path = "/leaveRequests" component = {ListLeaveRequestComponent}></Route>
                            <Route path = "/add-leaveRequest/:id" component = {CreateLeaveRequestComponent}></Route>
                            <Route path = "/view-leaveRequest/:id" component = {ViewLeaveRequestComponent}></Route>
                          {/* <Route path = "/update-leaveRequest/:id" component = {UpdateLeaveRequestComponent}></Route> */}
                            <Route path = "/performanceCycles" component = {ListPerformanceCycleComponent}></Route>
                            <Route path = "/add-performanceCycle/:id" component = {CreatePerformanceCycleComponent}></Route>
                            <Route path = "/view-performanceCycle/:id" component = {ViewPerformanceCycleComponent}></Route>
                          {/* <Route path = "/update-performanceCycle/:id" component = {UpdatePerformanceCycleComponent}></Route> */}
                            <Route path = "/goals" component = {ListGoalComponent}></Route>
                            <Route path = "/add-goal/:id" component = {CreateGoalComponent}></Route>
                            <Route path = "/view-goal/:id" component = {ViewGoalComponent}></Route>
                          {/* <Route path = "/update-goal/:id" component = {UpdateGoalComponent}></Route> */}
                            <Route path = "/performanceReviews" component = {ListPerformanceReviewComponent}></Route>
                            <Route path = "/add-performanceReview/:id" component = {CreatePerformanceReviewComponent}></Route>
                            <Route path = "/view-performanceReview/:id" component = {ViewPerformanceReviewComponent}></Route>
                          {/* <Route path = "/update-performanceReview/:id" component = {UpdatePerformanceReviewComponent}></Route> */}
                            <Route path = "/competencyRatings" component = {ListCompetencyRatingComponent}></Route>
                            <Route path = "/add-competencyRating/:id" component = {CreateCompetencyRatingComponent}></Route>
                            <Route path = "/view-competencyRating/:id" component = {ViewCompetencyRatingComponent}></Route>
                          {/* <Route path = "/update-competencyRating/:id" component = {UpdateCompetencyRatingComponent}></Route> */}
                            <Route path = "/trainingCourses" component = {ListTrainingCourseComponent}></Route>
                            <Route path = "/add-trainingCourse/:id" component = {CreateTrainingCourseComponent}></Route>
                            <Route path = "/view-trainingCourse/:id" component = {ViewTrainingCourseComponent}></Route>
                          {/* <Route path = "/update-trainingCourse/:id" component = {UpdateTrainingCourseComponent}></Route> */}
                            <Route path = "/trainingEnrollments" component = {ListTrainingEnrollmentComponent}></Route>
                            <Route path = "/add-trainingEnrollment/:id" component = {CreateTrainingEnrollmentComponent}></Route>
                            <Route path = "/view-trainingEnrollment/:id" component = {ViewTrainingEnrollmentComponent}></Route>
                          {/* <Route path = "/update-trainingEnrollment/:id" component = {UpdateTrainingEnrollmentComponent}></Route> */}
                            <Route path = "/certifications" component = {ListCertificationComponent}></Route>
                            <Route path = "/add-certification/:id" component = {CreateCertificationComponent}></Route>
                            <Route path = "/view-certification/:id" component = {ViewCertificationComponent}></Route>
                          {/* <Route path = "/update-certification/:id" component = {UpdateCertificationComponent}></Route> */}
                            <Route path = "/jobRequisitions" component = {ListJobRequisitionComponent}></Route>
                            <Route path = "/add-jobRequisition/:id" component = {CreateJobRequisitionComponent}></Route>
                            <Route path = "/view-jobRequisition/:id" component = {ViewJobRequisitionComponent}></Route>
                          {/* <Route path = "/update-jobRequisition/:id" component = {UpdateJobRequisitionComponent}></Route> */}
                            <Route path = "/candidates" component = {ListCandidateComponent}></Route>
                            <Route path = "/add-candidate/:id" component = {CreateCandidateComponent}></Route>
                            <Route path = "/view-candidate/:id" component = {ViewCandidateComponent}></Route>
                          {/* <Route path = "/update-candidate/:id" component = {UpdateCandidateComponent}></Route> */}
                            <Route path = "/jobApplications" component = {ListJobApplicationComponent}></Route>
                            <Route path = "/add-jobApplication/:id" component = {CreateJobApplicationComponent}></Route>
                            <Route path = "/view-jobApplication/:id" component = {ViewJobApplicationComponent}></Route>
                          {/* <Route path = "/update-jobApplication/:id" component = {UpdateJobApplicationComponent}></Route> */}
                            <Route path = "/interviews" component = {ListInterviewComponent}></Route>
                            <Route path = "/add-interview/:id" component = {CreateInterviewComponent}></Route>
                            <Route path = "/view-interview/:id" component = {ViewInterviewComponent}></Route>
                          {/* <Route path = "/update-interview/:id" component = {UpdateInterviewComponent}></Route> */}
                            <Route path = "/screenings" component = {ListScreeningComponent}></Route>
                            <Route path = "/add-screening/:id" component = {CreateScreeningComponent}></Route>
                            <Route path = "/view-screening/:id" component = {ViewScreeningComponent}></Route>
                          {/* <Route path = "/update-screening/:id" component = {UpdateScreeningComponent}></Route> */}
                            <Route path = "/offers" component = {ListOfferComponent}></Route>
                            <Route path = "/add-offer/:id" component = {CreateOfferComponent}></Route>
                            <Route path = "/view-offer/:id" component = {ViewOfferComponent}></Route>
                          {/* <Route path = "/update-offer/:id" component = {UpdateOfferComponent}></Route> */}
                            <Route path = "/onboardingTasks" component = {ListOnboardingTaskComponent}></Route>
                            <Route path = "/add-onboardingTask/:id" component = {CreateOnboardingTaskComponent}></Route>
                            <Route path = "/view-onboardingTask/:id" component = {ViewOnboardingTaskComponent}></Route>
                          {/* <Route path = "/update-onboardingTask/:id" component = {UpdateOnboardingTaskComponent}></Route> */}
                            <Route path = "/backgroundChecks" component = {ListBackgroundCheckComponent}></Route>
                            <Route path = "/add-backgroundCheck/:id" component = {CreateBackgroundCheckComponent}></Route>
                            <Route path = "/view-backgroundCheck/:id" component = {ViewBackgroundCheckComponent}></Route>
                          {/* <Route path = "/update-backgroundCheck/:id" component = {UpdateBackgroundCheckComponent}></Route> */}
                            <Route path = "/documents" component = {ListDocumentComponent}></Route>
                            <Route path = "/add-document/:id" component = {CreateDocumentComponent}></Route>
                            <Route path = "/view-document/:id" component = {ViewDocumentComponent}></Route>
                          {/* <Route path = "/update-document/:id" component = {UpdateDocumentComponent}></Route> */}
                            <Route path = "/policys" component = {ListPolicyComponent}></Route>
                            <Route path = "/add-policy/:id" component = {CreatePolicyComponent}></Route>
                            <Route path = "/view-policy/:id" component = {ViewPolicyComponent}></Route>
                          {/* <Route path = "/update-policy/:id" component = {UpdatePolicyComponent}></Route> */}
                            <Route path = "/policyAcknowledgements" component = {ListPolicyAcknowledgementComponent}></Route>
                            <Route path = "/add-policyAcknowledgement/:id" component = {CreatePolicyAcknowledgementComponent}></Route>
                            <Route path = "/view-policyAcknowledgement/:id" component = {ViewPolicyAcknowledgementComponent}></Route>
                          {/* <Route path = "/update-policyAcknowledgement/:id" component = {UpdatePolicyAcknowledgementComponent}></Route> */}
                            <Route path = "/terminations" component = {ListTerminationComponent}></Route>
                            <Route path = "/add-termination/:id" component = {CreateTerminationComponent}></Route>
                            <Route path = "/view-termination/:id" component = {ViewTerminationComponent}></Route>
                          {/* <Route path = "/update-termination/:id" component = {UpdateTerminationComponent}></Route> */}
                            <Route path = "/workAuthorizations" component = {ListWorkAuthorizationComponent}></Route>
                            <Route path = "/add-workAuthorization/:id" component = {CreateWorkAuthorizationComponent}></Route>
                            <Route path = "/view-workAuthorization/:id" component = {ViewWorkAuthorizationComponent}></Route>
                          {/* <Route path = "/update-workAuthorization/:id" component = {UpdateWorkAuthorizationComponent}></Route> */}
                            <Route path = "/bankAccounts" component = {ListBankAccountComponent}></Route>
                            <Route path = "/add-bankAccount/:id" component = {CreateBankAccountComponent}></Route>
                            <Route path = "/view-bankAccount/:id" component = {ViewBankAccountComponent}></Route>
                          {/* <Route path = "/update-bankAccount/:id" component = {UpdateBankAccountComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
