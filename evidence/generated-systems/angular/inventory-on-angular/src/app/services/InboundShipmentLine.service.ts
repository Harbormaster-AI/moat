import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InboundShipmentLine} from '../models/InboundShipmentLine';
import {InboundShipmentService} from '../services/InboundShipment.service';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {LotService} from '../services/Lot.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InboundShipmentLineService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inboundShipmentLine : InboundShipmentLine;

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
	// add a InboundShipmentLine
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInboundShipmentLine(lineNumber, quantity, InboundShipment, Sku, Lot, SerialNumbers, DestinationLocation, UnitOfMeasure, StockStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/InboundShipmentLine/create';
		const obj = {
			      		lineNumber: lineNumber,
      		quantity: quantity,
      		InboundShipment: InboundShipment != null && InboundShipment.length > 0 ? InboundShipment : null,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		DestinationLocation: DestinationLocation != null && DestinationLocation.length > 0 ? DestinationLocation : null,
      		UnitOfMeasure: UnitOfMeasure,
			StockStatus: StockStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InboundShipmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInboundShipmentLine(lineNumber, quantity, InboundShipment, Sku, Lot, SerialNumbers, DestinationLocation, UnitOfMeasure, StockStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InboundShipmentLine/update/' + id;
		const obj = {
				      		lineNumber: lineNumber,
      		quantity: quantity,
      		InboundShipment: InboundShipment != null && InboundShipment.length > 0 ? InboundShipment : null,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		DestinationLocation: DestinationLocation != null && DestinationLocation.length > 0 ? DestinationLocation : null,
      		UnitOfMeasure: UnitOfMeasure,
			StockStatus: StockStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InboundShipmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInboundShipmentLine(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InboundShipmentLine/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InboundShipmentLine
	// returns the results untouched as an Observable InboundShipmentLine
	// InboundShipmentLine model
	// delegates via URI
	//********************************************************************
	getInboundShipmentLine(id) : Observable<InboundShipmentLine> {
		const uri_ = this.apiUrl + '/InboundShipmentLine/load/' + id;

		return this.http.get<InboundShipmentLine>(uri_);
	}
	
	//********************************************************************
	// gets all InboundShipmentLine
	// returns the results untouched as JSON representation of an
	// Observable array of InboundShipmentLine models
	// delegates via URI
	//********************************************************************
	getInboundShipmentLines() : Observable<InboundShipmentLine[]> {
		const uri_ = this.apiUrl + '/InboundShipmentLine/';

		return this
			.http.get<InboundShipmentLine[]>(uri_);
	}
	
			//********************************************************************
	// assigns a InboundShipment on a InboundShipmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInboundShipment( inboundShipmentLineId, _inboundShipmentId ): Observable<any> {

		// get the InboundShipmentLine from storage
		this.loadHelper( inboundShipmentLineId );

	// get the InboundShipment from storage
	var tmp 	= new InboundShipmentService(this.http).getInboundShipment(_inboundShipmentId);

	// assign the InboundShipment
	this.inboundShipmentLine.inboundShipment = tmp;

	// save the InboundShipmentLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InboundShipment on a InboundShipmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInboundShipment( inboundShipmentLineId ): Observable<any> {

		// get the InboundShipmentLine from storage
		this.loadHelper( inboundShipmentLineId );

	// assign InboundShipment to null
	this.inboundShipmentLine.inboundShipment = null;

	// save the InboundShipmentLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Sku on a InboundShipmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( inboundShipmentLineId, _skuId ): Observable<any> {

		// get the InboundShipmentLine from storage
		this.loadHelper( inboundShipmentLineId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.inboundShipmentLine.sku = tmp;

	// save the InboundShipmentLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a InboundShipmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( inboundShipmentLineId ): Observable<any> {

		// get the InboundShipmentLine from storage
		this.loadHelper( inboundShipmentLineId );

	// assign Sku to null
	this.inboundShipmentLine.sku = null;

	// save the InboundShipmentLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lot on a InboundShipmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLot( inboundShipmentLineId, _lotId ): Observable<any> {

		// get the InboundShipmentLine from storage
		this.loadHelper( inboundShipmentLineId );

	// get the Lot from storage
	var tmp 	= new LotService(this.http).getLot(_lotId);

	// assign the Lot
	this.inboundShipmentLine.lot = tmp;

	// save the InboundShipmentLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lot on a InboundShipmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLot( inboundShipmentLineId ): Observable<any> {

		// get the InboundShipmentLine from storage
		this.loadHelper( inboundShipmentLineId );

	// assign Lot to null
	this.inboundShipmentLine.lot = null;

	// save the InboundShipmentLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a DestinationLocation on a InboundShipmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDestinationLocation( inboundShipmentLineId, _destinationLocationId ): Observable<any> {

		// get the InboundShipmentLine from storage
		this.loadHelper( inboundShipmentLineId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_destinationLocationId);

	// assign the DestinationLocation
	this.inboundShipmentLine.destinationLocation = tmp;

	// save the InboundShipmentLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a DestinationLocation on a InboundShipmentLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDestinationLocation( inboundShipmentLineId ): Observable<any> {

		// get the InboundShipmentLine from storage
		this.loadHelper( inboundShipmentLineId );

	// assign DestinationLocation to null
	this.inboundShipmentLine.destinationLocation = null;

	// save the InboundShipmentLine
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more serialNumbersIds as a SerialNumbers
	// to a InboundShipmentLine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSerialNumbers( inboundShipmentLineId, serialNumbersIds ): Observable<any> {

		// get the InboundShipmentLine
		this.loadHelper( inboundShipmentLineId );

	// split on a comma with no spaces
	var idList = serialNumbersIds.split(',')

	// iterate over array of serialNumbers ids
	idList.forEach(function (id) {
		// read the SerialNumber
		var serialNumber = new SerialNumberService(this.http).getSerialNumber(id);
		// add the SerialNumber if not already assigned
		if ( this.inboundShipmentLine.serialNumbers.indexOf(serialNumber) == -1 )
		this.inboundShipmentLine.serialNumbers.push(serialNumber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serialNumbersIds as a SerialNumbers
	// from a InboundShipmentLine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSerialNumbers( inboundShipmentLineId, serialNumbersIds ): Observable<any> {

		// get the InboundShipmentLine
		this.loadHelper( inboundShipmentLineId );


	// split on a comma with no spaces
	var idList 					= serialNumbersIds.split(',');
	var serialNumbers 	= this.inboundShipmentLine.serialNumbers;

	if ( serialNumbers != null && serialNumbersIds != null ) {

		// iterate over array of serialNumbers ids
		serialNumbers.forEach(function (obj) {
			if ( serialNumbersIds.indexOf(obj._id) > -1 ) {
				// remove the SerialNumber
				this.inboundShipmentLine.serialNumbers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InboundShipmentLine
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InboundShipmentLine/update/' + this.inboundShipmentLine;

	return  this.http.post(uri_, this.inboundShipmentLine );
}

	//********************************************************************
	// loadHelper - internal helper to load a InboundShipmentLine
	//********************************************************************	
	loadHelper( id ) {
		this.getInboundShipmentLine(id)
			.subscribe((res : InboundShipmentLine) => {
				this.inboundShipmentLine = res;
			});
	}
}