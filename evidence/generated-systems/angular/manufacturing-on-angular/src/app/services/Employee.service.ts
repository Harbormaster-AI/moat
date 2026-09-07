import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Employee} from '../models/Employee';
import {WorkCenterService} from '../services/WorkCenter.service';
import {ShiftAssignmentService} from '../services/ShiftAssignment.service';
import {CorrectiveActionService} from '../services/CorrectiveAction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EmployeeService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	employee : Employee;

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
	// add a Employee
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEmployee(firstName, lastName, WorkCenter, ShiftAssignments, CorrectiveActions, Role, SkillLevel) : Observable<any> {
		const uri_ = this.apiUrl + '/Employee/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		WorkCenter: WorkCenter != null && WorkCenter.length > 0 ? WorkCenter : null,
      		ShiftAssignments: ShiftAssignments != null && ShiftAssignments.length > 0 ? ShiftAssignments : null,
      		CorrectiveActions: CorrectiveActions != null && CorrectiveActions.length > 0 ? CorrectiveActions : null,
      		Role: Role,
			SkillLevel: SkillLevel
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEmployee(firstName, lastName, WorkCenter, ShiftAssignments, CorrectiveActions, Role, SkillLevel, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Employee/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		WorkCenter: WorkCenter != null && WorkCenter.length > 0 ? WorkCenter : null,
      		ShiftAssignments: ShiftAssignments != null && ShiftAssignments.length > 0 ? ShiftAssignments : null,
      		CorrectiveActions: CorrectiveActions != null && CorrectiveActions.length > 0 ? CorrectiveActions : null,
      		Role: Role,
			SkillLevel: SkillLevel
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEmployee(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Employee/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Employee
	// returns the results untouched as an Observable Employee
	// Employee model
	// delegates via URI
	//********************************************************************
	getEmployee(id) : Observable<Employee> {
		const uri_ = this.apiUrl + '/Employee/load/' + id;

		return this.http.get<Employee>(uri_);
	}
	
	//********************************************************************
	// gets all Employee
	// returns the results untouched as JSON representation of an
	// Observable array of Employee models
	// delegates via URI
	//********************************************************************
	getEmployees() : Observable<Employee[]> {
		const uri_ = this.apiUrl + '/Employee/';

		return this
			.http.get<Employee[]>(uri_);
	}
	
			//********************************************************************
	// assigns a WorkCenter on a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkCenter( employeeId, _workCenterId ): Observable<any> {

		// get the Employee from storage
		this.loadHelper( employeeId );

	// get the WorkCenter from storage
	var tmp 	= new WorkCenterService(this.http).getWorkCenter(_workCenterId);

	// assign the WorkCenter
	this.employee.workCenter = tmp;

	// save the Employee
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkCenter on a Employee
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkCenter( employeeId ): Observable<any> {

		// get the Employee from storage
		this.loadHelper( employeeId );

	// assign WorkCenter to null
	this.employee.workCenter = null;

	// save the Employee
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more shiftAssignmentsIds as a ShiftAssignments
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addShiftAssignments( employeeId, shiftAssignmentsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = shiftAssignmentsIds.split(',')

	// iterate over array of shiftAssignments ids
	idList.forEach(function (id) {
		// read the ShiftAssignment
		var shiftAssignment = new ShiftAssignmentService(this.http).getShiftAssignment(id);
		// add the ShiftAssignment if not already assigned
		if ( this.employee.shiftAssignments.indexOf(shiftAssignment) == -1 )
		this.employee.shiftAssignments.push(shiftAssignment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more shiftAssignmentsIds as a ShiftAssignments
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeShiftAssignments( employeeId, shiftAssignmentsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= shiftAssignmentsIds.split(',');
	var shiftAssignments 	= this.employee.shiftAssignments;

	if ( shiftAssignments != null && shiftAssignmentsIds != null ) {

		// iterate over array of shiftAssignments ids
		shiftAssignments.forEach(function (obj) {
			if ( shiftAssignmentsIds.indexOf(obj._id) > -1 ) {
				// remove the ShiftAssignment
				this.employee.shiftAssignments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more correctiveActionsIds as a CorrectiveActions
	// to a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCorrectiveActions( employeeId, correctiveActionsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );

	// split on a comma with no spaces
	var idList = correctiveActionsIds.split(',')

	// iterate over array of correctiveActions ids
	idList.forEach(function (id) {
		// read the CorrectiveAction
		var correctiveAction = new CorrectiveActionService(this.http).getCorrectiveAction(id);
		// add the CorrectiveAction if not already assigned
		if ( this.employee.correctiveActions.indexOf(correctiveAction) == -1 )
		this.employee.correctiveActions.push(correctiveAction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more correctiveActionsIds as a CorrectiveActions
	// from a Employee
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCorrectiveActions( employeeId, correctiveActionsIds ): Observable<any> {

		// get the Employee
		this.loadHelper( employeeId );


	// split on a comma with no spaces
	var idList 					= correctiveActionsIds.split(',');
	var correctiveActions 	= this.employee.correctiveActions;

	if ( correctiveActions != null && correctiveActionsIds != null ) {

		// iterate over array of correctiveActions ids
		correctiveActions.forEach(function (obj) {
			if ( correctiveActionsIds.indexOf(obj._id) > -1 ) {
				// remove the CorrectiveAction
				this.employee.correctiveActions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Employee
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Employee/update/' + this.employee;

	return  this.http.post(uri_, this.employee );
}

	//********************************************************************
	// loadHelper - internal helper to load a Employee
	//********************************************************************	
	loadHelper( id ) {
		this.getEmployee(id)
			.subscribe((res : Employee) => {
				this.employee = res;
			});
	}
}