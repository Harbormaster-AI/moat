import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Registration} from '../models/Registration';
import {AircraftService} from '../services/Aircraft.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RegistrationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	registration : Registration;

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
	// add a Registration
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRegistration(tailNumber, registryCountry, Aircraft) : Observable<any> {
		const uri_ = this.apiUrl + '/Registration/create';
		const obj = {
			      		tailNumber: tailNumber,
      		registryCountry: registryCountry,
			Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Registration
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRegistration(tailNumber, registryCountry, Aircraft, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Registration/update/' + id;
		const obj = {
				      		tailNumber: tailNumber,
      		registryCountry: registryCountry,
			Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Registration
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRegistration(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Registration/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Registration
	// returns the results untouched as an Observable Registration
	// Registration model
	// delegates via URI
	//********************************************************************
	getRegistration(id) : Observable<Registration> {
		const uri_ = this.apiUrl + '/Registration/load/' + id;

		return this.http.get<Registration>(uri_);
	}
	
	//********************************************************************
	// gets all Registration
	// returns the results untouched as JSON representation of an
	// Observable array of Registration models
	// delegates via URI
	//********************************************************************
	getRegistrations() : Observable<Registration[]> {
		const uri_ = this.apiUrl + '/Registration/';

		return this
			.http.get<Registration[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Aircraft on a Registration
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAircraft( registrationId, _aircraftId ): Observable<any> {

		// get the Registration from storage
		this.loadHelper( registrationId );

	// get the Aircraft from storage
	var tmp 	= new AircraftService(this.http).getAircraft(_aircraftId);

	// assign the Aircraft
	this.registration.aircraft = tmp;

	// save the Registration
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Aircraft on a Registration
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAircraft( registrationId ): Observable<any> {

		// get the Registration from storage
		this.loadHelper( registrationId );

	// assign Aircraft to null
	this.registration.aircraft = null;

	// save the Registration
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Registration
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Registration/update/' + this.registration;

	return  this.http.post(uri_, this.registration );
}

	//********************************************************************
	// loadHelper - internal helper to load a Registration
	//********************************************************************	
	loadHelper( id ) {
		this.getRegistration(id)
			.subscribe((res : Registration) => {
				this.registration = res;
			});
	}
}