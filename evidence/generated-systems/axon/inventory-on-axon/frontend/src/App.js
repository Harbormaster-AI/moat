import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListStockKeepingUnitComponent from './components/ListStockKeepingUnitComponent';
import CreateStockKeepingUnitComponent from './components/CreateStockKeepingUnitComponent';
import ViewStockKeepingUnitComponent from './components/ViewStockKeepingUnitComponent';
import ListWarehouseComponent from './components/ListWarehouseComponent';
import CreateWarehouseComponent from './components/CreateWarehouseComponent';
import ViewWarehouseComponent from './components/ViewWarehouseComponent';
import ListStorageLocationComponent from './components/ListStorageLocationComponent';
import CreateStorageLocationComponent from './components/CreateStorageLocationComponent';
import ViewStorageLocationComponent from './components/ViewStorageLocationComponent';
import ListInventoryItemComponent from './components/ListInventoryItemComponent';
import CreateInventoryItemComponent from './components/CreateInventoryItemComponent';
import ViewInventoryItemComponent from './components/ViewInventoryItemComponent';
import ListLotComponent from './components/ListLotComponent';
import CreateLotComponent from './components/CreateLotComponent';
import ViewLotComponent from './components/ViewLotComponent';
import ListSerialNumberComponent from './components/ListSerialNumberComponent';
import CreateSerialNumberComponent from './components/CreateSerialNumberComponent';
import ViewSerialNumberComponent from './components/ViewSerialNumberComponent';
import ListReservationComponent from './components/ListReservationComponent';
import CreateReservationComponent from './components/CreateReservationComponent';
import ViewReservationComponent from './components/ViewReservationComponent';
import ListDemandSignalComponent from './components/ListDemandSignalComponent';
import CreateDemandSignalComponent from './components/CreateDemandSignalComponent';
import ViewDemandSignalComponent from './components/ViewDemandSignalComponent';
import ListInventoryTransactionComponent from './components/ListInventoryTransactionComponent';
import CreateInventoryTransactionComponent from './components/CreateInventoryTransactionComponent';
import ViewInventoryTransactionComponent from './components/ViewInventoryTransactionComponent';
import ListTransferOrderComponent from './components/ListTransferOrderComponent';
import CreateTransferOrderComponent from './components/CreateTransferOrderComponent';
import ViewTransferOrderComponent from './components/ViewTransferOrderComponent';
import ListTransferOrderLineComponent from './components/ListTransferOrderLineComponent';
import CreateTransferOrderLineComponent from './components/CreateTransferOrderLineComponent';
import ViewTransferOrderLineComponent from './components/ViewTransferOrderLineComponent';
import ListStockAdjustmentComponent from './components/ListStockAdjustmentComponent';
import CreateStockAdjustmentComponent from './components/CreateStockAdjustmentComponent';
import ViewStockAdjustmentComponent from './components/ViewStockAdjustmentComponent';
import ListStockAdjustmentLineComponent from './components/ListStockAdjustmentLineComponent';
import CreateStockAdjustmentLineComponent from './components/CreateStockAdjustmentLineComponent';
import ViewStockAdjustmentLineComponent from './components/ViewStockAdjustmentLineComponent';
import ListCycleCountComponent from './components/ListCycleCountComponent';
import CreateCycleCountComponent from './components/CreateCycleCountComponent';
import ViewCycleCountComponent from './components/ViewCycleCountComponent';
import ListCycleCountEntryComponent from './components/ListCycleCountEntryComponent';
import CreateCycleCountEntryComponent from './components/CreateCycleCountEntryComponent';
import ViewCycleCountEntryComponent from './components/ViewCycleCountEntryComponent';
import ListReplenishmentPolicyComponent from './components/ListReplenishmentPolicyComponent';
import CreateReplenishmentPolicyComponent from './components/CreateReplenishmentPolicyComponent';
import ViewReplenishmentPolicyComponent from './components/ViewReplenishmentPolicyComponent';
import ListUoMConversionComponent from './components/ListUoMConversionComponent';
import CreateUoMConversionComponent from './components/CreateUoMConversionComponent';
import ViewUoMConversionComponent from './components/ViewUoMConversionComponent';
import ListInventoryThresholdAlertComponent from './components/ListInventoryThresholdAlertComponent';
import CreateInventoryThresholdAlertComponent from './components/CreateInventoryThresholdAlertComponent';
import ViewInventoryThresholdAlertComponent from './components/ViewInventoryThresholdAlertComponent';
import ListQuarantineComponent from './components/ListQuarantineComponent';
import CreateQuarantineComponent from './components/CreateQuarantineComponent';
import ViewQuarantineComponent from './components/ViewQuarantineComponent';
import ListExpirationPolicyComponent from './components/ListExpirationPolicyComponent';
import CreateExpirationPolicyComponent from './components/CreateExpirationPolicyComponent';
import ViewExpirationPolicyComponent from './components/ViewExpirationPolicyComponent';
import ListInboundShipmentComponent from './components/ListInboundShipmentComponent';
import CreateInboundShipmentComponent from './components/CreateInboundShipmentComponent';
import ViewInboundShipmentComponent from './components/ViewInboundShipmentComponent';
import ListInboundShipmentLineComponent from './components/ListInboundShipmentLineComponent';
import CreateInboundShipmentLineComponent from './components/CreateInboundShipmentLineComponent';
import ViewInboundShipmentLineComponent from './components/ViewInboundShipmentLineComponent';
import ListOutboundAllocationComponent from './components/ListOutboundAllocationComponent';
import CreateOutboundAllocationComponent from './components/CreateOutboundAllocationComponent';
import ViewOutboundAllocationComponent from './components/ViewOutboundAllocationComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/stockKeepingUnits" component = {ListStockKeepingUnitComponent}></Route>
                            <Route path = "/add-stockKeepingUnit/:id" component = {CreateStockKeepingUnitComponent}></Route>
                            <Route path = "/view-stockKeepingUnit/:id" component = {ViewStockKeepingUnitComponent}></Route>
                          {/* <Route path = "/update-stockKeepingUnit/:id" component = {UpdateStockKeepingUnitComponent}></Route> */}
                            <Route path = "/warehouses" component = {ListWarehouseComponent}></Route>
                            <Route path = "/add-warehouse/:id" component = {CreateWarehouseComponent}></Route>
                            <Route path = "/view-warehouse/:id" component = {ViewWarehouseComponent}></Route>
                          {/* <Route path = "/update-warehouse/:id" component = {UpdateWarehouseComponent}></Route> */}
                            <Route path = "/storageLocations" component = {ListStorageLocationComponent}></Route>
                            <Route path = "/add-storageLocation/:id" component = {CreateStorageLocationComponent}></Route>
                            <Route path = "/view-storageLocation/:id" component = {ViewStorageLocationComponent}></Route>
                          {/* <Route path = "/update-storageLocation/:id" component = {UpdateStorageLocationComponent}></Route> */}
                            <Route path = "/inventoryItems" component = {ListInventoryItemComponent}></Route>
                            <Route path = "/add-inventoryItem/:id" component = {CreateInventoryItemComponent}></Route>
                            <Route path = "/view-inventoryItem/:id" component = {ViewInventoryItemComponent}></Route>
                          {/* <Route path = "/update-inventoryItem/:id" component = {UpdateInventoryItemComponent}></Route> */}
                            <Route path = "/lots" component = {ListLotComponent}></Route>
                            <Route path = "/add-lot/:id" component = {CreateLotComponent}></Route>
                            <Route path = "/view-lot/:id" component = {ViewLotComponent}></Route>
                          {/* <Route path = "/update-lot/:id" component = {UpdateLotComponent}></Route> */}
                            <Route path = "/serialNumbers" component = {ListSerialNumberComponent}></Route>
                            <Route path = "/add-serialNumber/:id" component = {CreateSerialNumberComponent}></Route>
                            <Route path = "/view-serialNumber/:id" component = {ViewSerialNumberComponent}></Route>
                          {/* <Route path = "/update-serialNumber/:id" component = {UpdateSerialNumberComponent}></Route> */}
                            <Route path = "/reservations" component = {ListReservationComponent}></Route>
                            <Route path = "/add-reservation/:id" component = {CreateReservationComponent}></Route>
                            <Route path = "/view-reservation/:id" component = {ViewReservationComponent}></Route>
                          {/* <Route path = "/update-reservation/:id" component = {UpdateReservationComponent}></Route> */}
                            <Route path = "/demandSignals" component = {ListDemandSignalComponent}></Route>
                            <Route path = "/add-demandSignal/:id" component = {CreateDemandSignalComponent}></Route>
                            <Route path = "/view-demandSignal/:id" component = {ViewDemandSignalComponent}></Route>
                          {/* <Route path = "/update-demandSignal/:id" component = {UpdateDemandSignalComponent}></Route> */}
                            <Route path = "/inventoryTransactions" component = {ListInventoryTransactionComponent}></Route>
                            <Route path = "/add-inventoryTransaction/:id" component = {CreateInventoryTransactionComponent}></Route>
                            <Route path = "/view-inventoryTransaction/:id" component = {ViewInventoryTransactionComponent}></Route>
                          {/* <Route path = "/update-inventoryTransaction/:id" component = {UpdateInventoryTransactionComponent}></Route> */}
                            <Route path = "/transferOrders" component = {ListTransferOrderComponent}></Route>
                            <Route path = "/add-transferOrder/:id" component = {CreateTransferOrderComponent}></Route>
                            <Route path = "/view-transferOrder/:id" component = {ViewTransferOrderComponent}></Route>
                          {/* <Route path = "/update-transferOrder/:id" component = {UpdateTransferOrderComponent}></Route> */}
                            <Route path = "/transferOrderLines" component = {ListTransferOrderLineComponent}></Route>
                            <Route path = "/add-transferOrderLine/:id" component = {CreateTransferOrderLineComponent}></Route>
                            <Route path = "/view-transferOrderLine/:id" component = {ViewTransferOrderLineComponent}></Route>
                          {/* <Route path = "/update-transferOrderLine/:id" component = {UpdateTransferOrderLineComponent}></Route> */}
                            <Route path = "/stockAdjustments" component = {ListStockAdjustmentComponent}></Route>
                            <Route path = "/add-stockAdjustment/:id" component = {CreateStockAdjustmentComponent}></Route>
                            <Route path = "/view-stockAdjustment/:id" component = {ViewStockAdjustmentComponent}></Route>
                          {/* <Route path = "/update-stockAdjustment/:id" component = {UpdateStockAdjustmentComponent}></Route> */}
                            <Route path = "/stockAdjustmentLines" component = {ListStockAdjustmentLineComponent}></Route>
                            <Route path = "/add-stockAdjustmentLine/:id" component = {CreateStockAdjustmentLineComponent}></Route>
                            <Route path = "/view-stockAdjustmentLine/:id" component = {ViewStockAdjustmentLineComponent}></Route>
                          {/* <Route path = "/update-stockAdjustmentLine/:id" component = {UpdateStockAdjustmentLineComponent}></Route> */}
                            <Route path = "/cycleCounts" component = {ListCycleCountComponent}></Route>
                            <Route path = "/add-cycleCount/:id" component = {CreateCycleCountComponent}></Route>
                            <Route path = "/view-cycleCount/:id" component = {ViewCycleCountComponent}></Route>
                          {/* <Route path = "/update-cycleCount/:id" component = {UpdateCycleCountComponent}></Route> */}
                            <Route path = "/cycleCountEntrys" component = {ListCycleCountEntryComponent}></Route>
                            <Route path = "/add-cycleCountEntry/:id" component = {CreateCycleCountEntryComponent}></Route>
                            <Route path = "/view-cycleCountEntry/:id" component = {ViewCycleCountEntryComponent}></Route>
                          {/* <Route path = "/update-cycleCountEntry/:id" component = {UpdateCycleCountEntryComponent}></Route> */}
                            <Route path = "/replenishmentPolicys" component = {ListReplenishmentPolicyComponent}></Route>
                            <Route path = "/add-replenishmentPolicy/:id" component = {CreateReplenishmentPolicyComponent}></Route>
                            <Route path = "/view-replenishmentPolicy/:id" component = {ViewReplenishmentPolicyComponent}></Route>
                          {/* <Route path = "/update-replenishmentPolicy/:id" component = {UpdateReplenishmentPolicyComponent}></Route> */}
                            <Route path = "/uoMConversions" component = {ListUoMConversionComponent}></Route>
                            <Route path = "/add-uoMConversion/:id" component = {CreateUoMConversionComponent}></Route>
                            <Route path = "/view-uoMConversion/:id" component = {ViewUoMConversionComponent}></Route>
                          {/* <Route path = "/update-uoMConversion/:id" component = {UpdateUoMConversionComponent}></Route> */}
                            <Route path = "/inventoryThresholdAlerts" component = {ListInventoryThresholdAlertComponent}></Route>
                            <Route path = "/add-inventoryThresholdAlert/:id" component = {CreateInventoryThresholdAlertComponent}></Route>
                            <Route path = "/view-inventoryThresholdAlert/:id" component = {ViewInventoryThresholdAlertComponent}></Route>
                          {/* <Route path = "/update-inventoryThresholdAlert/:id" component = {UpdateInventoryThresholdAlertComponent}></Route> */}
                            <Route path = "/quarantines" component = {ListQuarantineComponent}></Route>
                            <Route path = "/add-quarantine/:id" component = {CreateQuarantineComponent}></Route>
                            <Route path = "/view-quarantine/:id" component = {ViewQuarantineComponent}></Route>
                          {/* <Route path = "/update-quarantine/:id" component = {UpdateQuarantineComponent}></Route> */}
                            <Route path = "/expirationPolicys" component = {ListExpirationPolicyComponent}></Route>
                            <Route path = "/add-expirationPolicy/:id" component = {CreateExpirationPolicyComponent}></Route>
                            <Route path = "/view-expirationPolicy/:id" component = {ViewExpirationPolicyComponent}></Route>
                          {/* <Route path = "/update-expirationPolicy/:id" component = {UpdateExpirationPolicyComponent}></Route> */}
                            <Route path = "/inboundShipments" component = {ListInboundShipmentComponent}></Route>
                            <Route path = "/add-inboundShipment/:id" component = {CreateInboundShipmentComponent}></Route>
                            <Route path = "/view-inboundShipment/:id" component = {ViewInboundShipmentComponent}></Route>
                          {/* <Route path = "/update-inboundShipment/:id" component = {UpdateInboundShipmentComponent}></Route> */}
                            <Route path = "/inboundShipmentLines" component = {ListInboundShipmentLineComponent}></Route>
                            <Route path = "/add-inboundShipmentLine/:id" component = {CreateInboundShipmentLineComponent}></Route>
                            <Route path = "/view-inboundShipmentLine/:id" component = {ViewInboundShipmentLineComponent}></Route>
                          {/* <Route path = "/update-inboundShipmentLine/:id" component = {UpdateInboundShipmentLineComponent}></Route> */}
                            <Route path = "/outboundAllocations" component = {ListOutboundAllocationComponent}></Route>
                            <Route path = "/add-outboundAllocation/:id" component = {CreateOutboundAllocationComponent}></Route>
                            <Route path = "/view-outboundAllocation/:id" component = {ViewOutboundAllocationComponent}></Route>
                          {/* <Route path = "/update-outboundAllocation/:id" component = {UpdateOutboundAllocationComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
