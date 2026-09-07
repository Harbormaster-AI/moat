import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ScheduleException} from '../models/ScheduleException';
import {WorkScheduleService} from '../services/WorkSchedule.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ScheduleExceptionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	scheduleException : ScheduleException;

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
	// add a ScheduleException
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addScheduleException(date, reason, hours, WorkSchedule, Employee) : Observable<any> {
		const uri_ = this.apiUrl + '/ScheduleException/create';
		const obj = {
			      		date: date,
      		reason: reason,
      		hours: hours,
      		WorkSchedule: WorkSchedule != null && WorkSchedule.length > 0 ? WorkSchedule : null,
			Employee: Employee != null && Employee.length > 0 ? Employee : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ScheduleException
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateScheduleException(date, reason, hours, WorkSchedule, Employee, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ScheduleException/update/' + id;
		const obj = {
				      		date: date,
      		reason: reason,
      		hours: hours,
      		WorkSchedule: WorkSchedule != null && WorkSchedule.length > 0 ? WorkSchedule : null,
			Employee: Employee != null && Employee.length > 0 ? Employee : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ScheduleException
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteScheduleException(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ScheduleException/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ScheduleException
	// returns the results untouched as an Observable ScheduleException
	// ScheduleException model
	// delegates via URI
	//********************************************************************
	getScheduleException(id) : Observable<ScheduleException> {
		const uri_ = this.apiUrl + '/ScheduleException/load/' + id;

		return this.http.get<ScheduleException>(uri_);
	}
	
	//********************************************************************
	// gets all ScheduleException
	// returns the results untouched as JSON representation of an
	// Observable array of ScheduleException models
	// delegates via URI
	//********************************************************************
	getScheduleExceptions() : Observable<ScheduleException[]> {
		const uri_ = this.apiUrl + '/ScheduleException/';

		return this
			.http.get<ScheduleException[]>(uri_);
	}
	
			//********************************************************************
	// assigns a WorkSchedule on a ScheduleException
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkSchedule( scheduleExceptionId, _workScheduleId ): Observable<any> {

		// get the ScheduleException from storage
		this.loadHelper( scheduleExceptionId );

	// get the WorkSchedule from storage
	var tmp 	= new WorkScheduleService(this.http).getWorkSchedule(_workScheduleId);

	// assign the WorkSchedule
	this.scheduleException.workSchedule = tmp;

	// save the ScheduleException
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkSchedule on a ScheduleException
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkSchedule( scheduleExceptionId ): Observable<any> {

		// get the ScheduleException from storage
		this.loadHelper( scheduleExceptionId );

	// assign WorkSchedule to null
	this.scheduleException.workSchedule = null;

	// save the ScheduleException
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Employee on a ScheduleException
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( scheduleExceptionId, _employeeId ): Observable<any> {

		// get the ScheduleException from storage
		this.loadHelper( scheduleExceptionId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.scheduleException.employee = tmp;

	// save the ScheduleException
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a ScheduleException
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( scheduleExceptionId ): Observable<any> {

		// get the ScheduleException from storage
		this.loadHelper( scheduleExceptionId );

	// assign Employee to null
	this.scheduleException.employee = null;

	// save the ScheduleException
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ScheduleException
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ScheduleException/update/' + this.scheduleException;

	return  this.http.post(uri_, this.scheduleException );
}

	//********************************************************************
	// loadHelper - internal helper to load a ScheduleException
	//********************************************************************	
	loadHelper( id ) {
		this.getScheduleException(id)
			.subscribe((res : ScheduleException) => {
				this.scheduleException = res;
			});
	}
}