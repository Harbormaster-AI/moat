import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PayrollCalendar} from '../models/PayrollCalendar';
import {OrganizationService} from '../services/Organization.service';
import {PayrollRunService} from '../services/PayrollRun.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PayrollCalendarService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	payrollCalendar : PayrollCalendar;

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
	// add a PayrollCalendar
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPayrollCalendar(name, country, Organization, PayrollRuns, Employees, PayFrequency) : Observable<any> {
		const uri_ = this.apiUrl + '/PayrollCalendar/create';
		const obj = {
			      		name: name,
      		country: country,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		PayrollRuns: PayrollRuns != null && PayrollRuns.length > 0 ? PayrollRuns : null,
      		Employees: Employees != null && Employees.length > 0 ? Employees : null,
			PayFrequency: PayFrequency
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PayrollCalendar
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePayrollCalendar(name, country, Organization, PayrollRuns, Employees, PayFrequency, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PayrollCalendar/update/' + id;
		const obj = {
				      		name: name,
      		country: country,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		PayrollRuns: PayrollRuns != null && PayrollRuns.length > 0 ? PayrollRuns : null,
      		Employees: Employees != null && Employees.length > 0 ? Employees : null,
			PayFrequency: PayFrequency
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PayrollCalendar
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePayrollCalendar(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PayrollCalendar/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PayrollCalendar
	// returns the results untouched as an Observable PayrollCalendar
	// PayrollCalendar model
	// delegates via URI
	//********************************************************************
	getPayrollCalendar(id) : Observable<PayrollCalendar> {
		const uri_ = this.apiUrl + '/PayrollCalendar/load/' + id;

		return this.http.get<PayrollCalendar>(uri_);
	}
	
	//********************************************************************
	// gets all PayrollCalendar
	// returns the results untouched as JSON representation of an
	// Observable array of PayrollCalendar models
	// delegates via URI
	//********************************************************************
	getPayrollCalendars() : Observable<PayrollCalendar[]> {
		const uri_ = this.apiUrl + '/PayrollCalendar/';

		return this
			.http.get<PayrollCalendar[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a PayrollCalendar
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( payrollCalendarId, _organizationId ): Observable<any> {

		// get the PayrollCalendar from storage
		this.loadHelper( payrollCalendarId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.payrollCalendar.organization = tmp;

	// save the PayrollCalendar
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a PayrollCalendar
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( payrollCalendarId ): Observable<any> {

		// get the PayrollCalendar from storage
		this.loadHelper( payrollCalendarId );

	// assign Organization to null
	this.payrollCalendar.organization = null;

	// save the PayrollCalendar
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more payrollRunsIds as a PayrollRuns
	// to a PayrollCalendar
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPayrollRuns( payrollCalendarId, payrollRunsIds ): Observable<any> {

		// get the PayrollCalendar
		this.loadHelper( payrollCalendarId );

	// split on a comma with no spaces
	var idList = payrollRunsIds.split(',')

	// iterate over array of payrollRuns ids
	idList.forEach(function (id) {
		// read the PayrollRun
		var payrollRun = new PayrollRunService(this.http).getPayrollRun(id);
		// add the PayrollRun if not already assigned
		if ( this.payrollCalendar.payrollRuns.indexOf(payrollRun) == -1 )
		this.payrollCalendar.payrollRuns.push(payrollRun);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more payrollRunsIds as a PayrollRuns
	// from a PayrollCalendar
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePayrollRuns( payrollCalendarId, payrollRunsIds ): Observable<any> {

		// get the PayrollCalendar
		this.loadHelper( payrollCalendarId );


	// split on a comma with no spaces
	var idList 					= payrollRunsIds.split(',');
	var payrollRuns 	= this.payrollCalendar.payrollRuns;

	if ( payrollRuns != null && payrollRunsIds != null ) {

		// iterate over array of payrollRuns ids
		payrollRuns.forEach(function (obj) {
			if ( payrollRunsIds.indexOf(obj._id) > -1 ) {
				// remove the PayrollRun
				this.payrollCalendar.payrollRuns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more employeesIds as a Employees
	// to a PayrollCalendar
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEmployees( payrollCalendarId, employeesIds ): Observable<any> {

		// get the PayrollCalendar
		this.loadHelper( payrollCalendarId );

	// split on a comma with no spaces
	var idList = employeesIds.split(',')

	// iterate over array of employees ids
	idList.forEach(function (id) {
		// read the Employee
		var employee = new EmployeeService(this.http).getEmployee(id);
		// add the Employee if not already assigned
		if ( this.payrollCalendar.employees.indexOf(employee) == -1 )
		this.payrollCalendar.employees.push(employee);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more employeesIds as a Employees
	// from a PayrollCalendar
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEmployees( payrollCalendarId, employeesIds ): Observable<any> {

		// get the PayrollCalendar
		this.loadHelper( payrollCalendarId );


	// split on a comma with no spaces
	var idList 					= employeesIds.split(',');
	var employees 	= this.payrollCalendar.employees;

	if ( employees != null && employeesIds != null ) {

		// iterate over array of employees ids
		employees.forEach(function (obj) {
			if ( employeesIds.indexOf(obj._id) > -1 ) {
				// remove the Employee
				this.payrollCalendar.employees.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PayrollCalendar
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PayrollCalendar/update/' + this.payrollCalendar;

	return  this.http.post(uri_, this.payrollCalendar );
}

	//********************************************************************
	// loadHelper - internal helper to load a PayrollCalendar
	//********************************************************************	
	loadHelper( id ) {
		this.getPayrollCalendar(id)
			.subscribe((res : PayrollCalendar) => {
				this.payrollCalendar = res;
			});
	}
}