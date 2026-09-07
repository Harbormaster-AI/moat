import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PurchaseOrder} from '../models/PurchaseOrder';
import {SupplierService} from '../services/Supplier.service';
import {PlantService} from '../services/Plant.service';
import {PurchaseOrderLineService} from '../services/PurchaseOrderLine.service';
import {GoodsReceiptService} from '../services/GoodsReceipt.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PurchaseOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	purchaseOrder : PurchaseOrder;

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
	// add a PurchaseOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPurchaseOrder(poNumber, orderDate, totalAmount, Supplier, Plant, Lines, GoodsReceipts, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/PurchaseOrder/create';
		const obj = {
			      		poNumber: poNumber,
      		orderDate: orderDate,
      		totalAmount: totalAmount,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
      		GoodsReceipts: GoodsReceipts != null && GoodsReceipts.length > 0 ? GoodsReceipts : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PurchaseOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePurchaseOrder(poNumber, orderDate, totalAmount, Supplier, Plant, Lines, GoodsReceipts, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PurchaseOrder/update/' + id;
		const obj = {
				      		poNumber: poNumber,
      		orderDate: orderDate,
      		totalAmount: totalAmount,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
      		GoodsReceipts: GoodsReceipts != null && GoodsReceipts.length > 0 ? GoodsReceipts : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PurchaseOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePurchaseOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PurchaseOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PurchaseOrder
	// returns the results untouched as an Observable PurchaseOrder
	// PurchaseOrder model
	// delegates via URI
	//********************************************************************
	getPurchaseOrder(id) : Observable<PurchaseOrder> {
		const uri_ = this.apiUrl + '/PurchaseOrder/load/' + id;

		return this.http.get<PurchaseOrder>(uri_);
	}
	
	//********************************************************************
	// gets all PurchaseOrder
	// returns the results untouched as JSON representation of an
	// Observable array of PurchaseOrder models
	// delegates via URI
	//********************************************************************
	getPurchaseOrders() : Observable<PurchaseOrder[]> {
		const uri_ = this.apiUrl + '/PurchaseOrder/';

		return this
			.http.get<PurchaseOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Supplier on a PurchaseOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSupplier( purchaseOrderId, _supplierId ): Observable<any> {

		// get the PurchaseOrder from storage
		this.loadHelper( purchaseOrderId );

	// get the Supplier from storage
	var tmp 	= new SupplierService(this.http).getSupplier(_supplierId);

	// assign the Supplier
	this.purchaseOrder.supplier = tmp;

	// save the PurchaseOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Supplier on a PurchaseOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSupplier( purchaseOrderId ): Observable<any> {

		// get the PurchaseOrder from storage
		this.loadHelper( purchaseOrderId );

	// assign Supplier to null
	this.purchaseOrder.supplier = null;

	// save the PurchaseOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Plant on a PurchaseOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( purchaseOrderId, _plantId ): Observable<any> {

		// get the PurchaseOrder from storage
		this.loadHelper( purchaseOrderId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.purchaseOrder.plant = tmp;

	// save the PurchaseOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a PurchaseOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( purchaseOrderId ): Observable<any> {

		// get the PurchaseOrder from storage
		this.loadHelper( purchaseOrderId );

	// assign Plant to null
	this.purchaseOrder.plant = null;

	// save the PurchaseOrder
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more linesIds as a Lines
	// to a PurchaseOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLines( purchaseOrderId, linesIds ): Observable<any> {

		// get the PurchaseOrder
		this.loadHelper( purchaseOrderId );

	// split on a comma with no spaces
	var idList = linesIds.split(',')

	// iterate over array of lines ids
	idList.forEach(function (id) {
		// read the PurchaseOrderLine
		var purchaseOrderLine = new PurchaseOrderLineService(this.http).getPurchaseOrderLine(id);
		// add the PurchaseOrderLine if not already assigned
		if ( this.purchaseOrder.lines.indexOf(purchaseOrderLine) == -1 )
		this.purchaseOrder.lines.push(purchaseOrderLine);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more linesIds as a Lines
	// from a PurchaseOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLines( purchaseOrderId, linesIds ): Observable<any> {

		// get the PurchaseOrder
		this.loadHelper( purchaseOrderId );


	// split on a comma with no spaces
	var idList 					= linesIds.split(',');
	var lines 	= this.purchaseOrder.lines;

	if ( lines != null && linesIds != null ) {

		// iterate over array of lines ids
		lines.forEach(function (obj) {
			if ( linesIds.indexOf(obj._id) > -1 ) {
				// remove the PurchaseOrderLine
				this.purchaseOrder.lines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more goodsReceiptsIds as a GoodsReceipts
	// to a PurchaseOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addGoodsReceipts( purchaseOrderId, goodsReceiptsIds ): Observable<any> {

		// get the PurchaseOrder
		this.loadHelper( purchaseOrderId );

	// split on a comma with no spaces
	var idList = goodsReceiptsIds.split(',')

	// iterate over array of goodsReceipts ids
	idList.forEach(function (id) {
		// read the GoodsReceipt
		var goodsReceipt = new GoodsReceiptService(this.http).getGoodsReceipt(id);
		// add the GoodsReceipt if not already assigned
		if ( this.purchaseOrder.goodsReceipts.indexOf(goodsReceipt) == -1 )
		this.purchaseOrder.goodsReceipts.push(goodsReceipt);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more goodsReceiptsIds as a GoodsReceipts
	// from a PurchaseOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeGoodsReceipts( purchaseOrderId, goodsReceiptsIds ): Observable<any> {

		// get the PurchaseOrder
		this.loadHelper( purchaseOrderId );


	// split on a comma with no spaces
	var idList 					= goodsReceiptsIds.split(',');
	var goodsReceipts 	= this.purchaseOrder.goodsReceipts;

	if ( goodsReceipts != null && goodsReceiptsIds != null ) {

		// iterate over array of goodsReceipts ids
		goodsReceipts.forEach(function (obj) {
			if ( goodsReceiptsIds.indexOf(obj._id) > -1 ) {
				// remove the GoodsReceipt
				this.purchaseOrder.goodsReceipts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PurchaseOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PurchaseOrder/update/' + this.purchaseOrder;

	return  this.http.post(uri_, this.purchaseOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a PurchaseOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getPurchaseOrder(id)
			.subscribe((res : PurchaseOrder) => {
				this.purchaseOrder = res;
			});
	}
}