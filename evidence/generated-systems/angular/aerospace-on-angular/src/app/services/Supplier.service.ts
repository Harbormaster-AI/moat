import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Supplier} from '../models/Supplier';
import {AerospaceManufacturerService} from '../services/AerospaceManufacturer.service';
import {Component_Service} from '../services/Component_.service';
import {EngineTypeService} from '../services/EngineType.service';
import {AvionicsSuiteService} from '../services/AvionicsSuite.service';
import {APUService} from '../services/APU.service';
import {LandingGearService} from '../services/LandingGear.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SupplierService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	supplier : Supplier;

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
	// add a Supplier
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSupplier(name, Manufacturers, Components, EngineTypes, AvionicsSuites, Apus, LandingGears, SupplierType, ApprovalStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/Supplier/create';
		const obj = {
			      		name: name,
      		Manufacturers: Manufacturers != null && Manufacturers.length > 0 ? Manufacturers : null,
      		Components: Components != null && Components.length > 0 ? Components : null,
      		EngineTypes: EngineTypes != null && EngineTypes.length > 0 ? EngineTypes : null,
      		AvionicsSuites: AvionicsSuites != null && AvionicsSuites.length > 0 ? AvionicsSuites : null,
      		Apus: Apus != null && Apus.length > 0 ? Apus : null,
      		LandingGears: LandingGears != null && LandingGears.length > 0 ? LandingGears : null,
      		SupplierType: SupplierType,
			ApprovalStatus: ApprovalStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Supplier
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSupplier(name, Manufacturers, Components, EngineTypes, AvionicsSuites, Apus, LandingGears, SupplierType, ApprovalStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Supplier/update/' + id;
		const obj = {
				      		name: name,
      		Manufacturers: Manufacturers != null && Manufacturers.length > 0 ? Manufacturers : null,
      		Components: Components != null && Components.length > 0 ? Components : null,
      		EngineTypes: EngineTypes != null && EngineTypes.length > 0 ? EngineTypes : null,
      		AvionicsSuites: AvionicsSuites != null && AvionicsSuites.length > 0 ? AvionicsSuites : null,
      		Apus: Apus != null && Apus.length > 0 ? Apus : null,
      		LandingGears: LandingGears != null && LandingGears.length > 0 ? LandingGears : null,
      		SupplierType: SupplierType,
			ApprovalStatus: ApprovalStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Supplier
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSupplier(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Supplier/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Supplier
	// returns the results untouched as an Observable Supplier
	// Supplier model
	// delegates via URI
	//********************************************************************
	getSupplier(id) : Observable<Supplier> {
		const uri_ = this.apiUrl + '/Supplier/load/' + id;

		return this.http.get<Supplier>(uri_);
	}
	
	//********************************************************************
	// gets all Supplier
	// returns the results untouched as JSON representation of an
	// Observable array of Supplier models
	// delegates via URI
	//********************************************************************
	getSuppliers() : Observable<Supplier[]> {
		const uri_ = this.apiUrl + '/Supplier/';

		return this
			.http.get<Supplier[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more manufacturersIds as a Manufacturers
	// to a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addManufacturers( supplierId, manufacturersIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );

	// split on a comma with no spaces
	var idList = manufacturersIds.split(',')

	// iterate over array of manufacturers ids
	idList.forEach(function (id) {
		// read the AerospaceManufacturer
		var aerospaceManufacturer = new AerospaceManufacturerService(this.http).getAerospaceManufacturer(id);
		// add the AerospaceManufacturer if not already assigned
		if ( this.supplier.manufacturers.indexOf(aerospaceManufacturer) == -1 )
		this.supplier.manufacturers.push(aerospaceManufacturer);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more manufacturersIds as a Manufacturers
	// from a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeManufacturers( supplierId, manufacturersIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );


	// split on a comma with no spaces
	var idList 					= manufacturersIds.split(',');
	var manufacturers 	= this.supplier.manufacturers;

	if ( manufacturers != null && manufacturersIds != null ) {

		// iterate over array of manufacturers ids
		manufacturers.forEach(function (obj) {
			if ( manufacturersIds.indexOf(obj._id) > -1 ) {
				// remove the AerospaceManufacturer
				this.supplier.manufacturers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more componentsIds as a Components
	// to a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addComponents( supplierId, componentsIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );

	// split on a comma with no spaces
	var idList = componentsIds.split(',')

	// iterate over array of components ids
	idList.forEach(function (id) {
		// read the Component_
		var component_ = new Component_Service(this.http).getComponent_(id);
		// add the Component_ if not already assigned
		if ( this.supplier.components.indexOf(component_) == -1 )
		this.supplier.components.push(component_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more componentsIds as a Components
	// from a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeComponents( supplierId, componentsIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );


	// split on a comma with no spaces
	var idList 					= componentsIds.split(',');
	var components 	= this.supplier.components;

	if ( components != null && componentsIds != null ) {

		// iterate over array of components ids
		components.forEach(function (obj) {
			if ( componentsIds.indexOf(obj._id) > -1 ) {
				// remove the Component_
				this.supplier.components.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more engineTypesIds as a EngineTypes
	// to a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEngineTypes( supplierId, engineTypesIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );

	// split on a comma with no spaces
	var idList = engineTypesIds.split(',')

	// iterate over array of engineTypes ids
	idList.forEach(function (id) {
		// read the EngineType
		var engineType = new EngineTypeService(this.http).getEngineType(id);
		// add the EngineType if not already assigned
		if ( this.supplier.engineTypes.indexOf(engineType) == -1 )
		this.supplier.engineTypes.push(engineType);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more engineTypesIds as a EngineTypes
	// from a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEngineTypes( supplierId, engineTypesIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );


	// split on a comma with no spaces
	var idList 					= engineTypesIds.split(',');
	var engineTypes 	= this.supplier.engineTypes;

	if ( engineTypes != null && engineTypesIds != null ) {

		// iterate over array of engineTypes ids
		engineTypes.forEach(function (obj) {
			if ( engineTypesIds.indexOf(obj._id) > -1 ) {
				// remove the EngineType
				this.supplier.engineTypes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more avionicsSuitesIds as a AvionicsSuites
	// to a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAvionicsSuites( supplierId, avionicsSuitesIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );

	// split on a comma with no spaces
	var idList = avionicsSuitesIds.split(',')

	// iterate over array of avionicsSuites ids
	idList.forEach(function (id) {
		// read the AvionicsSuite
		var avionicsSuite = new AvionicsSuiteService(this.http).getAvionicsSuite(id);
		// add the AvionicsSuite if not already assigned
		if ( this.supplier.avionicsSuites.indexOf(avionicsSuite) == -1 )
		this.supplier.avionicsSuites.push(avionicsSuite);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more avionicsSuitesIds as a AvionicsSuites
	// from a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAvionicsSuites( supplierId, avionicsSuitesIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );


	// split on a comma with no spaces
	var idList 					= avionicsSuitesIds.split(',');
	var avionicsSuites 	= this.supplier.avionicsSuites;

	if ( avionicsSuites != null && avionicsSuitesIds != null ) {

		// iterate over array of avionicsSuites ids
		avionicsSuites.forEach(function (obj) {
			if ( avionicsSuitesIds.indexOf(obj._id) > -1 ) {
				// remove the AvionicsSuite
				this.supplier.avionicsSuites.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more apusIds as a Apus
	// to a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addApus( supplierId, apusIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );

	// split on a comma with no spaces
	var idList = apusIds.split(',')

	// iterate over array of apus ids
	idList.forEach(function (id) {
		// read the APU
		var aPU = new APUService(this.http).getAPU(id);
		// add the APU if not already assigned
		if ( this.supplier.apus.indexOf(aPU) == -1 )
		this.supplier.apus.push(aPU);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more apusIds as a Apus
	// from a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeApus( supplierId, apusIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );


	// split on a comma with no spaces
	var idList 					= apusIds.split(',');
	var apus 	= this.supplier.apus;

	if ( apus != null && apusIds != null ) {

		// iterate over array of apus ids
		apus.forEach(function (obj) {
			if ( apusIds.indexOf(obj._id) > -1 ) {
				// remove the APU
				this.supplier.apus.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more landingGearsIds as a LandingGears
	// to a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLandingGears( supplierId, landingGearsIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );

	// split on a comma with no spaces
	var idList = landingGearsIds.split(',')

	// iterate over array of landingGears ids
	idList.forEach(function (id) {
		// read the LandingGear
		var landingGear = new LandingGearService(this.http).getLandingGear(id);
		// add the LandingGear if not already assigned
		if ( this.supplier.landingGears.indexOf(landingGear) == -1 )
		this.supplier.landingGears.push(landingGear);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more landingGearsIds as a LandingGears
	// from a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLandingGears( supplierId, landingGearsIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );


	// split on a comma with no spaces
	var idList 					= landingGearsIds.split(',');
	var landingGears 	= this.supplier.landingGears;

	if ( landingGears != null && landingGearsIds != null ) {

		// iterate over array of landingGears ids
		landingGears.forEach(function (obj) {
			if ( landingGearsIds.indexOf(obj._id) > -1 ) {
				// remove the LandingGear
				this.supplier.landingGears.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Supplier
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Supplier/update/' + this.supplier;

	return  this.http.post(uri_, this.supplier );
}

	//********************************************************************
	// loadHelper - internal helper to load a Supplier
	//********************************************************************	
	loadHelper( id ) {
		this.getSupplier(id)
			.subscribe((res : Supplier) => {
				this.supplier = res;
			});
	}
}