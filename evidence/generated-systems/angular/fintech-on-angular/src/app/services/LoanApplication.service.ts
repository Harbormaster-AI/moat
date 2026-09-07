import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {LoanApplication} from '../models/LoanApplication';
import {CustomerService} from '../services/Customer.service';
import {RiskAssessmentService} from '../services/RiskAssessment.service';
import {LoanService} from '../services/Loan.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LoanApplicationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	loanApplication : LoanApplication;

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
	// add a LoanApplication
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLoanApplication(applicationNumber, amountRequested, termMonths, submittedAt, Customer, RiskAssessment, Loan, Product, Purpose, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/LoanApplication/create';
		const obj = {
			      		applicationNumber: applicationNumber,
      		amountRequested: amountRequested,
      		termMonths: termMonths,
      		submittedAt: submittedAt,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		RiskAssessment: RiskAssessment != null && RiskAssessment.length > 0 ? RiskAssessment : null,
      		Loan: Loan != null && Loan.length > 0 ? Loan : null,
      		Product: Product,
      		Purpose: Purpose,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a LoanApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLoanApplication(applicationNumber, amountRequested, termMonths, submittedAt, Customer, RiskAssessment, Loan, Product, Purpose, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/LoanApplication/update/' + id;
		const obj = {
				      		applicationNumber: applicationNumber,
      		amountRequested: amountRequested,
      		termMonths: termMonths,
      		submittedAt: submittedAt,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		RiskAssessment: RiskAssessment != null && RiskAssessment.length > 0 ? RiskAssessment : null,
      		Loan: Loan != null && Loan.length > 0 ? Loan : null,
      		Product: Product,
      		Purpose: Purpose,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a LoanApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLoanApplication(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/LoanApplication/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a LoanApplication
	// returns the results untouched as an Observable LoanApplication
	// LoanApplication model
	// delegates via URI
	//********************************************************************
	getLoanApplication(id) : Observable<LoanApplication> {
		const uri_ = this.apiUrl + '/LoanApplication/load/' + id;

		return this.http.get<LoanApplication>(uri_);
	}
	
	//********************************************************************
	// gets all LoanApplication
	// returns the results untouched as JSON representation of an
	// Observable array of LoanApplication models
	// delegates via URI
	//********************************************************************
	getLoanApplications() : Observable<LoanApplication[]> {
		const uri_ = this.apiUrl + '/LoanApplication/';

		return this
			.http.get<LoanApplication[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a LoanApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( loanApplicationId, _customerId ): Observable<any> {

		// get the LoanApplication from storage
		this.loadHelper( loanApplicationId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.loanApplication.customer = tmp;

	// save the LoanApplication
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a LoanApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( loanApplicationId ): Observable<any> {

		// get the LoanApplication from storage
		this.loadHelper( loanApplicationId );

	// assign Customer to null
	this.loanApplication.customer = null;

	// save the LoanApplication
	return this.saveHelper();
}

		//********************************************************************
	// assigns a RiskAssessment on a LoanApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRiskAssessment( loanApplicationId, _riskAssessmentId ): Observable<any> {

		// get the LoanApplication from storage
		this.loadHelper( loanApplicationId );

	// get the RiskAssessment from storage
	var tmp 	= new RiskAssessmentService(this.http).getRiskAssessment(_riskAssessmentId);

	// assign the RiskAssessment
	this.loanApplication.riskAssessment = tmp;

	// save the LoanApplication
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a RiskAssessment on a LoanApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRiskAssessment( loanApplicationId ): Observable<any> {

		// get the LoanApplication from storage
		this.loadHelper( loanApplicationId );

	// assign RiskAssessment to null
	this.loanApplication.riskAssessment = null;

	// save the LoanApplication
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Loan on a LoanApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLoan( loanApplicationId, _loanId ): Observable<any> {

		// get the LoanApplication from storage
		this.loadHelper( loanApplicationId );

	// get the Loan from storage
	var tmp 	= new LoanService(this.http).getLoan(_loanId);

	// assign the Loan
	this.loanApplication.loan = tmp;

	// save the LoanApplication
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Loan on a LoanApplication
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLoan( loanApplicationId ): Observable<any> {

		// get the LoanApplication from storage
		this.loadHelper( loanApplicationId );

	// assign Loan to null
	this.loanApplication.loan = null;

	// save the LoanApplication
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a LoanApplication
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/LoanApplication/update/' + this.loanApplication;

	return  this.http.post(uri_, this.loanApplication );
}

	//********************************************************************
	// loadHelper - internal helper to load a LoanApplication
	//********************************************************************	
	loadHelper( id ) {
		this.getLoanApplication(id)
			.subscribe((res : LoanApplication) => {
				this.loanApplication = res;
			});
	}
}