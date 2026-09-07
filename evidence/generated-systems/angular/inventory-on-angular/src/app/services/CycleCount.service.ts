import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CycleCount} from '../models/CycleCount';
import {WarehouseService} from '../services/Warehouse.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import {CycleCountEntryService} from '../services/CycleCountEntry.service';
import {InventoryTransactionService} from '../services/InventoryTransaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CycleCountService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	cycleCount : CycleCount;

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
	// add a CycleCount
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCycleCount(countNumber, scheduledDate, performedDate, approvedBy, Warehouse, Locations, Entries, Transactions, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/CycleCount/create';
		const obj = {
			      		countNumber: countNumber,
      		scheduledDate: scheduledDate,
      		performedDate: performedDate,
      		approvedBy: approvedBy,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Locations: Locations != null && Locations.length > 0 ? Locations : null,
      		Entries: Entries != null && Entries.length > 0 ? Entries : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CycleCount
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCycleCount(countNumber, scheduledDate, performedDate, approvedBy, Warehouse, Locations, Entries, Transactions, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CycleCount/update/' + id;
		const obj = {
				      		countNumber: countNumber,
      		scheduledDate: scheduledDate,
      		performedDate: performedDate,
      		approvedBy: approvedBy,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Locations: Locations != null && Locations.length > 0 ? Locations : null,
      		Entries: Entries != null && Entries.length > 0 ? Entries : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CycleCount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCycleCount(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CycleCount/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CycleCount
	// returns the results untouched as an Observable CycleCount
	// CycleCount model
	// delegates via URI
	//********************************************************************
	getCycleCount(id) : Observable<CycleCount> {
		const uri_ = this.apiUrl + '/CycleCount/load/' + id;

		return this.http.get<CycleCount>(uri_);
	}
	
	//********************************************************************
	// gets all CycleCount
	// returns the results untouched as JSON representation of an
	// Observable array of CycleCount models
	// delegates via URI
	//********************************************************************
	getCycleCounts() : Observable<CycleCount[]> {
		const uri_ = this.apiUrl + '/CycleCount/';

		return this
			.http.get<CycleCount[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Warehouse on a CycleCount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( cycleCountId, _warehouseId ): Observable<any> {

		// get the CycleCount from storage
		this.loadHelper( cycleCountId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.cycleCount.warehouse = tmp;

	// save the CycleCount
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a CycleCount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( cycleCountId ): Observable<any> {

		// get the CycleCount from storage
		this.loadHelper( cycleCountId );

	// assign Warehouse to null
	this.cycleCount.warehouse = null;

	// save the CycleCount
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more locationsIds as a Locations
	// to a CycleCount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLocations( cycleCountId, locationsIds ): Observable<any> {

		// get the CycleCount
		this.loadHelper( cycleCountId );

	// split on a comma with no spaces
	var idList = locationsIds.split(',')

	// iterate over array of locations ids
	idList.forEach(function (id) {
		// read the StorageLocation
		var storageLocation = new StorageLocationService(this.http).getStorageLocation(id);
		// add the StorageLocation if not already assigned
		if ( this.cycleCount.locations.indexOf(storageLocation) == -1 )
		this.cycleCount.locations.push(storageLocation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more locationsIds as a Locations
	// from a CycleCount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLocations( cycleCountId, locationsIds ): Observable<any> {

		// get the CycleCount
		this.loadHelper( cycleCountId );


	// split on a comma with no spaces
	var idList 					= locationsIds.split(',');
	var locations 	= this.cycleCount.locations;

	if ( locations != null && locationsIds != null ) {

		// iterate over array of locations ids
		locations.forEach(function (obj) {
			if ( locationsIds.indexOf(obj._id) > -1 ) {
				// remove the StorageLocation
				this.cycleCount.locations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more entriesIds as a Entries
	// to a CycleCount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEntries( cycleCountId, entriesIds ): Observable<any> {

		// get the CycleCount
		this.loadHelper( cycleCountId );

	// split on a comma with no spaces
	var idList = entriesIds.split(',')

	// iterate over array of entries ids
	idList.forEach(function (id) {
		// read the CycleCountEntry
		var cycleCountEntry = new CycleCountEntryService(this.http).getCycleCountEntry(id);
		// add the CycleCountEntry if not already assigned
		if ( this.cycleCount.entries.indexOf(cycleCountEntry) == -1 )
		this.cycleCount.entries.push(cycleCountEntry);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more entriesIds as a Entries
	// from a CycleCount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEntries( cycleCountId, entriesIds ): Observable<any> {

		// get the CycleCount
		this.loadHelper( cycleCountId );


	// split on a comma with no spaces
	var idList 					= entriesIds.split(',');
	var entries 	= this.cycleCount.entries;

	if ( entries != null && entriesIds != null ) {

		// iterate over array of entries ids
		entries.forEach(function (obj) {
			if ( entriesIds.indexOf(obj._id) > -1 ) {
				// remove the CycleCountEntry
				this.cycleCount.entries.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more transactionsIds as a Transactions
	// to a CycleCount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTransactions( cycleCountId, transactionsIds ): Observable<any> {

		// get the CycleCount
		this.loadHelper( cycleCountId );

	// split on a comma with no spaces
	var idList = transactionsIds.split(',')

	// iterate over array of transactions ids
	idList.forEach(function (id) {
		// read the InventoryTransaction
		var inventoryTransaction = new InventoryTransactionService(this.http).getInventoryTransaction(id);
		// add the InventoryTransaction if not already assigned
		if ( this.cycleCount.transactions.indexOf(inventoryTransaction) == -1 )
		this.cycleCount.transactions.push(inventoryTransaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more transactionsIds as a Transactions
	// from a CycleCount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTransactions( cycleCountId, transactionsIds ): Observable<any> {

		// get the CycleCount
		this.loadHelper( cycleCountId );


	// split on a comma with no spaces
	var idList 					= transactionsIds.split(',');
	var transactions 	= this.cycleCount.transactions;

	if ( transactions != null && transactionsIds != null ) {

		// iterate over array of transactions ids
		transactions.forEach(function (obj) {
			if ( transactionsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryTransaction
				this.cycleCount.transactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a CycleCount
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CycleCount/update/' + this.cycleCount;

	return  this.http.post(uri_, this.cycleCount );
}

	//********************************************************************
	// loadHelper - internal helper to load a CycleCount
	//********************************************************************	
	loadHelper( id ) {
		this.getCycleCount(id)
			.subscribe((res : CycleCount) => {
				this.cycleCount = res;
			});
	}
}