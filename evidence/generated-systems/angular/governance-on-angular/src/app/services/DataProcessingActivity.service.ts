import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DataProcessingActivity} from '../models/DataProcessingActivity';
import {OrganizationService} from '../services/Organization.service';
import {DataCategoryService} from '../services/DataCategory.service';
import {System_Service} from '../services/System_.service';
import {Record_Service} from '../services/Record_.service';
import {PrivacyNoticeService} from '../services/PrivacyNotice.service';
import {ThirdPartyService} from '../services/ThirdParty.service';
import {ConsentService} from '../services/Consent.service';
import {DataBreachService} from '../services/DataBreach.service';
import {DataSubjectRequestService} from '../services/DataSubjectRequest.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DataProcessingActivityService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dataProcessingActivity : DataProcessingActivity;

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
	// add a DataProcessingActivity
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDataProcessingActivity(name, purpose, startDate, Organization, DataCategories, Systems, Records, PrivacyNotices, ThirdParties, Consents, DataBreaches, DataSubjectRequests, LawfulBasis) : Observable<any> {
		const uri_ = this.apiUrl + '/DataProcessingActivity/create';
		const obj = {
			      		name: name,
      		purpose: purpose,
      		startDate: startDate,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		DataCategories: DataCategories != null && DataCategories.length > 0 ? DataCategories : null,
      		Systems: Systems != null && Systems.length > 0 ? Systems : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		PrivacyNotices: PrivacyNotices != null && PrivacyNotices.length > 0 ? PrivacyNotices : null,
      		ThirdParties: ThirdParties != null && ThirdParties.length > 0 ? ThirdParties : null,
      		Consents: Consents != null && Consents.length > 0 ? Consents : null,
      		DataBreaches: DataBreaches != null && DataBreaches.length > 0 ? DataBreaches : null,
      		DataSubjectRequests: DataSubjectRequests != null && DataSubjectRequests.length > 0 ? DataSubjectRequests : null,
			LawfulBasis: LawfulBasis
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DataProcessingActivity
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDataProcessingActivity(name, purpose, startDate, Organization, DataCategories, Systems, Records, PrivacyNotices, ThirdParties, Consents, DataBreaches, DataSubjectRequests, LawfulBasis, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DataProcessingActivity/update/' + id;
		const obj = {
				      		name: name,
      		purpose: purpose,
      		startDate: startDate,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		DataCategories: DataCategories != null && DataCategories.length > 0 ? DataCategories : null,
      		Systems: Systems != null && Systems.length > 0 ? Systems : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		PrivacyNotices: PrivacyNotices != null && PrivacyNotices.length > 0 ? PrivacyNotices : null,
      		ThirdParties: ThirdParties != null && ThirdParties.length > 0 ? ThirdParties : null,
      		Consents: Consents != null && Consents.length > 0 ? Consents : null,
      		DataBreaches: DataBreaches != null && DataBreaches.length > 0 ? DataBreaches : null,
      		DataSubjectRequests: DataSubjectRequests != null && DataSubjectRequests.length > 0 ? DataSubjectRequests : null,
			LawfulBasis: LawfulBasis
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DataProcessingActivity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDataProcessingActivity(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DataProcessingActivity/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DataProcessingActivity
	// returns the results untouched as an Observable DataProcessingActivity
	// DataProcessingActivity model
	// delegates via URI
	//********************************************************************
	getDataProcessingActivity(id) : Observable<DataProcessingActivity> {
		const uri_ = this.apiUrl + '/DataProcessingActivity/load/' + id;

		return this.http.get<DataProcessingActivity>(uri_);
	}
	
	//********************************************************************
	// gets all DataProcessingActivity
	// returns the results untouched as JSON representation of an
	// Observable array of DataProcessingActivity models
	// delegates via URI
	//********************************************************************
	getDataProcessingActivitys() : Observable<DataProcessingActivity[]> {
		const uri_ = this.apiUrl + '/DataProcessingActivity/';

		return this
			.http.get<DataProcessingActivity[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a DataProcessingActivity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( dataProcessingActivityId, _organizationId ): Observable<any> {

		// get the DataProcessingActivity from storage
		this.loadHelper( dataProcessingActivityId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.dataProcessingActivity.organization = tmp;

	// save the DataProcessingActivity
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a DataProcessingActivity
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( dataProcessingActivityId ): Observable<any> {

		// get the DataProcessingActivity from storage
		this.loadHelper( dataProcessingActivityId );

	// assign Organization to null
	this.dataProcessingActivity.organization = null;

	// save the DataProcessingActivity
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more dataCategoriesIds as a DataCategories
	// to a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataCategories( dataProcessingActivityId, dataCategoriesIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );

	// split on a comma with no spaces
	var idList = dataCategoriesIds.split(',')

	// iterate over array of dataCategories ids
	idList.forEach(function (id) {
		// read the DataCategory
		var dataCategory = new DataCategoryService(this.http).getDataCategory(id);
		// add the DataCategory if not already assigned
		if ( this.dataProcessingActivity.dataCategories.indexOf(dataCategory) == -1 )
		this.dataProcessingActivity.dataCategories.push(dataCategory);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataCategoriesIds as a DataCategories
	// from a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataCategories( dataProcessingActivityId, dataCategoriesIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );


	// split on a comma with no spaces
	var idList 					= dataCategoriesIds.split(',');
	var dataCategories 	= this.dataProcessingActivity.dataCategories;

	if ( dataCategories != null && dataCategoriesIds != null ) {

		// iterate over array of dataCategories ids
		dataCategories.forEach(function (obj) {
			if ( dataCategoriesIds.indexOf(obj._id) > -1 ) {
				// remove the DataCategory
				this.dataProcessingActivity.dataCategories.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more systemsIds as a Systems
	// to a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSystems( dataProcessingActivityId, systemsIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );

	// split on a comma with no spaces
	var idList = systemsIds.split(',')

	// iterate over array of systems ids
	idList.forEach(function (id) {
		// read the System_
		var system_ = new System_Service(this.http).getSystem_(id);
		// add the System_ if not already assigned
		if ( this.dataProcessingActivity.systems.indexOf(system_) == -1 )
		this.dataProcessingActivity.systems.push(system_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more systemsIds as a Systems
	// from a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSystems( dataProcessingActivityId, systemsIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );


	// split on a comma with no spaces
	var idList 					= systemsIds.split(',');
	var systems 	= this.dataProcessingActivity.systems;

	if ( systems != null && systemsIds != null ) {

		// iterate over array of systems ids
		systems.forEach(function (obj) {
			if ( systemsIds.indexOf(obj._id) > -1 ) {
				// remove the System_
				this.dataProcessingActivity.systems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more recordsIds as a Records
	// to a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRecords( dataProcessingActivityId, recordsIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );

	// split on a comma with no spaces
	var idList = recordsIds.split(',')

	// iterate over array of records ids
	idList.forEach(function (id) {
		// read the Record_
		var record_ = new Record_Service(this.http).getRecord_(id);
		// add the Record_ if not already assigned
		if ( this.dataProcessingActivity.records.indexOf(record_) == -1 )
		this.dataProcessingActivity.records.push(record_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more recordsIds as a Records
	// from a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRecords( dataProcessingActivityId, recordsIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );


	// split on a comma with no spaces
	var idList 					= recordsIds.split(',');
	var records 	= this.dataProcessingActivity.records;

	if ( records != null && recordsIds != null ) {

		// iterate over array of records ids
		records.forEach(function (obj) {
			if ( recordsIds.indexOf(obj._id) > -1 ) {
				// remove the Record_
				this.dataProcessingActivity.records.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more privacyNoticesIds as a PrivacyNotices
	// to a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPrivacyNotices( dataProcessingActivityId, privacyNoticesIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );

	// split on a comma with no spaces
	var idList = privacyNoticesIds.split(',')

	// iterate over array of privacyNotices ids
	idList.forEach(function (id) {
		// read the PrivacyNotice
		var privacyNotice = new PrivacyNoticeService(this.http).getPrivacyNotice(id);
		// add the PrivacyNotice if not already assigned
		if ( this.dataProcessingActivity.privacyNotices.indexOf(privacyNotice) == -1 )
		this.dataProcessingActivity.privacyNotices.push(privacyNotice);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more privacyNoticesIds as a PrivacyNotices
	// from a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePrivacyNotices( dataProcessingActivityId, privacyNoticesIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );


	// split on a comma with no spaces
	var idList 					= privacyNoticesIds.split(',');
	var privacyNotices 	= this.dataProcessingActivity.privacyNotices;

	if ( privacyNotices != null && privacyNoticesIds != null ) {

		// iterate over array of privacyNotices ids
		privacyNotices.forEach(function (obj) {
			if ( privacyNoticesIds.indexOf(obj._id) > -1 ) {
				// remove the PrivacyNotice
				this.dataProcessingActivity.privacyNotices.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more thirdPartiesIds as a ThirdParties
	// to a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addThirdParties( dataProcessingActivityId, thirdPartiesIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );

	// split on a comma with no spaces
	var idList = thirdPartiesIds.split(',')

	// iterate over array of thirdParties ids
	idList.forEach(function (id) {
		// read the ThirdParty
		var thirdParty = new ThirdPartyService(this.http).getThirdParty(id);
		// add the ThirdParty if not already assigned
		if ( this.dataProcessingActivity.thirdParties.indexOf(thirdParty) == -1 )
		this.dataProcessingActivity.thirdParties.push(thirdParty);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more thirdPartiesIds as a ThirdParties
	// from a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeThirdParties( dataProcessingActivityId, thirdPartiesIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );


	// split on a comma with no spaces
	var idList 					= thirdPartiesIds.split(',');
	var thirdParties 	= this.dataProcessingActivity.thirdParties;

	if ( thirdParties != null && thirdPartiesIds != null ) {

		// iterate over array of thirdParties ids
		thirdParties.forEach(function (obj) {
			if ( thirdPartiesIds.indexOf(obj._id) > -1 ) {
				// remove the ThirdParty
				this.dataProcessingActivity.thirdParties.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more consentsIds as a Consents
	// to a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addConsents( dataProcessingActivityId, consentsIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );

	// split on a comma with no spaces
	var idList = consentsIds.split(',')

	// iterate over array of consents ids
	idList.forEach(function (id) {
		// read the Consent
		var consent = new ConsentService(this.http).getConsent(id);
		// add the Consent if not already assigned
		if ( this.dataProcessingActivity.consents.indexOf(consent) == -1 )
		this.dataProcessingActivity.consents.push(consent);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more consentsIds as a Consents
	// from a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeConsents( dataProcessingActivityId, consentsIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );


	// split on a comma with no spaces
	var idList 					= consentsIds.split(',');
	var consents 	= this.dataProcessingActivity.consents;

	if ( consents != null && consentsIds != null ) {

		// iterate over array of consents ids
		consents.forEach(function (obj) {
			if ( consentsIds.indexOf(obj._id) > -1 ) {
				// remove the Consent
				this.dataProcessingActivity.consents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataBreachesIds as a DataBreaches
	// to a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataBreaches( dataProcessingActivityId, dataBreachesIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );

	// split on a comma with no spaces
	var idList = dataBreachesIds.split(',')

	// iterate over array of dataBreaches ids
	idList.forEach(function (id) {
		// read the DataBreach
		var dataBreach = new DataBreachService(this.http).getDataBreach(id);
		// add the DataBreach if not already assigned
		if ( this.dataProcessingActivity.dataBreaches.indexOf(dataBreach) == -1 )
		this.dataProcessingActivity.dataBreaches.push(dataBreach);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataBreachesIds as a DataBreaches
	// from a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataBreaches( dataProcessingActivityId, dataBreachesIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );


	// split on a comma with no spaces
	var idList 					= dataBreachesIds.split(',');
	var dataBreaches 	= this.dataProcessingActivity.dataBreaches;

	if ( dataBreaches != null && dataBreachesIds != null ) {

		// iterate over array of dataBreaches ids
		dataBreaches.forEach(function (obj) {
			if ( dataBreachesIds.indexOf(obj._id) > -1 ) {
				// remove the DataBreach
				this.dataProcessingActivity.dataBreaches.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataSubjectRequestsIds as a DataSubjectRequests
	// to a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataSubjectRequests( dataProcessingActivityId, dataSubjectRequestsIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );

	// split on a comma with no spaces
	var idList = dataSubjectRequestsIds.split(',')

	// iterate over array of dataSubjectRequests ids
	idList.forEach(function (id) {
		// read the DataSubjectRequest
		var dataSubjectRequest = new DataSubjectRequestService(this.http).getDataSubjectRequest(id);
		// add the DataSubjectRequest if not already assigned
		if ( this.dataProcessingActivity.dataSubjectRequests.indexOf(dataSubjectRequest) == -1 )
		this.dataProcessingActivity.dataSubjectRequests.push(dataSubjectRequest);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataSubjectRequestsIds as a DataSubjectRequests
	// from a DataProcessingActivity
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataSubjectRequests( dataProcessingActivityId, dataSubjectRequestsIds ): Observable<any> {

		// get the DataProcessingActivity
		this.loadHelper( dataProcessingActivityId );


	// split on a comma with no spaces
	var idList 					= dataSubjectRequestsIds.split(',');
	var dataSubjectRequests 	= this.dataProcessingActivity.dataSubjectRequests;

	if ( dataSubjectRequests != null && dataSubjectRequestsIds != null ) {

		// iterate over array of dataSubjectRequests ids
		dataSubjectRequests.forEach(function (obj) {
			if ( dataSubjectRequestsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSubjectRequest
				this.dataProcessingActivity.dataSubjectRequests.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DataProcessingActivity
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DataProcessingActivity/update/' + this.dataProcessingActivity;

	return  this.http.post(uri_, this.dataProcessingActivity );
}

	//********************************************************************
	// loadHelper - internal helper to load a DataProcessingActivity
	//********************************************************************	
	loadHelper( id ) {
		this.getDataProcessingActivity(id)
			.subscribe((res : DataProcessingActivity) => {
				this.dataProcessingActivity = res;
			});
	}
}