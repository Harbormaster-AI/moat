import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListHealthSystemComponent from './components/ListHealthSystemComponent';
import CreateHealthSystemComponent from './components/CreateHealthSystemComponent';
import ViewHealthSystemComponent from './components/ViewHealthSystemComponent';
import ListFacilityComponent from './components/ListFacilityComponent';
import CreateFacilityComponent from './components/CreateFacilityComponent';
import ViewFacilityComponent from './components/ViewFacilityComponent';
import ListDepartmentComponent from './components/ListDepartmentComponent';
import CreateDepartmentComponent from './components/CreateDepartmentComponent';
import ViewDepartmentComponent from './components/ViewDepartmentComponent';
import ListCareTeamComponent from './components/ListCareTeamComponent';
import CreateCareTeamComponent from './components/CreateCareTeamComponent';
import ViewCareTeamComponent from './components/ViewCareTeamComponent';
import ListClinicianComponent from './components/ListClinicianComponent';
import CreateClinicianComponent from './components/CreateClinicianComponent';
import ViewClinicianComponent from './components/ViewClinicianComponent';
import ListPatientComponent from './components/ListPatientComponent';
import CreatePatientComponent from './components/CreatePatientComponent';
import ViewPatientComponent from './components/ViewPatientComponent';
import ListAppointmentComponent from './components/ListAppointmentComponent';
import CreateAppointmentComponent from './components/CreateAppointmentComponent';
import ViewAppointmentComponent from './components/ViewAppointmentComponent';
import ListEncounterComponent from './components/ListEncounterComponent';
import CreateEncounterComponent from './components/CreateEncounterComponent';
import ViewEncounterComponent from './components/ViewEncounterComponent';
import ListAdmissionComponent from './components/ListAdmissionComponent';
import CreateAdmissionComponent from './components/CreateAdmissionComponent';
import ViewAdmissionComponent from './components/ViewAdmissionComponent';
import ListDischargeComponent from './components/ListDischargeComponent';
import CreateDischargeComponent from './components/CreateDischargeComponent';
import ViewDischargeComponent from './components/ViewDischargeComponent';
import ListClinicalOrderComponent from './components/ListClinicalOrderComponent';
import CreateClinicalOrderComponent from './components/CreateClinicalOrderComponent';
import ViewClinicalOrderComponent from './components/ViewClinicalOrderComponent';
import ListMedicationOrderComponent from './components/ListMedicationOrderComponent';
import CreateMedicationOrderComponent from './components/CreateMedicationOrderComponent';
import ViewMedicationOrderComponent from './components/ViewMedicationOrderComponent';
import ListLaboratoryComponent from './components/ListLaboratoryComponent';
import CreateLaboratoryComponent from './components/CreateLaboratoryComponent';
import ViewLaboratoryComponent from './components/ViewLaboratoryComponent';
import ListLaboratoryOrderComponent from './components/ListLaboratoryOrderComponent';
import CreateLaboratoryOrderComponent from './components/CreateLaboratoryOrderComponent';
import ViewLaboratoryOrderComponent from './components/ViewLaboratoryOrderComponent';
import ListLabResultComponent from './components/ListLabResultComponent';
import CreateLabResultComponent from './components/CreateLabResultComponent';
import ViewLabResultComponent from './components/ViewLabResultComponent';
import ListImagingCenterComponent from './components/ListImagingCenterComponent';
import CreateImagingCenterComponent from './components/CreateImagingCenterComponent';
import ViewImagingCenterComponent from './components/ViewImagingCenterComponent';
import ListImagingOrderComponent from './components/ListImagingOrderComponent';
import CreateImagingOrderComponent from './components/CreateImagingOrderComponent';
import ViewImagingOrderComponent from './components/ViewImagingOrderComponent';
import ListImagingReportComponent from './components/ListImagingReportComponent';
import CreateImagingReportComponent from './components/CreateImagingReportComponent';
import ViewImagingReportComponent from './components/ViewImagingReportComponent';
import ListProcedureOrderComponent from './components/ListProcedureOrderComponent';
import CreateProcedureOrderComponent from './components/CreateProcedureOrderComponent';
import ViewProcedureOrderComponent from './components/ViewProcedureOrderComponent';
import ListProcedureComponent from './components/ListProcedureComponent';
import CreateProcedureComponent from './components/CreateProcedureComponent';
import ViewProcedureComponent from './components/ViewProcedureComponent';
import ListPharmacyComponent from './components/ListPharmacyComponent';
import CreatePharmacyComponent from './components/CreatePharmacyComponent';
import ViewPharmacyComponent from './components/ViewPharmacyComponent';
import ListMedicationDispenseComponent from './components/ListMedicationDispenseComponent';
import CreateMedicationDispenseComponent from './components/CreateMedicationDispenseComponent';
import ViewMedicationDispenseComponent from './components/ViewMedicationDispenseComponent';
import ListDiagnosisComponent from './components/ListDiagnosisComponent';
import CreateDiagnosisComponent from './components/CreateDiagnosisComponent';
import ViewDiagnosisComponent from './components/ViewDiagnosisComponent';
import ListObservationComponent from './components/ListObservationComponent';
import CreateObservationComponent from './components/CreateObservationComponent';
import ViewObservationComponent from './components/ViewObservationComponent';
import ListCarePlanComponent from './components/ListCarePlanComponent';
import CreateCarePlanComponent from './components/CreateCarePlanComponent';
import ViewCarePlanComponent from './components/ViewCarePlanComponent';
import ListCareTaskComponent from './components/ListCareTaskComponent';
import CreateCareTaskComponent from './components/CreateCareTaskComponent';
import ViewCareTaskComponent from './components/ViewCareTaskComponent';
import ListAllergyComponent from './components/ListAllergyComponent';
import CreateAllergyComponent from './components/CreateAllergyComponent';
import ViewAllergyComponent from './components/ViewAllergyComponent';
import ListConditionComponent from './components/ListConditionComponent';
import CreateConditionComponent from './components/CreateConditionComponent';
import ViewConditionComponent from './components/ViewConditionComponent';
import ListInsurancePayerComponent from './components/ListInsurancePayerComponent';
import CreateInsurancePayerComponent from './components/CreateInsurancePayerComponent';
import ViewInsurancePayerComponent from './components/ViewInsurancePayerComponent';
import ListInsurancePlanComponent from './components/ListInsurancePlanComponent';
import CreateInsurancePlanComponent from './components/CreateInsurancePlanComponent';
import ViewInsurancePlanComponent from './components/ViewInsurancePlanComponent';
import ListCoverageComponent from './components/ListCoverageComponent';
import CreateCoverageComponent from './components/CreateCoverageComponent';
import ViewCoverageComponent from './components/ViewCoverageComponent';
import ListClaimComponent from './components/ListClaimComponent';
import CreateClaimComponent from './components/CreateClaimComponent';
import ViewClaimComponent from './components/ViewClaimComponent';
import ListAuthorizationComponent from './components/ListAuthorizationComponent';
import CreateAuthorizationComponent from './components/CreateAuthorizationComponent';
import ViewAuthorizationComponent from './components/ViewAuthorizationComponent';
import ListInvoiceComponent from './components/ListInvoiceComponent';
import CreateInvoiceComponent from './components/CreateInvoiceComponent';
import ViewInvoiceComponent from './components/ViewInvoiceComponent';
import ListPaymentComponent from './components/ListPaymentComponent';
import CreatePaymentComponent from './components/CreatePaymentComponent';
import ViewPaymentComponent from './components/ViewPaymentComponent';
import ListMedicalDeviceComponent from './components/ListMedicalDeviceComponent';
import CreateMedicalDeviceComponent from './components/CreateMedicalDeviceComponent';
import ViewMedicalDeviceComponent from './components/ViewMedicalDeviceComponent';
import ListSoftwareUpdateComponent from './components/ListSoftwareUpdateComponent';
import CreateSoftwareUpdateComponent from './components/CreateSoftwareUpdateComponent';
import ViewSoftwareUpdateComponent from './components/ViewSoftwareUpdateComponent';
import ListMedicalSupplierComponent from './components/ListMedicalSupplierComponent';
import CreateMedicalSupplierComponent from './components/CreateMedicalSupplierComponent';
import ViewMedicalSupplierComponent from './components/ViewMedicalSupplierComponent';
import ListInventoryItemComponent from './components/ListInventoryItemComponent';
import CreateInventoryItemComponent from './components/CreateInventoryItemComponent';
import ViewInventoryItemComponent from './components/ViewInventoryItemComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/healthSystems" component = {ListHealthSystemComponent}></Route>
                            <Route path = "/add-healthSystem/:id" component = {CreateHealthSystemComponent}></Route>
                            <Route path = "/view-healthSystem/:id" component = {ViewHealthSystemComponent}></Route>
                          {/* <Route path = "/update-healthSystem/:id" component = {UpdateHealthSystemComponent}></Route> */}
                            <Route path = "/facilitys" component = {ListFacilityComponent}></Route>
                            <Route path = "/add-facility/:id" component = {CreateFacilityComponent}></Route>
                            <Route path = "/view-facility/:id" component = {ViewFacilityComponent}></Route>
                          {/* <Route path = "/update-facility/:id" component = {UpdateFacilityComponent}></Route> */}
                            <Route path = "/departments" component = {ListDepartmentComponent}></Route>
                            <Route path = "/add-department/:id" component = {CreateDepartmentComponent}></Route>
                            <Route path = "/view-department/:id" component = {ViewDepartmentComponent}></Route>
                          {/* <Route path = "/update-department/:id" component = {UpdateDepartmentComponent}></Route> */}
                            <Route path = "/careTeams" component = {ListCareTeamComponent}></Route>
                            <Route path = "/add-careTeam/:id" component = {CreateCareTeamComponent}></Route>
                            <Route path = "/view-careTeam/:id" component = {ViewCareTeamComponent}></Route>
                          {/* <Route path = "/update-careTeam/:id" component = {UpdateCareTeamComponent}></Route> */}
                            <Route path = "/clinicians" component = {ListClinicianComponent}></Route>
                            <Route path = "/add-clinician/:id" component = {CreateClinicianComponent}></Route>
                            <Route path = "/view-clinician/:id" component = {ViewClinicianComponent}></Route>
                          {/* <Route path = "/update-clinician/:id" component = {UpdateClinicianComponent}></Route> */}
                            <Route path = "/patients" component = {ListPatientComponent}></Route>
                            <Route path = "/add-patient/:id" component = {CreatePatientComponent}></Route>
                            <Route path = "/view-patient/:id" component = {ViewPatientComponent}></Route>
                          {/* <Route path = "/update-patient/:id" component = {UpdatePatientComponent}></Route> */}
                            <Route path = "/appointments" component = {ListAppointmentComponent}></Route>
                            <Route path = "/add-appointment/:id" component = {CreateAppointmentComponent}></Route>
                            <Route path = "/view-appointment/:id" component = {ViewAppointmentComponent}></Route>
                          {/* <Route path = "/update-appointment/:id" component = {UpdateAppointmentComponent}></Route> */}
                            <Route path = "/encounters" component = {ListEncounterComponent}></Route>
                            <Route path = "/add-encounter/:id" component = {CreateEncounterComponent}></Route>
                            <Route path = "/view-encounter/:id" component = {ViewEncounterComponent}></Route>
                          {/* <Route path = "/update-encounter/:id" component = {UpdateEncounterComponent}></Route> */}
                            <Route path = "/admissions" component = {ListAdmissionComponent}></Route>
                            <Route path = "/add-admission/:id" component = {CreateAdmissionComponent}></Route>
                            <Route path = "/view-admission/:id" component = {ViewAdmissionComponent}></Route>
                          {/* <Route path = "/update-admission/:id" component = {UpdateAdmissionComponent}></Route> */}
                            <Route path = "/discharges" component = {ListDischargeComponent}></Route>
                            <Route path = "/add-discharge/:id" component = {CreateDischargeComponent}></Route>
                            <Route path = "/view-discharge/:id" component = {ViewDischargeComponent}></Route>
                          {/* <Route path = "/update-discharge/:id" component = {UpdateDischargeComponent}></Route> */}
                            <Route path = "/clinicalOrders" component = {ListClinicalOrderComponent}></Route>
                            <Route path = "/add-clinicalOrder/:id" component = {CreateClinicalOrderComponent}></Route>
                            <Route path = "/view-clinicalOrder/:id" component = {ViewClinicalOrderComponent}></Route>
                          {/* <Route path = "/update-clinicalOrder/:id" component = {UpdateClinicalOrderComponent}></Route> */}
                            <Route path = "/medicationOrders" component = {ListMedicationOrderComponent}></Route>
                            <Route path = "/add-medicationOrder/:id" component = {CreateMedicationOrderComponent}></Route>
                            <Route path = "/view-medicationOrder/:id" component = {ViewMedicationOrderComponent}></Route>
                          {/* <Route path = "/update-medicationOrder/:id" component = {UpdateMedicationOrderComponent}></Route> */}
                            <Route path = "/laboratorys" component = {ListLaboratoryComponent}></Route>
                            <Route path = "/add-laboratory/:id" component = {CreateLaboratoryComponent}></Route>
                            <Route path = "/view-laboratory/:id" component = {ViewLaboratoryComponent}></Route>
                          {/* <Route path = "/update-laboratory/:id" component = {UpdateLaboratoryComponent}></Route> */}
                            <Route path = "/laboratoryOrders" component = {ListLaboratoryOrderComponent}></Route>
                            <Route path = "/add-laboratoryOrder/:id" component = {CreateLaboratoryOrderComponent}></Route>
                            <Route path = "/view-laboratoryOrder/:id" component = {ViewLaboratoryOrderComponent}></Route>
                          {/* <Route path = "/update-laboratoryOrder/:id" component = {UpdateLaboratoryOrderComponent}></Route> */}
                            <Route path = "/labResults" component = {ListLabResultComponent}></Route>
                            <Route path = "/add-labResult/:id" component = {CreateLabResultComponent}></Route>
                            <Route path = "/view-labResult/:id" component = {ViewLabResultComponent}></Route>
                          {/* <Route path = "/update-labResult/:id" component = {UpdateLabResultComponent}></Route> */}
                            <Route path = "/imagingCenters" component = {ListImagingCenterComponent}></Route>
                            <Route path = "/add-imagingCenter/:id" component = {CreateImagingCenterComponent}></Route>
                            <Route path = "/view-imagingCenter/:id" component = {ViewImagingCenterComponent}></Route>
                          {/* <Route path = "/update-imagingCenter/:id" component = {UpdateImagingCenterComponent}></Route> */}
                            <Route path = "/imagingOrders" component = {ListImagingOrderComponent}></Route>
                            <Route path = "/add-imagingOrder/:id" component = {CreateImagingOrderComponent}></Route>
                            <Route path = "/view-imagingOrder/:id" component = {ViewImagingOrderComponent}></Route>
                          {/* <Route path = "/update-imagingOrder/:id" component = {UpdateImagingOrderComponent}></Route> */}
                            <Route path = "/imagingReports" component = {ListImagingReportComponent}></Route>
                            <Route path = "/add-imagingReport/:id" component = {CreateImagingReportComponent}></Route>
                            <Route path = "/view-imagingReport/:id" component = {ViewImagingReportComponent}></Route>
                          {/* <Route path = "/update-imagingReport/:id" component = {UpdateImagingReportComponent}></Route> */}
                            <Route path = "/procedureOrders" component = {ListProcedureOrderComponent}></Route>
                            <Route path = "/add-procedureOrder/:id" component = {CreateProcedureOrderComponent}></Route>
                            <Route path = "/view-procedureOrder/:id" component = {ViewProcedureOrderComponent}></Route>
                          {/* <Route path = "/update-procedureOrder/:id" component = {UpdateProcedureOrderComponent}></Route> */}
                            <Route path = "/procedures" component = {ListProcedureComponent}></Route>
                            <Route path = "/add-procedure/:id" component = {CreateProcedureComponent}></Route>
                            <Route path = "/view-procedure/:id" component = {ViewProcedureComponent}></Route>
                          {/* <Route path = "/update-procedure/:id" component = {UpdateProcedureComponent}></Route> */}
                            <Route path = "/pharmacys" component = {ListPharmacyComponent}></Route>
                            <Route path = "/add-pharmacy/:id" component = {CreatePharmacyComponent}></Route>
                            <Route path = "/view-pharmacy/:id" component = {ViewPharmacyComponent}></Route>
                          {/* <Route path = "/update-pharmacy/:id" component = {UpdatePharmacyComponent}></Route> */}
                            <Route path = "/medicationDispenses" component = {ListMedicationDispenseComponent}></Route>
                            <Route path = "/add-medicationDispense/:id" component = {CreateMedicationDispenseComponent}></Route>
                            <Route path = "/view-medicationDispense/:id" component = {ViewMedicationDispenseComponent}></Route>
                          {/* <Route path = "/update-medicationDispense/:id" component = {UpdateMedicationDispenseComponent}></Route> */}
                            <Route path = "/diagnosiss" component = {ListDiagnosisComponent}></Route>
                            <Route path = "/add-diagnosis/:id" component = {CreateDiagnosisComponent}></Route>
                            <Route path = "/view-diagnosis/:id" component = {ViewDiagnosisComponent}></Route>
                          {/* <Route path = "/update-diagnosis/:id" component = {UpdateDiagnosisComponent}></Route> */}
                            <Route path = "/observations" component = {ListObservationComponent}></Route>
                            <Route path = "/add-observation/:id" component = {CreateObservationComponent}></Route>
                            <Route path = "/view-observation/:id" component = {ViewObservationComponent}></Route>
                          {/* <Route path = "/update-observation/:id" component = {UpdateObservationComponent}></Route> */}
                            <Route path = "/carePlans" component = {ListCarePlanComponent}></Route>
                            <Route path = "/add-carePlan/:id" component = {CreateCarePlanComponent}></Route>
                            <Route path = "/view-carePlan/:id" component = {ViewCarePlanComponent}></Route>
                          {/* <Route path = "/update-carePlan/:id" component = {UpdateCarePlanComponent}></Route> */}
                            <Route path = "/careTasks" component = {ListCareTaskComponent}></Route>
                            <Route path = "/add-careTask/:id" component = {CreateCareTaskComponent}></Route>
                            <Route path = "/view-careTask/:id" component = {ViewCareTaskComponent}></Route>
                          {/* <Route path = "/update-careTask/:id" component = {UpdateCareTaskComponent}></Route> */}
                            <Route path = "/allergys" component = {ListAllergyComponent}></Route>
                            <Route path = "/add-allergy/:id" component = {CreateAllergyComponent}></Route>
                            <Route path = "/view-allergy/:id" component = {ViewAllergyComponent}></Route>
                          {/* <Route path = "/update-allergy/:id" component = {UpdateAllergyComponent}></Route> */}
                            <Route path = "/conditions" component = {ListConditionComponent}></Route>
                            <Route path = "/add-condition/:id" component = {CreateConditionComponent}></Route>
                            <Route path = "/view-condition/:id" component = {ViewConditionComponent}></Route>
                          {/* <Route path = "/update-condition/:id" component = {UpdateConditionComponent}></Route> */}
                            <Route path = "/insurancePayers" component = {ListInsurancePayerComponent}></Route>
                            <Route path = "/add-insurancePayer/:id" component = {CreateInsurancePayerComponent}></Route>
                            <Route path = "/view-insurancePayer/:id" component = {ViewInsurancePayerComponent}></Route>
                          {/* <Route path = "/update-insurancePayer/:id" component = {UpdateInsurancePayerComponent}></Route> */}
                            <Route path = "/insurancePlans" component = {ListInsurancePlanComponent}></Route>
                            <Route path = "/add-insurancePlan/:id" component = {CreateInsurancePlanComponent}></Route>
                            <Route path = "/view-insurancePlan/:id" component = {ViewInsurancePlanComponent}></Route>
                          {/* <Route path = "/update-insurancePlan/:id" component = {UpdateInsurancePlanComponent}></Route> */}
                            <Route path = "/coverages" component = {ListCoverageComponent}></Route>
                            <Route path = "/add-coverage/:id" component = {CreateCoverageComponent}></Route>
                            <Route path = "/view-coverage/:id" component = {ViewCoverageComponent}></Route>
                          {/* <Route path = "/update-coverage/:id" component = {UpdateCoverageComponent}></Route> */}
                            <Route path = "/claims" component = {ListClaimComponent}></Route>
                            <Route path = "/add-claim/:id" component = {CreateClaimComponent}></Route>
                            <Route path = "/view-claim/:id" component = {ViewClaimComponent}></Route>
                          {/* <Route path = "/update-claim/:id" component = {UpdateClaimComponent}></Route> */}
                            <Route path = "/authorizations" component = {ListAuthorizationComponent}></Route>
                            <Route path = "/add-authorization/:id" component = {CreateAuthorizationComponent}></Route>
                            <Route path = "/view-authorization/:id" component = {ViewAuthorizationComponent}></Route>
                          {/* <Route path = "/update-authorization/:id" component = {UpdateAuthorizationComponent}></Route> */}
                            <Route path = "/invoices" component = {ListInvoiceComponent}></Route>
                            <Route path = "/add-invoice/:id" component = {CreateInvoiceComponent}></Route>
                            <Route path = "/view-invoice/:id" component = {ViewInvoiceComponent}></Route>
                          {/* <Route path = "/update-invoice/:id" component = {UpdateInvoiceComponent}></Route> */}
                            <Route path = "/payments" component = {ListPaymentComponent}></Route>
                            <Route path = "/add-payment/:id" component = {CreatePaymentComponent}></Route>
                            <Route path = "/view-payment/:id" component = {ViewPaymentComponent}></Route>
                          {/* <Route path = "/update-payment/:id" component = {UpdatePaymentComponent}></Route> */}
                            <Route path = "/medicalDevices" component = {ListMedicalDeviceComponent}></Route>
                            <Route path = "/add-medicalDevice/:id" component = {CreateMedicalDeviceComponent}></Route>
                            <Route path = "/view-medicalDevice/:id" component = {ViewMedicalDeviceComponent}></Route>
                          {/* <Route path = "/update-medicalDevice/:id" component = {UpdateMedicalDeviceComponent}></Route> */}
                            <Route path = "/softwareUpdates" component = {ListSoftwareUpdateComponent}></Route>
                            <Route path = "/add-softwareUpdate/:id" component = {CreateSoftwareUpdateComponent}></Route>
                            <Route path = "/view-softwareUpdate/:id" component = {ViewSoftwareUpdateComponent}></Route>
                          {/* <Route path = "/update-softwareUpdate/:id" component = {UpdateSoftwareUpdateComponent}></Route> */}
                            <Route path = "/medicalSuppliers" component = {ListMedicalSupplierComponent}></Route>
                            <Route path = "/add-medicalSupplier/:id" component = {CreateMedicalSupplierComponent}></Route>
                            <Route path = "/view-medicalSupplier/:id" component = {ViewMedicalSupplierComponent}></Route>
                          {/* <Route path = "/update-medicalSupplier/:id" component = {UpdateMedicalSupplierComponent}></Route> */}
                            <Route path = "/inventoryItems" component = {ListInventoryItemComponent}></Route>
                            <Route path = "/add-inventoryItem/:id" component = {CreateInventoryItemComponent}></Route>
                            <Route path = "/view-inventoryItem/:id" component = {ViewInventoryItemComponent}></Route>
                          {/* <Route path = "/update-inventoryItem/:id" component = {UpdateInventoryItemComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
