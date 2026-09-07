import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Transaction} from '../models/Transaction';
import {AccountService} from '../services/Account.service';
import {WalletService} from '../services/Wallet.service';
import {PaymentOrderService} from '../services/PaymentOrder.service';
import {MerchantService} from '../services/Merchant.service';
import {PaymentCardService} from '../services/PaymentCard.service';
import {ComplianceAlertService} from '../services/ComplianceAlert.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TransactionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	transaction : Transaction;

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
	// add a Transaction
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTransaction(amount, fee, exchangeRate, createdAt, completedAt, narrative, Account, Wallet, PaymentOrder, Merchant, Card, RelatedTransactions, Alerts, TransactionType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Transaction/create';
		const obj = {
			      		amount: amount,
      		fee: fee,
      		exchangeRate: exchangeRate,
      		createdAt: createdAt,
      		completedAt: completedAt,
      		narrative: narrative,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Wallet: Wallet != null && Wallet.length > 0 ? Wallet : null,
      		PaymentOrder: PaymentOrder != null && PaymentOrder.length > 0 ? PaymentOrder : null,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		Card: Card != null && Card.length > 0 ? Card : null,
      		RelatedTransactions: RelatedTransactions != null && RelatedTransactions.length > 0 ? RelatedTransactions : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
      		TransactionType: TransactionType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTransaction(amount, fee, exchangeRate, createdAt, completedAt, narrative, Account, Wallet, PaymentOrder, Merchant, Card, RelatedTransactions, Alerts, TransactionType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Transaction/update/' + id;
		const obj = {
				      		amount: amount,
      		fee: fee,
      		exchangeRate: exchangeRate,
      		createdAt: createdAt,
      		completedAt: completedAt,
      		narrative: narrative,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Wallet: Wallet != null && Wallet.length > 0 ? Wallet : null,
      		PaymentOrder: PaymentOrder != null && PaymentOrder.length > 0 ? PaymentOrder : null,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		Card: Card != null && Card.length > 0 ? Card : null,
      		RelatedTransactions: RelatedTransactions != null && RelatedTransactions.length > 0 ? RelatedTransactions : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
      		TransactionType: TransactionType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTransaction(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Transaction/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Transaction
	// returns the results untouched as an Observable Transaction
	// Transaction model
	// delegates via URI
	//********************************************************************
	getTransaction(id) : Observable<Transaction> {
		const uri_ = this.apiUrl + '/Transaction/load/' + id;

		return this.http.get<Transaction>(uri_);
	}
	
	//********************************************************************
	// gets all Transaction
	// returns the results untouched as JSON representation of an
	// Observable array of Transaction models
	// delegates via URI
	//********************************************************************
	getTransactions() : Observable<Transaction[]> {
		const uri_ = this.apiUrl + '/Transaction/';

		return this
			.http.get<Transaction[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Account on a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( transactionId, _accountId ): Observable<any> {

		// get the Transaction from storage
		this.loadHelper( transactionId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.transaction.account = tmp;

	// save the Transaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( transactionId ): Observable<any> {

		// get the Transaction from storage
		this.loadHelper( transactionId );

	// assign Account to null
	this.transaction.account = null;

	// save the Transaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Wallet on a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWallet( transactionId, _walletId ): Observable<any> {

		// get the Transaction from storage
		this.loadHelper( transactionId );

	// get the Wallet from storage
	var tmp 	= new WalletService(this.http).getWallet(_walletId);

	// assign the Wallet
	this.transaction.wallet = tmp;

	// save the Transaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Wallet on a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWallet( transactionId ): Observable<any> {

		// get the Transaction from storage
		this.loadHelper( transactionId );

	// assign Wallet to null
	this.transaction.wallet = null;

	// save the Transaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PaymentOrder on a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPaymentOrder( transactionId, _paymentOrderId ): Observable<any> {

		// get the Transaction from storage
		this.loadHelper( transactionId );

	// get the PaymentOrder from storage
	var tmp 	= new PaymentOrderService(this.http).getPaymentOrder(_paymentOrderId);

	// assign the PaymentOrder
	this.transaction.paymentOrder = tmp;

	// save the Transaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PaymentOrder on a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPaymentOrder( transactionId ): Observable<any> {

		// get the Transaction from storage
		this.loadHelper( transactionId );

	// assign PaymentOrder to null
	this.transaction.paymentOrder = null;

	// save the Transaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Merchant on a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMerchant( transactionId, _merchantId ): Observable<any> {

		// get the Transaction from storage
		this.loadHelper( transactionId );

	// get the Merchant from storage
	var tmp 	= new MerchantService(this.http).getMerchant(_merchantId);

	// assign the Merchant
	this.transaction.merchant = tmp;

	// save the Transaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Merchant on a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMerchant( transactionId ): Observable<any> {

		// get the Transaction from storage
		this.loadHelper( transactionId );

	// assign Merchant to null
	this.transaction.merchant = null;

	// save the Transaction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Card on a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCard( transactionId, _cardId ): Observable<any> {

		// get the Transaction from storage
		this.loadHelper( transactionId );

	// get the PaymentCard from storage
	var tmp 	= new PaymentCardService(this.http).getPaymentCard(_cardId);

	// assign the Card
	this.transaction.card = tmp;

	// save the Transaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Card on a Transaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCard( transactionId ): Observable<any> {

		// get the Transaction from storage
		this.loadHelper( transactionId );

	// assign Card to null
	this.transaction.card = null;

	// save the Transaction
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more relatedTransactionsIds as a RelatedTransactions
	// to a Transaction
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRelatedTransactions( transactionId, relatedTransactionsIds ): Observable<any> {

		// get the Transaction
		this.loadHelper( transactionId );

	// split on a comma with no spaces
	var idList = relatedTransactionsIds.split(',')

	// iterate over array of relatedTransactions ids
	idList.forEach(function (id) {
		// read the Transaction
		var transaction = new TransactionService(this.http).getTransaction(id);
		// add the Transaction if not already assigned
		if ( this.transaction.relatedTransactions.indexOf(transaction) == -1 )
		this.transaction.relatedTransactions.push(transaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more relatedTransactionsIds as a RelatedTransactions
	// from a Transaction
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRelatedTransactions( transactionId, relatedTransactionsIds ): Observable<any> {

		// get the Transaction
		this.loadHelper( transactionId );


	// split on a comma with no spaces
	var idList 					= relatedTransactionsIds.split(',');
	var relatedTransactions 	= this.transaction.relatedTransactions;

	if ( relatedTransactions != null && relatedTransactionsIds != null ) {

		// iterate over array of relatedTransactions ids
		relatedTransactions.forEach(function (obj) {
			if ( relatedTransactionsIds.indexOf(obj._id) > -1 ) {
				// remove the Transaction
				this.transaction.relatedTransactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more alertsIds as a Alerts
	// to a Transaction
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAlerts( transactionId, alertsIds ): Observable<any> {

		// get the Transaction
		this.loadHelper( transactionId );

	// split on a comma with no spaces
	var idList = alertsIds.split(',')

	// iterate over array of alerts ids
	idList.forEach(function (id) {
		// read the ComplianceAlert
		var complianceAlert = new ComplianceAlertService(this.http).getComplianceAlert(id);
		// add the ComplianceAlert if not already assigned
		if ( this.transaction.alerts.indexOf(complianceAlert) == -1 )
		this.transaction.alerts.push(complianceAlert);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more alertsIds as a Alerts
	// from a Transaction
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAlerts( transactionId, alertsIds ): Observable<any> {

		// get the Transaction
		this.loadHelper( transactionId );


	// split on a comma with no spaces
	var idList 					= alertsIds.split(',');
	var alerts 	= this.transaction.alerts;

	if ( alerts != null && alertsIds != null ) {

		// iterate over array of alerts ids
		alerts.forEach(function (obj) {
			if ( alertsIds.indexOf(obj._id) > -1 ) {
				// remove the ComplianceAlert
				this.transaction.alerts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Transaction
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Transaction/update/' + this.transaction;

	return  this.http.post(uri_, this.transaction );
}

	//********************************************************************
	// loadHelper - internal helper to load a Transaction
	//********************************************************************	
	loadHelper( id ) {
		this.getTransaction(id)
			.subscribe((res : Transaction) => {
				this.transaction = res;
			});
	}
}