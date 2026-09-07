import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AircraftVariant} from '../models/AircraftVariant';
import {AircraftModelService} from '../services/AircraftModel.service';
import {EngineTypeService} from '../services/EngineType.service';
import {AvionicsSuiteService} from '../services/AvionicsSuite.service';
import {APUService} from '../services/APU.service';
import {LandingGearService} from '../services/LandingGear.service';
import {CabinLayoutService} from '../services/CabinLayout.service';
import {AircraftOptionService} from '../services/AircraftOption.service';
import {AircraftPackageService} from '../services/AircraftPackage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AircraftVariantService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aircraftVariant : AircraftVariant;

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
	// add a AircraftVariant
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAircraftVariant(variantCode, rangeNm, maxTakeoffWeightKg, Model_, EngineType, AvionicsSuite, Apu, LandingGear, CabinLayouts, Options, Packages) : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftVariant/create';
		const obj = {
			      		variantCode: variantCode,
      		rangeNm: rangeNm,
      		maxTakeoffWeightKg: maxTakeoffWeightKg,
      		Model_: Model_ != null && Model_.length > 0 ? Model_ : null,
      		EngineType: EngineType != null && EngineType.length > 0 ? EngineType : null,
      		AvionicsSuite: AvionicsSuite != null && AvionicsSuite.length > 0 ? AvionicsSuite : null,
      		Apu: Apu != null && Apu.length > 0 ? Apu : null,
      		LandingGear: LandingGear != null && LandingGear.length > 0 ? LandingGear : null,
      		CabinLayouts: CabinLayouts != null && CabinLayouts.length > 0 ? CabinLayouts : null,
      		Options: Options != null && Options.length > 0 ? Options : null,
			Packages: Packages != null && Packages.length > 0 ? Packages : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAircraftVariant(variantCode, rangeNm, maxTakeoffWeightKg, Model_, EngineType, AvionicsSuite, Apu, LandingGear, CabinLayouts, Options, Packages, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AircraftVariant/update/' + id;
		const obj = {
				      		variantCode: variantCode,
      		rangeNm: rangeNm,
      		maxTakeoffWeightKg: maxTakeoffWeightKg,
      		Model_: Model_ != null && Model_.length > 0 ? Model_ : null,
      		EngineType: EngineType != null && EngineType.length > 0 ? EngineType : null,
      		AvionicsSuite: AvionicsSuite != null && AvionicsSuite.length > 0 ? AvionicsSuite : null,
      		Apu: Apu != null && Apu.length > 0 ? Apu : null,
      		LandingGear: LandingGear != null && LandingGear.length > 0 ? LandingGear : null,
      		CabinLayouts: CabinLayouts != null && CabinLayouts.length > 0 ? CabinLayouts : null,
      		Options: Options != null && Options.length > 0 ? Options : null,
			Packages: Packages != null && Packages.length > 0 ? Packages : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAircraftVariant(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftVariant/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AircraftVariant
	// returns the results untouched as an Observable AircraftVariant
	// AircraftVariant model
	// delegates via URI
	//********************************************************************
	getAircraftVariant(id) : Observable<AircraftVariant> {
		const uri_ = this.apiUrl + '/AircraftVariant/load/' + id;

		return this.http.get<AircraftVariant>(uri_);
	}
	
	//********************************************************************
	// gets all AircraftVariant
	// returns the results untouched as JSON representation of an
	// Observable array of AircraftVariant models
	// delegates via URI
	//********************************************************************
	getAircraftVariants() : Observable<AircraftVariant[]> {
		const uri_ = this.apiUrl + '/AircraftVariant/';

		return this
			.http.get<AircraftVariant[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Model_ on a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignModel_( aircraftVariantId, _model_Id ): Observable<any> {

		// get the AircraftVariant from storage
		this.loadHelper( aircraftVariantId );

	// get the AircraftModel from storage
	var tmp 	= new AircraftModelService(this.http).getAircraftModel(_model_Id);

	// assign the Model_
	this.aircraftVariant.model_ = tmp;

	// save the AircraftVariant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Model_ on a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignModel_( aircraftVariantId ): Observable<any> {

		// get the AircraftVariant from storage
		this.loadHelper( aircraftVariantId );

	// assign Model_ to null
	this.aircraftVariant.model_ = null;

	// save the AircraftVariant
	return this.saveHelper();
}

		//********************************************************************
	// assigns a EngineType on a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEngineType( aircraftVariantId, _engineTypeId ): Observable<any> {

		// get the AircraftVariant from storage
		this.loadHelper( aircraftVariantId );

	// get the EngineType from storage
	var tmp 	= new EngineTypeService(this.http).getEngineType(_engineTypeId);

	// assign the EngineType
	this.aircraftVariant.engineType = tmp;

	// save the AircraftVariant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a EngineType on a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEngineType( aircraftVariantId ): Observable<any> {

		// get the AircraftVariant from storage
		this.loadHelper( aircraftVariantId );

	// assign EngineType to null
	this.aircraftVariant.engineType = null;

	// save the AircraftVariant
	return this.saveHelper();
}

		//********************************************************************
	// assigns a AvionicsSuite on a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAvionicsSuite( aircraftVariantId, _avionicsSuiteId ): Observable<any> {

		// get the AircraftVariant from storage
		this.loadHelper( aircraftVariantId );

	// get the AvionicsSuite from storage
	var tmp 	= new AvionicsSuiteService(this.http).getAvionicsSuite(_avionicsSuiteId);

	// assign the AvionicsSuite
	this.aircraftVariant.avionicsSuite = tmp;

	// save the AircraftVariant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AvionicsSuite on a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAvionicsSuite( aircraftVariantId ): Observable<any> {

		// get the AircraftVariant from storage
		this.loadHelper( aircraftVariantId );

	// assign AvionicsSuite to null
	this.aircraftVariant.avionicsSuite = null;

	// save the AircraftVariant
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Apu on a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignApu( aircraftVariantId, _apuId ): Observable<any> {

		// get the AircraftVariant from storage
		this.loadHelper( aircraftVariantId );

	// get the APU from storage
	var tmp 	= new APUService(this.http).getAPU(_apuId);

	// assign the Apu
	this.aircraftVariant.apu = tmp;

	// save the AircraftVariant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Apu on a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignApu( aircraftVariantId ): Observable<any> {

		// get the AircraftVariant from storage
		this.loadHelper( aircraftVariantId );

	// assign Apu to null
	this.aircraftVariant.apu = null;

	// save the AircraftVariant
	return this.saveHelper();
}

		//********************************************************************
	// assigns a LandingGear on a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLandingGear( aircraftVariantId, _landingGearId ): Observable<any> {

		// get the AircraftVariant from storage
		this.loadHelper( aircraftVariantId );

	// get the LandingGear from storage
	var tmp 	= new LandingGearService(this.http).getLandingGear(_landingGearId);

	// assign the LandingGear
	this.aircraftVariant.landingGear = tmp;

	// save the AircraftVariant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a LandingGear on a AircraftVariant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLandingGear( aircraftVariantId ): Observable<any> {

		// get the AircraftVariant from storage
		this.loadHelper( aircraftVariantId );

	// assign LandingGear to null
	this.aircraftVariant.landingGear = null;

	// save the AircraftVariant
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more cabinLayoutsIds as a CabinLayouts
	// to a AircraftVariant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCabinLayouts( aircraftVariantId, cabinLayoutsIds ): Observable<any> {

		// get the AircraftVariant
		this.loadHelper( aircraftVariantId );

	// split on a comma with no spaces
	var idList = cabinLayoutsIds.split(',')

	// iterate over array of cabinLayouts ids
	idList.forEach(function (id) {
		// read the CabinLayout
		var cabinLayout = new CabinLayoutService(this.http).getCabinLayout(id);
		// add the CabinLayout if not already assigned
		if ( this.aircraftVariant.cabinLayouts.indexOf(cabinLayout) == -1 )
		this.aircraftVariant.cabinLayouts.push(cabinLayout);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more cabinLayoutsIds as a CabinLayouts
	// from a AircraftVariant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCabinLayouts( aircraftVariantId, cabinLayoutsIds ): Observable<any> {

		// get the AircraftVariant
		this.loadHelper( aircraftVariantId );


	// split on a comma with no spaces
	var idList 					= cabinLayoutsIds.split(',');
	var cabinLayouts 	= this.aircraftVariant.cabinLayouts;

	if ( cabinLayouts != null && cabinLayoutsIds != null ) {

		// iterate over array of cabinLayouts ids
		cabinLayouts.forEach(function (obj) {
			if ( cabinLayoutsIds.indexOf(obj._id) > -1 ) {
				// remove the CabinLayout
				this.aircraftVariant.cabinLayouts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more optionsIds as a Options
	// to a AircraftVariant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOptions( aircraftVariantId, optionsIds ): Observable<any> {

		// get the AircraftVariant
		this.loadHelper( aircraftVariantId );

	// split on a comma with no spaces
	var idList = optionsIds.split(',')

	// iterate over array of options ids
	idList.forEach(function (id) {
		// read the AircraftOption
		var aircraftOption = new AircraftOptionService(this.http).getAircraftOption(id);
		// add the AircraftOption if not already assigned
		if ( this.aircraftVariant.options.indexOf(aircraftOption) == -1 )
		this.aircraftVariant.options.push(aircraftOption);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more optionsIds as a Options
	// from a AircraftVariant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOptions( aircraftVariantId, optionsIds ): Observable<any> {

		// get the AircraftVariant
		this.loadHelper( aircraftVariantId );


	// split on a comma with no spaces
	var idList 					= optionsIds.split(',');
	var options 	= this.aircraftVariant.options;

	if ( options != null && optionsIds != null ) {

		// iterate over array of options ids
		options.forEach(function (obj) {
			if ( optionsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftOption
				this.aircraftVariant.options.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more packagesIds as a Packages
	// to a AircraftVariant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPackages( aircraftVariantId, packagesIds ): Observable<any> {

		// get the AircraftVariant
		this.loadHelper( aircraftVariantId );

	// split on a comma with no spaces
	var idList = packagesIds.split(',')

	// iterate over array of packages ids
	idList.forEach(function (id) {
		// read the AircraftPackage
		var aircraftPackage = new AircraftPackageService(this.http).getAircraftPackage(id);
		// add the AircraftPackage if not already assigned
		if ( this.aircraftVariant.packages.indexOf(aircraftPackage) == -1 )
		this.aircraftVariant.packages.push(aircraftPackage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more packagesIds as a Packages
	// from a AircraftVariant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePackages( aircraftVariantId, packagesIds ): Observable<any> {

		// get the AircraftVariant
		this.loadHelper( aircraftVariantId );


	// split on a comma with no spaces
	var idList 					= packagesIds.split(',');
	var packages 	= this.aircraftVariant.packages;

	if ( packages != null && packagesIds != null ) {

		// iterate over array of packages ids
		packages.forEach(function (obj) {
			if ( packagesIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftPackage
				this.aircraftVariant.packages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AircraftVariant
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AircraftVariant/update/' + this.aircraftVariant;

	return  this.http.post(uri_, this.aircraftVariant );
}

	//********************************************************************
	// loadHelper - internal helper to load a AircraftVariant
	//********************************************************************	
	loadHelper( id ) {
		this.getAircraftVariant(id)
			.subscribe((res : AircraftVariant) => {
				this.aircraftVariant = res;
			});
	}
}