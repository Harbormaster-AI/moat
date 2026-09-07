import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Termination} from '../models/Termination';
import {EmployeeService} from '../services/Employee.service';
import {EmploymentAssignmentService} from '../services/EmploymentAssignment.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TerminationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	termination : Termination;

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
	// add a Termination
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTermination(terminationNumber, terminationDate, notes, eligibleForRehire, Employee, Assignment, Reason, Type) : Observable<any> {
		const uri_ = this.apiUrl + '/Termination/create';
		const obj = {
			      		terminationNumber: terminationNumber,
      		terminationDate: terminationDate,
      		notes: notes,
      		eligibleForRehire: eligibleForRehire,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Assignment: Assignment != null && Assignment.length > 0 ? Assignment : null,
      		Reason: Reason,
			Type: Type
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Termination
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTermination(terminationNumber, terminationDate, notes, eligibleForRehire, Employee, Assignment, Reason, Type, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Termination/update/' + id;
		const obj = {
				      		terminationNumber: terminationNumber,
      		terminationDate: terminationDate,
      		notes: notes,
      		eligibleForRehire: eligibleForRehire,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Assignment: Assignment != null && Assignment.length > 0 ? Assignment : null,
      		Reason: Reason,
			Type: Type
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Termination
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTermination(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Termination/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Termination
	// returns the results untouched as an Observable Termination
	// Termination model
	// delegates via URI
	//********************************************************************
	getTermination(id) : Observable<Termination> {
		const uri_ = this.apiUrl + '/Termination/load/' + id;

		return this.http.get<Termination>(uri_);
	}
	
	//********************************************************************
	// gets all Termination
	// returns the results untouched as JSON representation of an
	// Observable array of Termination models
	// delegates via URI
	//********************************************************************
	getTerminations() : Observable<Termination[]> {
		const uri_ = this.apiUrl + '/Termination/';

		return this
			.http.get<Termination[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a Termination
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( terminationId, _employeeId ): Observable<any> {

		// get the Termination from storage
		this.loadHelper( terminationId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.termination.employee = tmp;

	// save the Termination
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a Termination
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( terminationId ): Observable<any> {

		// get the Termination from storage
		this.loadHelper( terminationId );

	// assign Employee to null
	this.termination.employee = null;

	// save the Termination
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Assignment on a Termination
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAssignment( terminationId, _assignmentId ): Observable<any> {

		// get the Termination from storage
		this.loadHelper( terminationId );

	// get the EmploymentAssignment from storage
	var tmp 	= new EmploymentAssignmentService(this.http).getEmploymentAssignment(_assignmentId);

	// assign the Assignment
	this.termination.assignment = tmp;

	// save the Termination
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Assignment on a Termination
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAssignment( terminationId ): Observable<any> {

		// get the Termination from storage
		this.loadHelper( terminationId );

	// assign Assignment to null
	this.termination.assignment = null;

	// save the Termination
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Termination
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Termination/update/' + this.termination;

	return  this.http.post(uri_, this.termination );
}

	//********************************************************************
	// loadHelper - internal helper to load a Termination
	//********************************************************************	
	loadHelper( id ) {
		this.getTermination(id)
			.subscribe((res : Termination) => {
				this.termination = res;
			});
	}
}