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

import {IndexHealthSystemComponent} from './components/HealthSystem/index/index.component';
import {CreateHealthSystemComponent} from './components/HealthSystem/create/create.component';
import {EditHealthSystemComponent} from './components/HealthSystem/edit/edit.component';
import {IndexFacilityComponent} from './components/Facility/index/index.component';
import {CreateFacilityComponent} from './components/Facility/create/create.component';
import {EditFacilityComponent} from './components/Facility/edit/edit.component';
import {IndexDepartmentComponent} from './components/Department/index/index.component';
import {CreateDepartmentComponent} from './components/Department/create/create.component';
import {EditDepartmentComponent} from './components/Department/edit/edit.component';
import {IndexCareTeamComponent} from './components/CareTeam/index/index.component';
import {CreateCareTeamComponent} from './components/CareTeam/create/create.component';
import {EditCareTeamComponent} from './components/CareTeam/edit/edit.component';
import {IndexClinicianComponent} from './components/Clinician/index/index.component';
import {CreateClinicianComponent} from './components/Clinician/create/create.component';
import {EditClinicianComponent} from './components/Clinician/edit/edit.component';
import {IndexPatientComponent} from './components/Patient/index/index.component';
import {CreatePatientComponent} from './components/Patient/create/create.component';
import {EditPatientComponent} from './components/Patient/edit/edit.component';
import {IndexAppointmentComponent} from './components/Appointment/index/index.component';
import {CreateAppointmentComponent} from './components/Appointment/create/create.component';
import {EditAppointmentComponent} from './components/Appointment/edit/edit.component';
import {IndexEncounterComponent} from './components/Encounter/index/index.component';
import {CreateEncounterComponent} from './components/Encounter/create/create.component';
import {EditEncounterComponent} from './components/Encounter/edit/edit.component';
import {IndexAdmissionComponent} from './components/Admission/index/index.component';
import {CreateAdmissionComponent} from './components/Admission/create/create.component';
import {EditAdmissionComponent} from './components/Admission/edit/edit.component';
import {IndexDischargeComponent} from './components/Discharge/index/index.component';
import {CreateDischargeComponent} from './components/Discharge/create/create.component';
import {EditDischargeComponent} from './components/Discharge/edit/edit.component';
import {IndexClinicalOrderComponent} from './components/ClinicalOrder/index/index.component';
import {CreateClinicalOrderComponent} from './components/ClinicalOrder/create/create.component';
import {EditClinicalOrderComponent} from './components/ClinicalOrder/edit/edit.component';
import {IndexMedicationOrderComponent} from './components/MedicationOrder/index/index.component';
import {CreateMedicationOrderComponent} from './components/MedicationOrder/create/create.component';
import {EditMedicationOrderComponent} from './components/MedicationOrder/edit/edit.component';
import {IndexLaboratoryComponent} from './components/Laboratory/index/index.component';
import {CreateLaboratoryComponent} from './components/Laboratory/create/create.component';
import {EditLaboratoryComponent} from './components/Laboratory/edit/edit.component';
import {IndexLaboratoryOrderComponent} from './components/LaboratoryOrder/index/index.component';
import {CreateLaboratoryOrderComponent} from './components/LaboratoryOrder/create/create.component';
import {EditLaboratoryOrderComponent} from './components/LaboratoryOrder/edit/edit.component';
import {IndexLabResultComponent} from './components/LabResult/index/index.component';
import {CreateLabResultComponent} from './components/LabResult/create/create.component';
import {EditLabResultComponent} from './components/LabResult/edit/edit.component';
import {IndexImagingCenterComponent} from './components/ImagingCenter/index/index.component';
import {CreateImagingCenterComponent} from './components/ImagingCenter/create/create.component';
import {EditImagingCenterComponent} from './components/ImagingCenter/edit/edit.component';
import {IndexImagingOrderComponent} from './components/ImagingOrder/index/index.component';
import {CreateImagingOrderComponent} from './components/ImagingOrder/create/create.component';
import {EditImagingOrderComponent} from './components/ImagingOrder/edit/edit.component';
import {IndexImagingReportComponent} from './components/ImagingReport/index/index.component';
import {CreateImagingReportComponent} from './components/ImagingReport/create/create.component';
import {EditImagingReportComponent} from './components/ImagingReport/edit/edit.component';
import {IndexProcedureOrderComponent} from './components/ProcedureOrder/index/index.component';
import {CreateProcedureOrderComponent} from './components/ProcedureOrder/create/create.component';
import {EditProcedureOrderComponent} from './components/ProcedureOrder/edit/edit.component';
import {IndexProcedureComponent} from './components/Procedure/index/index.component';
import {CreateProcedureComponent} from './components/Procedure/create/create.component';
import {EditProcedureComponent} from './components/Procedure/edit/edit.component';
import {IndexPharmacyComponent} from './components/Pharmacy/index/index.component';
import {CreatePharmacyComponent} from './components/Pharmacy/create/create.component';
import {EditPharmacyComponent} from './components/Pharmacy/edit/edit.component';
import {IndexMedicationDispenseComponent} from './components/MedicationDispense/index/index.component';
import {CreateMedicationDispenseComponent} from './components/MedicationDispense/create/create.component';
import {EditMedicationDispenseComponent} from './components/MedicationDispense/edit/edit.component';
import {IndexDiagnosisComponent} from './components/Diagnosis/index/index.component';
import {CreateDiagnosisComponent} from './components/Diagnosis/create/create.component';
import {EditDiagnosisComponent} from './components/Diagnosis/edit/edit.component';
import {IndexObservationComponent} from './components/Observation/index/index.component';
import {CreateObservationComponent} from './components/Observation/create/create.component';
import {EditObservationComponent} from './components/Observation/edit/edit.component';
import {IndexCarePlanComponent} from './components/CarePlan/index/index.component';
import {CreateCarePlanComponent} from './components/CarePlan/create/create.component';
import {EditCarePlanComponent} from './components/CarePlan/edit/edit.component';
import {IndexCareTaskComponent} from './components/CareTask/index/index.component';
import {CreateCareTaskComponent} from './components/CareTask/create/create.component';
import {EditCareTaskComponent} from './components/CareTask/edit/edit.component';
import {IndexAllergyComponent} from './components/Allergy/index/index.component';
import {CreateAllergyComponent} from './components/Allergy/create/create.component';
import {EditAllergyComponent} from './components/Allergy/edit/edit.component';
import {IndexConditionComponent} from './components/Condition/index/index.component';
import {CreateConditionComponent} from './components/Condition/create/create.component';
import {EditConditionComponent} from './components/Condition/edit/edit.component';
import {IndexInsurancePayerComponent} from './components/InsurancePayer/index/index.component';
import {CreateInsurancePayerComponent} from './components/InsurancePayer/create/create.component';
import {EditInsurancePayerComponent} from './components/InsurancePayer/edit/edit.component';
import {IndexInsurancePlanComponent} from './components/InsurancePlan/index/index.component';
import {CreateInsurancePlanComponent} from './components/InsurancePlan/create/create.component';
import {EditInsurancePlanComponent} from './components/InsurancePlan/edit/edit.component';
import {IndexCoverageComponent} from './components/Coverage/index/index.component';
import {CreateCoverageComponent} from './components/Coverage/create/create.component';
import {EditCoverageComponent} from './components/Coverage/edit/edit.component';
import {IndexClaimComponent} from './components/Claim/index/index.component';
import {CreateClaimComponent} from './components/Claim/create/create.component';
import {EditClaimComponent} from './components/Claim/edit/edit.component';
import {IndexAuthorizationComponent} from './components/Authorization/index/index.component';
import {CreateAuthorizationComponent} from './components/Authorization/create/create.component';
import {EditAuthorizationComponent} from './components/Authorization/edit/edit.component';
import {IndexInvoiceComponent} from './components/Invoice/index/index.component';
import {CreateInvoiceComponent} from './components/Invoice/create/create.component';
import {EditInvoiceComponent} from './components/Invoice/edit/edit.component';
import {IndexPaymentComponent} from './components/Payment/index/index.component';
import {CreatePaymentComponent} from './components/Payment/create/create.component';
import {EditPaymentComponent} from './components/Payment/edit/edit.component';
import {IndexMedicalDeviceComponent} from './components/MedicalDevice/index/index.component';
import {CreateMedicalDeviceComponent} from './components/MedicalDevice/create/create.component';
import {EditMedicalDeviceComponent} from './components/MedicalDevice/edit/edit.component';
import {IndexSoftwareUpdateComponent} from './components/SoftwareUpdate/index/index.component';
import {CreateSoftwareUpdateComponent} from './components/SoftwareUpdate/create/create.component';
import {EditSoftwareUpdateComponent} from './components/SoftwareUpdate/edit/edit.component';
import {IndexMedicalSupplierComponent} from './components/MedicalSupplier/index/index.component';
import {CreateMedicalSupplierComponent} from './components/MedicalSupplier/create/create.component';
import {EditMedicalSupplierComponent} from './components/MedicalSupplier/edit/edit.component';
import {IndexInventoryItemComponent} from './components/InventoryItem/index/index.component';
import {CreateInventoryItemComponent} from './components/InventoryItem/create/create.component';
import {EditInventoryItemComponent} from './components/InventoryItem/edit/edit.component';

