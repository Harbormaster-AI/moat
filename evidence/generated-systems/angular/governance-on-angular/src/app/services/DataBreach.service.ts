import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DataBreach} from '../models/DataBreach';
import {OrganizationService} from '../services/Organization.service';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {DataCategoryService} from '../services/DataCategory.service';
import {ThirdPartyService} from '../services/ThirdParty.service';
import {MatterService} from '../services/Matter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DataBreachService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dataBreach : DataBreach;

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
	// add a DataBreach
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDataBreach(incidentDate, description, recordsAffected, notificationRequired, Organization, ProcessingActivities, DataCategories, ThirdParties, Matter, Severity, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/DataBreach/create';
		const obj = {
			      		incidentDate: incidentDate,
      		description: description,
      		recordsAffected: recordsAffected,
      		notificationRequired: notificationRequired,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		DataCategories: DataCategories != null && DataCategories.length > 0 ? DataCategories : null,
      		ThirdParties: ThirdParties != null && ThirdParties.length > 0 ? ThirdParties : null,
      		Matter: Matter != null && Matter.length > 0 ? Matter : null,
      		Severity: Severity,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DataBreach
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDataBreach(incidentDate, description, recordsAffected, notificationRequired, Organization, ProcessingActivities, DataCategories, ThirdParties, Matter, Severity, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DataBreach/update/' + id;
		const obj = {
				      		incidentDate: incidentDate,
      		description: description,
      		recordsAffected: recordsAffected,
      		notificationRequired: notificationRequired,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		DataCategories: DataCategories != null && DataCategories.length > 0 ? DataCategories : null,
      		ThirdParties: ThirdParties != null && ThirdParties.length > 0 ? ThirdParties : null,
      		Matter: Matter != null && Matter.length > 0 ? Matter : null,
      		Severity: Severity,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DataBreach
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDataBreach(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DataBreach/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DataBreach
	// returns the results untouched as an Observable DataBreach
	// DataBreach model
	// delegates via URI
	//********************************************************************
	getDataBreach(id) : Observable<DataBreach> {
		const uri_ = this.apiUrl + '/DataBreach/load/' + id;

		return this.http.get<DataBreach>(uri_);
	}
	
	//********************************************************************
	// gets all DataBreach
	// returns the results untouched as JSON representation of an
	// Observable array of DataBreach models
	// delegates via URI
	//********************************************************************
	getDataBreachs() : Observable<DataBreach[]> {
		const uri_ = this.apiUrl + '/DataBreach/';

		return this
			.http.get<DataBreach[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a DataBreach
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( dataBreachId, _organizationId ): Observable<any> {

		// get the DataBreach from storage
		this.loadHelper( dataBreachId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.dataBreach.organization = tmp;

	// save the DataBreach
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a DataBreach
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( dataBreachId ): Observable<any> {

		// get the DataBreach from storage
		this.loadHelper( dataBreachId );

	// assign Organization to null
	this.dataBreach.organization = null;

	// save the DataBreach
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Matter on a DataBreach
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMatter( dataBreachId, _matterId ): Observable<any> {

		// get the DataBreach from storage
		this.loadHelper( dataBreachId );

	// get the Matter from storage
	var tmp 	= new MatterService(this.http).getMatter(_matterId);

	// assign the Matter
	this.dataBreach.matter = tmp;

	// save the DataBreach
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Matter on a DataBreach
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMatter( dataBreachId ): Observable<any> {

		// get the DataBreach from storage
		this.loadHelper( dataBreachId );

	// assign Matter to null
	this.dataBreach.matter = null;

	// save the DataBreach
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more processingActivitiesIds as a ProcessingActivities
	// to a DataBreach
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcessingActivities( dataBreachId, processingActivitiesIds ): Observable<any> {

		// get the DataBreach
		this.loadHelper( dataBreachId );

	// split on a comma with no spaces
	var idList = processingActivitiesIds.split(',')

	// iterate over array of processingActivities ids
	idList.forEach(function (id) {
		// read the DataProcessingActivity
		var dataProcessingActivity = new DataProcessingActivityService(this.http).getDataProcessingActivity(id);
		// add the DataProcessingActivity if not already assigned
		if ( this.dataBreach.processingActivities.indexOf(dataProcessingActivity) == -1 )
		this.dataBreach.processingActivities.push(dataProcessingActivity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more processingActivitiesIds as a ProcessingActivities
	// from a DataBreach
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcessingActivities( dataBreachId, processingActivitiesIds ): Observable<any> {

		// get the DataBreach
		this.loadHelper( dataBreachId );


	// split on a comma with no spaces
	var idList 					= processingActivitiesIds.split(',');
	var processingActivities 	= this.dataBreach.processingActivities;

	if ( processingActivities != null && processingActivitiesIds != null ) {

		// iterate over array of processingActivities ids
		processingActivities.forEach(function (obj) {
			if ( processingActivitiesIds.indexOf(obj._id) > -1 ) {
				// remove the DataProcessingActivity
				this.dataBreach.processingActivities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataCategoriesIds as a DataCategories
	// to a DataBreach
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataCategories( dataBreachId, dataCategoriesIds ): Observable<any> {

		// get the DataBreach
		this.loadHelper( dataBreachId );

	// split on a comma with no spaces
	var idList = dataCategoriesIds.split(',')

	// iterate over array of dataCategories ids
	idList.forEach(function (id) {
		// read the DataCategory
		var dataCategory = new DataCategoryService(this.http).getDataCategory(id);
		// add the DataCategory if not already assigned
		if ( this.dataBreach.dataCategories.indexOf(dataCategory) == -1 )
		this.dataBreach.dataCategories.push(dataCategory);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataCategoriesIds as a DataCategories
	// from a DataBreach
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataCategories( dataBreachId, dataCategoriesIds ): Observable<any> {

		// get the DataBreach
		this.loadHelper( dataBreachId );


	// split on a comma with no spaces
	var idList 					= dataCategoriesIds.split(',');
	var dataCategories 	= this.dataBreach.dataCategories;

	if ( dataCategories != null && dataCategoriesIds != null ) {

		// iterate over array of dataCategories ids
		dataCategories.forEach(function (obj) {
			if ( dataCategoriesIds.indexOf(obj._id) > -1 ) {
				// remove the DataCategory
				this.dataBreach.dataCategories.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more thirdPartiesIds as a ThirdParties
	// to a DataBreach
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addThirdParties( dataBreachId, thirdPartiesIds ): Observable<any> {

		// get the DataBreach
		this.loadHelper( dataBreachId );

	// split on a comma with no spaces
	var idList = thirdPartiesIds.split(',')

	// iterate over array of thirdParties ids
	idList.forEach(function (id) {
		// read the ThirdParty
		var thirdParty = new ThirdPartyService(this.http).getThirdParty(id);
		// add the ThirdParty if not already assigned
		if ( this.dataBreach.thirdParties.indexOf(thirdParty) == -1 )
		this.dataBreach.thirdParties.push(thirdParty);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more thirdPartiesIds as a ThirdParties
	// from a DataBreach
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeThirdParties( dataBreachId, thirdPartiesIds ): Observable<any> {

		// get the DataBreach
		this.loadHelper( dataBreachId );


	// split on a comma with no spaces
	var idList 					= thirdPartiesIds.split(',');
	var thirdParties 	= this.dataBreach.thirdParties;

	if ( thirdParties != null && thirdPartiesIds != null ) {

		// iterate over array of thirdParties ids
		thirdParties.forEach(function (obj) {
			if ( thirdPartiesIds.indexOf(obj._id) > -1 ) {
				// remove the ThirdParty
				this.dataBreach.thirdParties.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DataBreach
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DataBreach/update/' + this.dataBreach;

	return  this.http.post(uri_, this.dataBreach );
}

	//********************************************************************
	// loadHelper - internal helper to load a DataBreach
	//********************************************************************	
	loadHelper( id ) {
		this.getDataBreach(id)
			.subscribe((res : DataBreach) => {
				this.dataBreach = res;
			});
	}
}