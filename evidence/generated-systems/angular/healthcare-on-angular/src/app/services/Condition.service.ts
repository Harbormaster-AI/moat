import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Condition} from '../models/Condition';
import {PatientService} from '../services/Patient.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ConditionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	condition : Condition;

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
	// add a Condition
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCondition(code, onsetDate, abatementDate, Patient, ClinicalStatus, VerificationStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/Condition/create';
		const obj = {
			      		code: code,
      		onsetDate: onsetDate,
      		abatementDate: abatementDate,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		ClinicalStatus: ClinicalStatus,
			VerificationStatus: VerificationStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Condition
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCondition(code, onsetDate, abatementDate, Patient, ClinicalStatus, VerificationStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Condition/update/' + id;
		const obj = {
				      		code: code,
      		onsetDate: onsetDate,
      		abatementDate: abatementDate,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		ClinicalStatus: ClinicalStatus,
			VerificationStatus: VerificationStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Condition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCondition(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Condition/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Condition
	// returns the results untouched as an Observable Condition
	// Condition model
	// delegates via URI
	//********************************************************************
	getCondition(id) : Observable<Condition> {
		const uri_ = this.apiUrl + '/Condition/load/' + id;

		return this.http.get<Condition>(uri_);
	}
	
	//********************************************************************
	// gets all Condition
	// returns the results untouched as JSON representation of an
	// Observable array of Condition models
	// delegates via URI
	//********************************************************************
	getConditions() : Observable<Condition[]> {
		const uri_ = this.apiUrl + '/Condition/';

		return this
			.http.get<Condition[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Patient on a Condition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( conditionId, _patientId ): Observable<any> {

		// get the Condition from storage
		this.loadHelper( conditionId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.condition.patient = tmp;

	// save the Condition
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a Condition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( conditionId ): Observable<any> {

		// get the Condition from storage
		this.loadHelper( conditionId );

	// assign Patient to null
	this.condition.patient = null;

	// save the Condition
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Condition
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Condition/update/' + this.condition;

	return  this.http.post(uri_, this.condition );
}

	//********************************************************************
	// loadHelper - internal helper to load a Condition
	//********************************************************************	
	loadHelper( id ) {
		this.getCondition(id)
			.subscribe((res : Condition) => {
				this.condition = res;
			});
	}
}