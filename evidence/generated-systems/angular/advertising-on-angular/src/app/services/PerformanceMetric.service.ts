import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PerformanceMetric} from '../models/PerformanceMetric';
import {AdAccountService} from '../services/AdAccount.service';
import {CampaignService} from '../services/Campaign.service';
import {LineItemService} from '../services/LineItem.service';
import {PlacementService} from '../services/Placement.service';
import {CreativeAssetService} from '../services/CreativeAsset.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PerformanceMetricService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	performanceMetric : PerformanceMetric;

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
	// add a PerformanceMetric
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPerformanceMetric(date, value, AdAccount, Campaign, LineItem, Placement, CreativeAsset, MetricType) : Observable<any> {
		const uri_ = this.apiUrl + '/PerformanceMetric/create';
		const obj = {
			      		date: date,
      		value: value,
      		AdAccount: AdAccount != null && AdAccount.length > 0 ? AdAccount : null,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		LineItem: LineItem != null && LineItem.length > 0 ? LineItem : null,
      		Placement: Placement != null && Placement.length > 0 ? Placement : null,
      		CreativeAsset: CreativeAsset != null && CreativeAsset.length > 0 ? CreativeAsset : null,
			MetricType: MetricType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePerformanceMetric(date, value, AdAccount, Campaign, LineItem, Placement, CreativeAsset, MetricType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PerformanceMetric/update/' + id;
		const obj = {
				      		date: date,
      		value: value,
      		AdAccount: AdAccount != null && AdAccount.length > 0 ? AdAccount : null,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		LineItem: LineItem != null && LineItem.length > 0 ? LineItem : null,
      		Placement: Placement != null && Placement.length > 0 ? Placement : null,
      		CreativeAsset: CreativeAsset != null && CreativeAsset.length > 0 ? CreativeAsset : null,
			MetricType: MetricType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePerformanceMetric(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PerformanceMetric/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PerformanceMetric
	// returns the results untouched as an Observable PerformanceMetric
	// PerformanceMetric model
	// delegates via URI
	//********************************************************************
	getPerformanceMetric(id) : Observable<PerformanceMetric> {
		const uri_ = this.apiUrl + '/PerformanceMetric/load/' + id;

		return this.http.get<PerformanceMetric>(uri_);
	}
	
	//********************************************************************
	// gets all PerformanceMetric
	// returns the results untouched as JSON representation of an
	// Observable array of PerformanceMetric models
	// delegates via URI
	//********************************************************************
	getPerformanceMetrics() : Observable<PerformanceMetric[]> {
		const uri_ = this.apiUrl + '/PerformanceMetric/';

		return this
			.http.get<PerformanceMetric[]>(uri_);
	}
	
			//********************************************************************
	// assigns a AdAccount on a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdAccount( performanceMetricId, _adAccountId ): Observable<any> {

		// get the PerformanceMetric from storage
		this.loadHelper( performanceMetricId );

	// get the AdAccount from storage
	var tmp 	= new AdAccountService(this.http).getAdAccount(_adAccountId);

	// assign the AdAccount
	this.performanceMetric.adAccount = tmp;

	// save the PerformanceMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AdAccount on a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdAccount( performanceMetricId ): Observable<any> {

		// get the PerformanceMetric from storage
		this.loadHelper( performanceMetricId );

	// assign AdAccount to null
	this.performanceMetric.adAccount = null;

	// save the PerformanceMetric
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Campaign on a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCampaign( performanceMetricId, _campaignId ): Observable<any> {

		// get the PerformanceMetric from storage
		this.loadHelper( performanceMetricId );

	// get the Campaign from storage
	var tmp 	= new CampaignService(this.http).getCampaign(_campaignId);

	// assign the Campaign
	this.performanceMetric.campaign = tmp;

	// save the PerformanceMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Campaign on a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCampaign( performanceMetricId ): Observable<any> {

		// get the PerformanceMetric from storage
		this.loadHelper( performanceMetricId );

	// assign Campaign to null
	this.performanceMetric.campaign = null;

	// save the PerformanceMetric
	return this.saveHelper();
}

		//********************************************************************
	// assigns a LineItem on a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLineItem( performanceMetricId, _lineItemId ): Observable<any> {

		// get the PerformanceMetric from storage
		this.loadHelper( performanceMetricId );

	// get the LineItem from storage
	var tmp 	= new LineItemService(this.http).getLineItem(_lineItemId);

	// assign the LineItem
	this.performanceMetric.lineItem = tmp;

	// save the PerformanceMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LineItem on a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLineItem( performanceMetricId ): Observable<any> {

		// get the PerformanceMetric from storage
		this.loadHelper( performanceMetricId );

	// assign LineItem to null
	this.performanceMetric.lineItem = null;

	// save the PerformanceMetric
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Placement on a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlacement( performanceMetricId, _placementId ): Observable<any> {

		// get the PerformanceMetric from storage
		this.loadHelper( performanceMetricId );

	// get the Placement from storage
	var tmp 	= new PlacementService(this.http).getPlacement(_placementId);

	// assign the Placement
	this.performanceMetric.placement = tmp;

	// save the PerformanceMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Placement on a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlacement( performanceMetricId ): Observable<any> {

		// get the PerformanceMetric from storage
		this.loadHelper( performanceMetricId );

	// assign Placement to null
	this.performanceMetric.placement = null;

	// save the PerformanceMetric
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CreativeAsset on a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCreativeAsset( performanceMetricId, _creativeAssetId ): Observable<any> {

		// get the PerformanceMetric from storage
		this.loadHelper( performanceMetricId );

	// get the CreativeAsset from storage
	var tmp 	= new CreativeAssetService(this.http).getCreativeAsset(_creativeAssetId);

	// assign the CreativeAsset
	this.performanceMetric.creativeAsset = tmp;

	// save the PerformanceMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CreativeAsset on a PerformanceMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCreativeAsset( performanceMetricId ): Observable<any> {

		// get the PerformanceMetric from storage
		this.loadHelper( performanceMetricId );

	// assign CreativeAsset to null
	this.performanceMetric.creativeAsset = null;

	// save the PerformanceMetric
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a PerformanceMetric
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PerformanceMetric/update/' + this.performanceMetric;

	return  this.http.post(uri_, this.performanceMetric );
}

	//********************************************************************
	// loadHelper - internal helper to load a PerformanceMetric
	//********************************************************************	
	loadHelper( id ) {
		this.getPerformanceMetric(id)
			.subscribe((res : PerformanceMetric) => {
				this.performanceMetric = res;
			});
	}
}