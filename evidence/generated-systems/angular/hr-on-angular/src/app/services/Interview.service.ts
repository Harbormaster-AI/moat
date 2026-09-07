import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Interview} from '../models/Interview';
import {JobRequisitionService} from '../services/JobRequisition.service';
import {CandidateService} from '../services/Candidate.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InterviewService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	interview : Interview;

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
	// add a Interview
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInterview(interviewDate, feedback, Requisition, Candidate, Interviewers, Stage, Result) : Observable<any> {
		const uri_ = this.apiUrl + '/Interview/create';
		const obj = {
			      		interviewDate: interviewDate,
      		feedback: feedback,
      		Requisition: Requisition != null && Requisition.length > 0 ? Requisition : null,
      		Candidate: Candidate != null && Candidate.length > 0 ? Candidate : null,
      		Interviewers: Interviewers != null && Interviewers.length > 0 ? Interviewers : null,
      		Stage: Stage,
			Result: Result
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Interview
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInterview(interviewDate, feedback, Requisition, Candidate, Interviewers, Stage, Result, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Interview/update/' + id;
		const obj = {
				      		interviewDate: interviewDate,
      		feedback: feedback,
      		Requisition: Requisition != null && Requisition.length > 0 ? Requisition : null,
      		Candidate: Candidate != null && Candidate.length > 0 ? Candidate : null,
      		Interviewers: Interviewers != null && Interviewers.length > 0 ? Interviewers : null,
      		Stage: Stage,
			Result: Result
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Interview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInterview(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Interview/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Interview
	// returns the results untouched as an Observable Interview
	// Interview model
	// delegates via URI
	//********************************************************************
	getInterview(id) : Observable<Interview> {
		const uri_ = this.apiUrl + '/Interview/load/' + id;

		return this.http.get<Interview>(uri_);
	}
	
	//********************************************************************
	// gets all Interview
	// returns the results untouched as JSON representation of an
	// Observable array of Interview models
	// delegates via URI
	//********************************************************************
	getInterviews() : Observable<Interview[]> {
		const uri_ = this.apiUrl + '/Interview/';

		return this
			.http.get<Interview[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Requisition on a Interview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRequisition( interviewId, _requisitionId ): Observable<any> {

		// get the Interview from storage
		this.loadHelper( interviewId );

	// get the JobRequisition from storage
	var tmp 	= new JobRequisitionService(this.http).getJobRequisition(_requisitionId);

	// assign the Requisition
	this.interview.requisition = tmp;

	// save the Interview
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Requisition on a Interview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRequisition( interviewId ): Observable<any> {

		// get the Interview from storage
		this.loadHelper( interviewId );

	// assign Requisition to null
	this.interview.requisition = null;

	// save the Interview
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Candidate on a Interview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCandidate( interviewId, _candidateId ): Observable<any> {

		// get the Interview from storage
		this.loadHelper( interviewId );

	// get the Candidate from storage
	var tmp 	= new CandidateService(this.http).getCandidate(_candidateId);

	// assign the Candidate
	this.interview.candidate = tmp;

	// save the Interview
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Candidate on a Interview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCandidate( interviewId ): Observable<any> {

		// get the Interview from storage
		this.loadHelper( interviewId );

	// assign Candidate to null
	this.interview.candidate = null;

	// save the Interview
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more interviewersIds as a Interviewers
	// to a Interview
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInterviewers( interviewId, interviewersIds ): Observable<any> {

		// get the Interview
		this.loadHelper( interviewId );

	// split on a comma with no spaces
	var idList = interviewersIds.split(',')

	// iterate over array of interviewers ids
	idList.forEach(function (id) {
		// read the Employee
		var employee = new EmployeeService(this.http).getEmployee(id);
		// add the Employee if not already assigned
		if ( this.interview.interviewers.indexOf(employee) == -1 )
		this.interview.interviewers.push(employee);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more interviewersIds as a Interviewers
	// from a Interview
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInterviewers( interviewId, interviewersIds ): Observable<any> {

		// get the Interview
		this.loadHelper( interviewId );


	// split on a comma with no spaces
	var idList 					= interviewersIds.split(',');
	var interviewers 	= this.interview.interviewers;

	if ( interviewers != null && interviewersIds != null ) {

		// iterate over array of interviewers ids
		interviewers.forEach(function (obj) {
			if ( interviewersIds.indexOf(obj._id) > -1 ) {
				// remove the Employee
				this.interview.interviewers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Interview
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Interview/update/' + this.interview;

	return  this.http.post(uri_, this.interview );
}

	//********************************************************************
	// loadHelper - internal helper to load a Interview
	//********************************************************************	
	loadHelper( id ) {
		this.getInterview(id)
			.subscribe((res : Interview) => {
				this.interview = res;
			});
	}
}