import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {WorkSchedule} from '../models/WorkSchedule';
import {EmploymentContractService} from '../services/EmploymentContract.service';
import {WorkShiftService} from '../services/WorkShift.service';
import {ScheduleExceptionService} from '../services/ScheduleException.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class WorkScheduleService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	workSchedule : WorkSchedule;

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
	// add a WorkSchedule
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addWorkSchedule(name, standardHoursPerWeek, Contracts, Shifts, Exceptions, ScheduleType) : Observable<any> {
		const uri_ = this.apiUrl + '/WorkSchedule/create';
		const obj = {
			      		name: name,
      		standardHoursPerWeek: standardHoursPerWeek,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		Shifts: Shifts != null && Shifts.length > 0 ? Shifts : null,
      		Exceptions: Exceptions != null && Exceptions.length > 0 ? Exceptions : null,
			ScheduleType: ScheduleType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a WorkSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateWorkSchedule(name, standardHoursPerWeek, Contracts, Shifts, Exceptions, ScheduleType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/WorkSchedule/update/' + id;
		const obj = {
				      		name: name,
      		standardHoursPerWeek: standardHoursPerWeek,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		Shifts: Shifts != null && Shifts.length > 0 ? Shifts : null,
      		Exceptions: Exceptions != null && Exceptions.length > 0 ? Exceptions : null,
			ScheduleType: ScheduleType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a WorkSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteWorkSchedule(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/WorkSchedule/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a WorkSchedule
	// returns the results untouched as an Observable WorkSchedule
	// WorkSchedule model
	// delegates via URI
	//********************************************************************
	getWorkSchedule(id) : Observable<WorkSchedule> {
		const uri_ = this.apiUrl + '/WorkSchedule/load/' + id;

		return this.http.get<WorkSchedule>(uri_);
	}
	
	//********************************************************************
	// gets all WorkSchedule
	// returns the results untouched as JSON representation of an
	// Observable array of WorkSchedule models
	// delegates via URI
	//********************************************************************
	getWorkSchedules() : Observable<WorkSchedule[]> {
		const uri_ = this.apiUrl + '/WorkSchedule/';

		return this
			.http.get<WorkSchedule[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more contractsIds as a Contracts
	// to a WorkSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContracts( workScheduleId, contractsIds ): Observable<any> {

		// get the WorkSchedule
		this.loadHelper( workScheduleId );

	// split on a comma with no spaces
	var idList = contractsIds.split(',')

	// iterate over array of contracts ids
	idList.forEach(function (id) {
		// read the EmploymentContract
		var employmentContract = new EmploymentContractService(this.http).getEmploymentContract(id);
		// add the EmploymentContract if not already assigned
		if ( this.workSchedule.contracts.indexOf(employmentContract) == -1 )
		this.workSchedule.contracts.push(employmentContract);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contractsIds as a Contracts
	// from a WorkSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContracts( workScheduleId, contractsIds ): Observable<any> {

		// get the WorkSchedule
		this.loadHelper( workScheduleId );


	// split on a comma with no spaces
	var idList 					= contractsIds.split(',');
	var contracts 	= this.workSchedule.contracts;

	if ( contracts != null && contractsIds != null ) {

		// iterate over array of contracts ids
		contracts.forEach(function (obj) {
			if ( contractsIds.indexOf(obj._id) > -1 ) {
				// remove the EmploymentContract
				this.workSchedule.contracts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more shiftsIds as a Shifts
	// to a WorkSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addShifts( workScheduleId, shiftsIds ): Observable<any> {

		// get the WorkSchedule
		this.loadHelper( workScheduleId );

	// split on a comma with no spaces
	var idList = shiftsIds.split(',')

	// iterate over array of shifts ids
	idList.forEach(function (id) {
		// read the WorkShift
		var workShift = new WorkShiftService(this.http).getWorkShift(id);
		// add the WorkShift if not already assigned
		if ( this.workSchedule.shifts.indexOf(workShift) == -1 )
		this.workSchedule.shifts.push(workShift);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more shiftsIds as a Shifts
	// from a WorkSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeShifts( workScheduleId, shiftsIds ): Observable<any> {

		// get the WorkSchedule
		this.loadHelper( workScheduleId );


	// split on a comma with no spaces
	var idList 					= shiftsIds.split(',');
	var shifts 	= this.workSchedule.shifts;

	if ( shifts != null && shiftsIds != null ) {

		// iterate over array of shifts ids
		shifts.forEach(function (obj) {
			if ( shiftsIds.indexOf(obj._id) > -1 ) {
				// remove the WorkShift
				this.workSchedule.shifts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more exceptionsIds as a Exceptions
	// to a WorkSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addExceptions( workScheduleId, exceptionsIds ): Observable<any> {

		// get the WorkSchedule
		this.loadHelper( workScheduleId );

	// split on a comma with no spaces
	var idList = exceptionsIds.split(',')

	// iterate over array of exceptions ids
	idList.forEach(function (id) {
		// read the ScheduleException
		var scheduleException = new ScheduleExceptionService(this.http).getScheduleException(id);
		// add the ScheduleException if not already assigned
		if ( this.workSchedule.exceptions.indexOf(scheduleException) == -1 )
		this.workSchedule.exceptions.push(scheduleException);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more exceptionsIds as a Exceptions
	// from a WorkSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeExceptions( workScheduleId, exceptionsIds ): Observable<any> {

		// get the WorkSchedule
		this.loadHelper( workScheduleId );


	// split on a comma with no spaces
	var idList 					= exceptionsIds.split(',');
	var exceptions 	= this.workSchedule.exceptions;

	if ( exceptions != null && exceptionsIds != null ) {

		// iterate over array of exceptions ids
		exceptions.forEach(function (obj) {
			if ( exceptionsIds.indexOf(obj._id) > -1 ) {
				// remove the ScheduleException
				this.workSchedule.exceptions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a WorkSchedule
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/WorkSchedule/update/' + this.workSchedule;

	return  this.http.post(uri_, this.workSchedule );
}

	//********************************************************************
	// loadHelper - internal helper to load a WorkSchedule
	//********************************************************************	
	loadHelper( id ) {
		this.getWorkSchedule(id)
			.subscribe((res : WorkSchedule) => {
				this.workSchedule = res;
			});
	}
}