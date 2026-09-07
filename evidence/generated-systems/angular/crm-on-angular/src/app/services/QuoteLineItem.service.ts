import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {QuoteLineItem} from '../models/QuoteLineItem';
import {QuoteService} from '../services/Quote.service';
import {ProductService} from '../services/Product.service';
import {PriceBookEntryService} from '../services/PriceBookEntry.service';
import {OpportunityLineItemService} from '../services/OpportunityLineItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class QuoteLineItemService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	quoteLineItem : QuoteLineItem;

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
	// add a QuoteLineItem
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addQuoteLineItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Quote, Product, PriceBookEntry, OpportunityLineItem) : Observable<any> {
		const uri_ = this.apiUrl + '/QuoteLineItem/create';
		const obj = {
			      		quantity: quantity,
      		unitPrice: unitPrice,
      		discountAmount: discountAmount,
      		taxAmount: taxAmount,
      		totalAmount: totalAmount,
      		Quote: Quote != null && Quote.length > 0 ? Quote : null,
      		Product: Product != null && Product.length > 0 ? Product : null,
      		PriceBookEntry: PriceBookEntry != null && PriceBookEntry.length > 0 ? PriceBookEntry : null,
			OpportunityLineItem: OpportunityLineItem != null && OpportunityLineItem.length > 0 ? OpportunityLineItem : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a QuoteLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateQuoteLineItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Quote, Product, PriceBookEntry, OpportunityLineItem, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/QuoteLineItem/update/' + id;
		const obj = {
				      		quantity: quantity,
      		unitPrice: unitPrice,
      		discountAmount: discountAmount,
      		taxAmount: taxAmount,
      		totalAmount: totalAmount,
      		Quote: Quote != null && Quote.length > 0 ? Quote : null,
      		Product: Product != null && Product.length > 0 ? Product : null,
      		PriceBookEntry: PriceBookEntry != null && PriceBookEntry.length > 0 ? PriceBookEntry : null,
			OpportunityLineItem: OpportunityLineItem != null && OpportunityLineItem.length > 0 ? OpportunityLineItem : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a QuoteLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteQuoteLineItem(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/QuoteLineItem/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a QuoteLineItem
	// returns the results untouched as an Observable QuoteLineItem
	// QuoteLineItem model
	// delegates via URI
	//********************************************************************
	getQuoteLineItem(id) : Observable<QuoteLineItem> {
		const uri_ = this.apiUrl + '/QuoteLineItem/load/' + id;

		return this.http.get<QuoteLineItem>(uri_);
	}
	
	//********************************************************************
	// gets all QuoteLineItem
	// returns the results untouched as JSON representation of an
	// Observable array of QuoteLineItem models
	// delegates via URI
	//********************************************************************
	getQuoteLineItems() : Observable<QuoteLineItem[]> {
		const uri_ = this.apiUrl + '/QuoteLineItem/';

		return this
			.http.get<QuoteLineItem[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Quote on a QuoteLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignQuote( quoteLineItemId, _quoteId ): Observable<any> {

		// get the QuoteLineItem from storage
		this.loadHelper( quoteLineItemId );

	// get the Quote from storage
	var tmp 	= new QuoteService(this.http).getQuote(_quoteId);

	// assign the Quote
	this.quoteLineItem.quote = tmp;

	// save the QuoteLineItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Quote on a QuoteLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignQuote( quoteLineItemId ): Observable<any> {

		// get the QuoteLineItem from storage
		this.loadHelper( quoteLineItemId );

	// assign Quote to null
	this.quoteLineItem.quote = null;

	// save the QuoteLineItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Product on a QuoteLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProduct( quoteLineItemId, _productId ): Observable<any> {

		// get the QuoteLineItem from storage
		this.loadHelper( quoteLineItemId );

	// get the Product from storage
	var tmp 	= new ProductService(this.http).getProduct(_productId);

	// assign the Product
	this.quoteLineItem.product = tmp;

	// save the QuoteLineItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Product on a QuoteLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProduct( quoteLineItemId ): Observable<any> {

		// get the QuoteLineItem from storage
		this.loadHelper( quoteLineItemId );

	// assign Product to null
	this.quoteLineItem.product = null;

	// save the QuoteLineItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PriceBookEntry on a QuoteLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPriceBookEntry( quoteLineItemId, _priceBookEntryId ): Observable<any> {

		// get the QuoteLineItem from storage
		this.loadHelper( quoteLineItemId );

	// get the PriceBookEntry from storage
	var tmp 	= new PriceBookEntryService(this.http).getPriceBookEntry(_priceBookEntryId);

	// assign the PriceBookEntry
	this.quoteLineItem.priceBookEntry = tmp;

	// save the QuoteLineItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PriceBookEntry on a QuoteLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPriceBookEntry( quoteLineItemId ): Observable<any> {

		// get the QuoteLineItem from storage
		this.loadHelper( quoteLineItemId );

	// assign PriceBookEntry to null
	this.quoteLineItem.priceBookEntry = null;

	// save the QuoteLineItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a OpportunityLineItem on a QuoteLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOpportunityLineItem( quoteLineItemId, _opportunityLineItemId ): Observable<any> {

		// get the QuoteLineItem from storage
		this.loadHelper( quoteLineItemId );

	// get the OpportunityLineItem from storage
	var tmp 	= new OpportunityLineItemService(this.http).getOpportunityLineItem(_opportunityLineItemId);

	// assign the OpportunityLineItem
	this.quoteLineItem.opportunityLineItem = tmp;

	// save the QuoteLineItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a OpportunityLineItem on a QuoteLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOpportunityLineItem( quoteLineItemId ): Observable<any> {

		// get the QuoteLineItem from storage
		this.loadHelper( quoteLineItemId );

	// assign OpportunityLineItem to null
	this.quoteLineItem.opportunityLineItem = null;

	// save the QuoteLineItem
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a QuoteLineItem
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/QuoteLineItem/update/' + this.quoteLineItem;

	return  this.http.post(uri_, this.quoteLineItem );
}

	//********************************************************************
	// loadHelper - internal helper to load a QuoteLineItem
	//********************************************************************	
	loadHelper( id ) {
		this.getQuoteLineItem(id)
			.subscribe((res : QuoteLineItem) => {
				this.quoteLineItem = res;
			});
	}
}