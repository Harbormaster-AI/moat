import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Appointment} from '../models/Appointment';
import {PatientService} from '../services/Patient.service';
import {ClinicianService} from '../services/Clinician.service';
import {FacilityService} from '../services/Facility.service';
import {EncounterService} from '../services/Encounter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AppointmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	appointment : Appointment;

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
	// add a Appointment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAppointment(appointmentDate, reason, Patient, Clinician, Facility, Encounter, Status, Priority) : Observable<any> {
		const uri_ = this.apiUrl + '/Appointment/create';
		const obj = {
			      		appointmentDate: appointmentDate,
      		reason: reason,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Clinician: Clinician != null && Clinician.length > 0 ? Clinician : null,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Status: Status,
			Priority: Priority
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Appointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAppointment(appointmentDate, reason, Patient, Clinician, Facility, Encounter, Status, Priority, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Appointment/update/' + id;
		const obj = {
				      		appointmentDate: appointmentDate,
      		reason: reason,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Clinician: Clinician != null && Clinician.length > 0 ? Clinician : null,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Status: Status,
			Priority: Priority
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Appointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAppointment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Appointment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Appointment
	// returns the results untouched as an Observable Appointment
	// Appointment model
	// delegates via URI
	//********************************************************************
	getAppointment(id) : Observable<Appointment> {
		const uri_ = this.apiUrl + '/Appointment/load/' + id;

		return this.http.get<Appointment>(uri_);
	}
	
	//********************************************************************
	// gets all Appointment
	// returns the results untouched as JSON representation of an
	// Observable array of Appointment models
	// delegates via URI
	//********************************************************************
	getAppointments() : Observable<Appointment[]> {
		const uri_ = this.apiUrl + '/Appointment/';

		return this
			.http.get<Appointment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Patient on a Appointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( appointmentId, _patientId ): Observable<any> {

		// get the Appointment from storage
		this.loadHelper( appointmentId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.appointment.patient = tmp;

	// save the Appointment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a Appointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( appointmentId ): Observable<any> {

		// get the Appointment from storage
		this.loadHelper( appointmentId );

	// assign Patient to null
	this.appointment.patient = null;

	// save the Appointment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Clinician on a Appointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignClinician( appointmentId, _clinicianId ): Observable<any> {

		// get the Appointment from storage
		this.loadHelper( appointmentId );

	// get the Clinician from storage
	var tmp 	= new ClinicianService(this.http).getClinician(_clinicianId);

	// assign the Clinician
	this.appointment.clinician = tmp;

	// save the Appointment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Clinician on a Appointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignClinician( appointmentId ): Observable<any> {

		// get the Appointment from storage
		this.loadHelper( appointmentId );

	// assign Clinician to null
	this.appointment.clinician = null;

	// save the Appointment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Facility on a Appointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFacility( appointmentId, _facilityId ): Observable<any> {

		// get the Appointment from storage
		this.loadHelper( appointmentId );

	// get the Facility from storage
	var tmp 	= new FacilityService(this.http).getFacility(_facilityId);

	// assign the Facility
	this.appointment.facility = tmp;

	// save the Appointment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Facility on a Appointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFacility( appointmentId ): Observable<any> {

		// get the Appointment from storage
		this.loadHelper( appointmentId );

	// assign Facility to null
	this.appointment.facility = null;

	// save the Appointment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Encounter on a Appointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEncounter( appointmentId, _encounterId ): Observable<any> {

		// get the Appointment from storage
		this.loadHelper( appointmentId );

	// get the Encounter from storage
	var tmp 	= new EncounterService(this.http).getEncounter(_encounterId);

	// assign the Encounter
	this.appointment.encounter = tmp;

	// save the Appointment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Encounter on a Appointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEncounter( appointmentId ): Observable<any> {

		// get the Appointment from storage
		this.loadHelper( appointmentId );

	// assign Encounter to null
	this.appointment.encounter = null;

	// save the Appointment
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Appointment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Appointment/update/' + this.appointment;

	return  this.http.post(uri_, this.appointment );
}

	//********************************************************************
	// loadHelper - internal helper to load a Appointment
	//********************************************************************	
	loadHelper( id ) {
		this.getAppointment(id)
			.subscribe((res : Appointment) => {
				this.appointment = res;
			});
	}
}