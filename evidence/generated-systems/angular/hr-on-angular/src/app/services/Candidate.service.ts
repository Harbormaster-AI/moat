import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Candidate} from '../models/Candidate';
import {JobApplicationService} from '../services/JobApplication.service';
import {InterviewService} from '../services/Interview.service';
import {OfferService} from '../services/Offer.service';
import {DocumentService} from '../services/Document.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CandidateService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	candidate : Candidate;

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
	// add a Candidate
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCandidate(name, email, phone, Applications, Interviews, Offers, Documents, Source) : Observable<any> {
		const uri_ = this.apiUrl + '/Candidate/create';
		const obj = {
			      		name: name,
      		email: email,
      		phone: phone,
      		Applications: Applications != null && Applications.length > 0 ? Applications : null,
      		Interviews: Interviews != null && Interviews.length > 0 ? Interviews : null,
      		Offers: Offers != null && Offers.length > 0 ? Offers : null,
      		Documents: Documents != null && Documents.length > 0 ? Documents : null,
			Source: Source
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Candidate
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCandidate(name, email, phone, Applications, Interviews, Offers, Documents, Source, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Candidate/update/' + id;
		const obj = {
				      		name: name,
      		email: email,
      		phone: phone,
      		Applications: Applications != null && Applications.length > 0 ? Applications : null,
      		Interviews: Interviews != null && Interviews.length > 0 ? Interviews : null,
      		Offers: Offers != null && Offers.length > 0 ? Offers : null,
      		Documents: Documents != null && Documents.length > 0 ? Documents : null,
			Source: Source
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Candidate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCandidate(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Candidate/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Candidate
	// returns the results untouched as an Observable Candidate
	// Candidate model
	// delegates via URI
	//********************************************************************
	getCandidate(id) : Observable<Candidate> {
		const uri_ = this.apiUrl + '/Candidate/load/' + id;

		return this.http.get<Candidate>(uri_);
	}
	
	//********************************************************************
	// gets all Candidate
	// returns the results untouched as JSON representation of an
	// Observable array of Candidate models
	// delegates via URI
	//********************************************************************
	getCandidates() : Observable<Candidate[]> {
		const uri_ = this.apiUrl + '/Candidate/';

		return this
			.http.get<Candidate[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more applicationsIds as a Applications
	// to a Candidate
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addApplications( candidateId, applicationsIds ): Observable<any> {

		// get the Candidate
		this.loadHelper( candidateId );

	// split on a comma with no spaces
	var idList = applicationsIds.split(',')

	// iterate over array of applications ids
	idList.forEach(function (id) {
		// read the JobApplication
		var jobApplication = new JobApplicationService(this.http).getJobApplication(id);
		// add the JobApplication if not already assigned
		if ( this.candidate.applications.indexOf(jobApplication) == -1 )
		this.candidate.applications.push(jobApplication);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more applicationsIds as a Applications
	// from a Candidate
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeApplications( candidateId, applicationsIds ): Observable<any> {

		// get the Candidate
		this.loadHelper( candidateId );


	// split on a comma with no spaces
	var idList 					= applicationsIds.split(',');
	var applications 	= this.candidate.applications;

	if ( applications != null && applicationsIds != null ) {

		// iterate over array of applications ids
		applications.forEach(function (obj) {
			if ( applicationsIds.indexOf(obj._id) > -1 ) {
				// remove the JobApplication
				this.candidate.applications.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more interviewsIds as a Interviews
	// to a Candidate
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInterviews( candidateId, interviewsIds ): Observable<any> {

		// get the Candidate
		this.loadHelper( candidateId );

	// split on a comma with no spaces
	var idList = interviewsIds.split(',')

	// iterate over array of interviews ids
	idList.forEach(function (id) {
		// read the Interview
		var interview = new InterviewService(this.http).getInterview(id);
		// add the Interview if not already assigned
		if ( this.candidate.interviews.indexOf(interview) == -1 )
		this.candidate.interviews.push(interview);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more interviewsIds as a Interviews
	// from a Candidate
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInterviews( candidateId, interviewsIds ): Observable<any> {

		// get the Candidate
		this.loadHelper( candidateId );


	// split on a comma with no spaces
	var idList 					= interviewsIds.split(',');
	var interviews 	= this.candidate.interviews;

	if ( interviews != null && interviewsIds != null ) {

		// iterate over array of interviews ids
		interviews.forEach(function (obj) {
			if ( interviewsIds.indexOf(obj._id) > -1 ) {
				// remove the Interview
				this.candidate.interviews.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more offersIds as a Offers
	// to a Candidate
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOffers( candidateId, offersIds ): Observable<any> {

		// get the Candidate
		this.loadHelper( candidateId );

	// split on a comma with no spaces
	var idList = offersIds.split(',')

	// iterate over array of offers ids
	idList.forEach(function (id) {
		// read the Offer
		var offer = new OfferService(this.http).getOffer(id);
		// add the Offer if not already assigned
		if ( this.candidate.offers.indexOf(offer) == -1 )
		this.candidate.offers.push(offer);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more offersIds as a Offers
	// from a Candidate
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOffers( candidateId, offersIds ): Observable<any> {

		// get the Candidate
		this.loadHelper( candidateId );


	// split on a comma with no spaces
	var idList 					= offersIds.split(',');
	var offers 	= this.candidate.offers;

	if ( offers != null && offersIds != null ) {

		// iterate over array of offers ids
		offers.forEach(function (obj) {
			if ( offersIds.indexOf(obj._id) > -1 ) {
				// remove the Offer
				this.candidate.offers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more documentsIds as a Documents
	// to a Candidate
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDocuments( candidateId, documentsIds ): Observable<any> {

		// get the Candidate
		this.loadHelper( candidateId );

	// split on a comma with no spaces
	var idList = documentsIds.split(',')

	// iterate over array of documents ids
	idList.forEach(function (id) {
		// read the Document
		var document = new DocumentService(this.http).getDocument(id);
		// add the Document if not already assigned
		if ( this.candidate.documents.indexOf(document) == -1 )
		this.candidate.documents.push(document);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more documentsIds as a Documents
	// from a Candidate
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDocuments( candidateId, documentsIds ): Observable<any> {

		// get the Candidate
		this.loadHelper( candidateId );


	// split on a comma with no spaces
	var idList 					= documentsIds.split(',');
	var documents 	= this.candidate.documents;

	if ( documents != null && documentsIds != null ) {

		// iterate over array of documents ids
		documents.forEach(function (obj) {
			if ( documentsIds.indexOf(obj._id) > -1 ) {
				// remove the Document
				this.candidate.documents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Candidate
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Candidate/update/' + this.candidate;

	return  this.http.post(uri_, this.candidate );
}

	//********************************************************************
	// loadHelper - internal helper to load a Candidate
	//********************************************************************	
	loadHelper( id ) {
		this.getCandidate(id)
			.subscribe((res : Candidate) => {
				this.candidate = res;
			});
	}
}