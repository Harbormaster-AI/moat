import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PurchaseOrderLine} from '../models/PurchaseOrderLine';
import {PurchaseOrderService} from '../services/PurchaseOrder.service';
import {ItemService} from '../services/Item.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PurchaseOrderLineService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	purchaseOrderLine : PurchaseOrderLine;

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
	// add a PurchaseOrderLine
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPurchaseOrderLine(lineNumber, quantity, unitPrice, dueDate, PurchaseOrder, Item) : Observable<any> {
		const uri_ = this.apiUrl + '/PurchaseOrderLine/create';
		const obj = {
			      		lineNumber: lineNumber,
      		quantity: quantity,
      		unitPrice: unitPrice,
      		dueDate: dueDate,
      		PurchaseOrder: PurchaseOrder != null && PurchaseOrder.length > 0 ? PurchaseOrder : null,
			Item: Item != null && Item.length > 0 ? Item : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PurchaseOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePurchaseOrderLine(lineNumber, quantity, unitPrice, dueDate, PurchaseOrder, Item, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PurchaseOrderLine/update/' + id;
		const obj = {
				      		lineNumber: lineNumber,
      		quantity: quantity,
      		unitPrice: unitPrice,
      		dueDate: dueDate,
      		PurchaseOrder: PurchaseOrder != null && PurchaseOrder.length > 0 ? PurchaseOrder : null,
			Item: Item != null && Item.length > 0 ? Item : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PurchaseOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePurchaseOrderLine(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PurchaseOrderLine/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PurchaseOrderLine
	// returns the results untouched as an Observable PurchaseOrderLine
	// PurchaseOrderLine model
	// delegates via URI
	//********************************************************************
	getPurchaseOrderLine(id) : Observable<PurchaseOrderLine> {
		const uri_ = this.apiUrl + '/PurchaseOrderLine/load/' + id;

		return this.http.get<PurchaseOrderLine>(uri_);
	}
	
	//********************************************************************
	// gets all PurchaseOrderLine
	// returns the results untouched as JSON representation of an
	// Observable array of PurchaseOrderLine models
	// delegates via URI
	//********************************************************************
	getPurchaseOrderLines() : Observable<PurchaseOrderLine[]> {
		const uri_ = this.apiUrl + '/PurchaseOrderLine/';

		return this
			.http.get<PurchaseOrderLine[]>(uri_);
	}
	
			//********************************************************************
	// assigns a PurchaseOrder on a PurchaseOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPurchaseOrder( purchaseOrderLineId, _purchaseOrderId ): Observable<any> {

		// get the PurchaseOrderLine from storage
		this.loadHelper( purchaseOrderLineId );

	// get the PurchaseOrder from storage
	var tmp 	= new PurchaseOrderService(this.http).getPurchaseOrder(_purchaseOrderId);

	// assign the PurchaseOrder
	this.purchaseOrderLine.purchaseOrder = tmp;

	// save the PurchaseOrderLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PurchaseOrder on a PurchaseOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPurchaseOrder( purchaseOrderLineId ): Observable<any> {

		// get the PurchaseOrderLine from storage
		this.loadHelper( purchaseOrderLineId );

	// assign PurchaseOrder to null
	this.purchaseOrderLine.purchaseOrder = null;

	// save the PurchaseOrderLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Item on a PurchaseOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( purchaseOrderLineId, _itemId ): Observable<any> {

		// get the PurchaseOrderLine from storage
		this.loadHelper( purchaseOrderLineId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.purchaseOrderLine.item = tmp;

	// save the PurchaseOrderLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a PurchaseOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( purchaseOrderLineId ): Observable<any> {

		// get the PurchaseOrderLine from storage
		this.loadHelper( purchaseOrderLineId );

	// assign Item to null
	this.purchaseOrderLine.item = null;

	// save the PurchaseOrderLine
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a PurchaseOrderLine
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PurchaseOrderLine/update/' + this.purchaseOrderLine;

	return  this.http.post(uri_, this.purchaseOrderLine );
}

	//********************************************************************
	// loadHelper - internal helper to load a PurchaseOrderLine
	//********************************************************************	
	loadHelper( id ) {
		this.getPurchaseOrderLine(id)
			.subscribe((res : PurchaseOrderLine) => {
				this.purchaseOrderLine = res;
			});
	}
}