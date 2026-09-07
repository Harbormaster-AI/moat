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

import {IndexAerospaceManufacturerComponent} from './components/AerospaceManufacturer/index/index.component';
import {CreateAerospaceManufacturerComponent} from './components/AerospaceManufacturer/create/create.component';
import {EditAerospaceManufacturerComponent} from './components/AerospaceManufacturer/edit/edit.component';
import {IndexAircraftProgramComponent} from './components/AircraftProgram/index/index.component';
import {CreateAircraftProgramComponent} from './components/AircraftProgram/create/create.component';
import {EditAircraftProgramComponent} from './components/AircraftProgram/edit/edit.component';
import {IndexAircraftFamilyComponent} from './components/AircraftFamily/index/index.component';
import {CreateAircraftFamilyComponent} from './components/AircraftFamily/create/create.component';
import {EditAircraftFamilyComponent} from './components/AircraftFamily/edit/edit.component';
import {IndexAircraftModelComponent} from './components/AircraftModel/index/index.component';
import {CreateAircraftModelComponent} from './components/AircraftModel/create/create.component';
import {EditAircraftModelComponent} from './components/AircraftModel/edit/edit.component';
import {IndexEngineTypeComponent} from './components/EngineType/index/index.component';
import {CreateEngineTypeComponent} from './components/EngineType/create/create.component';
import {EditEngineTypeComponent} from './components/EngineType/edit/edit.component';
import {IndexAircraftVariantComponent} from './components/AircraftVariant/index/index.component';
import {CreateAircraftVariantComponent} from './components/AircraftVariant/create/create.component';
import {EditAircraftVariantComponent} from './components/AircraftVariant/edit/edit.component';
import {IndexAvionicsSuiteComponent} from './components/AvionicsSuite/index/index.component';
import {CreateAvionicsSuiteComponent} from './components/AvionicsSuite/create/create.component';
import {EditAvionicsSuiteComponent} from './components/AvionicsSuite/edit/edit.component';
import {IndexAPUComponent} from './components/APU/index/index.component';
import {CreateAPUComponent} from './components/APU/create/create.component';
import {EditAPUComponent} from './components/APU/edit/edit.component';
import {IndexLandingGearComponent} from './components/LandingGear/index/index.component';
import {CreateLandingGearComponent} from './components/LandingGear/create/create.component';
import {EditLandingGearComponent} from './components/LandingGear/edit/edit.component';
import {IndexAircraftOptionComponent} from './components/AircraftOption/index/index.component';
import {CreateAircraftOptionComponent} from './components/AircraftOption/create/create.component';
import {EditAircraftOptionComponent} from './components/AircraftOption/edit/edit.component';
import {IndexAircraftPackageComponent} from './components/AircraftPackage/index/index.component';
import {CreateAircraftPackageComponent} from './components/AircraftPackage/create/create.component';
import {EditAircraftPackageComponent} from './components/AircraftPackage/edit/edit.component';
import {IndexSupplierComponent} from './components/Supplier/index/index.component';
import {CreateSupplierComponent} from './components/Supplier/create/create.component';
import {EditSupplierComponent} from './components/Supplier/edit/edit.component';
import {IndexComponent_Component} from './components/Component_/index/index.component';
import {CreateComponent_Component} from './components/Component_/create/create.component';
import {EditComponent_Component} from './components/Component_/edit/edit.component';
import {IndexPlantComponent} from './components/Plant/index/index.component';
import {CreatePlantComponent} from './components/Plant/create/create.component';
import {EditPlantComponent} from './components/Plant/edit/edit.component';
import {IndexProductionLineComponent} from './components/ProductionLine/index/index.component';
import {CreateProductionLineComponent} from './components/ProductionLine/create/create.component';
import {EditProductionLineComponent} from './components/ProductionLine/edit/edit.component';
import {IndexWorkCenterComponent} from './components/WorkCenter/index/index.component';
import {CreateWorkCenterComponent} from './components/WorkCenter/create/create.component';
import {EditWorkCenterComponent} from './components/WorkCenter/edit/edit.component';
import {IndexProductionOrderComponent} from './components/ProductionOrder/index/index.component';
import {CreateProductionOrderComponent} from './components/ProductionOrder/create/create.component';
import {EditProductionOrderComponent} from './components/ProductionOrder/edit/edit.component';
import {IndexBuildScheduleComponent} from './components/BuildSchedule/index/index.component';
import {CreateBuildScheduleComponent} from './components/BuildSchedule/create/create.component';
import {EditBuildScheduleComponent} from './components/BuildSchedule/edit/edit.component';
import {IndexWarehouseComponent} from './components/Warehouse/index/index.component';
import {CreateWarehouseComponent} from './components/Warehouse/create/create.component';
import {EditWarehouseComponent} from './components/Warehouse/edit/edit.component';
import {IndexInventoryItemComponent} from './components/InventoryItem/index/index.component';
import {CreateInventoryItemComponent} from './components/InventoryItem/create/create.component';
import {EditInventoryItemComponent} from './components/InventoryItem/edit/edit.component';
import {IndexOperatorComponent} from './components/Operator/index/index.component';
import {CreateOperatorComponent} from './components/Operator/create/create.component';
import {EditOperatorComponent} from './components/Operator/edit/edit.component';
import {IndexAircraftOrderComponent} from './components/AircraftOrder/index/index.component';
import {CreateAircraftOrderComponent} from './components/AircraftOrder/create/create.component';
import {EditAircraftOrderComponent} from './components/AircraftOrder/edit/edit.component';
import {IndexQuoteComponent} from './components/Quote/index/index.component';
import {CreateQuoteComponent} from './components/Quote/create/create.component';
import {EditQuoteComponent} from './components/Quote/edit/edit.component';
import {IndexPurchaseAgreementComponent} from './components/PurchaseAgreement/index/index.component';
import {CreatePurchaseAgreementComponent} from './components/PurchaseAgreement/create/create.component';
import {EditPurchaseAgreementComponent} from './components/PurchaseAgreement/edit/edit.component';
import {IndexAircraftComponent} from './components/Aircraft/index/index.component';
import {CreateAircraftComponent} from './components/Aircraft/create/create.component';
import {EditAircraftComponent} from './components/Aircraft/edit/edit.component';
import {IndexRegistrationComponent} from './components/Registration/index/index.component';
import {CreateRegistrationComponent} from './components/Registration/create/create.component';
import {EditRegistrationComponent} from './components/Registration/edit/edit.component';
import {IndexWarrantyComponent} from './components/Warranty/index/index.component';
import {CreateWarrantyComponent} from './components/Warranty/create/create.component';
import {EditWarrantyComponent} from './components/Warranty/edit/edit.component';
import {IndexCabinLayoutComponent} from './components/CabinLayout/index/index.component';
import {CreateCabinLayoutComponent} from './components/CabinLayout/create/create.component';
import {EditCabinLayoutComponent} from './components/CabinLayout/edit/edit.component';
import {IndexMROFacilityComponent} from './components/MROFacility/index/index.component';
import {CreateMROFacilityComponent} from './components/MROFacility/create/create.component';
import {EditMROFacilityComponent} from './components/MROFacility/edit/edit.component';
import {IndexMaintenanceAppointmentComponent} from './components/MaintenanceAppointment/index/index.component';
import {CreateMaintenanceAppointmentComponent} from './components/MaintenanceAppointment/create/create.component';
import {EditMaintenanceAppointmentComponent} from './components/MaintenanceAppointment/edit/edit.component';
import {IndexMaintenanceWorkOrderComponent} from './components/MaintenanceWorkOrder/index/index.component';
import {CreateMaintenanceWorkOrderComponent} from './components/MaintenanceWorkOrder/create/create.component';
import {EditMaintenanceWorkOrderComponent} from './components/MaintenanceWorkOrder/edit/edit.component';
import {IndexAirworthinessDirectiveComponent} from './components/AirworthinessDirective/index/index.component';
import {CreateAirworthinessDirectiveComponent} from './components/AirworthinessDirective/create/create.component';
import {EditAirworthinessDirectiveComponent} from './components/AirworthinessDirective/edit/edit.component';
import {IndexServiceBulletinComponent} from './components/ServiceBulletin/index/index.component';
import {CreateServiceBulletinComponent} from './components/ServiceBulletin/create/create.component';
import {EditServiceBulletinComponent} from './components/ServiceBulletin/edit/edit.component';
import {IndexConnectedAircraftComponent} from './components/ConnectedAircraft/index/index.component';
import {CreateConnectedAircraftComponent} from './components/ConnectedAircraft/create/create.component';
import {EditConnectedAircraftComponent} from './components/ConnectedAircraft/edit/edit.component';
import {IndexFlightHealthEventComponent} from './components/FlightHealthEvent/index/index.component';
import {CreateFlightHealthEventComponent} from './components/FlightHealthEvent/create/create.component';
import {EditFlightHealthEventComponent} from './components/FlightHealthEvent/edit/edit.component';
import {IndexSoftwareLoadComponent} from './components/SoftwareLoad/index/index.component';
import {CreateSoftwareLoadComponent} from './components/SoftwareLoad/create/create.component';
import {EditSoftwareLoadComponent} from './components/SoftwareLoad/edit/edit.component';
import {IndexTypeCertificateComponent} from './components/TypeCertificate/index/index.component';
import {CreateTypeCertificateComponent} from './components/TypeCertificate/create/create.component';
import {EditTypeCertificateComponent} from './components/TypeCertificate/edit/edit.component';
import {IndexProductionCertificateComponent} from './components/ProductionCertificate/index/index.component';
import {CreateProductionCertificateComponent} from './components/ProductionCertificate/create/create.component';
import {EditProductionCertificateComponent} from './components/ProductionCertificate/edit/edit.component';
import {IndexSalesRegionComponent} from './components/SalesRegion/index/index.component';
import {CreateSalesRegionComponent} from './components/SalesRegion/create/create.component';
import {EditSalesRegionComponent} from './components/SalesRegion/edit/edit.component';
import {IndexSalesCampaignComponent} from './components/SalesCampaign/index/index.component';
import {CreateSalesCampaignComponent} from './components/SalesCampaign/create/create.component';
import {EditSalesCampaignComponent} from './components/SalesCampaign/edit/edit.component';

