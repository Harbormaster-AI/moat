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

import {IndexStockKeepingUnitComponent} from './components/StockKeepingUnit/index/index.component';
import {CreateStockKeepingUnitComponent} from './components/StockKeepingUnit/create/create.component';
import {EditStockKeepingUnitComponent} from './components/StockKeepingUnit/edit/edit.component';
import {IndexWarehouseComponent} from './components/Warehouse/index/index.component';
import {CreateWarehouseComponent} from './components/Warehouse/create/create.component';
import {EditWarehouseComponent} from './components/Warehouse/edit/edit.component';
import {IndexStorageLocationComponent} from './components/StorageLocation/index/index.component';
import {CreateStorageLocationComponent} from './components/StorageLocation/create/create.component';
import {EditStorageLocationComponent} from './components/StorageLocation/edit/edit.component';
import {IndexInventoryItemComponent} from './components/InventoryItem/index/index.component';
import {CreateInventoryItemComponent} from './components/InventoryItem/create/create.component';
import {EditInventoryItemComponent} from './components/InventoryItem/edit/edit.component';
import {IndexLotComponent} from './components/Lot/index/index.component';
import {CreateLotComponent} from './components/Lot/create/create.component';
import {EditLotComponent} from './components/Lot/edit/edit.component';
import {IndexSerialNumberComponent} from './components/SerialNumber/index/index.component';
import {CreateSerialNumberComponent} from './components/SerialNumber/create/create.component';
import {EditSerialNumberComponent} from './components/SerialNumber/edit/edit.component';
import {IndexReservationComponent} from './components/Reservation/index/index.component';
import {CreateReservationComponent} from './components/Reservation/create/create.component';
import {EditReservationComponent} from './components/Reservation/edit/edit.component';
import {IndexDemandSignalComponent} from './components/DemandSignal/index/index.component';
import {CreateDemandSignalComponent} from './components/DemandSignal/create/create.component';
import {EditDemandSignalComponent} from './components/DemandSignal/edit/edit.component';
import {IndexInventoryTransactionComponent} from './components/InventoryTransaction/index/index.component';
import {CreateInventoryTransactionComponent} from './components/InventoryTransaction/create/create.component';
import {EditInventoryTransactionComponent} from './components/InventoryTransaction/edit/edit.component';
import {IndexTransferOrderComponent} from './components/TransferOrder/index/index.component';
import {CreateTransferOrderComponent} from './components/TransferOrder/create/create.component';
import {EditTransferOrderComponent} from './components/TransferOrder/edit/edit.component';
import {IndexTransferOrderLineComponent} from './components/TransferOrderLine/index/index.component';
import {CreateTransferOrderLineComponent} from './components/TransferOrderLine/create/create.component';
import {EditTransferOrderLineComponent} from './components/TransferOrderLine/edit/edit.component';
import {IndexStockAdjustmentComponent} from './components/StockAdjustment/index/index.component';
import {CreateStockAdjustmentComponent} from './components/StockAdjustment/create/create.component';
import {EditStockAdjustmentComponent} from './components/StockAdjustment/edit/edit.component';
import {IndexStockAdjustmentLineComponent} from './components/StockAdjustmentLine/index/index.component';
import {CreateStockAdjustmentLineComponent} from './components/StockAdjustmentLine/create/create.component';
import {EditStockAdjustmentLineComponent} from './components/StockAdjustmentLine/edit/edit.component';
import {IndexCycleCountComponent} from './components/CycleCount/index/index.component';
import {CreateCycleCountComponent} from './components/CycleCount/create/create.component';
import {EditCycleCountComponent} from './components/CycleCount/edit/edit.component';
import {IndexCycleCountEntryComponent} from './components/CycleCountEntry/index/index.component';
import {CreateCycleCountEntryComponent} from './components/CycleCountEntry/create/create.component';
import {EditCycleCountEntryComponent} from './components/CycleCountEntry/edit/edit.component';
import {IndexReplenishmentPolicyComponent} from './components/ReplenishmentPolicy/index/index.component';
import {CreateReplenishmentPolicyComponent} from './components/ReplenishmentPolicy/create/create.component';
import {EditReplenishmentPolicyComponent} from './components/ReplenishmentPolicy/edit/edit.component';
import {IndexUoMConversionComponent} from './components/UoMConversion/index/index.component';
import {CreateUoMConversionComponent} from './components/UoMConversion/create/create.component';
import {EditUoMConversionComponent} from './components/UoMConversion/edit/edit.component';
import {IndexInventoryThresholdAlertComponent} from './components/InventoryThresholdAlert/index/index.component';
import {CreateInventoryThresholdAlertComponent} from './components/InventoryThresholdAlert/create/create.component';
import {EditInventoryThresholdAlertComponent} from './components/InventoryThresholdAlert/edit/edit.component';
import {IndexQuarantineComponent} from './components/Quarantine/index/index.component';
import {CreateQuarantineComponent} from './components/Quarantine/create/create.component';
import {EditQuarantineComponent} from './components/Quarantine/edit/edit.component';
import {IndexExpirationPolicyComponent} from './components/ExpirationPolicy/index/index.component';
import {CreateExpirationPolicyComponent} from './components/ExpirationPolicy/create/create.component';
import {EditExpirationPolicyComponent} from './components/ExpirationPolicy/edit/edit.component';
import {IndexInboundShipmentComponent} from './components/InboundShipment/index/index.component';
import {CreateInboundShipmentComponent} from './components/InboundShipment/create/create.component';
import {EditInboundShipmentComponent} from './components/InboundShipment/edit/edit.component';
import {IndexInboundShipmentLineComponent} from './components/InboundShipmentLine/index/index.component';
import {CreateInboundShipmentLineComponent} from './components/InboundShipmentLine/create/create.component';
import {EditInboundShipmentLineComponent} from './components/InboundShipmentLine/edit/edit.component';
import {IndexOutboundAllocationComponent} from './components/OutboundAllocation/index/index.component';
import {CreateOutboundAllocationComponent} from './components/OutboundAllocation/create/create.component';
import {EditOutboundAllocationComponent} from './components/OutboundAllocation/edit/edit.component';

