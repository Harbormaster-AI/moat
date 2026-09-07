import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BenefitPlan} from '../models/BenefitPlan';
import {OrganizationService} from '../services/Organization.service';
import {BenefitEnrollmentService} from '../services/BenefitEnrollment.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BenefitPlanService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	benefitPlan : BenefitPlan;

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
	// add a BenefitPlan
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBenefitPlan(name, providerName, employeeContributionRate, employerContributionRate, eligibilityRules, Organization, Enrollments, BenefitType) : Observable<any> {
		const uri_ = this.apiUrl + '/BenefitPlan/create';
		const obj = {
			      		name: name,
      		providerName: providerName,
      		employeeContributionRate: employeeContributionRate,
      		employerContributionRate: employerContributionRate,
      		eligibilityRules: eligibilityRules,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Enrollments: Enrollments != null && Enrollments.length > 0 ? Enrollments : null,
			BenefitType: BenefitType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BenefitPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBenefitPlan(name, providerName, employeeContributionRate, employerContributionRate, eligibilityRules, Organization, Enrollments, BenefitType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BenefitPlan/update/' + id;
		const obj = {
				      		name: name,
      		providerName: providerName,
      		employeeContributionRate: employeeContributionRate,
      		employerContributionRate: employerContributionRate,
      		eligibilityRules: eligibilityRules,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Enrollments: Enrollments != null && Enrollments.length > 0 ? Enrollments : null,
			BenefitType: BenefitType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BenefitPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBenefitPlan(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BenefitPlan/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BenefitPlan
	// returns the results untouched as an Observable BenefitPlan
	// BenefitPlan model
	// delegates via URI
	//********************************************************************
	getBenefitPlan(id) : Observable<BenefitPlan> {
		const uri_ = this.apiUrl + '/BenefitPlan/load/' + id;

		return this.http.get<BenefitPlan>(uri_);
	}
	
	//********************************************************************
	// gets all BenefitPlan
	// returns the results untouched as JSON representation of an
	// Observable array of BenefitPlan models
	// delegates via URI
	//********************************************************************
	getBenefitPlans() : Observable<BenefitPlan[]> {
		const uri_ = this.apiUrl + '/BenefitPlan/';

		return this
			.http.get<BenefitPlan[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a BenefitPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( benefitPlanId, _organizationId ): Observable<any> {

		// get the BenefitPlan from storage
		this.loadHelper( benefitPlanId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.benefitPlan.organization = tmp;

	// save the BenefitPlan
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a BenefitPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( benefitPlanId ): Observable<any> {

		// get the BenefitPlan from storage
		this.loadHelper( benefitPlanId );

	// assign Organization to null
	this.benefitPlan.organization = null;

	// save the BenefitPlan
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more enrollmentsIds as a Enrollments
	// to a BenefitPlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEnrollments( benefitPlanId, enrollmentsIds ): Observable<any> {

		// get the BenefitPlan
		this.loadHelper( benefitPlanId );

	// split on a comma with no spaces
	var idList = enrollmentsIds.split(',')

	// iterate over array of enrollments ids
	idList.forEach(function (id) {
		// read the BenefitEnrollment
		var benefitEnrollment = new BenefitEnrollmentService(this.http).getBenefitEnrollment(id);
		// add the BenefitEnrollment if not already assigned
		if ( this.benefitPlan.enrollments.indexOf(benefitEnrollment) == -1 )
		this.benefitPlan.enrollments.push(benefitEnrollment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more enrollmentsIds as a Enrollments
	// from a BenefitPlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEnrollments( benefitPlanId, enrollmentsIds ): Observable<any> {

		// get the BenefitPlan
		this.loadHelper( benefitPlanId );


	// split on a comma with no spaces
	var idList 					= enrollmentsIds.split(',');
	var enrollments 	= this.benefitPlan.enrollments;

	if ( enrollments != null && enrollmentsIds != null ) {

		// iterate over array of enrollments ids
		enrollments.forEach(function (obj) {
			if ( enrollmentsIds.indexOf(obj._id) > -1 ) {
				// remove the BenefitEnrollment
				this.benefitPlan.enrollments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BenefitPlan
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BenefitPlan/update/' + this.benefitPlan;

	return  this.http.post(uri_, this.benefitPlan );
}

	//********************************************************************
	// loadHelper - internal helper to load a BenefitPlan
	//********************************************************************	
	loadHelper( id ) {
		this.getBenefitPlan(id)
			.subscribe((res : BenefitPlan) => {
				this.benefitPlan = res;
			});
	}
}