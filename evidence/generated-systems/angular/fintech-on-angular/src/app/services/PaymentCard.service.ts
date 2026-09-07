import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PaymentCard} from '../models/PaymentCard';
import {CustomerService} from '../services/Customer.service';
import {AccountService} from '../services/Account.service';
import {CardTokenizationService} from '../services/CardTokenization.service';
import {DisputeService} from '../services/Dispute.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PaymentCardService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	paymentCard : PaymentCard;

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
	// add a PaymentCard
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPaymentCard(cardToken, maskedPan, expiryMonth, expiryYear, cardholderName, Customer, Account, Tokenizations, Disputes, Scheme, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentCard/create';
		const obj = {
			      		cardToken: cardToken,
      		maskedPan: maskedPan,
      		expiryMonth: expiryMonth,
      		expiryYear: expiryYear,
      		cardholderName: cardholderName,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Tokenizations: Tokenizations != null && Tokenizations.length > 0 ? Tokenizations : null,
      		Disputes: Disputes != null && Disputes.length > 0 ? Disputes : null,
      		Scheme: Scheme,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PaymentCard
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePaymentCard(cardToken, maskedPan, expiryMonth, expiryYear, cardholderName, Customer, Account, Tokenizations, Disputes, Scheme, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PaymentCard/update/' + id;
		const obj = {
				      		cardToken: cardToken,
      		maskedPan: maskedPan,
      		expiryMonth: expiryMonth,
      		expiryYear: expiryYear,
      		cardholderName: cardholderName,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Tokenizations: Tokenizations != null && Tokenizations.length > 0 ? Tokenizations : null,
      		Disputes: Disputes != null && Disputes.length > 0 ? Disputes : null,
      		Scheme: Scheme,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PaymentCard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePaymentCard(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentCard/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PaymentCard
	// returns the results untouched as an Observable PaymentCard
	// PaymentCard model
	// delegates via URI
	//********************************************************************
	getPaymentCard(id) : Observable<PaymentCard> {
		const uri_ = this.apiUrl + '/PaymentCard/load/' + id;

		return this.http.get<PaymentCard>(uri_);
	}
	
	//********************************************************************
	// gets all PaymentCard
	// returns the results untouched as JSON representation of an
	// Observable array of PaymentCard models
	// delegates via URI
	//********************************************************************
	getPaymentCards() : Observable<PaymentCard[]> {
		const uri_ = this.apiUrl + '/PaymentCard/';

		return this
			.http.get<PaymentCard[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a PaymentCard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( paymentCardId, _customerId ): Observable<any> {

		// get the PaymentCard from storage
		this.loadHelper( paymentCardId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.paymentCard.customer = tmp;

	// save the PaymentCard
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a PaymentCard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( paymentCardId ): Observable<any> {

		// get the PaymentCard from storage
		this.loadHelper( paymentCardId );

	// assign Customer to null
	this.paymentCard.customer = null;

	// save the PaymentCard
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Account on a PaymentCard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( paymentCardId, _accountId ): Observable<any> {

		// get the PaymentCard from storage
		this.loadHelper( paymentCardId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.paymentCard.account = tmp;

	// save the PaymentCard
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a PaymentCard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( paymentCardId ): Observable<any> {

		// get the PaymentCard from storage
		this.loadHelper( paymentCardId );

	// assign Account to null
	this.paymentCard.account = null;

	// save the PaymentCard
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more tokenizationsIds as a Tokenizations
	// to a PaymentCard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTokenizations( paymentCardId, tokenizationsIds ): Observable<any> {

		// get the PaymentCard
		this.loadHelper( paymentCardId );

	// split on a comma with no spaces
	var idList = tokenizationsIds.split(',')

	// iterate over array of tokenizations ids
	idList.forEach(function (id) {
		// read the CardTokenization
		var cardTokenization = new CardTokenizationService(this.http).getCardTokenization(id);
		// add the CardTokenization if not already assigned
		if ( this.paymentCard.tokenizations.indexOf(cardTokenization) == -1 )
		this.paymentCard.tokenizations.push(cardTokenization);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tokenizationsIds as a Tokenizations
	// from a PaymentCard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTokenizations( paymentCardId, tokenizationsIds ): Observable<any> {

		// get the PaymentCard
		this.loadHelper( paymentCardId );


	// split on a comma with no spaces
	var idList 					= tokenizationsIds.split(',');
	var tokenizations 	= this.paymentCard.tokenizations;

	if ( tokenizations != null && tokenizationsIds != null ) {

		// iterate over array of tokenizations ids
		tokenizations.forEach(function (obj) {
			if ( tokenizationsIds.indexOf(obj._id) > -1 ) {
				// remove the CardTokenization
				this.paymentCard.tokenizations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more disputesIds as a Disputes
	// to a PaymentCard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDisputes( paymentCardId, disputesIds ): Observable<any> {

		// get the PaymentCard
		this.loadHelper( paymentCardId );

	// split on a comma with no spaces
	var idList = disputesIds.split(',')

	// iterate over array of disputes ids
	idList.forEach(function (id) {
		// read the Dispute
		var dispute = new DisputeService(this.http).getDispute(id);
		// add the Dispute if not already assigned
		if ( this.paymentCard.disputes.indexOf(dispute) == -1 )
		this.paymentCard.disputes.push(dispute);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more disputesIds as a Disputes
	// from a PaymentCard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDisputes( paymentCardId, disputesIds ): Observable<any> {

		// get the PaymentCard
		this.loadHelper( paymentCardId );


	// split on a comma with no spaces
	var idList 					= disputesIds.split(',');
	var disputes 	= this.paymentCard.disputes;

	if ( disputes != null && disputesIds != null ) {

		// iterate over array of disputes ids
		disputes.forEach(function (obj) {
			if ( disputesIds.indexOf(obj._id) > -1 ) {
				// remove the Dispute
				this.paymentCard.disputes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PaymentCard
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PaymentCard/update/' + this.paymentCard;

	return  this.http.post(uri_, this.paymentCard );
}

	//********************************************************************
	// loadHelper - internal helper to load a PaymentCard
	//********************************************************************	
	loadHelper( id ) {
		this.getPaymentCard(id)
			.subscribe((res : PaymentCard) => {
				this.paymentCard = res;
			});
	}
}