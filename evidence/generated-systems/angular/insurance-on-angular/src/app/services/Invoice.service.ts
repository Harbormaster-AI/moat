import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Invoice} from '../models/Invoice';
import {BillingAccountService} from '../services/BillingAccount.service';
import {PolicyService} from '../services/Policy.service';
import {PaymentService} from '../services/Payment.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InvoiceService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	invoice : Invoice;

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
	// add a Invoice
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInvoice(invoiceNumber, dueDate, totalDue, BillingAccount, Policy, Payments, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Invoice/create';
		const obj = {
			      		invoiceNumber: invoiceNumber,
      		dueDate: dueDate,
      		totalDue: totalDue,
      		BillingAccount: BillingAccount != null && BillingAccount.length > 0 ? BillingAccount : null,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Payments: Payments != null && Payments.length > 0 ? Payments : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Invoice
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInvoice(invoiceNumber, dueDate, totalDue, BillingAccount, Policy, Payments, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Invoice/update/' + id;
		const obj = {
				      		invoiceNumber: invoiceNumber,
      		dueDate: dueDate,
      		totalDue: totalDue,
      		BillingAccount: BillingAccount != null && BillingAccount.length > 0 ? BillingAccount : null,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Payments: Payments != null && Payments.length > 0 ? Payments : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Invoice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInvoice(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Invoice/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Invoice
	// returns the results untouched as an Observable Invoice
	// Invoice model
	// delegates via URI
	//********************************************************************
	getInvoice(id) : Observable<Invoice> {
		const uri_ = this.apiUrl + '/Invoice/load/' + id;

		return this.http.get<Invoice>(uri_);
	}
	
	//********************************************************************
	// gets all Invoice
	// returns the results untouched as JSON representation of an
	// Observable array of Invoice models
	// delegates via URI
	//********************************************************************
	getInvoices() : Observable<Invoice[]> {
		const uri_ = this.apiUrl + '/Invoice/';

		return this
			.http.get<Invoice[]>(uri_);
	}
	
			//********************************************************************
	// assigns a BillingAccount on a Invoice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBillingAccount( invoiceId, _billingAccountId ): Observable<any> {

		// get the Invoice from storage
		this.loadHelper( invoiceId );

	// get the BillingAccount from storage
	var tmp 	= new BillingAccountService(this.http).getBillingAccount(_billingAccountId);

	// assign the BillingAccount
	this.invoice.billingAccount = tmp;

	// save the Invoice
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a BillingAccount on a Invoice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBillingAccount( invoiceId ): Observable<any> {

		// get the Invoice from storage
		this.loadHelper( invoiceId );

	// assign BillingAccount to null
	this.invoice.billingAccount = null;

	// save the Invoice
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Policy on a Invoice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( invoiceId, _policyId ): Observable<any> {

		// get the Invoice from storage
		this.loadHelper( invoiceId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.invoice.policy = tmp;

	// save the Invoice
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Invoice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( invoiceId ): Observable<any> {

		// get the Invoice from storage
		this.loadHelper( invoiceId );

	// assign Policy to null
	this.invoice.policy = null;

	// save the Invoice
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more paymentsIds as a Payments
	// to a Invoice
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPayments( invoiceId, paymentsIds ): Observable<any> {

		// get the Invoice
		this.loadHelper( invoiceId );

	// split on a comma with no spaces
	var idList = paymentsIds.split(',')

	// iterate over array of payments ids
	idList.forEach(function (id) {
		// read the Payment
		var payment = new PaymentService(this.http).getPayment(id);
		// add the Payment if not already assigned
		if ( this.invoice.payments.indexOf(payment) == -1 )
		this.invoice.payments.push(payment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more paymentsIds as a Payments
	// from a Invoice
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePayments( invoiceId, paymentsIds ): Observable<any> {

		// get the Invoice
		this.loadHelper( invoiceId );


	// split on a comma with no spaces
	var idList 					= paymentsIds.split(',');
	var payments 	= this.invoice.payments;

	if ( payments != null && paymentsIds != null ) {

		// iterate over array of payments ids
		payments.forEach(function (obj) {
			if ( paymentsIds.indexOf(obj._id) > -1 ) {
				// remove the Payment
				this.invoice.payments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Invoice
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Invoice/update/' + this.invoice;

	return  this.http.post(uri_, this.invoice );
}

	//********************************************************************
	// loadHelper - internal helper to load a Invoice
	//********************************************************************	
	loadHelper( id ) {
		this.getInvoice(id)
			.subscribe((res : Invoice) => {
				this.invoice = res;
			});
	}
}