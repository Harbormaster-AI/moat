import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ServiceBulletin} from '../models/ServiceBulletin';
import {MaintenanceWorkOrderService} from '../services/MaintenanceWorkOrder.service';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ServiceBulletinService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	serviceBulletin : ServiceBulletin;

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
	// add a ServiceBulletin
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addServiceBulletin(bulletinNumber, WorkOrders, Variants, Category) : Observable<any> {
		const uri_ = this.apiUrl + '/ServiceBulletin/create';
		const obj = {
			      		bulletinNumber: bulletinNumber,
      		WorkOrders: WorkOrders != null && WorkOrders.length > 0 ? WorkOrders : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
			Category: Category
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ServiceBulletin
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateServiceBulletin(bulletinNumber, WorkOrders, Variants, Category, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ServiceBulletin/update/' + id;
		const obj = {
				      		bulletinNumber: bulletinNumber,
      		WorkOrders: WorkOrders != null && WorkOrders.length > 0 ? WorkOrders : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
			Category: Category
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ServiceBulletin
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteServiceBulletin(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ServiceBulletin/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ServiceBulletin
	// returns the results untouched as an Observable ServiceBulletin
	// ServiceBulletin model
	// delegates via URI
	//********************************************************************
	getServiceBulletin(id) : Observable<ServiceBulletin> {
		const uri_ = this.apiUrl + '/ServiceBulletin/load/' + id;

		return this.http.get<ServiceBulletin>(uri_);
	}
	
	//********************************************************************
	// gets all ServiceBulletin
	// returns the results untouched as JSON representation of an
	// Observable array of ServiceBulletin models
	// delegates via URI
	//********************************************************************
	getServiceBulletins() : Observable<ServiceBulletin[]> {
		const uri_ = this.apiUrl + '/ServiceBulletin/';

		return this
			.http.get<ServiceBulletin[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more workOrdersIds as a WorkOrders
	// to a ServiceBulletin
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWorkOrders( serviceBulletinId, workOrdersIds ): Observable<any> {

		// get the ServiceBulletin
		this.loadHelper( serviceBulletinId );

	// split on a comma with no spaces
	var idList = workOrdersIds.split(',')

	// iterate over array of workOrders ids
	idList.forEach(function (id) {
		// read the MaintenanceWorkOrder
		var maintenanceWorkOrder = new MaintenanceWorkOrderService(this.http).getMaintenanceWorkOrder(id);
		// add the MaintenanceWorkOrder if not already assigned
		if ( this.serviceBulletin.workOrders.indexOf(maintenanceWorkOrder) == -1 )
		this.serviceBulletin.workOrders.push(maintenanceWorkOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more workOrdersIds as a WorkOrders
	// from a ServiceBulletin
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWorkOrders( serviceBulletinId, workOrdersIds ): Observable<any> {

		// get the ServiceBulletin
		this.loadHelper( serviceBulletinId );


	// split on a comma with no spaces
	var idList 					= workOrdersIds.split(',');
	var workOrders 	= this.serviceBulletin.workOrders;

	if ( workOrders != null && workOrdersIds != null ) {

		// iterate over array of workOrders ids
		workOrders.forEach(function (obj) {
			if ( workOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the MaintenanceWorkOrder
				this.serviceBulletin.workOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more variantsIds as a Variants
	// to a ServiceBulletin
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVariants( serviceBulletinId, variantsIds ): Observable<any> {

		// get the ServiceBulletin
		this.loadHelper( serviceBulletinId );

	// split on a comma with no spaces
	var idList = variantsIds.split(',')

	// iterate over array of variants ids
	idList.forEach(function (id) {
		// read the AircraftVariant
		var aircraftVariant = new AircraftVariantService(this.http).getAircraftVariant(id);
		// add the AircraftVariant if not already assigned
		if ( this.serviceBulletin.variants.indexOf(aircraftVariant) == -1 )
		this.serviceBulletin.variants.push(aircraftVariant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more variantsIds as a Variants
	// from a ServiceBulletin
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVariants( serviceBulletinId, variantsIds ): Observable<any> {

		// get the ServiceBulletin
		this.loadHelper( serviceBulletinId );


	// split on a comma with no spaces
	var idList 					= variantsIds.split(',');
	var variants 	= this.serviceBulletin.variants;

	if ( variants != null && variantsIds != null ) {

		// iterate over array of variants ids
		variants.forEach(function (obj) {
			if ( variantsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftVariant
				this.serviceBulletin.variants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ServiceBulletin
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ServiceBulletin/update/' + this.serviceBulletin;

	return  this.http.post(uri_, this.serviceBulletin );
}

	//********************************************************************
	// loadHelper - internal helper to load a ServiceBulletin
	//********************************************************************	
	loadHelper( id ) {
		this.getServiceBulletin(id)
			.subscribe((res : ServiceBulletin) => {
				this.serviceBulletin = res;
			});
	}
}