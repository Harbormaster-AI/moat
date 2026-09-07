import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AircraftPackage} from '../models/AircraftPackage';
import {AircraftOptionService} from '../services/AircraftOption.service';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AircraftPackageService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aircraftPackage : AircraftPackage;

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
	// add a AircraftPackage
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAircraftPackage(name, Options, Variants, PackageType) : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftPackage/create';
		const obj = {
			      		name: name,
      		Options: Options != null && Options.length > 0 ? Options : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
			PackageType: PackageType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AircraftPackage
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAircraftPackage(name, Options, Variants, PackageType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AircraftPackage/update/' + id;
		const obj = {
				      		name: name,
      		Options: Options != null && Options.length > 0 ? Options : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
			PackageType: PackageType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AircraftPackage
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAircraftPackage(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftPackage/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AircraftPackage
	// returns the results untouched as an Observable AircraftPackage
	// AircraftPackage model
	// delegates via URI
	//********************************************************************
	getAircraftPackage(id) : Observable<AircraftPackage> {
		const uri_ = this.apiUrl + '/AircraftPackage/load/' + id;

		return this.http.get<AircraftPackage>(uri_);
	}
	
	//********************************************************************
	// gets all AircraftPackage
	// returns the results untouched as JSON representation of an
	// Observable array of AircraftPackage models
	// delegates via URI
	//********************************************************************
	getAircraftPackages() : Observable<AircraftPackage[]> {
		const uri_ = this.apiUrl + '/AircraftPackage/';

		return this
			.http.get<AircraftPackage[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more optionsIds as a Options
	// to a AircraftPackage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOptions( aircraftPackageId, optionsIds ): Observable<any> {

		// get the AircraftPackage
		this.loadHelper( aircraftPackageId );

	// split on a comma with no spaces
	var idList = optionsIds.split(',')

	// iterate over array of options ids
	idList.forEach(function (id) {
		// read the AircraftOption
		var aircraftOption = new AircraftOptionService(this.http).getAircraftOption(id);
		// add the AircraftOption if not already assigned
		if ( this.aircraftPackage.options.indexOf(aircraftOption) == -1 )
		this.aircraftPackage.options.push(aircraftOption);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more optionsIds as a Options
	// from a AircraftPackage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOptions( aircraftPackageId, optionsIds ): Observable<any> {

		// get the AircraftPackage
		this.loadHelper( aircraftPackageId );


	// split on a comma with no spaces
	var idList 					= optionsIds.split(',');
	var options 	= this.aircraftPackage.options;

	if ( options != null && optionsIds != null ) {

		// iterate over array of options ids
		options.forEach(function (obj) {
			if ( optionsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftOption
				this.aircraftPackage.options.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more variantsIds as a Variants
	// to a AircraftPackage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVariants( aircraftPackageId, variantsIds ): Observable<any> {

		// get the AircraftPackage
		this.loadHelper( aircraftPackageId );

	// split on a comma with no spaces
	var idList = variantsIds.split(',')

	// iterate over array of variants ids
	idList.forEach(function (id) {
		// read the AircraftVariant
		var aircraftVariant = new AircraftVariantService(this.http).getAircraftVariant(id);
		// add the AircraftVariant if not already assigned
		if ( this.aircraftPackage.variants.indexOf(aircraftVariant) == -1 )
		this.aircraftPackage.variants.push(aircraftVariant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more variantsIds as a Variants
	// from a AircraftPackage
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVariants( aircraftPackageId, variantsIds ): Observable<any> {

		// get the AircraftPackage
		this.loadHelper( aircraftPackageId );


	// split on a comma with no spaces
	var idList 					= variantsIds.split(',');
	var variants 	= this.aircraftPackage.variants;

	if ( variants != null && variantsIds != null ) {

		// iterate over array of variants ids
		variants.forEach(function (obj) {
			if ( variantsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftVariant
				this.aircraftPackage.variants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AircraftPackage
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AircraftPackage/update/' + this.aircraftPackage;

	return  this.http.post(uri_, this.aircraftPackage );
}

	//********************************************************************
	// loadHelper - internal helper to load a AircraftPackage
	//********************************************************************	
	loadHelper( id ) {
		this.getAircraftPackage(id)
			.subscribe((res : AircraftPackage) => {
				this.aircraftPackage = res;
			});
	}
}