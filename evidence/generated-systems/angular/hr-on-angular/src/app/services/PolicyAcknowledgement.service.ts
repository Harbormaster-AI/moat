import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PolicyAcknowledgement} from '../models/PolicyAcknowledgement';
import {PolicyService} from '../services/Policy.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PolicyAcknowledgementService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	policyAcknowledgement : PolicyAcknowledgement;

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
	// add a PolicyAcknowledgement
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPolicyAcknowledgement(acknowledgementDate, Policy, Employee, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/PolicyAcknowledgement/create';
		const obj = {
			      		acknowledgementDate: acknowledgementDate,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PolicyAcknowledgement
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePolicyAcknowledgement(acknowledgementDate, Policy, Employee, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PolicyAcknowledgement/update/' + id;
		const obj = {
				      		acknowledgementDate: acknowledgementDate,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PolicyAcknowledgement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePolicyAcknowledgement(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PolicyAcknowledgement/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PolicyAcknowledgement
	// returns the results untouched as an Observable PolicyAcknowledgement
	// PolicyAcknowledgement model
	// delegates via URI
	//********************************************************************
	getPolicyAcknowledgement(id) : Observable<PolicyAcknowledgement> {
		const uri_ = this.apiUrl + '/PolicyAcknowledgement/load/' + id;

		return this.http.get<PolicyAcknowledgement>(uri_);
	}
	
	//********************************************************************
	// gets all PolicyAcknowledgement
	// returns the results untouched as JSON representation of an
	// Observable array of PolicyAcknowledgement models
	// delegates via URI
	//********************************************************************
	getPolicyAcknowledgements() : Observable<PolicyAcknowledgement[]> {
		const uri_ = this.apiUrl + '/PolicyAcknowledgement/';

		return this
			.http.get<PolicyAcknowledgement[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Policy on a PolicyAcknowledgement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( policyAcknowledgementId, _policyId ): Observable<any> {

		// get the PolicyAcknowledgement from storage
		this.loadHelper( policyAcknowledgementId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.policyAcknowledgement.policy = tmp;

	// save the PolicyAcknowledgement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a PolicyAcknowledgement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( policyAcknowledgementId ): Observable<any> {

		// get the PolicyAcknowledgement from storage
		this.loadHelper( policyAcknowledgementId );

	// assign Policy to null
	this.policyAcknowledgement.policy = null;

	// save the PolicyAcknowledgement
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Employee on a PolicyAcknowledgement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( policyAcknowledgementId, _employeeId ): Observable<any> {

		// get the PolicyAcknowledgement from storage
		this.loadHelper( policyAcknowledgementId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.policyAcknowledgement.employee = tmp;

	// save the PolicyAcknowledgement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a PolicyAcknowledgement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( policyAcknowledgementId ): Observable<any> {

		// get the PolicyAcknowledgement from storage
		this.loadHelper( policyAcknowledgementId );

	// assign Employee to null
	this.policyAcknowledgement.employee = null;

	// save the PolicyAcknowledgement
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a PolicyAcknowledgement
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PolicyAcknowledgement/update/' + this.policyAcknowledgement;

	return  this.http.post(uri_, this.policyAcknowledgement );
}

	//********************************************************************
	// loadHelper - internal helper to load a PolicyAcknowledgement
	//********************************************************************	
	loadHelper( id ) {
		this.getPolicyAcknowledgement(id)
			.subscribe((res : PolicyAcknowledgement) => {
				this.policyAcknowledgement = res;
			});
	}
}