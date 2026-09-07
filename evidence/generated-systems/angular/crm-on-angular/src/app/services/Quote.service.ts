import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Quote} from '../models/Quote';
import {OrganizationService} from '../services/Organization.service';
import {AccountService} from '../services/Account.service';
import {OpportunityService} from '../services/Opportunity.service';
import {UserService} from '../services/User.service';
import {QuoteLineItemService} from '../services/QuoteLineItem.service';
import {PriceBookService} from '../services/PriceBook.service';
import {OrderService} from '../services/Order.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class QuoteService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	quote : Quote;

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
	// add a Quote
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addQuote(quoteNumber, validityStart, validityEnd, totalAmount, discountPercent, taxAmount, shippingAmount, Organization, Account, Opportunity, Owner, LineItems, PriceBook, Order, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Quote/create';
		const obj = {
			      		quoteNumber: quoteNumber,
      		validityStart: validityStart,
      		validityEnd: validityEnd,
      		totalAmount: totalAmount,
      		discountPercent: discountPercent,
      		taxAmount: taxAmount,
      		shippingAmount: shippingAmount,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		LineItems: LineItems != null && LineItems.length > 0 ? LineItems : null,
      		PriceBook: PriceBook != null && PriceBook.length > 0 ? PriceBook : null,
      		Order: Order != null && Order.length > 0 ? Order : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateQuote(quoteNumber, validityStart, validityEnd, totalAmount, discountPercent, taxAmount, shippingAmount, Organization, Account, Opportunity, Owner, LineItems, PriceBook, Order, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Quote/update/' + id;
		const obj = {
				      		quoteNumber: quoteNumber,
      		validityStart: validityStart,
      		validityEnd: validityEnd,
      		totalAmount: totalAmount,
      		discountPercent: discountPercent,
      		taxAmount: taxAmount,
      		shippingAmount: shippingAmount,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		LineItems: LineItems != null && LineItems.length > 0 ? LineItems : null,
      		PriceBook: PriceBook != null && PriceBook.length > 0 ? PriceBook : null,
      		Order: Order != null && Order.length > 0 ? Order : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteQuote(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Quote/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Quote
	// returns the results untouched as an Observable Quote
	// Quote model
	// delegates via URI
	//********************************************************************
	getQuote(id) : Observable<Quote> {
		const uri_ = this.apiUrl + '/Quote/load/' + id;

		return this.http.get<Quote>(uri_);
	}
	
	//********************************************************************
	// gets all Quote
	// returns the results untouched as JSON representation of an
	// Observable array of Quote models
	// delegates via URI
	//********************************************************************
	getQuotes() : Observable<Quote[]> {
		const uri_ = this.apiUrl + '/Quote/';

		return this
			.http.get<Quote[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( quoteId, _organizationId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.quote.organization = tmp;

	// save the Quote
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( quoteId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// assign Organization to null
	this.quote.organization = null;

	// save the Quote
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Account on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( quoteId, _accountId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.quote.account = tmp;

	// save the Quote
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( quoteId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// assign Account to null
	this.quote.account = null;

	// save the Quote
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Opportunity on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOpportunity( quoteId, _opportunityId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// get the Opportunity from storage
	var tmp 	= new OpportunityService(this.http).getOpportunity(_opportunityId);

	// assign the Opportunity
	this.quote.opportunity = tmp;

	// save the Quote
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Opportunity on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOpportunity( quoteId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// assign Opportunity to null
	this.quote.opportunity = null;

	// save the Quote
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( quoteId, _ownerId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.quote.owner = tmp;

	// save the Quote
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( quoteId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// assign Owner to null
	this.quote.owner = null;

	// save the Quote
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PriceBook on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPriceBook( quoteId, _priceBookId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// get the PriceBook from storage
	var tmp 	= new PriceBookService(this.http).getPriceBook(_priceBookId);

	// assign the PriceBook
	this.quote.priceBook = tmp;

	// save the Quote
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PriceBook on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPriceBook( quoteId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// assign PriceBook to null
	this.quote.priceBook = null;

	// save the Quote
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Order on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrder( quoteId, _orderId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// get the Order from storage
	var tmp 	= new OrderService(this.http).getOrder(_orderId);

	// assign the Order
	this.quote.order = tmp;

	// save the Quote
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Order on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrder( quoteId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// assign Order to null
	this.quote.order = null;

	// save the Quote
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more lineItemsIds as a LineItems
	// to a Quote
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLineItems( quoteId, lineItemsIds ): Observable<any> {

		// get the Quote
		this.loadHelper( quoteId );

	// split on a comma with no spaces
	var idList = lineItemsIds.split(',')

	// iterate over array of lineItems ids
	idList.forEach(function (id) {
		// read the QuoteLineItem
		var quoteLineItem = new QuoteLineItemService(this.http).getQuoteLineItem(id);
		// add the QuoteLineItem if not already assigned
		if ( this.quote.lineItems.indexOf(quoteLineItem) == -1 )
		this.quote.lineItems.push(quoteLineItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more lineItemsIds as a LineItems
	// from a Quote
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLineItems( quoteId, lineItemsIds ): Observable<any> {

		// get the Quote
		this.loadHelper( quoteId );


	// split on a comma with no spaces
	var idList 					= lineItemsIds.split(',');
	var lineItems 	= this.quote.lineItems;

	if ( lineItems != null && lineItemsIds != null ) {

		// iterate over array of lineItems ids
		lineItems.forEach(function (obj) {
			if ( lineItemsIds.indexOf(obj._id) > -1 ) {
				// remove the QuoteLineItem
				this.quote.lineItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Quote
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Quote/update/' + this.quote;

	return  this.http.post(uri_, this.quote );
}

	//********************************************************************
	// loadHelper - internal helper to load a Quote
	//********************************************************************	
	loadHelper( id ) {
		this.getQuote(id)
			.subscribe((res : Quote) => {
				this.quote = res;
			});
	}
}