import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Facility} from '../models/Facility';
import {HealthSystemService} from '../services/HealthSystem.service';
import {DepartmentService} from '../services/Department.service';
import {CareTeamService} from '../services/CareTeam.service';
import {LaboratoryService} from '../services/Laboratory.service';
import {ImagingCenterService} from '../services/ImagingCenter.service';
import {PharmacyService} from '../services/Pharmacy.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class FacilityService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	facility : Facility;

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
	// add a Facility
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addFacility(name, facilityCode, address, HealthSystem, Departments, CareTeams, Laboratories, ImagingCenters, Pharmacies, InventoryItems, FacilityType) : Observable<any> {
		const uri_ = this.apiUrl + '/Facility/create';
		const obj = {
			      		name: name,
      		facilityCode: facilityCode,
      		address: address,
      		HealthSystem: HealthSystem != null && HealthSystem.length > 0 ? HealthSystem : null,
      		Departments: Departments != null && Departments.length > 0 ? Departments : null,
      		CareTeams: CareTeams != null && CareTeams.length > 0 ? CareTeams : null,
      		Laboratories: Laboratories != null && Laboratories.length > 0 ? Laboratories : null,
      		ImagingCenters: ImagingCenters != null && ImagingCenters.length > 0 ? ImagingCenters : null,
      		Pharmacies: Pharmacies != null && Pharmacies.length > 0 ? Pharmacies : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			FacilityType: FacilityType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Facility
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateFacility(name, facilityCode, address, HealthSystem, Departments, CareTeams, Laboratories, ImagingCenters, Pharmacies, InventoryItems, FacilityType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Facility/update/' + id;
		const obj = {
				      		name: name,
      		facilityCode: facilityCode,
      		address: address,
      		HealthSystem: HealthSystem != null && HealthSystem.length > 0 ? HealthSystem : null,
      		Departments: Departments != null && Departments.length > 0 ? Departments : null,
      		CareTeams: CareTeams != null && CareTeams.length > 0 ? CareTeams : null,
      		Laboratories: Laboratories != null && Laboratories.length > 0 ? Laboratories : null,
      		ImagingCenters: ImagingCenters != null && ImagingCenters.length > 0 ? ImagingCenters : null,
      		Pharmacies: Pharmacies != null && Pharmacies.length > 0 ? Pharmacies : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			FacilityType: FacilityType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Facility
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteFacility(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Facility/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Facility
	// returns the results untouched as an Observable Facility
	// Facility model
	// delegates via URI
	//********************************************************************
	getFacility(id) : Observable<Facility> {
		const uri_ = this.apiUrl + '/Facility/load/' + id;

		return this.http.get<Facility>(uri_);
	}
	
	//********************************************************************
	// gets all Facility
	// returns the results untouched as JSON representation of an
	// Observable array of Facility models
	// delegates via URI
	//********************************************************************
	getFacilitys() : Observable<Facility[]> {
		const uri_ = this.apiUrl + '/Facility/';

		return this
			.http.get<Facility[]>(uri_);
	}
	
			//********************************************************************
	// assigns a HealthSystem on a Facility
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignHealthSystem( facilityId, _healthSystemId ): Observable<any> {

		// get the Facility from storage
		this.loadHelper( facilityId );

	// get the HealthSystem from storage
	var tmp 	= new HealthSystemService(this.http).getHealthSystem(_healthSystemId);

	// assign the HealthSystem
	this.facility.healthSystem = tmp;

	// save the Facility
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a HealthSystem on a Facility
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignHealthSystem( facilityId ): Observable<any> {

		// get the Facility from storage
		this.loadHelper( facilityId );

	// assign HealthSystem to null
	this.facility.healthSystem = null;

	// save the Facility
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more departmentsIds as a Departments
	// to a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDepartments( facilityId, departmentsIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );

	// split on a comma with no spaces
	var idList = departmentsIds.split(',')

	// iterate over array of departments ids
	idList.forEach(function (id) {
		// read the Department
		var department = new DepartmentService(this.http).getDepartment(id);
		// add the Department if not already assigned
		if ( this.facility.departments.indexOf(department) == -1 )
		this.facility.departments.push(department);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more departmentsIds as a Departments
	// from a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDepartments( facilityId, departmentsIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );


	// split on a comma with no spaces
	var idList 					= departmentsIds.split(',');
	var departments 	= this.facility.departments;

	if ( departments != null && departmentsIds != null ) {

		// iterate over array of departments ids
		departments.forEach(function (obj) {
			if ( departmentsIds.indexOf(obj._id) > -1 ) {
				// remove the Department
				this.facility.departments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more careTeamsIds as a CareTeams
	// to a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCareTeams( facilityId, careTeamsIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );

	// split on a comma with no spaces
	var idList = careTeamsIds.split(',')

	// iterate over array of careTeams ids
	idList.forEach(function (id) {
		// read the CareTeam
		var careTeam = new CareTeamService(this.http).getCareTeam(id);
		// add the CareTeam if not already assigned
		if ( this.facility.careTeams.indexOf(careTeam) == -1 )
		this.facility.careTeams.push(careTeam);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more careTeamsIds as a CareTeams
	// from a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCareTeams( facilityId, careTeamsIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );


	// split on a comma with no spaces
	var idList 					= careTeamsIds.split(',');
	var careTeams 	= this.facility.careTeams;

	if ( careTeams != null && careTeamsIds != null ) {

		// iterate over array of careTeams ids
		careTeams.forEach(function (obj) {
			if ( careTeamsIds.indexOf(obj._id) > -1 ) {
				// remove the CareTeam
				this.facility.careTeams.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more laboratoriesIds as a Laboratories
	// to a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLaboratories( facilityId, laboratoriesIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );

	// split on a comma with no spaces
	var idList = laboratoriesIds.split(',')

	// iterate over array of laboratories ids
	idList.forEach(function (id) {
		// read the Laboratory
		var laboratory = new LaboratoryService(this.http).getLaboratory(id);
		// add the Laboratory if not already assigned
		if ( this.facility.laboratories.indexOf(laboratory) == -1 )
		this.facility.laboratories.push(laboratory);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more laboratoriesIds as a Laboratories
	// from a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLaboratories( facilityId, laboratoriesIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );


	// split on a comma with no spaces
	var idList 					= laboratoriesIds.split(',');
	var laboratories 	= this.facility.laboratories;

	if ( laboratories != null && laboratoriesIds != null ) {

		// iterate over array of laboratories ids
		laboratories.forEach(function (obj) {
			if ( laboratoriesIds.indexOf(obj._id) > -1 ) {
				// remove the Laboratory
				this.facility.laboratories.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more imagingCentersIds as a ImagingCenters
	// to a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addImagingCenters( facilityId, imagingCentersIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );

	// split on a comma with no spaces
	var idList = imagingCentersIds.split(',')

	// iterate over array of imagingCenters ids
	idList.forEach(function (id) {
		// read the ImagingCenter
		var imagingCenter = new ImagingCenterService(this.http).getImagingCenter(id);
		// add the ImagingCenter if not already assigned
		if ( this.facility.imagingCenters.indexOf(imagingCenter) == -1 )
		this.facility.imagingCenters.push(imagingCenter);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more imagingCentersIds as a ImagingCenters
	// from a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeImagingCenters( facilityId, imagingCentersIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );


	// split on a comma with no spaces
	var idList 					= imagingCentersIds.split(',');
	var imagingCenters 	= this.facility.imagingCenters;

	if ( imagingCenters != null && imagingCentersIds != null ) {

		// iterate over array of imagingCenters ids
		imagingCenters.forEach(function (obj) {
			if ( imagingCentersIds.indexOf(obj._id) > -1 ) {
				// remove the ImagingCenter
				this.facility.imagingCenters.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more pharmaciesIds as a Pharmacies
	// to a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPharmacies( facilityId, pharmaciesIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );

	// split on a comma with no spaces
	var idList = pharmaciesIds.split(',')

	// iterate over array of pharmacies ids
	idList.forEach(function (id) {
		// read the Pharmacy
		var pharmacy = new PharmacyService(this.http).getPharmacy(id);
		// add the Pharmacy if not already assigned
		if ( this.facility.pharmacies.indexOf(pharmacy) == -1 )
		this.facility.pharmacies.push(pharmacy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more pharmaciesIds as a Pharmacies
	// from a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePharmacies( facilityId, pharmaciesIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );


	// split on a comma with no spaces
	var idList 					= pharmaciesIds.split(',');
	var pharmacies 	= this.facility.pharmacies;

	if ( pharmacies != null && pharmaciesIds != null ) {

		// iterate over array of pharmacies ids
		pharmacies.forEach(function (obj) {
			if ( pharmaciesIds.indexOf(obj._id) > -1 ) {
				// remove the Pharmacy
				this.facility.pharmacies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more inventoryItemsIds as a InventoryItems
	// to a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventoryItems( facilityId, inventoryItemsIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );

	// split on a comma with no spaces
	var idList = inventoryItemsIds.split(',')

	// iterate over array of inventoryItems ids
	idList.forEach(function (id) {
		// read the InventoryItem
		var inventoryItem = new InventoryItemService(this.http).getInventoryItem(id);
		// add the InventoryItem if not already assigned
		if ( this.facility.inventoryItems.indexOf(inventoryItem) == -1 )
		this.facility.inventoryItems.push(inventoryItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventoryItemsIds as a InventoryItems
	// from a Facility
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventoryItems( facilityId, inventoryItemsIds ): Observable<any> {

		// get the Facility
		this.loadHelper( facilityId );


	// split on a comma with no spaces
	var idList 					= inventoryItemsIds.split(',');
	var inventoryItems 	= this.facility.inventoryItems;

	if ( inventoryItems != null && inventoryItemsIds != null ) {

		// iterate over array of inventoryItems ids
		inventoryItems.forEach(function (obj) {
			if ( inventoryItemsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryItem
				this.facility.inventoryItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Facility
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Facility/update/' + this.facility;

	return  this.http.post(uri_, this.facility );
}

	//********************************************************************
	// loadHelper - internal helper to load a Facility
	//********************************************************************	
	loadHelper( id ) {
		this.getFacility(id)
			.subscribe((res : Facility) => {
				this.facility = res;
			});
	}
}