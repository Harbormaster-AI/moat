import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ProductionSchedule} from '../models/ProductionSchedule';
import {PlantService} from '../services/Plant.service';
import {WorkOrderService} from '../services/WorkOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ProductionScheduleService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	productionSchedule : ProductionSchedule;

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
	// add a ProductionSchedule
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addProductionSchedule(scheduleNumber, horizonStart, horizonEnd, Plant, WorkOrders, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/ProductionSchedule/create';
		const obj = {
			      		scheduleNumber: scheduleNumber,
      		horizonStart: horizonStart,
      		horizonEnd: horizonEnd,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		WorkOrders: WorkOrders != null && WorkOrders.length > 0 ? WorkOrders : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ProductionSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateProductionSchedule(scheduleNumber, horizonStart, horizonEnd, Plant, WorkOrders, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ProductionSchedule/update/' + id;
		const obj = {
				      		scheduleNumber: scheduleNumber,
      		horizonStart: horizonStart,
      		horizonEnd: horizonEnd,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		WorkOrders: WorkOrders != null && WorkOrders.length > 0 ? WorkOrders : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ProductionSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteProductionSchedule(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ProductionSchedule/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ProductionSchedule
	// returns the results untouched as an Observable ProductionSchedule
	// ProductionSchedule model
	// delegates via URI
	//********************************************************************
	getProductionSchedule(id) : Observable<ProductionSchedule> {
		const uri_ = this.apiUrl + '/ProductionSchedule/load/' + id;

		return this.http.get<ProductionSchedule>(uri_);
	}
	
	//********************************************************************
	// gets all ProductionSchedule
	// returns the results untouched as JSON representation of an
	// Observable array of ProductionSchedule models
	// delegates via URI
	//********************************************************************
	getProductionSchedules() : Observable<ProductionSchedule[]> {
		const uri_ = this.apiUrl + '/ProductionSchedule/';

		return this
			.http.get<ProductionSchedule[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Plant on a ProductionSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( productionScheduleId, _plantId ): Observable<any> {

		// get the ProductionSchedule from storage
		this.loadHelper( productionScheduleId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.productionSchedule.plant = tmp;

	// save the ProductionSchedule
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a ProductionSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( productionScheduleId ): Observable<any> {

		// get the ProductionSchedule from storage
		this.loadHelper( productionScheduleId );

	// assign Plant to null
	this.productionSchedule.plant = null;

	// save the ProductionSchedule
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more workOrdersIds as a WorkOrders
	// to a ProductionSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWorkOrders( productionScheduleId, workOrdersIds ): Observable<any> {

		// get the ProductionSchedule
		this.loadHelper( productionScheduleId );

	// split on a comma with no spaces
	var idList = workOrdersIds.split(',')

	// iterate over array of workOrders ids
	idList.forEach(function (id) {
		// read the WorkOrder
		var workOrder = new WorkOrderService(this.http).getWorkOrder(id);
		// add the WorkOrder if not already assigned
		if ( this.productionSchedule.workOrders.indexOf(workOrder) == -1 )
		this.productionSchedule.workOrders.push(workOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more workOrdersIds as a WorkOrders
	// from a ProductionSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWorkOrders( productionScheduleId, workOrdersIds ): Observable<any> {

		// get the ProductionSchedule
		this.loadHelper( productionScheduleId );


	// split on a comma with no spaces
	var idList 					= workOrdersIds.split(',');
	var workOrders 	= this.productionSchedule.workOrders;

	if ( workOrders != null && workOrdersIds != null ) {

		// iterate over array of workOrders ids
		workOrders.forEach(function (obj) {
			if ( workOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the WorkOrder
				this.productionSchedule.workOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ProductionSchedule
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ProductionSchedule/update/' + this.productionSchedule;

	return  this.http.post(uri_, this.productionSchedule );
}

	//********************************************************************
	// loadHelper - internal helper to load a ProductionSchedule
	//********************************************************************	
	loadHelper( id ) {
		this.getProductionSchedule(id)
			.subscribe((res : ProductionSchedule) => {
				this.productionSchedule = res;
			});
	}
}