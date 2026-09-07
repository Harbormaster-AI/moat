import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InventoryItem} from '../models/InventoryItem';
import {ItemService} from '../services/Item.service';
import {LocationService} from '../services/Location.service';
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
	addInventoryItem(quantityOnHand, quantityReserved, lotNumber, serialNumber, Item, Location) : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryItem/create';
		const obj = {
			      		quantityOnHand: quantityOnHand,
      		quantityReserved: quantityReserved,
      		lotNumber: lotNumber,
      		serialNumber: serialNumber,
      		Item: Item != null && Item.length > 0 ? Item : null,
			Location: Location != null && Location.length > 0 ? Location : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInventoryItem(quantityOnHand, quantityReserved, lotNumber, serialNumber, Item, Location, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InventoryItem/update/' + id;
		const obj = {
				      		quantityOnHand: quantityOnHand,
      		quantityReserved: quantityReserved,
      		lotNumber: lotNumber,
      		serialNumber: serialNumber,
      		Item: Item != null && Item.length > 0 ? Item : null,
			Location: Location != null && Location.length > 0 ? Location : null
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
	// assigns a Item on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( inventoryItemId, _itemId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.inventoryItem.item = tmp;

	// save the InventoryItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( inventoryItemId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// assign Item to null
	this.inventoryItem.item = null;

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

	// get the Location from storage
	var tmp 	= new LocationService(this.http).getLocation(_locationId);

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