import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Dependent} from '../models/Dependent';
import {BenefitEnrollmentService} from '../services/BenefitEnrollment.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DependentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dependent : Dependent;

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
	// add a Dependent
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDependent(firstName, lastName, birthDate, BenefitEnrollment, Employee, Relationship) : Observable<any> {
		const uri_ = this.apiUrl + '/Dependent/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		birthDate: birthDate,
      		BenefitEnrollment: BenefitEnrollment != null && BenefitEnrollment.length > 0 ? BenefitEnrollment : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			Relationship: Relationship
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Dependent
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDependent(firstName, lastName, birthDate, BenefitEnrollment, Employee, Relationship, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Dependent/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		birthDate: birthDate,
      		BenefitEnrollment: BenefitEnrollment != null && BenefitEnrollment.length > 0 ? BenefitEnrollment : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			Relationship: Relationship
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Dependent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDependent(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Dependent/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Dependent
	// returns the results untouched as an Observable Dependent
	// Dependent model
	// delegates via URI
	//********************************************************************
	getDependent(id) : Observable<Dependent> {
		const uri_ = this.apiUrl + '/Dependent/load/' + id;

		return this.http.get<Dependent>(uri_);
	}
	
	//********************************************************************
	// gets all Dependent
	// returns the results untouched as JSON representation of an
	// Observable array of Dependent models
	// delegates via URI
	//********************************************************************
	getDependents() : Observable<Dependent[]> {
		const uri_ = this.apiUrl + '/Dependent/';

		return this
			.http.get<Dependent[]>(uri_);
	}
	
			//********************************************************************
	// assigns a BenefitEnrollment on a Dependent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBenefitEnrollment( dependentId, _benefitEnrollmentId ): Observable<any> {

		// get the Dependent from storage
		this.loadHelper( dependentId );

	// get the BenefitEnrollment from storage
	var tmp 	= new BenefitEnrollmentService(this.http).getBenefitEnrollment(_benefitEnrollmentId);

	// assign the BenefitEnrollment
	this.dependent.benefitEnrollment = tmp;

	// save the Dependent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a BenefitEnrollment on a Dependent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBenefitEnrollment( dependentId ): Observable<any> {

		// get the Dependent from storage
		this.loadHelper( dependentId );

	// assign BenefitEnrollment to null
	this.dependent.benefitEnrollment = null;

	// save the Dependent
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Employee on a Dependent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( dependentId, _employeeId ): Observable<any> {

		// get the Dependent from storage
		this.loadHelper( dependentId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.dependent.employee = tmp;

	// save the Dependent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a Dependent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( dependentId ): Observable<any> {

		// get the Dependent from storage
		this.loadHelper( dependentId );

	// assign Employee to null
	this.dependent.employee = null;

	// save the Dependent
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Dependent
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Dependent/update/' + this.dependent;

	return  this.http.post(uri_, this.dependent );
}

	//********************************************************************
	// loadHelper - internal helper to load a Dependent
	//********************************************************************	
	loadHelper( id ) {
		this.getDependent(id)
			.subscribe((res : Dependent) => {
				this.dependent = res;
			});
	}
}