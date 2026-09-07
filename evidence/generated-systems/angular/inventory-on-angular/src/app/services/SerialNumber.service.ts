import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SerialNumber} from '../models/SerialNumber';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import {LotService} from '../services/Lot.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SerialNumberService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	serialNumber : SerialNumber;

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
	// add a SerialNumber
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSerialNumber(serial, activationDate, Sku, CurrentInventoryItem, Lot, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/SerialNumber/create';
		const obj = {
			      		serial: serial,
      		activationDate: activationDate,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		CurrentInventoryItem: CurrentInventoryItem != null && CurrentInventoryItem.length > 0 ? CurrentInventoryItem : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SerialNumber
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSerialNumber(serial, activationDate, Sku, CurrentInventoryItem, Lot, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SerialNumber/update/' + id;
		const obj = {
				      		serial: serial,
      		activationDate: activationDate,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		CurrentInventoryItem: CurrentInventoryItem != null && CurrentInventoryItem.length > 0 ? CurrentInventoryItem : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SerialNumber
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSerialNumber(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SerialNumber/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SerialNumber
	// returns the results untouched as an Observable SerialNumber
	// SerialNumber model
	// delegates via URI
	//********************************************************************
	getSerialNumber(id) : Observable<SerialNumber> {
		const uri_ = this.apiUrl + '/SerialNumber/load/' + id;

		return this.http.get<SerialNumber>(uri_);
	}
	
	//********************************************************************
	// gets all SerialNumber
	// returns the results untouched as JSON representation of an
	// Observable array of SerialNumber models
	// delegates via URI
	//********************************************************************
	getSerialNumbers() : Observable<SerialNumber[]> {
		const uri_ = this.apiUrl + '/SerialNumber/';

		return this
			.http.get<SerialNumber[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Sku on a SerialNumber
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( serialNumberId, _skuId ): Observable<any> {

		// get the SerialNumber from storage
		this.loadHelper( serialNumberId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.serialNumber.sku = tmp;

	// save the SerialNumber
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a SerialNumber
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( serialNumberId ): Observable<any> {

		// get the SerialNumber from storage
		this.loadHelper( serialNumberId );

	// assign Sku to null
	this.serialNumber.sku = null;

	// save the SerialNumber
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CurrentInventoryItem on a SerialNumber
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCurrentInventoryItem( serialNumberId, _currentInventoryItemId ): Observable<any> {

		// get the SerialNumber from storage
		this.loadHelper( serialNumberId );

	// get the InventoryItem from storage
	var tmp 	= new InventoryItemService(this.http).getInventoryItem(_currentInventoryItemId);

	// assign the CurrentInventoryItem
	this.serialNumber.currentInventoryItem = tmp;

	// save the SerialNumber
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CurrentInventoryItem on a SerialNumber
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCurrentInventoryItem( serialNumberId ): Observable<any> {

		// get the SerialNumber from storage
		this.loadHelper( serialNumberId );

	// assign CurrentInventoryItem to null
	this.serialNumber.currentInventoryItem = null;

	// save the SerialNumber
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lot on a SerialNumber
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLot( serialNumberId, _lotId ): Observable<any> {

		// get the SerialNumber from storage
		this.loadHelper( serialNumberId );

	// get the Lot from storage
	var tmp 	= new LotService(this.http).getLot(_lotId);

	// assign the Lot
	this.serialNumber.lot = tmp;

	// save the SerialNumber
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lot on a SerialNumber
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLot( serialNumberId ): Observable<any> {

		// get the SerialNumber from storage
		this.loadHelper( serialNumberId );

	// assign Lot to null
	this.serialNumber.lot = null;

	// save the SerialNumber
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a SerialNumber
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SerialNumber/update/' + this.serialNumber;

	return  this.http.post(uri_, this.serialNumber );
}

	//********************************************************************
	// loadHelper - internal helper to load a SerialNumber
	//********************************************************************	
	loadHelper( id ) {
		this.getSerialNumber(id)
			.subscribe((res : SerialNumber) => {
				this.serialNumber = res;
			});
	}
}