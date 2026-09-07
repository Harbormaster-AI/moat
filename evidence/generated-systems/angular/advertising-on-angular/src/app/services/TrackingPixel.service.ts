import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TrackingPixel} from '../models/TrackingPixel';
import {CampaignService} from '../services/Campaign.service';
import {AdvertiserService} from '../services/Advertiser.service';
import {ConversionEventService} from '../services/ConversionEvent.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TrackingPixelService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	trackingPixel : TrackingPixel;

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
	// add a TrackingPixel
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTrackingPixel(name, url, Campaign, Advertiser, ConversionEvents, EventType, PixelType) : Observable<any> {
		const uri_ = this.apiUrl + '/TrackingPixel/create';
		const obj = {
			      		name: name,
      		url: url,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		Advertiser: Advertiser != null && Advertiser.length > 0 ? Advertiser : null,
      		ConversionEvents: ConversionEvents != null && ConversionEvents.length > 0 ? ConversionEvents : null,
      		EventType: EventType,
			PixelType: PixelType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TrackingPixel
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTrackingPixel(name, url, Campaign, Advertiser, ConversionEvents, EventType, PixelType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TrackingPixel/update/' + id;
		const obj = {
				      		name: name,
      		url: url,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		Advertiser: Advertiser != null && Advertiser.length > 0 ? Advertiser : null,
      		ConversionEvents: ConversionEvents != null && ConversionEvents.length > 0 ? ConversionEvents : null,
      		EventType: EventType,
			PixelType: PixelType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TrackingPixel
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTrackingPixel(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TrackingPixel/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TrackingPixel
	// returns the results untouched as an Observable TrackingPixel
	// TrackingPixel model
	// delegates via URI
	//********************************************************************
	getTrackingPixel(id) : Observable<TrackingPixel> {
		const uri_ = this.apiUrl + '/TrackingPixel/load/' + id;

		return this.http.get<TrackingPixel>(uri_);
	}
	
	//********************************************************************
	// gets all TrackingPixel
	// returns the results untouched as JSON representation of an
	// Observable array of TrackingPixel models
	// delegates via URI
	//********************************************************************
	getTrackingPixels() : Observable<TrackingPixel[]> {
		const uri_ = this.apiUrl + '/TrackingPixel/';

		return this
			.http.get<TrackingPixel[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Campaign on a TrackingPixel
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCampaign( trackingPixelId, _campaignId ): Observable<any> {

		// get the TrackingPixel from storage
		this.loadHelper( trackingPixelId );

	// get the Campaign from storage
	var tmp 	= new CampaignService(this.http).getCampaign(_campaignId);

	// assign the Campaign
	this.trackingPixel.campaign = tmp;

	// save the TrackingPixel
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Campaign on a TrackingPixel
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCampaign( trackingPixelId ): Observable<any> {

		// get the TrackingPixel from storage
		this.loadHelper( trackingPixelId );

	// assign Campaign to null
	this.trackingPixel.campaign = null;

	// save the TrackingPixel
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Advertiser on a TrackingPixel
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdvertiser( trackingPixelId, _advertiserId ): Observable<any> {

		// get the TrackingPixel from storage
		this.loadHelper( trackingPixelId );

	// get the Advertiser from storage
	var tmp 	= new AdvertiserService(this.http).getAdvertiser(_advertiserId);

	// assign the Advertiser
	this.trackingPixel.advertiser = tmp;

	// save the TrackingPixel
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Advertiser on a TrackingPixel
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdvertiser( trackingPixelId ): Observable<any> {

		// get the TrackingPixel from storage
		this.loadHelper( trackingPixelId );

	// assign Advertiser to null
	this.trackingPixel.advertiser = null;

	// save the TrackingPixel
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more conversionEventsIds as a ConversionEvents
	// to a TrackingPixel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addConversionEvents( trackingPixelId, conversionEventsIds ): Observable<any> {

		// get the TrackingPixel
		this.loadHelper( trackingPixelId );

	// split on a comma with no spaces
	var idList = conversionEventsIds.split(',')

	// iterate over array of conversionEvents ids
	idList.forEach(function (id) {
		// read the ConversionEvent
		var conversionEvent = new ConversionEventService(this.http).getConversionEvent(id);
		// add the ConversionEvent if not already assigned
		if ( this.trackingPixel.conversionEvents.indexOf(conversionEvent) == -1 )
		this.trackingPixel.conversionEvents.push(conversionEvent);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more conversionEventsIds as a ConversionEvents
	// from a TrackingPixel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeConversionEvents( trackingPixelId, conversionEventsIds ): Observable<any> {

		// get the TrackingPixel
		this.loadHelper( trackingPixelId );


	// split on a comma with no spaces
	var idList 					= conversionEventsIds.split(',');
	var conversionEvents 	= this.trackingPixel.conversionEvents;

	if ( conversionEvents != null && conversionEventsIds != null ) {

		// iterate over array of conversionEvents ids
		conversionEvents.forEach(function (obj) {
			if ( conversionEventsIds.indexOf(obj._id) > -1 ) {
				// remove the ConversionEvent
				this.trackingPixel.conversionEvents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a TrackingPixel
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TrackingPixel/update/' + this.trackingPixel;

	return  this.http.post(uri_, this.trackingPixel );
}

	//********************************************************************
	// loadHelper - internal helper to load a TrackingPixel
	//********************************************************************	
	loadHelper( id ) {
		this.getTrackingPixel(id)
			.subscribe((res : TrackingPixel) => {
				this.trackingPixel = res;
			});
	}
}