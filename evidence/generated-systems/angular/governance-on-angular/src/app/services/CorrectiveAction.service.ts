import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CorrectiveAction} from '../models/CorrectiveAction';
import {AuditFindingService} from '../services/AuditFinding.service';
import {IssueService} from '../services/Issue.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CorrectiveActionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	correctiveAction : CorrectiveAction;

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
	// add a CorrectiveAction
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCorrectiveAction(actionTitle, owner, targetDate, Finding, Issue, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/CorrectiveAction/create';
		const obj = {
			      		actionTitle: actionTitle,
      		owner: owner,
      		targetDate: targetDate,
      		Finding: Finding != null && Finding.length > 0 ? Finding : null,
      		Issue: Issue != null && Issue.length > 0 ? Issue : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCorrectiveAction(actionTitle, owner, targetDate, Finding, Issue, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CorrectiveAction/update/' + id;
		const obj = {
				      		actionTitle: actionTitle,
      		owner: owner,
      		targetDate: targetDate,
      		Finding: Finding != null && Finding.length > 0 ? Finding : null,
      		Issue: Issue != null && Issue.length > 0 ? Issue : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCorrectiveAction(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CorrectiveAction/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CorrectiveAction
	// returns the results untouched as an Observable CorrectiveAction
	// CorrectiveAction model
	// delegates via URI
	//********************************************************************
	getCorrectiveAction(id) : Observable<CorrectiveAction> {
		const uri_ = this.apiUrl + '/CorrectiveAction/load/' + id;

		return this.http.get<CorrectiveAction>(uri_);
	}
	
	//********************************************************************
	// gets all CorrectiveAction
	// returns the results untouched as JSON representation of an
	// Observable array of CorrectiveAction models
	// delegates via URI
	//********************************************************************
	getCorrectiveActions() : Observable<CorrectiveAction[]> {
		const uri_ = this.apiUrl + '/CorrectiveAction/';

		return this
			.http.get<CorrectiveAction[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Finding on a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFinding( correctiveActionId, _findingId ): Observable<any> {

		// get the CorrectiveAction from storage
		this.loadHelper( correctiveActionId );

	// get the AuditFinding from storage
	var tmp 	= new AuditFindingService(this.http).getAuditFinding(_findingId);

	// assign the Finding
	this.correctiveAction.finding = tmp;

	// save the CorrectiveAction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Finding on a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFinding( correctiveActionId ): Observable<any> {

		// get the CorrectiveAction from storage
		this.loadHelper( correctiveActionId );

	// assign Finding to null
	this.correctiveAction.finding = null;

	// save the CorrectiveAction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Issue on a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignIssue( correctiveActionId, _issueId ): Observable<any> {

		// get the CorrectiveAction from storage
		this.loadHelper( correctiveActionId );

	// get the Issue from storage
	var tmp 	= new IssueService(this.http).getIssue(_issueId);

	// assign the Issue
	this.correctiveAction.issue = tmp;

	// save the CorrectiveAction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Issue on a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignIssue( correctiveActionId ): Observable<any> {

		// get the CorrectiveAction from storage
		this.loadHelper( correctiveActionId );

	// assign Issue to null
	this.correctiveAction.issue = null;

	// save the CorrectiveAction
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CorrectiveAction
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CorrectiveAction/update/' + this.correctiveAction;

	return  this.http.post(uri_, this.correctiveAction );
}

	//********************************************************************
	// loadHelper - internal helper to load a CorrectiveAction
	//********************************************************************	
	loadHelper( id ) {
		this.getCorrectiveAction(id)
			.subscribe((res : CorrectiveAction) => {
				this.correctiveAction = res;
			});
	}
}