// routerConfig.ts

import { Routes } from '@angular/router';
import { CreateEnterpriseComponent } from './components/Enterprise/create/create.component';
import { EditEnterpriseComponent } from './components/Enterprise/edit/edit.component';
import { IndexEnterpriseComponent } from './components/Enterprise/index/index.component';
import { CreateBusinessUnitComponent } from './components/BusinessUnit/create/create.component';
import { EditBusinessUnitComponent } from './components/BusinessUnit/edit/edit.component';
import { IndexBusinessUnitComponent } from './components/BusinessUnit/index/index.component';
import { CreatePlantComponent } from './components/Plant/create/create.component';
import { EditPlantComponent } from './components/Plant/edit/edit.component';
import { IndexPlantComponent } from './components/Plant/index/index.component';
import { CreateProductionLineComponent } from './components/ProductionLine/create/create.component';
import { EditProductionLineComponent } from './components/ProductionLine/edit/edit.component';
import { IndexProductionLineComponent } from './components/ProductionLine/index/index.component';
import { CreateWorkCenterComponent } from './components/WorkCenter/create/create.component';
import { EditWorkCenterComponent } from './components/WorkCenter/edit/edit.component';
import { IndexWorkCenterComponent } from './components/WorkCenter/index/index.component';
import { CreateItemComponent } from './components/Item/create/create.component';
import { EditItemComponent } from './components/Item/edit/edit.component';
import { IndexItemComponent } from './components/Item/index/index.component';
import { CreateBOMComponent } from './components/BOM/create/create.component';
import { EditBOMComponent } from './components/BOM/edit/edit.component';
import { IndexBOMComponent } from './components/BOM/index/index.component';
import { CreateBOMItemComponent } from './components/BOMItem/create/create.component';
import { EditBOMItemComponent } from './components/BOMItem/edit/edit.component';
import { IndexBOMItemComponent } from './components/BOMItem/index/index.component';
import { CreateRoutingComponent } from './components/Routing/create/create.component';
import { EditRoutingComponent } from './components/Routing/edit/edit.component';
import { IndexRoutingComponent } from './components/Routing/index/index.component';
import { CreateOperationComponent } from './components/Operation/create/create.component';
import { EditOperationComponent } from './components/Operation/edit/edit.component';
import { IndexOperationComponent } from './components/Operation/index/index.component';
import { CreateWorkOrderComponent } from './components/WorkOrder/create/create.component';
import { EditWorkOrderComponent } from './components/WorkOrder/edit/edit.component';
import { IndexWorkOrderComponent } from './components/WorkOrder/index/index.component';
import { CreateProductionScheduleComponent } from './components/ProductionSchedule/create/create.component';
import { EditProductionScheduleComponent } from './components/ProductionSchedule/edit/edit.component';
import { IndexProductionScheduleComponent } from './components/ProductionSchedule/index/index.component';
import { CreateSupplierComponent } from './components/Supplier/create/create.component';
import { EditSupplierComponent } from './components/Supplier/edit/edit.component';
import { IndexSupplierComponent } from './components/Supplier/index/index.component';
import { CreatePurchaseOrderComponent } from './components/PurchaseOrder/create/create.component';
import { EditPurchaseOrderComponent } from './components/PurchaseOrder/edit/edit.component';
import { IndexPurchaseOrderComponent } from './components/PurchaseOrder/index/index.component';
import { CreatePurchaseOrderLineComponent } from './components/PurchaseOrderLine/create/create.component';
import { EditPurchaseOrderLineComponent } from './components/PurchaseOrderLine/edit/edit.component';
import { IndexPurchaseOrderLineComponent } from './components/PurchaseOrderLine/index/index.component';
import { CreateGoodsReceiptComponent } from './components/GoodsReceipt/create/create.component';
import { EditGoodsReceiptComponent } from './components/GoodsReceipt/edit/edit.component';
import { IndexGoodsReceiptComponent } from './components/GoodsReceipt/index/index.component';
import { CreateGoodsReceiptLineComponent } from './components/GoodsReceiptLine/create/create.component';
import { EditGoodsReceiptLineComponent } from './components/GoodsReceiptLine/edit/edit.component';
import { IndexGoodsReceiptLineComponent } from './components/GoodsReceiptLine/index/index.component';
import { CreateWarehouseComponent } from './components/Warehouse/create/create.component';
import { EditWarehouseComponent } from './components/Warehouse/edit/edit.component';
import { IndexWarehouseComponent } from './components/Warehouse/index/index.component';
import { CreateLocationComponent } from './components/Location/create/create.component';
import { EditLocationComponent } from './components/Location/edit/edit.component';
import { IndexLocationComponent } from './components/Location/index/index.component';
import { CreateInventoryItemComponent } from './components/InventoryItem/create/create.component';
import { EditInventoryItemComponent } from './components/InventoryItem/edit/edit.component';
import { IndexInventoryItemComponent } from './components/InventoryItem/index/index.component';
import { CreateInventoryTransactionComponent } from './components/InventoryTransaction/create/create.component';
import { EditInventoryTransactionComponent } from './components/InventoryTransaction/edit/edit.component';
import { IndexInventoryTransactionComponent } from './components/InventoryTransaction/index/index.component';
import { CreateCustomerComponent } from './components/Customer/create/create.component';
import { EditCustomerComponent } from './components/Customer/edit/edit.component';
import { IndexCustomerComponent } from './components/Customer/index/index.component';
import { CreateSalesOrderComponent } from './components/SalesOrder/create/create.component';
import { EditSalesOrderComponent } from './components/SalesOrder/edit/edit.component';
import { IndexSalesOrderComponent } from './components/SalesOrder/index/index.component';
import { CreateSalesOrderLineComponent } from './components/SalesOrderLine/create/create.component';
import { EditSalesOrderLineComponent } from './components/SalesOrderLine/edit/edit.component';
import { IndexSalesOrderLineComponent } from './components/SalesOrderLine/index/index.component';
import { CreateQualitySpecificationComponent } from './components/QualitySpecification/create/create.component';
import { EditQualitySpecificationComponent } from './components/QualitySpecification/edit/edit.component';
import { IndexQualitySpecificationComponent } from './components/QualitySpecification/index/index.component';
import { CreateInspectionPlanComponent } from './components/InspectionPlan/create/create.component';
import { EditInspectionPlanComponent } from './components/InspectionPlan/edit/edit.component';
import { IndexInspectionPlanComponent } from './components/InspectionPlan/index/index.component';
import { CreateInspectionCharacteristicComponent } from './components/InspectionCharacteristic/create/create.component';
import { EditInspectionCharacteristicComponent } from './components/InspectionCharacteristic/edit/edit.component';
import { IndexInspectionCharacteristicComponent } from './components/InspectionCharacteristic/index/index.component';
import { CreateInspectionLotComponent } from './components/InspectionLot/create/create.component';
import { EditInspectionLotComponent } from './components/InspectionLot/edit/edit.component';
import { IndexInspectionLotComponent } from './components/InspectionLot/index/index.component';
import { CreateInspectionResultComponent } from './components/InspectionResult/create/create.component';
import { EditInspectionResultComponent } from './components/InspectionResult/edit/edit.component';
import { IndexInspectionResultComponent } from './components/InspectionResult/index/index.component';
import { CreateNonconformanceComponent } from './components/Nonconformance/create/create.component';
import { EditNonconformanceComponent } from './components/Nonconformance/edit/edit.component';
import { IndexNonconformanceComponent } from './components/Nonconformance/index/index.component';
import { CreateCorrectiveActionComponent } from './components/CorrectiveAction/create/create.component';
import { EditCorrectiveActionComponent } from './components/CorrectiveAction/edit/edit.component';
import { IndexCorrectiveActionComponent } from './components/CorrectiveAction/index/index.component';
import { CreateAssetComponent } from './components/Asset/create/create.component';
import { EditAssetComponent } from './components/Asset/edit/edit.component';
import { IndexAssetComponent } from './components/Asset/index/index.component';
import { CreateMaintenancePlanComponent } from './components/MaintenancePlan/create/create.component';
import { EditMaintenancePlanComponent } from './components/MaintenancePlan/edit/edit.component';
import { IndexMaintenancePlanComponent } from './components/MaintenancePlan/index/index.component';
import { CreateMaintenanceOrderComponent } from './components/MaintenanceOrder/create/create.component';
import { EditMaintenanceOrderComponent } from './components/MaintenanceOrder/edit/edit.component';
import { IndexMaintenanceOrderComponent } from './components/MaintenanceOrder/index/index.component';
import { CreateEmployeeComponent } from './components/Employee/create/create.component';
import { EditEmployeeComponent } from './components/Employee/edit/edit.component';
import { IndexEmployeeComponent } from './components/Employee/index/index.component';
import { CreateShiftComponent } from './components/Shift/create/create.component';
import { EditShiftComponent } from './components/Shift/edit/edit.component';
import { IndexShiftComponent } from './components/Shift/index/index.component';
import { CreateShiftAssignmentComponent } from './components/ShiftAssignment/create/create.component';
import { EditShiftAssignmentComponent } from './components/ShiftAssignment/edit/edit.component';
import { IndexShiftAssignmentComponent } from './components/ShiftAssignment/index/index.component';
import { CreateForecastComponent } from './components/Forecast/create/create.component';
import { EditForecastComponent } from './components/Forecast/edit/edit.component';
import { IndexForecastComponent } from './components/Forecast/index/index.component';
import { CreateForecastLineComponent } from './components/ForecastLine/create/create.component';
import { EditForecastLineComponent } from './components/ForecastLine/edit/edit.component';
import { IndexForecastLineComponent } from './components/ForecastLine/index/index.component';
import { CreateMRPRunComponent } from './components/MRPRun/create/create.component';
import { EditMRPRunComponent } from './components/MRPRun/edit/edit.component';
import { IndexMRPRunComponent } from './components/MRPRun/index/index.component';
import { CreatePlannedOrderComponent } from './components/PlannedOrder/create/create.component';
import { EditPlannedOrderComponent } from './components/PlannedOrder/edit/edit.component';
import { IndexPlannedOrderComponent } from './components/PlannedOrder/index/index.component';

