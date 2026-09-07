import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Loan} from '../models/Loan';
import {CustomerService} from '../services/Customer.service';
import {RepaymentScheduleService} from '../services/RepaymentSchedule.service';
import {CollateralService} from '../services/Collateral.service';
import {LoanTransactionService} from '../services/LoanTransaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LoanService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	loan : Loan;

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
	// add a Loan
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLoan(loanNumber, principal, interestRate, originationDate, maturityDate, Customer, Schedule, Collateral, Transactions, RateType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Loan/create';
		const obj = {
			      		loanNumber: loanNumber,
      		principal: principal,
      		interestRate: interestRate,
      		originationDate: originationDate,
      		maturityDate: maturityDate,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Schedule: Schedule != null && Schedule.length > 0 ? Schedule : null,
      		Collateral: Collateral != null && Collateral.length > 0 ? Collateral : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
      		RateType: RateType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Loan
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLoan(loanNumber, principal, interestRate, originationDate, maturityDate, Customer, Schedule, Collateral, Transactions, RateType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Loan/update/' + id;
		const obj = {
				      		loanNumber: loanNumber,
      		principal: principal,
      		interestRate: interestRate,
      		originationDate: originationDate,
      		maturityDate: maturityDate,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Schedule: Schedule != null && Schedule.length > 0 ? Schedule : null,
      		Collateral: Collateral != null && Collateral.length > 0 ? Collateral : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
      		RateType: RateType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Loan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLoan(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Loan/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Loan
	// returns the results untouched as an Observable Loan
	// Loan model
	// delegates via URI
	//********************************************************************
	getLoan(id) : Observable<Loan> {
		const uri_ = this.apiUrl + '/Loan/load/' + id;

		return this.http.get<Loan>(uri_);
	}
	
	//********************************************************************
	// gets all Loan
	// returns the results untouched as JSON representation of an
	// Observable array of Loan models
	// delegates via URI
	//********************************************************************
	getLoans() : Observable<Loan[]> {
		const uri_ = this.apiUrl + '/Loan/';

		return this
			.http.get<Loan[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a Loan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( loanId, _customerId ): Observable<any> {

		// get the Loan from storage
		this.loadHelper( loanId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.loan.customer = tmp;

	// save the Loan
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a Loan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( loanId ): Observable<any> {

		// get the Loan from storage
		this.loadHelper( loanId );

	// assign Customer to null
	this.loan.customer = null;

	// save the Loan
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more scheduleIds as a Schedule
	// to a Loan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSchedule( loanId, scheduleIds ): Observable<any> {

		// get the Loan
		this.loadHelper( loanId );

	// split on a comma with no spaces
	var idList = scheduleIds.split(',')

	// iterate over array of schedule ids
	idList.forEach(function (id) {
		// read the RepaymentSchedule
		var repaymentSchedule = new RepaymentScheduleService(this.http).getRepaymentSchedule(id);
		// add the RepaymentSchedule if not already assigned
		if ( this.loan.schedule.indexOf(repaymentSchedule) == -1 )
		this.loan.schedule.push(repaymentSchedule);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more scheduleIds as a Schedule
	// from a Loan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSchedule( loanId, scheduleIds ): Observable<any> {

		// get the Loan
		this.loadHelper( loanId );


	// split on a comma with no spaces
	var idList 					= scheduleIds.split(',');
	var schedule 	= this.loan.schedule;

	if ( schedule != null && scheduleIds != null ) {

		// iterate over array of schedule ids
		schedule.forEach(function (obj) {
			if ( scheduleIds.indexOf(obj._id) > -1 ) {
				// remove the RepaymentSchedule
				this.loan.schedule.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more collateralIds as a Collateral
	// to a Loan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCollateral( loanId, collateralIds ): Observable<any> {

		// get the Loan
		this.loadHelper( loanId );

	// split on a comma with no spaces
	var idList = collateralIds.split(',')

	// iterate over array of collateral ids
	idList.forEach(function (id) {
		// read the Collateral
		var collateral = new CollateralService(this.http).getCollateral(id);
		// add the Collateral if not already assigned
		if ( this.loan.collateral.indexOf(collateral) == -1 )
		this.loan.collateral.push(collateral);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more collateralIds as a Collateral
	// from a Loan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCollateral( loanId, collateralIds ): Observable<any> {

		// get the Loan
		this.loadHelper( loanId );


	// split on a comma with no spaces
	var idList 					= collateralIds.split(',');
	var collateral 	= this.loan.collateral;

	if ( collateral != null && collateralIds != null ) {

		// iterate over array of collateral ids
		collateral.forEach(function (obj) {
			if ( collateralIds.indexOf(obj._id) > -1 ) {
				// remove the Collateral
				this.loan.collateral.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more transactionsIds as a Transactions
	// to a Loan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTransactions( loanId, transactionsIds ): Observable<any> {

		// get the Loan
		this.loadHelper( loanId );

	// split on a comma with no spaces
	var idList = transactionsIds.split(',')

	// iterate over array of transactions ids
	idList.forEach(function (id) {
		// read the LoanTransaction
		var loanTransaction = new LoanTransactionService(this.http).getLoanTransaction(id);
		// add the LoanTransaction if not already assigned
		if ( this.loan.transactions.indexOf(loanTransaction) == -1 )
		this.loan.transactions.push(loanTransaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more transactionsIds as a Transactions
	// from a Loan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTransactions( loanId, transactionsIds ): Observable<any> {

		// get the Loan
		this.loadHelper( loanId );


	// split on a comma with no spaces
	var idList 					= transactionsIds.split(',');
	var transactions 	= this.loan.transactions;

	if ( transactions != null && transactionsIds != null ) {

		// iterate over array of transactions ids
		transactions.forEach(function (obj) {
			if ( transactionsIds.indexOf(obj._id) > -1 ) {
				// remove the LoanTransaction
				this.loan.transactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Loan
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Loan/update/' + this.loan;

	return  this.http.post(uri_, this.loan );
}

	//********************************************************************
	// loadHelper - internal helper to load a Loan
	//********************************************************************	
	loadHelper( id ) {
		this.getLoan(id)
			.subscribe((res : Loan) => {
				this.loan = res;
			});
	}
}