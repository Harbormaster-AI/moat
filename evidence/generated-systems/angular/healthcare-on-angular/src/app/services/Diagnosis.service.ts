import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Diagnosis} from '../models/Diagnosis';
import {EncounterService} from '../services/Encounter.service';
import {PatientService} from '../services/Patient.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DiagnosisService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	diagnosis : Diagnosis;

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
	// add a Diagnosis
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDiagnosis(code, description, onsetDate, Encounter, Patient, Certainty) : Observable<any> {
		const uri_ = this.apiUrl + '/Diagnosis/create';
		const obj = {
			      		code: code,
      		description: description,
      		onsetDate: onsetDate,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
			Certainty: Certainty
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Diagnosis
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDiagnosis(code, description, onsetDate, Encounter, Patient, Certainty, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Diagnosis/update/' + id;
		const obj = {
				      		code: code,
      		description: description,
      		onsetDate: onsetDate,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
			Certainty: Certainty
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Diagnosis
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDiagnosis(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Diagnosis/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Diagnosis
	// returns the results untouched as an Observable Diagnosis
	// Diagnosis model
	// delegates via URI
	//********************************************************************
	getDiagnosis(id) : Observable<Diagnosis> {
		const uri_ = this.apiUrl + '/Diagnosis/load/' + id;

		return this.http.get<Diagnosis>(uri_);
	}
	
	//********************************************************************
	// gets all Diagnosis
	// returns the results untouched as JSON representation of an
	// Observable array of Diagnosis models
	// delegates via URI
	//********************************************************************
	getDiagnosiss() : Observable<Diagnosis[]> {
		const uri_ = this.apiUrl + '/Diagnosis/';

		return this
			.http.get<Diagnosis[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Encounter on a Diagnosis
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEncounter( diagnosisId, _encounterId ): Observable<any> {

		// get the Diagnosis from storage
		this.loadHelper( diagnosisId );

	// get the Encounter from storage
	var tmp 	= new EncounterService(this.http).getEncounter(_encounterId);

	// assign the Encounter
	this.diagnosis.encounter = tmp;

	// save the Diagnosis
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Encounter on a Diagnosis
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEncounter( diagnosisId ): Observable<any> {

		// get the Diagnosis from storage
		this.loadHelper( diagnosisId );

	// assign Encounter to null
	this.diagnosis.encounter = null;

	// save the Diagnosis
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Patient on a Diagnosis
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( diagnosisId, _patientId ): Observable<any> {

		// get the Diagnosis from storage
		this.loadHelper( diagnosisId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.diagnosis.patient = tmp;

	// save the Diagnosis
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a Diagnosis
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( diagnosisId ): Observable<any> {

		// get the Diagnosis from storage
		this.loadHelper( diagnosisId );

	// assign Patient to null
	this.diagnosis.patient = null;

	// save the Diagnosis
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Diagnosis
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Diagnosis/update/' + this.diagnosis;

	return  this.http.post(uri_, this.diagnosis );
}

	//********************************************************************
	// loadHelper - internal helper to load a Diagnosis
	//********************************************************************	
	loadHelper( id ) {
		this.getDiagnosis(id)
			.subscribe((res : Diagnosis) => {
				this.diagnosis = res;
			});
	}
}