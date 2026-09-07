import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Risk} from '../models/Risk';
import {OrganizationService} from '../services/Organization.service';
import {ControlService} from '../services/Control.service';
import {RiskAssessmentService} from '../services/RiskAssessment.service';
import {IssueService} from '../services/Issue.service';
import {AuditFindingService} from '../services/AuditFinding.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RiskService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	risk : Risk;

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
	// add a Risk
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRisk(name, description, inherentRiskScore, residualRiskScore, Organization, Controls, Assessments, Issues, Findings, Category, Impact, Likelihood, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Risk/create';
		const obj = {
			      		name: name,
      		description: description,
      		inherentRiskScore: inherentRiskScore,
      		residualRiskScore: residualRiskScore,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
      		Assessments: Assessments != null && Assessments.length > 0 ? Assessments : null,
      		Issues: Issues != null && Issues.length > 0 ? Issues : null,
      		Findings: Findings != null && Findings.length > 0 ? Findings : null,
      		Category: Category,
      		Impact: Impact,
      		Likelihood: Likelihood,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Risk
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRisk(name, description, inherentRiskScore, residualRiskScore, Organization, Controls, Assessments, Issues, Findings, Category, Impact, Likelihood, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Risk/update/' + id;
		const obj = {
				      		name: name,
      		description: description,
      		inherentRiskScore: inherentRiskScore,
      		residualRiskScore: residualRiskScore,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
      		Assessments: Assessments != null && Assessments.length > 0 ? Assessments : null,
      		Issues: Issues != null && Issues.length > 0 ? Issues : null,
      		Findings: Findings != null && Findings.length > 0 ? Findings : null,
      		Category: Category,
      		Impact: Impact,
      		Likelihood: Likelihood,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Risk
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRisk(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Risk/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Risk
	// returns the results untouched as an Observable Risk
	// Risk model
	// delegates via URI
	//********************************************************************
	getRisk(id) : Observable<Risk> {
		const uri_ = this.apiUrl + '/Risk/load/' + id;

		return this.http.get<Risk>(uri_);
	}
	
	//********************************************************************
	// gets all Risk
	// returns the results untouched as JSON representation of an
	// Observable array of Risk models
	// delegates via URI
	//********************************************************************
	getRisks() : Observable<Risk[]> {
		const uri_ = this.apiUrl + '/Risk/';

		return this
			.http.get<Risk[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Risk
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( riskId, _organizationId ): Observable<any> {

		// get the Risk from storage
		this.loadHelper( riskId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.risk.organization = tmp;

	// save the Risk
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Risk
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( riskId ): Observable<any> {

		// get the Risk from storage
		this.loadHelper( riskId );

	// assign Organization to null
	this.risk.organization = null;

	// save the Risk
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more controlsIds as a Controls
	// to a Risk
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addControls( riskId, controlsIds ): Observable<any> {

		// get the Risk
		this.loadHelper( riskId );

	// split on a comma with no spaces
	var idList = controlsIds.split(',')

	// iterate over array of controls ids
	idList.forEach(function (id) {
		// read the Control
		var control = new ControlService(this.http).getControl(id);
		// add the Control if not already assigned
		if ( this.risk.controls.indexOf(control) == -1 )
		this.risk.controls.push(control);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more controlsIds as a Controls
	// from a Risk
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeControls( riskId, controlsIds ): Observable<any> {

		// get the Risk
		this.loadHelper( riskId );


	// split on a comma with no spaces
	var idList 					= controlsIds.split(',');
	var controls 	= this.risk.controls;

	if ( controls != null && controlsIds != null ) {

		// iterate over array of controls ids
		controls.forEach(function (obj) {
			if ( controlsIds.indexOf(obj._id) > -1 ) {
				// remove the Control
				this.risk.controls.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more assessmentsIds as a Assessments
	// to a Risk
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAssessments( riskId, assessmentsIds ): Observable<any> {

		// get the Risk
		this.loadHelper( riskId );

	// split on a comma with no spaces
	var idList = assessmentsIds.split(',')

	// iterate over array of assessments ids
	idList.forEach(function (id) {
		// read the RiskAssessment
		var riskAssessment = new RiskAssessmentService(this.http).getRiskAssessment(id);
		// add the RiskAssessment if not already assigned
		if ( this.risk.assessments.indexOf(riskAssessment) == -1 )
		this.risk.assessments.push(riskAssessment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more assessmentsIds as a Assessments
	// from a Risk
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAssessments( riskId, assessmentsIds ): Observable<any> {

		// get the Risk
		this.loadHelper( riskId );


	// split on a comma with no spaces
	var idList 					= assessmentsIds.split(',');
	var assessments 	= this.risk.assessments;

	if ( assessments != null && assessmentsIds != null ) {

		// iterate over array of assessments ids
		assessments.forEach(function (obj) {
			if ( assessmentsIds.indexOf(obj._id) > -1 ) {
				// remove the RiskAssessment
				this.risk.assessments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more issuesIds as a Issues
	// to a Risk
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addIssues( riskId, issuesIds ): Observable<any> {

		// get the Risk
		this.loadHelper( riskId );

	// split on a comma with no spaces
	var idList = issuesIds.split(',')

	// iterate over array of issues ids
	idList.forEach(function (id) {
		// read the Issue
		var issue = new IssueService(this.http).getIssue(id);
		// add the Issue if not already assigned
		if ( this.risk.issues.indexOf(issue) == -1 )
		this.risk.issues.push(issue);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more issuesIds as a Issues
	// from a Risk
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeIssues( riskId, issuesIds ): Observable<any> {

		// get the Risk
		this.loadHelper( riskId );


	// split on a comma with no spaces
	var idList 					= issuesIds.split(',');
	var issues 	= this.risk.issues;

	if ( issues != null && issuesIds != null ) {

		// iterate over array of issues ids
		issues.forEach(function (obj) {
			if ( issuesIds.indexOf(obj._id) > -1 ) {
				// remove the Issue
				this.risk.issues.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more findingsIds as a Findings
	// to a Risk
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFindings( riskId, findingsIds ): Observable<any> {

		// get the Risk
		this.loadHelper( riskId );

	// split on a comma with no spaces
	var idList = findingsIds.split(',')

	// iterate over array of findings ids
	idList.forEach(function (id) {
		// read the AuditFinding
		var auditFinding = new AuditFindingService(this.http).getAuditFinding(id);
		// add the AuditFinding if not already assigned
		if ( this.risk.findings.indexOf(auditFinding) == -1 )
		this.risk.findings.push(auditFinding);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more findingsIds as a Findings
	// from a Risk
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFindings( riskId, findingsIds ): Observable<any> {

		// get the Risk
		this.loadHelper( riskId );


	// split on a comma with no spaces
	var idList 					= findingsIds.split(',');
	var findings 	= this.risk.findings;

	if ( findings != null && findingsIds != null ) {

		// iterate over array of findings ids
		findings.forEach(function (obj) {
			if ( findingsIds.indexOf(obj._id) > -1 ) {
				// remove the AuditFinding
				this.risk.findings.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Risk
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Risk/update/' + this.risk;

	return  this.http.post(uri_, this.risk );
}

	//********************************************************************
	// loadHelper - internal helper to load a Risk
	//********************************************************************	
	loadHelper( id ) {
		this.getRisk(id)
			.subscribe((res : Risk) => {
				this.risk = res;
			});
	}
}