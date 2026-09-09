import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListAerospaceManufacturerComponent from './components/ListAerospaceManufacturerComponent';
import CreateAerospaceManufacturerComponent from './components/CreateAerospaceManufacturerComponent';
import ViewAerospaceManufacturerComponent from './components/ViewAerospaceManufacturerComponent';
import ListAircraftProgramComponent from './components/ListAircraftProgramComponent';
import CreateAircraftProgramComponent from './components/CreateAircraftProgramComponent';
import ViewAircraftProgramComponent from './components/ViewAircraftProgramComponent';
import ListAircraftFamilyComponent from './components/ListAircraftFamilyComponent';
import CreateAircraftFamilyComponent from './components/CreateAircraftFamilyComponent';
import ViewAircraftFamilyComponent from './components/ViewAircraftFamilyComponent';
import ListAircraftModelComponent from './components/ListAircraftModelComponent';
import CreateAircraftModelComponent from './components/CreateAircraftModelComponent';
import ViewAircraftModelComponent from './components/ViewAircraftModelComponent';
import ListEngineTypeComponent from './components/ListEngineTypeComponent';
import CreateEngineTypeComponent from './components/CreateEngineTypeComponent';
import ViewEngineTypeComponent from './components/ViewEngineTypeComponent';
import ListAircraftVariantComponent from './components/ListAircraftVariantComponent';
import CreateAircraftVariantComponent from './components/CreateAircraftVariantComponent';
import ViewAircraftVariantComponent from './components/ViewAircraftVariantComponent';
import ListAvionicsSuiteComponent from './components/ListAvionicsSuiteComponent';
import CreateAvionicsSuiteComponent from './components/CreateAvionicsSuiteComponent';
import ViewAvionicsSuiteComponent from './components/ViewAvionicsSuiteComponent';
import ListAPUComponent from './components/ListAPUComponent';
import CreateAPUComponent from './components/CreateAPUComponent';
import ViewAPUComponent from './components/ViewAPUComponent';
import ListLandingGearComponent from './components/ListLandingGearComponent';
import CreateLandingGearComponent from './components/CreateLandingGearComponent';
import ViewLandingGearComponent from './components/ViewLandingGearComponent';
import ListAircraftOptionComponent from './components/ListAircraftOptionComponent';
import CreateAircraftOptionComponent from './components/CreateAircraftOptionComponent';
import ViewAircraftOptionComponent from './components/ViewAircraftOptionComponent';
import ListAircraftPackageComponent from './components/ListAircraftPackageComponent';
import CreateAircraftPackageComponent from './components/CreateAircraftPackageComponent';
import ViewAircraftPackageComponent from './components/ViewAircraftPackageComponent';
import ListSupplierComponent from './components/ListSupplierComponent';
import CreateSupplierComponent from './components/CreateSupplierComponent';
import ViewSupplierComponent from './components/ViewSupplierComponent';
import ListComponent_Component from './components/ListComponent_Component';
import CreateComponent_Component from './components/CreateComponent_Component';
import ViewComponent_Component from './components/ViewComponent_Component';
import ListPlantComponent from './components/ListPlantComponent';
import CreatePlantComponent from './components/CreatePlantComponent';
import ViewPlantComponent from './components/ViewPlantComponent';
import ListProductionLineComponent from './components/ListProductionLineComponent';
import CreateProductionLineComponent from './components/CreateProductionLineComponent';
import ViewProductionLineComponent from './components/ViewProductionLineComponent';
import ListWorkCenterComponent from './components/ListWorkCenterComponent';
import CreateWorkCenterComponent from './components/CreateWorkCenterComponent';
import ViewWorkCenterComponent from './components/ViewWorkCenterComponent';
import ListProductionOrderComponent from './components/ListProductionOrderComponent';
import CreateProductionOrderComponent from './components/CreateProductionOrderComponent';
import ViewProductionOrderComponent from './components/ViewProductionOrderComponent';
import ListBuildScheduleComponent from './components/ListBuildScheduleComponent';
import CreateBuildScheduleComponent from './components/CreateBuildScheduleComponent';
import ViewBuildScheduleComponent from './components/ViewBuildScheduleComponent';
import ListWarehouseComponent from './components/ListWarehouseComponent';
import CreateWarehouseComponent from './components/CreateWarehouseComponent';
import ViewWarehouseComponent from './components/ViewWarehouseComponent';
import ListInventoryItemComponent from './components/ListInventoryItemComponent';
import CreateInventoryItemComponent from './components/CreateInventoryItemComponent';
import ViewInventoryItemComponent from './components/ViewInventoryItemComponent';
import ListOperatorComponent from './components/ListOperatorComponent';
import CreateOperatorComponent from './components/CreateOperatorComponent';
import ViewOperatorComponent from './components/ViewOperatorComponent';
import ListAircraftOrderComponent from './components/ListAircraftOrderComponent';
import CreateAircraftOrderComponent from './components/CreateAircraftOrderComponent';
import ViewAircraftOrderComponent from './components/ViewAircraftOrderComponent';
import ListQuoteComponent from './components/ListQuoteComponent';
import CreateQuoteComponent from './components/CreateQuoteComponent';
import ViewQuoteComponent from './components/ViewQuoteComponent';
import ListPurchaseAgreementComponent from './components/ListPurchaseAgreementComponent';
import CreatePurchaseAgreementComponent from './components/CreatePurchaseAgreementComponent';
import ViewPurchaseAgreementComponent from './components/ViewPurchaseAgreementComponent';
import ListAircraftComponent from './components/ListAircraftComponent';
import CreateAircraftComponent from './components/CreateAircraftComponent';
import ViewAircraftComponent from './components/ViewAircraftComponent';
import ListRegistrationComponent from './components/ListRegistrationComponent';
import CreateRegistrationComponent from './components/CreateRegistrationComponent';
import ViewRegistrationComponent from './components/ViewRegistrationComponent';
import ListWarrantyComponent from './components/ListWarrantyComponent';
import CreateWarrantyComponent from './components/CreateWarrantyComponent';
import ViewWarrantyComponent from './components/ViewWarrantyComponent';
import ListCabinLayoutComponent from './components/ListCabinLayoutComponent';
import CreateCabinLayoutComponent from './components/CreateCabinLayoutComponent';
import ViewCabinLayoutComponent from './components/ViewCabinLayoutComponent';
import ListMROFacilityComponent from './components/ListMROFacilityComponent';
import CreateMROFacilityComponent from './components/CreateMROFacilityComponent';
import ViewMROFacilityComponent from './components/ViewMROFacilityComponent';
import ListMaintenanceAppointmentComponent from './components/ListMaintenanceAppointmentComponent';
import CreateMaintenanceAppointmentComponent from './components/CreateMaintenanceAppointmentComponent';
import ViewMaintenanceAppointmentComponent from './components/ViewMaintenanceAppointmentComponent';
import ListMaintenanceWorkOrderComponent from './components/ListMaintenanceWorkOrderComponent';
import CreateMaintenanceWorkOrderComponent from './components/CreateMaintenanceWorkOrderComponent';
import ViewMaintenanceWorkOrderComponent from './components/ViewMaintenanceWorkOrderComponent';
import ListAirworthinessDirectiveComponent from './components/ListAirworthinessDirectiveComponent';
import CreateAirworthinessDirectiveComponent from './components/CreateAirworthinessDirectiveComponent';
import ViewAirworthinessDirectiveComponent from './components/ViewAirworthinessDirectiveComponent';
import ListServiceBulletinComponent from './components/ListServiceBulletinComponent';
import CreateServiceBulletinComponent from './components/CreateServiceBulletinComponent';
import ViewServiceBulletinComponent from './components/ViewServiceBulletinComponent';
import ListConnectedAircraftComponent from './components/ListConnectedAircraftComponent';
import CreateConnectedAircraftComponent from './components/CreateConnectedAircraftComponent';
import ViewConnectedAircraftComponent from './components/ViewConnectedAircraftComponent';
import ListFlightHealthEventComponent from './components/ListFlightHealthEventComponent';
import CreateFlightHealthEventComponent from './components/CreateFlightHealthEventComponent';
import ViewFlightHealthEventComponent from './components/ViewFlightHealthEventComponent';
import ListSoftwareLoadComponent from './components/ListSoftwareLoadComponent';
import CreateSoftwareLoadComponent from './components/CreateSoftwareLoadComponent';
import ViewSoftwareLoadComponent from './components/ViewSoftwareLoadComponent';
import ListTypeCertificateComponent from './components/ListTypeCertificateComponent';
import CreateTypeCertificateComponent from './components/CreateTypeCertificateComponent';
import ViewTypeCertificateComponent from './components/ViewTypeCertificateComponent';
import ListProductionCertificateComponent from './components/ListProductionCertificateComponent';
import CreateProductionCertificateComponent from './components/CreateProductionCertificateComponent';
import ViewProductionCertificateComponent from './components/ViewProductionCertificateComponent';
import ListSalesRegionComponent from './components/ListSalesRegionComponent';
import CreateSalesRegionComponent from './components/CreateSalesRegionComponent';
import ViewSalesRegionComponent from './components/ViewSalesRegionComponent';
import ListSalesCampaignComponent from './components/ListSalesCampaignComponent';
import CreateSalesCampaignComponent from './components/CreateSalesCampaignComponent';
import ViewSalesCampaignComponent from './components/ViewSalesCampaignComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/aerospaceManufacturers" component = {ListAerospaceManufacturerComponent}></Route>
                            <Route path = "/add-aerospaceManufacturer/:id" component = {CreateAerospaceManufacturerComponent}></Route>
                            <Route path = "/view-aerospaceManufacturer/:id" component = {ViewAerospaceManufacturerComponent}></Route>
                          {/* <Route path = "/update-aerospaceManufacturer/:id" component = {UpdateAerospaceManufacturerComponent}></Route> */}
                            <Route path = "/aircraftPrograms" component = {ListAircraftProgramComponent}></Route>
                            <Route path = "/add-aircraftProgram/:id" component = {CreateAircraftProgramComponent}></Route>
                            <Route path = "/view-aircraftProgram/:id" component = {ViewAircraftProgramComponent}></Route>
                          {/* <Route path = "/update-aircraftProgram/:id" component = {UpdateAircraftProgramComponent}></Route> */}
                            <Route path = "/aircraftFamilys" component = {ListAircraftFamilyComponent}></Route>
                            <Route path = "/add-aircraftFamily/:id" component = {CreateAircraftFamilyComponent}></Route>
                            <Route path = "/view-aircraftFamily/:id" component = {ViewAircraftFamilyComponent}></Route>
                          {/* <Route path = "/update-aircraftFamily/:id" component = {UpdateAircraftFamilyComponent}></Route> */}
                            <Route path = "/aircraftModels" component = {ListAircraftModelComponent}></Route>
                            <Route path = "/add-aircraftModel/:id" component = {CreateAircraftModelComponent}></Route>
                            <Route path = "/view-aircraftModel/:id" component = {ViewAircraftModelComponent}></Route>
                          {/* <Route path = "/update-aircraftModel/:id" component = {UpdateAircraftModelComponent}></Route> */}
                            <Route path = "/engineTypes" component = {ListEngineTypeComponent}></Route>
                            <Route path = "/add-engineType/:id" component = {CreateEngineTypeComponent}></Route>
                            <Route path = "/view-engineType/:id" component = {ViewEngineTypeComponent}></Route>
                          {/* <Route path = "/update-engineType/:id" component = {UpdateEngineTypeComponent}></Route> */}
                            <Route path = "/aircraftVariants" component = {ListAircraftVariantComponent}></Route>
                            <Route path = "/add-aircraftVariant/:id" component = {CreateAircraftVariantComponent}></Route>
                            <Route path = "/view-aircraftVariant/:id" component = {ViewAircraftVariantComponent}></Route>
                          {/* <Route path = "/update-aircraftVariant/:id" component = {UpdateAircraftVariantComponent}></Route> */}
                            <Route path = "/avionicsSuites" component = {ListAvionicsSuiteComponent}></Route>
                            <Route path = "/add-avionicsSuite/:id" component = {CreateAvionicsSuiteComponent}></Route>
                            <Route path = "/view-avionicsSuite/:id" component = {ViewAvionicsSuiteComponent}></Route>
                          {/* <Route path = "/update-avionicsSuite/:id" component = {UpdateAvionicsSuiteComponent}></Route> */}
                            <Route path = "/aPUs" component = {ListAPUComponent}></Route>
                            <Route path = "/add-aPU/:id" component = {CreateAPUComponent}></Route>
                            <Route path = "/view-aPU/:id" component = {ViewAPUComponent}></Route>
                          {/* <Route path = "/update-aPU/:id" component = {UpdateAPUComponent}></Route> */}
                            <Route path = "/landingGears" component = {ListLandingGearComponent}></Route>
                            <Route path = "/add-landingGear/:id" component = {CreateLandingGearComponent}></Route>
                            <Route path = "/view-landingGear/:id" component = {ViewLandingGearComponent}></Route>
                          {/* <Route path = "/update-landingGear/:id" component = {UpdateLandingGearComponent}></Route> */}
                            <Route path = "/aircraftOptions" component = {ListAircraftOptionComponent}></Route>
                            <Route path = "/add-aircraftOption/:id" component = {CreateAircraftOptionComponent}></Route>
                            <Route path = "/view-aircraftOption/:id" component = {ViewAircraftOptionComponent}></Route>
                          {/* <Route path = "/update-aircraftOption/:id" component = {UpdateAircraftOptionComponent}></Route> */}
                            <Route path = "/aircraftPackages" component = {ListAircraftPackageComponent}></Route>
                            <Route path = "/add-aircraftPackage/:id" component = {CreateAircraftPackageComponent}></Route>
                            <Route path = "/view-aircraftPackage/:id" component = {ViewAircraftPackageComponent}></Route>
                          {/* <Route path = "/update-aircraftPackage/:id" component = {UpdateAircraftPackageComponent}></Route> */}
                            <Route path = "/suppliers" component = {ListSupplierComponent}></Route>
                            <Route path = "/add-supplier/:id" component = {CreateSupplierComponent}></Route>
                            <Route path = "/view-supplier/:id" component = {ViewSupplierComponent}></Route>
                          {/* <Route path = "/update-supplier/:id" component = {UpdateSupplierComponent}></Route> */}
                            <Route path = "/component_s" component = {ListComponent_Component}></Route>
                            <Route path = "/add-component_/:id" component = {CreateComponent_Component}></Route>
                            <Route path = "/view-component_/:id" component = {ViewComponent_Component}></Route>
                          {/* <Route path = "/update-component_/:id" component = {UpdateComponent_Component}></Route> */}
                            <Route path = "/plants" component = {ListPlantComponent}></Route>
                            <Route path = "/add-plant/:id" component = {CreatePlantComponent}></Route>
                            <Route path = "/view-plant/:id" component = {ViewPlantComponent}></Route>
                          {/* <Route path = "/update-plant/:id" component = {UpdatePlantComponent}></Route> */}
                            <Route path = "/productionLines" component = {ListProductionLineComponent}></Route>
                            <Route path = "/add-productionLine/:id" component = {CreateProductionLineComponent}></Route>
                            <Route path = "/view-productionLine/:id" component = {ViewProductionLineComponent}></Route>
                          {/* <Route path = "/update-productionLine/:id" component = {UpdateProductionLineComponent}></Route> */}
                            <Route path = "/workCenters" component = {ListWorkCenterComponent}></Route>
                            <Route path = "/add-workCenter/:id" component = {CreateWorkCenterComponent}></Route>
                            <Route path = "/view-workCenter/:id" component = {ViewWorkCenterComponent}></Route>
                          {/* <Route path = "/update-workCenter/:id" component = {UpdateWorkCenterComponent}></Route> */}
                            <Route path = "/productionOrders" component = {ListProductionOrderComponent}></Route>
                            <Route path = "/add-productionOrder/:id" component = {CreateProductionOrderComponent}></Route>
                            <Route path = "/view-productionOrder/:id" component = {ViewProductionOrderComponent}></Route>
                          {/* <Route path = "/update-productionOrder/:id" component = {UpdateProductionOrderComponent}></Route> */}
                            <Route path = "/buildSchedules" component = {ListBuildScheduleComponent}></Route>
                            <Route path = "/add-buildSchedule/:id" component = {CreateBuildScheduleComponent}></Route>
                            <Route path = "/view-buildSchedule/:id" component = {ViewBuildScheduleComponent}></Route>
                          {/* <Route path = "/update-buildSchedule/:id" component = {UpdateBuildScheduleComponent}></Route> */}
                            <Route path = "/warehouses" component = {ListWarehouseComponent}></Route>
                            <Route path = "/add-warehouse/:id" component = {CreateWarehouseComponent}></Route>
                            <Route path = "/view-warehouse/:id" component = {ViewWarehouseComponent}></Route>
                          {/* <Route path = "/update-warehouse/:id" component = {UpdateWarehouseComponent}></Route> */}
                            <Route path = "/inventoryItems" component = {ListInventoryItemComponent}></Route>
                            <Route path = "/add-inventoryItem/:id" component = {CreateInventoryItemComponent}></Route>
                            <Route path = "/view-inventoryItem/:id" component = {ViewInventoryItemComponent}></Route>
                          {/* <Route path = "/update-inventoryItem/:id" component = {UpdateInventoryItemComponent}></Route> */}
                            <Route path = "/operators" component = {ListOperatorComponent}></Route>
                            <Route path = "/add-operator/:id" component = {CreateOperatorComponent}></Route>
                            <Route path = "/view-operator/:id" component = {ViewOperatorComponent}></Route>
                          {/* <Route path = "/update-operator/:id" component = {UpdateOperatorComponent}></Route> */}
                            <Route path = "/aircraftOrders" component = {ListAircraftOrderComponent}></Route>
                            <Route path = "/add-aircraftOrder/:id" component = {CreateAircraftOrderComponent}></Route>
                            <Route path = "/view-aircraftOrder/:id" component = {ViewAircraftOrderComponent}></Route>
                          {/* <Route path = "/update-aircraftOrder/:id" component = {UpdateAircraftOrderComponent}></Route> */}
                            <Route path = "/quotes" component = {ListQuoteComponent}></Route>
                            <Route path = "/add-quote/:id" component = {CreateQuoteComponent}></Route>
                            <Route path = "/view-quote/:id" component = {ViewQuoteComponent}></Route>
                          {/* <Route path = "/update-quote/:id" component = {UpdateQuoteComponent}></Route> */}
                            <Route path = "/purchaseAgreements" component = {ListPurchaseAgreementComponent}></Route>
                            <Route path = "/add-purchaseAgreement/:id" component = {CreatePurchaseAgreementComponent}></Route>
                            <Route path = "/view-purchaseAgreement/:id" component = {ViewPurchaseAgreementComponent}></Route>
                          {/* <Route path = "/update-purchaseAgreement/:id" component = {UpdatePurchaseAgreementComponent}></Route> */}
                            <Route path = "/aircrafts" component = {ListAircraftComponent}></Route>
                            <Route path = "/add-aircraft/:id" component = {CreateAircraftComponent}></Route>
                            <Route path = "/view-aircraft/:id" component = {ViewAircraftComponent}></Route>
                          {/* <Route path = "/update-aircraft/:id" component = {UpdateAircraftComponent}></Route> */}
                            <Route path = "/registrations" component = {ListRegistrationComponent}></Route>
                            <Route path = "/add-registration/:id" component = {CreateRegistrationComponent}></Route>
                            <Route path = "/view-registration/:id" component = {ViewRegistrationComponent}></Route>
                          {/* <Route path = "/update-registration/:id" component = {UpdateRegistrationComponent}></Route> */}
                            <Route path = "/warrantys" component = {ListWarrantyComponent}></Route>
                            <Route path = "/add-warranty/:id" component = {CreateWarrantyComponent}></Route>
                            <Route path = "/view-warranty/:id" component = {ViewWarrantyComponent}></Route>
                          {/* <Route path = "/update-warranty/:id" component = {UpdateWarrantyComponent}></Route> */}
                            <Route path = "/cabinLayouts" component = {ListCabinLayoutComponent}></Route>
                            <Route path = "/add-cabinLayout/:id" component = {CreateCabinLayoutComponent}></Route>
                            <Route path = "/view-cabinLayout/:id" component = {ViewCabinLayoutComponent}></Route>
                          {/* <Route path = "/update-cabinLayout/:id" component = {UpdateCabinLayoutComponent}></Route> */}
                            <Route path = "/mROFacilitys" component = {ListMROFacilityComponent}></Route>
                            <Route path = "/add-mROFacility/:id" component = {CreateMROFacilityComponent}></Route>
                            <Route path = "/view-mROFacility/:id" component = {ViewMROFacilityComponent}></Route>
                          {/* <Route path = "/update-mROFacility/:id" component = {UpdateMROFacilityComponent}></Route> */}
                            <Route path = "/maintenanceAppointments" component = {ListMaintenanceAppointmentComponent}></Route>
                            <Route path = "/add-maintenanceAppointment/:id" component = {CreateMaintenanceAppointmentComponent}></Route>
                            <Route path = "/view-maintenanceAppointment/:id" component = {ViewMaintenanceAppointmentComponent}></Route>
                          {/* <Route path = "/update-maintenanceAppointment/:id" component = {UpdateMaintenanceAppointmentComponent}></Route> */}
                            <Route path = "/maintenanceWorkOrders" component = {ListMaintenanceWorkOrderComponent}></Route>
                            <Route path = "/add-maintenanceWorkOrder/:id" component = {CreateMaintenanceWorkOrderComponent}></Route>
                            <Route path = "/view-maintenanceWorkOrder/:id" component = {ViewMaintenanceWorkOrderComponent}></Route>
                          {/* <Route path = "/update-maintenanceWorkOrder/:id" component = {UpdateMaintenanceWorkOrderComponent}></Route> */}
                            <Route path = "/airworthinessDirectives" component = {ListAirworthinessDirectiveComponent}></Route>
                            <Route path = "/add-airworthinessDirective/:id" component = {CreateAirworthinessDirectiveComponent}></Route>
                            <Route path = "/view-airworthinessDirective/:id" component = {ViewAirworthinessDirectiveComponent}></Route>
                          {/* <Route path = "/update-airworthinessDirective/:id" component = {UpdateAirworthinessDirectiveComponent}></Route> */}
                            <Route path = "/serviceBulletins" component = {ListServiceBulletinComponent}></Route>
                            <Route path = "/add-serviceBulletin/:id" component = {CreateServiceBulletinComponent}></Route>
                            <Route path = "/view-serviceBulletin/:id" component = {ViewServiceBulletinComponent}></Route>
                          {/* <Route path = "/update-serviceBulletin/:id" component = {UpdateServiceBulletinComponent}></Route> */}
                            <Route path = "/connectedAircrafts" component = {ListConnectedAircraftComponent}></Route>
                            <Route path = "/add-connectedAircraft/:id" component = {CreateConnectedAircraftComponent}></Route>
                            <Route path = "/view-connectedAircraft/:id" component = {ViewConnectedAircraftComponent}></Route>
                          {/* <Route path = "/update-connectedAircraft/:id" component = {UpdateConnectedAircraftComponent}></Route> */}
                            <Route path = "/flightHealthEvents" component = {ListFlightHealthEventComponent}></Route>
                            <Route path = "/add-flightHealthEvent/:id" component = {CreateFlightHealthEventComponent}></Route>
                            <Route path = "/view-flightHealthEvent/:id" component = {ViewFlightHealthEventComponent}></Route>
                          {/* <Route path = "/update-flightHealthEvent/:id" component = {UpdateFlightHealthEventComponent}></Route> */}
                            <Route path = "/softwareLoads" component = {ListSoftwareLoadComponent}></Route>
                            <Route path = "/add-softwareLoad/:id" component = {CreateSoftwareLoadComponent}></Route>
                            <Route path = "/view-softwareLoad/:id" component = {ViewSoftwareLoadComponent}></Route>
                          {/* <Route path = "/update-softwareLoad/:id" component = {UpdateSoftwareLoadComponent}></Route> */}
                            <Route path = "/typeCertificates" component = {ListTypeCertificateComponent}></Route>
                            <Route path = "/add-typeCertificate/:id" component = {CreateTypeCertificateComponent}></Route>
                            <Route path = "/view-typeCertificate/:id" component = {ViewTypeCertificateComponent}></Route>
                          {/* <Route path = "/update-typeCertificate/:id" component = {UpdateTypeCertificateComponent}></Route> */}
                            <Route path = "/productionCertificates" component = {ListProductionCertificateComponent}></Route>
                            <Route path = "/add-productionCertificate/:id" component = {CreateProductionCertificateComponent}></Route>
                            <Route path = "/view-productionCertificate/:id" component = {ViewProductionCertificateComponent}></Route>
                          {/* <Route path = "/update-productionCertificate/:id" component = {UpdateProductionCertificateComponent}></Route> */}
                            <Route path = "/salesRegions" component = {ListSalesRegionComponent}></Route>
                            <Route path = "/add-salesRegion/:id" component = {CreateSalesRegionComponent}></Route>
                            <Route path = "/view-salesRegion/:id" component = {ViewSalesRegionComponent}></Route>
                          {/* <Route path = "/update-salesRegion/:id" component = {UpdateSalesRegionComponent}></Route> */}
                            <Route path = "/salesCampaigns" component = {ListSalesCampaignComponent}></Route>
                            <Route path = "/add-salesCampaign/:id" component = {CreateSalesCampaignComponent}></Route>
                            <Route path = "/view-salesCampaign/:id" component = {ViewSalesCampaignComponent}></Route>
                          {/* <Route path = "/update-salesCampaign/:id" component = {UpdateSalesCampaignComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
