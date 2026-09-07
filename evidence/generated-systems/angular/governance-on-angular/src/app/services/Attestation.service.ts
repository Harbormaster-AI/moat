import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Attestation} from '../models/Attestation';
import {ControlService} from '../services/Control.service';
import {PolicyService} from '../services/Policy.service';
import {ComplianceProgramService} from '../services/ComplianceProgram.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AttestationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	attestation : Attestation;

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
	// add a Attestation
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAttestation(statement, attestor, dateSigned, Control, Policy, ComplianceProgram, Result) : Observable<any> {
		const uri_ = this.apiUrl + '/Attestation/create';
		const obj = {
			      		statement: statement,
      		attestor: attestor,
      		dateSigned: dateSigned,
      		Control: Control != null && Control.length > 0 ? Control : null,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		ComplianceProgram: ComplianceProgram != null && ComplianceProgram.length > 0 ? ComplianceProgram : null,
			Result: Result
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Attestation
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAttestation(statement, attestor, dateSigned, Control, Policy, ComplianceProgram, Result, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Attestation/update/' + id;
		const obj = {
				      		statement: statement,
      		attestor: attestor,
      		dateSigned: dateSigned,
      		Control: Control != null && Control.length > 0 ? Control : null,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		ComplianceProgram: ComplianceProgram != null && ComplianceProgram.length > 0 ? ComplianceProgram : null,
			Result: Result
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Attestation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAttestation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Attestation/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Attestation
	// returns the results untouched as an Observable Attestation
	// Attestation model
	// delegates via URI
	//********************************************************************
	getAttestation(id) : Observable<Attestation> {
		const uri_ = this.apiUrl + '/Attestation/load/' + id;

		return this.http.get<Attestation>(uri_);
	}
	
	//********************************************************************
	// gets all Attestation
	// returns the results untouched as JSON representation of an
	// Observable array of Attestation models
	// delegates via URI
	//********************************************************************
	getAttestations() : Observable<Attestation[]> {
		const uri_ = this.apiUrl + '/Attestation/';

		return this
			.http.get<Attestation[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Control on a Attestation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignControl( attestationId, _controlId ): Observable<any> {

		// get the Attestation from storage
		this.loadHelper( attestationId );

	// get the Control from storage
	var tmp 	= new ControlService(this.http).getControl(_controlId);

	// assign the Control
	this.attestation.control = tmp;

	// save the Attestation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Control on a Attestation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignControl( attestationId ): Observable<any> {

		// get the Attestation from storage
		this.loadHelper( attestationId );

	// assign Control to null
	this.attestation.control = null;

	// save the Attestation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Policy on a Attestation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( attestationId, _policyId ): Observable<any> {

		// get the Attestation from storage
		this.loadHelper( attestationId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.attestation.policy = tmp;

	// save the Attestation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Attestation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( attestationId ): Observable<any> {

		// get the Attestation from storage
		this.loadHelper( attestationId );

	// assign Policy to null
	this.attestation.policy = null;

	// save the Attestation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ComplianceProgram on a Attestation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignComplianceProgram( attestationId, _complianceProgramId ): Observable<any> {

		// get the Attestation from storage
		this.loadHelper( attestationId );

	// get the ComplianceProgram from storage
	var tmp 	= new ComplianceProgramService(this.http).getComplianceProgram(_complianceProgramId);

	// assign the ComplianceProgram
	this.attestation.complianceProgram = tmp;

	// save the Attestation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ComplianceProgram on a Attestation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignComplianceProgram( attestationId ): Observable<any> {

		// get the Attestation from storage
		this.loadHelper( attestationId );

	// assign ComplianceProgram to null
	this.attestation.complianceProgram = null;

	// save the Attestation
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Attestation
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Attestation/update/' + this.attestation;

	return  this.http.post(uri_, this.attestation );
}

	//********************************************************************
	// loadHelper - internal helper to load a Attestation
	//********************************************************************	
	loadHelper( id ) {
		this.getAttestation(id)
			.subscribe((res : Attestation) => {
				this.attestation = res;
			});
	}
}