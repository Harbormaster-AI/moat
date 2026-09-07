import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {GoodsReceipt} from '../models/GoodsReceipt';
import {PurchaseOrderService} from '../services/PurchaseOrder.service';
import {WarehouseService} from '../services/Warehouse.service';
import {GoodsReceiptLineService} from '../services/GoodsReceiptLine.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class GoodsReceiptService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	goodsReceipt : GoodsReceipt;

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
	// add a GoodsReceipt
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addGoodsReceipt(receiptNumber, receiptDate, PurchaseOrder, Warehouse, Lines, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/GoodsReceipt/create';
		const obj = {
			      		receiptNumber: receiptNumber,
      		receiptDate: receiptDate,
      		PurchaseOrder: PurchaseOrder != null && PurchaseOrder.length > 0 ? PurchaseOrder : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a GoodsReceipt
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateGoodsReceipt(receiptNumber, receiptDate, PurchaseOrder, Warehouse, Lines, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/GoodsReceipt/update/' + id;
		const obj = {
				      		receiptNumber: receiptNumber,
      		receiptDate: receiptDate,
      		PurchaseOrder: PurchaseOrder != null && PurchaseOrder.length > 0 ? PurchaseOrder : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a GoodsReceipt
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteGoodsReceipt(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/GoodsReceipt/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a GoodsReceipt
	// returns the results untouched as an Observable GoodsReceipt
	// GoodsReceipt model
	// delegates via URI
	//********************************************************************
	getGoodsReceipt(id) : Observable<GoodsReceipt> {
		const uri_ = this.apiUrl + '/GoodsReceipt/load/' + id;

		return this.http.get<GoodsReceipt>(uri_);
	}
	
	//********************************************************************
	// gets all GoodsReceipt
	// returns the results untouched as JSON representation of an
	// Observable array of GoodsReceipt models
	// delegates via URI
	//********************************************************************
	getGoodsReceipts() : Observable<GoodsReceipt[]> {
		const uri_ = this.apiUrl + '/GoodsReceipt/';

		return this
			.http.get<GoodsReceipt[]>(uri_);
	}
	
			//********************************************************************
	// assigns a PurchaseOrder on a GoodsReceipt
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPurchaseOrder( goodsReceiptId, _purchaseOrderId ): Observable<any> {

		// get the GoodsReceipt from storage
		this.loadHelper( goodsReceiptId );

	// get the PurchaseOrder from storage
	var tmp 	= new PurchaseOrderService(this.http).getPurchaseOrder(_purchaseOrderId);

	// assign the PurchaseOrder
	this.goodsReceipt.purchaseOrder = tmp;

	// save the GoodsReceipt
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PurchaseOrder on a GoodsReceipt
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPurchaseOrder( goodsReceiptId ): Observable<any> {

		// get the GoodsReceipt from storage
		this.loadHelper( goodsReceiptId );

	// assign PurchaseOrder to null
	this.goodsReceipt.purchaseOrder = null;

	// save the GoodsReceipt
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Warehouse on a GoodsReceipt
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( goodsReceiptId, _warehouseId ): Observable<any> {

		// get the GoodsReceipt from storage
		this.loadHelper( goodsReceiptId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.goodsReceipt.warehouse = tmp;

	// save the GoodsReceipt
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a GoodsReceipt
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( goodsReceiptId ): Observable<any> {

		// get the GoodsReceipt from storage
		this.loadHelper( goodsReceiptId );

	// assign Warehouse to null
	this.goodsReceipt.warehouse = null;

	// save the GoodsReceipt
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more linesIds as a Lines
	// to a GoodsReceipt
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLines( goodsReceiptId, linesIds ): Observable<any> {

		// get the GoodsReceipt
		this.loadHelper( goodsReceiptId );

	// split on a comma with no spaces
	var idList = linesIds.split(',')

	// iterate over array of lines ids
	idList.forEach(function (id) {
		// read the GoodsReceiptLine
		var goodsReceiptLine = new GoodsReceiptLineService(this.http).getGoodsReceiptLine(id);
		// add the GoodsReceiptLine if not already assigned
		if ( this.goodsReceipt.lines.indexOf(goodsReceiptLine) == -1 )
		this.goodsReceipt.lines.push(goodsReceiptLine);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more linesIds as a Lines
	// from a GoodsReceipt
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLines( goodsReceiptId, linesIds ): Observable<any> {

		// get the GoodsReceipt
		this.loadHelper( goodsReceiptId );


	// split on a comma with no spaces
	var idList 					= linesIds.split(',');
	var lines 	= this.goodsReceipt.lines;

	if ( lines != null && linesIds != null ) {

		// iterate over array of lines ids
		lines.forEach(function (obj) {
			if ( linesIds.indexOf(obj._id) > -1 ) {
				// remove the GoodsReceiptLine
				this.goodsReceipt.lines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a GoodsReceipt
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/GoodsReceipt/update/' + this.goodsReceipt;

	return  this.http.post(uri_, this.goodsReceipt );
}

	//********************************************************************
	// loadHelper - internal helper to load a GoodsReceipt
	//********************************************************************	
	loadHelper( id ) {
		this.getGoodsReceipt(id)
			.subscribe((res : GoodsReceipt) => {
				this.goodsReceipt = res;
			});
	}
}