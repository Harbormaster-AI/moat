import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {LeaveRequest} from '../models/LeaveRequest';
import {EmployeeService} from '../services/Employee.service';
import {LeavePolicyService} from '../services/LeavePolicy.service';
import {ApprovalService} from '../services/Approval.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LeaveRequestService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	leaveRequest : LeaveRequest;

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
	// add a LeaveRequest
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLeaveRequest(requestNumber, startDate, endDate, reason, hours, Employee, LeavePolicy, Approvals, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/LeaveRequest/create';
		const obj = {
			      		requestNumber: requestNumber,
      		startDate: startDate,
      		endDate: endDate,
      		reason: reason,
      		hours: hours,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		LeavePolicy: LeavePolicy != null && LeavePolicy.length > 0 ? LeavePolicy : null,
      		Approvals: Approvals != null && Approvals.length > 0 ? Approvals : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a LeaveRequest
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLeaveRequest(requestNumber, startDate, endDate, reason, hours, Employee, LeavePolicy, Approvals, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/LeaveRequest/update/' + id;
		const obj = {
				      		requestNumber: requestNumber,
      		startDate: startDate,
      		endDate: endDate,
      		reason: reason,
      		hours: hours,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		LeavePolicy: LeavePolicy != null && LeavePolicy.length > 0 ? LeavePolicy : null,
      		Approvals: Approvals != null && Approvals.length > 0 ? Approvals : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a LeaveRequest
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLeaveRequest(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/LeaveRequest/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a LeaveRequest
	// returns the results untouched as an Observable LeaveRequest
	// LeaveRequest model
	// delegates via URI
	//********************************************************************
	getLeaveRequest(id) : Observable<LeaveRequest> {
		const uri_ = this.apiUrl + '/LeaveRequest/load/' + id;

		return this.http.get<LeaveRequest>(uri_);
	}
	
	//********************************************************************
	// gets all LeaveRequest
	// returns the results untouched as JSON representation of an
	// Observable array of LeaveRequest models
	// delegates via URI
	//********************************************************************
	getLeaveRequests() : Observable<LeaveRequest[]> {
		const uri_ = this.apiUrl + '/LeaveRequest/';

		return this
			.http.get<LeaveRequest[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a LeaveRequest
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( leaveRequestId, _employeeId ): Observable<any> {

		// get the LeaveRequest from storage
		this.loadHelper( leaveRequestId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.leaveRequest.employee = tmp;

	// save the LeaveRequest
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a LeaveRequest
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( leaveRequestId ): Observable<any> {

		// get the LeaveRequest from storage
		this.loadHelper( leaveRequestId );

	// assign Employee to null
	this.leaveRequest.employee = null;

	// save the LeaveRequest
	return this.saveHelper();
}

		//********************************************************************
	// assigns a LeavePolicy on a LeaveRequest
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLeavePolicy( leaveRequestId, _leavePolicyId ): Observable<any> {

		// get the LeaveRequest from storage
		this.loadHelper( leaveRequestId );

	// get the LeavePolicy from storage
	var tmp 	= new LeavePolicyService(this.http).getLeavePolicy(_leavePolicyId);

	// assign the LeavePolicy
	this.leaveRequest.leavePolicy = tmp;

	// save the LeaveRequest
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LeavePolicy on a LeaveRequest
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLeavePolicy( leaveRequestId ): Observable<any> {

		// get the LeaveRequest from storage
		this.loadHelper( leaveRequestId );

	// assign LeavePolicy to null
	this.leaveRequest.leavePolicy = null;

	// save the LeaveRequest
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more approvalsIds as a Approvals
	// to a LeaveRequest
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addApprovals( leaveRequestId, approvalsIds ): Observable<any> {

		// get the LeaveRequest
		this.loadHelper( leaveRequestId );

	// split on a comma with no spaces
	var idList = approvalsIds.split(',')

	// iterate over array of approvals ids
	idList.forEach(function (id) {
		// read the Approval
		var approval = new ApprovalService(this.http).getApproval(id);
		// add the Approval if not already assigned
		if ( this.leaveRequest.approvals.indexOf(approval) == -1 )
		this.leaveRequest.approvals.push(approval);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more approvalsIds as a Approvals
	// from a LeaveRequest
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeApprovals( leaveRequestId, approvalsIds ): Observable<any> {

		// get the LeaveRequest
		this.loadHelper( leaveRequestId );


	// split on a comma with no spaces
	var idList 					= approvalsIds.split(',');
	var approvals 	= this.leaveRequest.approvals;

	if ( approvals != null && approvalsIds != null ) {

		// iterate over array of approvals ids
		approvals.forEach(function (obj) {
			if ( approvalsIds.indexOf(obj._id) > -1 ) {
				// remove the Approval
				this.leaveRequest.approvals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a LeaveRequest
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/LeaveRequest/update/' + this.leaveRequest;

	return  this.http.post(uri_, this.leaveRequest );
}

	//********************************************************************
	// loadHelper - internal helper to load a LeaveRequest
	//********************************************************************	
	loadHelper( id ) {
		this.getLeaveRequest(id)
			.subscribe((res : LeaveRequest) => {
				this.leaveRequest = res;
			});
	}
}