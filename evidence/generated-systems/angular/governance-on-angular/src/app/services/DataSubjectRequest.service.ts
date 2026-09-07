import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DataSubjectRequest} from '../models/DataSubjectRequest';
import {OrganizationService} from '../services/Organization.service';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {Record_Service} from '../services/Record_.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DataSubjectRequestService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dataSubjectRequest : DataSubjectRequest;

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
	// add a DataSubjectRequest
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDataSubjectRequest(receivedDate, dueDate, requesterCountry, Organization, ProcessingActivities, Records, RequestType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/DataSubjectRequest/create';
		const obj = {
			      		receivedDate: receivedDate,
      		dueDate: dueDate,
      		requesterCountry: requesterCountry,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		RequestType: RequestType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DataSubjectRequest
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDataSubjectRequest(receivedDate, dueDate, requesterCountry, Organization, ProcessingActivities, Records, RequestType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DataSubjectRequest/update/' + id;
		const obj = {
				      		receivedDate: receivedDate,
      		dueDate: dueDate,
      		requesterCountry: requesterCountry,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		RequestType: RequestType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DataSubjectRequest
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDataSubjectRequest(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DataSubjectRequest/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DataSubjectRequest
	// returns the results untouched as an Observable DataSubjectRequest
	// DataSubjectRequest model
	// delegates via URI
	//********************************************************************
	getDataSubjectRequest(id) : Observable<DataSubjectRequest> {
		const uri_ = this.apiUrl + '/DataSubjectRequest/load/' + id;

		return this.http.get<DataSubjectRequest>(uri_);
	}
	
	//********************************************************************
	// gets all DataSubjectRequest
	// returns the results untouched as JSON representation of an
	// Observable array of DataSubjectRequest models
	// delegates via URI
	//********************************************************************
	getDataSubjectRequests() : Observable<DataSubjectRequest[]> {
		const uri_ = this.apiUrl + '/DataSubjectRequest/';

		return this
			.http.get<DataSubjectRequest[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a DataSubjectRequest
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( dataSubjectRequestId, _organizationId ): Observable<any> {

		// get the DataSubjectRequest from storage
		this.loadHelper( dataSubjectRequestId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.dataSubjectRequest.organization = tmp;

	// save the DataSubjectRequest
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a DataSubjectRequest
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( dataSubjectRequestId ): Observable<any> {

		// get the DataSubjectRequest from storage
		this.loadHelper( dataSubjectRequestId );

	// assign Organization to null
	this.dataSubjectRequest.organization = null;

	// save the DataSubjectRequest
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more processingActivitiesIds as a ProcessingActivities
	// to a DataSubjectRequest
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcessingActivities( dataSubjectRequestId, processingActivitiesIds ): Observable<any> {

		// get the DataSubjectRequest
		this.loadHelper( dataSubjectRequestId );

	// split on a comma with no spaces
	var idList = processingActivitiesIds.split(',')

	// iterate over array of processingActivities ids
	idList.forEach(function (id) {
		// read the DataProcessingActivity
		var dataProcessingActivity = new DataProcessingActivityService(this.http).getDataProcessingActivity(id);
		// add the DataProcessingActivity if not already assigned
		if ( this.dataSubjectRequest.processingActivities.indexOf(dataProcessingActivity) == -1 )
		this.dataSubjectRequest.processingActivities.push(dataProcessingActivity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more processingActivitiesIds as a ProcessingActivities
	// from a DataSubjectRequest
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcessingActivities( dataSubjectRequestId, processingActivitiesIds ): Observable<any> {

		// get the DataSubjectRequest
		this.loadHelper( dataSubjectRequestId );


	// split on a comma with no spaces
	var idList 					= processingActivitiesIds.split(',');
	var processingActivities 	= this.dataSubjectRequest.processingActivities;

	if ( processingActivities != null && processingActivitiesIds != null ) {

		// iterate over array of processingActivities ids
		processingActivities.forEach(function (obj) {
			if ( processingActivitiesIds.indexOf(obj._id) > -1 ) {
				// remove the DataProcessingActivity
				this.dataSubjectRequest.processingActivities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more recordsIds as a Records
	// to a DataSubjectRequest
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRecords( dataSubjectRequestId, recordsIds ): Observable<any> {

		// get the DataSubjectRequest
		this.loadHelper( dataSubjectRequestId );

	// split on a comma with no spaces
	var idList = recordsIds.split(',')

	// iterate over array of records ids
	idList.forEach(function (id) {
		// read the Record_
		var record_ = new Record_Service(this.http).getRecord_(id);
		// add the Record_ if not already assigned
		if ( this.dataSubjectRequest.records.indexOf(record_) == -1 )
		this.dataSubjectRequest.records.push(record_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more recordsIds as a Records
	// from a DataSubjectRequest
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRecords( dataSubjectRequestId, recordsIds ): Observable<any> {

		// get the DataSubjectRequest
		this.loadHelper( dataSubjectRequestId );


	// split on a comma with no spaces
	var idList 					= recordsIds.split(',');
	var records 	= this.dataSubjectRequest.records;

	if ( records != null && recordsIds != null ) {

		// iterate over array of records ids
		records.forEach(function (obj) {
			if ( recordsIds.indexOf(obj._id) > -1 ) {
				// remove the Record_
				this.dataSubjectRequest.records.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DataSubjectRequest
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DataSubjectRequest/update/' + this.dataSubjectRequest;

	return  this.http.post(uri_, this.dataSubjectRequest );
}

	//********************************************************************
	// loadHelper - internal helper to load a DataSubjectRequest
	//********************************************************************	
	loadHelper( id ) {
		this.getDataSubjectRequest(id)
			.subscribe((res : DataSubjectRequest) => {
				this.dataSubjectRequest = res;
			});
	}
}