import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {StockKeepingUnit} from '../models/StockKeepingUnit';
import {InventoryItemService} from '../services/InventoryItem.service';
import {UoMConversionService} from '../services/UoMConversion.service';
import {ReplenishmentPolicyService} from '../services/ReplenishmentPolicy.service';
import {LotService} from '../services/Lot.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class StockKeepingUnitService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	stockKeepingUnit : StockKeepingUnit;

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
	// add a StockKeepingUnit
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addStockKeepingUnit(skuCode, name, weight, weightUnit, volume, volumeUnit, shelfLifeDays, hazardousMaterial, InventoryItems, UomConversions, ReplenishmentPolicies, Lots, SerialNumbers, ItemType, UnitOfMeasure) : Observable<any> {
		const uri_ = this.apiUrl + '/StockKeepingUnit/create';
		const obj = {
			      		skuCode: skuCode,
      		name: name,
      		weight: weight,
      		weightUnit: weightUnit,
      		volume: volume,
      		volumeUnit: volumeUnit,
      		shelfLifeDays: shelfLifeDays,
      		hazardousMaterial: hazardousMaterial,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
      		UomConversions: UomConversions != null && UomConversions.length > 0 ? UomConversions : null,
      		ReplenishmentPolicies: ReplenishmentPolicies != null && ReplenishmentPolicies.length > 0 ? ReplenishmentPolicies : null,
      		Lots: Lots != null && Lots.length > 0 ? Lots : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		ItemType: ItemType,
			UnitOfMeasure: UnitOfMeasure
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a StockKeepingUnit
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateStockKeepingUnit(skuCode, name, weight, weightUnit, volume, volumeUnit, shelfLifeDays, hazardousMaterial, InventoryItems, UomConversions, ReplenishmentPolicies, Lots, SerialNumbers, ItemType, UnitOfMeasure, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/StockKeepingUnit/update/' + id;
		const obj = {
				      		skuCode: skuCode,
      		name: name,
      		weight: weight,
      		weightUnit: weightUnit,
      		volume: volume,
      		volumeUnit: volumeUnit,
      		shelfLifeDays: shelfLifeDays,
      		hazardousMaterial: hazardousMaterial,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
      		UomConversions: UomConversions != null && UomConversions.length > 0 ? UomConversions : null,
      		ReplenishmentPolicies: ReplenishmentPolicies != null && ReplenishmentPolicies.length > 0 ? ReplenishmentPolicies : null,
      		Lots: Lots != null && Lots.length > 0 ? Lots : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		ItemType: ItemType,
			UnitOfMeasure: UnitOfMeasure
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a StockKeepingUnit
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteStockKeepingUnit(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/StockKeepingUnit/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a StockKeepingUnit
	// returns the results untouched as an Observable StockKeepingUnit
	// StockKeepingUnit model
	// delegates via URI
	//********************************************************************
	getStockKeepingUnit(id) : Observable<StockKeepingUnit> {
		const uri_ = this.apiUrl + '/StockKeepingUnit/load/' + id;

		return this.http.get<StockKeepingUnit>(uri_);
	}
	
	//********************************************************************
	// gets all StockKeepingUnit
	// returns the results untouched as JSON representation of an
	// Observable array of StockKeepingUnit models
	// delegates via URI
	//********************************************************************
	getStockKeepingUnits() : Observable<StockKeepingUnit[]> {
		const uri_ = this.apiUrl + '/StockKeepingUnit/';

		return this
			.http.get<StockKeepingUnit[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more inventoryItemsIds as a InventoryItems
	// to a StockKeepingUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventoryItems( stockKeepingUnitId, inventoryItemsIds ): Observable<any> {

		// get the StockKeepingUnit
		this.loadHelper( stockKeepingUnitId );

	// split on a comma with no spaces
	var idList = inventoryItemsIds.split(',')

	// iterate over array of inventoryItems ids
	idList.forEach(function (id) {
		// read the InventoryItem
		var inventoryItem = new InventoryItemService(this.http).getInventoryItem(id);
		// add the InventoryItem if not already assigned
		if ( this.stockKeepingUnit.inventoryItems.indexOf(inventoryItem) == -1 )
		this.stockKeepingUnit.inventoryItems.push(inventoryItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventoryItemsIds as a InventoryItems
	// from a StockKeepingUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventoryItems( stockKeepingUnitId, inventoryItemsIds ): Observable<any> {

		// get the StockKeepingUnit
		this.loadHelper( stockKeepingUnitId );


	// split on a comma with no spaces
	var idList 					= inventoryItemsIds.split(',');
	var inventoryItems 	= this.stockKeepingUnit.inventoryItems;

	if ( inventoryItems != null && inventoryItemsIds != null ) {

		// iterate over array of inventoryItems ids
		inventoryItems.forEach(function (obj) {
			if ( inventoryItemsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryItem
				this.stockKeepingUnit.inventoryItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more uomConversionsIds as a UomConversions
	// to a StockKeepingUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addUomConversions( stockKeepingUnitId, uomConversionsIds ): Observable<any> {

		// get the StockKeepingUnit
		this.loadHelper( stockKeepingUnitId );

	// split on a comma with no spaces
	var idList = uomConversionsIds.split(',')

	// iterate over array of uomConversions ids
	idList.forEach(function (id) {
		// read the UoMConversion
		var uoMConversion = new UoMConversionService(this.http).getUoMConversion(id);
		// add the UoMConversion if not already assigned
		if ( this.stockKeepingUnit.uomConversions.indexOf(uoMConversion) == -1 )
		this.stockKeepingUnit.uomConversions.push(uoMConversion);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more uomConversionsIds as a UomConversions
	// from a StockKeepingUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeUomConversions( stockKeepingUnitId, uomConversionsIds ): Observable<any> {

		// get the StockKeepingUnit
		this.loadHelper( stockKeepingUnitId );


	// split on a comma with no spaces
	var idList 					= uomConversionsIds.split(',');
	var uomConversions 	= this.stockKeepingUnit.uomConversions;

	if ( uomConversions != null && uomConversionsIds != null ) {

		// iterate over array of uomConversions ids
		uomConversions.forEach(function (obj) {
			if ( uomConversionsIds.indexOf(obj._id) > -1 ) {
				// remove the UoMConversion
				this.stockKeepingUnit.uomConversions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more replenishmentPoliciesIds as a ReplenishmentPolicies
	// to a StockKeepingUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReplenishmentPolicies( stockKeepingUnitId, replenishmentPoliciesIds ): Observable<any> {

		// get the StockKeepingUnit
		this.loadHelper( stockKeepingUnitId );

	// split on a comma with no spaces
	var idList = replenishmentPoliciesIds.split(',')

	// iterate over array of replenishmentPolicies ids
	idList.forEach(function (id) {
		// read the ReplenishmentPolicy
		var replenishmentPolicy = new ReplenishmentPolicyService(this.http).getReplenishmentPolicy(id);
		// add the ReplenishmentPolicy if not already assigned
		if ( this.stockKeepingUnit.replenishmentPolicies.indexOf(replenishmentPolicy) == -1 )
		this.stockKeepingUnit.replenishmentPolicies.push(replenishmentPolicy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more replenishmentPoliciesIds as a ReplenishmentPolicies
	// from a StockKeepingUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReplenishmentPolicies( stockKeepingUnitId, replenishmentPoliciesIds ): Observable<any> {

		// get the StockKeepingUnit
		this.loadHelper( stockKeepingUnitId );


	// split on a comma with no spaces
	var idList 					= replenishmentPoliciesIds.split(',');
	var replenishmentPolicies 	= this.stockKeepingUnit.replenishmentPolicies;

	if ( replenishmentPolicies != null && replenishmentPoliciesIds != null ) {

		// iterate over array of replenishmentPolicies ids
		replenishmentPolicies.forEach(function (obj) {
			if ( replenishmentPoliciesIds.indexOf(obj._id) > -1 ) {
				// remove the ReplenishmentPolicy
				this.stockKeepingUnit.replenishmentPolicies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more lotsIds as a Lots
	// to a StockKeepingUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLots( stockKeepingUnitId, lotsIds ): Observable<any> {

		// get the StockKeepingUnit
		this.loadHelper( stockKeepingUnitId );

	// split on a comma with no spaces
	var idList = lotsIds.split(',')

	// iterate over array of lots ids
	idList.forEach(function (id) {
		// read the Lot
		var lot = new LotService(this.http).getLot(id);
		// add the Lot if not already assigned
		if ( this.stockKeepingUnit.lots.indexOf(lot) == -1 )
		this.stockKeepingUnit.lots.push(lot);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more lotsIds as a Lots
	// from a StockKeepingUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLots( stockKeepingUnitId, lotsIds ): Observable<any> {

		// get the StockKeepingUnit
		this.loadHelper( stockKeepingUnitId );


	// split on a comma with no spaces
	var idList 					= lotsIds.split(',');
	var lots 	= this.stockKeepingUnit.lots;

	if ( lots != null && lotsIds != null ) {

		// iterate over array of lots ids
		lots.forEach(function (obj) {
			if ( lotsIds.indexOf(obj._id) > -1 ) {
				// remove the Lot
				this.stockKeepingUnit.lots.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more serialNumbersIds as a SerialNumbers
	// to a StockKeepingUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSerialNumbers( stockKeepingUnitId, serialNumbersIds ): Observable<any> {

		// get the StockKeepingUnit
		this.loadHelper( stockKeepingUnitId );

	// split on a comma with no spaces
	var idList = serialNumbersIds.split(',')

	// iterate over array of serialNumbers ids
	idList.forEach(function (id) {
		// read the SerialNumber
		var serialNumber = new SerialNumberService(this.http).getSerialNumber(id);
		// add the SerialNumber if not already assigned
		if ( this.stockKeepingUnit.serialNumbers.indexOf(serialNumber) == -1 )
		this.stockKeepingUnit.serialNumbers.push(serialNumber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serialNumbersIds as a SerialNumbers
	// from a StockKeepingUnit
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSerialNumbers( stockKeepingUnitId, serialNumbersIds ): Observable<any> {

		// get the StockKeepingUnit
		this.loadHelper( stockKeepingUnitId );


	// split on a comma with no spaces
	var idList 					= serialNumbersIds.split(',');
	var serialNumbers 	= this.stockKeepingUnit.serialNumbers;

	if ( serialNumbers != null && serialNumbersIds != null ) {

		// iterate over array of serialNumbers ids
		serialNumbers.forEach(function (obj) {
			if ( serialNumbersIds.indexOf(obj._id) > -1 ) {
				// remove the SerialNumber
				this.stockKeepingUnit.serialNumbers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a StockKeepingUnit
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/StockKeepingUnit/update/' + this.stockKeepingUnit;

	return  this.http.post(uri_, this.stockKeepingUnit );
}

	//********************************************************************
	// loadHelper - internal helper to load a StockKeepingUnit
	//********************************************************************	
	loadHelper( id ) {
		this.getStockKeepingUnit(id)
			.subscribe((res : StockKeepingUnit) => {
				this.stockKeepingUnit = res;
			});
	}
}