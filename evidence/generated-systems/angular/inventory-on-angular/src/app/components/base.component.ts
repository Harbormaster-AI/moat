import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {WarehouseService} from '../services/Warehouse.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import {LotService} from '../services/Lot.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import {ReservationService} from '../services/Reservation.service';
import {DemandSignalService} from '../services/DemandSignal.service';
import {InventoryTransactionService} from '../services/InventoryTransaction.service';
import {TransferOrderService} from '../services/TransferOrder.service';
import {TransferOrderLineService} from '../services/TransferOrderLine.service';
import {StockAdjustmentService} from '../services/StockAdjustment.service';
import {StockAdjustmentLineService} from '../services/StockAdjustmentLine.service';
import {CycleCountService} from '../services/CycleCount.service';
import {CycleCountEntryService} from '../services/CycleCountEntry.service';
import {ReplenishmentPolicyService} from '../services/ReplenishmentPolicy.service';
import {UoMConversionService} from '../services/UoMConversion.service';
import {InventoryThresholdAlertService} from '../services/InventoryThresholdAlert.service';
import {QuarantineService} from '../services/Quarantine.service';
import {ExpirationPolicyService} from '../services/ExpirationPolicy.service';
import {InboundShipmentService} from '../services/InboundShipment.service';
import {InboundShipmentLineService} from '../services/InboundShipmentLine.service';
import {OutboundAllocationService} from '../services/OutboundAllocation.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    ItemTypes = Object.keys(enumTypes.ItemType);
    UnitOfMeasures = Object.keys(enumTypes.UnitOfMeasure);
    StockStatuss = Object.keys(enumTypes.StockStatus);
    LocationTypes = Object.keys(enumTypes.LocationType);
    LotStatuss = Object.keys(enumTypes.LotStatus);
    SerialStatuss = Object.keys(enumTypes.SerialStatus);
    ReservationStatuss = Object.keys(enumTypes.ReservationStatus);
    ReservationTypes = Object.keys(enumTypes.ReservationType);
    DemandTypes = Object.keys(enumTypes.DemandType);
    TransactionTypes = Object.keys(enumTypes.TransactionType);
    TransactionStatuss = Object.keys(enumTypes.TransactionStatus);
    TransferOrderStatuss = Object.keys(enumTypes.TransferOrderStatus);
    AdjustmentTypes = Object.keys(enumTypes.AdjustmentType);
    AdjustmentStatuss = Object.keys(enumTypes.AdjustmentStatus);
    CountStatuss = Object.keys(enumTypes.CountStatus);
    ReplenishmentPolicyTypes = Object.keys(enumTypes.ReplenishmentPolicyType);
    InventoryAlertTypes = Object.keys(enumTypes.InventoryAlertType);
    AlertStatuss = Object.keys(enumTypes.AlertStatus);
    Dispositions = Object.keys(enumTypes.Disposition);
    RotationMethods = Object.keys(enumTypes.RotationMethod);
    InboundShipmentStatuss = Object.keys(enumTypes.InboundShipmentStatus);
    AllocationStatuss = Object.keys(enumTypes.AllocationStatus);

// all collection instances
    stockKeepingUnits : any;
    warehouses : any;
    storageLocations : any;
    inventoryItems : any;
    lots : any;
    serialNumbers : any;
    reservations : any;
    demandSignals : any;
    inventoryTransactions : any;
    transferOrders : any;
    transferOrderLines : any;
    stockAdjustments : any;
    stockAdjustmentLines : any;
    cycleCounts : any;
    cycleCountEntrys : any;
    replenishmentPolicys : any;
    uoMConversions : any;
    inventoryThresholdAlerts : any;
    quarantines : any;
    expirationPolicys : any;
    inboundShipments : any;
    inboundShipmentLines : any;
    outboundAllocations : any;
  
