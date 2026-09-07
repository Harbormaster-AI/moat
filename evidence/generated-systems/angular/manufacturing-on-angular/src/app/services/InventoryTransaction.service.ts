import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InventoryTransaction} from '../models/InventoryTransaction';
import {ItemService} from '../services/Item.service';
import {LocationService} from '../services/Location.service';
import {WorkOrderService} from '../services/WorkOrder.service';
import {PurchaseOrderService} from '../services/PurchaseOrder.service';
import {SalesOrderService} from '../services/SalesOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InventoryTransactionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inventoryTransaction : InventoryTransaction;

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
	// add a InventoryTransaction
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInventoryTransaction(transactionNumber, quantity, transactionDateTime, referenceDocument, Item, Location, WorkOrder, PurchaseOrder, SalesOrder, TransactionType) : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryTransaction/create';
		const obj = {
			      		transactionNumber: transactionNumber,
      		quantity: quantity,
      		transactionDateTime: transactionDateTime,
      		referenceDocument: referenceDocument,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		WorkOrder: WorkOrder != null && WorkOrder.length > 0 ? WorkOrder : null,
      		PurchaseOrder: PurchaseOrder != null && PurchaseOrder.length > 0 ? PurchaseOrder : null,
      		SalesOrder: SalesOrder != null && SalesOrder.length > 0 ? SalesOrder : null,
			TransactionType: TransactionType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInventoryTransaction(transactionNumber, quantity, transactionDateTime, referenceDocument, Item, Location, WorkOrder, PurchaseOrder, SalesOrder, TransactionType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InventoryTransaction/update/' + id;
		const obj = {
				      		transactionNumber: transactionNumber,
      		quantity: quantity,
      		transactionDateTime: transactionDateTime,
      		referenceDocument: referenceDocument,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		WorkOrder: WorkOrder != null && WorkOrder.length > 0 ? WorkOrder : null,
      		PurchaseOrder: PurchaseOrder != null && PurchaseOrder.length > 0 ? PurchaseOrder : null,
      		SalesOrder: SalesOrder != null && SalesOrder.length > 0 ? SalesOrder : null,
			TransactionType: TransactionType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInventoryTransaction(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryTransaction/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InventoryTransaction
	// returns the results untouched as an Observable InventoryTransaction
	// InventoryTransaction model
	// delegates via URI
	//********************************************************************
	getInventoryTransaction(id) : Observable<InventoryTransaction> {
		const uri_ = this.apiUrl + '/InventoryTransaction/load/' + id;

		return this.http.get<InventoryTransaction>(uri_);
	}
	
	//********************************************************************
	// gets all InventoryTransaction
	// returns the results untouched as JSON representation of an
	// Observable array of InventoryTransaction models
	// delegates via URI
	//********************************************************************
	getInventoryTransactions() : Observable<InventoryTransaction[]> {
		const uri_ = this.apiUrl + '/InventoryTransaction/';

		return this
			.http.get<InventoryTransaction[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Item on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( inventoryTransactionId, _itemId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.inventoryTransaction.item = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign Item to null
	this.inventoryTransaction.item = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Location on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLocation( inventoryTransactionId, _locationId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the Location from storage
	var tmp 	= new LocationService(this.http).getLocation(_locationId);

	// assign the Location
	this.inventoryTransaction.location = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Location on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLocation( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign Location to null
	this.inventoryTransaction.location = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a WorkOrder on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkOrder( inventoryTransactionId, _workOrderId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the WorkOrder from storage
	var tmp 	= new WorkOrderService(this.http).getWorkOrder(_workOrderId);

	// assign the WorkOrder
	this.inventoryTransaction.workOrder = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkOrder on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkOrder( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign WorkOrder to null
	this.inventoryTransaction.workOrder = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PurchaseOrder on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPurchaseOrder( inventoryTransactionId, _purchaseOrderId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the PurchaseOrder from storage
	var tmp 	= new PurchaseOrderService(this.http).getPurchaseOrder(_purchaseOrderId);

	// assign the PurchaseOrder
	this.inventoryTransaction.purchaseOrder = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PurchaseOrder on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPurchaseOrder( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign PurchaseOrder to null
	this.inventoryTransaction.purchaseOrder = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a SalesOrder on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSalesOrder( inventoryTransactionId, _salesOrderId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// get the SalesOrder from storage
	var tmp 	= new SalesOrderService(this.http).getSalesOrder(_salesOrderId);

	// assign the SalesOrder
	this.inventoryTransaction.salesOrder = tmp;

	// save the InventoryTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SalesOrder on a InventoryTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSalesOrder( inventoryTransactionId ): Observable<any> {

		// get the InventoryTransaction from storage
		this.loadHelper( inventoryTransactionId );

	// assign SalesOrder to null
	this.inventoryTransaction.salesOrder = null;

	// save the InventoryTransaction
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a InventoryTransaction
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InventoryTransaction/update/' + this.inventoryTransaction;

	return  this.http.post(uri_, this.inventoryTransaction );
}

	//********************************************************************
	// loadHelper - internal helper to load a InventoryTransaction
	//********************************************************************	
	loadHelper( id ) {
		this.getInventoryTransaction(id)
			.subscribe((res : InventoryTransaction) => {
				this.inventoryTransaction = res;
			});
	}
}