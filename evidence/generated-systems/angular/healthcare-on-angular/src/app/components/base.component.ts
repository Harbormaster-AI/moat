import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {HealthSystemService} from '../services/HealthSystem.service';
import {FacilityService} from '../services/Facility.service';
import {DepartmentService} from '../services/Department.service';
import {CareTeamService} from '../services/CareTeam.service';
import {ClinicianService} from '../services/Clinician.service';
import {PatientService} from '../services/Patient.service';
import {AppointmentService} from '../services/Appointment.service';
import {EncounterService} from '../services/Encounter.service';
import {AdmissionService} from '../services/Admission.service';
import {DischargeService} from '../services/Discharge.service';
import {ClinicalOrderService} from '../services/ClinicalOrder.service';
import {MedicationOrderService} from '../services/MedicationOrder.service';
import {LaboratoryService} from '../services/Laboratory.service';
import {LaboratoryOrderService} from '../services/LaboratoryOrder.service';
import {LabResultService} from '../services/LabResult.service';
import {ImagingCenterService} from '../services/ImagingCenter.service';
import {ImagingOrderService} from '../services/ImagingOrder.service';
import {ImagingReportService} from '../services/ImagingReport.service';
import {ProcedureOrderService} from '../services/ProcedureOrder.service';
import {ProcedureService} from '../services/Procedure.service';
import {PharmacyService} from '../services/Pharmacy.service';
import {MedicationDispenseService} from '../services/MedicationDispense.service';
import {DiagnosisService} from '../services/Diagnosis.service';
import {ObservationService} from '../services/Observation.service';
import {CarePlanService} from '../services/CarePlan.service';
import {CareTaskService} from '../services/CareTask.service';
import {AllergyService} from '../services/Allergy.service';
import {ConditionService} from '../services/Condition.service';
import {InsurancePayerService} from '../services/InsurancePayer.service';
import {InsurancePlanService} from '../services/InsurancePlan.service';
import {CoverageService} from '../services/Coverage.service';
import {ClaimService} from '../services/Claim.service';
import {AuthorizationService} from '../services/Authorization.service';
import {InvoiceService} from '../services/Invoice.service';
import {PaymentService} from '../services/Payment.service';
import {MedicalDeviceService} from '../services/MedicalDevice.service';
import {SoftwareUpdateService} from '../services/SoftwareUpdate.service';
import {MedicalSupplierService} from '../services/MedicalSupplier.service';
import {InventoryItemService} from '../services/InventoryItem.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    FacilityTypes = Object.keys(enumTypes.FacilityType);
    DepartmentTypes = Object.keys(enumTypes.DepartmentType);
    CareSettingTypes = Object.keys(enumTypes.CareSettingType);
    ClinicianTypes = Object.keys(enumTypes.ClinicianType);
    ClinicianSpecialtys = Object.keys(enumTypes.ClinicianSpecialty);
    AdministrativeSexs = Object.keys(enumTypes.AdministrativeSex);
    BloodTypes = Object.keys(enumTypes.BloodType);
    AppointmentStatuss = Object.keys(enumTypes.AppointmentStatus);
    Prioritys = Object.keys(enumTypes.Priority);
    EncounterStatuss = Object.keys(enumTypes.EncounterStatus);
    EncounterTypes = Object.keys(enumTypes.EncounterType);
    AdmissionTypes = Object.keys(enumTypes.AdmissionType);
    DischargeDispositions = Object.keys(enumTypes.DischargeDisposition);
    OrderStatuss = Object.keys(enumTypes.OrderStatus);
    ClinicalOrderTypes = Object.keys(enumTypes.ClinicalOrderType);
    RouteOfAdministrations = Object.keys(enumTypes.RouteOfAdministration);
    SpecimenTypes = Object.keys(enumTypes.SpecimenType);
    ImagingModalitys = Object.keys(enumTypes.ImagingModality);
    AnesthesiaTypes = Object.keys(enumTypes.AnesthesiaType);
    DispenseStatuss = Object.keys(enumTypes.DispenseStatus);
    ResultStatuss = Object.keys(enumTypes.ResultStatus);
    ProcedureStatuss = Object.keys(enumTypes.ProcedureStatus);
    DiagnosisCertaintys = Object.keys(enumTypes.DiagnosisCertainty);
    ObservationInterpretations = Object.keys(enumTypes.ObservationInterpretation);
    CarePlanStatuss = Object.keys(enumTypes.CarePlanStatus);
    TaskStatuss = Object.keys(enumTypes.TaskStatus);
    AllergySeveritys = Object.keys(enumTypes.AllergySeverity);
    AllergyStatuss = Object.keys(enumTypes.AllergyStatus);
    ConditionStatuss = Object.keys(enumTypes.ConditionStatus);
    PayerTypes = Object.keys(enumTypes.PayerType);
    InsurancePlanTypes = Object.keys(enumTypes.InsurancePlanType);
    CoverageTypes = Object.keys(enumTypes.CoverageType);
    ClaimStatuss = Object.keys(enumTypes.ClaimStatus);
    AuthorizationStatuss = Object.keys(enumTypes.AuthorizationStatus);
    InvoiceStatuss = Object.keys(enumTypes.InvoiceStatus);
    PaymentMethods = Object.keys(enumTypes.PaymentMethod);
    DeviceTypes = Object.keys(enumTypes.DeviceType);
    DeviceConnectivityStatuss = Object.keys(enumTypes.DeviceConnectivityStatus);
    SoftwareUpdateTypes = Object.keys(enumTypes.SoftwareUpdateType);
    SupplierTiers = Object.keys(enumTypes.SupplierTier);

