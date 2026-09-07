import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {LineItem} from '../models/LineItem';
import {CampaignService} from '../services/Campaign.service';
import {PlacementService} from '../services/Placement.service';
import {TargetingProfileService} from '../services/TargetingProfile.service';
import {DealService} from '../services/Deal.service';
import {CreativeAssetService} from '../services/CreativeAsset.service';
import {PerformanceMetricService} from '../services/PerformanceMetric.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LineItemService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	lineItem : LineItem;

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
	// add a LineItem
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLineItem(name, bidAmount, dailyBudget, frequencyCap, Campaign, Placements, TargetingProfile, Deal, Creatives, PerformanceMetrics, Status, PricingModel, BidStrategy, Pacing) : Observable<any> {
		const uri_ = this.apiUrl + '/LineItem/create';
		const obj = {
			      		name: name,
      		bidAmount: bidAmount,
      		dailyBudget: dailyBudget,
      		frequencyCap: frequencyCap,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		Placements: Placements != null && Placements.length > 0 ? Placements : null,
      		TargetingProfile: TargetingProfile != null && TargetingProfile.length > 0 ? TargetingProfile : null,
      		Deal: Deal != null && Deal.length > 0 ? Deal : null,
      		Creatives: Creatives != null && Creatives.length > 0 ? Creatives : null,
      		PerformanceMetrics: PerformanceMetrics != null && PerformanceMetrics.length > 0 ? PerformanceMetrics : null,
      		Status: Status,
      		PricingModel: PricingModel,
      		BidStrategy: BidStrategy,
			Pacing: Pacing
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a LineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLineItem(name, bidAmount, dailyBudget, frequencyCap, Campaign, Placements, TargetingProfile, Deal, Creatives, PerformanceMetrics, Status, PricingModel, BidStrategy, Pacing, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/LineItem/update/' + id;
		const obj = {
				      		name: name,
      		bidAmount: bidAmount,
      		dailyBudget: dailyBudget,
      		frequencyCap: frequencyCap,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		Placements: Placements != null && Placements.length > 0 ? Placements : null,
      		TargetingProfile: TargetingProfile != null && TargetingProfile.length > 0 ? TargetingProfile : null,
      		Deal: Deal != null && Deal.length > 0 ? Deal : null,
      		Creatives: Creatives != null && Creatives.length > 0 ? Creatives : null,
      		PerformanceMetrics: PerformanceMetrics != null && PerformanceMetrics.length > 0 ? PerformanceMetrics : null,
      		Status: Status,
      		PricingModel: PricingModel,
      		BidStrategy: BidStrategy,
			Pacing: Pacing
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a LineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLineItem(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/LineItem/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a LineItem
	// returns the results untouched as an Observable LineItem
	// LineItem model
	// delegates via URI
	//********************************************************************
	getLineItem(id) : Observable<LineItem> {
		const uri_ = this.apiUrl + '/LineItem/load/' + id;

		return this.http.get<LineItem>(uri_);
	}
	
	//********************************************************************
	// gets all LineItem
	// returns the results untouched as JSON representation of an
	// Observable array of LineItem models
	// delegates via URI
	//********************************************************************
	getLineItems() : Observable<LineItem[]> {
		const uri_ = this.apiUrl + '/LineItem/';

		return this
			.http.get<LineItem[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Campaign on a LineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCampaign( lineItemId, _campaignId ): Observable<any> {

		// get the LineItem from storage
		this.loadHelper( lineItemId );

	// get the Campaign from storage
	var tmp 	= new CampaignService(this.http).getCampaign(_campaignId);

	// assign the Campaign
	this.lineItem.campaign = tmp;

	// save the LineItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Campaign on a LineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCampaign( lineItemId ): Observable<any> {

		// get the LineItem from storage
		this.loadHelper( lineItemId );

	// assign Campaign to null
	this.lineItem.campaign = null;

	// save the LineItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a TargetingProfile on a LineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTargetingProfile( lineItemId, _targetingProfileId ): Observable<any> {

		// get the LineItem from storage
		this.loadHelper( lineItemId );

	// get the TargetingProfile from storage
	var tmp 	= new TargetingProfileService(this.http).getTargetingProfile(_targetingProfileId);

	// assign the TargetingProfile
	this.lineItem.targetingProfile = tmp;

	// save the LineItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TargetingProfile on a LineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTargetingProfile( lineItemId ): Observable<any> {

		// get the LineItem from storage
		this.loadHelper( lineItemId );

	// assign TargetingProfile to null
	this.lineItem.targetingProfile = null;

	// save the LineItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Deal on a LineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDeal( lineItemId, _dealId ): Observable<any> {

		// get the LineItem from storage
		this.loadHelper( lineItemId );

	// get the Deal from storage
	var tmp 	= new DealService(this.http).getDeal(_dealId);

	// assign the Deal
	this.lineItem.deal = tmp;

	// save the LineItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Deal on a LineItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDeal( lineItemId ): Observable<any> {

		// get the LineItem from storage
		this.loadHelper( lineItemId );

	// assign Deal to null
	this.lineItem.deal = null;

	// save the LineItem
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more placementsIds as a Placements
	// to a LineItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPlacements( lineItemId, placementsIds ): Observable<any> {

		// get the LineItem
		this.loadHelper( lineItemId );

	// split on a comma with no spaces
	var idList = placementsIds.split(',')

	// iterate over array of placements ids
	idList.forEach(function (id) {
		// read the Placement
		var placement = new PlacementService(this.http).getPlacement(id);
		// add the Placement if not already assigned
		if ( this.lineItem.placements.indexOf(placement) == -1 )
		this.lineItem.placements.push(placement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more placementsIds as a Placements
	// from a LineItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePlacements( lineItemId, placementsIds ): Observable<any> {

		// get the LineItem
		this.loadHelper( lineItemId );


	// split on a comma with no spaces
	var idList 					= placementsIds.split(',');
	var placements 	= this.lineItem.placements;

	if ( placements != null && placementsIds != null ) {

		// iterate over array of placements ids
		placements.forEach(function (obj) {
			if ( placementsIds.indexOf(obj._id) > -1 ) {
				// remove the Placement
				this.lineItem.placements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more creativesIds as a Creatives
	// to a LineItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCreatives( lineItemId, creativesIds ): Observable<any> {

		// get the LineItem
		this.loadHelper( lineItemId );

	// split on a comma with no spaces
	var idList = creativesIds.split(',')

	// iterate over array of creatives ids
	idList.forEach(function (id) {
		// read the CreativeAsset
		var creativeAsset = new CreativeAssetService(this.http).getCreativeAsset(id);
		// add the CreativeAsset if not already assigned
		if ( this.lineItem.creatives.indexOf(creativeAsset) == -1 )
		this.lineItem.creatives.push(creativeAsset);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more creativesIds as a Creatives
	// from a LineItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCreatives( lineItemId, creativesIds ): Observable<any> {

		// get the LineItem
		this.loadHelper( lineItemId );


	// split on a comma with no spaces
	var idList 					= creativesIds.split(',');
	var creatives 	= this.lineItem.creatives;

	if ( creatives != null && creativesIds != null ) {

		// iterate over array of creatives ids
		creatives.forEach(function (obj) {
			if ( creativesIds.indexOf(obj._id) > -1 ) {
				// remove the CreativeAsset
				this.lineItem.creatives.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more performanceMetricsIds as a PerformanceMetrics
	// to a LineItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPerformanceMetrics( lineItemId, performanceMetricsIds ): Observable<any> {

		// get the LineItem
		this.loadHelper( lineItemId );

	// split on a comma with no spaces
	var idList = performanceMetricsIds.split(',')

	// iterate over array of performanceMetrics ids
	idList.forEach(function (id) {
		// read the PerformanceMetric
		var performanceMetric = new PerformanceMetricService(this.http).getPerformanceMetric(id);
		// add the PerformanceMetric if not already assigned
		if ( this.lineItem.performanceMetrics.indexOf(performanceMetric) == -1 )
		this.lineItem.performanceMetrics.push(performanceMetric);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more performanceMetricsIds as a PerformanceMetrics
	// from a LineItem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePerformanceMetrics( lineItemId, performanceMetricsIds ): Observable<any> {

		// get the LineItem
		this.loadHelper( lineItemId );


	// split on a comma with no spaces
	var idList 					= performanceMetricsIds.split(',');
	var performanceMetrics 	= this.lineItem.performanceMetrics;

	if ( performanceMetrics != null && performanceMetricsIds != null ) {

		// iterate over array of performanceMetrics ids
		performanceMetrics.forEach(function (obj) {
			if ( performanceMetricsIds.indexOf(obj._id) > -1 ) {
				// remove the PerformanceMetric
				this.lineItem.performanceMetrics.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a LineItem
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/LineItem/update/' + this.lineItem;

	return  this.http.post(uri_, this.lineItem );
}

	//********************************************************************
	// loadHelper - internal helper to load a LineItem
	//********************************************************************	
	loadHelper( id ) {
		this.getLineItem(id)
			.subscribe((res : LineItem) => {
				this.lineItem = res;
			});
	}
}