import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Quarantine} from '../models/Quarantine';
import {WarehouseService} from '../services/Warehouse.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import {LotService} from '../services/Lot.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class QuarantineService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	quarantine : Quarantine;

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
	// add a Quarantine
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addQuarantine(reason, startedAt, releasedAt, Warehouse, Items, Lot, SerialNumbers, Disposition) : Observable<any> {
		const uri_ = this.apiUrl + '/Quarantine/create';
		const obj = {
			      		reason: reason,
      		startedAt: startedAt,
      		releasedAt: releasedAt,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Items: Items != null && Items.length > 0 ? Items : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
			Disposition: Disposition
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Quarantine
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateQuarantine(reason, startedAt, releasedAt, Warehouse, Items, Lot, SerialNumbers, Disposition, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Quarantine/update/' + id;
		const obj = {
				      		reason: reason,
      		startedAt: startedAt,
      		releasedAt: releasedAt,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Items: Items != null && Items.length > 0 ? Items : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
			Disposition: Disposition
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Quarantine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteQuarantine(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Quarantine/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Quarantine
	// returns the results untouched as an Observable Quarantine
	// Quarantine model
	// delegates via URI
	//********************************************************************
	getQuarantine(id) : Observable<Quarantine> {
		const uri_ = this.apiUrl + '/Quarantine/load/' + id;

		return this.http.get<Quarantine>(uri_);
	}
	
	//********************************************************************
	// gets all Quarantine
	// returns the results untouched as JSON representation of an
	// Observable array of Quarantine models
	// delegates via URI
	//********************************************************************
	getQuarantines() : Observable<Quarantine[]> {
		const uri_ = this.apiUrl + '/Quarantine/';

		return this
			.http.get<Quarantine[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Warehouse on a Quarantine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( quarantineId, _warehouseId ): Observable<any> {

		// get the Quarantine from storage
		this.loadHelper( quarantineId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.quarantine.warehouse = tmp;

	// save the Quarantine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a Quarantine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( quarantineId ): Observable<any> {

		// get the Quarantine from storage
		this.loadHelper( quarantineId );

	// assign Warehouse to null
	this.quarantine.warehouse = null;

	// save the Quarantine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lot on a Quarantine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLot( quarantineId, _lotId ): Observable<any> {

		// get the Quarantine from storage
		this.loadHelper( quarantineId );

	// get the Lot from storage
	var tmp 	= new LotService(this.http).getLot(_lotId);

	// assign the Lot
	this.quarantine.lot = tmp;

	// save the Quarantine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lot on a Quarantine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLot( quarantineId ): Observable<any> {

		// get the Quarantine from storage
		this.loadHelper( quarantineId );

	// assign Lot to null
	this.quarantine.lot = null;

	// save the Quarantine
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more itemsIds as a Items
	// to a Quarantine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addItems( quarantineId, itemsIds ): Observable<any> {

		// get the Quarantine
		this.loadHelper( quarantineId );

	// split on a comma with no spaces
	var idList = itemsIds.split(',')

	// iterate over array of items ids
	idList.forEach(function (id) {
		// read the InventoryItem
		var inventoryItem = new InventoryItemService(this.http).getInventoryItem(id);
		// add the InventoryItem if not already assigned
		if ( this.quarantine.items.indexOf(inventoryItem) == -1 )
		this.quarantine.items.push(inventoryItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more itemsIds as a Items
	// from a Quarantine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeItems( quarantineId, itemsIds ): Observable<any> {

		// get the Quarantine
		this.loadHelper( quarantineId );


	// split on a comma with no spaces
	var idList 					= itemsIds.split(',');
	var items 	= this.quarantine.items;

	if ( items != null && itemsIds != null ) {

		// iterate over array of items ids
		items.forEach(function (obj) {
			if ( itemsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryItem
				this.quarantine.items.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more serialNumbersIds as a SerialNumbers
	// to a Quarantine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSerialNumbers( quarantineId, serialNumbersIds ): Observable<any> {

		// get the Quarantine
		this.loadHelper( quarantineId );

	// split on a comma with no spaces
	var idList = serialNumbersIds.split(',')

	// iterate over array of serialNumbers ids
	idList.forEach(function (id) {
		// read the SerialNumber
		var serialNumber = new SerialNumberService(this.http).getSerialNumber(id);
		// add the SerialNumber if not already assigned
		if ( this.quarantine.serialNumbers.indexOf(serialNumber) == -1 )
		this.quarantine.serialNumbers.push(serialNumber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serialNumbersIds as a SerialNumbers
	// from a Quarantine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSerialNumbers( quarantineId, serialNumbersIds ): Observable<any> {

		// get the Quarantine
		this.loadHelper( quarantineId );


	// split on a comma with no spaces
	var idList 					= serialNumbersIds.split(',');
	var serialNumbers 	= this.quarantine.serialNumbers;

	if ( serialNumbers != null && serialNumbersIds != null ) {

		// iterate over array of serialNumbers ids
		serialNumbers.forEach(function (obj) {
			if ( serialNumbersIds.indexOf(obj._id) > -1 ) {
				// remove the SerialNumber
				this.quarantine.serialNumbers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Quarantine
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Quarantine/update/' + this.quarantine;

	return  this.http.post(uri_, this.quarantine );
}

	//********************************************************************
	// loadHelper - internal helper to load a Quarantine
	//********************************************************************	
	loadHelper( id ) {
		this.getQuarantine(id)
			.subscribe((res : Quarantine) => {
				this.quarantine = res;
			});
	}
}