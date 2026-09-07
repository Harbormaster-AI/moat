import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {GeoRegion} from '../models/GeoRegion';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class GeoRegionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	geoRegion : GeoRegion;

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
	// add a GeoRegion
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addGeoRegion(code, name, Parent, Children, RegionType) : Observable<any> {
		const uri_ = this.apiUrl + '/GeoRegion/create';
		const obj = {
			      		code: code,
      		name: name,
      		Parent: Parent != null && Parent.length > 0 ? Parent : null,
      		Children: Children != null && Children.length > 0 ? Children : null,
			RegionType: RegionType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a GeoRegion
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateGeoRegion(code, name, Parent, Children, RegionType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/GeoRegion/update/' + id;
		const obj = {
				      		code: code,
      		name: name,
      		Parent: Parent != null && Parent.length > 0 ? Parent : null,
      		Children: Children != null && Children.length > 0 ? Children : null,
			RegionType: RegionType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a GeoRegion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteGeoRegion(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/GeoRegion/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a GeoRegion
	// returns the results untouched as an Observable GeoRegion
	// GeoRegion model
	// delegates via URI
	//********************************************************************
	getGeoRegion(id) : Observable<GeoRegion> {
		const uri_ = this.apiUrl + '/GeoRegion/load/' + id;

		return this.http.get<GeoRegion>(uri_);
	}
	
	//********************************************************************
	// gets all GeoRegion
	// returns the results untouched as JSON representation of an
	// Observable array of GeoRegion models
	// delegates via URI
	//********************************************************************
	getGeoRegions() : Observable<GeoRegion[]> {
		const uri_ = this.apiUrl + '/GeoRegion/';

		return this
			.http.get<GeoRegion[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Parent on a GeoRegion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignParent( geoRegionId, _parentId ): Observable<any> {

		// get the GeoRegion from storage
		this.loadHelper( geoRegionId );

	// get the GeoRegion from storage
	var tmp 	= new GeoRegionService(this.http).getGeoRegion(_parentId);

	// assign the Parent
	this.geoRegion.parent = tmp;

	// save the GeoRegion
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Parent on a GeoRegion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignParent( geoRegionId ): Observable<any> {

		// get the GeoRegion from storage
		this.loadHelper( geoRegionId );

	// assign Parent to null
	this.geoRegion.parent = null;

	// save the GeoRegion
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more childrenIds as a Children
	// to a GeoRegion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addChildren( geoRegionId, childrenIds ): Observable<any> {

		// get the GeoRegion
		this.loadHelper( geoRegionId );

	// split on a comma with no spaces
	var idList = childrenIds.split(',')

	// iterate over array of children ids
	idList.forEach(function (id) {
		// read the GeoRegion
		var geoRegion = new GeoRegionService(this.http).getGeoRegion(id);
		// add the GeoRegion if not already assigned
		if ( this.geoRegion.children.indexOf(geoRegion) == -1 )
		this.geoRegion.children.push(geoRegion);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more childrenIds as a Children
	// from a GeoRegion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeChildren( geoRegionId, childrenIds ): Observable<any> {

		// get the GeoRegion
		this.loadHelper( geoRegionId );


	// split on a comma with no spaces
	var idList 					= childrenIds.split(',');
	var children 	= this.geoRegion.children;

	if ( children != null && childrenIds != null ) {

		// iterate over array of children ids
		children.forEach(function (obj) {
			if ( childrenIds.indexOf(obj._id) > -1 ) {
				// remove the GeoRegion
				this.geoRegion.children.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a GeoRegion
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/GeoRegion/update/' + this.geoRegion;

	return  this.http.post(uri_, this.geoRegion );
}

	//********************************************************************
	// loadHelper - internal helper to load a GeoRegion
	//********************************************************************	
	loadHelper( id ) {
		this.getGeoRegion(id)
			.subscribe((res : GeoRegion) => {
				this.geoRegion = res;
			});
	}
}