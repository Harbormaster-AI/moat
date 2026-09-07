import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {StockAdjustment} from '../models/StockAdjustment';
import {WarehouseService} from '../services/Warehouse.service';
import {StockAdjustmentLineService} from '../services/StockAdjustmentLine.service';
import {InventoryTransactionService} from '../services/InventoryTransaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class StockAdjustmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	stockAdjustment : StockAdjustment;

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
	// add a StockAdjustment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addStockAdjustment(adjustmentNumber, reason, adjustmentDate, Warehouse, Lines, Transactions, AdjustmentType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/StockAdjustment/create';
		const obj = {
			      		adjustmentNumber: adjustmentNumber,
      		reason: reason,
      		adjustmentDate: adjustmentDate,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
      		AdjustmentType: AdjustmentType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a StockAdjustment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateStockAdjustment(adjustmentNumber, reason, adjustmentDate, Warehouse, Lines, Transactions, AdjustmentType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/StockAdjustment/update/' + id;
		const obj = {
				      		adjustmentNumber: adjustmentNumber,
      		reason: reason,
      		adjustmentDate: adjustmentDate,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
      		AdjustmentType: AdjustmentType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a StockAdjustment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteStockAdjustment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/StockAdjustment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a StockAdjustment
	// returns the results untouched as an Observable StockAdjustment
	// StockAdjustment model
	// delegates via URI
	//********************************************************************
	getStockAdjustment(id) : Observable<StockAdjustment> {
		const uri_ = this.apiUrl + '/StockAdjustment/load/' + id;

		return this.http.get<StockAdjustment>(uri_);
	}
	
	//********************************************************************
	// gets all StockAdjustment
	// returns the results untouched as JSON representation of an
	// Observable array of StockAdjustment models
	// delegates via URI
	//********************************************************************
	getStockAdjustments() : Observable<StockAdjustment[]> {
		const uri_ = this.apiUrl + '/StockAdjustment/';

		return this
			.http.get<StockAdjustment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Warehouse on a StockAdjustment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( stockAdjustmentId, _warehouseId ): Observable<any> {

		// get the StockAdjustment from storage
		this.loadHelper( stockAdjustmentId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.stockAdjustment.warehouse = tmp;

	// save the StockAdjustment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a StockAdjustment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( stockAdjustmentId ): Observable<any> {

		// get the StockAdjustment from storage
		this.loadHelper( stockAdjustmentId );

	// assign Warehouse to null
	this.stockAdjustment.warehouse = null;

	// save the StockAdjustment
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more linesIds as a Lines
	// to a StockAdjustment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLines( stockAdjustmentId, linesIds ): Observable<any> {

		// get the StockAdjustment
		this.loadHelper( stockAdjustmentId );

	// split on a comma with no spaces
	var idList = linesIds.split(',')

	// iterate over array of lines ids
	idList.forEach(function (id) {
		// read the StockAdjustmentLine
		var stockAdjustmentLine = new StockAdjustmentLineService(this.http).getStockAdjustmentLine(id);
		// add the StockAdjustmentLine if not already assigned
		if ( this.stockAdjustment.lines.indexOf(stockAdjustmentLine) == -1 )
		this.stockAdjustment.lines.push(stockAdjustmentLine);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more linesIds as a Lines
	// from a StockAdjustment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLines( stockAdjustmentId, linesIds ): Observable<any> {

		// get the StockAdjustment
		this.loadHelper( stockAdjustmentId );


	// split on a comma with no spaces
	var idList 					= linesIds.split(',');
	var lines 	= this.stockAdjustment.lines;

	if ( lines != null && linesIds != null ) {

		// iterate over array of lines ids
		lines.forEach(function (obj) {
			if ( linesIds.indexOf(obj._id) > -1 ) {
				// remove the StockAdjustmentLine
				this.stockAdjustment.lines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more transactionsIds as a Transactions
	// to a StockAdjustment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTransactions( stockAdjustmentId, transactionsIds ): Observable<any> {

		// get the StockAdjustment
		this.loadHelper( stockAdjustmentId );

	// split on a comma with no spaces
	var idList = transactionsIds.split(',')

	// iterate over array of transactions ids
	idList.forEach(function (id) {
		// read the InventoryTransaction
		var inventoryTransaction = new InventoryTransactionService(this.http).getInventoryTransaction(id);
		// add the InventoryTransaction if not already assigned
		if ( this.stockAdjustment.transactions.indexOf(inventoryTransaction) == -1 )
		this.stockAdjustment.transactions.push(inventoryTransaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more transactionsIds as a Transactions
	// from a StockAdjustment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTransactions( stockAdjustmentId, transactionsIds ): Observable<any> {

		// get the StockAdjustment
		this.loadHelper( stockAdjustmentId );


	// split on a comma with no spaces
	var idList 					= transactionsIds.split(',');
	var transactions 	= this.stockAdjustment.transactions;

	if ( transactions != null && transactionsIds != null ) {

		// iterate over array of transactions ids
		transactions.forEach(function (obj) {
			if ( transactionsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryTransaction
				this.stockAdjustment.transactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a StockAdjustment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/StockAdjustment/update/' + this.stockAdjustment;

	return  this.http.post(uri_, this.stockAdjustment );
}

	//********************************************************************
	// loadHelper - internal helper to load a StockAdjustment
	//********************************************************************	
	loadHelper( id ) {
		this.getStockAdjustment(id)
			.subscribe((res : StockAdjustment) => {
				this.stockAdjustment = res;
			});
	}
}