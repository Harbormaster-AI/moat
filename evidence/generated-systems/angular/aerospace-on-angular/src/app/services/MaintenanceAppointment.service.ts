import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {MaintenanceAppointment} from '../models/MaintenanceAppointment';
import {AircraftService} from '../services/Aircraft.service';
import {MROFacilityService} from '../services/MROFacility.service';
import {MaintenanceWorkOrderService} from '../services/MaintenanceWorkOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MaintenanceAppointmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	maintenanceAppointment : MaintenanceAppointment;

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
	// add a MaintenanceAppointment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMaintenanceAppointment(appointmentDate, Aircraft, MroFacility, WorkOrder, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/MaintenanceAppointment/create';
		const obj = {
			      		appointmentDate: appointmentDate,
      		Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null,
      		MroFacility: MroFacility != null && MroFacility.length > 0 ? MroFacility : null,
      		WorkOrder: WorkOrder != null && WorkOrder.length > 0 ? WorkOrder : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a MaintenanceAppointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMaintenanceAppointment(appointmentDate, Aircraft, MroFacility, WorkOrder, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/MaintenanceAppointment/update/' + id;
		const obj = {
				      		appointmentDate: appointmentDate,
      		Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null,
      		MroFacility: MroFacility != null && MroFacility.length > 0 ? MroFacility : null,
      		WorkOrder: WorkOrder != null && WorkOrder.length > 0 ? WorkOrder : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a MaintenanceAppointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMaintenanceAppointment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/MaintenanceAppointment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a MaintenanceAppointment
	// returns the results untouched as an Observable MaintenanceAppointment
	// MaintenanceAppointment model
	// delegates via URI
	//********************************************************************
	getMaintenanceAppointment(id) : Observable<MaintenanceAppointment> {
		const uri_ = this.apiUrl + '/MaintenanceAppointment/load/' + id;

		return this.http.get<MaintenanceAppointment>(uri_);
	}
	
	//********************************************************************
	// gets all MaintenanceAppointment
	// returns the results untouched as JSON representation of an
	// Observable array of MaintenanceAppointment models
	// delegates via URI
	//********************************************************************
	getMaintenanceAppointments() : Observable<MaintenanceAppointment[]> {
		const uri_ = this.apiUrl + '/MaintenanceAppointment/';

		return this
			.http.get<MaintenanceAppointment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Aircraft on a MaintenanceAppointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAircraft( maintenanceAppointmentId, _aircraftId ): Observable<any> {

		// get the MaintenanceAppointment from storage
		this.loadHelper( maintenanceAppointmentId );

	// get the Aircraft from storage
	var tmp 	= new AircraftService(this.http).getAircraft(_aircraftId);

	// assign the Aircraft
	this.maintenanceAppointment.aircraft = tmp;

	// save the MaintenanceAppointment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Aircraft on a MaintenanceAppointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAircraft( maintenanceAppointmentId ): Observable<any> {

		// get the MaintenanceAppointment from storage
		this.loadHelper( maintenanceAppointmentId );

	// assign Aircraft to null
	this.maintenanceAppointment.aircraft = null;

	// save the MaintenanceAppointment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a MroFacility on a MaintenanceAppointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMroFacility( maintenanceAppointmentId, _mroFacilityId ): Observable<any> {

		// get the MaintenanceAppointment from storage
		this.loadHelper( maintenanceAppointmentId );

	// get the MROFacility from storage
	var tmp 	= new MROFacilityService(this.http).getMROFacility(_mroFacilityId);

	// assign the MroFacility
	this.maintenanceAppointment.mroFacility = tmp;

	// save the MaintenanceAppointment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a MroFacility on a MaintenanceAppointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMroFacility( maintenanceAppointmentId ): Observable<any> {

		// get the MaintenanceAppointment from storage
		this.loadHelper( maintenanceAppointmentId );

	// assign MroFacility to null
	this.maintenanceAppointment.mroFacility = null;

	// save the MaintenanceAppointment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a WorkOrder on a MaintenanceAppointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkOrder( maintenanceAppointmentId, _workOrderId ): Observable<any> {

		// get the MaintenanceAppointment from storage
		this.loadHelper( maintenanceAppointmentId );

	// get the MaintenanceWorkOrder from storage
	var tmp 	= new MaintenanceWorkOrderService(this.http).getMaintenanceWorkOrder(_workOrderId);

	// assign the WorkOrder
	this.maintenanceAppointment.workOrder = tmp;

	// save the MaintenanceAppointment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkOrder on a MaintenanceAppointment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkOrder( maintenanceAppointmentId ): Observable<any> {

		// get the MaintenanceAppointment from storage
		this.loadHelper( maintenanceAppointmentId );

	// assign WorkOrder to null
	this.maintenanceAppointment.workOrder = null;

	// save the MaintenanceAppointment
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a MaintenanceAppointment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/MaintenanceAppointment/update/' + this.maintenanceAppointment;

	return  this.http.post(uri_, this.maintenanceAppointment );
}

	//********************************************************************
	// loadHelper - internal helper to load a MaintenanceAppointment
	//********************************************************************	
	loadHelper( id ) {
		this.getMaintenanceAppointment(id)
			.subscribe((res : MaintenanceAppointment) => {
				this.maintenanceAppointment = res;
			});
	}
}