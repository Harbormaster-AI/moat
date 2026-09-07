import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InsertionOrder} from '../models/InsertionOrder';
import {AdvertiserService} from '../services/Advertiser.service';
import {AgencyService} from '../services/Agency.service';
import {PublisherService} from '../services/Publisher.service';
import {CampaignService} from '../services/Campaign.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InsertionOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	insertionOrder : InsertionOrder;

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
	// add a InsertionOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInsertionOrder(ioNumber, agreedBudget, flight, Advertiser, Agency, Publisher, Campaigns, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/InsertionOrder/create';
		const obj = {
			      		ioNumber: ioNumber,
      		agreedBudget: agreedBudget,
      		flight: flight,
      		Advertiser: Advertiser != null && Advertiser.length > 0 ? Advertiser : null,
      		Agency: Agency != null && Agency.length > 0 ? Agency : null,
      		Publisher: Publisher != null && Publisher.length > 0 ? Publisher : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InsertionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInsertionOrder(ioNumber, agreedBudget, flight, Advertiser, Agency, Publisher, Campaigns, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InsertionOrder/update/' + id;
		const obj = {
				      		ioNumber: ioNumber,
      		agreedBudget: agreedBudget,
      		flight: flight,
      		Advertiser: Advertiser != null && Advertiser.length > 0 ? Advertiser : null,
      		Agency: Agency != null && Agency.length > 0 ? Agency : null,
      		Publisher: Publisher != null && Publisher.length > 0 ? Publisher : null,
      		Campaigns: Campaigns != null && Campaigns.length > 0 ? Campaigns : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InsertionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInsertionOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InsertionOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InsertionOrder
	// returns the results untouched as an Observable InsertionOrder
	// InsertionOrder model
	// delegates via URI
	//********************************************************************
	getInsertionOrder(id) : Observable<InsertionOrder> {
		const uri_ = this.apiUrl + '/InsertionOrder/load/' + id;

		return this.http.get<InsertionOrder>(uri_);
	}
	
	//********************************************************************
	// gets all InsertionOrder
	// returns the results untouched as JSON representation of an
	// Observable array of InsertionOrder models
	// delegates via URI
	//********************************************************************
	getInsertionOrders() : Observable<InsertionOrder[]> {
		const uri_ = this.apiUrl + '/InsertionOrder/';

		return this
			.http.get<InsertionOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Advertiser on a InsertionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdvertiser( insertionOrderId, _advertiserId ): Observable<any> {

		// get the InsertionOrder from storage
		this.loadHelper( insertionOrderId );

	// get the Advertiser from storage
	var tmp 	= new AdvertiserService(this.http).getAdvertiser(_advertiserId);

	// assign the Advertiser
	this.insertionOrder.advertiser = tmp;

	// save the InsertionOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Advertiser on a InsertionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdvertiser( insertionOrderId ): Observable<any> {

		// get the InsertionOrder from storage
		this.loadHelper( insertionOrderId );

	// assign Advertiser to null
	this.insertionOrder.advertiser = null;

	// save the InsertionOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Agency on a InsertionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAgency( insertionOrderId, _agencyId ): Observable<any> {

		// get the InsertionOrder from storage
		this.loadHelper( insertionOrderId );

	// get the Agency from storage
	var tmp 	= new AgencyService(this.http).getAgency(_agencyId);

	// assign the Agency
	this.insertionOrder.agency = tmp;

	// save the InsertionOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Agency on a InsertionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAgency( insertionOrderId ): Observable<any> {

		// get the InsertionOrder from storage
		this.loadHelper( insertionOrderId );

	// assign Agency to null
	this.insertionOrder.agency = null;

	// save the InsertionOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Publisher on a InsertionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPublisher( insertionOrderId, _publisherId ): Observable<any> {

		// get the InsertionOrder from storage
		this.loadHelper( insertionOrderId );

	// get the Publisher from storage
	var tmp 	= new PublisherService(this.http).getPublisher(_publisherId);

	// assign the Publisher
	this.insertionOrder.publisher = tmp;

	// save the InsertionOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Publisher on a InsertionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPublisher( insertionOrderId ): Observable<any> {

		// get the InsertionOrder from storage
		this.loadHelper( insertionOrderId );

	// assign Publisher to null
	this.insertionOrder.publisher = null;

	// save the InsertionOrder
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more campaignsIds as a Campaigns
	// to a InsertionOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCampaigns( insertionOrderId, campaignsIds ): Observable<any> {

		// get the InsertionOrder
		this.loadHelper( insertionOrderId );

	// split on a comma with no spaces
	var idList = campaignsIds.split(',')

	// iterate over array of campaigns ids
	idList.forEach(function (id) {
		// read the Campaign
		var campaign = new CampaignService(this.http).getCampaign(id);
		// add the Campaign if not already assigned
		if ( this.insertionOrder.campaigns.indexOf(campaign) == -1 )
		this.insertionOrder.campaigns.push(campaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more campaignsIds as a Campaigns
	// from a InsertionOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCampaigns( insertionOrderId, campaignsIds ): Observable<any> {

		// get the InsertionOrder
		this.loadHelper( insertionOrderId );


	// split on a comma with no spaces
	var idList 					= campaignsIds.split(',');
	var campaigns 	= this.insertionOrder.campaigns;

	if ( campaigns != null && campaignsIds != null ) {

		// iterate over array of campaigns ids
		campaigns.forEach(function (obj) {
			if ( campaignsIds.indexOf(obj._id) > -1 ) {
				// remove the Campaign
				this.insertionOrder.campaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InsertionOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InsertionOrder/update/' + this.insertionOrder;

	return  this.http.post(uri_, this.insertionOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a InsertionOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getInsertionOrder(id)
			.subscribe((res : InsertionOrder) => {
				this.insertionOrder = res;
			});
	}
}