import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InsurancePayer} from '../models/InsurancePayer';
import {InsurancePlanService} from '../services/InsurancePlan.service';
import {ClaimService} from '../services/Claim.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InsurancePayerService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	insurancePayer : InsurancePayer;

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
	// add a InsurancePayer
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInsurancePayer(name, website, Plans, Claims, PayerType) : Observable<any> {
		const uri_ = this.apiUrl + '/InsurancePayer/create';
		const obj = {
			      		name: name,
      		website: website,
      		Plans: Plans != null && Plans.length > 0 ? Plans : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
			PayerType: PayerType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InsurancePayer
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInsurancePayer(name, website, Plans, Claims, PayerType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InsurancePayer/update/' + id;
		const obj = {
				      		name: name,
      		website: website,
      		Plans: Plans != null && Plans.length > 0 ? Plans : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
			PayerType: PayerType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InsurancePayer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInsurancePayer(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InsurancePayer/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InsurancePayer
	// returns the results untouched as an Observable InsurancePayer
	// InsurancePayer model
	// delegates via URI
	//********************************************************************
	getInsurancePayer(id) : Observable<InsurancePayer> {
		const uri_ = this.apiUrl + '/InsurancePayer/load/' + id;

		return this.http.get<InsurancePayer>(uri_);
	}
	
	//********************************************************************
	// gets all InsurancePayer
	// returns the results untouched as JSON representation of an
	// Observable array of InsurancePayer models
	// delegates via URI
	//********************************************************************
	getInsurancePayers() : Observable<InsurancePayer[]> {
		const uri_ = this.apiUrl + '/InsurancePayer/';

		return this
			.http.get<InsurancePayer[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more plansIds as a Plans
	// to a InsurancePayer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPlans( insurancePayerId, plansIds ): Observable<any> {

		// get the InsurancePayer
		this.loadHelper( insurancePayerId );

	// split on a comma with no spaces
	var idList = plansIds.split(',')

	// iterate over array of plans ids
	idList.forEach(function (id) {
		// read the InsurancePlan
		var insurancePlan = new InsurancePlanService(this.http).getInsurancePlan(id);
		// add the InsurancePlan if not already assigned
		if ( this.insurancePayer.plans.indexOf(insurancePlan) == -1 )
		this.insurancePayer.plans.push(insurancePlan);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more plansIds as a Plans
	// from a InsurancePayer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePlans( insurancePayerId, plansIds ): Observable<any> {

		// get the InsurancePayer
		this.loadHelper( insurancePayerId );


	// split on a comma with no spaces
	var idList 					= plansIds.split(',');
	var plans 	= this.insurancePayer.plans;

	if ( plans != null && plansIds != null ) {

		// iterate over array of plans ids
		plans.forEach(function (obj) {
			if ( plansIds.indexOf(obj._id) > -1 ) {
				// remove the InsurancePlan
				this.insurancePayer.plans.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more claimsIds as a Claims
	// to a InsurancePayer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addClaims( insurancePayerId, claimsIds ): Observable<any> {

		// get the InsurancePayer
		this.loadHelper( insurancePayerId );

	// split on a comma with no spaces
	var idList = claimsIds.split(',')

	// iterate over array of claims ids
	idList.forEach(function (id) {
		// read the Claim
		var claim = new ClaimService(this.http).getClaim(id);
		// add the Claim if not already assigned
		if ( this.insurancePayer.claims.indexOf(claim) == -1 )
		this.insurancePayer.claims.push(claim);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more claimsIds as a Claims
	// from a InsurancePayer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeClaims( insurancePayerId, claimsIds ): Observable<any> {

		// get the InsurancePayer
		this.loadHelper( insurancePayerId );


	// split on a comma with no spaces
	var idList 					= claimsIds.split(',');
	var claims 	= this.insurancePayer.claims;

	if ( claims != null && claimsIds != null ) {

		// iterate over array of claims ids
		claims.forEach(function (obj) {
			if ( claimsIds.indexOf(obj._id) > -1 ) {
				// remove the Claim
				this.insurancePayer.claims.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InsurancePayer
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InsurancePayer/update/' + this.insurancePayer;

	return  this.http.post(uri_, this.insurancePayer );
}

	//********************************************************************
	// loadHelper - internal helper to load a InsurancePayer
	//********************************************************************	
	loadHelper( id ) {
		this.getInsurancePayer(id)
			.subscribe((res : InsurancePayer) => {
				this.insurancePayer = res;
			});
	}
}