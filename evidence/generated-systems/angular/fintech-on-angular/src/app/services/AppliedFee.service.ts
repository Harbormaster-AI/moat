import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AppliedFee} from '../models/AppliedFee';
import {PaymentOrderService} from '../services/PaymentOrder.service';
import {TransactionService} from '../services/Transaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AppliedFeeService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	appliedFee : AppliedFee;

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
	// add a AppliedFee
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAppliedFee(amount, description, PaymentOrder, Transaction, FeeType) : Observable<any> {
		const uri_ = this.apiUrl + '/AppliedFee/create';
		const obj = {
			      		amount: amount,
      		description: description,
      		PaymentOrder: PaymentOrder != null && PaymentOrder.length > 0 ? PaymentOrder : null,
      		Transaction: Transaction != null && Transaction.length > 0 ? Transaction : null,
			FeeType: FeeType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AppliedFee
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAppliedFee(amount, description, PaymentOrder, Transaction, FeeType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AppliedFee/update/' + id;
		const obj = {
				      		amount: amount,
      		description: description,
      		PaymentOrder: PaymentOrder != null && PaymentOrder.length > 0 ? PaymentOrder : null,
      		Transaction: Transaction != null && Transaction.length > 0 ? Transaction : null,
			FeeType: FeeType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AppliedFee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAppliedFee(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AppliedFee/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AppliedFee
	// returns the results untouched as an Observable AppliedFee
	// AppliedFee model
	// delegates via URI
	//********************************************************************
	getAppliedFee(id) : Observable<AppliedFee> {
		const uri_ = this.apiUrl + '/AppliedFee/load/' + id;

		return this.http.get<AppliedFee>(uri_);
	}
	
	//********************************************************************
	// gets all AppliedFee
	// returns the results untouched as JSON representation of an
	// Observable array of AppliedFee models
	// delegates via URI
	//********************************************************************
	getAppliedFees() : Observable<AppliedFee[]> {
		const uri_ = this.apiUrl + '/AppliedFee/';

		return this
			.http.get<AppliedFee[]>(uri_);
	}
	
			//********************************************************************
	// assigns a PaymentOrder on a AppliedFee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPaymentOrder( appliedFeeId, _paymentOrderId ): Observable<any> {

		// get the AppliedFee from storage
		this.loadHelper( appliedFeeId );

	// get the PaymentOrder from storage
	var tmp 	= new PaymentOrderService(this.http).getPaymentOrder(_paymentOrderId);

	// assign the PaymentOrder
	this.appliedFee.paymentOrder = tmp;

	// save the AppliedFee
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PaymentOrder on a AppliedFee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPaymentOrder( appliedFeeId ): Observable<any> {

		// get the AppliedFee from storage
		this.loadHelper( appliedFeeId );

	// assign PaymentOrder to null
	this.appliedFee.paymentOrder = null;

	// save the AppliedFee
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Transaction on a AppliedFee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTransaction( appliedFeeId, _transactionId ): Observable<any> {

		// get the AppliedFee from storage
		this.loadHelper( appliedFeeId );

	// get the Transaction from storage
	var tmp 	= new TransactionService(this.http).getTransaction(_transactionId);

	// assign the Transaction
	this.appliedFee.transaction = tmp;

	// save the AppliedFee
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Transaction on a AppliedFee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTransaction( appliedFeeId ): Observable<any> {

		// get the AppliedFee from storage
		this.loadHelper( appliedFeeId );

	// assign Transaction to null
	this.appliedFee.transaction = null;

	// save the AppliedFee
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a AppliedFee
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AppliedFee/update/' + this.appliedFee;

	return  this.http.post(uri_, this.appliedFee );
}

	//********************************************************************
	// loadHelper - internal helper to load a AppliedFee
	//********************************************************************	
	loadHelper( id ) {
		this.getAppliedFee(id)
			.subscribe((res : AppliedFee) => {
				this.appliedFee = res;
			});
	}
}