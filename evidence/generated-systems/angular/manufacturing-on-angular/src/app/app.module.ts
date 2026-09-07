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

import {IndexEnterpriseComponent} from './components/Enterprise/index/index.component';
import {CreateEnterpriseComponent} from './components/Enterprise/create/create.component';
import {EditEnterpriseComponent} from './components/Enterprise/edit/edit.component';
import {IndexBusinessUnitComponent} from './components/BusinessUnit/index/index.component';
import {CreateBusinessUnitComponent} from './components/BusinessUnit/create/create.component';
import {EditBusinessUnitComponent} from './components/BusinessUnit/edit/edit.component';
import {IndexPlantComponent} from './components/Plant/index/index.component';
import {CreatePlantComponent} from './components/Plant/create/create.component';
import {EditPlantComponent} from './components/Plant/edit/edit.component';
import {IndexProductionLineComponent} from './components/ProductionLine/index/index.component';
import {CreateProductionLineComponent} from './components/ProductionLine/create/create.component';
import {EditProductionLineComponent} from './components/ProductionLine/edit/edit.component';
import {IndexWorkCenterComponent} from './components/WorkCenter/index/index.component';
import {CreateWorkCenterComponent} from './components/WorkCenter/create/create.component';
import {EditWorkCenterComponent} from './components/WorkCenter/edit/edit.component';
import {IndexItemComponent} from './components/Item/index/index.component';
import {CreateItemComponent} from './components/Item/create/create.component';
import {EditItemComponent} from './components/Item/edit/edit.component';
import {IndexBOMComponent} from './components/BOM/index/index.component';
import {CreateBOMComponent} from './components/BOM/create/create.component';
import {EditBOMComponent} from './components/BOM/edit/edit.component';
import {IndexBOMItemComponent} from './components/BOMItem/index/index.component';
import {CreateBOMItemComponent} from './components/BOMItem/create/create.component';
import {EditBOMItemComponent} from './components/BOMItem/edit/edit.component';
import {IndexRoutingComponent} from './components/Routing/index/index.component';
import {CreateRoutingComponent} from './components/Routing/create/create.component';
import {EditRoutingComponent} from './components/Routing/edit/edit.component';
import {IndexOperationComponent} from './components/Operation/index/index.component';
import {CreateOperationComponent} from './components/Operation/create/create.component';
import {EditOperationComponent} from './components/Operation/edit/edit.component';
import {IndexWorkOrderComponent} from './components/WorkOrder/index/index.component';
import {CreateWorkOrderComponent} from './components/WorkOrder/create/create.component';
import {EditWorkOrderComponent} from './components/WorkOrder/edit/edit.component';
import {IndexProductionScheduleComponent} from './components/ProductionSchedule/index/index.component';
import {CreateProductionScheduleComponent} from './components/ProductionSchedule/create/create.component';
import {EditProductionScheduleComponent} from './components/ProductionSchedule/edit/edit.component';
import {IndexSupplierComponent} from './components/Supplier/index/index.component';
import {CreateSupplierComponent} from './components/Supplier/create/create.component';
import {EditSupplierComponent} from './components/Supplier/edit/edit.component';
import {IndexPurchaseOrderComponent} from './components/PurchaseOrder/index/index.component';
import {CreatePurchaseOrderComponent} from './components/PurchaseOrder/create/create.component';
import {EditPurchaseOrderComponent} from './components/PurchaseOrder/edit/edit.component';
import {IndexPurchaseOrderLineComponent} from './components/PurchaseOrderLine/index/index.component';
import {CreatePurchaseOrderLineComponent} from './components/PurchaseOrderLine/create/create.component';
import {EditPurchaseOrderLineComponent} from './components/PurchaseOrderLine/edit/edit.component';
import {IndexGoodsReceiptComponent} from './components/GoodsReceipt/index/index.component';
import {CreateGoodsReceiptComponent} from './components/GoodsReceipt/create/create.component';
import {EditGoodsReceiptComponent} from './components/GoodsReceipt/edit/edit.component';
import {IndexGoodsReceiptLineComponent} from './components/GoodsReceiptLine/index/index.component';
import {CreateGoodsReceiptLineComponent} from './components/GoodsReceiptLine/create/create.component';
import {EditGoodsReceiptLineComponent} from './components/GoodsReceiptLine/edit/edit.component';
import {IndexWarehouseComponent} from './components/Warehouse/index/index.component';
import {CreateWarehouseComponent} from './components/Warehouse/create/create.component';
import {EditWarehouseComponent} from './components/Warehouse/edit/edit.component';
import {IndexLocationComponent} from './components/Location/index/index.component';
import {CreateLocationComponent} from './components/Location/create/create.component';
import {EditLocationComponent} from './components/Location/edit/edit.component';
import {IndexInventoryItemComponent} from './components/InventoryItem/index/index.component';
import {CreateInventoryItemComponent} from './components/InventoryItem/create/create.component';
import {EditInventoryItemComponent} from './components/InventoryItem/edit/edit.component';
import {IndexInventoryTransactionComponent} from './components/InventoryTransaction/index/index.component';
import {CreateInventoryTransactionComponent} from './components/InventoryTransaction/create/create.component';
import {EditInventoryTransactionComponent} from './components/InventoryTransaction/edit/edit.component';
import {IndexCustomerComponent} from './components/Customer/index/index.component';
import {CreateCustomerComponent} from './components/Customer/create/create.component';
import {EditCustomerComponent} from './components/Customer/edit/edit.component';
import {IndexSalesOrderComponent} from './components/SalesOrder/index/index.component';
import {CreateSalesOrderComponent} from './components/SalesOrder/create/create.component';
import {EditSalesOrderComponent} from './components/SalesOrder/edit/edit.component';
import {IndexSalesOrderLineComponent} from './components/SalesOrderLine/index/index.component';
import {CreateSalesOrderLineComponent} from './components/SalesOrderLine/create/create.component';
import {EditSalesOrderLineComponent} from './components/SalesOrderLine/edit/edit.component';
import {IndexQualitySpecificationComponent} from './components/QualitySpecification/index/index.component';
import {CreateQualitySpecificationComponent} from './components/QualitySpecification/create/create.component';
import {EditQualitySpecificationComponent} from './components/QualitySpecification/edit/edit.component';
import {IndexInspectionPlanComponent} from './components/InspectionPlan/index/index.component';
import {CreateInspectionPlanComponent} from './components/InspectionPlan/create/create.component';
import {EditInspectionPlanComponent} from './components/InspectionPlan/edit/edit.component';
import {IndexInspectionCharacteristicComponent} from './components/InspectionCharacteristic/index/index.component';
import {CreateInspectionCharacteristicComponent} from './components/InspectionCharacteristic/create/create.component';
import {EditInspectionCharacteristicComponent} from './components/InspectionCharacteristic/edit/edit.component';
import {IndexInspectionLotComponent} from './components/InspectionLot/index/index.component';
import {CreateInspectionLotComponent} from './components/InspectionLot/create/create.component';
import {EditInspectionLotComponent} from './components/InspectionLot/edit/edit.component';
import {IndexInspectionResultComponent} from './components/InspectionResult/index/index.component';
import {CreateInspectionResultComponent} from './components/InspectionResult/create/create.component';
import {EditInspectionResultComponent} from './components/InspectionResult/edit/edit.component';
import {IndexNonconformanceComponent} from './components/Nonconformance/index/index.component';
import {CreateNonconformanceComponent} from './components/Nonconformance/create/create.component';
import {EditNonconformanceComponent} from './components/Nonconformance/edit/edit.component';
import {IndexCorrectiveActionComponent} from './components/CorrectiveAction/index/index.component';
import {CreateCorrectiveActionComponent} from './components/CorrectiveAction/create/create.component';
import {EditCorrectiveActionComponent} from './components/CorrectiveAction/edit/edit.component';
import {IndexAssetComponent} from './components/Asset/index/index.component';
import {CreateAssetComponent} from './components/Asset/create/create.component';
import {EditAssetComponent} from './components/Asset/edit/edit.component';
import {IndexMaintenancePlanComponent} from './components/MaintenancePlan/index/index.component';
import {CreateMaintenancePlanComponent} from './components/MaintenancePlan/create/create.component';
import {EditMaintenancePlanComponent} from './components/MaintenancePlan/edit/edit.component';
import {IndexMaintenanceOrderComponent} from './components/MaintenanceOrder/index/index.component';
import {CreateMaintenanceOrderComponent} from './components/MaintenanceOrder/create/create.component';
import {EditMaintenanceOrderComponent} from './components/MaintenanceOrder/edit/edit.component';
import {IndexEmployeeComponent} from './components/Employee/index/index.component';
import {CreateEmployeeComponent} from './components/Employee/create/create.component';
import {EditEmployeeComponent} from './components/Employee/edit/edit.component';
import {IndexShiftComponent} from './components/Shift/index/index.component';
import {CreateShiftComponent} from './components/Shift/create/create.component';
import {EditShiftComponent} from './components/Shift/edit/edit.component';
import {IndexShiftAssignmentComponent} from './components/ShiftAssignment/index/index.component';
import {CreateShiftAssignmentComponent} from './components/ShiftAssignment/create/create.component';
import {EditShiftAssignmentComponent} from './components/ShiftAssignment/edit/edit.component';
import {IndexForecastComponent} from './components/Forecast/index/index.component';
import {CreateForecastComponent} from './components/Forecast/create/create.component';
import {EditForecastComponent} from './components/Forecast/edit/edit.component';
import {IndexForecastLineComponent} from './components/ForecastLine/index/index.component';
import {CreateForecastLineComponent} from './components/ForecastLine/create/create.component';
import {EditForecastLineComponent} from './components/ForecastLine/edit/edit.component';
import {IndexMRPRunComponent} from './components/MRPRun/index/index.component';
import {CreateMRPRunComponent} from './components/MRPRun/create/create.component';
import {EditMRPRunComponent} from './components/MRPRun/edit/edit.component';
import {IndexPlannedOrderComponent} from './components/PlannedOrder/index/index.component';
import {CreatePlannedOrderComponent} from './components/PlannedOrder/create/create.component';
import {EditPlannedOrderComponent} from './components/PlannedOrder/edit/edit.component';

