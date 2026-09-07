import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Beneficiary} from '../models/Beneficiary';
import {PolicyService} from '../services/Policy.service';
import {CustomerService} from '../services/Customer.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BeneficiaryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	beneficiary : Beneficiary;

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
	// add a Beneficiary
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBeneficiary(name, share, Policy, Customer, Relationship) : Observable<any> {
		const uri_ = this.apiUrl + '/Beneficiary/create';
		const obj = {
			      		name: name,
      		share: share,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
			Relationship: Relationship
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Beneficiary
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBeneficiary(name, share, Policy, Customer, Relationship, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Beneficiary/update/' + id;
		const obj = {
				      		name: name,
      		share: share,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
			Relationship: Relationship
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Beneficiary
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBeneficiary(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Beneficiary/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Beneficiary
	// returns the results untouched as an Observable Beneficiary
	// Beneficiary model
	// delegates via URI
	//********************************************************************
	getBeneficiary(id) : Observable<Beneficiary> {
		const uri_ = this.apiUrl + '/Beneficiary/load/' + id;

		return this.http.get<Beneficiary>(uri_);
	}
	
	//********************************************************************
	// gets all Beneficiary
	// returns the results untouched as JSON representation of an
	// Observable array of Beneficiary models
	// delegates via URI
	//********************************************************************
	getBeneficiarys() : Observable<Beneficiary[]> {
		const uri_ = this.apiUrl + '/Beneficiary/';

		return this
			.http.get<Beneficiary[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Policy on a Beneficiary
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( beneficiaryId, _policyId ): Observable<any> {

		// get the Beneficiary from storage
		this.loadHelper( beneficiaryId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.beneficiary.policy = tmp;

	// save the Beneficiary
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Beneficiary
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( beneficiaryId ): Observable<any> {

		// get the Beneficiary from storage
		this.loadHelper( beneficiaryId );

	// assign Policy to null
	this.beneficiary.policy = null;

	// save the Beneficiary
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Customer on a Beneficiary
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( beneficiaryId, _customerId ): Observable<any> {

		// get the Beneficiary from storage
		this.loadHelper( beneficiaryId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.beneficiary.customer = tmp;

	// save the Beneficiary
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a Beneficiary
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( beneficiaryId ): Observable<any> {

		// get the Beneficiary from storage
		this.loadHelper( beneficiaryId );

	// assign Customer to null
	this.beneficiary.customer = null;

	// save the Beneficiary
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Beneficiary
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Beneficiary/update/' + this.beneficiary;

	return  this.http.post(uri_, this.beneficiary );
}

	//********************************************************************
	// loadHelper - internal helper to load a Beneficiary
	//********************************************************************	
	loadHelper( id ) {
		this.getBeneficiary(id)
			.subscribe((res : Beneficiary) => {
				this.beneficiary = res;
			});
	}
}