import * as appRoutes from './routerConfig';

import {AerospaceManufacturerService} from './services/AerospaceManufacturer.service';
import {AircraftProgramService} from './services/AircraftProgram.service';
import {AircraftFamilyService} from './services/AircraftFamily.service';
import {AircraftModelService} from './services/AircraftModel.service';
import {EngineTypeService} from './services/EngineType.service';
import {AircraftVariantService} from './services/AircraftVariant.service';
import {AvionicsSuiteService} from './services/AvionicsSuite.service';
import {APUService} from './services/APU.service';
import {LandingGearService} from './services/LandingGear.service';
import {AircraftOptionService} from './services/AircraftOption.service';
import {AircraftPackageService} from './services/AircraftPackage.service';
import {SupplierService} from './services/Supplier.service';
import {Component_Service} from './services/Component_.service';
import {PlantService} from './services/Plant.service';
import {ProductionLineService} from './services/ProductionLine.service';
import {WorkCenterService} from './services/WorkCenter.service';
import {ProductionOrderService} from './services/ProductionOrder.service';
import {BuildScheduleService} from './services/BuildSchedule.service';
import {WarehouseService} from './services/Warehouse.service';
import {InventoryItemService} from './services/InventoryItem.service';
import {OperatorService} from './services/Operator.service';
import {AircraftOrderService} from './services/AircraftOrder.service';
import {QuoteService} from './services/Quote.service';
import {PurchaseAgreementService} from './services/PurchaseAgreement.service';
import {AircraftService} from './services/Aircraft.service';
import {RegistrationService} from './services/Registration.service';
import {WarrantyService} from './services/Warranty.service';
import {CabinLayoutService} from './services/CabinLayout.service';
import {MROFacilityService} from './services/MROFacility.service';
import {MaintenanceAppointmentService} from './services/MaintenanceAppointment.service';
import {MaintenanceWorkOrderService} from './services/MaintenanceWorkOrder.service';
import {AirworthinessDirectiveService} from './services/AirworthinessDirective.service';
import {ServiceBulletinService} from './services/ServiceBulletin.service';
import {ConnectedAircraftService} from './services/ConnectedAircraft.service';
import {FlightHealthEventService} from './services/FlightHealthEvent.service';
import {SoftwareLoadService} from './services/SoftwareLoad.service';
import {TypeCertificateService} from './services/TypeCertificate.service';
import {ProductionCertificateService} from './services/ProductionCertificate.service';
import {SalesRegionService} from './services/SalesRegion.service';
import {SalesCampaignService} from './services/SalesCampaign.service';

