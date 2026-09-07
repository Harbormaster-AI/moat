import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {EnterpriseService} from '../services/Enterprise.service';
import {BusinessUnitService} from '../services/BusinessUnit.service';
import {PlantService} from '../services/Plant.service';
import {ProductionLineService} from '../services/ProductionLine.service';
import {WorkCenterService} from '../services/WorkCenter.service';
import {ItemService} from '../services/Item.service';
import {BOMService} from '../services/BOM.service';
import {BOMItemService} from '../services/BOMItem.service';
import {RoutingService} from '../services/Routing.service';
import {OperationService} from '../services/Operation.service';
import {WorkOrderService} from '../services/WorkOrder.service';
import {ProductionScheduleService} from '../services/ProductionSchedule.service';
import {SupplierService} from '../services/Supplier.service';
import {PurchaseOrderService} from '../services/PurchaseOrder.service';
import {PurchaseOrderLineService} from '../services/PurchaseOrderLine.service';
import {GoodsReceiptService} from '../services/GoodsReceipt.service';
import {GoodsReceiptLineService} from '../services/GoodsReceiptLine.service';
import {WarehouseService} from '../services/Warehouse.service';
import {LocationService} from '../services/Location.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import {InventoryTransactionService} from '../services/InventoryTransaction.service';
import {CustomerService} from '../services/Customer.service';
import {SalesOrderService} from '../services/SalesOrder.service';
import {SalesOrderLineService} from '../services/SalesOrderLine.service';
import {QualitySpecificationService} from '../services/QualitySpecification.service';
import {InspectionPlanService} from '../services/InspectionPlan.service';
import {InspectionCharacteristicService} from '../services/InspectionCharacteristic.service';
import {InspectionLotService} from '../services/InspectionLot.service';
import {InspectionResultService} from '../services/InspectionResult.service';
import {NonconformanceService} from '../services/Nonconformance.service';
import {CorrectiveActionService} from '../services/CorrectiveAction.service';
import {AssetService} from '../services/Asset.service';
import {MaintenancePlanService} from '../services/MaintenancePlan.service';
import {MaintenanceOrderService} from '../services/MaintenanceOrder.service';
import {EmployeeService} from '../services/Employee.service';
import {ShiftService} from '../services/Shift.service';
import {ShiftAssignmentService} from '../services/ShiftAssignment.service';
import {ForecastService} from '../services/Forecast.service';
import {ForecastLineService} from '../services/ForecastLine.service';
import {MRPRunService} from '../services/MRPRun.service';
import {PlannedOrderService} from '../services/PlannedOrder.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    BusinessUnitCategorys = Object.keys(enumTypes.BusinessUnitCategory);
    ProductionLineTypes = Object.keys(enumTypes.ProductionLineType);
    WorkCenterTypes = Object.keys(enumTypes.WorkCenterType);
    ItemTypes = Object.keys(enumTypes.ItemType);
    ProcurementTypes = Object.keys(enumTypes.ProcurementType);
    UnitOfMeasures = Object.keys(enumTypes.UnitOfMeasure);
    TimeUnits = Object.keys(enumTypes.TimeUnit);
    ProductLifecycleStatuss = Object.keys(enumTypes.ProductLifecycleStatus);
    BOMStatuss = Object.keys(enumTypes.BOMStatus);
    RoutingTypes = Object.keys(enumTypes.RoutingType);
    RoutingStatuss = Object.keys(enumTypes.RoutingStatus);
    OperationTypes = Object.keys(enumTypes.OperationType);
    WorkOrderStatuss = Object.keys(enumTypes.WorkOrderStatus);
    ScheduleStatuss = Object.keys(enumTypes.ScheduleStatus);
    SupplierTiers = Object.keys(enumTypes.SupplierTier);
    PurchaseOrderStatuss = Object.keys(enumTypes.PurchaseOrderStatus);
    ReceiptStatuss = Object.keys(enumTypes.ReceiptStatus);
    WarehouseTypes = Object.keys(enumTypes.WarehouseType);
    LocationTypes = Object.keys(enumTypes.LocationType);
    InventoryTransactionTypes = Object.keys(enumTypes.InventoryTransactionType);
    CustomerTypes = Object.keys(enumTypes.CustomerType);
    SalesOrderStatuss = Object.keys(enumTypes.SalesOrderStatus);
    SamplingPlanTypes = Object.keys(enumTypes.SamplingPlanType);
    QualityPlanStatuss = Object.keys(enumTypes.QualityPlanStatus);
    MeasurementTypes = Object.keys(enumTypes.MeasurementType);
    InspectionTypes = Object.keys(enumTypes.InspectionType);
    InspectionStatuss = Object.keys(enumTypes.InspectionStatus);
    InspectionResultStatuss = Object.keys(enumTypes.InspectionResultStatus);
    NonconformanceTypes = Object.keys(enumTypes.NonconformanceType);
    QualitySeveritys = Object.keys(enumTypes.QualitySeverity);
    NonconformanceStatuss = Object.keys(enumTypes.NonconformanceStatus);
    CAPAStatuss = Object.keys(enumTypes.CAPAStatus);
    AssetStatuss = Object.keys(enumTypes.AssetStatus);
    MaintenanceStrategys = Object.keys(enumTypes.MaintenanceStrategy);
    MaintenanceOrderStatuss = Object.keys(enumTypes.MaintenanceOrderStatus);
    EmployeeRoles = Object.keys(enumTypes.EmployeeRole);
    SkillLevels = Object.keys(enumTypes.SkillLevel);
    ShiftTypes = Object.keys(enumTypes.ShiftType);
    ForecastMethods = Object.keys(enumTypes.ForecastMethod);
    MRPRunStatuss = Object.keys(enumTypes.MRPRunStatus);
    PlannedOrderTypes = Object.keys(enumTypes.PlannedOrderType);
    PlannedOrderStatuss = Object.keys(enumTypes.PlannedOrderStatus);
    PaymentTermss = Object.keys(enumTypes.PaymentTerms);

