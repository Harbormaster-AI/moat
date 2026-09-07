import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ClaimPayment} from '../models/ClaimPayment';
import {ClaimService} from '../services/Claim.service';
import {ExposureService} from '../services/Exposure.service';
import {BeneficiaryService} from '../services/Beneficiary.service';
import {ServiceProviderService} from '../services/ServiceProvider.service';
import {CustomerService} from '../services/Customer.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ClaimPaymentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	claimPayment : ClaimPayment;

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
	// add a ClaimPayment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addClaimPayment(paymentNumber, amount, paymentDate, Claim, Exposure, Beneficiary, ServiceProvider, Customer, PayeeType, Method, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/ClaimPayment/create';
		const obj = {
			      		paymentNumber: paymentNumber,
      		amount: amount,
      		paymentDate: paymentDate,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		Exposure: Exposure != null && Exposure.length > 0 ? Exposure : null,
      		Beneficiary: Beneficiary != null && Beneficiary.length > 0 ? Beneficiary : null,
      		ServiceProvider: ServiceProvider != null && ServiceProvider.length > 0 ? ServiceProvider : null,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		PayeeType: PayeeType,
      		Method: Method,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateClaimPayment(paymentNumber, amount, paymentDate, Claim, Exposure, Beneficiary, ServiceProvider, Customer, PayeeType, Method, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ClaimPayment/update/' + id;
		const obj = {
				      		paymentNumber: paymentNumber,
      		amount: amount,
      		paymentDate: paymentDate,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		Exposure: Exposure != null && Exposure.length > 0 ? Exposure : null,
      		Beneficiary: Beneficiary != null && Beneficiary.length > 0 ? Beneficiary : null,
      		ServiceProvider: ServiceProvider != null && ServiceProvider.length > 0 ? ServiceProvider : null,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		PayeeType: PayeeType,
      		Method: Method,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteClaimPayment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ClaimPayment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ClaimPayment
	// returns the results untouched as an Observable ClaimPayment
	// ClaimPayment model
	// delegates via URI
	//********************************************************************
	getClaimPayment(id) : Observable<ClaimPayment> {
		const uri_ = this.apiUrl + '/ClaimPayment/load/' + id;

		return this.http.get<ClaimPayment>(uri_);
	}
	
	//********************************************************************
	// gets all ClaimPayment
	// returns the results untouched as JSON representation of an
	// Observable array of ClaimPayment models
	// delegates via URI
	//********************************************************************
	getClaimPayments() : Observable<ClaimPayment[]> {
		const uri_ = this.apiUrl + '/ClaimPayment/';

		return this
			.http.get<ClaimPayment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Claim on a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignClaim( claimPaymentId, _claimId ): Observable<any> {

		// get the ClaimPayment from storage
		this.loadHelper( claimPaymentId );

	// get the Claim from storage
	var tmp 	= new ClaimService(this.http).getClaim(_claimId);

	// assign the Claim
	this.claimPayment.claim = tmp;

	// save the ClaimPayment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Claim on a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignClaim( claimPaymentId ): Observable<any> {

		// get the ClaimPayment from storage
		this.loadHelper( claimPaymentId );

	// assign Claim to null
	this.claimPayment.claim = null;

	// save the ClaimPayment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Exposure on a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignExposure( claimPaymentId, _exposureId ): Observable<any> {

		// get the ClaimPayment from storage
		this.loadHelper( claimPaymentId );

	// get the Exposure from storage
	var tmp 	= new ExposureService(this.http).getExposure(_exposureId);

	// assign the Exposure
	this.claimPayment.exposure = tmp;

	// save the ClaimPayment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Exposure on a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignExposure( claimPaymentId ): Observable<any> {

		// get the ClaimPayment from storage
		this.loadHelper( claimPaymentId );

	// assign Exposure to null
	this.claimPayment.exposure = null;

	// save the ClaimPayment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Beneficiary on a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBeneficiary( claimPaymentId, _beneficiaryId ): Observable<any> {

		// get the ClaimPayment from storage
		this.loadHelper( claimPaymentId );

	// get the Beneficiary from storage
	var tmp 	= new BeneficiaryService(this.http).getBeneficiary(_beneficiaryId);

	// assign the Beneficiary
	this.claimPayment.beneficiary = tmp;

	// save the ClaimPayment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Beneficiary on a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBeneficiary( claimPaymentId ): Observable<any> {

		// get the ClaimPayment from storage
		this.loadHelper( claimPaymentId );

	// assign Beneficiary to null
	this.claimPayment.beneficiary = null;

	// save the ClaimPayment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ServiceProvider on a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignServiceProvider( claimPaymentId, _serviceProviderId ): Observable<any> {

		// get the ClaimPayment from storage
		this.loadHelper( claimPaymentId );

	// get the ServiceProvider from storage
	var tmp 	= new ServiceProviderService(this.http).getServiceProvider(_serviceProviderId);

	// assign the ServiceProvider
	this.claimPayment.serviceProvider = tmp;

	// save the ClaimPayment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ServiceProvider on a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignServiceProvider( claimPaymentId ): Observable<any> {

		// get the ClaimPayment from storage
		this.loadHelper( claimPaymentId );

	// assign ServiceProvider to null
	this.claimPayment.serviceProvider = null;

	// save the ClaimPayment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Customer on a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( claimPaymentId, _customerId ): Observable<any> {

		// get the ClaimPayment from storage
		this.loadHelper( claimPaymentId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.claimPayment.customer = tmp;

	// save the ClaimPayment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a ClaimPayment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( claimPaymentId ): Observable<any> {

		// get the ClaimPayment from storage
		this.loadHelper( claimPaymentId );

	// assign Customer to null
	this.claimPayment.customer = null;

	// save the ClaimPayment
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ClaimPayment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ClaimPayment/update/' + this.claimPayment;

	return  this.http.post(uri_, this.claimPayment );
}

	//********************************************************************
	// loadHelper - internal helper to load a ClaimPayment
	//********************************************************************	
	loadHelper( id ) {
		this.getClaimPayment(id)
			.subscribe((res : ClaimPayment) => {
				this.claimPayment = res;
			});
	}
}