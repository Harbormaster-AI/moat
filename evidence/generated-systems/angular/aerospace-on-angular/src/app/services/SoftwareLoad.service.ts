import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SoftwareLoad} from '../models/SoftwareLoad';
import {ConnectedAircraftService} from '../services/ConnectedAircraft.service';
import {AvionicsSuiteService} from '../services/AvionicsSuite.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SoftwareLoadService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	softwareLoad : SoftwareLoad;

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
	// add a SoftwareLoad
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSoftwareLoad(version, ConnectedAircraft, AvionicsSuite, LoadType) : Observable<any> {
		const uri_ = this.apiUrl + '/SoftwareLoad/create';
		const obj = {
			      		version: version,
      		ConnectedAircraft: ConnectedAircraft != null && ConnectedAircraft.length > 0 ? ConnectedAircraft : null,
      		AvionicsSuite: AvionicsSuite != null && AvionicsSuite.length > 0 ? AvionicsSuite : null,
			LoadType: LoadType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SoftwareLoad
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSoftwareLoad(version, ConnectedAircraft, AvionicsSuite, LoadType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SoftwareLoad/update/' + id;
		const obj = {
				      		version: version,
      		ConnectedAircraft: ConnectedAircraft != null && ConnectedAircraft.length > 0 ? ConnectedAircraft : null,
      		AvionicsSuite: AvionicsSuite != null && AvionicsSuite.length > 0 ? AvionicsSuite : null,
			LoadType: LoadType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SoftwareLoad
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSoftwareLoad(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SoftwareLoad/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SoftwareLoad
	// returns the results untouched as an Observable SoftwareLoad
	// SoftwareLoad model
	// delegates via URI
	//********************************************************************
	getSoftwareLoad(id) : Observable<SoftwareLoad> {
		const uri_ = this.apiUrl + '/SoftwareLoad/load/' + id;

		return this.http.get<SoftwareLoad>(uri_);
	}
	
	//********************************************************************
	// gets all SoftwareLoad
	// returns the results untouched as JSON representation of an
	// Observable array of SoftwareLoad models
	// delegates via URI
	//********************************************************************
	getSoftwareLoads() : Observable<SoftwareLoad[]> {
		const uri_ = this.apiUrl + '/SoftwareLoad/';

		return this
			.http.get<SoftwareLoad[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ConnectedAircraft on a SoftwareLoad
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignConnectedAircraft( softwareLoadId, _connectedAircraftId ): Observable<any> {

		// get the SoftwareLoad from storage
		this.loadHelper( softwareLoadId );

	// get the ConnectedAircraft from storage
	var tmp 	= new ConnectedAircraftService(this.http).getConnectedAircraft(_connectedAircraftId);

	// assign the ConnectedAircraft
	this.softwareLoad.connectedAircraft = tmp;

	// save the SoftwareLoad
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ConnectedAircraft on a SoftwareLoad
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignConnectedAircraft( softwareLoadId ): Observable<any> {

		// get the SoftwareLoad from storage
		this.loadHelper( softwareLoadId );

	// assign ConnectedAircraft to null
	this.softwareLoad.connectedAircraft = null;

	// save the SoftwareLoad
	return this.saveHelper();
}

		//********************************************************************
	// assigns a AvionicsSuite on a SoftwareLoad
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAvionicsSuite( softwareLoadId, _avionicsSuiteId ): Observable<any> {

		// get the SoftwareLoad from storage
		this.loadHelper( softwareLoadId );

	// get the AvionicsSuite from storage
	var tmp 	= new AvionicsSuiteService(this.http).getAvionicsSuite(_avionicsSuiteId);

	// assign the AvionicsSuite
	this.softwareLoad.avionicsSuite = tmp;

	// save the SoftwareLoad
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AvionicsSuite on a SoftwareLoad
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAvionicsSuite( softwareLoadId ): Observable<any> {

		// get the SoftwareLoad from storage
		this.loadHelper( softwareLoadId );

	// assign AvionicsSuite to null
	this.softwareLoad.avionicsSuite = null;

	// save the SoftwareLoad
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a SoftwareLoad
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SoftwareLoad/update/' + this.softwareLoad;

	return  this.http.post(uri_, this.softwareLoad );
}

	//********************************************************************
	// loadHelper - internal helper to load a SoftwareLoad
	//********************************************************************	
	loadHelper( id ) {
		this.getSoftwareLoad(id)
			.subscribe((res : SoftwareLoad) => {
				this.softwareLoad = res;
			});
	}
}