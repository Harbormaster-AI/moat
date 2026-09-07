import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AirworthinessDirective} from '../models/AirworthinessDirective';
import {MaintenanceWorkOrderService} from '../services/MaintenanceWorkOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AirworthinessDirectiveService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	airworthinessDirective : AirworthinessDirective;

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
	// add a AirworthinessDirective
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAirworthinessDirective(directiveNumber, title, WorkOrders) : Observable<any> {
		const uri_ = this.apiUrl + '/AirworthinessDirective/create';
		const obj = {
			      		directiveNumber: directiveNumber,
      		title: title,
			WorkOrders: WorkOrders != null && WorkOrders.length > 0 ? WorkOrders : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AirworthinessDirective
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAirworthinessDirective(directiveNumber, title, WorkOrders, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AirworthinessDirective/update/' + id;
		const obj = {
				      		directiveNumber: directiveNumber,
      		title: title,
			WorkOrders: WorkOrders != null && WorkOrders.length > 0 ? WorkOrders : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AirworthinessDirective
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAirworthinessDirective(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AirworthinessDirective/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AirworthinessDirective
	// returns the results untouched as an Observable AirworthinessDirective
	// AirworthinessDirective model
	// delegates via URI
	//********************************************************************
	getAirworthinessDirective(id) : Observable<AirworthinessDirective> {
		const uri_ = this.apiUrl + '/AirworthinessDirective/load/' + id;

		return this.http.get<AirworthinessDirective>(uri_);
	}
	
	//********************************************************************
	// gets all AirworthinessDirective
	// returns the results untouched as JSON representation of an
	// Observable array of AirworthinessDirective models
	// delegates via URI
	//********************************************************************
	getAirworthinessDirectives() : Observable<AirworthinessDirective[]> {
		const uri_ = this.apiUrl + '/AirworthinessDirective/';

		return this
			.http.get<AirworthinessDirective[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more workOrdersIds as a WorkOrders
	// to a AirworthinessDirective
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWorkOrders( airworthinessDirectiveId, workOrdersIds ): Observable<any> {

		// get the AirworthinessDirective
		this.loadHelper( airworthinessDirectiveId );

	// split on a comma with no spaces
	var idList = workOrdersIds.split(',')

	// iterate over array of workOrders ids
	idList.forEach(function (id) {
		// read the MaintenanceWorkOrder
		var maintenanceWorkOrder = new MaintenanceWorkOrderService(this.http).getMaintenanceWorkOrder(id);
		// add the MaintenanceWorkOrder if not already assigned
		if ( this.airworthinessDirective.workOrders.indexOf(maintenanceWorkOrder) == -1 )
		this.airworthinessDirective.workOrders.push(maintenanceWorkOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more workOrdersIds as a WorkOrders
	// from a AirworthinessDirective
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWorkOrders( airworthinessDirectiveId, workOrdersIds ): Observable<any> {

		// get the AirworthinessDirective
		this.loadHelper( airworthinessDirectiveId );


	// split on a comma with no spaces
	var idList 					= workOrdersIds.split(',');
	var workOrders 	= this.airworthinessDirective.workOrders;

	if ( workOrders != null && workOrdersIds != null ) {

		// iterate over array of workOrders ids
		workOrders.forEach(function (obj) {
			if ( workOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the MaintenanceWorkOrder
				this.airworthinessDirective.workOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AirworthinessDirective
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AirworthinessDirective/update/' + this.airworthinessDirective;

	return  this.http.post(uri_, this.airworthinessDirective );
}

	//********************************************************************
	// loadHelper - internal helper to load a AirworthinessDirective
	//********************************************************************	
	loadHelper( id ) {
		this.getAirworthinessDirective(id)
			.subscribe((res : AirworthinessDirective) => {
				this.airworthinessDirective = res;
			});
	}
}