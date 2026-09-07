import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PlannedOrder} from '../models/PlannedOrder';
import {MRPRunService} from '../services/MRPRun.service';
import {ItemService} from '../services/Item.service';
import {PlantService} from '../services/Plant.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PlannedOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	plannedOrder : PlannedOrder;

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
	// add a PlannedOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPlannedOrder(plannedOrderNumber, quantity, dueDate, MrpRun, Item, Plant, OrderType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/PlannedOrder/create';
		const obj = {
			      		plannedOrderNumber: plannedOrderNumber,
      		quantity: quantity,
      		dueDate: dueDate,
      		MrpRun: MrpRun != null && MrpRun.length > 0 ? MrpRun : null,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		OrderType: OrderType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PlannedOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePlannedOrder(plannedOrderNumber, quantity, dueDate, MrpRun, Item, Plant, OrderType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PlannedOrder/update/' + id;
		const obj = {
				      		plannedOrderNumber: plannedOrderNumber,
      		quantity: quantity,
      		dueDate: dueDate,
      		MrpRun: MrpRun != null && MrpRun.length > 0 ? MrpRun : null,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		OrderType: OrderType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PlannedOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePlannedOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PlannedOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PlannedOrder
	// returns the results untouched as an Observable PlannedOrder
	// PlannedOrder model
	// delegates via URI
	//********************************************************************
	getPlannedOrder(id) : Observable<PlannedOrder> {
		const uri_ = this.apiUrl + '/PlannedOrder/load/' + id;

		return this.http.get<PlannedOrder>(uri_);
	}
	
	//********************************************************************
	// gets all PlannedOrder
	// returns the results untouched as JSON representation of an
	// Observable array of PlannedOrder models
	// delegates via URI
	//********************************************************************
	getPlannedOrders() : Observable<PlannedOrder[]> {
		const uri_ = this.apiUrl + '/PlannedOrder/';

		return this
			.http.get<PlannedOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a MrpRun on a PlannedOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMrpRun( plannedOrderId, _mrpRunId ): Observable<any> {

		// get the PlannedOrder from storage
		this.loadHelper( plannedOrderId );

	// get the MRPRun from storage
	var tmp 	= new MRPRunService(this.http).getMRPRun(_mrpRunId);

	// assign the MrpRun
	this.plannedOrder.mrpRun = tmp;

	// save the PlannedOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a MrpRun on a PlannedOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMrpRun( plannedOrderId ): Observable<any> {

		// get the PlannedOrder from storage
		this.loadHelper( plannedOrderId );

	// assign MrpRun to null
	this.plannedOrder.mrpRun = null;

	// save the PlannedOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Item on a PlannedOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( plannedOrderId, _itemId ): Observable<any> {

		// get the PlannedOrder from storage
		this.loadHelper( plannedOrderId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.plannedOrder.item = tmp;

	// save the PlannedOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a PlannedOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( plannedOrderId ): Observable<any> {

		// get the PlannedOrder from storage
		this.loadHelper( plannedOrderId );

	// assign Item to null
	this.plannedOrder.item = null;

	// save the PlannedOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Plant on a PlannedOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( plannedOrderId, _plantId ): Observable<any> {

		// get the PlannedOrder from storage
		this.loadHelper( plannedOrderId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.plannedOrder.plant = tmp;

	// save the PlannedOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a PlannedOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( plannedOrderId ): Observable<any> {

		// get the PlannedOrder from storage
		this.loadHelper( plannedOrderId );

	// assign Plant to null
	this.plannedOrder.plant = null;

	// save the PlannedOrder
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a PlannedOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PlannedOrder/update/' + this.plannedOrder;

	return  this.http.post(uri_, this.plannedOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a PlannedOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getPlannedOrder(id)
			.subscribe((res : PlannedOrder) => {
				this.plannedOrder = res;
			});
	}
}