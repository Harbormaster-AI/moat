import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AircraftProgram} from '../models/AircraftProgram';
import {AerospaceManufacturerService} from '../services/AerospaceManufacturer.service';
import {AircraftFamilyService} from '../services/AircraftFamily.service';
import {TypeCertificateService} from '../services/TypeCertificate.service';
import {SupplierService} from '../services/Supplier.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AircraftProgramService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aircraftProgram : AircraftProgram;

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
	// add a AircraftProgram
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAircraftProgram(name, programCode, entryIntoServiceYear, Manufacturer, AircraftFamilies, TypeCertificate, KeySuppliers, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftProgram/create';
		const obj = {
			      		name: name,
      		programCode: programCode,
      		entryIntoServiceYear: entryIntoServiceYear,
      		Manufacturer: Manufacturer != null && Manufacturer.length > 0 ? Manufacturer : null,
      		AircraftFamilies: AircraftFamilies != null && AircraftFamilies.length > 0 ? AircraftFamilies : null,
      		TypeCertificate: TypeCertificate != null && TypeCertificate.length > 0 ? TypeCertificate : null,
      		KeySuppliers: KeySuppliers != null && KeySuppliers.length > 0 ? KeySuppliers : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AircraftProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAircraftProgram(name, programCode, entryIntoServiceYear, Manufacturer, AircraftFamilies, TypeCertificate, KeySuppliers, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AircraftProgram/update/' + id;
		const obj = {
				      		name: name,
      		programCode: programCode,
      		entryIntoServiceYear: entryIntoServiceYear,
      		Manufacturer: Manufacturer != null && Manufacturer.length > 0 ? Manufacturer : null,
      		AircraftFamilies: AircraftFamilies != null && AircraftFamilies.length > 0 ? AircraftFamilies : null,
      		TypeCertificate: TypeCertificate != null && TypeCertificate.length > 0 ? TypeCertificate : null,
      		KeySuppliers: KeySuppliers != null && KeySuppliers.length > 0 ? KeySuppliers : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AircraftProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAircraftProgram(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftProgram/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AircraftProgram
	// returns the results untouched as an Observable AircraftProgram
	// AircraftProgram model
	// delegates via URI
	//********************************************************************
	getAircraftProgram(id) : Observable<AircraftProgram> {
		const uri_ = this.apiUrl + '/AircraftProgram/load/' + id;

		return this.http.get<AircraftProgram>(uri_);
	}
	
	//********************************************************************
	// gets all AircraftProgram
	// returns the results untouched as JSON representation of an
	// Observable array of AircraftProgram models
	// delegates via URI
	//********************************************************************
	getAircraftPrograms() : Observable<AircraftProgram[]> {
		const uri_ = this.apiUrl + '/AircraftProgram/';

		return this
			.http.get<AircraftProgram[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Manufacturer on a AircraftProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignManufacturer( aircraftProgramId, _manufacturerId ): Observable<any> {

		// get the AircraftProgram from storage
		this.loadHelper( aircraftProgramId );

	// get the AerospaceManufacturer from storage
	var tmp 	= new AerospaceManufacturerService(this.http).getAerospaceManufacturer(_manufacturerId);

	// assign the Manufacturer
	this.aircraftProgram.manufacturer = tmp;

	// save the AircraftProgram
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Manufacturer on a AircraftProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignManufacturer( aircraftProgramId ): Observable<any> {

		// get the AircraftProgram from storage
		this.loadHelper( aircraftProgramId );

	// assign Manufacturer to null
	this.aircraftProgram.manufacturer = null;

	// save the AircraftProgram
	return this.saveHelper();
}

		//********************************************************************
	// assigns a TypeCertificate on a AircraftProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTypeCertificate( aircraftProgramId, _typeCertificateId ): Observable<any> {

		// get the AircraftProgram from storage
		this.loadHelper( aircraftProgramId );

	// get the TypeCertificate from storage
	var tmp 	= new TypeCertificateService(this.http).getTypeCertificate(_typeCertificateId);

	// assign the TypeCertificate
	this.aircraftProgram.typeCertificate = tmp;

	// save the AircraftProgram
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TypeCertificate on a AircraftProgram
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTypeCertificate( aircraftProgramId ): Observable<any> {

		// get the AircraftProgram from storage
		this.loadHelper( aircraftProgramId );

	// assign TypeCertificate to null
	this.aircraftProgram.typeCertificate = null;

	// save the AircraftProgram
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more aircraftFamiliesIds as a AircraftFamilies
	// to a AircraftProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAircraftFamilies( aircraftProgramId, aircraftFamiliesIds ): Observable<any> {

		// get the AircraftProgram
		this.loadHelper( aircraftProgramId );

	// split on a comma with no spaces
	var idList = aircraftFamiliesIds.split(',')

	// iterate over array of aircraftFamilies ids
	idList.forEach(function (id) {
		// read the AircraftFamily
		var aircraftFamily = new AircraftFamilyService(this.http).getAircraftFamily(id);
		// add the AircraftFamily if not already assigned
		if ( this.aircraftProgram.aircraftFamilies.indexOf(aircraftFamily) == -1 )
		this.aircraftProgram.aircraftFamilies.push(aircraftFamily);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more aircraftFamiliesIds as a AircraftFamilies
	// from a AircraftProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAircraftFamilies( aircraftProgramId, aircraftFamiliesIds ): Observable<any> {

		// get the AircraftProgram
		this.loadHelper( aircraftProgramId );


	// split on a comma with no spaces
	var idList 					= aircraftFamiliesIds.split(',');
	var aircraftFamilies 	= this.aircraftProgram.aircraftFamilies;

	if ( aircraftFamilies != null && aircraftFamiliesIds != null ) {

		// iterate over array of aircraftFamilies ids
		aircraftFamilies.forEach(function (obj) {
			if ( aircraftFamiliesIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftFamily
				this.aircraftProgram.aircraftFamilies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more keySuppliersIds as a KeySuppliers
	// to a AircraftProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addKeySuppliers( aircraftProgramId, keySuppliersIds ): Observable<any> {

		// get the AircraftProgram
		this.loadHelper( aircraftProgramId );

	// split on a comma with no spaces
	var idList = keySuppliersIds.split(',')

	// iterate over array of keySuppliers ids
	idList.forEach(function (id) {
		// read the Supplier
		var supplier = new SupplierService(this.http).getSupplier(id);
		// add the Supplier if not already assigned
		if ( this.aircraftProgram.keySuppliers.indexOf(supplier) == -1 )
		this.aircraftProgram.keySuppliers.push(supplier);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more keySuppliersIds as a KeySuppliers
	// from a AircraftProgram
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeKeySuppliers( aircraftProgramId, keySuppliersIds ): Observable<any> {

		// get the AircraftProgram
		this.loadHelper( aircraftProgramId );


	// split on a comma with no spaces
	var idList 					= keySuppliersIds.split(',');
	var keySuppliers 	= this.aircraftProgram.keySuppliers;

	if ( keySuppliers != null && keySuppliersIds != null ) {

		// iterate over array of keySuppliers ids
		keySuppliers.forEach(function (obj) {
			if ( keySuppliersIds.indexOf(obj._id) > -1 ) {
				// remove the Supplier
				this.aircraftProgram.keySuppliers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AircraftProgram
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AircraftProgram/update/' + this.aircraftProgram;

	return  this.http.post(uri_, this.aircraftProgram );
}

	//********************************************************************
	// loadHelper - internal helper to load a AircraftProgram
	//********************************************************************	
	loadHelper( id ) {
		this.getAircraftProgram(id)
			.subscribe((res : AircraftProgram) => {
				this.aircraftProgram = res;
			});
	}
}