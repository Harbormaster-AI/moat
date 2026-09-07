import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DataProvider} from '../models/DataProvider';
import {AudienceSegmentService} from '../services/AudienceSegment.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DataProviderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dataProvider : DataProvider;

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
	// add a DataProvider
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDataProvider(name, website, AudienceSegments, ProviderType) : Observable<any> {
		const uri_ = this.apiUrl + '/DataProvider/create';
		const obj = {
			      		name: name,
      		website: website,
      		AudienceSegments: AudienceSegments != null && AudienceSegments.length > 0 ? AudienceSegments : null,
			ProviderType: ProviderType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DataProvider
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDataProvider(name, website, AudienceSegments, ProviderType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DataProvider/update/' + id;
		const obj = {
				      		name: name,
      		website: website,
      		AudienceSegments: AudienceSegments != null && AudienceSegments.length > 0 ? AudienceSegments : null,
			ProviderType: ProviderType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DataProvider
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDataProvider(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DataProvider/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DataProvider
	// returns the results untouched as an Observable DataProvider
	// DataProvider model
	// delegates via URI
	//********************************************************************
	getDataProvider(id) : Observable<DataProvider> {
		const uri_ = this.apiUrl + '/DataProvider/load/' + id;

		return this.http.get<DataProvider>(uri_);
	}
	
	//********************************************************************
	// gets all DataProvider
	// returns the results untouched as JSON representation of an
	// Observable array of DataProvider models
	// delegates via URI
	//********************************************************************
	getDataProviders() : Observable<DataProvider[]> {
		const uri_ = this.apiUrl + '/DataProvider/';

		return this
			.http.get<DataProvider[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more audienceSegmentsIds as a AudienceSegments
	// to a DataProvider
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAudienceSegments( dataProviderId, audienceSegmentsIds ): Observable<any> {

		// get the DataProvider
		this.loadHelper( dataProviderId );

	// split on a comma with no spaces
	var idList = audienceSegmentsIds.split(',')

	// iterate over array of audienceSegments ids
	idList.forEach(function (id) {
		// read the AudienceSegment
		var audienceSegment = new AudienceSegmentService(this.http).getAudienceSegment(id);
		// add the AudienceSegment if not already assigned
		if ( this.dataProvider.audienceSegments.indexOf(audienceSegment) == -1 )
		this.dataProvider.audienceSegments.push(audienceSegment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more audienceSegmentsIds as a AudienceSegments
	// from a DataProvider
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAudienceSegments( dataProviderId, audienceSegmentsIds ): Observable<any> {

		// get the DataProvider
		this.loadHelper( dataProviderId );


	// split on a comma with no spaces
	var idList 					= audienceSegmentsIds.split(',');
	var audienceSegments 	= this.dataProvider.audienceSegments;

	if ( audienceSegments != null && audienceSegmentsIds != null ) {

		// iterate over array of audienceSegments ids
		audienceSegments.forEach(function (obj) {
			if ( audienceSegmentsIds.indexOf(obj._id) > -1 ) {
				// remove the AudienceSegment
				this.dataProvider.audienceSegments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DataProvider
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DataProvider/update/' + this.dataProvider;

	return  this.http.post(uri_, this.dataProvider );
}

	//********************************************************************
	// loadHelper - internal helper to load a DataProvider
	//********************************************************************	
	loadHelper( id ) {
		this.getDataProvider(id)
			.subscribe((res : DataProvider) => {
				this.dataProvider = res;
			});
	}
}