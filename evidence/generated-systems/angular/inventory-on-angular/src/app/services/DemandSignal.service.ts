import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DemandSignal} from '../models/DemandSignal';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {ReservationService} from '../services/Reservation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DemandSignalService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	demandSignal : DemandSignal;

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
	// add a DemandSignal
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDemandSignal(externalReference, requestedDate, quantity, Sku, Reservations, DemandType) : Observable<any> {
		const uri_ = this.apiUrl + '/DemandSignal/create';
		const obj = {
			      		externalReference: externalReference,
      		requestedDate: requestedDate,
      		quantity: quantity,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Reservations: Reservations != null && Reservations.length > 0 ? Reservations : null,
			DemandType: DemandType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DemandSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDemandSignal(externalReference, requestedDate, quantity, Sku, Reservations, DemandType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DemandSignal/update/' + id;
		const obj = {
				      		externalReference: externalReference,
      		requestedDate: requestedDate,
      		quantity: quantity,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Reservations: Reservations != null && Reservations.length > 0 ? Reservations : null,
			DemandType: DemandType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DemandSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDemandSignal(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DemandSignal/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DemandSignal
	// returns the results untouched as an Observable DemandSignal
	// DemandSignal model
	// delegates via URI
	//********************************************************************
	getDemandSignal(id) : Observable<DemandSignal> {
		const uri_ = this.apiUrl + '/DemandSignal/load/' + id;

		return this.http.get<DemandSignal>(uri_);
	}
	
	//********************************************************************
	// gets all DemandSignal
	// returns the results untouched as JSON representation of an
	// Observable array of DemandSignal models
	// delegates via URI
	//********************************************************************
	getDemandSignals() : Observable<DemandSignal[]> {
		const uri_ = this.apiUrl + '/DemandSignal/';

		return this
			.http.get<DemandSignal[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Sku on a DemandSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( demandSignalId, _skuId ): Observable<any> {

		// get the DemandSignal from storage
		this.loadHelper( demandSignalId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.demandSignal.sku = tmp;

	// save the DemandSignal
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a DemandSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( demandSignalId ): Observable<any> {

		// get the DemandSignal from storage
		this.loadHelper( demandSignalId );

	// assign Sku to null
	this.demandSignal.sku = null;

	// save the DemandSignal
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more reservationsIds as a Reservations
	// to a DemandSignal
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReservations( demandSignalId, reservationsIds ): Observable<any> {

		// get the DemandSignal
		this.loadHelper( demandSignalId );

	// split on a comma with no spaces
	var idList = reservationsIds.split(',')

	// iterate over array of reservations ids
	idList.forEach(function (id) {
		// read the Reservation
		var reservation = new ReservationService(this.http).getReservation(id);
		// add the Reservation if not already assigned
		if ( this.demandSignal.reservations.indexOf(reservation) == -1 )
		this.demandSignal.reservations.push(reservation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reservationsIds as a Reservations
	// from a DemandSignal
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReservations( demandSignalId, reservationsIds ): Observable<any> {

		// get the DemandSignal
		this.loadHelper( demandSignalId );


	// split on a comma with no spaces
	var idList 					= reservationsIds.split(',');
	var reservations 	= this.demandSignal.reservations;

	if ( reservations != null && reservationsIds != null ) {

		// iterate over array of reservations ids
		reservations.forEach(function (obj) {
			if ( reservationsIds.indexOf(obj._id) > -1 ) {
				// remove the Reservation
				this.demandSignal.reservations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DemandSignal
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DemandSignal/update/' + this.demandSignal;

	return  this.http.post(uri_, this.demandSignal );
}

	//********************************************************************
	// loadHelper - internal helper to load a DemandSignal
	//********************************************************************	
	loadHelper( id ) {
		this.getDemandSignal(id)
			.subscribe((res : DemandSignal) => {
				this.demandSignal = res;
			});
	}
}