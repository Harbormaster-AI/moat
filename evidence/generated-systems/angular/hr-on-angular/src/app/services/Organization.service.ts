import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Organization} from '../models/Organization';
import {DepartmentService} from '../services/Department.service';
import {LocationService} from '../services/Location.service';
import {JobFamilyService} from '../services/JobFamily.service';
import {BenefitPlanService} from '../services/BenefitPlan.service';
import {CostCenterService} from '../services/CostCenter.service';
import {PayrollCalendarService} from '../services/PayrollCalendar.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OrganizationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	organization : Organization;

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
	// add a Organization
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOrganization(name, legalName, registrationCountry, website, Departments, Locations, JobFamilies, BenefitPlans, CostCenters, PayrollCalendars) : Observable<any> {
		const uri_ = this.apiUrl + '/Organization/create';
		const obj = {
			      		name: name,
      		legalName: legalName,
      		registrationCountry: registrationCountry,
      		website: website,
      		Departments: Departments != null && Departments.length > 0 ? Departments : null,
      		Locations: Locations != null && Locations.length > 0 ? Locations : null,
      		JobFamilies: JobFamilies != null && JobFamilies.length > 0 ? JobFamilies : null,
      		BenefitPlans: BenefitPlans != null && BenefitPlans.length > 0 ? BenefitPlans : null,
      		CostCenters: CostCenters != null && CostCenters.length > 0 ? CostCenters : null,
			PayrollCalendars: PayrollCalendars != null && PayrollCalendars.length > 0 ? PayrollCalendars : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Organization
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOrganization(name, legalName, registrationCountry, website, Departments, Locations, JobFamilies, BenefitPlans, CostCenters, PayrollCalendars, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Organization/update/' + id;
		const obj = {
				      		name: name,
      		legalName: legalName,
      		registrationCountry: registrationCountry,
      		website: website,
      		Departments: Departments != null && Departments.length > 0 ? Departments : null,
      		Locations: Locations != null && Locations.length > 0 ? Locations : null,
      		JobFamilies: JobFamilies != null && JobFamilies.length > 0 ? JobFamilies : null,
      		BenefitPlans: BenefitPlans != null && BenefitPlans.length > 0 ? BenefitPlans : null,
      		CostCenters: CostCenters != null && CostCenters.length > 0 ? CostCenters : null,
			PayrollCalendars: PayrollCalendars != null && PayrollCalendars.length > 0 ? PayrollCalendars : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Organization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOrganization(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Organization/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Organization
	// returns the results untouched as an Observable Organization
	// Organization model
	// delegates via URI
	//********************************************************************
	getOrganization(id) : Observable<Organization> {
		const uri_ = this.apiUrl + '/Organization/load/' + id;

		return this.http.get<Organization>(uri_);
	}
	
	//********************************************************************
	// gets all Organization
	// returns the results untouched as JSON representation of an
	// Observable array of Organization models
	// delegates via URI
	//********************************************************************
	getOrganizations() : Observable<Organization[]> {
		const uri_ = this.apiUrl + '/Organization/';

		return this
			.http.get<Organization[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more departmentsIds as a Departments
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDepartments( organizationId, departmentsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = departmentsIds.split(',')

	// iterate over array of departments ids
	idList.forEach(function (id) {
		// read the Department
		var department = new DepartmentService(this.http).getDepartment(id);
		// add the Department if not already assigned
		if ( this.organization.departments.indexOf(department) == -1 )
		this.organization.departments.push(department);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more departmentsIds as a Departments
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDepartments( organizationId, departmentsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= departmentsIds.split(',');
	var departments 	= this.organization.departments;

	if ( departments != null && departmentsIds != null ) {

		// iterate over array of departments ids
		departments.forEach(function (obj) {
			if ( departmentsIds.indexOf(obj._id) > -1 ) {
				// remove the Department
				this.organization.departments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more locationsIds as a Locations
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLocations( organizationId, locationsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = locationsIds.split(',')

	// iterate over array of locations ids
	idList.forEach(function (id) {
		// read the Location
		var location = new LocationService(this.http).getLocation(id);
		// add the Location if not already assigned
		if ( this.organization.locations.indexOf(location) == -1 )
		this.organization.locations.push(location);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more locationsIds as a Locations
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLocations( organizationId, locationsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= locationsIds.split(',');
	var locations 	= this.organization.locations;

	if ( locations != null && locationsIds != null ) {

		// iterate over array of locations ids
		locations.forEach(function (obj) {
			if ( locationsIds.indexOf(obj._id) > -1 ) {
				// remove the Location
				this.organization.locations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more jobFamiliesIds as a JobFamilies
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addJobFamilies( organizationId, jobFamiliesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = jobFamiliesIds.split(',')

	// iterate over array of jobFamilies ids
	idList.forEach(function (id) {
		// read the JobFamily
		var jobFamily = new JobFamilyService(this.http).getJobFamily(id);
		// add the JobFamily if not already assigned
		if ( this.organization.jobFamilies.indexOf(jobFamily) == -1 )
		this.organization.jobFamilies.push(jobFamily);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more jobFamiliesIds as a JobFamilies
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeJobFamilies( organizationId, jobFamiliesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= jobFamiliesIds.split(',');
	var jobFamilies 	= this.organization.jobFamilies;

	if ( jobFamilies != null && jobFamiliesIds != null ) {

		// iterate over array of jobFamilies ids
		jobFamilies.forEach(function (obj) {
			if ( jobFamiliesIds.indexOf(obj._id) > -1 ) {
				// remove the JobFamily
				this.organization.jobFamilies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more benefitPlansIds as a BenefitPlans
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBenefitPlans( organizationId, benefitPlansIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = benefitPlansIds.split(',')

	// iterate over array of benefitPlans ids
	idList.forEach(function (id) {
		// read the BenefitPlan
		var benefitPlan = new BenefitPlanService(this.http).getBenefitPlan(id);
		// add the BenefitPlan if not already assigned
		if ( this.organization.benefitPlans.indexOf(benefitPlan) == -1 )
		this.organization.benefitPlans.push(benefitPlan);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more benefitPlansIds as a BenefitPlans
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBenefitPlans( organizationId, benefitPlansIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= benefitPlansIds.split(',');
	var benefitPlans 	= this.organization.benefitPlans;

	if ( benefitPlans != null && benefitPlansIds != null ) {

		// iterate over array of benefitPlans ids
		benefitPlans.forEach(function (obj) {
			if ( benefitPlansIds.indexOf(obj._id) > -1 ) {
				// remove the BenefitPlan
				this.organization.benefitPlans.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more costCentersIds as a CostCenters
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCostCenters( organizationId, costCentersIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = costCentersIds.split(',')

	// iterate over array of costCenters ids
	idList.forEach(function (id) {
		// read the CostCenter
		var costCenter = new CostCenterService(this.http).getCostCenter(id);
		// add the CostCenter if not already assigned
		if ( this.organization.costCenters.indexOf(costCenter) == -1 )
		this.organization.costCenters.push(costCenter);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more costCentersIds as a CostCenters
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCostCenters( organizationId, costCentersIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= costCentersIds.split(',');
	var costCenters 	= this.organization.costCenters;

	if ( costCenters != null && costCentersIds != null ) {

		// iterate over array of costCenters ids
		costCenters.forEach(function (obj) {
			if ( costCentersIds.indexOf(obj._id) > -1 ) {
				// remove the CostCenter
				this.organization.costCenters.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more payrollCalendarsIds as a PayrollCalendars
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPayrollCalendars( organizationId, payrollCalendarsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = payrollCalendarsIds.split(',')

	// iterate over array of payrollCalendars ids
	idList.forEach(function (id) {
		// read the PayrollCalendar
		var payrollCalendar = new PayrollCalendarService(this.http).getPayrollCalendar(id);
		// add the PayrollCalendar if not already assigned
		if ( this.organization.payrollCalendars.indexOf(payrollCalendar) == -1 )
		this.organization.payrollCalendars.push(payrollCalendar);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more payrollCalendarsIds as a PayrollCalendars
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePayrollCalendars( organizationId, payrollCalendarsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= payrollCalendarsIds.split(',');
	var payrollCalendars 	= this.organization.payrollCalendars;

	if ( payrollCalendars != null && payrollCalendarsIds != null ) {

		// iterate over array of payrollCalendars ids
		payrollCalendars.forEach(function (obj) {
			if ( payrollCalendarsIds.indexOf(obj._id) > -1 ) {
				// remove the PayrollCalendar
				this.organization.payrollCalendars.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Organization
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Organization/update/' + this.organization;

	return  this.http.post(uri_, this.organization );
}

	//********************************************************************
	// loadHelper - internal helper to load a Organization
	//********************************************************************	
	loadHelper( id ) {
		this.getOrganization(id)
			.subscribe((res : Organization) => {
				this.organization = res;
			});
	}
}