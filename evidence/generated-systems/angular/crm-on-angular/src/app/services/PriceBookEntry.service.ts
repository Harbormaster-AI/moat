import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PriceBookEntry} from '../models/PriceBookEntry';
import {PriceBookService} from '../services/PriceBook.service';
import {ProductService} from '../services/Product.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PriceBookEntryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	priceBookEntry : PriceBookEntry;

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
	// add a PriceBookEntry
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPriceBookEntry(unitPrice, effectiveDate, expirationDate, asActive, PriceBook, Product) : Observable<any> {
		const uri_ = this.apiUrl + '/PriceBookEntry/create';
		const obj = {
			      		unitPrice: unitPrice,
      		effectiveDate: effectiveDate,
      		expirationDate: expirationDate,
      		asActive: asActive,
      		PriceBook: PriceBook != null && PriceBook.length > 0 ? PriceBook : null,
			Product: Product != null && Product.length > 0 ? Product : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PriceBookEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePriceBookEntry(unitPrice, effectiveDate, expirationDate, asActive, PriceBook, Product, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PriceBookEntry/update/' + id;
		const obj = {
				      		unitPrice: unitPrice,
      		effectiveDate: effectiveDate,
      		expirationDate: expirationDate,
      		asActive: asActive,
      		PriceBook: PriceBook != null && PriceBook.length > 0 ? PriceBook : null,
			Product: Product != null && Product.length > 0 ? Product : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PriceBookEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePriceBookEntry(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PriceBookEntry/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PriceBookEntry
	// returns the results untouched as an Observable PriceBookEntry
	// PriceBookEntry model
	// delegates via URI
	//********************************************************************
	getPriceBookEntry(id) : Observable<PriceBookEntry> {
		const uri_ = this.apiUrl + '/PriceBookEntry/load/' + id;

		return this.http.get<PriceBookEntry>(uri_);
	}
	
	//********************************************************************
	// gets all PriceBookEntry
	// returns the results untouched as JSON representation of an
	// Observable array of PriceBookEntry models
	// delegates via URI
	//********************************************************************
	getPriceBookEntrys() : Observable<PriceBookEntry[]> {
		const uri_ = this.apiUrl + '/PriceBookEntry/';

		return this
			.http.get<PriceBookEntry[]>(uri_);
	}
	
			//********************************************************************
	// assigns a PriceBook on a PriceBookEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPriceBook( priceBookEntryId, _priceBookId ): Observable<any> {

		// get the PriceBookEntry from storage
		this.loadHelper( priceBookEntryId );

	// get the PriceBook from storage
	var tmp 	= new PriceBookService(this.http).getPriceBook(_priceBookId);

	// assign the PriceBook
	this.priceBookEntry.priceBook = tmp;

	// save the PriceBookEntry
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PriceBook on a PriceBookEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPriceBook( priceBookEntryId ): Observable<any> {

		// get the PriceBookEntry from storage
		this.loadHelper( priceBookEntryId );

	// assign PriceBook to null
	this.priceBookEntry.priceBook = null;

	// save the PriceBookEntry
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Product on a PriceBookEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProduct( priceBookEntryId, _productId ): Observable<any> {

		// get the PriceBookEntry from storage
		this.loadHelper( priceBookEntryId );

	// get the Product from storage
	var tmp 	= new ProductService(this.http).getProduct(_productId);

	// assign the Product
	this.priceBookEntry.product = tmp;

	// save the PriceBookEntry
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Product on a PriceBookEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProduct( priceBookEntryId ): Observable<any> {

		// get the PriceBookEntry from storage
		this.loadHelper( priceBookEntryId );

	// assign Product to null
	this.priceBookEntry.product = null;

	// save the PriceBookEntry
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a PriceBookEntry
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PriceBookEntry/update/' + this.priceBookEntry;

	return  this.http.post(uri_, this.priceBookEntry );
}

	//********************************************************************
	// loadHelper - internal helper to load a PriceBookEntry
	//********************************************************************	
	loadHelper( id ) {
		this.getPriceBookEntry(id)
			.subscribe((res : PriceBookEntry) => {
				this.priceBookEntry = res;
			});
	}
}