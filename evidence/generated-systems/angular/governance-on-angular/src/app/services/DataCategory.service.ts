import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DataCategory} from '../models/DataCategory';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {Record_Service} from '../services/Record_.service';
import {DataBreachService} from '../services/DataBreach.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DataCategoryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dataCategory : DataCategory;

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
	// add a DataCategory
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDataCategory(name, description, ProcessingActivities, Records, DataBreaches, Classification) : Observable<any> {
		const uri_ = this.apiUrl + '/DataCategory/create';
		const obj = {
			      		name: name,
      		description: description,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		DataBreaches: DataBreaches != null && DataBreaches.length > 0 ? DataBreaches : null,
			Classification: Classification
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DataCategory
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDataCategory(name, description, ProcessingActivities, Records, DataBreaches, Classification, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DataCategory/update/' + id;
		const obj = {
				      		name: name,
      		description: description,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		DataBreaches: DataBreaches != null && DataBreaches.length > 0 ? DataBreaches : null,
			Classification: Classification
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DataCategory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDataCategory(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DataCategory/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DataCategory
	// returns the results untouched as an Observable DataCategory
	// DataCategory model
	// delegates via URI
	//********************************************************************
	getDataCategory(id) : Observable<DataCategory> {
		const uri_ = this.apiUrl + '/DataCategory/load/' + id;

		return this.http.get<DataCategory>(uri_);
	}
	
	//********************************************************************
	// gets all DataCategory
	// returns the results untouched as JSON representation of an
	// Observable array of DataCategory models
	// delegates via URI
	//********************************************************************
	getDataCategorys() : Observable<DataCategory[]> {
		const uri_ = this.apiUrl + '/DataCategory/';

		return this
			.http.get<DataCategory[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more processingActivitiesIds as a ProcessingActivities
	// to a DataCategory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcessingActivities( dataCategoryId, processingActivitiesIds ): Observable<any> {

		// get the DataCategory
		this.loadHelper( dataCategoryId );

	// split on a comma with no spaces
	var idList = processingActivitiesIds.split(',')

	// iterate over array of processingActivities ids
	idList.forEach(function (id) {
		// read the DataProcessingActivity
		var dataProcessingActivity = new DataProcessingActivityService(this.http).getDataProcessingActivity(id);
		// add the DataProcessingActivity if not already assigned
		if ( this.dataCategory.processingActivities.indexOf(dataProcessingActivity) == -1 )
		this.dataCategory.processingActivities.push(dataProcessingActivity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more processingActivitiesIds as a ProcessingActivities
	// from a DataCategory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcessingActivities( dataCategoryId, processingActivitiesIds ): Observable<any> {

		// get the DataCategory
		this.loadHelper( dataCategoryId );


	// split on a comma with no spaces
	var idList 					= processingActivitiesIds.split(',');
	var processingActivities 	= this.dataCategory.processingActivities;

	if ( processingActivities != null && processingActivitiesIds != null ) {

		// iterate over array of processingActivities ids
		processingActivities.forEach(function (obj) {
			if ( processingActivitiesIds.indexOf(obj._id) > -1 ) {
				// remove the DataProcessingActivity
				this.dataCategory.processingActivities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more recordsIds as a Records
	// to a DataCategory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRecords( dataCategoryId, recordsIds ): Observable<any> {

		// get the DataCategory
		this.loadHelper( dataCategoryId );

	// split on a comma with no spaces
	var idList = recordsIds.split(',')

	// iterate over array of records ids
	idList.forEach(function (id) {
		// read the Record_
		var record_ = new Record_Service(this.http).getRecord_(id);
		// add the Record_ if not already assigned
		if ( this.dataCategory.records.indexOf(record_) == -1 )
		this.dataCategory.records.push(record_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more recordsIds as a Records
	// from a DataCategory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRecords( dataCategoryId, recordsIds ): Observable<any> {

		// get the DataCategory
		this.loadHelper( dataCategoryId );


	// split on a comma with no spaces
	var idList 					= recordsIds.split(',');
	var records 	= this.dataCategory.records;

	if ( records != null && recordsIds != null ) {

		// iterate over array of records ids
		records.forEach(function (obj) {
			if ( recordsIds.indexOf(obj._id) > -1 ) {
				// remove the Record_
				this.dataCategory.records.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataBreachesIds as a DataBreaches
	// to a DataCategory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataBreaches( dataCategoryId, dataBreachesIds ): Observable<any> {

		// get the DataCategory
		this.loadHelper( dataCategoryId );

	// split on a comma with no spaces
	var idList = dataBreachesIds.split(',')

	// iterate over array of dataBreaches ids
	idList.forEach(function (id) {
		// read the DataBreach
		var dataBreach = new DataBreachService(this.http).getDataBreach(id);
		// add the DataBreach if not already assigned
		if ( this.dataCategory.dataBreaches.indexOf(dataBreach) == -1 )
		this.dataCategory.dataBreaches.push(dataBreach);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataBreachesIds as a DataBreaches
	// from a DataCategory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataBreaches( dataCategoryId, dataBreachesIds ): Observable<any> {

		// get the DataCategory
		this.loadHelper( dataCategoryId );


	// split on a comma with no spaces
	var idList 					= dataBreachesIds.split(',');
	var dataBreaches 	= this.dataCategory.dataBreaches;

	if ( dataBreaches != null && dataBreachesIds != null ) {

		// iterate over array of dataBreaches ids
		dataBreaches.forEach(function (obj) {
			if ( dataBreachesIds.indexOf(obj._id) > -1 ) {
				// remove the DataBreach
				this.dataCategory.dataBreaches.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DataCategory
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DataCategory/update/' + this.dataCategory;

	return  this.http.post(uri_, this.dataCategory );
}

	//********************************************************************
	// loadHelper - internal helper to load a DataCategory
	//********************************************************************	
	loadHelper( id ) {
		this.getDataCategory(id)
			.subscribe((res : DataCategory) => {
				this.dataCategory = res;
			});
	}
}