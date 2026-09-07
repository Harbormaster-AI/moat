// routerConfig.ts

import { Routes } from '@angular/router';
import { CreateHealthSystemComponent } from './components/HealthSystem/create/create.component';
import { EditHealthSystemComponent } from './components/HealthSystem/edit/edit.component';
import { IndexHealthSystemComponent } from './components/HealthSystem/index/index.component';
import { CreateFacilityComponent } from './components/Facility/create/create.component';
import { EditFacilityComponent } from './components/Facility/edit/edit.component';
import { IndexFacilityComponent } from './components/Facility/index/index.component';
import { CreateDepartmentComponent } from './components/Department/create/create.component';
import { EditDepartmentComponent } from './components/Department/edit/edit.component';
import { IndexDepartmentComponent } from './components/Department/index/index.component';
import { CreateCareTeamComponent } from './components/CareTeam/create/create.component';
import { EditCareTeamComponent } from './components/CareTeam/edit/edit.component';
import { IndexCareTeamComponent } from './components/CareTeam/index/index.component';
import { CreateClinicianComponent } from './components/Clinician/create/create.component';
import { EditClinicianComponent } from './components/Clinician/edit/edit.component';
import { IndexClinicianComponent } from './components/Clinician/index/index.component';
import { CreatePatientComponent } from './components/Patient/create/create.component';
import { EditPatientComponent } from './components/Patient/edit/edit.component';
import { IndexPatientComponent } from './components/Patient/index/index.component';
import { CreateAppointmentComponent } from './components/Appointment/create/create.component';
import { EditAppointmentComponent } from './components/Appointment/edit/edit.component';
import { IndexAppointmentComponent } from './components/Appointment/index/index.component';
import { CreateEncounterComponent } from './components/Encounter/create/create.component';
import { EditEncounterComponent } from './components/Encounter/edit/edit.component';
import { IndexEncounterComponent } from './components/Encounter/index/index.component';
import { CreateAdmissionComponent } from './components/Admission/create/create.component';
import { EditAdmissionComponent } from './components/Admission/edit/edit.component';
import { IndexAdmissionComponent } from './components/Admission/index/index.component';
import { CreateDischargeComponent } from './components/Discharge/create/create.component';
import { EditDischargeComponent } from './components/Discharge/edit/edit.component';
import { IndexDischargeComponent } from './components/Discharge/index/index.component';
import { CreateClinicalOrderComponent } from './components/ClinicalOrder/create/create.component';
import { EditClinicalOrderComponent } from './components/ClinicalOrder/edit/edit.component';
import { IndexClinicalOrderComponent } from './components/ClinicalOrder/index/index.component';
import { CreateMedicationOrderComponent } from './components/MedicationOrder/create/create.component';
import { EditMedicationOrderComponent } from './components/MedicationOrder/edit/edit.component';
import { IndexMedicationOrderComponent } from './components/MedicationOrder/index/index.component';
import { CreateLaboratoryComponent } from './components/Laboratory/create/create.component';
import { EditLaboratoryComponent } from './components/Laboratory/edit/edit.component';
import { IndexLaboratoryComponent } from './components/Laboratory/index/index.component';
import { CreateLaboratoryOrderComponent } from './components/LaboratoryOrder/create/create.component';
import { EditLaboratoryOrderComponent } from './components/LaboratoryOrder/edit/edit.component';
import { IndexLaboratoryOrderComponent } from './components/LaboratoryOrder/index/index.component';
import { CreateLabResultComponent } from './components/LabResult/create/create.component';
import { EditLabResultComponent } from './components/LabResult/edit/edit.component';
import { IndexLabResultComponent } from './components/LabResult/index/index.component';
import { CreateImagingCenterComponent } from './components/ImagingCenter/create/create.component';
import { EditImagingCenterComponent } from './components/ImagingCenter/edit/edit.component';
import { IndexImagingCenterComponent } from './components/ImagingCenter/index/index.component';
import { CreateImagingOrderComponent } from './components/ImagingOrder/create/create.component';
import { EditImagingOrderComponent } from './components/ImagingOrder/edit/edit.component';
import { IndexImagingOrderComponent } from './components/ImagingOrder/index/index.component';
import { CreateImagingReportComponent } from './components/ImagingReport/create/create.component';
import { EditImagingReportComponent } from './components/ImagingReport/edit/edit.component';
import { IndexImagingReportComponent } from './components/ImagingReport/index/index.component';
import { CreateProcedureOrderComponent } from './components/ProcedureOrder/create/create.component';
import { EditProcedureOrderComponent } from './components/ProcedureOrder/edit/edit.component';
import { IndexProcedureOrderComponent } from './components/ProcedureOrder/index/index.component';
import { CreateProcedureComponent } from './components/Procedure/create/create.component';
import { EditProcedureComponent } from './components/Procedure/edit/edit.component';
import { IndexProcedureComponent } from './components/Procedure/index/index.component';
import { CreatePharmacyComponent } from './components/Pharmacy/create/create.component';
import { EditPharmacyComponent } from './components/Pharmacy/edit/edit.component';
import { IndexPharmacyComponent } from './components/Pharmacy/index/index.component';
import { CreateMedicationDispenseComponent } from './components/MedicationDispense/create/create.component';
import { EditMedicationDispenseComponent } from './components/MedicationDispense/edit/edit.component';
import { IndexMedicationDispenseComponent } from './components/MedicationDispense/index/index.component';
import { CreateDiagnosisComponent } from './components/Diagnosis/create/create.component';
import { EditDiagnosisComponent } from './components/Diagnosis/edit/edit.component';
import { IndexDiagnosisComponent } from './components/Diagnosis/index/index.component';
import { CreateObservationComponent } from './components/Observation/create/create.component';
import { EditObservationComponent } from './components/Observation/edit/edit.component';
import { IndexObservationComponent } from './components/Observation/index/index.component';
import { CreateCarePlanComponent } from './components/CarePlan/create/create.component';
import { EditCarePlanComponent } from './components/CarePlan/edit/edit.component';
import { IndexCarePlanComponent } from './components/CarePlan/index/index.component';
import { CreateCareTaskComponent } from './components/CareTask/create/create.component';
import { EditCareTaskComponent } from './components/CareTask/edit/edit.component';
import { IndexCareTaskComponent } from './components/CareTask/index/index.component';
import { CreateAllergyComponent } from './components/Allergy/create/create.component';
import { EditAllergyComponent } from './components/Allergy/edit/edit.component';
import { IndexAllergyComponent } from './components/Allergy/index/index.component';
import { CreateConditionComponent } from './components/Condition/create/create.component';
import { EditConditionComponent } from './components/Condition/edit/edit.component';
import { IndexConditionComponent } from './components/Condition/index/index.component';
import { CreateInsurancePayerComponent } from './components/InsurancePayer/create/create.component';
import { EditInsurancePayerComponent } from './components/InsurancePayer/edit/edit.component';
import { IndexInsurancePayerComponent } from './components/InsurancePayer/index/index.component';
import { CreateInsurancePlanComponent } from './components/InsurancePlan/create/create.component';
import { EditInsurancePlanComponent } from './components/InsurancePlan/edit/edit.component';
import { IndexInsurancePlanComponent } from './components/InsurancePlan/index/index.component';
import { CreateCoverageComponent } from './components/Coverage/create/create.component';
import { EditCoverageComponent } from './components/Coverage/edit/edit.component';
import { IndexCoverageComponent } from './components/Coverage/index/index.component';
import { CreateClaimComponent } from './components/Claim/create/create.component';
import { EditClaimComponent } from './components/Claim/edit/edit.component';
import { IndexClaimComponent } from './components/Claim/index/index.component';
import { CreateAuthorizationComponent } from './components/Authorization/create/create.component';
import { EditAuthorizationComponent } from './components/Authorization/edit/edit.component';
import { IndexAuthorizationComponent } from './components/Authorization/index/index.component';
import { CreateInvoiceComponent } from './components/Invoice/create/create.component';
import { EditInvoiceComponent } from './components/Invoice/edit/edit.component';
import { IndexInvoiceComponent } from './components/Invoice/index/index.component';
import { CreatePaymentComponent } from './components/Payment/create/create.component';
import { EditPaymentComponent } from './components/Payment/edit/edit.component';
import { IndexPaymentComponent } from './components/Payment/index/index.component';
import { CreateMedicalDeviceComponent } from './components/MedicalDevice/create/create.component';
import { EditMedicalDeviceComponent } from './components/MedicalDevice/edit/edit.component';
import { IndexMedicalDeviceComponent } from './components/MedicalDevice/index/index.component';
import { CreateSoftwareUpdateComponent } from './components/SoftwareUpdate/create/create.component';
import { EditSoftwareUpdateComponent } from './components/SoftwareUpdate/edit/edit.component';
import { IndexSoftwareUpdateComponent } from './components/SoftwareUpdate/index/index.component';
import { CreateMedicalSupplierComponent } from './components/MedicalSupplier/create/create.component';
import { EditMedicalSupplierComponent } from './components/MedicalSupplier/edit/edit.component';
import { IndexMedicalSupplierComponent } from './components/MedicalSupplier/index/index.component';
import { CreateInventoryItemComponent } from './components/InventoryItem/create/create.component';
import { EditInventoryItemComponent } from './components/InventoryItem/edit/edit.component';
import { IndexInventoryItemComponent } from './components/InventoryItem/index/index.component';

