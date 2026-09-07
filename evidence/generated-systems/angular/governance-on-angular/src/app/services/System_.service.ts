import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {System_} from '../models/System_';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {RecordsRepositoryService} from '../services/RecordsRepository.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class System_Service extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	system_ : System_;

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
	// add a System_
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSystem_(name, ownerDepartment, ProcessingActivities, RecordsRepositories, SystemType) : Observable<any> {
		const uri_ = this.apiUrl + '/System_/create';
		const obj = {
			      		name: name,
      		ownerDepartment: ownerDepartment,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		RecordsRepositories: RecordsRepositories != null && RecordsRepositories.length > 0 ? RecordsRepositories : null,
			SystemType: SystemType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a System_
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSystem_(name, ownerDepartment, ProcessingActivities, RecordsRepositories, SystemType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/System_/update/' + id;
		const obj = {
				      		name: name,
      		ownerDepartment: ownerDepartment,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		RecordsRepositories: RecordsRepositories != null && RecordsRepositories.length > 0 ? RecordsRepositories : null,
			SystemType: SystemType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a System_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSystem_(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/System_/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a System_
	// returns the results untouched as an Observable System_
	// System_ model
	// delegates via URI
	//********************************************************************
	getSystem_(id) : Observable<System_> {
		const uri_ = this.apiUrl + '/System_/load/' + id;

		return this.http.get<System_>(uri_);
	}
	
	//********************************************************************
	// gets all System_
	// returns the results untouched as JSON representation of an
	// Observable array of System_ models
	// delegates via URI
	//********************************************************************
	getSystem_s() : Observable<System_[]> {
		const uri_ = this.apiUrl + '/System_/';

		return this
			.http.get<System_[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more processingActivitiesIds as a ProcessingActivities
	// to a System_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcessingActivities( system_Id, processingActivitiesIds ): Observable<any> {

		// get the System_
		this.loadHelper( system_Id );

	// split on a comma with no spaces
	var idList = processingActivitiesIds.split(',')

	// iterate over array of processingActivities ids
	idList.forEach(function (id) {
		// read the DataProcessingActivity
		var dataProcessingActivity = new DataProcessingActivityService(this.http).getDataProcessingActivity(id);
		// add the DataProcessingActivity if not already assigned
		if ( this.system_.processingActivities.indexOf(dataProcessingActivity) == -1 )
		this.system_.processingActivities.push(dataProcessingActivity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more processingActivitiesIds as a ProcessingActivities
	// from a System_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcessingActivities( system_Id, processingActivitiesIds ): Observable<any> {

		// get the System_
		this.loadHelper( system_Id );


	// split on a comma with no spaces
	var idList 					= processingActivitiesIds.split(',');
	var processingActivities 	= this.system_.processingActivities;

	if ( processingActivities != null && processingActivitiesIds != null ) {

		// iterate over array of processingActivities ids
		processingActivities.forEach(function (obj) {
			if ( processingActivitiesIds.indexOf(obj._id) > -1 ) {
				// remove the DataProcessingActivity
				this.system_.processingActivities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more recordsRepositoriesIds as a RecordsRepositories
	// to a System_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRecordsRepositories( system_Id, recordsRepositoriesIds ): Observable<any> {

		// get the System_
		this.loadHelper( system_Id );

	// split on a comma with no spaces
	var idList = recordsRepositoriesIds.split(',')

	// iterate over array of recordsRepositories ids
	idList.forEach(function (id) {
		// read the RecordsRepository
		var recordsRepository = new RecordsRepositoryService(this.http).getRecordsRepository(id);
		// add the RecordsRepository if not already assigned
		if ( this.system_.recordsRepositories.indexOf(recordsRepository) == -1 )
		this.system_.recordsRepositories.push(recordsRepository);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more recordsRepositoriesIds as a RecordsRepositories
	// from a System_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRecordsRepositories( system_Id, recordsRepositoriesIds ): Observable<any> {

		// get the System_
		this.loadHelper( system_Id );


	// split on a comma with no spaces
	var idList 					= recordsRepositoriesIds.split(',');
	var recordsRepositories 	= this.system_.recordsRepositories;

	if ( recordsRepositories != null && recordsRepositoriesIds != null ) {

		// iterate over array of recordsRepositories ids
		recordsRepositories.forEach(function (obj) {
			if ( recordsRepositoriesIds.indexOf(obj._id) > -1 ) {
				// remove the RecordsRepository
				this.system_.recordsRepositories.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a System_
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/System_/update/' + this.system_;

	return  this.http.post(uri_, this.system_ );
}

	//********************************************************************
	// loadHelper - internal helper to load a System_
	//********************************************************************	
	loadHelper( id ) {
		this.getSystem_(id)
			.subscribe((res : System_) => {
				this.system_ = res;
			});
	}
}