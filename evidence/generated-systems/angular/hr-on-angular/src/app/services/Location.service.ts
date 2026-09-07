import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Location} from '../models/Location';
import {OrganizationService} from '../services/Organization.service';
import {DepartmentService} from '../services/Department.service';
import {PositionService} from '../services/Position.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LocationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	location : Location;

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
	// add a Location
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLocation(name, address, timezone, Organization, Departments, Positions, Employees) : Observable<any> {
		const uri_ = this.apiUrl + '/Location/create';
		const obj = {
			      		name: name,
      		address: address,
      		timezone: timezone,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Departments: Departments != null && Departments.length > 0 ? Departments : null,
      		Positions: Positions != null && Positions.length > 0 ? Positions : null,
			Employees: Employees != null && Employees.length > 0 ? Employees : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Location
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLocation(name, address, timezone, Organization, Departments, Positions, Employees, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Location/update/' + id;
		const obj = {
				      		name: name,
      		address: address,
      		timezone: timezone,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Departments: Departments != null && Departments.length > 0 ? Departments : null,
      		Positions: Positions != null && Positions.length > 0 ? Positions : null,
			Employees: Employees != null && Employees.length > 0 ? Employees : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Location
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLocation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Location/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Location
	// returns the results untouched as an Observable Location
	// Location model
	// delegates via URI
	//********************************************************************
	getLocation(id) : Observable<Location> {
		const uri_ = this.apiUrl + '/Location/load/' + id;

		return this.http.get<Location>(uri_);
	}
	
	//********************************************************************
	// gets all Location
	// returns the results untouched as JSON representation of an
	// Observable array of Location models
	// delegates via URI
	//********************************************************************
	getLocations() : Observable<Location[]> {
		const uri_ = this.apiUrl + '/Location/';

		return this
			.http.get<Location[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Location
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( locationId, _organizationId ): Observable<any> {

		// get the Location from storage
		this.loadHelper( locationId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.location.organization = tmp;

	// save the Location
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Location
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( locationId ): Observable<any> {

		// get the Location from storage
		this.loadHelper( locationId );

	// assign Organization to null
	this.location.organization = null;

	// save the Location
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more departmentsIds as a Departments
	// to a Location
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDepartments( locationId, departmentsIds ): Observable<any> {

		// get the Location
		this.loadHelper( locationId );

	// split on a comma with no spaces
	var idList = departmentsIds.split(',')

	// iterate over array of departments ids
	idList.forEach(function (id) {
		// read the Department
		var department = new DepartmentService(this.http).getDepartment(id);
		// add the Department if not already assigned
		if ( this.location.departments.indexOf(department) == -1 )
		this.location.departments.push(department);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more departmentsIds as a Departments
	// from a Location
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDepartments( locationId, departmentsIds ): Observable<any> {

		// get the Location
		this.loadHelper( locationId );


	// split on a comma with no spaces
	var idList 					= departmentsIds.split(',');
	var departments 	= this.location.departments;

	if ( departments != null && departmentsIds != null ) {

		// iterate over array of departments ids
		departments.forEach(function (obj) {
			if ( departmentsIds.indexOf(obj._id) > -1 ) {
				// remove the Department
				this.location.departments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more positionsIds as a Positions
	// to a Location
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPositions( locationId, positionsIds ): Observable<any> {

		// get the Location
		this.loadHelper( locationId );

	// split on a comma with no spaces
	var idList = positionsIds.split(',')

	// iterate over array of positions ids
	idList.forEach(function (id) {
		// read the Position
		var position = new PositionService(this.http).getPosition(id);
		// add the Position if not already assigned
		if ( this.location.positions.indexOf(position) == -1 )
		this.location.positions.push(position);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more positionsIds as a Positions
	// from a Location
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePositions( locationId, positionsIds ): Observable<any> {

		// get the Location
		this.loadHelper( locationId );


	// split on a comma with no spaces
	var idList 					= positionsIds.split(',');
	var positions 	= this.location.positions;

	if ( positions != null && positionsIds != null ) {

		// iterate over array of positions ids
		positions.forEach(function (obj) {
			if ( positionsIds.indexOf(obj._id) > -1 ) {
				// remove the Position
				this.location.positions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more employeesIds as a Employees
	// to a Location
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEmployees( locationId, employeesIds ): Observable<any> {

		// get the Location
		this.loadHelper( locationId );

	// split on a comma with no spaces
	var idList = employeesIds.split(',')

	// iterate over array of employees ids
	idList.forEach(function (id) {
		// read the Employee
		var employee = new EmployeeService(this.http).getEmployee(id);
		// add the Employee if not already assigned
		if ( this.location.employees.indexOf(employee) == -1 )
		this.location.employees.push(employee);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more employeesIds as a Employees
	// from a Location
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEmployees( locationId, employeesIds ): Observable<any> {

		// get the Location
		this.loadHelper( locationId );


	// split on a comma with no spaces
	var idList 					= employeesIds.split(',');
	var employees 	= this.location.employees;

	if ( employees != null && employeesIds != null ) {

		// iterate over array of employees ids
		employees.forEach(function (obj) {
			if ( employeesIds.indexOf(obj._id) > -1 ) {
				// remove the Employee
				this.location.employees.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Location
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Location/update/' + this.location;

	return  this.http.post(uri_, this.location );
}

	//********************************************************************
	// loadHelper - internal helper to load a Location
	//********************************************************************	
	loadHelper( id ) {
		this.getLocation(id)
			.subscribe((res : Location) => {
				this.location = res;
			});
	}
}