// all collection instances
    enterprises : any;
    businessUnits : any;
    plants : any;
    productionLines : any;
    workCenters : any;
    items : any;
    bOMs : any;
    bOMItems : any;
    routings : any;
    operations : any;
    workOrders : any;
    productionSchedules : any;
    suppliers : any;
    purchaseOrders : any;
    purchaseOrderLines : any;
    goodsReceipts : any;
    goodsReceiptLines : any;
    warehouses : any;
    locations : any;
    inventoryItems : any;
    inventoryTransactions : any;
    customers : any;
    salesOrders : any;
    salesOrderLines : any;
    qualitySpecifications : any;
    inspectionPlans : any;
    inspectionCharacteristics : any;
    inspectionLots : any;
    inspectionResults : any;
    nonconformances : any;
    correctiveActions : any;
    assets : any;
    maintenancePlans : any;
    maintenanceOrders : any;
    employees : any;
    shifts : any;
    shiftAssignments : any;
    forecasts : any;
    forecastLines : any;
    mRPRuns : any;
    plannedOrders : any;
  
// initialization  
    ngOnInit() {
    }

    initEnterpriseList() {
        if ( this.enterprises == null ) {
            new EnterpriseService(this.http).getEnterprises().subscribe(res => {
                this.enterprises = res;
            });
        }
    }
    
    initBusinessUnitList() {
        if ( this.businessUnits == null ) {
            new BusinessUnitService(this.http).getBusinessUnits().subscribe(res => {
                this.businessUnits = res;
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
    
    initItemList() {
        if ( this.items == null ) {
            new ItemService(this.http).getItems().subscribe(res => {
                this.items = res;
            });
        }
    }
    
    initBOMList() {
        if ( this.bOMs == null ) {
            new BOMService(this.http).getBOMs().subscribe(res => {
                this.bOMs = res;
            });
        }
    }
    
    initBOMItemList() {
        if ( this.bOMItems == null ) {
            new BOMItemService(this.http).getBOMItems().subscribe(res => {
                this.bOMItems = res;
            });
        }
    }
    
    initRoutingList() {
        if ( this.routings == null ) {
            new RoutingService(this.http).getRoutings().subscribe(res => {
                this.routings = res;
            });
        }
    }
    
    initOperationList() {
        if ( this.operations == null ) {
            new OperationService(this.http).getOperations().subscribe(res => {
                this.operations = res;
            });
        }
    }
    
    initWorkOrderList() {
        if ( this.workOrders == null ) {
            new WorkOrderService(this.http).getWorkOrders().subscribe(res => {
                this.workOrders = res;
            });
        }
    }
    
    initProductionScheduleList() {
        if ( this.productionSchedules == null ) {
            new ProductionScheduleService(this.http).getProductionSchedules().subscribe(res => {
                this.productionSchedules = res;
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
    
    initPurchaseOrderList() {
        if ( this.purchaseOrders == null ) {
            new PurchaseOrderService(this.http).getPurchaseOrders().subscribe(res => {
                this.purchaseOrders = res;
            });
        }
    }
    
    initPurchaseOrderLineList() {
        if ( this.purchaseOrderLines == null ) {
            new PurchaseOrderLineService(this.http).getPurchaseOrderLines().subscribe(res => {
                this.purchaseOrderLines = res;
            });
        }
    }
    
    initGoodsReceiptList() {
        if ( this.goodsReceipts == null ) {
            new GoodsReceiptService(this.http).getGoodsReceipts().subscribe(res => {
                this.goodsReceipts = res;
            });
        }
    }
    
    initGoodsReceiptLineList() {
        if ( this.goodsReceiptLines == null ) {
            new GoodsReceiptLineService(this.http).getGoodsReceiptLines().subscribe(res => {
                this.goodsReceiptLines = res;
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
    
    initLocationList() {
        if ( this.locations == null ) {
            new LocationService(this.http).getLocations().subscribe(res => {
                this.locations = res;
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
    
    initInventoryTransactionList() {
        if ( this.inventoryTransactions == null ) {
            new InventoryTransactionService(this.http).getInventoryTransactions().subscribe(res => {
                this.inventoryTransactions = res;
            });
        }
    }
    
    initCustomerList() {
        if ( this.customers == null ) {
            new CustomerService(this.http).getCustomers().subscribe(res => {
                this.customers = res;
            });
        }
    }
    
    initSalesOrderList() {
        if ( this.salesOrders == null ) {
            new SalesOrderService(this.http).getSalesOrders().subscribe(res => {
                this.salesOrders = res;
            });
        }
    }
    
    initSalesOrderLineList() {
        if ( this.salesOrderLines == null ) {
            new SalesOrderLineService(this.http).getSalesOrderLines().subscribe(res => {
                this.salesOrderLines = res;
            });
        }
    }
    
    initQualitySpecificationList() {
        if ( this.qualitySpecifications == null ) {
            new QualitySpecificationService(this.http).getQualitySpecifications().subscribe(res => {
                this.qualitySpecifications = res;
            });
        }
    }
    
    initInspectionPlanList() {
        if ( this.inspectionPlans == null ) {
            new InspectionPlanService(this.http).getInspectionPlans().subscribe(res => {
                this.inspectionPlans = res;
            });
        }
    }
    
    initInspectionCharacteristicList() {
        if ( this.inspectionCharacteristics == null ) {
            new InspectionCharacteristicService(this.http).getInspectionCharacteristics().subscribe(res => {
                this.inspectionCharacteristics = res;
            });
        }
    }
    
    initInspectionLotList() {
        if ( this.inspectionLots == null ) {
            new InspectionLotService(this.http).getInspectionLots().subscribe(res => {
                this.inspectionLots = res;
            });
        }
    }
    
    initInspectionResultList() {
        if ( this.inspectionResults == null ) {
            new InspectionResultService(this.http).getInspectionResults().subscribe(res => {
                this.inspectionResults = res;
            });
        }
    }
    
    initNonconformanceList() {
        if ( this.nonconformances == null ) {
            new NonconformanceService(this.http).getNonconformances().subscribe(res => {
                this.nonconformances = res;
            });
        }
    }
    
    initCorrectiveActionList() {
        if ( this.correctiveActions == null ) {
            new CorrectiveActionService(this.http).getCorrectiveActions().subscribe(res => {
                this.correctiveActions = res;
            });
        }
    }
    
    initAssetList() {
        if ( this.assets == null ) {
            new AssetService(this.http).getAssets().subscribe(res => {
                this.assets = res;
            });
        }
    }
    
    initMaintenancePlanList() {
        if ( this.maintenancePlans == null ) {
            new MaintenancePlanService(this.http).getMaintenancePlans().subscribe(res => {
                this.maintenancePlans = res;
            });
        }
    }
    
    initMaintenanceOrderList() {
        if ( this.maintenanceOrders == null ) {
            new MaintenanceOrderService(this.http).getMaintenanceOrders().subscribe(res => {
                this.maintenanceOrders = res;
            });
        }
    }
    
    initEmployeeList() {
        if ( this.employees == null ) {
            new EmployeeService(this.http).getEmployees().subscribe(res => {
                this.employees = res;
            });
        }
    }
    
    initShiftList() {
        if ( this.shifts == null ) {
            new ShiftService(this.http).getShifts().subscribe(res => {
                this.shifts = res;
            });
        }
    }
    
    initShiftAssignmentList() {
        if ( this.shiftAssignments == null ) {
            new ShiftAssignmentService(this.http).getShiftAssignments().subscribe(res => {
                this.shiftAssignments = res;
            });
        }
    }
    
    initForecastList() {
        if ( this.forecasts == null ) {
            new ForecastService(this.http).getForecasts().subscribe(res => {
                this.forecasts = res;
            });
        }
    }
    
    initForecastLineList() {
        if ( this.forecastLines == null ) {
            new ForecastLineService(this.http).getForecastLines().subscribe(res => {
                this.forecastLines = res;
            });
        }
    }
    
    initMRPRunList() {
        if ( this.mRPRuns == null ) {
            new MRPRunService(this.http).getMRPRuns().subscribe(res => {
                this.mRPRuns = res;
            });
        }
    }
    
    initPlannedOrderList() {
        if ( this.plannedOrders == null ) {
            new PlannedOrderService(this.http).getPlannedOrders().subscribe(res => {
                this.plannedOrders = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
