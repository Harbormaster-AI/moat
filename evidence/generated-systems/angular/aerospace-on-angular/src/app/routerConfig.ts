// routerConfig.ts

import { Routes } from '@angular/router';
import { CreateAerospaceManufacturerComponent } from './components/AerospaceManufacturer/create/create.component';
import { EditAerospaceManufacturerComponent } from './components/AerospaceManufacturer/edit/edit.component';
import { IndexAerospaceManufacturerComponent } from './components/AerospaceManufacturer/index/index.component';
import { CreateAircraftProgramComponent } from './components/AircraftProgram/create/create.component';
import { EditAircraftProgramComponent } from './components/AircraftProgram/edit/edit.component';
import { IndexAircraftProgramComponent } from './components/AircraftProgram/index/index.component';
import { CreateAircraftFamilyComponent } from './components/AircraftFamily/create/create.component';
import { EditAircraftFamilyComponent } from './components/AircraftFamily/edit/edit.component';
import { IndexAircraftFamilyComponent } from './components/AircraftFamily/index/index.component';
import { CreateAircraftModelComponent } from './components/AircraftModel/create/create.component';
import { EditAircraftModelComponent } from './components/AircraftModel/edit/edit.component';
import { IndexAircraftModelComponent } from './components/AircraftModel/index/index.component';
import { CreateEngineTypeComponent } from './components/EngineType/create/create.component';
import { EditEngineTypeComponent } from './components/EngineType/edit/edit.component';
import { IndexEngineTypeComponent } from './components/EngineType/index/index.component';
import { CreateAircraftVariantComponent } from './components/AircraftVariant/create/create.component';
import { EditAircraftVariantComponent } from './components/AircraftVariant/edit/edit.component';
import { IndexAircraftVariantComponent } from './components/AircraftVariant/index/index.component';
import { CreateAvionicsSuiteComponent } from './components/AvionicsSuite/create/create.component';
import { EditAvionicsSuiteComponent } from './components/AvionicsSuite/edit/edit.component';
import { IndexAvionicsSuiteComponent } from './components/AvionicsSuite/index/index.component';
import { CreateAPUComponent } from './components/APU/create/create.component';
import { EditAPUComponent } from './components/APU/edit/edit.component';
import { IndexAPUComponent } from './components/APU/index/index.component';
import { CreateLandingGearComponent } from './components/LandingGear/create/create.component';
import { EditLandingGearComponent } from './components/LandingGear/edit/edit.component';
import { IndexLandingGearComponent } from './components/LandingGear/index/index.component';
import { CreateAircraftOptionComponent } from './components/AircraftOption/create/create.component';
import { EditAircraftOptionComponent } from './components/AircraftOption/edit/edit.component';
import { IndexAircraftOptionComponent } from './components/AircraftOption/index/index.component';
import { CreateAircraftPackageComponent } from './components/AircraftPackage/create/create.component';
import { EditAircraftPackageComponent } from './components/AircraftPackage/edit/edit.component';
import { IndexAircraftPackageComponent } from './components/AircraftPackage/index/index.component';
import { CreateSupplierComponent } from './components/Supplier/create/create.component';
import { EditSupplierComponent } from './components/Supplier/edit/edit.component';
import { IndexSupplierComponent } from './components/Supplier/index/index.component';
import { CreateComponent_Component } from './components/Component_/create/create.component';
import { EditComponent_Component } from './components/Component_/edit/edit.component';
import { IndexComponent_Component } from './components/Component_/index/index.component';
import { CreatePlantComponent } from './components/Plant/create/create.component';
import { EditPlantComponent } from './components/Plant/edit/edit.component';
import { IndexPlantComponent } from './components/Plant/index/index.component';
import { CreateProductionLineComponent } from './components/ProductionLine/create/create.component';
import { EditProductionLineComponent } from './components/ProductionLine/edit/edit.component';
import { IndexProductionLineComponent } from './components/ProductionLine/index/index.component';
import { CreateWorkCenterComponent } from './components/WorkCenter/create/create.component';
import { EditWorkCenterComponent } from './components/WorkCenter/edit/edit.component';
import { IndexWorkCenterComponent } from './components/WorkCenter/index/index.component';
import { CreateProductionOrderComponent } from './components/ProductionOrder/create/create.component';
import { EditProductionOrderComponent } from './components/ProductionOrder/edit/edit.component';
import { IndexProductionOrderComponent } from './components/ProductionOrder/index/index.component';
import { CreateBuildScheduleComponent } from './components/BuildSchedule/create/create.component';
import { EditBuildScheduleComponent } from './components/BuildSchedule/edit/edit.component';
import { IndexBuildScheduleComponent } from './components/BuildSchedule/index/index.component';
import { CreateWarehouseComponent } from './components/Warehouse/create/create.component';
import { EditWarehouseComponent } from './components/Warehouse/edit/edit.component';
import { IndexWarehouseComponent } from './components/Warehouse/index/index.component';
import { CreateInventoryItemComponent } from './components/InventoryItem/create/create.component';
import { EditInventoryItemComponent } from './components/InventoryItem/edit/edit.component';
import { IndexInventoryItemComponent } from './components/InventoryItem/index/index.component';
import { CreateOperatorComponent } from './components/Operator/create/create.component';
import { EditOperatorComponent } from './components/Operator/edit/edit.component';
import { IndexOperatorComponent } from './components/Operator/index/index.component';
import { CreateAircraftOrderComponent } from './components/AircraftOrder/create/create.component';
import { EditAircraftOrderComponent } from './components/AircraftOrder/edit/edit.component';
import { IndexAircraftOrderComponent } from './components/AircraftOrder/index/index.component';
import { CreateQuoteComponent } from './components/Quote/create/create.component';
import { EditQuoteComponent } from './components/Quote/edit/edit.component';
import { IndexQuoteComponent } from './components/Quote/index/index.component';
import { CreatePurchaseAgreementComponent } from './components/PurchaseAgreement/create/create.component';
import { EditPurchaseAgreementComponent } from './components/PurchaseAgreement/edit/edit.component';
import { IndexPurchaseAgreementComponent } from './components/PurchaseAgreement/index/index.component';
import { CreateAircraftComponent } from './components/Aircraft/create/create.component';
import { EditAircraftComponent } from './components/Aircraft/edit/edit.component';
import { IndexAircraftComponent } from './components/Aircraft/index/index.component';
import { CreateRegistrationComponent } from './components/Registration/create/create.component';
import { EditRegistrationComponent } from './components/Registration/edit/edit.component';
import { IndexRegistrationComponent } from './components/Registration/index/index.component';
import { CreateWarrantyComponent } from './components/Warranty/create/create.component';
import { EditWarrantyComponent } from './components/Warranty/edit/edit.component';
import { IndexWarrantyComponent } from './components/Warranty/index/index.component';
import { CreateCabinLayoutComponent } from './components/CabinLayout/create/create.component';
import { EditCabinLayoutComponent } from './components/CabinLayout/edit/edit.component';
import { IndexCabinLayoutComponent } from './components/CabinLayout/index/index.component';
import { CreateMROFacilityComponent } from './components/MROFacility/create/create.component';
import { EditMROFacilityComponent } from './components/MROFacility/edit/edit.component';
import { IndexMROFacilityComponent } from './components/MROFacility/index/index.component';
import { CreateMaintenanceAppointmentComponent } from './components/MaintenanceAppointment/create/create.component';
import { EditMaintenanceAppointmentComponent } from './components/MaintenanceAppointment/edit/edit.component';
import { IndexMaintenanceAppointmentComponent } from './components/MaintenanceAppointment/index/index.component';
import { CreateMaintenanceWorkOrderComponent } from './components/MaintenanceWorkOrder/create/create.component';
import { EditMaintenanceWorkOrderComponent } from './components/MaintenanceWorkOrder/edit/edit.component';
import { IndexMaintenanceWorkOrderComponent } from './components/MaintenanceWorkOrder/index/index.component';
import { CreateAirworthinessDirectiveComponent } from './components/AirworthinessDirective/create/create.component';
import { EditAirworthinessDirectiveComponent } from './components/AirworthinessDirective/edit/edit.component';
import { IndexAirworthinessDirectiveComponent } from './components/AirworthinessDirective/index/index.component';
import { CreateServiceBulletinComponent } from './components/ServiceBulletin/create/create.component';
import { EditServiceBulletinComponent } from './components/ServiceBulletin/edit/edit.component';
import { IndexServiceBulletinComponent } from './components/ServiceBulletin/index/index.component';
import { CreateConnectedAircraftComponent } from './components/ConnectedAircraft/create/create.component';
import { EditConnectedAircraftComponent } from './components/ConnectedAircraft/edit/edit.component';
import { IndexConnectedAircraftComponent } from './components/ConnectedAircraft/index/index.component';
import { CreateFlightHealthEventComponent } from './components/FlightHealthEvent/create/create.component';
import { EditFlightHealthEventComponent } from './components/FlightHealthEvent/edit/edit.component';
import { IndexFlightHealthEventComponent } from './components/FlightHealthEvent/index/index.component';
import { CreateSoftwareLoadComponent } from './components/SoftwareLoad/create/create.component';
import { EditSoftwareLoadComponent } from './components/SoftwareLoad/edit/edit.component';
import { IndexSoftwareLoadComponent } from './components/SoftwareLoad/index/index.component';
import { CreateTypeCertificateComponent } from './components/TypeCertificate/create/create.component';
import { EditTypeCertificateComponent } from './components/TypeCertificate/edit/edit.component';
import { IndexTypeCertificateComponent } from './components/TypeCertificate/index/index.component';
import { CreateProductionCertificateComponent } from './components/ProductionCertificate/create/create.component';
import { EditProductionCertificateComponent } from './components/ProductionCertificate/edit/edit.component';
import { IndexProductionCertificateComponent } from './components/ProductionCertificate/index/index.component';
import { CreateSalesRegionComponent } from './components/SalesRegion/create/create.component';
import { EditSalesRegionComponent } from './components/SalesRegion/edit/edit.component';
import { IndexSalesRegionComponent } from './components/SalesRegion/index/index.component';
import { CreateSalesCampaignComponent } from './components/SalesCampaign/create/create.component';
import { EditSalesCampaignComponent } from './components/SalesCampaign/edit/edit.component';
import { IndexSalesCampaignComponent } from './components/SalesCampaign/index/index.component';

