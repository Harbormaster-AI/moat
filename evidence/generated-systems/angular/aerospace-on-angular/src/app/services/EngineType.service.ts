import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {EngineType} from '../models/EngineType';
import {SupplierService} from '../services/Supplier.service';
import {AircraftModelService} from '../services/AircraftModel.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EngineTypeService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	engineType : EngineType;

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
	// add a EngineType
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEngineType(engineModelCode, maxThrustKn, Supplier, CompatibleModels, Category) : Observable<any> {
		const uri_ = this.apiUrl + '/EngineType/create';
		const obj = {
			      		engineModelCode: engineModelCode,
      		maxThrustKn: maxThrustKn,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
      		CompatibleModels: CompatibleModels != null && CompatibleModels.length > 0 ? CompatibleModels : null,
			Category: Category
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a EngineType
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEngineType(engineModelCode, maxThrustKn, Supplier, CompatibleModels, Category, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/EngineType/update/' + id;
		const obj = {
				      		engineModelCode: engineModelCode,
      		maxThrustKn: maxThrustKn,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
      		CompatibleModels: CompatibleModels != null && CompatibleModels.length > 0 ? CompatibleModels : null,
			Category: Category
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a EngineType
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEngineType(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/EngineType/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a EngineType
	// returns the results untouched as an Observable EngineType
	// EngineType model
	// delegates via URI
	//********************************************************************
	getEngineType(id) : Observable<EngineType> {
		const uri_ = this.apiUrl + '/EngineType/load/' + id;

		return this.http.get<EngineType>(uri_);
	}
	
	//********************************************************************
	// gets all EngineType
	// returns the results untouched as JSON representation of an
	// Observable array of EngineType models
	// delegates via URI
	//********************************************************************
	getEngineTypes() : Observable<EngineType[]> {
		const uri_ = this.apiUrl + '/EngineType/';

		return this
			.http.get<EngineType[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Supplier on a EngineType
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSupplier( engineTypeId, _supplierId ): Observable<any> {

		// get the EngineType from storage
		this.loadHelper( engineTypeId );

	// get the Supplier from storage
	var tmp 	= new SupplierService(this.http).getSupplier(_supplierId);

	// assign the Supplier
	this.engineType.supplier = tmp;

	// save the EngineType
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Supplier on a EngineType
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSupplier( engineTypeId ): Observable<any> {

		// get the EngineType from storage
		this.loadHelper( engineTypeId );

	// assign Supplier to null
	this.engineType.supplier = null;

	// save the EngineType
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more compatibleModelsIds as a CompatibleModels
	// to a EngineType
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCompatibleModels( engineTypeId, compatibleModelsIds ): Observable<any> {

		// get the EngineType
		this.loadHelper( engineTypeId );

	// split on a comma with no spaces
	var idList = compatibleModelsIds.split(',')

	// iterate over array of compatibleModels ids
	idList.forEach(function (id) {
		// read the AircraftModel
		var aircraftModel = new AircraftModelService(this.http).getAircraftModel(id);
		// add the AircraftModel if not already assigned
		if ( this.engineType.compatibleModels.indexOf(aircraftModel) == -1 )
		this.engineType.compatibleModels.push(aircraftModel);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more compatibleModelsIds as a CompatibleModels
	// from a EngineType
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCompatibleModels( engineTypeId, compatibleModelsIds ): Observable<any> {

		// get the EngineType
		this.loadHelper( engineTypeId );


	// split on a comma with no spaces
	var idList 					= compatibleModelsIds.split(',');
	var compatibleModels 	= this.engineType.compatibleModels;

	if ( compatibleModels != null && compatibleModelsIds != null ) {

		// iterate over array of compatibleModels ids
		compatibleModels.forEach(function (obj) {
			if ( compatibleModelsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftModel
				this.engineType.compatibleModels.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a EngineType
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/EngineType/update/' + this.engineType;

	return  this.http.post(uri_, this.engineType );
}

	//********************************************************************
	// loadHelper - internal helper to load a EngineType
	//********************************************************************	
	loadHelper( id ) {
		this.getEngineType(id)
			.subscribe((res : EngineType) => {
				this.engineType = res;
			});
	}
}