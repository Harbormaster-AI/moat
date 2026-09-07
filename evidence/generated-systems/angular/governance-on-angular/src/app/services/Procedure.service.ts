import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Procedure} from '../models/Procedure';
import {PolicyService} from '../services/Policy.service';
import {ControlService} from '../services/Control.service';
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
	addProcedure(title, versionLabel, Policy, Controls, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Procedure/create';
		const obj = {
			      		title: title,
      		versionLabel: versionLabel,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateProcedure(title, versionLabel, Policy, Controls, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Procedure/update/' + id;
		const obj = {
				      		title: title,
      		versionLabel: versionLabel,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
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
	// assigns a Policy on a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( procedureId, _policyId ): Observable<any> {

		// get the Procedure from storage
		this.loadHelper( procedureId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.procedure.policy = tmp;

	// save the Procedure
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Procedure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( procedureId ): Observable<any> {

		// get the Procedure from storage
		this.loadHelper( procedureId );

	// assign Policy to null
	this.procedure.policy = null;

	// save the Procedure
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more controlsIds as a Controls
	// to a Procedure
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addControls( procedureId, controlsIds ): Observable<any> {

		// get the Procedure
		this.loadHelper( procedureId );

	// split on a comma with no spaces
	var idList = controlsIds.split(',')

	// iterate over array of controls ids
	idList.forEach(function (id) {
		// read the Control
		var control = new ControlService(this.http).getControl(id);
		// add the Control if not already assigned
		if ( this.procedure.controls.indexOf(control) == -1 )
		this.procedure.controls.push(control);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more controlsIds as a Controls
	// from a Procedure
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeControls( procedureId, controlsIds ): Observable<any> {

		// get the Procedure
		this.loadHelper( procedureId );


	// split on a comma with no spaces
	var idList 					= controlsIds.split(',');
	var controls 	= this.procedure.controls;

	if ( controls != null && controlsIds != null ) {

		// iterate over array of controls ids
		controls.forEach(function (obj) {
			if ( controlsIds.indexOf(obj._id) > -1 ) {
				// remove the Control
				this.procedure.controls.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
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