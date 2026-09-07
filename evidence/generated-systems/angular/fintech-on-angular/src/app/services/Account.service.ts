import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Account} from '../models/Account';
import {CustomerService} from '../services/Customer.service';
import {FinancialInstitutionService} from '../services/FinancialInstitution.service';
import {TransactionService} from '../services/Transaction.service';
import {PaymentCardService} from '../services/PaymentCard.service';
import {AccountStatementService} from '../services/AccountStatement.service';
import {DirectDebitMandateService} from '../services/DirectDebitMandate.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AccountService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	account : Account;

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
	// add a Account
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAccount(accountNumber, iban, bic, openedDate, currency, balance, availableBalance, Customer, Institution, Transactions, Cards, Statements, Mandates, AccountType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Account/create';
		const obj = {
			      		accountNumber: accountNumber,
      		iban: iban,
      		bic: bic,
      		openedDate: openedDate,
      		currency: currency,
      		balance: balance,
      		availableBalance: availableBalance,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Institution: Institution != null && Institution.length > 0 ? Institution : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
      		Cards: Cards != null && Cards.length > 0 ? Cards : null,
      		Statements: Statements != null && Statements.length > 0 ? Statements : null,
      		Mandates: Mandates != null && Mandates.length > 0 ? Mandates : null,
      		AccountType: AccountType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAccount(accountNumber, iban, bic, openedDate, currency, balance, availableBalance, Customer, Institution, Transactions, Cards, Statements, Mandates, AccountType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Account/update/' + id;
		const obj = {
				      		accountNumber: accountNumber,
      		iban: iban,
      		bic: bic,
      		openedDate: openedDate,
      		currency: currency,
      		balance: balance,
      		availableBalance: availableBalance,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Institution: Institution != null && Institution.length > 0 ? Institution : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
      		Cards: Cards != null && Cards.length > 0 ? Cards : null,
      		Statements: Statements != null && Statements.length > 0 ? Statements : null,
      		Mandates: Mandates != null && Mandates.length > 0 ? Mandates : null,
      		AccountType: AccountType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAccount(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Account/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Account
	// returns the results untouched as an Observable Account
	// Account model
	// delegates via URI
	//********************************************************************
	getAccount(id) : Observable<Account> {
		const uri_ = this.apiUrl + '/Account/load/' + id;

		return this.http.get<Account>(uri_);
	}
	
	//********************************************************************
	// gets all Account
	// returns the results untouched as JSON representation of an
	// Observable array of Account models
	// delegates via URI
	//********************************************************************
	getAccounts() : Observable<Account[]> {
		const uri_ = this.apiUrl + '/Account/';

		return this
			.http.get<Account[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( accountId, _customerId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.account.customer = tmp;

	// save the Account
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( accountId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// assign Customer to null
	this.account.customer = null;

	// save the Account
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Institution on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInstitution( accountId, _institutionId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// get the FinancialInstitution from storage
	var tmp 	= new FinancialInstitutionService(this.http).getFinancialInstitution(_institutionId);

	// assign the Institution
	this.account.institution = tmp;

	// save the Account
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Institution on a Account
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInstitution( accountId ): Observable<any> {

		// get the Account from storage
		this.loadHelper( accountId );

	// assign Institution to null
	this.account.institution = null;

	// save the Account
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more transactionsIds as a Transactions
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTransactions( accountId, transactionsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = transactionsIds.split(',')

	// iterate over array of transactions ids
	idList.forEach(function (id) {
		// read the Transaction
		var transaction = new TransactionService(this.http).getTransaction(id);
		// add the Transaction if not already assigned
		if ( this.account.transactions.indexOf(transaction) == -1 )
		this.account.transactions.push(transaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more transactionsIds as a Transactions
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTransactions( accountId, transactionsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= transactionsIds.split(',');
	var transactions 	= this.account.transactions;

	if ( transactions != null && transactionsIds != null ) {

		// iterate over array of transactions ids
		transactions.forEach(function (obj) {
			if ( transactionsIds.indexOf(obj._id) > -1 ) {
				// remove the Transaction
				this.account.transactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more cardsIds as a Cards
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCards( accountId, cardsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = cardsIds.split(',')

	// iterate over array of cards ids
	idList.forEach(function (id) {
		// read the PaymentCard
		var paymentCard = new PaymentCardService(this.http).getPaymentCard(id);
		// add the PaymentCard if not already assigned
		if ( this.account.cards.indexOf(paymentCard) == -1 )
		this.account.cards.push(paymentCard);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more cardsIds as a Cards
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCards( accountId, cardsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= cardsIds.split(',');
	var cards 	= this.account.cards;

	if ( cards != null && cardsIds != null ) {

		// iterate over array of cards ids
		cards.forEach(function (obj) {
			if ( cardsIds.indexOf(obj._id) > -1 ) {
				// remove the PaymentCard
				this.account.cards.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more statementsIds as a Statements
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addStatements( accountId, statementsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = statementsIds.split(',')

	// iterate over array of statements ids
	idList.forEach(function (id) {
		// read the AccountStatement
		var accountStatement = new AccountStatementService(this.http).getAccountStatement(id);
		// add the AccountStatement if not already assigned
		if ( this.account.statements.indexOf(accountStatement) == -1 )
		this.account.statements.push(accountStatement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more statementsIds as a Statements
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeStatements( accountId, statementsIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= statementsIds.split(',');
	var statements 	= this.account.statements;

	if ( statements != null && statementsIds != null ) {

		// iterate over array of statements ids
		statements.forEach(function (obj) {
			if ( statementsIds.indexOf(obj._id) > -1 ) {
				// remove the AccountStatement
				this.account.statements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more mandatesIds as a Mandates
	// to a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMandates( accountId, mandatesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );

	// split on a comma with no spaces
	var idList = mandatesIds.split(',')

	// iterate over array of mandates ids
	idList.forEach(function (id) {
		// read the DirectDebitMandate
		var directDebitMandate = new DirectDebitMandateService(this.http).getDirectDebitMandate(id);
		// add the DirectDebitMandate if not already assigned
		if ( this.account.mandates.indexOf(directDebitMandate) == -1 )
		this.account.mandates.push(directDebitMandate);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more mandatesIds as a Mandates
	// from a Account
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMandates( accountId, mandatesIds ): Observable<any> {

		// get the Account
		this.loadHelper( accountId );


	// split on a comma with no spaces
	var idList 					= mandatesIds.split(',');
	var mandates 	= this.account.mandates;

	if ( mandates != null && mandatesIds != null ) {

		// iterate over array of mandates ids
		mandates.forEach(function (obj) {
			if ( mandatesIds.indexOf(obj._id) > -1 ) {
				// remove the DirectDebitMandate
				this.account.mandates.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Account
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Account/update/' + this.account;

	return  this.http.post(uri_, this.account );
}

	//********************************************************************
	// loadHelper - internal helper to load a Account
	//********************************************************************	
	loadHelper( id ) {
		this.getAccount(id)
			.subscribe((res : Account) => {
				this.account = res;
			});
	}
}