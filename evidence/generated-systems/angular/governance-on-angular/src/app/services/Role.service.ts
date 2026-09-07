import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Role} from '../models/Role';
import {RoleAssignmentService} from '../services/RoleAssignment.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RoleService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	role : Role;

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
	// add a Role
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRole(name, responsibility, Assignments) : Observable<any> {
		const uri_ = this.apiUrl + '/Role/create';
		const obj = {
			      		name: name,
      		responsibility: responsibility,
			Assignments: Assignments != null && Assignments.length > 0 ? Assignments : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Role
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRole(name, responsibility, Assignments, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Role/update/' + id;
		const obj = {
				      		name: name,
      		responsibility: responsibility,
			Assignments: Assignments != null && Assignments.length > 0 ? Assignments : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Role
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRole(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Role/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Role
	// returns the results untouched as an Observable Role
	// Role model
	// delegates via URI
	//********************************************************************
	getRole(id) : Observable<Role> {
		const uri_ = this.apiUrl + '/Role/load/' + id;

		return this.http.get<Role>(uri_);
	}
	
	//********************************************************************
	// gets all Role
	// returns the results untouched as JSON representation of an
	// Observable array of Role models
	// delegates via URI
	//********************************************************************
	getRoles() : Observable<Role[]> {
		const uri_ = this.apiUrl + '/Role/';

		return this
			.http.get<Role[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more assignmentsIds as a Assignments
	// to a Role
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAssignments( roleId, assignmentsIds ): Observable<any> {

		// get the Role
		this.loadHelper( roleId );

	// split on a comma with no spaces
	var idList = assignmentsIds.split(',')

	// iterate over array of assignments ids
	idList.forEach(function (id) {
		// read the RoleAssignment
		var roleAssignment = new RoleAssignmentService(this.http).getRoleAssignment(id);
		// add the RoleAssignment if not already assigned
		if ( this.role.assignments.indexOf(roleAssignment) == -1 )
		this.role.assignments.push(roleAssignment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more assignmentsIds as a Assignments
	// from a Role
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAssignments( roleId, assignmentsIds ): Observable<any> {

		// get the Role
		this.loadHelper( roleId );


	// split on a comma with no spaces
	var idList 					= assignmentsIds.split(',');
	var assignments 	= this.role.assignments;

	if ( assignments != null && assignmentsIds != null ) {

		// iterate over array of assignments ids
		assignments.forEach(function (obj) {
			if ( assignmentsIds.indexOf(obj._id) > -1 ) {
				// remove the RoleAssignment
				this.role.assignments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Role
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Role/update/' + this.role;

	return  this.http.post(uri_, this.role );
}

	//********************************************************************
	// loadHelper - internal helper to load a Role
	//********************************************************************	
	loadHelper( id ) {
		this.getRole(id)
			.subscribe((res : Role) => {
				this.role = res;
			});
	}
}