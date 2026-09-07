import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InventoryItem} from '../models/InventoryItem';
import {Component_Service} from '../services/Component_.service';
import {WarehouseService} from '../services/Warehouse.service';
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
	addInventoryItem(quantityOnHand, quantityReserved, lotNumber, Component, Warehouse) : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryItem/create';
		const obj = {
			      		quantityOnHand: quantityOnHand,
      		quantityReserved: quantityReserved,
      		lotNumber: lotNumber,
      		Component: Component != null && Component.length > 0 ? Component : null,
			Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInventoryItem(quantityOnHand, quantityReserved, lotNumber, Component, Warehouse, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InventoryItem/update/' + id;
		const obj = {
				      		quantityOnHand: quantityOnHand,
      		quantityReserved: quantityReserved,
      		lotNumber: lotNumber,
      		Component: Component != null && Component.length > 0 ? Component : null,
			Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null
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
	// assigns a Component on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignComponent( inventoryItemId, _componentId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// get the Component_ from storage
	var tmp 	= new Component_Service(this.http).getComponent_(_componentId);

	// assign the Component
	this.inventoryItem.component = tmp;

	// save the InventoryItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Component on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignComponent( inventoryItemId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// assign Component to null
	this.inventoryItem.component = null;

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