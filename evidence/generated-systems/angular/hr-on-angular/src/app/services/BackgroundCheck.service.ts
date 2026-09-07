import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BackgroundCheck} from '../models/BackgroundCheck';
import {CandidateService} from '../services/Candidate.service';
import {JobRequisitionService} from '../services/JobRequisition.service';
import {DocumentService} from '../services/Document.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BackgroundCheckService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	backgroundCheck : BackgroundCheck;

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
	// add a BackgroundCheck
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBackgroundCheck(checkNumber, provider, completedDate, Candidate, Requisition, Report, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/BackgroundCheck/create';
		const obj = {
			      		checkNumber: checkNumber,
      		provider: provider,
      		completedDate: completedDate,
      		Candidate: Candidate != null && Candidate.length > 0 ? Candidate : null,
      		Requisition: Requisition != null && Requisition.length > 0 ? Requisition : null,
      		Report: Report != null && Report.length > 0 ? Report : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BackgroundCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBackgroundCheck(checkNumber, provider, completedDate, Candidate, Requisition, Report, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BackgroundCheck/update/' + id;
		const obj = {
				      		checkNumber: checkNumber,
      		provider: provider,
      		completedDate: completedDate,
      		Candidate: Candidate != null && Candidate.length > 0 ? Candidate : null,
      		Requisition: Requisition != null && Requisition.length > 0 ? Requisition : null,
      		Report: Report != null && Report.length > 0 ? Report : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BackgroundCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBackgroundCheck(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BackgroundCheck/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BackgroundCheck
	// returns the results untouched as an Observable BackgroundCheck
	// BackgroundCheck model
	// delegates via URI
	//********************************************************************
	getBackgroundCheck(id) : Observable<BackgroundCheck> {
		const uri_ = this.apiUrl + '/BackgroundCheck/load/' + id;

		return this.http.get<BackgroundCheck>(uri_);
	}
	
	//********************************************************************
	// gets all BackgroundCheck
	// returns the results untouched as JSON representation of an
	// Observable array of BackgroundCheck models
	// delegates via URI
	//********************************************************************
	getBackgroundChecks() : Observable<BackgroundCheck[]> {
		const uri_ = this.apiUrl + '/BackgroundCheck/';

		return this
			.http.get<BackgroundCheck[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Candidate on a BackgroundCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCandidate( backgroundCheckId, _candidateId ): Observable<any> {

		// get the BackgroundCheck from storage
		this.loadHelper( backgroundCheckId );

	// get the Candidate from storage
	var tmp 	= new CandidateService(this.http).getCandidate(_candidateId);

	// assign the Candidate
	this.backgroundCheck.candidate = tmp;

	// save the BackgroundCheck
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Candidate on a BackgroundCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCandidate( backgroundCheckId ): Observable<any> {

		// get the BackgroundCheck from storage
		this.loadHelper( backgroundCheckId );

	// assign Candidate to null
	this.backgroundCheck.candidate = null;

	// save the BackgroundCheck
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Requisition on a BackgroundCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRequisition( backgroundCheckId, _requisitionId ): Observable<any> {

		// get the BackgroundCheck from storage
		this.loadHelper( backgroundCheckId );

	// get the JobRequisition from storage
	var tmp 	= new JobRequisitionService(this.http).getJobRequisition(_requisitionId);

	// assign the Requisition
	this.backgroundCheck.requisition = tmp;

	// save the BackgroundCheck
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Requisition on a BackgroundCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRequisition( backgroundCheckId ): Observable<any> {

		// get the BackgroundCheck from storage
		this.loadHelper( backgroundCheckId );

	// assign Requisition to null
	this.backgroundCheck.requisition = null;

	// save the BackgroundCheck
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Report on a BackgroundCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignReport( backgroundCheckId, _reportId ): Observable<any> {

		// get the BackgroundCheck from storage
		this.loadHelper( backgroundCheckId );

	// get the Document from storage
	var tmp 	= new DocumentService(this.http).getDocument(_reportId);

	// assign the Report
	this.backgroundCheck.report = tmp;

	// save the BackgroundCheck
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Report on a BackgroundCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignReport( backgroundCheckId ): Observable<any> {

		// get the BackgroundCheck from storage
		this.loadHelper( backgroundCheckId );

	// assign Report to null
	this.backgroundCheck.report = null;

	// save the BackgroundCheck
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a BackgroundCheck
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BackgroundCheck/update/' + this.backgroundCheck;

	return  this.http.post(uri_, this.backgroundCheck );
}

	//********************************************************************
	// loadHelper - internal helper to load a BackgroundCheck
	//********************************************************************	
	loadHelper( id ) {
		this.getBackgroundCheck(id)
			.subscribe((res : BackgroundCheck) => {
				this.backgroundCheck = res;
			});
	}
}