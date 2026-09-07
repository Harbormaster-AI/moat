import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InventoryTransaction} from '../models/InventoryTransaction';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {WarehouseService} from '../services/Warehouse.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import {LotService} from '../services/Lot.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import {ReservationService} from '../services/Reservation.service';
import {TransferOrderService} from '../services/TransferOrder.service';
import {StockAdjustmentService} from '../services/StockAdjustment.service';
import {CycleCountService} from '../services/CycleCount.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InventoryTransactionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inventoryTransaction : InventoryTransaction;

	//********************************************************************
	// Catch all for the return value of a service call
	//********************************************************************
	result: any;

	//********************************************************************
	// sole constructor, injected with the HttpClient
	//********************************************************************
	constructor(private http: HttpClient) {
		super();
	}

		//********************************************************************
	// add a InventoryTransaction
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInventoryTransaction(transactionNumber, quantity, unitCost, transactionDate, reasonCode, Sku, Warehouse, Location, Lot, SerialNumbers, RelatedReservation, TransferOrder, Adjustment, CycleCount, TransactionType, UnitOfMeasure, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryTransaction/create';
		const obj = {
			      		transactionNumber: transactionNumber,
      		quantity: quantity,
      		unitCost: unitCost,
      		transactionDate: transactionDate,
      		reasonCode: reasonCode,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		RelatedReservation: RelatedReservation != null && RelatedReservation.length > 0 ? RelatedReservation : null,
      		TransferOrder: TransferOrder != null && TransferOrder.length > 0 ? TransferOrder : null,
      		Adjustment: Adjustment != null && Adjustment.length > 0 ? Adjustment : null,
      		CycleCount: CycleCount != null && CycleCount.length > 0 ? CycleCount : null,
      		TransactionType: TransactionType,
      		UnitOfMeasure: UnitOfMeasure,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInventoryTransaction(transactionNumber, quantity, unitCost, transactionDate, reasonCode, Sku, Warehouse, Location, Lot, SerialNumbers, RelatedReservation, TransferOrder, Adjustment, CycleCount, TransactionType, UnitOfMeasure, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InventoryTransaction/update/' + id;
		const obj = {
				      		transactionNumber: transactionNumber,
      		quantity: quantity,
      		unitCost: unitCost,
      		transactionDate: transactionDate,
      		reasonCode: reasonCode,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		RelatedReservation: RelatedReservation != null && RelatedReservation.length > 0 ? RelatedReservation : null,
      		TransferOrder: TransferOrder != null && TransferOrder.length > 0 ? TransferOrder : null,
      		Adjustment: Adjustment != null && Adjustment.length > 0 ? Adjustment : null,
      		CycleCount: CycleCount != null && CycleCount.length > 0 ? CycleCount : null,
      		TransactionType: TransactionType,
      		UnitOfMeasure: UnitOfMeasure,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInventoryTransaction(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryTransaction/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InventoryTransaction
	// returns the results untouched as an Observable InventoryTransaction
	// InventoryTransaction model
	// delegates via URI
	//********************************************************************
	getInventoryTransaction(id) : Observable<InventoryTransaction> {
		const uri_ = this.apiUrl + '/InventoryTransaction/load/' + id;

		return this.http.get<InventoryTransaction>(uri_);
	}
	
	//********************************************************************
	// gets all InventoryTransaction
	// returns the results untouched as JSON representation of an
	// Observable array of InventoryTransaction models
	// delegates via URI
	//********************************************************************
	getInventoryTransactions() : Observable<InventoryTransaction[]> {
		const uri_ = this.apiUrl + '/InventoryTransaction/';

		return this
			.http.get<InventoryTransaction[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Sku on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( inventoryTransactionId, _skuId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.inventoryTransaction.sku = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign Sku to null
	this.inventoryTransaction.sku = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Warehouse on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( inventoryTransactionId, _warehouseId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.inventoryTransaction.warehouse = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign Warehouse to null
	this.inventoryTransaction.warehouse = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Location on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLocation( inventoryTransactionId, _locationId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_locationId);

	// assign the Location
	this.inventoryTransaction.location = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Location on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLocation( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign Location to null
	this.inventoryTransaction.location = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lot on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLot( inventoryTransactionId, _lotId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the Lot from storage
	var tmp 	= new LotService(this.http).getLot(_lotId);

	// assign the Lot
	this.inventoryTransaction.lot = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lot on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLot( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign Lot to null
	this.inventoryTransaction.lot = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a RelatedReservation on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRelatedReservation( inventoryTransactionId, _relatedReservationId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the Reservation from storage
	var tmp 	= new ReservationService(this.http).getReservation(_relatedReservationId);

	// assign the RelatedReservation
	this.inventoryTransaction.relatedReservation = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a RelatedReservation on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRelatedReservation( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign RelatedReservation to null
	this.inventoryTransaction.relatedReservation = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a TransferOrder on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTransferOrder( inventoryTransactionId, _transferOrderId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the TransferOrder from storage
	var tmp 	= new TransferOrderService(this.http).getTransferOrder(_transferOrderId);

	// assign the TransferOrder
	this.inventoryTransaction.transferOrder = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TransferOrder on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTransferOrder( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign TransferOrder to null
	this.inventoryTransaction.transferOrder = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Adjustment on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdjustment( inventoryTransactionId, _adjustmentId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the StockAdjustment from storage
	var tmp 	= new StockAdjustmentService(this.http).getStockAdjustment(_adjustmentId);

	// assign the Adjustment
	this.inventoryTransaction.adjustment = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Adjustment on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdjustment( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign Adjustment to null
	this.inventoryTransaction.adjustment = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CycleCount on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCycleCount( inventoryTransactionId, _cycleCountId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the CycleCount from storage
	var tmp 	= new CycleCountService(this.http).getCycleCount(_cycleCountId);

	// assign the CycleCount
	this.inventoryTransaction.cycleCount = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CycleCount on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCycleCount( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign CycleCount to null
	this.inventoryTransaction.cycleCount = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more serialNumbersIds as a SerialNumbers
	// to a InventoryTransaction
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSerialNumbers( inventoryTransactionId, serialNumbersIds ): Observable<any> {

		// get the InventoryTransaction
		this.loadHelper( inventoryTransactionId );

	// split on a comma with no spaces
	var idList = serialNumbersIds.split(',')

	// iterate over array of serialNumbers ids
	idList.forEach(function (id) {
		// read the SerialNumber
		var serialNumber = new SerialNumberService(this.http).getSerialNumber(id);
		// add the SerialNumber if not already assigned
		if ( this.inventoryTransaction.serialNumbers.indexOf(serialNumber) == -1 )
		this.inventoryTransaction.serialNumbers.push(serialNumber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serialNumbersIds as a SerialNumbers
	// from a InventoryTransaction
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSerialNumbers( inventoryTransactionId, serialNumbersIds ): Observable<any> {

		// get the InventoryTransaction
		this.loadHelper( inventoryTransactionId );


	// split on a comma with no spaces
	var idList 					= serialNumbersIds.split(',');
	var serialNumbers 	= this.inventoryTransaction.serialNumbers;

	if ( serialNumbers != null && serialNumbersIds != null ) {

		// iterate over array of serialNumbers ids
		serialNumbers.forEach(function (obj) {
			if ( serialNumbersIds.indexOf(obj._id) > -1 ) {
				// remove the SerialNumber
				this.inventoryTransaction.serialNumbers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InventoryTransaction
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InventoryTransaction/update/' + this.inventoryTransaction;

	return  this.http.post(uri_, this.inventoryTransaction );
}

	//********************************************************************
	// loadHelper - internal helper to load a InventoryTransaction
	//********************************************************************	
	loadHelper( id ) {
		this.getInventoryTransaction(id)
			.subscribe((res : InventoryTransaction) => {
				this.inventoryTransaction = res;
			});
	}
}