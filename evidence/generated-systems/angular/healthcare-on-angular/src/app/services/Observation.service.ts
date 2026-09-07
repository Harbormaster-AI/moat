import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Observation} from '../models/Observation';
import {EncounterService} from '../services/Encounter.service';
import {PatientService} from '../services/Patient.service';
import {MedicalDeviceService} from '../services/MedicalDevice.service';
import {LabResultService} from '../services/LabResult.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ObservationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	observation : Observation;

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
	// add a Observation
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addObservation(code, value, unit, effectiveDateTime, Encounter, Patient, Device, LabResult, Interpretation) : Observable<any> {
		const uri_ = this.apiUrl + '/Observation/create';
		const obj = {
			      		code: code,
      		value: value,
      		unit: unit,
      		effectiveDateTime: effectiveDateTime,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Device: Device != null && Device.length > 0 ? Device : null,
      		LabResult: LabResult != null && LabResult.length > 0 ? LabResult : null,
			Interpretation: Interpretation
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Observation
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateObservation(code, value, unit, effectiveDateTime, Encounter, Patient, Device, LabResult, Interpretation, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Observation/update/' + id;
		const obj = {
				      		code: code,
      		value: value,
      		unit: unit,
      		effectiveDateTime: effectiveDateTime,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Device: Device != null && Device.length > 0 ? Device : null,
      		LabResult: LabResult != null && LabResult.length > 0 ? LabResult : null,
			Interpretation: Interpretation
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Observation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteObservation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Observation/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Observation
	// returns the results untouched as an Observable Observation
	// Observation model
	// delegates via URI
	//********************************************************************
	getObservation(id) : Observable<Observation> {
		const uri_ = this.apiUrl + '/Observation/load/' + id;

		return this.http.get<Observation>(uri_);
	}
	
	//********************************************************************
	// gets all Observation
	// returns the results untouched as JSON representation of an
	// Observable array of Observation models
	// delegates via URI
	//********************************************************************
	getObservations() : Observable<Observation[]> {
		const uri_ = this.apiUrl + '/Observation/';

		return this
			.http.get<Observation[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Encounter on a Observation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEncounter( observationId, _encounterId ): Observable<any> {

		// get the Observation from storage
		this.loadHelper( observationId );

	// get the Encounter from storage
	var tmp 	= new EncounterService(this.http).getEncounter(_encounterId);

	// assign the Encounter
	this.observation.encounter = tmp;

	// save the Observation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Encounter on a Observation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEncounter( observationId ): Observable<any> {

		// get the Observation from storage
		this.loadHelper( observationId );

	// assign Encounter to null
	this.observation.encounter = null;

	// save the Observation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Patient on a Observation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( observationId, _patientId ): Observable<any> {

		// get the Observation from storage
		this.loadHelper( observationId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.observation.patient = tmp;

	// save the Observation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a Observation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( observationId ): Observable<any> {

		// get the Observation from storage
		this.loadHelper( observationId );

	// assign Patient to null
	this.observation.patient = null;

	// save the Observation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Device on a Observation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDevice( observationId, _deviceId ): Observable<any> {

		// get the Observation from storage
		this.loadHelper( observationId );

	// get the MedicalDevice from storage
	var tmp 	= new MedicalDeviceService(this.http).getMedicalDevice(_deviceId);

	// assign the Device
	this.observation.device = tmp;

	// save the Observation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Device on a Observation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDevice( observationId ): Observable<any> {

		// get the Observation from storage
		this.loadHelper( observationId );

	// assign Device to null
	this.observation.device = null;

	// save the Observation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a LabResult on a Observation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLabResult( observationId, _labResultId ): Observable<any> {

		// get the Observation from storage
		this.loadHelper( observationId );

	// get the LabResult from storage
	var tmp 	= new LabResultService(this.http).getLabResult(_labResultId);

	// assign the LabResult
	this.observation.labResult = tmp;

	// save the Observation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LabResult on a Observation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLabResult( observationId ): Observable<any> {

		// get the Observation from storage
		this.loadHelper( observationId );

	// assign LabResult to null
	this.observation.labResult = null;

	// save the Observation
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Observation
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Observation/update/' + this.observation;

	return  this.http.post(uri_, this.observation );
}

	//********************************************************************
	// loadHelper - internal helper to load a Observation
	//********************************************************************	
	loadHelper( id ) {
		this.getObservation(id)
			.subscribe((res : Observation) => {
				this.observation = res;
			});
	}
}