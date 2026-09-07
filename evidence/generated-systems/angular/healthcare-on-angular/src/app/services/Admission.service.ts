import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Admission} from '../models/Admission';
import {EncounterService} from '../services/Encounter.service';
import {FacilityService} from '../services/Facility.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AdmissionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	admission : Admission;

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
	// add a Admission
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAdmission(admitDateTime, bed, Encounter, Facility, AdmissionType) : Observable<any> {
		const uri_ = this.apiUrl + '/Admission/create';
		const obj = {
			      		admitDateTime: admitDateTime,
      		bed: bed,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
			AdmissionType: AdmissionType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Admission
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAdmission(admitDateTime, bed, Encounter, Facility, AdmissionType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Admission/update/' + id;
		const obj = {
				      		admitDateTime: admitDateTime,
      		bed: bed,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
			AdmissionType: AdmissionType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Admission
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAdmission(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Admission/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Admission
	// returns the results untouched as an Observable Admission
	// Admission model
	// delegates via URI
	//********************************************************************
	getAdmission(id) : Observable<Admission> {
		const uri_ = this.apiUrl + '/Admission/load/' + id;

		return this.http.get<Admission>(uri_);
	}
	
	//********************************************************************
	// gets all Admission
	// returns the results untouched as JSON representation of an
	// Observable array of Admission models
	// delegates via URI
	//********************************************************************
	getAdmissions() : Observable<Admission[]> {
		const uri_ = this.apiUrl + '/Admission/';

		return this
			.http.get<Admission[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Encounter on a Admission
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEncounter( admissionId, _encounterId ): Observable<any> {

		// get the Admission from storage
		this.loadHelper( admissionId );

	// get the Encounter from storage
	var tmp 	= new EncounterService(this.http).getEncounter(_encounterId);

	// assign the Encounter
	this.admission.encounter = tmp;

	// save the Admission
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Encounter on a Admission
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEncounter( admissionId ): Observable<any> {

		// get the Admission from storage
		this.loadHelper( admissionId );

	// assign Encounter to null
	this.admission.encounter = null;

	// save the Admission
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Facility on a Admission
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFacility( admissionId, _facilityId ): Observable<any> {

		// get the Admission from storage
		this.loadHelper( admissionId );

	// get the Facility from storage
	var tmp 	= new FacilityService(this.http).getFacility(_facilityId);

	// assign the Facility
	this.admission.facility = tmp;

	// save the Admission
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Facility on a Admission
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFacility( admissionId ): Observable<any> {

		// get the Admission from storage
		this.loadHelper( admissionId );

	// assign Facility to null
	this.admission.facility = null;

	// save the Admission
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Admission
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Admission/update/' + this.admission;

	return  this.http.post(uri_, this.admission );
}

	//********************************************************************
	// loadHelper - internal helper to load a Admission
	//********************************************************************	
	loadHelper( id ) {
		this.getAdmission(id)
			.subscribe((res : Admission) => {
				this.admission = res;
			});
	}
}