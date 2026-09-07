import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SettlementBatch} from '../models/SettlementBatch';
import {PaymentProcessorService} from '../services/PaymentProcessor.service';
import {MerchantService} from '../services/Merchant.service';
import {PayoutService} from '../services/Payout.service';
import {TransactionService} from '../services/Transaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SettlementBatchService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	settlementBatch : SettlementBatch;

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
	// add a SettlementBatch
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSettlementBatch(batchId, periodStart, periodEnd, totalVolume, totalCount, Processor, Merchant, Payouts, Transactions, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/SettlementBatch/create';
		const obj = {
			      		batchId: batchId,
      		periodStart: periodStart,
      		periodEnd: periodEnd,
      		totalVolume: totalVolume,
      		totalCount: totalCount,
      		Processor: Processor != null && Processor.length > 0 ? Processor : null,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		Payouts: Payouts != null && Payouts.length > 0 ? Payouts : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SettlementBatch
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSettlementBatch(batchId, periodStart, periodEnd, totalVolume, totalCount, Processor, Merchant, Payouts, Transactions, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SettlementBatch/update/' + id;
		const obj = {
				      		batchId: batchId,
      		periodStart: periodStart,
      		periodEnd: periodEnd,
      		totalVolume: totalVolume,
      		totalCount: totalCount,
      		Processor: Processor != null && Processor.length > 0 ? Processor : null,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		Payouts: Payouts != null && Payouts.length > 0 ? Payouts : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SettlementBatch
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSettlementBatch(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SettlementBatch/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SettlementBatch
	// returns the results untouched as an Observable SettlementBatch
	// SettlementBatch model
	// delegates via URI
	//********************************************************************
	getSettlementBatch(id) : Observable<SettlementBatch> {
		const uri_ = this.apiUrl + '/SettlementBatch/load/' + id;

		return this.http.get<SettlementBatch>(uri_);
	}
	
	//********************************************************************
	// gets all SettlementBatch
	// returns the results untouched as JSON representation of an
	// Observable array of SettlementBatch models
	// delegates via URI
	//********************************************************************
	getSettlementBatchs() : Observable<SettlementBatch[]> {
		const uri_ = this.apiUrl + '/SettlementBatch/';

		return this
			.http.get<SettlementBatch[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Processor on a SettlementBatch
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProcessor( settlementBatchId, _processorId ): Observable<any> {

		// get the SettlementBatch from storage
		this.loadHelper( settlementBatchId );

	// get the PaymentProcessor from storage
	var tmp 	= new PaymentProcessorService(this.http).getPaymentProcessor(_processorId);

	// assign the Processor
	this.settlementBatch.processor = tmp;

	// save the SettlementBatch
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Processor on a SettlementBatch
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProcessor( settlementBatchId ): Observable<any> {

		// get the SettlementBatch from storage
		this.loadHelper( settlementBatchId );

	// assign Processor to null
	this.settlementBatch.processor = null;

	// save the SettlementBatch
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Merchant on a SettlementBatch
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMerchant( settlementBatchId, _merchantId ): Observable<any> {

		// get the SettlementBatch from storage
		this.loadHelper( settlementBatchId );

	// get the Merchant from storage
	var tmp 	= new MerchantService(this.http).getMerchant(_merchantId);

	// assign the Merchant
	this.settlementBatch.merchant = tmp;

	// save the SettlementBatch
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Merchant on a SettlementBatch
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMerchant( settlementBatchId ): Observable<any> {

		// get the SettlementBatch from storage
		this.loadHelper( settlementBatchId );

	// assign Merchant to null
	this.settlementBatch.merchant = null;

	// save the SettlementBatch
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more payoutsIds as a Payouts
	// to a SettlementBatch
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPayouts( settlementBatchId, payoutsIds ): Observable<any> {

		// get the SettlementBatch
		this.loadHelper( settlementBatchId );

	// split on a comma with no spaces
	var idList = payoutsIds.split(',')

	// iterate over array of payouts ids
	idList.forEach(function (id) {
		// read the Payout
		var payout = new PayoutService(this.http).getPayout(id);
		// add the Payout if not already assigned
		if ( this.settlementBatch.payouts.indexOf(payout) == -1 )
		this.settlementBatch.payouts.push(payout);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more payoutsIds as a Payouts
	// from a SettlementBatch
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePayouts( settlementBatchId, payoutsIds ): Observable<any> {

		// get the SettlementBatch
		this.loadHelper( settlementBatchId );


	// split on a comma with no spaces
	var idList 					= payoutsIds.split(',');
	var payouts 	= this.settlementBatch.payouts;

	if ( payouts != null && payoutsIds != null ) {

		// iterate over array of payouts ids
		payouts.forEach(function (obj) {
			if ( payoutsIds.indexOf(obj._id) > -1 ) {
				// remove the Payout
				this.settlementBatch.payouts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more transactionsIds as a Transactions
	// to a SettlementBatch
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTransactions( settlementBatchId, transactionsIds ): Observable<any> {

		// get the SettlementBatch
		this.loadHelper( settlementBatchId );

	// split on a comma with no spaces
	var idList = transactionsIds.split(',')

	// iterate over array of transactions ids
	idList.forEach(function (id) {
		// read the Transaction
		var transaction = new TransactionService(this.http).getTransaction(id);
		// add the Transaction if not already assigned
		if ( this.settlementBatch.transactions.indexOf(transaction) == -1 )
		this.settlementBatch.transactions.push(transaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more transactionsIds as a Transactions
	// from a SettlementBatch
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTransactions( settlementBatchId, transactionsIds ): Observable<any> {

		// get the SettlementBatch
		this.loadHelper( settlementBatchId );


	// split on a comma with no spaces
	var idList 					= transactionsIds.split(',');
	var transactions 	= this.settlementBatch.transactions;

	if ( transactions != null && transactionsIds != null ) {

		// iterate over array of transactions ids
		transactions.forEach(function (obj) {
			if ( transactionsIds.indexOf(obj._id) > -1 ) {
				// remove the Transaction
				this.settlementBatch.transactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a SettlementBatch
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SettlementBatch/update/' + this.settlementBatch;

	return  this.http.post(uri_, this.settlementBatch );
}

	//********************************************************************
	// loadHelper - internal helper to load a SettlementBatch
	//********************************************************************	
	loadHelper( id ) {
		this.getSettlementBatch(id)
			.subscribe((res : SettlementBatch) => {
				this.settlementBatch = res;
			});
	}
}