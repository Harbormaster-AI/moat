import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {GoodsReceiptLine} from '../models/GoodsReceiptLine';
import {GoodsReceiptService} from '../services/GoodsReceipt.service';
import {ItemService} from '../services/Item.service';
import {InventoryTransactionService} from '../services/InventoryTransaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class GoodsReceiptLineService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	goodsReceiptLine : GoodsReceiptLine;

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
	// add a GoodsReceiptLine
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addGoodsReceiptLine(lineNumber, receivedQuantity, acceptedQuantity, rejectedQuantity, lot, GoodsReceipt, Item, InventoryTransaction) : Observable<any> {
		const uri_ = this.apiUrl + '/GoodsReceiptLine/create';
		const obj = {
			      		lineNumber: lineNumber,
      		receivedQuantity: receivedQuantity,
      		acceptedQuantity: acceptedQuantity,
      		rejectedQuantity: rejectedQuantity,
      		lot: lot,
      		GoodsReceipt: GoodsReceipt != null && GoodsReceipt.length > 0 ? GoodsReceipt : null,
      		Item: Item != null && Item.length > 0 ? Item : null,
			InventoryTransaction: InventoryTransaction != null && InventoryTransaction.length > 0 ? InventoryTransaction : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a GoodsReceiptLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateGoodsReceiptLine(lineNumber, receivedQuantity, acceptedQuantity, rejectedQuantity, lot, GoodsReceipt, Item, InventoryTransaction, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/GoodsReceiptLine/update/' + id;
		const obj = {
				      		lineNumber: lineNumber,
      		receivedQuantity: receivedQuantity,
      		acceptedQuantity: acceptedQuantity,
      		rejectedQuantity: rejectedQuantity,
      		lot: lot,
      		GoodsReceipt: GoodsReceipt != null && GoodsReceipt.length > 0 ? GoodsReceipt : null,
      		Item: Item != null && Item.length > 0 ? Item : null,
			InventoryTransaction: InventoryTransaction != null && InventoryTransaction.length > 0 ? InventoryTransaction : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a GoodsReceiptLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteGoodsReceiptLine(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/GoodsReceiptLine/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a GoodsReceiptLine
	// returns the results untouched as an Observable GoodsReceiptLine
	// GoodsReceiptLine model
	// delegates via URI
	//********************************************************************
	getGoodsReceiptLine(id) : Observable<GoodsReceiptLine> {
		const uri_ = this.apiUrl + '/GoodsReceiptLine/load/' + id;

		return this.http.get<GoodsReceiptLine>(uri_);
	}
	
	//********************************************************************
	// gets all GoodsReceiptLine
	// returns the results untouched as JSON representation of an
	// Observable array of GoodsReceiptLine models
	// delegates via URI
	//********************************************************************
	getGoodsReceiptLines() : Observable<GoodsReceiptLine[]> {
		const uri_ = this.apiUrl + '/GoodsReceiptLine/';

		return this
			.http.get<GoodsReceiptLine[]>(uri_);
	}
	
			//********************************************************************
	// assigns a GoodsReceipt on a GoodsReceiptLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignGoodsReceipt( goodsReceiptLineId, _goodsReceiptId ): Observable<any> {

		// get the GoodsReceiptLine from storage
		this.loadHelper( goodsReceiptLineId );

	// get the GoodsReceipt from storage
	var tmp 	= new GoodsReceiptService(this.http).getGoodsReceipt(_goodsReceiptId);

	// assign the GoodsReceipt
	this.goodsReceiptLine.goodsReceipt = tmp;

	// save the GoodsReceiptLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a GoodsReceipt on a GoodsReceiptLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignGoodsReceipt( goodsReceiptLineId ): Observable<any> {

		// get the GoodsReceiptLine from storage
		this.loadHelper( goodsReceiptLineId );

	// assign GoodsReceipt to null
	this.goodsReceiptLine.goodsReceipt = null;

	// save the GoodsReceiptLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Item on a GoodsReceiptLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( goodsReceiptLineId, _itemId ): Observable<any> {

		// get the GoodsReceiptLine from storage
		this.loadHelper( goodsReceiptLineId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.goodsReceiptLine.item = tmp;

	// save the GoodsReceiptLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a GoodsReceiptLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( goodsReceiptLineId ): Observable<any> {

		// get the GoodsReceiptLine from storage
		this.loadHelper( goodsReceiptLineId );

	// assign Item to null
	this.goodsReceiptLine.item = null;

	// save the GoodsReceiptLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a InventoryTransaction on a GoodsReceiptLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInventoryTransaction( goodsReceiptLineId, _inventoryTransactionId ): Observable<any> {

		// get the GoodsReceiptLine from storage
		this.loadHelper( goodsReceiptLineId );

	// get the InventoryTransaction from storage
	var tmp 	= new InventoryTransactionService(this.http).getInventoryTransaction(_inventoryTransactionId);

	// assign the InventoryTransaction
	this.goodsReceiptLine.inventoryTransaction = tmp;

	// save the GoodsReceiptLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InventoryTransaction on a GoodsReceiptLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInventoryTransaction( goodsReceiptLineId ): Observable<any> {

		// get the GoodsReceiptLine from storage
		this.loadHelper( goodsReceiptLineId );

	// assign InventoryTransaction to null
	this.goodsReceiptLine.inventoryTransaction = null;

	// save the GoodsReceiptLine
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a GoodsReceiptLine
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/GoodsReceiptLine/update/' + this.goodsReceiptLine;

	return  this.http.post(uri_, this.goodsReceiptLine );
}

	//********************************************************************
	// loadHelper - internal helper to load a GoodsReceiptLine
	//********************************************************************	
	loadHelper( id ) {
		this.getGoodsReceiptLine(id)
			.subscribe((res : GoodsReceiptLine) => {
				this.goodsReceiptLine = res;
			});
	}
}