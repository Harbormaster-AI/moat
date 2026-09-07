import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CompliancePolicy} from '../models/CompliancePolicy';
import {FinancialInstitutionService} from '../services/FinancialInstitution.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CompliancePolicyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	compliancePolicy : CompliancePolicy;

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
	// add a CompliancePolicy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCompliancePolicy(name, policyCode, description, Institution, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/CompliancePolicy/create';
		const obj = {
			      		name: name,
      		policyCode: policyCode,
      		description: description,
      		Institution: Institution != null && Institution.length > 0 ? Institution : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CompliancePolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCompliancePolicy(name, policyCode, description, Institution, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CompliancePolicy/update/' + id;
		const obj = {
				      		name: name,
      		policyCode: policyCode,
      		description: description,
      		Institution: Institution != null && Institution.length > 0 ? Institution : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CompliancePolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCompliancePolicy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CompliancePolicy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CompliancePolicy
	// returns the results untouched as an Observable CompliancePolicy
	// CompliancePolicy model
	// delegates via URI
	//********************************************************************
	getCompliancePolicy(id) : Observable<CompliancePolicy> {
		const uri_ = this.apiUrl + '/CompliancePolicy/load/' + id;

		return this.http.get<CompliancePolicy>(uri_);
	}
	
	//********************************************************************
	// gets all CompliancePolicy
	// returns the results untouched as JSON representation of an
	// Observable array of CompliancePolicy models
	// delegates via URI
	//********************************************************************
	getCompliancePolicys() : Observable<CompliancePolicy[]> {
		const uri_ = this.apiUrl + '/CompliancePolicy/';

		return this
			.http.get<CompliancePolicy[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Institution on a CompliancePolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInstitution( compliancePolicyId, _institutionId ): Observable<any> {

		// get the CompliancePolicy from storage
		this.loadHelper( compliancePolicyId );

	// get the FinancialInstitution from storage
	var tmp 	= new FinancialInstitutionService(this.http).getFinancialInstitution(_institutionId);

	// assign the Institution
	this.compliancePolicy.institution = tmp;

	// save the CompliancePolicy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Institution on a CompliancePolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInstitution( compliancePolicyId ): Observable<any> {

		// get the CompliancePolicy from storage
		this.loadHelper( compliancePolicyId );

	// assign Institution to null
	this.compliancePolicy.institution = null;

	// save the CompliancePolicy
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CompliancePolicy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CompliancePolicy/update/' + this.compliancePolicy;

	return  this.http.post(uri_, this.compliancePolicy );
}

	//********************************************************************
	// loadHelper - internal helper to load a CompliancePolicy
	//********************************************************************	
	loadHelper( id ) {
		this.getCompliancePolicy(id)
			.subscribe((res : CompliancePolicy) => {
				this.compliancePolicy = res;
			});
	}
}