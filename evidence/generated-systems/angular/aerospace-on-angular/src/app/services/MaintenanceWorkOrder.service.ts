import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {MaintenanceWorkOrder} from '../models/MaintenanceWorkOrder';
import {AircraftService} from '../services/Aircraft.service';
import {AirworthinessDirectiveService} from '../services/AirworthinessDirective.service';
import {ServiceBulletinService} from '../services/ServiceBulletin.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MaintenanceWorkOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	maintenanceWorkOrder : MaintenanceWorkOrder;

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
	// add a MaintenanceWorkOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMaintenanceWorkOrder(workOrderNumber, Aircraft, AirworthinessDirective, ServiceBulletin, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/MaintenanceWorkOrder/create';
		const obj = {
			      		workOrderNumber: workOrderNumber,
      		Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null,
      		AirworthinessDirective: AirworthinessDirective != null && AirworthinessDirective.length > 0 ? AirworthinessDirective : null,
      		ServiceBulletin: ServiceBulletin != null && ServiceBulletin.length > 0 ? ServiceBulletin : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a MaintenanceWorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMaintenanceWorkOrder(workOrderNumber, Aircraft, AirworthinessDirective, ServiceBulletin, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/MaintenanceWorkOrder/update/' + id;
		const obj = {
				      		workOrderNumber: workOrderNumber,
      		Aircraft: Aircraft != null && Aircraft.length > 0 ? Aircraft : null,
      		AirworthinessDirective: AirworthinessDirective != null && AirworthinessDirective.length > 0 ? AirworthinessDirective : null,
      		ServiceBulletin: ServiceBulletin != null && ServiceBulletin.length > 0 ? ServiceBulletin : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a MaintenanceWorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMaintenanceWorkOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/MaintenanceWorkOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a MaintenanceWorkOrder
	// returns the results untouched as an Observable MaintenanceWorkOrder
	// MaintenanceWorkOrder model
	// delegates via URI
	//********************************************************************
	getMaintenanceWorkOrder(id) : Observable<MaintenanceWorkOrder> {
		const uri_ = this.apiUrl + '/MaintenanceWorkOrder/load/' + id;

		return this.http.get<MaintenanceWorkOrder>(uri_);
	}
	
	//********************************************************************
	// gets all MaintenanceWorkOrder
	// returns the results untouched as JSON representation of an
	// Observable array of MaintenanceWorkOrder models
	// delegates via URI
	//********************************************************************
	getMaintenanceWorkOrders() : Observable<MaintenanceWorkOrder[]> {
		const uri_ = this.apiUrl + '/MaintenanceWorkOrder/';

		return this
			.http.get<MaintenanceWorkOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Aircraft on a MaintenanceWorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAircraft( maintenanceWorkOrderId, _aircraftId ): Observable<any> {

		// get the MaintenanceWorkOrder from storage
		this.loadHelper( maintenanceWorkOrderId );

	// get the Aircraft from storage
	var tmp 	= new AircraftService(this.http).getAircraft(_aircraftId);

	// assign the Aircraft
	this.maintenanceWorkOrder.aircraft = tmp;

	// save the MaintenanceWorkOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Aircraft on a MaintenanceWorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAircraft( maintenanceWorkOrderId ): Observable<any> {

		// get the MaintenanceWorkOrder from storage
		this.loadHelper( maintenanceWorkOrderId );

	// assign Aircraft to null
	this.maintenanceWorkOrder.aircraft = null;

	// save the MaintenanceWorkOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a AirworthinessDirective on a MaintenanceWorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAirworthinessDirective( maintenanceWorkOrderId, _airworthinessDirectiveId ): Observable<any> {

		// get the MaintenanceWorkOrder from storage
		this.loadHelper( maintenanceWorkOrderId );

	// get the AirworthinessDirective from storage
	var tmp 	= new AirworthinessDirectiveService(this.http).getAirworthinessDirective(_airworthinessDirectiveId);

	// assign the AirworthinessDirective
	this.maintenanceWorkOrder.airworthinessDirective = tmp;

	// save the MaintenanceWorkOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AirworthinessDirective on a MaintenanceWorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAirworthinessDirective( maintenanceWorkOrderId ): Observable<any> {

		// get the MaintenanceWorkOrder from storage
		this.loadHelper( maintenanceWorkOrderId );

	// assign AirworthinessDirective to null
	this.maintenanceWorkOrder.airworthinessDirective = null;

	// save the MaintenanceWorkOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ServiceBulletin on a MaintenanceWorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignServiceBulletin( maintenanceWorkOrderId, _serviceBulletinId ): Observable<any> {

		// get the MaintenanceWorkOrder from storage
		this.loadHelper( maintenanceWorkOrderId );

	// get the ServiceBulletin from storage
	var tmp 	= new ServiceBulletinService(this.http).getServiceBulletin(_serviceBulletinId);

	// assign the ServiceBulletin
	this.maintenanceWorkOrder.serviceBulletin = tmp;

	// save the MaintenanceWorkOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ServiceBulletin on a MaintenanceWorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignServiceBulletin( maintenanceWorkOrderId ): Observable<any> {

		// get the MaintenanceWorkOrder from storage
		this.loadHelper( maintenanceWorkOrderId );

	// assign ServiceBulletin to null
	this.maintenanceWorkOrder.serviceBulletin = null;

	// save the MaintenanceWorkOrder
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a MaintenanceWorkOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/MaintenanceWorkOrder/update/' + this.maintenanceWorkOrder;

	return  this.http.post(uri_, this.maintenanceWorkOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a MaintenanceWorkOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getMaintenanceWorkOrder(id)
			.subscribe((res : MaintenanceWorkOrder) => {
				this.maintenanceWorkOrder = res;
			});
	}
}