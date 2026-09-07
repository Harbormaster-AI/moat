import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InsurancePlan} from '../models/InsurancePlan';
import {InsurancePayerService} from '../services/InsurancePayer.service';
import {CoverageService} from '../services/Coverage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InsurancePlanService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	insurancePlan : InsurancePlan;

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
	// add a InsurancePlan
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInsurancePlan(name, planCode, Payer, Coverages, PlanType) : Observable<any> {
		const uri_ = this.apiUrl + '/InsurancePlan/create';
		const obj = {
			      		name: name,
      		planCode: planCode,
      		Payer: Payer != null && Payer.length > 0 ? Payer : null,
      		Coverages: Coverages != null && Coverages.length > 0 ? Coverages : null,
			PlanType: PlanType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InsurancePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInsurancePlan(name, planCode, Payer, Coverages, PlanType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InsurancePlan/update/' + id;
		const obj = {
				      		name: name,
      		planCode: planCode,
      		Payer: Payer != null && Payer.length > 0 ? Payer : null,
      		Coverages: Coverages != null && Coverages.length > 0 ? Coverages : null,
			PlanType: PlanType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InsurancePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInsurancePlan(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InsurancePlan/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InsurancePlan
	// returns the results untouched as an Observable InsurancePlan
	// InsurancePlan model
	// delegates via URI
	//********************************************************************
	getInsurancePlan(id) : Observable<InsurancePlan> {
		const uri_ = this.apiUrl + '/InsurancePlan/load/' + id;

		return this.http.get<InsurancePlan>(uri_);
	}
	
	//********************************************************************
	// gets all InsurancePlan
	// returns the results untouched as JSON representation of an
	// Observable array of InsurancePlan models
	// delegates via URI
	//********************************************************************
	getInsurancePlans() : Observable<InsurancePlan[]> {
		const uri_ = this.apiUrl + '/InsurancePlan/';

		return this
			.http.get<InsurancePlan[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Payer on a InsurancePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPayer( insurancePlanId, _payerId ): Observable<any> {

		// get the InsurancePlan from storage
		this.loadHelper( insurancePlanId );

	// get the InsurancePayer from storage
	var tmp 	= new InsurancePayerService(this.http).getInsurancePayer(_payerId);

	// assign the Payer
	this.insurancePlan.payer = tmp;

	// save the InsurancePlan
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Payer on a InsurancePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPayer( insurancePlanId ): Observable<any> {

		// get the InsurancePlan from storage
		this.loadHelper( insurancePlanId );

	// assign Payer to null
	this.insurancePlan.payer = null;

	// save the InsurancePlan
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more coveragesIds as a Coverages
	// to a InsurancePlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCoverages( insurancePlanId, coveragesIds ): Observable<any> {

		// get the InsurancePlan
		this.loadHelper( insurancePlanId );

	// split on a comma with no spaces
	var idList = coveragesIds.split(',')

	// iterate over array of coverages ids
	idList.forEach(function (id) {
		// read the Coverage
		var coverage = new CoverageService(this.http).getCoverage(id);
		// add the Coverage if not already assigned
		if ( this.insurancePlan.coverages.indexOf(coverage) == -1 )
		this.insurancePlan.coverages.push(coverage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more coveragesIds as a Coverages
	// from a InsurancePlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCoverages( insurancePlanId, coveragesIds ): Observable<any> {

		// get the InsurancePlan
		this.loadHelper( insurancePlanId );


	// split on a comma with no spaces
	var idList 					= coveragesIds.split(',');
	var coverages 	= this.insurancePlan.coverages;

	if ( coverages != null && coveragesIds != null ) {

		// iterate over array of coverages ids
		coverages.forEach(function (obj) {
			if ( coveragesIds.indexOf(obj._id) > -1 ) {
				// remove the Coverage
				this.insurancePlan.coverages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InsurancePlan
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InsurancePlan/update/' + this.insurancePlan;

	return  this.http.post(uri_, this.insurancePlan );
}

	//********************************************************************
	// loadHelper - internal helper to load a InsurancePlan
	//********************************************************************	
	loadHelper( id ) {
		this.getInsurancePlan(id)
			.subscribe((res : InsurancePlan) => {
				this.insurancePlan = res;
			});
	}
}