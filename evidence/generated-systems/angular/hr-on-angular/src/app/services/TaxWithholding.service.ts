import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TaxWithholding} from '../models/TaxWithholding';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TaxWithholdingService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	taxWithholding : TaxWithholding;

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
	// add a TaxWithholding
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTaxWithholding(taxId, allowances, additionalAmount, Employee, FilingStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/TaxWithholding/create';
		const obj = {
			      		taxId: taxId,
      		allowances: allowances,
      		additionalAmount: additionalAmount,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			FilingStatus: FilingStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TaxWithholding
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTaxWithholding(taxId, allowances, additionalAmount, Employee, FilingStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TaxWithholding/update/' + id;
		const obj = {
				      		taxId: taxId,
      		allowances: allowances,
      		additionalAmount: additionalAmount,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			FilingStatus: FilingStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TaxWithholding
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTaxWithholding(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TaxWithholding/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TaxWithholding
	// returns the results untouched as an Observable TaxWithholding
	// TaxWithholding model
	// delegates via URI
	//********************************************************************
	getTaxWithholding(id) : Observable<TaxWithholding> {
		const uri_ = this.apiUrl + '/TaxWithholding/load/' + id;

		return this.http.get<TaxWithholding>(uri_);
	}
	
	//********************************************************************
	// gets all TaxWithholding
	// returns the results untouched as JSON representation of an
	// Observable array of TaxWithholding models
	// delegates via URI
	//********************************************************************
	getTaxWithholdings() : Observable<TaxWithholding[]> {
		const uri_ = this.apiUrl + '/TaxWithholding/';

		return this
			.http.get<TaxWithholding[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a TaxWithholding
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( taxWithholdingId, _employeeId ): Observable<any> {

		// get the TaxWithholding from storage
		this.loadHelper( taxWithholdingId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.taxWithholding.employee = tmp;

	// save the TaxWithholding
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a TaxWithholding
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( taxWithholdingId ): Observable<any> {

		// get the TaxWithholding from storage
		this.loadHelper( taxWithholdingId );

	// assign Employee to null
	this.taxWithholding.employee = null;

	// save the TaxWithholding
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a TaxWithholding
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TaxWithholding/update/' + this.taxWithholding;

	return  this.http.post(uri_, this.taxWithholding );
}

	//********************************************************************
	// loadHelper - internal helper to load a TaxWithholding
	//********************************************************************	
	loadHelper( id ) {
		this.getTaxWithholding(id)
			.subscribe((res : TaxWithholding) => {
				this.taxWithholding = res;
			});
	}
}