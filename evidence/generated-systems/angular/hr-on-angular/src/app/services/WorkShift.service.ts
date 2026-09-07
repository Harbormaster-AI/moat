import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {WorkShift} from '../models/WorkShift';
import {WorkScheduleService} from '../services/WorkSchedule.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class WorkShiftService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	workShift : WorkShift;

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
	// add a WorkShift
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addWorkShift(startTime, endTime, breakMinutes, WorkSchedule, DayOfWeek) : Observable<any> {
		const uri_ = this.apiUrl + '/WorkShift/create';
		const obj = {
			      		startTime: startTime,
      		endTime: endTime,
      		breakMinutes: breakMinutes,
      		WorkSchedule: WorkSchedule != null && WorkSchedule.length > 0 ? WorkSchedule : null,
			DayOfWeek: DayOfWeek
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a WorkShift
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateWorkShift(startTime, endTime, breakMinutes, WorkSchedule, DayOfWeek, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/WorkShift/update/' + id;
		const obj = {
				      		startTime: startTime,
      		endTime: endTime,
      		breakMinutes: breakMinutes,
      		WorkSchedule: WorkSchedule != null && WorkSchedule.length > 0 ? WorkSchedule : null,
			DayOfWeek: DayOfWeek
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a WorkShift
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteWorkShift(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/WorkShift/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a WorkShift
	// returns the results untouched as an Observable WorkShift
	// WorkShift model
	// delegates via URI
	//********************************************************************
	getWorkShift(id) : Observable<WorkShift> {
		const uri_ = this.apiUrl + '/WorkShift/load/' + id;

		return this.http.get<WorkShift>(uri_);
	}
	
	//********************************************************************
	// gets all WorkShift
	// returns the results untouched as JSON representation of an
	// Observable array of WorkShift models
	// delegates via URI
	//********************************************************************
	getWorkShifts() : Observable<WorkShift[]> {
		const uri_ = this.apiUrl + '/WorkShift/';

		return this
			.http.get<WorkShift[]>(uri_);
	}
	
			//********************************************************************
	// assigns a WorkSchedule on a WorkShift
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkSchedule( workShiftId, _workScheduleId ): Observable<any> {

		// get the WorkShift from storage
		this.loadHelper( workShiftId );

	// get the WorkSchedule from storage
	var tmp 	= new WorkScheduleService(this.http).getWorkSchedule(_workScheduleId);

	// assign the WorkSchedule
	this.workShift.workSchedule = tmp;

	// save the WorkShift
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkSchedule on a WorkShift
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkSchedule( workShiftId ): Observable<any> {

		// get the WorkShift from storage
		this.loadHelper( workShiftId );

	// assign WorkSchedule to null
	this.workShift.workSchedule = null;

	// save the WorkShift
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a WorkShift
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/WorkShift/update/' + this.workShift;

	return  this.http.post(uri_, this.workShift );
}

	//********************************************************************
	// loadHelper - internal helper to load a WorkShift
	//********************************************************************	
	loadHelper( id ) {
		this.getWorkShift(id)
			.subscribe((res : WorkShift) => {
				this.workShift = res;
			});
	}
}