export const AerospaceManufacturerRoutes: Routes = [
  { path: 'createAerospaceManufacturer',
    component: CreateAerospaceManufacturerComponent
  },
  {
    path: 'editAerospaceManufacturer/:id',
    component: EditAerospaceManufacturerComponent
  },
  { path: 'indexAerospaceManufacturer',
    component: IndexAerospaceManufacturerComponent
  }
];
export const AircraftProgramRoutes: Routes = [
  { path: 'createAircraftProgram',
    component: CreateAircraftProgramComponent
  },
  {
    path: 'editAircraftProgram/:id',
    component: EditAircraftProgramComponent
  },
  { path: 'indexAircraftProgram',
    component: IndexAircraftProgramComponent
  }
];
export const AircraftFamilyRoutes: Routes = [
  { path: 'createAircraftFamily',
    component: CreateAircraftFamilyComponent
  },
  {
    path: 'editAircraftFamily/:id',
    component: EditAircraftFamilyComponent
  },
  { path: 'indexAircraftFamily',
    component: IndexAircraftFamilyComponent
  }
];
export const AircraftModelRoutes: Routes = [
  { path: 'createAircraftModel',
    component: CreateAircraftModelComponent
  },
  {
    path: 'editAircraftModel/:id',
    component: EditAircraftModelComponent
  },
  { path: 'indexAircraftModel',
    component: IndexAircraftModelComponent
  }
];
export const EngineTypeRoutes: Routes = [
  { path: 'createEngineType',
    component: CreateEngineTypeComponent
  },
  {
    path: 'editEngineType/:id',
    component: EditEngineTypeComponent
  },
  { path: 'indexEngineType',
    component: IndexEngineTypeComponent
  }
];
export const AircraftVariantRoutes: Routes = [
  { path: 'createAircraftVariant',
    component: CreateAircraftVariantComponent
  },
  {
    path: 'editAircraftVariant/:id',
    component: EditAircraftVariantComponent
  },
  { path: 'indexAircraftVariant',
    component: IndexAircraftVariantComponent
  }
];
export const AvionicsSuiteRoutes: Routes = [
  { path: 'createAvionicsSuite',
    component: CreateAvionicsSuiteComponent
  },
  {
    path: 'editAvionicsSuite/:id',
    component: EditAvionicsSuiteComponent
  },
  { path: 'indexAvionicsSuite',
    component: IndexAvionicsSuiteComponent
  }
];
export const APURoutes: Routes = [
  { path: 'createAPU',
    component: CreateAPUComponent
  },
  {
    path: 'editAPU/:id',
    component: EditAPUComponent
  },
  { path: 'indexAPU',
    component: IndexAPUComponent
  }
];
export const LandingGearRoutes: Routes = [
  { path: 'createLandingGear',
    component: CreateLandingGearComponent
  },
  {
    path: 'editLandingGear/:id',
    component: EditLandingGearComponent
  },
  { path: 'indexLandingGear',
    component: IndexLandingGearComponent
  }
];
export const AircraftOptionRoutes: Routes = [
  { path: 'createAircraftOption',
    component: CreateAircraftOptionComponent
  },
  {
    path: 'editAircraftOption/:id',
    component: EditAircraftOptionComponent
  },
  { path: 'indexAircraftOption',
    component: IndexAircraftOptionComponent
  }
];
export const AircraftPackageRoutes: Routes = [
  { path: 'createAircraftPackage',
    component: CreateAircraftPackageComponent
  },
  {
    path: 'editAircraftPackage/:id',
    component: EditAircraftPackageComponent
  },
  { path: 'indexAircraftPackage',
    component: IndexAircraftPackageComponent
  }
];
export const SupplierRoutes: Routes = [
  { path: 'createSupplier',
    component: CreateSupplierComponent
  },
  {
    path: 'editSupplier/:id',
    component: EditSupplierComponent
  },
  { path: 'indexSupplier',
    component: IndexSupplierComponent
  }
];
export const Component_Routes: Routes = [
  { path: 'createComponent_',
    component: CreateComponent_Component
  },
  {
    path: 'editComponent_/:id',
    component: EditComponent_Component
  },
  { path: 'indexComponent_',
    component: IndexComponent_Component
  }
];
export const PlantRoutes: Routes = [
  { path: 'createPlant',
    component: CreatePlantComponent
  },
  {
    path: 'editPlant/:id',
    component: EditPlantComponent
  },
  { path: 'indexPlant',
    component: IndexPlantComponent
  }
];
export const ProductionLineRoutes: Routes = [
  { path: 'createProductionLine',
    component: CreateProductionLineComponent
  },
  {
    path: 'editProductionLine/:id',
    component: EditProductionLineComponent
  },
  { path: 'indexProductionLine',
    component: IndexProductionLineComponent
  }
];
export const WorkCenterRoutes: Routes = [
  { path: 'createWorkCenter',
    component: CreateWorkCenterComponent
  },
  {
    path: 'editWorkCenter/:id',
    component: EditWorkCenterComponent
  },
  { path: 'indexWorkCenter',
    component: IndexWorkCenterComponent
  }
];
export const ProductionOrderRoutes: Routes = [
  { path: 'createProductionOrder',
    component: CreateProductionOrderComponent
  },
  {
    path: 'editProductionOrder/:id',
    component: EditProductionOrderComponent
  },
  { path: 'indexProductionOrder',
    component: IndexProductionOrderComponent
  }
];
export const BuildScheduleRoutes: Routes = [
  { path: 'createBuildSchedule',
    component: CreateBuildScheduleComponent
  },
  {
    path: 'editBuildSchedule/:id',
    component: EditBuildScheduleComponent
  },
  { path: 'indexBuildSchedule',
    component: IndexBuildScheduleComponent
  }
];
export const WarehouseRoutes: Routes = [
  { path: 'createWarehouse',
    component: CreateWarehouseComponent
  },
  {
    path: 'editWarehouse/:id',
    component: EditWarehouseComponent
  },
  { path: 'indexWarehouse',
    component: IndexWarehouseComponent
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
export const OperatorRoutes: Routes = [
  { path: 'createOperator',
    component: CreateOperatorComponent
  },
  {
    path: 'editOperator/:id',
    component: EditOperatorComponent
  },
  { path: 'indexOperator',
    component: IndexOperatorComponent
  }
];
export const AircraftOrderRoutes: Routes = [
  { path: 'createAircraftOrder',
    component: CreateAircraftOrderComponent
  },
  {
    path: 'editAircraftOrder/:id',
    component: EditAircraftOrderComponent
  },
  { path: 'indexAircraftOrder',
    component: IndexAircraftOrderComponent
  }
];
export const QuoteRoutes: Routes = [
  { path: 'createQuote',
    component: CreateQuoteComponent
  },
  {
    path: 'editQuote/:id',
    component: EditQuoteComponent
  },
  { path: 'indexQuote',
    component: IndexQuoteComponent
  }
];
export const PurchaseAgreementRoutes: Routes = [
  { path: 'createPurchaseAgreement',
    component: CreatePurchaseAgreementComponent
  },
  {
    path: 'editPurchaseAgreement/:id',
    component: EditPurchaseAgreementComponent
  },
  { path: 'indexPurchaseAgreement',
    component: IndexPurchaseAgreementComponent
  }
];
export const AircraftRoutes: Routes = [
  { path: 'createAircraft',
    component: CreateAircraftComponent
  },
  {
    path: 'editAircraft/:id',
    component: EditAircraftComponent
  },
  { path: 'indexAircraft',
    component: IndexAircraftComponent
  }
];
export const RegistrationRoutes: Routes = [
  { path: 'createRegistration',
    component: CreateRegistrationComponent
  },
  {
    path: 'editRegistration/:id',
    component: EditRegistrationComponent
  },
  { path: 'indexRegistration',
    component: IndexRegistrationComponent
  }
];
export const WarrantyRoutes: Routes = [
  { path: 'createWarranty',
    component: CreateWarrantyComponent
  },
  {
    path: 'editWarranty/:id',
    component: EditWarrantyComponent
  },
  { path: 'indexWarranty',
    component: IndexWarrantyComponent
  }
];
export const CabinLayoutRoutes: Routes = [
  { path: 'createCabinLayout',
    component: CreateCabinLayoutComponent
  },
  {
    path: 'editCabinLayout/:id',
    component: EditCabinLayoutComponent
  },
  { path: 'indexCabinLayout',
    component: IndexCabinLayoutComponent
  }
];
export const MROFacilityRoutes: Routes = [
  { path: 'createMROFacility',
    component: CreateMROFacilityComponent
  },
  {
    path: 'editMROFacility/:id',
    component: EditMROFacilityComponent
  },
  { path: 'indexMROFacility',
    component: IndexMROFacilityComponent
  }
];
export const MaintenanceAppointmentRoutes: Routes = [
  { path: 'createMaintenanceAppointment',
    component: CreateMaintenanceAppointmentComponent
  },
  {
    path: 'editMaintenanceAppointment/:id',
    component: EditMaintenanceAppointmentComponent
  },
  { path: 'indexMaintenanceAppointment',
    component: IndexMaintenanceAppointmentComponent
  }
];
export const MaintenanceWorkOrderRoutes: Routes = [
  { path: 'createMaintenanceWorkOrder',
    component: CreateMaintenanceWorkOrderComponent
  },
  {
    path: 'editMaintenanceWorkOrder/:id',
    component: EditMaintenanceWorkOrderComponent
  },
  { path: 'indexMaintenanceWorkOrder',
    component: IndexMaintenanceWorkOrderComponent
  }
];
export const AirworthinessDirectiveRoutes: Routes = [
  { path: 'createAirworthinessDirective',
    component: CreateAirworthinessDirectiveComponent
  },
  {
    path: 'editAirworthinessDirective/:id',
    component: EditAirworthinessDirectiveComponent
  },
  { path: 'indexAirworthinessDirective',
    component: IndexAirworthinessDirectiveComponent
  }
];
export const ServiceBulletinRoutes: Routes = [
  { path: 'createServiceBulletin',
    component: CreateServiceBulletinComponent
  },
  {
    path: 'editServiceBulletin/:id',
    component: EditServiceBulletinComponent
  },
  { path: 'indexServiceBulletin',
    component: IndexServiceBulletinComponent
  }
];
export const ConnectedAircraftRoutes: Routes = [
  { path: 'createConnectedAircraft',
    component: CreateConnectedAircraftComponent
  },
  {
    path: 'editConnectedAircraft/:id',
    component: EditConnectedAircraftComponent
  },
  { path: 'indexConnectedAircraft',
    component: IndexConnectedAircraftComponent
  }
];
export const FlightHealthEventRoutes: Routes = [
  { path: 'createFlightHealthEvent',
    component: CreateFlightHealthEventComponent
  },
  {
    path: 'editFlightHealthEvent/:id',
    component: EditFlightHealthEventComponent
  },
  { path: 'indexFlightHealthEvent',
    component: IndexFlightHealthEventComponent
  }
];
export const SoftwareLoadRoutes: Routes = [
  { path: 'createSoftwareLoad',
    component: CreateSoftwareLoadComponent
  },
  {
    path: 'editSoftwareLoad/:id',
    component: EditSoftwareLoadComponent
  },
  { path: 'indexSoftwareLoad',
    component: IndexSoftwareLoadComponent
  }
];
export const TypeCertificateRoutes: Routes = [
  { path: 'createTypeCertificate',
    component: CreateTypeCertificateComponent
  },
  {
    path: 'editTypeCertificate/:id',
    component: EditTypeCertificateComponent
  },
  { path: 'indexTypeCertificate',
    component: IndexTypeCertificateComponent
  }
];
export const ProductionCertificateRoutes: Routes = [
  { path: 'createProductionCertificate',
    component: CreateProductionCertificateComponent
  },
  {
    path: 'editProductionCertificate/:id',
    component: EditProductionCertificateComponent
  },
  { path: 'indexProductionCertificate',
    component: IndexProductionCertificateComponent
  }
];
export const SalesRegionRoutes: Routes = [
  { path: 'createSalesRegion',
    component: CreateSalesRegionComponent
  },
  {
    path: 'editSalesRegion/:id',
    component: EditSalesRegionComponent
  },
  { path: 'indexSalesRegion',
    component: IndexSalesRegionComponent
  }
];
export const SalesCampaignRoutes: Routes = [
  { path: 'createSalesCampaign',
    component: CreateSalesCampaignComponent
  },
  {
    path: 'editSalesCampaign/:id',
    component: EditSalesCampaignComponent
  },
  { path: 'indexSalesCampaign',
    component: IndexSalesCampaignComponent
  }
];
