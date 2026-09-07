import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {GovernanceBody} from '../models/GovernanceBody';
import {OrganizationService} from '../services/Organization.service';
import {RoleAssignmentService} from '../services/RoleAssignment.service';
import {PolicyService} from '../services/Policy.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class GovernanceBodyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	governanceBody : GovernanceBody;

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
	// add a GovernanceBody
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addGovernanceBody(name, charterUrl, chair, Organization, RoleAssignments, Policies, BodyType) : Observable<any> {
		const uri_ = this.apiUrl + '/GovernanceBody/create';
		const obj = {
			      		name: name,
      		charterUrl: charterUrl,
      		chair: chair,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		RoleAssignments: RoleAssignments != null && RoleAssignments.length > 0 ? RoleAssignments : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
			BodyType: BodyType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a GovernanceBody
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateGovernanceBody(name, charterUrl, chair, Organization, RoleAssignments, Policies, BodyType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/GovernanceBody/update/' + id;
		const obj = {
				      		name: name,
      		charterUrl: charterUrl,
      		chair: chair,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		RoleAssignments: RoleAssignments != null && RoleAssignments.length > 0 ? RoleAssignments : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
			BodyType: BodyType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a GovernanceBody
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteGovernanceBody(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/GovernanceBody/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a GovernanceBody
	// returns the results untouched as an Observable GovernanceBody
	// GovernanceBody model
	// delegates via URI
	//********************************************************************
	getGovernanceBody(id) : Observable<GovernanceBody> {
		const uri_ = this.apiUrl + '/GovernanceBody/load/' + id;

		return this.http.get<GovernanceBody>(uri_);
	}
	
	//********************************************************************
	// gets all GovernanceBody
	// returns the results untouched as JSON representation of an
	// Observable array of GovernanceBody models
	// delegates via URI
	//********************************************************************
	getGovernanceBodys() : Observable<GovernanceBody[]> {
		const uri_ = this.apiUrl + '/GovernanceBody/';

		return this
			.http.get<GovernanceBody[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a GovernanceBody
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( governanceBodyId, _organizationId ): Observable<any> {

		// get the GovernanceBody from storage
		this.loadHelper( governanceBodyId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.governanceBody.organization = tmp;

	// save the GovernanceBody
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a GovernanceBody
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( governanceBodyId ): Observable<any> {

		// get the GovernanceBody from storage
		this.loadHelper( governanceBodyId );

	// assign Organization to null
	this.governanceBody.organization = null;

	// save the GovernanceBody
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more roleAssignmentsIds as a RoleAssignments
	// to a GovernanceBody
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRoleAssignments( governanceBodyId, roleAssignmentsIds ): Observable<any> {

		// get the GovernanceBody
		this.loadHelper( governanceBodyId );

	// split on a comma with no spaces
	var idList = roleAssignmentsIds.split(',')

	// iterate over array of roleAssignments ids
	idList.forEach(function (id) {
		// read the RoleAssignment
		var roleAssignment = new RoleAssignmentService(this.http).getRoleAssignment(id);
		// add the RoleAssignment if not already assigned
		if ( this.governanceBody.roleAssignments.indexOf(roleAssignment) == -1 )
		this.governanceBody.roleAssignments.push(roleAssignment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more roleAssignmentsIds as a RoleAssignments
	// from a GovernanceBody
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRoleAssignments( governanceBodyId, roleAssignmentsIds ): Observable<any> {

		// get the GovernanceBody
		this.loadHelper( governanceBodyId );


	// split on a comma with no spaces
	var idList 					= roleAssignmentsIds.split(',');
	var roleAssignments 	= this.governanceBody.roleAssignments;

	if ( roleAssignments != null && roleAssignmentsIds != null ) {

		// iterate over array of roleAssignments ids
		roleAssignments.forEach(function (obj) {
			if ( roleAssignmentsIds.indexOf(obj._id) > -1 ) {
				// remove the RoleAssignment
				this.governanceBody.roleAssignments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a GovernanceBody
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( governanceBodyId, policiesIds ): Observable<any> {

		// get the GovernanceBody
		this.loadHelper( governanceBodyId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.governanceBody.policies.indexOf(policy) == -1 )
		this.governanceBody.policies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a GovernanceBody
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( governanceBodyId, policiesIds ): Observable<any> {

		// get the GovernanceBody
		this.loadHelper( governanceBodyId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.governanceBody.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.governanceBody.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a GovernanceBody
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/GovernanceBody/update/' + this.governanceBody;

	return  this.http.post(uri_, this.governanceBody );
}

	//********************************************************************
	// loadHelper - internal helper to load a GovernanceBody
	//********************************************************************	
	loadHelper( id ) {
		this.getGovernanceBody(id)
			.subscribe((res : GovernanceBody) => {
				this.governanceBody = res;
			});
	}
}