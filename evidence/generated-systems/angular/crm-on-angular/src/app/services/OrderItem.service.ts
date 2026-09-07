import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {OrderItem} from '../models/OrderItem';
import {OrderService} from '../services/Order.service';
import {ProductService} from '../services/Product.service';
import {PriceBookEntryService} from '../services/PriceBookEntry.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OrderItemService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	orderItem : OrderItem;

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
	// add a OrderItem
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOrderItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Order, Product, PriceBookEntry) : Observable<any> {
		const uri_ = this.apiUrl + '/OrderItem/create';
		const obj = {
			      		quantity: quantity,
      		unitPrice: unitPrice,
      		discountAmount: discountAmount,
      		taxAmount: taxAmount,
      		totalAmount: totalAmount,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		Product: Product != null && Product.length > 0 ? Product : null,
			PriceBookEntry: PriceBookEntry != null && PriceBookEntry.length > 0 ? PriceBookEntry : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a OrderItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOrderItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Order, Product, PriceBookEntry, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/OrderItem/update/' + id;
		const obj = {
				      		quantity: quantity,
      		unitPrice: unitPrice,
      		discountAmount: discountAmount,
      		taxAmount: taxAmount,
      		totalAmount: totalAmount,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		Product: Product != null && Product.length > 0 ? Product : null,
			PriceBookEntry: PriceBookEntry != null && PriceBookEntry.length > 0 ? PriceBookEntry : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a OrderItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOrderItem(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/OrderItem/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a OrderItem
	// returns the results untouched as an Observable OrderItem
	// OrderItem model
	// delegates via URI
	//********************************************************************
	getOrderItem(id) : Observable<OrderItem> {
		const uri_ = this.apiUrl + '/OrderItem/load/' + id;

		return this.http.get<OrderItem>(uri_);
	}
	
	//********************************************************************
	// gets all OrderItem
	// returns the results untouched as JSON representation of an
	// Observable array of OrderItem models
	// delegates via URI
	//********************************************************************
	getOrderItems() : Observable<OrderItem[]> {
		const uri_ = this.apiUrl + '/OrderItem/';

		return this
			.http.get<OrderItem[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Order on a OrderItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrder( orderItemId, _orderId ): Observable<any> {

		// get the OrderItem from storage
		this.loadHelper( orderItemId );

	// get the Order from storage
	var tmp 	= new OrderService(this.http).getOrder(_orderId);

	// assign the Order
	this.orderItem.order = tmp;

	// save the OrderItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Order on a OrderItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrder( orderItemId ): Observable<any> {

		// get the OrderItem from storage
		this.loadHelper( orderItemId );

	// assign Order to null
	this.orderItem.order = null;

	// save the OrderItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Product on a OrderItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProduct( orderItemId, _productId ): Observable<any> {

		// get the OrderItem from storage
		this.loadHelper( orderItemId );

	// get the Product from storage
	var tmp 	= new ProductService(this.http).getProduct(_productId);

	// assign the Product
	this.orderItem.product = tmp;

	// save the OrderItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Product on a OrderItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProduct( orderItemId ): Observable<any> {

		// get the OrderItem from storage
		this.loadHelper( orderItemId );

	// assign Product to null
	this.orderItem.product = null;

	// save the OrderItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PriceBookEntry on a OrderItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPriceBookEntry( orderItemId, _priceBookEntryId ): Observable<any> {

		// get the OrderItem from storage
		this.loadHelper( orderItemId );

	// get the PriceBookEntry from storage
	var tmp 	= new PriceBookEntryService(this.http).getPriceBookEntry(_priceBookEntryId);

	// assign the PriceBookEntry
	this.orderItem.priceBookEntry = tmp;

	// save the OrderItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PriceBookEntry on a OrderItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPriceBookEntry( orderItemId ): Observable<any> {

		// get the OrderItem from storage
		this.loadHelper( orderItemId );

	// assign PriceBookEntry to null
	this.orderItem.priceBookEntry = null;

	// save the OrderItem
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a OrderItem
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/OrderItem/update/' + this.orderItem;

	return  this.http.post(uri_, this.orderItem );
}

	//********************************************************************
	// loadHelper - internal helper to load a OrderItem
	//********************************************************************	
	loadHelper( id ) {
		this.getOrderItem(id)
			.subscribe((res : OrderItem) => {
				this.orderItem = res;
			});
	}
}