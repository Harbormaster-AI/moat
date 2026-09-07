import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {LoanTransaction} from '../models/LoanTransaction';
import {LoanService} from '../services/Loan.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LoanTransactionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	loanTransaction : LoanTransaction;

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
	// add a LoanTransaction
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLoanTransaction(transactionId, amount, postingDate, Loan, Type, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/LoanTransaction/create';
		const obj = {
			      		transactionId: transactionId,
      		amount: amount,
      		postingDate: postingDate,
      		Loan: Loan != null && Loan.length > 0 ? Loan : null,
      		Type: Type,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a LoanTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLoanTransaction(transactionId, amount, postingDate, Loan, Type, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/LoanTransaction/update/' + id;
		const obj = {
				      		transactionId: transactionId,
      		amount: amount,
      		postingDate: postingDate,
      		Loan: Loan != null && Loan.length > 0 ? Loan : null,
      		Type: Type,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a LoanTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLoanTransaction(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/LoanTransaction/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a LoanTransaction
	// returns the results untouched as an Observable LoanTransaction
	// LoanTransaction model
	// delegates via URI
	//********************************************************************
	getLoanTransaction(id) : Observable<LoanTransaction> {
		const uri_ = this.apiUrl + '/LoanTransaction/load/' + id;

		return this.http.get<LoanTransaction>(uri_);
	}
	
	//********************************************************************
	// gets all LoanTransaction
	// returns the results untouched as JSON representation of an
	// Observable array of LoanTransaction models
	// delegates via URI
	//********************************************************************
	getLoanTransactions() : Observable<LoanTransaction[]> {
		const uri_ = this.apiUrl + '/LoanTransaction/';

		return this
			.http.get<LoanTransaction[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Loan on a LoanTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLoan( loanTransactionId, _loanId ): Observable<any> {

		// get the LoanTransaction from storage
		this.loadHelper( loanTransactionId );

	// get the Loan from storage
	var tmp 	= new LoanService(this.http).getLoan(_loanId);

	// assign the Loan
	this.loanTransaction.loan = tmp;

	// save the LoanTransaction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Loan on a LoanTransaction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLoan( loanTransactionId ): Observable<any> {

		// get the LoanTransaction from storage
		this.loadHelper( loanTransactionId );

	// assign Loan to null
	this.loanTransaction.loan = null;

	// save the LoanTransaction
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a LoanTransaction
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/LoanTransaction/update/' + this.loanTransaction;

	return  this.http.post(uri_, this.loanTransaction );
}

	//********************************************************************
	// loadHelper - internal helper to load a LoanTransaction
	//********************************************************************	
	loadHelper( id ) {
		this.getLoanTransaction(id)
			.subscribe((res : LoanTransaction) => {
				this.loanTransaction = res;
			});
	}
}