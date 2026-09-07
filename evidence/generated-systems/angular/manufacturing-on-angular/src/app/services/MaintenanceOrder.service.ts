import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {MaintenanceOrder} from '../models/MaintenanceOrder';
import {AssetService} from '../services/Asset.service';
import {MaintenancePlanService} from '../services/MaintenancePlan.service';
import {WorkCenterService} from '../services/WorkCenter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MaintenanceOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	maintenanceOrder : MaintenanceOrder;

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
	// add a MaintenanceOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMaintenanceOrder(orderNumber, priority, requestedDate, completionDate, Asset, Plan, WorkCenter, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/MaintenanceOrder/create';
		const obj = {
			      		orderNumber: orderNumber,
      		priority: priority,
      		requestedDate: requestedDate,
      		completionDate: completionDate,
      		Asset: Asset != null && Asset.length > 0 ? Asset : null,
      		Plan: Plan != null && Plan.length > 0 ? Plan : null,
      		WorkCenter: WorkCenter != null && WorkCenter.length > 0 ? WorkCenter : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a MaintenanceOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMaintenanceOrder(orderNumber, priority, requestedDate, completionDate, Asset, Plan, WorkCenter, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/MaintenanceOrder/update/' + id;
		const obj = {
				      		orderNumber: orderNumber,
      		priority: priority,
      		requestedDate: requestedDate,
      		completionDate: completionDate,
      		Asset: Asset != null && Asset.length > 0 ? Asset : null,
      		Plan: Plan != null && Plan.length > 0 ? Plan : null,
      		WorkCenter: WorkCenter != null && WorkCenter.length > 0 ? WorkCenter : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a MaintenanceOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMaintenanceOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/MaintenanceOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a MaintenanceOrder
	// returns the results untouched as an Observable MaintenanceOrder
	// MaintenanceOrder model
	// delegates via URI
	//********************************************************************
	getMaintenanceOrder(id) : Observable<MaintenanceOrder> {
		const uri_ = this.apiUrl + '/MaintenanceOrder/load/' + id;

		return this.http.get<MaintenanceOrder>(uri_);
	}
	
	//********************************************************************
	// gets all MaintenanceOrder
	// returns the results untouched as JSON representation of an
	// Observable array of MaintenanceOrder models
	// delegates via URI
	//********************************************************************
	getMaintenanceOrders() : Observable<MaintenanceOrder[]> {
		const uri_ = this.apiUrl + '/MaintenanceOrder/';

		return this
			.http.get<MaintenanceOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Asset on a MaintenanceOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAsset( maintenanceOrderId, _assetId ): Observable<any> {

		// get the MaintenanceOrder from storage
		this.loadHelper( maintenanceOrderId );

	// get the Asset from storage
	var tmp 	= new AssetService(this.http).getAsset(_assetId);

	// assign the Asset
	this.maintenanceOrder.asset = tmp;

	// save the MaintenanceOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Asset on a MaintenanceOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAsset( maintenanceOrderId ): Observable<any> {

		// get the MaintenanceOrder from storage
		this.loadHelper( maintenanceOrderId );

	// assign Asset to null
	this.maintenanceOrder.asset = null;

	// save the MaintenanceOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Plan on a MaintenanceOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlan( maintenanceOrderId, _planId ): Observable<any> {

		// get the MaintenanceOrder from storage
		this.loadHelper( maintenanceOrderId );

	// get the MaintenancePlan from storage
	var tmp 	= new MaintenancePlanService(this.http).getMaintenancePlan(_planId);

	// assign the Plan
	this.maintenanceOrder.plan = tmp;

	// save the MaintenanceOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plan on a MaintenanceOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlan( maintenanceOrderId ): Observable<any> {

		// get the MaintenanceOrder from storage
		this.loadHelper( maintenanceOrderId );

	// assign Plan to null
	this.maintenanceOrder.plan = null;

	// save the MaintenanceOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a WorkCenter on a MaintenanceOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkCenter( maintenanceOrderId, _workCenterId ): Observable<any> {

		// get the MaintenanceOrder from storage
		this.loadHelper( maintenanceOrderId );

	// get the WorkCenter from storage
	var tmp 	= new WorkCenterService(this.http).getWorkCenter(_workCenterId);

	// assign the WorkCenter
	this.maintenanceOrder.workCenter = tmp;

	// save the MaintenanceOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkCenter on a MaintenanceOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkCenter( maintenanceOrderId ): Observable<any> {

		// get the MaintenanceOrder from storage
		this.loadHelper( maintenanceOrderId );

	// assign WorkCenter to null
	this.maintenanceOrder.workCenter = null;

	// save the MaintenanceOrder
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a MaintenanceOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/MaintenanceOrder/update/' + this.maintenanceOrder;

	return  this.http.post(uri_, this.maintenanceOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a MaintenanceOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getMaintenanceOrder(id)
			.subscribe((res : MaintenanceOrder) => {
				this.maintenanceOrder = res;
			});
	}
}