import * as appRoutes from './routerConfig';

import {HealthSystemService} from './services/HealthSystem.service';
import {FacilityService} from './services/Facility.service';
import {DepartmentService} from './services/Department.service';
import {CareTeamService} from './services/CareTeam.service';
import {ClinicianService} from './services/Clinician.service';
import {PatientService} from './services/Patient.service';
import {AppointmentService} from './services/Appointment.service';
import {EncounterService} from './services/Encounter.service';
import {AdmissionService} from './services/Admission.service';
import {DischargeService} from './services/Discharge.service';
import {ClinicalOrderService} from './services/ClinicalOrder.service';
import {MedicationOrderService} from './services/MedicationOrder.service';
import {LaboratoryService} from './services/Laboratory.service';
import {LaboratoryOrderService} from './services/LaboratoryOrder.service';
import {LabResultService} from './services/LabResult.service';
import {ImagingCenterService} from './services/ImagingCenter.service';
import {ImagingOrderService} from './services/ImagingOrder.service';
import {ImagingReportService} from './services/ImagingReport.service';
import {ProcedureOrderService} from './services/ProcedureOrder.service';
import {ProcedureService} from './services/Procedure.service';
import {PharmacyService} from './services/Pharmacy.service';
import {MedicationDispenseService} from './services/MedicationDispense.service';
import {DiagnosisService} from './services/Diagnosis.service';
import {ObservationService} from './services/Observation.service';
import {CarePlanService} from './services/CarePlan.service';
import {CareTaskService} from './services/CareTask.service';
import {AllergyService} from './services/Allergy.service';
import {ConditionService} from './services/Condition.service';
import {InsurancePayerService} from './services/InsurancePayer.service';
import {InsurancePlanService} from './services/InsurancePlan.service';
import {CoverageService} from './services/Coverage.service';
import {ClaimService} from './services/Claim.service';
import {AuthorizationService} from './services/Authorization.service';
import {InvoiceService} from './services/Invoice.service';
import {PaymentService} from './services/Payment.service';
import {MedicalDeviceService} from './services/MedicalDevice.service';
import {SoftwareUpdateService} from './services/SoftwareUpdate.service';
import {MedicalSupplierService} from './services/MedicalSupplier.service';
import {InventoryItemService} from './services/InventoryItem.service';

