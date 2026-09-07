import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Shift} from '../models/Shift';
import {PlantService} from '../services/Plant.service';
import {ShiftAssignmentService} from '../services/ShiftAssignment.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ShiftService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	shift : Shift;

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
	// add a Shift
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addShift(shiftName, startTime, endTime, Plant, Assignments, ShiftType) : Observable<any> {
		const uri_ = this.apiUrl + '/Shift/create';
		const obj = {
			      		shiftName: shiftName,
      		startTime: startTime,
      		endTime: endTime,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		Assignments: Assignments != null && Assignments.length > 0 ? Assignments : null,
			ShiftType: ShiftType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Shift
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateShift(shiftName, startTime, endTime, Plant, Assignments, ShiftType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Shift/update/' + id;
		const obj = {
				      		shiftName: shiftName,
      		startTime: startTime,
      		endTime: endTime,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		Assignments: Assignments != null && Assignments.length > 0 ? Assignments : null,
			ShiftType: ShiftType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Shift
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteShift(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Shift/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Shift
	// returns the results untouched as an Observable Shift
	// Shift model
	// delegates via URI
	//********************************************************************
	getShift(id) : Observable<Shift> {
		const uri_ = this.apiUrl + '/Shift/load/' + id;

		return this.http.get<Shift>(uri_);
	}
	
	//********************************************************************
	// gets all Shift
	// returns the results untouched as JSON representation of an
	// Observable array of Shift models
	// delegates via URI
	//********************************************************************
	getShifts() : Observable<Shift[]> {
		const uri_ = this.apiUrl + '/Shift/';

		return this
			.http.get<Shift[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Plant on a Shift
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( shiftId, _plantId ): Observable<any> {

		// get the Shift from storage
		this.loadHelper( shiftId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.shift.plant = tmp;

	// save the Shift
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a Shift
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( shiftId ): Observable<any> {

		// get the Shift from storage
		this.loadHelper( shiftId );

	// assign Plant to null
	this.shift.plant = null;

	// save the Shift
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more assignmentsIds as a Assignments
	// to a Shift
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAssignments( shiftId, assignmentsIds ): Observable<any> {

		// get the Shift
		this.loadHelper( shiftId );

	// split on a comma with no spaces
	var idList = assignmentsIds.split(',')

	// iterate over array of assignments ids
	idList.forEach(function (id) {
		// read the ShiftAssignment
		var shiftAssignment = new ShiftAssignmentService(this.http).getShiftAssignment(id);
		// add the ShiftAssignment if not already assigned
		if ( this.shift.assignments.indexOf(shiftAssignment) == -1 )
		this.shift.assignments.push(shiftAssignment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more assignmentsIds as a Assignments
	// from a Shift
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAssignments( shiftId, assignmentsIds ): Observable<any> {

		// get the Shift
		this.loadHelper( shiftId );


	// split on a comma with no spaces
	var idList 					= assignmentsIds.split(',');
	var assignments 	= this.shift.assignments;

	if ( assignments != null && assignmentsIds != null ) {

		// iterate over array of assignments ids
		assignments.forEach(function (obj) {
			if ( assignmentsIds.indexOf(obj._id) > -1 ) {
				// remove the ShiftAssignment
				this.shift.assignments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Shift
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Shift/update/' + this.shift;

	return  this.http.post(uri_, this.shift );
}

	//********************************************************************
	// loadHelper - internal helper to load a Shift
	//********************************************************************	
	loadHelper( id ) {
		this.getShift(id)
			.subscribe((res : Shift) => {
				this.shift = res;
			});
	}
}