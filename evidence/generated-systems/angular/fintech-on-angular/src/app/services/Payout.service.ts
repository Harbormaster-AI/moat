import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Payout} from '../models/Payout';
import {MerchantService} from '../services/Merchant.service';
import {SettlementBatchService} from '../services/SettlementBatch.service';
import {AccountService} from '../services/Account.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PayoutService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	payout : Payout;

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
	// add a Payout
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPayout(payoutReference, amount, currency, scheduledDate, paidDate, Merchant, SettlementBatch, DestinationAccount, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Payout/create';
		const obj = {
			      		payoutReference: payoutReference,
      		amount: amount,
      		currency: currency,
      		scheduledDate: scheduledDate,
      		paidDate: paidDate,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		SettlementBatch: SettlementBatch != null && SettlementBatch.length > 0 ? SettlementBatch : null,
      		DestinationAccount: DestinationAccount != null && DestinationAccount.length > 0 ? DestinationAccount : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Payout
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePayout(payoutReference, amount, currency, scheduledDate, paidDate, Merchant, SettlementBatch, DestinationAccount, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Payout/update/' + id;
		const obj = {
				      		payoutReference: payoutReference,
      		amount: amount,
      		currency: currency,
      		scheduledDate: scheduledDate,
      		paidDate: paidDate,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		SettlementBatch: SettlementBatch != null && SettlementBatch.length > 0 ? SettlementBatch : null,
      		DestinationAccount: DestinationAccount != null && DestinationAccount.length > 0 ? DestinationAccount : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Payout
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePayout(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Payout/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Payout
	// returns the results untouched as an Observable Payout
	// Payout model
	// delegates via URI
	//********************************************************************
	getPayout(id) : Observable<Payout> {
		const uri_ = this.apiUrl + '/Payout/load/' + id;

		return this.http.get<Payout>(uri_);
	}
	
	//********************************************************************
	// gets all Payout
	// returns the results untouched as JSON representation of an
	// Observable array of Payout models
	// delegates via URI
	//********************************************************************
	getPayouts() : Observable<Payout[]> {
		const uri_ = this.apiUrl + '/Payout/';

		return this
			.http.get<Payout[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Merchant on a Payout
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMerchant( payoutId, _merchantId ): Observable<any> {

		// get the Payout from storage
		this.loadHelper( payoutId );

	// get the Merchant from storage
	var tmp 	= new MerchantService(this.http).getMerchant(_merchantId);

	// assign the Merchant
	this.payout.merchant = tmp;

	// save the Payout
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Merchant on a Payout
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMerchant( payoutId ): Observable<any> {

		// get the Payout from storage
		this.loadHelper( payoutId );

	// assign Merchant to null
	this.payout.merchant = null;

	// save the Payout
	return this.saveHelper();
}

		//********************************************************************
	// assigns a SettlementBatch on a Payout
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSettlementBatch( payoutId, _settlementBatchId ): Observable<any> {

		// get the Payout from storage
		this.loadHelper( payoutId );

	// get the SettlementBatch from storage
	var tmp 	= new SettlementBatchService(this.http).getSettlementBatch(_settlementBatchId);

	// assign the SettlementBatch
	this.payout.settlementBatch = tmp;

	// save the Payout
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SettlementBatch on a Payout
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSettlementBatch( payoutId ): Observable<any> {

		// get the Payout from storage
		this.loadHelper( payoutId );

	// assign SettlementBatch to null
	this.payout.settlementBatch = null;

	// save the Payout
	return this.saveHelper();
}

		//********************************************************************
	// assigns a DestinationAccount on a Payout
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDestinationAccount( payoutId, _destinationAccountId ): Observable<any> {

		// get the Payout from storage
		this.loadHelper( payoutId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_destinationAccountId);

	// assign the DestinationAccount
	this.payout.destinationAccount = tmp;

	// save the Payout
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a DestinationAccount on a Payout
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDestinationAccount( payoutId ): Observable<any> {

		// get the Payout from storage
		this.loadHelper( payoutId );

	// assign DestinationAccount to null
	this.payout.destinationAccount = null;

	// save the Payout
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Payout
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Payout/update/' + this.payout;

	return  this.http.post(uri_, this.payout );
}

	//********************************************************************
	// loadHelper - internal helper to load a Payout
	//********************************************************************	
	loadHelper( id ) {
		this.getPayout(id)
			.subscribe((res : Payout) => {
				this.payout = res;
			});
	}
}