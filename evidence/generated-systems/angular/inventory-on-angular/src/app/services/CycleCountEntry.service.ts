import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CycleCountEntry} from '../models/CycleCountEntry';
import {CycleCountService} from '../services/CycleCount.service';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {LotService} from '../services/Lot.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CycleCountEntryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	cycleCountEntry : CycleCountEntry;

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
	// add a CycleCountEntry
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCycleCountEntry(lineNumber, systemQuantity, countedQuantity, varianceQuantity, recountRequired, CycleCount, Sku, Lot, Location, SerialNumbers, StockStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/CycleCountEntry/create';
		const obj = {
			      		lineNumber: lineNumber,
      		systemQuantity: systemQuantity,
      		countedQuantity: countedQuantity,
      		varianceQuantity: varianceQuantity,
      		recountRequired: recountRequired,
      		CycleCount: CycleCount != null && CycleCount.length > 0 ? CycleCount : null,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
			StockStatus: StockStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CycleCountEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCycleCountEntry(lineNumber, systemQuantity, countedQuantity, varianceQuantity, recountRequired, CycleCount, Sku, Lot, Location, SerialNumbers, StockStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CycleCountEntry/update/' + id;
		const obj = {
				      		lineNumber: lineNumber,
      		systemQuantity: systemQuantity,
      		countedQuantity: countedQuantity,
      		varianceQuantity: varianceQuantity,
      		recountRequired: recountRequired,
      		CycleCount: CycleCount != null && CycleCount.length > 0 ? CycleCount : null,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
			StockStatus: StockStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CycleCountEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCycleCountEntry(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CycleCountEntry/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CycleCountEntry
	// returns the results untouched as an Observable CycleCountEntry
	// CycleCountEntry model
	// delegates via URI
	//********************************************************************
	getCycleCountEntry(id) : Observable<CycleCountEntry> {
		const uri_ = this.apiUrl + '/CycleCountEntry/load/' + id;

		return this.http.get<CycleCountEntry>(uri_);
	}
	
	//********************************************************************
	// gets all CycleCountEntry
	// returns the results untouched as JSON representation of an
	// Observable array of CycleCountEntry models
	// delegates via URI
	//********************************************************************
	getCycleCountEntrys() : Observable<CycleCountEntry[]> {
		const uri_ = this.apiUrl + '/CycleCountEntry/';

		return this
			.http.get<CycleCountEntry[]>(uri_);
	}
	
			//********************************************************************
	// assigns a CycleCount on a CycleCountEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCycleCount( cycleCountEntryId, _cycleCountId ): Observable<any> {

		// get the CycleCountEntry from storage
		this.loadHelper( cycleCountEntryId );

	// get the CycleCount from storage
	var tmp 	= new CycleCountService(this.http).getCycleCount(_cycleCountId);

	// assign the CycleCount
	this.cycleCountEntry.cycleCount = tmp;

	// save the CycleCountEntry
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CycleCount on a CycleCountEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCycleCount( cycleCountEntryId ): Observable<any> {

		// get the CycleCountEntry from storage
		this.loadHelper( cycleCountEntryId );

	// assign CycleCount to null
	this.cycleCountEntry.cycleCount = null;

	// save the CycleCountEntry
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Sku on a CycleCountEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( cycleCountEntryId, _skuId ): Observable<any> {

		// get the CycleCountEntry from storage
		this.loadHelper( cycleCountEntryId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.cycleCountEntry.sku = tmp;

	// save the CycleCountEntry
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a CycleCountEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( cycleCountEntryId ): Observable<any> {

		// get the CycleCountEntry from storage
		this.loadHelper( cycleCountEntryId );

	// assign Sku to null
	this.cycleCountEntry.sku = null;

	// save the CycleCountEntry
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lot on a CycleCountEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLot( cycleCountEntryId, _lotId ): Observable<any> {

		// get the CycleCountEntry from storage
		this.loadHelper( cycleCountEntryId );

	// get the Lot from storage
	var tmp 	= new LotService(this.http).getLot(_lotId);

	// assign the Lot
	this.cycleCountEntry.lot = tmp;

	// save the CycleCountEntry
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lot on a CycleCountEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLot( cycleCountEntryId ): Observable<any> {

		// get the CycleCountEntry from storage
		this.loadHelper( cycleCountEntryId );

	// assign Lot to null
	this.cycleCountEntry.lot = null;

	// save the CycleCountEntry
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Location on a CycleCountEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLocation( cycleCountEntryId, _locationId ): Observable<any> {

		// get the CycleCountEntry from storage
		this.loadHelper( cycleCountEntryId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_locationId);

	// assign the Location
	this.cycleCountEntry.location = tmp;

	// save the CycleCountEntry
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Location on a CycleCountEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLocation( cycleCountEntryId ): Observable<any> {

		// get the CycleCountEntry from storage
		this.loadHelper( cycleCountEntryId );

	// assign Location to null
	this.cycleCountEntry.location = null;

	// save the CycleCountEntry
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more serialNumbersIds as a SerialNumbers
	// to a CycleCountEntry
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSerialNumbers( cycleCountEntryId, serialNumbersIds ): Observable<any> {

		// get the CycleCountEntry
		this.loadHelper( cycleCountEntryId );

	// split on a comma with no spaces
	var idList = serialNumbersIds.split(',')

	// iterate over array of serialNumbers ids
	idList.forEach(function (id) {
		// read the SerialNumber
		var serialNumber = new SerialNumberService(this.http).getSerialNumber(id);
		// add the SerialNumber if not already assigned
		if ( this.cycleCountEntry.serialNumbers.indexOf(serialNumber) == -1 )
		this.cycleCountEntry.serialNumbers.push(serialNumber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serialNumbersIds as a SerialNumbers
	// from a CycleCountEntry
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSerialNumbers( cycleCountEntryId, serialNumbersIds ): Observable<any> {

		// get the CycleCountEntry
		this.loadHelper( cycleCountEntryId );


	// split on a comma with no spaces
	var idList 					= serialNumbersIds.split(',');
	var serialNumbers 	= this.cycleCountEntry.serialNumbers;

	if ( serialNumbers != null && serialNumbersIds != null ) {

		// iterate over array of serialNumbers ids
		serialNumbers.forEach(function (obj) {
			if ( serialNumbersIds.indexOf(obj._id) > -1 ) {
				// remove the SerialNumber
				this.cycleCountEntry.serialNumbers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a CycleCountEntry
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CycleCountEntry/update/' + this.cycleCountEntry;

	return  this.http.post(uri_, this.cycleCountEntry );
}

	//********************************************************************
	// loadHelper - internal helper to load a CycleCountEntry
	//********************************************************************	
	loadHelper( id ) {
		this.getCycleCountEntry(id)
			.subscribe((res : CycleCountEntry) => {
				this.cycleCountEntry = res;
			});
	}
}