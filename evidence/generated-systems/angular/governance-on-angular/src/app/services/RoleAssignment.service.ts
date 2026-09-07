import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {RoleAssignment} from '../models/RoleAssignment';
import {PersonService} from '../services/Person.service';
import {RoleService} from '../services/Role.service';
import {GovernanceBodyService} from '../services/GovernanceBody.service';
import {OrganizationService} from '../services/Organization.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RoleAssignmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	roleAssignment : RoleAssignment;

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
	// add a RoleAssignment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRoleAssignment(effectiveFrom, effectiveTo, Person, Role, GovernanceBody, Organization) : Observable<any> {
		const uri_ = this.apiUrl + '/RoleAssignment/create';
		const obj = {
			      		effectiveFrom: effectiveFrom,
      		effectiveTo: effectiveTo,
      		Person: Person != null && Person.length > 0 ? Person : null,
      		Role: Role != null && Role.length > 0 ? Role : null,
      		GovernanceBody: GovernanceBody != null && GovernanceBody.length > 0 ? GovernanceBody : null,
			Organization: Organization != null && Organization.length > 0 ? Organization : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a RoleAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRoleAssignment(effectiveFrom, effectiveTo, Person, Role, GovernanceBody, Organization, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/RoleAssignment/update/' + id;
		const obj = {
				      		effectiveFrom: effectiveFrom,
      		effectiveTo: effectiveTo,
      		Person: Person != null && Person.length > 0 ? Person : null,
      		Role: Role != null && Role.length > 0 ? Role : null,
      		GovernanceBody: GovernanceBody != null && GovernanceBody.length > 0 ? GovernanceBody : null,
			Organization: Organization != null && Organization.length > 0 ? Organization : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a RoleAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRoleAssignment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/RoleAssignment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a RoleAssignment
	// returns the results untouched as an Observable RoleAssignment
	// RoleAssignment model
	// delegates via URI
	//********************************************************************
	getRoleAssignment(id) : Observable<RoleAssignment> {
		const uri_ = this.apiUrl + '/RoleAssignment/load/' + id;

		return this.http.get<RoleAssignment>(uri_);
	}
	
	//********************************************************************
	// gets all RoleAssignment
	// returns the results untouched as JSON representation of an
	// Observable array of RoleAssignment models
	// delegates via URI
	//********************************************************************
	getRoleAssignments() : Observable<RoleAssignment[]> {
		const uri_ = this.apiUrl + '/RoleAssignment/';

		return this
			.http.get<RoleAssignment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Person on a RoleAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPerson( roleAssignmentId, _personId ): Observable<any> {

		// get the RoleAssignment from storage
		this.loadHelper( roleAssignmentId );

	// get the Person from storage
	var tmp 	= new PersonService(this.http).getPerson(_personId);

	// assign the Person
	this.roleAssignment.person = tmp;

	// save the RoleAssignment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Person on a RoleAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPerson( roleAssignmentId ): Observable<any> {

		// get the RoleAssignment from storage
		this.loadHelper( roleAssignmentId );

	// assign Person to null
	this.roleAssignment.person = null;

	// save the RoleAssignment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Role on a RoleAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRole( roleAssignmentId, _roleId ): Observable<any> {

		// get the RoleAssignment from storage
		this.loadHelper( roleAssignmentId );

	// get the Role from storage
	var tmp 	= new RoleService(this.http).getRole(_roleId);

	// assign the Role
	this.roleAssignment.role = tmp;

	// save the RoleAssignment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Role on a RoleAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRole( roleAssignmentId ): Observable<any> {

		// get the RoleAssignment from storage
		this.loadHelper( roleAssignmentId );

	// assign Role to null
	this.roleAssignment.role = null;

	// save the RoleAssignment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a GovernanceBody on a RoleAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignGovernanceBody( roleAssignmentId, _governanceBodyId ): Observable<any> {

		// get the RoleAssignment from storage
		this.loadHelper( roleAssignmentId );

	// get the GovernanceBody from storage
	var tmp 	= new GovernanceBodyService(this.http).getGovernanceBody(_governanceBodyId);

	// assign the GovernanceBody
	this.roleAssignment.governanceBody = tmp;

	// save the RoleAssignment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a GovernanceBody on a RoleAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignGovernanceBody( roleAssignmentId ): Observable<any> {

		// get the RoleAssignment from storage
		this.loadHelper( roleAssignmentId );

	// assign GovernanceBody to null
	this.roleAssignment.governanceBody = null;

	// save the RoleAssignment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Organization on a RoleAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( roleAssignmentId, _organizationId ): Observable<any> {

		// get the RoleAssignment from storage
		this.loadHelper( roleAssignmentId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.roleAssignment.organization = tmp;

	// save the RoleAssignment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a RoleAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( roleAssignmentId ): Observable<any> {

		// get the RoleAssignment from storage
		this.loadHelper( roleAssignmentId );

	// assign Organization to null
	this.roleAssignment.organization = null;

	// save the RoleAssignment
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a RoleAssignment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/RoleAssignment/update/' + this.roleAssignment;

	return  this.http.post(uri_, this.roleAssignment );
}

	//********************************************************************
	// loadHelper - internal helper to load a RoleAssignment
	//********************************************************************	
	loadHelper( id ) {
		this.getRoleAssignment(id)
			.subscribe((res : RoleAssignment) => {
				this.roleAssignment = res;
			});
	}
}