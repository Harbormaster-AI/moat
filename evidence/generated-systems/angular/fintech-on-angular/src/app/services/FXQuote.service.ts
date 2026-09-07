import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {FXQuote} from '../models/FXQuote';
import {CustomerService} from '../services/Customer.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class FXQuoteService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	fXQuote : FXQuote;

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
	// add a FXQuote
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addFXQuote(baseCurrency, quoteCurrency, rate, quotedAt, expiresAt, RequestedBy, PriceType) : Observable<any> {
		const uri_ = this.apiUrl + '/FXQuote/create';
		const obj = {
			      		baseCurrency: baseCurrency,
      		quoteCurrency: quoteCurrency,
      		rate: rate,
      		quotedAt: quotedAt,
      		expiresAt: expiresAt,
      		RequestedBy: RequestedBy != null && RequestedBy.length > 0 ? RequestedBy : null,
			PriceType: PriceType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a FXQuote
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateFXQuote(baseCurrency, quoteCurrency, rate, quotedAt, expiresAt, RequestedBy, PriceType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/FXQuote/update/' + id;
		const obj = {
				      		baseCurrency: baseCurrency,
      		quoteCurrency: quoteCurrency,
      		rate: rate,
      		quotedAt: quotedAt,
      		expiresAt: expiresAt,
      		RequestedBy: RequestedBy != null && RequestedBy.length > 0 ? RequestedBy : null,
			PriceType: PriceType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a FXQuote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteFXQuote(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/FXQuote/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a FXQuote
	// returns the results untouched as an Observable FXQuote
	// FXQuote model
	// delegates via URI
	//********************************************************************
	getFXQuote(id) : Observable<FXQuote> {
		const uri_ = this.apiUrl + '/FXQuote/load/' + id;

		return this.http.get<FXQuote>(uri_);
	}
	
	//********************************************************************
	// gets all FXQuote
	// returns the results untouched as JSON representation of an
	// Observable array of FXQuote models
	// delegates via URI
	//********************************************************************
	getFXQuotes() : Observable<FXQuote[]> {
		const uri_ = this.apiUrl + '/FXQuote/';

		return this
			.http.get<FXQuote[]>(uri_);
	}
	
			//********************************************************************
	// assigns a RequestedBy on a FXQuote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRequestedBy( fXQuoteId, _requestedById ): Observable<any> {

		// get the FXQuote from storage
		this.loadHelper( fXQuoteId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_requestedById);

	// assign the RequestedBy
	this.fXQuote.requestedBy = tmp;

	// save the FXQuote
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a RequestedBy on a FXQuote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRequestedBy( fXQuoteId ): Observable<any> {

		// get the FXQuote from storage
		this.loadHelper( fXQuoteId );

	// assign RequestedBy to null
	this.fXQuote.requestedBy = null;

	// save the FXQuote
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a FXQuote
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/FXQuote/update/' + this.fXQuote;

	return  this.http.post(uri_, this.fXQuote );
}

	//********************************************************************
	// loadHelper - internal helper to load a FXQuote
	//********************************************************************	
	loadHelper( id ) {
		this.getFXQuote(id)
			.subscribe((res : FXQuote) => {
				this.fXQuote = res;
			});
	}
}