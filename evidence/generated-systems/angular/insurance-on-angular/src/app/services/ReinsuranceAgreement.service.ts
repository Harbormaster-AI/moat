import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ReinsuranceAgreement} from '../models/ReinsuranceAgreement';
import {InsurerService} from '../services/Insurer.service';
import {PolicyService} from '../services/Policy.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ReinsuranceAgreementService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	reinsuranceAgreement : ReinsuranceAgreement;

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
	// add a ReinsuranceAgreement
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addReinsuranceAgreement(agreementNumber, effectivePeriod, retention, limit, cessionPercentage, Insurer, Policies, ReinsuranceType, TreatyType) : Observable<any> {
		const uri_ = this.apiUrl + '/ReinsuranceAgreement/create';
		const obj = {
			      		agreementNumber: agreementNumber,
      		effectivePeriod: effectivePeriod,
      		retention: retention,
      		limit: limit,
      		cessionPercentage: cessionPercentage,
      		Insurer: Insurer != null && Insurer.length > 0 ? Insurer : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		ReinsuranceType: ReinsuranceType,
			TreatyType: TreatyType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ReinsuranceAgreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateReinsuranceAgreement(agreementNumber, effectivePeriod, retention, limit, cessionPercentage, Insurer, Policies, ReinsuranceType, TreatyType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ReinsuranceAgreement/update/' + id;
		const obj = {
				      		agreementNumber: agreementNumber,
      		effectivePeriod: effectivePeriod,
      		retention: retention,
      		limit: limit,
      		cessionPercentage: cessionPercentage,
      		Insurer: Insurer != null && Insurer.length > 0 ? Insurer : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		ReinsuranceType: ReinsuranceType,
			TreatyType: TreatyType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ReinsuranceAgreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteReinsuranceAgreement(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ReinsuranceAgreement/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ReinsuranceAgreement
	// returns the results untouched as an Observable ReinsuranceAgreement
	// ReinsuranceAgreement model
	// delegates via URI
	//********************************************************************
	getReinsuranceAgreement(id) : Observable<ReinsuranceAgreement> {
		const uri_ = this.apiUrl + '/ReinsuranceAgreement/load/' + id;

		return this.http.get<ReinsuranceAgreement>(uri_);
	}
	
	//********************************************************************
	// gets all ReinsuranceAgreement
	// returns the results untouched as JSON representation of an
	// Observable array of ReinsuranceAgreement models
	// delegates via URI
	//********************************************************************
	getReinsuranceAgreements() : Observable<ReinsuranceAgreement[]> {
		const uri_ = this.apiUrl + '/ReinsuranceAgreement/';

		return this
			.http.get<ReinsuranceAgreement[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Insurer on a ReinsuranceAgreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInsurer( reinsuranceAgreementId, _insurerId ): Observable<any> {

		// get the ReinsuranceAgreement from storage
		this.loadHelper( reinsuranceAgreementId );

	// get the Insurer from storage
	var tmp 	= new InsurerService(this.http).getInsurer(_insurerId);

	// assign the Insurer
	this.reinsuranceAgreement.insurer = tmp;

	// save the ReinsuranceAgreement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Insurer on a ReinsuranceAgreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInsurer( reinsuranceAgreementId ): Observable<any> {

		// get the ReinsuranceAgreement from storage
		this.loadHelper( reinsuranceAgreementId );

	// assign Insurer to null
	this.reinsuranceAgreement.insurer = null;

	// save the ReinsuranceAgreement
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a ReinsuranceAgreement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( reinsuranceAgreementId, policiesIds ): Observable<any> {

		// get the ReinsuranceAgreement
		this.loadHelper( reinsuranceAgreementId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.reinsuranceAgreement.policies.indexOf(policy) == -1 )
		this.reinsuranceAgreement.policies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a ReinsuranceAgreement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( reinsuranceAgreementId, policiesIds ): Observable<any> {

		// get the ReinsuranceAgreement
		this.loadHelper( reinsuranceAgreementId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.reinsuranceAgreement.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.reinsuranceAgreement.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ReinsuranceAgreement
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ReinsuranceAgreement/update/' + this.reinsuranceAgreement;

	return  this.http.post(uri_, this.reinsuranceAgreement );
}

	//********************************************************************
	// loadHelper - internal helper to load a ReinsuranceAgreement
	//********************************************************************	
	loadHelper( id ) {
		this.getReinsuranceAgreement(id)
			.subscribe((res : ReinsuranceAgreement) => {
				this.reinsuranceAgreement = res;
			});
	}
}