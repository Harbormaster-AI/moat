import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InventorySource} from '../models/InventorySource';
import {PublisherService} from '../services/Publisher.service';
import {AdSlotService} from '../services/AdSlot.service';
import {DealService} from '../services/Deal.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InventorySourceService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inventorySource : InventorySource;

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
	// add a InventorySource
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInventorySource(name, domain, Publisher, AdSlots, Deals, Channel, PrimaryFormat) : Observable<any> {
		const uri_ = this.apiUrl + '/InventorySource/create';
		const obj = {
			      		name: name,
      		domain: domain,
      		Publisher: Publisher != null && Publisher.length > 0 ? Publisher : null,
      		AdSlots: AdSlots != null && AdSlots.length > 0 ? AdSlots : null,
      		Deals: Deals != null && Deals.length > 0 ? Deals : null,
      		Channel: Channel,
			PrimaryFormat: PrimaryFormat
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InventorySource
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInventorySource(name, domain, Publisher, AdSlots, Deals, Channel, PrimaryFormat, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InventorySource/update/' + id;
		const obj = {
				      		name: name,
      		domain: domain,
      		Publisher: Publisher != null && Publisher.length > 0 ? Publisher : null,
      		AdSlots: AdSlots != null && AdSlots.length > 0 ? AdSlots : null,
      		Deals: Deals != null && Deals.length > 0 ? Deals : null,
      		Channel: Channel,
			PrimaryFormat: PrimaryFormat
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InventorySource
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInventorySource(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InventorySource/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InventorySource
	// returns the results untouched as an Observable InventorySource
	// InventorySource model
	// delegates via URI
	//********************************************************************
	getInventorySource(id) : Observable<InventorySource> {
		const uri_ = this.apiUrl + '/InventorySource/load/' + id;

		return this.http.get<InventorySource>(uri_);
	}
	
	//********************************************************************
	// gets all InventorySource
	// returns the results untouched as JSON representation of an
	// Observable array of InventorySource models
	// delegates via URI
	//********************************************************************
	getInventorySources() : Observable<InventorySource[]> {
		const uri_ = this.apiUrl + '/InventorySource/';

		return this
			.http.get<InventorySource[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Publisher on a InventorySource
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPublisher( inventorySourceId, _publisherId ): Observable<any> {

		// get the InventorySource from storage
		this.loadHelper( inventorySourceId );

	// get the Publisher from storage
	var tmp 	= new PublisherService(this.http).getPublisher(_publisherId);

	// assign the Publisher
	this.inventorySource.publisher = tmp;

	// save the InventorySource
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Publisher on a InventorySource
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPublisher( inventorySourceId ): Observable<any> {

		// get the InventorySource from storage
		this.loadHelper( inventorySourceId );

	// assign Publisher to null
	this.inventorySource.publisher = null;

	// save the InventorySource
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more adSlotsIds as a AdSlots
	// to a InventorySource
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAdSlots( inventorySourceId, adSlotsIds ): Observable<any> {

		// get the InventorySource
		this.loadHelper( inventorySourceId );

	// split on a comma with no spaces
	var idList = adSlotsIds.split(',')

	// iterate over array of adSlots ids
	idList.forEach(function (id) {
		// read the AdSlot
		var adSlot = new AdSlotService(this.http).getAdSlot(id);
		// add the AdSlot if not already assigned
		if ( this.inventorySource.adSlots.indexOf(adSlot) == -1 )
		this.inventorySource.adSlots.push(adSlot);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more adSlotsIds as a AdSlots
	// from a InventorySource
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAdSlots( inventorySourceId, adSlotsIds ): Observable<any> {

		// get the InventorySource
		this.loadHelper( inventorySourceId );


	// split on a comma with no spaces
	var idList 					= adSlotsIds.split(',');
	var adSlots 	= this.inventorySource.adSlots;

	if ( adSlots != null && adSlotsIds != null ) {

		// iterate over array of adSlots ids
		adSlots.forEach(function (obj) {
			if ( adSlotsIds.indexOf(obj._id) > -1 ) {
				// remove the AdSlot
				this.inventorySource.adSlots.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dealsIds as a Deals
	// to a InventorySource
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDeals( inventorySourceId, dealsIds ): Observable<any> {

		// get the InventorySource
		this.loadHelper( inventorySourceId );

	// split on a comma with no spaces
	var idList = dealsIds.split(',')

	// iterate over array of deals ids
	idList.forEach(function (id) {
		// read the Deal
		var deal = new DealService(this.http).getDeal(id);
		// add the Deal if not already assigned
		if ( this.inventorySource.deals.indexOf(deal) == -1 )
		this.inventorySource.deals.push(deal);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dealsIds as a Deals
	// from a InventorySource
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDeals( inventorySourceId, dealsIds ): Observable<any> {

		// get the InventorySource
		this.loadHelper( inventorySourceId );


	// split on a comma with no spaces
	var idList 					= dealsIds.split(',');
	var deals 	= this.inventorySource.deals;

	if ( deals != null && dealsIds != null ) {

		// iterate over array of deals ids
		deals.forEach(function (obj) {
			if ( dealsIds.indexOf(obj._id) > -1 ) {
				// remove the Deal
				this.inventorySource.deals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InventorySource
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InventorySource/update/' + this.inventorySource;

	return  this.http.post(uri_, this.inventorySource );
}

	//********************************************************************
	// loadHelper - internal helper to load a InventorySource
	//********************************************************************	
	loadHelper( id ) {
		this.getInventorySource(id)
			.subscribe((res : InventorySource) => {
				this.inventorySource = res;
			});
	}
}