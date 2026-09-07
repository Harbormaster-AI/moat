import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TransferOrder} from '../models/TransferOrder';
import {WarehouseService} from '../services/Warehouse.service';
import {TransferOrderLineService} from '../services/TransferOrderLine.service';
import {InventoryTransactionService} from '../services/InventoryTransaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TransferOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	transferOrder : TransferOrder;

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
	// add a TransferOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTransferOrder(orderNumber, requestedShipDate, requestedReceiveDate, shippedDate, receivedDate, OriginWarehouse, DestinationWarehouse, Lines, Transactions, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/TransferOrder/create';
		const obj = {
			      		orderNumber: orderNumber,
      		requestedShipDate: requestedShipDate,
      		requestedReceiveDate: requestedReceiveDate,
      		shippedDate: shippedDate,
      		receivedDate: receivedDate,
      		OriginWarehouse: OriginWarehouse != null && OriginWarehouse.length > 0 ? OriginWarehouse : null,
      		DestinationWarehouse: DestinationWarehouse != null && DestinationWarehouse.length > 0 ? DestinationWarehouse : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TransferOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTransferOrder(orderNumber, requestedShipDate, requestedReceiveDate, shippedDate, receivedDate, OriginWarehouse, DestinationWarehouse, Lines, Transactions, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TransferOrder/update/' + id;
		const obj = {
				      		orderNumber: orderNumber,
      		requestedShipDate: requestedShipDate,
      		requestedReceiveDate: requestedReceiveDate,
      		shippedDate: shippedDate,
      		receivedDate: receivedDate,
      		OriginWarehouse: OriginWarehouse != null && OriginWarehouse.length > 0 ? OriginWarehouse : null,
      		DestinationWarehouse: DestinationWarehouse != null && DestinationWarehouse.length > 0 ? DestinationWarehouse : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TransferOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTransferOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TransferOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TransferOrder
	// returns the results untouched as an Observable TransferOrder
	// TransferOrder model
	// delegates via URI
	//********************************************************************
	getTransferOrder(id) : Observable<TransferOrder> {
		const uri_ = this.apiUrl + '/TransferOrder/load/' + id;

		return this.http.get<TransferOrder>(uri_);
	}
	
	//********************************************************************
	// gets all TransferOrder
	// returns the results untouched as JSON representation of an
	// Observable array of TransferOrder models
	// delegates via URI
	//********************************************************************
	getTransferOrders() : Observable<TransferOrder[]> {
		const uri_ = this.apiUrl + '/TransferOrder/';

		return this
			.http.get<TransferOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a OriginWarehouse on a TransferOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOriginWarehouse( transferOrderId, _originWarehouseId ): Observable<any> {

		// get the TransferOrder from storage
		this.loadHelper( transferOrderId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_originWarehouseId);

	// assign the OriginWarehouse
	this.transferOrder.originWarehouse = tmp;

	// save the TransferOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a OriginWarehouse on a TransferOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOriginWarehouse( transferOrderId ): Observable<any> {

		// get the TransferOrder from storage
		this.loadHelper( transferOrderId );

	// assign OriginWarehouse to null
	this.transferOrder.originWarehouse = null;

	// save the TransferOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a DestinationWarehouse on a TransferOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDestinationWarehouse( transferOrderId, _destinationWarehouseId ): Observable<any> {

		// get the TransferOrder from storage
		this.loadHelper( transferOrderId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_destinationWarehouseId);

	// assign the DestinationWarehouse
	this.transferOrder.destinationWarehouse = tmp;

	// save the TransferOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a DestinationWarehouse on a TransferOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDestinationWarehouse( transferOrderId ): Observable<any> {

		// get the TransferOrder from storage
		this.loadHelper( transferOrderId );

	// assign DestinationWarehouse to null
	this.transferOrder.destinationWarehouse = null;

	// save the TransferOrder
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more linesIds as a Lines
	// to a TransferOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLines( transferOrderId, linesIds ): Observable<any> {

		// get the TransferOrder
		this.loadHelper( transferOrderId );

	// split on a comma with no spaces
	var idList = linesIds.split(',')

	// iterate over array of lines ids
	idList.forEach(function (id) {
		// read the TransferOrderLine
		var transferOrderLine = new TransferOrderLineService(this.http).getTransferOrderLine(id);
		// add the TransferOrderLine if not already assigned
		if ( this.transferOrder.lines.indexOf(transferOrderLine) == -1 )
		this.transferOrder.lines.push(transferOrderLine);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more linesIds as a Lines
	// from a TransferOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLines( transferOrderId, linesIds ): Observable<any> {

		// get the TransferOrder
		this.loadHelper( transferOrderId );


	// split on a comma with no spaces
	var idList 					= linesIds.split(',');
	var lines 	= this.transferOrder.lines;

	if ( lines != null && linesIds != null ) {

		// iterate over array of lines ids
		lines.forEach(function (obj) {
			if ( linesIds.indexOf(obj._id) > -1 ) {
				// remove the TransferOrderLine
				this.transferOrder.lines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more transactionsIds as a Transactions
	// to a TransferOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTransactions( transferOrderId, transactionsIds ): Observable<any> {

		// get the TransferOrder
		this.loadHelper( transferOrderId );

	// split on a comma with no spaces
	var idList = transactionsIds.split(',')

	// iterate over array of transactions ids
	idList.forEach(function (id) {
		// read the InventoryTransaction
		var inventoryTransaction = new InventoryTransactionService(this.http).getInventoryTransaction(id);
		// add the InventoryTransaction if not already assigned
		if ( this.transferOrder.transactions.indexOf(inventoryTransaction) == -1 )
		this.transferOrder.transactions.push(inventoryTransaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more transactionsIds as a Transactions
	// from a TransferOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTransactions( transferOrderId, transactionsIds ): Observable<any> {

		// get the TransferOrder
		this.loadHelper( transferOrderId );


	// split on a comma with no spaces
	var idList 					= transactionsIds.split(',');
	var transactions 	= this.transferOrder.transactions;

	if ( transactions != null && transactionsIds != null ) {

		// iterate over array of transactions ids
		transactions.forEach(function (obj) {
			if ( transactionsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryTransaction
				this.transferOrder.transactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a TransferOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TransferOrder/update/' + this.transferOrder;

	return  this.http.post(uri_, this.transferOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a TransferOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getTransferOrder(id)
			.subscribe((res : TransferOrder) => {
				this.transferOrder = res;
			});
	}
}