import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TimeEntry} from '../models/TimeEntry';
import {TimesheetService} from '../services/Timesheet.service';
import {EmployeeService} from '../services/Employee.service';
import {CostCenterService} from '../services/CostCenter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TimeEntryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	timeEntry : TimeEntry;

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
	// add a TimeEntry
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTimeEntry(entryDate, hoursWorked, Timesheet, Employee, CostCenter, EntryType) : Observable<any> {
		const uri_ = this.apiUrl + '/TimeEntry/create';
		const obj = {
			      		entryDate: entryDate,
      		hoursWorked: hoursWorked,
      		Timesheet: Timesheet != null && Timesheet.length > 0 ? Timesheet : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		CostCenter: CostCenter != null && CostCenter.length > 0 ? CostCenter : null,
			EntryType: EntryType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TimeEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTimeEntry(entryDate, hoursWorked, Timesheet, Employee, CostCenter, EntryType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TimeEntry/update/' + id;
		const obj = {
				      		entryDate: entryDate,
      		hoursWorked: hoursWorked,
      		Timesheet: Timesheet != null && Timesheet.length > 0 ? Timesheet : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		CostCenter: CostCenter != null && CostCenter.length > 0 ? CostCenter : null,
			EntryType: EntryType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TimeEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTimeEntry(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TimeEntry/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TimeEntry
	// returns the results untouched as an Observable TimeEntry
	// TimeEntry model
	// delegates via URI
	//********************************************************************
	getTimeEntry(id) : Observable<TimeEntry> {
		const uri_ = this.apiUrl + '/TimeEntry/load/' + id;

		return this.http.get<TimeEntry>(uri_);
	}
	
	//********************************************************************
	// gets all TimeEntry
	// returns the results untouched as JSON representation of an
	// Observable array of TimeEntry models
	// delegates via URI
	//********************************************************************
	getTimeEntrys() : Observable<TimeEntry[]> {
		const uri_ = this.apiUrl + '/TimeEntry/';

		return this
			.http.get<TimeEntry[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Timesheet on a TimeEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTimesheet( timeEntryId, _timesheetId ): Observable<any> {

		// get the TimeEntry from storage
		this.loadHelper( timeEntryId );

	// get the Timesheet from storage
	var tmp 	= new TimesheetService(this.http).getTimesheet(_timesheetId);

	// assign the Timesheet
	this.timeEntry.timesheet = tmp;

	// save the TimeEntry
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Timesheet on a TimeEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTimesheet( timeEntryId ): Observable<any> {

		// get the TimeEntry from storage
		this.loadHelper( timeEntryId );

	// assign Timesheet to null
	this.timeEntry.timesheet = null;

	// save the TimeEntry
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Employee on a TimeEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( timeEntryId, _employeeId ): Observable<any> {

		// get the TimeEntry from storage
		this.loadHelper( timeEntryId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.timeEntry.employee = tmp;

	// save the TimeEntry
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a TimeEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( timeEntryId ): Observable<any> {

		// get the TimeEntry from storage
		this.loadHelper( timeEntryId );

	// assign Employee to null
	this.timeEntry.employee = null;

	// save the TimeEntry
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CostCenter on a TimeEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCostCenter( timeEntryId, _costCenterId ): Observable<any> {

		// get the TimeEntry from storage
		this.loadHelper( timeEntryId );

	// get the CostCenter from storage
	var tmp 	= new CostCenterService(this.http).getCostCenter(_costCenterId);

	// assign the CostCenter
	this.timeEntry.costCenter = tmp;

	// save the TimeEntry
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CostCenter on a TimeEntry
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCostCenter( timeEntryId ): Observable<any> {

		// get the TimeEntry from storage
		this.loadHelper( timeEntryId );

	// assign CostCenter to null
	this.timeEntry.costCenter = null;

	// save the TimeEntry
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a TimeEntry
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TimeEntry/update/' + this.timeEntry;

	return  this.http.post(uri_, this.timeEntry );
}

	//********************************************************************
	// loadHelper - internal helper to load a TimeEntry
	//********************************************************************	
	loadHelper( id ) {
		this.getTimeEntry(id)
			.subscribe((res : TimeEntry) => {
				this.timeEntry = res;
			});
	}
}