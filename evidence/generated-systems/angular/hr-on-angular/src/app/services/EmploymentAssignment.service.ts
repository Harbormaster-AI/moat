import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {EmploymentAssignment} from '../models/EmploymentAssignment';
import {EmployeeService} from '../services/Employee.service';
import {PositionService} from '../services/Position.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EmploymentAssignmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	employmentAssignment : EmploymentAssignment;

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
	// add a EmploymentAssignment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEmploymentAssignment(startDate, endDate, primary, Employee, Position, Supervisor, AssignmentType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/EmploymentAssignment/create';
		const obj = {
			      		startDate: startDate,
      		endDate: endDate,
      		primary: primary,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Position: Position != null && Position.length > 0 ? Position : null,
      		Supervisor: Supervisor != null && Supervisor.length > 0 ? Supervisor : null,
      		AssignmentType: AssignmentType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a EmploymentAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEmploymentAssignment(startDate, endDate, primary, Employee, Position, Supervisor, AssignmentType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/EmploymentAssignment/update/' + id;
		const obj = {
				      		startDate: startDate,
      		endDate: endDate,
      		primary: primary,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Position: Position != null && Position.length > 0 ? Position : null,
      		Supervisor: Supervisor != null && Supervisor.length > 0 ? Supervisor : null,
      		AssignmentType: AssignmentType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a EmploymentAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEmploymentAssignment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/EmploymentAssignment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a EmploymentAssignment
	// returns the results untouched as an Observable EmploymentAssignment
	// EmploymentAssignment model
	// delegates via URI
	//********************************************************************
	getEmploymentAssignment(id) : Observable<EmploymentAssignment> {
		const uri_ = this.apiUrl + '/EmploymentAssignment/load/' + id;

		return this.http.get<EmploymentAssignment>(uri_);
	}
	
	//********************************************************************
	// gets all EmploymentAssignment
	// returns the results untouched as JSON representation of an
	// Observable array of EmploymentAssignment models
	// delegates via URI
	//********************************************************************
	getEmploymentAssignments() : Observable<EmploymentAssignment[]> {
		const uri_ = this.apiUrl + '/EmploymentAssignment/';

		return this
			.http.get<EmploymentAssignment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a EmploymentAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( employmentAssignmentId, _employeeId ): Observable<any> {

		// get the EmploymentAssignment from storage
		this.loadHelper( employmentAssignmentId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.employmentAssignment.employee = tmp;

	// save the EmploymentAssignment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a EmploymentAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( employmentAssignmentId ): Observable<any> {

		// get the EmploymentAssignment from storage
		this.loadHelper( employmentAssignmentId );

	// assign Employee to null
	this.employmentAssignment.employee = null;

	// save the EmploymentAssignment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Position on a EmploymentAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPosition( employmentAssignmentId, _positionId ): Observable<any> {

		// get the EmploymentAssignment from storage
		this.loadHelper( employmentAssignmentId );

	// get the Position from storage
	var tmp 	= new PositionService(this.http).getPosition(_positionId);

	// assign the Position
	this.employmentAssignment.position = tmp;

	// save the EmploymentAssignment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Position on a EmploymentAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPosition( employmentAssignmentId ): Observable<any> {

		// get the EmploymentAssignment from storage
		this.loadHelper( employmentAssignmentId );

	// assign Position to null
	this.employmentAssignment.position = null;

	// save the EmploymentAssignment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Supervisor on a EmploymentAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSupervisor( employmentAssignmentId, _supervisorId ): Observable<any> {

		// get the EmploymentAssignment from storage
		this.loadHelper( employmentAssignmentId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_supervisorId);

	// assign the Supervisor
	this.employmentAssignment.supervisor = tmp;

	// save the EmploymentAssignment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Supervisor on a EmploymentAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSupervisor( employmentAssignmentId ): Observable<any> {

		// get the EmploymentAssignment from storage
		this.loadHelper( employmentAssignmentId );

	// assign Supervisor to null
	this.employmentAssignment.supervisor = null;

	// save the EmploymentAssignment
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a EmploymentAssignment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/EmploymentAssignment/update/' + this.employmentAssignment;

	return  this.http.post(uri_, this.employmentAssignment );
}

	//********************************************************************
	// loadHelper - internal helper to load a EmploymentAssignment
	//********************************************************************	
	loadHelper( id ) {
		this.getEmploymentAssignment(id)
			.subscribe((res : EmploymentAssignment) => {
				this.employmentAssignment = res;
			});
	}
}