// all collection instances
    healthSystems : any;
    facilitys : any;
    departments : any;
    careTeams : any;
    clinicians : any;
    patients : any;
    appointments : any;
    encounters : any;
    admissions : any;
    discharges : any;
    clinicalOrders : any;
    medicationOrders : any;
    laboratorys : any;
    laboratoryOrders : any;
    labResults : any;
    imagingCenters : any;
    imagingOrders : any;
    imagingReports : any;
    procedureOrders : any;
    procedures : any;
    pharmacys : any;
    medicationDispenses : any;
    diagnosiss : any;
    observations : any;
    carePlans : any;
    careTasks : any;
    allergys : any;
    conditions : any;
    insurancePayers : any;
    insurancePlans : any;
    coverages : any;
    claims : any;
    authorizations : any;
    invoices : any;
    payments : any;
    medicalDevices : any;
    softwareUpdates : any;
    medicalSuppliers : any;
    inventoryItems : any;
  
// initialization  
    ngOnInit() {
    }

    initHealthSystemList() {
        if ( this.healthSystems == null ) {
            new HealthSystemService(this.http).getHealthSystems().subscribe(res => {
                this.healthSystems = res;
            });
        }
    }
    
    initFacilityList() {
        if ( this.facilitys == null ) {
            new FacilityService(this.http).getFacilitys().subscribe(res => {
                this.facilitys = res;
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
    
    initCareTeamList() {
        if ( this.careTeams == null ) {
            new CareTeamService(this.http).getCareTeams().subscribe(res => {
                this.careTeams = res;
            });
        }
    }
    
    initClinicianList() {
        if ( this.clinicians == null ) {
            new ClinicianService(this.http).getClinicians().subscribe(res => {
                this.clinicians = res;
            });
        }
    }
    
    initPatientList() {
        if ( this.patients == null ) {
            new PatientService(this.http).getPatients().subscribe(res => {
                this.patients = res;
            });
        }
    }
    
    initAppointmentList() {
        if ( this.appointments == null ) {
            new AppointmentService(this.http).getAppointments().subscribe(res => {
                this.appointments = res;
            });
        }
    }
    
    initEncounterList() {
        if ( this.encounters == null ) {
            new EncounterService(this.http).getEncounters().subscribe(res => {
                this.encounters = res;
            });
        }
    }
    
    initAdmissionList() {
        if ( this.admissions == null ) {
            new AdmissionService(this.http).getAdmissions().subscribe(res => {
                this.admissions = res;
            });
        }
    }
    
    initDischargeList() {
        if ( this.discharges == null ) {
            new DischargeService(this.http).getDischarges().subscribe(res => {
                this.discharges = res;
            });
        }
    }
    
    initClinicalOrderList() {
        if ( this.clinicalOrders == null ) {
            new ClinicalOrderService(this.http).getClinicalOrders().subscribe(res => {
                this.clinicalOrders = res;
            });
        }
    }
    
    initMedicationOrderList() {
        if ( this.medicationOrders == null ) {
            new MedicationOrderService(this.http).getMedicationOrders().subscribe(res => {
                this.medicationOrders = res;
            });
        }
    }
    
    initLaboratoryList() {
        if ( this.laboratorys == null ) {
            new LaboratoryService(this.http).getLaboratorys().subscribe(res => {
                this.laboratorys = res;
            });
        }
    }
    
    initLaboratoryOrderList() {
        if ( this.laboratoryOrders == null ) {
            new LaboratoryOrderService(this.http).getLaboratoryOrders().subscribe(res => {
                this.laboratoryOrders = res;
            });
        }
    }
    
    initLabResultList() {
        if ( this.labResults == null ) {
            new LabResultService(this.http).getLabResults().subscribe(res => {
                this.labResults = res;
            });
        }
    }
    
    initImagingCenterList() {
        if ( this.imagingCenters == null ) {
            new ImagingCenterService(this.http).getImagingCenters().subscribe(res => {
                this.imagingCenters = res;
            });
        }
    }
    
    initImagingOrderList() {
        if ( this.imagingOrders == null ) {
            new ImagingOrderService(this.http).getImagingOrders().subscribe(res => {
                this.imagingOrders = res;
            });
        }
    }
    
    initImagingReportList() {
        if ( this.imagingReports == null ) {
            new ImagingReportService(this.http).getImagingReports().subscribe(res => {
                this.imagingReports = res;
            });
        }
    }
    
    initProcedureOrderList() {
        if ( this.procedureOrders == null ) {
            new ProcedureOrderService(this.http).getProcedureOrders().subscribe(res => {
                this.procedureOrders = res;
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
    
    initPharmacyList() {
        if ( this.pharmacys == null ) {
            new PharmacyService(this.http).getPharmacys().subscribe(res => {
                this.pharmacys = res;
            });
        }
    }
    
    initMedicationDispenseList() {
        if ( this.medicationDispenses == null ) {
            new MedicationDispenseService(this.http).getMedicationDispenses().subscribe(res => {
                this.medicationDispenses = res;
            });
        }
    }
    
    initDiagnosisList() {
        if ( this.diagnosiss == null ) {
            new DiagnosisService(this.http).getDiagnosiss().subscribe(res => {
                this.diagnosiss = res;
            });
        }
    }
    
    initObservationList() {
        if ( this.observations == null ) {
            new ObservationService(this.http).getObservations().subscribe(res => {
                this.observations = res;
            });
        }
    }
    
    initCarePlanList() {
        if ( this.carePlans == null ) {
            new CarePlanService(this.http).getCarePlans().subscribe(res => {
                this.carePlans = res;
            });
        }
    }
    
    initCareTaskList() {
        if ( this.careTasks == null ) {
            new CareTaskService(this.http).getCareTasks().subscribe(res => {
                this.careTasks = res;
            });
        }
    }
    
    initAllergyList() {
        if ( this.allergys == null ) {
            new AllergyService(this.http).getAllergys().subscribe(res => {
                this.allergys = res;
            });
        }
    }
    
    initConditionList() {
        if ( this.conditions == null ) {
            new ConditionService(this.http).getConditions().subscribe(res => {
                this.conditions = res;
            });
        }
    }
    
    initInsurancePayerList() {
        if ( this.insurancePayers == null ) {
            new InsurancePayerService(this.http).getInsurancePayers().subscribe(res => {
                this.insurancePayers = res;
            });
        }
    }
    
    initInsurancePlanList() {
        if ( this.insurancePlans == null ) {
            new InsurancePlanService(this.http).getInsurancePlans().subscribe(res => {
                this.insurancePlans = res;
            });
        }
    }
    
    initCoverageList() {
        if ( this.coverages == null ) {
            new CoverageService(this.http).getCoverages().subscribe(res => {
                this.coverages = res;
            });
        }
    }
    
    initClaimList() {
        if ( this.claims == null ) {
            new ClaimService(this.http).getClaims().subscribe(res => {
                this.claims = res;
            });
        }
    }
    
    initAuthorizationList() {
        if ( this.authorizations == null ) {
            new AuthorizationService(this.http).getAuthorizations().subscribe(res => {
                this.authorizations = res;
            });
        }
    }
    
    initInvoiceList() {
        if ( this.invoices == null ) {
            new InvoiceService(this.http).getInvoices().subscribe(res => {
                this.invoices = res;
            });
        }
    }
    
    initPaymentList() {
        if ( this.payments == null ) {
            new PaymentService(this.http).getPayments().subscribe(res => {
                this.payments = res;
            });
        }
    }
    
    initMedicalDeviceList() {
        if ( this.medicalDevices == null ) {
            new MedicalDeviceService(this.http).getMedicalDevices().subscribe(res => {
                this.medicalDevices = res;
            });
        }
    }
    
    initSoftwareUpdateList() {
        if ( this.softwareUpdates == null ) {
            new SoftwareUpdateService(this.http).getSoftwareUpdates().subscribe(res => {
                this.softwareUpdates = res;
            });
        }
    }
    
    initMedicalSupplierList() {
        if ( this.medicalSuppliers == null ) {
            new MedicalSupplierService(this.http).getMedicalSuppliers().subscribe(res => {
                this.medicalSuppliers = res;
            });
        }
    }
    
    initInventoryItemList() {
        if ( this.inventoryItems == null ) {
            new InventoryItemService(this.http).getInventoryItems().subscribe(res => {
                this.inventoryItems = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
