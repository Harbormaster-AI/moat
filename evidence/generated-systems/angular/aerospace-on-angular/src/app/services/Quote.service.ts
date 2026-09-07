import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Quote} from '../models/Quote';
import {AircraftOrderService} from '../services/AircraftOrder.service';
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
	addQuote(quoteNumber, totalAmount, AircraftOrder) : Observable<any> {
		const uri_ = this.apiUrl + '/Quote/create';
		const obj = {
			      		quoteNumber: quoteNumber,
      		totalAmount: totalAmount,
			AircraftOrder: AircraftOrder != null && AircraftOrder.length > 0 ? AircraftOrder : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateQuote(quoteNumber, totalAmount, AircraftOrder, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Quote/update/' + id;
		const obj = {
				      		quoteNumber: quoteNumber,
      		totalAmount: totalAmount,
			AircraftOrder: AircraftOrder != null && AircraftOrder.length > 0 ? AircraftOrder : null
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
	// assigns a AircraftOrder on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAircraftOrder( quoteId, _aircraftOrderId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// get the AircraftOrder from storage
	var tmp 	= new AircraftOrderService(this.http).getAircraftOrder(_aircraftOrderId);

	// assign the AircraftOrder
	this.quote.aircraftOrder = tmp;

	// save the Quote
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AircraftOrder on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAircraftOrder( quoteId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// assign AircraftOrder to null
	this.quote.aircraftOrder = null;

	// save the Quote
	return this.saveHelper();
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