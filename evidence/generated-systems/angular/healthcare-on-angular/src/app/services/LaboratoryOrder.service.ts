import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {LaboratoryOrder} from '../models/LaboratoryOrder';
import {ClinicalOrderService} from '../services/ClinicalOrder.service';
import {LaboratoryService} from '../services/Laboratory.service';
import {LabResultService} from '../services/LabResult.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LaboratoryOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	laboratoryOrder : LaboratoryOrder;

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
	// add a LaboratoryOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLaboratoryOrder(testCode, fastingRequired, Order, Laboratory, Results, SpecimenType) : Observable<any> {
		const uri_ = this.apiUrl + '/LaboratoryOrder/create';
		const obj = {
			      		testCode: testCode,
      		fastingRequired: fastingRequired,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		Laboratory: Laboratory != null && Laboratory.length > 0 ? Laboratory : null,
      		Results: Results != null && Results.length > 0 ? Results : null,
			SpecimenType: SpecimenType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a LaboratoryOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLaboratoryOrder(testCode, fastingRequired, Order, Laboratory, Results, SpecimenType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/LaboratoryOrder/update/' + id;
		const obj = {
				      		testCode: testCode,
      		fastingRequired: fastingRequired,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		Laboratory: Laboratory != null && Laboratory.length > 0 ? Laboratory : null,
      		Results: Results != null && Results.length > 0 ? Results : null,
			SpecimenType: SpecimenType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a LaboratoryOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLaboratoryOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/LaboratoryOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a LaboratoryOrder
	// returns the results untouched as an Observable LaboratoryOrder
	// LaboratoryOrder model
	// delegates via URI
	//********************************************************************
	getLaboratoryOrder(id) : Observable<LaboratoryOrder> {
		const uri_ = this.apiUrl + '/LaboratoryOrder/load/' + id;

		return this.http.get<LaboratoryOrder>(uri_);
	}
	
	//********************************************************************
	// gets all LaboratoryOrder
	// returns the results untouched as JSON representation of an
	// Observable array of LaboratoryOrder models
	// delegates via URI
	//********************************************************************
	getLaboratoryOrders() : Observable<LaboratoryOrder[]> {
		const uri_ = this.apiUrl + '/LaboratoryOrder/';

		return this
			.http.get<LaboratoryOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Order on a LaboratoryOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrder( laboratoryOrderId, _orderId ): Observable<any> {

		// get the LaboratoryOrder from storage
		this.loadHelper( laboratoryOrderId );

	// get the ClinicalOrder from storage
	var tmp 	= new ClinicalOrderService(this.http).getClinicalOrder(_orderId);

	// assign the Order
	this.laboratoryOrder.order = tmp;

	// save the LaboratoryOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Order on a LaboratoryOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrder( laboratoryOrderId ): Observable<any> {

		// get the LaboratoryOrder from storage
		this.loadHelper( laboratoryOrderId );

	// assign Order to null
	this.laboratoryOrder.order = null;

	// save the LaboratoryOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Laboratory on a LaboratoryOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLaboratory( laboratoryOrderId, _laboratoryId ): Observable<any> {

		// get the LaboratoryOrder from storage
		this.loadHelper( laboratoryOrderId );

	// get the Laboratory from storage
	var tmp 	= new LaboratoryService(this.http).getLaboratory(_laboratoryId);

	// assign the Laboratory
	this.laboratoryOrder.laboratory = tmp;

	// save the LaboratoryOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Laboratory on a LaboratoryOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLaboratory( laboratoryOrderId ): Observable<any> {

		// get the LaboratoryOrder from storage
		this.loadHelper( laboratoryOrderId );

	// assign Laboratory to null
	this.laboratoryOrder.laboratory = null;

	// save the LaboratoryOrder
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more resultsIds as a Results
	// to a LaboratoryOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addResults( laboratoryOrderId, resultsIds ): Observable<any> {

		// get the LaboratoryOrder
		this.loadHelper( laboratoryOrderId );

	// split on a comma with no spaces
	var idList = resultsIds.split(',')

	// iterate over array of results ids
	idList.forEach(function (id) {
		// read the LabResult
		var labResult = new LabResultService(this.http).getLabResult(id);
		// add the LabResult if not already assigned
		if ( this.laboratoryOrder.results.indexOf(labResult) == -1 )
		this.laboratoryOrder.results.push(labResult);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more resultsIds as a Results
	// from a LaboratoryOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeResults( laboratoryOrderId, resultsIds ): Observable<any> {

		// get the LaboratoryOrder
		this.loadHelper( laboratoryOrderId );


	// split on a comma with no spaces
	var idList 					= resultsIds.split(',');
	var results 	= this.laboratoryOrder.results;

	if ( results != null && resultsIds != null ) {

		// iterate over array of results ids
		results.forEach(function (obj) {
			if ( resultsIds.indexOf(obj._id) > -1 ) {
				// remove the LabResult
				this.laboratoryOrder.results.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a LaboratoryOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/LaboratoryOrder/update/' + this.laboratoryOrder;

	return  this.http.post(uri_, this.laboratoryOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a LaboratoryOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getLaboratoryOrder(id)
			.subscribe((res : LaboratoryOrder) => {
				this.laboratoryOrder = res;
			});
	}
}