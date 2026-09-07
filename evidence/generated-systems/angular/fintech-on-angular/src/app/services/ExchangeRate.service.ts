import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ExchangeRate} from '../models/ExchangeRate';
import {FXQuoteService} from '../services/FXQuote.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ExchangeRateService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	exchangeRate : ExchangeRate;

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
	// add a ExchangeRate
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addExchangeRate(baseCurrency, quoteCurrency, rate, asOf, source, UsedByQuotes) : Observable<any> {
		const uri_ = this.apiUrl + '/ExchangeRate/create';
		const obj = {
			      		baseCurrency: baseCurrency,
      		quoteCurrency: quoteCurrency,
      		rate: rate,
      		asOf: asOf,
      		source: source,
			UsedByQuotes: UsedByQuotes != null && UsedByQuotes.length > 0 ? UsedByQuotes : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ExchangeRate
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateExchangeRate(baseCurrency, quoteCurrency, rate, asOf, source, UsedByQuotes, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ExchangeRate/update/' + id;
		const obj = {
				      		baseCurrency: baseCurrency,
      		quoteCurrency: quoteCurrency,
      		rate: rate,
      		asOf: asOf,
      		source: source,
			UsedByQuotes: UsedByQuotes != null && UsedByQuotes.length > 0 ? UsedByQuotes : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ExchangeRate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteExchangeRate(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ExchangeRate/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ExchangeRate
	// returns the results untouched as an Observable ExchangeRate
	// ExchangeRate model
	// delegates via URI
	//********************************************************************
	getExchangeRate(id) : Observable<ExchangeRate> {
		const uri_ = this.apiUrl + '/ExchangeRate/load/' + id;

		return this.http.get<ExchangeRate>(uri_);
	}
	
	//********************************************************************
	// gets all ExchangeRate
	// returns the results untouched as JSON representation of an
	// Observable array of ExchangeRate models
	// delegates via URI
	//********************************************************************
	getExchangeRates() : Observable<ExchangeRate[]> {
		const uri_ = this.apiUrl + '/ExchangeRate/';

		return this
			.http.get<ExchangeRate[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more usedByQuotesIds as a UsedByQuotes
	// to a ExchangeRate
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addUsedByQuotes( exchangeRateId, usedByQuotesIds ): Observable<any> {

		// get the ExchangeRate
		this.loadHelper( exchangeRateId );

	// split on a comma with no spaces
	var idList = usedByQuotesIds.split(',')

	// iterate over array of usedByQuotes ids
	idList.forEach(function (id) {
		// read the FXQuote
		var fXQuote = new FXQuoteService(this.http).getFXQuote(id);
		// add the FXQuote if not already assigned
		if ( this.exchangeRate.usedByQuotes.indexOf(fXQuote) == -1 )
		this.exchangeRate.usedByQuotes.push(fXQuote);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more usedByQuotesIds as a UsedByQuotes
	// from a ExchangeRate
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeUsedByQuotes( exchangeRateId, usedByQuotesIds ): Observable<any> {

		// get the ExchangeRate
		this.loadHelper( exchangeRateId );


	// split on a comma with no spaces
	var idList 					= usedByQuotesIds.split(',');
	var usedByQuotes 	= this.exchangeRate.usedByQuotes;

	if ( usedByQuotes != null && usedByQuotesIds != null ) {

		// iterate over array of usedByQuotes ids
		usedByQuotes.forEach(function (obj) {
			if ( usedByQuotesIds.indexOf(obj._id) > -1 ) {
				// remove the FXQuote
				this.exchangeRate.usedByQuotes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ExchangeRate
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ExchangeRate/update/' + this.exchangeRate;

	return  this.http.post(uri_, this.exchangeRate );
}

	//********************************************************************
	// loadHelper - internal helper to load a ExchangeRate
	//********************************************************************	
	loadHelper( id ) {
		this.getExchangeRate(id)
			.subscribe((res : ExchangeRate) => {
				this.exchangeRate = res;
			});
	}
}