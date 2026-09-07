import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TransferOrderLine} from '../models/TransferOrderLine';
import {TransferOrderService} from '../services/TransferOrder.service';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {LotService} from '../services/Lot.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TransferOrderLineService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	transferOrderLine : TransferOrderLine;

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
	// add a TransferOrderLine
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTransferOrderLine(lineNumber, quantity, TransferOrder, Sku, Lot, SerialNumbers, FromLocation, ToLocation, UnitOfMeasure, StockStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/TransferOrderLine/create';
		const obj = {
			      		lineNumber: lineNumber,
      		quantity: quantity,
      		TransferOrder: TransferOrder != null && TransferOrder.length > 0 ? TransferOrder : null,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		FromLocation: FromLocation != null && FromLocation.length > 0 ? FromLocation : null,
      		ToLocation: ToLocation != null && ToLocation.length > 0 ? ToLocation : null,
      		UnitOfMeasure: UnitOfMeasure,
			StockStatus: StockStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTransferOrderLine(lineNumber, quantity, TransferOrder, Sku, Lot, SerialNumbers, FromLocation, ToLocation, UnitOfMeasure, StockStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TransferOrderLine/update/' + id;
		const obj = {
				      		lineNumber: lineNumber,
      		quantity: quantity,
      		TransferOrder: TransferOrder != null && TransferOrder.length > 0 ? TransferOrder : null,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		FromLocation: FromLocation != null && FromLocation.length > 0 ? FromLocation : null,
      		ToLocation: ToLocation != null && ToLocation.length > 0 ? ToLocation : null,
      		UnitOfMeasure: UnitOfMeasure,
			StockStatus: StockStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTransferOrderLine(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TransferOrderLine/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TransferOrderLine
	// returns the results untouched as an Observable TransferOrderLine
	// TransferOrderLine model
	// delegates via URI
	//********************************************************************
	getTransferOrderLine(id) : Observable<TransferOrderLine> {
		const uri_ = this.apiUrl + '/TransferOrderLine/load/' + id;

		return this.http.get<TransferOrderLine>(uri_);
	}
	
	//********************************************************************
	// gets all TransferOrderLine
	// returns the results untouched as JSON representation of an
	// Observable array of TransferOrderLine models
	// delegates via URI
	//********************************************************************
	getTransferOrderLines() : Observable<TransferOrderLine[]> {
		const uri_ = this.apiUrl + '/TransferOrderLine/';

		return this
			.http.get<TransferOrderLine[]>(uri_);
	}
	
			//********************************************************************
	// assigns a TransferOrder on a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTransferOrder( transferOrderLineId, _transferOrderId ): Observable<any> {

		// get the TransferOrderLine from storage
		this.loadHelper( transferOrderLineId );

	// get the TransferOrder from storage
	var tmp 	= new TransferOrderService(this.http).getTransferOrder(_transferOrderId);

	// assign the TransferOrder
	this.transferOrderLine.transferOrder = tmp;

	// save the TransferOrderLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TransferOrder on a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTransferOrder( transferOrderLineId ): Observable<any> {

		// get the TransferOrderLine from storage
		this.loadHelper( transferOrderLineId );

	// assign TransferOrder to null
	this.transferOrderLine.transferOrder = null;

	// save the TransferOrderLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Sku on a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( transferOrderLineId, _skuId ): Observable<any> {

		// get the TransferOrderLine from storage
		this.loadHelper( transferOrderLineId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.transferOrderLine.sku = tmp;

	// save the TransferOrderLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( transferOrderLineId ): Observable<any> {

		// get the TransferOrderLine from storage
		this.loadHelper( transferOrderLineId );

	// assign Sku to null
	this.transferOrderLine.sku = null;

	// save the TransferOrderLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lot on a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLot( transferOrderLineId, _lotId ): Observable<any> {

		// get the TransferOrderLine from storage
		this.loadHelper( transferOrderLineId );

	// get the Lot from storage
	var tmp 	= new LotService(this.http).getLot(_lotId);

	// assign the Lot
	this.transferOrderLine.lot = tmp;

	// save the TransferOrderLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lot on a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLot( transferOrderLineId ): Observable<any> {

		// get the TransferOrderLine from storage
		this.loadHelper( transferOrderLineId );

	// assign Lot to null
	this.transferOrderLine.lot = null;

	// save the TransferOrderLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a FromLocation on a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFromLocation( transferOrderLineId, _fromLocationId ): Observable<any> {

		// get the TransferOrderLine from storage
		this.loadHelper( transferOrderLineId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_fromLocationId);

	// assign the FromLocation
	this.transferOrderLine.fromLocation = tmp;

	// save the TransferOrderLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a FromLocation on a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFromLocation( transferOrderLineId ): Observable<any> {

		// get the TransferOrderLine from storage
		this.loadHelper( transferOrderLineId );

	// assign FromLocation to null
	this.transferOrderLine.fromLocation = null;

	// save the TransferOrderLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ToLocation on a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignToLocation( transferOrderLineId, _toLocationId ): Observable<any> {

		// get the TransferOrderLine from storage
		this.loadHelper( transferOrderLineId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_toLocationId);

	// assign the ToLocation
	this.transferOrderLine.toLocation = tmp;

	// save the TransferOrderLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ToLocation on a TransferOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignToLocation( transferOrderLineId ): Observable<any> {

		// get the TransferOrderLine from storage
		this.loadHelper( transferOrderLineId );

	// assign ToLocation to null
	this.transferOrderLine.toLocation = null;

	// save the TransferOrderLine
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more serialNumbersIds as a SerialNumbers
	// to a TransferOrderLine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSerialNumbers( transferOrderLineId, serialNumbersIds ): Observable<any> {

		// get the TransferOrderLine
		this.loadHelper( transferOrderLineId );

	// split on a comma with no spaces
	var idList = serialNumbersIds.split(',')

	// iterate over array of serialNumbers ids
	idList.forEach(function (id) {
		// read the SerialNumber
		var serialNumber = new SerialNumberService(this.http).getSerialNumber(id);
		// add the SerialNumber if not already assigned
		if ( this.transferOrderLine.serialNumbers.indexOf(serialNumber) == -1 )
		this.transferOrderLine.serialNumbers.push(serialNumber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serialNumbersIds as a SerialNumbers
	// from a TransferOrderLine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSerialNumbers( transferOrderLineId, serialNumbersIds ): Observable<any> {

		// get the TransferOrderLine
		this.loadHelper( transferOrderLineId );


	// split on a comma with no spaces
	var idList 					= serialNumbersIds.split(',');
	var serialNumbers 	= this.transferOrderLine.serialNumbers;

	if ( serialNumbers != null && serialNumbersIds != null ) {

		// iterate over array of serialNumbers ids
		serialNumbers.forEach(function (obj) {
			if ( serialNumbersIds.indexOf(obj._id) > -1 ) {
				// remove the SerialNumber
				this.transferOrderLine.serialNumbers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a TransferOrderLine
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TransferOrderLine/update/' + this.transferOrderLine;

	return  this.http.post(uri_, this.transferOrderLine );
}

	//********************************************************************
	// loadHelper - internal helper to load a TransferOrderLine
	//********************************************************************	
	loadHelper( id ) {
		this.getTransferOrderLine(id)
			.subscribe((res : TransferOrderLine) => {
				this.transferOrderLine = res;
			});
	}
}