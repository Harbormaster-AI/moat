import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {EmploymentContract} from '../models/EmploymentContract';
import {EmployeeService} from '../services/Employee.service';
import {CompensationPackageService} from '../services/CompensationPackage.service';
import {WorkScheduleService} from '../services/WorkSchedule.service';
import {LocationService} from '../services/Location.service';
import {PayrollCalendarService} from '../services/PayrollCalendar.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EmploymentContractService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	employmentContract : EmploymentContract;

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
	// add a EmploymentContract
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEmploymentContract(contractNumber, startDate, endDate, workHoursPerWeek, Employee, CompensationPackage, WorkSchedule, Location, PayrollCalendar, EmploymentType, Status, PayFrequency) : Observable<any> {
		const uri_ = this.apiUrl + '/EmploymentContract/create';
		const obj = {
			      		contractNumber: contractNumber,
      		startDate: startDate,
      		endDate: endDate,
      		workHoursPerWeek: workHoursPerWeek,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		CompensationPackage: CompensationPackage != null && CompensationPackage.length > 0 ? CompensationPackage : null,
      		WorkSchedule: WorkSchedule != null && WorkSchedule.length > 0 ? WorkSchedule : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		PayrollCalendar: PayrollCalendar != null && PayrollCalendar.length > 0 ? PayrollCalendar : null,
      		EmploymentType: EmploymentType,
      		Status: Status,
			PayFrequency: PayFrequency
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEmploymentContract(contractNumber, startDate, endDate, workHoursPerWeek, Employee, CompensationPackage, WorkSchedule, Location, PayrollCalendar, EmploymentType, Status, PayFrequency, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/EmploymentContract/update/' + id;
		const obj = {
				      		contractNumber: contractNumber,
      		startDate: startDate,
      		endDate: endDate,
      		workHoursPerWeek: workHoursPerWeek,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		CompensationPackage: CompensationPackage != null && CompensationPackage.length > 0 ? CompensationPackage : null,
      		WorkSchedule: WorkSchedule != null && WorkSchedule.length > 0 ? WorkSchedule : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		PayrollCalendar: PayrollCalendar != null && PayrollCalendar.length > 0 ? PayrollCalendar : null,
      		EmploymentType: EmploymentType,
      		Status: Status,
			PayFrequency: PayFrequency
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEmploymentContract(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/EmploymentContract/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a EmploymentContract
	// returns the results untouched as an Observable EmploymentContract
	// EmploymentContract model
	// delegates via URI
	//********************************************************************
	getEmploymentContract(id) : Observable<EmploymentContract> {
		const uri_ = this.apiUrl + '/EmploymentContract/load/' + id;

		return this.http.get<EmploymentContract>(uri_);
	}
	
	//********************************************************************
	// gets all EmploymentContract
	// returns the results untouched as JSON representation of an
	// Observable array of EmploymentContract models
	// delegates via URI
	//********************************************************************
	getEmploymentContracts() : Observable<EmploymentContract[]> {
		const uri_ = this.apiUrl + '/EmploymentContract/';

		return this
			.http.get<EmploymentContract[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( employmentContractId, _employeeId ): Observable<any> {

		// get the EmploymentContract from storage
		this.loadHelper( employmentContractId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.employmentContract.employee = tmp;

	// save the EmploymentContract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( employmentContractId ): Observable<any> {

		// get the EmploymentContract from storage
		this.loadHelper( employmentContractId );

	// assign Employee to null
	this.employmentContract.employee = null;

	// save the EmploymentContract
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CompensationPackage on a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCompensationPackage( employmentContractId, _compensationPackageId ): Observable<any> {

		// get the EmploymentContract from storage
		this.loadHelper( employmentContractId );

	// get the CompensationPackage from storage
	var tmp 	= new CompensationPackageService(this.http).getCompensationPackage(_compensationPackageId);

	// assign the CompensationPackage
	this.employmentContract.compensationPackage = tmp;

	// save the EmploymentContract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CompensationPackage on a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCompensationPackage( employmentContractId ): Observable<any> {

		// get the EmploymentContract from storage
		this.loadHelper( employmentContractId );

	// assign CompensationPackage to null
	this.employmentContract.compensationPackage = null;

	// save the EmploymentContract
	return this.saveHelper();
}

		//********************************************************************
	// assigns a WorkSchedule on a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkSchedule( employmentContractId, _workScheduleId ): Observable<any> {

		// get the EmploymentContract from storage
		this.loadHelper( employmentContractId );

	// get the WorkSchedule from storage
	var tmp 	= new WorkScheduleService(this.http).getWorkSchedule(_workScheduleId);

	// assign the WorkSchedule
	this.employmentContract.workSchedule = tmp;

	// save the EmploymentContract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkSchedule on a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkSchedule( employmentContractId ): Observable<any> {

		// get the EmploymentContract from storage
		this.loadHelper( employmentContractId );

	// assign WorkSchedule to null
	this.employmentContract.workSchedule = null;

	// save the EmploymentContract
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Location on a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLocation( employmentContractId, _locationId ): Observable<any> {

		// get the EmploymentContract from storage
		this.loadHelper( employmentContractId );

	// get the Location from storage
	var tmp 	= new LocationService(this.http).getLocation(_locationId);

	// assign the Location
	this.employmentContract.location = tmp;

	// save the EmploymentContract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Location on a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLocation( employmentContractId ): Observable<any> {

		// get the EmploymentContract from storage
		this.loadHelper( employmentContractId );

	// assign Location to null
	this.employmentContract.location = null;

	// save the EmploymentContract
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PayrollCalendar on a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPayrollCalendar( employmentContractId, _payrollCalendarId ): Observable<any> {

		// get the EmploymentContract from storage
		this.loadHelper( employmentContractId );

	// get the PayrollCalendar from storage
	var tmp 	= new PayrollCalendarService(this.http).getPayrollCalendar(_payrollCalendarId);

	// assign the PayrollCalendar
	this.employmentContract.payrollCalendar = tmp;

	// save the EmploymentContract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PayrollCalendar on a EmploymentContract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPayrollCalendar( employmentContractId ): Observable<any> {

		// get the EmploymentContract from storage
		this.loadHelper( employmentContractId );

	// assign PayrollCalendar to null
	this.employmentContract.payrollCalendar = null;

	// save the EmploymentContract
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a EmploymentContract
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/EmploymentContract/update/' + this.employmentContract;

	return  this.http.post(uri_, this.employmentContract );
}

	//********************************************************************
	// loadHelper - internal helper to load a EmploymentContract
	//********************************************************************	
	loadHelper( id ) {
		this.getEmploymentContract(id)
			.subscribe((res : EmploymentContract) => {
				this.employmentContract = res;
			});
	}
}