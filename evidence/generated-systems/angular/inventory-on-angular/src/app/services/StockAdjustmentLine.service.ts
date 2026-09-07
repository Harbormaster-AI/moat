import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {StockAdjustmentLine} from '../models/StockAdjustmentLine';
import {StockAdjustmentService} from '../services/StockAdjustment.service';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {LotService} from '../services/Lot.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class StockAdjustmentLineService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	stockAdjustmentLine : StockAdjustmentLine;

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
	// add a StockAdjustmentLine
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addStockAdjustmentLine(lineNumber, quantity, Adjustment, Sku, Lot, Location, SerialNumbers, UnitOfMeasure, StockStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/StockAdjustmentLine/create';
		const obj = {
			      		lineNumber: lineNumber,
      		quantity: quantity,
      		Adjustment: Adjustment != null && Adjustment.length > 0 ? Adjustment : null,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		UnitOfMeasure: UnitOfMeasure,
			StockStatus: StockStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a StockAdjustmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateStockAdjustmentLine(lineNumber, quantity, Adjustment, Sku, Lot, Location, SerialNumbers, UnitOfMeasure, StockStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/StockAdjustmentLine/update/' + id;
		const obj = {
				      		lineNumber: lineNumber,
      		quantity: quantity,
      		Adjustment: Adjustment != null && Adjustment.length > 0 ? Adjustment : null,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		UnitOfMeasure: UnitOfMeasure,
			StockStatus: StockStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a StockAdjustmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteStockAdjustmentLine(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/StockAdjustmentLine/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a StockAdjustmentLine
	// returns the results untouched as an Observable StockAdjustmentLine
	// StockAdjustmentLine model
	// delegates via URI
	//********************************************************************
	getStockAdjustmentLine(id) : Observable<StockAdjustmentLine> {
		const uri_ = this.apiUrl + '/StockAdjustmentLine/load/' + id;

		return this.http.get<StockAdjustmentLine>(uri_);
	}
	
	//********************************************************************
	// gets all StockAdjustmentLine
	// returns the results untouched as JSON representation of an
	// Observable array of StockAdjustmentLine models
	// delegates via URI
	//********************************************************************
	getStockAdjustmentLines() : Observable<StockAdjustmentLine[]> {
		const uri_ = this.apiUrl + '/StockAdjustmentLine/';

		return this
			.http.get<StockAdjustmentLine[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Adjustment on a StockAdjustmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdjustment( stockAdjustmentLineId, _adjustmentId ): Observable<any> {

		// get the StockAdjustmentLine from storage
		this.loadHelper( stockAdjustmentLineId );

	// get the StockAdjustment from storage
	var tmp 	= new StockAdjustmentService(this.http).getStockAdjustment(_adjustmentId);

	// assign the Adjustment
	this.stockAdjustmentLine.adjustment = tmp;

	// save the StockAdjustmentLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Adjustment on a StockAdjustmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdjustment( stockAdjustmentLineId ): Observable<any> {

		// get the StockAdjustmentLine from storage
		this.loadHelper( stockAdjustmentLineId );

	// assign Adjustment to null
	this.stockAdjustmentLine.adjustment = null;

	// save the StockAdjustmentLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Sku on a StockAdjustmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( stockAdjustmentLineId, _skuId ): Observable<any> {

		// get the StockAdjustmentLine from storage
		this.loadHelper( stockAdjustmentLineId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.stockAdjustmentLine.sku = tmp;

	// save the StockAdjustmentLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a StockAdjustmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( stockAdjustmentLineId ): Observable<any> {

		// get the StockAdjustmentLine from storage
		this.loadHelper( stockAdjustmentLineId );

	// assign Sku to null
	this.stockAdjustmentLine.sku = null;

	// save the StockAdjustmentLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lot on a StockAdjustmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLot( stockAdjustmentLineId, _lotId ): Observable<any> {

		// get the StockAdjustmentLine from storage
		this.loadHelper( stockAdjustmentLineId );

	// get the Lot from storage
	var tmp 	= new LotService(this.http).getLot(_lotId);

	// assign the Lot
	this.stockAdjustmentLine.lot = tmp;

	// save the StockAdjustmentLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lot on a StockAdjustmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLot( stockAdjustmentLineId ): Observable<any> {

		// get the StockAdjustmentLine from storage
		this.loadHelper( stockAdjustmentLineId );

	// assign Lot to null
	this.stockAdjustmentLine.lot = null;

	// save the StockAdjustmentLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Location on a StockAdjustmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLocation( stockAdjustmentLineId, _locationId ): Observable<any> {

		// get the StockAdjustmentLine from storage
		this.loadHelper( stockAdjustmentLineId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_locationId);

	// assign the Location
	this.stockAdjustmentLine.location = tmp;

	// save the StockAdjustmentLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Location on a StockAdjustmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLocation( stockAdjustmentLineId ): Observable<any> {

		// get the StockAdjustmentLine from storage
		this.loadHelper( stockAdjustmentLineId );

	// assign Location to null
	this.stockAdjustmentLine.location = null;

	// save the StockAdjustmentLine
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more serialNumbersIds as a SerialNumbers
	// to a StockAdjustmentLine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSerialNumbers( stockAdjustmentLineId, serialNumbersIds ): Observable<any> {

		// get the StockAdjustmentLine
		this.loadHelper( stockAdjustmentLineId );

	// split on a comma with no spaces
	var idList = serialNumbersIds.split(',')

	// iterate over array of serialNumbers ids
	idList.forEach(function (id) {
		// read the SerialNumber
		var serialNumber = new SerialNumberService(this.http).getSerialNumber(id);
		// add the SerialNumber if not already assigned
		if ( this.stockAdjustmentLine.serialNumbers.indexOf(serialNumber) == -1 )
		this.stockAdjustmentLine.serialNumbers.push(serialNumber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serialNumbersIds as a SerialNumbers
	// from a StockAdjustmentLine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSerialNumbers( stockAdjustmentLineId, serialNumbersIds ): Observable<any> {

		// get the StockAdjustmentLine
		this.loadHelper( stockAdjustmentLineId );


	// split on a comma with no spaces
	var idList 					= serialNumbersIds.split(',');
	var serialNumbers 	= this.stockAdjustmentLine.serialNumbers;

	if ( serialNumbers != null && serialNumbersIds != null ) {

		// iterate over array of serialNumbers ids
		serialNumbers.forEach(function (obj) {
			if ( serialNumbersIds.indexOf(obj._id) > -1 ) {
				// remove the SerialNumber
				this.stockAdjustmentLine.serialNumbers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a StockAdjustmentLine
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/StockAdjustmentLine/update/' + this.stockAdjustmentLine;

	return  this.http.post(uri_, this.stockAdjustmentLine );
}

	//********************************************************************
	// loadHelper - internal helper to load a StockAdjustmentLine
	//********************************************************************	
	loadHelper( id ) {
		this.getStockAdjustmentLine(id)
			.subscribe((res : StockAdjustmentLine) => {
				this.stockAdjustmentLine = res;
			});
	}
}