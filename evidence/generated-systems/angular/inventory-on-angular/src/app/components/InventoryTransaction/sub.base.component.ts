import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all InventoryTransaction Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initStockKeepingUnitList();
	this.initWarehouseList();
	this.initStorageLocationList();
	this.initLotList();
	this.initSerialNumberList();
	this.initReservationList();
	this.initTransferOrderList();
	this.initStockAdjustmentList();
	this.initCycleCountList();
  }
}
