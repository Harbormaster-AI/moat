import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AdSlot} from '../models/AdSlot';
import {InventorySourceService} from '../services/InventorySource.service';
import {PlacementService} from '../services/Placement.service';
import {RateService} from '../services/Rate.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AdSlotService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	adSlot : AdSlot;

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
	// add a AdSlot
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAdSlot(slotCode, width, height, floorPrice, InventorySource, Placements, Rates, Format) : Observable<any> {
		const uri_ = this.apiUrl + '/AdSlot/create';
		const obj = {
			      		slotCode: slotCode,
      		width: width,
      		height: height,
      		floorPrice: floorPrice,
      		InventorySource: InventorySource != null && InventorySource.length > 0 ? InventorySource : null,
      		Placements: Placements != null && Placements.length > 0 ? Placements : null,
      		Rates: Rates != null && Rates.length > 0 ? Rates : null,
			Format: Format
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AdSlot
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAdSlot(slotCode, width, height, floorPrice, InventorySource, Placements, Rates, Format, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AdSlot/update/' + id;
		const obj = {
				      		slotCode: slotCode,
      		width: width,
      		height: height,
      		floorPrice: floorPrice,
      		InventorySource: InventorySource != null && InventorySource.length > 0 ? InventorySource : null,
      		Placements: Placements != null && Placements.length > 0 ? Placements : null,
      		Rates: Rates != null && Rates.length > 0 ? Rates : null,
			Format: Format
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AdSlot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAdSlot(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AdSlot/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AdSlot
	// returns the results untouched as an Observable AdSlot
	// AdSlot model
	// delegates via URI
	//********************************************************************
	getAdSlot(id) : Observable<AdSlot> {
		const uri_ = this.apiUrl + '/AdSlot/load/' + id;

		return this.http.get<AdSlot>(uri_);
	}
	
	//********************************************************************
	// gets all AdSlot
	// returns the results untouched as JSON representation of an
	// Observable array of AdSlot models
	// delegates via URI
	//********************************************************************
	getAdSlots() : Observable<AdSlot[]> {
		const uri_ = this.apiUrl + '/AdSlot/';

		return this
			.http.get<AdSlot[]>(uri_);
	}
	
			//********************************************************************
	// assigns a InventorySource on a AdSlot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInventorySource( adSlotId, _inventorySourceId ): Observable<any> {

		// get the AdSlot from storage
		this.loadHelper( adSlotId );

	// get the InventorySource from storage
	var tmp 	= new InventorySourceService(this.http).getInventorySource(_inventorySourceId);

	// assign the InventorySource
	this.adSlot.inventorySource = tmp;

	// save the AdSlot
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InventorySource on a AdSlot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInventorySource( adSlotId ): Observable<any> {

		// get the AdSlot from storage
		this.loadHelper( adSlotId );

	// assign InventorySource to null
	this.adSlot.inventorySource = null;

	// save the AdSlot
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more placementsIds as a Placements
	// to a AdSlot
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPlacements( adSlotId, placementsIds ): Observable<any> {

		// get the AdSlot
		this.loadHelper( adSlotId );

	// split on a comma with no spaces
	var idList = placementsIds.split(',')

	// iterate over array of placements ids
	idList.forEach(function (id) {
		// read the Placement
		var placement = new PlacementService(this.http).getPlacement(id);
		// add the Placement if not already assigned
		if ( this.adSlot.placements.indexOf(placement) == -1 )
		this.adSlot.placements.push(placement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more placementsIds as a Placements
	// from a AdSlot
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePlacements( adSlotId, placementsIds ): Observable<any> {

		// get the AdSlot
		this.loadHelper( adSlotId );


	// split on a comma with no spaces
	var idList 					= placementsIds.split(',');
	var placements 	= this.adSlot.placements;

	if ( placements != null && placementsIds != null ) {

		// iterate over array of placements ids
		placements.forEach(function (obj) {
			if ( placementsIds.indexOf(obj._id) > -1 ) {
				// remove the Placement
				this.adSlot.placements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ratesIds as a Rates
	// to a AdSlot
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRates( adSlotId, ratesIds ): Observable<any> {

		// get the AdSlot
		this.loadHelper( adSlotId );

	// split on a comma with no spaces
	var idList = ratesIds.split(',')

	// iterate over array of rates ids
	idList.forEach(function (id) {
		// read the Rate
		var rate = new RateService(this.http).getRate(id);
		// add the Rate if not already assigned
		if ( this.adSlot.rates.indexOf(rate) == -1 )
		this.adSlot.rates.push(rate);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ratesIds as a Rates
	// from a AdSlot
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRates( adSlotId, ratesIds ): Observable<any> {

		// get the AdSlot
		this.loadHelper( adSlotId );


	// split on a comma with no spaces
	var idList 					= ratesIds.split(',');
	var rates 	= this.adSlot.rates;

	if ( rates != null && ratesIds != null ) {

		// iterate over array of rates ids
		rates.forEach(function (obj) {
			if ( ratesIds.indexOf(obj._id) > -1 ) {
				// remove the Rate
				this.adSlot.rates.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AdSlot
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AdSlot/update/' + this.adSlot;

	return  this.http.post(uri_, this.adSlot );
}

	//********************************************************************
	// loadHelper - internal helper to load a AdSlot
	//********************************************************************	
	loadHelper( id ) {
		this.getAdSlot(id)
			.subscribe((res : AdSlot) => {
				this.adSlot = res;
			});
	}
}