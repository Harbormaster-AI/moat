import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Warranty} from '../models/Warranty';
import {AircraftService} from '../services/Aircraft.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class WarrantyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	warranty : Warranty;

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
	// add a Warranty
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addWarranty(coverageMonths, Aircraft, WarrantyType) : Observable<any> {
		const uri_ = this.apiUrl + '/Warranty/create';
		const obj = {
			      		coverageMonths: coverageMonths,
      		Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null,
			WarrantyType: WarrantyType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Warranty
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateWarranty(coverageMonths, Aircraft, WarrantyType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Warranty/update/' + id;
		const obj = {
				      		coverageMonths: coverageMonths,
      		Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null,
			WarrantyType: WarrantyType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Warranty
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteWarranty(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Warranty/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Warranty
	// returns the results untouched as an Observable Warranty
	// Warranty model
	// delegates via URI
	//********************************************************************
	getWarranty(id) : Observable<Warranty> {
		const uri_ = this.apiUrl + '/Warranty/load/' + id;

		return this.http.get<Warranty>(uri_);
	}
	
	//********************************************************************
	// gets all Warranty
	// returns the results untouched as JSON representation of an
	// Observable array of Warranty models
	// delegates via URI
	//********************************************************************
	getWarrantys() : Observable<Warranty[]> {
		const uri_ = this.apiUrl + '/Warranty/';

		return this
			.http.get<Warranty[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Aircraft on a Warranty
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAircraft( warrantyId, _aircraftId ): Observable<any> {

		// get the Warranty from storage
		this.loadHelper( warrantyId );

	// get the Aircraft from storage
	var tmp 	= new AircraftService(this.http).getAircraft(_aircraftId);

	// assign the Aircraft
	this.warranty.aircraft = tmp;

	// save the Warranty
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Aircraft on a Warranty
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAircraft( warrantyId ): Observable<any> {

		// get the Warranty from storage
		this.loadHelper( warrantyId );

	// assign Aircraft to null
	this.warranty.aircraft = null;

	// save the Warranty
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Warranty
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Warranty/update/' + this.warranty;

	return  this.http.post(uri_, this.warranty );
}

	//********************************************************************
	// loadHelper - internal helper to load a Warranty
	//********************************************************************	
	loadHelper( id ) {
		this.getWarranty(id)
			.subscribe((res : Warranty) => {
				this.warranty = res;
			});
	}
}