import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Wallet} from '../models/Wallet';
import {CustomerService} from '../services/Customer.service';
import {TransactionService} from '../services/Transaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class WalletService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	wallet : Wallet;

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
	// add a Wallet
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addWallet(currency, balance, Customer, Transactions, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Wallet/create';
		const obj = {
			      		currency: currency,
      		balance: balance,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Wallet
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateWallet(currency, balance, Customer, Transactions, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Wallet/update/' + id;
		const obj = {
				      		currency: currency,
      		balance: balance,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Wallet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteWallet(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Wallet/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Wallet
	// returns the results untouched as an Observable Wallet
	// Wallet model
	// delegates via URI
	//********************************************************************
	getWallet(id) : Observable<Wallet> {
		const uri_ = this.apiUrl + '/Wallet/load/' + id;

		return this.http.get<Wallet>(uri_);
	}
	
	//********************************************************************
	// gets all Wallet
	// returns the results untouched as JSON representation of an
	// Observable array of Wallet models
	// delegates via URI
	//********************************************************************
	getWallets() : Observable<Wallet[]> {
		const uri_ = this.apiUrl + '/Wallet/';

		return this
			.http.get<Wallet[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a Wallet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( walletId, _customerId ): Observable<any> {

		// get the Wallet from storage
		this.loadHelper( walletId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.wallet.customer = tmp;

	// save the Wallet
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a Wallet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( walletId ): Observable<any> {

		// get the Wallet from storage
		this.loadHelper( walletId );

	// assign Customer to null
	this.wallet.customer = null;

	// save the Wallet
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more transactionsIds as a Transactions
	// to a Wallet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTransactions( walletId, transactionsIds ): Observable<any> {

		// get the Wallet
		this.loadHelper( walletId );

	// split on a comma with no spaces
	var idList = transactionsIds.split(',')

	// iterate over array of transactions ids
	idList.forEach(function (id) {
		// read the Transaction
		var transaction = new TransactionService(this.http).getTransaction(id);
		// add the Transaction if not already assigned
		if ( this.wallet.transactions.indexOf(transaction) == -1 )
		this.wallet.transactions.push(transaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more transactionsIds as a Transactions
	// from a Wallet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTransactions( walletId, transactionsIds ): Observable<any> {

		// get the Wallet
		this.loadHelper( walletId );


	// split on a comma with no spaces
	var idList 					= transactionsIds.split(',');
	var transactions 	= this.wallet.transactions;

	if ( transactions != null && transactionsIds != null ) {

		// iterate over array of transactions ids
		transactions.forEach(function (obj) {
			if ( transactionsIds.indexOf(obj._id) > -1 ) {
				// remove the Transaction
				this.wallet.transactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Wallet
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Wallet/update/' + this.wallet;

	return  this.http.post(uri_, this.wallet );
}

	//********************************************************************
	// loadHelper - internal helper to load a Wallet
	//********************************************************************	
	loadHelper( id ) {
		this.getWallet(id)
			.subscribe((res : Wallet) => {
				this.wallet = res;
			});
	}
}