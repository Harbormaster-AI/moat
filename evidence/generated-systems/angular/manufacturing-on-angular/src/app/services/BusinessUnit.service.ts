import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BusinessUnit} from '../models/BusinessUnit';
import {EnterpriseService} from '../services/Enterprise.service';
import {ItemService} from '../services/Item.service';
import {PlantService} from '../services/Plant.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BusinessUnitService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	businessUnit : BusinessUnit;

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
	// add a BusinessUnit
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBusinessUnit(name, code, Enterprise, Items, Plants, Category) : Observable<any> {
		const uri_ = this.apiUrl + '/BusinessUnit/create';
		const obj = {
			      		name: name,
      		code: code,
      		Enterprise: Enterprise != null && Enterprise.length > 0 ? Enterprise : null,
      		Items: Items != null && Items.length > 0 ? Items : null,
      		Plants: Plants != null && Plants.length > 0 ? Plants : null,
			Category: Category
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BusinessUnit
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBusinessUnit(name, code, Enterprise, Items, Plants, Category, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BusinessUnit/update/' + id;
		const obj = {
				      		name: name,
      		code: code,
      		Enterprise: Enterprise != null && Enterprise.length > 0 ? Enterprise : null,
      		Items: Items != null && Items.length > 0 ? Items : null,
      		Plants: Plants != null && Plants.length > 0 ? Plants : null,
			Category: Category
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BusinessUnit
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBusinessUnit(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BusinessUnit/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BusinessUnit
	// returns the results untouched as an Observable BusinessUnit
	// BusinessUnit model
	// delegates via URI
	//********************************************************************
	getBusinessUnit(id) : Observable<BusinessUnit> {
		const uri_ = this.apiUrl + '/BusinessUnit/load/' + id;

		return this.http.get<BusinessUnit>(uri_);
	}
	
	//********************************************************************
	// gets all BusinessUnit
	// returns the results untouched as JSON representation of an
	// Observable array of BusinessUnit models
	// delegates via URI
	//********************************************************************
	getBusinessUnits() : Observable<BusinessUnit[]> {
		const uri_ = this.apiUrl + '/BusinessUnit/';

		return this
			.http.get<BusinessUnit[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Enterprise on a BusinessUnit
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEnterprise( businessUnitId, _enterpriseId ): Observable<any> {

		// get the BusinessUnit from storage
		this.loadHelper( businessUnitId );

	// get the Enterprise from storage
	var tmp 	= new EnterpriseService(this.http).getEnterprise(_enterpriseId);

	// assign the Enterprise
	this.businessUnit.enterprise = tmp;

	// save the BusinessUnit
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Enterprise on a BusinessUnit
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEnterprise( businessUnitId ): Observable<any> {

		// get the BusinessUnit from storage
		this.loadHelper( businessUnitId );

	// assign Enterprise to null
	this.businessUnit.enterprise = null;

	// save the BusinessUnit
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more itemsIds as a Items
	// to a BusinessUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addItems( businessUnitId, itemsIds ): Observable<any> {

		// get the BusinessUnit
		this.loadHelper( businessUnitId );

	// split on a comma with no spaces
	var idList = itemsIds.split(',')

	// iterate over array of items ids
	idList.forEach(function (id) {
		// read the Item
		var item = new ItemService(this.http).getItem(id);
		// add the Item if not already assigned
		if ( this.businessUnit.items.indexOf(item) == -1 )
		this.businessUnit.items.push(item);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more itemsIds as a Items
	// from a BusinessUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeItems( businessUnitId, itemsIds ): Observable<any> {

		// get the BusinessUnit
		this.loadHelper( businessUnitId );


	// split on a comma with no spaces
	var idList 					= itemsIds.split(',');
	var items 	= this.businessUnit.items;

	if ( items != null && itemsIds != null ) {

		// iterate over array of items ids
		items.forEach(function (obj) {
			if ( itemsIds.indexOf(obj._id) > -1 ) {
				// remove the Item
				this.businessUnit.items.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more plantsIds as a Plants
	// to a BusinessUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPlants( businessUnitId, plantsIds ): Observable<any> {

		// get the BusinessUnit
		this.loadHelper( businessUnitId );

	// split on a comma with no spaces
	var idList = plantsIds.split(',')

	// iterate over array of plants ids
	idList.forEach(function (id) {
		// read the Plant
		var plant = new PlantService(this.http).getPlant(id);
		// add the Plant if not already assigned
		if ( this.businessUnit.plants.indexOf(plant) == -1 )
		this.businessUnit.plants.push(plant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more plantsIds as a Plants
	// from a BusinessUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePlants( businessUnitId, plantsIds ): Observable<any> {

		// get the BusinessUnit
		this.loadHelper( businessUnitId );


	// split on a comma with no spaces
	var idList 					= plantsIds.split(',');
	var plants 	= this.businessUnit.plants;

	if ( plants != null && plantsIds != null ) {

		// iterate over array of plants ids
		plants.forEach(function (obj) {
			if ( plantsIds.indexOf(obj._id) > -1 ) {
				// remove the Plant
				this.businessUnit.plants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BusinessUnit
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BusinessUnit/update/' + this.businessUnit;

	return  this.http.post(uri_, this.businessUnit );
}

	//********************************************************************
	// loadHelper - internal helper to load a BusinessUnit
	//********************************************************************	
	loadHelper( id ) {
		this.getBusinessUnit(id)
			.subscribe((res : BusinessUnit) => {
				this.businessUnit = res;
			});
	}
}