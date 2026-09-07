import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {RepaymentSchedule} from '../models/RepaymentSchedule';
import {LoanService} from '../services/Loan.service';
import {TransactionService} from '../services/Transaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RepaymentScheduleService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	repaymentSchedule : RepaymentSchedule;

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
	// add a RepaymentSchedule
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRepaymentSchedule(installmentNumber, dueDate, amountDue, principalDue, interestDue, Loan, Payments, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/RepaymentSchedule/create';
		const obj = {
			      		installmentNumber: installmentNumber,
      		dueDate: dueDate,
      		amountDue: amountDue,
      		principalDue: principalDue,
      		interestDue: interestDue,
      		Loan: Loan != null && Loan.length > 0 ? Loan : null,
      		Payments: Payments != null && Payments.length > 0 ? Payments : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a RepaymentSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRepaymentSchedule(installmentNumber, dueDate, amountDue, principalDue, interestDue, Loan, Payments, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/RepaymentSchedule/update/' + id;
		const obj = {
				      		installmentNumber: installmentNumber,
      		dueDate: dueDate,
      		amountDue: amountDue,
      		principalDue: principalDue,
      		interestDue: interestDue,
      		Loan: Loan != null && Loan.length > 0 ? Loan : null,
      		Payments: Payments != null && Payments.length > 0 ? Payments : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a RepaymentSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRepaymentSchedule(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/RepaymentSchedule/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a RepaymentSchedule
	// returns the results untouched as an Observable RepaymentSchedule
	// RepaymentSchedule model
	// delegates via URI
	//********************************************************************
	getRepaymentSchedule(id) : Observable<RepaymentSchedule> {
		const uri_ = this.apiUrl + '/RepaymentSchedule/load/' + id;

		return this.http.get<RepaymentSchedule>(uri_);
	}
	
	//********************************************************************
	// gets all RepaymentSchedule
	// returns the results untouched as JSON representation of an
	// Observable array of RepaymentSchedule models
	// delegates via URI
	//********************************************************************
	getRepaymentSchedules() : Observable<RepaymentSchedule[]> {
		const uri_ = this.apiUrl + '/RepaymentSchedule/';

		return this
			.http.get<RepaymentSchedule[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Loan on a RepaymentSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLoan( repaymentScheduleId, _loanId ): Observable<any> {

		// get the RepaymentSchedule from storage
		this.loadHelper( repaymentScheduleId );

	// get the Loan from storage
	var tmp 	= new LoanService(this.http).getLoan(_loanId);

	// assign the Loan
	this.repaymentSchedule.loan = tmp;

	// save the RepaymentSchedule
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Loan on a RepaymentSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLoan( repaymentScheduleId ): Observable<any> {

		// get the RepaymentSchedule from storage
		this.loadHelper( repaymentScheduleId );

	// assign Loan to null
	this.repaymentSchedule.loan = null;

	// save the RepaymentSchedule
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more paymentsIds as a Payments
	// to a RepaymentSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPayments( repaymentScheduleId, paymentsIds ): Observable<any> {

		// get the RepaymentSchedule
		this.loadHelper( repaymentScheduleId );

	// split on a comma with no spaces
	var idList = paymentsIds.split(',')

	// iterate over array of payments ids
	idList.forEach(function (id) {
		// read the Transaction
		var transaction = new TransactionService(this.http).getTransaction(id);
		// add the Transaction if not already assigned
		if ( this.repaymentSchedule.payments.indexOf(transaction) == -1 )
		this.repaymentSchedule.payments.push(transaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more paymentsIds as a Payments
	// from a RepaymentSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePayments( repaymentScheduleId, paymentsIds ): Observable<any> {

		// get the RepaymentSchedule
		this.loadHelper( repaymentScheduleId );


	// split on a comma with no spaces
	var idList 					= paymentsIds.split(',');
	var payments 	= this.repaymentSchedule.payments;

	if ( payments != null && paymentsIds != null ) {

		// iterate over array of payments ids
		payments.forEach(function (obj) {
			if ( paymentsIds.indexOf(obj._id) > -1 ) {
				// remove the Transaction
				this.repaymentSchedule.payments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a RepaymentSchedule
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/RepaymentSchedule/update/' + this.repaymentSchedule;

	return  this.http.post(uri_, this.repaymentSchedule );
}

	//********************************************************************
	// loadHelper - internal helper to load a RepaymentSchedule
	//********************************************************************	
	loadHelper( id ) {
		this.getRepaymentSchedule(id)
			.subscribe((res : RepaymentSchedule) => {
				this.repaymentSchedule = res;
			});
	}
}