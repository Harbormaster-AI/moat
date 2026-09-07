import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ShiftAssignment} from '../models/ShiftAssignment';
import {ShiftService} from '../services/Shift.service';
import {EmployeeService} from '../services/Employee.service';
import {WorkCenterService} from '../services/WorkCenter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ShiftAssignmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	shiftAssignment : ShiftAssignment;

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
	// add a ShiftAssignment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addShiftAssignment(assignmentDate, Shift, Employee, WorkCenter) : Observable<any> {
		const uri_ = this.apiUrl + '/ShiftAssignment/create';
		const obj = {
			      		assignmentDate: assignmentDate,
      		Shift: Shift != null && Shift.length > 0 ? Shift : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			WorkCenter: WorkCenter != null && WorkCenter.length > 0 ? WorkCenter : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ShiftAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateShiftAssignment(assignmentDate, Shift, Employee, WorkCenter, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ShiftAssignment/update/' + id;
		const obj = {
				      		assignmentDate: assignmentDate,
      		Shift: Shift != null && Shift.length > 0 ? Shift : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			WorkCenter: WorkCenter != null && WorkCenter.length > 0 ? WorkCenter : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ShiftAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteShiftAssignment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ShiftAssignment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ShiftAssignment
	// returns the results untouched as an Observable ShiftAssignment
	// ShiftAssignment model
	// delegates via URI
	//********************************************************************
	getShiftAssignment(id) : Observable<ShiftAssignment> {
		const uri_ = this.apiUrl + '/ShiftAssignment/load/' + id;

		return this.http.get<ShiftAssignment>(uri_);
	}
	
	//********************************************************************
	// gets all ShiftAssignment
	// returns the results untouched as JSON representation of an
	// Observable array of ShiftAssignment models
	// delegates via URI
	//********************************************************************
	getShiftAssignments() : Observable<ShiftAssignment[]> {
		const uri_ = this.apiUrl + '/ShiftAssignment/';

		return this
			.http.get<ShiftAssignment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Shift on a ShiftAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignShift( shiftAssignmentId, _shiftId ): Observable<any> {

		// get the ShiftAssignment from storage
		this.loadHelper( shiftAssignmentId );

	// get the Shift from storage
	var tmp 	= new ShiftService(this.http).getShift(_shiftId);

	// assign the Shift
	this.shiftAssignment.shift = tmp;

	// save the ShiftAssignment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Shift on a ShiftAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignShift( shiftAssignmentId ): Observable<any> {

		// get the ShiftAssignment from storage
		this.loadHelper( shiftAssignmentId );

	// assign Shift to null
	this.shiftAssignment.shift = null;

	// save the ShiftAssignment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Employee on a ShiftAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( shiftAssignmentId, _employeeId ): Observable<any> {

		// get the ShiftAssignment from storage
		this.loadHelper( shiftAssignmentId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.shiftAssignment.employee = tmp;

	// save the ShiftAssignment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a ShiftAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( shiftAssignmentId ): Observable<any> {

		// get the ShiftAssignment from storage
		this.loadHelper( shiftAssignmentId );

	// assign Employee to null
	this.shiftAssignment.employee = null;

	// save the ShiftAssignment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a WorkCenter on a ShiftAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkCenter( shiftAssignmentId, _workCenterId ): Observable<any> {

		// get the ShiftAssignment from storage
		this.loadHelper( shiftAssignmentId );

	// get the WorkCenter from storage
	var tmp 	= new WorkCenterService(this.http).getWorkCenter(_workCenterId);

	// assign the WorkCenter
	this.shiftAssignment.workCenter = tmp;

	// save the ShiftAssignment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkCenter on a ShiftAssignment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkCenter( shiftAssignmentId ): Observable<any> {

		// get the ShiftAssignment from storage
		this.loadHelper( shiftAssignmentId );

	// assign WorkCenter to null
	this.shiftAssignment.workCenter = null;

	// save the ShiftAssignment
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ShiftAssignment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ShiftAssignment/update/' + this.shiftAssignment;

	return  this.http.post(uri_, this.shiftAssignment );
}

	//********************************************************************
	// loadHelper - internal helper to load a ShiftAssignment
	//********************************************************************	
	loadHelper( id ) {
		this.getShiftAssignment(id)
			.subscribe((res : ShiftAssignment) => {
				this.shiftAssignment = res;
			});
	}
}