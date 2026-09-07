import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InsuredObject} from '../models/InsuredObject';
import {PolicyService} from '../services/Policy.service';
import {PolicyCoverageService} from '../services/PolicyCoverage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InsuredObjectService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	insuredObject : InsuredObject;

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
	// add a InsuredObject
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInsuredObject(description, serialOrId, primaryAddress, Policy, Coverages, ObjectType) : Observable<any> {
		const uri_ = this.apiUrl + '/InsuredObject/create';
		const obj = {
			      		description: description,
      		serialOrId: serialOrId,
      		primaryAddress: primaryAddress,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Coverages: Coverages != null && Coverages.length > 0 ? Coverages : null,
			ObjectType: ObjectType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InsuredObject
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInsuredObject(description, serialOrId, primaryAddress, Policy, Coverages, ObjectType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InsuredObject/update/' + id;
		const obj = {
				      		description: description,
      		serialOrId: serialOrId,
      		primaryAddress: primaryAddress,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Coverages: Coverages != null && Coverages.length > 0 ? Coverages : null,
			ObjectType: ObjectType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InsuredObject
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInsuredObject(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InsuredObject/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InsuredObject
	// returns the results untouched as an Observable InsuredObject
	// InsuredObject model
	// delegates via URI
	//********************************************************************
	getInsuredObject(id) : Observable<InsuredObject> {
		const uri_ = this.apiUrl + '/InsuredObject/load/' + id;

		return this.http.get<InsuredObject>(uri_);
	}
	
	//********************************************************************
	// gets all InsuredObject
	// returns the results untouched as JSON representation of an
	// Observable array of InsuredObject models
	// delegates via URI
	//********************************************************************
	getInsuredObjects() : Observable<InsuredObject[]> {
		const uri_ = this.apiUrl + '/InsuredObject/';

		return this
			.http.get<InsuredObject[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Policy on a InsuredObject
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( insuredObjectId, _policyId ): Observable<any> {

		// get the InsuredObject from storage
		this.loadHelper( insuredObjectId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.insuredObject.policy = tmp;

	// save the InsuredObject
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a InsuredObject
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( insuredObjectId ): Observable<any> {

		// get the InsuredObject from storage
		this.loadHelper( insuredObjectId );

	// assign Policy to null
	this.insuredObject.policy = null;

	// save the InsuredObject
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more coveragesIds as a Coverages
	// to a InsuredObject
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCoverages( insuredObjectId, coveragesIds ): Observable<any> {

		// get the InsuredObject
		this.loadHelper( insuredObjectId );

	// split on a comma with no spaces
	var idList = coveragesIds.split(',')

	// iterate over array of coverages ids
	idList.forEach(function (id) {
		// read the PolicyCoverage
		var policyCoverage = new PolicyCoverageService(this.http).getPolicyCoverage(id);
		// add the PolicyCoverage if not already assigned
		if ( this.insuredObject.coverages.indexOf(policyCoverage) == -1 )
		this.insuredObject.coverages.push(policyCoverage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more coveragesIds as a Coverages
	// from a InsuredObject
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCoverages( insuredObjectId, coveragesIds ): Observable<any> {

		// get the InsuredObject
		this.loadHelper( insuredObjectId );


	// split on a comma with no spaces
	var idList 					= coveragesIds.split(',');
	var coverages 	= this.insuredObject.coverages;

	if ( coverages != null && coveragesIds != null ) {

		// iterate over array of coverages ids
		coverages.forEach(function (obj) {
			if ( coveragesIds.indexOf(obj._id) > -1 ) {
				// remove the PolicyCoverage
				this.insuredObject.coverages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InsuredObject
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InsuredObject/update/' + this.insuredObject;

	return  this.http.post(uri_, this.insuredObject );
}

	//********************************************************************
	// loadHelper - internal helper to load a InsuredObject
	//********************************************************************	
	loadHelper( id ) {
		this.getInsuredObject(id)
			.subscribe((res : InsuredObject) => {
				this.insuredObject = res;
			});
	}
}