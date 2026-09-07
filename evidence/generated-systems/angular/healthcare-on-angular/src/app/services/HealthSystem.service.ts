import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {HealthSystem} from '../models/HealthSystem';
import {FacilityService} from '../services/Facility.service';
import {MedicalSupplierService} from '../services/MedicalSupplier.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class HealthSystemService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	healthSystem : HealthSystem;

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
	// add a HealthSystem
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addHealthSystem(name, legalName, headquartersCountry, website, Facilities, Suppliers) : Observable<any> {
		const uri_ = this.apiUrl + '/HealthSystem/create';
		const obj = {
			      		name: name,
      		legalName: legalName,
      		headquartersCountry: headquartersCountry,
      		website: website,
      		Facilities: Facilities != null && Facilities.length > 0 ? Facilities : null,
			Suppliers: Suppliers != null && Suppliers.length > 0 ? Suppliers : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a HealthSystem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateHealthSystem(name, legalName, headquartersCountry, website, Facilities, Suppliers, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/HealthSystem/update/' + id;
		const obj = {
				      		name: name,
      		legalName: legalName,
      		headquartersCountry: headquartersCountry,
      		website: website,
      		Facilities: Facilities != null && Facilities.length > 0 ? Facilities : null,
			Suppliers: Suppliers != null && Suppliers.length > 0 ? Suppliers : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a HealthSystem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteHealthSystem(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/HealthSystem/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a HealthSystem
	// returns the results untouched as an Observable HealthSystem
	// HealthSystem model
	// delegates via URI
	//********************************************************************
	getHealthSystem(id) : Observable<HealthSystem> {
		const uri_ = this.apiUrl + '/HealthSystem/load/' + id;

		return this.http.get<HealthSystem>(uri_);
	}
	
	//********************************************************************
	// gets all HealthSystem
	// returns the results untouched as JSON representation of an
	// Observable array of HealthSystem models
	// delegates via URI
	//********************************************************************
	getHealthSystems() : Observable<HealthSystem[]> {
		const uri_ = this.apiUrl + '/HealthSystem/';

		return this
			.http.get<HealthSystem[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more facilitiesIds as a Facilities
	// to a HealthSystem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFacilities( healthSystemId, facilitiesIds ): Observable<any> {

		// get the HealthSystem
		this.loadHelper( healthSystemId );

	// split on a comma with no spaces
	var idList = facilitiesIds.split(',')

	// iterate over array of facilities ids
	idList.forEach(function (id) {
		// read the Facility
		var facility = new FacilityService(this.http).getFacility(id);
		// add the Facility if not already assigned
		if ( this.healthSystem.facilities.indexOf(facility) == -1 )
		this.healthSystem.facilities.push(facility);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more facilitiesIds as a Facilities
	// from a HealthSystem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFacilities( healthSystemId, facilitiesIds ): Observable<any> {

		// get the HealthSystem
		this.loadHelper( healthSystemId );


	// split on a comma with no spaces
	var idList 					= facilitiesIds.split(',');
	var facilities 	= this.healthSystem.facilities;

	if ( facilities != null && facilitiesIds != null ) {

		// iterate over array of facilities ids
		facilities.forEach(function (obj) {
			if ( facilitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Facility
				this.healthSystem.facilities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more suppliersIds as a Suppliers
	// to a HealthSystem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSuppliers( healthSystemId, suppliersIds ): Observable<any> {

		// get the HealthSystem
		this.loadHelper( healthSystemId );

	// split on a comma with no spaces
	var idList = suppliersIds.split(',')

	// iterate over array of suppliers ids
	idList.forEach(function (id) {
		// read the MedicalSupplier
		var medicalSupplier = new MedicalSupplierService(this.http).getMedicalSupplier(id);
		// add the MedicalSupplier if not already assigned
		if ( this.healthSystem.suppliers.indexOf(medicalSupplier) == -1 )
		this.healthSystem.suppliers.push(medicalSupplier);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more suppliersIds as a Suppliers
	// from a HealthSystem
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSuppliers( healthSystemId, suppliersIds ): Observable<any> {

		// get the HealthSystem
		this.loadHelper( healthSystemId );


	// split on a comma with no spaces
	var idList 					= suppliersIds.split(',');
	var suppliers 	= this.healthSystem.suppliers;

	if ( suppliers != null && suppliersIds != null ) {

		// iterate over array of suppliers ids
		suppliers.forEach(function (obj) {
			if ( suppliersIds.indexOf(obj._id) > -1 ) {
				// remove the MedicalSupplier
				this.healthSystem.suppliers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a HealthSystem
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/HealthSystem/update/' + this.healthSystem;

	return  this.http.post(uri_, this.healthSystem );
}

	//********************************************************************
	// loadHelper - internal helper to load a HealthSystem
	//********************************************************************	
	loadHelper( id ) {
		this.getHealthSystem(id)
			.subscribe((res : HealthSystem) => {
				this.healthSystem = res;
			});
	}
}