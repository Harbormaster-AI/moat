import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Approval} from '../models/Approval';
import {EmployeeService} from '../services/Employee.service';
import {TimesheetService} from '../services/Timesheet.service';
import {LeaveRequestService} from '../services/LeaveRequest.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ApprovalService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	approval : Approval;

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
	// add a Approval
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addApproval(approverComment, actionDate, Approver, Timesheet, LeaveRequest, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Approval/create';
		const obj = {
			      		approverComment: approverComment,
      		actionDate: actionDate,
      		Approver: Approver != null && Approver.length > 0 ? Approver : null,
      		Timesheet: Timesheet != null && Timesheet.length > 0 ? Timesheet : null,
      		LeaveRequest: LeaveRequest != null && LeaveRequest.length > 0 ? LeaveRequest : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Approval
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateApproval(approverComment, actionDate, Approver, Timesheet, LeaveRequest, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Approval/update/' + id;
		const obj = {
				      		approverComment: approverComment,
      		actionDate: actionDate,
      		Approver: Approver != null && Approver.length > 0 ? Approver : null,
      		Timesheet: Timesheet != null && Timesheet.length > 0 ? Timesheet : null,
      		LeaveRequest: LeaveRequest != null && LeaveRequest.length > 0 ? LeaveRequest : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Approval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteApproval(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Approval/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Approval
	// returns the results untouched as an Observable Approval
	// Approval model
	// delegates via URI
	//********************************************************************
	getApproval(id) : Observable<Approval> {
		const uri_ = this.apiUrl + '/Approval/load/' + id;

		return this.http.get<Approval>(uri_);
	}
	
	//********************************************************************
	// gets all Approval
	// returns the results untouched as JSON representation of an
	// Observable array of Approval models
	// delegates via URI
	//********************************************************************
	getApprovals() : Observable<Approval[]> {
		const uri_ = this.apiUrl + '/Approval/';

		return this
			.http.get<Approval[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Approver on a Approval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignApprover( approvalId, _approverId ): Observable<any> {

		// get the Approval from storage
		this.loadHelper( approvalId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_approverId);

	// assign the Approver
	this.approval.approver = tmp;

	// save the Approval
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Approver on a Approval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignApprover( approvalId ): Observable<any> {

		// get the Approval from storage
		this.loadHelper( approvalId );

	// assign Approver to null
	this.approval.approver = null;

	// save the Approval
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Timesheet on a Approval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTimesheet( approvalId, _timesheetId ): Observable<any> {

		// get the Approval from storage
		this.loadHelper( approvalId );

	// get the Timesheet from storage
	var tmp 	= new TimesheetService(this.http).getTimesheet(_timesheetId);

	// assign the Timesheet
	this.approval.timesheet = tmp;

	// save the Approval
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Timesheet on a Approval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTimesheet( approvalId ): Observable<any> {

		// get the Approval from storage
		this.loadHelper( approvalId );

	// assign Timesheet to null
	this.approval.timesheet = null;

	// save the Approval
	return this.saveHelper();
}

		//********************************************************************
	// assigns a LeaveRequest on a Approval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLeaveRequest( approvalId, _leaveRequestId ): Observable<any> {

		// get the Approval from storage
		this.loadHelper( approvalId );

	// get the LeaveRequest from storage
	var tmp 	= new LeaveRequestService(this.http).getLeaveRequest(_leaveRequestId);

	// assign the LeaveRequest
	this.approval.leaveRequest = tmp;

	// save the Approval
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LeaveRequest on a Approval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLeaveRequest( approvalId ): Observable<any> {

		// get the Approval from storage
		this.loadHelper( approvalId );

	// assign LeaveRequest to null
	this.approval.leaveRequest = null;

	// save the Approval
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Approval
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Approval/update/' + this.approval;

	return  this.http.post(uri_, this.approval );
}

	//********************************************************************
	// loadHelper - internal helper to load a Approval
	//********************************************************************	
	loadHelper( id ) {
		this.getApproval(id)
			.subscribe((res : Approval) => {
				this.approval = res;
			});
	}
}