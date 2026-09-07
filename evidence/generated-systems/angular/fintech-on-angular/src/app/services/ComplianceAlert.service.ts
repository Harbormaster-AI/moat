import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ComplianceAlert} from '../models/ComplianceAlert';
import {ScreeningService} from '../services/Screening.service';
import {TransactionService} from '../services/Transaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ComplianceAlertService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	complianceAlert : ComplianceAlert;

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
	// add a ComplianceAlert
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addComplianceAlert(alertCode, raisedAt, notes, Screening, Transaction, Severity, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/ComplianceAlert/create';
		const obj = {
			      		alertCode: alertCode,
      		raisedAt: raisedAt,
      		notes: notes,
      		Screening: Screening != null && Screening.length > 0 ? Screening : null,
      		Transaction: Transaction != null && Transaction.length > 0 ? Transaction : null,
      		Severity: Severity,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ComplianceAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateComplianceAlert(alertCode, raisedAt, notes, Screening, Transaction, Severity, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ComplianceAlert/update/' + id;
		const obj = {
				      		alertCode: alertCode,
      		raisedAt: raisedAt,
      		notes: notes,
      		Screening: Screening != null && Screening.length > 0 ? Screening : null,
      		Transaction: Transaction != null && Transaction.length > 0 ? Transaction : null,
      		Severity: Severity,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ComplianceAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteComplianceAlert(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ComplianceAlert/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ComplianceAlert
	// returns the results untouched as an Observable ComplianceAlert
	// ComplianceAlert model
	// delegates via URI
	//********************************************************************
	getComplianceAlert(id) : Observable<ComplianceAlert> {
		const uri_ = this.apiUrl + '/ComplianceAlert/load/' + id;

		return this.http.get<ComplianceAlert>(uri_);
	}
	
	//********************************************************************
	// gets all ComplianceAlert
	// returns the results untouched as JSON representation of an
	// Observable array of ComplianceAlert models
	// delegates via URI
	//********************************************************************
	getComplianceAlerts() : Observable<ComplianceAlert[]> {
		const uri_ = this.apiUrl + '/ComplianceAlert/';

		return this
			.http.get<ComplianceAlert[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Screening on a ComplianceAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignScreening( complianceAlertId, _screeningId ): Observable<any> {

		// get the ComplianceAlert from storage
		this.loadHelper( complianceAlertId );

	// get the Screening from storage
	var tmp 	= new ScreeningService(this.http).getScreening(_screeningId);

	// assign the Screening
	this.complianceAlert.screening = tmp;

	// save the ComplianceAlert
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Screening on a ComplianceAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignScreening( complianceAlertId ): Observable<any> {

		// get the ComplianceAlert from storage
		this.loadHelper( complianceAlertId );

	// assign Screening to null
	this.complianceAlert.screening = null;

	// save the ComplianceAlert
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Transaction on a ComplianceAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTransaction( complianceAlertId, _transactionId ): Observable<any> {

		// get the ComplianceAlert from storage
		this.loadHelper( complianceAlertId );

	// get the Transaction from storage
	var tmp 	= new TransactionService(this.http).getTransaction(_transactionId);

	// assign the Transaction
	this.complianceAlert.transaction = tmp;

	// save the ComplianceAlert
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Transaction on a ComplianceAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTransaction( complianceAlertId ): Observable<any> {

		// get the ComplianceAlert from storage
		this.loadHelper( complianceAlertId );

	// assign Transaction to null
	this.complianceAlert.transaction = null;

	// save the ComplianceAlert
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ComplianceAlert
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ComplianceAlert/update/' + this.complianceAlert;

	return  this.http.post(uri_, this.complianceAlert );
}

	//********************************************************************
	// loadHelper - internal helper to load a ComplianceAlert
	//********************************************************************	
	loadHelper( id ) {
		this.getComplianceAlert(id)
			.subscribe((res : ComplianceAlert) => {
				this.complianceAlert = res;
			});
	}
}