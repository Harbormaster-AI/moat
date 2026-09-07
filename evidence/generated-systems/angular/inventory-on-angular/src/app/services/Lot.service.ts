import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Lot} from '../models/Lot';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LotService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	lot : Lot;

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
	// add a Lot
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLot(batchNumber, manufactureDate, expirationDate, Sku, InventoryItems, LotStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/Lot/create';
		const obj = {
			      		batchNumber: batchNumber,
      		manufactureDate: manufactureDate,
      		expirationDate: expirationDate,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			LotStatus: LotStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Lot
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLot(batchNumber, manufactureDate, expirationDate, Sku, InventoryItems, LotStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Lot/update/' + id;
		const obj = {
				      		batchNumber: batchNumber,
      		manufactureDate: manufactureDate,
      		expirationDate: expirationDate,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			LotStatus: LotStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Lot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLot(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Lot/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Lot
	// returns the results untouched as an Observable Lot
	// Lot model
	// delegates via URI
	//********************************************************************
	getLot(id) : Observable<Lot> {
		const uri_ = this.apiUrl + '/Lot/load/' + id;

		return this.http.get<Lot>(uri_);
	}
	
	//********************************************************************
	// gets all Lot
	// returns the results untouched as JSON representation of an
	// Observable array of Lot models
	// delegates via URI
	//********************************************************************
	getLots() : Observable<Lot[]> {
		const uri_ = this.apiUrl + '/Lot/';

		return this
			.http.get<Lot[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Sku on a Lot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( lotId, _skuId ): Observable<any> {

		// get the Lot from storage
		this.loadHelper( lotId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.lot.sku = tmp;

	// save the Lot
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a Lot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( lotId ): Observable<any> {

		// get the Lot from storage
		this.loadHelper( lotId );

	// assign Sku to null
	this.lot.sku = null;

	// save the Lot
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more inventoryItemsIds as a InventoryItems
	// to a Lot
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventoryItems( lotId, inventoryItemsIds ): Observable<any> {

		// get the Lot
		this.loadHelper( lotId );

	// split on a comma with no spaces
	var idList = inventoryItemsIds.split(',')

	// iterate over array of inventoryItems ids
	idList.forEach(function (id) {
		// read the InventoryItem
		var inventoryItem = new InventoryItemService(this.http).getInventoryItem(id);
		// add the InventoryItem if not already assigned
		if ( this.lot.inventoryItems.indexOf(inventoryItem) == -1 )
		this.lot.inventoryItems.push(inventoryItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventoryItemsIds as a InventoryItems
	// from a Lot
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventoryItems( lotId, inventoryItemsIds ): Observable<any> {

		// get the Lot
		this.loadHelper( lotId );


	// split on a comma with no spaces
	var idList 					= inventoryItemsIds.split(',');
	var inventoryItems 	= this.lot.inventoryItems;

	if ( inventoryItems != null && inventoryItemsIds != null ) {

		// iterate over array of inventoryItems ids
		inventoryItems.forEach(function (obj) {
			if ( inventoryItemsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryItem
				this.lot.inventoryItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Lot
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Lot/update/' + this.lot;

	return  this.http.post(uri_, this.lot );
}

	//********************************************************************
	// loadHelper - internal helper to load a Lot
	//********************************************************************	
	loadHelper( id ) {
		this.getLot(id)
			.subscribe((res : Lot) => {
				this.lot = res;
			});
	}
}