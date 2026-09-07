import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Discharge} from '../models/Discharge';
import {EncounterService} from '../services/Encounter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DischargeService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	discharge : Discharge;

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
	// add a Discharge
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDischarge(dischargeDateTime, Encounter, Disposition) : Observable<any> {
		const uri_ = this.apiUrl + '/Discharge/create';
		const obj = {
			      		dischargeDateTime: dischargeDateTime,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
			Disposition: Disposition
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Discharge
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDischarge(dischargeDateTime, Encounter, Disposition, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Discharge/update/' + id;
		const obj = {
				      		dischargeDateTime: dischargeDateTime,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
			Disposition: Disposition
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Discharge
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDischarge(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Discharge/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Discharge
	// returns the results untouched as an Observable Discharge
	// Discharge model
	// delegates via URI
	//********************************************************************
	getDischarge(id) : Observable<Discharge> {
		const uri_ = this.apiUrl + '/Discharge/load/' + id;

		return this.http.get<Discharge>(uri_);
	}
	
	//********************************************************************
	// gets all Discharge
	// returns the results untouched as JSON representation of an
	// Observable array of Discharge models
	// delegates via URI
	//********************************************************************
	getDischarges() : Observable<Discharge[]> {
		const uri_ = this.apiUrl + '/Discharge/';

		return this
			.http.get<Discharge[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Encounter on a Discharge
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEncounter( dischargeId, _encounterId ): Observable<any> {

		// get the Discharge from storage
		this.loadHelper( dischargeId );

	// get the Encounter from storage
	var tmp 	= new EncounterService(this.http).getEncounter(_encounterId);

	// assign the Encounter
	this.discharge.encounter = tmp;

	// save the Discharge
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Encounter on a Discharge
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEncounter( dischargeId ): Observable<any> {

		// get the Discharge from storage
		this.loadHelper( dischargeId );

	// assign Encounter to null
	this.discharge.encounter = null;

	// save the Discharge
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Discharge
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Discharge/update/' + this.discharge;

	return  this.http.post(uri_, this.discharge );
}

	//********************************************************************
	// loadHelper - internal helper to load a Discharge
	//********************************************************************	
	loadHelper( id ) {
		this.getDischarge(id)
			.subscribe((res : Discharge) => {
				this.discharge = res;
			});
	}
}