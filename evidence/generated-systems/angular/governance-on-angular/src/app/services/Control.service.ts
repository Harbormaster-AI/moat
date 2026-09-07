import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Control} from '../models/Control';
import {PolicyService} from '../services/Policy.service';
import {ControlTest_Service} from '../services/ControlTest_.service';
import {EvidenceService} from '../services/Evidence.service';
import {RiskService} from '../services/Risk.service';
import {ObligationService} from '../services/Obligation.service';
import {ProcedureService} from '../services/Procedure.service';
import {IssueService} from '../services/Issue.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ControlService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	control : Control;

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
	// add a Control
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addControl(name, objective, ownerDepartment, Policy, ControlTests, Evidence, Risks, Obligations, Procedures, Issues, ControlType, Frequency, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Control/create';
		const obj = {
			      		name: name,
      		objective: objective,
      		ownerDepartment: ownerDepartment,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		ControlTests: ControlTests != null && ControlTests.length > 0 ? ControlTests : null,
      		Evidence: Evidence != null && Evidence.length > 0 ? Evidence : null,
      		Risks: Risks != null && Risks.length > 0 ? Risks : null,
      		Obligations: Obligations != null && Obligations.length > 0 ? Obligations : null,
      		Procedures: Procedures != null && Procedures.length > 0 ? Procedures : null,
      		Issues: Issues != null && Issues.length > 0 ? Issues : null,
      		ControlType: ControlType,
      		Frequency: Frequency,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Control
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateControl(name, objective, ownerDepartment, Policy, ControlTests, Evidence, Risks, Obligations, Procedures, Issues, ControlType, Frequency, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Control/update/' + id;
		const obj = {
				      		name: name,
      		objective: objective,
      		ownerDepartment: ownerDepartment,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		ControlTests: ControlTests != null && ControlTests.length > 0 ? ControlTests : null,
      		Evidence: Evidence != null && Evidence.length > 0 ? Evidence : null,
      		Risks: Risks != null && Risks.length > 0 ? Risks : null,
      		Obligations: Obligations != null && Obligations.length > 0 ? Obligations : null,
      		Procedures: Procedures != null && Procedures.length > 0 ? Procedures : null,
      		Issues: Issues != null && Issues.length > 0 ? Issues : null,
      		ControlType: ControlType,
      		Frequency: Frequency,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Control
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteControl(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Control/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Control
	// returns the results untouched as an Observable Control
	// Control model
	// delegates via URI
	//********************************************************************
	getControl(id) : Observable<Control> {
		const uri_ = this.apiUrl + '/Control/load/' + id;

		return this.http.get<Control>(uri_);
	}
	
	//********************************************************************
	// gets all Control
	// returns the results untouched as JSON representation of an
	// Observable array of Control models
	// delegates via URI
	//********************************************************************
	getControls() : Observable<Control[]> {
		const uri_ = this.apiUrl + '/Control/';

		return this
			.http.get<Control[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Policy on a Control
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( controlId, _policyId ): Observable<any> {

		// get the Control from storage
		this.loadHelper( controlId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.control.policy = tmp;

	// save the Control
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Control
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( controlId ): Observable<any> {

		// get the Control from storage
		this.loadHelper( controlId );

	// assign Policy to null
	this.control.policy = null;

	// save the Control
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more controlTestsIds as a ControlTests
	// to a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addControlTests( controlId, controlTestsIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );

	// split on a comma with no spaces
	var idList = controlTestsIds.split(',')

	// iterate over array of controlTests ids
	idList.forEach(function (id) {
		// read the ControlTest_
		var controlTest_ = new ControlTest_Service(this.http).getControlTest_(id);
		// add the ControlTest_ if not already assigned
		if ( this.control.controlTests.indexOf(controlTest_) == -1 )
		this.control.controlTests.push(controlTest_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more controlTestsIds as a ControlTests
	// from a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeControlTests( controlId, controlTestsIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );


	// split on a comma with no spaces
	var idList 					= controlTestsIds.split(',');
	var controlTests 	= this.control.controlTests;

	if ( controlTests != null && controlTestsIds != null ) {

		// iterate over array of controlTests ids
		controlTests.forEach(function (obj) {
			if ( controlTestsIds.indexOf(obj._id) > -1 ) {
				// remove the ControlTest_
				this.control.controlTests.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more evidenceIds as a Evidence
	// to a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEvidence( controlId, evidenceIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );

	// split on a comma with no spaces
	var idList = evidenceIds.split(',')

	// iterate over array of evidence ids
	idList.forEach(function (id) {
		// read the Evidence
		var evidence = new EvidenceService(this.http).getEvidence(id);
		// add the Evidence if not already assigned
		if ( this.control.evidence.indexOf(evidence) == -1 )
		this.control.evidence.push(evidence);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more evidenceIds as a Evidence
	// from a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEvidence( controlId, evidenceIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );


	// split on a comma with no spaces
	var idList 					= evidenceIds.split(',');
	var evidence 	= this.control.evidence;

	if ( evidence != null && evidenceIds != null ) {

		// iterate over array of evidence ids
		evidence.forEach(function (obj) {
			if ( evidenceIds.indexOf(obj._id) > -1 ) {
				// remove the Evidence
				this.control.evidence.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more risksIds as a Risks
	// to a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRisks( controlId, risksIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );

	// split on a comma with no spaces
	var idList = risksIds.split(',')

	// iterate over array of risks ids
	idList.forEach(function (id) {
		// read the Risk
		var risk = new RiskService(this.http).getRisk(id);
		// add the Risk if not already assigned
		if ( this.control.risks.indexOf(risk) == -1 )
		this.control.risks.push(risk);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more risksIds as a Risks
	// from a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRisks( controlId, risksIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );


	// split on a comma with no spaces
	var idList 					= risksIds.split(',');
	var risks 	= this.control.risks;

	if ( risks != null && risksIds != null ) {

		// iterate over array of risks ids
		risks.forEach(function (obj) {
			if ( risksIds.indexOf(obj._id) > -1 ) {
				// remove the Risk
				this.control.risks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more obligationsIds as a Obligations
	// to a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addObligations( controlId, obligationsIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );

	// split on a comma with no spaces
	var idList = obligationsIds.split(',')

	// iterate over array of obligations ids
	idList.forEach(function (id) {
		// read the Obligation
		var obligation = new ObligationService(this.http).getObligation(id);
		// add the Obligation if not already assigned
		if ( this.control.obligations.indexOf(obligation) == -1 )
		this.control.obligations.push(obligation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more obligationsIds as a Obligations
	// from a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeObligations( controlId, obligationsIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );


	// split on a comma with no spaces
	var idList 					= obligationsIds.split(',');
	var obligations 	= this.control.obligations;

	if ( obligations != null && obligationsIds != null ) {

		// iterate over array of obligations ids
		obligations.forEach(function (obj) {
			if ( obligationsIds.indexOf(obj._id) > -1 ) {
				// remove the Obligation
				this.control.obligations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more proceduresIds as a Procedures
	// to a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcedures( controlId, proceduresIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );

	// split on a comma with no spaces
	var idList = proceduresIds.split(',')

	// iterate over array of procedures ids
	idList.forEach(function (id) {
		// read the Procedure
		var procedure = new ProcedureService(this.http).getProcedure(id);
		// add the Procedure if not already assigned
		if ( this.control.procedures.indexOf(procedure) == -1 )
		this.control.procedures.push(procedure);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more proceduresIds as a Procedures
	// from a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcedures( controlId, proceduresIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );


	// split on a comma with no spaces
	var idList 					= proceduresIds.split(',');
	var procedures 	= this.control.procedures;

	if ( procedures != null && proceduresIds != null ) {

		// iterate over array of procedures ids
		procedures.forEach(function (obj) {
			if ( proceduresIds.indexOf(obj._id) > -1 ) {
				// remove the Procedure
				this.control.procedures.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more issuesIds as a Issues
	// to a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addIssues( controlId, issuesIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );

	// split on a comma with no spaces
	var idList = issuesIds.split(',')

	// iterate over array of issues ids
	idList.forEach(function (id) {
		// read the Issue
		var issue = new IssueService(this.http).getIssue(id);
		// add the Issue if not already assigned
		if ( this.control.issues.indexOf(issue) == -1 )
		this.control.issues.push(issue);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more issuesIds as a Issues
	// from a Control
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeIssues( controlId, issuesIds ): Observable<any> {

		// get the Control
		this.loadHelper( controlId );


	// split on a comma with no spaces
	var idList 					= issuesIds.split(',');
	var issues 	= this.control.issues;

	if ( issues != null && issuesIds != null ) {

		// iterate over array of issues ids
		issues.forEach(function (obj) {
			if ( issuesIds.indexOf(obj._id) > -1 ) {
				// remove the Issue
				this.control.issues.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Control
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Control/update/' + this.control;

	return  this.http.post(uri_, this.control );
}

	//********************************************************************
	// loadHelper - internal helper to load a Control
	//********************************************************************	
	loadHelper( id ) {
		this.getControl(id)
			.subscribe((res : Control) => {
				this.control = res;
			});
	}
}