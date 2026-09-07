import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ThirdPartyAssessment} from '../models/ThirdPartyAssessment';
import {ThirdPartyService} from '../services/ThirdParty.service';
import {IssueService} from '../services/Issue.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ThirdPartyAssessmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	thirdPartyAssessment : ThirdPartyAssessment;

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
	// add a ThirdPartyAssessment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addThirdPartyAssessment(assessmentDate, assessor, ThirdParty, Issues, AssessmentType, Result) : Observable<any> {
		const uri_ = this.apiUrl + '/ThirdPartyAssessment/create';
		const obj = {
			      		assessmentDate: assessmentDate,
      		assessor: assessor,
      		ThirdParty: ThirdParty != null && ThirdParty.length > 0 ? ThirdParty : null,
      		Issues: Issues != null && Issues.length > 0 ? Issues : null,
      		AssessmentType: AssessmentType,
			Result: Result
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ThirdPartyAssessment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateThirdPartyAssessment(assessmentDate, assessor, ThirdParty, Issues, AssessmentType, Result, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ThirdPartyAssessment/update/' + id;
		const obj = {
				      		assessmentDate: assessmentDate,
      		assessor: assessor,
      		ThirdParty: ThirdParty != null && ThirdParty.length > 0 ? ThirdParty : null,
      		Issues: Issues != null && Issues.length > 0 ? Issues : null,
      		AssessmentType: AssessmentType,
			Result: Result
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ThirdPartyAssessment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteThirdPartyAssessment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ThirdPartyAssessment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ThirdPartyAssessment
	// returns the results untouched as an Observable ThirdPartyAssessment
	// ThirdPartyAssessment model
	// delegates via URI
	//********************************************************************
	getThirdPartyAssessment(id) : Observable<ThirdPartyAssessment> {
		const uri_ = this.apiUrl + '/ThirdPartyAssessment/load/' + id;

		return this.http.get<ThirdPartyAssessment>(uri_);
	}
	
	//********************************************************************
	// gets all ThirdPartyAssessment
	// returns the results untouched as JSON representation of an
	// Observable array of ThirdPartyAssessment models
	// delegates via URI
	//********************************************************************
	getThirdPartyAssessments() : Observable<ThirdPartyAssessment[]> {
		const uri_ = this.apiUrl + '/ThirdPartyAssessment/';

		return this
			.http.get<ThirdPartyAssessment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ThirdParty on a ThirdPartyAssessment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignThirdParty( thirdPartyAssessmentId, _thirdPartyId ): Observable<any> {

		// get the ThirdPartyAssessment from storage
		this.loadHelper( thirdPartyAssessmentId );

	// get the ThirdParty from storage
	var tmp 	= new ThirdPartyService(this.http).getThirdParty(_thirdPartyId);

	// assign the ThirdParty
	this.thirdPartyAssessment.thirdParty = tmp;

	// save the ThirdPartyAssessment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ThirdParty on a ThirdPartyAssessment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignThirdParty( thirdPartyAssessmentId ): Observable<any> {

		// get the ThirdPartyAssessment from storage
		this.loadHelper( thirdPartyAssessmentId );

	// assign ThirdParty to null
	this.thirdPartyAssessment.thirdParty = null;

	// save the ThirdPartyAssessment
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more issuesIds as a Issues
	// to a ThirdPartyAssessment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addIssues( thirdPartyAssessmentId, issuesIds ): Observable<any> {

		// get the ThirdPartyAssessment
		this.loadHelper( thirdPartyAssessmentId );

	// split on a comma with no spaces
	var idList = issuesIds.split(',')

	// iterate over array of issues ids
	idList.forEach(function (id) {
		// read the Issue
		var issue = new IssueService(this.http).getIssue(id);
		// add the Issue if not already assigned
		if ( this.thirdPartyAssessment.issues.indexOf(issue) == -1 )
		this.thirdPartyAssessment.issues.push(issue);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more issuesIds as a Issues
	// from a ThirdPartyAssessment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeIssues( thirdPartyAssessmentId, issuesIds ): Observable<any> {

		// get the ThirdPartyAssessment
		this.loadHelper( thirdPartyAssessmentId );


	// split on a comma with no spaces
	var idList 					= issuesIds.split(',');
	var issues 	= this.thirdPartyAssessment.issues;

	if ( issues != null && issuesIds != null ) {

		// iterate over array of issues ids
		issues.forEach(function (obj) {
			if ( issuesIds.indexOf(obj._id) > -1 ) {
				// remove the Issue
				this.thirdPartyAssessment.issues.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ThirdPartyAssessment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ThirdPartyAssessment/update/' + this.thirdPartyAssessment;

	return  this.http.post(uri_, this.thirdPartyAssessment );
}

	//********************************************************************
	// loadHelper - internal helper to load a ThirdPartyAssessment
	//********************************************************************	
	loadHelper( id ) {
		this.getThirdPartyAssessment(id)
			.subscribe((res : ThirdPartyAssessment) => {
				this.thirdPartyAssessment = res;
			});
	}
}