import * as appRoutes from './routerConfig';

import {StockKeepingUnitService} from './services/StockKeepingUnit.service';
import {WarehouseService} from './services/Warehouse.service';
import {StorageLocationService} from './services/StorageLocation.service';
import {InventoryItemService} from './services/InventoryItem.service';
import {LotService} from './services/Lot.service';
import {SerialNumberService} from './services/SerialNumber.service';
import {ReservationService} from './services/Reservation.service';
import {DemandSignalService} from './services/DemandSignal.service';
import {InventoryTransactionService} from './services/InventoryTransaction.service';
import {TransferOrderService} from './services/TransferOrder.service';
import {TransferOrderLineService} from './services/TransferOrderLine.service';
import {StockAdjustmentService} from './services/StockAdjustment.service';
import {StockAdjustmentLineService} from './services/StockAdjustmentLine.service';
import {CycleCountService} from './services/CycleCount.service';
import {CycleCountEntryService} from './services/CycleCountEntry.service';
import {ReplenishmentPolicyService} from './services/ReplenishmentPolicy.service';
import {UoMConversionService} from './services/UoMConversion.service';
import {InventoryThresholdAlertService} from './services/InventoryThresholdAlert.service';
import {QuarantineService} from './services/Quarantine.service';
import {ExpirationPolicyService} from './services/ExpirationPolicy.service';
import {InboundShipmentService} from './services/InboundShipment.service';
import {InboundShipmentLineService} from './services/InboundShipmentLine.service';
import {OutboundAllocationService} from './services/OutboundAllocation.service';

