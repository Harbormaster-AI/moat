import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BOMItem} from '../models/BOMItem';
import {BOMService} from '../services/BOM.service';
import {ItemService} from '../services/Item.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BOMItemService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	bOMItem : BOMItem;

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
	// add a BOMItem
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBOMItem(lineNumber, quantity, scrapPercent, Bom, Component) : Observable<any> {
		const uri_ = this.apiUrl + '/BOMItem/create';
		const obj = {
			      		lineNumber: lineNumber,
      		quantity: quantity,
      		scrapPercent: scrapPercent,
      		Bom: Bom != null && Bom.length > 0 ? Bom : null,
			Component: Component != null && Component.length > 0 ? Component : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BOMItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBOMItem(lineNumber, quantity, scrapPercent, Bom, Component, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BOMItem/update/' + id;
		const obj = {
				      		lineNumber: lineNumber,
      		quantity: quantity,
      		scrapPercent: scrapPercent,
      		Bom: Bom != null && Bom.length > 0 ? Bom : null,
			Component: Component != null && Component.length > 0 ? Component : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BOMItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBOMItem(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BOMItem/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BOMItem
	// returns the results untouched as an Observable BOMItem
	// BOMItem model
	// delegates via URI
	//********************************************************************
	getBOMItem(id) : Observable<BOMItem> {
		const uri_ = this.apiUrl + '/BOMItem/load/' + id;

		return this.http.get<BOMItem>(uri_);
	}
	
	//********************************************************************
	// gets all BOMItem
	// returns the results untouched as JSON representation of an
	// Observable array of BOMItem models
	// delegates via URI
	//********************************************************************
	getBOMItems() : Observable<BOMItem[]> {
		const uri_ = this.apiUrl + '/BOMItem/';

		return this
			.http.get<BOMItem[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Bom on a BOMItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBom( bOMItemId, _bomId ): Observable<any> {

		// get the BOMItem from storage
		this.loadHelper( bOMItemId );

	// get the BOM from storage
	var tmp 	= new BOMService(this.http).getBOM(_bomId);

	// assign the Bom
	this.bOMItem.bom = tmp;

	// save the BOMItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Bom on a BOMItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBom( bOMItemId ): Observable<any> {

		// get the BOMItem from storage
		this.loadHelper( bOMItemId );

	// assign Bom to null
	this.bOMItem.bom = null;

	// save the BOMItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Component on a BOMItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignComponent( bOMItemId, _componentId ): Observable<any> {

		// get the BOMItem from storage
		this.loadHelper( bOMItemId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_componentId);

	// assign the Component
	this.bOMItem.component = tmp;

	// save the BOMItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Component on a BOMItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignComponent( bOMItemId ): Observable<any> {

		// get the BOMItem from storage
		this.loadHelper( bOMItemId );

	// assign Component to null
	this.bOMItem.component = null;

	// save the BOMItem
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a BOMItem
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BOMItem/update/' + this.bOMItem;

	return  this.http.post(uri_, this.bOMItem );
}

	//********************************************************************
	// loadHelper - internal helper to load a BOMItem
	//********************************************************************	
	loadHelper( id ) {
		this.getBOMItem(id)
			.subscribe((res : BOMItem) => {
				this.bOMItem = res;
			});
	}
}