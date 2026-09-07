import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Policy} from '../models/Policy';
import {OrganizationService} from '../services/Organization.service';
import {PersonService} from '../services/Person.service';
import {ComplianceRequirementService} from '../services/ComplianceRequirement.service';
import {ControlService} from '../services/Control.service';
import {ProcedureService} from '../services/Procedure.service';
import {Exception_Service} from '../services/Exception_.service';
import {AttestationService} from '../services/Attestation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PolicyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	policy : Policy;

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
	// add a Policy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPolicy(title, versionLabel, approvalDate, nextReviewDate, documentUrl, Organization, Owners, RelatedRequirements, Controls, Procedures, Exceptions, Attestations, PolicyType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Policy/create';
		const obj = {
			      		title: title,
      		versionLabel: versionLabel,
      		approvalDate: approvalDate,
      		nextReviewDate: nextReviewDate,
      		documentUrl: documentUrl,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Owners: Owners != null && Owners.length > 0 ? Owners : null,
      		RelatedRequirements: RelatedRequirements != null && RelatedRequirements.length > 0 ? RelatedRequirements : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
      		Procedures: Procedures != null && Procedures.length > 0 ? Procedures : null,
      		Exceptions: Exceptions != null && Exceptions.length > 0 ? Exceptions : null,
      		Attestations: Attestations != null && Attestations.length > 0 ? Attestations : null,
      		PolicyType: PolicyType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePolicy(title, versionLabel, approvalDate, nextReviewDate, documentUrl, Organization, Owners, RelatedRequirements, Controls, Procedures, Exceptions, Attestations, PolicyType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Policy/update/' + id;
		const obj = {
				      		title: title,
      		versionLabel: versionLabel,
      		approvalDate: approvalDate,
      		nextReviewDate: nextReviewDate,
      		documentUrl: documentUrl,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Owners: Owners != null && Owners.length > 0 ? Owners : null,
      		RelatedRequirements: RelatedRequirements != null && RelatedRequirements.length > 0 ? RelatedRequirements : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
      		Procedures: Procedures != null && Procedures.length > 0 ? Procedures : null,
      		Exceptions: Exceptions != null && Exceptions.length > 0 ? Exceptions : null,
      		Attestations: Attestations != null && Attestations.length > 0 ? Attestations : null,
      		PolicyType: PolicyType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePolicy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Policy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Policy
	// returns the results untouched as an Observable Policy
	// Policy model
	// delegates via URI
	//********************************************************************
	getPolicy(id) : Observable<Policy> {
		const uri_ = this.apiUrl + '/Policy/load/' + id;

		return this.http.get<Policy>(uri_);
	}
	
	//********************************************************************
	// gets all Policy
	// returns the results untouched as JSON representation of an
	// Observable array of Policy models
	// delegates via URI
	//********************************************************************
	getPolicys() : Observable<Policy[]> {
		const uri_ = this.apiUrl + '/Policy/';

		return this
			.http.get<Policy[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( policyId, _organizationId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.policy.organization = tmp;

	// save the Policy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Policy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( policyId ): Observable<any> {

		// get the Policy from storage
		this.loadHelper( policyId );

	// assign Organization to null
	this.policy.organization = null;

	// save the Policy
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more ownersIds as a Owners
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOwners( policyId, ownersIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = ownersIds.split(',')

	// iterate over array of owners ids
	idList.forEach(function (id) {
		// read the Person
		var person = new PersonService(this.http).getPerson(id);
		// add the Person if not already assigned
		if ( this.policy.owners.indexOf(person) == -1 )
		this.policy.owners.push(person);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ownersIds as a Owners
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOwners( policyId, ownersIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= ownersIds.split(',');
	var owners 	= this.policy.owners;

	if ( owners != null && ownersIds != null ) {

		// iterate over array of owners ids
		owners.forEach(function (obj) {
			if ( ownersIds.indexOf(obj._id) > -1 ) {
				// remove the Person
				this.policy.owners.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more relatedRequirementsIds as a RelatedRequirements
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRelatedRequirements( policyId, relatedRequirementsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = relatedRequirementsIds.split(',')

	// iterate over array of relatedRequirements ids
	idList.forEach(function (id) {
		// read the ComplianceRequirement
		var complianceRequirement = new ComplianceRequirementService(this.http).getComplianceRequirement(id);
		// add the ComplianceRequirement if not already assigned
		if ( this.policy.relatedRequirements.indexOf(complianceRequirement) == -1 )
		this.policy.relatedRequirements.push(complianceRequirement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more relatedRequirementsIds as a RelatedRequirements
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRelatedRequirements( policyId, relatedRequirementsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= relatedRequirementsIds.split(',');
	var relatedRequirements 	= this.policy.relatedRequirements;

	if ( relatedRequirements != null && relatedRequirementsIds != null ) {

		// iterate over array of relatedRequirements ids
		relatedRequirements.forEach(function (obj) {
			if ( relatedRequirementsIds.indexOf(obj._id) > -1 ) {
				// remove the ComplianceRequirement
				this.policy.relatedRequirements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more controlsIds as a Controls
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addControls( policyId, controlsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = controlsIds.split(',')

	// iterate over array of controls ids
	idList.forEach(function (id) {
		// read the Control
		var control = new ControlService(this.http).getControl(id);
		// add the Control if not already assigned
		if ( this.policy.controls.indexOf(control) == -1 )
		this.policy.controls.push(control);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more controlsIds as a Controls
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeControls( policyId, controlsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= controlsIds.split(',');
	var controls 	= this.policy.controls;

	if ( controls != null && controlsIds != null ) {

		// iterate over array of controls ids
		controls.forEach(function (obj) {
			if ( controlsIds.indexOf(obj._id) > -1 ) {
				// remove the Control
				this.policy.controls.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more proceduresIds as a Procedures
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcedures( policyId, proceduresIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = proceduresIds.split(',')

	// iterate over array of procedures ids
	idList.forEach(function (id) {
		// read the Procedure
		var procedure = new ProcedureService(this.http).getProcedure(id);
		// add the Procedure if not already assigned
		if ( this.policy.procedures.indexOf(procedure) == -1 )
		this.policy.procedures.push(procedure);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more proceduresIds as a Procedures
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcedures( policyId, proceduresIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= proceduresIds.split(',');
	var procedures 	= this.policy.procedures;

	if ( procedures != null && proceduresIds != null ) {

		// iterate over array of procedures ids
		procedures.forEach(function (obj) {
			if ( proceduresIds.indexOf(obj._id) > -1 ) {
				// remove the Procedure
				this.policy.procedures.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more exceptionsIds as a Exceptions
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addExceptions( policyId, exceptionsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = exceptionsIds.split(',')

	// iterate over array of exceptions ids
	idList.forEach(function (id) {
		// read the Exception_
		var exception_ = new Exception_Service(this.http).getException_(id);
		// add the Exception_ if not already assigned
		if ( this.policy.exceptions.indexOf(exception_) == -1 )
		this.policy.exceptions.push(exception_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more exceptionsIds as a Exceptions
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeExceptions( policyId, exceptionsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= exceptionsIds.split(',');
	var exceptions 	= this.policy.exceptions;

	if ( exceptions != null && exceptionsIds != null ) {

		// iterate over array of exceptions ids
		exceptions.forEach(function (obj) {
			if ( exceptionsIds.indexOf(obj._id) > -1 ) {
				// remove the Exception_
				this.policy.exceptions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more attestationsIds as a Attestations
	// to a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAttestations( policyId, attestationsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );

	// split on a comma with no spaces
	var idList = attestationsIds.split(',')

	// iterate over array of attestations ids
	idList.forEach(function (id) {
		// read the Attestation
		var attestation = new AttestationService(this.http).getAttestation(id);
		// add the Attestation if not already assigned
		if ( this.policy.attestations.indexOf(attestation) == -1 )
		this.policy.attestations.push(attestation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more attestationsIds as a Attestations
	// from a Policy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAttestations( policyId, attestationsIds ): Observable<any> {

		// get the Policy
		this.loadHelper( policyId );


	// split on a comma with no spaces
	var idList 					= attestationsIds.split(',');
	var attestations 	= this.policy.attestations;

	if ( attestations != null && attestationsIds != null ) {

		// iterate over array of attestations ids
		attestations.forEach(function (obj) {
			if ( attestationsIds.indexOf(obj._id) > -1 ) {
				// remove the Attestation
				this.policy.attestations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Policy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Policy/update/' + this.policy;

	return  this.http.post(uri_, this.policy );
}

	//********************************************************************
	// loadHelper - internal helper to load a Policy
	//********************************************************************	
	loadHelper( id ) {
		this.getPolicy(id)
			.subscribe((res : Policy) => {
				this.policy = res;
			});
	}
}