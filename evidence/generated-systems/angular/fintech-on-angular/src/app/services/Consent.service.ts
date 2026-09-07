import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Consent} from '../models/Consent';
import {CustomerService} from '../services/Customer.service';
import {APIClientService} from '../services/APIClient.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ConsentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	consent : Consent;

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
	// add a Consent
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addConsent(grantedAt, expiresAt, scope, Customer, ApiClient, ConsentType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Consent/create';
		const obj = {
			      		grantedAt: grantedAt,
      		expiresAt: expiresAt,
      		scope: scope,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		ApiClient: ApiClient != null && ApiClient.length > 0 ? ApiClient : null,
      		ConsentType: ConsentType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Consent
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateConsent(grantedAt, expiresAt, scope, Customer, ApiClient, ConsentType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Consent/update/' + id;
		const obj = {
				      		grantedAt: grantedAt,
      		expiresAt: expiresAt,
      		scope: scope,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		ApiClient: ApiClient != null && ApiClient.length > 0 ? ApiClient : null,
      		ConsentType: ConsentType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Consent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteConsent(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Consent/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Consent
	// returns the results untouched as an Observable Consent
	// Consent model
	// delegates via URI
	//********************************************************************
	getConsent(id) : Observable<Consent> {
		const uri_ = this.apiUrl + '/Consent/load/' + id;

		return this.http.get<Consent>(uri_);
	}
	
	//********************************************************************
	// gets all Consent
	// returns the results untouched as JSON representation of an
	// Observable array of Consent models
	// delegates via URI
	//********************************************************************
	getConsents() : Observable<Consent[]> {
		const uri_ = this.apiUrl + '/Consent/';

		return this
			.http.get<Consent[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a Consent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( consentId, _customerId ): Observable<any> {

		// get the Consent from storage
		this.loadHelper( consentId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.consent.customer = tmp;

	// save the Consent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a Consent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( consentId ): Observable<any> {

		// get the Consent from storage
		this.loadHelper( consentId );

	// assign Customer to null
	this.consent.customer = null;

	// save the Consent
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ApiClient on a Consent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignApiClient( consentId, _apiClientId ): Observable<any> {

		// get the Consent from storage
		this.loadHelper( consentId );

	// get the APIClient from storage
	var tmp 	= new APIClientService(this.http).getAPIClient(_apiClientId);

	// assign the ApiClient
	this.consent.apiClient = tmp;

	// save the Consent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ApiClient on a Consent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignApiClient( consentId ): Observable<any> {

		// get the Consent from storage
		this.loadHelper( consentId );

	// assign ApiClient to null
	this.consent.apiClient = null;

	// save the Consent
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Consent
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Consent/update/' + this.consent;

	return  this.http.post(uri_, this.consent );
}

	//********************************************************************
	// loadHelper - internal helper to load a Consent
	//********************************************************************	
	loadHelper( id ) {
		this.getConsent(id)
			.subscribe((res : Consent) => {
				this.consent = res;
			});
	}
}