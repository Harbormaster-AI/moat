import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PriceBook} from '../models/PriceBook';
import {OrganizationService} from '../services/Organization.service';
import {PriceBookEntryService} from '../services/PriceBookEntry.service';
import {QuoteService} from '../services/Quote.service';
import {OrderService} from '../services/Order.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PriceBookService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	priceBook : PriceBook;

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
	// add a PriceBook
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPriceBook(name, asActive, description, Organization, Entries, Quotes, Orders) : Observable<any> {
		const uri_ = this.apiUrl + '/PriceBook/create';
		const obj = {
			      		name: name,
      		asActive: asActive,
      		description: description,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Entries: Entries != null && Entries.length > 0 ? Entries : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
			Orders: Orders != null && Orders.length > 0 ? Orders : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PriceBook
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePriceBook(name, asActive, description, Organization, Entries, Quotes, Orders, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PriceBook/update/' + id;
		const obj = {
				      		name: name,
      		asActive: asActive,
      		description: description,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Entries: Entries != null && Entries.length > 0 ? Entries : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
			Orders: Orders != null && Orders.length > 0 ? Orders : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PriceBook
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePriceBook(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PriceBook/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PriceBook
	// returns the results untouched as an Observable PriceBook
	// PriceBook model
	// delegates via URI
	//********************************************************************
	getPriceBook(id) : Observable<PriceBook> {
		const uri_ = this.apiUrl + '/PriceBook/load/' + id;

		return this.http.get<PriceBook>(uri_);
	}
	
	//********************************************************************
	// gets all PriceBook
	// returns the results untouched as JSON representation of an
	// Observable array of PriceBook models
	// delegates via URI
	//********************************************************************
	getPriceBooks() : Observable<PriceBook[]> {
		const uri_ = this.apiUrl + '/PriceBook/';

		return this
			.http.get<PriceBook[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a PriceBook
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( priceBookId, _organizationId ): Observable<any> {

		// get the PriceBook from storage
		this.loadHelper( priceBookId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.priceBook.organization = tmp;

	// save the PriceBook
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a PriceBook
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( priceBookId ): Observable<any> {

		// get the PriceBook from storage
		this.loadHelper( priceBookId );

	// assign Organization to null
	this.priceBook.organization = null;

	// save the PriceBook
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more entriesIds as a Entries
	// to a PriceBook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEntries( priceBookId, entriesIds ): Observable<any> {

		// get the PriceBook
		this.loadHelper( priceBookId );

	// split on a comma with no spaces
	var idList = entriesIds.split(',')

	// iterate over array of entries ids
	idList.forEach(function (id) {
		// read the PriceBookEntry
		var priceBookEntry = new PriceBookEntryService(this.http).getPriceBookEntry(id);
		// add the PriceBookEntry if not already assigned
		if ( this.priceBook.entries.indexOf(priceBookEntry) == -1 )
		this.priceBook.entries.push(priceBookEntry);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more entriesIds as a Entries
	// from a PriceBook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEntries( priceBookId, entriesIds ): Observable<any> {

		// get the PriceBook
		this.loadHelper( priceBookId );


	// split on a comma with no spaces
	var idList 					= entriesIds.split(',');
	var entries 	= this.priceBook.entries;

	if ( entries != null && entriesIds != null ) {

		// iterate over array of entries ids
		entries.forEach(function (obj) {
			if ( entriesIds.indexOf(obj._id) > -1 ) {
				// remove the PriceBookEntry
				this.priceBook.entries.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more quotesIds as a Quotes
	// to a PriceBook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQuotes( priceBookId, quotesIds ): Observable<any> {

		// get the PriceBook
		this.loadHelper( priceBookId );

	// split on a comma with no spaces
	var idList = quotesIds.split(',')

	// iterate over array of quotes ids
	idList.forEach(function (id) {
		// read the Quote
		var quote = new QuoteService(this.http).getQuote(id);
		// add the Quote if not already assigned
		if ( this.priceBook.quotes.indexOf(quote) == -1 )
		this.priceBook.quotes.push(quote);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more quotesIds as a Quotes
	// from a PriceBook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQuotes( priceBookId, quotesIds ): Observable<any> {

		// get the PriceBook
		this.loadHelper( priceBookId );


	// split on a comma with no spaces
	var idList 					= quotesIds.split(',');
	var quotes 	= this.priceBook.quotes;

	if ( quotes != null && quotesIds != null ) {

		// iterate over array of quotes ids
		quotes.forEach(function (obj) {
			if ( quotesIds.indexOf(obj._id) > -1 ) {
				// remove the Quote
				this.priceBook.quotes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ordersIds as a Orders
	// to a PriceBook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOrders( priceBookId, ordersIds ): Observable<any> {

		// get the PriceBook
		this.loadHelper( priceBookId );

	// split on a comma with no spaces
	var idList = ordersIds.split(',')

	// iterate over array of orders ids
	idList.forEach(function (id) {
		// read the Order
		var order = new OrderService(this.http).getOrder(id);
		// add the Order if not already assigned
		if ( this.priceBook.orders.indexOf(order) == -1 )
		this.priceBook.orders.push(order);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ordersIds as a Orders
	// from a PriceBook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOrders( priceBookId, ordersIds ): Observable<any> {

		// get the PriceBook
		this.loadHelper( priceBookId );


	// split on a comma with no spaces
	var idList 					= ordersIds.split(',');
	var orders 	= this.priceBook.orders;

	if ( orders != null && ordersIds != null ) {

		// iterate over array of orders ids
		orders.forEach(function (obj) {
			if ( ordersIds.indexOf(obj._id) > -1 ) {
				// remove the Order
				this.priceBook.orders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PriceBook
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PriceBook/update/' + this.priceBook;

	return  this.http.post(uri_, this.priceBook );
}

	//********************************************************************
	// loadHelper - internal helper to load a PriceBook
	//********************************************************************	
	loadHelper( id ) {
		this.getPriceBook(id)
			.subscribe((res : PriceBook) => {
				this.priceBook = res;
			});
	}
}