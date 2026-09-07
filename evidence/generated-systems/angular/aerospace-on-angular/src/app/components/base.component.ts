import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {AerospaceManufacturerService} from '../services/AerospaceManufacturer.service';
import {AircraftProgramService} from '../services/AircraftProgram.service';
import {AircraftFamilyService} from '../services/AircraftFamily.service';
import {AircraftModelService} from '../services/AircraftModel.service';
import {EngineTypeService} from '../services/EngineType.service';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import {AvionicsSuiteService} from '../services/AvionicsSuite.service';
import {APUService} from '../services/APU.service';
import {LandingGearService} from '../services/LandingGear.service';
import {AircraftOptionService} from '../services/AircraftOption.service';
import {AircraftPackageService} from '../services/AircraftPackage.service';
import {SupplierService} from '../services/Supplier.service';
import {Component_Service} from '../services/Component_.service';
import {PlantService} from '../services/Plant.service';
import {ProductionLineService} from '../services/ProductionLine.service';
import {WorkCenterService} from '../services/WorkCenter.service';
import {ProductionOrderService} from '../services/ProductionOrder.service';
import {BuildScheduleService} from '../services/BuildSchedule.service';
import {WarehouseService} from '../services/Warehouse.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import {OperatorService} from '../services/Operator.service';
import {AircraftOrderService} from '../services/AircraftOrder.service';
import {QuoteService} from '../services/Quote.service';
import {PurchaseAgreementService} from '../services/PurchaseAgreement.service';
import {AircraftService} from '../services/Aircraft.service';
import {RegistrationService} from '../services/Registration.service';
import {WarrantyService} from '../services/Warranty.service';
import {CabinLayoutService} from '../services/CabinLayout.service';
import {MROFacilityService} from '../services/MROFacility.service';
import {MaintenanceAppointmentService} from '../services/MaintenanceAppointment.service';
import {MaintenanceWorkOrderService} from '../services/MaintenanceWorkOrder.service';
import {AirworthinessDirectiveService} from '../services/AirworthinessDirective.service';
import {ServiceBulletinService} from '../services/ServiceBulletin.service';
import {ConnectedAircraftService} from '../services/ConnectedAircraft.service';
import {FlightHealthEventService} from '../services/FlightHealthEvent.service';
import {SoftwareLoadService} from '../services/SoftwareLoad.service';
import {TypeCertificateService} from '../services/TypeCertificate.service';
import {ProductionCertificateService} from '../services/ProductionCertificate.service';
import {SalesRegionService} from '../services/SalesRegion.service';
import {SalesCampaignService} from '../services/SalesCampaign.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    AircraftTypes = Object.keys(enumTypes.AircraftType);
    EngineCategorys = Object.keys(enumTypes.EngineCategory);
    LandingGearTypes = Object.keys(enumTypes.LandingGearType);
    OptionCategorys = Object.keys(enumTypes.OptionCategory);
    PackageTypes = Object.keys(enumTypes.PackageType);
    SupplierTypes = Object.keys(enumTypes.SupplierType);
    SupplierApprovalStatuss = Object.keys(enumTypes.SupplierApprovalStatus);
    ComponentCategorys = Object.keys(enumTypes.ComponentCategory);
    SerializationMethods = Object.keys(enumTypes.SerializationMethod);
    ProductionLineTypes = Object.keys(enumTypes.ProductionLineType);
    ProductionOrderStatuss = Object.keys(enumTypes.ProductionOrderStatus);
    ScheduleStatuss = Object.keys(enumTypes.ScheduleStatus);
    OperatorTypes = Object.keys(enumTypes.OperatorType);
    AircraftOrderStatuss = Object.keys(enumTypes.AircraftOrderStatus);
    WarrantyTypes = Object.keys(enumTypes.WarrantyType);
    AppointmentStatuss = Object.keys(enumTypes.AppointmentStatus);
    WorkOrderStatuss = Object.keys(enumTypes.WorkOrderStatus);
    ServiceBulletinCategorys = Object.keys(enumTypes.ServiceBulletinCategory);
    ConnectivityStatuss = Object.keys(enumTypes.ConnectivityStatus);
    EventSeveritys = Object.keys(enumTypes.EventSeverity);
    SoftwareLoadTypes = Object.keys(enumTypes.SoftwareLoadType);
    SalesCampaignStatuss = Object.keys(enumTypes.SalesCampaignStatus);
    ProgramStatuss = Object.keys(enumTypes.ProgramStatus);

