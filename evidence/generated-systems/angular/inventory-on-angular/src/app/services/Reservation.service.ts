import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Reservation} from '../models/Reservation';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {WarehouseService} from '../services/Warehouse.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import {LotService} from '../services/Lot.service';
import {SerialNumberService} from '../services/SerialNumber.service';
import {DemandSignalService} from '../services/DemandSignal.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ReservationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	reservation : Reservation;

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
	// add a Reservation
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addReservation(referenceNumber, reservedQuantity, promisedDate, Sku, Warehouse, Location, InventoryItem, Lot, SerialNumbers, DemandSignal, ReservationStatus, ReservationType) : Observable<any> {
		const uri_ = this.apiUrl + '/Reservation/create';
		const obj = {
			      		referenceNumber: referenceNumber,
      		reservedQuantity: reservedQuantity,
      		promisedDate: promisedDate,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		InventoryItem: InventoryItem != null && InventoryItem.length > 0 ? InventoryItem : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		DemandSignal: DemandSignal != null && DemandSignal.length > 0 ? DemandSignal : null,
      		ReservationStatus: ReservationStatus,
			ReservationType: ReservationType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateReservation(referenceNumber, reservedQuantity, promisedDate, Sku, Warehouse, Location, InventoryItem, Lot, SerialNumbers, DemandSignal, ReservationStatus, ReservationType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Reservation/update/' + id;
		const obj = {
				      		referenceNumber: referenceNumber,
      		reservedQuantity: reservedQuantity,
      		promisedDate: promisedDate,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		InventoryItem: InventoryItem != null && InventoryItem.length > 0 ? InventoryItem : null,
      		Lot: Lot != null && Lot.length > 0 ? Lot : null,
      		SerialNumbers: SerialNumbers != null && SerialNumbers.length > 0 ? SerialNumbers : null,
      		DemandSignal: DemandSignal != null && DemandSignal.length > 0 ? DemandSignal : null,
      		ReservationStatus: ReservationStatus,
			ReservationType: ReservationType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteReservation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Reservation/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Reservation
	// returns the results untouched as an Observable Reservation
	// Reservation model
	// delegates via URI
	//********************************************************************
	getReservation(id) : Observable<Reservation> {
		const uri_ = this.apiUrl + '/Reservation/load/' + id;

		return this.http.get<Reservation>(uri_);
	}
	
	//********************************************************************
	// gets all Reservation
	// returns the results untouched as JSON representation of an
	// Observable array of Reservation models
	// delegates via URI
	//********************************************************************
	getReservations() : Observable<Reservation[]> {
		const uri_ = this.apiUrl + '/Reservation/';

		return this
			.http.get<Reservation[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Sku on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( reservationId, _skuId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.reservation.sku = tmp;

	// save the Reservation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( reservationId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// assign Sku to null
	this.reservation.sku = null;

	// save the Reservation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Warehouse on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( reservationId, _warehouseId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.reservation.warehouse = tmp;

	// save the Reservation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( reservationId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// assign Warehouse to null
	this.reservation.warehouse = null;

	// save the Reservation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Location on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLocation( reservationId, _locationId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_locationId);

	// assign the Location
	this.reservation.location = tmp;

	// save the Reservation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Location on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLocation( reservationId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// assign Location to null
	this.reservation.location = null;

	// save the Reservation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a InventoryItem on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInventoryItem( reservationId, _inventoryItemId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// get the InventoryItem from storage
	var tmp 	= new InventoryItemService(this.http).getInventoryItem(_inventoryItemId);

	// assign the InventoryItem
	this.reservation.inventoryItem = tmp;

	// save the Reservation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InventoryItem on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInventoryItem( reservationId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// assign InventoryItem to null
	this.reservation.inventoryItem = null;

	// save the Reservation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lot on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLot( reservationId, _lotId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// get the Lot from storage
	var tmp 	= new LotService(this.http).getLot(_lotId);

	// assign the Lot
	this.reservation.lot = tmp;

	// save the Reservation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lot on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLot( reservationId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// assign Lot to null
	this.reservation.lot = null;

	// save the Reservation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a DemandSignal on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDemandSignal( reservationId, _demandSignalId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// get the DemandSignal from storage
	var tmp 	= new DemandSignalService(this.http).getDemandSignal(_demandSignalId);

	// assign the DemandSignal
	this.reservation.demandSignal = tmp;

	// save the Reservation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a DemandSignal on a Reservation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDemandSignal( reservationId ): Observable<any> {

		// get the Reservation from storage
		this.loadHelper( reservationId );

	// assign DemandSignal to null
	this.reservation.demandSignal = null;

	// save the Reservation
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more serialNumbersIds as a SerialNumbers
	// to a Reservation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSerialNumbers( reservationId, serialNumbersIds ): Observable<any> {

		// get the Reservation
		this.loadHelper( reservationId );

	// split on a comma with no spaces
	var idList = serialNumbersIds.split(',')

	// iterate over array of serialNumbers ids
	idList.forEach(function (id) {
		// read the SerialNumber
		var serialNumber = new SerialNumberService(this.http).getSerialNumber(id);
		// add the SerialNumber if not already assigned
		if ( this.reservation.serialNumbers.indexOf(serialNumber) == -1 )
		this.reservation.serialNumbers.push(serialNumber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serialNumbersIds as a SerialNumbers
	// from a Reservation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSerialNumbers( reservationId, serialNumbersIds ): Observable<any> {

		// get the Reservation
		this.loadHelper( reservationId );


	// split on a comma with no spaces
	var idList 					= serialNumbersIds.split(',');
	var serialNumbers 	= this.reservation.serialNumbers;

	if ( serialNumbers != null && serialNumbersIds != null ) {

		// iterate over array of serialNumbers ids
		serialNumbers.forEach(function (obj) {
			if ( serialNumbersIds.indexOf(obj._id) > -1 ) {
				// remove the SerialNumber
				this.reservation.serialNumbers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Reservation
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Reservation/update/' + this.reservation;

	return  this.http.post(uri_, this.reservation );
}

	//********************************************************************
	// loadHelper - internal helper to load a Reservation
	//********************************************************************	
	loadHelper( id ) {
		this.getReservation(id)
			.subscribe((res : Reservation) => {
				this.reservation = res;
			});
	}
}