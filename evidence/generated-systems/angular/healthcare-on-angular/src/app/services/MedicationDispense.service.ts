import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {MedicationDispense} from '../models/MedicationDispense';
import {MedicationOrderService} from '../services/MedicationOrder.service';
import {PharmacyService} from '../services/Pharmacy.service';
import {PatientService} from '../services/Patient.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MedicationDispenseService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	medicationDispense : MedicationDispense;

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
	// add a MedicationDispense
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMedicationDispense(dispenseNumber, quantity, whenPrepared, MedicationOrder, Pharmacy, Patient, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/MedicationDispense/create';
		const obj = {
			      		dispenseNumber: dispenseNumber,
      		quantity: quantity,
      		whenPrepared: whenPrepared,
      		MedicationOrder: MedicationOrder != null && MedicationOrder.length > 0 ? MedicationOrder : null,
      		Pharmacy: Pharmacy != null && Pharmacy.length > 0 ? Pharmacy : null,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a MedicationDispense
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMedicationDispense(dispenseNumber, quantity, whenPrepared, MedicationOrder, Pharmacy, Patient, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/MedicationDispense/update/' + id;
		const obj = {
				      		dispenseNumber: dispenseNumber,
      		quantity: quantity,
      		whenPrepared: whenPrepared,
      		MedicationOrder: MedicationOrder != null && MedicationOrder.length > 0 ? MedicationOrder : null,
      		Pharmacy: Pharmacy != null && Pharmacy.length > 0 ? Pharmacy : null,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a MedicationDispense
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMedicationDispense(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/MedicationDispense/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a MedicationDispense
	// returns the results untouched as an Observable MedicationDispense
	// MedicationDispense model
	// delegates via URI
	//********************************************************************
	getMedicationDispense(id) : Observable<MedicationDispense> {
		const uri_ = this.apiUrl + '/MedicationDispense/load/' + id;

		return this.http.get<MedicationDispense>(uri_);
	}
	
	//********************************************************************
	// gets all MedicationDispense
	// returns the results untouched as JSON representation of an
	// Observable array of MedicationDispense models
	// delegates via URI
	//********************************************************************
	getMedicationDispenses() : Observable<MedicationDispense[]> {
		const uri_ = this.apiUrl + '/MedicationDispense/';

		return this
			.http.get<MedicationDispense[]>(uri_);
	}
	
			//********************************************************************
	// assigns a MedicationOrder on a MedicationDispense
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMedicationOrder( medicationDispenseId, _medicationOrderId ): Observable<any> {

		// get the MedicationDispense from storage
		this.loadHelper( medicationDispenseId );

	// get the MedicationOrder from storage
	var tmp 	= new MedicationOrderService(this.http).getMedicationOrder(_medicationOrderId);

	// assign the MedicationOrder
	this.medicationDispense.medicationOrder = tmp;

	// save the MedicationDispense
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a MedicationOrder on a MedicationDispense
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMedicationOrder( medicationDispenseId ): Observable<any> {

		// get the MedicationDispense from storage
		this.loadHelper( medicationDispenseId );

	// assign MedicationOrder to null
	this.medicationDispense.medicationOrder = null;

	// save the MedicationDispense
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Pharmacy on a MedicationDispense
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPharmacy( medicationDispenseId, _pharmacyId ): Observable<any> {

		// get the MedicationDispense from storage
		this.loadHelper( medicationDispenseId );

	// get the Pharmacy from storage
	var tmp 	= new PharmacyService(this.http).getPharmacy(_pharmacyId);

	// assign the Pharmacy
	this.medicationDispense.pharmacy = tmp;

	// save the MedicationDispense
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Pharmacy on a MedicationDispense
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPharmacy( medicationDispenseId ): Observable<any> {

		// get the MedicationDispense from storage
		this.loadHelper( medicationDispenseId );

	// assign Pharmacy to null
	this.medicationDispense.pharmacy = null;

	// save the MedicationDispense
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Patient on a MedicationDispense
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( medicationDispenseId, _patientId ): Observable<any> {

		// get the MedicationDispense from storage
		this.loadHelper( medicationDispenseId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.medicationDispense.patient = tmp;

	// save the MedicationDispense
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a MedicationDispense
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( medicationDispenseId ): Observable<any> {

		// get the MedicationDispense from storage
		this.loadHelper( medicationDispenseId );

	// assign Patient to null
	this.medicationDispense.patient = null;

	// save the MedicationDispense
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a MedicationDispense
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/MedicationDispense/update/' + this.medicationDispense;

	return  this.http.post(uri_, this.medicationDispense );
}

	//********************************************************************
	// loadHelper - internal helper to load a MedicationDispense
	//********************************************************************	
	loadHelper( id ) {
		this.getMedicationDispense(id)
			.subscribe((res : MedicationDispense) => {
				this.medicationDispense = res;
			});
	}
}