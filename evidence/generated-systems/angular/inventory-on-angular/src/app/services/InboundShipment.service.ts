import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InboundShipment} from '../models/InboundShipment';
import {WarehouseService} from '../services/Warehouse.service';
import {InboundShipmentLineService} from '../services/InboundShipmentLine.service';
import {InventoryTransactionService} from '../services/InventoryTransaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InboundShipmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inboundShipment : InboundShipment;

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
	// add a InboundShipment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInboundShipment(shipmentNumber, expectedArrivalDate, arrivalDate, carrierName, Warehouse, Lines, Transactions, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/InboundShipment/create';
		const obj = {
			      		shipmentNumber: shipmentNumber,
      		expectedArrivalDate: expectedArrivalDate,
      		arrivalDate: arrivalDate,
      		carrierName: carrierName,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InboundShipment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInboundShipment(shipmentNumber, expectedArrivalDate, arrivalDate, carrierName, Warehouse, Lines, Transactions, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InboundShipment/update/' + id;
		const obj = {
				      		shipmentNumber: shipmentNumber,
      		expectedArrivalDate: expectedArrivalDate,
      		arrivalDate: arrivalDate,
      		carrierName: carrierName,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InboundShipment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInboundShipment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InboundShipment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InboundShipment
	// returns the results untouched as an Observable InboundShipment
	// InboundShipment model
	// delegates via URI
	//********************************************************************
	getInboundShipment(id) : Observable<InboundShipment> {
		const uri_ = this.apiUrl + '/InboundShipment/load/' + id;

		return this.http.get<InboundShipment>(uri_);
	}
	
	//********************************************************************
	// gets all InboundShipment
	// returns the results untouched as JSON representation of an
	// Observable array of InboundShipment models
	// delegates via URI
	//********************************************************************
	getInboundShipments() : Observable<InboundShipment[]> {
		const uri_ = this.apiUrl + '/InboundShipment/';

		return this
			.http.get<InboundShipment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Warehouse on a InboundShipment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( inboundShipmentId, _warehouseId ): Observable<any> {

		// get the InboundShipment from storage
		this.loadHelper( inboundShipmentId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.inboundShipment.warehouse = tmp;

	// save the InboundShipment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a InboundShipment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( inboundShipmentId ): Observable<any> {

		// get the InboundShipment from storage
		this.loadHelper( inboundShipmentId );

	// assign Warehouse to null
	this.inboundShipment.warehouse = null;

	// save the InboundShipment
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more linesIds as a Lines
	// to a InboundShipment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLines( inboundShipmentId, linesIds ): Observable<any> {

		// get the InboundShipment
		this.loadHelper( inboundShipmentId );

	// split on a comma with no spaces
	var idList = linesIds.split(',')

	// iterate over array of lines ids
	idList.forEach(function (id) {
		// read the InboundShipmentLine
		var inboundShipmentLine = new InboundShipmentLineService(this.http).getInboundShipmentLine(id);
		// add the InboundShipmentLine if not already assigned
		if ( this.inboundShipment.lines.indexOf(inboundShipmentLine) == -1 )
		this.inboundShipment.lines.push(inboundShipmentLine);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more linesIds as a Lines
	// from a InboundShipment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLines( inboundShipmentId, linesIds ): Observable<any> {

		// get the InboundShipment
		this.loadHelper( inboundShipmentId );


	// split on a comma with no spaces
	var idList 					= linesIds.split(',');
	var lines 	= this.inboundShipment.lines;

	if ( lines != null && linesIds != null ) {

		// iterate over array of lines ids
		lines.forEach(function (obj) {
			if ( linesIds.indexOf(obj._id) > -1 ) {
				// remove the InboundShipmentLine
				this.inboundShipment.lines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more transactionsIds as a Transactions
	// to a InboundShipment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTransactions( inboundShipmentId, transactionsIds ): Observable<any> {

		// get the InboundShipment
		this.loadHelper( inboundShipmentId );

	// split on a comma with no spaces
	var idList = transactionsIds.split(',')

	// iterate over array of transactions ids
	idList.forEach(function (id) {
		// read the InventoryTransaction
		var inventoryTransaction = new InventoryTransactionService(this.http).getInventoryTransaction(id);
		// add the InventoryTransaction if not already assigned
		if ( this.inboundShipment.transactions.indexOf(inventoryTransaction) == -1 )
		this.inboundShipment.transactions.push(inventoryTransaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more transactionsIds as a Transactions
	// from a InboundShipment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTransactions( inboundShipmentId, transactionsIds ): Observable<any> {

		// get the InboundShipment
		this.loadHelper( inboundShipmentId );


	// split on a comma with no spaces
	var idList 					= transactionsIds.split(',');
	var transactions 	= this.inboundShipment.transactions;

	if ( transactions != null && transactionsIds != null ) {

		// iterate over array of transactions ids
		transactions.forEach(function (obj) {
			if ( transactionsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryTransaction
				this.inboundShipment.transactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InboundShipment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InboundShipment/update/' + this.inboundShipment;

	return  this.http.post(uri_, this.inboundShipment );
}

	//********************************************************************
	// loadHelper - internal helper to load a InboundShipment
	//********************************************************************	
	loadHelper( id ) {
		this.getInboundShipment(id)
			.subscribe((res : InboundShipment) => {
				this.inboundShipment = res;
			});
	}
}