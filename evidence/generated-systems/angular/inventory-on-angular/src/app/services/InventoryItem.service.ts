import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InventoryItem} from '../models/InventoryItem';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {WarehouseService} from '../services/Warehouse.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import {LotService} from '../services/Lot.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import {InventoryTransactionService} from '../services/InventoryTransaction.service';
import {ReservationService} from '../services/Reservation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InventoryItemService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inventoryItem : InventoryItem;

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
	// add a InventoryItem
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInventoryItem(quantityOnHand, quantityAvailable, quantityReserved, unitCost, lastUpdated, Sku, Warehouse, Location, Lot, SerialNumbers, Transactions, Reservations, StockStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryItem/create';
		const obj = {
			      		quantityOnHand: quantityOnHand,
      		quantityAvailable: quantityAvailable,
      		quantityReserved: quantityReserved,
      		unitCost: unitCost,
      		lastUpdated: lastUpdated,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
      		Reservations: Reservations != null && Reservations.length > 0 ? Reservations : null,
			StockStatus: StockStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInventoryItem(quantityOnHand, quantityAvailable, quantityReserved, unitCost, lastUpdated, Sku, Warehouse, Location, Lot, SerialNumbers, Transactions, Reservations, StockStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InventoryItem/update/' + id;
		const obj = {
				      		quantityOnHand: quantityOnHand,
      		quantityAvailable: quantityAvailable,
      		quantityReserved: quantityReserved,
      		unitCost: unitCost,
      		lastUpdated: lastUpdated,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
      		Reservations: Reservations != null && Reservations.length > 0 ? Reservations : null,
			StockStatus: StockStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInventoryItem(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryItem/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InventoryItem
	// returns the results untouched as an Observable InventoryItem
	// InventoryItem model
	// delegates via URI
	//********************************************************************
	getInventoryItem(id) : Observable<InventoryItem> {
		const uri_ = this.apiUrl + '/InventoryItem/load/' + id;

		return this.http.get<InventoryItem>(uri_);
	}
	
	//********************************************************************
	// gets all InventoryItem
	// returns the results untouched as JSON representation of an
	// Observable array of InventoryItem models
	// delegates via URI
	//********************************************************************
	getInventoryItems() : Observable<InventoryItem[]> {
		const uri_ = this.apiUrl + '/InventoryItem/';

		return this
			.http.get<InventoryItem[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Sku on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( inventoryItemId, _skuId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.inventoryItem.sku = tmp;

	// save the InventoryItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( inventoryItemId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// assign Sku to null
	this.inventoryItem.sku = null;

	// save the InventoryItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Warehouse on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( inventoryItemId, _warehouseId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.inventoryItem.warehouse = tmp;

	// save the InventoryItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( inventoryItemId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// assign Warehouse to null
	this.inventoryItem.warehouse = null;

	// save the InventoryItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Location on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLocation( inventoryItemId, _locationId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_locationId);

	// assign the Location
	this.inventoryItem.location = tmp;

	// save the InventoryItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Location on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLocation( inventoryItemId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// assign Location to null
	this.inventoryItem.location = null;

	// save the InventoryItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lot on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLot( inventoryItemId, _lotId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// get the Lot from storage
	var tmp 	= new LotService(this.http).getLot(_lotId);

	// assign the Lot
	this.inventoryItem.lot = tmp;

	// save the InventoryItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lot on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLot( inventoryItemId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// assign Lot to null
	this.inventoryItem.lot = null;

	// save the InventoryItem
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more serialNumbersIds as a SerialNumbers
	// to a InventoryItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSerialNumbers( inventoryItemId, serialNumbersIds ): Observable<any> {

		// get the InventoryItem
		this.loadHelper( inventoryItemId );

	// split on a comma with no spaces
	var idList = serialNumbersIds.split(',')

	// iterate over array of serialNumbers ids
	idList.forEach(function (id) {
		// read the SerialNumber
		var serialNumber = new SerialNumberService(this.http).getSerialNumber(id);
		// add the SerialNumber if not already assigned
		if ( this.inventoryItem.serialNumbers.indexOf(serialNumber) == -1 )
		this.inventoryItem.serialNumbers.push(serialNumber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serialNumbersIds as a SerialNumbers
	// from a InventoryItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSerialNumbers( inventoryItemId, serialNumbersIds ): Observable<any> {

		// get the InventoryItem
		this.loadHelper( inventoryItemId );


	// split on a comma with no spaces
	var idList 					= serialNumbersIds.split(',');
	var serialNumbers 	= this.inventoryItem.serialNumbers;

	if ( serialNumbers != null && serialNumbersIds != null ) {

		// iterate over array of serialNumbers ids
		serialNumbers.forEach(function (obj) {
			if ( serialNumbersIds.indexOf(obj._id) > -1 ) {
				// remove the SerialNumber
				this.inventoryItem.serialNumbers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more transactionsIds as a Transactions
	// to a InventoryItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTransactions( inventoryItemId, transactionsIds ): Observable<any> {

		// get the InventoryItem
		this.loadHelper( inventoryItemId );

	// split on a comma with no spaces
	var idList = transactionsIds.split(',')

	// iterate over array of transactions ids
	idList.forEach(function (id) {
		// read the InventoryTransaction
		var inventoryTransaction = new InventoryTransactionService(this.http).getInventoryTransaction(id);
		// add the InventoryTransaction if not already assigned
		if ( this.inventoryItem.transactions.indexOf(inventoryTransaction) == -1 )
		this.inventoryItem.transactions.push(inventoryTransaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more transactionsIds as a Transactions
	// from a InventoryItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTransactions( inventoryItemId, transactionsIds ): Observable<any> {

		// get the InventoryItem
		this.loadHelper( inventoryItemId );


	// split on a comma with no spaces
	var idList 					= transactionsIds.split(',');
	var transactions 	= this.inventoryItem.transactions;

	if ( transactions != null && transactionsIds != null ) {

		// iterate over array of transactions ids
		transactions.forEach(function (obj) {
			if ( transactionsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryTransaction
				this.inventoryItem.transactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reservationsIds as a Reservations
	// to a InventoryItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReservations( inventoryItemId, reservationsIds ): Observable<any> {

		// get the InventoryItem
		this.loadHelper( inventoryItemId );

	// split on a comma with no spaces
	var idList = reservationsIds.split(',')

	// iterate over array of reservations ids
	idList.forEach(function (id) {
		// read the Reservation
		var reservation = new ReservationService(this.http).getReservation(id);
		// add the Reservation if not already assigned
		if ( this.inventoryItem.reservations.indexOf(reservation) == -1 )
		this.inventoryItem.reservations.push(reservation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reservationsIds as a Reservations
	// from a InventoryItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReservations( inventoryItemId, reservationsIds ): Observable<any> {

		// get the InventoryItem
		this.loadHelper( inventoryItemId );


	// split on a comma with no spaces
	var idList 					= reservationsIds.split(',');
	var reservations 	= this.inventoryItem.reservations;

	if ( reservations != null && reservationsIds != null ) {

		// iterate over array of reservations ids
		reservations.forEach(function (obj) {
			if ( reservationsIds.indexOf(obj._id) > -1 ) {
				// remove the Reservation
				this.inventoryItem.reservations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InventoryItem
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InventoryItem/update/' + this.inventoryItem;

	return  this.http.post(uri_, this.inventoryItem );
}

	//********************************************************************
	// loadHelper - internal helper to load a InventoryItem
	//********************************************************************	
	loadHelper( id ) {
		this.getInventoryItem(id)
			.subscribe((res : InventoryItem) => {
				this.inventoryItem = res;
			});
	}
}