import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {OutboundAllocation} from '../models/OutboundAllocation';
import {WarehouseService} from '../services/Warehouse.service';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import {ReservationService} from '../services/Reservation.service';
import {LotService} from '../services/Lot.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OutboundAllocationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	outboundAllocation : OutboundAllocation;

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
	// add a OutboundAllocation
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOutboundAllocation(allocationNumber, allocatedQuantity, allocationDate, Warehouse, Sku, InventoryItem, Reservation, Lot, SerialNumbers, SourceLocation, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/OutboundAllocation/create';
		const obj = {
			      		allocationNumber: allocationNumber,
      		allocatedQuantity: allocatedQuantity,
      		allocationDate: allocationDate,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		InventoryItem: InventoryItem != null && InventoryItem.length > 0 ? InventoryItem : null,
      		Reservation: Reservation != null && Reservation.length > 0 ? Reservation : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		SourceLocation: SourceLocation != null && SourceLocation.length > 0 ? SourceLocation : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOutboundAllocation(allocationNumber, allocatedQuantity, allocationDate, Warehouse, Sku, InventoryItem, Reservation, Lot, SerialNumbers, SourceLocation, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/OutboundAllocation/update/' + id;
		const obj = {
				      		allocationNumber: allocationNumber,
      		allocatedQuantity: allocatedQuantity,
      		allocationDate: allocationDate,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		InventoryItem: InventoryItem != null && InventoryItem.length > 0 ? InventoryItem : null,
      		Reservation: Reservation != null && Reservation.length > 0 ? Reservation : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		SourceLocation: SourceLocation != null && SourceLocation.length > 0 ? SourceLocation : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOutboundAllocation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/OutboundAllocation/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a OutboundAllocation
	// returns the results untouched as an Observable OutboundAllocation
	// OutboundAllocation model
	// delegates via URI
	//********************************************************************
	getOutboundAllocation(id) : Observable<OutboundAllocation> {
		const uri_ = this.apiUrl + '/OutboundAllocation/load/' + id;

		return this.http.get<OutboundAllocation>(uri_);
	}
	
	//********************************************************************
	// gets all OutboundAllocation
	// returns the results untouched as JSON representation of an
	// Observable array of OutboundAllocation models
	// delegates via URI
	//********************************************************************
	getOutboundAllocations() : Observable<OutboundAllocation[]> {
		const uri_ = this.apiUrl + '/OutboundAllocation/';

		return this
			.http.get<OutboundAllocation[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Warehouse on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( outboundAllocationId, _warehouseId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.outboundAllocation.warehouse = tmp;

	// save the OutboundAllocation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( outboundAllocationId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// assign Warehouse to null
	this.outboundAllocation.warehouse = null;

	// save the OutboundAllocation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Sku on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( outboundAllocationId, _skuId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.outboundAllocation.sku = tmp;

	// save the OutboundAllocation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( outboundAllocationId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// assign Sku to null
	this.outboundAllocation.sku = null;

	// save the OutboundAllocation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a InventoryItem on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInventoryItem( outboundAllocationId, _inventoryItemId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// get the InventoryItem from storage
	var tmp 	= new InventoryItemService(this.http).getInventoryItem(_inventoryItemId);

	// assign the InventoryItem
	this.outboundAllocation.inventoryItem = tmp;

	// save the OutboundAllocation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InventoryItem on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInventoryItem( outboundAllocationId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// assign InventoryItem to null
	this.outboundAllocation.inventoryItem = null;

	// save the OutboundAllocation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Reservation on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignReservation( outboundAllocationId, _reservationId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// get the Reservation from storage
	var tmp 	= new ReservationService(this.http).getReservation(_reservationId);

	// assign the Reservation
	this.outboundAllocation.reservation = tmp;

	// save the OutboundAllocation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Reservation on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignReservation( outboundAllocationId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// assign Reservation to null
	this.outboundAllocation.reservation = null;

	// save the OutboundAllocation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lot on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLot( outboundAllocationId, _lotId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// get the Lot from storage
	var tmp 	= new LotService(this.http).getLot(_lotId);

	// assign the Lot
	this.outboundAllocation.lot = tmp;

	// save the OutboundAllocation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lot on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLot( outboundAllocationId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// assign Lot to null
	this.outboundAllocation.lot = null;

	// save the OutboundAllocation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a SourceLocation on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSourceLocation( outboundAllocationId, _sourceLocationId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_sourceLocationId);

	// assign the SourceLocation
	this.outboundAllocation.sourceLocation = tmp;

	// save the OutboundAllocation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SourceLocation on a OutboundAllocation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSourceLocation( outboundAllocationId ): Observable<any> {

		// get the OutboundAllocation from storage
		this.loadHelper( outboundAllocationId );

	// assign SourceLocation to null
	this.outboundAllocation.sourceLocation = null;

	// save the OutboundAllocation
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more serialNumbersIds as a SerialNumbers
	// to a OutboundAllocation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSerialNumbers( outboundAllocationId, serialNumbersIds ): Observable<any> {

		// get the OutboundAllocation
		this.loadHelper( outboundAllocationId );

	// split on a comma with no spaces
	var idList = serialNumbersIds.split(',')

	// iterate over array of serialNumbers ids
	idList.forEach(function (id) {
		// read the SerialNumber
		var serialNumber = new SerialNumberService(this.http).getSerialNumber(id);
		// add the SerialNumber if not already assigned
		if ( this.outboundAllocation.serialNumbers.indexOf(serialNumber) == -1 )
		this.outboundAllocation.serialNumbers.push(serialNumber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serialNumbersIds as a SerialNumbers
	// from a OutboundAllocation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSerialNumbers( outboundAllocationId, serialNumbersIds ): Observable<any> {

		// get the OutboundAllocation
		this.loadHelper( outboundAllocationId );


	// split on a comma with no spaces
	var idList 					= serialNumbersIds.split(',');
	var serialNumbers 	= this.outboundAllocation.serialNumbers;

	if ( serialNumbers != null && serialNumbersIds != null ) {

		// iterate over array of serialNumbers ids
		serialNumbers.forEach(function (obj) {
			if ( serialNumbersIds.indexOf(obj._id) > -1 ) {
				// remove the SerialNumber
				this.outboundAllocation.serialNumbers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a OutboundAllocation
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/OutboundAllocation/update/' + this.outboundAllocation;

	return  this.http.post(uri_, this.outboundAllocation );
}

	//********************************************************************
	// loadHelper - internal helper to load a OutboundAllocation
	//********************************************************************	
	loadHelper( id ) {
		this.getOutboundAllocation(id)
			.subscribe((res : OutboundAllocation) => {
				this.outboundAllocation = res;
			});
	}
}