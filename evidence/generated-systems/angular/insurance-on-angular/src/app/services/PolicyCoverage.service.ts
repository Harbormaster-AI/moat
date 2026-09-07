import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PolicyCoverage} from '../models/PolicyCoverage';
import {PolicyService} from '../services/Policy.service';
import {InsuredObjectService} from '../services/InsuredObject.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PolicyCoverageService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	policyCoverage : PolicyCoverage;

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
	// add a PolicyCoverage
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPolicyCoverage(limit, deductible, premium, Policy, InsuredObjects, CoverageType) : Observable<any> {
		const uri_ = this.apiUrl + '/PolicyCoverage/create';
		const obj = {
			      		limit: limit,
      		deductible: deductible,
      		premium: premium,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		InsuredObjects: InsuredObjects != null && InsuredObjects.length > 0 ? InsuredObjects : null,
			CoverageType: CoverageType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PolicyCoverage
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePolicyCoverage(limit, deductible, premium, Policy, InsuredObjects, CoverageType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PolicyCoverage/update/' + id;
		const obj = {
				      		limit: limit,
      		deductible: deductible,
      		premium: premium,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		InsuredObjects: InsuredObjects != null && InsuredObjects.length > 0 ? InsuredObjects : null,
			CoverageType: CoverageType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PolicyCoverage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePolicyCoverage(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PolicyCoverage/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PolicyCoverage
	// returns the results untouched as an Observable PolicyCoverage
	// PolicyCoverage model
	// delegates via URI
	//********************************************************************
	getPolicyCoverage(id) : Observable<PolicyCoverage> {
		const uri_ = this.apiUrl + '/PolicyCoverage/load/' + id;

		return this.http.get<PolicyCoverage>(uri_);
	}
	
	//********************************************************************
	// gets all PolicyCoverage
	// returns the results untouched as JSON representation of an
	// Observable array of PolicyCoverage models
	// delegates via URI
	//********************************************************************
	getPolicyCoverages() : Observable<PolicyCoverage[]> {
		const uri_ = this.apiUrl + '/PolicyCoverage/';

		return this
			.http.get<PolicyCoverage[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Policy on a PolicyCoverage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( policyCoverageId, _policyId ): Observable<any> {

		// get the PolicyCoverage from storage
		this.loadHelper( policyCoverageId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.policyCoverage.policy = tmp;

	// save the PolicyCoverage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a PolicyCoverage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( policyCoverageId ): Observable<any> {

		// get the PolicyCoverage from storage
		this.loadHelper( policyCoverageId );

	// assign Policy to null
	this.policyCoverage.policy = null;

	// save the PolicyCoverage
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more insuredObjectsIds as a InsuredObjects
	// to a PolicyCoverage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInsuredObjects( policyCoverageId, insuredObjectsIds ): Observable<any> {

		// get the PolicyCoverage
		this.loadHelper( policyCoverageId );

	// split on a comma with no spaces
	var idList = insuredObjectsIds.split(',')

	// iterate over array of insuredObjects ids
	idList.forEach(function (id) {
		// read the InsuredObject
		var insuredObject = new InsuredObjectService(this.http).getInsuredObject(id);
		// add the InsuredObject if not already assigned
		if ( this.policyCoverage.insuredObjects.indexOf(insuredObject) == -1 )
		this.policyCoverage.insuredObjects.push(insuredObject);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more insuredObjectsIds as a InsuredObjects
	// from a PolicyCoverage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInsuredObjects( policyCoverageId, insuredObjectsIds ): Observable<any> {

		// get the PolicyCoverage
		this.loadHelper( policyCoverageId );


	// split on a comma with no spaces
	var idList 					= insuredObjectsIds.split(',');
	var insuredObjects 	= this.policyCoverage.insuredObjects;

	if ( insuredObjects != null && insuredObjectsIds != null ) {

		// iterate over array of insuredObjects ids
		insuredObjects.forEach(function (obj) {
			if ( insuredObjectsIds.indexOf(obj._id) > -1 ) {
				// remove the InsuredObject
				this.policyCoverage.insuredObjects.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PolicyCoverage
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PolicyCoverage/update/' + this.policyCoverage;

	return  this.http.post(uri_, this.policyCoverage );
}

	//********************************************************************
	// loadHelper - internal helper to load a PolicyCoverage
	//********************************************************************	
	loadHelper( id ) {
		this.getPolicyCoverage(id)
			.subscribe((res : PolicyCoverage) => {
				this.policyCoverage = res;
			});
	}
}