// all collection instances
    aerospaceManufacturers : any;
    aircraftPrograms : any;
    aircraftFamilys : any;
    aircraftModels : any;
    engineTypes : any;
    aircraftVariants : any;
    avionicsSuites : any;
    aPUs : any;
    landingGears : any;
    aircraftOptions : any;
    aircraftPackages : any;
    suppliers : any;
    component_s : any;
    plants : any;
    productionLines : any;
    workCenters : any;
    productionOrders : any;
    buildSchedules : any;
    warehouses : any;
    inventoryItems : any;
    operators : any;
    aircraftOrders : any;
    quotes : any;
    purchaseAgreements : any;
    aircrafts : any;
    registrations : any;
    warrantys : any;
    cabinLayouts : any;
    mROFacilitys : any;
    maintenanceAppointments : any;
    maintenanceWorkOrders : any;
    airworthinessDirectives : any;
    serviceBulletins : any;
    connectedAircrafts : any;
    flightHealthEvents : any;
    softwareLoads : any;
    typeCertificates : any;
    productionCertificates : any;
    salesRegions : any;
    salesCampaigns : any;
  
// initialization  
    ngOnInit() {
    }

    initAerospaceManufacturerList() {
        if ( this.aerospaceManufacturers == null ) {
            new AerospaceManufacturerService(this.http).getAerospaceManufacturers().subscribe(res => {
                this.aerospaceManufacturers = res;
            });
        }
    }
    
    initAircraftProgramList() {
        if ( this.aircraftPrograms == null ) {
            new AircraftProgramService(this.http).getAircraftPrograms().subscribe(res => {
                this.aircraftPrograms = res;
            });
        }
    }
    
    initAircraftFamilyList() {
        if ( this.aircraftFamilys == null ) {
            new AircraftFamilyService(this.http).getAircraftFamilys().subscribe(res => {
                this.aircraftFamilys = res;
            });
        }
    }
    
    initAircraftModelList() {
        if ( this.aircraftModels == null ) {
            new AircraftModelService(this.http).getAircraftModels().subscribe(res => {
                this.aircraftModels = res;
            });
        }
    }
    
    initEngineTypeList() {
        if ( this.engineTypes == null ) {
            new EngineTypeService(this.http).getEngineTypes().subscribe(res => {
                this.engineTypes = res;
            });
        }
    }
    
    initAircraftVariantList() {
        if ( this.aircraftVariants == null ) {
            new AircraftVariantService(this.http).getAircraftVariants().subscribe(res => {
                this.aircraftVariants = res;
            });
        }
    }
    
    initAvionicsSuiteList() {
        if ( this.avionicsSuites == null ) {
            new AvionicsSuiteService(this.http).getAvionicsSuites().subscribe(res => {
                this.avionicsSuites = res;
            });
        }
    }
    
    initAPUList() {
        if ( this.aPUs == null ) {
            new APUService(this.http).getAPUs().subscribe(res => {
                this.aPUs = res;
            });
        }
    }
    
    initLandingGearList() {
        if ( this.landingGears == null ) {
            new LandingGearService(this.http).getLandingGears().subscribe(res => {
                this.landingGears = res;
            });
        }
    }
    
    initAircraftOptionList() {
        if ( this.aircraftOptions == null ) {
            new AircraftOptionService(this.http).getAircraftOptions().subscribe(res => {
                this.aircraftOptions = res;
            });
        }
    }
    
    initAircraftPackageList() {
        if ( this.aircraftPackages == null ) {
            new AircraftPackageService(this.http).getAircraftPackages().subscribe(res => {
                this.aircraftPackages = res;
            });
        }
    }
    
    initSupplierList() {
        if ( this.suppliers == null ) {
            new SupplierService(this.http).getSuppliers().subscribe(res => {
                this.suppliers = res;
            });
        }
    }
    
    initComponent_List() {
        if ( this.component_s == null ) {
            new Component_Service(this.http).getComponent_s().subscribe(res => {
                this.component_s = res;
            });
        }
    }
    
    initPlantList() {
        if ( this.plants == null ) {
            new PlantService(this.http).getPlants().subscribe(res => {
                this.plants = res;
            });
        }
    }
    
    initProductionLineList() {
        if ( this.productionLines == null ) {
            new ProductionLineService(this.http).getProductionLines().subscribe(res => {
                this.productionLines = res;
            });
        }
    }
    
    initWorkCenterList() {
        if ( this.workCenters == null ) {
            new WorkCenterService(this.http).getWorkCenters().subscribe(res => {
                this.workCenters = res;
            });
        }
    }
    
    initProductionOrderList() {
        if ( this.productionOrders == null ) {
            new ProductionOrderService(this.http).getProductionOrders().subscribe(res => {
                this.productionOrders = res;
            });
        }
    }
    
    initBuildScheduleList() {
        if ( this.buildSchedules == null ) {
            new BuildScheduleService(this.http).getBuildSchedules().subscribe(res => {
                this.buildSchedules = res;
            });
        }
    }
    
    initWarehouseList() {
        if ( this.warehouses == null ) {
            new WarehouseService(this.http).getWarehouses().subscribe(res => {
                this.warehouses = res;
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
    
    initOperatorList() {
        if ( this.operators == null ) {
            new OperatorService(this.http).getOperators().subscribe(res => {
                this.operators = res;
            });
        }
    }
    
    initAircraftOrderList() {
        if ( this.aircraftOrders == null ) {
            new AircraftOrderService(this.http).getAircraftOrders().subscribe(res => {
                this.aircraftOrders = res;
            });
        }
    }
    
    initQuoteList() {
        if ( this.quotes == null ) {
            new QuoteService(this.http).getQuotes().subscribe(res => {
                this.quotes = res;
            });
        }
    }
    
    initPurchaseAgreementList() {
        if ( this.purchaseAgreements == null ) {
            new PurchaseAgreementService(this.http).getPurchaseAgreements().subscribe(res => {
                this.purchaseAgreements = res;
            });
        }
    }
    
    initAircraftList() {
        if ( this.aircrafts == null ) {
            new AircraftService(this.http).getAircrafts().subscribe(res => {
                this.aircrafts = res;
            });
        }
    }
    
    initRegistrationList() {
        if ( this.registrations == null ) {
            new RegistrationService(this.http).getRegistrations().subscribe(res => {
                this.registrations = res;
            });
        }
    }
    
    initWarrantyList() {
        if ( this.warrantys == null ) {
            new WarrantyService(this.http).getWarrantys().subscribe(res => {
                this.warrantys = res;
            });
        }
    }
    
    initCabinLayoutList() {
        if ( this.cabinLayouts == null ) {
            new CabinLayoutService(this.http).getCabinLayouts().subscribe(res => {
                this.cabinLayouts = res;
            });
        }
    }
    
    initMROFacilityList() {
        if ( this.mROFacilitys == null ) {
            new MROFacilityService(this.http).getMROFacilitys().subscribe(res => {
                this.mROFacilitys = res;
            });
        }
    }
    
    initMaintenanceAppointmentList() {
        if ( this.maintenanceAppointments == null ) {
            new MaintenanceAppointmentService(this.http).getMaintenanceAppointments().subscribe(res => {
                this.maintenanceAppointments = res;
            });
        }
    }
    
    initMaintenanceWorkOrderList() {
        if ( this.maintenanceWorkOrders == null ) {
            new MaintenanceWorkOrderService(this.http).getMaintenanceWorkOrders().subscribe(res => {
                this.maintenanceWorkOrders = res;
            });
        }
    }
    
    initAirworthinessDirectiveList() {
        if ( this.airworthinessDirectives == null ) {
            new AirworthinessDirectiveService(this.http).getAirworthinessDirectives().subscribe(res => {
                this.airworthinessDirectives = res;
            });
        }
    }
    
    initServiceBulletinList() {
        if ( this.serviceBulletins == null ) {
            new ServiceBulletinService(this.http).getServiceBulletins().subscribe(res => {
                this.serviceBulletins = res;
            });
        }
    }
    
    initConnectedAircraftList() {
        if ( this.connectedAircrafts == null ) {
            new ConnectedAircraftService(this.http).getConnectedAircrafts().subscribe(res => {
                this.connectedAircrafts = res;
            });
        }
    }
    
    initFlightHealthEventList() {
        if ( this.flightHealthEvents == null ) {
            new FlightHealthEventService(this.http).getFlightHealthEvents().subscribe(res => {
                this.flightHealthEvents = res;
            });
        }
    }
    
    initSoftwareLoadList() {
        if ( this.softwareLoads == null ) {
            new SoftwareLoadService(this.http).getSoftwareLoads().subscribe(res => {
                this.softwareLoads = res;
            });
        }
    }
    
    initTypeCertificateList() {
        if ( this.typeCertificates == null ) {
            new TypeCertificateService(this.http).getTypeCertificates().subscribe(res => {
                this.typeCertificates = res;
            });
        }
    }
    
    initProductionCertificateList() {
        if ( this.productionCertificates == null ) {
            new ProductionCertificateService(this.http).getProductionCertificates().subscribe(res => {
                this.productionCertificates = res;
            });
        }
    }
    
    initSalesRegionList() {
        if ( this.salesRegions == null ) {
            new SalesRegionService(this.http).getSalesRegions().subscribe(res => {
                this.salesRegions = res;
            });
        }
    }
    
    initSalesCampaignList() {
        if ( this.salesCampaigns == null ) {
            new SalesCampaignService(this.http).getSalesCampaigns().subscribe(res => {
                this.salesCampaigns = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
