// routerConfig.ts

import { Routes } from '@angular/router';
import { CreateStockKeepingUnitComponent } from './components/StockKeepingUnit/create/create.component';
import { EditStockKeepingUnitComponent } from './components/StockKeepingUnit/edit/edit.component';
import { IndexStockKeepingUnitComponent } from './components/StockKeepingUnit/index/index.component';
import { CreateWarehouseComponent } from './components/Warehouse/create/create.component';
import { EditWarehouseComponent } from './components/Warehouse/edit/edit.component';
import { IndexWarehouseComponent } from './components/Warehouse/index/index.component';
import { CreateStorageLocationComponent } from './components/StorageLocation/create/create.component';
import { EditStorageLocationComponent } from './components/StorageLocation/edit/edit.component';
import { IndexStorageLocationComponent } from './components/StorageLocation/index/index.component';
import { CreateInventoryItemComponent } from './components/InventoryItem/create/create.component';
import { EditInventoryItemComponent } from './components/InventoryItem/edit/edit.component';
import { IndexInventoryItemComponent } from './components/InventoryItem/index/index.component';
import { CreateLotComponent } from './components/Lot/create/create.component';
import { EditLotComponent } from './components/Lot/edit/edit.component';
import { IndexLotComponent } from './components/Lot/index/index.component';
import { CreateSerialNumberComponent } from './components/SerialNumber/create/create.component';
import { EditSerialNumberComponent } from './components/SerialNumber/edit/edit.component';
import { IndexSerialNumberComponent } from './components/SerialNumber/index/index.component';
import { CreateReservationComponent } from './components/Reservation/create/create.component';
import { EditReservationComponent } from './components/Reservation/edit/edit.component';
import { IndexReservationComponent } from './components/Reservation/index/index.component';
import { CreateDemandSignalComponent } from './components/DemandSignal/create/create.component';
import { EditDemandSignalComponent } from './components/DemandSignal/edit/edit.component';
import { IndexDemandSignalComponent } from './components/DemandSignal/index/index.component';
import { CreateInventoryTransactionComponent } from './components/InventoryTransaction/create/create.component';
import { EditInventoryTransactionComponent } from './components/InventoryTransaction/edit/edit.component';
import { IndexInventoryTransactionComponent } from './components/InventoryTransaction/index/index.component';
import { CreateTransferOrderComponent } from './components/TransferOrder/create/create.component';
import { EditTransferOrderComponent } from './components/TransferOrder/edit/edit.component';
import { IndexTransferOrderComponent } from './components/TransferOrder/index/index.component';
import { CreateTransferOrderLineComponent } from './components/TransferOrderLine/create/create.component';
import { EditTransferOrderLineComponent } from './components/TransferOrderLine/edit/edit.component';
import { IndexTransferOrderLineComponent } from './components/TransferOrderLine/index/index.component';
import { CreateStockAdjustmentComponent } from './components/StockAdjustment/create/create.component';
import { EditStockAdjustmentComponent } from './components/StockAdjustment/edit/edit.component';
import { IndexStockAdjustmentComponent } from './components/StockAdjustment/index/index.component';
import { CreateStockAdjustmentLineComponent } from './components/StockAdjustmentLine/create/create.component';
import { EditStockAdjustmentLineComponent } from './components/StockAdjustmentLine/edit/edit.component';
import { IndexStockAdjustmentLineComponent } from './components/StockAdjustmentLine/index/index.component';
import { CreateCycleCountComponent } from './components/CycleCount/create/create.component';
import { EditCycleCountComponent } from './components/CycleCount/edit/edit.component';
import { IndexCycleCountComponent } from './components/CycleCount/index/index.component';
import { CreateCycleCountEntryComponent } from './components/CycleCountEntry/create/create.component';
import { EditCycleCountEntryComponent } from './components/CycleCountEntry/edit/edit.component';
import { IndexCycleCountEntryComponent } from './components/CycleCountEntry/index/index.component';
import { CreateReplenishmentPolicyComponent } from './components/ReplenishmentPolicy/create/create.component';
import { EditReplenishmentPolicyComponent } from './components/ReplenishmentPolicy/edit/edit.component';
import { IndexReplenishmentPolicyComponent } from './components/ReplenishmentPolicy/index/index.component';
import { CreateUoMConversionComponent } from './components/UoMConversion/create/create.component';
import { EditUoMConversionComponent } from './components/UoMConversion/edit/edit.component';
import { IndexUoMConversionComponent } from './components/UoMConversion/index/index.component';
import { CreateInventoryThresholdAlertComponent } from './components/InventoryThresholdAlert/create/create.component';
import { EditInventoryThresholdAlertComponent } from './components/InventoryThresholdAlert/edit/edit.component';
import { IndexInventoryThresholdAlertComponent } from './components/InventoryThresholdAlert/index/index.component';
import { CreateQuarantineComponent } from './components/Quarantine/create/create.component';
import { EditQuarantineComponent } from './components/Quarantine/edit/edit.component';
import { IndexQuarantineComponent } from './components/Quarantine/index/index.component';
import { CreateExpirationPolicyComponent } from './components/ExpirationPolicy/create/create.component';
import { EditExpirationPolicyComponent } from './components/ExpirationPolicy/edit/edit.component';
import { IndexExpirationPolicyComponent } from './components/ExpirationPolicy/index/index.component';
import { CreateInboundShipmentComponent } from './components/InboundShipment/create/create.component';
import { EditInboundShipmentComponent } from './components/InboundShipment/edit/edit.component';
import { IndexInboundShipmentComponent } from './components/InboundShipment/index/index.component';
import { CreateInboundShipmentLineComponent } from './components/InboundShipmentLine/create/create.component';
import { EditInboundShipmentLineComponent } from './components/InboundShipmentLine/edit/edit.component';
import { IndexInboundShipmentLineComponent } from './components/InboundShipmentLine/index/index.component';
import { CreateOutboundAllocationComponent } from './components/OutboundAllocation/create/create.component';
import { EditOutboundAllocationComponent } from './components/OutboundAllocation/edit/edit.component';
import { IndexOutboundAllocationComponent } from './components/OutboundAllocation/index/index.component';

