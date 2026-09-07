import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ClinicalOrder} from '../models/ClinicalOrder';
import {PatientService} from '../services/Patient.service';
import {EncounterService} from '../services/Encounter.service';
import {ClinicianService} from '../services/Clinician.service';
import {MedicationOrderService} from '../services/MedicationOrder.service';
import {LaboratoryOrderService} from '../services/LaboratoryOrder.service';
import {ImagingOrderService} from '../services/ImagingOrder.service';
import {ProcedureOrderService} from '../services/ProcedureOrder.service';
import {AuthorizationService} from '../services/Authorization.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ClinicalOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	clinicalOrder : ClinicalOrder;

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
	// add a ClinicalOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addClinicalOrder(orderNumber, Patient, Encounter, OrderingClinician, MedicationOrders, LaboratoryOrders, ImagingOrders, ProcedureOrders, Authorizations, Status, OrderType, Priority) : Observable<any> {
		const uri_ = this.apiUrl + '/ClinicalOrder/create';
		const obj = {
			      		orderNumber: orderNumber,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		OrderingClinician: OrderingClinician != null && OrderingClinician.length > 0 ? OrderingClinician : null,
      		MedicationOrders: MedicationOrders != null && MedicationOrders.length > 0 ? MedicationOrders : null,
      		LaboratoryOrders: LaboratoryOrders != null && LaboratoryOrders.length > 0 ? LaboratoryOrders : null,
      		ImagingOrders: ImagingOrders != null && ImagingOrders.length > 0 ? ImagingOrders : null,
      		ProcedureOrders: ProcedureOrders != null && ProcedureOrders.length > 0 ? ProcedureOrders : null,
      		Authorizations: Authorizations != null && Authorizations.length > 0 ? Authorizations : null,
      		Status: Status,
      		OrderType: OrderType,
			Priority: Priority
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ClinicalOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateClinicalOrder(orderNumber, Patient, Encounter, OrderingClinician, MedicationOrders, LaboratoryOrders, ImagingOrders, ProcedureOrders, Authorizations, Status, OrderType, Priority, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ClinicalOrder/update/' + id;
		const obj = {
				      		orderNumber: orderNumber,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		OrderingClinician: OrderingClinician != null && OrderingClinician.length > 0 ? OrderingClinician : null,
      		MedicationOrders: MedicationOrders != null && MedicationOrders.length > 0 ? MedicationOrders : null,
      		LaboratoryOrders: LaboratoryOrders != null && LaboratoryOrders.length > 0 ? LaboratoryOrders : null,
      		ImagingOrders: ImagingOrders != null && ImagingOrders.length > 0 ? ImagingOrders : null,
      		ProcedureOrders: ProcedureOrders != null && ProcedureOrders.length > 0 ? ProcedureOrders : null,
      		Authorizations: Authorizations != null && Authorizations.length > 0 ? Authorizations : null,
      		Status: Status,
      		OrderType: OrderType,
			Priority: Priority
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ClinicalOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteClinicalOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ClinicalOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ClinicalOrder
	// returns the results untouched as an Observable ClinicalOrder
	// ClinicalOrder model
	// delegates via URI
	//********************************************************************
	getClinicalOrder(id) : Observable<ClinicalOrder> {
		const uri_ = this.apiUrl + '/ClinicalOrder/load/' + id;

		return this.http.get<ClinicalOrder>(uri_);
	}
	
	//********************************************************************
	// gets all ClinicalOrder
	// returns the results untouched as JSON representation of an
	// Observable array of ClinicalOrder models
	// delegates via URI
	//********************************************************************
	getClinicalOrders() : Observable<ClinicalOrder[]> {
		const uri_ = this.apiUrl + '/ClinicalOrder/';

		return this
			.http.get<ClinicalOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Patient on a ClinicalOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( clinicalOrderId, _patientId ): Observable<any> {

		// get the ClinicalOrder from storage
		this.loadHelper( clinicalOrderId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.clinicalOrder.patient = tmp;

	// save the ClinicalOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a ClinicalOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( clinicalOrderId ): Observable<any> {

		// get the ClinicalOrder from storage
		this.loadHelper( clinicalOrderId );

	// assign Patient to null
	this.clinicalOrder.patient = null;

	// save the ClinicalOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Encounter on a ClinicalOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEncounter( clinicalOrderId, _encounterId ): Observable<any> {

		// get the ClinicalOrder from storage
		this.loadHelper( clinicalOrderId );

	// get the Encounter from storage
	var tmp 	= new EncounterService(this.http).getEncounter(_encounterId);

	// assign the Encounter
	this.clinicalOrder.encounter = tmp;

	// save the ClinicalOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Encounter on a ClinicalOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEncounter( clinicalOrderId ): Observable<any> {

		// get the ClinicalOrder from storage
		this.loadHelper( clinicalOrderId );

	// assign Encounter to null
	this.clinicalOrder.encounter = null;

	// save the ClinicalOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a OrderingClinician on a ClinicalOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrderingClinician( clinicalOrderId, _orderingClinicianId ): Observable<any> {

		// get the ClinicalOrder from storage
		this.loadHelper( clinicalOrderId );

	// get the Clinician from storage
	var tmp 	= new ClinicianService(this.http).getClinician(_orderingClinicianId);

	// assign the OrderingClinician
	this.clinicalOrder.orderingClinician = tmp;

	// save the ClinicalOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a OrderingClinician on a ClinicalOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrderingClinician( clinicalOrderId ): Observable<any> {

		// get the ClinicalOrder from storage
		this.loadHelper( clinicalOrderId );

	// assign OrderingClinician to null
	this.clinicalOrder.orderingClinician = null;

	// save the ClinicalOrder
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more medicationOrdersIds as a MedicationOrders
	// to a ClinicalOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMedicationOrders( clinicalOrderId, medicationOrdersIds ): Observable<any> {

		// get the ClinicalOrder
		this.loadHelper( clinicalOrderId );

	// split on a comma with no spaces
	var idList = medicationOrdersIds.split(',')

	// iterate over array of medicationOrders ids
	idList.forEach(function (id) {
		// read the MedicationOrder
		var medicationOrder = new MedicationOrderService(this.http).getMedicationOrder(id);
		// add the MedicationOrder if not already assigned
		if ( this.clinicalOrder.medicationOrders.indexOf(medicationOrder) == -1 )
		this.clinicalOrder.medicationOrders.push(medicationOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more medicationOrdersIds as a MedicationOrders
	// from a ClinicalOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMedicationOrders( clinicalOrderId, medicationOrdersIds ): Observable<any> {

		// get the ClinicalOrder
		this.loadHelper( clinicalOrderId );


	// split on a comma with no spaces
	var idList 					= medicationOrdersIds.split(',');
	var medicationOrders 	= this.clinicalOrder.medicationOrders;

	if ( medicationOrders != null && medicationOrdersIds != null ) {

		// iterate over array of medicationOrders ids
		medicationOrders.forEach(function (obj) {
			if ( medicationOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the MedicationOrder
				this.clinicalOrder.medicationOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more laboratoryOrdersIds as a LaboratoryOrders
	// to a ClinicalOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLaboratoryOrders( clinicalOrderId, laboratoryOrdersIds ): Observable<any> {

		// get the ClinicalOrder
		this.loadHelper( clinicalOrderId );

	// split on a comma with no spaces
	var idList = laboratoryOrdersIds.split(',')

	// iterate over array of laboratoryOrders ids
	idList.forEach(function (id) {
		// read the LaboratoryOrder
		var laboratoryOrder = new LaboratoryOrderService(this.http).getLaboratoryOrder(id);
		// add the LaboratoryOrder if not already assigned
		if ( this.clinicalOrder.laboratoryOrders.indexOf(laboratoryOrder) == -1 )
		this.clinicalOrder.laboratoryOrders.push(laboratoryOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more laboratoryOrdersIds as a LaboratoryOrders
	// from a ClinicalOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLaboratoryOrders( clinicalOrderId, laboratoryOrdersIds ): Observable<any> {

		// get the ClinicalOrder
		this.loadHelper( clinicalOrderId );


	// split on a comma with no spaces
	var idList 					= laboratoryOrdersIds.split(',');
	var laboratoryOrders 	= this.clinicalOrder.laboratoryOrders;

	if ( laboratoryOrders != null && laboratoryOrdersIds != null ) {

		// iterate over array of laboratoryOrders ids
		laboratoryOrders.forEach(function (obj) {
			if ( laboratoryOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the LaboratoryOrder
				this.clinicalOrder.laboratoryOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more imagingOrdersIds as a ImagingOrders
	// to a ClinicalOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addImagingOrders( clinicalOrderId, imagingOrdersIds ): Observable<any> {

		// get the ClinicalOrder
		this.loadHelper( clinicalOrderId );

	// split on a comma with no spaces
	var idList = imagingOrdersIds.split(',')

	// iterate over array of imagingOrders ids
	idList.forEach(function (id) {
		// read the ImagingOrder
		var imagingOrder = new ImagingOrderService(this.http).getImagingOrder(id);
		// add the ImagingOrder if not already assigned
		if ( this.clinicalOrder.imagingOrders.indexOf(imagingOrder) == -1 )
		this.clinicalOrder.imagingOrders.push(imagingOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more imagingOrdersIds as a ImagingOrders
	// from a ClinicalOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeImagingOrders( clinicalOrderId, imagingOrdersIds ): Observable<any> {

		// get the ClinicalOrder
		this.loadHelper( clinicalOrderId );


	// split on a comma with no spaces
	var idList 					= imagingOrdersIds.split(',');
	var imagingOrders 	= this.clinicalOrder.imagingOrders;

	if ( imagingOrders != null && imagingOrdersIds != null ) {

		// iterate over array of imagingOrders ids
		imagingOrders.forEach(function (obj) {
			if ( imagingOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the ImagingOrder
				this.clinicalOrder.imagingOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more procedureOrdersIds as a ProcedureOrders
	// to a ClinicalOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcedureOrders( clinicalOrderId, procedureOrdersIds ): Observable<any> {

		// get the ClinicalOrder
		this.loadHelper( clinicalOrderId );

	// split on a comma with no spaces
	var idList = procedureOrdersIds.split(',')

	// iterate over array of procedureOrders ids
	idList.forEach(function (id) {
		// read the ProcedureOrder
		var procedureOrder = new ProcedureOrderService(this.http).getProcedureOrder(id);
		// add the ProcedureOrder if not already assigned
		if ( this.clinicalOrder.procedureOrders.indexOf(procedureOrder) == -1 )
		this.clinicalOrder.procedureOrders.push(procedureOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more procedureOrdersIds as a ProcedureOrders
	// from a ClinicalOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcedureOrders( clinicalOrderId, procedureOrdersIds ): Observable<any> {

		// get the ClinicalOrder
		this.loadHelper( clinicalOrderId );


	// split on a comma with no spaces
	var idList 					= procedureOrdersIds.split(',');
	var procedureOrders 	= this.clinicalOrder.procedureOrders;

	if ( procedureOrders != null && procedureOrdersIds != null ) {

		// iterate over array of procedureOrders ids
		procedureOrders.forEach(function (obj) {
			if ( procedureOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the ProcedureOrder
				this.clinicalOrder.procedureOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more authorizationsIds as a Authorizations
	// to a ClinicalOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAuthorizations( clinicalOrderId, authorizationsIds ): Observable<any> {

		// get the ClinicalOrder
		this.loadHelper( clinicalOrderId );

	// split on a comma with no spaces
	var idList = authorizationsIds.split(',')

	// iterate over array of authorizations ids
	idList.forEach(function (id) {
		// read the Authorization
		var authorization = new AuthorizationService(this.http).getAuthorization(id);
		// add the Authorization if not already assigned
		if ( this.clinicalOrder.authorizations.indexOf(authorization) == -1 )
		this.clinicalOrder.authorizations.push(authorization);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more authorizationsIds as a Authorizations
	// from a ClinicalOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAuthorizations( clinicalOrderId, authorizationsIds ): Observable<any> {

		// get the ClinicalOrder
		this.loadHelper( clinicalOrderId );


	// split on a comma with no spaces
	var idList 					= authorizationsIds.split(',');
	var authorizations 	= this.clinicalOrder.authorizations;

	if ( authorizations != null && authorizationsIds != null ) {

		// iterate over array of authorizations ids
		authorizations.forEach(function (obj) {
			if ( authorizationsIds.indexOf(obj._id) > -1 ) {
				// remove the Authorization
				this.clinicalOrder.authorizations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ClinicalOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ClinicalOrder/update/' + this.clinicalOrder;

	return  this.http.post(uri_, this.clinicalOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a ClinicalOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getClinicalOrder(id)
			.subscribe((res : ClinicalOrder) => {
				this.clinicalOrder = res;
			});
	}
}