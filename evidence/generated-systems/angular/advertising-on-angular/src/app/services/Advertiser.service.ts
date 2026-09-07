import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Advertiser} from '../models/Advertiser';
import {AgencyService} from '../services/Agency.service';
import {AdAccountService} from '../services/AdAccount.service';
import {BillingProfileService} from '../services/BillingProfile.service';
import {CampaignService} from '../services/Campaign.service';
import {TrackingPixelService} from '../services/TrackingPixel.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AdvertiserService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	advertiser : Advertiser;

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
	// add a Advertiser
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAdvertiser(name, legalName, industry, website, Agency, AdAccounts, BillingProfiles, Campaigns, TrackingPixels) : Observable<any> {
		const uri_ = this.apiUrl + '/Advertiser/create';
		const obj = {
			      		name: name,
      		legalName: legalName,
      		industry: industry,
      		website: website,
      		Agency: Agency != null && Agency.length > 0 ? Agency : null,
      		AdAccounts: AdAccounts != null && AdAccounts.length > 0 ? AdAccounts : null,
      		BillingProfiles: BillingProfiles != null && BillingProfiles.length > 0 ? BillingProfiles : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
			TrackingPixels: TrackingPixels != null && TrackingPixels.length > 0 ? TrackingPixels : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Advertiser
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAdvertiser(name, legalName, industry, website, Agency, AdAccounts, BillingProfiles, Campaigns, TrackingPixels, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Advertiser/update/' + id;
		const obj = {
				      		name: name,
      		legalName: legalName,
      		industry: industry,
      		website: website,
      		Agency: Agency != null && Agency.length > 0 ? Agency : null,
      		AdAccounts: AdAccounts != null && AdAccounts.length > 0 ? AdAccounts : null,
      		BillingProfiles: BillingProfiles != null && BillingProfiles.length > 0 ? BillingProfiles : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
			TrackingPixels: TrackingPixels != null && TrackingPixels.length > 0 ? TrackingPixels : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Advertiser
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAdvertiser(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Advertiser/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Advertiser
	// returns the results untouched as an Observable Advertiser
	// Advertiser model
	// delegates via URI
	//********************************************************************
	getAdvertiser(id) : Observable<Advertiser> {
		const uri_ = this.apiUrl + '/Advertiser/load/' + id;

		return this.http.get<Advertiser>(uri_);
	}
	
	//********************************************************************
	// gets all Advertiser
	// returns the results untouched as JSON representation of an
	// Observable array of Advertiser models
	// delegates via URI
	//********************************************************************
	getAdvertisers() : Observable<Advertiser[]> {
		const uri_ = this.apiUrl + '/Advertiser/';

		return this
			.http.get<Advertiser[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Agency on a Advertiser
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAgency( advertiserId, _agencyId ): Observable<any> {

		// get the Advertiser from storage
		this.loadHelper( advertiserId );

	// get the Agency from storage
	var tmp 	= new AgencyService(this.http).getAgency(_agencyId);

	// assign the Agency
	this.advertiser.agency = tmp;

	// save the Advertiser
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Agency on a Advertiser
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAgency( advertiserId ): Observable<any> {

		// get the Advertiser from storage
		this.loadHelper( advertiserId );

	// assign Agency to null
	this.advertiser.agency = null;

	// save the Advertiser
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more adAccountsIds as a AdAccounts
	// to a Advertiser
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAdAccounts( advertiserId, adAccountsIds ): Observable<any> {

		// get the Advertiser
		this.loadHelper( advertiserId );

	// split on a comma with no spaces
	var idList = adAccountsIds.split(',')

	// iterate over array of adAccounts ids
	idList.forEach(function (id) {
		// read the AdAccount
		var adAccount = new AdAccountService(this.http).getAdAccount(id);
		// add the AdAccount if not already assigned
		if ( this.advertiser.adAccounts.indexOf(adAccount) == -1 )
		this.advertiser.adAccounts.push(adAccount);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more adAccountsIds as a AdAccounts
	// from a Advertiser
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAdAccounts( advertiserId, adAccountsIds ): Observable<any> {

		// get the Advertiser
		this.loadHelper( advertiserId );


	// split on a comma with no spaces
	var idList 					= adAccountsIds.split(',');
	var adAccounts 	= this.advertiser.adAccounts;

	if ( adAccounts != null && adAccountsIds != null ) {

		// iterate over array of adAccounts ids
		adAccounts.forEach(function (obj) {
			if ( adAccountsIds.indexOf(obj._id) > -1 ) {
				// remove the AdAccount
				this.advertiser.adAccounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more billingProfilesIds as a BillingProfiles
	// to a Advertiser
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBillingProfiles( advertiserId, billingProfilesIds ): Observable<any> {

		// get the Advertiser
		this.loadHelper( advertiserId );

	// split on a comma with no spaces
	var idList = billingProfilesIds.split(',')

	// iterate over array of billingProfiles ids
	idList.forEach(function (id) {
		// read the BillingProfile
		var billingProfile = new BillingProfileService(this.http).getBillingProfile(id);
		// add the BillingProfile if not already assigned
		if ( this.advertiser.billingProfiles.indexOf(billingProfile) == -1 )
		this.advertiser.billingProfiles.push(billingProfile);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more billingProfilesIds as a BillingProfiles
	// from a Advertiser
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBillingProfiles( advertiserId, billingProfilesIds ): Observable<any> {

		// get the Advertiser
		this.loadHelper( advertiserId );


	// split on a comma with no spaces
	var idList 					= billingProfilesIds.split(',');
	var billingProfiles 	= this.advertiser.billingProfiles;

	if ( billingProfiles != null && billingProfilesIds != null ) {

		// iterate over array of billingProfiles ids
		billingProfiles.forEach(function (obj) {
			if ( billingProfilesIds.indexOf(obj._id) > -1 ) {
				// remove the BillingProfile
				this.advertiser.billingProfiles.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more campaignsIds as a Campaigns
	// to a Advertiser
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCampaigns( advertiserId, campaignsIds ): Observable<any> {

		// get the Advertiser
		this.loadHelper( advertiserId );

	// split on a comma with no spaces
	var idList = campaignsIds.split(',')

	// iterate over array of campaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.advertiser.campaigns.indexOf(campaign) == -1 )
		this.advertiser.campaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more campaignsIds as a Campaigns
	// from a Advertiser
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCampaigns( advertiserId, campaignsIds ): Observable<any> {

		// get the Advertiser
		this.loadHelper( advertiserId );


	// split on a comma with no spaces
	var idList 					= campaignsIds.split(',');
	var campaigns 	= this.advertiser.campaigns;

	if ( campaigns != null && campaignsIds != null ) {

		// iterate over array of campaigns ids
		campaigns.forEach(function (obj) {
			if ( campaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.advertiser.campaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more trackingPixelsIds as a TrackingPixels
	// to a Advertiser
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTrackingPixels( advertiserId, trackingPixelsIds ): Observable<any> {

		// get the Advertiser
		this.loadHelper( advertiserId );

	// split on a comma with no spaces
	var idList = trackingPixelsIds.split(',')

	// iterate over array of trackingPixels ids
	idList.forEach(function (id) {
		// read the TrackingPixel
		var trackingPixel = new TrackingPixelService(this.http).getTrackingPixel(id);
		// add the TrackingPixel if not already assigned
		if ( this.advertiser.trackingPixels.indexOf(trackingPixel) == -1 )
		this.advertiser.trackingPixels.push(trackingPixel);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more trackingPixelsIds as a TrackingPixels
	// from a Advertiser
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTrackingPixels( advertiserId, trackingPixelsIds ): Observable<any> {

		// get the Advertiser
		this.loadHelper( advertiserId );


	// split on a comma with no spaces
	var idList 					= trackingPixelsIds.split(',');
	var trackingPixels 	= this.advertiser.trackingPixels;

	if ( trackingPixels != null && trackingPixelsIds != null ) {

		// iterate over array of trackingPixels ids
		trackingPixels.forEach(function (obj) {
			if ( trackingPixelsIds.indexOf(obj._id) > -1 ) {
				// remove the TrackingPixel
				this.advertiser.trackingPixels.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Advertiser
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Advertiser/update/' + this.advertiser;

	return  this.http.post(uri_, this.advertiser );
}

	//********************************************************************
	// loadHelper - internal helper to load a Advertiser
	//********************************************************************	
	loadHelper( id ) {
		this.getAdvertiser(id)
			.subscribe((res : Advertiser) => {
				this.advertiser = res;
			});
	}
}