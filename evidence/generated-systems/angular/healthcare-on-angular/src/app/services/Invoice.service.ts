import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Invoice} from '../models/Invoice';
import {PatientService} from '../services/Patient.service';
import {ClaimService} from '../services/Claim.service';
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
	addInvoice(invoiceNumber, totalAmount, dueDate, Patient, Claim, Payments, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Invoice/create';
		const obj = {
			      		invoiceNumber: invoiceNumber,
      		totalAmount: totalAmount,
      		dueDate: dueDate,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
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
		updateInvoice(invoiceNumber, totalAmount, dueDate, Patient, Claim, Payments, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Invoice/update/' + id;
		const obj = {
				      		invoiceNumber: invoiceNumber,
      		totalAmount: totalAmount,
      		dueDate: dueDate,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
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
	// assigns a Patient on a Invoice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( invoiceId, _patientId ): Observable<any> {

		// get the Invoice from storage
		this.loadHelper( invoiceId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.invoice.patient = tmp;

	// save the Invoice
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a Invoice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( invoiceId ): Observable<any> {

		// get the Invoice from storage
		this.loadHelper( invoiceId );

	// assign Patient to null
	this.invoice.patient = null;

	// save the Invoice
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Claim on a Invoice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignClaim( invoiceId, _claimId ): Observable<any> {

		// get the Invoice from storage
		this.loadHelper( invoiceId );

	// get the Claim from storage
	var tmp 	= new ClaimService(this.http).getClaim(_claimId);

	// assign the Claim
	this.invoice.claim = tmp;

	// save the Invoice
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Claim on a Invoice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignClaim( invoiceId ): Observable<any> {

		// get the Invoice from storage
		this.loadHelper( invoiceId );

	// assign Claim to null
	this.invoice.claim = null;

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