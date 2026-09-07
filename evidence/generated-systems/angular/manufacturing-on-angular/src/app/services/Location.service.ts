import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Location} from '../models/Location';
import {WarehouseService} from '../services/Warehouse.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LocationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	location : Location;

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
	// add a Location
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLocation(locationCode, description, Warehouse, InventoryItems, LocationType) : Observable<any> {
		const uri_ = this.apiUrl + '/Location/create';
		const obj = {
			      		locationCode: locationCode,
      		description: description,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			LocationType: LocationType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Location
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLocation(locationCode, description, Warehouse, InventoryItems, LocationType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Location/update/' + id;
		const obj = {
				      		locationCode: locationCode,
      		description: description,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			LocationType: LocationType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Location
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLocation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Location/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Location
	// returns the results untouched as an Observable Location
	// Location model
	// delegates via URI
	//********************************************************************
	getLocation(id) : Observable<Location> {
		const uri_ = this.apiUrl + '/Location/load/' + id;

		return this.http.get<Location>(uri_);
	}
	
	//********************************************************************
	// gets all Location
	// returns the results untouched as JSON representation of an
	// Observable array of Location models
	// delegates via URI
	//********************************************************************
	getLocations() : Observable<Location[]> {
		const uri_ = this.apiUrl + '/Location/';

		return this
			.http.get<Location[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Warehouse on a Location
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( locationId, _warehouseId ): Observable<any> {

		// get the Location from storage
		this.loadHelper( locationId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.location.warehouse = tmp;

	// save the Location
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a Location
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( locationId ): Observable<any> {

		// get the Location from storage
		this.loadHelper( locationId );

	// assign Warehouse to null
	this.location.warehouse = null;

	// save the Location
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more inventoryItemsIds as a InventoryItems
	// to a Location
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventoryItems( locationId, inventoryItemsIds ): Observable<any> {

		// get the Location
		this.loadHelper( locationId );

	// split on a comma with no spaces
	var idList = inventoryItemsIds.split(',')

	// iterate over array of inventoryItems ids
	idList.forEach(function (id) {
		// read the InventoryItem
		var inventoryItem = new InventoryItemService(this.http).getInventoryItem(id);
		// add the InventoryItem if not already assigned
		if ( this.location.inventoryItems.indexOf(inventoryItem) == -1 )
		this.location.inventoryItems.push(inventoryItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventoryItemsIds as a InventoryItems
	// from a Location
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventoryItems( locationId, inventoryItemsIds ): Observable<any> {

		// get the Location
		this.loadHelper( locationId );


	// split on a comma with no spaces
	var idList 					= inventoryItemsIds.split(',');
	var inventoryItems 	= this.location.inventoryItems;

	if ( inventoryItems != null && inventoryItemsIds != null ) {

		// iterate over array of inventoryItems ids
		inventoryItems.forEach(function (obj) {
			if ( inventoryItemsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryItem
				this.location.inventoryItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Location
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Location/update/' + this.location;

	return  this.http.post(uri_, this.location );
}

	//********************************************************************
	// loadHelper - internal helper to load a Location
	//********************************************************************	
	loadHelper( id ) {
		this.getLocation(id)
			.subscribe((res : Location) => {
				this.location = res;
			});
	}
}