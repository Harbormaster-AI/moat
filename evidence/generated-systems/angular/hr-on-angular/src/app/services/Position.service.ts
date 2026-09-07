import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Position} from '../models/Position';
import {DepartmentService} from '../services/Department.service';
import {JobProfileService} from '../services/JobProfile.service';
import {CostCenterService} from '../services/CostCenter.service';
import {LocationService} from '../services/Location.service';
import {EmploymentAssignmentService} from '../services/EmploymentAssignment.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PositionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	position : Position;

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
	// add a Position
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPosition(positionCode, fte, Department, JobProfile, CostCenter, Location, ManagerPosition, DirectReports, Assignments, Status, WorkLocationType) : Observable<any> {
		const uri_ = this.apiUrl + '/Position/create';
		const obj = {
			      		positionCode: positionCode,
      		fte: fte,
      		Department: Department != null && Department.length > 0 ? Department : null,
      		JobProfile: JobProfile != null && JobProfile.length > 0 ? JobProfile : null,
      		CostCenter: CostCenter != null && CostCenter.length > 0 ? CostCenter : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		ManagerPosition: ManagerPosition != null && ManagerPosition.length > 0 ? ManagerPosition : null,
      		DirectReports: DirectReports != null && DirectReports.length > 0 ? DirectReports : null,
      		Assignments: Assignments != null && Assignments.length > 0 ? Assignments : null,
      		Status: Status,
			WorkLocationType: WorkLocationType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePosition(positionCode, fte, Department, JobProfile, CostCenter, Location, ManagerPosition, DirectReports, Assignments, Status, WorkLocationType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Position/update/' + id;
		const obj = {
				      		positionCode: positionCode,
      		fte: fte,
      		Department: Department != null && Department.length > 0 ? Department : null,
      		JobProfile: JobProfile != null && JobProfile.length > 0 ? JobProfile : null,
      		CostCenter: CostCenter != null && CostCenter.length > 0 ? CostCenter : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		ManagerPosition: ManagerPosition != null && ManagerPosition.length > 0 ? ManagerPosition : null,
      		DirectReports: DirectReports != null && DirectReports.length > 0 ? DirectReports : null,
      		Assignments: Assignments != null && Assignments.length > 0 ? Assignments : null,
      		Status: Status,
			WorkLocationType: WorkLocationType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePosition(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Position/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Position
	// returns the results untouched as an Observable Position
	// Position model
	// delegates via URI
	//********************************************************************
	getPosition(id) : Observable<Position> {
		const uri_ = this.apiUrl + '/Position/load/' + id;

		return this.http.get<Position>(uri_);
	}
	
	//********************************************************************
	// gets all Position
	// returns the results untouched as JSON representation of an
	// Observable array of Position models
	// delegates via URI
	//********************************************************************
	getPositions() : Observable<Position[]> {
		const uri_ = this.apiUrl + '/Position/';

		return this
			.http.get<Position[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Department on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDepartment( positionId, _departmentId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// get the Department from storage
	var tmp 	= new DepartmentService(this.http).getDepartment(_departmentId);

	// assign the Department
	this.position.department = tmp;

	// save the Position
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Department on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDepartment( positionId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// assign Department to null
	this.position.department = null;

	// save the Position
	return this.saveHelper();
}

		//********************************************************************
	// assigns a JobProfile on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignJobProfile( positionId, _jobProfileId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// get the JobProfile from storage
	var tmp 	= new JobProfileService(this.http).getJobProfile(_jobProfileId);

	// assign the JobProfile
	this.position.jobProfile = tmp;

	// save the Position
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a JobProfile on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignJobProfile( positionId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// assign JobProfile to null
	this.position.jobProfile = null;

	// save the Position
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CostCenter on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCostCenter( positionId, _costCenterId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// get the CostCenter from storage
	var tmp 	= new CostCenterService(this.http).getCostCenter(_costCenterId);

	// assign the CostCenter
	this.position.costCenter = tmp;

	// save the Position
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CostCenter on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCostCenter( positionId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// assign CostCenter to null
	this.position.costCenter = null;

	// save the Position
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Location on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLocation( positionId, _locationId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// get the Location from storage
	var tmp 	= new LocationService(this.http).getLocation(_locationId);

	// assign the Location
	this.position.location = tmp;

	// save the Position
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Location on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLocation( positionId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// assign Location to null
	this.position.location = null;

	// save the Position
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ManagerPosition on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignManagerPosition( positionId, _managerPositionId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// get the Position from storage
	var tmp 	= new PositionService(this.http).getPosition(_managerPositionId);

	// assign the ManagerPosition
	this.position.managerPosition = tmp;

	// save the Position
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ManagerPosition on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignManagerPosition( positionId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// assign ManagerPosition to null
	this.position.managerPosition = null;

	// save the Position
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more directReportsIds as a DirectReports
	// to a Position
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDirectReports( positionId, directReportsIds ): Observable<any> {

		// get the Position
		this.loadHelper( positionId );

	// split on a comma with no spaces
	var idList = directReportsIds.split(',')

	// iterate over array of directReports ids
	idList.forEach(function (id) {
		// read the Position
		var position = new PositionService(this.http).getPosition(id);
		// add the Position if not already assigned
		if ( this.position.directReports.indexOf(position) == -1 )
		this.position.directReports.push(position);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more directReportsIds as a DirectReports
	// from a Position
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDirectReports( positionId, directReportsIds ): Observable<any> {

		// get the Position
		this.loadHelper( positionId );


	// split on a comma with no spaces
	var idList 					= directReportsIds.split(',');
	var directReports 	= this.position.directReports;

	if ( directReports != null && directReportsIds != null ) {

		// iterate over array of directReports ids
		directReports.forEach(function (obj) {
			if ( directReportsIds.indexOf(obj._id) > -1 ) {
				// remove the Position
				this.position.directReports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more assignmentsIds as a Assignments
	// to a Position
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAssignments( positionId, assignmentsIds ): Observable<any> {

		// get the Position
		this.loadHelper( positionId );

	// split on a comma with no spaces
	var idList = assignmentsIds.split(',')

	// iterate over array of assignments ids
	idList.forEach(function (id) {
		// read the EmploymentAssignment
		var employmentAssignment = new EmploymentAssignmentService(this.http).getEmploymentAssignment(id);
		// add the EmploymentAssignment if not already assigned
		if ( this.position.assignments.indexOf(employmentAssignment) == -1 )
		this.position.assignments.push(employmentAssignment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more assignmentsIds as a Assignments
	// from a Position
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAssignments( positionId, assignmentsIds ): Observable<any> {

		// get the Position
		this.loadHelper( positionId );


	// split on a comma with no spaces
	var idList 					= assignmentsIds.split(',');
	var assignments 	= this.position.assignments;

	if ( assignments != null && assignmentsIds != null ) {

		// iterate over array of assignments ids
		assignments.forEach(function (obj) {
			if ( assignmentsIds.indexOf(obj._id) > -1 ) {
				// remove the EmploymentAssignment
				this.position.assignments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Position
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Position/update/' + this.position;

	return  this.http.post(uri_, this.position );
}

	//********************************************************************
	// loadHelper - internal helper to load a Position
	//********************************************************************	
	loadHelper( id ) {
		this.getPosition(id)
			.subscribe((res : Position) => {
				this.position = res;
			});
	}
}