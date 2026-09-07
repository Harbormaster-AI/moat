import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Campaign} from '../models/Campaign';
import {AdAccountService} from '../services/AdAccount.service';
import {LineItemService} from '../services/LineItem.service';
import {KPIService} from '../services/KPI.service';
import {TrackingPixelService} from '../services/TrackingPixel.service';
import {AudienceSegmentService} from '../services/AudienceSegment.service';
import {ReportService} from '../services/Report.service';
import {InsertionOrderService} from '../services/InsertionOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CampaignService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	campaign : Campaign;

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
	// add a Campaign
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCampaign(name, totalBudget, flight, AdAccount, LineItems, Kpis, TrackingPixels, Audiences, Reports, InsertionOrder, Objective, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Campaign/create';
		const obj = {
			      		name: name,
      		totalBudget: totalBudget,
      		flight: flight,
      		AdAccount: AdAccount != null && AdAccount.length > 0 ? AdAccount : null,
      		LineItems: LineItems != null && LineItems.length > 0 ? LineItems : null,
      		Kpis: Kpis != null && Kpis.length > 0 ? Kpis : null,
      		TrackingPixels: TrackingPixels != null && TrackingPixels.length > 0 ? TrackingPixels : null,
      		Audiences: Audiences != null && Audiences.length > 0 ? Audiences : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		InsertionOrder: InsertionOrder != null && InsertionOrder.length > 0 ? InsertionOrder : null,
      		Objective: Objective,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCampaign(name, totalBudget, flight, AdAccount, LineItems, Kpis, TrackingPixels, Audiences, Reports, InsertionOrder, Objective, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Campaign/update/' + id;
		const obj = {
				      		name: name,
      		totalBudget: totalBudget,
      		flight: flight,
      		AdAccount: AdAccount != null && AdAccount.length > 0 ? AdAccount : null,
      		LineItems: LineItems != null && LineItems.length > 0 ? LineItems : null,
      		Kpis: Kpis != null && Kpis.length > 0 ? Kpis : null,
      		TrackingPixels: TrackingPixels != null && TrackingPixels.length > 0 ? TrackingPixels : null,
      		Audiences: Audiences != null && Audiences.length > 0 ? Audiences : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		InsertionOrder: InsertionOrder != null && InsertionOrder.length > 0 ? InsertionOrder : null,
      		Objective: Objective,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCampaign(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Campaign/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Campaign
	// returns the results untouched as an Observable Campaign
	// Campaign model
	// delegates via URI
	//********************************************************************
	getCampaign(id) : Observable<Campaign> {
		const uri_ = this.apiUrl + '/Campaign/load/' + id;

		return this.http.get<Campaign>(uri_);
	}
	
	//********************************************************************
	// gets all Campaign
	// returns the results untouched as JSON representation of an
	// Observable array of Campaign models
	// delegates via URI
	//********************************************************************
	getCampaigns() : Observable<Campaign[]> {
		const uri_ = this.apiUrl + '/Campaign/';

		return this
			.http.get<Campaign[]>(uri_);
	}
	
			//********************************************************************
	// assigns a AdAccount on a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdAccount( campaignId, _adAccountId ): Observable<any> {

		// get the Campaign from storage
		this.loadHelper( campaignId );

	// get the AdAccount from storage
	var tmp 	= new AdAccountService(this.http).getAdAccount(_adAccountId);

	// assign the AdAccount
	this.campaign.adAccount = tmp;

	// save the Campaign
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AdAccount on a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdAccount( campaignId ): Observable<any> {

		// get the Campaign from storage
		this.loadHelper( campaignId );

	// assign AdAccount to null
	this.campaign.adAccount = null;

	// save the Campaign
	return this.saveHelper();
}

		//********************************************************************
	// assigns a InsertionOrder on a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInsertionOrder( campaignId, _insertionOrderId ): Observable<any> {

		// get the Campaign from storage
		this.loadHelper( campaignId );

	// get the InsertionOrder from storage
	var tmp 	= new InsertionOrderService(this.http).getInsertionOrder(_insertionOrderId);

	// assign the InsertionOrder
	this.campaign.insertionOrder = tmp;

	// save the Campaign
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InsertionOrder on a Campaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInsertionOrder( campaignId ): Observable<any> {

		// get the Campaign from storage
		this.loadHelper( campaignId );

	// assign InsertionOrder to null
	this.campaign.insertionOrder = null;

	// save the Campaign
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more lineItemsIds as a LineItems
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLineItems( campaignId, lineItemsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = lineItemsIds.split(',')

	// iterate over array of lineItems ids
	idList.forEach(function (id) {
		// read the LineItem
		var lineItem = new LineItemService(this.http).getLineItem(id);
		// add the LineItem if not already assigned
		if ( this.campaign.lineItems.indexOf(lineItem) == -1 )
		this.campaign.lineItems.push(lineItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more lineItemsIds as a LineItems
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLineItems( campaignId, lineItemsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= lineItemsIds.split(',');
	var lineItems 	= this.campaign.lineItems;

	if ( lineItems != null && lineItemsIds != null ) {

		// iterate over array of lineItems ids
		lineItems.forEach(function (obj) {
			if ( lineItemsIds.indexOf(obj._id) > -1 ) {
				// remove the LineItem
				this.campaign.lineItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more kpisIds as a Kpis
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addKpis( campaignId, kpisIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = kpisIds.split(',')

	// iterate over array of kpis ids
	idList.forEach(function (id) {
		// read the KPI
		var kPI = new KPIService(this.http).getKPI(id);
		// add the KPI if not already assigned
		if ( this.campaign.kpis.indexOf(kPI) == -1 )
		this.campaign.kpis.push(kPI);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more kpisIds as a Kpis
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeKpis( campaignId, kpisIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= kpisIds.split(',');
	var kpis 	= this.campaign.kpis;

	if ( kpis != null && kpisIds != null ) {

		// iterate over array of kpis ids
		kpis.forEach(function (obj) {
			if ( kpisIds.indexOf(obj._id) > -1 ) {
				// remove the KPI
				this.campaign.kpis.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more trackingPixelsIds as a TrackingPixels
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTrackingPixels( campaignId, trackingPixelsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = trackingPixelsIds.split(',')

	// iterate over array of trackingPixels ids
	idList.forEach(function (id) {
		// read the TrackingPixel
		var trackingPixel = new TrackingPixelService(this.http).getTrackingPixel(id);
		// add the TrackingPixel if not already assigned
		if ( this.campaign.trackingPixels.indexOf(trackingPixel) == -1 )
		this.campaign.trackingPixels.push(trackingPixel);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more trackingPixelsIds as a TrackingPixels
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTrackingPixels( campaignId, trackingPixelsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= trackingPixelsIds.split(',');
	var trackingPixels 	= this.campaign.trackingPixels;

	if ( trackingPixels != null && trackingPixelsIds != null ) {

		// iterate over array of trackingPixels ids
		trackingPixels.forEach(function (obj) {
			if ( trackingPixelsIds.indexOf(obj._id) > -1 ) {
				// remove the TrackingPixel
				this.campaign.trackingPixels.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more audiencesIds as a Audiences
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAudiences( campaignId, audiencesIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = audiencesIds.split(',')

	// iterate over array of audiences ids
	idList.forEach(function (id) {
		// read the AudienceSegment
		var audienceSegment = new AudienceSegmentService(this.http).getAudienceSegment(id);
		// add the AudienceSegment if not already assigned
		if ( this.campaign.audiences.indexOf(audienceSegment) == -1 )
		this.campaign.audiences.push(audienceSegment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more audiencesIds as a Audiences
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAudiences( campaignId, audiencesIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= audiencesIds.split(',');
	var audiences 	= this.campaign.audiences;

	if ( audiences != null && audiencesIds != null ) {

		// iterate over array of audiences ids
		audiences.forEach(function (obj) {
			if ( audiencesIds.indexOf(obj._id) > -1 ) {
				// remove the AudienceSegment
				this.campaign.audiences.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reportsIds as a Reports
	// to a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReports( campaignId, reportsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );

	// split on a comma with no spaces
	var idList = reportsIds.split(',')

	// iterate over array of reports ids
	idList.forEach(function (id) {
		// read the Report
		var report = new ReportService(this.http).getReport(id);
		// add the Report if not already assigned
		if ( this.campaign.reports.indexOf(report) == -1 )
		this.campaign.reports.push(report);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reportsIds as a Reports
	// from a Campaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReports( campaignId, reportsIds ): Observable<any> {

		// get the Campaign
		this.loadHelper( campaignId );


	// split on a comma with no spaces
	var idList 					= reportsIds.split(',');
	var reports 	= this.campaign.reports;

	if ( reports != null && reportsIds != null ) {

		// iterate over array of reports ids
		reports.forEach(function (obj) {
			if ( reportsIds.indexOf(obj._id) > -1 ) {
				// remove the Report
				this.campaign.reports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Campaign
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Campaign/update/' + this.campaign;

	return  this.http.post(uri_, this.campaign );
}

	//********************************************************************
	// loadHelper - internal helper to load a Campaign
	//********************************************************************	
	loadHelper( id ) {
		this.getCampaign(id)
			.subscribe((res : Campaign) => {
				this.campaign = res;
			});
	}
}