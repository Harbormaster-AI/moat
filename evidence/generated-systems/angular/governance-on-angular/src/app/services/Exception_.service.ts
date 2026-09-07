import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Exception_} from '../models/Exception_';
import {RetentionScheduleService} from '../services/RetentionSchedule.service';
import {PolicyService} from '../services/Policy.service';
import {ControlService} from '../services/Control.service';
import {RiskService} from '../services/Risk.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class Exception_Service extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	exception_ : Exception_;

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
	// add a Exception_
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addException_(title, justification, startDate, endDate, RetentionSchedule, Policy, Control, Risk, ExceptionType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Exception_/create';
		const obj = {
			      		title: title,
      		justification: justification,
      		startDate: startDate,
      		endDate: endDate,
      		RetentionSchedule: RetentionSchedule != null && RetentionSchedule.length > 0 ? RetentionSchedule : null,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Control: Control != null && Control.length > 0 ? Control : null,
      		Risk: Risk != null && Risk.length > 0 ? Risk : null,
      		ExceptionType: ExceptionType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Exception_
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateException_(title, justification, startDate, endDate, RetentionSchedule, Policy, Control, Risk, ExceptionType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Exception_/update/' + id;
		const obj = {
				      		title: title,
      		justification: justification,
      		startDate: startDate,
      		endDate: endDate,
      		RetentionSchedule: RetentionSchedule != null && RetentionSchedule.length > 0 ? RetentionSchedule : null,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Control: Control != null && Control.length > 0 ? Control : null,
      		Risk: Risk != null && Risk.length > 0 ? Risk : null,
      		ExceptionType: ExceptionType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Exception_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteException_(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Exception_/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Exception_
	// returns the results untouched as an Observable Exception_
	// Exception_ model
	// delegates via URI
	//********************************************************************
	getException_(id) : Observable<Exception_> {
		const uri_ = this.apiUrl + '/Exception_/load/' + id;

		return this.http.get<Exception_>(uri_);
	}
	
	//********************************************************************
	// gets all Exception_
	// returns the results untouched as JSON representation of an
	// Observable array of Exception_ models
	// delegates via URI
	//********************************************************************
	getException_s() : Observable<Exception_[]> {
		const uri_ = this.apiUrl + '/Exception_/';

		return this
			.http.get<Exception_[]>(uri_);
	}
	
			//********************************************************************
	// assigns a RetentionSchedule on a Exception_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRetentionSchedule( exception_Id, _retentionScheduleId ): Observable<any> {

		// get the Exception_ from storage
		this.loadHelper( exception_Id );

	// get the RetentionSchedule from storage
	var tmp 	= new RetentionScheduleService(this.http).getRetentionSchedule(_retentionScheduleId);

	// assign the RetentionSchedule
	this.exception_.retentionSchedule = tmp;

	// save the Exception_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a RetentionSchedule on a Exception_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRetentionSchedule( exception_Id ): Observable<any> {

		// get the Exception_ from storage
		this.loadHelper( exception_Id );

	// assign RetentionSchedule to null
	this.exception_.retentionSchedule = null;

	// save the Exception_
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Policy on a Exception_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( exception_Id, _policyId ): Observable<any> {

		// get the Exception_ from storage
		this.loadHelper( exception_Id );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.exception_.policy = tmp;

	// save the Exception_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Exception_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( exception_Id ): Observable<any> {

		// get the Exception_ from storage
		this.loadHelper( exception_Id );

	// assign Policy to null
	this.exception_.policy = null;

	// save the Exception_
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Control on a Exception_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignControl( exception_Id, _controlId ): Observable<any> {

		// get the Exception_ from storage
		this.loadHelper( exception_Id );

	// get the Control from storage
	var tmp 	= new ControlService(this.http).getControl(_controlId);

	// assign the Control
	this.exception_.control = tmp;

	// save the Exception_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Control on a Exception_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignControl( exception_Id ): Observable<any> {

		// get the Exception_ from storage
		this.loadHelper( exception_Id );

	// assign Control to null
	this.exception_.control = null;

	// save the Exception_
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Risk on a Exception_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRisk( exception_Id, _riskId ): Observable<any> {

		// get the Exception_ from storage
		this.loadHelper( exception_Id );

	// get the Risk from storage
	var tmp 	= new RiskService(this.http).getRisk(_riskId);

	// assign the Risk
	this.exception_.risk = tmp;

	// save the Exception_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Risk on a Exception_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRisk( exception_Id ): Observable<any> {

		// get the Exception_ from storage
		this.loadHelper( exception_Id );

	// assign Risk to null
	this.exception_.risk = null;

	// save the Exception_
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Exception_
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Exception_/update/' + this.exception_;

	return  this.http.post(uri_, this.exception_ );
}

	//********************************************************************
	// loadHelper - internal helper to load a Exception_
	//********************************************************************	
	loadHelper( id ) {
		this.getException_(id)
			.subscribe((res : Exception_) => {
				this.exception_ = res;
			});
	}
}