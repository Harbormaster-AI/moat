import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AudienceSegment} from '../models/AudienceSegment';
import {DataProviderService} from '../services/DataProvider.service';
import {CampaignService} from '../services/Campaign.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AudienceSegmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	audienceSegment : AudienceSegment;

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
	// add a AudienceSegment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAudienceSegment(name, estimatedReach, description, Provider, Campaigns, ProviderType) : Observable<any> {
		const uri_ = this.apiUrl + '/AudienceSegment/create';
		const obj = {
			      		name: name,
      		estimatedReach: estimatedReach,
      		description: description,
      		Provider: Provider != null && Provider.length > 0 ? Provider : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
			ProviderType: ProviderType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AudienceSegment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAudienceSegment(name, estimatedReach, description, Provider, Campaigns, ProviderType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AudienceSegment/update/' + id;
		const obj = {
				      		name: name,
      		estimatedReach: estimatedReach,
      		description: description,
      		Provider: Provider != null && Provider.length > 0 ? Provider : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
			ProviderType: ProviderType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AudienceSegment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAudienceSegment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AudienceSegment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AudienceSegment
	// returns the results untouched as an Observable AudienceSegment
	// AudienceSegment model
	// delegates via URI
	//********************************************************************
	getAudienceSegment(id) : Observable<AudienceSegment> {
		const uri_ = this.apiUrl + '/AudienceSegment/load/' + id;

		return this.http.get<AudienceSegment>(uri_);
	}
	
	//********************************************************************
	// gets all AudienceSegment
	// returns the results untouched as JSON representation of an
	// Observable array of AudienceSegment models
	// delegates via URI
	//********************************************************************
	getAudienceSegments() : Observable<AudienceSegment[]> {
		const uri_ = this.apiUrl + '/AudienceSegment/';

		return this
			.http.get<AudienceSegment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Provider on a AudienceSegment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProvider( audienceSegmentId, _providerId ): Observable<any> {

		// get the AudienceSegment from storage
		this.loadHelper( audienceSegmentId );

	// get the DataProvider from storage
	var tmp 	= new DataProviderService(this.http).getDataProvider(_providerId);

	// assign the Provider
	this.audienceSegment.provider = tmp;

	// save the AudienceSegment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Provider on a AudienceSegment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProvider( audienceSegmentId ): Observable<any> {

		// get the AudienceSegment from storage
		this.loadHelper( audienceSegmentId );

	// assign Provider to null
	this.audienceSegment.provider = null;

	// save the AudienceSegment
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more campaignsIds as a Campaigns
	// to a AudienceSegment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCampaigns( audienceSegmentId, campaignsIds ): Observable<any> {

		// get the AudienceSegment
		this.loadHelper( audienceSegmentId );

	// split on a comma with no spaces
	var idList = campaignsIds.split(',')

	// iterate over array of campaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.audienceSegment.campaigns.indexOf(campaign) == -1 )
		this.audienceSegment.campaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more campaignsIds as a Campaigns
	// from a AudienceSegment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCampaigns( audienceSegmentId, campaignsIds ): Observable<any> {

		// get the AudienceSegment
		this.loadHelper( audienceSegmentId );


	// split on a comma with no spaces
	var idList 					= campaignsIds.split(',');
	var campaigns 	= this.audienceSegment.campaigns;

	if ( campaigns != null && campaignsIds != null ) {

		// iterate over array of campaigns ids
		campaigns.forEach(function (obj) {
			if ( campaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.audienceSegment.campaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AudienceSegment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AudienceSegment/update/' + this.audienceSegment;

	return  this.http.post(uri_, this.audienceSegment );
}

	//********************************************************************
	// loadHelper - internal helper to load a AudienceSegment
	//********************************************************************	
	loadHelper( id ) {
		this.getAudienceSegment(id)
			.subscribe((res : AudienceSegment) => {
				this.audienceSegment = res;
			});
	}
}