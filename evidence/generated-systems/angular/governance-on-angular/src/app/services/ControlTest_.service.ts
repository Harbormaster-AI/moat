import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ControlTest_} from '../models/ControlTest_';
import {ControlService} from '../services/Control.service';
import {EvidenceService} from '../services/Evidence.service';
import {AuditEngagementService} from '../services/AuditEngagement.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ControlTest_Service extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	controlTest_ : ControlTest_;

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
	// add a ControlTest_
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addControlTest_(name, testPeriodStart, testPeriodEnd, sampleSize, Control, Evidence, Engagement, TestType, Effectiveness, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/ControlTest_/create';
		const obj = {
			      		name: name,
      		testPeriodStart: testPeriodStart,
      		testPeriodEnd: testPeriodEnd,
      		sampleSize: sampleSize,
      		Control: Control != null && Control.length > 0 ? Control : null,
      		Evidence: Evidence != null && Evidence.length > 0 ? Evidence : null,
      		Engagement: Engagement != null && Engagement.length > 0 ? Engagement : null,
      		TestType: TestType,
      		Effectiveness: Effectiveness,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ControlTest_
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateControlTest_(name, testPeriodStart, testPeriodEnd, sampleSize, Control, Evidence, Engagement, TestType, Effectiveness, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ControlTest_/update/' + id;
		const obj = {
				      		name: name,
      		testPeriodStart: testPeriodStart,
      		testPeriodEnd: testPeriodEnd,
      		sampleSize: sampleSize,
      		Control: Control != null && Control.length > 0 ? Control : null,
      		Evidence: Evidence != null && Evidence.length > 0 ? Evidence : null,
      		Engagement: Engagement != null && Engagement.length > 0 ? Engagement : null,
      		TestType: TestType,
      		Effectiveness: Effectiveness,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ControlTest_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteControlTest_(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ControlTest_/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ControlTest_
	// returns the results untouched as an Observable ControlTest_
	// ControlTest_ model
	// delegates via URI
	//********************************************************************
	getControlTest_(id) : Observable<ControlTest_> {
		const uri_ = this.apiUrl + '/ControlTest_/load/' + id;

		return this.http.get<ControlTest_>(uri_);
	}
	
	//********************************************************************
	// gets all ControlTest_
	// returns the results untouched as JSON representation of an
	// Observable array of ControlTest_ models
	// delegates via URI
	//********************************************************************
	getControlTest_s() : Observable<ControlTest_[]> {
		const uri_ = this.apiUrl + '/ControlTest_/';

		return this
			.http.get<ControlTest_[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Control on a ControlTest_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignControl( controlTest_Id, _controlId ): Observable<any> {

		// get the ControlTest_ from storage
		this.loadHelper( controlTest_Id );

	// get the Control from storage
	var tmp 	= new ControlService(this.http).getControl(_controlId);

	// assign the Control
	this.controlTest_.control = tmp;

	// save the ControlTest_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Control on a ControlTest_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignControl( controlTest_Id ): Observable<any> {

		// get the ControlTest_ from storage
		this.loadHelper( controlTest_Id );

	// assign Control to null
	this.controlTest_.control = null;

	// save the ControlTest_
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Engagement on a ControlTest_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEngagement( controlTest_Id, _engagementId ): Observable<any> {

		// get the ControlTest_ from storage
		this.loadHelper( controlTest_Id );

	// get the AuditEngagement from storage
	var tmp 	= new AuditEngagementService(this.http).getAuditEngagement(_engagementId);

	// assign the Engagement
	this.controlTest_.engagement = tmp;

	// save the ControlTest_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Engagement on a ControlTest_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEngagement( controlTest_Id ): Observable<any> {

		// get the ControlTest_ from storage
		this.loadHelper( controlTest_Id );

	// assign Engagement to null
	this.controlTest_.engagement = null;

	// save the ControlTest_
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more evidenceIds as a Evidence
	// to a ControlTest_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEvidence( controlTest_Id, evidenceIds ): Observable<any> {

		// get the ControlTest_
		this.loadHelper( controlTest_Id );

	// split on a comma with no spaces
	var idList = evidenceIds.split(',')

	// iterate over array of evidence ids
	idList.forEach(function (id) {
		// read the Evidence
		var evidence = new EvidenceService(this.http).getEvidence(id);
		// add the Evidence if not already assigned
		if ( this.controlTest_.evidence.indexOf(evidence) == -1 )
		this.controlTest_.evidence.push(evidence);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more evidenceIds as a Evidence
	// from a ControlTest_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEvidence( controlTest_Id, evidenceIds ): Observable<any> {

		// get the ControlTest_
		this.loadHelper( controlTest_Id );


	// split on a comma with no spaces
	var idList 					= evidenceIds.split(',');
	var evidence 	= this.controlTest_.evidence;

	if ( evidence != null && evidenceIds != null ) {

		// iterate over array of evidence ids
		evidence.forEach(function (obj) {
			if ( evidenceIds.indexOf(obj._id) > -1 ) {
				// remove the Evidence
				this.controlTest_.evidence.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ControlTest_
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ControlTest_/update/' + this.controlTest_;

	return  this.http.post(uri_, this.controlTest_ );
}

	//********************************************************************
	// loadHelper - internal helper to load a ControlTest_
	//********************************************************************	
	loadHelper( id ) {
		this.getControlTest_(id)
			.subscribe((res : ControlTest_) => {
				this.controlTest_ = res;
			});
	}
}