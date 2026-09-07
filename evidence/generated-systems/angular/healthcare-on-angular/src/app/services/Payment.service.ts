import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Payment} from '../models/Payment';
import {InvoiceService} from '../services/Invoice.service';
import {InsurancePayerService} from '../services/InsurancePayer.service';
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
	addPayment(paymentNumber, amount, paymentDate, Invoice, Payer, Method) : Observable<any> {
		const uri_ = this.apiUrl + '/Payment/create';
		const obj = {
			      		paymentNumber: paymentNumber,
      		amount: amount,
      		paymentDate: paymentDate,
      		Invoice: Invoice != null && Invoice.length > 0 ? Invoice : null,
      		Payer: Payer != null && Payer.length > 0 ? Payer : null,
			Method: Method
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePayment(paymentNumber, amount, paymentDate, Invoice, Payer, Method, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Payment/update/' + id;
		const obj = {
				      		paymentNumber: paymentNumber,
      		amount: amount,
      		paymentDate: paymentDate,
      		Invoice: Invoice != null && Invoice.length > 0 ? Invoice : null,
      		Payer: Payer != null && Payer.length > 0 ? Payer : null,
			Method: Method
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
	// assigns a Payer on a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPayer( paymentId, _payerId ): Observable<any> {

		// get the Payment from storage
		this.loadHelper( paymentId );

	// get the InsurancePayer from storage
	var tmp 	= new InsurancePayerService(this.http).getInsurancePayer(_payerId);

	// assign the Payer
	this.payment.payer = tmp;

	// save the Payment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Payer on a Payment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPayer( paymentId ): Observable<any> {

		// get the Payment from storage
		this.loadHelper( paymentId );

	// assign Payer to null
	this.payment.payer = null;

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