import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Evidence} from '../models/Evidence';
import {ControlTest_Service} from '../services/ControlTest_.service';
import {ControlService} from '../services/Control.service';
import {ObligationService} from '../services/Obligation.service';
import {AuditWorkpaperService} from '../services/AuditWorkpaper.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EvidenceService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	evidence : Evidence;

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
	// add a Evidence
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEvidence(title, locationUrl, receivedDate, ControlTest, Control, Obligation, Workpaper, EvidenceType) : Observable<any> {
		const uri_ = this.apiUrl + '/Evidence/create';
		const obj = {
			      		title: title,
      		locationUrl: locationUrl,
      		receivedDate: receivedDate,
      		ControlTest: ControlTest != null && ControlTest.length > 0 ? ControlTest : null,
      		Control: Control != null && Control.length > 0 ? Control : null,
      		Obligation: Obligation != null && Obligation.length > 0 ? Obligation : null,
      		Workpaper: Workpaper != null && Workpaper.length > 0 ? Workpaper : null,
			EvidenceType: EvidenceType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Evidence
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEvidence(title, locationUrl, receivedDate, ControlTest, Control, Obligation, Workpaper, EvidenceType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Evidence/update/' + id;
		const obj = {
				      		title: title,
      		locationUrl: locationUrl,
      		receivedDate: receivedDate,
      		ControlTest: ControlTest != null && ControlTest.length > 0 ? ControlTest : null,
      		Control: Control != null && Control.length > 0 ? Control : null,
      		Obligation: Obligation != null && Obligation.length > 0 ? Obligation : null,
      		Workpaper: Workpaper != null && Workpaper.length > 0 ? Workpaper : null,
			EvidenceType: EvidenceType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Evidence
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEvidence(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Evidence/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Evidence
	// returns the results untouched as an Observable Evidence
	// Evidence model
	// delegates via URI
	//********************************************************************
	getEvidence(id) : Observable<Evidence> {
		const uri_ = this.apiUrl + '/Evidence/load/' + id;

		return this.http.get<Evidence>(uri_);
	}
	
	//********************************************************************
	// gets all Evidence
	// returns the results untouched as JSON representation of an
	// Observable array of Evidence models
	// delegates via URI
	//********************************************************************
	getEvidences() : Observable<Evidence[]> {
		const uri_ = this.apiUrl + '/Evidence/';

		return this
			.http.get<Evidence[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ControlTest on a Evidence
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignControlTest( evidenceId, _controlTestId ): Observable<any> {

		// get the Evidence from storage
		this.loadHelper( evidenceId );

	// get the ControlTest_ from storage
	var tmp 	= new ControlTest_Service(this.http).getControlTest_(_controlTestId);

	// assign the ControlTest
	this.evidence.controlTest = tmp;

	// save the Evidence
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ControlTest on a Evidence
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignControlTest( evidenceId ): Observable<any> {

		// get the Evidence from storage
		this.loadHelper( evidenceId );

	// assign ControlTest to null
	this.evidence.controlTest = null;

	// save the Evidence
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Control on a Evidence
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignControl( evidenceId, _controlId ): Observable<any> {

		// get the Evidence from storage
		this.loadHelper( evidenceId );

	// get the Control from storage
	var tmp 	= new ControlService(this.http).getControl(_controlId);

	// assign the Control
	this.evidence.control = tmp;

	// save the Evidence
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Control on a Evidence
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignControl( evidenceId ): Observable<any> {

		// get the Evidence from storage
		this.loadHelper( evidenceId );

	// assign Control to null
	this.evidence.control = null;

	// save the Evidence
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Obligation on a Evidence
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignObligation( evidenceId, _obligationId ): Observable<any> {

		// get the Evidence from storage
		this.loadHelper( evidenceId );

	// get the Obligation from storage
	var tmp 	= new ObligationService(this.http).getObligation(_obligationId);

	// assign the Obligation
	this.evidence.obligation = tmp;

	// save the Evidence
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Obligation on a Evidence
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignObligation( evidenceId ): Observable<any> {

		// get the Evidence from storage
		this.loadHelper( evidenceId );

	// assign Obligation to null
	this.evidence.obligation = null;

	// save the Evidence
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Workpaper on a Evidence
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkpaper( evidenceId, _workpaperId ): Observable<any> {

		// get the Evidence from storage
		this.loadHelper( evidenceId );

	// get the AuditWorkpaper from storage
	var tmp 	= new AuditWorkpaperService(this.http).getAuditWorkpaper(_workpaperId);

	// assign the Workpaper
	this.evidence.workpaper = tmp;

	// save the Evidence
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workpaper on a Evidence
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkpaper( evidenceId ): Observable<any> {

		// get the Evidence from storage
		this.loadHelper( evidenceId );

	// assign Workpaper to null
	this.evidence.workpaper = null;

	// save the Evidence
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Evidence
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Evidence/update/' + this.evidence;

	return  this.http.post(uri_, this.evidence );
}

	//********************************************************************
	// loadHelper - internal helper to load a Evidence
	//********************************************************************	
	loadHelper( id ) {
		this.getEvidence(id)
			.subscribe((res : Evidence) => {
				this.evidence = res;
			});
	}
}