import * as appRoutes from './routerConfig';

import {EnterpriseService} from './services/Enterprise.service';
import {BusinessUnitService} from './services/BusinessUnit.service';
import {PlantService} from './services/Plant.service';
import {ProductionLineService} from './services/ProductionLine.service';
import {WorkCenterService} from './services/WorkCenter.service';
import {ItemService} from './services/Item.service';
import {BOMService} from './services/BOM.service';
import {BOMItemService} from './services/BOMItem.service';
import {RoutingService} from './services/Routing.service';
import {OperationService} from './services/Operation.service';
import {WorkOrderService} from './services/WorkOrder.service';
import {ProductionScheduleService} from './services/ProductionSchedule.service';
import {SupplierService} from './services/Supplier.service';
import {PurchaseOrderService} from './services/PurchaseOrder.service';
import {PurchaseOrderLineService} from './services/PurchaseOrderLine.service';
import {GoodsReceiptService} from './services/GoodsReceipt.service';
import {GoodsReceiptLineService} from './services/GoodsReceiptLine.service';
import {WarehouseService} from './services/Warehouse.service';
import {LocationService} from './services/Location.service';
import {InventoryItemService} from './services/InventoryItem.service';
import {InventoryTransactionService} from './services/InventoryTransaction.service';
import {CustomerService} from './services/Customer.service';
import {SalesOrderService} from './services/SalesOrder.service';
import {SalesOrderLineService} from './services/SalesOrderLine.service';
import {QualitySpecificationService} from './services/QualitySpecification.service';
import {InspectionPlanService} from './services/InspectionPlan.service';
import {InspectionCharacteristicService} from './services/InspectionCharacteristic.service';
import {InspectionLotService} from './services/InspectionLot.service';
import {InspectionResultService} from './services/InspectionResult.service';
import {NonconformanceService} from './services/Nonconformance.service';
import {CorrectiveActionService} from './services/CorrectiveAction.service';
import {AssetService} from './services/Asset.service';
import {MaintenancePlanService} from './services/MaintenancePlan.service';
import {MaintenanceOrderService} from './services/MaintenanceOrder.service';
import {EmployeeService} from './services/Employee.service';
import {ShiftService} from './services/Shift.service';
import {ShiftAssignmentService} from './services/ShiftAssignment.service';
import {ForecastService} from './services/Forecast.service';
import {ForecastLineService} from './services/ForecastLine.service';
import {MRPRunService} from './services/MRPRun.service';
import {PlannedOrderService} from './services/PlannedOrder.service';

