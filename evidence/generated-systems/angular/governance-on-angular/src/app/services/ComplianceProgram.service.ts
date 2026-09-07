import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ComplianceProgram} from '../models/ComplianceProgram';
import {OrganizationService} from '../services/Organization.service';
import {ComplianceRequirementService} from '../services/ComplianceRequirement.service';
import {ControlService} from '../services/Control.service';
import {AttestationService} from '../services/Attestation.service';
import {RegulationService} from '../services/Regulation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ComplianceProgramService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	complianceProgram : ComplianceProgram;

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
	// add a ComplianceProgram
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addComplianceProgram(name, framework, Organization, Requirements, Controls, Attestations, Regulations, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/ComplianceProgram/create';
		const obj = {
			      		name: name,
      		framework: framework,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Requirements: Requirements != null && Requirements.length > 0 ? Requirements : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
      		Attestations: Attestations != null && Attestations.length > 0 ? Attestations : null,
      		Regulations: Regulations != null && Regulations.length > 0 ? Regulations : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ComplianceProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateComplianceProgram(name, framework, Organization, Requirements, Controls, Attestations, Regulations, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ComplianceProgram/update/' + id;
		const obj = {
				      		name: name,
      		framework: framework,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Requirements: Requirements != null && Requirements.length > 0 ? Requirements : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
      		Attestations: Attestations != null && Attestations.length > 0 ? Attestations : null,
      		Regulations: Regulations != null && Regulations.length > 0 ? Regulations : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ComplianceProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteComplianceProgram(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ComplianceProgram/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ComplianceProgram
	// returns the results untouched as an Observable ComplianceProgram
	// ComplianceProgram model
	// delegates via URI
	//********************************************************************
	getComplianceProgram(id) : Observable<ComplianceProgram> {
		const uri_ = this.apiUrl + '/ComplianceProgram/load/' + id;

		return this.http.get<ComplianceProgram>(uri_);
	}
	
	//********************************************************************
	// gets all ComplianceProgram
	// returns the results untouched as JSON representation of an
	// Observable array of ComplianceProgram models
	// delegates via URI
	//********************************************************************
	getCompliancePrograms() : Observable<ComplianceProgram[]> {
		const uri_ = this.apiUrl + '/ComplianceProgram/';

		return this
			.http.get<ComplianceProgram[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a ComplianceProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( complianceProgramId, _organizationId ): Observable<any> {

		// get the ComplianceProgram from storage
		this.loadHelper( complianceProgramId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.complianceProgram.organization = tmp;

	// save the ComplianceProgram
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a ComplianceProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( complianceProgramId ): Observable<any> {

		// get the ComplianceProgram from storage
		this.loadHelper( complianceProgramId );

	// assign Organization to null
	this.complianceProgram.organization = null;

	// save the ComplianceProgram
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more requirementsIds as a Requirements
	// to a ComplianceProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRequirements( complianceProgramId, requirementsIds ): Observable<any> {

		// get the ComplianceProgram
		this.loadHelper( complianceProgramId );

	// split on a comma with no spaces
	var idList = requirementsIds.split(',')

	// iterate over array of requirements ids
	idList.forEach(function (id) {
		// read the ComplianceRequirement
		var complianceRequirement = new ComplianceRequirementService(this.http).getComplianceRequirement(id);
		// add the ComplianceRequirement if not already assigned
		if ( this.complianceProgram.requirements.indexOf(complianceRequirement) == -1 )
		this.complianceProgram.requirements.push(complianceRequirement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more requirementsIds as a Requirements
	// from a ComplianceProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRequirements( complianceProgramId, requirementsIds ): Observable<any> {

		// get the ComplianceProgram
		this.loadHelper( complianceProgramId );


	// split on a comma with no spaces
	var idList 					= requirementsIds.split(',');
	var requirements 	= this.complianceProgram.requirements;

	if ( requirements != null && requirementsIds != null ) {

		// iterate over array of requirements ids
		requirements.forEach(function (obj) {
			if ( requirementsIds.indexOf(obj._id) > -1 ) {
				// remove the ComplianceRequirement
				this.complianceProgram.requirements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more controlsIds as a Controls
	// to a ComplianceProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addControls( complianceProgramId, controlsIds ): Observable<any> {

		// get the ComplianceProgram
		this.loadHelper( complianceProgramId );

	// split on a comma with no spaces
	var idList = controlsIds.split(',')

	// iterate over array of controls ids
	idList.forEach(function (id) {
		// read the Control
		var control = new ControlService(this.http).getControl(id);
		// add the Control if not already assigned
		if ( this.complianceProgram.controls.indexOf(control) == -1 )
		this.complianceProgram.controls.push(control);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more controlsIds as a Controls
	// from a ComplianceProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeControls( complianceProgramId, controlsIds ): Observable<any> {

		// get the ComplianceProgram
		this.loadHelper( complianceProgramId );


	// split on a comma with no spaces
	var idList 					= controlsIds.split(',');
	var controls 	= this.complianceProgram.controls;

	if ( controls != null && controlsIds != null ) {

		// iterate over array of controls ids
		controls.forEach(function (obj) {
			if ( controlsIds.indexOf(obj._id) > -1 ) {
				// remove the Control
				this.complianceProgram.controls.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more attestationsIds as a Attestations
	// to a ComplianceProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAttestations( complianceProgramId, attestationsIds ): Observable<any> {

		// get the ComplianceProgram
		this.loadHelper( complianceProgramId );

	// split on a comma with no spaces
	var idList = attestationsIds.split(',')

	// iterate over array of attestations ids
	idList.forEach(function (id) {
		// read the Attestation
		var attestation = new AttestationService(this.http).getAttestation(id);
		// add the Attestation if not already assigned
		if ( this.complianceProgram.attestations.indexOf(attestation) == -1 )
		this.complianceProgram.attestations.push(attestation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more attestationsIds as a Attestations
	// from a ComplianceProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAttestations( complianceProgramId, attestationsIds ): Observable<any> {

		// get the ComplianceProgram
		this.loadHelper( complianceProgramId );


	// split on a comma with no spaces
	var idList 					= attestationsIds.split(',');
	var attestations 	= this.complianceProgram.attestations;

	if ( attestations != null && attestationsIds != null ) {

		// iterate over array of attestations ids
		attestations.forEach(function (obj) {
			if ( attestationsIds.indexOf(obj._id) > -1 ) {
				// remove the Attestation
				this.complianceProgram.attestations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more regulationsIds as a Regulations
	// to a ComplianceProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRegulations( complianceProgramId, regulationsIds ): Observable<any> {

		// get the ComplianceProgram
		this.loadHelper( complianceProgramId );

	// split on a comma with no spaces
	var idList = regulationsIds.split(',')

	// iterate over array of regulations ids
	idList.forEach(function (id) {
		// read the Regulation
		var regulation = new RegulationService(this.http).getRegulation(id);
		// add the Regulation if not already assigned
		if ( this.complianceProgram.regulations.indexOf(regulation) == -1 )
		this.complianceProgram.regulations.push(regulation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more regulationsIds as a Regulations
	// from a ComplianceProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRegulations( complianceProgramId, regulationsIds ): Observable<any> {

		// get the ComplianceProgram
		this.loadHelper( complianceProgramId );


	// split on a comma with no spaces
	var idList 					= regulationsIds.split(',');
	var regulations 	= this.complianceProgram.regulations;

	if ( regulations != null && regulationsIds != null ) {

		// iterate over array of regulations ids
		regulations.forEach(function (obj) {
			if ( regulationsIds.indexOf(obj._id) > -1 ) {
				// remove the Regulation
				this.complianceProgram.regulations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ComplianceProgram
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ComplianceProgram/update/' + this.complianceProgram;

	return  this.http.post(uri_, this.complianceProgram );
}

	//********************************************************************
	// loadHelper - internal helper to load a ComplianceProgram
	//********************************************************************	
	loadHelper( id ) {
		this.getComplianceProgram(id)
			.subscribe((res : ComplianceProgram) => {
				this.complianceProgram = res;
			});
	}
}