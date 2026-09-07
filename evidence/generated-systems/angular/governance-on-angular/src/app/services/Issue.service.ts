import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Issue} from '../models/Issue';
import {RiskService} from '../services/Risk.service';
import {AuditFindingService} from '../services/AuditFinding.service';
import {CorrectiveActionService} from '../services/CorrectiveAction.service';
import {ControlService} from '../services/Control.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class IssueService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	issue : Issue;

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
	// add a Issue
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addIssue(title, openedDate, closedDate, Risk, Finding, CorrectiveActions, Control, IssueType, Priority, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Issue/create';
		const obj = {
			      		title: title,
      		openedDate: openedDate,
      		closedDate: closedDate,
      		Risk: Risk != null && Risk.length > 0 ? Risk : null,
      		Finding: Finding != null && Finding.length > 0 ? Finding : null,
      		CorrectiveActions: CorrectiveActions != null && CorrectiveActions.length > 0 ? CorrectiveActions : null,
      		Control: Control != null && Control.length > 0 ? Control : null,
      		IssueType: IssueType,
      		Priority: Priority,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Issue
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateIssue(title, openedDate, closedDate, Risk, Finding, CorrectiveActions, Control, IssueType, Priority, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Issue/update/' + id;
		const obj = {
				      		title: title,
      		openedDate: openedDate,
      		closedDate: closedDate,
      		Risk: Risk != null && Risk.length > 0 ? Risk : null,
      		Finding: Finding != null && Finding.length > 0 ? Finding : null,
      		CorrectiveActions: CorrectiveActions != null && CorrectiveActions.length > 0 ? CorrectiveActions : null,
      		Control: Control != null && Control.length > 0 ? Control : null,
      		IssueType: IssueType,
      		Priority: Priority,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Issue
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteIssue(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Issue/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Issue
	// returns the results untouched as an Observable Issue
	// Issue model
	// delegates via URI
	//********************************************************************
	getIssue(id) : Observable<Issue> {
		const uri_ = this.apiUrl + '/Issue/load/' + id;

		return this.http.get<Issue>(uri_);
	}
	
	//********************************************************************
	// gets all Issue
	// returns the results untouched as JSON representation of an
	// Observable array of Issue models
	// delegates via URI
	//********************************************************************
	getIssues() : Observable<Issue[]> {
		const uri_ = this.apiUrl + '/Issue/';

		return this
			.http.get<Issue[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Risk on a Issue
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRisk( issueId, _riskId ): Observable<any> {

		// get the Issue from storage
		this.loadHelper( issueId );

	// get the Risk from storage
	var tmp 	= new RiskService(this.http).getRisk(_riskId);

	// assign the Risk
	this.issue.risk = tmp;

	// save the Issue
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Risk on a Issue
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRisk( issueId ): Observable<any> {

		// get the Issue from storage
		this.loadHelper( issueId );

	// assign Risk to null
	this.issue.risk = null;

	// save the Issue
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Finding on a Issue
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFinding( issueId, _findingId ): Observable<any> {

		// get the Issue from storage
		this.loadHelper( issueId );

	// get the AuditFinding from storage
	var tmp 	= new AuditFindingService(this.http).getAuditFinding(_findingId);

	// assign the Finding
	this.issue.finding = tmp;

	// save the Issue
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Finding on a Issue
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFinding( issueId ): Observable<any> {

		// get the Issue from storage
		this.loadHelper( issueId );

	// assign Finding to null
	this.issue.finding = null;

	// save the Issue
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Control on a Issue
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignControl( issueId, _controlId ): Observable<any> {

		// get the Issue from storage
		this.loadHelper( issueId );

	// get the Control from storage
	var tmp 	= new ControlService(this.http).getControl(_controlId);

	// assign the Control
	this.issue.control = tmp;

	// save the Issue
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Control on a Issue
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignControl( issueId ): Observable<any> {

		// get the Issue from storage
		this.loadHelper( issueId );

	// assign Control to null
	this.issue.control = null;

	// save the Issue
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more correctiveActionsIds as a CorrectiveActions
	// to a Issue
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCorrectiveActions( issueId, correctiveActionsIds ): Observable<any> {

		// get the Issue
		this.loadHelper( issueId );

	// split on a comma with no spaces
	var idList = correctiveActionsIds.split(',')

	// iterate over array of correctiveActions ids
	idList.forEach(function (id) {
		// read the CorrectiveAction
		var correctiveAction = new CorrectiveActionService(this.http).getCorrectiveAction(id);
		// add the CorrectiveAction if not already assigned
		if ( this.issue.correctiveActions.indexOf(correctiveAction) == -1 )
		this.issue.correctiveActions.push(correctiveAction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more correctiveActionsIds as a CorrectiveActions
	// from a Issue
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCorrectiveActions( issueId, correctiveActionsIds ): Observable<any> {

		// get the Issue
		this.loadHelper( issueId );


	// split on a comma with no spaces
	var idList 					= correctiveActionsIds.split(',');
	var correctiveActions 	= this.issue.correctiveActions;

	if ( correctiveActions != null && correctiveActionsIds != null ) {

		// iterate over array of correctiveActions ids
		correctiveActions.forEach(function (obj) {
			if ( correctiveActionsIds.indexOf(obj._id) > -1 ) {
				// remove the CorrectiveAction
				this.issue.correctiveActions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Issue
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Issue/update/' + this.issue;

	return  this.http.post(uri_, this.issue );
}

	//********************************************************************
	// loadHelper - internal helper to load a Issue
	//********************************************************************	
	loadHelper( id ) {
		this.getIssue(id)
			.subscribe((res : Issue) => {
				this.issue = res;
			});
	}
}