import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PrivacyNotice} from '../models/PrivacyNotice';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {OrganizationService} from '../services/Organization.service';
import {ConsentService} from '../services/Consent.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PrivacyNoticeService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	privacyNotice : PrivacyNotice;

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
	// add a PrivacyNotice
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPrivacyNotice(title, audience, versionLabel, publicationDate, publicationUrl, ProcessingActivities, Organization, Consents, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/PrivacyNotice/create';
		const obj = {
			      		title: title,
      		audience: audience,
      		versionLabel: versionLabel,
      		publicationDate: publicationDate,
      		publicationUrl: publicationUrl,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Consents: Consents != null && Consents.length > 0 ? Consents : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PrivacyNotice
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePrivacyNotice(title, audience, versionLabel, publicationDate, publicationUrl, ProcessingActivities, Organization, Consents, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PrivacyNotice/update/' + id;
		const obj = {
				      		title: title,
      		audience: audience,
      		versionLabel: versionLabel,
      		publicationDate: publicationDate,
      		publicationUrl: publicationUrl,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Consents: Consents != null && Consents.length > 0 ? Consents : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PrivacyNotice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePrivacyNotice(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PrivacyNotice/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PrivacyNotice
	// returns the results untouched as an Observable PrivacyNotice
	// PrivacyNotice model
	// delegates via URI
	//********************************************************************
	getPrivacyNotice(id) : Observable<PrivacyNotice> {
		const uri_ = this.apiUrl + '/PrivacyNotice/load/' + id;

		return this.http.get<PrivacyNotice>(uri_);
	}
	
	//********************************************************************
	// gets all PrivacyNotice
	// returns the results untouched as JSON representation of an
	// Observable array of PrivacyNotice models
	// delegates via URI
	//********************************************************************
	getPrivacyNotices() : Observable<PrivacyNotice[]> {
		const uri_ = this.apiUrl + '/PrivacyNotice/';

		return this
			.http.get<PrivacyNotice[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a PrivacyNotice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( privacyNoticeId, _organizationId ): Observable<any> {

		// get the PrivacyNotice from storage
		this.loadHelper( privacyNoticeId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.privacyNotice.organization = tmp;

	// save the PrivacyNotice
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a PrivacyNotice
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( privacyNoticeId ): Observable<any> {

		// get the PrivacyNotice from storage
		this.loadHelper( privacyNoticeId );

	// assign Organization to null
	this.privacyNotice.organization = null;

	// save the PrivacyNotice
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more processingActivitiesIds as a ProcessingActivities
	// to a PrivacyNotice
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcessingActivities( privacyNoticeId, processingActivitiesIds ): Observable<any> {

		// get the PrivacyNotice
		this.loadHelper( privacyNoticeId );

	// split on a comma with no spaces
	var idList = processingActivitiesIds.split(',')

	// iterate over array of processingActivities ids
	idList.forEach(function (id) {
		// read the DataProcessingActivity
		var dataProcessingActivity = new DataProcessingActivityService(this.http).getDataProcessingActivity(id);
		// add the DataProcessingActivity if not already assigned
		if ( this.privacyNotice.processingActivities.indexOf(dataProcessingActivity) == -1 )
		this.privacyNotice.processingActivities.push(dataProcessingActivity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more processingActivitiesIds as a ProcessingActivities
	// from a PrivacyNotice
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcessingActivities( privacyNoticeId, processingActivitiesIds ): Observable<any> {

		// get the PrivacyNotice
		this.loadHelper( privacyNoticeId );


	// split on a comma with no spaces
	var idList 					= processingActivitiesIds.split(',');
	var processingActivities 	= this.privacyNotice.processingActivities;

	if ( processingActivities != null && processingActivitiesIds != null ) {

		// iterate over array of processingActivities ids
		processingActivities.forEach(function (obj) {
			if ( processingActivitiesIds.indexOf(obj._id) > -1 ) {
				// remove the DataProcessingActivity
				this.privacyNotice.processingActivities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more consentsIds as a Consents
	// to a PrivacyNotice
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addConsents( privacyNoticeId, consentsIds ): Observable<any> {

		// get the PrivacyNotice
		this.loadHelper( privacyNoticeId );

	// split on a comma with no spaces
	var idList = consentsIds.split(',')

	// iterate over array of consents ids
	idList.forEach(function (id) {
		// read the Consent
		var consent = new ConsentService(this.http).getConsent(id);
		// add the Consent if not already assigned
		if ( this.privacyNotice.consents.indexOf(consent) == -1 )
		this.privacyNotice.consents.push(consent);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more consentsIds as a Consents
	// from a PrivacyNotice
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeConsents( privacyNoticeId, consentsIds ): Observable<any> {

		// get the PrivacyNotice
		this.loadHelper( privacyNoticeId );


	// split on a comma with no spaces
	var idList 					= consentsIds.split(',');
	var consents 	= this.privacyNotice.consents;

	if ( consents != null && consentsIds != null ) {

		// iterate over array of consents ids
		consents.forEach(function (obj) {
			if ( consentsIds.indexOf(obj._id) > -1 ) {
				// remove the Consent
				this.privacyNotice.consents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PrivacyNotice
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PrivacyNotice/update/' + this.privacyNotice;

	return  this.http.post(uri_, this.privacyNotice );
}

	//********************************************************************
	// loadHelper - internal helper to load a PrivacyNotice
	//********************************************************************	
	loadHelper( id ) {
		this.getPrivacyNotice(id)
			.subscribe((res : PrivacyNotice) => {
				this.privacyNotice = res;
			});
	}
}