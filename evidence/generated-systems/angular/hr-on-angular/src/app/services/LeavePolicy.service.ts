import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {LeavePolicy} from '../models/LeavePolicy';
import {OrganizationService} from '../services/Organization.service';
import {LeaveRequestService} from '../services/LeaveRequest.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LeavePolicyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	leavePolicy : LeavePolicy;

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
	// add a LeavePolicy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLeavePolicy(name, accrualRate, carryoverAllowed, maxBalance, Organization, LeaveRequests, LeaveCategory, AccrualUnit) : Observable<any> {
		const uri_ = this.apiUrl + '/LeavePolicy/create';
		const obj = {
			      		name: name,
      		accrualRate: accrualRate,
      		carryoverAllowed: carryoverAllowed,
      		maxBalance: maxBalance,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		LeaveRequests: LeaveRequests != null && LeaveRequests.length > 0 ? LeaveRequests : null,
      		LeaveCategory: LeaveCategory,
			AccrualUnit: AccrualUnit
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a LeavePolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLeavePolicy(name, accrualRate, carryoverAllowed, maxBalance, Organization, LeaveRequests, LeaveCategory, AccrualUnit, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/LeavePolicy/update/' + id;
		const obj = {
				      		name: name,
      		accrualRate: accrualRate,
      		carryoverAllowed: carryoverAllowed,
      		maxBalance: maxBalance,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		LeaveRequests: LeaveRequests != null && LeaveRequests.length > 0 ? LeaveRequests : null,
      		LeaveCategory: LeaveCategory,
			AccrualUnit: AccrualUnit
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a LeavePolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLeavePolicy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/LeavePolicy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a LeavePolicy
	// returns the results untouched as an Observable LeavePolicy
	// LeavePolicy model
	// delegates via URI
	//********************************************************************
	getLeavePolicy(id) : Observable<LeavePolicy> {
		const uri_ = this.apiUrl + '/LeavePolicy/load/' + id;

		return this.http.get<LeavePolicy>(uri_);
	}
	
	//********************************************************************
	// gets all LeavePolicy
	// returns the results untouched as JSON representation of an
	// Observable array of LeavePolicy models
	// delegates via URI
	//********************************************************************
	getLeavePolicys() : Observable<LeavePolicy[]> {
		const uri_ = this.apiUrl + '/LeavePolicy/';

		return this
			.http.get<LeavePolicy[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a LeavePolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( leavePolicyId, _organizationId ): Observable<any> {

		// get the LeavePolicy from storage
		this.loadHelper( leavePolicyId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.leavePolicy.organization = tmp;

	// save the LeavePolicy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a LeavePolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( leavePolicyId ): Observable<any> {

		// get the LeavePolicy from storage
		this.loadHelper( leavePolicyId );

	// assign Organization to null
	this.leavePolicy.organization = null;

	// save the LeavePolicy
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more leaveRequestsIds as a LeaveRequests
	// to a LeavePolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLeaveRequests( leavePolicyId, leaveRequestsIds ): Observable<any> {

		// get the LeavePolicy
		this.loadHelper( leavePolicyId );

	// split on a comma with no spaces
	var idList = leaveRequestsIds.split(',')

	// iterate over array of leaveRequests ids
	idList.forEach(function (id) {
		// read the LeaveRequest
		var leaveRequest = new LeaveRequestService(this.http).getLeaveRequest(id);
		// add the LeaveRequest if not already assigned
		if ( this.leavePolicy.leaveRequests.indexOf(leaveRequest) == -1 )
		this.leavePolicy.leaveRequests.push(leaveRequest);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more leaveRequestsIds as a LeaveRequests
	// from a LeavePolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLeaveRequests( leavePolicyId, leaveRequestsIds ): Observable<any> {

		// get the LeavePolicy
		this.loadHelper( leavePolicyId );


	// split on a comma with no spaces
	var idList 					= leaveRequestsIds.split(',');
	var leaveRequests 	= this.leavePolicy.leaveRequests;

	if ( leaveRequests != null && leaveRequestsIds != null ) {

		// iterate over array of leaveRequests ids
		leaveRequests.forEach(function (obj) {
			if ( leaveRequestsIds.indexOf(obj._id) > -1 ) {
				// remove the LeaveRequest
				this.leavePolicy.leaveRequests.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a LeavePolicy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/LeavePolicy/update/' + this.leavePolicy;

	return  this.http.post(uri_, this.leavePolicy );
}

	//********************************************************************
	// loadHelper - internal helper to load a LeavePolicy
	//********************************************************************	
	loadHelper( id ) {
		this.getLeavePolicy(id)
			.subscribe((res : LeavePolicy) => {
				this.leavePolicy = res;
			});
	}
}