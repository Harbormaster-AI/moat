import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PaymentContract} from '../models/PaymentContract';
import {MerchantService} from '../services/Merchant.service';
import {PaymentProcessorService} from '../services/PaymentProcessor.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PaymentContractService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	paymentContract : PaymentContract;

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
	// add a PaymentContract
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPaymentContract(contractNumber, pricingPlanCode, Merchant, Acquirer, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentContract/create';
		const obj = {
			      		contractNumber: contractNumber,
      		pricingPlanCode: pricingPlanCode,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		Acquirer: Acquirer != null && Acquirer.length > 0 ? Acquirer : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PaymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePaymentContract(contractNumber, pricingPlanCode, Merchant, Acquirer, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PaymentContract/update/' + id;
		const obj = {
				      		contractNumber: contractNumber,
      		pricingPlanCode: pricingPlanCode,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		Acquirer: Acquirer != null && Acquirer.length > 0 ? Acquirer : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PaymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePaymentContract(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentContract/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PaymentContract
	// returns the results untouched as an Observable PaymentContract
	// PaymentContract model
	// delegates via URI
	//********************************************************************
	getPaymentContract(id) : Observable<PaymentContract> {
		const uri_ = this.apiUrl + '/PaymentContract/load/' + id;

		return this.http.get<PaymentContract>(uri_);
	}
	
	//********************************************************************
	// gets all PaymentContract
	// returns the results untouched as JSON representation of an
	// Observable array of PaymentContract models
	// delegates via URI
	//********************************************************************
	getPaymentContracts() : Observable<PaymentContract[]> {
		const uri_ = this.apiUrl + '/PaymentContract/';

		return this
			.http.get<PaymentContract[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Merchant on a PaymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMerchant( paymentContractId, _merchantId ): Observable<any> {

		// get the PaymentContract from storage
		this.loadHelper( paymentContractId );

	// get the Merchant from storage
	var tmp 	= new MerchantService(this.http).getMerchant(_merchantId);

	// assign the Merchant
	this.paymentContract.merchant = tmp;

	// save the PaymentContract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Merchant on a PaymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMerchant( paymentContractId ): Observable<any> {

		// get the PaymentContract from storage
		this.loadHelper( paymentContractId );

	// assign Merchant to null
	this.paymentContract.merchant = null;

	// save the PaymentContract
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Acquirer on a PaymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAcquirer( paymentContractId, _acquirerId ): Observable<any> {

		// get the PaymentContract from storage
		this.loadHelper( paymentContractId );

	// get the PaymentProcessor from storage
	var tmp 	= new PaymentProcessorService(this.http).getPaymentProcessor(_acquirerId);

	// assign the Acquirer
	this.paymentContract.acquirer = tmp;

	// save the PaymentContract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Acquirer on a PaymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAcquirer( paymentContractId ): Observable<any> {

		// get the PaymentContract from storage
		this.loadHelper( paymentContractId );

	// assign Acquirer to null
	this.paymentContract.acquirer = null;

	// save the PaymentContract
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a PaymentContract
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PaymentContract/update/' + this.paymentContract;

	return  this.http.post(uri_, this.paymentContract );
}

	//********************************************************************
	// loadHelper - internal helper to load a PaymentContract
	//********************************************************************	
	loadHelper( id ) {
		this.getPaymentContract(id)
			.subscribe((res : PaymentContract) => {
				this.paymentContract = res;
			});
	}
}