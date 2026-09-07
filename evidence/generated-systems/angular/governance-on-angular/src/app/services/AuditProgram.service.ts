import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AuditProgram} from '../models/AuditProgram';
import {OrganizationService} from '../services/Organization.service';
import {AuditEngagementService} from '../services/AuditEngagement.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AuditProgramService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	auditProgram : AuditProgram;

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
	// add a AuditProgram
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAuditProgram(name, scope, Organization, Engagements, Cycle, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/AuditProgram/create';
		const obj = {
			      		name: name,
      		scope: scope,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Engagements: Engagements != null && Engagements.length > 0 ? Engagements : null,
      		Cycle: Cycle,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AuditProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAuditProgram(name, scope, Organization, Engagements, Cycle, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AuditProgram/update/' + id;
		const obj = {
				      		name: name,
      		scope: scope,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Engagements: Engagements != null && Engagements.length > 0 ? Engagements : null,
      		Cycle: Cycle,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AuditProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAuditProgram(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AuditProgram/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AuditProgram
	// returns the results untouched as an Observable AuditProgram
	// AuditProgram model
	// delegates via URI
	//********************************************************************
	getAuditProgram(id) : Observable<AuditProgram> {
		const uri_ = this.apiUrl + '/AuditProgram/load/' + id;

		return this.http.get<AuditProgram>(uri_);
	}
	
	//********************************************************************
	// gets all AuditProgram
	// returns the results untouched as JSON representation of an
	// Observable array of AuditProgram models
	// delegates via URI
	//********************************************************************
	getAuditPrograms() : Observable<AuditProgram[]> {
		const uri_ = this.apiUrl + '/AuditProgram/';

		return this
			.http.get<AuditProgram[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a AuditProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( auditProgramId, _organizationId ): Observable<any> {

		// get the AuditProgram from storage
		this.loadHelper( auditProgramId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.auditProgram.organization = tmp;

	// save the AuditProgram
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a AuditProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( auditProgramId ): Observable<any> {

		// get the AuditProgram from storage
		this.loadHelper( auditProgramId );

	// assign Organization to null
	this.auditProgram.organization = null;

	// save the AuditProgram
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more engagementsIds as a Engagements
	// to a AuditProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEngagements( auditProgramId, engagementsIds ): Observable<any> {

		// get the AuditProgram
		this.loadHelper( auditProgramId );

	// split on a comma with no spaces
	var idList = engagementsIds.split(',')

	// iterate over array of engagements ids
	idList.forEach(function (id) {
		// read the AuditEngagement
		var auditEngagement = new AuditEngagementService(this.http).getAuditEngagement(id);
		// add the AuditEngagement if not already assigned
		if ( this.auditProgram.engagements.indexOf(auditEngagement) == -1 )
		this.auditProgram.engagements.push(auditEngagement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more engagementsIds as a Engagements
	// from a AuditProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEngagements( auditProgramId, engagementsIds ): Observable<any> {

		// get the AuditProgram
		this.loadHelper( auditProgramId );


	// split on a comma with no spaces
	var idList 					= engagementsIds.split(',');
	var engagements 	= this.auditProgram.engagements;

	if ( engagements != null && engagementsIds != null ) {

		// iterate over array of engagements ids
		engagements.forEach(function (obj) {
			if ( engagementsIds.indexOf(obj._id) > -1 ) {
				// remove the AuditEngagement
				this.auditProgram.engagements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AuditProgram
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AuditProgram/update/' + this.auditProgram;

	return  this.http.post(uri_, this.auditProgram );
}

	//********************************************************************
	// loadHelper - internal helper to load a AuditProgram
	//********************************************************************	
	loadHelper( id ) {
		this.getAuditProgram(id)
			.subscribe((res : AuditProgram) => {
				this.auditProgram = res;
			});
	}
}