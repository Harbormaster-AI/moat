import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {StorageLocation} from '../models/StorageLocation';
import {WarehouseService} from '../services/Warehouse.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class StorageLocationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	storageLocation : StorageLocation;

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
	// add a StorageLocation
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addStorageLocation(code, temperatureControlled, capacity, capacityUnit, Warehouse, ParentLocation, ChildLocations, InventoryItems, LocationType) : Observable<any> {
		const uri_ = this.apiUrl + '/StorageLocation/create';
		const obj = {
			      		code: code,
      		temperatureControlled: temperatureControlled,
      		capacity: capacity,
      		capacityUnit: capacityUnit,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		ParentLocation: ParentLocation != null && ParentLocation.length > 0 ? ParentLocation : null,
      		ChildLocations: ChildLocations != null && ChildLocations.length > 0 ? ChildLocations : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			LocationType: LocationType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a StorageLocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateStorageLocation(code, temperatureControlled, capacity, capacityUnit, Warehouse, ParentLocation, ChildLocations, InventoryItems, LocationType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/StorageLocation/update/' + id;
		const obj = {
				      		code: code,
      		temperatureControlled: temperatureControlled,
      		capacity: capacity,
      		capacityUnit: capacityUnit,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		ParentLocation: ParentLocation != null && ParentLocation.length > 0 ? ParentLocation : null,
      		ChildLocations: ChildLocations != null && ChildLocations.length > 0 ? ChildLocations : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			LocationType: LocationType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a StorageLocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteStorageLocation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/StorageLocation/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a StorageLocation
	// returns the results untouched as an Observable StorageLocation
	// StorageLocation model
	// delegates via URI
	//********************************************************************
	getStorageLocation(id) : Observable<StorageLocation> {
		const uri_ = this.apiUrl + '/StorageLocation/load/' + id;

		return this.http.get<StorageLocation>(uri_);
	}
	
	//********************************************************************
	// gets all StorageLocation
	// returns the results untouched as JSON representation of an
	// Observable array of StorageLocation models
	// delegates via URI
	//********************************************************************
	getStorageLocations() : Observable<StorageLocation[]> {
		const uri_ = this.apiUrl + '/StorageLocation/';

		return this
			.http.get<StorageLocation[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Warehouse on a StorageLocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( storageLocationId, _warehouseId ): Observable<any> {

		// get the StorageLocation from storage
		this.loadHelper( storageLocationId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.storageLocation.warehouse = tmp;

	// save the StorageLocation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a StorageLocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( storageLocationId ): Observable<any> {

		// get the StorageLocation from storage
		this.loadHelper( storageLocationId );

	// assign Warehouse to null
	this.storageLocation.warehouse = null;

	// save the StorageLocation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ParentLocation on a StorageLocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignParentLocation( storageLocationId, _parentLocationId ): Observable<any> {

		// get the StorageLocation from storage
		this.loadHelper( storageLocationId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_parentLocationId);

	// assign the ParentLocation
	this.storageLocation.parentLocation = tmp;

	// save the StorageLocation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ParentLocation on a StorageLocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignParentLocation( storageLocationId ): Observable<any> {

		// get the StorageLocation from storage
		this.loadHelper( storageLocationId );

	// assign ParentLocation to null
	this.storageLocation.parentLocation = null;

	// save the StorageLocation
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more childLocationsIds as a ChildLocations
	// to a StorageLocation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addChildLocations( storageLocationId, childLocationsIds ): Observable<any> {

		// get the StorageLocation
		this.loadHelper( storageLocationId );

	// split on a comma with no spaces
	var idList = childLocationsIds.split(',')

	// iterate over array of childLocations ids
	idList.forEach(function (id) {
		// read the StorageLocation
		var storageLocation = new StorageLocationService(this.http).getStorageLocation(id);
		// add the StorageLocation if not already assigned
		if ( this.storageLocation.childLocations.indexOf(storageLocation) == -1 )
		this.storageLocation.childLocations.push(storageLocation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more childLocationsIds as a ChildLocations
	// from a StorageLocation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeChildLocations( storageLocationId, childLocationsIds ): Observable<any> {

		// get the StorageLocation
		this.loadHelper( storageLocationId );


	// split on a comma with no spaces
	var idList 					= childLocationsIds.split(',');
	var childLocations 	= this.storageLocation.childLocations;

	if ( childLocations != null && childLocationsIds != null ) {

		// iterate over array of childLocations ids
		childLocations.forEach(function (obj) {
			if ( childLocationsIds.indexOf(obj._id) > -1 ) {
				// remove the StorageLocation
				this.storageLocation.childLocations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more inventoryItemsIds as a InventoryItems
	// to a StorageLocation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventoryItems( storageLocationId, inventoryItemsIds ): Observable<any> {

		// get the StorageLocation
		this.loadHelper( storageLocationId );

	// split on a comma with no spaces
	var idList = inventoryItemsIds.split(',')

	// iterate over array of inventoryItems ids
	idList.forEach(function (id) {
		// read the InventoryItem
		var inventoryItem = new InventoryItemService(this.http).getInventoryItem(id);
		// add the InventoryItem if not already assigned
		if ( this.storageLocation.inventoryItems.indexOf(inventoryItem) == -1 )
		this.storageLocation.inventoryItems.push(inventoryItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventoryItemsIds as a InventoryItems
	// from a StorageLocation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventoryItems( storageLocationId, inventoryItemsIds ): Observable<any> {

		// get the StorageLocation
		this.loadHelper( storageLocationId );


	// split on a comma with no spaces
	var idList 					= inventoryItemsIds.split(',');
	var inventoryItems 	= this.storageLocation.inventoryItems;

	if ( inventoryItems != null && inventoryItemsIds != null ) {

		// iterate over array of inventoryItems ids
		inventoryItems.forEach(function (obj) {
			if ( inventoryItemsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryItem
				this.storageLocation.inventoryItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a StorageLocation
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/StorageLocation/update/' + this.storageLocation;

	return  this.http.post(uri_, this.storageLocation );
}

	//********************************************************************
	// loadHelper - internal helper to load a StorageLocation
	//********************************************************************	
	loadHelper( id ) {
		this.getStorageLocation(id)
			.subscribe((res : StorageLocation) => {
				this.storageLocation = res;
			});
	}
}