import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PaymentMethod} from '../models/PaymentMethod';
import {EmployeeService} from '../services/Employee.service';
import {BankAccountService} from '../services/BankAccount.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PaymentMethodService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	paymentMethod : PaymentMethod;

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
	// add a PaymentMethod
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPaymentMethod(preferred, Employee, BankAccount, MethodType) : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentMethod/create';
		const obj = {
			      		preferred: preferred,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		BankAccount: BankAccount != null && BankAccount.length > 0 ? BankAccount : null,
			MethodType: MethodType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PaymentMethod
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePaymentMethod(preferred, Employee, BankAccount, MethodType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PaymentMethod/update/' + id;
		const obj = {
				      		preferred: preferred,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		BankAccount: BankAccount != null && BankAccount.length > 0 ? BankAccount : null,
			MethodType: MethodType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PaymentMethod
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePaymentMethod(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentMethod/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PaymentMethod
	// returns the results untouched as an Observable PaymentMethod
	// PaymentMethod model
	// delegates via URI
	//********************************************************************
	getPaymentMethod(id) : Observable<PaymentMethod> {
		const uri_ = this.apiUrl + '/PaymentMethod/load/' + id;

		return this.http.get<PaymentMethod>(uri_);
	}
	
	//********************************************************************
	// gets all PaymentMethod
	// returns the results untouched as JSON representation of an
	// Observable array of PaymentMethod models
	// delegates via URI
	//********************************************************************
	getPaymentMethods() : Observable<PaymentMethod[]> {
		const uri_ = this.apiUrl + '/PaymentMethod/';

		return this
			.http.get<PaymentMethod[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a PaymentMethod
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( paymentMethodId, _employeeId ): Observable<any> {

		// get the PaymentMethod from storage
		this.loadHelper( paymentMethodId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.paymentMethod.employee = tmp;

	// save the PaymentMethod
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a PaymentMethod
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( paymentMethodId ): Observable<any> {

		// get the PaymentMethod from storage
		this.loadHelper( paymentMethodId );

	// assign Employee to null
	this.paymentMethod.employee = null;

	// save the PaymentMethod
	return this.saveHelper();
}

		//********************************************************************
	// assigns a BankAccount on a PaymentMethod
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBankAccount( paymentMethodId, _bankAccountId ): Observable<any> {

		// get the PaymentMethod from storage
		this.loadHelper( paymentMethodId );

	// get the BankAccount from storage
	var tmp 	= new BankAccountService(this.http).getBankAccount(_bankAccountId);

	// assign the BankAccount
	this.paymentMethod.bankAccount = tmp;

	// save the PaymentMethod
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a BankAccount on a PaymentMethod
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBankAccount( paymentMethodId ): Observable<any> {

		// get the PaymentMethod from storage
		this.loadHelper( paymentMethodId );

	// assign BankAccount to null
	this.paymentMethod.bankAccount = null;

	// save the PaymentMethod
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a PaymentMethod
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PaymentMethod/update/' + this.paymentMethod;

	return  this.http.post(uri_, this.paymentMethod );
}

	//********************************************************************
	// loadHelper - internal helper to load a PaymentMethod
	//********************************************************************	
	loadHelper( id ) {
		this.getPaymentMethod(id)
			.subscribe((res : PaymentMethod) => {
				this.paymentMethod = res;
			});
	}
}