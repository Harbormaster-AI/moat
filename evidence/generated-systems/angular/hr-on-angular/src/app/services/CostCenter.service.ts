import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CostCenter} from '../models/CostCenter';
import {OrganizationService} from '../services/Organization.service';
import {DepartmentService} from '../services/Department.service';
import {PositionService} from '../services/Position.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CostCenterService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	costCenter : CostCenter;

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
	// add a CostCenter
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCostCenter(code, name, Organization, Departments, Positions, Employees) : Observable<any> {
		const uri_ = this.apiUrl + '/CostCenter/create';
		const obj = {
			      		code: code,
      		name: name,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Departments: Departments != null && Departments.length > 0 ? Departments : null,
      		Positions: Positions != null && Positions.length > 0 ? Positions : null,
			Employees: Employees != null && Employees.length > 0 ? Employees : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CostCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCostCenter(code, name, Organization, Departments, Positions, Employees, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CostCenter/update/' + id;
		const obj = {
				      		code: code,
      		name: name,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Departments: Departments != null && Departments.length > 0 ? Departments : null,
      		Positions: Positions != null && Positions.length > 0 ? Positions : null,
			Employees: Employees != null && Employees.length > 0 ? Employees : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CostCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCostCenter(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CostCenter/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CostCenter
	// returns the results untouched as an Observable CostCenter
	// CostCenter model
	// delegates via URI
	//********************************************************************
	getCostCenter(id) : Observable<CostCenter> {
		const uri_ = this.apiUrl + '/CostCenter/load/' + id;

		return this.http.get<CostCenter>(uri_);
	}
	
	//********************************************************************
	// gets all CostCenter
	// returns the results untouched as JSON representation of an
	// Observable array of CostCenter models
	// delegates via URI
	//********************************************************************
	getCostCenters() : Observable<CostCenter[]> {
		const uri_ = this.apiUrl + '/CostCenter/';

		return this
			.http.get<CostCenter[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a CostCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( costCenterId, _organizationId ): Observable<any> {

		// get the CostCenter from storage
		this.loadHelper( costCenterId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.costCenter.organization = tmp;

	// save the CostCenter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a CostCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( costCenterId ): Observable<any> {

		// get the CostCenter from storage
		this.loadHelper( costCenterId );

	// assign Organization to null
	this.costCenter.organization = null;

	// save the CostCenter
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more departmentsIds as a Departments
	// to a CostCenter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDepartments( costCenterId, departmentsIds ): Observable<any> {

		// get the CostCenter
		this.loadHelper( costCenterId );

	// split on a comma with no spaces
	var idList = departmentsIds.split(',')

	// iterate over array of departments ids
	idList.forEach(function (id) {
		// read the Department
		var department = new DepartmentService(this.http).getDepartment(id);
		// add the Department if not already assigned
		if ( this.costCenter.departments.indexOf(department) == -1 )
		this.costCenter.departments.push(department);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more departmentsIds as a Departments
	// from a CostCenter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDepartments( costCenterId, departmentsIds ): Observable<any> {

		// get the CostCenter
		this.loadHelper( costCenterId );


	// split on a comma with no spaces
	var idList 					= departmentsIds.split(',');
	var departments 	= this.costCenter.departments;

	if ( departments != null && departmentsIds != null ) {

		// iterate over array of departments ids
		departments.forEach(function (obj) {
			if ( departmentsIds.indexOf(obj._id) > -1 ) {
				// remove the Department
				this.costCenter.departments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more positionsIds as a Positions
	// to a CostCenter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPositions( costCenterId, positionsIds ): Observable<any> {

		// get the CostCenter
		this.loadHelper( costCenterId );

	// split on a comma with no spaces
	var idList = positionsIds.split(',')

	// iterate over array of positions ids
	idList.forEach(function (id) {
		// read the Position
		var position = new PositionService(this.http).getPosition(id);
		// add the Position if not already assigned
		if ( this.costCenter.positions.indexOf(position) == -1 )
		this.costCenter.positions.push(position);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more positionsIds as a Positions
	// from a CostCenter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePositions( costCenterId, positionsIds ): Observable<any> {

		// get the CostCenter
		this.loadHelper( costCenterId );


	// split on a comma with no spaces
	var idList 					= positionsIds.split(',');
	var positions 	= this.costCenter.positions;

	if ( positions != null && positionsIds != null ) {

		// iterate over array of positions ids
		positions.forEach(function (obj) {
			if ( positionsIds.indexOf(obj._id) > -1 ) {
				// remove the Position
				this.costCenter.positions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more employeesIds as a Employees
	// to a CostCenter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEmployees( costCenterId, employeesIds ): Observable<any> {

		// get the CostCenter
		this.loadHelper( costCenterId );

	// split on a comma with no spaces
	var idList = employeesIds.split(',')

	// iterate over array of employees ids
	idList.forEach(function (id) {
		// read the Employee
		var employee = new EmployeeService(this.http).getEmployee(id);
		// add the Employee if not already assigned
		if ( this.costCenter.employees.indexOf(employee) == -1 )
		this.costCenter.employees.push(employee);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more employeesIds as a Employees
	// from a CostCenter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEmployees( costCenterId, employeesIds ): Observable<any> {

		// get the CostCenter
		this.loadHelper( costCenterId );


	// split on a comma with no spaces
	var idList 					= employeesIds.split(',');
	var employees 	= this.costCenter.employees;

	if ( employees != null && employeesIds != null ) {

		// iterate over array of employees ids
		employees.forEach(function (obj) {
			if ( employeesIds.indexOf(obj._id) > -1 ) {
				// remove the Employee
				this.costCenter.employees.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a CostCenter
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CostCenter/update/' + this.costCenter;

	return  this.http.post(uri_, this.costCenter );
}

	//********************************************************************
	// loadHelper - internal helper to load a CostCenter
	//********************************************************************	
	loadHelper( id ) {
		this.getCostCenter(id)
			.subscribe((res : CostCenter) => {
				this.costCenter = res;
			});
	}
}