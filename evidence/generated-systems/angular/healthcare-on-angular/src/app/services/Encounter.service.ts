import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Encounter} from '../models/Encounter';
import {PatientService} from '../services/Patient.service';
import {ClinicianService} from '../services/Clinician.service';
import {FacilityService} from '../services/Facility.service';
import {AppointmentService} from '../services/Appointment.service';
import {DiagnosisService} from '../services/Diagnosis.service';
import {ProcedureService} from '../services/Procedure.service';
import {ObservationService} from '../services/Observation.service';
import {ClinicalOrderService} from '../services/ClinicalOrder.service';
import {AdmissionService} from '../services/Admission.service';
import {DischargeService} from '../services/Discharge.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EncounterService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	encounter : Encounter;

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
	// add a Encounter
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEncounter(encounterNumber, startDateTime, endDateTime, Patient, Clinician, Facility, Appointment, Diagnoses, Procedures, Observations, Orders, Admission, Discharge, Status, EncounterType) : Observable<any> {
		const uri_ = this.apiUrl + '/Encounter/create';
		const obj = {
			      		encounterNumber: encounterNumber,
      		startDateTime: startDateTime,
      		endDateTime: endDateTime,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Clinician: Clinician != null && Clinician.length > 0 ? Clinician : null,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		Appointment: Appointment != null && Appointment.length > 0 ? Appointment : null,
      		Diagnoses: Diagnoses != null && Diagnoses.length > 0 ? Diagnoses : null,
      		Procedures: Procedures != null && Procedures.length > 0 ? Procedures : null,
      		Observations: Observations != null && Observations.length > 0 ? Observations : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Admission: Admission != null && Admission.length > 0 ? Admission : null,
      		Discharge: Discharge != null && Discharge.length > 0 ? Discharge : null,
      		Status: Status,
			EncounterType: EncounterType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEncounter(encounterNumber, startDateTime, endDateTime, Patient, Clinician, Facility, Appointment, Diagnoses, Procedures, Observations, Orders, Admission, Discharge, Status, EncounterType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Encounter/update/' + id;
		const obj = {
				      		encounterNumber: encounterNumber,
      		startDateTime: startDateTime,
      		endDateTime: endDateTime,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Clinician: Clinician != null && Clinician.length > 0 ? Clinician : null,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		Appointment: Appointment != null && Appointment.length > 0 ? Appointment : null,
      		Diagnoses: Diagnoses != null && Diagnoses.length > 0 ? Diagnoses : null,
      		Procedures: Procedures != null && Procedures.length > 0 ? Procedures : null,
      		Observations: Observations != null && Observations.length > 0 ? Observations : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Admission: Admission != null && Admission.length > 0 ? Admission : null,
      		Discharge: Discharge != null && Discharge.length > 0 ? Discharge : null,
      		Status: Status,
			EncounterType: EncounterType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEncounter(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Encounter/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Encounter
	// returns the results untouched as an Observable Encounter
	// Encounter model
	// delegates via URI
	//********************************************************************
	getEncounter(id) : Observable<Encounter> {
		const uri_ = this.apiUrl + '/Encounter/load/' + id;

		return this.http.get<Encounter>(uri_);
	}
	
	//********************************************************************
	// gets all Encounter
	// returns the results untouched as JSON representation of an
	// Observable array of Encounter models
	// delegates via URI
	//********************************************************************
	getEncounters() : Observable<Encounter[]> {
		const uri_ = this.apiUrl + '/Encounter/';

		return this
			.http.get<Encounter[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Patient on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( encounterId, _patientId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.encounter.patient = tmp;

	// save the Encounter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( encounterId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// assign Patient to null
	this.encounter.patient = null;

	// save the Encounter
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Clinician on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignClinician( encounterId, _clinicianId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// get the Clinician from storage
	var tmp 	= new ClinicianService(this.http).getClinician(_clinicianId);

	// assign the Clinician
	this.encounter.clinician = tmp;

	// save the Encounter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Clinician on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignClinician( encounterId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// assign Clinician to null
	this.encounter.clinician = null;

	// save the Encounter
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Facility on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFacility( encounterId, _facilityId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// get the Facility from storage
	var tmp 	= new FacilityService(this.http).getFacility(_facilityId);

	// assign the Facility
	this.encounter.facility = tmp;

	// save the Encounter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Facility on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFacility( encounterId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// assign Facility to null
	this.encounter.facility = null;

	// save the Encounter
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Appointment on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAppointment( encounterId, _appointmentId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// get the Appointment from storage
	var tmp 	= new AppointmentService(this.http).getAppointment(_appointmentId);

	// assign the Appointment
	this.encounter.appointment = tmp;

	// save the Encounter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Appointment on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAppointment( encounterId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// assign Appointment to null
	this.encounter.appointment = null;

	// save the Encounter
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Admission on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdmission( encounterId, _admissionId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// get the Admission from storage
	var tmp 	= new AdmissionService(this.http).getAdmission(_admissionId);

	// assign the Admission
	this.encounter.admission = tmp;

	// save the Encounter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Admission on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdmission( encounterId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// assign Admission to null
	this.encounter.admission = null;

	// save the Encounter
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Discharge on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDischarge( encounterId, _dischargeId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// get the Discharge from storage
	var tmp 	= new DischargeService(this.http).getDischarge(_dischargeId);

	// assign the Discharge
	this.encounter.discharge = tmp;

	// save the Encounter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Discharge on a Encounter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDischarge( encounterId ): Observable<any> {

		// get the Encounter from storage
		this.loadHelper( encounterId );

	// assign Discharge to null
	this.encounter.discharge = null;

	// save the Encounter
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more diagnosesIds as a Diagnoses
	// to a Encounter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDiagnoses( encounterId, diagnosesIds ): Observable<any> {

		// get the Encounter
		this.loadHelper( encounterId );

	// split on a comma with no spaces
	var idList = diagnosesIds.split(',')

	// iterate over array of diagnoses ids
	idList.forEach(function (id) {
		// read the Diagnosis
		var diagnosis = new DiagnosisService(this.http).getDiagnosis(id);
		// add the Diagnosis if not already assigned
		if ( this.encounter.diagnoses.indexOf(diagnosis) == -1 )
		this.encounter.diagnoses.push(diagnosis);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more diagnosesIds as a Diagnoses
	// from a Encounter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDiagnoses( encounterId, diagnosesIds ): Observable<any> {

		// get the Encounter
		this.loadHelper( encounterId );


	// split on a comma with no spaces
	var idList 					= diagnosesIds.split(',');
	var diagnoses 	= this.encounter.diagnoses;

	if ( diagnoses != null && diagnosesIds != null ) {

		// iterate over array of diagnoses ids
		diagnoses.forEach(function (obj) {
			if ( diagnosesIds.indexOf(obj._id) > -1 ) {
				// remove the Diagnosis
				this.encounter.diagnoses.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more proceduresIds as a Procedures
	// to a Encounter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcedures( encounterId, proceduresIds ): Observable<any> {

		// get the Encounter
		this.loadHelper( encounterId );

	// split on a comma with no spaces
	var idList = proceduresIds.split(',')

	// iterate over array of procedures ids
	idList.forEach(function (id) {
		// read the Procedure
		var procedure = new ProcedureService(this.http).getProcedure(id);
		// add the Procedure if not already assigned
		if ( this.encounter.procedures.indexOf(procedure) == -1 )
		this.encounter.procedures.push(procedure);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more proceduresIds as a Procedures
	// from a Encounter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcedures( encounterId, proceduresIds ): Observable<any> {

		// get the Encounter
		this.loadHelper( encounterId );


	// split on a comma with no spaces
	var idList 					= proceduresIds.split(',');
	var procedures 	= this.encounter.procedures;

	if ( procedures != null && proceduresIds != null ) {

		// iterate over array of procedures ids
		procedures.forEach(function (obj) {
			if ( proceduresIds.indexOf(obj._id) > -1 ) {
				// remove the Procedure
				this.encounter.procedures.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more observationsIds as a Observations
	// to a Encounter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addObservations( encounterId, observationsIds ): Observable<any> {

		// get the Encounter
		this.loadHelper( encounterId );

	// split on a comma with no spaces
	var idList = observationsIds.split(',')

	// iterate over array of observations ids
	idList.forEach(function (id) {
		// read the Observation
		var observation = new ObservationService(this.http).getObservation(id);
		// add the Observation if not already assigned
		if ( this.encounter.observations.indexOf(observation) == -1 )
		this.encounter.observations.push(observation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more observationsIds as a Observations
	// from a Encounter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeObservations( encounterId, observationsIds ): Observable<any> {

		// get the Encounter
		this.loadHelper( encounterId );


	// split on a comma with no spaces
	var idList 					= observationsIds.split(',');
	var observations 	= this.encounter.observations;

	if ( observations != null && observationsIds != null ) {

		// iterate over array of observations ids
		observations.forEach(function (obj) {
			if ( observationsIds.indexOf(obj._id) > -1 ) {
				// remove the Observation
				this.encounter.observations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ordersIds as a Orders
	// to a Encounter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOrders( encounterId, ordersIds ): Observable<any> {

		// get the Encounter
		this.loadHelper( encounterId );

	// split on a comma with no spaces
	var idList = ordersIds.split(',')

	// iterate over array of orders ids
	idList.forEach(function (id) {
		// read the ClinicalOrder
		var clinicalOrder = new ClinicalOrderService(this.http).getClinicalOrder(id);
		// add the ClinicalOrder if not already assigned
		if ( this.encounter.orders.indexOf(clinicalOrder) == -1 )
		this.encounter.orders.push(clinicalOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ordersIds as a Orders
	// from a Encounter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOrders( encounterId, ordersIds ): Observable<any> {

		// get the Encounter
		this.loadHelper( encounterId );


	// split on a comma with no spaces
	var idList 					= ordersIds.split(',');
	var orders 	= this.encounter.orders;

	if ( orders != null && ordersIds != null ) {

		// iterate over array of orders ids
		orders.forEach(function (obj) {
			if ( ordersIds.indexOf(obj._id) > -1 ) {
				// remove the ClinicalOrder
				this.encounter.orders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Encounter
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Encounter/update/' + this.encounter;

	return  this.http.post(uri_, this.encounter );
}

	//********************************************************************
	// loadHelper - internal helper to load a Encounter
	//********************************************************************	
	loadHelper( id ) {
		this.getEncounter(id)
			.subscribe((res : Encounter) => {
				this.encounter = res;
			});
	}
}