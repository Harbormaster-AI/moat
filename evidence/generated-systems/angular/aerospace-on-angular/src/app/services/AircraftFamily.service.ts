import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AircraftFamily} from '../models/AircraftFamily';
import {AircraftProgramService} from '../services/AircraftProgram.service';
import {AircraftModelService} from '../services/AircraftModel.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AircraftFamilyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aircraftFamily : AircraftFamily;

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
	// add a AircraftFamily
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAircraftFamily(name, familyCode, Program, AircraftModels) : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftFamily/create';
		const obj = {
			      		name: name,
      		familyCode: familyCode,
      		Program: Program != null && Program.length > 0 ? Program : null,
			AircraftModels: AircraftModels != null && AircraftModels.length > 0 ? AircraftModels : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AircraftFamily
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAircraftFamily(name, familyCode, Program, AircraftModels, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AircraftFamily/update/' + id;
		const obj = {
				      		name: name,
      		familyCode: familyCode,
      		Program: Program != null && Program.length > 0 ? Program : null,
			AircraftModels: AircraftModels != null && AircraftModels.length > 0 ? AircraftModels : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AircraftFamily
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAircraftFamily(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftFamily/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AircraftFamily
	// returns the results untouched as an Observable AircraftFamily
	// AircraftFamily model
	// delegates via URI
	//********************************************************************
	getAircraftFamily(id) : Observable<AircraftFamily> {
		const uri_ = this.apiUrl + '/AircraftFamily/load/' + id;

		return this.http.get<AircraftFamily>(uri_);
	}
	
	//********************************************************************
	// gets all AircraftFamily
	// returns the results untouched as JSON representation of an
	// Observable array of AircraftFamily models
	// delegates via URI
	//********************************************************************
	getAircraftFamilys() : Observable<AircraftFamily[]> {
		const uri_ = this.apiUrl + '/AircraftFamily/';

		return this
			.http.get<AircraftFamily[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Program on a AircraftFamily
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProgram( aircraftFamilyId, _programId ): Observable<any> {

		// get the AircraftFamily from storage
		this.loadHelper( aircraftFamilyId );

	// get the AircraftProgram from storage
	var tmp 	= new AircraftProgramService(this.http).getAircraftProgram(_programId);

	// assign the Program
	this.aircraftFamily.program = tmp;

	// save the AircraftFamily
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Program on a AircraftFamily
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProgram( aircraftFamilyId ): Observable<any> {

		// get the AircraftFamily from storage
		this.loadHelper( aircraftFamilyId );

	// assign Program to null
	this.aircraftFamily.program = null;

	// save the AircraftFamily
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more aircraftModelsIds as a AircraftModels
	// to a AircraftFamily
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAircraftModels( aircraftFamilyId, aircraftModelsIds ): Observable<any> {

		// get the AircraftFamily
		this.loadHelper( aircraftFamilyId );

	// split on a comma with no spaces
	var idList = aircraftModelsIds.split(',')

	// iterate over array of aircraftModels ids
	idList.forEach(function (id) {
		// read the AircraftModel
		var aircraftModel = new AircraftModelService(this.http).getAircraftModel(id);
		// add the AircraftModel if not already assigned
		if ( this.aircraftFamily.aircraftModels.indexOf(aircraftModel) == -1 )
		this.aircraftFamily.aircraftModels.push(aircraftModel);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more aircraftModelsIds as a AircraftModels
	// from a AircraftFamily
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAircraftModels( aircraftFamilyId, aircraftModelsIds ): Observable<any> {

		// get the AircraftFamily
		this.loadHelper( aircraftFamilyId );


	// split on a comma with no spaces
	var idList 					= aircraftModelsIds.split(',');
	var aircraftModels 	= this.aircraftFamily.aircraftModels;

	if ( aircraftModels != null && aircraftModelsIds != null ) {

		// iterate over array of aircraftModels ids
		aircraftModels.forEach(function (obj) {
			if ( aircraftModelsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftModel
				this.aircraftFamily.aircraftModels.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AircraftFamily
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AircraftFamily/update/' + this.aircraftFamily;

	return  this.http.post(uri_, this.aircraftFamily );
}

	//********************************************************************
	// loadHelper - internal helper to load a AircraftFamily
	//********************************************************************	
	loadHelper( id ) {
		this.getAircraftFamily(id)
			.subscribe((res : AircraftFamily) => {
				this.aircraftFamily = res;
			});
	}
}