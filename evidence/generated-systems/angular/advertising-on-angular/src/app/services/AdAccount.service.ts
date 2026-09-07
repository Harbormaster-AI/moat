import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AdAccount} from '../models/AdAccount';
import {AdvertiserService} from '../services/Advertiser.service';
import {UserService} from '../services/User.service';
import {CampaignService} from '../services/Campaign.service';
import {BillingProfileService} from '../services/BillingProfile.service';
import {DSPService} from '../services/DSP.service';
import {PerformanceMetricService} from '../services/PerformanceMetric.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AdAccountService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	adAccount : AdAccount;

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
	// add a AdAccount
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAdAccount(name, accountCode, defaultCurrency, defaultTimezone, Advertiser, Users, Campaigns, BillingProfile, Dsp, PerformanceMetrics) : Observable<any> {
		const uri_ = this.apiUrl + '/AdAccount/create';
		const obj = {
			      		name: name,
      		accountCode: accountCode,
      		defaultCurrency: defaultCurrency,
      		defaultTimezone: defaultTimezone,
      		Advertiser: Advertiser != null && Advertiser.length > 0 ? Advertiser : null,
      		Users: Users != null && Users.length > 0 ? Users : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
      		BillingProfile: BillingProfile != null && BillingProfile.length > 0 ? BillingProfile : null,
      		Dsp: Dsp != null && Dsp.length > 0 ? Dsp : null,
			PerformanceMetrics: PerformanceMetrics != null && PerformanceMetrics.length > 0 ? PerformanceMetrics : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AdAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAdAccount(name, accountCode, defaultCurrency, defaultTimezone, Advertiser, Users, Campaigns, BillingProfile, Dsp, PerformanceMetrics, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AdAccount/update/' + id;
		const obj = {
				      		name: name,
      		accountCode: accountCode,
      		defaultCurrency: defaultCurrency,
      		defaultTimezone: defaultTimezone,
      		Advertiser: Advertiser != null && Advertiser.length > 0 ? Advertiser : null,
      		Users: Users != null && Users.length > 0 ? Users : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
      		BillingProfile: BillingProfile != null && BillingProfile.length > 0 ? BillingProfile : null,
      		Dsp: Dsp != null && Dsp.length > 0 ? Dsp : null,
			PerformanceMetrics: PerformanceMetrics != null && PerformanceMetrics.length > 0 ? PerformanceMetrics : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AdAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAdAccount(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AdAccount/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AdAccount
	// returns the results untouched as an Observable AdAccount
	// AdAccount model
	// delegates via URI
	//********************************************************************
	getAdAccount(id) : Observable<AdAccount> {
		const uri_ = this.apiUrl + '/AdAccount/load/' + id;

		return this.http.get<AdAccount>(uri_);
	}
	
	//********************************************************************
	// gets all AdAccount
	// returns the results untouched as JSON representation of an
	// Observable array of AdAccount models
	// delegates via URI
	//********************************************************************
	getAdAccounts() : Observable<AdAccount[]> {
		const uri_ = this.apiUrl + '/AdAccount/';

		return this
			.http.get<AdAccount[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Advertiser on a AdAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdvertiser( adAccountId, _advertiserId ): Observable<any> {

		// get the AdAccount from storage
		this.loadHelper( adAccountId );

	// get the Advertiser from storage
	var tmp 	= new AdvertiserService(this.http).getAdvertiser(_advertiserId);

	// assign the Advertiser
	this.adAccount.advertiser = tmp;

	// save the AdAccount
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Advertiser on a AdAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdvertiser( adAccountId ): Observable<any> {

		// get the AdAccount from storage
		this.loadHelper( adAccountId );

	// assign Advertiser to null
	this.adAccount.advertiser = null;

	// save the AdAccount
	return this.saveHelper();
}

		//********************************************************************
	// assigns a BillingProfile on a AdAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBillingProfile( adAccountId, _billingProfileId ): Observable<any> {

		// get the AdAccount from storage
		this.loadHelper( adAccountId );

	// get the BillingProfile from storage
	var tmp 	= new BillingProfileService(this.http).getBillingProfile(_billingProfileId);

	// assign the BillingProfile
	this.adAccount.billingProfile = tmp;

	// save the AdAccount
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a BillingProfile on a AdAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBillingProfile( adAccountId ): Observable<any> {

		// get the AdAccount from storage
		this.loadHelper( adAccountId );

	// assign BillingProfile to null
	this.adAccount.billingProfile = null;

	// save the AdAccount
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Dsp on a AdAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDsp( adAccountId, _dspId ): Observable<any> {

		// get the AdAccount from storage
		this.loadHelper( adAccountId );

	// get the DSP from storage
	var tmp 	= new DSPService(this.http).getDSP(_dspId);

	// assign the Dsp
	this.adAccount.dsp = tmp;

	// save the AdAccount
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dsp on a AdAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDsp( adAccountId ): Observable<any> {

		// get the AdAccount from storage
		this.loadHelper( adAccountId );

	// assign Dsp to null
	this.adAccount.dsp = null;

	// save the AdAccount
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more usersIds as a Users
	// to a AdAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addUsers( adAccountId, usersIds ): Observable<any> {

		// get the AdAccount
		this.loadHelper( adAccountId );

	// split on a comma with no spaces
	var idList = usersIds.split(',')

	// iterate over array of users ids
	idList.forEach(function (id) {
		// read the User
		var user = new UserService(this.http).getUser(id);
		// add the User if not already assigned
		if ( this.adAccount.users.indexOf(user) == -1 )
		this.adAccount.users.push(user);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more usersIds as a Users
	// from a AdAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeUsers( adAccountId, usersIds ): Observable<any> {

		// get the AdAccount
		this.loadHelper( adAccountId );


	// split on a comma with no spaces
	var idList 					= usersIds.split(',');
	var users 	= this.adAccount.users;

	if ( users != null && usersIds != null ) {

		// iterate over array of users ids
		users.forEach(function (obj) {
			if ( usersIds.indexOf(obj._id) > -1 ) {
				// remove the User
				this.adAccount.users.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more campaignsIds as a Campaigns
	// to a AdAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCampaigns( adAccountId, campaignsIds ): Observable<any> {

		// get the AdAccount
		this.loadHelper( adAccountId );

	// split on a comma with no spaces
	var idList = campaignsIds.split(',')

	// iterate over array of campaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.adAccount.campaigns.indexOf(campaign) == -1 )
		this.adAccount.campaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more campaignsIds as a Campaigns
	// from a AdAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCampaigns( adAccountId, campaignsIds ): Observable<any> {

		// get the AdAccount
		this.loadHelper( adAccountId );


	// split on a comma with no spaces
	var idList 					= campaignsIds.split(',');
	var campaigns 	= this.adAccount.campaigns;

	if ( campaigns != null && campaignsIds != null ) {

		// iterate over array of campaigns ids
		campaigns.forEach(function (obj) {
			if ( campaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.adAccount.campaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more performanceMetricsIds as a PerformanceMetrics
	// to a AdAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPerformanceMetrics( adAccountId, performanceMetricsIds ): Observable<any> {

		// get the AdAccount
		this.loadHelper( adAccountId );

	// split on a comma with no spaces
	var idList = performanceMetricsIds.split(',')

	// iterate over array of performanceMetrics ids
	idList.forEach(function (id) {
		// read the PerformanceMetric
		var performanceMetric = new PerformanceMetricService(this.http).getPerformanceMetric(id);
		// add the PerformanceMetric if not already assigned
		if ( this.adAccount.performanceMetrics.indexOf(performanceMetric) == -1 )
		this.adAccount.performanceMetrics.push(performanceMetric);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more performanceMetricsIds as a PerformanceMetrics
	// from a AdAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePerformanceMetrics( adAccountId, performanceMetricsIds ): Observable<any> {

		// get the AdAccount
		this.loadHelper( adAccountId );


	// split on a comma with no spaces
	var idList 					= performanceMetricsIds.split(',');
	var performanceMetrics 	= this.adAccount.performanceMetrics;

	if ( performanceMetrics != null && performanceMetricsIds != null ) {

		// iterate over array of performanceMetrics ids
		performanceMetrics.forEach(function (obj) {
			if ( performanceMetricsIds.indexOf(obj._id) > -1 ) {
				// remove the PerformanceMetric
				this.adAccount.performanceMetrics.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AdAccount
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AdAccount/update/' + this.adAccount;

	return  this.http.post(uri_, this.adAccount );
}

	//********************************************************************
	// loadHelper - internal helper to load a AdAccount
	//********************************************************************	
	loadHelper( id ) {
		this.getAdAccount(id)
			.subscribe((res : AdAccount) => {
				this.adAccount = res;
			});
	}
}