import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Product} from '../models/Product';
import {OrganizationService} from '../services/Organization.service';
import {PriceBookEntryService} from '../services/PriceBookEntry.service';
import {OpportunityLineItemService} from '../services/OpportunityLineItem.service';
import {QuoteLineItemService} from '../services/QuoteLineItem.service';
import {OrderItemService} from '../services/OrderItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ProductService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	product : Product;

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
	// add a Product
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addProduct(sku, name, asActive, standardPrice, description, Organization, PriceBookEntries, OpportunityLineItems, QuoteLineItems, OrderItems, ProductType, Uom) : Observable<any> {
		const uri_ = this.apiUrl + '/Product/create';
		const obj = {
			      		sku: sku,
      		name: name,
      		asActive: asActive,
      		standardPrice: standardPrice,
      		description: description,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		PriceBookEntries: PriceBookEntries != null && PriceBookEntries.length > 0 ? PriceBookEntries : null,
      		OpportunityLineItems: OpportunityLineItems != null && OpportunityLineItems.length > 0 ? OpportunityLineItems : null,
      		QuoteLineItems: QuoteLineItems != null && QuoteLineItems.length > 0 ? QuoteLineItems : null,
      		OrderItems: OrderItems != null && OrderItems.length > 0 ? OrderItems : null,
      		ProductType: ProductType,
			Uom: Uom
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Product
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateProduct(sku, name, asActive, standardPrice, description, Organization, PriceBookEntries, OpportunityLineItems, QuoteLineItems, OrderItems, ProductType, Uom, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Product/update/' + id;
		const obj = {
				      		sku: sku,
      		name: name,
      		asActive: asActive,
      		standardPrice: standardPrice,
      		description: description,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		PriceBookEntries: PriceBookEntries != null && PriceBookEntries.length > 0 ? PriceBookEntries : null,
      		OpportunityLineItems: OpportunityLineItems != null && OpportunityLineItems.length > 0 ? OpportunityLineItems : null,
      		QuoteLineItems: QuoteLineItems != null && QuoteLineItems.length > 0 ? QuoteLineItems : null,
      		OrderItems: OrderItems != null && OrderItems.length > 0 ? OrderItems : null,
      		ProductType: ProductType,
			Uom: Uom
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Product
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteProduct(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Product/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Product
	// returns the results untouched as an Observable Product
	// Product model
	// delegates via URI
	//********************************************************************
	getProduct(id) : Observable<Product> {
		const uri_ = this.apiUrl + '/Product/load/' + id;

		return this.http.get<Product>(uri_);
	}
	
	//********************************************************************
	// gets all Product
	// returns the results untouched as JSON representation of an
	// Observable array of Product models
	// delegates via URI
	//********************************************************************
	getProducts() : Observable<Product[]> {
		const uri_ = this.apiUrl + '/Product/';

		return this
			.http.get<Product[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Product
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( productId, _organizationId ): Observable<any> {

		// get the Product from storage
		this.loadHelper( productId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.product.organization = tmp;

	// save the Product
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Product
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( productId ): Observable<any> {

		// get the Product from storage
		this.loadHelper( productId );

	// assign Organization to null
	this.product.organization = null;

	// save the Product
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more priceBookEntriesIds as a PriceBookEntries
	// to a Product
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPriceBookEntries( productId, priceBookEntriesIds ): Observable<any> {

		// get the Product
		this.loadHelper( productId );

	// split on a comma with no spaces
	var idList = priceBookEntriesIds.split(',')

	// iterate over array of priceBookEntries ids
	idList.forEach(function (id) {
		// read the PriceBookEntry
		var priceBookEntry = new PriceBookEntryService(this.http).getPriceBookEntry(id);
		// add the PriceBookEntry if not already assigned
		if ( this.product.priceBookEntries.indexOf(priceBookEntry) == -1 )
		this.product.priceBookEntries.push(priceBookEntry);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more priceBookEntriesIds as a PriceBookEntries
	// from a Product
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePriceBookEntries( productId, priceBookEntriesIds ): Observable<any> {

		// get the Product
		this.loadHelper( productId );


	// split on a comma with no spaces
	var idList 					= priceBookEntriesIds.split(',');
	var priceBookEntries 	= this.product.priceBookEntries;

	if ( priceBookEntries != null && priceBookEntriesIds != null ) {

		// iterate over array of priceBookEntries ids
		priceBookEntries.forEach(function (obj) {
			if ( priceBookEntriesIds.indexOf(obj._id) > -1 ) {
				// remove the PriceBookEntry
				this.product.priceBookEntries.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more opportunityLineItemsIds as a OpportunityLineItems
	// to a Product
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOpportunityLineItems( productId, opportunityLineItemsIds ): Observable<any> {

		// get the Product
		this.loadHelper( productId );

	// split on a comma with no spaces
	var idList = opportunityLineItemsIds.split(',')

	// iterate over array of opportunityLineItems ids
	idList.forEach(function (id) {
		// read the OpportunityLineItem
		var opportunityLineItem = new OpportunityLineItemService(this.http).getOpportunityLineItem(id);
		// add the OpportunityLineItem if not already assigned
		if ( this.product.opportunityLineItems.indexOf(opportunityLineItem) == -1 )
		this.product.opportunityLineItems.push(opportunityLineItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more opportunityLineItemsIds as a OpportunityLineItems
	// from a Product
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOpportunityLineItems( productId, opportunityLineItemsIds ): Observable<any> {

		// get the Product
		this.loadHelper( productId );


	// split on a comma with no spaces
	var idList 					= opportunityLineItemsIds.split(',');
	var opportunityLineItems 	= this.product.opportunityLineItems;

	if ( opportunityLineItems != null && opportunityLineItemsIds != null ) {

		// iterate over array of opportunityLineItems ids
		opportunityLineItems.forEach(function (obj) {
			if ( opportunityLineItemsIds.indexOf(obj._id) > -1 ) {
				// remove the OpportunityLineItem
				this.product.opportunityLineItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more quoteLineItemsIds as a QuoteLineItems
	// to a Product
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQuoteLineItems( productId, quoteLineItemsIds ): Observable<any> {

		// get the Product
		this.loadHelper( productId );

	// split on a comma with no spaces
	var idList = quoteLineItemsIds.split(',')

	// iterate over array of quoteLineItems ids
	idList.forEach(function (id) {
		// read the QuoteLineItem
		var quoteLineItem = new QuoteLineItemService(this.http).getQuoteLineItem(id);
		// add the QuoteLineItem if not already assigned
		if ( this.product.quoteLineItems.indexOf(quoteLineItem) == -1 )
		this.product.quoteLineItems.push(quoteLineItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more quoteLineItemsIds as a QuoteLineItems
	// from a Product
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQuoteLineItems( productId, quoteLineItemsIds ): Observable<any> {

		// get the Product
		this.loadHelper( productId );


	// split on a comma with no spaces
	var idList 					= quoteLineItemsIds.split(',');
	var quoteLineItems 	= this.product.quoteLineItems;

	if ( quoteLineItems != null && quoteLineItemsIds != null ) {

		// iterate over array of quoteLineItems ids
		quoteLineItems.forEach(function (obj) {
			if ( quoteLineItemsIds.indexOf(obj._id) > -1 ) {
				// remove the QuoteLineItem
				this.product.quoteLineItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more orderItemsIds as a OrderItems
	// to a Product
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOrderItems( productId, orderItemsIds ): Observable<any> {

		// get the Product
		this.loadHelper( productId );

	// split on a comma with no spaces
	var idList = orderItemsIds.split(',')

	// iterate over array of orderItems ids
	idList.forEach(function (id) {
		// read the OrderItem
		var orderItem = new OrderItemService(this.http).getOrderItem(id);
		// add the OrderItem if not already assigned
		if ( this.product.orderItems.indexOf(orderItem) == -1 )
		this.product.orderItems.push(orderItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more orderItemsIds as a OrderItems
	// from a Product
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOrderItems( productId, orderItemsIds ): Observable<any> {

		// get the Product
		this.loadHelper( productId );


	// split on a comma with no spaces
	var idList 					= orderItemsIds.split(',');
	var orderItems 	= this.product.orderItems;

	if ( orderItems != null && orderItemsIds != null ) {

		// iterate over array of orderItems ids
		orderItems.forEach(function (obj) {
			if ( orderItemsIds.indexOf(obj._id) > -1 ) {
				// remove the OrderItem
				this.product.orderItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Product
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Product/update/' + this.product;

	return  this.http.post(uri_, this.product );
}

	//********************************************************************
	// loadHelper - internal helper to load a Product
	//********************************************************************	
	loadHelper( id ) {
		this.getProduct(id)
			.subscribe((res : Product) => {
				this.product = res;
			});
	}
}