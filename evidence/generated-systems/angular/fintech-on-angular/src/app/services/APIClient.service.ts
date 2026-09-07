import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {APIClient} from '../models/APIClient';
import {ConsentService} from '../services/Consent.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class APIClientService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aPIClient : APIClient;

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
	// add a APIClient
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAPIClient(name, clientId, redirectUri, Consents, ClientType) : Observable<any> {
		const uri_ = this.apiUrl + '/APIClient/create';
		const obj = {
			      		name: name,
      		clientId: clientId,
      		redirectUri: redirectUri,
      		Consents: Consents != null && Consents.length > 0 ? Consents : null,
			ClientType: ClientType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a APIClient
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAPIClient(name, clientId, redirectUri, Consents, ClientType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/APIClient/update/' + id;
		const obj = {
				      		name: name,
      		clientId: clientId,
      		redirectUri: redirectUri,
      		Consents: Consents != null && Consents.length > 0 ? Consents : null,
			ClientType: ClientType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a APIClient
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAPIClient(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/APIClient/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a APIClient
	// returns the results untouched as an Observable APIClient
	// APIClient model
	// delegates via URI
	//********************************************************************
	getAPIClient(id) : Observable<APIClient> {
		const uri_ = this.apiUrl + '/APIClient/load/' + id;

		return this.http.get<APIClient>(uri_);
	}
	
	//********************************************************************
	// gets all APIClient
	// returns the results untouched as JSON representation of an
	// Observable array of APIClient models
	// delegates via URI
	//********************************************************************
	getAPIClients() : Observable<APIClient[]> {
		const uri_ = this.apiUrl + '/APIClient/';

		return this
			.http.get<APIClient[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more consentsIds as a Consents
	// to a APIClient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addConsents( aPIClientId, consentsIds ): Observable<any> {

		// get the APIClient
		this.loadHelper( aPIClientId );

	// split on a comma with no spaces
	var idList = consentsIds.split(',')

	// iterate over array of consents ids
	idList.forEach(function (id) {
		// read the Consent
		var consent = new ConsentService(this.http).getConsent(id);
		// add the Consent if not already assigned
		if ( this.aPIClient.consents.indexOf(consent) == -1 )
		this.aPIClient.consents.push(consent);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more consentsIds as a Consents
	// from a APIClient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeConsents( aPIClientId, consentsIds ): Observable<any> {

		// get the APIClient
		this.loadHelper( aPIClientId );


	// split on a comma with no spaces
	var idList 					= consentsIds.split(',');
	var consents 	= this.aPIClient.consents;

	if ( consents != null && consentsIds != null ) {

		// iterate over array of consents ids
		consents.forEach(function (obj) {
			if ( consentsIds.indexOf(obj._id) > -1 ) {
				// remove the Consent
				this.aPIClient.consents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a APIClient
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/APIClient/update/' + this.aPIClient;

	return  this.http.post(uri_, this.aPIClient );
}

	//********************************************************************
	// loadHelper - internal helper to load a APIClient
	//********************************************************************	
	loadHelper( id ) {
		this.getAPIClient(id)
			.subscribe((res : APIClient) => {
				this.aPIClient = res;
			});
	}
}