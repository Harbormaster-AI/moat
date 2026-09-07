import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BusinessUnit} from '../models/BusinessUnit';
import {OrganizationService} from '../services/Organization.service';
import {AuditEngagementService} from '../services/AuditEngagement.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BusinessUnitService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	businessUnit : BusinessUnit;

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
	// add a BusinessUnit
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBusinessUnit(name, leader, Organization, Audits) : Observable<any> {
		const uri_ = this.apiUrl + '/BusinessUnit/create';
		const obj = {
			      		name: name,
      		leader: leader,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
			Audits: Audits != null && Audits.length > 0 ? Audits : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BusinessUnit
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBusinessUnit(name, leader, Organization, Audits, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BusinessUnit/update/' + id;
		const obj = {
				      		name: name,
      		leader: leader,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
			Audits: Audits != null && Audits.length > 0 ? Audits : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BusinessUnit
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBusinessUnit(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BusinessUnit/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BusinessUnit
	// returns the results untouched as an Observable BusinessUnit
	// BusinessUnit model
	// delegates via URI
	//********************************************************************
	getBusinessUnit(id) : Observable<BusinessUnit> {
		const uri_ = this.apiUrl + '/BusinessUnit/load/' + id;

		return this.http.get<BusinessUnit>(uri_);
	}
	
	//********************************************************************
	// gets all BusinessUnit
	// returns the results untouched as JSON representation of an
	// Observable array of BusinessUnit models
	// delegates via URI
	//********************************************************************
	getBusinessUnits() : Observable<BusinessUnit[]> {
		const uri_ = this.apiUrl + '/BusinessUnit/';

		return this
			.http.get<BusinessUnit[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a BusinessUnit
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( businessUnitId, _organizationId ): Observable<any> {

		// get the BusinessUnit from storage
		this.loadHelper( businessUnitId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.businessUnit.organization = tmp;

	// save the BusinessUnit
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a BusinessUnit
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( businessUnitId ): Observable<any> {

		// get the BusinessUnit from storage
		this.loadHelper( businessUnitId );

	// assign Organization to null
	this.businessUnit.organization = null;

	// save the BusinessUnit
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more auditsIds as a Audits
	// to a BusinessUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAudits( businessUnitId, auditsIds ): Observable<any> {

		// get the BusinessUnit
		this.loadHelper( businessUnitId );

	// split on a comma with no spaces
	var idList = auditsIds.split(',')

	// iterate over array of audits ids
	idList.forEach(function (id) {
		// read the AuditEngagement
		var auditEngagement = new AuditEngagementService(this.http).getAuditEngagement(id);
		// add the AuditEngagement if not already assigned
		if ( this.businessUnit.audits.indexOf(auditEngagement) == -1 )
		this.businessUnit.audits.push(auditEngagement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more auditsIds as a Audits
	// from a BusinessUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAudits( businessUnitId, auditsIds ): Observable<any> {

		// get the BusinessUnit
		this.loadHelper( businessUnitId );


	// split on a comma with no spaces
	var idList 					= auditsIds.split(',');
	var audits 	= this.businessUnit.audits;

	if ( audits != null && auditsIds != null ) {

		// iterate over array of audits ids
		audits.forEach(function (obj) {
			if ( auditsIds.indexOf(obj._id) > -1 ) {
				// remove the AuditEngagement
				this.businessUnit.audits.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BusinessUnit
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BusinessUnit/update/' + this.businessUnit;

	return  this.http.post(uri_, this.businessUnit );
}

	//********************************************************************
	// loadHelper - internal helper to load a BusinessUnit
	//********************************************************************	
	loadHelper( id ) {
		this.getBusinessUnit(id)
			.subscribe((res : BusinessUnit) => {
				this.businessUnit = res;
			});
	}
}