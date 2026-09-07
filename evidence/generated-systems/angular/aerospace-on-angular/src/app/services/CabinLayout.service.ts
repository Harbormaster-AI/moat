import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CabinLayout} from '../models/CabinLayout';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import {AircraftService} from '../services/Aircraft.service';
import {AircraftOptionService} from '../services/AircraftOption.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CabinLayoutService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	cabinLayout : CabinLayout;

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
	// add a CabinLayout
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCabinLayout(layoutCode, totalSeats, classLayout, Variant, Aircraft, Options) : Observable<any> {
		const uri_ = this.apiUrl + '/CabinLayout/create';
		const obj = {
			      		layoutCode: layoutCode,
      		totalSeats: totalSeats,
      		classLayout: classLayout,
      		Variant: Variant != null && Variant.length > 0 ? Variant : null,
      		Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null,
			Options: Options != null && Options.length > 0 ? Options : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CabinLayout
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCabinLayout(layoutCode, totalSeats, classLayout, Variant, Aircraft, Options, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CabinLayout/update/' + id;
		const obj = {
				      		layoutCode: layoutCode,
      		totalSeats: totalSeats,
      		classLayout: classLayout,
      		Variant: Variant != null && Variant.length > 0 ? Variant : null,
      		Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null,
			Options: Options != null && Options.length > 0 ? Options : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CabinLayout
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCabinLayout(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CabinLayout/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CabinLayout
	// returns the results untouched as an Observable CabinLayout
	// CabinLayout model
	// delegates via URI
	//********************************************************************
	getCabinLayout(id) : Observable<CabinLayout> {
		const uri_ = this.apiUrl + '/CabinLayout/load/' + id;

		return this.http.get<CabinLayout>(uri_);
	}
	
	//********************************************************************
	// gets all CabinLayout
	// returns the results untouched as JSON representation of an
	// Observable array of CabinLayout models
	// delegates via URI
	//********************************************************************
	getCabinLayouts() : Observable<CabinLayout[]> {
		const uri_ = this.apiUrl + '/CabinLayout/';

		return this
			.http.get<CabinLayout[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Variant on a CabinLayout
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignVariant( cabinLayoutId, _variantId ): Observable<any> {

		// get the CabinLayout from storage
		this.loadHelper( cabinLayoutId );

	// get the AircraftVariant from storage
	var tmp 	= new AircraftVariantService(this.http).getAircraftVariant(_variantId);

	// assign the Variant
	this.cabinLayout.variant = tmp;

	// save the CabinLayout
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Variant on a CabinLayout
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignVariant( cabinLayoutId ): Observable<any> {

		// get the CabinLayout from storage
		this.loadHelper( cabinLayoutId );

	// assign Variant to null
	this.cabinLayout.variant = null;

	// save the CabinLayout
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more aircraftIds as a Aircraft
	// to a CabinLayout
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAircraft( cabinLayoutId, aircraftIds ): Observable<any> {

		// get the CabinLayout
		this.loadHelper( cabinLayoutId );

	// split on a comma with no spaces
	var idList = aircraftIds.split(',')

	// iterate over array of aircraft ids
	idList.forEach(function (id) {
		// read the Aircraft
		var aircraft = new AircraftService(this.http).getAircraft(id);
		// add the Aircraft if not already assigned
		if ( this.cabinLayout.aircraft.indexOf(aircraft) == -1 )
		this.cabinLayout.aircraft.push(aircraft);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more aircraftIds as a Aircraft
	// from a CabinLayout
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAircraft( cabinLayoutId, aircraftIds ): Observable<any> {

		// get the CabinLayout
		this.loadHelper( cabinLayoutId );


	// split on a comma with no spaces
	var idList 					= aircraftIds.split(',');
	var aircraft 	= this.cabinLayout.aircraft;

	if ( aircraft != null && aircraftIds != null ) {

		// iterate over array of aircraft ids
		aircraft.forEach(function (obj) {
			if ( aircraftIds.indexOf(obj._id) > -1 ) {
				// remove the Aircraft
				this.cabinLayout.aircraft.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more optionsIds as a Options
	// to a CabinLayout
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOptions( cabinLayoutId, optionsIds ): Observable<any> {

		// get the CabinLayout
		this.loadHelper( cabinLayoutId );

	// split on a comma with no spaces
	var idList = optionsIds.split(',')

	// iterate over array of options ids
	idList.forEach(function (id) {
		// read the AircraftOption
		var aircraftOption = new AircraftOptionService(this.http).getAircraftOption(id);
		// add the AircraftOption if not already assigned
		if ( this.cabinLayout.options.indexOf(aircraftOption) == -1 )
		this.cabinLayout.options.push(aircraftOption);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more optionsIds as a Options
	// from a CabinLayout
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOptions( cabinLayoutId, optionsIds ): Observable<any> {

		// get the CabinLayout
		this.loadHelper( cabinLayoutId );


	// split on a comma with no spaces
	var idList 					= optionsIds.split(',');
	var options 	= this.cabinLayout.options;

	if ( options != null && optionsIds != null ) {

		// iterate over array of options ids
		options.forEach(function (obj) {
			if ( optionsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftOption
				this.cabinLayout.options.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a CabinLayout
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CabinLayout/update/' + this.cabinLayout;

	return  this.http.post(uri_, this.cabinLayout );
}

	//********************************************************************
	// loadHelper - internal helper to load a CabinLayout
	//********************************************************************	
	loadHelper( id ) {
		this.getCabinLayout(id)
			.subscribe((res : CabinLayout) => {
				this.cabinLayout = res;
			});
	}
}