import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CorrectiveAction} from '../models/CorrectiveAction';
import {NonconformanceService} from '../services/Nonconformance.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CorrectiveActionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	correctiveAction : CorrectiveAction;

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
	// add a CorrectiveAction
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCorrectiveAction(capaNumber, rootCause, correctiveAction, verificationDate, Nonconformance, Owner, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/CorrectiveAction/create';
		const obj = {
			      		capaNumber: capaNumber,
      		rootCause: rootCause,
      		correctiveAction: correctiveAction,
      		verificationDate: verificationDate,
      		Nonconformance: Nonconformance != null && Nonconformance.length > 0 ? Nonconformance : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCorrectiveAction(capaNumber, rootCause, correctiveAction, verificationDate, Nonconformance, Owner, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CorrectiveAction/update/' + id;
		const obj = {
				      		capaNumber: capaNumber,
      		rootCause: rootCause,
      		correctiveAction: correctiveAction,
      		verificationDate: verificationDate,
      		Nonconformance: Nonconformance != null && Nonconformance.length > 0 ? Nonconformance : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCorrectiveAction(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CorrectiveAction/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CorrectiveAction
	// returns the results untouched as an Observable CorrectiveAction
	// CorrectiveAction model
	// delegates via URI
	//********************************************************************
	getCorrectiveAction(id) : Observable<CorrectiveAction> {
		const uri_ = this.apiUrl + '/CorrectiveAction/load/' + id;

		return this.http.get<CorrectiveAction>(uri_);
	}
	
	//********************************************************************
	// gets all CorrectiveAction
	// returns the results untouched as JSON representation of an
	// Observable array of CorrectiveAction models
	// delegates via URI
	//********************************************************************
	getCorrectiveActions() : Observable<CorrectiveAction[]> {
		const uri_ = this.apiUrl + '/CorrectiveAction/';

		return this
			.http.get<CorrectiveAction[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Nonconformance on a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignNonconformance( correctiveActionId, _nonconformanceId ): Observable<any> {

		// get the CorrectiveAction from storage
		this.loadHelper( correctiveActionId );

	// get the Nonconformance from storage
	var tmp 	= new NonconformanceService(this.http).getNonconformance(_nonconformanceId);

	// assign the Nonconformance
	this.correctiveAction.nonconformance = tmp;

	// save the CorrectiveAction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Nonconformance on a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignNonconformance( correctiveActionId ): Observable<any> {

		// get the CorrectiveAction from storage
		this.loadHelper( correctiveActionId );

	// assign Nonconformance to null
	this.correctiveAction.nonconformance = null;

	// save the CorrectiveAction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( correctiveActionId, _ownerId ): Observable<any> {

		// get the CorrectiveAction from storage
		this.loadHelper( correctiveActionId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_ownerId);

	// assign the Owner
	this.correctiveAction.owner = tmp;

	// save the CorrectiveAction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a CorrectiveAction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( correctiveActionId ): Observable<any> {

		// get the CorrectiveAction from storage
		this.loadHelper( correctiveActionId );

	// assign Owner to null
	this.correctiveAction.owner = null;

	// save the CorrectiveAction
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CorrectiveAction
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CorrectiveAction/update/' + this.correctiveAction;

	return  this.http.post(uri_, this.correctiveAction );
}

	//********************************************************************
	// loadHelper - internal helper to load a CorrectiveAction
	//********************************************************************	
	loadHelper( id ) {
		this.getCorrectiveAction(id)
			.subscribe((res : CorrectiveAction) => {
				this.correctiveAction = res;
			});
	}
}