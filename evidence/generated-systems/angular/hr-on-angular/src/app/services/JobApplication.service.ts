import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {JobApplication} from '../models/JobApplication';
import {CandidateService} from '../services/Candidate.service';
import {JobRequisitionService} from '../services/JobRequisition.service';
import {ScreeningService} from '../services/Screening.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class JobApplicationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	jobApplication : JobApplication;

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
	// add a JobApplication
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addJobApplication(applicationNumber, appliedDate, resumeUrl, Candidate, Requisition, Screenings, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/JobApplication/create';
		const obj = {
			      		applicationNumber: applicationNumber,
      		appliedDate: appliedDate,
      		resumeUrl: resumeUrl,
      		Candidate: Candidate != null && Candidate.length > 0 ? Candidate : null,
      		Requisition: Requisition != null && Requisition.length > 0 ? Requisition : null,
      		Screenings: Screenings != null && Screenings.length > 0 ? Screenings : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a JobApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateJobApplication(applicationNumber, appliedDate, resumeUrl, Candidate, Requisition, Screenings, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/JobApplication/update/' + id;
		const obj = {
				      		applicationNumber: applicationNumber,
      		appliedDate: appliedDate,
      		resumeUrl: resumeUrl,
      		Candidate: Candidate != null && Candidate.length > 0 ? Candidate : null,
      		Requisition: Requisition != null && Requisition.length > 0 ? Requisition : null,
      		Screenings: Screenings != null && Screenings.length > 0 ? Screenings : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a JobApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteJobApplication(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/JobApplication/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a JobApplication
	// returns the results untouched as an Observable JobApplication
	// JobApplication model
	// delegates via URI
	//********************************************************************
	getJobApplication(id) : Observable<JobApplication> {
		const uri_ = this.apiUrl + '/JobApplication/load/' + id;

		return this.http.get<JobApplication>(uri_);
	}
	
	//********************************************************************
	// gets all JobApplication
	// returns the results untouched as JSON representation of an
	// Observable array of JobApplication models
	// delegates via URI
	//********************************************************************
	getJobApplications() : Observable<JobApplication[]> {
		const uri_ = this.apiUrl + '/JobApplication/';

		return this
			.http.get<JobApplication[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Candidate on a JobApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCandidate( jobApplicationId, _candidateId ): Observable<any> {

		// get the JobApplication from storage
		this.loadHelper( jobApplicationId );

	// get the Candidate from storage
	var tmp 	= new CandidateService(this.http).getCandidate(_candidateId);

	// assign the Candidate
	this.jobApplication.candidate = tmp;

	// save the JobApplication
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Candidate on a JobApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCandidate( jobApplicationId ): Observable<any> {

		// get the JobApplication from storage
		this.loadHelper( jobApplicationId );

	// assign Candidate to null
	this.jobApplication.candidate = null;

	// save the JobApplication
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Requisition on a JobApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRequisition( jobApplicationId, _requisitionId ): Observable<any> {

		// get the JobApplication from storage
		this.loadHelper( jobApplicationId );

	// get the JobRequisition from storage
	var tmp 	= new JobRequisitionService(this.http).getJobRequisition(_requisitionId);

	// assign the Requisition
	this.jobApplication.requisition = tmp;

	// save the JobApplication
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Requisition on a JobApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRequisition( jobApplicationId ): Observable<any> {

		// get the JobApplication from storage
		this.loadHelper( jobApplicationId );

	// assign Requisition to null
	this.jobApplication.requisition = null;

	// save the JobApplication
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more screeningsIds as a Screenings
	// to a JobApplication
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addScreenings( jobApplicationId, screeningsIds ): Observable<any> {

		// get the JobApplication
		this.loadHelper( jobApplicationId );

	// split on a comma with no spaces
	var idList = screeningsIds.split(',')

	// iterate over array of screenings ids
	idList.forEach(function (id) {
		// read the Screening
		var screening = new ScreeningService(this.http).getScreening(id);
		// add the Screening if not already assigned
		if ( this.jobApplication.screenings.indexOf(screening) == -1 )
		this.jobApplication.screenings.push(screening);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more screeningsIds as a Screenings
	// from a JobApplication
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeScreenings( jobApplicationId, screeningsIds ): Observable<any> {

		// get the JobApplication
		this.loadHelper( jobApplicationId );


	// split on a comma with no spaces
	var idList 					= screeningsIds.split(',');
	var screenings 	= this.jobApplication.screenings;

	if ( screenings != null && screeningsIds != null ) {

		// iterate over array of screenings ids
		screenings.forEach(function (obj) {
			if ( screeningsIds.indexOf(obj._id) > -1 ) {
				// remove the Screening
				this.jobApplication.screenings.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a JobApplication
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/JobApplication/update/' + this.jobApplication;

	return  this.http.post(uri_, this.jobApplication );
}

	//********************************************************************
	// loadHelper - internal helper to load a JobApplication
	//********************************************************************	
	loadHelper( id ) {
		this.getJobApplication(id)
			.subscribe((res : JobApplication) => {
				this.jobApplication = res;
			});
	}
}