import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Payment} from '../models/Payment';
import {InvoiceService} from '../services/Invoice.service';
import {BillingAccountService} from '../services/BillingAccount.service';
import {PolicyService} from '../services/Policy.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PaymentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	payment : Payment;

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
	// add a Payment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPayment(paymentReference, amount, paymentDate, Invoice, BillingAccount, Policy, Method, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Payment/create';
		const obj = {
			      		paymentReference: paymentReference,
      		amount: amount,
      		paymentDate: paymentDate,
      		Invoice: Invoice != null && Invoice.length > 0 ? Invoice : null,
      		BillingAccount: BillingAccount != null && BillingAccount.length > 0 ? BillingAccount : null,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Method: Method,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePayment(paymentReference, amount, paymentDate, Invoice, BillingAccount, Policy, Method, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Payment/update/' + id;
		const obj = {
				      		paymentReference: paymentReference,
      		amount: amount,
      		paymentDate: paymentDate,
      		Invoice: Invoice != null && Invoice.length > 0 ? Invoice : null,
      		BillingAccount: BillingAccount != null && BillingAccount.length > 0 ? BillingAccount : null,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Method: Method,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePayment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Payment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Payment
	// returns the results untouched as an Observable Payment
	// Payment model
	// delegates via URI
	//********************************************************************
	getPayment(id) : Observable<Payment> {
		const uri_ = this.apiUrl + '/Payment/load/' + id;

		return this.http.get<Payment>(uri_);
	}
	
	//********************************************************************
	// gets all Payment
	// returns the results untouched as JSON representation of an
	// Observable array of Payment models
	// delegates via URI
	//********************************************************************
	getPayments() : Observable<Payment[]> {
		const uri_ = this.apiUrl + '/Payment/';

		return this
			.http.get<Payment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Invoice on a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInvoice( paymentId, _invoiceId ): Observable<any> {

		// get the Payment from storage
		this.loadHelper( paymentId );

	// get the Invoice from storage
	var tmp 	= new InvoiceService(this.http).getInvoice(_invoiceId);

	// assign the Invoice
	this.payment.invoice = tmp;

	// save the Payment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Invoice on a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInvoice( paymentId ): Observable<any> {

		// get the Payment from storage
		this.loadHelper( paymentId );

	// assign Invoice to null
	this.payment.invoice = null;

	// save the Payment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a BillingAccount on a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBillingAccount( paymentId, _billingAccountId ): Observable<any> {

		// get the Payment from storage
		this.loadHelper( paymentId );

	// get the BillingAccount from storage
	var tmp 	= new BillingAccountService(this.http).getBillingAccount(_billingAccountId);

	// assign the BillingAccount
	this.payment.billingAccount = tmp;

	// save the Payment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a BillingAccount on a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBillingAccount( paymentId ): Observable<any> {

		// get the Payment from storage
		this.loadHelper( paymentId );

	// assign BillingAccount to null
	this.payment.billingAccount = null;

	// save the Payment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Policy on a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( paymentId, _policyId ): Observable<any> {

		// get the Payment from storage
		this.loadHelper( paymentId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.payment.policy = tmp;

	// save the Payment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( paymentId ): Observable<any> {

		// get the Payment from storage
		this.loadHelper( paymentId );

	// assign Policy to null
	this.payment.policy = null;

	// save the Payment
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Payment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Payment/update/' + this.payment;

	return  this.http.post(uri_, this.payment );
}

	//********************************************************************
	// loadHelper - internal helper to load a Payment
	//********************************************************************	
	loadHelper( id ) {
		this.getPayment(id)
			.subscribe((res : Payment) => {
				this.payment = res;
			});
	}
}