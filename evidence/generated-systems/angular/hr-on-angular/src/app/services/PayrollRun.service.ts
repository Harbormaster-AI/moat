import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PayrollRun} from '../models/PayrollRun';
import {PayrollCalendarService} from '../services/PayrollCalendar.service';
import {PayrollItemService} from '../services/PayrollItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PayrollRunService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	payrollRun : PayrollRun;

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
	// add a PayrollRun
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPayrollRun(runNumber, periodStart, periodEnd, paymentDate, PayrollCalendar, PayrollItems, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/PayrollRun/create';
		const obj = {
			      		runNumber: runNumber,
      		periodStart: periodStart,
      		periodEnd: periodEnd,
      		paymentDate: paymentDate,
      		PayrollCalendar: PayrollCalendar != null && PayrollCalendar.length > 0 ? PayrollCalendar : null,
      		PayrollItems: PayrollItems != null && PayrollItems.length > 0 ? PayrollItems : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PayrollRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePayrollRun(runNumber, periodStart, periodEnd, paymentDate, PayrollCalendar, PayrollItems, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PayrollRun/update/' + id;
		const obj = {
				      		runNumber: runNumber,
      		periodStart: periodStart,
      		periodEnd: periodEnd,
      		paymentDate: paymentDate,
      		PayrollCalendar: PayrollCalendar != null && PayrollCalendar.length > 0 ? PayrollCalendar : null,
      		PayrollItems: PayrollItems != null && PayrollItems.length > 0 ? PayrollItems : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PayrollRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePayrollRun(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PayrollRun/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PayrollRun
	// returns the results untouched as an Observable PayrollRun
	// PayrollRun model
	// delegates via URI
	//********************************************************************
	getPayrollRun(id) : Observable<PayrollRun> {
		const uri_ = this.apiUrl + '/PayrollRun/load/' + id;

		return this.http.get<PayrollRun>(uri_);
	}
	
	//********************************************************************
	// gets all PayrollRun
	// returns the results untouched as JSON representation of an
	// Observable array of PayrollRun models
	// delegates via URI
	//********************************************************************
	getPayrollRuns() : Observable<PayrollRun[]> {
		const uri_ = this.apiUrl + '/PayrollRun/';

		return this
			.http.get<PayrollRun[]>(uri_);
	}
	
			//********************************************************************
	// assigns a PayrollCalendar on a PayrollRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPayrollCalendar( payrollRunId, _payrollCalendarId ): Observable<any> {

		// get the PayrollRun from storage
		this.loadHelper( payrollRunId );

	// get the PayrollCalendar from storage
	var tmp 	= new PayrollCalendarService(this.http).getPayrollCalendar(_payrollCalendarId);

	// assign the PayrollCalendar
	this.payrollRun.payrollCalendar = tmp;

	// save the PayrollRun
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PayrollCalendar on a PayrollRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPayrollCalendar( payrollRunId ): Observable<any> {

		// get the PayrollRun from storage
		this.loadHelper( payrollRunId );

	// assign PayrollCalendar to null
	this.payrollRun.payrollCalendar = null;

	// save the PayrollRun
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more payrollItemsIds as a PayrollItems
	// to a PayrollRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPayrollItems( payrollRunId, payrollItemsIds ): Observable<any> {

		// get the PayrollRun
		this.loadHelper( payrollRunId );

	// split on a comma with no spaces
	var idList = payrollItemsIds.split(',')

	// iterate over array of payrollItems ids
	idList.forEach(function (id) {
		// read the PayrollItem
		var payrollItem = new PayrollItemService(this.http).getPayrollItem(id);
		// add the PayrollItem if not already assigned
		if ( this.payrollRun.payrollItems.indexOf(payrollItem) == -1 )
		this.payrollRun.payrollItems.push(payrollItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more payrollItemsIds as a PayrollItems
	// from a PayrollRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePayrollItems( payrollRunId, payrollItemsIds ): Observable<any> {

		// get the PayrollRun
		this.loadHelper( payrollRunId );


	// split on a comma with no spaces
	var idList 					= payrollItemsIds.split(',');
	var payrollItems 	= this.payrollRun.payrollItems;

	if ( payrollItems != null && payrollItemsIds != null ) {

		// iterate over array of payrollItems ids
		payrollItems.forEach(function (obj) {
			if ( payrollItemsIds.indexOf(obj._id) > -1 ) {
				// remove the PayrollItem
				this.payrollRun.payrollItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PayrollRun
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PayrollRun/update/' + this.payrollRun;

	return  this.http.post(uri_, this.payrollRun );
}

	//********************************************************************
	// loadHelper - internal helper to load a PayrollRun
	//********************************************************************	
	loadHelper( id ) {
		this.getPayrollRun(id)
			.subscribe((res : PayrollRun) => {
				this.payrollRun = res;
			});
	}
}