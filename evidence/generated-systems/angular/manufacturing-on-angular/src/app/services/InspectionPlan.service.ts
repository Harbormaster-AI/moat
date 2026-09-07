import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InspectionPlan} from '../models/InspectionPlan';
import {ItemService} from '../services/Item.service';
import {InspectionCharacteristicService} from '../services/InspectionCharacteristic.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InspectionPlanService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inspectionPlan : InspectionPlan;

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
	// add a InspectionPlan
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInspectionPlan(planNumber, revision, Item, Characteristics, SamplingPlan, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/InspectionPlan/create';
		const obj = {
			      		planNumber: planNumber,
      		revision: revision,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		Characteristics: Characteristics != null && Characteristics.length > 0 ? Characteristics : null,
      		SamplingPlan: SamplingPlan,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InspectionPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInspectionPlan(planNumber, revision, Item, Characteristics, SamplingPlan, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InspectionPlan/update/' + id;
		const obj = {
				      		planNumber: planNumber,
      		revision: revision,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		Characteristics: Characteristics != null && Characteristics.length > 0 ? Characteristics : null,
      		SamplingPlan: SamplingPlan,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InspectionPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInspectionPlan(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InspectionPlan/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InspectionPlan
	// returns the results untouched as an Observable InspectionPlan
	// InspectionPlan model
	// delegates via URI
	//********************************************************************
	getInspectionPlan(id) : Observable<InspectionPlan> {
		const uri_ = this.apiUrl + '/InspectionPlan/load/' + id;

		return this.http.get<InspectionPlan>(uri_);
	}
	
	//********************************************************************
	// gets all InspectionPlan
	// returns the results untouched as JSON representation of an
	// Observable array of InspectionPlan models
	// delegates via URI
	//********************************************************************
	getInspectionPlans() : Observable<InspectionPlan[]> {
		const uri_ = this.apiUrl + '/InspectionPlan/';

		return this
			.http.get<InspectionPlan[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Item on a InspectionPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( inspectionPlanId, _itemId ): Observable<any> {

		// get the InspectionPlan from storage
		this.loadHelper( inspectionPlanId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.inspectionPlan.item = tmp;

	// save the InspectionPlan
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a InspectionPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( inspectionPlanId ): Observable<any> {

		// get the InspectionPlan from storage
		this.loadHelper( inspectionPlanId );

	// assign Item to null
	this.inspectionPlan.item = null;

	// save the InspectionPlan
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more characteristicsIds as a Characteristics
	// to a InspectionPlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCharacteristics( inspectionPlanId, characteristicsIds ): Observable<any> {

		// get the InspectionPlan
		this.loadHelper( inspectionPlanId );

	// split on a comma with no spaces
	var idList = characteristicsIds.split(',')

	// iterate over array of characteristics ids
	idList.forEach(function (id) {
		// read the InspectionCharacteristic
		var inspectionCharacteristic = new InspectionCharacteristicService(this.http).getInspectionCharacteristic(id);
		// add the InspectionCharacteristic if not already assigned
		if ( this.inspectionPlan.characteristics.indexOf(inspectionCharacteristic) == -1 )
		this.inspectionPlan.characteristics.push(inspectionCharacteristic);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more characteristicsIds as a Characteristics
	// from a InspectionPlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCharacteristics( inspectionPlanId, characteristicsIds ): Observable<any> {

		// get the InspectionPlan
		this.loadHelper( inspectionPlanId );


	// split on a comma with no spaces
	var idList 					= characteristicsIds.split(',');
	var characteristics 	= this.inspectionPlan.characteristics;

	if ( characteristics != null && characteristicsIds != null ) {

		// iterate over array of characteristics ids
		characteristics.forEach(function (obj) {
			if ( characteristicsIds.indexOf(obj._id) > -1 ) {
				// remove the InspectionCharacteristic
				this.inspectionPlan.characteristics.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InspectionPlan
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InspectionPlan/update/' + this.inspectionPlan;

	return  this.http.post(uri_, this.inspectionPlan );
}

	//********************************************************************
	// loadHelper - internal helper to load a InspectionPlan
	//********************************************************************	
	loadHelper( id ) {
		this.getInspectionPlan(id)
			.subscribe((res : InspectionPlan) => {
				this.inspectionPlan = res;
			});
	}
}