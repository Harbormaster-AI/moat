import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Person} from '../models/Person';
import {RoleAssignmentService} from '../services/RoleAssignment.service';
import {PolicyService} from '../services/Policy.service';
import {CorrectiveActionService} from '../services/CorrectiveAction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PersonService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	person : Person;

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
	// add a Person
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPerson(firstName, lastName, email, department, RoleAssignments, OwnedPolicies, CorrectiveActions) : Observable<any> {
		const uri_ = this.apiUrl + '/Person/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		email: email,
      		department: department,
      		RoleAssignments: RoleAssignments != null && RoleAssignments.length > 0 ? RoleAssignments : null,
      		OwnedPolicies: OwnedPolicies != null && OwnedPolicies.length > 0 ? OwnedPolicies : null,
			CorrectiveActions: CorrectiveActions != null && CorrectiveActions.length > 0 ? CorrectiveActions : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Person
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePerson(firstName, lastName, email, department, RoleAssignments, OwnedPolicies, CorrectiveActions, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Person/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		email: email,
      		department: department,
      		RoleAssignments: RoleAssignments != null && RoleAssignments.length > 0 ? RoleAssignments : null,
      		OwnedPolicies: OwnedPolicies != null && OwnedPolicies.length > 0 ? OwnedPolicies : null,
			CorrectiveActions: CorrectiveActions != null && CorrectiveActions.length > 0 ? CorrectiveActions : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Person
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePerson(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Person/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Person
	// returns the results untouched as an Observable Person
	// Person model
	// delegates via URI
	//********************************************************************
	getPerson(id) : Observable<Person> {
		const uri_ = this.apiUrl + '/Person/load/' + id;

		return this.http.get<Person>(uri_);
	}
	
	//********************************************************************
	// gets all Person
	// returns the results untouched as JSON representation of an
	// Observable array of Person models
	// delegates via URI
	//********************************************************************
	getPersons() : Observable<Person[]> {
		const uri_ = this.apiUrl + '/Person/';

		return this
			.http.get<Person[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more roleAssignmentsIds as a RoleAssignments
	// to a Person
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRoleAssignments( personId, roleAssignmentsIds ): Observable<any> {

		// get the Person
		this.loadHelper( personId );

	// split on a comma with no spaces
	var idList = roleAssignmentsIds.split(',')

	// iterate over array of roleAssignments ids
	idList.forEach(function (id) {
		// read the RoleAssignment
		var roleAssignment = new RoleAssignmentService(this.http).getRoleAssignment(id);
		// add the RoleAssignment if not already assigned
		if ( this.person.roleAssignments.indexOf(roleAssignment) == -1 )
		this.person.roleAssignments.push(roleAssignment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more roleAssignmentsIds as a RoleAssignments
	// from a Person
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRoleAssignments( personId, roleAssignmentsIds ): Observable<any> {

		// get the Person
		this.loadHelper( personId );


	// split on a comma with no spaces
	var idList 					= roleAssignmentsIds.split(',');
	var roleAssignments 	= this.person.roleAssignments;

	if ( roleAssignments != null && roleAssignmentsIds != null ) {

		// iterate over array of roleAssignments ids
		roleAssignments.forEach(function (obj) {
			if ( roleAssignmentsIds.indexOf(obj._id) > -1 ) {
				// remove the RoleAssignment
				this.person.roleAssignments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ownedPoliciesIds as a OwnedPolicies
	// to a Person
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOwnedPolicies( personId, ownedPoliciesIds ): Observable<any> {

		// get the Person
		this.loadHelper( personId );

	// split on a comma with no spaces
	var idList = ownedPoliciesIds.split(',')

	// iterate over array of ownedPolicies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.person.ownedPolicies.indexOf(policy) == -1 )
		this.person.ownedPolicies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ownedPoliciesIds as a OwnedPolicies
	// from a Person
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOwnedPolicies( personId, ownedPoliciesIds ): Observable<any> {

		// get the Person
		this.loadHelper( personId );


	// split on a comma with no spaces
	var idList 					= ownedPoliciesIds.split(',');
	var ownedPolicies 	= this.person.ownedPolicies;

	if ( ownedPolicies != null && ownedPoliciesIds != null ) {

		// iterate over array of ownedPolicies ids
		ownedPolicies.forEach(function (obj) {
			if ( ownedPoliciesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.person.ownedPolicies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more correctiveActionsIds as a CorrectiveActions
	// to a Person
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCorrectiveActions( personId, correctiveActionsIds ): Observable<any> {

		// get the Person
		this.loadHelper( personId );

	// split on a comma with no spaces
	var idList = correctiveActionsIds.split(',')

	// iterate over array of correctiveActions ids
	idList.forEach(function (id) {
		// read the CorrectiveAction
		var correctiveAction = new CorrectiveActionService(this.http).getCorrectiveAction(id);
		// add the CorrectiveAction if not already assigned
		if ( this.person.correctiveActions.indexOf(correctiveAction) == -1 )
		this.person.correctiveActions.push(correctiveAction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more correctiveActionsIds as a CorrectiveActions
	// from a Person
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCorrectiveActions( personId, correctiveActionsIds ): Observable<any> {

		// get the Person
		this.loadHelper( personId );


	// split on a comma with no spaces
	var idList 					= correctiveActionsIds.split(',');
	var correctiveActions 	= this.person.correctiveActions;

	if ( correctiveActions != null && correctiveActionsIds != null ) {

		// iterate over array of correctiveActions ids
		correctiveActions.forEach(function (obj) {
			if ( correctiveActionsIds.indexOf(obj._id) > -1 ) {
				// remove the CorrectiveAction
				this.person.correctiveActions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Person
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Person/update/' + this.person;

	return  this.http.post(uri_, this.person );
}

	//********************************************************************
	// loadHelper - internal helper to load a Person
	//********************************************************************	
	loadHelper( id ) {
		this.getPerson(id)
			.subscribe((res : Person) => {
				this.person = res;
			});
	}
}