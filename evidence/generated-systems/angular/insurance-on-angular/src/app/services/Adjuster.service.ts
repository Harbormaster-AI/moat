import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Adjuster} from '../models/Adjuster';
import {ClaimService} from '../services/Claim.service';
import {ServiceProviderService} from '../services/ServiceProvider.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AdjusterService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	adjuster : Adjuster;

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
	// add a Adjuster
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAdjuster(firstName, lastName, licenseNumber, Claims, ServiceProviders, AdjusterType) : Observable<any> {
		const uri_ = this.apiUrl + '/Adjuster/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		licenseNumber: licenseNumber,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		ServiceProviders: ServiceProviders != null && ServiceProviders.length > 0 ? ServiceProviders : null,
			AdjusterType: AdjusterType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Adjuster
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAdjuster(firstName, lastName, licenseNumber, Claims, ServiceProviders, AdjusterType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Adjuster/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		licenseNumber: licenseNumber,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		ServiceProviders: ServiceProviders != null && ServiceProviders.length > 0 ? ServiceProviders : null,
			AdjusterType: AdjusterType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Adjuster
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAdjuster(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Adjuster/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Adjuster
	// returns the results untouched as an Observable Adjuster
	// Adjuster model
	// delegates via URI
	//********************************************************************
	getAdjuster(id) : Observable<Adjuster> {
		const uri_ = this.apiUrl + '/Adjuster/load/' + id;

		return this.http.get<Adjuster>(uri_);
	}
	
	//********************************************************************
	// gets all Adjuster
	// returns the results untouched as JSON representation of an
	// Observable array of Adjuster models
	// delegates via URI
	//********************************************************************
	getAdjusters() : Observable<Adjuster[]> {
		const uri_ = this.apiUrl + '/Adjuster/';

		return this
			.http.get<Adjuster[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more claimsIds as a Claims
	// to a Adjuster
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addClaims( adjusterId, claimsIds ): Observable<any> {

		// get the Adjuster
		this.loadHelper( adjusterId );

	// split on a comma with no spaces
	var idList = claimsIds.split(',')

	// iterate over array of claims ids
	idList.forEach(function (id) {
		// read the Claim
		var claim = new ClaimService(this.http).getClaim(id);
		// add the Claim if not already assigned
		if ( this.adjuster.claims.indexOf(claim) == -1 )
		this.adjuster.claims.push(claim);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more claimsIds as a Claims
	// from a Adjuster
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeClaims( adjusterId, claimsIds ): Observable<any> {

		// get the Adjuster
		this.loadHelper( adjusterId );


	// split on a comma with no spaces
	var idList 					= claimsIds.split(',');
	var claims 	= this.adjuster.claims;

	if ( claims != null && claimsIds != null ) {

		// iterate over array of claims ids
		claims.forEach(function (obj) {
			if ( claimsIds.indexOf(obj._id) > -1 ) {
				// remove the Claim
				this.adjuster.claims.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more serviceProvidersIds as a ServiceProviders
	// to a Adjuster
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addServiceProviders( adjusterId, serviceProvidersIds ): Observable<any> {

		// get the Adjuster
		this.loadHelper( adjusterId );

	// split on a comma with no spaces
	var idList = serviceProvidersIds.split(',')

	// iterate over array of serviceProviders ids
	idList.forEach(function (id) {
		// read the ServiceProvider
		var serviceProvider = new ServiceProviderService(this.http).getServiceProvider(id);
		// add the ServiceProvider if not already assigned
		if ( this.adjuster.serviceProviders.indexOf(serviceProvider) == -1 )
		this.adjuster.serviceProviders.push(serviceProvider);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serviceProvidersIds as a ServiceProviders
	// from a Adjuster
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeServiceProviders( adjusterId, serviceProvidersIds ): Observable<any> {

		// get the Adjuster
		this.loadHelper( adjusterId );


	// split on a comma with no spaces
	var idList 					= serviceProvidersIds.split(',');
	var serviceProviders 	= this.adjuster.serviceProviders;

	if ( serviceProviders != null && serviceProvidersIds != null ) {

		// iterate over array of serviceProviders ids
		serviceProviders.forEach(function (obj) {
			if ( serviceProvidersIds.indexOf(obj._id) > -1 ) {
				// remove the ServiceProvider
				this.adjuster.serviceProviders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Adjuster
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Adjuster/update/' + this.adjuster;

	return  this.http.post(uri_, this.adjuster );
}

	//********************************************************************
	// loadHelper - internal helper to load a Adjuster
	//********************************************************************	
	loadHelper( id ) {
		this.getAdjuster(id)
			.subscribe((res : Adjuster) => {
				this.adjuster = res;
			});
	}
}