import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AircraftOption} from '../models/AircraftOption';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import {AircraftPackageService} from '../services/AircraftPackage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AircraftOptionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aircraftOption : AircraftOption;

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
	// add a AircraftOption
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAircraftOption(code, name, Variants, Packages, OptionCategory) : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftOption/create';
		const obj = {
			      		code: code,
      		name: name,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
      		Packages: Packages != null && Packages.length > 0 ? Packages : null,
			OptionCategory: OptionCategory
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AircraftOption
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAircraftOption(code, name, Variants, Packages, OptionCategory, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AircraftOption/update/' + id;
		const obj = {
				      		code: code,
      		name: name,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
      		Packages: Packages != null && Packages.length > 0 ? Packages : null,
			OptionCategory: OptionCategory
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AircraftOption
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAircraftOption(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftOption/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AircraftOption
	// returns the results untouched as an Observable AircraftOption
	// AircraftOption model
	// delegates via URI
	//********************************************************************
	getAircraftOption(id) : Observable<AircraftOption> {
		const uri_ = this.apiUrl + '/AircraftOption/load/' + id;

		return this.http.get<AircraftOption>(uri_);
	}
	
	//********************************************************************
	// gets all AircraftOption
	// returns the results untouched as JSON representation of an
	// Observable array of AircraftOption models
	// delegates via URI
	//********************************************************************
	getAircraftOptions() : Observable<AircraftOption[]> {
		const uri_ = this.apiUrl + '/AircraftOption/';

		return this
			.http.get<AircraftOption[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more variantsIds as a Variants
	// to a AircraftOption
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVariants( aircraftOptionId, variantsIds ): Observable<any> {

		// get the AircraftOption
		this.loadHelper( aircraftOptionId );

	// split on a comma with no spaces
	var idList = variantsIds.split(',')

	// iterate over array of variants ids
	idList.forEach(function (id) {
		// read the AircraftVariant
		var aircraftVariant = new AircraftVariantService(this.http).getAircraftVariant(id);
		// add the AircraftVariant if not already assigned
		if ( this.aircraftOption.variants.indexOf(aircraftVariant) == -1 )
		this.aircraftOption.variants.push(aircraftVariant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more variantsIds as a Variants
	// from a AircraftOption
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVariants( aircraftOptionId, variantsIds ): Observable<any> {

		// get the AircraftOption
		this.loadHelper( aircraftOptionId );


	// split on a comma with no spaces
	var idList 					= variantsIds.split(',');
	var variants 	= this.aircraftOption.variants;

	if ( variants != null && variantsIds != null ) {

		// iterate over array of variants ids
		variants.forEach(function (obj) {
			if ( variantsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftVariant
				this.aircraftOption.variants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more packagesIds as a Packages
	// to a AircraftOption
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPackages( aircraftOptionId, packagesIds ): Observable<any> {

		// get the AircraftOption
		this.loadHelper( aircraftOptionId );

	// split on a comma with no spaces
	var idList = packagesIds.split(',')

	// iterate over array of packages ids
	idList.forEach(function (id) {
		// read the AircraftPackage
		var aircraftPackage = new AircraftPackageService(this.http).getAircraftPackage(id);
		// add the AircraftPackage if not already assigned
		if ( this.aircraftOption.packages.indexOf(aircraftPackage) == -1 )
		this.aircraftOption.packages.push(aircraftPackage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more packagesIds as a Packages
	// from a AircraftOption
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePackages( aircraftOptionId, packagesIds ): Observable<any> {

		// get the AircraftOption
		this.loadHelper( aircraftOptionId );


	// split on a comma with no spaces
	var idList 					= packagesIds.split(',');
	var packages 	= this.aircraftOption.packages;

	if ( packages != null && packagesIds != null ) {

		// iterate over array of packages ids
		packages.forEach(function (obj) {
			if ( packagesIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftPackage
				this.aircraftOption.packages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AircraftOption
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AircraftOption/update/' + this.aircraftOption;

	return  this.http.post(uri_, this.aircraftOption );
}

	//********************************************************************
	// loadHelper - internal helper to load a AircraftOption
	//********************************************************************	
	loadHelper( id ) {
		this.getAircraftOption(id)
			.subscribe((res : AircraftOption) => {
				this.aircraftOption = res;
			});
	}
}