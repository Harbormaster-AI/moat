import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {JobRequisition} from '../models/JobRequisition';
import {DepartmentService} from '../services/Department.service';
import {EmployeeService} from '../services/Employee.service';
import {JobProfileService} from '../services/JobProfile.service';
import {CandidateService} from '../services/Candidate.service';
import {InterviewService} from '../services/Interview.service';
import {OfferService} from '../services/Offer.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class JobRequisitionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	jobRequisition : JobRequisition;

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
	// add a JobRequisition
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addJobRequisition(requisitionNumber, title, openings, targetStartDate, Department, HiringManager, Recruiter, JobProfile, Candidates, Interviews, Offers, Status, Priority) : Observable<any> {
		const uri_ = this.apiUrl + '/JobRequisition/create';
		const obj = {
			      		requisitionNumber: requisitionNumber,
      		title: title,
      		openings: openings,
      		targetStartDate: targetStartDate,
      		Department: Department != null && Department.length > 0 ? Department : null,
      		HiringManager: HiringManager != null && HiringManager.length > 0 ? HiringManager : null,
      		Recruiter: Recruiter != null && Recruiter.length > 0 ? Recruiter : null,
      		JobProfile: JobProfile != null && JobProfile.length > 0 ? JobProfile : null,
      		Candidates: Candidates != null && Candidates.length > 0 ? Candidates : null,
      		Interviews: Interviews != null && Interviews.length > 0 ? Interviews : null,
      		Offers: Offers != null && Offers.length > 0 ? Offers : null,
      		Status: Status,
			Priority: Priority
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a JobRequisition
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateJobRequisition(requisitionNumber, title, openings, targetStartDate, Department, HiringManager, Recruiter, JobProfile, Candidates, Interviews, Offers, Status, Priority, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/JobRequisition/update/' + id;
		const obj = {
				      		requisitionNumber: requisitionNumber,
      		title: title,
      		openings: openings,
      		targetStartDate: targetStartDate,
      		Department: Department != null && Department.length > 0 ? Department : null,
      		HiringManager: HiringManager != null && HiringManager.length > 0 ? HiringManager : null,
      		Recruiter: Recruiter != null && Recruiter.length > 0 ? Recruiter : null,
      		JobProfile: JobProfile != null && JobProfile.length > 0 ? JobProfile : null,
      		Candidates: Candidates != null && Candidates.length > 0 ? Candidates : null,
      		Interviews: Interviews != null && Interviews.length > 0 ? Interviews : null,
      		Offers: Offers != null && Offers.length > 0 ? Offers : null,
      		Status: Status,
			Priority: Priority
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a JobRequisition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteJobRequisition(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/JobRequisition/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a JobRequisition
	// returns the results untouched as an Observable JobRequisition
	// JobRequisition model
	// delegates via URI
	//********************************************************************
	getJobRequisition(id) : Observable<JobRequisition> {
		const uri_ = this.apiUrl + '/JobRequisition/load/' + id;

		return this.http.get<JobRequisition>(uri_);
	}
	
	//********************************************************************
	// gets all JobRequisition
	// returns the results untouched as JSON representation of an
	// Observable array of JobRequisition models
	// delegates via URI
	//********************************************************************
	getJobRequisitions() : Observable<JobRequisition[]> {
		const uri_ = this.apiUrl + '/JobRequisition/';

		return this
			.http.get<JobRequisition[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Department on a JobRequisition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDepartment( jobRequisitionId, _departmentId ): Observable<any> {

		// get the JobRequisition from storage
		this.loadHelper( jobRequisitionId );

	// get the Department from storage
	var tmp 	= new DepartmentService(this.http).getDepartment(_departmentId);

	// assign the Department
	this.jobRequisition.department = tmp;

	// save the JobRequisition
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Department on a JobRequisition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDepartment( jobRequisitionId ): Observable<any> {

		// get the JobRequisition from storage
		this.loadHelper( jobRequisitionId );

	// assign Department to null
	this.jobRequisition.department = null;

	// save the JobRequisition
	return this.saveHelper();
}

		//********************************************************************
	// assigns a HiringManager on a JobRequisition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignHiringManager( jobRequisitionId, _hiringManagerId ): Observable<any> {

		// get the JobRequisition from storage
		this.loadHelper( jobRequisitionId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_hiringManagerId);

	// assign the HiringManager
	this.jobRequisition.hiringManager = tmp;

	// save the JobRequisition
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a HiringManager on a JobRequisition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignHiringManager( jobRequisitionId ): Observable<any> {

		// get the JobRequisition from storage
		this.loadHelper( jobRequisitionId );

	// assign HiringManager to null
	this.jobRequisition.hiringManager = null;

	// save the JobRequisition
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Recruiter on a JobRequisition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRecruiter( jobRequisitionId, _recruiterId ): Observable<any> {

		// get the JobRequisition from storage
		this.loadHelper( jobRequisitionId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_recruiterId);

	// assign the Recruiter
	this.jobRequisition.recruiter = tmp;

	// save the JobRequisition
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Recruiter on a JobRequisition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRecruiter( jobRequisitionId ): Observable<any> {

		// get the JobRequisition from storage
		this.loadHelper( jobRequisitionId );

	// assign Recruiter to null
	this.jobRequisition.recruiter = null;

	// save the JobRequisition
	return this.saveHelper();
}

		//********************************************************************
	// assigns a JobProfile on a JobRequisition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignJobProfile( jobRequisitionId, _jobProfileId ): Observable<any> {

		// get the JobRequisition from storage
		this.loadHelper( jobRequisitionId );

	// get the JobProfile from storage
	var tmp 	= new JobProfileService(this.http).getJobProfile(_jobProfileId);

	// assign the JobProfile
	this.jobRequisition.jobProfile = tmp;

	// save the JobRequisition
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a JobProfile on a JobRequisition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignJobProfile( jobRequisitionId ): Observable<any> {

		// get the JobRequisition from storage
		this.loadHelper( jobRequisitionId );

	// assign JobProfile to null
	this.jobRequisition.jobProfile = null;

	// save the JobRequisition
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more candidatesIds as a Candidates
	// to a JobRequisition
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCandidates( jobRequisitionId, candidatesIds ): Observable<any> {

		// get the JobRequisition
		this.loadHelper( jobRequisitionId );

	// split on a comma with no spaces
	var idList = candidatesIds.split(',')

	// iterate over array of candidates ids
	idList.forEach(function (id) {
		// read the Candidate
		var candidate = new CandidateService(this.http).getCandidate(id);
		// add the Candidate if not already assigned
		if ( this.jobRequisition.candidates.indexOf(candidate) == -1 )
		this.jobRequisition.candidates.push(candidate);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more candidatesIds as a Candidates
	// from a JobRequisition
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCandidates( jobRequisitionId, candidatesIds ): Observable<any> {

		// get the JobRequisition
		this.loadHelper( jobRequisitionId );


	// split on a comma with no spaces
	var idList 					= candidatesIds.split(',');
	var candidates 	= this.jobRequisition.candidates;

	if ( candidates != null && candidatesIds != null ) {

		// iterate over array of candidates ids
		candidates.forEach(function (obj) {
			if ( candidatesIds.indexOf(obj._id) > -1 ) {
				// remove the Candidate
				this.jobRequisition.candidates.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more interviewsIds as a Interviews
	// to a JobRequisition
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInterviews( jobRequisitionId, interviewsIds ): Observable<any> {

		// get the JobRequisition
		this.loadHelper( jobRequisitionId );

	// split on a comma with no spaces
	var idList = interviewsIds.split(',')

	// iterate over array of interviews ids
	idList.forEach(function (id) {
		// read the Interview
		var interview = new InterviewService(this.http).getInterview(id);
		// add the Interview if not already assigned
		if ( this.jobRequisition.interviews.indexOf(interview) == -1 )
		this.jobRequisition.interviews.push(interview);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more interviewsIds as a Interviews
	// from a JobRequisition
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInterviews( jobRequisitionId, interviewsIds ): Observable<any> {

		// get the JobRequisition
		this.loadHelper( jobRequisitionId );


	// split on a comma with no spaces
	var idList 					= interviewsIds.split(',');
	var interviews 	= this.jobRequisition.interviews;

	if ( interviews != null && interviewsIds != null ) {

		// iterate over array of interviews ids
		interviews.forEach(function (obj) {
			if ( interviewsIds.indexOf(obj._id) > -1 ) {
				// remove the Interview
				this.jobRequisition.interviews.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more offersIds as a Offers
	// to a JobRequisition
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOffers( jobRequisitionId, offersIds ): Observable<any> {

		// get the JobRequisition
		this.loadHelper( jobRequisitionId );

	// split on a comma with no spaces
	var idList = offersIds.split(',')

	// iterate over array of offers ids
	idList.forEach(function (id) {
		// read the Offer
		var offer = new OfferService(this.http).getOffer(id);
		// add the Offer if not already assigned
		if ( this.jobRequisition.offers.indexOf(offer) == -1 )
		this.jobRequisition.offers.push(offer);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more offersIds as a Offers
	// from a JobRequisition
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOffers( jobRequisitionId, offersIds ): Observable<any> {

		// get the JobRequisition
		this.loadHelper( jobRequisitionId );


	// split on a comma with no spaces
	var idList 					= offersIds.split(',');
	var offers 	= this.jobRequisition.offers;

	if ( offers != null && offersIds != null ) {

		// iterate over array of offers ids
		offers.forEach(function (obj) {
			if ( offersIds.indexOf(obj._id) > -1 ) {
				// remove the Offer
				this.jobRequisition.offers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a JobRequisition
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/JobRequisition/update/' + this.jobRequisition;

	return  this.http.post(uri_, this.jobRequisition );
}

	//********************************************************************
	// loadHelper - internal helper to load a JobRequisition
	//********************************************************************	
	loadHelper( id ) {
		this.getJobRequisition(id)
			.subscribe((res : JobRequisition) => {
				this.jobRequisition = res;
			});
	}
}