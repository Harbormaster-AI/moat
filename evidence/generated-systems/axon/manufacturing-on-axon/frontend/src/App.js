import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListEnterpriseComponent from './components/ListEnterpriseComponent';
import CreateEnterpriseComponent from './components/CreateEnterpriseComponent';
import ViewEnterpriseComponent from './components/ViewEnterpriseComponent';
import ListBusinessUnitComponent from './components/ListBusinessUnitComponent';
import CreateBusinessUnitComponent from './components/CreateBusinessUnitComponent';
import ViewBusinessUnitComponent from './components/ViewBusinessUnitComponent';
import ListPlantComponent from './components/ListPlantComponent';
import CreatePlantComponent from './components/CreatePlantComponent';
import ViewPlantComponent from './components/ViewPlantComponent';
import ListProductionLineComponent from './components/ListProductionLineComponent';
import CreateProductionLineComponent from './components/CreateProductionLineComponent';
import ViewProductionLineComponent from './components/ViewProductionLineComponent';
import ListWorkCenterComponent from './components/ListWorkCenterComponent';
import CreateWorkCenterComponent from './components/CreateWorkCenterComponent';
import ViewWorkCenterComponent from './components/ViewWorkCenterComponent';
import ListItemComponent from './components/ListItemComponent';
import CreateItemComponent from './components/CreateItemComponent';
import ViewItemComponent from './components/ViewItemComponent';
import ListBOMComponent from './components/ListBOMComponent';
import CreateBOMComponent from './components/CreateBOMComponent';
import ViewBOMComponent from './components/ViewBOMComponent';
import ListBOMItemComponent from './components/ListBOMItemComponent';
import CreateBOMItemComponent from './components/CreateBOMItemComponent';
import ViewBOMItemComponent from './components/ViewBOMItemComponent';
import ListRoutingComponent from './components/ListRoutingComponent';
import CreateRoutingComponent from './components/CreateRoutingComponent';
import ViewRoutingComponent from './components/ViewRoutingComponent';
import ListOperationComponent from './components/ListOperationComponent';
import CreateOperationComponent from './components/CreateOperationComponent';
import ViewOperationComponent from './components/ViewOperationComponent';
import ListWorkOrderComponent from './components/ListWorkOrderComponent';
import CreateWorkOrderComponent from './components/CreateWorkOrderComponent';
import ViewWorkOrderComponent from './components/ViewWorkOrderComponent';
import ListProductionScheduleComponent from './components/ListProductionScheduleComponent';
import CreateProductionScheduleComponent from './components/CreateProductionScheduleComponent';
import ViewProductionScheduleComponent from './components/ViewProductionScheduleComponent';
import ListSupplierComponent from './components/ListSupplierComponent';
import CreateSupplierComponent from './components/CreateSupplierComponent';
import ViewSupplierComponent from './components/ViewSupplierComponent';
import ListPurchaseOrderComponent from './components/ListPurchaseOrderComponent';
import CreatePurchaseOrderComponent from './components/CreatePurchaseOrderComponent';
import ViewPurchaseOrderComponent from './components/ViewPurchaseOrderComponent';
import ListPurchaseOrderLineComponent from './components/ListPurchaseOrderLineComponent';
import CreatePurchaseOrderLineComponent from './components/CreatePurchaseOrderLineComponent';
import ViewPurchaseOrderLineComponent from './components/ViewPurchaseOrderLineComponent';
import ListGoodsReceiptComponent from './components/ListGoodsReceiptComponent';
import CreateGoodsReceiptComponent from './components/CreateGoodsReceiptComponent';
import ViewGoodsReceiptComponent from './components/ViewGoodsReceiptComponent';
import ListGoodsReceiptLineComponent from './components/ListGoodsReceiptLineComponent';
import CreateGoodsReceiptLineComponent from './components/CreateGoodsReceiptLineComponent';
import ViewGoodsReceiptLineComponent from './components/ViewGoodsReceiptLineComponent';
import ListWarehouseComponent from './components/ListWarehouseComponent';
import CreateWarehouseComponent from './components/CreateWarehouseComponent';
import ViewWarehouseComponent from './components/ViewWarehouseComponent';
import ListLocationComponent from './components/ListLocationComponent';
import CreateLocationComponent from './components/CreateLocationComponent';
import ViewLocationComponent from './components/ViewLocationComponent';
import ListInventoryItemComponent from './components/ListInventoryItemComponent';
import CreateInventoryItemComponent from './components/CreateInventoryItemComponent';
import ViewInventoryItemComponent from './components/ViewInventoryItemComponent';
import ListInventoryTransactionComponent from './components/ListInventoryTransactionComponent';
import CreateInventoryTransactionComponent from './components/CreateInventoryTransactionComponent';
import ViewInventoryTransactionComponent from './components/ViewInventoryTransactionComponent';
import ListCustomerComponent from './components/ListCustomerComponent';
import CreateCustomerComponent from './components/CreateCustomerComponent';
import ViewCustomerComponent from './components/ViewCustomerComponent';
import ListSalesOrderComponent from './components/ListSalesOrderComponent';
import CreateSalesOrderComponent from './components/CreateSalesOrderComponent';
import ViewSalesOrderComponent from './components/ViewSalesOrderComponent';
import ListSalesOrderLineComponent from './components/ListSalesOrderLineComponent';
import CreateSalesOrderLineComponent from './components/CreateSalesOrderLineComponent';
import ViewSalesOrderLineComponent from './components/ViewSalesOrderLineComponent';
import ListQualitySpecificationComponent from './components/ListQualitySpecificationComponent';
import CreateQualitySpecificationComponent from './components/CreateQualitySpecificationComponent';
import ViewQualitySpecificationComponent from './components/ViewQualitySpecificationComponent';
import ListInspectionPlanComponent from './components/ListInspectionPlanComponent';
import CreateInspectionPlanComponent from './components/CreateInspectionPlanComponent';
import ViewInspectionPlanComponent from './components/ViewInspectionPlanComponent';
import ListInspectionCharacteristicComponent from './components/ListInspectionCharacteristicComponent';
import CreateInspectionCharacteristicComponent from './components/CreateInspectionCharacteristicComponent';
import ViewInspectionCharacteristicComponent from './components/ViewInspectionCharacteristicComponent';
import ListInspectionLotComponent from './components/ListInspectionLotComponent';
import CreateInspectionLotComponent from './components/CreateInspectionLotComponent';
import ViewInspectionLotComponent from './components/ViewInspectionLotComponent';
import ListInspectionResultComponent from './components/ListInspectionResultComponent';
import CreateInspectionResultComponent from './components/CreateInspectionResultComponent';
import ViewInspectionResultComponent from './components/ViewInspectionResultComponent';
import ListNonconformanceComponent from './components/ListNonconformanceComponent';
import CreateNonconformanceComponent from './components/CreateNonconformanceComponent';
import ViewNonconformanceComponent from './components/ViewNonconformanceComponent';
import ListCorrectiveActionComponent from './components/ListCorrectiveActionComponent';
import CreateCorrectiveActionComponent from './components/CreateCorrectiveActionComponent';
import ViewCorrectiveActionComponent from './components/ViewCorrectiveActionComponent';
import ListAssetComponent from './components/ListAssetComponent';
import CreateAssetComponent from './components/CreateAssetComponent';
import ViewAssetComponent from './components/ViewAssetComponent';
import ListMaintenancePlanComponent from './components/ListMaintenancePlanComponent';
import CreateMaintenancePlanComponent from './components/CreateMaintenancePlanComponent';
import ViewMaintenancePlanComponent from './components/ViewMaintenancePlanComponent';
import ListMaintenanceOrderComponent from './components/ListMaintenanceOrderComponent';
import CreateMaintenanceOrderComponent from './components/CreateMaintenanceOrderComponent';
import ViewMaintenanceOrderComponent from './components/ViewMaintenanceOrderComponent';
import ListEmployeeComponent from './components/ListEmployeeComponent';
import CreateEmployeeComponent from './components/CreateEmployeeComponent';
import ViewEmployeeComponent from './components/ViewEmployeeComponent';
import ListShiftComponent from './components/ListShiftComponent';
import CreateShiftComponent from './components/CreateShiftComponent';
import ViewShiftComponent from './components/ViewShiftComponent';
import ListShiftAssignmentComponent from './components/ListShiftAssignmentComponent';
import CreateShiftAssignmentComponent from './components/CreateShiftAssignmentComponent';
import ViewShiftAssignmentComponent from './components/ViewShiftAssignmentComponent';
import ListForecastComponent from './components/ListForecastComponent';
import CreateForecastComponent from './components/CreateForecastComponent';
import ViewForecastComponent from './components/ViewForecastComponent';
import ListForecastLineComponent from './components/ListForecastLineComponent';
import CreateForecastLineComponent from './components/CreateForecastLineComponent';
import ViewForecastLineComponent from './components/ViewForecastLineComponent';
import ListMRPRunComponent from './components/ListMRPRunComponent';
import CreateMRPRunComponent from './components/CreateMRPRunComponent';
import ViewMRPRunComponent from './components/ViewMRPRunComponent';
import ListPlannedOrderComponent from './components/ListPlannedOrderComponent';
import CreatePlannedOrderComponent from './components/CreatePlannedOrderComponent';
import ViewPlannedOrderComponent from './components/ViewPlannedOrderComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/enterprises" component = {ListEnterpriseComponent}></Route>
                            <Route path = "/add-enterprise/:id" component = {CreateEnterpriseComponent}></Route>
                            <Route path = "/view-enterprise/:id" component = {ViewEnterpriseComponent}></Route>
                          {/* <Route path = "/update-enterprise/:id" component = {UpdateEnterpriseComponent}></Route> */}
                            <Route path = "/businessUnits" component = {ListBusinessUnitComponent}></Route>
                            <Route path = "/add-businessUnit/:id" component = {CreateBusinessUnitComponent}></Route>
                            <Route path = "/view-businessUnit/:id" component = {ViewBusinessUnitComponent}></Route>
                          {/* <Route path = "/update-businessUnit/:id" component = {UpdateBusinessUnitComponent}></Route> */}
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
                            <Route path = "/items" component = {ListItemComponent}></Route>
                            <Route path = "/add-item/:id" component = {CreateItemComponent}></Route>
                            <Route path = "/view-item/:id" component = {ViewItemComponent}></Route>
                          {/* <Route path = "/update-item/:id" component = {UpdateItemComponent}></Route> */}
                            <Route path = "/bOMs" component = {ListBOMComponent}></Route>
                            <Route path = "/add-bOM/:id" component = {CreateBOMComponent}></Route>
                            <Route path = "/view-bOM/:id" component = {ViewBOMComponent}></Route>
                          {/* <Route path = "/update-bOM/:id" component = {UpdateBOMComponent}></Route> */}
                            <Route path = "/bOMItems" component = {ListBOMItemComponent}></Route>
                            <Route path = "/add-bOMItem/:id" component = {CreateBOMItemComponent}></Route>
                            <Route path = "/view-bOMItem/:id" component = {ViewBOMItemComponent}></Route>
                          {/* <Route path = "/update-bOMItem/:id" component = {UpdateBOMItemComponent}></Route> */}
                            <Route path = "/routings" component = {ListRoutingComponent}></Route>
                            <Route path = "/add-routing/:id" component = {CreateRoutingComponent}></Route>
                            <Route path = "/view-routing/:id" component = {ViewRoutingComponent}></Route>
                          {/* <Route path = "/update-routing/:id" component = {UpdateRoutingComponent}></Route> */}
                            <Route path = "/operations" component = {ListOperationComponent}></Route>
                            <Route path = "/add-operation/:id" component = {CreateOperationComponent}></Route>
                            <Route path = "/view-operation/:id" component = {ViewOperationComponent}></Route>
                          {/* <Route path = "/update-operation/:id" component = {UpdateOperationComponent}></Route> */}
                            <Route path = "/workOrders" component = {ListWorkOrderComponent}></Route>
                            <Route path = "/add-workOrder/:id" component = {CreateWorkOrderComponent}></Route>
                            <Route path = "/view-workOrder/:id" component = {ViewWorkOrderComponent}></Route>
                          {/* <Route path = "/update-workOrder/:id" component = {UpdateWorkOrderComponent}></Route> */}
                            <Route path = "/productionSchedules" component = {ListProductionScheduleComponent}></Route>
                            <Route path = "/add-productionSchedule/:id" component = {CreateProductionScheduleComponent}></Route>
                            <Route path = "/view-productionSchedule/:id" component = {ViewProductionScheduleComponent}></Route>
                          {/* <Route path = "/update-productionSchedule/:id" component = {UpdateProductionScheduleComponent}></Route> */}
                            <Route path = "/suppliers" component = {ListSupplierComponent}></Route>
                            <Route path = "/add-supplier/:id" component = {CreateSupplierComponent}></Route>
                            <Route path = "/view-supplier/:id" component = {ViewSupplierComponent}></Route>
                          {/* <Route path = "/update-supplier/:id" component = {UpdateSupplierComponent}></Route> */}
                            <Route path = "/purchaseOrders" component = {ListPurchaseOrderComponent}></Route>
                            <Route path = "/add-purchaseOrder/:id" component = {CreatePurchaseOrderComponent}></Route>
                            <Route path = "/view-purchaseOrder/:id" component = {ViewPurchaseOrderComponent}></Route>
                          {/* <Route path = "/update-purchaseOrder/:id" component = {UpdatePurchaseOrderComponent}></Route> */}
                            <Route path = "/purchaseOrderLines" component = {ListPurchaseOrderLineComponent}></Route>
                            <Route path = "/add-purchaseOrderLine/:id" component = {CreatePurchaseOrderLineComponent}></Route>
                            <Route path = "/view-purchaseOrderLine/:id" component = {ViewPurchaseOrderLineComponent}></Route>
                          {/* <Route path = "/update-purchaseOrderLine/:id" component = {UpdatePurchaseOrderLineComponent}></Route> */}
                            <Route path = "/goodsReceipts" component = {ListGoodsReceiptComponent}></Route>
                            <Route path = "/add-goodsReceipt/:id" component = {CreateGoodsReceiptComponent}></Route>
                            <Route path = "/view-goodsReceipt/:id" component = {ViewGoodsReceiptComponent}></Route>
                          {/* <Route path = "/update-goodsReceipt/:id" component = {UpdateGoodsReceiptComponent}></Route> */}
                            <Route path = "/goodsReceiptLines" component = {ListGoodsReceiptLineComponent}></Route>
                            <Route path = "/add-goodsReceiptLine/:id" component = {CreateGoodsReceiptLineComponent}></Route>
                            <Route path = "/view-goodsReceiptLine/:id" component = {ViewGoodsReceiptLineComponent}></Route>
                          {/* <Route path = "/update-goodsReceiptLine/:id" component = {UpdateGoodsReceiptLineComponent}></Route> */}
                            <Route path = "/warehouses" component = {ListWarehouseComponent}></Route>
                            <Route path = "/add-warehouse/:id" component = {CreateWarehouseComponent}></Route>
                            <Route path = "/view-warehouse/:id" component = {ViewWarehouseComponent}></Route>
                          {/* <Route path = "/update-warehouse/:id" component = {UpdateWarehouseComponent}></Route> */}
                            <Route path = "/locations" component = {ListLocationComponent}></Route>
                            <Route path = "/add-location/:id" component = {CreateLocationComponent}></Route>
                            <Route path = "/view-location/:id" component = {ViewLocationComponent}></Route>
                          {/* <Route path = "/update-location/:id" component = {UpdateLocationComponent}></Route> */}
                            <Route path = "/inventoryItems" component = {ListInventoryItemComponent}></Route>
                            <Route path = "/add-inventoryItem/:id" component = {CreateInventoryItemComponent}></Route>
                            <Route path = "/view-inventoryItem/:id" component = {ViewInventoryItemComponent}></Route>
                          {/* <Route path = "/update-inventoryItem/:id" component = {UpdateInventoryItemComponent}></Route> */}
                            <Route path = "/inventoryTransactions" component = {ListInventoryTransactionComponent}></Route>
                            <Route path = "/add-inventoryTransaction/:id" component = {CreateInventoryTransactionComponent}></Route>
                            <Route path = "/view-inventoryTransaction/:id" component = {ViewInventoryTransactionComponent}></Route>
                          {/* <Route path = "/update-inventoryTransaction/:id" component = {UpdateInventoryTransactionComponent}></Route> */}
                            <Route path = "/customers" component = {ListCustomerComponent}></Route>
                            <Route path = "/add-customer/:id" component = {CreateCustomerComponent}></Route>
                            <Route path = "/view-customer/:id" component = {ViewCustomerComponent}></Route>
                          {/* <Route path = "/update-customer/:id" component = {UpdateCustomerComponent}></Route> */}
                            <Route path = "/salesOrders" component = {ListSalesOrderComponent}></Route>
                            <Route path = "/add-salesOrder/:id" component = {CreateSalesOrderComponent}></Route>
                            <Route path = "/view-salesOrder/:id" component = {ViewSalesOrderComponent}></Route>
                          {/* <Route path = "/update-salesOrder/:id" component = {UpdateSalesOrderComponent}></Route> */}
                            <Route path = "/salesOrderLines" component = {ListSalesOrderLineComponent}></Route>
                            <Route path = "/add-salesOrderLine/:id" component = {CreateSalesOrderLineComponent}></Route>
                            <Route path = "/view-salesOrderLine/:id" component = {ViewSalesOrderLineComponent}></Route>
                          {/* <Route path = "/update-salesOrderLine/:id" component = {UpdateSalesOrderLineComponent}></Route> */}
                            <Route path = "/qualitySpecifications" component = {ListQualitySpecificationComponent}></Route>
                            <Route path = "/add-qualitySpecification/:id" component = {CreateQualitySpecificationComponent}></Route>
                            <Route path = "/view-qualitySpecification/:id" component = {ViewQualitySpecificationComponent}></Route>
                          {/* <Route path = "/update-qualitySpecification/:id" component = {UpdateQualitySpecificationComponent}></Route> */}
                            <Route path = "/inspectionPlans" component = {ListInspectionPlanComponent}></Route>
                            <Route path = "/add-inspectionPlan/:id" component = {CreateInspectionPlanComponent}></Route>
                            <Route path = "/view-inspectionPlan/:id" component = {ViewInspectionPlanComponent}></Route>
                          {/* <Route path = "/update-inspectionPlan/:id" component = {UpdateInspectionPlanComponent}></Route> */}
                            <Route path = "/inspectionCharacteristics" component = {ListInspectionCharacteristicComponent}></Route>
                            <Route path = "/add-inspectionCharacteristic/:id" component = {CreateInspectionCharacteristicComponent}></Route>
                            <Route path = "/view-inspectionCharacteristic/:id" component = {ViewInspectionCharacteristicComponent}></Route>
                          {/* <Route path = "/update-inspectionCharacteristic/:id" component = {UpdateInspectionCharacteristicComponent}></Route> */}
                            <Route path = "/inspectionLots" component = {ListInspectionLotComponent}></Route>
                            <Route path = "/add-inspectionLot/:id" component = {CreateInspectionLotComponent}></Route>
                            <Route path = "/view-inspectionLot/:id" component = {ViewInspectionLotComponent}></Route>
                          {/* <Route path = "/update-inspectionLot/:id" component = {UpdateInspectionLotComponent}></Route> */}
                            <Route path = "/inspectionResults" component = {ListInspectionResultComponent}></Route>
                            <Route path = "/add-inspectionResult/:id" component = {CreateInspectionResultComponent}></Route>
                            <Route path = "/view-inspectionResult/:id" component = {ViewInspectionResultComponent}></Route>
                          {/* <Route path = "/update-inspectionResult/:id" component = {UpdateInspectionResultComponent}></Route> */}
                            <Route path = "/nonconformances" component = {ListNonconformanceComponent}></Route>
                            <Route path = "/add-nonconformance/:id" component = {CreateNonconformanceComponent}></Route>
                            <Route path = "/view-nonconformance/:id" component = {ViewNonconformanceComponent}></Route>
                          {/* <Route path = "/update-nonconformance/:id" component = {UpdateNonconformanceComponent}></Route> */}
                            <Route path = "/correctiveActions" component = {ListCorrectiveActionComponent}></Route>
                            <Route path = "/add-correctiveAction/:id" component = {CreateCorrectiveActionComponent}></Route>
                            <Route path = "/view-correctiveAction/:id" component = {ViewCorrectiveActionComponent}></Route>
                          {/* <Route path = "/update-correctiveAction/:id" component = {UpdateCorrectiveActionComponent}></Route> */}
                            <Route path = "/assets" component = {ListAssetComponent}></Route>
                            <Route path = "/add-asset/:id" component = {CreateAssetComponent}></Route>
                            <Route path = "/view-asset/:id" component = {ViewAssetComponent}></Route>
                          {/* <Route path = "/update-asset/:id" component = {UpdateAssetComponent}></Route> */}
                            <Route path = "/maintenancePlans" component = {ListMaintenancePlanComponent}></Route>
                            <Route path = "/add-maintenancePlan/:id" component = {CreateMaintenancePlanComponent}></Route>
                            <Route path = "/view-maintenancePlan/:id" component = {ViewMaintenancePlanComponent}></Route>
                          {/* <Route path = "/update-maintenancePlan/:id" component = {UpdateMaintenancePlanComponent}></Route> */}
                            <Route path = "/maintenanceOrders" component = {ListMaintenanceOrderComponent}></Route>
                            <Route path = "/add-maintenanceOrder/:id" component = {CreateMaintenanceOrderComponent}></Route>
                            <Route path = "/view-maintenanceOrder/:id" component = {ViewMaintenanceOrderComponent}></Route>
                          {/* <Route path = "/update-maintenanceOrder/:id" component = {UpdateMaintenanceOrderComponent}></Route> */}
                            <Route path = "/employees" component = {ListEmployeeComponent}></Route>
                            <Route path = "/add-employee/:id" component = {CreateEmployeeComponent}></Route>
                            <Route path = "/view-employee/:id" component = {ViewEmployeeComponent}></Route>
                          {/* <Route path = "/update-employee/:id" component = {UpdateEmployeeComponent}></Route> */}
                            <Route path = "/shifts" component = {ListShiftComponent}></Route>
                            <Route path = "/add-shift/:id" component = {CreateShiftComponent}></Route>
                            <Route path = "/view-shift/:id" component = {ViewShiftComponent}></Route>
                          {/* <Route path = "/update-shift/:id" component = {UpdateShiftComponent}></Route> */}
                            <Route path = "/shiftAssignments" component = {ListShiftAssignmentComponent}></Route>
                            <Route path = "/add-shiftAssignment/:id" component = {CreateShiftAssignmentComponent}></Route>
                            <Route path = "/view-shiftAssignment/:id" component = {ViewShiftAssignmentComponent}></Route>
                          {/* <Route path = "/update-shiftAssignment/:id" component = {UpdateShiftAssignmentComponent}></Route> */}
                            <Route path = "/forecasts" component = {ListForecastComponent}></Route>
                            <Route path = "/add-forecast/:id" component = {CreateForecastComponent}></Route>
                            <Route path = "/view-forecast/:id" component = {ViewForecastComponent}></Route>
                          {/* <Route path = "/update-forecast/:id" component = {UpdateForecastComponent}></Route> */}
                            <Route path = "/forecastLines" component = {ListForecastLineComponent}></Route>
                            <Route path = "/add-forecastLine/:id" component = {CreateForecastLineComponent}></Route>
                            <Route path = "/view-forecastLine/:id" component = {ViewForecastLineComponent}></Route>
                          {/* <Route path = "/update-forecastLine/:id" component = {UpdateForecastLineComponent}></Route> */}
                            <Route path = "/mRPRuns" component = {ListMRPRunComponent}></Route>
                            <Route path = "/add-mRPRun/:id" component = {CreateMRPRunComponent}></Route>
                            <Route path = "/view-mRPRun/:id" component = {ViewMRPRunComponent}></Route>
                          {/* <Route path = "/update-mRPRun/:id" component = {UpdateMRPRunComponent}></Route> */}
                            <Route path = "/plannedOrders" component = {ListPlannedOrderComponent}></Route>
                            <Route path = "/add-plannedOrder/:id" component = {CreatePlannedOrderComponent}></Route>
                            <Route path = "/view-plannedOrder/:id" component = {ViewPlannedOrderComponent}></Route>
                          {/* <Route path = "/update-plannedOrder/:id" component = {UpdatePlannedOrderComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
