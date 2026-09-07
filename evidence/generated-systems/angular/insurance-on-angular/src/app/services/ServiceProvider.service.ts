import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ServiceProvider} from '../models/ServiceProvider';
import {ClaimService} from '../services/Claim.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ServiceProviderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	serviceProvider : ServiceProvider;

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
	// add a ServiceProvider
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addServiceProvider(name, taxId, Claims, ProviderType, NetworkStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/ServiceProvider/create';
		const obj = {
			      		name: name,
      		taxId: taxId,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		ProviderType: ProviderType,
			NetworkStatus: NetworkStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ServiceProvider
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateServiceProvider(name, taxId, Claims, ProviderType, NetworkStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ServiceProvider/update/' + id;
		const obj = {
				      		name: name,
      		taxId: taxId,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		ProviderType: ProviderType,
			NetworkStatus: NetworkStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ServiceProvider
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteServiceProvider(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ServiceProvider/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ServiceProvider
	// returns the results untouched as an Observable ServiceProvider
	// ServiceProvider model
	// delegates via URI
	//********************************************************************
	getServiceProvider(id) : Observable<ServiceProvider> {
		const uri_ = this.apiUrl + '/ServiceProvider/load/' + id;

		return this.http.get<ServiceProvider>(uri_);
	}
	
	//********************************************************************
	// gets all ServiceProvider
	// returns the results untouched as JSON representation of an
	// Observable array of ServiceProvider models
	// delegates via URI
	//********************************************************************
	getServiceProviders() : Observable<ServiceProvider[]> {
		const uri_ = this.apiUrl + '/ServiceProvider/';

		return this
			.http.get<ServiceProvider[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more claimsIds as a Claims
	// to a ServiceProvider
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addClaims( serviceProviderId, claimsIds ): Observable<any> {

		// get the ServiceProvider
		this.loadHelper( serviceProviderId );

	// split on a comma with no spaces
	var idList = claimsIds.split(',')

	// iterate over array of claims ids
	idList.forEach(function (id) {
		// read the Claim
		var claim = new ClaimService(this.http).getClaim(id);
		// add the Claim if not already assigned
		if ( this.serviceProvider.claims.indexOf(claim) == -1 )
		this.serviceProvider.claims.push(claim);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more claimsIds as a Claims
	// from a ServiceProvider
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeClaims( serviceProviderId, claimsIds ): Observable<any> {

		// get the ServiceProvider
		this.loadHelper( serviceProviderId );


	// split on a comma with no spaces
	var idList 					= claimsIds.split(',');
	var claims 	= this.serviceProvider.claims;

	if ( claims != null && claimsIds != null ) {

		// iterate over array of claims ids
		claims.forEach(function (obj) {
			if ( claimsIds.indexOf(obj._id) > -1 ) {
				// remove the Claim
				this.serviceProvider.claims.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ServiceProvider
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ServiceProvider/update/' + this.serviceProvider;

	return  this.http.post(uri_, this.serviceProvider );
}

	//********************************************************************
	// loadHelper - internal helper to load a ServiceProvider
	//********************************************************************	
	loadHelper( id ) {
		this.getServiceProvider(id)
			.subscribe((res : ServiceProvider) => {
				this.serviceProvider = res;
			});
	}
}