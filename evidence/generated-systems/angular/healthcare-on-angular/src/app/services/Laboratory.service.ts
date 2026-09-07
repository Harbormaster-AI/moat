import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Laboratory} from '../models/Laboratory';
import {FacilityService} from '../services/Facility.service';
import {LaboratoryOrderService} from '../services/LaboratoryOrder.service';
import {LabResultService} from '../services/LabResult.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LaboratoryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	laboratory : Laboratory;

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
	// add a Laboratory
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLaboratory(name, cliaNumber, Facility, LaboratoryOrders, LabResults) : Observable<any> {
		const uri_ = this.apiUrl + '/Laboratory/create';
		const obj = {
			      		name: name,
      		cliaNumber: cliaNumber,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		LaboratoryOrders: LaboratoryOrders != null && LaboratoryOrders.length > 0 ? LaboratoryOrders : null,
			LabResults: LabResults != null && LabResults.length > 0 ? LabResults : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Laboratory
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLaboratory(name, cliaNumber, Facility, LaboratoryOrders, LabResults, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Laboratory/update/' + id;
		const obj = {
				      		name: name,
      		cliaNumber: cliaNumber,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		LaboratoryOrders: LaboratoryOrders != null && LaboratoryOrders.length > 0 ? LaboratoryOrders : null,
			LabResults: LabResults != null && LabResults.length > 0 ? LabResults : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Laboratory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLaboratory(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Laboratory/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Laboratory
	// returns the results untouched as an Observable Laboratory
	// Laboratory model
	// delegates via URI
	//********************************************************************
	getLaboratory(id) : Observable<Laboratory> {
		const uri_ = this.apiUrl + '/Laboratory/load/' + id;

		return this.http.get<Laboratory>(uri_);
	}
	
	//********************************************************************
	// gets all Laboratory
	// returns the results untouched as JSON representation of an
	// Observable array of Laboratory models
	// delegates via URI
	//********************************************************************
	getLaboratorys() : Observable<Laboratory[]> {
		const uri_ = this.apiUrl + '/Laboratory/';

		return this
			.http.get<Laboratory[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Facility on a Laboratory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFacility( laboratoryId, _facilityId ): Observable<any> {

		// get the Laboratory from storage
		this.loadHelper( laboratoryId );

	// get the Facility from storage
	var tmp 	= new FacilityService(this.http).getFacility(_facilityId);

	// assign the Facility
	this.laboratory.facility = tmp;

	// save the Laboratory
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Facility on a Laboratory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFacility( laboratoryId ): Observable<any> {

		// get the Laboratory from storage
		this.loadHelper( laboratoryId );

	// assign Facility to null
	this.laboratory.facility = null;

	// save the Laboratory
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more laboratoryOrdersIds as a LaboratoryOrders
	// to a Laboratory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLaboratoryOrders( laboratoryId, laboratoryOrdersIds ): Observable<any> {

		// get the Laboratory
		this.loadHelper( laboratoryId );

	// split on a comma with no spaces
	var idList = laboratoryOrdersIds.split(',')

	// iterate over array of laboratoryOrders ids
	idList.forEach(function (id) {
		// read the LaboratoryOrder
		var laboratoryOrder = new LaboratoryOrderService(this.http).getLaboratoryOrder(id);
		// add the LaboratoryOrder if not already assigned
		if ( this.laboratory.laboratoryOrders.indexOf(laboratoryOrder) == -1 )
		this.laboratory.laboratoryOrders.push(laboratoryOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more laboratoryOrdersIds as a LaboratoryOrders
	// from a Laboratory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLaboratoryOrders( laboratoryId, laboratoryOrdersIds ): Observable<any> {

		// get the Laboratory
		this.loadHelper( laboratoryId );


	// split on a comma with no spaces
	var idList 					= laboratoryOrdersIds.split(',');
	var laboratoryOrders 	= this.laboratory.laboratoryOrders;

	if ( laboratoryOrders != null && laboratoryOrdersIds != null ) {

		// iterate over array of laboratoryOrders ids
		laboratoryOrders.forEach(function (obj) {
			if ( laboratoryOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the LaboratoryOrder
				this.laboratory.laboratoryOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more labResultsIds as a LabResults
	// to a Laboratory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLabResults( laboratoryId, labResultsIds ): Observable<any> {

		// get the Laboratory
		this.loadHelper( laboratoryId );

	// split on a comma with no spaces
	var idList = labResultsIds.split(',')

	// iterate over array of labResults ids
	idList.forEach(function (id) {
		// read the LabResult
		var labResult = new LabResultService(this.http).getLabResult(id);
		// add the LabResult if not already assigned
		if ( this.laboratory.labResults.indexOf(labResult) == -1 )
		this.laboratory.labResults.push(labResult);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more labResultsIds as a LabResults
	// from a Laboratory
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLabResults( laboratoryId, labResultsIds ): Observable<any> {

		// get the Laboratory
		this.loadHelper( laboratoryId );


	// split on a comma with no spaces
	var idList 					= labResultsIds.split(',');
	var labResults 	= this.laboratory.labResults;

	if ( labResults != null && labResultsIds != null ) {

		// iterate over array of labResults ids
		labResults.forEach(function (obj) {
			if ( labResultsIds.indexOf(obj._id) > -1 ) {
				// remove the LabResult
				this.laboratory.labResults.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Laboratory
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Laboratory/update/' + this.laboratory;

	return  this.http.post(uri_, this.laboratory );
}

	//********************************************************************
	// loadHelper - internal helper to load a Laboratory
	//********************************************************************	
	loadHelper( id ) {
		this.getLaboratory(id)
			.subscribe((res : Laboratory) => {
				this.laboratory = res;
			});
	}
}