// initialization  
    ngOnInit() {
    }

    initStockKeepingUnitList() {
        if ( this.stockKeepingUnits == null ) {
            new StockKeepingUnitService(this.http).getStockKeepingUnits().subscribe(res => {
                this.stockKeepingUnits = res;
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
    
    initStorageLocationList() {
        if ( this.storageLocations == null ) {
            new StorageLocationService(this.http).getStorageLocations().subscribe(res => {
                this.storageLocations = res;
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
    
    initLotList() {
        if ( this.lots == null ) {
            new LotService(this.http).getLots().subscribe(res => {
                this.lots = res;
            });
        }
    }
    
    initSerialNumberList() {
        if ( this.serialNumbers == null ) {
            new SerialNumberService(this.http).getSerialNumbers().subscribe(res => {
                this.serialNumbers = res;
            });
        }
    }
    
    initReservationList() {
        if ( this.reservations == null ) {
            new ReservationService(this.http).getReservations().subscribe(res => {
                this.reservations = res;
            });
        }
    }
    
    initDemandSignalList() {
        if ( this.demandSignals == null ) {
            new DemandSignalService(this.http).getDemandSignals().subscribe(res => {
                this.demandSignals = res;
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
    
    initTransferOrderList() {
        if ( this.transferOrders == null ) {
            new TransferOrderService(this.http).getTransferOrders().subscribe(res => {
                this.transferOrders = res;
            });
        }
    }
    
    initTransferOrderLineList() {
        if ( this.transferOrderLines == null ) {
            new TransferOrderLineService(this.http).getTransferOrderLines().subscribe(res => {
                this.transferOrderLines = res;
            });
        }
    }
    
    initStockAdjustmentList() {
        if ( this.stockAdjustments == null ) {
            new StockAdjustmentService(this.http).getStockAdjustments().subscribe(res => {
                this.stockAdjustments = res;
            });
        }
    }
    
    initStockAdjustmentLineList() {
        if ( this.stockAdjustmentLines == null ) {
            new StockAdjustmentLineService(this.http).getStockAdjustmentLines().subscribe(res => {
                this.stockAdjustmentLines = res;
            });
        }
    }
    
    initCycleCountList() {
        if ( this.cycleCounts == null ) {
            new CycleCountService(this.http).getCycleCounts().subscribe(res => {
                this.cycleCounts = res;
            });
        }
    }
    
    initCycleCountEntryList() {
        if ( this.cycleCountEntrys == null ) {
            new CycleCountEntryService(this.http).getCycleCountEntrys().subscribe(res => {
                this.cycleCountEntrys = res;
            });
        }
    }
    
    initReplenishmentPolicyList() {
        if ( this.replenishmentPolicys == null ) {
            new ReplenishmentPolicyService(this.http).getReplenishmentPolicys().subscribe(res => {
                this.replenishmentPolicys = res;
            });
        }
    }
    
    initUoMConversionList() {
        if ( this.uoMConversions == null ) {
            new UoMConversionService(this.http).getUoMConversions().subscribe(res => {
                this.uoMConversions = res;
            });
        }
    }
    
    initInventoryThresholdAlertList() {
        if ( this.inventoryThresholdAlerts == null ) {
            new InventoryThresholdAlertService(this.http).getInventoryThresholdAlerts().subscribe(res => {
                this.inventoryThresholdAlerts = res;
            });
        }
    }
    
    initQuarantineList() {
        if ( this.quarantines == null ) {
            new QuarantineService(this.http).getQuarantines().subscribe(res => {
                this.quarantines = res;
            });
        }
    }
    
    initExpirationPolicyList() {
        if ( this.expirationPolicys == null ) {
            new ExpirationPolicyService(this.http).getExpirationPolicys().subscribe(res => {
                this.expirationPolicys = res;
            });
        }
    }
    
    initInboundShipmentList() {
        if ( this.inboundShipments == null ) {
            new InboundShipmentService(this.http).getInboundShipments().subscribe(res => {
                this.inboundShipments = res;
            });
        }
    }
    
    initInboundShipmentLineList() {
        if ( this.inboundShipmentLines == null ) {
            new InboundShipmentLineService(this.http).getInboundShipmentLines().subscribe(res => {
                this.inboundShipmentLines = res;
            });
        }
    }
    
    initOutboundAllocationList() {
        if ( this.outboundAllocations == null ) {
            new OutboundAllocationService(this.http).getOutboundAllocations().subscribe(res => {
                this.outboundAllocations = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
