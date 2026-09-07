import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AircraftModel} from '../models/AircraftModel';
import {AircraftFamilyService} from '../services/AircraftFamily.service';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import {EngineTypeService} from '../services/EngineType.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AircraftModelService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aircraftModel : AircraftModel;

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
	// add a AircraftModel
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAircraftModel(name, modelDesignation, Family, Variants, EngineTypes, AircraftType) : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftModel/create';
		const obj = {
			      		name: name,
      		modelDesignation: modelDesignation,
      		Family: Family != null && Family.length > 0 ? Family : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
      		EngineTypes: EngineTypes != null && EngineTypes.length > 0 ? EngineTypes : null,
			AircraftType: AircraftType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AircraftModel
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAircraftModel(name, modelDesignation, Family, Variants, EngineTypes, AircraftType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AircraftModel/update/' + id;
		const obj = {
				      		name: name,
      		modelDesignation: modelDesignation,
      		Family: Family != null && Family.length > 0 ? Family : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
      		EngineTypes: EngineTypes != null && EngineTypes.length > 0 ? EngineTypes : null,
			AircraftType: AircraftType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AircraftModel
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAircraftModel(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftModel/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AircraftModel
	// returns the results untouched as an Observable AircraftModel
	// AircraftModel model
	// delegates via URI
	//********************************************************************
	getAircraftModel(id) : Observable<AircraftModel> {
		const uri_ = this.apiUrl + '/AircraftModel/load/' + id;

		return this.http.get<AircraftModel>(uri_);
	}
	
	//********************************************************************
	// gets all AircraftModel
	// returns the results untouched as JSON representation of an
	// Observable array of AircraftModel models
	// delegates via URI
	//********************************************************************
	getAircraftModels() : Observable<AircraftModel[]> {
		const uri_ = this.apiUrl + '/AircraftModel/';

		return this
			.http.get<AircraftModel[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Family on a AircraftModel
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFamily( aircraftModelId, _familyId ): Observable<any> {

		// get the AircraftModel from storage
		this.loadHelper( aircraftModelId );

	// get the AircraftFamily from storage
	var tmp 	= new AircraftFamilyService(this.http).getAircraftFamily(_familyId);

	// assign the Family
	this.aircraftModel.family = tmp;

	// save the AircraftModel
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Family on a AircraftModel
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFamily( aircraftModelId ): Observable<any> {

		// get the AircraftModel from storage
		this.loadHelper( aircraftModelId );

	// assign Family to null
	this.aircraftModel.family = null;

	// save the AircraftModel
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more variantsIds as a Variants
	// to a AircraftModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVariants( aircraftModelId, variantsIds ): Observable<any> {

		// get the AircraftModel
		this.loadHelper( aircraftModelId );

	// split on a comma with no spaces
	var idList = variantsIds.split(',')

	// iterate over array of variants ids
	idList.forEach(function (id) {
		// read the AircraftVariant
		var aircraftVariant = new AircraftVariantService(this.http).getAircraftVariant(id);
		// add the AircraftVariant if not already assigned
		if ( this.aircraftModel.variants.indexOf(aircraftVariant) == -1 )
		this.aircraftModel.variants.push(aircraftVariant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more variantsIds as a Variants
	// from a AircraftModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVariants( aircraftModelId, variantsIds ): Observable<any> {

		// get the AircraftModel
		this.loadHelper( aircraftModelId );


	// split on a comma with no spaces
	var idList 					= variantsIds.split(',');
	var variants 	= this.aircraftModel.variants;

	if ( variants != null && variantsIds != null ) {

		// iterate over array of variants ids
		variants.forEach(function (obj) {
			if ( variantsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftVariant
				this.aircraftModel.variants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more engineTypesIds as a EngineTypes
	// to a AircraftModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEngineTypes( aircraftModelId, engineTypesIds ): Observable<any> {

		// get the AircraftModel
		this.loadHelper( aircraftModelId );

	// split on a comma with no spaces
	var idList = engineTypesIds.split(',')

	// iterate over array of engineTypes ids
	idList.forEach(function (id) {
		// read the EngineType
		var engineType = new EngineTypeService(this.http).getEngineType(id);
		// add the EngineType if not already assigned
		if ( this.aircraftModel.engineTypes.indexOf(engineType) == -1 )
		this.aircraftModel.engineTypes.push(engineType);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more engineTypesIds as a EngineTypes
	// from a AircraftModel
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEngineTypes( aircraftModelId, engineTypesIds ): Observable<any> {

		// get the AircraftModel
		this.loadHelper( aircraftModelId );


	// split on a comma with no spaces
	var idList 					= engineTypesIds.split(',');
	var engineTypes 	= this.aircraftModel.engineTypes;

	if ( engineTypes != null && engineTypesIds != null ) {

		// iterate over array of engineTypes ids
		engineTypes.forEach(function (obj) {
			if ( engineTypesIds.indexOf(obj._id) > -1 ) {
				// remove the EngineType
				this.aircraftModel.engineTypes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AircraftModel
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AircraftModel/update/' + this.aircraftModel;

	return  this.http.post(uri_, this.aircraftModel );
}

	//********************************************************************
	// loadHelper - internal helper to load a AircraftModel
	//********************************************************************	
	loadHelper( id ) {
		this.getAircraftModel(id)
			.subscribe((res : AircraftModel) => {
				this.aircraftModel = res;
			});
	}
}