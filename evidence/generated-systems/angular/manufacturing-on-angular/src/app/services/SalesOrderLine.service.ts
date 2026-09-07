import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SalesOrderLine} from '../models/SalesOrderLine';
import {SalesOrderService} from '../services/SalesOrder.service';
import {ItemService} from '../services/Item.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SalesOrderLineService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	salesOrderLine : SalesOrderLine;

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
	// add a SalesOrderLine
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSalesOrderLine(lineNumber, quantity, unitPrice, dueDate, SalesOrder, Item) : Observable<any> {
		const uri_ = this.apiUrl + '/SalesOrderLine/create';
		const obj = {
			      		lineNumber: lineNumber,
      		quantity: quantity,
      		unitPrice: unitPrice,
      		dueDate: dueDate,
      		SalesOrder: SalesOrder != null && SalesOrder.length > 0 ? SalesOrder : null,
			Item: Item != null && Item.length > 0 ? Item : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SalesOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSalesOrderLine(lineNumber, quantity, unitPrice, dueDate, SalesOrder, Item, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SalesOrderLine/update/' + id;
		const obj = {
				      		lineNumber: lineNumber,
      		quantity: quantity,
      		unitPrice: unitPrice,
      		dueDate: dueDate,
      		SalesOrder: SalesOrder != null && SalesOrder.length > 0 ? SalesOrder : null,
			Item: Item != null && Item.length > 0 ? Item : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SalesOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSalesOrderLine(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SalesOrderLine/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SalesOrderLine
	// returns the results untouched as an Observable SalesOrderLine
	// SalesOrderLine model
	// delegates via URI
	//********************************************************************
	getSalesOrderLine(id) : Observable<SalesOrderLine> {
		const uri_ = this.apiUrl + '/SalesOrderLine/load/' + id;

		return this.http.get<SalesOrderLine>(uri_);
	}
	
	//********************************************************************
	// gets all SalesOrderLine
	// returns the results untouched as JSON representation of an
	// Observable array of SalesOrderLine models
	// delegates via URI
	//********************************************************************
	getSalesOrderLines() : Observable<SalesOrderLine[]> {
		const uri_ = this.apiUrl + '/SalesOrderLine/';

		return this
			.http.get<SalesOrderLine[]>(uri_);
	}
	
			//********************************************************************
	// assigns a SalesOrder on a SalesOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSalesOrder( salesOrderLineId, _salesOrderId ): Observable<any> {

		// get the SalesOrderLine from storage
		this.loadHelper( salesOrderLineId );

	// get the SalesOrder from storage
	var tmp 	= new SalesOrderService(this.http).getSalesOrder(_salesOrderId);

	// assign the SalesOrder
	this.salesOrderLine.salesOrder = tmp;

	// save the SalesOrderLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SalesOrder on a SalesOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSalesOrder( salesOrderLineId ): Observable<any> {

		// get the SalesOrderLine from storage
		this.loadHelper( salesOrderLineId );

	// assign SalesOrder to null
	this.salesOrderLine.salesOrder = null;

	// save the SalesOrderLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Item on a SalesOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( salesOrderLineId, _itemId ): Observable<any> {

		// get the SalesOrderLine from storage
		this.loadHelper( salesOrderLineId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.salesOrderLine.item = tmp;

	// save the SalesOrderLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a SalesOrderLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( salesOrderLineId ): Observable<any> {

		// get the SalesOrderLine from storage
		this.loadHelper( salesOrderLineId );

	// assign Item to null
	this.salesOrderLine.item = null;

	// save the SalesOrderLine
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a SalesOrderLine
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SalesOrderLine/update/' + this.salesOrderLine;

	return  this.http.post(uri_, this.salesOrderLine );
}

	//********************************************************************
	// loadHelper - internal helper to load a SalesOrderLine
	//********************************************************************	
	loadHelper( id ) {
		this.getSalesOrderLine(id)
			.subscribe((res : SalesOrderLine) => {
				this.salesOrderLine = res;
			});
	}
}