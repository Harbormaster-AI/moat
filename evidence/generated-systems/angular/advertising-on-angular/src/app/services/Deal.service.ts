import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Deal} from '../models/Deal';
import {PublisherService} from '../services/Publisher.service';
import {InventorySourceService} from '../services/InventorySource.service';
import {PlacementService} from '../services/Placement.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DealService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	deal : Deal;

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
	// add a Deal
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDeal(floorPrice, Publisher, InventorySources, Placements, DealType) : Observable<any> {
		const uri_ = this.apiUrl + '/Deal/create';
		const obj = {
			      		floorPrice: floorPrice,
      		Publisher: Publisher != null && Publisher.length > 0 ? Publisher : null,
      		InventorySources: InventorySources != null && InventorySources.length > 0 ? InventorySources : null,
      		Placements: Placements != null && Placements.length > 0 ? Placements : null,
			DealType: DealType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Deal
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDeal(floorPrice, Publisher, InventorySources, Placements, DealType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Deal/update/' + id;
		const obj = {
				      		floorPrice: floorPrice,
      		Publisher: Publisher != null && Publisher.length > 0 ? Publisher : null,
      		InventorySources: InventorySources != null && InventorySources.length > 0 ? InventorySources : null,
      		Placements: Placements != null && Placements.length > 0 ? Placements : null,
			DealType: DealType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Deal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDeal(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Deal/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Deal
	// returns the results untouched as an Observable Deal
	// Deal model
	// delegates via URI
	//********************************************************************
	getDeal(id) : Observable<Deal> {
		const uri_ = this.apiUrl + '/Deal/load/' + id;

		return this.http.get<Deal>(uri_);
	}
	
	//********************************************************************
	// gets all Deal
	// returns the results untouched as JSON representation of an
	// Observable array of Deal models
	// delegates via URI
	//********************************************************************
	getDeals() : Observable<Deal[]> {
		const uri_ = this.apiUrl + '/Deal/';

		return this
			.http.get<Deal[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Publisher on a Deal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPublisher( dealId, _publisherId ): Observable<any> {

		// get the Deal from storage
		this.loadHelper( dealId );

	// get the Publisher from storage
	var tmp 	= new PublisherService(this.http).getPublisher(_publisherId);

	// assign the Publisher
	this.deal.publisher = tmp;

	// save the Deal
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Publisher on a Deal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPublisher( dealId ): Observable<any> {

		// get the Deal from storage
		this.loadHelper( dealId );

	// assign Publisher to null
	this.deal.publisher = null;

	// save the Deal
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more inventorySourcesIds as a InventorySources
	// to a Deal
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventorySources( dealId, inventorySourcesIds ): Observable<any> {

		// get the Deal
		this.loadHelper( dealId );

	// split on a comma with no spaces
	var idList = inventorySourcesIds.split(',')

	// iterate over array of inventorySources ids
	idList.forEach(function (id) {
		// read the InventorySource
		var inventorySource = new InventorySourceService(this.http).getInventorySource(id);
		// add the InventorySource if not already assigned
		if ( this.deal.inventorySources.indexOf(inventorySource) == -1 )
		this.deal.inventorySources.push(inventorySource);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventorySourcesIds as a InventorySources
	// from a Deal
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventorySources( dealId, inventorySourcesIds ): Observable<any> {

		// get the Deal
		this.loadHelper( dealId );


	// split on a comma with no spaces
	var idList 					= inventorySourcesIds.split(',');
	var inventorySources 	= this.deal.inventorySources;

	if ( inventorySources != null && inventorySourcesIds != null ) {

		// iterate over array of inventorySources ids
		inventorySources.forEach(function (obj) {
			if ( inventorySourcesIds.indexOf(obj._id) > -1 ) {
				// remove the InventorySource
				this.deal.inventorySources.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more placementsIds as a Placements
	// to a Deal
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPlacements( dealId, placementsIds ): Observable<any> {

		// get the Deal
		this.loadHelper( dealId );

	// split on a comma with no spaces
	var idList = placementsIds.split(',')

	// iterate over array of placements ids
	idList.forEach(function (id) {
		// read the Placement
		var placement = new PlacementService(this.http).getPlacement(id);
		// add the Placement if not already assigned
		if ( this.deal.placements.indexOf(placement) == -1 )
		this.deal.placements.push(placement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more placementsIds as a Placements
	// from a Deal
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePlacements( dealId, placementsIds ): Observable<any> {

		// get the Deal
		this.loadHelper( dealId );


	// split on a comma with no spaces
	var idList 					= placementsIds.split(',');
	var placements 	= this.deal.placements;

	if ( placements != null && placementsIds != null ) {

		// iterate over array of placements ids
		placements.forEach(function (obj) {
			if ( placementsIds.indexOf(obj._id) > -1 ) {
				// remove the Placement
				this.deal.placements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Deal
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Deal/update/' + this.deal;

	return  this.http.post(uri_, this.deal );
}

	//********************************************************************
	// loadHelper - internal helper to load a Deal
	//********************************************************************	
	loadHelper( id ) {
		this.getDeal(id)
			.subscribe((res : Deal) => {
				this.deal = res;
			});
	}
}