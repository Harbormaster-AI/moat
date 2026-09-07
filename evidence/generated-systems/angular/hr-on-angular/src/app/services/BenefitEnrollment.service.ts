import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BenefitEnrollment} from '../models/BenefitEnrollment';
import {BenefitPlanService} from '../services/BenefitPlan.service';
import {EmployeeService} from '../services/Employee.service';
import {DependentService} from '../services/Dependent.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BenefitEnrollmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	benefitEnrollment : BenefitEnrollment;

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
	// add a BenefitEnrollment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBenefitEnrollment(enrollmentId, effectiveFrom, effectiveTo, BenefitPlan, Employee, Dependents, Status, CoverageLevel) : Observable<any> {
		const uri_ = this.apiUrl + '/BenefitEnrollment/create';
		const obj = {
			      		enrollmentId: enrollmentId,
      		effectiveFrom: effectiveFrom,
      		effectiveTo: effectiveTo,
      		BenefitPlan: BenefitPlan != null && BenefitPlan.length > 0 ? BenefitPlan : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Dependents: Dependents != null && Dependents.length > 0 ? Dependents : null,
      		Status: Status,
			CoverageLevel: CoverageLevel
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BenefitEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBenefitEnrollment(enrollmentId, effectiveFrom, effectiveTo, BenefitPlan, Employee, Dependents, Status, CoverageLevel, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BenefitEnrollment/update/' + id;
		const obj = {
				      		enrollmentId: enrollmentId,
      		effectiveFrom: effectiveFrom,
      		effectiveTo: effectiveTo,
      		BenefitPlan: BenefitPlan != null && BenefitPlan.length > 0 ? BenefitPlan : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Dependents: Dependents != null && Dependents.length > 0 ? Dependents : null,
      		Status: Status,
			CoverageLevel: CoverageLevel
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BenefitEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBenefitEnrollment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BenefitEnrollment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BenefitEnrollment
	// returns the results untouched as an Observable BenefitEnrollment
	// BenefitEnrollment model
	// delegates via URI
	//********************************************************************
	getBenefitEnrollment(id) : Observable<BenefitEnrollment> {
		const uri_ = this.apiUrl + '/BenefitEnrollment/load/' + id;

		return this.http.get<BenefitEnrollment>(uri_);
	}
	
	//********************************************************************
	// gets all BenefitEnrollment
	// returns the results untouched as JSON representation of an
	// Observable array of BenefitEnrollment models
	// delegates via URI
	//********************************************************************
	getBenefitEnrollments() : Observable<BenefitEnrollment[]> {
		const uri_ = this.apiUrl + '/BenefitEnrollment/';

		return this
			.http.get<BenefitEnrollment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a BenefitPlan on a BenefitEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBenefitPlan( benefitEnrollmentId, _benefitPlanId ): Observable<any> {

		// get the BenefitEnrollment from storage
		this.loadHelper( benefitEnrollmentId );

	// get the BenefitPlan from storage
	var tmp 	= new BenefitPlanService(this.http).getBenefitPlan(_benefitPlanId);

	// assign the BenefitPlan
	this.benefitEnrollment.benefitPlan = tmp;

	// save the BenefitEnrollment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a BenefitPlan on a BenefitEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBenefitPlan( benefitEnrollmentId ): Observable<any> {

		// get the BenefitEnrollment from storage
		this.loadHelper( benefitEnrollmentId );

	// assign BenefitPlan to null
	this.benefitEnrollment.benefitPlan = null;

	// save the BenefitEnrollment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Employee on a BenefitEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( benefitEnrollmentId, _employeeId ): Observable<any> {

		// get the BenefitEnrollment from storage
		this.loadHelper( benefitEnrollmentId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.benefitEnrollment.employee = tmp;

	// save the BenefitEnrollment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a BenefitEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( benefitEnrollmentId ): Observable<any> {

		// get the BenefitEnrollment from storage
		this.loadHelper( benefitEnrollmentId );

	// assign Employee to null
	this.benefitEnrollment.employee = null;

	// save the BenefitEnrollment
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more dependentsIds as a Dependents
	// to a BenefitEnrollment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDependents( benefitEnrollmentId, dependentsIds ): Observable<any> {

		// get the BenefitEnrollment
		this.loadHelper( benefitEnrollmentId );

	// split on a comma with no spaces
	var idList = dependentsIds.split(',')

	// iterate over array of dependents ids
	idList.forEach(function (id) {
		// read the Dependent
		var dependent = new DependentService(this.http).getDependent(id);
		// add the Dependent if not already assigned
		if ( this.benefitEnrollment.dependents.indexOf(dependent) == -1 )
		this.benefitEnrollment.dependents.push(dependent);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dependentsIds as a Dependents
	// from a BenefitEnrollment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDependents( benefitEnrollmentId, dependentsIds ): Observable<any> {

		// get the BenefitEnrollment
		this.loadHelper( benefitEnrollmentId );


	// split on a comma with no spaces
	var idList 					= dependentsIds.split(',');
	var dependents 	= this.benefitEnrollment.dependents;

	if ( dependents != null && dependentsIds != null ) {

		// iterate over array of dependents ids
		dependents.forEach(function (obj) {
			if ( dependentsIds.indexOf(obj._id) > -1 ) {
				// remove the Dependent
				this.benefitEnrollment.dependents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BenefitEnrollment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BenefitEnrollment/update/' + this.benefitEnrollment;

	return  this.http.post(uri_, this.benefitEnrollment );
}

	//********************************************************************
	// loadHelper - internal helper to load a BenefitEnrollment
	//********************************************************************	
	loadHelper( id ) {
		this.getBenefitEnrollment(id)
			.subscribe((res : BenefitEnrollment) => {
				this.benefitEnrollment = res;
			});
	}
}