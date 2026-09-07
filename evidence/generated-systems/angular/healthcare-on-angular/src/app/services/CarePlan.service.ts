import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CarePlan} from '../models/CarePlan';
import {PatientService} from '../services/Patient.service';
import {EncounterService} from '../services/Encounter.service';
import {CareTaskService} from '../services/CareTask.service';
import {CareTeamService} from '../services/CareTeam.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CarePlanService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	carePlan : CarePlan;

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
	// add a CarePlan
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCarePlan(planNumber, goalSummary, Patient, Encounters, Tasks, CareTeam, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/CarePlan/create';
		const obj = {
			      		planNumber: planNumber,
      		goalSummary: goalSummary,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Encounters: Encounters != null && Encounters.length > 0 ? Encounters : null,
      		Tasks: Tasks != null && Tasks.length > 0 ? Tasks : null,
      		CareTeam: CareTeam != null && CareTeam.length > 0 ? CareTeam : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CarePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCarePlan(planNumber, goalSummary, Patient, Encounters, Tasks, CareTeam, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CarePlan/update/' + id;
		const obj = {
				      		planNumber: planNumber,
      		goalSummary: goalSummary,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Encounters: Encounters != null && Encounters.length > 0 ? Encounters : null,
      		Tasks: Tasks != null && Tasks.length > 0 ? Tasks : null,
      		CareTeam: CareTeam != null && CareTeam.length > 0 ? CareTeam : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CarePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCarePlan(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CarePlan/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CarePlan
	// returns the results untouched as an Observable CarePlan
	// CarePlan model
	// delegates via URI
	//********************************************************************
	getCarePlan(id) : Observable<CarePlan> {
		const uri_ = this.apiUrl + '/CarePlan/load/' + id;

		return this.http.get<CarePlan>(uri_);
	}
	
	//********************************************************************
	// gets all CarePlan
	// returns the results untouched as JSON representation of an
	// Observable array of CarePlan models
	// delegates via URI
	//********************************************************************
	getCarePlans() : Observable<CarePlan[]> {
		const uri_ = this.apiUrl + '/CarePlan/';

		return this
			.http.get<CarePlan[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Patient on a CarePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( carePlanId, _patientId ): Observable<any> {

		// get the CarePlan from storage
		this.loadHelper( carePlanId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.carePlan.patient = tmp;

	// save the CarePlan
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a CarePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( carePlanId ): Observable<any> {

		// get the CarePlan from storage
		this.loadHelper( carePlanId );

	// assign Patient to null
	this.carePlan.patient = null;

	// save the CarePlan
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CareTeam on a CarePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCareTeam( carePlanId, _careTeamId ): Observable<any> {

		// get the CarePlan from storage
		this.loadHelper( carePlanId );

	// get the CareTeam from storage
	var tmp 	= new CareTeamService(this.http).getCareTeam(_careTeamId);

	// assign the CareTeam
	this.carePlan.careTeam = tmp;

	// save the CarePlan
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CareTeam on a CarePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCareTeam( carePlanId ): Observable<any> {

		// get the CarePlan from storage
		this.loadHelper( carePlanId );

	// assign CareTeam to null
	this.carePlan.careTeam = null;

	// save the CarePlan
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more encountersIds as a Encounters
	// to a CarePlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEncounters( carePlanId, encountersIds ): Observable<any> {

		// get the CarePlan
		this.loadHelper( carePlanId );

	// split on a comma with no spaces
	var idList = encountersIds.split(',')

	// iterate over array of encounters ids
	idList.forEach(function (id) {
		// read the Encounter
		var encounter = new EncounterService(this.http).getEncounter(id);
		// add the Encounter if not already assigned
		if ( this.carePlan.encounters.indexOf(encounter) == -1 )
		this.carePlan.encounters.push(encounter);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more encountersIds as a Encounters
	// from a CarePlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEncounters( carePlanId, encountersIds ): Observable<any> {

		// get the CarePlan
		this.loadHelper( carePlanId );


	// split on a comma with no spaces
	var idList 					= encountersIds.split(',');
	var encounters 	= this.carePlan.encounters;

	if ( encounters != null && encountersIds != null ) {

		// iterate over array of encounters ids
		encounters.forEach(function (obj) {
			if ( encountersIds.indexOf(obj._id) > -1 ) {
				// remove the Encounter
				this.carePlan.encounters.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more tasksIds as a Tasks
	// to a CarePlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTasks( carePlanId, tasksIds ): Observable<any> {

		// get the CarePlan
		this.loadHelper( carePlanId );

	// split on a comma with no spaces
	var idList = tasksIds.split(',')

	// iterate over array of tasks ids
	idList.forEach(function (id) {
		// read the CareTask
		var careTask = new CareTaskService(this.http).getCareTask(id);
		// add the CareTask if not already assigned
		if ( this.carePlan.tasks.indexOf(careTask) == -1 )
		this.carePlan.tasks.push(careTask);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tasksIds as a Tasks
	// from a CarePlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTasks( carePlanId, tasksIds ): Observable<any> {

		// get the CarePlan
		this.loadHelper( carePlanId );


	// split on a comma with no spaces
	var idList 					= tasksIds.split(',');
	var tasks 	= this.carePlan.tasks;

	if ( tasks != null && tasksIds != null ) {

		// iterate over array of tasks ids
		tasks.forEach(function (obj) {
			if ( tasksIds.indexOf(obj._id) > -1 ) {
				// remove the CareTask
				this.carePlan.tasks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a CarePlan
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CarePlan/update/' + this.carePlan;

	return  this.http.post(uri_, this.carePlan );
}

	//********************************************************************
	// loadHelper - internal helper to load a CarePlan
	//********************************************************************	
	loadHelper( id ) {
		this.getCarePlan(id)
			.subscribe((res : CarePlan) => {
				this.carePlan = res;
			});
	}
}