@NgModule({
  declarations: [
    IndexAerospaceManufacturerComponent,
    CreateAerospaceManufacturerComponent,
    EditAerospaceManufacturerComponent,
    IndexAircraftProgramComponent,
    CreateAircraftProgramComponent,
    EditAircraftProgramComponent,
    IndexAircraftFamilyComponent,
    CreateAircraftFamilyComponent,
    EditAircraftFamilyComponent,
    IndexAircraftModelComponent,
    CreateAircraftModelComponent,
    EditAircraftModelComponent,
    IndexEngineTypeComponent,
    CreateEngineTypeComponent,
    EditEngineTypeComponent,
    IndexAircraftVariantComponent,
    CreateAircraftVariantComponent,
    EditAircraftVariantComponent,
    IndexAvionicsSuiteComponent,
    CreateAvionicsSuiteComponent,
    EditAvionicsSuiteComponent,
    IndexAPUComponent,
    CreateAPUComponent,
    EditAPUComponent,
    IndexLandingGearComponent,
    CreateLandingGearComponent,
    EditLandingGearComponent,
    IndexAircraftOptionComponent,
    CreateAircraftOptionComponent,
    EditAircraftOptionComponent,
    IndexAircraftPackageComponent,
    CreateAircraftPackageComponent,
    EditAircraftPackageComponent,
    IndexSupplierComponent,
    CreateSupplierComponent,
    EditSupplierComponent,
    IndexComponent_Component,
    CreateComponent_Component,
    EditComponent_Component,
    IndexPlantComponent,
    CreatePlantComponent,
    EditPlantComponent,
    IndexProductionLineComponent,
    CreateProductionLineComponent,
    EditProductionLineComponent,
    IndexWorkCenterComponent,
    CreateWorkCenterComponent,
    EditWorkCenterComponent,
    IndexProductionOrderComponent,
    CreateProductionOrderComponent,
    EditProductionOrderComponent,
    IndexBuildScheduleComponent,
    CreateBuildScheduleComponent,
    EditBuildScheduleComponent,
    IndexWarehouseComponent,
    CreateWarehouseComponent,
    EditWarehouseComponent,
    IndexInventoryItemComponent,
    CreateInventoryItemComponent,
    EditInventoryItemComponent,
    IndexOperatorComponent,
    CreateOperatorComponent,
    EditOperatorComponent,
    IndexAircraftOrderComponent,
    CreateAircraftOrderComponent,
    EditAircraftOrderComponent,
    IndexQuoteComponent,
    CreateQuoteComponent,
    EditQuoteComponent,
    IndexPurchaseAgreementComponent,
    CreatePurchaseAgreementComponent,
    EditPurchaseAgreementComponent,
    IndexAircraftComponent,
    CreateAircraftComponent,
    EditAircraftComponent,
    IndexRegistrationComponent,
    CreateRegistrationComponent,
    EditRegistrationComponent,
    IndexWarrantyComponent,
    CreateWarrantyComponent,
    EditWarrantyComponent,
    IndexCabinLayoutComponent,
    CreateCabinLayoutComponent,
    EditCabinLayoutComponent,
    IndexMROFacilityComponent,
    CreateMROFacilityComponent,
    EditMROFacilityComponent,
    IndexMaintenanceAppointmentComponent,
    CreateMaintenanceAppointmentComponent,
    EditMaintenanceAppointmentComponent,
    IndexMaintenanceWorkOrderComponent,
    CreateMaintenanceWorkOrderComponent,
    EditMaintenanceWorkOrderComponent,
    IndexAirworthinessDirectiveComponent,
    CreateAirworthinessDirectiveComponent,
    EditAirworthinessDirectiveComponent,
    IndexServiceBulletinComponent,
    CreateServiceBulletinComponent,
    EditServiceBulletinComponent,
    IndexConnectedAircraftComponent,
    CreateConnectedAircraftComponent,
    EditConnectedAircraftComponent,
    IndexFlightHealthEventComponent,
    CreateFlightHealthEventComponent,
    EditFlightHealthEventComponent,
    IndexSoftwareLoadComponent,
    CreateSoftwareLoadComponent,
    EditSoftwareLoadComponent,
    IndexTypeCertificateComponent,
    CreateTypeCertificateComponent,
    EditTypeCertificateComponent,
    IndexProductionCertificateComponent,
    CreateProductionCertificateComponent,
    EditProductionCertificateComponent,
    IndexSalesRegionComponent,
    CreateSalesRegionComponent,
    EditSalesRegionComponent,
    IndexSalesCampaignComponent,
    CreateSalesCampaignComponent,
    EditSalesCampaignComponent,
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
    RouterModule.forRoot(appRoutes.AerospaceManufacturerRoutes), 
    RouterModule.forRoot(appRoutes.AircraftProgramRoutes), 
    RouterModule.forRoot(appRoutes.AircraftFamilyRoutes), 
    RouterModule.forRoot(appRoutes.AircraftModelRoutes), 
    RouterModule.forRoot(appRoutes.EngineTypeRoutes), 
    RouterModule.forRoot(appRoutes.AircraftVariantRoutes), 
    RouterModule.forRoot(appRoutes.AvionicsSuiteRoutes), 
    RouterModule.forRoot(appRoutes.APURoutes), 
    RouterModule.forRoot(appRoutes.LandingGearRoutes), 
    RouterModule.forRoot(appRoutes.AircraftOptionRoutes), 
    RouterModule.forRoot(appRoutes.AircraftPackageRoutes), 
    RouterModule.forRoot(appRoutes.SupplierRoutes), 
    RouterModule.forRoot(appRoutes.Component_Routes), 
    RouterModule.forRoot(appRoutes.PlantRoutes), 
    RouterModule.forRoot(appRoutes.ProductionLineRoutes), 
    RouterModule.forRoot(appRoutes.WorkCenterRoutes), 
    RouterModule.forRoot(appRoutes.ProductionOrderRoutes), 
    RouterModule.forRoot(appRoutes.BuildScheduleRoutes), 
    RouterModule.forRoot(appRoutes.WarehouseRoutes), 
    RouterModule.forRoot(appRoutes.InventoryItemRoutes), 
    RouterModule.forRoot(appRoutes.OperatorRoutes), 
    RouterModule.forRoot(appRoutes.AircraftOrderRoutes), 
    RouterModule.forRoot(appRoutes.QuoteRoutes), 
    RouterModule.forRoot(appRoutes.PurchaseAgreementRoutes), 
    RouterModule.forRoot(appRoutes.AircraftRoutes), 
    RouterModule.forRoot(appRoutes.RegistrationRoutes), 
    RouterModule.forRoot(appRoutes.WarrantyRoutes), 
    RouterModule.forRoot(appRoutes.CabinLayoutRoutes), 
    RouterModule.forRoot(appRoutes.MROFacilityRoutes), 
    RouterModule.forRoot(appRoutes.MaintenanceAppointmentRoutes), 
    RouterModule.forRoot(appRoutes.MaintenanceWorkOrderRoutes), 
    RouterModule.forRoot(appRoutes.AirworthinessDirectiveRoutes), 
    RouterModule.forRoot(appRoutes.ServiceBulletinRoutes), 
    RouterModule.forRoot(appRoutes.ConnectedAircraftRoutes), 
    RouterModule.forRoot(appRoutes.FlightHealthEventRoutes), 
    RouterModule.forRoot(appRoutes.SoftwareLoadRoutes), 
    RouterModule.forRoot(appRoutes.TypeCertificateRoutes), 
    RouterModule.forRoot(appRoutes.ProductionCertificateRoutes), 
    RouterModule.forRoot(appRoutes.SalesRegionRoutes), 
    RouterModule.forRoot(appRoutes.SalesCampaignRoutes), 
  ],
  providers: [AerospaceManufacturerService,AircraftProgramService,AircraftFamilyService,AircraftModelService,EngineTypeService,AircraftVariantService,AvionicsSuiteService,APUService,LandingGearService,AircraftOptionService,AircraftPackageService,SupplierService,Component_Service,PlantService,ProductionLineService,WorkCenterService,ProductionOrderService,BuildScheduleService,WarehouseService,InventoryItemService,OperatorService,AircraftOrderService,QuoteService,PurchaseAgreementService,AircraftService,RegistrationService,WarrantyService,CabinLayoutService,MROFacilityService,MaintenanceAppointmentService,MaintenanceWorkOrderService,AirworthinessDirectiveService,ServiceBulletinService,ConnectedAircraftService,FlightHealthEventService,SoftwareLoadService,TypeCertificateService,ProductionCertificateService,SalesRegionService,SalesCampaignService],
  bootstrap: [AppComponent]
})
export class AppModule { }
