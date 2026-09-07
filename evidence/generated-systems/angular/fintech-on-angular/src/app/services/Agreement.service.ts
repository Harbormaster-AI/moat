import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Agreement} from '../models/Agreement';
import {CustomerService} from '../services/Customer.service';
import {ProductOfferingService} from '../services/ProductOffering.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AgreementService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	agreement : Agreement;

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
	// add a Agreement
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAgreement(agreementNumber, effectiveDate, Customer, ProductOffering, AgreementType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Agreement/create';
		const obj = {
			      		agreementNumber: agreementNumber,
      		effectiveDate: effectiveDate,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		ProductOffering: ProductOffering != null && ProductOffering.length > 0 ? ProductOffering : null,
      		AgreementType: AgreementType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Agreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAgreement(agreementNumber, effectiveDate, Customer, ProductOffering, AgreementType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Agreement/update/' + id;
		const obj = {
				      		agreementNumber: agreementNumber,
      		effectiveDate: effectiveDate,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		ProductOffering: ProductOffering != null && ProductOffering.length > 0 ? ProductOffering : null,
      		AgreementType: AgreementType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Agreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAgreement(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Agreement/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Agreement
	// returns the results untouched as an Observable Agreement
	// Agreement model
	// delegates via URI
	//********************************************************************
	getAgreement(id) : Observable<Agreement> {
		const uri_ = this.apiUrl + '/Agreement/load/' + id;

		return this.http.get<Agreement>(uri_);
	}
	
	//********************************************************************
	// gets all Agreement
	// returns the results untouched as JSON representation of an
	// Observable array of Agreement models
	// delegates via URI
	//********************************************************************
	getAgreements() : Observable<Agreement[]> {
		const uri_ = this.apiUrl + '/Agreement/';

		return this
			.http.get<Agreement[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a Agreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( agreementId, _customerId ): Observable<any> {

		// get the Agreement from storage
		this.loadHelper( agreementId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.agreement.customer = tmp;

	// save the Agreement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a Agreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( agreementId ): Observable<any> {

		// get the Agreement from storage
		this.loadHelper( agreementId );

	// assign Customer to null
	this.agreement.customer = null;

	// save the Agreement
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ProductOffering on a Agreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProductOffering( agreementId, _productOfferingId ): Observable<any> {

		// get the Agreement from storage
		this.loadHelper( agreementId );

	// get the ProductOffering from storage
	var tmp 	= new ProductOfferingService(this.http).getProductOffering(_productOfferingId);

	// assign the ProductOffering
	this.agreement.productOffering = tmp;

	// save the Agreement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ProductOffering on a Agreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProductOffering( agreementId ): Observable<any> {

		// get the Agreement from storage
		this.loadHelper( agreementId );

	// assign ProductOffering to null
	this.agreement.productOffering = null;

	// save the Agreement
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Agreement
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Agreement/update/' + this.agreement;

	return  this.http.post(uri_, this.agreement );
}

	//********************************************************************
	// loadHelper - internal helper to load a Agreement
	//********************************************************************	
	loadHelper( id ) {
		this.getAgreement(id)
			.subscribe((res : Agreement) => {
				this.agreement = res;
			});
	}
}