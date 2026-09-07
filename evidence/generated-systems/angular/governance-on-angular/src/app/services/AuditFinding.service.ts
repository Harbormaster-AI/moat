import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AuditFinding} from '../models/AuditFinding';
import {AuditEngagementService} from '../services/AuditEngagement.service';
import {AuditWorkpaperService} from '../services/AuditWorkpaper.service';
import {CorrectiveActionService} from '../services/CorrectiveAction.service';
import {RiskService} from '../services/Risk.service';
import {ControlService} from '../services/Control.service';
import {IssueService} from '../services/Issue.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AuditFindingService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	auditFinding : AuditFinding;

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
	// add a AuditFinding
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAuditFinding(title, description, dueDate, Engagement, Workpaper, CorrectiveActions, RelatedRisks, RelatedControls, Issues, Severity, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/AuditFinding/create';
		const obj = {
			      		title: title,
      		description: description,
      		dueDate: dueDate,
      		Engagement: Engagement != null && Engagement.length > 0 ? Engagement : null,
      		Workpaper: Workpaper != null && Workpaper.length > 0 ? Workpaper : null,
      		CorrectiveActions: CorrectiveActions != null && CorrectiveActions.length > 0 ? CorrectiveActions : null,
      		RelatedRisks: RelatedRisks != null && RelatedRisks.length > 0 ? RelatedRisks : null,
      		RelatedControls: RelatedControls != null && RelatedControls.length > 0 ? RelatedControls : null,
      		Issues: Issues != null && Issues.length > 0 ? Issues : null,
      		Severity: Severity,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AuditFinding
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAuditFinding(title, description, dueDate, Engagement, Workpaper, CorrectiveActions, RelatedRisks, RelatedControls, Issues, Severity, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AuditFinding/update/' + id;
		const obj = {
				      		title: title,
      		description: description,
      		dueDate: dueDate,
      		Engagement: Engagement != null && Engagement.length > 0 ? Engagement : null,
      		Workpaper: Workpaper != null && Workpaper.length > 0 ? Workpaper : null,
      		CorrectiveActions: CorrectiveActions != null && CorrectiveActions.length > 0 ? CorrectiveActions : null,
      		RelatedRisks: RelatedRisks != null && RelatedRisks.length > 0 ? RelatedRisks : null,
      		RelatedControls: RelatedControls != null && RelatedControls.length > 0 ? RelatedControls : null,
      		Issues: Issues != null && Issues.length > 0 ? Issues : null,
      		Severity: Severity,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AuditFinding
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAuditFinding(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AuditFinding/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AuditFinding
	// returns the results untouched as an Observable AuditFinding
	// AuditFinding model
	// delegates via URI
	//********************************************************************
	getAuditFinding(id) : Observable<AuditFinding> {
		const uri_ = this.apiUrl + '/AuditFinding/load/' + id;

		return this.http.get<AuditFinding>(uri_);
	}
	
	//********************************************************************
	// gets all AuditFinding
	// returns the results untouched as JSON representation of an
	// Observable array of AuditFinding models
	// delegates via URI
	//********************************************************************
	getAuditFindings() : Observable<AuditFinding[]> {
		const uri_ = this.apiUrl + '/AuditFinding/';

		return this
			.http.get<AuditFinding[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Engagement on a AuditFinding
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEngagement( auditFindingId, _engagementId ): Observable<any> {

		// get the AuditFinding from storage
		this.loadHelper( auditFindingId );

	// get the AuditEngagement from storage
	var tmp 	= new AuditEngagementService(this.http).getAuditEngagement(_engagementId);

	// assign the Engagement
	this.auditFinding.engagement = tmp;

	// save the AuditFinding
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Engagement on a AuditFinding
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEngagement( auditFindingId ): Observable<any> {

		// get the AuditFinding from storage
		this.loadHelper( auditFindingId );

	// assign Engagement to null
	this.auditFinding.engagement = null;

	// save the AuditFinding
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Workpaper on a AuditFinding
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkpaper( auditFindingId, _workpaperId ): Observable<any> {

		// get the AuditFinding from storage
		this.loadHelper( auditFindingId );

	// get the AuditWorkpaper from storage
	var tmp 	= new AuditWorkpaperService(this.http).getAuditWorkpaper(_workpaperId);

	// assign the Workpaper
	this.auditFinding.workpaper = tmp;

	// save the AuditFinding
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workpaper on a AuditFinding
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkpaper( auditFindingId ): Observable<any> {

		// get the AuditFinding from storage
		this.loadHelper( auditFindingId );

	// assign Workpaper to null
	this.auditFinding.workpaper = null;

	// save the AuditFinding
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more correctiveActionsIds as a CorrectiveActions
	// to a AuditFinding
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCorrectiveActions( auditFindingId, correctiveActionsIds ): Observable<any> {

		// get the AuditFinding
		this.loadHelper( auditFindingId );

	// split on a comma with no spaces
	var idList = correctiveActionsIds.split(',')

	// iterate over array of correctiveActions ids
	idList.forEach(function (id) {
		// read the CorrectiveAction
		var correctiveAction = new CorrectiveActionService(this.http).getCorrectiveAction(id);
		// add the CorrectiveAction if not already assigned
		if ( this.auditFinding.correctiveActions.indexOf(correctiveAction) == -1 )
		this.auditFinding.correctiveActions.push(correctiveAction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more correctiveActionsIds as a CorrectiveActions
	// from a AuditFinding
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCorrectiveActions( auditFindingId, correctiveActionsIds ): Observable<any> {

		// get the AuditFinding
		this.loadHelper( auditFindingId );


	// split on a comma with no spaces
	var idList 					= correctiveActionsIds.split(',');
	var correctiveActions 	= this.auditFinding.correctiveActions;

	if ( correctiveActions != null && correctiveActionsIds != null ) {

		// iterate over array of correctiveActions ids
		correctiveActions.forEach(function (obj) {
			if ( correctiveActionsIds.indexOf(obj._id) > -1 ) {
				// remove the CorrectiveAction
				this.auditFinding.correctiveActions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more relatedRisksIds as a RelatedRisks
	// to a AuditFinding
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRelatedRisks( auditFindingId, relatedRisksIds ): Observable<any> {

		// get the AuditFinding
		this.loadHelper( auditFindingId );

	// split on a comma with no spaces
	var idList = relatedRisksIds.split(',')

	// iterate over array of relatedRisks ids
	idList.forEach(function (id) {
		// read the Risk
		var risk = new RiskService(this.http).getRisk(id);
		// add the Risk if not already assigned
		if ( this.auditFinding.relatedRisks.indexOf(risk) == -1 )
		this.auditFinding.relatedRisks.push(risk);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more relatedRisksIds as a RelatedRisks
	// from a AuditFinding
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRelatedRisks( auditFindingId, relatedRisksIds ): Observable<any> {

		// get the AuditFinding
		this.loadHelper( auditFindingId );


	// split on a comma with no spaces
	var idList 					= relatedRisksIds.split(',');
	var relatedRisks 	= this.auditFinding.relatedRisks;

	if ( relatedRisks != null && relatedRisksIds != null ) {

		// iterate over array of relatedRisks ids
		relatedRisks.forEach(function (obj) {
			if ( relatedRisksIds.indexOf(obj._id) > -1 ) {
				// remove the Risk
				this.auditFinding.relatedRisks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more relatedControlsIds as a RelatedControls
	// to a AuditFinding
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRelatedControls( auditFindingId, relatedControlsIds ): Observable<any> {

		// get the AuditFinding
		this.loadHelper( auditFindingId );

	// split on a comma with no spaces
	var idList = relatedControlsIds.split(',')

	// iterate over array of relatedControls ids
	idList.forEach(function (id) {
		// read the Control
		var control = new ControlService(this.http).getControl(id);
		// add the Control if not already assigned
		if ( this.auditFinding.relatedControls.indexOf(control) == -1 )
		this.auditFinding.relatedControls.push(control);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more relatedControlsIds as a RelatedControls
	// from a AuditFinding
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRelatedControls( auditFindingId, relatedControlsIds ): Observable<any> {

		// get the AuditFinding
		this.loadHelper( auditFindingId );


	// split on a comma with no spaces
	var idList 					= relatedControlsIds.split(',');
	var relatedControls 	= this.auditFinding.relatedControls;

	if ( relatedControls != null && relatedControlsIds != null ) {

		// iterate over array of relatedControls ids
		relatedControls.forEach(function (obj) {
			if ( relatedControlsIds.indexOf(obj._id) > -1 ) {
				// remove the Control
				this.auditFinding.relatedControls.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more issuesIds as a Issues
	// to a AuditFinding
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addIssues( auditFindingId, issuesIds ): Observable<any> {

		// get the AuditFinding
		this.loadHelper( auditFindingId );

	// split on a comma with no spaces
	var idList = issuesIds.split(',')

	// iterate over array of issues ids
	idList.forEach(function (id) {
		// read the Issue
		var issue = new IssueService(this.http).getIssue(id);
		// add the Issue if not already assigned
		if ( this.auditFinding.issues.indexOf(issue) == -1 )
		this.auditFinding.issues.push(issue);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more issuesIds as a Issues
	// from a AuditFinding
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeIssues( auditFindingId, issuesIds ): Observable<any> {

		// get the AuditFinding
		this.loadHelper( auditFindingId );


	// split on a comma with no spaces
	var idList 					= issuesIds.split(',');
	var issues 	= this.auditFinding.issues;

	if ( issues != null && issuesIds != null ) {

		// iterate over array of issues ids
		issues.forEach(function (obj) {
			if ( issuesIds.indexOf(obj._id) > -1 ) {
				// remove the Issue
				this.auditFinding.issues.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AuditFinding
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AuditFinding/update/' + this.auditFinding;

	return  this.http.post(uri_, this.auditFinding );
}

	//********************************************************************
	// loadHelper - internal helper to load a AuditFinding
	//********************************************************************	
	loadHelper( id ) {
		this.getAuditFinding(id)
			.subscribe((res : AuditFinding) => {
				this.auditFinding = res;
			});
	}
}