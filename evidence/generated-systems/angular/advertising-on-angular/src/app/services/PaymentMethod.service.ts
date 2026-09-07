import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PaymentMethod} from '../models/PaymentMethod';
import {BillingProfileService} from '../services/BillingProfile.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PaymentMethodService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	paymentMethod : PaymentMethod;

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
	// add a PaymentMethod
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPaymentMethod(last4, cardholderName, billingAddress, BillingProfile, MethodType) : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentMethod/create';
		const obj = {
			      		last4: last4,
      		cardholderName: cardholderName,
      		billingAddress: billingAddress,
      		BillingProfile: BillingProfile != null && BillingProfile.length > 0 ? BillingProfile : null,
			MethodType: MethodType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PaymentMethod
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePaymentMethod(last4, cardholderName, billingAddress, BillingProfile, MethodType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PaymentMethod/update/' + id;
		const obj = {
				      		last4: last4,
      		cardholderName: cardholderName,
      		billingAddress: billingAddress,
      		BillingProfile: BillingProfile != null && BillingProfile.length > 0 ? BillingProfile : null,
			MethodType: MethodType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PaymentMethod
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePaymentMethod(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentMethod/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PaymentMethod
	// returns the results untouched as an Observable PaymentMethod
	// PaymentMethod model
	// delegates via URI
	//********************************************************************
	getPaymentMethod(id) : Observable<PaymentMethod> {
		const uri_ = this.apiUrl + '/PaymentMethod/load/' + id;

		return this.http.get<PaymentMethod>(uri_);
	}
	
	//********************************************************************
	// gets all PaymentMethod
	// returns the results untouched as JSON representation of an
	// Observable array of PaymentMethod models
	// delegates via URI
	//********************************************************************
	getPaymentMethods() : Observable<PaymentMethod[]> {
		const uri_ = this.apiUrl + '/PaymentMethod/';

		return this
			.http.get<PaymentMethod[]>(uri_);
	}
	
			//********************************************************************
	// assigns a BillingProfile on a PaymentMethod
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBillingProfile( paymentMethodId, _billingProfileId ): Observable<any> {

		// get the PaymentMethod from storage
		this.loadHelper( paymentMethodId );

	// get the BillingProfile from storage
	var tmp 	= new BillingProfileService(this.http).getBillingProfile(_billingProfileId);

	// assign the BillingProfile
	this.paymentMethod.billingProfile = tmp;

	// save the PaymentMethod
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a BillingProfile on a PaymentMethod
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBillingProfile( paymentMethodId ): Observable<any> {

		// get the PaymentMethod from storage
		this.loadHelper( paymentMethodId );

	// assign BillingProfile to null
	this.paymentMethod.billingProfile = null;

	// save the PaymentMethod
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a PaymentMethod
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PaymentMethod/update/' + this.paymentMethod;

	return  this.http.post(uri_, this.paymentMethod );
}

	//********************************************************************
	// loadHelper - internal helper to load a PaymentMethod
	//********************************************************************	
	loadHelper( id ) {
		this.getPaymentMethod(id)
			.subscribe((res : PaymentMethod) => {
				this.paymentMethod = res;
			});
	}
}