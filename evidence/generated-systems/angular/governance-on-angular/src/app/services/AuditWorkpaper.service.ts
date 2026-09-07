import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AuditWorkpaper} from '../models/AuditWorkpaper';
import {AuditEngagementService} from '../services/AuditEngagement.service';
import {EvidenceService} from '../services/Evidence.service';
import {AuditFindingService} from '../services/AuditFinding.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AuditWorkpaperService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	auditWorkpaper : AuditWorkpaper;

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
	// add a AuditWorkpaper
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAuditWorkpaper(workpaperRef, subject, workpaperUrl, Engagement, Evidence, Findings) : Observable<any> {
		const uri_ = this.apiUrl + '/AuditWorkpaper/create';
		const obj = {
			      		workpaperRef: workpaperRef,
      		subject: subject,
      		workpaperUrl: workpaperUrl,
      		Engagement: Engagement != null && Engagement.length > 0 ? Engagement : null,
      		Evidence: Evidence != null && Evidence.length > 0 ? Evidence : null,
			Findings: Findings != null && Findings.length > 0 ? Findings : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AuditWorkpaper
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAuditWorkpaper(workpaperRef, subject, workpaperUrl, Engagement, Evidence, Findings, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AuditWorkpaper/update/' + id;
		const obj = {
				      		workpaperRef: workpaperRef,
      		subject: subject,
      		workpaperUrl: workpaperUrl,
      		Engagement: Engagement != null && Engagement.length > 0 ? Engagement : null,
      		Evidence: Evidence != null && Evidence.length > 0 ? Evidence : null,
			Findings: Findings != null && Findings.length > 0 ? Findings : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AuditWorkpaper
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAuditWorkpaper(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AuditWorkpaper/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AuditWorkpaper
	// returns the results untouched as an Observable AuditWorkpaper
	// AuditWorkpaper model
	// delegates via URI
	//********************************************************************
	getAuditWorkpaper(id) : Observable<AuditWorkpaper> {
		const uri_ = this.apiUrl + '/AuditWorkpaper/load/' + id;

		return this.http.get<AuditWorkpaper>(uri_);
	}
	
	//********************************************************************
	// gets all AuditWorkpaper
	// returns the results untouched as JSON representation of an
	// Observable array of AuditWorkpaper models
	// delegates via URI
	//********************************************************************
	getAuditWorkpapers() : Observable<AuditWorkpaper[]> {
		const uri_ = this.apiUrl + '/AuditWorkpaper/';

		return this
			.http.get<AuditWorkpaper[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Engagement on a AuditWorkpaper
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEngagement( auditWorkpaperId, _engagementId ): Observable<any> {

		// get the AuditWorkpaper from storage
		this.loadHelper( auditWorkpaperId );

	// get the AuditEngagement from storage
	var tmp 	= new AuditEngagementService(this.http).getAuditEngagement(_engagementId);

	// assign the Engagement
	this.auditWorkpaper.engagement = tmp;

	// save the AuditWorkpaper
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Engagement on a AuditWorkpaper
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEngagement( auditWorkpaperId ): Observable<any> {

		// get the AuditWorkpaper from storage
		this.loadHelper( auditWorkpaperId );

	// assign Engagement to null
	this.auditWorkpaper.engagement = null;

	// save the AuditWorkpaper
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more evidenceIds as a Evidence
	// to a AuditWorkpaper
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEvidence( auditWorkpaperId, evidenceIds ): Observable<any> {

		// get the AuditWorkpaper
		this.loadHelper( auditWorkpaperId );

	// split on a comma with no spaces
	var idList = evidenceIds.split(',')

	// iterate over array of evidence ids
	idList.forEach(function (id) {
		// read the Evidence
		var evidence = new EvidenceService(this.http).getEvidence(id);
		// add the Evidence if not already assigned
		if ( this.auditWorkpaper.evidence.indexOf(evidence) == -1 )
		this.auditWorkpaper.evidence.push(evidence);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more evidenceIds as a Evidence
	// from a AuditWorkpaper
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEvidence( auditWorkpaperId, evidenceIds ): Observable<any> {

		// get the AuditWorkpaper
		this.loadHelper( auditWorkpaperId );


	// split on a comma with no spaces
	var idList 					= evidenceIds.split(',');
	var evidence 	= this.auditWorkpaper.evidence;

	if ( evidence != null && evidenceIds != null ) {

		// iterate over array of evidence ids
		evidence.forEach(function (obj) {
			if ( evidenceIds.indexOf(obj._id) > -1 ) {
				// remove the Evidence
				this.auditWorkpaper.evidence.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more findingsIds as a Findings
	// to a AuditWorkpaper
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFindings( auditWorkpaperId, findingsIds ): Observable<any> {

		// get the AuditWorkpaper
		this.loadHelper( auditWorkpaperId );

	// split on a comma with no spaces
	var idList = findingsIds.split(',')

	// iterate over array of findings ids
	idList.forEach(function (id) {
		// read the AuditFinding
		var auditFinding = new AuditFindingService(this.http).getAuditFinding(id);
		// add the AuditFinding if not already assigned
		if ( this.auditWorkpaper.findings.indexOf(auditFinding) == -1 )
		this.auditWorkpaper.findings.push(auditFinding);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more findingsIds as a Findings
	// from a AuditWorkpaper
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFindings( auditWorkpaperId, findingsIds ): Observable<any> {

		// get the AuditWorkpaper
		this.loadHelper( auditWorkpaperId );


	// split on a comma with no spaces
	var idList 					= findingsIds.split(',');
	var findings 	= this.auditWorkpaper.findings;

	if ( findings != null && findingsIds != null ) {

		// iterate over array of findings ids
		findings.forEach(function (obj) {
			if ( findingsIds.indexOf(obj._id) > -1 ) {
				// remove the AuditFinding
				this.auditWorkpaper.findings.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AuditWorkpaper
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AuditWorkpaper/update/' + this.auditWorkpaper;

	return  this.http.post(uri_, this.auditWorkpaper );
}

	//********************************************************************
	// loadHelper - internal helper to load a AuditWorkpaper
	//********************************************************************	
	loadHelper( id ) {
		this.getAuditWorkpaper(id)
			.subscribe((res : AuditWorkpaper) => {
				this.auditWorkpaper = res;
			});
	}
}