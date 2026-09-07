import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {RetentionSchedule} from '../models/RetentionSchedule';
import {RecordsRepositoryService} from '../services/RecordsRepository.service';
import {Record_Service} from '../services/Record_.service';
import {Exception_Service} from '../services/Exception_.service';
import {DispositionReviewService} from '../services/DispositionReview.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RetentionScheduleService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	retentionSchedule : RetentionSchedule;

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
	// add a RetentionSchedule
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRetentionSchedule(name, retentionPeriodMonths, Repositories, Records, Exceptions, DispositionReviews, RetentionTrigger, DispositionAction, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/RetentionSchedule/create';
		const obj = {
			      		name: name,
      		retentionPeriodMonths: retentionPeriodMonths,
      		Repositories: Repositories != null && Repositories.length > 0 ? Repositories : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		Exceptions: Exceptions != null && Exceptions.length > 0 ? Exceptions : null,
      		DispositionReviews: DispositionReviews != null && DispositionReviews.length > 0 ? DispositionReviews : null,
      		RetentionTrigger: RetentionTrigger,
      		DispositionAction: DispositionAction,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a RetentionSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRetentionSchedule(name, retentionPeriodMonths, Repositories, Records, Exceptions, DispositionReviews, RetentionTrigger, DispositionAction, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/RetentionSchedule/update/' + id;
		const obj = {
				      		name: name,
      		retentionPeriodMonths: retentionPeriodMonths,
      		Repositories: Repositories != null && Repositories.length > 0 ? Repositories : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		Exceptions: Exceptions != null && Exceptions.length > 0 ? Exceptions : null,
      		DispositionReviews: DispositionReviews != null && DispositionReviews.length > 0 ? DispositionReviews : null,
      		RetentionTrigger: RetentionTrigger,
      		DispositionAction: DispositionAction,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a RetentionSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRetentionSchedule(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/RetentionSchedule/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a RetentionSchedule
	// returns the results untouched as an Observable RetentionSchedule
	// RetentionSchedule model
	// delegates via URI
	//********************************************************************
	getRetentionSchedule(id) : Observable<RetentionSchedule> {
		const uri_ = this.apiUrl + '/RetentionSchedule/load/' + id;

		return this.http.get<RetentionSchedule>(uri_);
	}
	
	//********************************************************************
	// gets all RetentionSchedule
	// returns the results untouched as JSON representation of an
	// Observable array of RetentionSchedule models
	// delegates via URI
	//********************************************************************
	getRetentionSchedules() : Observable<RetentionSchedule[]> {
		const uri_ = this.apiUrl + '/RetentionSchedule/';

		return this
			.http.get<RetentionSchedule[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more repositoriesIds as a Repositories
	// to a RetentionSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRepositories( retentionScheduleId, repositoriesIds ): Observable<any> {

		// get the RetentionSchedule
		this.loadHelper( retentionScheduleId );

	// split on a comma with no spaces
	var idList = repositoriesIds.split(',')

	// iterate over array of repositories ids
	idList.forEach(function (id) {
		// read the RecordsRepository
		var recordsRepository = new RecordsRepositoryService(this.http).getRecordsRepository(id);
		// add the RecordsRepository if not already assigned
		if ( this.retentionSchedule.repositories.indexOf(recordsRepository) == -1 )
		this.retentionSchedule.repositories.push(recordsRepository);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more repositoriesIds as a Repositories
	// from a RetentionSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRepositories( retentionScheduleId, repositoriesIds ): Observable<any> {

		// get the RetentionSchedule
		this.loadHelper( retentionScheduleId );


	// split on a comma with no spaces
	var idList 					= repositoriesIds.split(',');
	var repositories 	= this.retentionSchedule.repositories;

	if ( repositories != null && repositoriesIds != null ) {

		// iterate over array of repositories ids
		repositories.forEach(function (obj) {
			if ( repositoriesIds.indexOf(obj._id) > -1 ) {
				// remove the RecordsRepository
				this.retentionSchedule.repositories.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more recordsIds as a Records
	// to a RetentionSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRecords( retentionScheduleId, recordsIds ): Observable<any> {

		// get the RetentionSchedule
		this.loadHelper( retentionScheduleId );

	// split on a comma with no spaces
	var idList = recordsIds.split(',')

	// iterate over array of records ids
	idList.forEach(function (id) {
		// read the Record_
		var record_ = new Record_Service(this.http).getRecord_(id);
		// add the Record_ if not already assigned
		if ( this.retentionSchedule.records.indexOf(record_) == -1 )
		this.retentionSchedule.records.push(record_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more recordsIds as a Records
	// from a RetentionSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRecords( retentionScheduleId, recordsIds ): Observable<any> {

		// get the RetentionSchedule
		this.loadHelper( retentionScheduleId );


	// split on a comma with no spaces
	var idList 					= recordsIds.split(',');
	var records 	= this.retentionSchedule.records;

	if ( records != null && recordsIds != null ) {

		// iterate over array of records ids
		records.forEach(function (obj) {
			if ( recordsIds.indexOf(obj._id) > -1 ) {
				// remove the Record_
				this.retentionSchedule.records.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more exceptionsIds as a Exceptions
	// to a RetentionSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addExceptions( retentionScheduleId, exceptionsIds ): Observable<any> {

		// get the RetentionSchedule
		this.loadHelper( retentionScheduleId );

	// split on a comma with no spaces
	var idList = exceptionsIds.split(',')

	// iterate over array of exceptions ids
	idList.forEach(function (id) {
		// read the Exception_
		var exception_ = new Exception_Service(this.http).getException_(id);
		// add the Exception_ if not already assigned
		if ( this.retentionSchedule.exceptions.indexOf(exception_) == -1 )
		this.retentionSchedule.exceptions.push(exception_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more exceptionsIds as a Exceptions
	// from a RetentionSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeExceptions( retentionScheduleId, exceptionsIds ): Observable<any> {

		// get the RetentionSchedule
		this.loadHelper( retentionScheduleId );


	// split on a comma with no spaces
	var idList 					= exceptionsIds.split(',');
	var exceptions 	= this.retentionSchedule.exceptions;

	if ( exceptions != null && exceptionsIds != null ) {

		// iterate over array of exceptions ids
		exceptions.forEach(function (obj) {
			if ( exceptionsIds.indexOf(obj._id) > -1 ) {
				// remove the Exception_
				this.retentionSchedule.exceptions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dispositionReviewsIds as a DispositionReviews
	// to a RetentionSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDispositionReviews( retentionScheduleId, dispositionReviewsIds ): Observable<any> {

		// get the RetentionSchedule
		this.loadHelper( retentionScheduleId );

	// split on a comma with no spaces
	var idList = dispositionReviewsIds.split(',')

	// iterate over array of dispositionReviews ids
	idList.forEach(function (id) {
		// read the DispositionReview
		var dispositionReview = new DispositionReviewService(this.http).getDispositionReview(id);
		// add the DispositionReview if not already assigned
		if ( this.retentionSchedule.dispositionReviews.indexOf(dispositionReview) == -1 )
		this.retentionSchedule.dispositionReviews.push(dispositionReview);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dispositionReviewsIds as a DispositionReviews
	// from a RetentionSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDispositionReviews( retentionScheduleId, dispositionReviewsIds ): Observable<any> {

		// get the RetentionSchedule
		this.loadHelper( retentionScheduleId );


	// split on a comma with no spaces
	var idList 					= dispositionReviewsIds.split(',');
	var dispositionReviews 	= this.retentionSchedule.dispositionReviews;

	if ( dispositionReviews != null && dispositionReviewsIds != null ) {

		// iterate over array of dispositionReviews ids
		dispositionReviews.forEach(function (obj) {
			if ( dispositionReviewsIds.indexOf(obj._id) > -1 ) {
				// remove the DispositionReview
				this.retentionSchedule.dispositionReviews.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a RetentionSchedule
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/RetentionSchedule/update/' + this.retentionSchedule;

	return  this.http.post(uri_, this.retentionSchedule );
}

	//********************************************************************
	// loadHelper - internal helper to load a RetentionSchedule
	//********************************************************************	
	loadHelper( id ) {
		this.getRetentionSchedule(id)
			.subscribe((res : RetentionSchedule) => {
				this.retentionSchedule = res;
			});
	}
}