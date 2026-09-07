import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CareTask} from '../models/CareTask';
import {CarePlanService} from '../services/CarePlan.service';
import {ClinicianService} from '../services/Clinician.service';
import {EncounterService} from '../services/Encounter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CareTaskService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	careTask : CareTask;

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
	// add a CareTask
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCareTask(description, dueDate, CarePlan, AssignedTo, Encounter, Status, Priority) : Observable<any> {
		const uri_ = this.apiUrl + '/CareTask/create';
		const obj = {
			      		description: description,
      		dueDate: dueDate,
      		CarePlan: CarePlan != null && CarePlan.length > 0 ? CarePlan : null,
      		AssignedTo: AssignedTo != null && AssignedTo.length > 0 ? AssignedTo : null,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Status: Status,
			Priority: Priority
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CareTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCareTask(description, dueDate, CarePlan, AssignedTo, Encounter, Status, Priority, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CareTask/update/' + id;
		const obj = {
				      		description: description,
      		dueDate: dueDate,
      		CarePlan: CarePlan != null && CarePlan.length > 0 ? CarePlan : null,
      		AssignedTo: AssignedTo != null && AssignedTo.length > 0 ? AssignedTo : null,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Status: Status,
			Priority: Priority
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CareTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCareTask(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CareTask/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CareTask
	// returns the results untouched as an Observable CareTask
	// CareTask model
	// delegates via URI
	//********************************************************************
	getCareTask(id) : Observable<CareTask> {
		const uri_ = this.apiUrl + '/CareTask/load/' + id;

		return this.http.get<CareTask>(uri_);
	}
	
	//********************************************************************
	// gets all CareTask
	// returns the results untouched as JSON representation of an
	// Observable array of CareTask models
	// delegates via URI
	//********************************************************************
	getCareTasks() : Observable<CareTask[]> {
		const uri_ = this.apiUrl + '/CareTask/';

		return this
			.http.get<CareTask[]>(uri_);
	}
	
			//********************************************************************
	// assigns a CarePlan on a CareTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCarePlan( careTaskId, _carePlanId ): Observable<any> {

		// get the CareTask from storage
		this.loadHelper( careTaskId );

	// get the CarePlan from storage
	var tmp 	= new CarePlanService(this.http).getCarePlan(_carePlanId);

	// assign the CarePlan
	this.careTask.carePlan = tmp;

	// save the CareTask
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CarePlan on a CareTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCarePlan( careTaskId ): Observable<any> {

		// get the CareTask from storage
		this.loadHelper( careTaskId );

	// assign CarePlan to null
	this.careTask.carePlan = null;

	// save the CareTask
	return this.saveHelper();
}

		//********************************************************************
	// assigns a AssignedTo on a CareTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAssignedTo( careTaskId, _assignedToId ): Observable<any> {

		// get the CareTask from storage
		this.loadHelper( careTaskId );

	// get the Clinician from storage
	var tmp 	= new ClinicianService(this.http).getClinician(_assignedToId);

	// assign the AssignedTo
	this.careTask.assignedTo = tmp;

	// save the CareTask
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AssignedTo on a CareTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAssignedTo( careTaskId ): Observable<any> {

		// get the CareTask from storage
		this.loadHelper( careTaskId );

	// assign AssignedTo to null
	this.careTask.assignedTo = null;

	// save the CareTask
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Encounter on a CareTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEncounter( careTaskId, _encounterId ): Observable<any> {

		// get the CareTask from storage
		this.loadHelper( careTaskId );

	// get the Encounter from storage
	var tmp 	= new EncounterService(this.http).getEncounter(_encounterId);

	// assign the Encounter
	this.careTask.encounter = tmp;

	// save the CareTask
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Encounter on a CareTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEncounter( careTaskId ): Observable<any> {

		// get the CareTask from storage
		this.loadHelper( careTaskId );

	// assign Encounter to null
	this.careTask.encounter = null;

	// save the CareTask
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CareTask
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CareTask/update/' + this.careTask;

	return  this.http.post(uri_, this.careTask );
}

	//********************************************************************
	// loadHelper - internal helper to load a CareTask
	//********************************************************************	
	loadHelper( id ) {
		this.getCareTask(id)
			.subscribe((res : CareTask) => {
				this.careTask = res;
			});
	}
}