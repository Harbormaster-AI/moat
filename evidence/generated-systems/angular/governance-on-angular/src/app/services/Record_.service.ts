import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Record_} from '../models/Record_';
import {RecordsRepositoryService} from '../services/RecordsRepository.service';
import {RetentionScheduleService} from '../services/RetentionSchedule.service';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {DataCategoryService} from '../services/DataCategory.service';
import {LegalHoldService} from '../services/LegalHold.service';
import {DataSubjectRequestService} from '../services/DataSubjectRequest.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class Record_Service extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	record_ : Record_;

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
	// add a Record_
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRecord_(title, creationDate, Repository, RetentionSchedule, ProcessingActivities, DataCategories, LegalHolds, DataSubjectRequests, RecordType, Classification, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Record_/create';
		const obj = {
			      		title: title,
      		creationDate: creationDate,
      		Repository: Repository != null && Repository.length > 0 ? Repository : null,
      		RetentionSchedule: RetentionSchedule != null && RetentionSchedule.length > 0 ? RetentionSchedule : null,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		DataCategories: DataCategories != null && DataCategories.length > 0 ? DataCategories : null,
      		LegalHolds: LegalHolds != null && LegalHolds.length > 0 ? LegalHolds : null,
      		DataSubjectRequests: DataSubjectRequests != null && DataSubjectRequests.length > 0 ? DataSubjectRequests : null,
      		RecordType: RecordType,
      		Classification: Classification,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Record_
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRecord_(title, creationDate, Repository, RetentionSchedule, ProcessingActivities, DataCategories, LegalHolds, DataSubjectRequests, RecordType, Classification, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Record_/update/' + id;
		const obj = {
				      		title: title,
      		creationDate: creationDate,
      		Repository: Repository != null && Repository.length > 0 ? Repository : null,
      		RetentionSchedule: RetentionSchedule != null && RetentionSchedule.length > 0 ? RetentionSchedule : null,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		DataCategories: DataCategories != null && DataCategories.length > 0 ? DataCategories : null,
      		LegalHolds: LegalHolds != null && LegalHolds.length > 0 ? LegalHolds : null,
      		DataSubjectRequests: DataSubjectRequests != null && DataSubjectRequests.length > 0 ? DataSubjectRequests : null,
      		RecordType: RecordType,
      		Classification: Classification,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Record_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRecord_(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Record_/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Record_
	// returns the results untouched as an Observable Record_
	// Record_ model
	// delegates via URI
	//********************************************************************
	getRecord_(id) : Observable<Record_> {
		const uri_ = this.apiUrl + '/Record_/load/' + id;

		return this.http.get<Record_>(uri_);
	}
	
	//********************************************************************
	// gets all Record_
	// returns the results untouched as JSON representation of an
	// Observable array of Record_ models
	// delegates via URI
	//********************************************************************
	getRecord_s() : Observable<Record_[]> {
		const uri_ = this.apiUrl + '/Record_/';

		return this
			.http.get<Record_[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Repository on a Record_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRepository( record_Id, _repositoryId ): Observable<any> {

		// get the Record_ from storage
		this.loadHelper( record_Id );

	// get the RecordsRepository from storage
	var tmp 	= new RecordsRepositoryService(this.http).getRecordsRepository(_repositoryId);

	// assign the Repository
	this.record_.repository = tmp;

	// save the Record_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Repository on a Record_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRepository( record_Id ): Observable<any> {

		// get the Record_ from storage
		this.loadHelper( record_Id );

	// assign Repository to null
	this.record_.repository = null;

	// save the Record_
	return this.saveHelper();
}

		//********************************************************************
	// assigns a RetentionSchedule on a Record_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRetentionSchedule( record_Id, _retentionScheduleId ): Observable<any> {

		// get the Record_ from storage
		this.loadHelper( record_Id );

	// get the RetentionSchedule from storage
	var tmp 	= new RetentionScheduleService(this.http).getRetentionSchedule(_retentionScheduleId);

	// assign the RetentionSchedule
	this.record_.retentionSchedule = tmp;

	// save the Record_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a RetentionSchedule on a Record_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRetentionSchedule( record_Id ): Observable<any> {

		// get the Record_ from storage
		this.loadHelper( record_Id );

	// assign RetentionSchedule to null
	this.record_.retentionSchedule = null;

	// save the Record_
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more processingActivitiesIds as a ProcessingActivities
	// to a Record_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcessingActivities( record_Id, processingActivitiesIds ): Observable<any> {

		// get the Record_
		this.loadHelper( record_Id );

	// split on a comma with no spaces
	var idList = processingActivitiesIds.split(',')

	// iterate over array of processingActivities ids
	idList.forEach(function (id) {
		// read the DataProcessingActivity
		var dataProcessingActivity = new DataProcessingActivityService(this.http).getDataProcessingActivity(id);
		// add the DataProcessingActivity if not already assigned
		if ( this.record_.processingActivities.indexOf(dataProcessingActivity) == -1 )
		this.record_.processingActivities.push(dataProcessingActivity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more processingActivitiesIds as a ProcessingActivities
	// from a Record_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcessingActivities( record_Id, processingActivitiesIds ): Observable<any> {

		// get the Record_
		this.loadHelper( record_Id );


	// split on a comma with no spaces
	var idList 					= processingActivitiesIds.split(',');
	var processingActivities 	= this.record_.processingActivities;

	if ( processingActivities != null && processingActivitiesIds != null ) {

		// iterate over array of processingActivities ids
		processingActivities.forEach(function (obj) {
			if ( processingActivitiesIds.indexOf(obj._id) > -1 ) {
				// remove the DataProcessingActivity
				this.record_.processingActivities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataCategoriesIds as a DataCategories
	// to a Record_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataCategories( record_Id, dataCategoriesIds ): Observable<any> {

		// get the Record_
		this.loadHelper( record_Id );

	// split on a comma with no spaces
	var idList = dataCategoriesIds.split(',')

	// iterate over array of dataCategories ids
	idList.forEach(function (id) {
		// read the DataCategory
		var dataCategory = new DataCategoryService(this.http).getDataCategory(id);
		// add the DataCategory if not already assigned
		if ( this.record_.dataCategories.indexOf(dataCategory) == -1 )
		this.record_.dataCategories.push(dataCategory);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataCategoriesIds as a DataCategories
	// from a Record_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataCategories( record_Id, dataCategoriesIds ): Observable<any> {

		// get the Record_
		this.loadHelper( record_Id );


	// split on a comma with no spaces
	var idList 					= dataCategoriesIds.split(',');
	var dataCategories 	= this.record_.dataCategories;

	if ( dataCategories != null && dataCategoriesIds != null ) {

		// iterate over array of dataCategories ids
		dataCategories.forEach(function (obj) {
			if ( dataCategoriesIds.indexOf(obj._id) > -1 ) {
				// remove the DataCategory
				this.record_.dataCategories.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more legalHoldsIds as a LegalHolds
	// to a Record_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLegalHolds( record_Id, legalHoldsIds ): Observable<any> {

		// get the Record_
		this.loadHelper( record_Id );

	// split on a comma with no spaces
	var idList = legalHoldsIds.split(',')

	// iterate over array of legalHolds ids
	idList.forEach(function (id) {
		// read the LegalHold
		var legalHold = new LegalHoldService(this.http).getLegalHold(id);
		// add the LegalHold if not already assigned
		if ( this.record_.legalHolds.indexOf(legalHold) == -1 )
		this.record_.legalHolds.push(legalHold);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more legalHoldsIds as a LegalHolds
	// from a Record_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLegalHolds( record_Id, legalHoldsIds ): Observable<any> {

		// get the Record_
		this.loadHelper( record_Id );


	// split on a comma with no spaces
	var idList 					= legalHoldsIds.split(',');
	var legalHolds 	= this.record_.legalHolds;

	if ( legalHolds != null && legalHoldsIds != null ) {

		// iterate over array of legalHolds ids
		legalHolds.forEach(function (obj) {
			if ( legalHoldsIds.indexOf(obj._id) > -1 ) {
				// remove the LegalHold
				this.record_.legalHolds.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataSubjectRequestsIds as a DataSubjectRequests
	// to a Record_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataSubjectRequests( record_Id, dataSubjectRequestsIds ): Observable<any> {

		// get the Record_
		this.loadHelper( record_Id );

	// split on a comma with no spaces
	var idList = dataSubjectRequestsIds.split(',')

	// iterate over array of dataSubjectRequests ids
	idList.forEach(function (id) {
		// read the DataSubjectRequest
		var dataSubjectRequest = new DataSubjectRequestService(this.http).getDataSubjectRequest(id);
		// add the DataSubjectRequest if not already assigned
		if ( this.record_.dataSubjectRequests.indexOf(dataSubjectRequest) == -1 )
		this.record_.dataSubjectRequests.push(dataSubjectRequest);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataSubjectRequestsIds as a DataSubjectRequests
	// from a Record_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataSubjectRequests( record_Id, dataSubjectRequestsIds ): Observable<any> {

		// get the Record_
		this.loadHelper( record_Id );


	// split on a comma with no spaces
	var idList 					= dataSubjectRequestsIds.split(',');
	var dataSubjectRequests 	= this.record_.dataSubjectRequests;

	if ( dataSubjectRequests != null && dataSubjectRequestsIds != null ) {

		// iterate over array of dataSubjectRequests ids
		dataSubjectRequests.forEach(function (obj) {
			if ( dataSubjectRequestsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSubjectRequest
				this.record_.dataSubjectRequests.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Record_
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Record_/update/' + this.record_;

	return  this.http.post(uri_, this.record_ );
}

	//********************************************************************
	// loadHelper - internal helper to load a Record_
	//********************************************************************	
	loadHelper( id ) {
		this.getRecord_(id)
			.subscribe((res : Record_) => {
				this.record_ = res;
			});
	}
}