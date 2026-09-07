import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {WorkOrder} from '../models/WorkOrder';
import {ItemService} from '../services/Item.service';
import {PlantService} from '../services/Plant.service';
import {RoutingService} from '../services/Routing.service';
import {BOMService} from '../services/BOM.service';
import {ProductionScheduleService} from '../services/ProductionSchedule.service';
import {SalesOrderService} from '../services/SalesOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class WorkOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	workOrder : WorkOrder;

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
	// add a WorkOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addWorkOrder(workOrderNumber, plannedStart, plannedEnd, quantity, priority, Item, Plant, Routing, Bom, ProductionSchedule, SalesOrder, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/WorkOrder/create';
		const obj = {
			      		workOrderNumber: workOrderNumber,
      		plannedStart: plannedStart,
      		plannedEnd: plannedEnd,
      		quantity: quantity,
      		priority: priority,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		Routing: Routing != null && Routing.length > 0 ? Routing : null,
      		Bom: Bom != null && Bom.length > 0 ? Bom : null,
      		ProductionSchedule: ProductionSchedule != null && ProductionSchedule.length > 0 ? ProductionSchedule : null,
      		SalesOrder: SalesOrder != null && SalesOrder.length > 0 ? SalesOrder : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateWorkOrder(workOrderNumber, plannedStart, plannedEnd, quantity, priority, Item, Plant, Routing, Bom, ProductionSchedule, SalesOrder, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/WorkOrder/update/' + id;
		const obj = {
				      		workOrderNumber: workOrderNumber,
      		plannedStart: plannedStart,
      		plannedEnd: plannedEnd,
      		quantity: quantity,
      		priority: priority,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		Routing: Routing != null && Routing.length > 0 ? Routing : null,
      		Bom: Bom != null && Bom.length > 0 ? Bom : null,
      		ProductionSchedule: ProductionSchedule != null && ProductionSchedule.length > 0 ? ProductionSchedule : null,
      		SalesOrder: SalesOrder != null && SalesOrder.length > 0 ? SalesOrder : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteWorkOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/WorkOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a WorkOrder
	// returns the results untouched as an Observable WorkOrder
	// WorkOrder model
	// delegates via URI
	//********************************************************************
	getWorkOrder(id) : Observable<WorkOrder> {
		const uri_ = this.apiUrl + '/WorkOrder/load/' + id;

		return this.http.get<WorkOrder>(uri_);
	}
	
	//********************************************************************
	// gets all WorkOrder
	// returns the results untouched as JSON representation of an
	// Observable array of WorkOrder models
	// delegates via URI
	//********************************************************************
	getWorkOrders() : Observable<WorkOrder[]> {
		const uri_ = this.apiUrl + '/WorkOrder/';

		return this
			.http.get<WorkOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Item on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( workOrderId, _itemId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.workOrder.item = tmp;

	// save the WorkOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( workOrderId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// assign Item to null
	this.workOrder.item = null;

	// save the WorkOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Plant on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( workOrderId, _plantId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.workOrder.plant = tmp;

	// save the WorkOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( workOrderId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// assign Plant to null
	this.workOrder.plant = null;

	// save the WorkOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Routing on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRouting( workOrderId, _routingId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// get the Routing from storage
	var tmp 	= new RoutingService(this.http).getRouting(_routingId);

	// assign the Routing
	this.workOrder.routing = tmp;

	// save the WorkOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Routing on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRouting( workOrderId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// assign Routing to null
	this.workOrder.routing = null;

	// save the WorkOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Bom on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBom( workOrderId, _bomId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// get the BOM from storage
	var tmp 	= new BOMService(this.http).getBOM(_bomId);

	// assign the Bom
	this.workOrder.bom = tmp;

	// save the WorkOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Bom on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBom( workOrderId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// assign Bom to null
	this.workOrder.bom = null;

	// save the WorkOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ProductionSchedule on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProductionSchedule( workOrderId, _productionScheduleId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// get the ProductionSchedule from storage
	var tmp 	= new ProductionScheduleService(this.http).getProductionSchedule(_productionScheduleId);

	// assign the ProductionSchedule
	this.workOrder.productionSchedule = tmp;

	// save the WorkOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ProductionSchedule on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProductionSchedule( workOrderId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// assign ProductionSchedule to null
	this.workOrder.productionSchedule = null;

	// save the WorkOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a SalesOrder on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSalesOrder( workOrderId, _salesOrderId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// get the SalesOrder from storage
	var tmp 	= new SalesOrderService(this.http).getSalesOrder(_salesOrderId);

	// assign the SalesOrder
	this.workOrder.salesOrder = tmp;

	// save the WorkOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SalesOrder on a WorkOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSalesOrder( workOrderId ): Observable<any> {

		// get the WorkOrder from storage
		this.loadHelper( workOrderId );

	// assign SalesOrder to null
	this.workOrder.salesOrder = null;

	// save the WorkOrder
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a WorkOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/WorkOrder/update/' + this.workOrder;

	return  this.http.post(uri_, this.workOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a WorkOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getWorkOrder(id)
			.subscribe((res : WorkOrder) => {
				this.workOrder = res;
			});
	}
}