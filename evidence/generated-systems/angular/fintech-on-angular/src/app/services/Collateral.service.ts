import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Collateral} from '../models/Collateral';
import {LoanService} from '../services/Loan.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CollateralService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	collateral : Collateral;

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
	// add a Collateral
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCollateral(description, value, Loan, CollateralType) : Observable<any> {
		const uri_ = this.apiUrl + '/Collateral/create';
		const obj = {
			      		description: description,
      		value: value,
      		Loan: Loan != null && Loan.length > 0 ? Loan : null,
			CollateralType: CollateralType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Collateral
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCollateral(description, value, Loan, CollateralType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Collateral/update/' + id;
		const obj = {
				      		description: description,
      		value: value,
      		Loan: Loan != null && Loan.length > 0 ? Loan : null,
			CollateralType: CollateralType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Collateral
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCollateral(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Collateral/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Collateral
	// returns the results untouched as an Observable Collateral
	// Collateral model
	// delegates via URI
	//********************************************************************
	getCollateral(id) : Observable<Collateral> {
		const uri_ = this.apiUrl + '/Collateral/load/' + id;

		return this.http.get<Collateral>(uri_);
	}
	
	//********************************************************************
	// gets all Collateral
	// returns the results untouched as JSON representation of an
	// Observable array of Collateral models
	// delegates via URI
	//********************************************************************
	getCollaterals() : Observable<Collateral[]> {
		const uri_ = this.apiUrl + '/Collateral/';

		return this
			.http.get<Collateral[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Loan on a Collateral
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLoan( collateralId, _loanId ): Observable<any> {

		// get the Collateral from storage
		this.loadHelper( collateralId );

	// get the Loan from storage
	var tmp 	= new LoanService(this.http).getLoan(_loanId);

	// assign the Loan
	this.collateral.loan = tmp;

	// save the Collateral
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Loan on a Collateral
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLoan( collateralId ): Observable<any> {

		// get the Collateral from storage
		this.loadHelper( collateralId );

	// assign Loan to null
	this.collateral.loan = null;

	// save the Collateral
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Collateral
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Collateral/update/' + this.collateral;

	return  this.http.post(uri_, this.collateral );
}

	//********************************************************************
	// loadHelper - internal helper to load a Collateral
	//********************************************************************	
	loadHelper( id ) {
		this.getCollateral(id)
			.subscribe((res : Collateral) => {
				this.collateral = res;
			});
	}
}