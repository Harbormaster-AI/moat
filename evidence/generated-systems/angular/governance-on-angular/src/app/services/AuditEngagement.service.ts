import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AuditEngagement} from '../models/AuditEngagement';
import {AuditProgramService} from '../services/AuditProgram.service';
import {BusinessUnitService} from '../services/BusinessUnit.service';
import {ControlTest_Service} from '../services/ControlTest_.service';
import {AuditWorkpaperService} from '../services/AuditWorkpaper.service';
import {AuditFindingService} from '../services/AuditFinding.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AuditEngagementService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	auditEngagement : AuditEngagement;

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
	// add a AuditEngagement
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAuditEngagement(title, startDate, endDate, AuditProgram, BusinessUnits, ControlTests, Workpapers, Findings, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/AuditEngagement/create';
		const obj = {
			      		title: title,
      		startDate: startDate,
      		endDate: endDate,
      		AuditProgram: AuditProgram != null && AuditProgram.length > 0 ? AuditProgram : null,
      		BusinessUnits: BusinessUnits != null && BusinessUnits.length > 0 ? BusinessUnits : null,
      		ControlTests: ControlTests != null && ControlTests.length > 0 ? ControlTests : null,
      		Workpapers: Workpapers != null && Workpapers.length > 0 ? Workpapers : null,
      		Findings: Findings != null && Findings.length > 0 ? Findings : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AuditEngagement
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAuditEngagement(title, startDate, endDate, AuditProgram, BusinessUnits, ControlTests, Workpapers, Findings, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AuditEngagement/update/' + id;
		const obj = {
				      		title: title,
      		startDate: startDate,
      		endDate: endDate,
      		AuditProgram: AuditProgram != null && AuditProgram.length > 0 ? AuditProgram : null,
      		BusinessUnits: BusinessUnits != null && BusinessUnits.length > 0 ? BusinessUnits : null,
      		ControlTests: ControlTests != null && ControlTests.length > 0 ? ControlTests : null,
      		Workpapers: Workpapers != null && Workpapers.length > 0 ? Workpapers : null,
      		Findings: Findings != null && Findings.length > 0 ? Findings : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AuditEngagement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAuditEngagement(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AuditEngagement/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AuditEngagement
	// returns the results untouched as an Observable AuditEngagement
	// AuditEngagement model
	// delegates via URI
	//********************************************************************
	getAuditEngagement(id) : Observable<AuditEngagement> {
		const uri_ = this.apiUrl + '/AuditEngagement/load/' + id;

		return this.http.get<AuditEngagement>(uri_);
	}
	
	//********************************************************************
	// gets all AuditEngagement
	// returns the results untouched as JSON representation of an
	// Observable array of AuditEngagement models
	// delegates via URI
	//********************************************************************
	getAuditEngagements() : Observable<AuditEngagement[]> {
		const uri_ = this.apiUrl + '/AuditEngagement/';

		return this
			.http.get<AuditEngagement[]>(uri_);
	}
	
			//********************************************************************
	// assigns a AuditProgram on a AuditEngagement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAuditProgram( auditEngagementId, _auditProgramId ): Observable<any> {

		// get the AuditEngagement from storage
		this.loadHelper( auditEngagementId );

	// get the AuditProgram from storage
	var tmp 	= new AuditProgramService(this.http).getAuditProgram(_auditProgramId);

	// assign the AuditProgram
	this.auditEngagement.auditProgram = tmp;

	// save the AuditEngagement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AuditProgram on a AuditEngagement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAuditProgram( auditEngagementId ): Observable<any> {

		// get the AuditEngagement from storage
		this.loadHelper( auditEngagementId );

	// assign AuditProgram to null
	this.auditEngagement.auditProgram = null;

	// save the AuditEngagement
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more businessUnitsIds as a BusinessUnits
	// to a AuditEngagement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBusinessUnits( auditEngagementId, businessUnitsIds ): Observable<any> {

		// get the AuditEngagement
		this.loadHelper( auditEngagementId );

	// split on a comma with no spaces
	var idList = businessUnitsIds.split(',')

	// iterate over array of businessUnits ids
	idList.forEach(function (id) {
		// read the BusinessUnit
		var businessUnit = new BusinessUnitService(this.http).getBusinessUnit(id);
		// add the BusinessUnit if not already assigned
		if ( this.auditEngagement.businessUnits.indexOf(businessUnit) == -1 )
		this.auditEngagement.businessUnits.push(businessUnit);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more businessUnitsIds as a BusinessUnits
	// from a AuditEngagement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBusinessUnits( auditEngagementId, businessUnitsIds ): Observable<any> {

		// get the AuditEngagement
		this.loadHelper( auditEngagementId );


	// split on a comma with no spaces
	var idList 					= businessUnitsIds.split(',');
	var businessUnits 	= this.auditEngagement.businessUnits;

	if ( businessUnits != null && businessUnitsIds != null ) {

		// iterate over array of businessUnits ids
		businessUnits.forEach(function (obj) {
			if ( businessUnitsIds.indexOf(obj._id) > -1 ) {
				// remove the BusinessUnit
				this.auditEngagement.businessUnits.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more controlTestsIds as a ControlTests
	// to a AuditEngagement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addControlTests( auditEngagementId, controlTestsIds ): Observable<any> {

		// get the AuditEngagement
		this.loadHelper( auditEngagementId );

	// split on a comma with no spaces
	var idList = controlTestsIds.split(',')

	// iterate over array of controlTests ids
	idList.forEach(function (id) {
		// read the ControlTest_
		var controlTest_ = new ControlTest_Service(this.http).getControlTest_(id);
		// add the ControlTest_ if not already assigned
		if ( this.auditEngagement.controlTests.indexOf(controlTest_) == -1 )
		this.auditEngagement.controlTests.push(controlTest_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more controlTestsIds as a ControlTests
	// from a AuditEngagement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeControlTests( auditEngagementId, controlTestsIds ): Observable<any> {

		// get the AuditEngagement
		this.loadHelper( auditEngagementId );


	// split on a comma with no spaces
	var idList 					= controlTestsIds.split(',');
	var controlTests 	= this.auditEngagement.controlTests;

	if ( controlTests != null && controlTestsIds != null ) {

		// iterate over array of controlTests ids
		controlTests.forEach(function (obj) {
			if ( controlTestsIds.indexOf(obj._id) > -1 ) {
				// remove the ControlTest_
				this.auditEngagement.controlTests.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more workpapersIds as a Workpapers
	// to a AuditEngagement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWorkpapers( auditEngagementId, workpapersIds ): Observable<any> {

		// get the AuditEngagement
		this.loadHelper( auditEngagementId );

	// split on a comma with no spaces
	var idList = workpapersIds.split(',')

	// iterate over array of workpapers ids
	idList.forEach(function (id) {
		// read the AuditWorkpaper
		var auditWorkpaper = new AuditWorkpaperService(this.http).getAuditWorkpaper(id);
		// add the AuditWorkpaper if not already assigned
		if ( this.auditEngagement.workpapers.indexOf(auditWorkpaper) == -1 )
		this.auditEngagement.workpapers.push(auditWorkpaper);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more workpapersIds as a Workpapers
	// from a AuditEngagement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWorkpapers( auditEngagementId, workpapersIds ): Observable<any> {

		// get the AuditEngagement
		this.loadHelper( auditEngagementId );


	// split on a comma with no spaces
	var idList 					= workpapersIds.split(',');
	var workpapers 	= this.auditEngagement.workpapers;

	if ( workpapers != null && workpapersIds != null ) {

		// iterate over array of workpapers ids
		workpapers.forEach(function (obj) {
			if ( workpapersIds.indexOf(obj._id) > -1 ) {
				// remove the AuditWorkpaper
				this.auditEngagement.workpapers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more findingsIds as a Findings
	// to a AuditEngagement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFindings( auditEngagementId, findingsIds ): Observable<any> {

		// get the AuditEngagement
		this.loadHelper( auditEngagementId );

	// split on a comma with no spaces
	var idList = findingsIds.split(',')

	// iterate over array of findings ids
	idList.forEach(function (id) {
		// read the AuditFinding
		var auditFinding = new AuditFindingService(this.http).getAuditFinding(id);
		// add the AuditFinding if not already assigned
		if ( this.auditEngagement.findings.indexOf(auditFinding) == -1 )
		this.auditEngagement.findings.push(auditFinding);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more findingsIds as a Findings
	// from a AuditEngagement
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFindings( auditEngagementId, findingsIds ): Observable<any> {

		// get the AuditEngagement
		this.loadHelper( auditEngagementId );


	// split on a comma with no spaces
	var idList 					= findingsIds.split(',');
	var findings 	= this.auditEngagement.findings;

	if ( findings != null && findingsIds != null ) {

		// iterate over array of findings ids
		findings.forEach(function (obj) {
			if ( findingsIds.indexOf(obj._id) > -1 ) {
				// remove the AuditFinding
				this.auditEngagement.findings.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AuditEngagement
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AuditEngagement/update/' + this.auditEngagement;

	return  this.http.post(uri_, this.auditEngagement );
}

	//********************************************************************
	// loadHelper - internal helper to load a AuditEngagement
	//********************************************************************	
	loadHelper( id ) {
		this.getAuditEngagement(id)
			.subscribe((res : AuditEngagement) => {
				this.auditEngagement = res;
			});
	}
}