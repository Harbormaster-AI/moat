import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Policy} from '../models/Policy';
import {OrganizationService} from '../services/Organization.service';
import {PolicyAcknowledgementService} from '../services/PolicyAcknowledgement.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PolicyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	policy : Policy;

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
	// add a Policy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPolicy(policyNumber, name, effectiveDate, description, Organization, Acknowledgements) : Observable<any> {
		const uri_ = this.apiUrl + '/Policy/create';
		const obj = {
			      		policyNumber: policyNumber,
      		name: name,
      		effectiveDate: effectiveDate,
      		description: description,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
			Acknowledgements: Acknowledgements != null && Acknowledgements.length > 0 ? Acknowledgements : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePolicy(policyNumber, name, effectiveDate, description, Organization, Acknowledgements, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Policy/update/' + id;
		const obj = {
				      		policyNumber: policyNumber,
      		name: name,
      		effectiveDate: effectiveDate,
      		description: description,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
			Acknowledgements: Acknowledgements != null && Acknowledgements.length > 0 ? Acknowledgements : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePolicy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Policy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Policy
	// returns the results untouched as an Observable Policy
	// Policy model
	// delegates via URI
	//********************************************************************
	getPolicy(id) : Observable<Policy> {
		const uri_ = this.apiUrl + '/Policy/load/' + id;

		return this.http.get<Policy>(uri_);
	}
	
	//********************************************************************
	// gets all Policy
	// returns the results untouched as JSON representation of an
	// Observable array of Policy models
	// delegates via URI
	//********************************************************************
	getPolicys() : Observable<Policy[]> {
		const uri_ = this.apiUrl + '/Policy/';

		return this
			.http.get<Policy[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( policyId, _organizationId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.policy.organization = tmp;

	// save the Policy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( policyId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// assign Organization to null
	this.policy.organization = null;

	// save the Policy
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more acknowledgementsIds as a Acknowledgements
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAcknowledgements( policyId, acknowledgementsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = acknowledgementsIds.split(',')

	// iterate over array of acknowledgements ids
	idList.forEach(function (id) {
		// read the PolicyAcknowledgement
		var policyAcknowledgement = new PolicyAcknowledgementService(this.http).getPolicyAcknowledgement(id);
		// add the PolicyAcknowledgement if not already assigned
		if ( this.policy.acknowledgements.indexOf(policyAcknowledgement) == -1 )
		this.policy.acknowledgements.push(policyAcknowledgement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more acknowledgementsIds as a Acknowledgements
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAcknowledgements( policyId, acknowledgementsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= acknowledgementsIds.split(',');
	var acknowledgements 	= this.policy.acknowledgements;

	if ( acknowledgements != null && acknowledgementsIds != null ) {

		// iterate over array of acknowledgements ids
		acknowledgements.forEach(function (obj) {
			if ( acknowledgementsIds.indexOf(obj._id) > -1 ) {
				// remove the PolicyAcknowledgement
				this.policy.acknowledgements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Policy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Policy/update/' + this.policy;

	return  this.http.post(uri_, this.policy );
}

	//********************************************************************
	// loadHelper - internal helper to load a Policy
	//********************************************************************	
	loadHelper( id ) {
		this.getPolicy(id)
			.subscribe((res : Policy) => {
				this.policy = res;
			});
	}
}