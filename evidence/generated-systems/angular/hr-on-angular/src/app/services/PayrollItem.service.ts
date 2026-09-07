import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PayrollItem} from '../models/PayrollItem';
import {PayrollRunService} from '../services/PayrollRun.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PayrollItemService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	payrollItem : PayrollItem;

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
	// add a PayrollItem
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPayrollItem(amount, taxable, PayrollRun, Employee, ItemType) : Observable<any> {
		const uri_ = this.apiUrl + '/PayrollItem/create';
		const obj = {
			      		amount: amount,
      		taxable: taxable,
      		PayrollRun: PayrollRun != null && PayrollRun.length > 0 ? PayrollRun : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			ItemType: ItemType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PayrollItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePayrollItem(amount, taxable, PayrollRun, Employee, ItemType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PayrollItem/update/' + id;
		const obj = {
				      		amount: amount,
      		taxable: taxable,
      		PayrollRun: PayrollRun != null && PayrollRun.length > 0 ? PayrollRun : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			ItemType: ItemType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PayrollItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePayrollItem(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PayrollItem/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PayrollItem
	// returns the results untouched as an Observable PayrollItem
	// PayrollItem model
	// delegates via URI
	//********************************************************************
	getPayrollItem(id) : Observable<PayrollItem> {
		const uri_ = this.apiUrl + '/PayrollItem/load/' + id;

		return this.http.get<PayrollItem>(uri_);
	}
	
	//********************************************************************
	// gets all PayrollItem
	// returns the results untouched as JSON representation of an
	// Observable array of PayrollItem models
	// delegates via URI
	//********************************************************************
	getPayrollItems() : Observable<PayrollItem[]> {
		const uri_ = this.apiUrl + '/PayrollItem/';

		return this
			.http.get<PayrollItem[]>(uri_);
	}
	
			//********************************************************************
	// assigns a PayrollRun on a PayrollItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPayrollRun( payrollItemId, _payrollRunId ): Observable<any> {

		// get the PayrollItem from storage
		this.loadHelper( payrollItemId );

	// get the PayrollRun from storage
	var tmp 	= new PayrollRunService(this.http).getPayrollRun(_payrollRunId);

	// assign the PayrollRun
	this.payrollItem.payrollRun = tmp;

	// save the PayrollItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PayrollRun on a PayrollItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPayrollRun( payrollItemId ): Observable<any> {

		// get the PayrollItem from storage
		this.loadHelper( payrollItemId );

	// assign PayrollRun to null
	this.payrollItem.payrollRun = null;

	// save the PayrollItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Employee on a PayrollItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( payrollItemId, _employeeId ): Observable<any> {

		// get the PayrollItem from storage
		this.loadHelper( payrollItemId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.payrollItem.employee = tmp;

	// save the PayrollItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a PayrollItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( payrollItemId ): Observable<any> {

		// get the PayrollItem from storage
		this.loadHelper( payrollItemId );

	// assign Employee to null
	this.payrollItem.employee = null;

	// save the PayrollItem
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a PayrollItem
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PayrollItem/update/' + this.payrollItem;

	return  this.http.post(uri_, this.payrollItem );
}

	//********************************************************************
	// loadHelper - internal helper to load a PayrollItem
	//********************************************************************	
	loadHelper( id ) {
		this.getPayrollItem(id)
			.subscribe((res : PayrollItem) => {
				this.payrollItem = res;
			});
	}
}