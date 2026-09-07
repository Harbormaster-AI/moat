import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Endorsement} from '../models/Endorsement';
import {PolicyService} from '../services/Policy.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EndorsementService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	endorsement : Endorsement;

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
	// add a Endorsement
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEndorsement(endorsementNumber, effectiveDate, description, Policy) : Observable<any> {
		const uri_ = this.apiUrl + '/Endorsement/create';
		const obj = {
			      		endorsementNumber: endorsementNumber,
      		effectiveDate: effectiveDate,
      		description: description,
			Policy: Policy != null && Policy.length > 0 ? Policy : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Endorsement
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEndorsement(endorsementNumber, effectiveDate, description, Policy, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Endorsement/update/' + id;
		const obj = {
				      		endorsementNumber: endorsementNumber,
      		effectiveDate: effectiveDate,
      		description: description,
			Policy: Policy != null && Policy.length > 0 ? Policy : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Endorsement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEndorsement(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Endorsement/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Endorsement
	// returns the results untouched as an Observable Endorsement
	// Endorsement model
	// delegates via URI
	//********************************************************************
	getEndorsement(id) : Observable<Endorsement> {
		const uri_ = this.apiUrl + '/Endorsement/load/' + id;

		return this.http.get<Endorsement>(uri_);
	}
	
	//********************************************************************
	// gets all Endorsement
	// returns the results untouched as JSON representation of an
	// Observable array of Endorsement models
	// delegates via URI
	//********************************************************************
	getEndorsements() : Observable<Endorsement[]> {
		const uri_ = this.apiUrl + '/Endorsement/';

		return this
			.http.get<Endorsement[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Policy on a Endorsement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( endorsementId, _policyId ): Observable<any> {

		// get the Endorsement from storage
		this.loadHelper( endorsementId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.endorsement.policy = tmp;

	// save the Endorsement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Endorsement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( endorsementId ): Observable<any> {

		// get the Endorsement from storage
		this.loadHelper( endorsementId );

	// assign Policy to null
	this.endorsement.policy = null;

	// save the Endorsement
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Endorsement
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Endorsement/update/' + this.endorsement;

	return  this.http.post(uri_, this.endorsement );
}

	//********************************************************************
	// loadHelper - internal helper to load a Endorsement
	//********************************************************************	
	loadHelper( id ) {
		this.getEndorsement(id)
			.subscribe((res : Endorsement) => {
				this.endorsement = res;
			});
	}
}