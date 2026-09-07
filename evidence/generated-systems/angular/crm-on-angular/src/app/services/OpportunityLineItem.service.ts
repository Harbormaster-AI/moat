import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {OpportunityLineItem} from '../models/OpportunityLineItem';
import {OpportunityService} from '../services/Opportunity.service';
import {ProductService} from '../services/Product.service';
import {PriceBookEntryService} from '../services/PriceBookEntry.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OpportunityLineItemService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	opportunityLineItem : OpportunityLineItem;

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
	// add a OpportunityLineItem
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOpportunityLineItem(quantity, unitPrice, discountPercent, totalPrice, Opportunity, Product, PriceBookEntry) : Observable<any> {
		const uri_ = this.apiUrl + '/OpportunityLineItem/create';
		const obj = {
			      		quantity: quantity,
      		unitPrice: unitPrice,
      		discountPercent: discountPercent,
      		totalPrice: totalPrice,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Product: Product != null && Product.length > 0 ? Product : null,
			PriceBookEntry: PriceBookEntry != null && PriceBookEntry.length > 0 ? PriceBookEntry : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a OpportunityLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOpportunityLineItem(quantity, unitPrice, discountPercent, totalPrice, Opportunity, Product, PriceBookEntry, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/OpportunityLineItem/update/' + id;
		const obj = {
				      		quantity: quantity,
      		unitPrice: unitPrice,
      		discountPercent: discountPercent,
      		totalPrice: totalPrice,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Product: Product != null && Product.length > 0 ? Product : null,
			PriceBookEntry: PriceBookEntry != null && PriceBookEntry.length > 0 ? PriceBookEntry : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a OpportunityLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOpportunityLineItem(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/OpportunityLineItem/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a OpportunityLineItem
	// returns the results untouched as an Observable OpportunityLineItem
	// OpportunityLineItem model
	// delegates via URI
	//********************************************************************
	getOpportunityLineItem(id) : Observable<OpportunityLineItem> {
		const uri_ = this.apiUrl + '/OpportunityLineItem/load/' + id;

		return this.http.get<OpportunityLineItem>(uri_);
	}
	
	//********************************************************************
	// gets all OpportunityLineItem
	// returns the results untouched as JSON representation of an
	// Observable array of OpportunityLineItem models
	// delegates via URI
	//********************************************************************
	getOpportunityLineItems() : Observable<OpportunityLineItem[]> {
		const uri_ = this.apiUrl + '/OpportunityLineItem/';

		return this
			.http.get<OpportunityLineItem[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Opportunity on a OpportunityLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOpportunity( opportunityLineItemId, _opportunityId ): Observable<any> {

		// get the OpportunityLineItem from storage
		this.loadHelper( opportunityLineItemId );

	// get the Opportunity from storage
	var tmp 	= new OpportunityService(this.http).getOpportunity(_opportunityId);

	// assign the Opportunity
	this.opportunityLineItem.opportunity = tmp;

	// save the OpportunityLineItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Opportunity on a OpportunityLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOpportunity( opportunityLineItemId ): Observable<any> {

		// get the OpportunityLineItem from storage
		this.loadHelper( opportunityLineItemId );

	// assign Opportunity to null
	this.opportunityLineItem.opportunity = null;

	// save the OpportunityLineItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Product on a OpportunityLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProduct( opportunityLineItemId, _productId ): Observable<any> {

		// get the OpportunityLineItem from storage
		this.loadHelper( opportunityLineItemId );

	// get the Product from storage
	var tmp 	= new ProductService(this.http).getProduct(_productId);

	// assign the Product
	this.opportunityLineItem.product = tmp;

	// save the OpportunityLineItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Product on a OpportunityLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProduct( opportunityLineItemId ): Observable<any> {

		// get the OpportunityLineItem from storage
		this.loadHelper( opportunityLineItemId );

	// assign Product to null
	this.opportunityLineItem.product = null;

	// save the OpportunityLineItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PriceBookEntry on a OpportunityLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPriceBookEntry( opportunityLineItemId, _priceBookEntryId ): Observable<any> {

		// get the OpportunityLineItem from storage
		this.loadHelper( opportunityLineItemId );

	// get the PriceBookEntry from storage
	var tmp 	= new PriceBookEntryService(this.http).getPriceBookEntry(_priceBookEntryId);

	// assign the PriceBookEntry
	this.opportunityLineItem.priceBookEntry = tmp;

	// save the OpportunityLineItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PriceBookEntry on a OpportunityLineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPriceBookEntry( opportunityLineItemId ): Observable<any> {

		// get the OpportunityLineItem from storage
		this.loadHelper( opportunityLineItemId );

	// assign PriceBookEntry to null
	this.opportunityLineItem.priceBookEntry = null;

	// save the OpportunityLineItem
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a OpportunityLineItem
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/OpportunityLineItem/update/' + this.opportunityLineItem;

	return  this.http.post(uri_, this.opportunityLineItem );
}

	//********************************************************************
	// loadHelper - internal helper to load a OpportunityLineItem
	//********************************************************************	
	loadHelper( id ) {
		this.getOpportunityLineItem(id)
			.subscribe((res : OpportunityLineItem) => {
				this.opportunityLineItem = res;
			});
	}
}