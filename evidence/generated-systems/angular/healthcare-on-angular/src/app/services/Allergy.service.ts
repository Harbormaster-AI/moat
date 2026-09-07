import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Allergy} from '../models/Allergy';
import {PatientService} from '../services/Patient.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AllergyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	allergy : Allergy;

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
	// add a Allergy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAllergy(substance, reaction, Patient, Severity, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Allergy/create';
		const obj = {
			      		substance: substance,
      		reaction: reaction,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Severity: Severity,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Allergy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAllergy(substance, reaction, Patient, Severity, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Allergy/update/' + id;
		const obj = {
				      		substance: substance,
      		reaction: reaction,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Severity: Severity,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Allergy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAllergy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Allergy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Allergy
	// returns the results untouched as an Observable Allergy
	// Allergy model
	// delegates via URI
	//********************************************************************
	getAllergy(id) : Observable<Allergy> {
		const uri_ = this.apiUrl + '/Allergy/load/' + id;

		return this.http.get<Allergy>(uri_);
	}
	
	//********************************************************************
	// gets all Allergy
	// returns the results untouched as JSON representation of an
	// Observable array of Allergy models
	// delegates via URI
	//********************************************************************
	getAllergys() : Observable<Allergy[]> {
		const uri_ = this.apiUrl + '/Allergy/';

		return this
			.http.get<Allergy[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Patient on a Allergy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( allergyId, _patientId ): Observable<any> {

		// get the Allergy from storage
		this.loadHelper( allergyId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.allergy.patient = tmp;

	// save the Allergy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a Allergy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( allergyId ): Observable<any> {

		// get the Allergy from storage
		this.loadHelper( allergyId );

	// assign Patient to null
	this.allergy.patient = null;

	// save the Allergy
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Allergy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Allergy/update/' + this.allergy;

	return  this.http.post(uri_, this.allergy );
}

	//********************************************************************
	// loadHelper - internal helper to load a Allergy
	//********************************************************************	
	loadHelper( id ) {
		this.getAllergy(id)
			.subscribe((res : Allergy) => {
				this.allergy = res;
			});
	}
}