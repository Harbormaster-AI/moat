import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SalesOrder} from '../models/SalesOrder';
import {CustomerService} from '../services/Customer.service';
import {PlantService} from '../services/Plant.service';
import {SalesOrderLineService} from '../services/SalesOrderLine.service';
import {WorkOrderService} from '../services/WorkOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SalesOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	salesOrder : SalesOrder;

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
	// add a SalesOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSalesOrder(orderNumber, orderDate, totalAmount, Customer, Plant, Lines, WorkOrders, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/SalesOrder/create';
		const obj = {
			      		orderNumber: orderNumber,
      		orderDate: orderDate,
      		totalAmount: totalAmount,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
      		WorkOrders: WorkOrders != null && WorkOrders.length > 0 ? WorkOrders : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SalesOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSalesOrder(orderNumber, orderDate, totalAmount, Customer, Plant, Lines, WorkOrders, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SalesOrder/update/' + id;
		const obj = {
				      		orderNumber: orderNumber,
      		orderDate: orderDate,
      		totalAmount: totalAmount,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
      		WorkOrders: WorkOrders != null && WorkOrders.length > 0 ? WorkOrders : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SalesOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSalesOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SalesOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SalesOrder
	// returns the results untouched as an Observable SalesOrder
	// SalesOrder model
	// delegates via URI
	//********************************************************************
	getSalesOrder(id) : Observable<SalesOrder> {
		const uri_ = this.apiUrl + '/SalesOrder/load/' + id;

		return this.http.get<SalesOrder>(uri_);
	}
	
	//********************************************************************
	// gets all SalesOrder
	// returns the results untouched as JSON representation of an
	// Observable array of SalesOrder models
	// delegates via URI
	//********************************************************************
	getSalesOrders() : Observable<SalesOrder[]> {
		const uri_ = this.apiUrl + '/SalesOrder/';

		return this
			.http.get<SalesOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a SalesOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( salesOrderId, _customerId ): Observable<any> {

		// get the SalesOrder from storage
		this.loadHelper( salesOrderId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.salesOrder.customer = tmp;

	// save the SalesOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a SalesOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( salesOrderId ): Observable<any> {

		// get the SalesOrder from storage
		this.loadHelper( salesOrderId );

	// assign Customer to null
	this.salesOrder.customer = null;

	// save the SalesOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Plant on a SalesOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( salesOrderId, _plantId ): Observable<any> {

		// get the SalesOrder from storage
		this.loadHelper( salesOrderId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.salesOrder.plant = tmp;

	// save the SalesOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a SalesOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( salesOrderId ): Observable<any> {

		// get the SalesOrder from storage
		this.loadHelper( salesOrderId );

	// assign Plant to null
	this.salesOrder.plant = null;

	// save the SalesOrder
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more linesIds as a Lines
	// to a SalesOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLines( salesOrderId, linesIds ): Observable<any> {

		// get the SalesOrder
		this.loadHelper( salesOrderId );

	// split on a comma with no spaces
	var idList = linesIds.split(',')

	// iterate over array of lines ids
	idList.forEach(function (id) {
		// read the SalesOrderLine
		var salesOrderLine = new SalesOrderLineService(this.http).getSalesOrderLine(id);
		// add the SalesOrderLine if not already assigned
		if ( this.salesOrder.lines.indexOf(salesOrderLine) == -1 )
		this.salesOrder.lines.push(salesOrderLine);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more linesIds as a Lines
	// from a SalesOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLines( salesOrderId, linesIds ): Observable<any> {

		// get the SalesOrder
		this.loadHelper( salesOrderId );


	// split on a comma with no spaces
	var idList 					= linesIds.split(',');
	var lines 	= this.salesOrder.lines;

	if ( lines != null && linesIds != null ) {

		// iterate over array of lines ids
		lines.forEach(function (obj) {
			if ( linesIds.indexOf(obj._id) > -1 ) {
				// remove the SalesOrderLine
				this.salesOrder.lines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more workOrdersIds as a WorkOrders
	// to a SalesOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWorkOrders( salesOrderId, workOrdersIds ): Observable<any> {

		// get the SalesOrder
		this.loadHelper( salesOrderId );

	// split on a comma with no spaces
	var idList = workOrdersIds.split(',')

	// iterate over array of workOrders ids
	idList.forEach(function (id) {
		// read the WorkOrder
		var workOrder = new WorkOrderService(this.http).getWorkOrder(id);
		// add the WorkOrder if not already assigned
		if ( this.salesOrder.workOrders.indexOf(workOrder) == -1 )
		this.salesOrder.workOrders.push(workOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more workOrdersIds as a WorkOrders
	// from a SalesOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWorkOrders( salesOrderId, workOrdersIds ): Observable<any> {

		// get the SalesOrder
		this.loadHelper( salesOrderId );


	// split on a comma with no spaces
	var idList 					= workOrdersIds.split(',');
	var workOrders 	= this.salesOrder.workOrders;

	if ( workOrders != null && workOrdersIds != null ) {

		// iterate over array of workOrders ids
		workOrders.forEach(function (obj) {
			if ( workOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the WorkOrder
				this.salesOrder.workOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a SalesOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SalesOrder/update/' + this.salesOrder;

	return  this.http.post(uri_, this.salesOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a SalesOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getSalesOrder(id)
			.subscribe((res : SalesOrder) => {
				this.salesOrder = res;
			});
	}
}