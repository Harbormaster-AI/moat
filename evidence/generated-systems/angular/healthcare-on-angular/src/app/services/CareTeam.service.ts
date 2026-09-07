import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CareTeam} from '../models/CareTeam';
import {DepartmentService} from '../services/Department.service';
import {ClinicianService} from '../services/Clinician.service';
import {PatientService} from '../services/Patient.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CareTeamService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	careTeam : CareTeam;

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
	// add a CareTeam
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCareTeam(name, Department, Clinicians, Patients, CareSetting) : Observable<any> {
		const uri_ = this.apiUrl + '/CareTeam/create';
		const obj = {
			      		name: name,
      		Department: Department != null && Department.length > 0 ? Department : null,
      		Clinicians: Clinicians != null && Clinicians.length > 0 ? Clinicians : null,
      		Patients: Patients != null && Patients.length > 0 ? Patients : null,
			CareSetting: CareSetting
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CareTeam
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCareTeam(name, Department, Clinicians, Patients, CareSetting, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CareTeam/update/' + id;
		const obj = {
				      		name: name,
      		Department: Department != null && Department.length > 0 ? Department : null,
      		Clinicians: Clinicians != null && Clinicians.length > 0 ? Clinicians : null,
      		Patients: Patients != null && Patients.length > 0 ? Patients : null,
			CareSetting: CareSetting
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CareTeam
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCareTeam(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CareTeam/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CareTeam
	// returns the results untouched as an Observable CareTeam
	// CareTeam model
	// delegates via URI
	//********************************************************************
	getCareTeam(id) : Observable<CareTeam> {
		const uri_ = this.apiUrl + '/CareTeam/load/' + id;

		return this.http.get<CareTeam>(uri_);
	}
	
	//********************************************************************
	// gets all CareTeam
	// returns the results untouched as JSON representation of an
	// Observable array of CareTeam models
	// delegates via URI
	//********************************************************************
	getCareTeams() : Observable<CareTeam[]> {
		const uri_ = this.apiUrl + '/CareTeam/';

		return this
			.http.get<CareTeam[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Department on a CareTeam
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDepartment( careTeamId, _departmentId ): Observable<any> {

		// get the CareTeam from storage
		this.loadHelper( careTeamId );

	// get the Department from storage
	var tmp 	= new DepartmentService(this.http).getDepartment(_departmentId);

	// assign the Department
	this.careTeam.department = tmp;

	// save the CareTeam
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Department on a CareTeam
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDepartment( careTeamId ): Observable<any> {

		// get the CareTeam from storage
		this.loadHelper( careTeamId );

	// assign Department to null
	this.careTeam.department = null;

	// save the CareTeam
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more cliniciansIds as a Clinicians
	// to a CareTeam
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addClinicians( careTeamId, cliniciansIds ): Observable<any> {

		// get the CareTeam
		this.loadHelper( careTeamId );

	// split on a comma with no spaces
	var idList = cliniciansIds.split(',')

	// iterate over array of clinicians ids
	idList.forEach(function (id) {
		// read the Clinician
		var clinician = new ClinicianService(this.http).getClinician(id);
		// add the Clinician if not already assigned
		if ( this.careTeam.clinicians.indexOf(clinician) == -1 )
		this.careTeam.clinicians.push(clinician);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more cliniciansIds as a Clinicians
	// from a CareTeam
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeClinicians( careTeamId, cliniciansIds ): Observable<any> {

		// get the CareTeam
		this.loadHelper( careTeamId );


	// split on a comma with no spaces
	var idList 					= cliniciansIds.split(',');
	var clinicians 	= this.careTeam.clinicians;

	if ( clinicians != null && cliniciansIds != null ) {

		// iterate over array of clinicians ids
		clinicians.forEach(function (obj) {
			if ( cliniciansIds.indexOf(obj._id) > -1 ) {
				// remove the Clinician
				this.careTeam.clinicians.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more patientsIds as a Patients
	// to a CareTeam
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPatients( careTeamId, patientsIds ): Observable<any> {

		// get the CareTeam
		this.loadHelper( careTeamId );

	// split on a comma with no spaces
	var idList = patientsIds.split(',')

	// iterate over array of patients ids
	idList.forEach(function (id) {
		// read the Patient
		var patient = new PatientService(this.http).getPatient(id);
		// add the Patient if not already assigned
		if ( this.careTeam.patients.indexOf(patient) == -1 )
		this.careTeam.patients.push(patient);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more patientsIds as a Patients
	// from a CareTeam
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePatients( careTeamId, patientsIds ): Observable<any> {

		// get the CareTeam
		this.loadHelper( careTeamId );


	// split on a comma with no spaces
	var idList 					= patientsIds.split(',');
	var patients 	= this.careTeam.patients;

	if ( patients != null && patientsIds != null ) {

		// iterate over array of patients ids
		patients.forEach(function (obj) {
			if ( patientsIds.indexOf(obj._id) > -1 ) {
				// remove the Patient
				this.careTeam.patients.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a CareTeam
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CareTeam/update/' + this.careTeam;

	return  this.http.post(uri_, this.careTeam );
}

	//********************************************************************
	// loadHelper - internal helper to load a CareTeam
	//********************************************************************	
	loadHelper( id ) {
		this.getCareTeam(id)
			.subscribe((res : CareTeam) => {
				this.careTeam = res;
			});
	}
}