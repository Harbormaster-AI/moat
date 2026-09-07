import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {FlightHealthEvent} from '../models/FlightHealthEvent';
import {ConnectedAircraftService} from '../services/ConnectedAircraft.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class FlightHealthEventService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	flightHealthEvent : FlightHealthEvent;

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
	// add a FlightHealthEvent
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addFlightHealthEvent(eventCode, ConnectedAircraft, Severity) : Observable<any> {
		const uri_ = this.apiUrl + '/FlightHealthEvent/create';
		const obj = {
			      		eventCode: eventCode,
      		ConnectedAircraft: ConnectedAircraft != null && ConnectedAircraft.length > 0 ? ConnectedAircraft : null,
			Severity: Severity
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a FlightHealthEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateFlightHealthEvent(eventCode, ConnectedAircraft, Severity, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/FlightHealthEvent/update/' + id;
		const obj = {
				      		eventCode: eventCode,
      		ConnectedAircraft: ConnectedAircraft != null && ConnectedAircraft.length > 0 ? ConnectedAircraft : null,
			Severity: Severity
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a FlightHealthEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteFlightHealthEvent(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/FlightHealthEvent/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a FlightHealthEvent
	// returns the results untouched as an Observable FlightHealthEvent
	// FlightHealthEvent model
	// delegates via URI
	//********************************************************************
	getFlightHealthEvent(id) : Observable<FlightHealthEvent> {
		const uri_ = this.apiUrl + '/FlightHealthEvent/load/' + id;

		return this.http.get<FlightHealthEvent>(uri_);
	}
	
	//********************************************************************
	// gets all FlightHealthEvent
	// returns the results untouched as JSON representation of an
	// Observable array of FlightHealthEvent models
	// delegates via URI
	//********************************************************************
	getFlightHealthEvents() : Observable<FlightHealthEvent[]> {
		const uri_ = this.apiUrl + '/FlightHealthEvent/';

		return this
			.http.get<FlightHealthEvent[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ConnectedAircraft on a FlightHealthEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignConnectedAircraft( flightHealthEventId, _connectedAircraftId ): Observable<any> {

		// get the FlightHealthEvent from storage
		this.loadHelper( flightHealthEventId );

	// get the ConnectedAircraft from storage
	var tmp 	= new ConnectedAircraftService(this.http).getConnectedAircraft(_connectedAircraftId);

	// assign the ConnectedAircraft
	this.flightHealthEvent.connectedAircraft = tmp;

	// save the FlightHealthEvent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ConnectedAircraft on a FlightHealthEvent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignConnectedAircraft( flightHealthEventId ): Observable<any> {

		// get the FlightHealthEvent from storage
		this.loadHelper( flightHealthEventId );

	// assign ConnectedAircraft to null
	this.flightHealthEvent.connectedAircraft = null;

	// save the FlightHealthEvent
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a FlightHealthEvent
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/FlightHealthEvent/update/' + this.flightHealthEvent;

	return  this.http.post(uri_, this.flightHealthEvent );
}

	//********************************************************************
	// loadHelper - internal helper to load a FlightHealthEvent
	//********************************************************************	
	loadHelper( id ) {
		this.getFlightHealthEvent(id)
			.subscribe((res : FlightHealthEvent) => {
				this.flightHealthEvent = res;
			});
	}
}