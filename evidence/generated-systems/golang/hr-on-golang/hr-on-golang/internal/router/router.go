package router

import (

    OrganizationController "hr-on-golang/internal/controller"
    DepartmentController "hr-on-golang/internal/controller"
    LocationController "hr-on-golang/internal/controller"
    CostCenterController "hr-on-golang/internal/controller"
    JobFamilyController "hr-on-golang/internal/controller"
    JobProfileController "hr-on-golang/internal/controller"
    CompetencyController "hr-on-golang/internal/controller"
    PositionController "hr-on-golang/internal/controller"
    EmployeeController "hr-on-golang/internal/controller"
    EmploymentAssignmentController "hr-on-golang/internal/controller"
    EmploymentContractController "hr-on-golang/internal/controller"
    WorkScheduleController "hr-on-golang/internal/controller"
    WorkShiftController "hr-on-golang/internal/controller"
    ScheduleExceptionController "hr-on-golang/internal/controller"
    CompensationPackageController "hr-on-golang/internal/controller"
    SalaryComponentController "hr-on-golang/internal/controller"
    BonusPlanController "hr-on-golang/internal/controller"
    EquityGrantController "hr-on-golang/internal/controller"
    BenefitPlanController "hr-on-golang/internal/controller"
    BenefitEnrollmentController "hr-on-golang/internal/controller"
    DependentController "hr-on-golang/internal/controller"
    PayrollCalendarController "hr-on-golang/internal/controller"
    PayrollRunController "hr-on-golang/internal/controller"
    PayrollItemController "hr-on-golang/internal/controller"
    TaxWithholdingController "hr-on-golang/internal/controller"
    PaymentMethodController "hr-on-golang/internal/controller"
    TimesheetController "hr-on-golang/internal/controller"
    TimeEntryController "hr-on-golang/internal/controller"
    ApprovalController "hr-on-golang/internal/controller"
    LeavePolicyController "hr-on-golang/internal/controller"
    LeaveRequestController "hr-on-golang/internal/controller"
    PerformanceCycleController "hr-on-golang/internal/controller"
    GoalController "hr-on-golang/internal/controller"
    PerformanceReviewController "hr-on-golang/internal/controller"
    CompetencyRatingController "hr-on-golang/internal/controller"
    TrainingCourseController "hr-on-golang/internal/controller"
    TrainingEnrollmentController "hr-on-golang/internal/controller"
    CertificationController "hr-on-golang/internal/controller"
    JobRequisitionController "hr-on-golang/internal/controller"
    CandidateController "hr-on-golang/internal/controller"
    JobApplicationController "hr-on-golang/internal/controller"
    InterviewController "hr-on-golang/internal/controller"
    ScreeningController "hr-on-golang/internal/controller"
    OfferController "hr-on-golang/internal/controller"
    OnboardingTaskController "hr-on-golang/internal/controller"
    BackgroundCheckController "hr-on-golang/internal/controller"
    DocumentController "hr-on-golang/internal/controller"
    PolicyController "hr-on-golang/internal/controller"
    PolicyAcknowledgementController "hr-on-golang/internal/controller"
    TerminationController "hr-on-golang/internal/controller"
    WorkAuthorizationController "hr-on-golang/internal/controller"
    BankAccountController "hr-on-golang/internal/controller"
    jsonResponseFormatter "hr-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "hr-on-golang/internal/controller"

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
    router.HandleFunc("/api/AddDepartmentsToOrganization/{parentId}/departmentsId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddDepartmentsToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDepartmentsFromOrganization/{parentId}/departmentsIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveDepartmentsFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLocationsToOrganization/{parentId}/locationsId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddLocationsToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLocationsFromOrganization/{parentId}/locationsIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveLocationsFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddJobFamiliesToOrganization/{parentId}/jobFamiliesId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddJobFamiliesToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveJobFamiliesFromOrganization/{parentId}/jobFamiliesIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveJobFamiliesFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddBenefitPlansToOrganization/{parentId}/benefitPlansId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddBenefitPlansToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBenefitPlansFromOrganization/{parentId}/benefitPlansIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveBenefitPlansFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCostCentersToOrganization/{parentId}/costCentersId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddCostCentersToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCostCentersFromOrganization/{parentId}/costCentersIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemoveCostCentersFromOrganization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPayrollCalendarsToOrganization/{parentId}/payrollCalendarsId", jsonResponseFormatter.FormatToJSON(OrganizationController.AddPayrollCalendarsToOrganization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePayrollCalendarsFromOrganization/{parentId}/payrollCalendarsIds", jsonResponseFormatter.FormatToJSON(OrganizationController.RemovePayrollCalendarsFromOrganization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Department Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Department/{id}", jsonResponseFormatter.FormatToJSON(DepartmentController.GetDepartment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Department", jsonResponseFormatter.FormatToJSON(DepartmentController.GetAllDepartment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDepartment", jsonResponseFormatter.FormatToJSON(DepartmentController.CreateDepartment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Department/{id}", jsonResponseFormatter.FormatToJSON(DepartmentController.UpdateDepartment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDepartment/{id}", jsonResponseFormatter.FormatToJSON(DepartmentController.DeleteDepartment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToDepartment/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(DepartmentController.AssignOrganizationToDepartment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromDepartment/{parentId}", jsonResponseFormatter.FormatToJSON(DepartmentController.UnassignOrganizationFromDepartment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignManagerToDepartment/{parentId}/managerId", jsonResponseFormatter.FormatToJSON(DepartmentController.AssignManagerToDepartment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignManagerFromDepartment/{parentId}", jsonResponseFormatter.FormatToJSON(DepartmentController.UnassignManagerFromDepartment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCostCenterToDepartment/{parentId}/costCenterId", jsonResponseFormatter.FormatToJSON(DepartmentController.AssignCostCenterToDepartment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCostCenterFromDepartment/{parentId}", jsonResponseFormatter.FormatToJSON(DepartmentController.UnassignCostCenterFromDepartment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPositionsToDepartment/{parentId}/positionsId", jsonResponseFormatter.FormatToJSON(DepartmentController.AddPositionsToDepartment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePositionsFromDepartment/{parentId}/positionsIds", jsonResponseFormatter.FormatToJSON(DepartmentController.RemovePositionsFromDepartment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEmployeesToDepartment/{parentId}/employeesId", jsonResponseFormatter.FormatToJSON(DepartmentController.AddEmployeesToDepartment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEmployeesFromDepartment/{parentId}/employeesIds", jsonResponseFormatter.FormatToJSON(DepartmentController.RemoveEmployeesFromDepartment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Location Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Location/{id}", jsonResponseFormatter.FormatToJSON(LocationController.GetLocation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Location", jsonResponseFormatter.FormatToJSON(LocationController.GetAllLocation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLocation", jsonResponseFormatter.FormatToJSON(LocationController.CreateLocation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Location/{id}", jsonResponseFormatter.FormatToJSON(LocationController.UpdateLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLocation/{id}", jsonResponseFormatter.FormatToJSON(LocationController.DeleteLocation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToLocation/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(LocationController.AssignOrganizationToLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromLocation/{parentId}", jsonResponseFormatter.FormatToJSON(LocationController.UnassignOrganizationFromLocation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDepartmentsToLocation/{parentId}/departmentsId", jsonResponseFormatter.FormatToJSON(LocationController.AddDepartmentsToLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDepartmentsFromLocation/{parentId}/departmentsIds", jsonResponseFormatter.FormatToJSON(LocationController.RemoveDepartmentsFromLocation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPositionsToLocation/{parentId}/positionsId", jsonResponseFormatter.FormatToJSON(LocationController.AddPositionsToLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePositionsFromLocation/{parentId}/positionsIds", jsonResponseFormatter.FormatToJSON(LocationController.RemovePositionsFromLocation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEmployeesToLocation/{parentId}/employeesId", jsonResponseFormatter.FormatToJSON(LocationController.AddEmployeesToLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEmployeesFromLocation/{parentId}/employeesIds", jsonResponseFormatter.FormatToJSON(LocationController.RemoveEmployeesFromLocation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CostCenter Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CostCenter/{id}", jsonResponseFormatter.FormatToJSON(CostCenterController.GetCostCenter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CostCenter", jsonResponseFormatter.FormatToJSON(CostCenterController.GetAllCostCenter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCostCenter", jsonResponseFormatter.FormatToJSON(CostCenterController.CreateCostCenter)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CostCenter/{id}", jsonResponseFormatter.FormatToJSON(CostCenterController.UpdateCostCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCostCenter/{id}", jsonResponseFormatter.FormatToJSON(CostCenterController.DeleteCostCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToCostCenter/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(CostCenterController.AssignOrganizationToCostCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromCostCenter/{parentId}", jsonResponseFormatter.FormatToJSON(CostCenterController.UnassignOrganizationFromCostCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDepartmentsToCostCenter/{parentId}/departmentsId", jsonResponseFormatter.FormatToJSON(CostCenterController.AddDepartmentsToCostCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDepartmentsFromCostCenter/{parentId}/departmentsIds", jsonResponseFormatter.FormatToJSON(CostCenterController.RemoveDepartmentsFromCostCenter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPositionsToCostCenter/{parentId}/positionsId", jsonResponseFormatter.FormatToJSON(CostCenterController.AddPositionsToCostCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePositionsFromCostCenter/{parentId}/positionsIds", jsonResponseFormatter.FormatToJSON(CostCenterController.RemovePositionsFromCostCenter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEmployeesToCostCenter/{parentId}/employeesId", jsonResponseFormatter.FormatToJSON(CostCenterController.AddEmployeesToCostCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEmployeesFromCostCenter/{parentId}/employeesIds", jsonResponseFormatter.FormatToJSON(CostCenterController.RemoveEmployeesFromCostCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // JobFamily Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/JobFamily/{id}", jsonResponseFormatter.FormatToJSON(JobFamilyController.GetJobFamily)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/JobFamily", jsonResponseFormatter.FormatToJSON(JobFamilyController.GetAllJobFamily)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewJobFamily", jsonResponseFormatter.FormatToJSON(JobFamilyController.CreateJobFamily)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/JobFamily/{id}", jsonResponseFormatter.FormatToJSON(JobFamilyController.UpdateJobFamily)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteJobFamily/{id}", jsonResponseFormatter.FormatToJSON(JobFamilyController.DeleteJobFamily)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToJobFamily/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(JobFamilyController.AssignOrganizationToJobFamily)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromJobFamily/{parentId}", jsonResponseFormatter.FormatToJSON(JobFamilyController.UnassignOrganizationFromJobFamily)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddJobProfilesToJobFamily/{parentId}/jobProfilesId", jsonResponseFormatter.FormatToJSON(JobFamilyController.AddJobProfilesToJobFamily)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveJobProfilesFromJobFamily/{parentId}/jobProfilesIds", jsonResponseFormatter.FormatToJSON(JobFamilyController.RemoveJobProfilesFromJobFamily)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // JobProfile Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/JobProfile/{id}", jsonResponseFormatter.FormatToJSON(JobProfileController.GetJobProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/JobProfile", jsonResponseFormatter.FormatToJSON(JobProfileController.GetAllJobProfile)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewJobProfile", jsonResponseFormatter.FormatToJSON(JobProfileController.CreateJobProfile)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/JobProfile/{id}", jsonResponseFormatter.FormatToJSON(JobProfileController.UpdateJobProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteJobProfile/{id}", jsonResponseFormatter.FormatToJSON(JobProfileController.DeleteJobProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignJobFamilyToJobProfile/{parentId}/jobFamilyId", jsonResponseFormatter.FormatToJSON(JobProfileController.AssignJobFamilyToJobProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignJobFamilyFromJobProfile/{parentId}", jsonResponseFormatter.FormatToJSON(JobProfileController.UnassignJobFamilyFromJobProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCompetenciesToJobProfile/{parentId}/competenciesId", jsonResponseFormatter.FormatToJSON(JobProfileController.AddCompetenciesToJobProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCompetenciesFromJobProfile/{parentId}/competenciesIds", jsonResponseFormatter.FormatToJSON(JobProfileController.RemoveCompetenciesFromJobProfile)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTrainingRecommendationsToJobProfile/{parentId}/trainingRecommendationsId", jsonResponseFormatter.FormatToJSON(JobProfileController.AddTrainingRecommendationsToJobProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTrainingRecommendationsFromJobProfile/{parentId}/trainingRecommendationsIds", jsonResponseFormatter.FormatToJSON(JobProfileController.RemoveTrainingRecommendationsFromJobProfile)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPositionsToJobProfile/{parentId}/positionsId", jsonResponseFormatter.FormatToJSON(JobProfileController.AddPositionsToJobProfile)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePositionsFromJobProfile/{parentId}/positionsIds", jsonResponseFormatter.FormatToJSON(JobProfileController.RemovePositionsFromJobProfile)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Competency Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Competency/{id}", jsonResponseFormatter.FormatToJSON(CompetencyController.GetCompetency)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Competency", jsonResponseFormatter.FormatToJSON(CompetencyController.GetAllCompetency)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCompetency", jsonResponseFormatter.FormatToJSON(CompetencyController.CreateCompetency)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Competency/{id}", jsonResponseFormatter.FormatToJSON(CompetencyController.UpdateCompetency)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCompetency/{id}", jsonResponseFormatter.FormatToJSON(CompetencyController.DeleteCompetency)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddJobProfilesToCompetency/{parentId}/jobProfilesId", jsonResponseFormatter.FormatToJSON(CompetencyController.AddJobProfilesToCompetency)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveJobProfilesFromCompetency/{parentId}/jobProfilesIds", jsonResponseFormatter.FormatToJSON(CompetencyController.RemoveJobProfilesFromCompetency)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCompetencyRatingsToCompetency/{parentId}/competencyRatingsId", jsonResponseFormatter.FormatToJSON(CompetencyController.AddCompetencyRatingsToCompetency)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCompetencyRatingsFromCompetency/{parentId}/competencyRatingsIds", jsonResponseFormatter.FormatToJSON(CompetencyController.RemoveCompetencyRatingsFromCompetency)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Position Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Position/{id}", jsonResponseFormatter.FormatToJSON(PositionController.GetPosition)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Position", jsonResponseFormatter.FormatToJSON(PositionController.GetAllPosition)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPosition", jsonResponseFormatter.FormatToJSON(PositionController.CreatePosition)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Position/{id}", jsonResponseFormatter.FormatToJSON(PositionController.UpdatePosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePosition/{id}", jsonResponseFormatter.FormatToJSON(PositionController.DeletePosition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignDepartmentToPosition/{parentId}/departmentId", jsonResponseFormatter.FormatToJSON(PositionController.AssignDepartmentToPosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDepartmentFromPosition/{parentId}", jsonResponseFormatter.FormatToJSON(PositionController.UnassignDepartmentFromPosition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignJobProfileToPosition/{parentId}/jobProfileId", jsonResponseFormatter.FormatToJSON(PositionController.AssignJobProfileToPosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignJobProfileFromPosition/{parentId}", jsonResponseFormatter.FormatToJSON(PositionController.UnassignJobProfileFromPosition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCostCenterToPosition/{parentId}/costCenterId", jsonResponseFormatter.FormatToJSON(PositionController.AssignCostCenterToPosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCostCenterFromPosition/{parentId}", jsonResponseFormatter.FormatToJSON(PositionController.UnassignCostCenterFromPosition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToPosition/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(PositionController.AssignLocationToPosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromPosition/{parentId}", jsonResponseFormatter.FormatToJSON(PositionController.UnassignLocationFromPosition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignManagerPositionToPosition/{parentId}/managerPositionId", jsonResponseFormatter.FormatToJSON(PositionController.AssignManagerPositionToPosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignManagerPositionFromPosition/{parentId}", jsonResponseFormatter.FormatToJSON(PositionController.UnassignManagerPositionFromPosition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDirectReportsToPosition/{parentId}/directReportsId", jsonResponseFormatter.FormatToJSON(PositionController.AddDirectReportsToPosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDirectReportsFromPosition/{parentId}/directReportsIds", jsonResponseFormatter.FormatToJSON(PositionController.RemoveDirectReportsFromPosition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAssignmentsToPosition/{parentId}/assignmentsId", jsonResponseFormatter.FormatToJSON(PositionController.AddAssignmentsToPosition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAssignmentsFromPosition/{parentId}/assignmentsIds", jsonResponseFormatter.FormatToJSON(PositionController.RemoveAssignmentsFromPosition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Employee Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Employee/{id}", jsonResponseFormatter.FormatToJSON(EmployeeController.GetEmployee)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Employee", jsonResponseFormatter.FormatToJSON(EmployeeController.GetAllEmployee)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEmployee", jsonResponseFormatter.FormatToJSON(EmployeeController.CreateEmployee)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Employee/{id}", jsonResponseFormatter.FormatToJSON(EmployeeController.UpdateEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEmployee/{id}", jsonResponseFormatter.FormatToJSON(EmployeeController.DeleteEmployee)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignManagerToEmployee/{parentId}/managerId", jsonResponseFormatter.FormatToJSON(EmployeeController.AssignManagerToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignManagerFromEmployee/{parentId}", jsonResponseFormatter.FormatToJSON(EmployeeController.UnassignManagerFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDepartmentToEmployee/{parentId}/departmentId", jsonResponseFormatter.FormatToJSON(EmployeeController.AssignDepartmentToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDepartmentFromEmployee/{parentId}", jsonResponseFormatter.FormatToJSON(EmployeeController.UnassignDepartmentFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPrimaryLocationToEmployee/{parentId}/primaryLocationId", jsonResponseFormatter.FormatToJSON(EmployeeController.AssignPrimaryLocationToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPrimaryLocationFromEmployee/{parentId}", jsonResponseFormatter.FormatToJSON(EmployeeController.UnassignPrimaryLocationFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCostCenterToEmployee/{parentId}/costCenterId", jsonResponseFormatter.FormatToJSON(EmployeeController.AssignCostCenterToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCostCenterFromEmployee/{parentId}", jsonResponseFormatter.FormatToJSON(EmployeeController.UnassignCostCenterFromEmployee)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDirectReportsToEmployee/{parentId}/directReportsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddDirectReportsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDirectReportsFromEmployee/{parentId}/directReportsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemoveDirectReportsFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEmploymentAssignmentsToEmployee/{parentId}/employmentAssignmentsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddEmploymentAssignmentsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEmploymentAssignmentsFromEmployee/{parentId}/employmentAssignmentsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemoveEmploymentAssignmentsFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddContractsToEmployee/{parentId}/contractsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddContractsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContractsFromEmployee/{parentId}/contractsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemoveContractsFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddBenefitEnrollmentsToEmployee/{parentId}/benefitEnrollmentsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddBenefitEnrollmentsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBenefitEnrollmentsFromEmployee/{parentId}/benefitEnrollmentsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemoveBenefitEnrollmentsFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTimesheetsToEmployee/{parentId}/timesheetsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddTimesheetsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTimesheetsFromEmployee/{parentId}/timesheetsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemoveTimesheetsFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLeaveRequestsToEmployee/{parentId}/leaveRequestsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddLeaveRequestsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLeaveRequestsFromEmployee/{parentId}/leaveRequestsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemoveLeaveRequestsFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPerformanceReviewsToEmployee/{parentId}/performanceReviewsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddPerformanceReviewsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePerformanceReviewsFromEmployee/{parentId}/performanceReviewsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemovePerformanceReviewsFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTrainingEnrollmentsToEmployee/{parentId}/trainingEnrollmentsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddTrainingEnrollmentsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTrainingEnrollmentsFromEmployee/{parentId}/trainingEnrollmentsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemoveTrainingEnrollmentsFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddWorkAuthorizationsToEmployee/{parentId}/workAuthorizationsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddWorkAuthorizationsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWorkAuthorizationsFromEmployee/{parentId}/workAuthorizationsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemoveWorkAuthorizationsFromEmployee)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // EmploymentAssignment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/EmploymentAssignment/{id}", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.GetEmploymentAssignment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/EmploymentAssignment", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.GetAllEmploymentAssignment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEmploymentAssignment", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.CreateEmploymentAssignment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/EmploymentAssignment/{id}", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.UpdateEmploymentAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEmploymentAssignment/{id}", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.DeleteEmploymentAssignment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToEmploymentAssignment/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.AssignEmployeeToEmploymentAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromEmploymentAssignment/{parentId}", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.UnassignEmployeeFromEmploymentAssignment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPositionToEmploymentAssignment/{parentId}/positionId", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.AssignPositionToEmploymentAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPositionFromEmploymentAssignment/{parentId}", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.UnassignPositionFromEmploymentAssignment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSupervisorToEmploymentAssignment/{parentId}/supervisorId", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.AssignSupervisorToEmploymentAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSupervisorFromEmploymentAssignment/{parentId}", jsonResponseFormatter.FormatToJSON(EmploymentAssignmentController.UnassignSupervisorFromEmploymentAssignment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // EmploymentContract Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/EmploymentContract/{id}", jsonResponseFormatter.FormatToJSON(EmploymentContractController.GetEmploymentContract)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/EmploymentContract", jsonResponseFormatter.FormatToJSON(EmploymentContractController.GetAllEmploymentContract)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEmploymentContract", jsonResponseFormatter.FormatToJSON(EmploymentContractController.CreateEmploymentContract)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/EmploymentContract/{id}", jsonResponseFormatter.FormatToJSON(EmploymentContractController.UpdateEmploymentContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEmploymentContract/{id}", jsonResponseFormatter.FormatToJSON(EmploymentContractController.DeleteEmploymentContract)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToEmploymentContract/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(EmploymentContractController.AssignEmployeeToEmploymentContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromEmploymentContract/{parentId}", jsonResponseFormatter.FormatToJSON(EmploymentContractController.UnassignEmployeeFromEmploymentContract)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCompensationPackageToEmploymentContract/{parentId}/compensationPackageId", jsonResponseFormatter.FormatToJSON(EmploymentContractController.AssignCompensationPackageToEmploymentContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCompensationPackageFromEmploymentContract/{parentId}", jsonResponseFormatter.FormatToJSON(EmploymentContractController.UnassignCompensationPackageFromEmploymentContract)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkScheduleToEmploymentContract/{parentId}/workScheduleId", jsonResponseFormatter.FormatToJSON(EmploymentContractController.AssignWorkScheduleToEmploymentContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkScheduleFromEmploymentContract/{parentId}", jsonResponseFormatter.FormatToJSON(EmploymentContractController.UnassignWorkScheduleFromEmploymentContract)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToEmploymentContract/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(EmploymentContractController.AssignLocationToEmploymentContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromEmploymentContract/{parentId}", jsonResponseFormatter.FormatToJSON(EmploymentContractController.UnassignLocationFromEmploymentContract)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPayrollCalendarToEmploymentContract/{parentId}/payrollCalendarId", jsonResponseFormatter.FormatToJSON(EmploymentContractController.AssignPayrollCalendarToEmploymentContract)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPayrollCalendarFromEmploymentContract/{parentId}", jsonResponseFormatter.FormatToJSON(EmploymentContractController.UnassignPayrollCalendarFromEmploymentContract)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // WorkSchedule Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/WorkSchedule/{id}", jsonResponseFormatter.FormatToJSON(WorkScheduleController.GetWorkSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/WorkSchedule", jsonResponseFormatter.FormatToJSON(WorkScheduleController.GetAllWorkSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewWorkSchedule", jsonResponseFormatter.FormatToJSON(WorkScheduleController.CreateWorkSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/WorkSchedule/{id}", jsonResponseFormatter.FormatToJSON(WorkScheduleController.UpdateWorkSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteWorkSchedule/{id}", jsonResponseFormatter.FormatToJSON(WorkScheduleController.DeleteWorkSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddContractsToWorkSchedule/{parentId}/contractsId", jsonResponseFormatter.FormatToJSON(WorkScheduleController.AddContractsToWorkSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveContractsFromWorkSchedule/{parentId}/contractsIds", jsonResponseFormatter.FormatToJSON(WorkScheduleController.RemoveContractsFromWorkSchedule)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddShiftsToWorkSchedule/{parentId}/shiftsId", jsonResponseFormatter.FormatToJSON(WorkScheduleController.AddShiftsToWorkSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveShiftsFromWorkSchedule/{parentId}/shiftsIds", jsonResponseFormatter.FormatToJSON(WorkScheduleController.RemoveShiftsFromWorkSchedule)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddExceptionsToWorkSchedule/{parentId}/exceptionsId", jsonResponseFormatter.FormatToJSON(WorkScheduleController.AddExceptionsToWorkSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveExceptionsFromWorkSchedule/{parentId}/exceptionsIds", jsonResponseFormatter.FormatToJSON(WorkScheduleController.RemoveExceptionsFromWorkSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // WorkShift Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/WorkShift/{id}", jsonResponseFormatter.FormatToJSON(WorkShiftController.GetWorkShift)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/WorkShift", jsonResponseFormatter.FormatToJSON(WorkShiftController.GetAllWorkShift)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewWorkShift", jsonResponseFormatter.FormatToJSON(WorkShiftController.CreateWorkShift)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/WorkShift/{id}", jsonResponseFormatter.FormatToJSON(WorkShiftController.UpdateWorkShift)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteWorkShift/{id}", jsonResponseFormatter.FormatToJSON(WorkShiftController.DeleteWorkShift)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkScheduleToWorkShift/{parentId}/workScheduleId", jsonResponseFormatter.FormatToJSON(WorkShiftController.AssignWorkScheduleToWorkShift)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkScheduleFromWorkShift/{parentId}", jsonResponseFormatter.FormatToJSON(WorkShiftController.UnassignWorkScheduleFromWorkShift)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ScheduleException Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ScheduleException/{id}", jsonResponseFormatter.FormatToJSON(ScheduleExceptionController.GetScheduleException)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ScheduleException", jsonResponseFormatter.FormatToJSON(ScheduleExceptionController.GetAllScheduleException)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewScheduleException", jsonResponseFormatter.FormatToJSON(ScheduleExceptionController.CreateScheduleException)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ScheduleException/{id}", jsonResponseFormatter.FormatToJSON(ScheduleExceptionController.UpdateScheduleException)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteScheduleException/{id}", jsonResponseFormatter.FormatToJSON(ScheduleExceptionController.DeleteScheduleException)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkScheduleToScheduleException/{parentId}/workScheduleId", jsonResponseFormatter.FormatToJSON(ScheduleExceptionController.AssignWorkScheduleToScheduleException)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkScheduleFromScheduleException/{parentId}", jsonResponseFormatter.FormatToJSON(ScheduleExceptionController.UnassignWorkScheduleFromScheduleException)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEmployeeToScheduleException/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(ScheduleExceptionController.AssignEmployeeToScheduleException)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromScheduleException/{parentId}", jsonResponseFormatter.FormatToJSON(ScheduleExceptionController.UnassignEmployeeFromScheduleException)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // CompensationPackage Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CompensationPackage/{id}", jsonResponseFormatter.FormatToJSON(CompensationPackageController.GetCompensationPackage)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CompensationPackage", jsonResponseFormatter.FormatToJSON(CompensationPackageController.GetAllCompensationPackage)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCompensationPackage", jsonResponseFormatter.FormatToJSON(CompensationPackageController.CreateCompensationPackage)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CompensationPackage/{id}", jsonResponseFormatter.FormatToJSON(CompensationPackageController.UpdateCompensationPackage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCompensationPackage/{id}", jsonResponseFormatter.FormatToJSON(CompensationPackageController.DeleteCompensationPackage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignContractToCompensationPackage/{parentId}/contractId", jsonResponseFormatter.FormatToJSON(CompensationPackageController.AssignContractToCompensationPackage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignContractFromCompensationPackage/{parentId}", jsonResponseFormatter.FormatToJSON(CompensationPackageController.UnassignContractFromCompensationPackage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSalaryComponentsToCompensationPackage/{parentId}/salaryComponentsId", jsonResponseFormatter.FormatToJSON(CompensationPackageController.AddSalaryComponentsToCompensationPackage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSalaryComponentsFromCompensationPackage/{parentId}/salaryComponentsIds", jsonResponseFormatter.FormatToJSON(CompensationPackageController.RemoveSalaryComponentsFromCompensationPackage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddBonusPlansToCompensationPackage/{parentId}/bonusPlansId", jsonResponseFormatter.FormatToJSON(CompensationPackageController.AddBonusPlansToCompensationPackage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBonusPlansFromCompensationPackage/{parentId}/bonusPlansIds", jsonResponseFormatter.FormatToJSON(CompensationPackageController.RemoveBonusPlansFromCompensationPackage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEquityGrantsToCompensationPackage/{parentId}/equityGrantsId", jsonResponseFormatter.FormatToJSON(CompensationPackageController.AddEquityGrantsToCompensationPackage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEquityGrantsFromCompensationPackage/{parentId}/equityGrantsIds", jsonResponseFormatter.FormatToJSON(CompensationPackageController.RemoveEquityGrantsFromCompensationPackage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // SalaryComponent Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SalaryComponent/{id}", jsonResponseFormatter.FormatToJSON(SalaryComponentController.GetSalaryComponent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SalaryComponent", jsonResponseFormatter.FormatToJSON(SalaryComponentController.GetAllSalaryComponent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSalaryComponent", jsonResponseFormatter.FormatToJSON(SalaryComponentController.CreateSalaryComponent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SalaryComponent/{id}", jsonResponseFormatter.FormatToJSON(SalaryComponentController.UpdateSalaryComponent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSalaryComponent/{id}", jsonResponseFormatter.FormatToJSON(SalaryComponentController.DeleteSalaryComponent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCompensationPackageToSalaryComponent/{parentId}/compensationPackageId", jsonResponseFormatter.FormatToJSON(SalaryComponentController.AssignCompensationPackageToSalaryComponent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCompensationPackageFromSalaryComponent/{parentId}", jsonResponseFormatter.FormatToJSON(SalaryComponentController.UnassignCompensationPackageFromSalaryComponent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // BonusPlan Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BonusPlan/{id}", jsonResponseFormatter.FormatToJSON(BonusPlanController.GetBonusPlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BonusPlan", jsonResponseFormatter.FormatToJSON(BonusPlanController.GetAllBonusPlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBonusPlan", jsonResponseFormatter.FormatToJSON(BonusPlanController.CreateBonusPlan)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BonusPlan/{id}", jsonResponseFormatter.FormatToJSON(BonusPlanController.UpdateBonusPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBonusPlan/{id}", jsonResponseFormatter.FormatToJSON(BonusPlanController.DeleteBonusPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCompensationPackagesToBonusPlan/{parentId}/compensationPackagesId", jsonResponseFormatter.FormatToJSON(BonusPlanController.AddCompensationPackagesToBonusPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCompensationPackagesFromBonusPlan/{parentId}/compensationPackagesIds", jsonResponseFormatter.FormatToJSON(BonusPlanController.RemoveCompensationPackagesFromBonusPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // EquityGrant Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/EquityGrant/{id}", jsonResponseFormatter.FormatToJSON(EquityGrantController.GetEquityGrant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/EquityGrant", jsonResponseFormatter.FormatToJSON(EquityGrantController.GetAllEquityGrant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEquityGrant", jsonResponseFormatter.FormatToJSON(EquityGrantController.CreateEquityGrant)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/EquityGrant/{id}", jsonResponseFormatter.FormatToJSON(EquityGrantController.UpdateEquityGrant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEquityGrant/{id}", jsonResponseFormatter.FormatToJSON(EquityGrantController.DeleteEquityGrant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCompensationPackageToEquityGrant/{parentId}/compensationPackageId", jsonResponseFormatter.FormatToJSON(EquityGrantController.AssignCompensationPackageToEquityGrant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCompensationPackageFromEquityGrant/{parentId}", jsonResponseFormatter.FormatToJSON(EquityGrantController.UnassignCompensationPackageFromEquityGrant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // BenefitPlan Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BenefitPlan/{id}", jsonResponseFormatter.FormatToJSON(BenefitPlanController.GetBenefitPlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BenefitPlan", jsonResponseFormatter.FormatToJSON(BenefitPlanController.GetAllBenefitPlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBenefitPlan", jsonResponseFormatter.FormatToJSON(BenefitPlanController.CreateBenefitPlan)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BenefitPlan/{id}", jsonResponseFormatter.FormatToJSON(BenefitPlanController.UpdateBenefitPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBenefitPlan/{id}", jsonResponseFormatter.FormatToJSON(BenefitPlanController.DeleteBenefitPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToBenefitPlan/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(BenefitPlanController.AssignOrganizationToBenefitPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromBenefitPlan/{parentId}", jsonResponseFormatter.FormatToJSON(BenefitPlanController.UnassignOrganizationFromBenefitPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddEnrollmentsToBenefitPlan/{parentId}/enrollmentsId", jsonResponseFormatter.FormatToJSON(BenefitPlanController.AddEnrollmentsToBenefitPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEnrollmentsFromBenefitPlan/{parentId}/enrollmentsIds", jsonResponseFormatter.FormatToJSON(BenefitPlanController.RemoveEnrollmentsFromBenefitPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // BenefitEnrollment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BenefitEnrollment/{id}", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.GetBenefitEnrollment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BenefitEnrollment", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.GetAllBenefitEnrollment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBenefitEnrollment", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.CreateBenefitEnrollment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BenefitEnrollment/{id}", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.UpdateBenefitEnrollment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBenefitEnrollment/{id}", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.DeleteBenefitEnrollment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBenefitPlanToBenefitEnrollment/{parentId}/benefitPlanId", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.AssignBenefitPlanToBenefitEnrollment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBenefitPlanFromBenefitEnrollment/{parentId}", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.UnassignBenefitPlanFromBenefitEnrollment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEmployeeToBenefitEnrollment/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.AssignEmployeeToBenefitEnrollment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromBenefitEnrollment/{parentId}", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.UnassignEmployeeFromBenefitEnrollment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDependentsToBenefitEnrollment/{parentId}/dependentsId", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.AddDependentsToBenefitEnrollment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDependentsFromBenefitEnrollment/{parentId}/dependentsIds", jsonResponseFormatter.FormatToJSON(BenefitEnrollmentController.RemoveDependentsFromBenefitEnrollment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Dependent Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Dependent/{id}", jsonResponseFormatter.FormatToJSON(DependentController.GetDependent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Dependent", jsonResponseFormatter.FormatToJSON(DependentController.GetAllDependent)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDependent", jsonResponseFormatter.FormatToJSON(DependentController.CreateDependent)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Dependent/{id}", jsonResponseFormatter.FormatToJSON(DependentController.UpdateDependent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDependent/{id}", jsonResponseFormatter.FormatToJSON(DependentController.DeleteDependent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBenefitEnrollmentToDependent/{parentId}/benefitEnrollmentId", jsonResponseFormatter.FormatToJSON(DependentController.AssignBenefitEnrollmentToDependent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBenefitEnrollmentFromDependent/{parentId}", jsonResponseFormatter.FormatToJSON(DependentController.UnassignBenefitEnrollmentFromDependent)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEmployeeToDependent/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(DependentController.AssignEmployeeToDependent)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromDependent/{parentId}", jsonResponseFormatter.FormatToJSON(DependentController.UnassignEmployeeFromDependent)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // PayrollCalendar Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PayrollCalendar/{id}", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.GetPayrollCalendar)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PayrollCalendar", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.GetAllPayrollCalendar)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPayrollCalendar", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.CreatePayrollCalendar)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PayrollCalendar/{id}", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.UpdatePayrollCalendar)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePayrollCalendar/{id}", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.DeletePayrollCalendar)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToPayrollCalendar/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.AssignOrganizationToPayrollCalendar)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromPayrollCalendar/{parentId}", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.UnassignOrganizationFromPayrollCalendar)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPayrollRunsToPayrollCalendar/{parentId}/payrollRunsId", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.AddPayrollRunsToPayrollCalendar)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePayrollRunsFromPayrollCalendar/{parentId}/payrollRunsIds", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.RemovePayrollRunsFromPayrollCalendar)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEmployeesToPayrollCalendar/{parentId}/employeesId", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.AddEmployeesToPayrollCalendar)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEmployeesFromPayrollCalendar/{parentId}/employeesIds", jsonResponseFormatter.FormatToJSON(PayrollCalendarController.RemoveEmployeesFromPayrollCalendar)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PayrollRun Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PayrollRun/{id}", jsonResponseFormatter.FormatToJSON(PayrollRunController.GetPayrollRun)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PayrollRun", jsonResponseFormatter.FormatToJSON(PayrollRunController.GetAllPayrollRun)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPayrollRun", jsonResponseFormatter.FormatToJSON(PayrollRunController.CreatePayrollRun)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PayrollRun/{id}", jsonResponseFormatter.FormatToJSON(PayrollRunController.UpdatePayrollRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePayrollRun/{id}", jsonResponseFormatter.FormatToJSON(PayrollRunController.DeletePayrollRun)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPayrollCalendarToPayrollRun/{parentId}/payrollCalendarId", jsonResponseFormatter.FormatToJSON(PayrollRunController.AssignPayrollCalendarToPayrollRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPayrollCalendarFromPayrollRun/{parentId}", jsonResponseFormatter.FormatToJSON(PayrollRunController.UnassignPayrollCalendarFromPayrollRun)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPayrollItemsToPayrollRun/{parentId}/payrollItemsId", jsonResponseFormatter.FormatToJSON(PayrollRunController.AddPayrollItemsToPayrollRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePayrollItemsFromPayrollRun/{parentId}/payrollItemsIds", jsonResponseFormatter.FormatToJSON(PayrollRunController.RemovePayrollItemsFromPayrollRun)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PayrollItem Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PayrollItem/{id}", jsonResponseFormatter.FormatToJSON(PayrollItemController.GetPayrollItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PayrollItem", jsonResponseFormatter.FormatToJSON(PayrollItemController.GetAllPayrollItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPayrollItem", jsonResponseFormatter.FormatToJSON(PayrollItemController.CreatePayrollItem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PayrollItem/{id}", jsonResponseFormatter.FormatToJSON(PayrollItemController.UpdatePayrollItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePayrollItem/{id}", jsonResponseFormatter.FormatToJSON(PayrollItemController.DeletePayrollItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPayrollRunToPayrollItem/{parentId}/payrollRunId", jsonResponseFormatter.FormatToJSON(PayrollItemController.AssignPayrollRunToPayrollItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPayrollRunFromPayrollItem/{parentId}", jsonResponseFormatter.FormatToJSON(PayrollItemController.UnassignPayrollRunFromPayrollItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEmployeeToPayrollItem/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(PayrollItemController.AssignEmployeeToPayrollItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromPayrollItem/{parentId}", jsonResponseFormatter.FormatToJSON(PayrollItemController.UnassignEmployeeFromPayrollItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // TaxWithholding Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TaxWithholding/{id}", jsonResponseFormatter.FormatToJSON(TaxWithholdingController.GetTaxWithholding)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TaxWithholding", jsonResponseFormatter.FormatToJSON(TaxWithholdingController.GetAllTaxWithholding)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTaxWithholding", jsonResponseFormatter.FormatToJSON(TaxWithholdingController.CreateTaxWithholding)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TaxWithholding/{id}", jsonResponseFormatter.FormatToJSON(TaxWithholdingController.UpdateTaxWithholding)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTaxWithholding/{id}", jsonResponseFormatter.FormatToJSON(TaxWithholdingController.DeleteTaxWithholding)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToTaxWithholding/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(TaxWithholdingController.AssignEmployeeToTaxWithholding)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromTaxWithholding/{parentId}", jsonResponseFormatter.FormatToJSON(TaxWithholdingController.UnassignEmployeeFromTaxWithholding)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // PaymentMethod Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PaymentMethod/{id}", jsonResponseFormatter.FormatToJSON(PaymentMethodController.GetPaymentMethod)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PaymentMethod", jsonResponseFormatter.FormatToJSON(PaymentMethodController.GetAllPaymentMethod)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPaymentMethod", jsonResponseFormatter.FormatToJSON(PaymentMethodController.CreatePaymentMethod)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PaymentMethod/{id}", jsonResponseFormatter.FormatToJSON(PaymentMethodController.UpdatePaymentMethod)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePaymentMethod/{id}", jsonResponseFormatter.FormatToJSON(PaymentMethodController.DeletePaymentMethod)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToPaymentMethod/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(PaymentMethodController.AssignEmployeeToPaymentMethod)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromPaymentMethod/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentMethodController.UnassignEmployeeFromPaymentMethod)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBankAccountToPaymentMethod/{parentId}/bankAccountId", jsonResponseFormatter.FormatToJSON(PaymentMethodController.AssignBankAccountToPaymentMethod)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBankAccountFromPaymentMethod/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentMethodController.UnassignBankAccountFromPaymentMethod)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Timesheet Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Timesheet/{id}", jsonResponseFormatter.FormatToJSON(TimesheetController.GetTimesheet)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Timesheet", jsonResponseFormatter.FormatToJSON(TimesheetController.GetAllTimesheet)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTimesheet", jsonResponseFormatter.FormatToJSON(TimesheetController.CreateTimesheet)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Timesheet/{id}", jsonResponseFormatter.FormatToJSON(TimesheetController.UpdateTimesheet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTimesheet/{id}", jsonResponseFormatter.FormatToJSON(TimesheetController.DeleteTimesheet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToTimesheet/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(TimesheetController.AssignEmployeeToTimesheet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromTimesheet/{parentId}", jsonResponseFormatter.FormatToJSON(TimesheetController.UnassignEmployeeFromTimesheet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTimeEntriesToTimesheet/{parentId}/timeEntriesId", jsonResponseFormatter.FormatToJSON(TimesheetController.AddTimeEntriesToTimesheet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTimeEntriesFromTimesheet/{parentId}/timeEntriesIds", jsonResponseFormatter.FormatToJSON(TimesheetController.RemoveTimeEntriesFromTimesheet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddApprovalsToTimesheet/{parentId}/approvalsId", jsonResponseFormatter.FormatToJSON(TimesheetController.AddApprovalsToTimesheet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveApprovalsFromTimesheet/{parentId}/approvalsIds", jsonResponseFormatter.FormatToJSON(TimesheetController.RemoveApprovalsFromTimesheet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // TimeEntry Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TimeEntry/{id}", jsonResponseFormatter.FormatToJSON(TimeEntryController.GetTimeEntry)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TimeEntry", jsonResponseFormatter.FormatToJSON(TimeEntryController.GetAllTimeEntry)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTimeEntry", jsonResponseFormatter.FormatToJSON(TimeEntryController.CreateTimeEntry)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TimeEntry/{id}", jsonResponseFormatter.FormatToJSON(TimeEntryController.UpdateTimeEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTimeEntry/{id}", jsonResponseFormatter.FormatToJSON(TimeEntryController.DeleteTimeEntry)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignTimesheetToTimeEntry/{parentId}/timesheetId", jsonResponseFormatter.FormatToJSON(TimeEntryController.AssignTimesheetToTimeEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTimesheetFromTimeEntry/{parentId}", jsonResponseFormatter.FormatToJSON(TimeEntryController.UnassignTimesheetFromTimeEntry)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEmployeeToTimeEntry/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(TimeEntryController.AssignEmployeeToTimeEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromTimeEntry/{parentId}", jsonResponseFormatter.FormatToJSON(TimeEntryController.UnassignEmployeeFromTimeEntry)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCostCenterToTimeEntry/{parentId}/costCenterId", jsonResponseFormatter.FormatToJSON(TimeEntryController.AssignCostCenterToTimeEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCostCenterFromTimeEntry/{parentId}", jsonResponseFormatter.FormatToJSON(TimeEntryController.UnassignCostCenterFromTimeEntry)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Approval Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Approval/{id}", jsonResponseFormatter.FormatToJSON(ApprovalController.GetApproval)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Approval", jsonResponseFormatter.FormatToJSON(ApprovalController.GetAllApproval)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewApproval", jsonResponseFormatter.FormatToJSON(ApprovalController.CreateApproval)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Approval/{id}", jsonResponseFormatter.FormatToJSON(ApprovalController.UpdateApproval)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteApproval/{id}", jsonResponseFormatter.FormatToJSON(ApprovalController.DeleteApproval)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignApproverToApproval/{parentId}/approverId", jsonResponseFormatter.FormatToJSON(ApprovalController.AssignApproverToApproval)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignApproverFromApproval/{parentId}", jsonResponseFormatter.FormatToJSON(ApprovalController.UnassignApproverFromApproval)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTimesheetToApproval/{parentId}/timesheetId", jsonResponseFormatter.FormatToJSON(ApprovalController.AssignTimesheetToApproval)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTimesheetFromApproval/{parentId}", jsonResponseFormatter.FormatToJSON(ApprovalController.UnassignTimesheetFromApproval)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLeaveRequestToApproval/{parentId}/leaveRequestId", jsonResponseFormatter.FormatToJSON(ApprovalController.AssignLeaveRequestToApproval)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLeaveRequestFromApproval/{parentId}", jsonResponseFormatter.FormatToJSON(ApprovalController.UnassignLeaveRequestFromApproval)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // LeavePolicy Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LeavePolicy/{id}", jsonResponseFormatter.FormatToJSON(LeavePolicyController.GetLeavePolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LeavePolicy", jsonResponseFormatter.FormatToJSON(LeavePolicyController.GetAllLeavePolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLeavePolicy", jsonResponseFormatter.FormatToJSON(LeavePolicyController.CreateLeavePolicy)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LeavePolicy/{id}", jsonResponseFormatter.FormatToJSON(LeavePolicyController.UpdateLeavePolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLeavePolicy/{id}", jsonResponseFormatter.FormatToJSON(LeavePolicyController.DeleteLeavePolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToLeavePolicy/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(LeavePolicyController.AssignOrganizationToLeavePolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromLeavePolicy/{parentId}", jsonResponseFormatter.FormatToJSON(LeavePolicyController.UnassignOrganizationFromLeavePolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLeaveRequestsToLeavePolicy/{parentId}/leaveRequestsId", jsonResponseFormatter.FormatToJSON(LeavePolicyController.AddLeaveRequestsToLeavePolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLeaveRequestsFromLeavePolicy/{parentId}/leaveRequestsIds", jsonResponseFormatter.FormatToJSON(LeavePolicyController.RemoveLeaveRequestsFromLeavePolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // LeaveRequest Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LeaveRequest/{id}", jsonResponseFormatter.FormatToJSON(LeaveRequestController.GetLeaveRequest)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LeaveRequest", jsonResponseFormatter.FormatToJSON(LeaveRequestController.GetAllLeaveRequest)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLeaveRequest", jsonResponseFormatter.FormatToJSON(LeaveRequestController.CreateLeaveRequest)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LeaveRequest/{id}", jsonResponseFormatter.FormatToJSON(LeaveRequestController.UpdateLeaveRequest)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLeaveRequest/{id}", jsonResponseFormatter.FormatToJSON(LeaveRequestController.DeleteLeaveRequest)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToLeaveRequest/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(LeaveRequestController.AssignEmployeeToLeaveRequest)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromLeaveRequest/{parentId}", jsonResponseFormatter.FormatToJSON(LeaveRequestController.UnassignEmployeeFromLeaveRequest)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLeavePolicyToLeaveRequest/{parentId}/leavePolicyId", jsonResponseFormatter.FormatToJSON(LeaveRequestController.AssignLeavePolicyToLeaveRequest)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLeavePolicyFromLeaveRequest/{parentId}", jsonResponseFormatter.FormatToJSON(LeaveRequestController.UnassignLeavePolicyFromLeaveRequest)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddApprovalsToLeaveRequest/{parentId}/approvalsId", jsonResponseFormatter.FormatToJSON(LeaveRequestController.AddApprovalsToLeaveRequest)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveApprovalsFromLeaveRequest/{parentId}/approvalsIds", jsonResponseFormatter.FormatToJSON(LeaveRequestController.RemoveApprovalsFromLeaveRequest)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PerformanceCycle Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PerformanceCycle/{id}", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.GetPerformanceCycle)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PerformanceCycle", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.GetAllPerformanceCycle)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPerformanceCycle", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.CreatePerformanceCycle)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PerformanceCycle/{id}", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.UpdatePerformanceCycle)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePerformanceCycle/{id}", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.DeletePerformanceCycle)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrganizationToPerformanceCycle/{parentId}/organizationId", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.AssignOrganizationToPerformanceCycle)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrganizationFromPerformanceCycle/{parentId}", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.UnassignOrganizationFromPerformanceCycle)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddReviewsToPerformanceCycle/{parentId}/reviewsId", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.AddReviewsToPerformanceCycle)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReviewsFromPerformanceCycle/{parentId}/reviewsIds", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.RemoveReviewsFromPerformanceCycle)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddGoalsToPerformanceCycle/{parentId}/goalsId", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.AddGoalsToPerformanceCycle)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveGoalsFromPerformanceCycle/{parentId}/goalsIds", jsonResponseFormatter.FormatToJSON(PerformanceCycleController.RemoveGoalsFromPerformanceCycle)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Goal Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Goal/{id}", jsonResponseFormatter.FormatToJSON(GoalController.GetGoal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Goal", jsonResponseFormatter.FormatToJSON(GoalController.GetAllGoal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewGoal", jsonResponseFormatter.FormatToJSON(GoalController.CreateGoal)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Goal/{id}", jsonResponseFormatter.FormatToJSON(GoalController.UpdateGoal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteGoal/{id}", jsonResponseFormatter.FormatToJSON(GoalController.DeleteGoal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToGoal/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(GoalController.AssignEmployeeToGoal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromGoal/{parentId}", jsonResponseFormatter.FormatToJSON(GoalController.UnassignEmployeeFromGoal)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCycleToGoal/{parentId}/cycleId", jsonResponseFormatter.FormatToJSON(GoalController.AssignCycleToGoal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCycleFromGoal/{parentId}", jsonResponseFormatter.FormatToJSON(GoalController.UnassignCycleFromGoal)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignParentGoalToGoal/{parentId}/parentGoalId", jsonResponseFormatter.FormatToJSON(GoalController.AssignParentGoalToGoal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignParentGoalFromGoal/{parentId}", jsonResponseFormatter.FormatToJSON(GoalController.UnassignParentGoalFromGoal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddChildGoalsToGoal/{parentId}/childGoalsId", jsonResponseFormatter.FormatToJSON(GoalController.AddChildGoalsToGoal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveChildGoalsFromGoal/{parentId}/childGoalsIds", jsonResponseFormatter.FormatToJSON(GoalController.RemoveChildGoalsFromGoal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PerformanceReview Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PerformanceReview/{id}", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.GetPerformanceReview)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PerformanceReview", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.GetAllPerformanceReview)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPerformanceReview", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.CreatePerformanceReview)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PerformanceReview/{id}", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.UpdatePerformanceReview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePerformanceReview/{id}", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.DeletePerformanceReview)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToPerformanceReview/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.AssignEmployeeToPerformanceReview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromPerformanceReview/{parentId}", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.UnassignEmployeeFromPerformanceReview)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignReviewerToPerformanceReview/{parentId}/reviewerId", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.AssignReviewerToPerformanceReview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignReviewerFromPerformanceReview/{parentId}", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.UnassignReviewerFromPerformanceReview)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCycleToPerformanceReview/{parentId}/cycleId", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.AssignCycleToPerformanceReview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCycleFromPerformanceReview/{parentId}", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.UnassignCycleFromPerformanceReview)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCompetencyRatingsToPerformanceReview/{parentId}/competencyRatingsId", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.AddCompetencyRatingsToPerformanceReview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCompetencyRatingsFromPerformanceReview/{parentId}/competencyRatingsIds", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.RemoveCompetencyRatingsFromPerformanceReview)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddGoalsToPerformanceReview/{parentId}/goalsId", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.AddGoalsToPerformanceReview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveGoalsFromPerformanceReview/{parentId}/goalsIds", jsonResponseFormatter.FormatToJSON(PerformanceReviewController.RemoveGoalsFromPerformanceReview)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CompetencyRating Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CompetencyRating/{id}", jsonResponseFormatter.FormatToJSON(CompetencyRatingController.GetCompetencyRating)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CompetencyRating", jsonResponseFormatter.FormatToJSON(CompetencyRatingController.GetAllCompetencyRating)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCompetencyRating", jsonResponseFormatter.FormatToJSON(CompetencyRatingController.CreateCompetencyRating)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CompetencyRating/{id}", jsonResponseFormatter.FormatToJSON(CompetencyRatingController.UpdateCompetencyRating)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCompetencyRating/{id}", jsonResponseFormatter.FormatToJSON(CompetencyRatingController.DeleteCompetencyRating)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignReviewToCompetencyRating/{parentId}/reviewId", jsonResponseFormatter.FormatToJSON(CompetencyRatingController.AssignReviewToCompetencyRating)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignReviewFromCompetencyRating/{parentId}", jsonResponseFormatter.FormatToJSON(CompetencyRatingController.UnassignReviewFromCompetencyRating)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCompetencyToCompetencyRating/{parentId}/competencyId", jsonResponseFormatter.FormatToJSON(CompetencyRatingController.AssignCompetencyToCompetencyRating)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCompetencyFromCompetencyRating/{parentId}", jsonResponseFormatter.FormatToJSON(CompetencyRatingController.UnassignCompetencyFromCompetencyRating)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // TrainingCourse Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TrainingCourse/{id}", jsonResponseFormatter.FormatToJSON(TrainingCourseController.GetTrainingCourse)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TrainingCourse", jsonResponseFormatter.FormatToJSON(TrainingCourseController.GetAllTrainingCourse)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTrainingCourse", jsonResponseFormatter.FormatToJSON(TrainingCourseController.CreateTrainingCourse)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TrainingCourse/{id}", jsonResponseFormatter.FormatToJSON(TrainingCourseController.UpdateTrainingCourse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTrainingCourse/{id}", jsonResponseFormatter.FormatToJSON(TrainingCourseController.DeleteTrainingCourse)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPrerequisitesToTrainingCourse/{parentId}/prerequisitesId", jsonResponseFormatter.FormatToJSON(TrainingCourseController.AddPrerequisitesToTrainingCourse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePrerequisitesFromTrainingCourse/{parentId}/prerequisitesIds", jsonResponseFormatter.FormatToJSON(TrainingCourseController.RemovePrerequisitesFromTrainingCourse)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEnrollmentsToTrainingCourse/{parentId}/enrollmentsId", jsonResponseFormatter.FormatToJSON(TrainingCourseController.AddEnrollmentsToTrainingCourse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEnrollmentsFromTrainingCourse/{parentId}/enrollmentsIds", jsonResponseFormatter.FormatToJSON(TrainingCourseController.RemoveEnrollmentsFromTrainingCourse)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddJobProfilesToTrainingCourse/{parentId}/jobProfilesId", jsonResponseFormatter.FormatToJSON(TrainingCourseController.AddJobProfilesToTrainingCourse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveJobProfilesFromTrainingCourse/{parentId}/jobProfilesIds", jsonResponseFormatter.FormatToJSON(TrainingCourseController.RemoveJobProfilesFromTrainingCourse)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // TrainingEnrollment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TrainingEnrollment/{id}", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.GetTrainingEnrollment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TrainingEnrollment", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.GetAllTrainingEnrollment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTrainingEnrollment", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.CreateTrainingEnrollment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TrainingEnrollment/{id}", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.UpdateTrainingEnrollment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTrainingEnrollment/{id}", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.DeleteTrainingEnrollment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCourseToTrainingEnrollment/{parentId}/courseId", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.AssignCourseToTrainingEnrollment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCourseFromTrainingEnrollment/{parentId}", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.UnassignCourseFromTrainingEnrollment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEmployeeToTrainingEnrollment/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.AssignEmployeeToTrainingEnrollment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromTrainingEnrollment/{parentId}", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.UnassignEmployeeFromTrainingEnrollment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInstructorToTrainingEnrollment/{parentId}/instructorId", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.AssignInstructorToTrainingEnrollment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInstructorFromTrainingEnrollment/{parentId}", jsonResponseFormatter.FormatToJSON(TrainingEnrollmentController.UnassignInstructorFromTrainingEnrollment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Certification Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Certification/{id}", jsonResponseFormatter.FormatToJSON(CertificationController.GetCertification)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Certification", jsonResponseFormatter.FormatToJSON(CertificationController.GetAllCertification)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCertification", jsonResponseFormatter.FormatToJSON(CertificationController.CreateCertification)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Certification/{id}", jsonResponseFormatter.FormatToJSON(CertificationController.UpdateCertification)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCertification/{id}", jsonResponseFormatter.FormatToJSON(CertificationController.DeleteCertification)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToCertification/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(CertificationController.AssignEmployeeToCertification)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromCertification/{parentId}", jsonResponseFormatter.FormatToJSON(CertificationController.UnassignEmployeeFromCertification)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCourseToCertification/{parentId}/courseId", jsonResponseFormatter.FormatToJSON(CertificationController.AssignCourseToCertification)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCourseFromCertification/{parentId}", jsonResponseFormatter.FormatToJSON(CertificationController.UnassignCourseFromCertification)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // JobRequisition Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/JobRequisition/{id}", jsonResponseFormatter.FormatToJSON(JobRequisitionController.GetJobRequisition)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/JobRequisition", jsonResponseFormatter.FormatToJSON(JobRequisitionController.GetAllJobRequisition)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewJobRequisition", jsonResponseFormatter.FormatToJSON(JobRequisitionController.CreateJobRequisition)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/JobRequisition/{id}", jsonResponseFormatter.FormatToJSON(JobRequisitionController.UpdateJobRequisition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteJobRequisition/{id}", jsonResponseFormatter.FormatToJSON(JobRequisitionController.DeleteJobRequisition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignDepartmentToJobRequisition/{parentId}/departmentId", jsonResponseFormatter.FormatToJSON(JobRequisitionController.AssignDepartmentToJobRequisition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDepartmentFromJobRequisition/{parentId}", jsonResponseFormatter.FormatToJSON(JobRequisitionController.UnassignDepartmentFromJobRequisition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignHiringManagerToJobRequisition/{parentId}/hiringManagerId", jsonResponseFormatter.FormatToJSON(JobRequisitionController.AssignHiringManagerToJobRequisition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignHiringManagerFromJobRequisition/{parentId}", jsonResponseFormatter.FormatToJSON(JobRequisitionController.UnassignHiringManagerFromJobRequisition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRecruiterToJobRequisition/{parentId}/recruiterId", jsonResponseFormatter.FormatToJSON(JobRequisitionController.AssignRecruiterToJobRequisition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRecruiterFromJobRequisition/{parentId}", jsonResponseFormatter.FormatToJSON(JobRequisitionController.UnassignRecruiterFromJobRequisition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignJobProfileToJobRequisition/{parentId}/jobProfileId", jsonResponseFormatter.FormatToJSON(JobRequisitionController.AssignJobProfileToJobRequisition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignJobProfileFromJobRequisition/{parentId}", jsonResponseFormatter.FormatToJSON(JobRequisitionController.UnassignJobProfileFromJobRequisition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCandidatesToJobRequisition/{parentId}/candidatesId", jsonResponseFormatter.FormatToJSON(JobRequisitionController.AddCandidatesToJobRequisition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCandidatesFromJobRequisition/{parentId}/candidatesIds", jsonResponseFormatter.FormatToJSON(JobRequisitionController.RemoveCandidatesFromJobRequisition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInterviewsToJobRequisition/{parentId}/interviewsId", jsonResponseFormatter.FormatToJSON(JobRequisitionController.AddInterviewsToJobRequisition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInterviewsFromJobRequisition/{parentId}/interviewsIds", jsonResponseFormatter.FormatToJSON(JobRequisitionController.RemoveInterviewsFromJobRequisition)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOffersToJobRequisition/{parentId}/offersId", jsonResponseFormatter.FormatToJSON(JobRequisitionController.AddOffersToJobRequisition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOffersFromJobRequisition/{parentId}/offersIds", jsonResponseFormatter.FormatToJSON(JobRequisitionController.RemoveOffersFromJobRequisition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Candidate Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Candidate/{id}", jsonResponseFormatter.FormatToJSON(CandidateController.GetCandidate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Candidate", jsonResponseFormatter.FormatToJSON(CandidateController.GetAllCandidate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCandidate", jsonResponseFormatter.FormatToJSON(CandidateController.CreateCandidate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Candidate/{id}", jsonResponseFormatter.FormatToJSON(CandidateController.UpdateCandidate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCandidate/{id}", jsonResponseFormatter.FormatToJSON(CandidateController.DeleteCandidate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddApplicationsToCandidate/{parentId}/applicationsId", jsonResponseFormatter.FormatToJSON(CandidateController.AddApplicationsToCandidate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveApplicationsFromCandidate/{parentId}/applicationsIds", jsonResponseFormatter.FormatToJSON(CandidateController.RemoveApplicationsFromCandidate)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInterviewsToCandidate/{parentId}/interviewsId", jsonResponseFormatter.FormatToJSON(CandidateController.AddInterviewsToCandidate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInterviewsFromCandidate/{parentId}/interviewsIds", jsonResponseFormatter.FormatToJSON(CandidateController.RemoveInterviewsFromCandidate)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOffersToCandidate/{parentId}/offersId", jsonResponseFormatter.FormatToJSON(CandidateController.AddOffersToCandidate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOffersFromCandidate/{parentId}/offersIds", jsonResponseFormatter.FormatToJSON(CandidateController.RemoveOffersFromCandidate)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDocumentsToCandidate/{parentId}/documentsId", jsonResponseFormatter.FormatToJSON(CandidateController.AddDocumentsToCandidate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDocumentsFromCandidate/{parentId}/documentsIds", jsonResponseFormatter.FormatToJSON(CandidateController.RemoveDocumentsFromCandidate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // JobApplication Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/JobApplication/{id}", jsonResponseFormatter.FormatToJSON(JobApplicationController.GetJobApplication)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/JobApplication", jsonResponseFormatter.FormatToJSON(JobApplicationController.GetAllJobApplication)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewJobApplication", jsonResponseFormatter.FormatToJSON(JobApplicationController.CreateJobApplication)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/JobApplication/{id}", jsonResponseFormatter.FormatToJSON(JobApplicationController.UpdateJobApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteJobApplication/{id}", jsonResponseFormatter.FormatToJSON(JobApplicationController.DeleteJobApplication)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCandidateToJobApplication/{parentId}/candidateId", jsonResponseFormatter.FormatToJSON(JobApplicationController.AssignCandidateToJobApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCandidateFromJobApplication/{parentId}", jsonResponseFormatter.FormatToJSON(JobApplicationController.UnassignCandidateFromJobApplication)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRequisitionToJobApplication/{parentId}/requisitionId", jsonResponseFormatter.FormatToJSON(JobApplicationController.AssignRequisitionToJobApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRequisitionFromJobApplication/{parentId}", jsonResponseFormatter.FormatToJSON(JobApplicationController.UnassignRequisitionFromJobApplication)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddScreeningsToJobApplication/{parentId}/screeningsId", jsonResponseFormatter.FormatToJSON(JobApplicationController.AddScreeningsToJobApplication)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveScreeningsFromJobApplication/{parentId}/screeningsIds", jsonResponseFormatter.FormatToJSON(JobApplicationController.RemoveScreeningsFromJobApplication)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Interview Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Interview/{id}", jsonResponseFormatter.FormatToJSON(InterviewController.GetInterview)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Interview", jsonResponseFormatter.FormatToJSON(InterviewController.GetAllInterview)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInterview", jsonResponseFormatter.FormatToJSON(InterviewController.CreateInterview)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Interview/{id}", jsonResponseFormatter.FormatToJSON(InterviewController.UpdateInterview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInterview/{id}", jsonResponseFormatter.FormatToJSON(InterviewController.DeleteInterview)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRequisitionToInterview/{parentId}/requisitionId", jsonResponseFormatter.FormatToJSON(InterviewController.AssignRequisitionToInterview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRequisitionFromInterview/{parentId}", jsonResponseFormatter.FormatToJSON(InterviewController.UnassignRequisitionFromInterview)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCandidateToInterview/{parentId}/candidateId", jsonResponseFormatter.FormatToJSON(InterviewController.AssignCandidateToInterview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCandidateFromInterview/{parentId}", jsonResponseFormatter.FormatToJSON(InterviewController.UnassignCandidateFromInterview)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInterviewersToInterview/{parentId}/interviewersId", jsonResponseFormatter.FormatToJSON(InterviewController.AddInterviewersToInterview)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInterviewersFromInterview/{parentId}/interviewersIds", jsonResponseFormatter.FormatToJSON(InterviewController.RemoveInterviewersFromInterview)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Screening Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Screening/{id}", jsonResponseFormatter.FormatToJSON(ScreeningController.GetScreening)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Screening", jsonResponseFormatter.FormatToJSON(ScreeningController.GetAllScreening)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewScreening", jsonResponseFormatter.FormatToJSON(ScreeningController.CreateScreening)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Screening/{id}", jsonResponseFormatter.FormatToJSON(ScreeningController.UpdateScreening)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteScreening/{id}", jsonResponseFormatter.FormatToJSON(ScreeningController.DeleteScreening)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignApplicationToScreening/{parentId}/applicationId", jsonResponseFormatter.FormatToJSON(ScreeningController.AssignApplicationToScreening)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignApplicationFromScreening/{parentId}", jsonResponseFormatter.FormatToJSON(ScreeningController.UnassignApplicationFromScreening)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Offer Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Offer/{id}", jsonResponseFormatter.FormatToJSON(OfferController.GetOffer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Offer", jsonResponseFormatter.FormatToJSON(OfferController.GetAllOffer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOffer", jsonResponseFormatter.FormatToJSON(OfferController.CreateOffer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Offer/{id}", jsonResponseFormatter.FormatToJSON(OfferController.UpdateOffer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOffer/{id}", jsonResponseFormatter.FormatToJSON(OfferController.DeleteOffer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRequisitionToOffer/{parentId}/requisitionId", jsonResponseFormatter.FormatToJSON(OfferController.AssignRequisitionToOffer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRequisitionFromOffer/{parentId}", jsonResponseFormatter.FormatToJSON(OfferController.UnassignRequisitionFromOffer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCandidateToOffer/{parentId}/candidateId", jsonResponseFormatter.FormatToJSON(OfferController.AssignCandidateToOffer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCandidateFromOffer/{parentId}", jsonResponseFormatter.FormatToJSON(OfferController.UnassignCandidateFromOffer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignApprovedByToOffer/{parentId}/approvedById", jsonResponseFormatter.FormatToJSON(OfferController.AssignApprovedByToOffer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignApprovedByFromOffer/{parentId}", jsonResponseFormatter.FormatToJSON(OfferController.UnassignApprovedByFromOffer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignContractToOffer/{parentId}/contractId", jsonResponseFormatter.FormatToJSON(OfferController.AssignContractToOffer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignContractFromOffer/{parentId}", jsonResponseFormatter.FormatToJSON(OfferController.UnassignContractFromOffer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // OnboardingTask Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/OnboardingTask/{id}", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.GetOnboardingTask)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/OnboardingTask", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.GetAllOnboardingTask)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOnboardingTask", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.CreateOnboardingTask)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/OnboardingTask/{id}", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.UpdateOnboardingTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOnboardingTask/{id}", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.DeleteOnboardingTask)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToOnboardingTask/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.AssignEmployeeToOnboardingTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromOnboardingTask/{parentId}", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.UnassignEmployeeFromOnboardingTask)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAssignedToToOnboardingTask/{parentId}/assignedToId", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.AssignAssignedToToOnboardingTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAssignedToFromOnboardingTask/{parentId}", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.UnassignAssignedToFromOnboardingTask)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRelatedOfferToOnboardingTask/{parentId}/relatedOfferId", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.AssignRelatedOfferToOnboardingTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRelatedOfferFromOnboardingTask/{parentId}", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.UnassignRelatedOfferFromOnboardingTask)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDependenciesToOnboardingTask/{parentId}/dependenciesId", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.AddDependenciesToOnboardingTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDependenciesFromOnboardingTask/{parentId}/dependenciesIds", jsonResponseFormatter.FormatToJSON(OnboardingTaskController.RemoveDependenciesFromOnboardingTask)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // BackgroundCheck Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BackgroundCheck/{id}", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.GetBackgroundCheck)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BackgroundCheck", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.GetAllBackgroundCheck)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBackgroundCheck", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.CreateBackgroundCheck)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BackgroundCheck/{id}", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.UpdateBackgroundCheck)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBackgroundCheck/{id}", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.DeleteBackgroundCheck)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCandidateToBackgroundCheck/{parentId}/candidateId", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.AssignCandidateToBackgroundCheck)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCandidateFromBackgroundCheck/{parentId}", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.UnassignCandidateFromBackgroundCheck)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRequisitionToBackgroundCheck/{parentId}/requisitionId", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.AssignRequisitionToBackgroundCheck)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRequisitionFromBackgroundCheck/{parentId}", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.UnassignRequisitionFromBackgroundCheck)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignReportToBackgroundCheck/{parentId}/reportId", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.AssignReportToBackgroundCheck)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignReportFromBackgroundCheck/{parentId}", jsonResponseFormatter.FormatToJSON(BackgroundCheckController.UnassignReportFromBackgroundCheck)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Document Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Document/{id}", jsonResponseFormatter.FormatToJSON(DocumentController.GetDocument)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Document", jsonResponseFormatter.FormatToJSON(DocumentController.GetAllDocument)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDocument", jsonResponseFormatter.FormatToJSON(DocumentController.CreateDocument)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Document/{id}", jsonResponseFormatter.FormatToJSON(DocumentController.UpdateDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDocument/{id}", jsonResponseFormatter.FormatToJSON(DocumentController.DeleteDocument)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCandidateToDocument/{parentId}/candidateId", jsonResponseFormatter.FormatToJSON(DocumentController.AssignCandidateToDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCandidateFromDocument/{parentId}", jsonResponseFormatter.FormatToJSON(DocumentController.UnassignCandidateFromDocument)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEmployeeToDocument/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(DocumentController.AssignEmployeeToDocument)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromDocument/{parentId}", jsonResponseFormatter.FormatToJSON(DocumentController.UnassignEmployeeFromDocument)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AddAcknowledgementsToPolicy/{parentId}/acknowledgementsId", jsonResponseFormatter.FormatToJSON(PolicyController.AddAcknowledgementsToPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAcknowledgementsFromPolicy/{parentId}/acknowledgementsIds", jsonResponseFormatter.FormatToJSON(PolicyController.RemoveAcknowledgementsFromPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PolicyAcknowledgement Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PolicyAcknowledgement/{id}", jsonResponseFormatter.FormatToJSON(PolicyAcknowledgementController.GetPolicyAcknowledgement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PolicyAcknowledgement", jsonResponseFormatter.FormatToJSON(PolicyAcknowledgementController.GetAllPolicyAcknowledgement)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPolicyAcknowledgement", jsonResponseFormatter.FormatToJSON(PolicyAcknowledgementController.CreatePolicyAcknowledgement)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PolicyAcknowledgement/{id}", jsonResponseFormatter.FormatToJSON(PolicyAcknowledgementController.UpdatePolicyAcknowledgement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePolicyAcknowledgement/{id}", jsonResponseFormatter.FormatToJSON(PolicyAcknowledgementController.DeletePolicyAcknowledgement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPolicyToPolicyAcknowledgement/{parentId}/policyId", jsonResponseFormatter.FormatToJSON(PolicyAcknowledgementController.AssignPolicyToPolicyAcknowledgement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPolicyFromPolicyAcknowledgement/{parentId}", jsonResponseFormatter.FormatToJSON(PolicyAcknowledgementController.UnassignPolicyFromPolicyAcknowledgement)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEmployeeToPolicyAcknowledgement/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(PolicyAcknowledgementController.AssignEmployeeToPolicyAcknowledgement)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromPolicyAcknowledgement/{parentId}", jsonResponseFormatter.FormatToJSON(PolicyAcknowledgementController.UnassignEmployeeFromPolicyAcknowledgement)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Termination Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Termination/{id}", jsonResponseFormatter.FormatToJSON(TerminationController.GetTermination)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Termination", jsonResponseFormatter.FormatToJSON(TerminationController.GetAllTermination)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTermination", jsonResponseFormatter.FormatToJSON(TerminationController.CreateTermination)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Termination/{id}", jsonResponseFormatter.FormatToJSON(TerminationController.UpdateTermination)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTermination/{id}", jsonResponseFormatter.FormatToJSON(TerminationController.DeleteTermination)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToTermination/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(TerminationController.AssignEmployeeToTermination)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromTermination/{parentId}", jsonResponseFormatter.FormatToJSON(TerminationController.UnassignEmployeeFromTermination)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAssignmentToTermination/{parentId}/assignmentId", jsonResponseFormatter.FormatToJSON(TerminationController.AssignAssignmentToTermination)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAssignmentFromTermination/{parentId}", jsonResponseFormatter.FormatToJSON(TerminationController.UnassignAssignmentFromTermination)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // WorkAuthorization Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/WorkAuthorization/{id}", jsonResponseFormatter.FormatToJSON(WorkAuthorizationController.GetWorkAuthorization)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/WorkAuthorization", jsonResponseFormatter.FormatToJSON(WorkAuthorizationController.GetAllWorkAuthorization)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewWorkAuthorization", jsonResponseFormatter.FormatToJSON(WorkAuthorizationController.CreateWorkAuthorization)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/WorkAuthorization/{id}", jsonResponseFormatter.FormatToJSON(WorkAuthorizationController.UpdateWorkAuthorization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteWorkAuthorization/{id}", jsonResponseFormatter.FormatToJSON(WorkAuthorizationController.DeleteWorkAuthorization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEmployeeToWorkAuthorization/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(WorkAuthorizationController.AssignEmployeeToWorkAuthorization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromWorkAuthorization/{parentId}", jsonResponseFormatter.FormatToJSON(WorkAuthorizationController.UnassignEmployeeFromWorkAuthorization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDocumentsToWorkAuthorization/{parentId}/documentsId", jsonResponseFormatter.FormatToJSON(WorkAuthorizationController.AddDocumentsToWorkAuthorization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDocumentsFromWorkAuthorization/{parentId}/documentsIds", jsonResponseFormatter.FormatToJSON(WorkAuthorizationController.RemoveDocumentsFromWorkAuthorization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // BankAccount Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BankAccount/{id}", jsonResponseFormatter.FormatToJSON(BankAccountController.GetBankAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BankAccount", jsonResponseFormatter.FormatToJSON(BankAccountController.GetAllBankAccount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBankAccount", jsonResponseFormatter.FormatToJSON(BankAccountController.CreateBankAccount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BankAccount/{id}", jsonResponseFormatter.FormatToJSON(BankAccountController.UpdateBankAccount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBankAccount/{id}", jsonResponseFormatter.FormatToJSON(BankAccountController.DeleteBankAccount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    return router
}