@NgModule({
  declarations: [
    IndexStockKeepingUnitComponent,
    CreateStockKeepingUnitComponent,
    EditStockKeepingUnitComponent,
    IndexWarehouseComponent,
    CreateWarehouseComponent,
    EditWarehouseComponent,
    IndexStorageLocationComponent,
    CreateStorageLocationComponent,
    EditStorageLocationComponent,
    IndexInventoryItemComponent,
    CreateInventoryItemComponent,
    EditInventoryItemComponent,
    IndexLotComponent,
    CreateLotComponent,
    EditLotComponent,
    IndexSerialNumberComponent,
    CreateSerialNumberComponent,
    EditSerialNumberComponent,
    IndexReservationComponent,
    CreateReservationComponent,
    EditReservationComponent,
    IndexDemandSignalComponent,
    CreateDemandSignalComponent,
    EditDemandSignalComponent,
    IndexInventoryTransactionComponent,
    CreateInventoryTransactionComponent,
    EditInventoryTransactionComponent,
    IndexTransferOrderComponent,
    CreateTransferOrderComponent,
    EditTransferOrderComponent,
    IndexTransferOrderLineComponent,
    CreateTransferOrderLineComponent,
    EditTransferOrderLineComponent,
    IndexStockAdjustmentComponent,
    CreateStockAdjustmentComponent,
    EditStockAdjustmentComponent,
    IndexStockAdjustmentLineComponent,
    CreateStockAdjustmentLineComponent,
    EditStockAdjustmentLineComponent,
    IndexCycleCountComponent,
    CreateCycleCountComponent,
    EditCycleCountComponent,
    IndexCycleCountEntryComponent,
    CreateCycleCountEntryComponent,
    EditCycleCountEntryComponent,
    IndexReplenishmentPolicyComponent,
    CreateReplenishmentPolicyComponent,
    EditReplenishmentPolicyComponent,
    IndexUoMConversionComponent,
    CreateUoMConversionComponent,
    EditUoMConversionComponent,
    IndexInventoryThresholdAlertComponent,
    CreateInventoryThresholdAlertComponent,
    EditInventoryThresholdAlertComponent,
    IndexQuarantineComponent,
    CreateQuarantineComponent,
    EditQuarantineComponent,
    IndexExpirationPolicyComponent,
    CreateExpirationPolicyComponent,
    EditExpirationPolicyComponent,
    IndexInboundShipmentComponent,
    CreateInboundShipmentComponent,
    EditInboundShipmentComponent,
    IndexInboundShipmentLineComponent,
    CreateInboundShipmentLineComponent,
    EditInboundShipmentLineComponent,
    IndexOutboundAllocationComponent,
    CreateOutboundAllocationComponent,
    EditOutboundAllocationComponent,
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
    RouterModule.forRoot(appRoutes.StockKeepingUnitRoutes), 
    RouterModule.forRoot(appRoutes.WarehouseRoutes), 
    RouterModule.forRoot(appRoutes.StorageLocationRoutes), 
    RouterModule.forRoot(appRoutes.InventoryItemRoutes), 
    RouterModule.forRoot(appRoutes.LotRoutes), 
    RouterModule.forRoot(appRoutes.SerialNumberRoutes), 
    RouterModule.forRoot(appRoutes.ReservationRoutes), 
    RouterModule.forRoot(appRoutes.DemandSignalRoutes), 
    RouterModule.forRoot(appRoutes.InventoryTransactionRoutes), 
    RouterModule.forRoot(appRoutes.TransferOrderRoutes), 
    RouterModule.forRoot(appRoutes.TransferOrderLineRoutes), 
    RouterModule.forRoot(appRoutes.StockAdjustmentRoutes), 
    RouterModule.forRoot(appRoutes.StockAdjustmentLineRoutes), 
    RouterModule.forRoot(appRoutes.CycleCountRoutes), 
    RouterModule.forRoot(appRoutes.CycleCountEntryRoutes), 
    RouterModule.forRoot(appRoutes.ReplenishmentPolicyRoutes), 
    RouterModule.forRoot(appRoutes.UoMConversionRoutes), 
    RouterModule.forRoot(appRoutes.InventoryThresholdAlertRoutes), 
    RouterModule.forRoot(appRoutes.QuarantineRoutes), 
    RouterModule.forRoot(appRoutes.ExpirationPolicyRoutes), 
    RouterModule.forRoot(appRoutes.InboundShipmentRoutes), 
    RouterModule.forRoot(appRoutes.InboundShipmentLineRoutes), 
    RouterModule.forRoot(appRoutes.OutboundAllocationRoutes), 
  ],
  providers: [StockKeepingUnitService,WarehouseService,StorageLocationService,InventoryItemService,LotService,SerialNumberService,ReservationService,DemandSignalService,InventoryTransactionService,TransferOrderService,TransferOrderLineService,StockAdjustmentService,StockAdjustmentLineService,CycleCountService,CycleCountEntryService,ReplenishmentPolicyService,UoMConversionService,InventoryThresholdAlertService,QuarantineService,ExpirationPolicyService,InboundShipmentService,InboundShipmentLineService,OutboundAllocationService],
  bootstrap: [AppComponent]
})
export class AppModule { }