export const HealthSystemRoutes: Routes = [
  { path: 'createHealthSystem',
    component: CreateHealthSystemComponent
  },
  {
    path: 'editHealthSystem/:id',
    component: EditHealthSystemComponent
  },
  { path: 'indexHealthSystem',
    component: IndexHealthSystemComponent
  }
];
export const FacilityRoutes: Routes = [
  { path: 'createFacility',
    component: CreateFacilityComponent
  },
  {
    path: 'editFacility/:id',
    component: EditFacilityComponent
  },
  { path: 'indexFacility',
    component: IndexFacilityComponent
  }
];
export const DepartmentRoutes: Routes = [
  { path: 'createDepartment',
    component: CreateDepartmentComponent
  },
  {
    path: 'editDepartment/:id',
    component: EditDepartmentComponent
  },
  { path: 'indexDepartment',
    component: IndexDepartmentComponent
  }
];
export const CareTeamRoutes: Routes = [
  { path: 'createCareTeam',
    component: CreateCareTeamComponent
  },
  {
    path: 'editCareTeam/:id',
    component: EditCareTeamComponent
  },
  { path: 'indexCareTeam',
    component: IndexCareTeamComponent
  }
];
export const ClinicianRoutes: Routes = [
  { path: 'createClinician',
    component: CreateClinicianComponent
  },
  {
    path: 'editClinician/:id',
    component: EditClinicianComponent
  },
  { path: 'indexClinician',
    component: IndexClinicianComponent
  }
];
export const PatientRoutes: Routes = [
  { path: 'createPatient',
    component: CreatePatientComponent
  },
  {
    path: 'editPatient/:id',
    component: EditPatientComponent
  },
  { path: 'indexPatient',
    component: IndexPatientComponent
  }
];
export const AppointmentRoutes: Routes = [
  { path: 'createAppointment',
    component: CreateAppointmentComponent
  },
  {
    path: 'editAppointment/:id',
    component: EditAppointmentComponent
  },
  { path: 'indexAppointment',
    component: IndexAppointmentComponent
  }
];
export const EncounterRoutes: Routes = [
  { path: 'createEncounter',
    component: CreateEncounterComponent
  },
  {
    path: 'editEncounter/:id',
    component: EditEncounterComponent
  },
  { path: 'indexEncounter',
    component: IndexEncounterComponent
  }
];
export const AdmissionRoutes: Routes = [
  { path: 'createAdmission',
    component: CreateAdmissionComponent
  },
  {
    path: 'editAdmission/:id',
    component: EditAdmissionComponent
  },
  { path: 'indexAdmission',
    component: IndexAdmissionComponent
  }
];
export const DischargeRoutes: Routes = [
  { path: 'createDischarge',
    component: CreateDischargeComponent
  },
  {
    path: 'editDischarge/:id',
    component: EditDischargeComponent
  },
  { path: 'indexDischarge',
    component: IndexDischargeComponent
  }
];
export const ClinicalOrderRoutes: Routes = [
  { path: 'createClinicalOrder',
    component: CreateClinicalOrderComponent
  },
  {
    path: 'editClinicalOrder/:id',
    component: EditClinicalOrderComponent
  },
  { path: 'indexClinicalOrder',
    component: IndexClinicalOrderComponent
  }
];
export const MedicationOrderRoutes: Routes = [
  { path: 'createMedicationOrder',
    component: CreateMedicationOrderComponent
  },
  {
    path: 'editMedicationOrder/:id',
    component: EditMedicationOrderComponent
  },
  { path: 'indexMedicationOrder',
    component: IndexMedicationOrderComponent
  }
];
export const LaboratoryRoutes: Routes = [
  { path: 'createLaboratory',
    component: CreateLaboratoryComponent
  },
  {
    path: 'editLaboratory/:id',
    component: EditLaboratoryComponent
  },
  { path: 'indexLaboratory',
    component: IndexLaboratoryComponent
  }
];
export const LaboratoryOrderRoutes: Routes = [
  { path: 'createLaboratoryOrder',
    component: CreateLaboratoryOrderComponent
  },
  {
    path: 'editLaboratoryOrder/:id',
    component: EditLaboratoryOrderComponent
  },
  { path: 'indexLaboratoryOrder',
    component: IndexLaboratoryOrderComponent
  }
];
export const LabResultRoutes: Routes = [
  { path: 'createLabResult',
    component: CreateLabResultComponent
  },
  {
    path: 'editLabResult/:id',
    component: EditLabResultComponent
  },
  { path: 'indexLabResult',
    component: IndexLabResultComponent
  }
];
export const ImagingCenterRoutes: Routes = [
  { path: 'createImagingCenter',
    component: CreateImagingCenterComponent
  },
  {
    path: 'editImagingCenter/:id',
    component: EditImagingCenterComponent
  },
  { path: 'indexImagingCenter',
    component: IndexImagingCenterComponent
  }
];
export const ImagingOrderRoutes: Routes = [
  { path: 'createImagingOrder',
    component: CreateImagingOrderComponent
  },
  {
    path: 'editImagingOrder/:id',
    component: EditImagingOrderComponent
  },
  { path: 'indexImagingOrder',
    component: IndexImagingOrderComponent
  }
];
export const ImagingReportRoutes: Routes = [
  { path: 'createImagingReport',
    component: CreateImagingReportComponent
  },
  {
    path: 'editImagingReport/:id',
    component: EditImagingReportComponent
  },
  { path: 'indexImagingReport',
    component: IndexImagingReportComponent
  }
];
export const ProcedureOrderRoutes: Routes = [
  { path: 'createProcedureOrder',
    component: CreateProcedureOrderComponent
  },
  {
    path: 'editProcedureOrder/:id',
    component: EditProcedureOrderComponent
  },
  { path: 'indexProcedureOrder',
    component: IndexProcedureOrderComponent
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
export const PharmacyRoutes: Routes = [
  { path: 'createPharmacy',
    component: CreatePharmacyComponent
  },
  {
    path: 'editPharmacy/:id',
    component: EditPharmacyComponent
  },
  { path: 'indexPharmacy',
    component: IndexPharmacyComponent
  }
];
export const MedicationDispenseRoutes: Routes = [
  { path: 'createMedicationDispense',
    component: CreateMedicationDispenseComponent
  },
  {
    path: 'editMedicationDispense/:id',
    component: EditMedicationDispenseComponent
  },
  { path: 'indexMedicationDispense',
    component: IndexMedicationDispenseComponent
  }
];
export const DiagnosisRoutes: Routes = [
  { path: 'createDiagnosis',
    component: CreateDiagnosisComponent
  },
  {
    path: 'editDiagnosis/:id',
    component: EditDiagnosisComponent
  },
  { path: 'indexDiagnosis',
    component: IndexDiagnosisComponent
  }
];
export const ObservationRoutes: Routes = [
  { path: 'createObservation',
    component: CreateObservationComponent
  },
  {
    path: 'editObservation/:id',
    component: EditObservationComponent
  },
  { path: 'indexObservation',
    component: IndexObservationComponent
  }
];
export const CarePlanRoutes: Routes = [
  { path: 'createCarePlan',
    component: CreateCarePlanComponent
  },
  {
    path: 'editCarePlan/:id',
    component: EditCarePlanComponent
  },
  { path: 'indexCarePlan',
    component: IndexCarePlanComponent
  }
];
export const CareTaskRoutes: Routes = [
  { path: 'createCareTask',
    component: CreateCareTaskComponent
  },
  {
    path: 'editCareTask/:id',
    component: EditCareTaskComponent
  },
  { path: 'indexCareTask',
    component: IndexCareTaskComponent
  }
];
export const AllergyRoutes: Routes = [
  { path: 'createAllergy',
    component: CreateAllergyComponent
  },
  {
    path: 'editAllergy/:id',
    component: EditAllergyComponent
  },
  { path: 'indexAllergy',
    component: IndexAllergyComponent
  }
];
export const ConditionRoutes: Routes = [
  { path: 'createCondition',
    component: CreateConditionComponent
  },
  {
    path: 'editCondition/:id',
    component: EditConditionComponent
  },
  { path: 'indexCondition',
    component: IndexConditionComponent
  }
];
export const InsurancePayerRoutes: Routes = [
  { path: 'createInsurancePayer',
    component: CreateInsurancePayerComponent
  },
  {
    path: 'editInsurancePayer/:id',
    component: EditInsurancePayerComponent
  },
  { path: 'indexInsurancePayer',
    component: IndexInsurancePayerComponent
  }
];
export const InsurancePlanRoutes: Routes = [
  { path: 'createInsurancePlan',
    component: CreateInsurancePlanComponent
  },
  {
    path: 'editInsurancePlan/:id',
    component: EditInsurancePlanComponent
  },
  { path: 'indexInsurancePlan',
    component: IndexInsurancePlanComponent
  }
];
export const CoverageRoutes: Routes = [
  { path: 'createCoverage',
    component: CreateCoverageComponent
  },
  {
    path: 'editCoverage/:id',
    component: EditCoverageComponent
  },
  { path: 'indexCoverage',
    component: IndexCoverageComponent
  }
];
export const ClaimRoutes: Routes = [
  { path: 'createClaim',
    component: CreateClaimComponent
  },
  {
    path: 'editClaim/:id',
    component: EditClaimComponent
  },
  { path: 'indexClaim',
    component: IndexClaimComponent
  }
];
export const AuthorizationRoutes: Routes = [
  { path: 'createAuthorization',
    component: CreateAuthorizationComponent
  },
  {
    path: 'editAuthorization/:id',
    component: EditAuthorizationComponent
  },
  { path: 'indexAuthorization',
    component: IndexAuthorizationComponent
  }
];
export const InvoiceRoutes: Routes = [
  { path: 'createInvoice',
    component: CreateInvoiceComponent
  },
  {
    path: 'editInvoice/:id',
    component: EditInvoiceComponent
  },
  { path: 'indexInvoice',
    component: IndexInvoiceComponent
  }
];
export const PaymentRoutes: Routes = [
  { path: 'createPayment',
    component: CreatePaymentComponent
  },
  {
    path: 'editPayment/:id',
    component: EditPaymentComponent
  },
  { path: 'indexPayment',
    component: IndexPaymentComponent
  }
];
export const MedicalDeviceRoutes: Routes = [
  { path: 'createMedicalDevice',
    component: CreateMedicalDeviceComponent
  },
  {
    path: 'editMedicalDevice/:id',
    component: EditMedicalDeviceComponent
  },
  { path: 'indexMedicalDevice',
    component: IndexMedicalDeviceComponent
  }
];
export const SoftwareUpdateRoutes: Routes = [
  { path: 'createSoftwareUpdate',
    component: CreateSoftwareUpdateComponent
  },
  {
    path: 'editSoftwareUpdate/:id',
    component: EditSoftwareUpdateComponent
  },
  { path: 'indexSoftwareUpdate',
    component: IndexSoftwareUpdateComponent
  }
];
export const MedicalSupplierRoutes: Routes = [
  { path: 'createMedicalSupplier',
    component: CreateMedicalSupplierComponent
  },
  {
    path: 'editMedicalSupplier/:id',
    component: EditMedicalSupplierComponent
  },
  { path: 'indexMedicalSupplier',
    component: IndexMedicalSupplierComponent
  }
];
export const InventoryItemRoutes: Routes = [
  { path: 'createInventoryItem',
    component: CreateInventoryItemComponent
  },
  {
    path: 'editInventoryItem/:id',
    component: EditInventoryItemComponent
  },
  { path: 'indexInventoryItem',
    component: IndexInventoryItemComponent
  }
];