export const StockKeepingUnitRoutes: Routes = [
  { path: 'createStockKeepingUnit',
    component: CreateStockKeepingUnitComponent
  },
  {
    path: 'editStockKeepingUnit/:id',
    component: EditStockKeepingUnitComponent
  },
  { path: 'indexStockKeepingUnit',
    component: IndexStockKeepingUnitComponent
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
export const StorageLocationRoutes: Routes = [
  { path: 'createStorageLocation',
    component: CreateStorageLocationComponent
  },
  {
    path: 'editStorageLocation/:id',
    component: EditStorageLocationComponent
  },
  { path: 'indexStorageLocation',
    component: IndexStorageLocationComponent
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
export const LotRoutes: Routes = [
  { path: 'createLot',
    component: CreateLotComponent
  },
  {
    path: 'editLot/:id',
    component: EditLotComponent
  },
  { path: 'indexLot',
    component: IndexLotComponent
  }
];
export const SerialNumberRoutes: Routes = [
  { path: 'createSerialNumber',
    component: CreateSerialNumberComponent
  },
  {
    path: 'editSerialNumber/:id',
    component: EditSerialNumberComponent
  },
  { path: 'indexSerialNumber',
    component: IndexSerialNumberComponent
  }
];
export const ReservationRoutes: Routes = [
  { path: 'createReservation',
    component: CreateReservationComponent
  },
  {
    path: 'editReservation/:id',
    component: EditReservationComponent
  },
  { path: 'indexReservation',
    component: IndexReservationComponent
  }
];
export const DemandSignalRoutes: Routes = [
  { path: 'createDemandSignal',
    component: CreateDemandSignalComponent
  },
  {
    path: 'editDemandSignal/:id',
    component: EditDemandSignalComponent
  },
  { path: 'indexDemandSignal',
    component: IndexDemandSignalComponent
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
export const TransferOrderRoutes: Routes = [
  { path: 'createTransferOrder',
    component: CreateTransferOrderComponent
  },
  {
    path: 'editTransferOrder/:id',
    component: EditTransferOrderComponent
  },
  { path: 'indexTransferOrder',
    component: IndexTransferOrderComponent
  }
];
export const TransferOrderLineRoutes: Routes = [
  { path: 'createTransferOrderLine',
    component: CreateTransferOrderLineComponent
  },
  {
    path: 'editTransferOrderLine/:id',
    component: EditTransferOrderLineComponent
  },
  { path: 'indexTransferOrderLine',
    component: IndexTransferOrderLineComponent
  }
];
export const StockAdjustmentRoutes: Routes = [
  { path: 'createStockAdjustment',
    component: CreateStockAdjustmentComponent
  },
  {
    path: 'editStockAdjustment/:id',
    component: EditStockAdjustmentComponent
  },
  { path: 'indexStockAdjustment',
    component: IndexStockAdjustmentComponent
  }
];
export const StockAdjustmentLineRoutes: Routes = [
  { path: 'createStockAdjustmentLine',
    component: CreateStockAdjustmentLineComponent
  },
  {
    path: 'editStockAdjustmentLine/:id',
    component: EditStockAdjustmentLineComponent
  },
  { path: 'indexStockAdjustmentLine',
    component: IndexStockAdjustmentLineComponent
  }
];
export const CycleCountRoutes: Routes = [
  { path: 'createCycleCount',
    component: CreateCycleCountComponent
  },
  {
    path: 'editCycleCount/:id',
    component: EditCycleCountComponent
  },
  { path: 'indexCycleCount',
    component: IndexCycleCountComponent
  }
];
export const CycleCountEntryRoutes: Routes = [
  { path: 'createCycleCountEntry',
    component: CreateCycleCountEntryComponent
  },
  {
    path: 'editCycleCountEntry/:id',
    component: EditCycleCountEntryComponent
  },
  { path: 'indexCycleCountEntry',
    component: IndexCycleCountEntryComponent
  }
];
export const ReplenishmentPolicyRoutes: Routes = [
  { path: 'createReplenishmentPolicy',
    component: CreateReplenishmentPolicyComponent
  },
  {
    path: 'editReplenishmentPolicy/:id',
    component: EditReplenishmentPolicyComponent
  },
  { path: 'indexReplenishmentPolicy',
    component: IndexReplenishmentPolicyComponent
  }
];
export const UoMConversionRoutes: Routes = [
  { path: 'createUoMConversion',
    component: CreateUoMConversionComponent
  },
  {
    path: 'editUoMConversion/:id',
    component: EditUoMConversionComponent
  },
  { path: 'indexUoMConversion',
    component: IndexUoMConversionComponent
  }
];
export const InventoryThresholdAlertRoutes: Routes = [
  { path: 'createInventoryThresholdAlert',
    component: CreateInventoryThresholdAlertComponent
  },
  {
    path: 'editInventoryThresholdAlert/:id',
    component: EditInventoryThresholdAlertComponent
  },
  { path: 'indexInventoryThresholdAlert',
    component: IndexInventoryThresholdAlertComponent
  }
];
export const QuarantineRoutes: Routes = [
  { path: 'createQuarantine',
    component: CreateQuarantineComponent
  },
  {
    path: 'editQuarantine/:id',
    component: EditQuarantineComponent
  },
  { path: 'indexQuarantine',
    component: IndexQuarantineComponent
  }
];
export const ExpirationPolicyRoutes: Routes = [
  { path: 'createExpirationPolicy',
    component: CreateExpirationPolicyComponent
  },
  {
    path: 'editExpirationPolicy/:id',
    component: EditExpirationPolicyComponent
  },
  { path: 'indexExpirationPolicy',
    component: IndexExpirationPolicyComponent
  }
];
export const InboundShipmentRoutes: Routes = [
  { path: 'createInboundShipment',
    component: CreateInboundShipmentComponent
  },
  {
    path: 'editInboundShipment/:id',
    component: EditInboundShipmentComponent
  },
  { path: 'indexInboundShipment',
    component: IndexInboundShipmentComponent
  }
];
export const InboundShipmentLineRoutes: Routes = [
  { path: 'createInboundShipmentLine',
    component: CreateInboundShipmentLineComponent
  },
  {
    path: 'editInboundShipmentLine/:id',
    component: EditInboundShipmentLineComponent
  },
  { path: 'indexInboundShipmentLine',
    component: IndexInboundShipmentLineComponent
  }
];
export const OutboundAllocationRoutes: Routes = [
  { path: 'createOutboundAllocation',
    component: CreateOutboundAllocationComponent
  },
  {
    path: 'editOutboundAllocation/:id',
    component: EditOutboundAllocationComponent
  },
  { path: 'indexOutboundAllocation',
    component: IndexOutboundAllocationComponent
  }
];
