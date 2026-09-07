import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Procedure} from '../models/Procedure';
import {EncounterService} from '../services/Encounter.service';
import {ClinicianService} from '../services/Clinician.service';
import {ProcedureOrderService} from '../services/ProcedureOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ProcedureService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	procedure : Procedure;

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
	// add a Procedure
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addProcedure(procedureCode, startDateTime, endDateTime, Encounter, Performer, ProcedureOrder, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Procedure/create';
		const obj = {
			      		procedureCode: procedureCode,
      		startDateTime: startDateTime,
      		endDateTime: endDateTime,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Performer: Performer != null && Performer.length > 0 ? Performer : null,
      		ProcedureOrder: ProcedureOrder != null && ProcedureOrder.length > 0 ? ProcedureOrder : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateProcedure(procedureCode, startDateTime, endDateTime, Encounter, Performer, ProcedureOrder, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Procedure/update/' + id;
		const obj = {
				      		procedureCode: procedureCode,
      		startDateTime: startDateTime,
      		endDateTime: endDateTime,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Performer: Performer != null && Performer.length > 0 ? Performer : null,
      		ProcedureOrder: ProcedureOrder != null && ProcedureOrder.length > 0 ? ProcedureOrder : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteProcedure(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Procedure/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Procedure
	// returns the results untouched as an Observable Procedure
	// Procedure model
	// delegates via URI
	//********************************************************************
	getProcedure(id) : Observable<Procedure> {
		const uri_ = this.apiUrl + '/Procedure/load/' + id;

		return this.http.get<Procedure>(uri_);
	}
	
	//********************************************************************
	// gets all Procedure
	// returns the results untouched as JSON representation of an
	// Observable array of Procedure models
	// delegates via URI
	//********************************************************************
	getProcedures() : Observable<Procedure[]> {
		const uri_ = this.apiUrl + '/Procedure/';

		return this
			.http.get<Procedure[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Encounter on a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEncounter( procedureId, _encounterId ): Observable<any> {

		// get the Procedure from storage
		this.loadHelper( procedureId );

	// get the Encounter from storage
	var tmp 	= new EncounterService(this.http).getEncounter(_encounterId);

	// assign the Encounter
	this.procedure.encounter = tmp;

	// save the Procedure
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Encounter on a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEncounter( procedureId ): Observable<any> {

		// get the Procedure from storage
		this.loadHelper( procedureId );

	// assign Encounter to null
	this.procedure.encounter = null;

	// save the Procedure
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Performer on a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPerformer( procedureId, _performerId ): Observable<any> {

		// get the Procedure from storage
		this.loadHelper( procedureId );

	// get the Clinician from storage
	var tmp 	= new ClinicianService(this.http).getClinician(_performerId);

	// assign the Performer
	this.procedure.performer = tmp;

	// save the Procedure
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Performer on a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPerformer( procedureId ): Observable<any> {

		// get the Procedure from storage
		this.loadHelper( procedureId );

	// assign Performer to null
	this.procedure.performer = null;

	// save the Procedure
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ProcedureOrder on a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProcedureOrder( procedureId, _procedureOrderId ): Observable<any> {

		// get the Procedure from storage
		this.loadHelper( procedureId );

	// get the ProcedureOrder from storage
	var tmp 	= new ProcedureOrderService(this.http).getProcedureOrder(_procedureOrderId);

	// assign the ProcedureOrder
	this.procedure.procedureOrder = tmp;

	// save the Procedure
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ProcedureOrder on a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProcedureOrder( procedureId ): Observable<any> {

		// get the Procedure from storage
		this.loadHelper( procedureId );

	// assign ProcedureOrder to null
	this.procedure.procedureOrder = null;

	// save the Procedure
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Procedure
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Procedure/update/' + this.procedure;

	return  this.http.post(uri_, this.procedure );
}

	//********************************************************************
	// loadHelper - internal helper to load a Procedure
	//********************************************************************	
	loadHelper( id ) {
		this.getProcedure(id)
			.subscribe((res : Procedure) => {
				this.procedure = res;
			});
	}
}