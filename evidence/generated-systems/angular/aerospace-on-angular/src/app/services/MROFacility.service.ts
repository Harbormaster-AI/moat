import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {MROFacility} from '../models/MROFacility';
import {MaintenanceAppointmentService} from '../services/MaintenanceAppointment.service';
import {MaintenanceWorkOrderService} from '../services/MaintenanceWorkOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MROFacilityService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	mROFacility : MROFacility;

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
	// add a MROFacility
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMROFacility(name, approvalScope, address, Appointments, WorkOrders) : Observable<any> {
		const uri_ = this.apiUrl + '/MROFacility/create';
		const obj = {
			      		name: name,
      		approvalScope: approvalScope,
      		address: address,
      		Appointments: Appointments != null && Appointments.length > 0 ? Appointments : null,
			WorkOrders: WorkOrders != null && WorkOrders.length > 0 ? WorkOrders : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a MROFacility
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMROFacility(name, approvalScope, address, Appointments, WorkOrders, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/MROFacility/update/' + id;
		const obj = {
				      		name: name,
      		approvalScope: approvalScope,
      		address: address,
      		Appointments: Appointments != null && Appointments.length > 0 ? Appointments : null,
			WorkOrders: WorkOrders != null && WorkOrders.length > 0 ? WorkOrders : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a MROFacility
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMROFacility(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/MROFacility/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a MROFacility
	// returns the results untouched as an Observable MROFacility
	// MROFacility model
	// delegates via URI
	//********************************************************************
	getMROFacility(id) : Observable<MROFacility> {
		const uri_ = this.apiUrl + '/MROFacility/load/' + id;

		return this.http.get<MROFacility>(uri_);
	}
	
	//********************************************************************
	// gets all MROFacility
	// returns the results untouched as JSON representation of an
	// Observable array of MROFacility models
	// delegates via URI
	//********************************************************************
	getMROFacilitys() : Observable<MROFacility[]> {
		const uri_ = this.apiUrl + '/MROFacility/';

		return this
			.http.get<MROFacility[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more appointmentsIds as a Appointments
	// to a MROFacility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAppointments( mROFacilityId, appointmentsIds ): Observable<any> {

		// get the MROFacility
		this.loadHelper( mROFacilityId );

	// split on a comma with no spaces
	var idList = appointmentsIds.split(',')

	// iterate over array of appointments ids
	idList.forEach(function (id) {
		// read the MaintenanceAppointment
		var maintenanceAppointment = new MaintenanceAppointmentService(this.http).getMaintenanceAppointment(id);
		// add the MaintenanceAppointment if not already assigned
		if ( this.mROFacility.appointments.indexOf(maintenanceAppointment) == -1 )
		this.mROFacility.appointments.push(maintenanceAppointment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more appointmentsIds as a Appointments
	// from a MROFacility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAppointments( mROFacilityId, appointmentsIds ): Observable<any> {

		// get the MROFacility
		this.loadHelper( mROFacilityId );


	// split on a comma with no spaces
	var idList 					= appointmentsIds.split(',');
	var appointments 	= this.mROFacility.appointments;

	if ( appointments != null && appointmentsIds != null ) {

		// iterate over array of appointments ids
		appointments.forEach(function (obj) {
			if ( appointmentsIds.indexOf(obj._id) > -1 ) {
				// remove the MaintenanceAppointment
				this.mROFacility.appointments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more workOrdersIds as a WorkOrders
	// to a MROFacility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWorkOrders( mROFacilityId, workOrdersIds ): Observable<any> {

		// get the MROFacility
		this.loadHelper( mROFacilityId );

	// split on a comma with no spaces
	var idList = workOrdersIds.split(',')

	// iterate over array of workOrders ids
	idList.forEach(function (id) {
		// read the MaintenanceWorkOrder
		var maintenanceWorkOrder = new MaintenanceWorkOrderService(this.http).getMaintenanceWorkOrder(id);
		// add the MaintenanceWorkOrder if not already assigned
		if ( this.mROFacility.workOrders.indexOf(maintenanceWorkOrder) == -1 )
		this.mROFacility.workOrders.push(maintenanceWorkOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more workOrdersIds as a WorkOrders
	// from a MROFacility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWorkOrders( mROFacilityId, workOrdersIds ): Observable<any> {

		// get the MROFacility
		this.loadHelper( mROFacilityId );


	// split on a comma with no spaces
	var idList 					= workOrdersIds.split(',');
	var workOrders 	= this.mROFacility.workOrders;

	if ( workOrders != null && workOrdersIds != null ) {

		// iterate over array of workOrders ids
		workOrders.forEach(function (obj) {
			if ( workOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the MaintenanceWorkOrder
				this.mROFacility.workOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a MROFacility
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/MROFacility/update/' + this.mROFacility;

	return  this.http.post(uri_, this.mROFacility );
}

	//********************************************************************
	// loadHelper - internal helper to load a MROFacility
	//********************************************************************	
	loadHelper( id ) {
		this.getMROFacility(id)
			.subscribe((res : MROFacility) => {
				this.mROFacility = res;
			});
	}
}