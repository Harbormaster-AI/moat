import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ConnectedAircraft} from '../models/ConnectedAircraft';
import {AircraftService} from '../services/Aircraft.service';
import {FlightHealthEventService} from '../services/FlightHealthEvent.service';
import {SoftwareLoadService} from '../services/SoftwareLoad.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ConnectedAircraftService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	connectedAircraft : ConnectedAircraft;

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
	// add a ConnectedAircraft
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addConnectedAircraft(communicationsProvider, Aircraft, FlightHealthEvents, SoftwareLoads, ConnectivityStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/ConnectedAircraft/create';
		const obj = {
			      		communicationsProvider: communicationsProvider,
      		Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null,
      		FlightHealthEvents: FlightHealthEvents != null && FlightHealthEvents.length > 0 ? FlightHealthEvents : null,
      		SoftwareLoads: SoftwareLoads != null && SoftwareLoads.length > 0 ? SoftwareLoads : null,
			ConnectivityStatus: ConnectivityStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ConnectedAircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateConnectedAircraft(communicationsProvider, Aircraft, FlightHealthEvents, SoftwareLoads, ConnectivityStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ConnectedAircraft/update/' + id;
		const obj = {
				      		communicationsProvider: communicationsProvider,
      		Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null,
      		FlightHealthEvents: FlightHealthEvents != null && FlightHealthEvents.length > 0 ? FlightHealthEvents : null,
      		SoftwareLoads: SoftwareLoads != null && SoftwareLoads.length > 0 ? SoftwareLoads : null,
			ConnectivityStatus: ConnectivityStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ConnectedAircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteConnectedAircraft(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ConnectedAircraft/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ConnectedAircraft
	// returns the results untouched as an Observable ConnectedAircraft
	// ConnectedAircraft model
	// delegates via URI
	//********************************************************************
	getConnectedAircraft(id) : Observable<ConnectedAircraft> {
		const uri_ = this.apiUrl + '/ConnectedAircraft/load/' + id;

		return this.http.get<ConnectedAircraft>(uri_);
	}
	
	//********************************************************************
	// gets all ConnectedAircraft
	// returns the results untouched as JSON representation of an
	// Observable array of ConnectedAircraft models
	// delegates via URI
	//********************************************************************
	getConnectedAircrafts() : Observable<ConnectedAircraft[]> {
		const uri_ = this.apiUrl + '/ConnectedAircraft/';

		return this
			.http.get<ConnectedAircraft[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Aircraft on a ConnectedAircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAircraft( connectedAircraftId, _aircraftId ): Observable<any> {

		// get the ConnectedAircraft from storage
		this.loadHelper( connectedAircraftId );

	// get the Aircraft from storage
	var tmp 	= new AircraftService(this.http).getAircraft(_aircraftId);

	// assign the Aircraft
	this.connectedAircraft.aircraft = tmp;

	// save the ConnectedAircraft
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Aircraft on a ConnectedAircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAircraft( connectedAircraftId ): Observable<any> {

		// get the ConnectedAircraft from storage
		this.loadHelper( connectedAircraftId );

	// assign Aircraft to null
	this.connectedAircraft.aircraft = null;

	// save the ConnectedAircraft
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more flightHealthEventsIds as a FlightHealthEvents
	// to a ConnectedAircraft
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFlightHealthEvents( connectedAircraftId, flightHealthEventsIds ): Observable<any> {

		// get the ConnectedAircraft
		this.loadHelper( connectedAircraftId );

	// split on a comma with no spaces
	var idList = flightHealthEventsIds.split(',')

	// iterate over array of flightHealthEvents ids
	idList.forEach(function (id) {
		// read the FlightHealthEvent
		var flightHealthEvent = new FlightHealthEventService(this.http).getFlightHealthEvent(id);
		// add the FlightHealthEvent if not already assigned
		if ( this.connectedAircraft.flightHealthEvents.indexOf(flightHealthEvent) == -1 )
		this.connectedAircraft.flightHealthEvents.push(flightHealthEvent);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more flightHealthEventsIds as a FlightHealthEvents
	// from a ConnectedAircraft
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFlightHealthEvents( connectedAircraftId, flightHealthEventsIds ): Observable<any> {

		// get the ConnectedAircraft
		this.loadHelper( connectedAircraftId );


	// split on a comma with no spaces
	var idList 					= flightHealthEventsIds.split(',');
	var flightHealthEvents 	= this.connectedAircraft.flightHealthEvents;

	if ( flightHealthEvents != null && flightHealthEventsIds != null ) {

		// iterate over array of flightHealthEvents ids
		flightHealthEvents.forEach(function (obj) {
			if ( flightHealthEventsIds.indexOf(obj._id) > -1 ) {
				// remove the FlightHealthEvent
				this.connectedAircraft.flightHealthEvents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more softwareLoadsIds as a SoftwareLoads
	// to a ConnectedAircraft
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSoftwareLoads( connectedAircraftId, softwareLoadsIds ): Observable<any> {

		// get the ConnectedAircraft
		this.loadHelper( connectedAircraftId );

	// split on a comma with no spaces
	var idList = softwareLoadsIds.split(',')

	// iterate over array of softwareLoads ids
	idList.forEach(function (id) {
		// read the SoftwareLoad
		var softwareLoad = new SoftwareLoadService(this.http).getSoftwareLoad(id);
		// add the SoftwareLoad if not already assigned
		if ( this.connectedAircraft.softwareLoads.indexOf(softwareLoad) == -1 )
		this.connectedAircraft.softwareLoads.push(softwareLoad);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more softwareLoadsIds as a SoftwareLoads
	// from a ConnectedAircraft
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSoftwareLoads( connectedAircraftId, softwareLoadsIds ): Observable<any> {

		// get the ConnectedAircraft
		this.loadHelper( connectedAircraftId );


	// split on a comma with no spaces
	var idList 					= softwareLoadsIds.split(',');
	var softwareLoads 	= this.connectedAircraft.softwareLoads;

	if ( softwareLoads != null && softwareLoadsIds != null ) {

		// iterate over array of softwareLoads ids
		softwareLoads.forEach(function (obj) {
			if ( softwareLoadsIds.indexOf(obj._id) > -1 ) {
				// remove the SoftwareLoad
				this.connectedAircraft.softwareLoads.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ConnectedAircraft
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ConnectedAircraft/update/' + this.connectedAircraft;

	return  this.http.post(uri_, this.connectedAircraft );
}

	//********************************************************************
	// loadHelper - internal helper to load a ConnectedAircraft
	//********************************************************************	
	loadHelper( id ) {
		this.getConnectedAircraft(id)
			.subscribe((res : ConnectedAircraft) => {
				this.connectedAircraft = res;
			});
	}
}