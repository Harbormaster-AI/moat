import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Dispute} from '../models/Dispute';
import {TransactionService} from '../services/Transaction.service';
import {PaymentCardService} from '../services/PaymentCard.service';
import {MerchantService} from '../services/Merchant.service';
import {ChargebackService} from '../services/Chargeback.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DisputeService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dispute : Dispute;

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
	// add a Dispute
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDispute(disputeReference, openedAt, closedAt, Transaction, Card, Merchant, Chargebacks, Reason, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Dispute/create';
		const obj = {
			      		disputeReference: disputeReference,
      		openedAt: openedAt,
      		closedAt: closedAt,
      		Transaction: Transaction != null && Transaction.length > 0 ? Transaction : null,
      		Card: Card != null && Card.length > 0 ? Card : null,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		Chargebacks: Chargebacks != null && Chargebacks.length > 0 ? Chargebacks : null,
      		Reason: Reason,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Dispute
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDispute(disputeReference, openedAt, closedAt, Transaction, Card, Merchant, Chargebacks, Reason, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Dispute/update/' + id;
		const obj = {
				      		disputeReference: disputeReference,
      		openedAt: openedAt,
      		closedAt: closedAt,
      		Transaction: Transaction != null && Transaction.length > 0 ? Transaction : null,
      		Card: Card != null && Card.length > 0 ? Card : null,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		Chargebacks: Chargebacks != null && Chargebacks.length > 0 ? Chargebacks : null,
      		Reason: Reason,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Dispute
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDispute(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Dispute/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Dispute
	// returns the results untouched as an Observable Dispute
	// Dispute model
	// delegates via URI
	//********************************************************************
	getDispute(id) : Observable<Dispute> {
		const uri_ = this.apiUrl + '/Dispute/load/' + id;

		return this.http.get<Dispute>(uri_);
	}
	
	//********************************************************************
	// gets all Dispute
	// returns the results untouched as JSON representation of an
	// Observable array of Dispute models
	// delegates via URI
	//********************************************************************
	getDisputes() : Observable<Dispute[]> {
		const uri_ = this.apiUrl + '/Dispute/';

		return this
			.http.get<Dispute[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Transaction on a Dispute
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTransaction( disputeId, _transactionId ): Observable<any> {

		// get the Dispute from storage
		this.loadHelper( disputeId );

	// get the Transaction from storage
	var tmp 	= new TransactionService(this.http).getTransaction(_transactionId);

	// assign the Transaction
	this.dispute.transaction = tmp;

	// save the Dispute
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Transaction on a Dispute
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTransaction( disputeId ): Observable<any> {

		// get the Dispute from storage
		this.loadHelper( disputeId );

	// assign Transaction to null
	this.dispute.transaction = null;

	// save the Dispute
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Card on a Dispute
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCard( disputeId, _cardId ): Observable<any> {

		// get the Dispute from storage
		this.loadHelper( disputeId );

	// get the PaymentCard from storage
	var tmp 	= new PaymentCardService(this.http).getPaymentCard(_cardId);

	// assign the Card
	this.dispute.card = tmp;

	// save the Dispute
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Card on a Dispute
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCard( disputeId ): Observable<any> {

		// get the Dispute from storage
		this.loadHelper( disputeId );

	// assign Card to null
	this.dispute.card = null;

	// save the Dispute
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Merchant on a Dispute
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMerchant( disputeId, _merchantId ): Observable<any> {

		// get the Dispute from storage
		this.loadHelper( disputeId );

	// get the Merchant from storage
	var tmp 	= new MerchantService(this.http).getMerchant(_merchantId);

	// assign the Merchant
	this.dispute.merchant = tmp;

	// save the Dispute
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Merchant on a Dispute
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMerchant( disputeId ): Observable<any> {

		// get the Dispute from storage
		this.loadHelper( disputeId );

	// assign Merchant to null
	this.dispute.merchant = null;

	// save the Dispute
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more chargebacksIds as a Chargebacks
	// to a Dispute
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addChargebacks( disputeId, chargebacksIds ): Observable<any> {

		// get the Dispute
		this.loadHelper( disputeId );

	// split on a comma with no spaces
	var idList = chargebacksIds.split(',')

	// iterate over array of chargebacks ids
	idList.forEach(function (id) {
		// read the Chargeback
		var chargeback = new ChargebackService(this.http).getChargeback(id);
		// add the Chargeback if not already assigned
		if ( this.dispute.chargebacks.indexOf(chargeback) == -1 )
		this.dispute.chargebacks.push(chargeback);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more chargebacksIds as a Chargebacks
	// from a Dispute
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeChargebacks( disputeId, chargebacksIds ): Observable<any> {

		// get the Dispute
		this.loadHelper( disputeId );


	// split on a comma with no spaces
	var idList 					= chargebacksIds.split(',');
	var chargebacks 	= this.dispute.chargebacks;

	if ( chargebacks != null && chargebacksIds != null ) {

		// iterate over array of chargebacks ids
		chargebacks.forEach(function (obj) {
			if ( chargebacksIds.indexOf(obj._id) > -1 ) {
				// remove the Chargeback
				this.dispute.chargebacks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Dispute
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Dispute/update/' + this.dispute;

	return  this.http.post(uri_, this.dispute );
}

	//********************************************************************
	// loadHelper - internal helper to load a Dispute
	//********************************************************************	
	loadHelper( id ) {
		this.getDispute(id)
			.subscribe((res : Dispute) => {
				this.dispute = res;
			});
	}
}