export const EnterpriseRoutes: Routes = [
  { path: 'createEnterprise',
    component: CreateEnterpriseComponent
  },
  {
    path: 'editEnterprise/:id',
    component: EditEnterpriseComponent
  },
  { path: 'indexEnterprise',
    component: IndexEnterpriseComponent
  }
];
export const BusinessUnitRoutes: Routes = [
  { path: 'createBusinessUnit',
    component: CreateBusinessUnitComponent
  },
  {
    path: 'editBusinessUnit/:id',
    component: EditBusinessUnitComponent
  },
  { path: 'indexBusinessUnit',
    component: IndexBusinessUnitComponent
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
export const ItemRoutes: Routes = [
  { path: 'createItem',
    component: CreateItemComponent
  },
  {
    path: 'editItem/:id',
    component: EditItemComponent
  },
  { path: 'indexItem',
    component: IndexItemComponent
  }
];
export const BOMRoutes: Routes = [
  { path: 'createBOM',
    component: CreateBOMComponent
  },
  {
    path: 'editBOM/:id',
    component: EditBOMComponent
  },
  { path: 'indexBOM',
    component: IndexBOMComponent
  }
];
export const BOMItemRoutes: Routes = [
  { path: 'createBOMItem',
    component: CreateBOMItemComponent
  },
  {
    path: 'editBOMItem/:id',
    component: EditBOMItemComponent
  },
  { path: 'indexBOMItem',
    component: IndexBOMItemComponent
  }
];
export const RoutingRoutes: Routes = [
  { path: 'createRouting',
    component: CreateRoutingComponent
  },
  {
    path: 'editRouting/:id',
    component: EditRoutingComponent
  },
  { path: 'indexRouting',
    component: IndexRoutingComponent
  }
];
export const OperationRoutes: Routes = [
  { path: 'createOperation',
    component: CreateOperationComponent
  },
  {
    path: 'editOperation/:id',
    component: EditOperationComponent
  },
  { path: 'indexOperation',
    component: IndexOperationComponent
  }
];
export const WorkOrderRoutes: Routes = [
  { path: 'createWorkOrder',
    component: CreateWorkOrderComponent
  },
  {
    path: 'editWorkOrder/:id',
    component: EditWorkOrderComponent
  },
  { path: 'indexWorkOrder',
    component: IndexWorkOrderComponent
  }
];
export const ProductionScheduleRoutes: Routes = [
  { path: 'createProductionSchedule',
    component: CreateProductionScheduleComponent
  },
  {
    path: 'editProductionSchedule/:id',
    component: EditProductionScheduleComponent
  },
  { path: 'indexProductionSchedule',
    component: IndexProductionScheduleComponent
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
export const PurchaseOrderRoutes: Routes = [
  { path: 'createPurchaseOrder',
    component: CreatePurchaseOrderComponent
  },
  {
    path: 'editPurchaseOrder/:id',
    component: EditPurchaseOrderComponent
  },
  { path: 'indexPurchaseOrder',
    component: IndexPurchaseOrderComponent
  }
];
export const PurchaseOrderLineRoutes: Routes = [
  { path: 'createPurchaseOrderLine',
    component: CreatePurchaseOrderLineComponent
  },
  {
    path: 'editPurchaseOrderLine/:id',
    component: EditPurchaseOrderLineComponent
  },
  { path: 'indexPurchaseOrderLine',
    component: IndexPurchaseOrderLineComponent
  }
];
export const GoodsReceiptRoutes: Routes = [
  { path: 'createGoodsReceipt',
    component: CreateGoodsReceiptComponent
  },
  {
    path: 'editGoodsReceipt/:id',
    component: EditGoodsReceiptComponent
  },
  { path: 'indexGoodsReceipt',
    component: IndexGoodsReceiptComponent
  }
];
export const GoodsReceiptLineRoutes: Routes = [
  { path: 'createGoodsReceiptLine',
    component: CreateGoodsReceiptLineComponent
  },
  {
    path: 'editGoodsReceiptLine/:id',
    component: EditGoodsReceiptLineComponent
  },
  { path: 'indexGoodsReceiptLine',
    component: IndexGoodsReceiptLineComponent
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
export const LocationRoutes: Routes = [
  { path: 'createLocation',
    component: CreateLocationComponent
  },
  {
    path: 'editLocation/:id',
    component: EditLocationComponent
  },
  { path: 'indexLocation',
    component: IndexLocationComponent
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
export const InventoryTransactionRoutes: Routes = [
  { path: 'createInventoryTransaction',
    component: CreateInventoryTransactionComponent
  },
  {
    path: 'editInventoryTransaction/:id',
    component: EditInventoryTransactionComponent
  },
  { path: 'indexInventoryTransaction',
    component: IndexInventoryTransactionComponent
  }
];
export const CustomerRoutes: Routes = [
  { path: 'createCustomer',
    component: CreateCustomerComponent
  },
  {
    path: 'editCustomer/:id',
    component: EditCustomerComponent
  },
  { path: 'indexCustomer',
    component: IndexCustomerComponent
  }
];
export const SalesOrderRoutes: Routes = [
  { path: 'createSalesOrder',
    component: CreateSalesOrderComponent
  },
  {
    path: 'editSalesOrder/:id',
    component: EditSalesOrderComponent
  },
  { path: 'indexSalesOrder',
    component: IndexSalesOrderComponent
  }
];
export const SalesOrderLineRoutes: Routes = [
  { path: 'createSalesOrderLine',
    component: CreateSalesOrderLineComponent
  },
  {
    path: 'editSalesOrderLine/:id',
    component: EditSalesOrderLineComponent
  },
  { path: 'indexSalesOrderLine',
    component: IndexSalesOrderLineComponent
  }
];
export const QualitySpecificationRoutes: Routes = [
  { path: 'createQualitySpecification',
    component: CreateQualitySpecificationComponent
  },
  {
    path: 'editQualitySpecification/:id',
    component: EditQualitySpecificationComponent
  },
  { path: 'indexQualitySpecification',
    component: IndexQualitySpecificationComponent
  }
];
export const InspectionPlanRoutes: Routes = [
  { path: 'createInspectionPlan',
    component: CreateInspectionPlanComponent
  },
  {
    path: 'editInspectionPlan/:id',
    component: EditInspectionPlanComponent
  },
  { path: 'indexInspectionPlan',
    component: IndexInspectionPlanComponent
  }
];
export const InspectionCharacteristicRoutes: Routes = [
  { path: 'createInspectionCharacteristic',
    component: CreateInspectionCharacteristicComponent
  },
  {
    path: 'editInspectionCharacteristic/:id',
    component: EditInspectionCharacteristicComponent
  },
  { path: 'indexInspectionCharacteristic',
    component: IndexInspectionCharacteristicComponent
  }
];
export const InspectionLotRoutes: Routes = [
  { path: 'createInspectionLot',
    component: CreateInspectionLotComponent
  },
  {
    path: 'editInspectionLot/:id',
    component: EditInspectionLotComponent
  },
  { path: 'indexInspectionLot',
    component: IndexInspectionLotComponent
  }
];
export const InspectionResultRoutes: Routes = [
  { path: 'createInspectionResult',
    component: CreateInspectionResultComponent
  },
  {
    path: 'editInspectionResult/:id',
    component: EditInspectionResultComponent
  },
  { path: 'indexInspectionResult',
    component: IndexInspectionResultComponent
  }
];
export const NonconformanceRoutes: Routes = [
  { path: 'createNonconformance',
    component: CreateNonconformanceComponent
  },
  {
    path: 'editNonconformance/:id',
    component: EditNonconformanceComponent
  },
  { path: 'indexNonconformance',
    component: IndexNonconformanceComponent
  }
];
export const CorrectiveActionRoutes: Routes = [
  { path: 'createCorrectiveAction',
    component: CreateCorrectiveActionComponent
  },
  {
    path: 'editCorrectiveAction/:id',
    component: EditCorrectiveActionComponent
  },
  { path: 'indexCorrectiveAction',
    component: IndexCorrectiveActionComponent
  }
];
export const AssetRoutes: Routes = [
  { path: 'createAsset',
    component: CreateAssetComponent
  },
  {
    path: 'editAsset/:id',
    component: EditAssetComponent
  },
  { path: 'indexAsset',
    component: IndexAssetComponent
  }
];
export const MaintenancePlanRoutes: Routes = [
  { path: 'createMaintenancePlan',
    component: CreateMaintenancePlanComponent
  },
  {
    path: 'editMaintenancePlan/:id',
    component: EditMaintenancePlanComponent
  },
  { path: 'indexMaintenancePlan',
    component: IndexMaintenancePlanComponent
  }
];
export const MaintenanceOrderRoutes: Routes = [
  { path: 'createMaintenanceOrder',
    component: CreateMaintenanceOrderComponent
  },
  {
    path: 'editMaintenanceOrder/:id',
    component: EditMaintenanceOrderComponent
  },
  { path: 'indexMaintenanceOrder',
    component: IndexMaintenanceOrderComponent
  }
];
export const EmployeeRoutes: Routes = [
  { path: 'createEmployee',
    component: CreateEmployeeComponent
  },
  {
    path: 'editEmployee/:id',
    component: EditEmployeeComponent
  },
  { path: 'indexEmployee',
    component: IndexEmployeeComponent
  }
];
export const ShiftRoutes: Routes = [
  { path: 'createShift',
    component: CreateShiftComponent
  },
  {
    path: 'editShift/:id',
    component: EditShiftComponent
  },
  { path: 'indexShift',
    component: IndexShiftComponent
  }
];
export const ShiftAssignmentRoutes: Routes = [
  { path: 'createShiftAssignment',
    component: CreateShiftAssignmentComponent
  },
  {
    path: 'editShiftAssignment/:id',
    component: EditShiftAssignmentComponent
  },
  { path: 'indexShiftAssignment',
    component: IndexShiftAssignmentComponent
  }
];
export const ForecastRoutes: Routes = [
  { path: 'createForecast',
    component: CreateForecastComponent
  },
  {
    path: 'editForecast/:id',
    component: EditForecastComponent
  },
  { path: 'indexForecast',
    component: IndexForecastComponent
  }
];
export const ForecastLineRoutes: Routes = [
  { path: 'createForecastLine',
    component: CreateForecastLineComponent
  },
  {
    path: 'editForecastLine/:id',
    component: EditForecastLineComponent
  },
  { path: 'indexForecastLine',
    component: IndexForecastLineComponent
  }
];
export const MRPRunRoutes: Routes = [
  { path: 'createMRPRun',
    component: CreateMRPRunComponent
  },
  {
    path: 'editMRPRun/:id',
    component: EditMRPRunComponent
  },
  { path: 'indexMRPRun',
    component: IndexMRPRunComponent
  }
];
export const PlannedOrderRoutes: Routes = [
  { path: 'createPlannedOrder',
    component: CreatePlannedOrderComponent
  },
  {
    path: 'editPlannedOrder/:id',
    component: EditPlannedOrderComponent
  },
  { path: 'indexPlannedOrder',
    component: IndexPlannedOrderComponent
  }
];
