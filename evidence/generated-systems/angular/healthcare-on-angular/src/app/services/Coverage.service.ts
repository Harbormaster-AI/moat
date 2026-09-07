import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Coverage} from '../models/Coverage';
import {PatientService} from '../services/Patient.service';
import {InsurancePlanService} from '../services/InsurancePlan.service';
import {ClaimService} from '../services/Claim.service';
import {AuthorizationService} from '../services/Authorization.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CoverageService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	coverage : Coverage;

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
	// add a Coverage
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCoverage(memberId, groupNumber, effectiveDate, endDate, Patient, Plan, Claims, Authorizations, CoverageType) : Observable<any> {
		const uri_ = this.apiUrl + '/Coverage/create';
		const obj = {
			      		memberId: memberId,
      		groupNumber: groupNumber,
      		effectiveDate: effectiveDate,
      		endDate: endDate,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Plan: Plan != null && Plan.length > 0 ? Plan : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		Authorizations: Authorizations != null && Authorizations.length > 0 ? Authorizations : null,
			CoverageType: CoverageType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Coverage
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCoverage(memberId, groupNumber, effectiveDate, endDate, Patient, Plan, Claims, Authorizations, CoverageType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Coverage/update/' + id;
		const obj = {
				      		memberId: memberId,
      		groupNumber: groupNumber,
      		effectiveDate: effectiveDate,
      		endDate: endDate,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Plan: Plan != null && Plan.length > 0 ? Plan : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		Authorizations: Authorizations != null && Authorizations.length > 0 ? Authorizations : null,
			CoverageType: CoverageType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Coverage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCoverage(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Coverage/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Coverage
	// returns the results untouched as an Observable Coverage
	// Coverage model
	// delegates via URI
	//********************************************************************
	getCoverage(id) : Observable<Coverage> {
		const uri_ = this.apiUrl + '/Coverage/load/' + id;

		return this.http.get<Coverage>(uri_);
	}
	
	//********************************************************************
	// gets all Coverage
	// returns the results untouched as JSON representation of an
	// Observable array of Coverage models
	// delegates via URI
	//********************************************************************
	getCoverages() : Observable<Coverage[]> {
		const uri_ = this.apiUrl + '/Coverage/';

		return this
			.http.get<Coverage[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Patient on a Coverage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( coverageId, _patientId ): Observable<any> {

		// get the Coverage from storage
		this.loadHelper( coverageId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.coverage.patient = tmp;

	// save the Coverage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a Coverage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( coverageId ): Observable<any> {

		// get the Coverage from storage
		this.loadHelper( coverageId );

	// assign Patient to null
	this.coverage.patient = null;

	// save the Coverage
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Plan on a Coverage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlan( coverageId, _planId ): Observable<any> {

		// get the Coverage from storage
		this.loadHelper( coverageId );

	// get the InsurancePlan from storage
	var tmp 	= new InsurancePlanService(this.http).getInsurancePlan(_planId);

	// assign the Plan
	this.coverage.plan = tmp;

	// save the Coverage
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plan on a Coverage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlan( coverageId ): Observable<any> {

		// get the Coverage from storage
		this.loadHelper( coverageId );

	// assign Plan to null
	this.coverage.plan = null;

	// save the Coverage
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more claimsIds as a Claims
	// to a Coverage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addClaims( coverageId, claimsIds ): Observable<any> {

		// get the Coverage
		this.loadHelper( coverageId );

	// split on a comma with no spaces
	var idList = claimsIds.split(',')

	// iterate over array of claims ids
	idList.forEach(function (id) {
		// read the Claim
		var claim = new ClaimService(this.http).getClaim(id);
		// add the Claim if not already assigned
		if ( this.coverage.claims.indexOf(claim) == -1 )
		this.coverage.claims.push(claim);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more claimsIds as a Claims
	// from a Coverage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeClaims( coverageId, claimsIds ): Observable<any> {

		// get the Coverage
		this.loadHelper( coverageId );


	// split on a comma with no spaces
	var idList 					= claimsIds.split(',');
	var claims 	= this.coverage.claims;

	if ( claims != null && claimsIds != null ) {

		// iterate over array of claims ids
		claims.forEach(function (obj) {
			if ( claimsIds.indexOf(obj._id) > -1 ) {
				// remove the Claim
				this.coverage.claims.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more authorizationsIds as a Authorizations
	// to a Coverage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAuthorizations( coverageId, authorizationsIds ): Observable<any> {

		// get the Coverage
		this.loadHelper( coverageId );

	// split on a comma with no spaces
	var idList = authorizationsIds.split(',')

	// iterate over array of authorizations ids
	idList.forEach(function (id) {
		// read the Authorization
		var authorization = new AuthorizationService(this.http).getAuthorization(id);
		// add the Authorization if not already assigned
		if ( this.coverage.authorizations.indexOf(authorization) == -1 )
		this.coverage.authorizations.push(authorization);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more authorizationsIds as a Authorizations
	// from a Coverage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAuthorizations( coverageId, authorizationsIds ): Observable<any> {

		// get the Coverage
		this.loadHelper( coverageId );


	// split on a comma with no spaces
	var idList 					= authorizationsIds.split(',');
	var authorizations 	= this.coverage.authorizations;

	if ( authorizations != null && authorizationsIds != null ) {

		// iterate over array of authorizations ids
		authorizations.forEach(function (obj) {
			if ( authorizationsIds.indexOf(obj._id) > -1 ) {
				// remove the Authorization
				this.coverage.authorizations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Coverage
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Coverage/update/' + this.coverage;

	return  this.http.post(uri_, this.coverage );
}

	//********************************************************************
	// loadHelper - internal helper to load a Coverage
	//********************************************************************	
	loadHelper( id ) {
		this.getCoverage(id)
			.subscribe((res : Coverage) => {
				this.coverage = res;
			});
	}
}