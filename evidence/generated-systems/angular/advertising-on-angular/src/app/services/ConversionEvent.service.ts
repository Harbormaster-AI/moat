import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ConversionEvent} from '../models/ConversionEvent';
import {CampaignService} from '../services/Campaign.service';
import {LineItemService} from '../services/LineItem.service';
import {TrackingPixelService} from '../services/TrackingPixel.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ConversionEventService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	conversionEvent : ConversionEvent;

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
	// add a ConversionEvent
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addConversionEvent(timestamp, value, Campaign, LineItem, TrackingPixel, EventType, AttributionModel) : Observable<any> {
		const uri_ = this.apiUrl + '/ConversionEvent/create';
		const obj = {
			      		timestamp: timestamp,
      		value: value,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		LineItem: LineItem != null && LineItem.length > 0 ? LineItem : null,
      		TrackingPixel: TrackingPixel != null && TrackingPixel.length > 0 ? TrackingPixel : null,
      		EventType: EventType,
			AttributionModel: AttributionModel
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ConversionEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateConversionEvent(timestamp, value, Campaign, LineItem, TrackingPixel, EventType, AttributionModel, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ConversionEvent/update/' + id;
		const obj = {
				      		timestamp: timestamp,
      		value: value,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		LineItem: LineItem != null && LineItem.length > 0 ? LineItem : null,
      		TrackingPixel: TrackingPixel != null && TrackingPixel.length > 0 ? TrackingPixel : null,
      		EventType: EventType,
			AttributionModel: AttributionModel
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ConversionEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteConversionEvent(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ConversionEvent/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ConversionEvent
	// returns the results untouched as an Observable ConversionEvent
	// ConversionEvent model
	// delegates via URI
	//********************************************************************
	getConversionEvent(id) : Observable<ConversionEvent> {
		const uri_ = this.apiUrl + '/ConversionEvent/load/' + id;

		return this.http.get<ConversionEvent>(uri_);
	}
	
	//********************************************************************
	// gets all ConversionEvent
	// returns the results untouched as JSON representation of an
	// Observable array of ConversionEvent models
	// delegates via URI
	//********************************************************************
	getConversionEvents() : Observable<ConversionEvent[]> {
		const uri_ = this.apiUrl + '/ConversionEvent/';

		return this
			.http.get<ConversionEvent[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Campaign on a ConversionEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCampaign( conversionEventId, _campaignId ): Observable<any> {

		// get the ConversionEvent from storage
		this.loadHelper( conversionEventId );

	// get the Campaign from storage
	var tmp 	= new CampaignService(this.http).getCampaign(_campaignId);

	// assign the Campaign
	this.conversionEvent.campaign = tmp;

	// save the ConversionEvent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Campaign on a ConversionEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCampaign( conversionEventId ): Observable<any> {

		// get the ConversionEvent from storage
		this.loadHelper( conversionEventId );

	// assign Campaign to null
	this.conversionEvent.campaign = null;

	// save the ConversionEvent
	return this.saveHelper();
}

		//********************************************************************
	// assigns a LineItem on a ConversionEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLineItem( conversionEventId, _lineItemId ): Observable<any> {

		// get the ConversionEvent from storage
		this.loadHelper( conversionEventId );

	// get the LineItem from storage
	var tmp 	= new LineItemService(this.http).getLineItem(_lineItemId);

	// assign the LineItem
	this.conversionEvent.lineItem = tmp;

	// save the ConversionEvent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LineItem on a ConversionEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLineItem( conversionEventId ): Observable<any> {

		// get the ConversionEvent from storage
		this.loadHelper( conversionEventId );

	// assign LineItem to null
	this.conversionEvent.lineItem = null;

	// save the ConversionEvent
	return this.saveHelper();
}

		//********************************************************************
	// assigns a TrackingPixel on a ConversionEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTrackingPixel( conversionEventId, _trackingPixelId ): Observable<any> {

		// get the ConversionEvent from storage
		this.loadHelper( conversionEventId );

	// get the TrackingPixel from storage
	var tmp 	= new TrackingPixelService(this.http).getTrackingPixel(_trackingPixelId);

	// assign the TrackingPixel
	this.conversionEvent.trackingPixel = tmp;

	// save the ConversionEvent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TrackingPixel on a ConversionEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTrackingPixel( conversionEventId ): Observable<any> {

		// get the ConversionEvent from storage
		this.loadHelper( conversionEventId );

	// assign TrackingPixel to null
	this.conversionEvent.trackingPixel = null;

	// save the ConversionEvent
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ConversionEvent
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ConversionEvent/update/' + this.conversionEvent;

	return  this.http.post(uri_, this.conversionEvent );
}

	//********************************************************************
	// loadHelper - internal helper to load a ConversionEvent
	//********************************************************************	
	loadHelper( id ) {
		this.getConversionEvent(id)
			.subscribe((res : ConversionEvent) => {
				this.conversionEvent = res;
			});
	}
}