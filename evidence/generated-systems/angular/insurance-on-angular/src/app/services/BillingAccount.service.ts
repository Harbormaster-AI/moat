import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BillingAccount} from '../models/BillingAccount';
import {CustomerService} from '../services/Customer.service';
import {PolicyService} from '../services/Policy.service';
import {InvoiceService} from '../services/Invoice.service';
import {PaymentService} from '../services/Payment.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BillingAccountService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	billingAccount : BillingAccount;

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
	// add a BillingAccount
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBillingAccount(accountNumber, balance, Customer, Policies, Invoices, Payments, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/BillingAccount/create';
		const obj = {
			      		accountNumber: accountNumber,
      		balance: balance,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Invoices: Invoices != null && Invoices.length > 0 ? Invoices : null,
      		Payments: Payments != null && Payments.length > 0 ? Payments : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BillingAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBillingAccount(accountNumber, balance, Customer, Policies, Invoices, Payments, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BillingAccount/update/' + id;
		const obj = {
				      		accountNumber: accountNumber,
      		balance: balance,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Invoices: Invoices != null && Invoices.length > 0 ? Invoices : null,
      		Payments: Payments != null && Payments.length > 0 ? Payments : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BillingAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBillingAccount(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BillingAccount/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BillingAccount
	// returns the results untouched as an Observable BillingAccount
	// BillingAccount model
	// delegates via URI
	//********************************************************************
	getBillingAccount(id) : Observable<BillingAccount> {
		const uri_ = this.apiUrl + '/BillingAccount/load/' + id;

		return this.http.get<BillingAccount>(uri_);
	}
	
	//********************************************************************
	// gets all BillingAccount
	// returns the results untouched as JSON representation of an
	// Observable array of BillingAccount models
	// delegates via URI
	//********************************************************************
	getBillingAccounts() : Observable<BillingAccount[]> {
		const uri_ = this.apiUrl + '/BillingAccount/';

		return this
			.http.get<BillingAccount[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a BillingAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( billingAccountId, _customerId ): Observable<any> {

		// get the BillingAccount from storage
		this.loadHelper( billingAccountId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.billingAccount.customer = tmp;

	// save the BillingAccount
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a BillingAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( billingAccountId ): Observable<any> {

		// get the BillingAccount from storage
		this.loadHelper( billingAccountId );

	// assign Customer to null
	this.billingAccount.customer = null;

	// save the BillingAccount
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a BillingAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( billingAccountId, policiesIds ): Observable<any> {

		// get the BillingAccount
		this.loadHelper( billingAccountId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.billingAccount.policies.indexOf(policy) == -1 )
		this.billingAccount.policies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a BillingAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( billingAccountId, policiesIds ): Observable<any> {

		// get the BillingAccount
		this.loadHelper( billingAccountId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.billingAccount.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.billingAccount.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more invoicesIds as a Invoices
	// to a BillingAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInvoices( billingAccountId, invoicesIds ): Observable<any> {

		// get the BillingAccount
		this.loadHelper( billingAccountId );

	// split on a comma with no spaces
	var idList = invoicesIds.split(',')

	// iterate over array of invoices ids
	idList.forEach(function (id) {
		// read the Invoice
		var invoice = new InvoiceService(this.http).getInvoice(id);
		// add the Invoice if not already assigned
		if ( this.billingAccount.invoices.indexOf(invoice) == -1 )
		this.billingAccount.invoices.push(invoice);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more invoicesIds as a Invoices
	// from a BillingAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInvoices( billingAccountId, invoicesIds ): Observable<any> {

		// get the BillingAccount
		this.loadHelper( billingAccountId );


	// split on a comma with no spaces
	var idList 					= invoicesIds.split(',');
	var invoices 	= this.billingAccount.invoices;

	if ( invoices != null && invoicesIds != null ) {

		// iterate over array of invoices ids
		invoices.forEach(function (obj) {
			if ( invoicesIds.indexOf(obj._id) > -1 ) {
				// remove the Invoice
				this.billingAccount.invoices.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more paymentsIds as a Payments
	// to a BillingAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPayments( billingAccountId, paymentsIds ): Observable<any> {

		// get the BillingAccount
		this.loadHelper( billingAccountId );

	// split on a comma with no spaces
	var idList = paymentsIds.split(',')

	// iterate over array of payments ids
	idList.forEach(function (id) {
		// read the Payment
		var payment = new PaymentService(this.http).getPayment(id);
		// add the Payment if not already assigned
		if ( this.billingAccount.payments.indexOf(payment) == -1 )
		this.billingAccount.payments.push(payment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more paymentsIds as a Payments
	// from a BillingAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePayments( billingAccountId, paymentsIds ): Observable<any> {

		// get the BillingAccount
		this.loadHelper( billingAccountId );


	// split on a comma with no spaces
	var idList 					= paymentsIds.split(',');
	var payments 	= this.billingAccount.payments;

	if ( payments != null && paymentsIds != null ) {

		// iterate over array of payments ids
		payments.forEach(function (obj) {
			if ( paymentsIds.indexOf(obj._id) > -1 ) {
				// remove the Payment
				this.billingAccount.payments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BillingAccount
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BillingAccount/update/' + this.billingAccount;

	return  this.http.post(uri_, this.billingAccount );
}

	//********************************************************************
	// loadHelper - internal helper to load a BillingAccount
	//********************************************************************	
	loadHelper( id ) {
		this.getBillingAccount(id)
			.subscribe((res : BillingAccount) => {
				this.billingAccount = res;
			});
	}
}