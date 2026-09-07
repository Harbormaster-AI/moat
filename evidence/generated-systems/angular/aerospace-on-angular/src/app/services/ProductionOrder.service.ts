import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ProductionOrder} from '../models/ProductionOrder';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import {PlantService} from '../services/Plant.service';
import {AircraftOrderService} from '../services/AircraftOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ProductionOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	productionOrder : ProductionOrder;

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
	// add a ProductionOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addProductionOrder(orderNumber, Variant, Plant, AircraftOrder, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/ProductionOrder/create';
		const obj = {
			      		orderNumber: orderNumber,
      		Variant: Variant != null && Variant.length > 0 ? Variant : null,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		AircraftOrder: AircraftOrder != null && AircraftOrder.length > 0 ? AircraftOrder : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ProductionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateProductionOrder(orderNumber, Variant, Plant, AircraftOrder, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ProductionOrder/update/' + id;
		const obj = {
				      		orderNumber: orderNumber,
      		Variant: Variant != null && Variant.length > 0 ? Variant : null,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		AircraftOrder: AircraftOrder != null && AircraftOrder.length > 0 ? AircraftOrder : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ProductionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteProductionOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ProductionOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ProductionOrder
	// returns the results untouched as an Observable ProductionOrder
	// ProductionOrder model
	// delegates via URI
	//********************************************************************
	getProductionOrder(id) : Observable<ProductionOrder> {
		const uri_ = this.apiUrl + '/ProductionOrder/load/' + id;

		return this.http.get<ProductionOrder>(uri_);
	}
	
	//********************************************************************
	// gets all ProductionOrder
	// returns the results untouched as JSON representation of an
	// Observable array of ProductionOrder models
	// delegates via URI
	//********************************************************************
	getProductionOrders() : Observable<ProductionOrder[]> {
		const uri_ = this.apiUrl + '/ProductionOrder/';

		return this
			.http.get<ProductionOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Variant on a ProductionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignVariant( productionOrderId, _variantId ): Observable<any> {

		// get the ProductionOrder from storage
		this.loadHelper( productionOrderId );

	// get the AircraftVariant from storage
	var tmp 	= new AircraftVariantService(this.http).getAircraftVariant(_variantId);

	// assign the Variant
	this.productionOrder.variant = tmp;

	// save the ProductionOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Variant on a ProductionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignVariant( productionOrderId ): Observable<any> {

		// get the ProductionOrder from storage
		this.loadHelper( productionOrderId );

	// assign Variant to null
	this.productionOrder.variant = null;

	// save the ProductionOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Plant on a ProductionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( productionOrderId, _plantId ): Observable<any> {

		// get the ProductionOrder from storage
		this.loadHelper( productionOrderId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.productionOrder.plant = tmp;

	// save the ProductionOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a ProductionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( productionOrderId ): Observable<any> {

		// get the ProductionOrder from storage
		this.loadHelper( productionOrderId );

	// assign Plant to null
	this.productionOrder.plant = null;

	// save the ProductionOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a AircraftOrder on a ProductionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAircraftOrder( productionOrderId, _aircraftOrderId ): Observable<any> {

		// get the ProductionOrder from storage
		this.loadHelper( productionOrderId );

	// get the AircraftOrder from storage
	var tmp 	= new AircraftOrderService(this.http).getAircraftOrder(_aircraftOrderId);

	// assign the AircraftOrder
	this.productionOrder.aircraftOrder = tmp;

	// save the ProductionOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AircraftOrder on a ProductionOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAircraftOrder( productionOrderId ): Observable<any> {

		// get the ProductionOrder from storage
		this.loadHelper( productionOrderId );

	// assign AircraftOrder to null
	this.productionOrder.aircraftOrder = null;

	// save the ProductionOrder
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ProductionOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ProductionOrder/update/' + this.productionOrder;

	return  this.http.post(uri_, this.productionOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a ProductionOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getProductionOrder(id)
			.subscribe((res : ProductionOrder) => {
				this.productionOrder = res;
			});
	}
}