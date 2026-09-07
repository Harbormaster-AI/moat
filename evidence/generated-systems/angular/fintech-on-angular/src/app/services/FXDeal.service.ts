import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {FXDeal} from '../models/FXDeal';
import {FXQuoteService} from '../services/FXQuote.service';
import {PaymentOrderService} from '../services/PaymentOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class FXDealService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	fXDeal : FXDeal;

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
	// add a FXDeal
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addFXDeal(dealReference, baseCurrency, quoteCurrency, rate, amount, settlementDate, Quote, PaymentOrders, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/FXDeal/create';
		const obj = {
			      		dealReference: dealReference,
      		baseCurrency: baseCurrency,
      		quoteCurrency: quoteCurrency,
      		rate: rate,
      		amount: amount,
      		settlementDate: settlementDate,
      		Quote: Quote != null && Quote.length > 0 ? Quote : null,
      		PaymentOrders: PaymentOrders != null && PaymentOrders.length > 0 ? PaymentOrders : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a FXDeal
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateFXDeal(dealReference, baseCurrency, quoteCurrency, rate, amount, settlementDate, Quote, PaymentOrders, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/FXDeal/update/' + id;
		const obj = {
				      		dealReference: dealReference,
      		baseCurrency: baseCurrency,
      		quoteCurrency: quoteCurrency,
      		rate: rate,
      		amount: amount,
      		settlementDate: settlementDate,
      		Quote: Quote != null && Quote.length > 0 ? Quote : null,
      		PaymentOrders: PaymentOrders != null && PaymentOrders.length > 0 ? PaymentOrders : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a FXDeal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteFXDeal(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/FXDeal/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a FXDeal
	// returns the results untouched as an Observable FXDeal
	// FXDeal model
	// delegates via URI
	//********************************************************************
	getFXDeal(id) : Observable<FXDeal> {
		const uri_ = this.apiUrl + '/FXDeal/load/' + id;

		return this.http.get<FXDeal>(uri_);
	}
	
	//********************************************************************
	// gets all FXDeal
	// returns the results untouched as JSON representation of an
	// Observable array of FXDeal models
	// delegates via URI
	//********************************************************************
	getFXDeals() : Observable<FXDeal[]> {
		const uri_ = this.apiUrl + '/FXDeal/';

		return this
			.http.get<FXDeal[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Quote on a FXDeal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignQuote( fXDealId, _quoteId ): Observable<any> {

		// get the FXDeal from storage
		this.loadHelper( fXDealId );

	// get the FXQuote from storage
	var tmp 	= new FXQuoteService(this.http).getFXQuote(_quoteId);

	// assign the Quote
	this.fXDeal.quote = tmp;

	// save the FXDeal
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Quote on a FXDeal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignQuote( fXDealId ): Observable<any> {

		// get the FXDeal from storage
		this.loadHelper( fXDealId );

	// assign Quote to null
	this.fXDeal.quote = null;

	// save the FXDeal
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more paymentOrdersIds as a PaymentOrders
	// to a FXDeal
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPaymentOrders( fXDealId, paymentOrdersIds ): Observable<any> {

		// get the FXDeal
		this.loadHelper( fXDealId );

	// split on a comma with no spaces
	var idList = paymentOrdersIds.split(',')

	// iterate over array of paymentOrders ids
	idList.forEach(function (id) {
		// read the PaymentOrder
		var paymentOrder = new PaymentOrderService(this.http).getPaymentOrder(id);
		// add the PaymentOrder if not already assigned
		if ( this.fXDeal.paymentOrders.indexOf(paymentOrder) == -1 )
		this.fXDeal.paymentOrders.push(paymentOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more paymentOrdersIds as a PaymentOrders
	// from a FXDeal
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePaymentOrders( fXDealId, paymentOrdersIds ): Observable<any> {

		// get the FXDeal
		this.loadHelper( fXDealId );


	// split on a comma with no spaces
	var idList 					= paymentOrdersIds.split(',');
	var paymentOrders 	= this.fXDeal.paymentOrders;

	if ( paymentOrders != null && paymentOrdersIds != null ) {

		// iterate over array of paymentOrders ids
		paymentOrders.forEach(function (obj) {
			if ( paymentOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the PaymentOrder
				this.fXDeal.paymentOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a FXDeal
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/FXDeal/update/' + this.fXDeal;

	return  this.http.post(uri_, this.fXDeal );
}

	//********************************************************************
	// loadHelper - internal helper to load a FXDeal
	//********************************************************************	
	loadHelper( id ) {
		this.getFXDeal(id)
			.subscribe((res : FXDeal) => {
				this.fXDeal = res;
			});
	}
}