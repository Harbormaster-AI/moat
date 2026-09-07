import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ComplianceRequirement} from '../models/ComplianceRequirement';
import {ComplianceProgramService} from '../services/ComplianceProgram.service';
import {PolicyService} from '../services/Policy.service';
import {ControlService} from '../services/Control.service';
import {ObligationService} from '../services/Obligation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ComplianceRequirementService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	complianceRequirement : ComplianceRequirement;

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
	// add a ComplianceRequirement
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addComplianceRequirement(name, source, citation, ComplianceProgram, Policies, Controls, Obligations, Applicability, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/ComplianceRequirement/create';
		const obj = {
			      		name: name,
      		source: source,
      		citation: citation,
      		ComplianceProgram: ComplianceProgram != null && ComplianceProgram.length > 0 ? ComplianceProgram : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
      		Obligations: Obligations != null && Obligations.length > 0 ? Obligations : null,
      		Applicability: Applicability,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ComplianceRequirement
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateComplianceRequirement(name, source, citation, ComplianceProgram, Policies, Controls, Obligations, Applicability, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ComplianceRequirement/update/' + id;
		const obj = {
				      		name: name,
      		source: source,
      		citation: citation,
      		ComplianceProgram: ComplianceProgram != null && ComplianceProgram.length > 0 ? ComplianceProgram : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
      		Obligations: Obligations != null && Obligations.length > 0 ? Obligations : null,
      		Applicability: Applicability,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ComplianceRequirement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteComplianceRequirement(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ComplianceRequirement/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ComplianceRequirement
	// returns the results untouched as an Observable ComplianceRequirement
	// ComplianceRequirement model
	// delegates via URI
	//********************************************************************
	getComplianceRequirement(id) : Observable<ComplianceRequirement> {
		const uri_ = this.apiUrl + '/ComplianceRequirement/load/' + id;

		return this.http.get<ComplianceRequirement>(uri_);
	}
	
	//********************************************************************
	// gets all ComplianceRequirement
	// returns the results untouched as JSON representation of an
	// Observable array of ComplianceRequirement models
	// delegates via URI
	//********************************************************************
	getComplianceRequirements() : Observable<ComplianceRequirement[]> {
		const uri_ = this.apiUrl + '/ComplianceRequirement/';

		return this
			.http.get<ComplianceRequirement[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ComplianceProgram on a ComplianceRequirement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignComplianceProgram( complianceRequirementId, _complianceProgramId ): Observable<any> {

		// get the ComplianceRequirement from storage
		this.loadHelper( complianceRequirementId );

	// get the ComplianceProgram from storage
	var tmp 	= new ComplianceProgramService(this.http).getComplianceProgram(_complianceProgramId);

	// assign the ComplianceProgram
	this.complianceRequirement.complianceProgram = tmp;

	// save the ComplianceRequirement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ComplianceProgram on a ComplianceRequirement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignComplianceProgram( complianceRequirementId ): Observable<any> {

		// get the ComplianceRequirement from storage
		this.loadHelper( complianceRequirementId );

	// assign ComplianceProgram to null
	this.complianceRequirement.complianceProgram = null;

	// save the ComplianceRequirement
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a ComplianceRequirement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( complianceRequirementId, policiesIds ): Observable<any> {

		// get the ComplianceRequirement
		this.loadHelper( complianceRequirementId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.complianceRequirement.policies.indexOf(policy) == -1 )
		this.complianceRequirement.policies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a ComplianceRequirement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( complianceRequirementId, policiesIds ): Observable<any> {

		// get the ComplianceRequirement
		this.loadHelper( complianceRequirementId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.complianceRequirement.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.complianceRequirement.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more controlsIds as a Controls
	// to a ComplianceRequirement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addControls( complianceRequirementId, controlsIds ): Observable<any> {

		// get the ComplianceRequirement
		this.loadHelper( complianceRequirementId );

	// split on a comma with no spaces
	var idList = controlsIds.split(',')

	// iterate over array of controls ids
	idList.forEach(function (id) {
		// read the Control
		var control = new ControlService(this.http).getControl(id);
		// add the Control if not already assigned
		if ( this.complianceRequirement.controls.indexOf(control) == -1 )
		this.complianceRequirement.controls.push(control);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more controlsIds as a Controls
	// from a ComplianceRequirement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeControls( complianceRequirementId, controlsIds ): Observable<any> {

		// get the ComplianceRequirement
		this.loadHelper( complianceRequirementId );


	// split on a comma with no spaces
	var idList 					= controlsIds.split(',');
	var controls 	= this.complianceRequirement.controls;

	if ( controls != null && controlsIds != null ) {

		// iterate over array of controls ids
		controls.forEach(function (obj) {
			if ( controlsIds.indexOf(obj._id) > -1 ) {
				// remove the Control
				this.complianceRequirement.controls.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more obligationsIds as a Obligations
	// to a ComplianceRequirement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addObligations( complianceRequirementId, obligationsIds ): Observable<any> {

		// get the ComplianceRequirement
		this.loadHelper( complianceRequirementId );

	// split on a comma with no spaces
	var idList = obligationsIds.split(',')

	// iterate over array of obligations ids
	idList.forEach(function (id) {
		// read the Obligation
		var obligation = new ObligationService(this.http).getObligation(id);
		// add the Obligation if not already assigned
		if ( this.complianceRequirement.obligations.indexOf(obligation) == -1 )
		this.complianceRequirement.obligations.push(obligation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more obligationsIds as a Obligations
	// from a ComplianceRequirement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeObligations( complianceRequirementId, obligationsIds ): Observable<any> {

		// get the ComplianceRequirement
		this.loadHelper( complianceRequirementId );


	// split on a comma with no spaces
	var idList 					= obligationsIds.split(',');
	var obligations 	= this.complianceRequirement.obligations;

	if ( obligations != null && obligationsIds != null ) {

		// iterate over array of obligations ids
		obligations.forEach(function (obj) {
			if ( obligationsIds.indexOf(obj._id) > -1 ) {
				// remove the Obligation
				this.complianceRequirement.obligations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ComplianceRequirement
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ComplianceRequirement/update/' + this.complianceRequirement;

	return  this.http.post(uri_, this.complianceRequirement );
}

	//********************************************************************
	// loadHelper - internal helper to load a ComplianceRequirement
	//********************************************************************	
	loadHelper( id ) {
		this.getComplianceRequirement(id)
			.subscribe((res : ComplianceRequirement) => {
				this.complianceRequirement = res;
			});
	}
}