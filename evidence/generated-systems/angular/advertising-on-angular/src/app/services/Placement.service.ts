import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Placement} from '../models/Placement';
import {LineItemService} from '../services/LineItem.service';
import {AdSlotService} from '../services/AdSlot.service';
import {DealService} from '../services/Deal.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PlacementService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	placement : Placement;

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
	// add a Placement
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPlacement(name, flight, goalImpressions, LineItem, AdSlot, Deal) : Observable<any> {
		const uri_ = this.apiUrl + '/Placement/create';
		const obj = {
			      		name: name,
      		flight: flight,
      		goalImpressions: goalImpressions,
      		LineItem: LineItem != null && LineItem.length > 0 ? LineItem : null,
      		AdSlot: AdSlot != null && AdSlot.length > 0 ? AdSlot : null,
			Deal: Deal != null && Deal.length > 0 ? Deal : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Placement
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePlacement(name, flight, goalImpressions, LineItem, AdSlot, Deal, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Placement/update/' + id;
		const obj = {
				      		name: name,
      		flight: flight,
      		goalImpressions: goalImpressions,
      		LineItem: LineItem != null && LineItem.length > 0 ? LineItem : null,
      		AdSlot: AdSlot != null && AdSlot.length > 0 ? AdSlot : null,
			Deal: Deal != null && Deal.length > 0 ? Deal : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Placement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePlacement(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Placement/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Placement
	// returns the results untouched as an Observable Placement
	// Placement model
	// delegates via URI
	//********************************************************************
	getPlacement(id) : Observable<Placement> {
		const uri_ = this.apiUrl + '/Placement/load/' + id;

		return this.http.get<Placement>(uri_);
	}
	
	//********************************************************************
	// gets all Placement
	// returns the results untouched as JSON representation of an
	// Observable array of Placement models
	// delegates via URI
	//********************************************************************
	getPlacements() : Observable<Placement[]> {
		const uri_ = this.apiUrl + '/Placement/';

		return this
			.http.get<Placement[]>(uri_);
	}
	
			//********************************************************************
	// assigns a LineItem on a Placement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLineItem( placementId, _lineItemId ): Observable<any> {

		// get the Placement from storage
		this.loadHelper( placementId );

	// get the LineItem from storage
	var tmp 	= new LineItemService(this.http).getLineItem(_lineItemId);

	// assign the LineItem
	this.placement.lineItem = tmp;

	// save the Placement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LineItem on a Placement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLineItem( placementId ): Observable<any> {

		// get the Placement from storage
		this.loadHelper( placementId );

	// assign LineItem to null
	this.placement.lineItem = null;

	// save the Placement
	return this.saveHelper();
}

		//********************************************************************
	// assigns a AdSlot on a Placement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdSlot( placementId, _adSlotId ): Observable<any> {

		// get the Placement from storage
		this.loadHelper( placementId );

	// get the AdSlot from storage
	var tmp 	= new AdSlotService(this.http).getAdSlot(_adSlotId);

	// assign the AdSlot
	this.placement.adSlot = tmp;

	// save the Placement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AdSlot on a Placement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdSlot( placementId ): Observable<any> {

		// get the Placement from storage
		this.loadHelper( placementId );

	// assign AdSlot to null
	this.placement.adSlot = null;

	// save the Placement
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Deal on a Placement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDeal( placementId, _dealId ): Observable<any> {

		// get the Placement from storage
		this.loadHelper( placementId );

	// get the Deal from storage
	var tmp 	= new DealService(this.http).getDeal(_dealId);

	// assign the Deal
	this.placement.deal = tmp;

	// save the Placement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Deal on a Placement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDeal( placementId ): Observable<any> {

		// get the Placement from storage
		this.loadHelper( placementId );

	// assign Deal to null
	this.placement.deal = null;

	// save the Placement
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Placement
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Placement/update/' + this.placement;

	return  this.http.post(uri_, this.placement );
}

	//********************************************************************
	// loadHelper - internal helper to load a Placement
	//********************************************************************	
	loadHelper( id ) {
		this.getPlacement(id)
			.subscribe((res : Placement) => {
				this.placement = res;
			});
	}
}