import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ProcedureOrder} from '../models/ProcedureOrder';
import {ClinicalOrderService} from '../services/ClinicalOrder.service';
import {FacilityService} from '../services/Facility.service';
import {ProcedureService} from '../services/Procedure.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ProcedureOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	procedureOrder : ProcedureOrder;

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
	// add a ProcedureOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addProcedureOrder(procedureCode, consentObtained, Order, Facility, Procedure, AnesthesiaType) : Observable<any> {
		const uri_ = this.apiUrl + '/ProcedureOrder/create';
		const obj = {
			      		procedureCode: procedureCode,
      		consentObtained: consentObtained,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		Procedure: Procedure != null && Procedure.length > 0 ? Procedure : null,
			AnesthesiaType: AnesthesiaType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ProcedureOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateProcedureOrder(procedureCode, consentObtained, Order, Facility, Procedure, AnesthesiaType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ProcedureOrder/update/' + id;
		const obj = {
				      		procedureCode: procedureCode,
      		consentObtained: consentObtained,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		Procedure: Procedure != null && Procedure.length > 0 ? Procedure : null,
			AnesthesiaType: AnesthesiaType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ProcedureOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteProcedureOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ProcedureOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ProcedureOrder
	// returns the results untouched as an Observable ProcedureOrder
	// ProcedureOrder model
	// delegates via URI
	//********************************************************************
	getProcedureOrder(id) : Observable<ProcedureOrder> {
		const uri_ = this.apiUrl + '/ProcedureOrder/load/' + id;

		return this.http.get<ProcedureOrder>(uri_);
	}
	
	//********************************************************************
	// gets all ProcedureOrder
	// returns the results untouched as JSON representation of an
	// Observable array of ProcedureOrder models
	// delegates via URI
	//********************************************************************
	getProcedureOrders() : Observable<ProcedureOrder[]> {
		const uri_ = this.apiUrl + '/ProcedureOrder/';

		return this
			.http.get<ProcedureOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Order on a ProcedureOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrder( procedureOrderId, _orderId ): Observable<any> {

		// get the ProcedureOrder from storage
		this.loadHelper( procedureOrderId );

	// get the ClinicalOrder from storage
	var tmp 	= new ClinicalOrderService(this.http).getClinicalOrder(_orderId);

	// assign the Order
	this.procedureOrder.order = tmp;

	// save the ProcedureOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Order on a ProcedureOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrder( procedureOrderId ): Observable<any> {

		// get the ProcedureOrder from storage
		this.loadHelper( procedureOrderId );

	// assign Order to null
	this.procedureOrder.order = null;

	// save the ProcedureOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Facility on a ProcedureOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFacility( procedureOrderId, _facilityId ): Observable<any> {

		// get the ProcedureOrder from storage
		this.loadHelper( procedureOrderId );

	// get the Facility from storage
	var tmp 	= new FacilityService(this.http).getFacility(_facilityId);

	// assign the Facility
	this.procedureOrder.facility = tmp;

	// save the ProcedureOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Facility on a ProcedureOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFacility( procedureOrderId ): Observable<any> {

		// get the ProcedureOrder from storage
		this.loadHelper( procedureOrderId );

	// assign Facility to null
	this.procedureOrder.facility = null;

	// save the ProcedureOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Procedure on a ProcedureOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProcedure( procedureOrderId, _procedureId ): Observable<any> {

		// get the ProcedureOrder from storage
		this.loadHelper( procedureOrderId );

	// get the Procedure from storage
	var tmp 	= new ProcedureService(this.http).getProcedure(_procedureId);

	// assign the Procedure
	this.procedureOrder.procedure = tmp;

	// save the ProcedureOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Procedure on a ProcedureOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProcedure( procedureOrderId ): Observable<any> {

		// get the ProcedureOrder from storage
		this.loadHelper( procedureOrderId );

	// assign Procedure to null
	this.procedureOrder.procedure = null;

	// save the ProcedureOrder
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ProcedureOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ProcedureOrder/update/' + this.procedureOrder;

	return  this.http.post(uri_, this.procedureOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a ProcedureOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getProcedureOrder(id)
			.subscribe((res : ProcedureOrder) => {
				this.procedureOrder = res;
			});
	}
}