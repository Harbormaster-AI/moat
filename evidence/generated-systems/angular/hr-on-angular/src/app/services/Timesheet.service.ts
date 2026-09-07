import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Timesheet} from '../models/Timesheet';
import {EmployeeService} from '../services/Employee.service';
import {TimeEntryService} from '../services/TimeEntry.service';
import {ApprovalService} from '../services/Approval.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TimesheetService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	timesheet : Timesheet;

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
	// add a Timesheet
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTimesheet(periodStart, periodEnd, submissionDate, Employee, TimeEntries, Approvals, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Timesheet/create';
		const obj = {
			      		periodStart: periodStart,
      		periodEnd: periodEnd,
      		submissionDate: submissionDate,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		TimeEntries: TimeEntries != null && TimeEntries.length > 0 ? TimeEntries : null,
      		Approvals: Approvals != null && Approvals.length > 0 ? Approvals : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Timesheet
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTimesheet(periodStart, periodEnd, submissionDate, Employee, TimeEntries, Approvals, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Timesheet/update/' + id;
		const obj = {
				      		periodStart: periodStart,
      		periodEnd: periodEnd,
      		submissionDate: submissionDate,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		TimeEntries: TimeEntries != null && TimeEntries.length > 0 ? TimeEntries : null,
      		Approvals: Approvals != null && Approvals.length > 0 ? Approvals : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Timesheet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTimesheet(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Timesheet/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Timesheet
	// returns the results untouched as an Observable Timesheet
	// Timesheet model
	// delegates via URI
	//********************************************************************
	getTimesheet(id) : Observable<Timesheet> {
		const uri_ = this.apiUrl + '/Timesheet/load/' + id;

		return this.http.get<Timesheet>(uri_);
	}
	
	//********************************************************************
	// gets all Timesheet
	// returns the results untouched as JSON representation of an
	// Observable array of Timesheet models
	// delegates via URI
	//********************************************************************
	getTimesheets() : Observable<Timesheet[]> {
		const uri_ = this.apiUrl + '/Timesheet/';

		return this
			.http.get<Timesheet[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a Timesheet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( timesheetId, _employeeId ): Observable<any> {

		// get the Timesheet from storage
		this.loadHelper( timesheetId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.timesheet.employee = tmp;

	// save the Timesheet
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a Timesheet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( timesheetId ): Observable<any> {

		// get the Timesheet from storage
		this.loadHelper( timesheetId );

	// assign Employee to null
	this.timesheet.employee = null;

	// save the Timesheet
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more timeEntriesIds as a TimeEntries
	// to a Timesheet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTimeEntries( timesheetId, timeEntriesIds ): Observable<any> {

		// get the Timesheet
		this.loadHelper( timesheetId );

	// split on a comma with no spaces
	var idList = timeEntriesIds.split(',')

	// iterate over array of timeEntries ids
	idList.forEach(function (id) {
		// read the TimeEntry
		var timeEntry = new TimeEntryService(this.http).getTimeEntry(id);
		// add the TimeEntry if not already assigned
		if ( this.timesheet.timeEntries.indexOf(timeEntry) == -1 )
		this.timesheet.timeEntries.push(timeEntry);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more timeEntriesIds as a TimeEntries
	// from a Timesheet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTimeEntries( timesheetId, timeEntriesIds ): Observable<any> {

		// get the Timesheet
		this.loadHelper( timesheetId );


	// split on a comma with no spaces
	var idList 					= timeEntriesIds.split(',');
	var timeEntries 	= this.timesheet.timeEntries;

	if ( timeEntries != null && timeEntriesIds != null ) {

		// iterate over array of timeEntries ids
		timeEntries.forEach(function (obj) {
			if ( timeEntriesIds.indexOf(obj._id) > -1 ) {
				// remove the TimeEntry
				this.timesheet.timeEntries.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more approvalsIds as a Approvals
	// to a Timesheet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addApprovals( timesheetId, approvalsIds ): Observable<any> {

		// get the Timesheet
		this.loadHelper( timesheetId );

	// split on a comma with no spaces
	var idList = approvalsIds.split(',')

	// iterate over array of approvals ids
	idList.forEach(function (id) {
		// read the Approval
		var approval = new ApprovalService(this.http).getApproval(id);
		// add the Approval if not already assigned
		if ( this.timesheet.approvals.indexOf(approval) == -1 )
		this.timesheet.approvals.push(approval);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more approvalsIds as a Approvals
	// from a Timesheet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeApprovals( timesheetId, approvalsIds ): Observable<any> {

		// get the Timesheet
		this.loadHelper( timesheetId );


	// split on a comma with no spaces
	var idList 					= approvalsIds.split(',');
	var approvals 	= this.timesheet.approvals;

	if ( approvals != null && approvalsIds != null ) {

		// iterate over array of approvals ids
		approvals.forEach(function (obj) {
			if ( approvalsIds.indexOf(obj._id) > -1 ) {
				// remove the Approval
				this.timesheet.approvals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Timesheet
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Timesheet/update/' + this.timesheet;

	return  this.http.post(uri_, this.timesheet );
}

	//********************************************************************
	// loadHelper - internal helper to load a Timesheet
	//********************************************************************	
	loadHelper( id ) {
		this.getTimesheet(id)
			.subscribe((res : Timesheet) => {
				this.timesheet = res;
			});
	}
}