import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InspectionCharacteristic} from '../models/InspectionCharacteristic';
import {InspectionPlanService} from '../services/InspectionPlan.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InspectionCharacteristicService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inspectionCharacteristic : InspectionCharacteristic;

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
	// add a InspectionCharacteristic
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInspectionCharacteristic(characteristicCode, name, lowerSpecLimit, upperSpecLimit, target, InspectionPlan, MeasurementType) : Observable<any> {
		const uri_ = this.apiUrl + '/InspectionCharacteristic/create';
		const obj = {
			      		characteristicCode: characteristicCode,
      		name: name,
      		lowerSpecLimit: lowerSpecLimit,
      		upperSpecLimit: upperSpecLimit,
      		target: target,
      		InspectionPlan: InspectionPlan != null && InspectionPlan.length > 0 ? InspectionPlan : null,
			MeasurementType: MeasurementType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InspectionCharacteristic
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInspectionCharacteristic(characteristicCode, name, lowerSpecLimit, upperSpecLimit, target, InspectionPlan, MeasurementType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InspectionCharacteristic/update/' + id;
		const obj = {
				      		characteristicCode: characteristicCode,
      		name: name,
      		lowerSpecLimit: lowerSpecLimit,
      		upperSpecLimit: upperSpecLimit,
      		target: target,
      		InspectionPlan: InspectionPlan != null && InspectionPlan.length > 0 ? InspectionPlan : null,
			MeasurementType: MeasurementType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InspectionCharacteristic
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInspectionCharacteristic(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InspectionCharacteristic/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InspectionCharacteristic
	// returns the results untouched as an Observable InspectionCharacteristic
	// InspectionCharacteristic model
	// delegates via URI
	//********************************************************************
	getInspectionCharacteristic(id) : Observable<InspectionCharacteristic> {
		const uri_ = this.apiUrl + '/InspectionCharacteristic/load/' + id;

		return this.http.get<InspectionCharacteristic>(uri_);
	}
	
	//********************************************************************
	// gets all InspectionCharacteristic
	// returns the results untouched as JSON representation of an
	// Observable array of InspectionCharacteristic models
	// delegates via URI
	//********************************************************************
	getInspectionCharacteristics() : Observable<InspectionCharacteristic[]> {
		const uri_ = this.apiUrl + '/InspectionCharacteristic/';

		return this
			.http.get<InspectionCharacteristic[]>(uri_);
	}
	
			//********************************************************************
	// assigns a InspectionPlan on a InspectionCharacteristic
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInspectionPlan( inspectionCharacteristicId, _inspectionPlanId ): Observable<any> {

		// get the InspectionCharacteristic from storage
		this.loadHelper( inspectionCharacteristicId );

	// get the InspectionPlan from storage
	var tmp 	= new InspectionPlanService(this.http).getInspectionPlan(_inspectionPlanId);

	// assign the InspectionPlan
	this.inspectionCharacteristic.inspectionPlan = tmp;

	// save the InspectionCharacteristic
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InspectionPlan on a InspectionCharacteristic
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInspectionPlan( inspectionCharacteristicId ): Observable<any> {

		// get the InspectionCharacteristic from storage
		this.loadHelper( inspectionCharacteristicId );

	// assign InspectionPlan to null
	this.inspectionCharacteristic.inspectionPlan = null;

	// save the InspectionCharacteristic
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a InspectionCharacteristic
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InspectionCharacteristic/update/' + this.inspectionCharacteristic;

	return  this.http.post(uri_, this.inspectionCharacteristic );
}

	//********************************************************************
	// loadHelper - internal helper to load a InspectionCharacteristic
	//********************************************************************	
	loadHelper( id ) {
		this.getInspectionCharacteristic(id)
			.subscribe((res : InspectionCharacteristic) => {
				this.inspectionCharacteristic = res;
			});
	}
}