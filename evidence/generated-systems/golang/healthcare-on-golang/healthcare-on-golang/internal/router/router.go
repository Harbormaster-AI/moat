package router

import (

    HealthSystemController "healthcare-on-golang/internal/controller"
    FacilityController "healthcare-on-golang/internal/controller"
    DepartmentController "healthcare-on-golang/internal/controller"
    CareTeamController "healthcare-on-golang/internal/controller"
    ClinicianController "healthcare-on-golang/internal/controller"
    PatientController "healthcare-on-golang/internal/controller"
    AppointmentController "healthcare-on-golang/internal/controller"
    EncounterController "healthcare-on-golang/internal/controller"
    AdmissionController "healthcare-on-golang/internal/controller"
    DischargeController "healthcare-on-golang/internal/controller"
    ClinicalOrderController "healthcare-on-golang/internal/controller"
    MedicationOrderController "healthcare-on-golang/internal/controller"
    LaboratoryController "healthcare-on-golang/internal/controller"
    LaboratoryOrderController "healthcare-on-golang/internal/controller"
    LabResultController "healthcare-on-golang/internal/controller"
    ImagingCenterController "healthcare-on-golang/internal/controller"
    ImagingOrderController "healthcare-on-golang/internal/controller"
    ImagingReportController "healthcare-on-golang/internal/controller"
    ProcedureOrderController "healthcare-on-golang/internal/controller"
    ProcedureController "healthcare-on-golang/internal/controller"
    PharmacyController "healthcare-on-golang/internal/controller"
    MedicationDispenseController "healthcare-on-golang/internal/controller"
    DiagnosisController "healthcare-on-golang/internal/controller"
    ObservationController "healthcare-on-golang/internal/controller"
    CarePlanController "healthcare-on-golang/internal/controller"
    CareTaskController "healthcare-on-golang/internal/controller"
    AllergyController "healthcare-on-golang/internal/controller"
    ConditionController "healthcare-on-golang/internal/controller"
    InsurancePayerController "healthcare-on-golang/internal/controller"
    InsurancePlanController "healthcare-on-golang/internal/controller"
    CoverageController "healthcare-on-golang/internal/controller"
    ClaimController "healthcare-on-golang/internal/controller"
    AuthorizationController "healthcare-on-golang/internal/controller"
    InvoiceController "healthcare-on-golang/internal/controller"
    PaymentController "healthcare-on-golang/internal/controller"
    MedicalDeviceController "healthcare-on-golang/internal/controller"
    SoftwareUpdateController "healthcare-on-golang/internal/controller"
    MedicalSupplierController "healthcare-on-golang/internal/controller"
    InventoryItemController "healthcare-on-golang/internal/controller"
    jsonResponseFormatter "healthcare-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "healthcare-on-golang/internal/controller"

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
    // HealthSystem Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/HealthSystem/{id}", jsonResponseFormatter.FormatToJSON(HealthSystemController.GetHealthSystem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/HealthSystem", jsonResponseFormatter.FormatToJSON(HealthSystemController.GetAllHealthSystem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewHealthSystem", jsonResponseFormatter.FormatToJSON(HealthSystemController.CreateHealthSystem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/HealthSystem/{id}", jsonResponseFormatter.FormatToJSON(HealthSystemController.UpdateHealthSystem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteHealthSystem/{id}", jsonResponseFormatter.FormatToJSON(HealthSystemController.DeleteHealthSystem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddFacilitiesToHealthSystem/{parentId}/facilitiesId", jsonResponseFormatter.FormatToJSON(HealthSystemController.AddFacilitiesToHealthSystem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFacilitiesFromHealthSystem/{parentId}/facilitiesIds", jsonResponseFormatter.FormatToJSON(HealthSystemController.RemoveFacilitiesFromHealthSystem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSuppliersToHealthSystem/{parentId}/suppliersId", jsonResponseFormatter.FormatToJSON(HealthSystemController.AddSuppliersToHealthSystem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSuppliersFromHealthSystem/{parentId}/suppliersIds", jsonResponseFormatter.FormatToJSON(HealthSystemController.RemoveSuppliersFromHealthSystem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Facility Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Facility/{id}", jsonResponseFormatter.FormatToJSON(FacilityController.GetFacility)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Facility", jsonResponseFormatter.FormatToJSON(FacilityController.GetAllFacility)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFacility", jsonResponseFormatter.FormatToJSON(FacilityController.CreateFacility)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Facility/{id}", jsonResponseFormatter.FormatToJSON(FacilityController.UpdateFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFacility/{id}", jsonResponseFormatter.FormatToJSON(FacilityController.DeleteFacility)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignHealthSystemToFacility/{parentId}/healthSystemId", jsonResponseFormatter.FormatToJSON(FacilityController.AssignHealthSystemToFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignHealthSystemFromFacility/{parentId}", jsonResponseFormatter.FormatToJSON(FacilityController.UnassignHealthSystemFromFacility)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDepartmentsToFacility/{parentId}/departmentsId", jsonResponseFormatter.FormatToJSON(FacilityController.AddDepartmentsToFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDepartmentsFromFacility/{parentId}/departmentsIds", jsonResponseFormatter.FormatToJSON(FacilityController.RemoveDepartmentsFromFacility)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCareTeamsToFacility/{parentId}/careTeamsId", jsonResponseFormatter.FormatToJSON(FacilityController.AddCareTeamsToFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCareTeamsFromFacility/{parentId}/careTeamsIds", jsonResponseFormatter.FormatToJSON(FacilityController.RemoveCareTeamsFromFacility)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLaboratoriesToFacility/{parentId}/laboratoriesId", jsonResponseFormatter.FormatToJSON(FacilityController.AddLaboratoriesToFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLaboratoriesFromFacility/{parentId}/laboratoriesIds", jsonResponseFormatter.FormatToJSON(FacilityController.RemoveLaboratoriesFromFacility)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddImagingCentersToFacility/{parentId}/imagingCentersId", jsonResponseFormatter.FormatToJSON(FacilityController.AddImagingCentersToFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveImagingCentersFromFacility/{parentId}/imagingCentersIds", jsonResponseFormatter.FormatToJSON(FacilityController.RemoveImagingCentersFromFacility)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPharmaciesToFacility/{parentId}/pharmaciesId", jsonResponseFormatter.FormatToJSON(FacilityController.AddPharmaciesToFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePharmaciesFromFacility/{parentId}/pharmaciesIds", jsonResponseFormatter.FormatToJSON(FacilityController.RemovePharmaciesFromFacility)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInventoryItemsToFacility/{parentId}/inventoryItemsId", jsonResponseFormatter.FormatToJSON(FacilityController.AddInventoryItemsToFacility)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventoryItemsFromFacility/{parentId}/inventoryItemsIds", jsonResponseFormatter.FormatToJSON(FacilityController.RemoveInventoryItemsFromFacility)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AssignFacilityToDepartment/{parentId}/facilityId", jsonResponseFormatter.FormatToJSON(DepartmentController.AssignFacilityToDepartment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFacilityFromDepartment/{parentId}", jsonResponseFormatter.FormatToJSON(DepartmentController.UnassignFacilityFromDepartment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCareTeamsToDepartment/{parentId}/careTeamsId", jsonResponseFormatter.FormatToJSON(DepartmentController.AddCareTeamsToDepartment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCareTeamsFromDepartment/{parentId}/careTeamsIds", jsonResponseFormatter.FormatToJSON(DepartmentController.RemoveCareTeamsFromDepartment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CareTeam Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CareTeam/{id}", jsonResponseFormatter.FormatToJSON(CareTeamController.GetCareTeam)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CareTeam", jsonResponseFormatter.FormatToJSON(CareTeamController.GetAllCareTeam)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCareTeam", jsonResponseFormatter.FormatToJSON(CareTeamController.CreateCareTeam)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CareTeam/{id}", jsonResponseFormatter.FormatToJSON(CareTeamController.UpdateCareTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCareTeam/{id}", jsonResponseFormatter.FormatToJSON(CareTeamController.DeleteCareTeam)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignDepartmentToCareTeam/{parentId}/departmentId", jsonResponseFormatter.FormatToJSON(CareTeamController.AssignDepartmentToCareTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDepartmentFromCareTeam/{parentId}", jsonResponseFormatter.FormatToJSON(CareTeamController.UnassignDepartmentFromCareTeam)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCliniciansToCareTeam/{parentId}/cliniciansId", jsonResponseFormatter.FormatToJSON(CareTeamController.AddCliniciansToCareTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCliniciansFromCareTeam/{parentId}/cliniciansIds", jsonResponseFormatter.FormatToJSON(CareTeamController.RemoveCliniciansFromCareTeam)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPatientsToCareTeam/{parentId}/patientsId", jsonResponseFormatter.FormatToJSON(CareTeamController.AddPatientsToCareTeam)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePatientsFromCareTeam/{parentId}/patientsIds", jsonResponseFormatter.FormatToJSON(CareTeamController.RemovePatientsFromCareTeam)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Clinician Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Clinician/{id}", jsonResponseFormatter.FormatToJSON(ClinicianController.GetClinician)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Clinician", jsonResponseFormatter.FormatToJSON(ClinicianController.GetAllClinician)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewClinician", jsonResponseFormatter.FormatToJSON(ClinicianController.CreateClinician)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Clinician/{id}", jsonResponseFormatter.FormatToJSON(ClinicianController.UpdateClinician)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteClinician/{id}", jsonResponseFormatter.FormatToJSON(ClinicianController.DeleteClinician)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCareTeamsToClinician/{parentId}/careTeamsId", jsonResponseFormatter.FormatToJSON(ClinicianController.AddCareTeamsToClinician)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCareTeamsFromClinician/{parentId}/careTeamsIds", jsonResponseFormatter.FormatToJSON(ClinicianController.RemoveCareTeamsFromClinician)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAppointmentsToClinician/{parentId}/appointmentsId", jsonResponseFormatter.FormatToJSON(ClinicianController.AddAppointmentsToClinician)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAppointmentsFromClinician/{parentId}/appointmentsIds", jsonResponseFormatter.FormatToJSON(ClinicianController.RemoveAppointmentsFromClinician)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEncountersToClinician/{parentId}/encountersId", jsonResponseFormatter.FormatToJSON(ClinicianController.AddEncountersToClinician)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEncountersFromClinician/{parentId}/encountersIds", jsonResponseFormatter.FormatToJSON(ClinicianController.RemoveEncountersFromClinician)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddProceduresToClinician/{parentId}/proceduresId", jsonResponseFormatter.FormatToJSON(ClinicianController.AddProceduresToClinician)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProceduresFromClinician/{parentId}/proceduresIds", jsonResponseFormatter.FormatToJSON(ClinicianController.RemoveProceduresFromClinician)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddImagingReportsToClinician/{parentId}/imagingReportsId", jsonResponseFormatter.FormatToJSON(ClinicianController.AddImagingReportsToClinician)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveImagingReportsFromClinician/{parentId}/imagingReportsIds", jsonResponseFormatter.FormatToJSON(ClinicianController.RemoveImagingReportsFromClinician)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Patient Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Patient/{id}", jsonResponseFormatter.FormatToJSON(PatientController.GetPatient)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Patient", jsonResponseFormatter.FormatToJSON(PatientController.GetAllPatient)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPatient", jsonResponseFormatter.FormatToJSON(PatientController.CreatePatient)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Patient/{id}", jsonResponseFormatter.FormatToJSON(PatientController.UpdatePatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePatient/{id}", jsonResponseFormatter.FormatToJSON(PatientController.DeletePatient)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAppointmentsToPatient/{parentId}/appointmentsId", jsonResponseFormatter.FormatToJSON(PatientController.AddAppointmentsToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAppointmentsFromPatient/{parentId}/appointmentsIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveAppointmentsFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEncountersToPatient/{parentId}/encountersId", jsonResponseFormatter.FormatToJSON(PatientController.AddEncountersToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEncountersFromPatient/{parentId}/encountersIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveEncountersFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCarePlansToPatient/{parentId}/carePlansId", jsonResponseFormatter.FormatToJSON(PatientController.AddCarePlansToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCarePlansFromPatient/{parentId}/carePlansIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveCarePlansFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAllergiesToPatient/{parentId}/allergiesId", jsonResponseFormatter.FormatToJSON(PatientController.AddAllergiesToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAllergiesFromPatient/{parentId}/allergiesIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveAllergiesFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddConditionsToPatient/{parentId}/conditionsId", jsonResponseFormatter.FormatToJSON(PatientController.AddConditionsToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveConditionsFromPatient/{parentId}/conditionsIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveConditionsFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMedicationOrdersToPatient/{parentId}/medicationOrdersId", jsonResponseFormatter.FormatToJSON(PatientController.AddMedicationOrdersToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMedicationOrdersFromPatient/{parentId}/medicationOrdersIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveMedicationOrdersFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLabOrdersToPatient/{parentId}/labOrdersId", jsonResponseFormatter.FormatToJSON(PatientController.AddLabOrdersToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLabOrdersFromPatient/{parentId}/labOrdersIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveLabOrdersFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddImagingOrdersToPatient/{parentId}/imagingOrdersId", jsonResponseFormatter.FormatToJSON(PatientController.AddImagingOrdersToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveImagingOrdersFromPatient/{parentId}/imagingOrdersIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveImagingOrdersFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCoveragesToPatient/{parentId}/coveragesId", jsonResponseFormatter.FormatToJSON(PatientController.AddCoveragesToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCoveragesFromPatient/{parentId}/coveragesIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveCoveragesFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddClaimsToPatient/{parentId}/claimsId", jsonResponseFormatter.FormatToJSON(PatientController.AddClaimsToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveClaimsFromPatient/{parentId}/claimsIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveClaimsFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDevicesToPatient/{parentId}/devicesId", jsonResponseFormatter.FormatToJSON(PatientController.AddDevicesToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDevicesFromPatient/{parentId}/devicesIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveDevicesFromPatient)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddObservationsToPatient/{parentId}/observationsId", jsonResponseFormatter.FormatToJSON(PatientController.AddObservationsToPatient)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveObservationsFromPatient/{parentId}/observationsIds", jsonResponseFormatter.FormatToJSON(PatientController.RemoveObservationsFromPatient)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Appointment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Appointment/{id}", jsonResponseFormatter.FormatToJSON(AppointmentController.GetAppointment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Appointment", jsonResponseFormatter.FormatToJSON(AppointmentController.GetAllAppointment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAppointment", jsonResponseFormatter.FormatToJSON(AppointmentController.CreateAppointment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Appointment/{id}", jsonResponseFormatter.FormatToJSON(AppointmentController.UpdateAppointment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAppointment/{id}", jsonResponseFormatter.FormatToJSON(AppointmentController.DeleteAppointment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPatientToAppointment/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(AppointmentController.AssignPatientToAppointment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromAppointment/{parentId}", jsonResponseFormatter.FormatToJSON(AppointmentController.UnassignPatientFromAppointment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignClinicianToAppointment/{parentId}/clinicianId", jsonResponseFormatter.FormatToJSON(AppointmentController.AssignClinicianToAppointment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignClinicianFromAppointment/{parentId}", jsonResponseFormatter.FormatToJSON(AppointmentController.UnassignClinicianFromAppointment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignFacilityToAppointment/{parentId}/facilityId", jsonResponseFormatter.FormatToJSON(AppointmentController.AssignFacilityToAppointment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFacilityFromAppointment/{parentId}", jsonResponseFormatter.FormatToJSON(AppointmentController.UnassignFacilityFromAppointment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEncounterToAppointment/{parentId}/encounterId", jsonResponseFormatter.FormatToJSON(AppointmentController.AssignEncounterToAppointment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEncounterFromAppointment/{parentId}", jsonResponseFormatter.FormatToJSON(AppointmentController.UnassignEncounterFromAppointment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Encounter Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Encounter/{id}", jsonResponseFormatter.FormatToJSON(EncounterController.GetEncounter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Encounter", jsonResponseFormatter.FormatToJSON(EncounterController.GetAllEncounter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEncounter", jsonResponseFormatter.FormatToJSON(EncounterController.CreateEncounter)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Encounter/{id}", jsonResponseFormatter.FormatToJSON(EncounterController.UpdateEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEncounter/{id}", jsonResponseFormatter.FormatToJSON(EncounterController.DeleteEncounter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPatientToEncounter/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(EncounterController.AssignPatientToEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromEncounter/{parentId}", jsonResponseFormatter.FormatToJSON(EncounterController.UnassignPatientFromEncounter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignClinicianToEncounter/{parentId}/clinicianId", jsonResponseFormatter.FormatToJSON(EncounterController.AssignClinicianToEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignClinicianFromEncounter/{parentId}", jsonResponseFormatter.FormatToJSON(EncounterController.UnassignClinicianFromEncounter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignFacilityToEncounter/{parentId}/facilityId", jsonResponseFormatter.FormatToJSON(EncounterController.AssignFacilityToEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFacilityFromEncounter/{parentId}", jsonResponseFormatter.FormatToJSON(EncounterController.UnassignFacilityFromEncounter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAppointmentToEncounter/{parentId}/appointmentId", jsonResponseFormatter.FormatToJSON(EncounterController.AssignAppointmentToEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAppointmentFromEncounter/{parentId}", jsonResponseFormatter.FormatToJSON(EncounterController.UnassignAppointmentFromEncounter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAdmissionToEncounter/{parentId}/admissionId", jsonResponseFormatter.FormatToJSON(EncounterController.AssignAdmissionToEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdmissionFromEncounter/{parentId}", jsonResponseFormatter.FormatToJSON(EncounterController.UnassignAdmissionFromEncounter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDischargeToEncounter/{parentId}/dischargeId", jsonResponseFormatter.FormatToJSON(EncounterController.AssignDischargeToEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDischargeFromEncounter/{parentId}", jsonResponseFormatter.FormatToJSON(EncounterController.UnassignDischargeFromEncounter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDiagnosesToEncounter/{parentId}/diagnosesId", jsonResponseFormatter.FormatToJSON(EncounterController.AddDiagnosesToEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDiagnosesFromEncounter/{parentId}/diagnosesIds", jsonResponseFormatter.FormatToJSON(EncounterController.RemoveDiagnosesFromEncounter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddProceduresToEncounter/{parentId}/proceduresId", jsonResponseFormatter.FormatToJSON(EncounterController.AddProceduresToEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProceduresFromEncounter/{parentId}/proceduresIds", jsonResponseFormatter.FormatToJSON(EncounterController.RemoveProceduresFromEncounter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddObservationsToEncounter/{parentId}/observationsId", jsonResponseFormatter.FormatToJSON(EncounterController.AddObservationsToEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveObservationsFromEncounter/{parentId}/observationsIds", jsonResponseFormatter.FormatToJSON(EncounterController.RemoveObservationsFromEncounter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOrdersToEncounter/{parentId}/ordersId", jsonResponseFormatter.FormatToJSON(EncounterController.AddOrdersToEncounter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOrdersFromEncounter/{parentId}/ordersIds", jsonResponseFormatter.FormatToJSON(EncounterController.RemoveOrdersFromEncounter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Admission Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Admission/{id}", jsonResponseFormatter.FormatToJSON(AdmissionController.GetAdmission)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Admission", jsonResponseFormatter.FormatToJSON(AdmissionController.GetAllAdmission)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAdmission", jsonResponseFormatter.FormatToJSON(AdmissionController.CreateAdmission)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Admission/{id}", jsonResponseFormatter.FormatToJSON(AdmissionController.UpdateAdmission)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAdmission/{id}", jsonResponseFormatter.FormatToJSON(AdmissionController.DeleteAdmission)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEncounterToAdmission/{parentId}/encounterId", jsonResponseFormatter.FormatToJSON(AdmissionController.AssignEncounterToAdmission)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEncounterFromAdmission/{parentId}", jsonResponseFormatter.FormatToJSON(AdmissionController.UnassignEncounterFromAdmission)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignFacilityToAdmission/{parentId}/facilityId", jsonResponseFormatter.FormatToJSON(AdmissionController.AssignFacilityToAdmission)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFacilityFromAdmission/{parentId}", jsonResponseFormatter.FormatToJSON(AdmissionController.UnassignFacilityFromAdmission)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Discharge Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Discharge/{id}", jsonResponseFormatter.FormatToJSON(DischargeController.GetDischarge)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Discharge", jsonResponseFormatter.FormatToJSON(DischargeController.GetAllDischarge)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDischarge", jsonResponseFormatter.FormatToJSON(DischargeController.CreateDischarge)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Discharge/{id}", jsonResponseFormatter.FormatToJSON(DischargeController.UpdateDischarge)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDischarge/{id}", jsonResponseFormatter.FormatToJSON(DischargeController.DeleteDischarge)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEncounterToDischarge/{parentId}/encounterId", jsonResponseFormatter.FormatToJSON(DischargeController.AssignEncounterToDischarge)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEncounterFromDischarge/{parentId}", jsonResponseFormatter.FormatToJSON(DischargeController.UnassignEncounterFromDischarge)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ClinicalOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ClinicalOrder/{id}", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.GetClinicalOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ClinicalOrder", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.GetAllClinicalOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewClinicalOrder", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.CreateClinicalOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ClinicalOrder/{id}", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.UpdateClinicalOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteClinicalOrder/{id}", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.DeleteClinicalOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPatientToClinicalOrder/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.AssignPatientToClinicalOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromClinicalOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.UnassignPatientFromClinicalOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEncounterToClinicalOrder/{parentId}/encounterId", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.AssignEncounterToClinicalOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEncounterFromClinicalOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.UnassignEncounterFromClinicalOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOrderingClinicianToClinicalOrder/{parentId}/orderingClinicianId", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.AssignOrderingClinicianToClinicalOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrderingClinicianFromClinicalOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.UnassignOrderingClinicianFromClinicalOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddMedicationOrdersToClinicalOrder/{parentId}/medicationOrdersId", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.AddMedicationOrdersToClinicalOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMedicationOrdersFromClinicalOrder/{parentId}/medicationOrdersIds", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.RemoveMedicationOrdersFromClinicalOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLaboratoryOrdersToClinicalOrder/{parentId}/laboratoryOrdersId", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.AddLaboratoryOrdersToClinicalOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLaboratoryOrdersFromClinicalOrder/{parentId}/laboratoryOrdersIds", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.RemoveLaboratoryOrdersFromClinicalOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddImagingOrdersToClinicalOrder/{parentId}/imagingOrdersId", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.AddImagingOrdersToClinicalOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveImagingOrdersFromClinicalOrder/{parentId}/imagingOrdersIds", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.RemoveImagingOrdersFromClinicalOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddProcedureOrdersToClinicalOrder/{parentId}/procedureOrdersId", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.AddProcedureOrdersToClinicalOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProcedureOrdersFromClinicalOrder/{parentId}/procedureOrdersIds", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.RemoveProcedureOrdersFromClinicalOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAuthorizationsToClinicalOrder/{parentId}/authorizationsId", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.AddAuthorizationsToClinicalOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAuthorizationsFromClinicalOrder/{parentId}/authorizationsIds", jsonResponseFormatter.FormatToJSON(ClinicalOrderController.RemoveAuthorizationsFromClinicalOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // MedicationOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/MedicationOrder/{id}", jsonResponseFormatter.FormatToJSON(MedicationOrderController.GetMedicationOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/MedicationOrder", jsonResponseFormatter.FormatToJSON(MedicationOrderController.GetAllMedicationOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMedicationOrder", jsonResponseFormatter.FormatToJSON(MedicationOrderController.CreateMedicationOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/MedicationOrder/{id}", jsonResponseFormatter.FormatToJSON(MedicationOrderController.UpdateMedicationOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMedicationOrder/{id}", jsonResponseFormatter.FormatToJSON(MedicationOrderController.DeleteMedicationOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrderToMedicationOrder/{parentId}/orderId", jsonResponseFormatter.FormatToJSON(MedicationOrderController.AssignOrderToMedicationOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrderFromMedicationOrder/{parentId}", jsonResponseFormatter.FormatToJSON(MedicationOrderController.UnassignOrderFromMedicationOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPharmacyToMedicationOrder/{parentId}/pharmacyId", jsonResponseFormatter.FormatToJSON(MedicationOrderController.AssignPharmacyToMedicationOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPharmacyFromMedicationOrder/{parentId}", jsonResponseFormatter.FormatToJSON(MedicationOrderController.UnassignPharmacyFromMedicationOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDispensesToMedicationOrder/{parentId}/dispensesId", jsonResponseFormatter.FormatToJSON(MedicationOrderController.AddDispensesToMedicationOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDispensesFromMedicationOrder/{parentId}/dispensesIds", jsonResponseFormatter.FormatToJSON(MedicationOrderController.RemoveDispensesFromMedicationOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Laboratory Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Laboratory/{id}", jsonResponseFormatter.FormatToJSON(LaboratoryController.GetLaboratory)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Laboratory", jsonResponseFormatter.FormatToJSON(LaboratoryController.GetAllLaboratory)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLaboratory", jsonResponseFormatter.FormatToJSON(LaboratoryController.CreateLaboratory)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Laboratory/{id}", jsonResponseFormatter.FormatToJSON(LaboratoryController.UpdateLaboratory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLaboratory/{id}", jsonResponseFormatter.FormatToJSON(LaboratoryController.DeleteLaboratory)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignFacilityToLaboratory/{parentId}/facilityId", jsonResponseFormatter.FormatToJSON(LaboratoryController.AssignFacilityToLaboratory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFacilityFromLaboratory/{parentId}", jsonResponseFormatter.FormatToJSON(LaboratoryController.UnassignFacilityFromLaboratory)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLaboratoryOrdersToLaboratory/{parentId}/laboratoryOrdersId", jsonResponseFormatter.FormatToJSON(LaboratoryController.AddLaboratoryOrdersToLaboratory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLaboratoryOrdersFromLaboratory/{parentId}/laboratoryOrdersIds", jsonResponseFormatter.FormatToJSON(LaboratoryController.RemoveLaboratoryOrdersFromLaboratory)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLabResultsToLaboratory/{parentId}/labResultsId", jsonResponseFormatter.FormatToJSON(LaboratoryController.AddLabResultsToLaboratory)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLabResultsFromLaboratory/{parentId}/labResultsIds", jsonResponseFormatter.FormatToJSON(LaboratoryController.RemoveLabResultsFromLaboratory)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // LaboratoryOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LaboratoryOrder/{id}", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.GetLaboratoryOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LaboratoryOrder", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.GetAllLaboratoryOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLaboratoryOrder", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.CreateLaboratoryOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LaboratoryOrder/{id}", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.UpdateLaboratoryOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLaboratoryOrder/{id}", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.DeleteLaboratoryOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrderToLaboratoryOrder/{parentId}/orderId", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.AssignOrderToLaboratoryOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrderFromLaboratoryOrder/{parentId}", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.UnassignOrderFromLaboratoryOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLaboratoryToLaboratoryOrder/{parentId}/laboratoryId", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.AssignLaboratoryToLaboratoryOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLaboratoryFromLaboratoryOrder/{parentId}", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.UnassignLaboratoryFromLaboratoryOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddResultsToLaboratoryOrder/{parentId}/resultsId", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.AddResultsToLaboratoryOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveResultsFromLaboratoryOrder/{parentId}/resultsIds", jsonResponseFormatter.FormatToJSON(LaboratoryOrderController.RemoveResultsFromLaboratoryOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // LabResult Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LabResult/{id}", jsonResponseFormatter.FormatToJSON(LabResultController.GetLabResult)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LabResult", jsonResponseFormatter.FormatToJSON(LabResultController.GetAllLabResult)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLabResult", jsonResponseFormatter.FormatToJSON(LabResultController.CreateLabResult)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LabResult/{id}", jsonResponseFormatter.FormatToJSON(LabResultController.UpdateLabResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLabResult/{id}", jsonResponseFormatter.FormatToJSON(LabResultController.DeleteLabResult)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignLaboratoryOrderToLabResult/{parentId}/laboratoryOrderId", jsonResponseFormatter.FormatToJSON(LabResultController.AssignLaboratoryOrderToLabResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLaboratoryOrderFromLabResult/{parentId}", jsonResponseFormatter.FormatToJSON(LabResultController.UnassignLaboratoryOrderFromLabResult)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLaboratoryToLabResult/{parentId}/laboratoryId", jsonResponseFormatter.FormatToJSON(LabResultController.AssignLaboratoryToLabResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLaboratoryFromLabResult/{parentId}", jsonResponseFormatter.FormatToJSON(LabResultController.UnassignLaboratoryFromLabResult)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddObservationsToLabResult/{parentId}/observationsId", jsonResponseFormatter.FormatToJSON(LabResultController.AddObservationsToLabResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveObservationsFromLabResult/{parentId}/observationsIds", jsonResponseFormatter.FormatToJSON(LabResultController.RemoveObservationsFromLabResult)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ImagingCenter Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ImagingCenter/{id}", jsonResponseFormatter.FormatToJSON(ImagingCenterController.GetImagingCenter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ImagingCenter", jsonResponseFormatter.FormatToJSON(ImagingCenterController.GetAllImagingCenter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewImagingCenter", jsonResponseFormatter.FormatToJSON(ImagingCenterController.CreateImagingCenter)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ImagingCenter/{id}", jsonResponseFormatter.FormatToJSON(ImagingCenterController.UpdateImagingCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteImagingCenter/{id}", jsonResponseFormatter.FormatToJSON(ImagingCenterController.DeleteImagingCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignFacilityToImagingCenter/{parentId}/facilityId", jsonResponseFormatter.FormatToJSON(ImagingCenterController.AssignFacilityToImagingCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFacilityFromImagingCenter/{parentId}", jsonResponseFormatter.FormatToJSON(ImagingCenterController.UnassignFacilityFromImagingCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddImagingOrdersToImagingCenter/{parentId}/imagingOrdersId", jsonResponseFormatter.FormatToJSON(ImagingCenterController.AddImagingOrdersToImagingCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveImagingOrdersFromImagingCenter/{parentId}/imagingOrdersIds", jsonResponseFormatter.FormatToJSON(ImagingCenterController.RemoveImagingOrdersFromImagingCenter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddImagingReportsToImagingCenter/{parentId}/imagingReportsId", jsonResponseFormatter.FormatToJSON(ImagingCenterController.AddImagingReportsToImagingCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveImagingReportsFromImagingCenter/{parentId}/imagingReportsIds", jsonResponseFormatter.FormatToJSON(ImagingCenterController.RemoveImagingReportsFromImagingCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ImagingOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ImagingOrder/{id}", jsonResponseFormatter.FormatToJSON(ImagingOrderController.GetImagingOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ImagingOrder", jsonResponseFormatter.FormatToJSON(ImagingOrderController.GetAllImagingOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewImagingOrder", jsonResponseFormatter.FormatToJSON(ImagingOrderController.CreateImagingOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ImagingOrder/{id}", jsonResponseFormatter.FormatToJSON(ImagingOrderController.UpdateImagingOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteImagingOrder/{id}", jsonResponseFormatter.FormatToJSON(ImagingOrderController.DeleteImagingOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrderToImagingOrder/{parentId}/orderId", jsonResponseFormatter.FormatToJSON(ImagingOrderController.AssignOrderToImagingOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrderFromImagingOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ImagingOrderController.UnassignOrderFromImagingOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignImagingCenterToImagingOrder/{parentId}/imagingCenterId", jsonResponseFormatter.FormatToJSON(ImagingOrderController.AssignImagingCenterToImagingOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignImagingCenterFromImagingOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ImagingOrderController.UnassignImagingCenterFromImagingOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddReportsToImagingOrder/{parentId}/reportsId", jsonResponseFormatter.FormatToJSON(ImagingOrderController.AddReportsToImagingOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReportsFromImagingOrder/{parentId}/reportsIds", jsonResponseFormatter.FormatToJSON(ImagingOrderController.RemoveReportsFromImagingOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ImagingReport Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ImagingReport/{id}", jsonResponseFormatter.FormatToJSON(ImagingReportController.GetImagingReport)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ImagingReport", jsonResponseFormatter.FormatToJSON(ImagingReportController.GetAllImagingReport)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewImagingReport", jsonResponseFormatter.FormatToJSON(ImagingReportController.CreateImagingReport)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ImagingReport/{id}", jsonResponseFormatter.FormatToJSON(ImagingReportController.UpdateImagingReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteImagingReport/{id}", jsonResponseFormatter.FormatToJSON(ImagingReportController.DeleteImagingReport)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignImagingOrderToImagingReport/{parentId}/imagingOrderId", jsonResponseFormatter.FormatToJSON(ImagingReportController.AssignImagingOrderToImagingReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignImagingOrderFromImagingReport/{parentId}", jsonResponseFormatter.FormatToJSON(ImagingReportController.UnassignImagingOrderFromImagingReport)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignClinicianToImagingReport/{parentId}/clinicianId", jsonResponseFormatter.FormatToJSON(ImagingReportController.AssignClinicianToImagingReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignClinicianFromImagingReport/{parentId}", jsonResponseFormatter.FormatToJSON(ImagingReportController.UnassignClinicianFromImagingReport)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEncounterToImagingReport/{parentId}/encounterId", jsonResponseFormatter.FormatToJSON(ImagingReportController.AssignEncounterToImagingReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEncounterFromImagingReport/{parentId}", jsonResponseFormatter.FormatToJSON(ImagingReportController.UnassignEncounterFromImagingReport)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignImagingCenterToImagingReport/{parentId}/imagingCenterId", jsonResponseFormatter.FormatToJSON(ImagingReportController.AssignImagingCenterToImagingReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignImagingCenterFromImagingReport/{parentId}", jsonResponseFormatter.FormatToJSON(ImagingReportController.UnassignImagingCenterFromImagingReport)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ProcedureOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ProcedureOrder/{id}", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.GetProcedureOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ProcedureOrder", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.GetAllProcedureOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewProcedureOrder", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.CreateProcedureOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ProcedureOrder/{id}", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.UpdateProcedureOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteProcedureOrder/{id}", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.DeleteProcedureOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOrderToProcedureOrder/{parentId}/orderId", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.AssignOrderToProcedureOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrderFromProcedureOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.UnassignOrderFromProcedureOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignFacilityToProcedureOrder/{parentId}/facilityId", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.AssignFacilityToProcedureOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFacilityFromProcedureOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.UnassignFacilityFromProcedureOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProcedureToProcedureOrder/{parentId}/procedureId", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.AssignProcedureToProcedureOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProcedureFromProcedureOrder/{parentId}", jsonResponseFormatter.FormatToJSON(ProcedureOrderController.UnassignProcedureFromProcedureOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

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
    router.HandleFunc("/api/AssignEncounterToProcedure/{parentId}/encounterId", jsonResponseFormatter.FormatToJSON(ProcedureController.AssignEncounterToProcedure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEncounterFromProcedure/{parentId}", jsonResponseFormatter.FormatToJSON(ProcedureController.UnassignEncounterFromProcedure)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPerformerToProcedure/{parentId}/performerId", jsonResponseFormatter.FormatToJSON(ProcedureController.AssignPerformerToProcedure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPerformerFromProcedure/{parentId}", jsonResponseFormatter.FormatToJSON(ProcedureController.UnassignPerformerFromProcedure)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProcedureOrderToProcedure/{parentId}/procedureOrderId", jsonResponseFormatter.FormatToJSON(ProcedureController.AssignProcedureOrderToProcedure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProcedureOrderFromProcedure/{parentId}", jsonResponseFormatter.FormatToJSON(ProcedureController.UnassignProcedureOrderFromProcedure)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Pharmacy Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Pharmacy/{id}", jsonResponseFormatter.FormatToJSON(PharmacyController.GetPharmacy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Pharmacy", jsonResponseFormatter.FormatToJSON(PharmacyController.GetAllPharmacy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPharmacy", jsonResponseFormatter.FormatToJSON(PharmacyController.CreatePharmacy)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Pharmacy/{id}", jsonResponseFormatter.FormatToJSON(PharmacyController.UpdatePharmacy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePharmacy/{id}", jsonResponseFormatter.FormatToJSON(PharmacyController.DeletePharmacy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignFacilityToPharmacy/{parentId}/facilityId", jsonResponseFormatter.FormatToJSON(PharmacyController.AssignFacilityToPharmacy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFacilityFromPharmacy/{parentId}", jsonResponseFormatter.FormatToJSON(PharmacyController.UnassignFacilityFromPharmacy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddMedicationDispensesToPharmacy/{parentId}/medicationDispensesId", jsonResponseFormatter.FormatToJSON(PharmacyController.AddMedicationDispensesToPharmacy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMedicationDispensesFromPharmacy/{parentId}/medicationDispensesIds", jsonResponseFormatter.FormatToJSON(PharmacyController.RemoveMedicationDispensesFromPharmacy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMedicationOrdersToPharmacy/{parentId}/medicationOrdersId", jsonResponseFormatter.FormatToJSON(PharmacyController.AddMedicationOrdersToPharmacy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMedicationOrdersFromPharmacy/{parentId}/medicationOrdersIds", jsonResponseFormatter.FormatToJSON(PharmacyController.RemoveMedicationOrdersFromPharmacy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // MedicationDispense Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/MedicationDispense/{id}", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.GetMedicationDispense)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/MedicationDispense", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.GetAllMedicationDispense)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMedicationDispense", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.CreateMedicationDispense)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/MedicationDispense/{id}", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.UpdateMedicationDispense)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMedicationDispense/{id}", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.DeleteMedicationDispense)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignMedicationOrderToMedicationDispense/{parentId}/medicationOrderId", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.AssignMedicationOrderToMedicationDispense)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMedicationOrderFromMedicationDispense/{parentId}", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.UnassignMedicationOrderFromMedicationDispense)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPharmacyToMedicationDispense/{parentId}/pharmacyId", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.AssignPharmacyToMedicationDispense)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPharmacyFromMedicationDispense/{parentId}", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.UnassignPharmacyFromMedicationDispense)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPatientToMedicationDispense/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.AssignPatientToMedicationDispense)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromMedicationDispense/{parentId}", jsonResponseFormatter.FormatToJSON(MedicationDispenseController.UnassignPatientFromMedicationDispense)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Diagnosis Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Diagnosis/{id}", jsonResponseFormatter.FormatToJSON(DiagnosisController.GetDiagnosis)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Diagnosis", jsonResponseFormatter.FormatToJSON(DiagnosisController.GetAllDiagnosis)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDiagnosis", jsonResponseFormatter.FormatToJSON(DiagnosisController.CreateDiagnosis)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Diagnosis/{id}", jsonResponseFormatter.FormatToJSON(DiagnosisController.UpdateDiagnosis)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDiagnosis/{id}", jsonResponseFormatter.FormatToJSON(DiagnosisController.DeleteDiagnosis)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEncounterToDiagnosis/{parentId}/encounterId", jsonResponseFormatter.FormatToJSON(DiagnosisController.AssignEncounterToDiagnosis)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEncounterFromDiagnosis/{parentId}", jsonResponseFormatter.FormatToJSON(DiagnosisController.UnassignEncounterFromDiagnosis)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPatientToDiagnosis/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(DiagnosisController.AssignPatientToDiagnosis)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromDiagnosis/{parentId}", jsonResponseFormatter.FormatToJSON(DiagnosisController.UnassignPatientFromDiagnosis)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Observation Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Observation/{id}", jsonResponseFormatter.FormatToJSON(ObservationController.GetObservation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Observation", jsonResponseFormatter.FormatToJSON(ObservationController.GetAllObservation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewObservation", jsonResponseFormatter.FormatToJSON(ObservationController.CreateObservation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Observation/{id}", jsonResponseFormatter.FormatToJSON(ObservationController.UpdateObservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteObservation/{id}", jsonResponseFormatter.FormatToJSON(ObservationController.DeleteObservation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEncounterToObservation/{parentId}/encounterId", jsonResponseFormatter.FormatToJSON(ObservationController.AssignEncounterToObservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEncounterFromObservation/{parentId}", jsonResponseFormatter.FormatToJSON(ObservationController.UnassignEncounterFromObservation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPatientToObservation/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(ObservationController.AssignPatientToObservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromObservation/{parentId}", jsonResponseFormatter.FormatToJSON(ObservationController.UnassignPatientFromObservation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDeviceToObservation/{parentId}/deviceId", jsonResponseFormatter.FormatToJSON(ObservationController.AssignDeviceToObservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDeviceFromObservation/{parentId}", jsonResponseFormatter.FormatToJSON(ObservationController.UnassignDeviceFromObservation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLabResultToObservation/{parentId}/labResultId", jsonResponseFormatter.FormatToJSON(ObservationController.AssignLabResultToObservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLabResultFromObservation/{parentId}", jsonResponseFormatter.FormatToJSON(ObservationController.UnassignLabResultFromObservation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // CarePlan Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CarePlan/{id}", jsonResponseFormatter.FormatToJSON(CarePlanController.GetCarePlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CarePlan", jsonResponseFormatter.FormatToJSON(CarePlanController.GetAllCarePlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCarePlan", jsonResponseFormatter.FormatToJSON(CarePlanController.CreateCarePlan)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CarePlan/{id}", jsonResponseFormatter.FormatToJSON(CarePlanController.UpdateCarePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCarePlan/{id}", jsonResponseFormatter.FormatToJSON(CarePlanController.DeleteCarePlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPatientToCarePlan/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(CarePlanController.AssignPatientToCarePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromCarePlan/{parentId}", jsonResponseFormatter.FormatToJSON(CarePlanController.UnassignPatientFromCarePlan)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCareTeamToCarePlan/{parentId}/careTeamId", jsonResponseFormatter.FormatToJSON(CarePlanController.AssignCareTeamToCarePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCareTeamFromCarePlan/{parentId}", jsonResponseFormatter.FormatToJSON(CarePlanController.UnassignCareTeamFromCarePlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddEncountersToCarePlan/{parentId}/encountersId", jsonResponseFormatter.FormatToJSON(CarePlanController.AddEncountersToCarePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEncountersFromCarePlan/{parentId}/encountersIds", jsonResponseFormatter.FormatToJSON(CarePlanController.RemoveEncountersFromCarePlan)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTasksToCarePlan/{parentId}/tasksId", jsonResponseFormatter.FormatToJSON(CarePlanController.AddTasksToCarePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTasksFromCarePlan/{parentId}/tasksIds", jsonResponseFormatter.FormatToJSON(CarePlanController.RemoveTasksFromCarePlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CareTask Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CareTask/{id}", jsonResponseFormatter.FormatToJSON(CareTaskController.GetCareTask)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CareTask", jsonResponseFormatter.FormatToJSON(CareTaskController.GetAllCareTask)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCareTask", jsonResponseFormatter.FormatToJSON(CareTaskController.CreateCareTask)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CareTask/{id}", jsonResponseFormatter.FormatToJSON(CareTaskController.UpdateCareTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCareTask/{id}", jsonResponseFormatter.FormatToJSON(CareTaskController.DeleteCareTask)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCarePlanToCareTask/{parentId}/carePlanId", jsonResponseFormatter.FormatToJSON(CareTaskController.AssignCarePlanToCareTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCarePlanFromCareTask/{parentId}", jsonResponseFormatter.FormatToJSON(CareTaskController.UnassignCarePlanFromCareTask)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAssignedToToCareTask/{parentId}/assignedToId", jsonResponseFormatter.FormatToJSON(CareTaskController.AssignAssignedToToCareTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAssignedToFromCareTask/{parentId}", jsonResponseFormatter.FormatToJSON(CareTaskController.UnassignAssignedToFromCareTask)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEncounterToCareTask/{parentId}/encounterId", jsonResponseFormatter.FormatToJSON(CareTaskController.AssignEncounterToCareTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEncounterFromCareTask/{parentId}", jsonResponseFormatter.FormatToJSON(CareTaskController.UnassignEncounterFromCareTask)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Allergy Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Allergy/{id}", jsonResponseFormatter.FormatToJSON(AllergyController.GetAllergy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Allergy", jsonResponseFormatter.FormatToJSON(AllergyController.GetAllAllergy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAllergy", jsonResponseFormatter.FormatToJSON(AllergyController.CreateAllergy)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Allergy/{id}", jsonResponseFormatter.FormatToJSON(AllergyController.UpdateAllergy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAllergy/{id}", jsonResponseFormatter.FormatToJSON(AllergyController.DeleteAllergy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPatientToAllergy/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(AllergyController.AssignPatientToAllergy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromAllergy/{parentId}", jsonResponseFormatter.FormatToJSON(AllergyController.UnassignPatientFromAllergy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Condition Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Condition/{id}", jsonResponseFormatter.FormatToJSON(ConditionController.GetCondition)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Condition", jsonResponseFormatter.FormatToJSON(ConditionController.GetAllCondition)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCondition", jsonResponseFormatter.FormatToJSON(ConditionController.CreateCondition)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Condition/{id}", jsonResponseFormatter.FormatToJSON(ConditionController.UpdateCondition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCondition/{id}", jsonResponseFormatter.FormatToJSON(ConditionController.DeleteCondition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPatientToCondition/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(ConditionController.AssignPatientToCondition)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromCondition/{parentId}", jsonResponseFormatter.FormatToJSON(ConditionController.UnassignPatientFromCondition)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // InsurancePayer Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InsurancePayer/{id}", jsonResponseFormatter.FormatToJSON(InsurancePayerController.GetInsurancePayer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InsurancePayer", jsonResponseFormatter.FormatToJSON(InsurancePayerController.GetAllInsurancePayer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInsurancePayer", jsonResponseFormatter.FormatToJSON(InsurancePayerController.CreateInsurancePayer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InsurancePayer/{id}", jsonResponseFormatter.FormatToJSON(InsurancePayerController.UpdateInsurancePayer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInsurancePayer/{id}", jsonResponseFormatter.FormatToJSON(InsurancePayerController.DeleteInsurancePayer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPlansToInsurancePayer/{parentId}/plansId", jsonResponseFormatter.FormatToJSON(InsurancePayerController.AddPlansToInsurancePayer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePlansFromInsurancePayer/{parentId}/plansIds", jsonResponseFormatter.FormatToJSON(InsurancePayerController.RemovePlansFromInsurancePayer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddClaimsToInsurancePayer/{parentId}/claimsId", jsonResponseFormatter.FormatToJSON(InsurancePayerController.AddClaimsToInsurancePayer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveClaimsFromInsurancePayer/{parentId}/claimsIds", jsonResponseFormatter.FormatToJSON(InsurancePayerController.RemoveClaimsFromInsurancePayer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InsurancePlan Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InsurancePlan/{id}", jsonResponseFormatter.FormatToJSON(InsurancePlanController.GetInsurancePlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InsurancePlan", jsonResponseFormatter.FormatToJSON(InsurancePlanController.GetAllInsurancePlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInsurancePlan", jsonResponseFormatter.FormatToJSON(InsurancePlanController.CreateInsurancePlan)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InsurancePlan/{id}", jsonResponseFormatter.FormatToJSON(InsurancePlanController.UpdateInsurancePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInsurancePlan/{id}", jsonResponseFormatter.FormatToJSON(InsurancePlanController.DeleteInsurancePlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPayerToInsurancePlan/{parentId}/payerId", jsonResponseFormatter.FormatToJSON(InsurancePlanController.AssignPayerToInsurancePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPayerFromInsurancePlan/{parentId}", jsonResponseFormatter.FormatToJSON(InsurancePlanController.UnassignPayerFromInsurancePlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCoveragesToInsurancePlan/{parentId}/coveragesId", jsonResponseFormatter.FormatToJSON(InsurancePlanController.AddCoveragesToInsurancePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCoveragesFromInsurancePlan/{parentId}/coveragesIds", jsonResponseFormatter.FormatToJSON(InsurancePlanController.RemoveCoveragesFromInsurancePlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Coverage Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Coverage/{id}", jsonResponseFormatter.FormatToJSON(CoverageController.GetCoverage)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Coverage", jsonResponseFormatter.FormatToJSON(CoverageController.GetAllCoverage)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCoverage", jsonResponseFormatter.FormatToJSON(CoverageController.CreateCoverage)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Coverage/{id}", jsonResponseFormatter.FormatToJSON(CoverageController.UpdateCoverage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCoverage/{id}", jsonResponseFormatter.FormatToJSON(CoverageController.DeleteCoverage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPatientToCoverage/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(CoverageController.AssignPatientToCoverage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromCoverage/{parentId}", jsonResponseFormatter.FormatToJSON(CoverageController.UnassignPatientFromCoverage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPlanToCoverage/{parentId}/planId", jsonResponseFormatter.FormatToJSON(CoverageController.AssignPlanToCoverage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlanFromCoverage/{parentId}", jsonResponseFormatter.FormatToJSON(CoverageController.UnassignPlanFromCoverage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddClaimsToCoverage/{parentId}/claimsId", jsonResponseFormatter.FormatToJSON(CoverageController.AddClaimsToCoverage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveClaimsFromCoverage/{parentId}/claimsIds", jsonResponseFormatter.FormatToJSON(CoverageController.RemoveClaimsFromCoverage)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAuthorizationsToCoverage/{parentId}/authorizationsId", jsonResponseFormatter.FormatToJSON(CoverageController.AddAuthorizationsToCoverage)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAuthorizationsFromCoverage/{parentId}/authorizationsIds", jsonResponseFormatter.FormatToJSON(CoverageController.RemoveAuthorizationsFromCoverage)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Claim Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Claim/{id}", jsonResponseFormatter.FormatToJSON(ClaimController.GetClaim)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Claim", jsonResponseFormatter.FormatToJSON(ClaimController.GetAllClaim)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewClaim", jsonResponseFormatter.FormatToJSON(ClaimController.CreateClaim)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Claim/{id}", jsonResponseFormatter.FormatToJSON(ClaimController.UpdateClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteClaim/{id}", jsonResponseFormatter.FormatToJSON(ClaimController.DeleteClaim)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPatientToClaim/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(ClaimController.AssignPatientToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromClaim/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimController.UnassignPatientFromClaim)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCoverageToClaim/{parentId}/coverageId", jsonResponseFormatter.FormatToJSON(ClaimController.AssignCoverageToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCoverageFromClaim/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimController.UnassignCoverageFromClaim)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEncounterToClaim/{parentId}/encounterId", jsonResponseFormatter.FormatToJSON(ClaimController.AssignEncounterToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEncounterFromClaim/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimController.UnassignEncounterFromClaim)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPayerToClaim/{parentId}/payerId", jsonResponseFormatter.FormatToJSON(ClaimController.AssignPayerToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPayerFromClaim/{parentId}", jsonResponseFormatter.FormatToJSON(ClaimController.UnassignPayerFromClaim)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInvoicesToClaim/{parentId}/invoicesId", jsonResponseFormatter.FormatToJSON(ClaimController.AddInvoicesToClaim)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInvoicesFromClaim/{parentId}/invoicesIds", jsonResponseFormatter.FormatToJSON(ClaimController.RemoveInvoicesFromClaim)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Authorization Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Authorization/{id}", jsonResponseFormatter.FormatToJSON(AuthorizationController.GetAuthorization)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Authorization", jsonResponseFormatter.FormatToJSON(AuthorizationController.GetAllAuthorization)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAuthorization", jsonResponseFormatter.FormatToJSON(AuthorizationController.CreateAuthorization)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Authorization/{id}", jsonResponseFormatter.FormatToJSON(AuthorizationController.UpdateAuthorization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAuthorization/{id}", jsonResponseFormatter.FormatToJSON(AuthorizationController.DeleteAuthorization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCoverageToAuthorization/{parentId}/coverageId", jsonResponseFormatter.FormatToJSON(AuthorizationController.AssignCoverageToAuthorization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCoverageFromAuthorization/{parentId}", jsonResponseFormatter.FormatToJSON(AuthorizationController.UnassignCoverageFromAuthorization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOrderToAuthorization/{parentId}/orderId", jsonResponseFormatter.FormatToJSON(AuthorizationController.AssignOrderToAuthorization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOrderFromAuthorization/{parentId}", jsonResponseFormatter.FormatToJSON(AuthorizationController.UnassignOrderFromAuthorization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Invoice Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Invoice/{id}", jsonResponseFormatter.FormatToJSON(InvoiceController.GetInvoice)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Invoice", jsonResponseFormatter.FormatToJSON(InvoiceController.GetAllInvoice)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInvoice", jsonResponseFormatter.FormatToJSON(InvoiceController.CreateInvoice)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Invoice/{id}", jsonResponseFormatter.FormatToJSON(InvoiceController.UpdateInvoice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInvoice/{id}", jsonResponseFormatter.FormatToJSON(InvoiceController.DeleteInvoice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPatientToInvoice/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(InvoiceController.AssignPatientToInvoice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromInvoice/{parentId}", jsonResponseFormatter.FormatToJSON(InvoiceController.UnassignPatientFromInvoice)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignClaimToInvoice/{parentId}/claimId", jsonResponseFormatter.FormatToJSON(InvoiceController.AssignClaimToInvoice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignClaimFromInvoice/{parentId}", jsonResponseFormatter.FormatToJSON(InvoiceController.UnassignClaimFromInvoice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPaymentsToInvoice/{parentId}/paymentsId", jsonResponseFormatter.FormatToJSON(InvoiceController.AddPaymentsToInvoice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePaymentsFromInvoice/{parentId}/paymentsIds", jsonResponseFormatter.FormatToJSON(InvoiceController.RemovePaymentsFromInvoice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Payment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Payment/{id}", jsonResponseFormatter.FormatToJSON(PaymentController.GetPayment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Payment", jsonResponseFormatter.FormatToJSON(PaymentController.GetAllPayment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPayment", jsonResponseFormatter.FormatToJSON(PaymentController.CreatePayment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Payment/{id}", jsonResponseFormatter.FormatToJSON(PaymentController.UpdatePayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePayment/{id}", jsonResponseFormatter.FormatToJSON(PaymentController.DeletePayment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInvoiceToPayment/{parentId}/invoiceId", jsonResponseFormatter.FormatToJSON(PaymentController.AssignInvoiceToPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInvoiceFromPayment/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentController.UnassignInvoiceFromPayment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPayerToPayment/{parentId}/payerId", jsonResponseFormatter.FormatToJSON(PaymentController.AssignPayerToPayment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPayerFromPayment/{parentId}", jsonResponseFormatter.FormatToJSON(PaymentController.UnassignPayerFromPayment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // MedicalDevice Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/MedicalDevice/{id}", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.GetMedicalDevice)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/MedicalDevice", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.GetAllMedicalDevice)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMedicalDevice", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.CreateMedicalDevice)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/MedicalDevice/{id}", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.UpdateMedicalDevice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMedicalDevice/{id}", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.DeleteMedicalDevice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPatientToMedicalDevice/{parentId}/patientId", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.AssignPatientToMedicalDevice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPatientFromMedicalDevice/{parentId}", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.UnassignPatientFromMedicalDevice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddObservationsToMedicalDevice/{parentId}/observationsId", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.AddObservationsToMedicalDevice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveObservationsFromMedicalDevice/{parentId}/observationsIds", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.RemoveObservationsFromMedicalDevice)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSoftwareUpdatesToMedicalDevice/{parentId}/softwareUpdatesId", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.AddSoftwareUpdatesToMedicalDevice)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSoftwareUpdatesFromMedicalDevice/{parentId}/softwareUpdatesIds", jsonResponseFormatter.FormatToJSON(MedicalDeviceController.RemoveSoftwareUpdatesFromMedicalDevice)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // SoftwareUpdate Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SoftwareUpdate/{id}", jsonResponseFormatter.FormatToJSON(SoftwareUpdateController.GetSoftwareUpdate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SoftwareUpdate", jsonResponseFormatter.FormatToJSON(SoftwareUpdateController.GetAllSoftwareUpdate)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSoftwareUpdate", jsonResponseFormatter.FormatToJSON(SoftwareUpdateController.CreateSoftwareUpdate)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SoftwareUpdate/{id}", jsonResponseFormatter.FormatToJSON(SoftwareUpdateController.UpdateSoftwareUpdate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSoftwareUpdate/{id}", jsonResponseFormatter.FormatToJSON(SoftwareUpdateController.DeleteSoftwareUpdate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignDeviceToSoftwareUpdate/{parentId}/deviceId", jsonResponseFormatter.FormatToJSON(SoftwareUpdateController.AssignDeviceToSoftwareUpdate)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDeviceFromSoftwareUpdate/{parentId}", jsonResponseFormatter.FormatToJSON(SoftwareUpdateController.UnassignDeviceFromSoftwareUpdate)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // MedicalSupplier Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/MedicalSupplier/{id}", jsonResponseFormatter.FormatToJSON(MedicalSupplierController.GetMedicalSupplier)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/MedicalSupplier", jsonResponseFormatter.FormatToJSON(MedicalSupplierController.GetAllMedicalSupplier)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMedicalSupplier", jsonResponseFormatter.FormatToJSON(MedicalSupplierController.CreateMedicalSupplier)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/MedicalSupplier/{id}", jsonResponseFormatter.FormatToJSON(MedicalSupplierController.UpdateMedicalSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMedicalSupplier/{id}", jsonResponseFormatter.FormatToJSON(MedicalSupplierController.DeleteMedicalSupplier)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddFacilitiesToMedicalSupplier/{parentId}/facilitiesId", jsonResponseFormatter.FormatToJSON(MedicalSupplierController.AddFacilitiesToMedicalSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFacilitiesFromMedicalSupplier/{parentId}/facilitiesIds", jsonResponseFormatter.FormatToJSON(MedicalSupplierController.RemoveFacilitiesFromMedicalSupplier)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInventoryItemsToMedicalSupplier/{parentId}/inventoryItemsId", jsonResponseFormatter.FormatToJSON(MedicalSupplierController.AddInventoryItemsToMedicalSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventoryItemsFromMedicalSupplier/{parentId}/inventoryItemsIds", jsonResponseFormatter.FormatToJSON(MedicalSupplierController.RemoveInventoryItemsFromMedicalSupplier)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InventoryItem Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InventoryItem/{id}", jsonResponseFormatter.FormatToJSON(InventoryItemController.GetInventoryItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InventoryItem", jsonResponseFormatter.FormatToJSON(InventoryItemController.GetAllInventoryItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInventoryItem", jsonResponseFormatter.FormatToJSON(InventoryItemController.CreateInventoryItem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InventoryItem/{id}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UpdateInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInventoryItem/{id}", jsonResponseFormatter.FormatToJSON(InventoryItemController.DeleteInventoryItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignFacilityToInventoryItem/{parentId}/facilityId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AssignFacilityToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFacilityFromInventoryItem/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UnassignFacilityFromInventoryItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSupplierToInventoryItem/{parentId}/supplierId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AssignSupplierToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSupplierFromInventoryItem/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UnassignSupplierFromInventoryItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    return router
}
