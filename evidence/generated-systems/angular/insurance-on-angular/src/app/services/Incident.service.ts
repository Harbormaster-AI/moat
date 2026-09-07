import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Incident} from '../models/Incident';
import {ClaimService} from '../services/Claim.service';
import {InsuredObjectService} from '../services/InsuredObject.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class IncidentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	incident : Incident;

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
	// add a Incident
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addIncident(location, description, Claim, InsuredObjects, IncidentType) : Observable<any> {
		const uri_ = this.apiUrl + '/Incident/create';
		const obj = {
			      		location: location,
      		description: description,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		InsuredObjects: InsuredObjects != null && InsuredObjects.length > 0 ? InsuredObjects : null,
			IncidentType: IncidentType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Incident
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateIncident(location, description, Claim, InsuredObjects, IncidentType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Incident/update/' + id;
		const obj = {
				      		location: location,
      		description: description,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		InsuredObjects: InsuredObjects != null && InsuredObjects.length > 0 ? InsuredObjects : null,
			IncidentType: IncidentType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Incident
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteIncident(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Incident/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Incident
	// returns the results untouched as an Observable Incident
	// Incident model
	// delegates via URI
	//********************************************************************
	getIncident(id) : Observable<Incident> {
		const uri_ = this.apiUrl + '/Incident/load/' + id;

		return this.http.get<Incident>(uri_);
	}
	
	//********************************************************************
	// gets all Incident
	// returns the results untouched as JSON representation of an
	// Observable array of Incident models
	// delegates via URI
	//********************************************************************
	getIncidents() : Observable<Incident[]> {
		const uri_ = this.apiUrl + '/Incident/';

		return this
			.http.get<Incident[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Claim on a Incident
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignClaim( incidentId, _claimId ): Observable<any> {

		// get the Incident from storage
		this.loadHelper( incidentId );

	// get the Claim from storage
	var tmp 	= new ClaimService(this.http).getClaim(_claimId);

	// assign the Claim
	this.incident.claim = tmp;

	// save the Incident
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Claim on a Incident
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignClaim( incidentId ): Observable<any> {

		// get the Incident from storage
		this.loadHelper( incidentId );

	// assign Claim to null
	this.incident.claim = null;

	// save the Incident
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more insuredObjectsIds as a InsuredObjects
	// to a Incident
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInsuredObjects( incidentId, insuredObjectsIds ): Observable<any> {

		// get the Incident
		this.loadHelper( incidentId );

	// split on a comma with no spaces
	var idList = insuredObjectsIds.split(',')

	// iterate over array of insuredObjects ids
	idList.forEach(function (id) {
		// read the InsuredObject
		var insuredObject = new InsuredObjectService(this.http).getInsuredObject(id);
		// add the InsuredObject if not already assigned
		if ( this.incident.insuredObjects.indexOf(insuredObject) == -1 )
		this.incident.insuredObjects.push(insuredObject);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more insuredObjectsIds as a InsuredObjects
	// from a Incident
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInsuredObjects( incidentId, insuredObjectsIds ): Observable<any> {

		// get the Incident
		this.loadHelper( incidentId );


	// split on a comma with no spaces
	var idList 					= insuredObjectsIds.split(',');
	var insuredObjects 	= this.incident.insuredObjects;

	if ( insuredObjects != null && insuredObjectsIds != null ) {

		// iterate over array of insuredObjects ids
		insuredObjects.forEach(function (obj) {
			if ( insuredObjectsIds.indexOf(obj._id) > -1 ) {
				// remove the InsuredObject
				this.incident.insuredObjects.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Incident
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Incident/update/' + this.incident;

	return  this.http.post(uri_, this.incident );
}

	//********************************************************************
	// loadHelper - internal helper to load a Incident
	//********************************************************************	
	loadHelper( id ) {
		this.getIncident(id)
			.subscribe((res : Incident) => {
				this.incident = res;
			});
	}
}