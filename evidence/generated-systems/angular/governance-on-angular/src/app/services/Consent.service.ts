import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Consent} from '../models/Consent';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {PrivacyNoticeService} from '../services/PrivacyNotice.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ConsentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	consent : Consent;

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
	// add a Consent
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addConsent(subjectIdentifier, captureDate, expiryDate, ProcessingActivities, PrivacyNotice, ConsentType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Consent/create';
		const obj = {
			      		subjectIdentifier: subjectIdentifier,
      		captureDate: captureDate,
      		expiryDate: expiryDate,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		PrivacyNotice: PrivacyNotice != null && PrivacyNotice.length > 0 ? PrivacyNotice : null,
      		ConsentType: ConsentType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Consent
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateConsent(subjectIdentifier, captureDate, expiryDate, ProcessingActivities, PrivacyNotice, ConsentType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Consent/update/' + id;
		const obj = {
				      		subjectIdentifier: subjectIdentifier,
      		captureDate: captureDate,
      		expiryDate: expiryDate,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		PrivacyNotice: PrivacyNotice != null && PrivacyNotice.length > 0 ? PrivacyNotice : null,
      		ConsentType: ConsentType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Consent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteConsent(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Consent/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Consent
	// returns the results untouched as an Observable Consent
	// Consent model
	// delegates via URI
	//********************************************************************
	getConsent(id) : Observable<Consent> {
		const uri_ = this.apiUrl + '/Consent/load/' + id;

		return this.http.get<Consent>(uri_);
	}
	
	//********************************************************************
	// gets all Consent
	// returns the results untouched as JSON representation of an
	// Observable array of Consent models
	// delegates via URI
	//********************************************************************
	getConsents() : Observable<Consent[]> {
		const uri_ = this.apiUrl + '/Consent/';

		return this
			.http.get<Consent[]>(uri_);
	}
	
			//********************************************************************
	// assigns a PrivacyNotice on a Consent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPrivacyNotice( consentId, _privacyNoticeId ): Observable<any> {

		// get the Consent from storage
		this.loadHelper( consentId );

	// get the PrivacyNotice from storage
	var tmp 	= new PrivacyNoticeService(this.http).getPrivacyNotice(_privacyNoticeId);

	// assign the PrivacyNotice
	this.consent.privacyNotice = tmp;

	// save the Consent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PrivacyNotice on a Consent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPrivacyNotice( consentId ): Observable<any> {

		// get the Consent from storage
		this.loadHelper( consentId );

	// assign PrivacyNotice to null
	this.consent.privacyNotice = null;

	// save the Consent
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more processingActivitiesIds as a ProcessingActivities
	// to a Consent
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcessingActivities( consentId, processingActivitiesIds ): Observable<any> {

		// get the Consent
		this.loadHelper( consentId );

	// split on a comma with no spaces
	var idList = processingActivitiesIds.split(',')

	// iterate over array of processingActivities ids
	idList.forEach(function (id) {
		// read the DataProcessingActivity
		var dataProcessingActivity = new DataProcessingActivityService(this.http).getDataProcessingActivity(id);
		// add the DataProcessingActivity if not already assigned
		if ( this.consent.processingActivities.indexOf(dataProcessingActivity) == -1 )
		this.consent.processingActivities.push(dataProcessingActivity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more processingActivitiesIds as a ProcessingActivities
	// from a Consent
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcessingActivities( consentId, processingActivitiesIds ): Observable<any> {

		// get the Consent
		this.loadHelper( consentId );


	// split on a comma with no spaces
	var idList 					= processingActivitiesIds.split(',');
	var processingActivities 	= this.consent.processingActivities;

	if ( processingActivities != null && processingActivitiesIds != null ) {

		// iterate over array of processingActivities ids
		processingActivities.forEach(function (obj) {
			if ( processingActivitiesIds.indexOf(obj._id) > -1 ) {
				// remove the DataProcessingActivity
				this.consent.processingActivities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Consent
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Consent/update/' + this.consent;

	return  this.http.post(uri_, this.consent );
}

	//********************************************************************
	// loadHelper - internal helper to load a Consent
	//********************************************************************	
	loadHelper( id ) {
		this.getConsent(id)
			.subscribe((res : Consent) => {
				this.consent = res;
			});
	}
}