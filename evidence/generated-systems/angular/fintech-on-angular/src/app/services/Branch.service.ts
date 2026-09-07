import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Branch} from '../models/Branch';
import {FinancialInstitutionService} from '../services/FinancialInstitution.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BranchService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	branch : Branch;

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
	// add a Branch
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBranch(name, branchCode, address, Institution) : Observable<any> {
		const uri_ = this.apiUrl + '/Branch/create';
		const obj = {
			      		name: name,
      		branchCode: branchCode,
      		address: address,
			Institution: Institution != null && Institution.length > 0 ? Institution : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Branch
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBranch(name, branchCode, address, Institution, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Branch/update/' + id;
		const obj = {
				      		name: name,
      		branchCode: branchCode,
      		address: address,
			Institution: Institution != null && Institution.length > 0 ? Institution : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Branch
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBranch(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Branch/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Branch
	// returns the results untouched as an Observable Branch
	// Branch model
	// delegates via URI
	//********************************************************************
	getBranch(id) : Observable<Branch> {
		const uri_ = this.apiUrl + '/Branch/load/' + id;

		return this.http.get<Branch>(uri_);
	}
	
	//********************************************************************
	// gets all Branch
	// returns the results untouched as JSON representation of an
	// Observable array of Branch models
	// delegates via URI
	//********************************************************************
	getBranchs() : Observable<Branch[]> {
		const uri_ = this.apiUrl + '/Branch/';

		return this
			.http.get<Branch[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Institution on a Branch
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInstitution( branchId, _institutionId ): Observable<any> {

		// get the Branch from storage
		this.loadHelper( branchId );

	// get the FinancialInstitution from storage
	var tmp 	= new FinancialInstitutionService(this.http).getFinancialInstitution(_institutionId);

	// assign the Institution
	this.branch.institution = tmp;

	// save the Branch
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Institution on a Branch
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInstitution( branchId ): Observable<any> {

		// get the Branch from storage
		this.loadHelper( branchId );

	// assign Institution to null
	this.branch.institution = null;

	// save the Branch
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Branch
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Branch/update/' + this.branch;

	return  this.http.post(uri_, this.branch );
}

	//********************************************************************
	// loadHelper - internal helper to load a Branch
	//********************************************************************	
	loadHelper( id ) {
		this.getBranch(id)
			.subscribe((res : Branch) => {
				this.branch = res;
			});
	}
}