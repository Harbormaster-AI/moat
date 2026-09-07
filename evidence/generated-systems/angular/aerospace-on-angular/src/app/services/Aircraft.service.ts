import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Aircraft} from '../models/Aircraft';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import {OperatorService} from '../services/Operator.service';
import {RegistrationService} from '../services/Registration.service';
import {WarrantyService} from '../services/Warranty.service';
import {MaintenanceWorkOrderService} from '../services/MaintenanceWorkOrder.service';
import {ConnectedAircraftService} from '../services/ConnectedAircraft.service';
import {CabinLayoutService} from '../services/CabinLayout.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AircraftService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aircraft : Aircraft;

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
	// add a Aircraft
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAircraft(msn, deliveryDate, Variant, Operator, Registration, Warranty, MaintenanceRecords, ConnectedAircraft, CabinLayout) : Observable<any> {
		const uri_ = this.apiUrl + '/Aircraft/create';
		const obj = {
			      		msn: msn,
      		deliveryDate: deliveryDate,
      		Variant: Variant != null && Variant.length > 0 ? Variant : null,
      		Operator: Operator != null && Operator.length > 0 ? Operator : null,
      		Registration: Registration != null && Registration.length > 0 ? Registration : null,
      		Warranty: Warranty != null && Warranty.length > 0 ? Warranty : null,
      		MaintenanceRecords: MaintenanceRecords != null && MaintenanceRecords.length > 0 ? MaintenanceRecords : null,
      		ConnectedAircraft: ConnectedAircraft != null && ConnectedAircraft.length > 0 ? ConnectedAircraft : null,
			CabinLayout: CabinLayout != null && CabinLayout.length > 0 ? CabinLayout : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAircraft(msn, deliveryDate, Variant, Operator, Registration, Warranty, MaintenanceRecords, ConnectedAircraft, CabinLayout, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Aircraft/update/' + id;
		const obj = {
				      		msn: msn,
      		deliveryDate: deliveryDate,
      		Variant: Variant != null && Variant.length > 0 ? Variant : null,
      		Operator: Operator != null && Operator.length > 0 ? Operator : null,
      		Registration: Registration != null && Registration.length > 0 ? Registration : null,
      		Warranty: Warranty != null && Warranty.length > 0 ? Warranty : null,
      		MaintenanceRecords: MaintenanceRecords != null && MaintenanceRecords.length > 0 ? MaintenanceRecords : null,
      		ConnectedAircraft: ConnectedAircraft != null && ConnectedAircraft.length > 0 ? ConnectedAircraft : null,
			CabinLayout: CabinLayout != null && CabinLayout.length > 0 ? CabinLayout : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAircraft(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Aircraft/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Aircraft
	// returns the results untouched as an Observable Aircraft
	// Aircraft model
	// delegates via URI
	//********************************************************************
	getAircraft(id) : Observable<Aircraft> {
		const uri_ = this.apiUrl + '/Aircraft/load/' + id;

		return this.http.get<Aircraft>(uri_);
	}
	
	//********************************************************************
	// gets all Aircraft
	// returns the results untouched as JSON representation of an
	// Observable array of Aircraft models
	// delegates via URI
	//********************************************************************
	getAircrafts() : Observable<Aircraft[]> {
		const uri_ = this.apiUrl + '/Aircraft/';

		return this
			.http.get<Aircraft[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Variant on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignVariant( aircraftId, _variantId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// get the AircraftVariant from storage
	var tmp 	= new AircraftVariantService(this.http).getAircraftVariant(_variantId);

	// assign the Variant
	this.aircraft.variant = tmp;

	// save the Aircraft
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Variant on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignVariant( aircraftId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// assign Variant to null
	this.aircraft.variant = null;

	// save the Aircraft
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Operator on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOperator( aircraftId, _operatorId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// get the Operator from storage
	var tmp 	= new OperatorService(this.http).getOperator(_operatorId);

	// assign the Operator
	this.aircraft.operator = tmp;

	// save the Aircraft
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Operator on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOperator( aircraftId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// assign Operator to null
	this.aircraft.operator = null;

	// save the Aircraft
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Registration on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRegistration( aircraftId, _registrationId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// get the Registration from storage
	var tmp 	= new RegistrationService(this.http).getRegistration(_registrationId);

	// assign the Registration
	this.aircraft.registration = tmp;

	// save the Aircraft
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Registration on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRegistration( aircraftId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// assign Registration to null
	this.aircraft.registration = null;

	// save the Aircraft
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Warranty on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarranty( aircraftId, _warrantyId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// get the Warranty from storage
	var tmp 	= new WarrantyService(this.http).getWarranty(_warrantyId);

	// assign the Warranty
	this.aircraft.warranty = tmp;

	// save the Aircraft
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warranty on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarranty( aircraftId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// assign Warranty to null
	this.aircraft.warranty = null;

	// save the Aircraft
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ConnectedAircraft on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignConnectedAircraft( aircraftId, _connectedAircraftId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// get the ConnectedAircraft from storage
	var tmp 	= new ConnectedAircraftService(this.http).getConnectedAircraft(_connectedAircraftId);

	// assign the ConnectedAircraft
	this.aircraft.connectedAircraft = tmp;

	// save the Aircraft
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ConnectedAircraft on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignConnectedAircraft( aircraftId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// assign ConnectedAircraft to null
	this.aircraft.connectedAircraft = null;

	// save the Aircraft
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CabinLayout on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCabinLayout( aircraftId, _cabinLayoutId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// get the CabinLayout from storage
	var tmp 	= new CabinLayoutService(this.http).getCabinLayout(_cabinLayoutId);

	// assign the CabinLayout
	this.aircraft.cabinLayout = tmp;

	// save the Aircraft
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CabinLayout on a Aircraft
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCabinLayout( aircraftId ): Observable<any> {

		// get the Aircraft from storage
		this.loadHelper( aircraftId );

	// assign CabinLayout to null
	this.aircraft.cabinLayout = null;

	// save the Aircraft
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more maintenanceRecordsIds as a MaintenanceRecords
	// to a Aircraft
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMaintenanceRecords( aircraftId, maintenanceRecordsIds ): Observable<any> {

		// get the Aircraft
		this.loadHelper( aircraftId );

	// split on a comma with no spaces
	var idList = maintenanceRecordsIds.split(',')

	// iterate over array of maintenanceRecords ids
	idList.forEach(function (id) {
		// read the MaintenanceWorkOrder
		var maintenanceWorkOrder = new MaintenanceWorkOrderService(this.http).getMaintenanceWorkOrder(id);
		// add the MaintenanceWorkOrder if not already assigned
		if ( this.aircraft.maintenanceRecords.indexOf(maintenanceWorkOrder) == -1 )
		this.aircraft.maintenanceRecords.push(maintenanceWorkOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more maintenanceRecordsIds as a MaintenanceRecords
	// from a Aircraft
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMaintenanceRecords( aircraftId, maintenanceRecordsIds ): Observable<any> {

		// get the Aircraft
		this.loadHelper( aircraftId );


	// split on a comma with no spaces
	var idList 					= maintenanceRecordsIds.split(',');
	var maintenanceRecords 	= this.aircraft.maintenanceRecords;

	if ( maintenanceRecords != null && maintenanceRecordsIds != null ) {

		// iterate over array of maintenanceRecords ids
		maintenanceRecords.forEach(function (obj) {
			if ( maintenanceRecordsIds.indexOf(obj._id) > -1 ) {
				// remove the MaintenanceWorkOrder
				this.aircraft.maintenanceRecords.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Aircraft
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Aircraft/update/' + this.aircraft;

	return  this.http.post(uri_, this.aircraft );
}

	//********************************************************************
	// loadHelper - internal helper to load a Aircraft
	//********************************************************************	
	loadHelper( id ) {
		this.getAircraft(id)
			.subscribe((res : Aircraft) => {
				this.aircraft = res;
			});
	}
}