@NgModule({
  declarations: [
    IndexEnterpriseComponent,
    CreateEnterpriseComponent,
    EditEnterpriseComponent,
    IndexBusinessUnitComponent,
    CreateBusinessUnitComponent,
    EditBusinessUnitComponent,
    IndexPlantComponent,
    CreatePlantComponent,
    EditPlantComponent,
    IndexProductionLineComponent,
    CreateProductionLineComponent,
    EditProductionLineComponent,
    IndexWorkCenterComponent,
    CreateWorkCenterComponent,
    EditWorkCenterComponent,
    IndexItemComponent,
    CreateItemComponent,
    EditItemComponent,
    IndexBOMComponent,
    CreateBOMComponent,
    EditBOMComponent,
    IndexBOMItemComponent,
    CreateBOMItemComponent,
    EditBOMItemComponent,
    IndexRoutingComponent,
    CreateRoutingComponent,
    EditRoutingComponent,
    IndexOperationComponent,
    CreateOperationComponent,
    EditOperationComponent,
    IndexWorkOrderComponent,
    CreateWorkOrderComponent,
    EditWorkOrderComponent,
    IndexProductionScheduleComponent,
    CreateProductionScheduleComponent,
    EditProductionScheduleComponent,
    IndexSupplierComponent,
    CreateSupplierComponent,
    EditSupplierComponent,
    IndexPurchaseOrderComponent,
    CreatePurchaseOrderComponent,
    EditPurchaseOrderComponent,
    IndexPurchaseOrderLineComponent,
    CreatePurchaseOrderLineComponent,
    EditPurchaseOrderLineComponent,
    IndexGoodsReceiptComponent,
    CreateGoodsReceiptComponent,
    EditGoodsReceiptComponent,
    IndexGoodsReceiptLineComponent,
    CreateGoodsReceiptLineComponent,
    EditGoodsReceiptLineComponent,
    IndexWarehouseComponent,
    CreateWarehouseComponent,
    EditWarehouseComponent,
    IndexLocationComponent,
    CreateLocationComponent,
    EditLocationComponent,
    IndexInventoryItemComponent,
    CreateInventoryItemComponent,
    EditInventoryItemComponent,
    IndexInventoryTransactionComponent,
    CreateInventoryTransactionComponent,
    EditInventoryTransactionComponent,
    IndexCustomerComponent,
    CreateCustomerComponent,
    EditCustomerComponent,
    IndexSalesOrderComponent,
    CreateSalesOrderComponent,
    EditSalesOrderComponent,
    IndexSalesOrderLineComponent,
    CreateSalesOrderLineComponent,
    EditSalesOrderLineComponent,
    IndexQualitySpecificationComponent,
    CreateQualitySpecificationComponent,
    EditQualitySpecificationComponent,
    IndexInspectionPlanComponent,
    CreateInspectionPlanComponent,
    EditInspectionPlanComponent,
    IndexInspectionCharacteristicComponent,
    CreateInspectionCharacteristicComponent,
    EditInspectionCharacteristicComponent,
    IndexInspectionLotComponent,
    CreateInspectionLotComponent,
    EditInspectionLotComponent,
    IndexInspectionResultComponent,
    CreateInspectionResultComponent,
    EditInspectionResultComponent,
    IndexNonconformanceComponent,
    CreateNonconformanceComponent,
    EditNonconformanceComponent,
    IndexCorrectiveActionComponent,
    CreateCorrectiveActionComponent,
    EditCorrectiveActionComponent,
    IndexAssetComponent,
    CreateAssetComponent,
    EditAssetComponent,
    IndexMaintenancePlanComponent,
    CreateMaintenancePlanComponent,
    EditMaintenancePlanComponent,
    IndexMaintenanceOrderComponent,
    CreateMaintenanceOrderComponent,
    EditMaintenanceOrderComponent,
    IndexEmployeeComponent,
    CreateEmployeeComponent,
    EditEmployeeComponent,
    IndexShiftComponent,
    CreateShiftComponent,
    EditShiftComponent,
    IndexShiftAssignmentComponent,
    CreateShiftAssignmentComponent,
    EditShiftAssignmentComponent,
    IndexForecastComponent,
    CreateForecastComponent,
    EditForecastComponent,
    IndexForecastLineComponent,
    CreateForecastLineComponent,
    EditForecastLineComponent,
    IndexMRPRunComponent,
    CreateMRPRunComponent,
    EditMRPRunComponent,
    IndexPlannedOrderComponent,
    CreatePlannedOrderComponent,
    EditPlannedOrderComponent,
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
    RouterModule.forRoot(appRoutes.EnterpriseRoutes), 
    RouterModule.forRoot(appRoutes.BusinessUnitRoutes), 
    RouterModule.forRoot(appRoutes.PlantRoutes), 
    RouterModule.forRoot(appRoutes.ProductionLineRoutes), 
    RouterModule.forRoot(appRoutes.WorkCenterRoutes), 
    RouterModule.forRoot(appRoutes.ItemRoutes), 
    RouterModule.forRoot(appRoutes.BOMRoutes), 
    RouterModule.forRoot(appRoutes.BOMItemRoutes), 
    RouterModule.forRoot(appRoutes.RoutingRoutes), 
    RouterModule.forRoot(appRoutes.OperationRoutes), 
    RouterModule.forRoot(appRoutes.WorkOrderRoutes), 
    RouterModule.forRoot(appRoutes.ProductionScheduleRoutes), 
    RouterModule.forRoot(appRoutes.SupplierRoutes), 
    RouterModule.forRoot(appRoutes.PurchaseOrderRoutes), 
    RouterModule.forRoot(appRoutes.PurchaseOrderLineRoutes), 
    RouterModule.forRoot(appRoutes.GoodsReceiptRoutes), 
    RouterModule.forRoot(appRoutes.GoodsReceiptLineRoutes), 
    RouterModule.forRoot(appRoutes.WarehouseRoutes), 
    RouterModule.forRoot(appRoutes.LocationRoutes), 
    RouterModule.forRoot(appRoutes.InventoryItemRoutes), 
    RouterModule.forRoot(appRoutes.InventoryTransactionRoutes), 
    RouterModule.forRoot(appRoutes.CustomerRoutes), 
    RouterModule.forRoot(appRoutes.SalesOrderRoutes), 
    RouterModule.forRoot(appRoutes.SalesOrderLineRoutes), 
    RouterModule.forRoot(appRoutes.QualitySpecificationRoutes), 
    RouterModule.forRoot(appRoutes.InspectionPlanRoutes), 
    RouterModule.forRoot(appRoutes.InspectionCharacteristicRoutes), 
    RouterModule.forRoot(appRoutes.InspectionLotRoutes), 
    RouterModule.forRoot(appRoutes.InspectionResultRoutes), 
    RouterModule.forRoot(appRoutes.NonconformanceRoutes), 
    RouterModule.forRoot(appRoutes.CorrectiveActionRoutes), 
    RouterModule.forRoot(appRoutes.AssetRoutes), 
    RouterModule.forRoot(appRoutes.MaintenancePlanRoutes), 
    RouterModule.forRoot(appRoutes.MaintenanceOrderRoutes), 
    RouterModule.forRoot(appRoutes.EmployeeRoutes), 
    RouterModule.forRoot(appRoutes.ShiftRoutes), 
    RouterModule.forRoot(appRoutes.ShiftAssignmentRoutes), 
    RouterModule.forRoot(appRoutes.ForecastRoutes), 
    RouterModule.forRoot(appRoutes.ForecastLineRoutes), 
    RouterModule.forRoot(appRoutes.MRPRunRoutes), 
    RouterModule.forRoot(appRoutes.PlannedOrderRoutes), 
  ],
  providers: [EnterpriseService,BusinessUnitService,PlantService,ProductionLineService,WorkCenterService,ItemService,BOMService,BOMItemService,RoutingService,OperationService,WorkOrderService,ProductionScheduleService,SupplierService,PurchaseOrderService,PurchaseOrderLineService,GoodsReceiptService,GoodsReceiptLineService,WarehouseService,LocationService,InventoryItemService,InventoryTransactionService,CustomerService,SalesOrderService,SalesOrderLineService,QualitySpecificationService,InspectionPlanService,InspectionCharacteristicService,InspectionLotService,InspectionResultService,NonconformanceService,CorrectiveActionService,AssetService,MaintenancePlanService,MaintenanceOrderService,EmployeeService,ShiftService,ShiftAssignmentService,ForecastService,ForecastLineService,MRPRunService,PlannedOrderService],
  bootstrap: [AppComponent]
})
export class AppModule { }
