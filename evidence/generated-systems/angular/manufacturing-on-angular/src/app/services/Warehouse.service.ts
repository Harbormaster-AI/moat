import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Warehouse} from '../models/Warehouse';
import {PlantService} from '../services/Plant.service';
import {LocationService} from '../services/Location.service';
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
	addWarehouse(name, warehouseCode, address, Plant, Locations, InventoryItems, WarehouseType) : Observable<any> {
		const uri_ = this.apiUrl + '/Warehouse/create';
		const obj = {
			      		name: name,
      		warehouseCode: warehouseCode,
      		address: address,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		Locations: Locations != null && Locations.length > 0 ? Locations : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			WarehouseType: WarehouseType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Warehouse
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateWarehouse(name, warehouseCode, address, Plant, Locations, InventoryItems, WarehouseType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Warehouse/update/' + id;
		const obj = {
				      		name: name,
      		warehouseCode: warehouseCode,
      		address: address,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		Locations: Locations != null && Locations.length > 0 ? Locations : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			WarehouseType: WarehouseType
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
	// assigns a Plant on a Warehouse
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( warehouseId, _plantId ): Observable<any> {

		// get the Warehouse from storage
		this.loadHelper( warehouseId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.warehouse.plant = tmp;

	// save the Warehouse
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a Warehouse
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( warehouseId ): Observable<any> {

		// get the Warehouse from storage
		this.loadHelper( warehouseId );

	// assign Plant to null
	this.warehouse.plant = null;

	// save the Warehouse
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more locationsIds as a Locations
	// to a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLocations( warehouseId, locationsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );

	// split on a comma with no spaces
	var idList = locationsIds.split(',')

	// iterate over array of locations ids
	idList.forEach(function (id) {
		// read the Location
		var location = new LocationService(this.http).getLocation(id);
		// add the Location if not already assigned
		if ( this.warehouse.locations.indexOf(location) == -1 )
		this.warehouse.locations.push(location);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more locationsIds as a Locations
	// from a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLocations( warehouseId, locationsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );


	// split on a comma with no spaces
	var idList 					= locationsIds.split(',');
	var locations 	= this.warehouse.locations;

	if ( locations != null && locationsIds != null ) {

		// iterate over array of locations ids
		locations.forEach(function (obj) {
			if ( locationsIds.indexOf(obj._id) > -1 ) {
				// remove the Location
				this.warehouse.locations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
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