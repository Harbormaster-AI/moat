import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Department} from '../models/Department';
import {OrganizationService} from '../services/Organization.service';
import {EmployeeService} from '../services/Employee.service';
import {PositionService} from '../services/Position.service';
import {CostCenterService} from '../services/CostCenter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DepartmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	department : Department;

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
	// add a Department
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDepartment(name, code, Organization, Manager, Positions, Employees, CostCenter) : Observable<any> {
		const uri_ = this.apiUrl + '/Department/create';
		const obj = {
			      		name: name,
      		code: code,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Manager: Manager != null && Manager.length > 0 ? Manager : null,
      		Positions: Positions != null && Positions.length > 0 ? Positions : null,
      		Employees: Employees != null && Employees.length > 0 ? Employees : null,
			CostCenter: CostCenter != null && CostCenter.length > 0 ? CostCenter : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDepartment(name, code, Organization, Manager, Positions, Employees, CostCenter, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Department/update/' + id;
		const obj = {
				      		name: name,
      		code: code,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Manager: Manager != null && Manager.length > 0 ? Manager : null,
      		Positions: Positions != null && Positions.length > 0 ? Positions : null,
      		Employees: Employees != null && Employees.length > 0 ? Employees : null,
			CostCenter: CostCenter != null && CostCenter.length > 0 ? CostCenter : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDepartment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Department/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Department
	// returns the results untouched as an Observable Department
	// Department model
	// delegates via URI
	//********************************************************************
	getDepartment(id) : Observable<Department> {
		const uri_ = this.apiUrl + '/Department/load/' + id;

		return this.http.get<Department>(uri_);
	}
	
	//********************************************************************
	// gets all Department
	// returns the results untouched as JSON representation of an
	// Observable array of Department models
	// delegates via URI
	//********************************************************************
	getDepartments() : Observable<Department[]> {
		const uri_ = this.apiUrl + '/Department/';

		return this
			.http.get<Department[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( departmentId, _organizationId ): Observable<any> {

		// get the Department from storage
		this.loadHelper( departmentId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.department.organization = tmp;

	// save the Department
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( departmentId ): Observable<any> {

		// get the Department from storage
		this.loadHelper( departmentId );

	// assign Organization to null
	this.department.organization = null;

	// save the Department
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Manager on a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignManager( departmentId, _managerId ): Observable<any> {

		// get the Department from storage
		this.loadHelper( departmentId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_managerId);

	// assign the Manager
	this.department.manager = tmp;

	// save the Department
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Manager on a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignManager( departmentId ): Observable<any> {

		// get the Department from storage
		this.loadHelper( departmentId );

	// assign Manager to null
	this.department.manager = null;

	// save the Department
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CostCenter on a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCostCenter( departmentId, _costCenterId ): Observable<any> {

		// get the Department from storage
		this.loadHelper( departmentId );

	// get the CostCenter from storage
	var tmp 	= new CostCenterService(this.http).getCostCenter(_costCenterId);

	// assign the CostCenter
	this.department.costCenter = tmp;

	// save the Department
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CostCenter on a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCostCenter( departmentId ): Observable<any> {

		// get the Department from storage
		this.loadHelper( departmentId );

	// assign CostCenter to null
	this.department.costCenter = null;

	// save the Department
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more positionsIds as a Positions
	// to a Department
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPositions( departmentId, positionsIds ): Observable<any> {

		// get the Department
		this.loadHelper( departmentId );

	// split on a comma with no spaces
	var idList = positionsIds.split(',')

	// iterate over array of positions ids
	idList.forEach(function (id) {
		// read the Position
		var position = new PositionService(this.http).getPosition(id);
		// add the Position if not already assigned
		if ( this.department.positions.indexOf(position) == -1 )
		this.department.positions.push(position);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more positionsIds as a Positions
	// from a Department
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePositions( departmentId, positionsIds ): Observable<any> {

		// get the Department
		this.loadHelper( departmentId );


	// split on a comma with no spaces
	var idList 					= positionsIds.split(',');
	var positions 	= this.department.positions;

	if ( positions != null && positionsIds != null ) {

		// iterate over array of positions ids
		positions.forEach(function (obj) {
			if ( positionsIds.indexOf(obj._id) > -1 ) {
				// remove the Position
				this.department.positions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more employeesIds as a Employees
	// to a Department
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEmployees( departmentId, employeesIds ): Observable<any> {

		// get the Department
		this.loadHelper( departmentId );

	// split on a comma with no spaces
	var idList = employeesIds.split(',')

	// iterate over array of employees ids
	idList.forEach(function (id) {
		// read the Employee
		var employee = new EmployeeService(this.http).getEmployee(id);
		// add the Employee if not already assigned
		if ( this.department.employees.indexOf(employee) == -1 )
		this.department.employees.push(employee);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more employeesIds as a Employees
	// from a Department
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEmployees( departmentId, employeesIds ): Observable<any> {

		// get the Department
		this.loadHelper( departmentId );


	// split on a comma with no spaces
	var idList 					= employeesIds.split(',');
	var employees 	= this.department.employees;

	if ( employees != null && employeesIds != null ) {

		// iterate over array of employees ids
		employees.forEach(function (obj) {
			if ( employeesIds.indexOf(obj._id) > -1 ) {
				// remove the Employee
				this.department.employees.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Department
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Department/update/' + this.department;

	return  this.http.post(uri_, this.department );
}

	//********************************************************************
	// loadHelper - internal helper to load a Department
	//********************************************************************	
	loadHelper( id ) {
		this.getDepartment(id)
			.subscribe((res : Department) => {
				this.department = res;
			});
	}
}