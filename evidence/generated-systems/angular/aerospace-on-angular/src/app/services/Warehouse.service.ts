import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Warehouse} from '../models/Warehouse';
import {InventoryItemService} from '../services/InventoryItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class WarehouseService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	warehouse : Warehouse;

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
	// add a Warehouse
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addWarehouse(name, InventoryItems) : Observable<any> {
		const uri_ = this.apiUrl + '/Warehouse/create';
		const obj = {
			      		name: name,
			InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Warehouse
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateWarehouse(name, InventoryItems, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Warehouse/update/' + id;
		const obj = {
				      		name: name,
			InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Warehouse
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteWarehouse(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Warehouse/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Warehouse
	// returns the results untouched as an Observable Warehouse
	// Warehouse model
	// delegates via URI
	//********************************************************************
	getWarehouse(id) : Observable<Warehouse> {
		const uri_ = this.apiUrl + '/Warehouse/load/' + id;

		return this.http.get<Warehouse>(uri_);
	}
	
	//********************************************************************
	// gets all Warehouse
	// returns the results untouched as JSON representation of an
	// Observable array of Warehouse models
	// delegates via URI
	//********************************************************************
	getWarehouses() : Observable<Warehouse[]> {
		const uri_ = this.apiUrl + '/Warehouse/';

		return this
			.http.get<Warehouse[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more inventoryItemsIds as a InventoryItems
	// to a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventoryItems( warehouseId, inventoryItemsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );

	// split on a comma with no spaces
	var idList = inventoryItemsIds.split(',')

	// iterate over array of inventoryItems ids
	idList.forEach(function (id) {
		// read the InventoryItem
		var inventoryItem = new InventoryItemService(this.http).getInventoryItem(id);
		// add the InventoryItem if not already assigned
		if ( this.warehouse.inventoryItems.indexOf(inventoryItem) == -1 )
		this.warehouse.inventoryItems.push(inventoryItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventoryItemsIds as a InventoryItems
	// from a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventoryItems( warehouseId, inventoryItemsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );


	// split on a comma with no spaces
	var idList 					= inventoryItemsIds.split(',');
	var inventoryItems 	= this.warehouse.inventoryItems;

	if ( inventoryItems != null && inventoryItemsIds != null ) {

		// iterate over array of inventoryItems ids
		inventoryItems.forEach(function (obj) {
			if ( inventoryItemsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryItem
				this.warehouse.inventoryItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Warehouse
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Warehouse/update/' + this.warehouse;

	return  this.http.post(uri_, this.warehouse );
}

	//********************************************************************
	// loadHelper - internal helper to load a Warehouse
	//********************************************************************	
	loadHelper( id ) {
		this.getWarehouse(id)
			.subscribe((res : Warehouse) => {
				this.warehouse = res;
			});
	}
}