@NgModule({
  declarations: [
    IndexHealthSystemComponent,
    CreateHealthSystemComponent,
    EditHealthSystemComponent,
    IndexFacilityComponent,
    CreateFacilityComponent,
    EditFacilityComponent,
    IndexDepartmentComponent,
    CreateDepartmentComponent,
    EditDepartmentComponent,
    IndexCareTeamComponent,
    CreateCareTeamComponent,
    EditCareTeamComponent,
    IndexClinicianComponent,
    CreateClinicianComponent,
    EditClinicianComponent,
    IndexPatientComponent,
    CreatePatientComponent,
    EditPatientComponent,
    IndexAppointmentComponent,
    CreateAppointmentComponent,
    EditAppointmentComponent,
    IndexEncounterComponent,
    CreateEncounterComponent,
    EditEncounterComponent,
    IndexAdmissionComponent,
    CreateAdmissionComponent,
    EditAdmissionComponent,
    IndexDischargeComponent,
    CreateDischargeComponent,
    EditDischargeComponent,
    IndexClinicalOrderComponent,
    CreateClinicalOrderComponent,
    EditClinicalOrderComponent,
    IndexMedicationOrderComponent,
    CreateMedicationOrderComponent,
    EditMedicationOrderComponent,
    IndexLaboratoryComponent,
    CreateLaboratoryComponent,
    EditLaboratoryComponent,
    IndexLaboratoryOrderComponent,
    CreateLaboratoryOrderComponent,
    EditLaboratoryOrderComponent,
    IndexLabResultComponent,
    CreateLabResultComponent,
    EditLabResultComponent,
    IndexImagingCenterComponent,
    CreateImagingCenterComponent,
    EditImagingCenterComponent,
    IndexImagingOrderComponent,
    CreateImagingOrderComponent,
    EditImagingOrderComponent,
    IndexImagingReportComponent,
    CreateImagingReportComponent,
    EditImagingReportComponent,
    IndexProcedureOrderComponent,
    CreateProcedureOrderComponent,
    EditProcedureOrderComponent,
    IndexProcedureComponent,
    CreateProcedureComponent,
    EditProcedureComponent,
    IndexPharmacyComponent,
    CreatePharmacyComponent,
    EditPharmacyComponent,
    IndexMedicationDispenseComponent,
    CreateMedicationDispenseComponent,
    EditMedicationDispenseComponent,
    IndexDiagnosisComponent,
    CreateDiagnosisComponent,
    EditDiagnosisComponent,
    IndexObservationComponent,
    CreateObservationComponent,
    EditObservationComponent,
    IndexCarePlanComponent,
    CreateCarePlanComponent,
    EditCarePlanComponent,
    IndexCareTaskComponent,
    CreateCareTaskComponent,
    EditCareTaskComponent,
    IndexAllergyComponent,
    CreateAllergyComponent,
    EditAllergyComponent,
    IndexConditionComponent,
    CreateConditionComponent,
    EditConditionComponent,
    IndexInsurancePayerComponent,
    CreateInsurancePayerComponent,
    EditInsurancePayerComponent,
    IndexInsurancePlanComponent,
    CreateInsurancePlanComponent,
    EditInsurancePlanComponent,
    IndexCoverageComponent,
    CreateCoverageComponent,
    EditCoverageComponent,
    IndexClaimComponent,
    CreateClaimComponent,
    EditClaimComponent,
    IndexAuthorizationComponent,
    CreateAuthorizationComponent,
    EditAuthorizationComponent,
    IndexInvoiceComponent,
    CreateInvoiceComponent,
    EditInvoiceComponent,
    IndexPaymentComponent,
    CreatePaymentComponent,
    EditPaymentComponent,
    IndexMedicalDeviceComponent,
    CreateMedicalDeviceComponent,
    EditMedicalDeviceComponent,
    IndexSoftwareUpdateComponent,
    CreateSoftwareUpdateComponent,
    EditSoftwareUpdateComponent,
    IndexMedicalSupplierComponent,
    CreateMedicalSupplierComponent,
    EditMedicalSupplierComponent,
    IndexInventoryItemComponent,
    CreateInventoryItemComponent,
    EditInventoryItemComponent,
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
    RouterModule.forRoot(appRoutes.HealthSystemRoutes), 
    RouterModule.forRoot(appRoutes.FacilityRoutes), 
    RouterModule.forRoot(appRoutes.DepartmentRoutes), 
    RouterModule.forRoot(appRoutes.CareTeamRoutes), 
    RouterModule.forRoot(appRoutes.ClinicianRoutes), 
    RouterModule.forRoot(appRoutes.PatientRoutes), 
    RouterModule.forRoot(appRoutes.AppointmentRoutes), 
    RouterModule.forRoot(appRoutes.EncounterRoutes), 
    RouterModule.forRoot(appRoutes.AdmissionRoutes), 
    RouterModule.forRoot(appRoutes.DischargeRoutes), 
    RouterModule.forRoot(appRoutes.ClinicalOrderRoutes), 
    RouterModule.forRoot(appRoutes.MedicationOrderRoutes), 
    RouterModule.forRoot(appRoutes.LaboratoryRoutes), 
    RouterModule.forRoot(appRoutes.LaboratoryOrderRoutes), 
    RouterModule.forRoot(appRoutes.LabResultRoutes), 
    RouterModule.forRoot(appRoutes.ImagingCenterRoutes), 
    RouterModule.forRoot(appRoutes.ImagingOrderRoutes), 
    RouterModule.forRoot(appRoutes.ImagingReportRoutes), 
    RouterModule.forRoot(appRoutes.ProcedureOrderRoutes), 
    RouterModule.forRoot(appRoutes.ProcedureRoutes), 
    RouterModule.forRoot(appRoutes.PharmacyRoutes), 
    RouterModule.forRoot(appRoutes.MedicationDispenseRoutes), 
    RouterModule.forRoot(appRoutes.DiagnosisRoutes), 
    RouterModule.forRoot(appRoutes.ObservationRoutes), 
    RouterModule.forRoot(appRoutes.CarePlanRoutes), 
    RouterModule.forRoot(appRoutes.CareTaskRoutes), 
    RouterModule.forRoot(appRoutes.AllergyRoutes), 
    RouterModule.forRoot(appRoutes.ConditionRoutes), 
    RouterModule.forRoot(appRoutes.InsurancePayerRoutes), 
    RouterModule.forRoot(appRoutes.InsurancePlanRoutes), 
    RouterModule.forRoot(appRoutes.CoverageRoutes), 
    RouterModule.forRoot(appRoutes.ClaimRoutes), 
    RouterModule.forRoot(appRoutes.AuthorizationRoutes), 
    RouterModule.forRoot(appRoutes.InvoiceRoutes), 
    RouterModule.forRoot(appRoutes.PaymentRoutes), 
    RouterModule.forRoot(appRoutes.MedicalDeviceRoutes), 
    RouterModule.forRoot(appRoutes.SoftwareUpdateRoutes), 
    RouterModule.forRoot(appRoutes.MedicalSupplierRoutes), 
    RouterModule.forRoot(appRoutes.InventoryItemRoutes), 
  ],
  providers: [HealthSystemService,FacilityService,DepartmentService,CareTeamService,ClinicianService,PatientService,AppointmentService,EncounterService,AdmissionService,DischargeService,ClinicalOrderService,MedicationOrderService,LaboratoryService,LaboratoryOrderService,LabResultService,ImagingCenterService,ImagingOrderService,ImagingReportService,ProcedureOrderService,ProcedureService,PharmacyService,MedicationDispenseService,DiagnosisService,ObservationService,CarePlanService,CareTaskService,AllergyService,ConditionService,InsurancePayerService,InsurancePlanService,CoverageService,ClaimService,AuthorizationService,InvoiceService,PaymentService,MedicalDeviceService,SoftwareUpdateService,MedicalSupplierService,InventoryItemService],
  bootstrap: [AppComponent]
})
export class AppModule { }
