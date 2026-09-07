import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {LabResult} from '../models/LabResult';
import {LaboratoryOrderService} from '../services/LaboratoryOrder.service';
import {ObservationService} from '../services/Observation.service';
import {LaboratoryService} from '../services/Laboratory.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LabResultService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	labResult : LabResult;

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
	// add a LabResult
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLabResult(resultCode, issuedDate, LaboratoryOrder, Observations, Laboratory, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/LabResult/create';
		const obj = {
			      		resultCode: resultCode,
      		issuedDate: issuedDate,
      		LaboratoryOrder: LaboratoryOrder != null && LaboratoryOrder.length > 0 ? LaboratoryOrder : null,
      		Observations: Observations != null && Observations.length > 0 ? Observations : null,
      		Laboratory: Laboratory != null && Laboratory.length > 0 ? Laboratory : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a LabResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLabResult(resultCode, issuedDate, LaboratoryOrder, Observations, Laboratory, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/LabResult/update/' + id;
		const obj = {
				      		resultCode: resultCode,
      		issuedDate: issuedDate,
      		LaboratoryOrder: LaboratoryOrder != null && LaboratoryOrder.length > 0 ? LaboratoryOrder : null,
      		Observations: Observations != null && Observations.length > 0 ? Observations : null,
      		Laboratory: Laboratory != null && Laboratory.length > 0 ? Laboratory : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a LabResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLabResult(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/LabResult/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a LabResult
	// returns the results untouched as an Observable LabResult
	// LabResult model
	// delegates via URI
	//********************************************************************
	getLabResult(id) : Observable<LabResult> {
		const uri_ = this.apiUrl + '/LabResult/load/' + id;

		return this.http.get<LabResult>(uri_);
	}
	
	//********************************************************************
	// gets all LabResult
	// returns the results untouched as JSON representation of an
	// Observable array of LabResult models
	// delegates via URI
	//********************************************************************
	getLabResults() : Observable<LabResult[]> {
		const uri_ = this.apiUrl + '/LabResult/';

		return this
			.http.get<LabResult[]>(uri_);
	}
	
			//********************************************************************
	// assigns a LaboratoryOrder on a LabResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLaboratoryOrder( labResultId, _laboratoryOrderId ): Observable<any> {

		// get the LabResult from storage
		this.loadHelper( labResultId );

	// get the LaboratoryOrder from storage
	var tmp 	= new LaboratoryOrderService(this.http).getLaboratoryOrder(_laboratoryOrderId);

	// assign the LaboratoryOrder
	this.labResult.laboratoryOrder = tmp;

	// save the LabResult
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LaboratoryOrder on a LabResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLaboratoryOrder( labResultId ): Observable<any> {

		// get the LabResult from storage
		this.loadHelper( labResultId );

	// assign LaboratoryOrder to null
	this.labResult.laboratoryOrder = null;

	// save the LabResult
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Laboratory on a LabResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLaboratory( labResultId, _laboratoryId ): Observable<any> {

		// get the LabResult from storage
		this.loadHelper( labResultId );

	// get the Laboratory from storage
	var tmp 	= new LaboratoryService(this.http).getLaboratory(_laboratoryId);

	// assign the Laboratory
	this.labResult.laboratory = tmp;

	// save the LabResult
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Laboratory on a LabResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLaboratory( labResultId ): Observable<any> {

		// get the LabResult from storage
		this.loadHelper( labResultId );

	// assign Laboratory to null
	this.labResult.laboratory = null;

	// save the LabResult
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more observationsIds as a Observations
	// to a LabResult
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addObservations( labResultId, observationsIds ): Observable<any> {

		// get the LabResult
		this.loadHelper( labResultId );

	// split on a comma with no spaces
	var idList = observationsIds.split(',')

	// iterate over array of observations ids
	idList.forEach(function (id) {
		// read the Observation
		var observation = new ObservationService(this.http).getObservation(id);
		// add the Observation if not already assigned
		if ( this.labResult.observations.indexOf(observation) == -1 )
		this.labResult.observations.push(observation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more observationsIds as a Observations
	// from a LabResult
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeObservations( labResultId, observationsIds ): Observable<any> {

		// get the LabResult
		this.loadHelper( labResultId );


	// split on a comma with no spaces
	var idList 					= observationsIds.split(',');
	var observations 	= this.labResult.observations;

	if ( observations != null && observationsIds != null ) {

		// iterate over array of observations ids
		observations.forEach(function (obj) {
			if ( observationsIds.indexOf(obj._id) > -1 ) {
				// remove the Observation
				this.labResult.observations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a LabResult
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/LabResult/update/' + this.labResult;

	return  this.http.post(uri_, this.labResult );
}

	//********************************************************************
	// loadHelper - internal helper to load a LabResult
	//********************************************************************	
	loadHelper( id ) {
		this.getLabResult(id)
			.subscribe((res : LabResult) => {
				this.labResult = res;
			});
	}
}