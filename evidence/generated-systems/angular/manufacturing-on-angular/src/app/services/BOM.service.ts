import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BOM} from '../models/BOM';
import {ItemService} from '../services/Item.service';
import {BOMItemService} from '../services/BOMItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BOMService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	bOM : BOM;

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
	// add a BOM
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBOM(bomNumber, revision, effectivityStart, effectivityEnd, ParentItem, BomItems, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/BOM/create';
		const obj = {
			      		bomNumber: bomNumber,
      		revision: revision,
      		effectivityStart: effectivityStart,
      		effectivityEnd: effectivityEnd,
      		ParentItem: ParentItem != null && ParentItem.length > 0 ? ParentItem : null,
      		BomItems: BomItems != null && BomItems.length > 0 ? BomItems : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BOM
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBOM(bomNumber, revision, effectivityStart, effectivityEnd, ParentItem, BomItems, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BOM/update/' + id;
		const obj = {
				      		bomNumber: bomNumber,
      		revision: revision,
      		effectivityStart: effectivityStart,
      		effectivityEnd: effectivityEnd,
      		ParentItem: ParentItem != null && ParentItem.length > 0 ? ParentItem : null,
      		BomItems: BomItems != null && BomItems.length > 0 ? BomItems : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BOM
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBOM(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BOM/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BOM
	// returns the results untouched as an Observable BOM
	// BOM model
	// delegates via URI
	//********************************************************************
	getBOM(id) : Observable<BOM> {
		const uri_ = this.apiUrl + '/BOM/load/' + id;

		return this.http.get<BOM>(uri_);
	}
	
	//********************************************************************
	// gets all BOM
	// returns the results untouched as JSON representation of an
	// Observable array of BOM models
	// delegates via URI
	//********************************************************************
	getBOMs() : Observable<BOM[]> {
		const uri_ = this.apiUrl + '/BOM/';

		return this
			.http.get<BOM[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ParentItem on a BOM
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignParentItem( bOMId, _parentItemId ): Observable<any> {

		// get the BOM from storage
		this.loadHelper( bOMId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_parentItemId);

	// assign the ParentItem
	this.bOM.parentItem = tmp;

	// save the BOM
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ParentItem on a BOM
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignParentItem( bOMId ): Observable<any> {

		// get the BOM from storage
		this.loadHelper( bOMId );

	// assign ParentItem to null
	this.bOM.parentItem = null;

	// save the BOM
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more bomItemsIds as a BomItems
	// to a BOM
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBomItems( bOMId, bomItemsIds ): Observable<any> {

		// get the BOM
		this.loadHelper( bOMId );

	// split on a comma with no spaces
	var idList = bomItemsIds.split(',')

	// iterate over array of bomItems ids
	idList.forEach(function (id) {
		// read the BOMItem
		var bOMItem = new BOMItemService(this.http).getBOMItem(id);
		// add the BOMItem if not already assigned
		if ( this.bOM.bomItems.indexOf(bOMItem) == -1 )
		this.bOM.bomItems.push(bOMItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more bomItemsIds as a BomItems
	// from a BOM
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBomItems( bOMId, bomItemsIds ): Observable<any> {

		// get the BOM
		this.loadHelper( bOMId );


	// split on a comma with no spaces
	var idList 					= bomItemsIds.split(',');
	var bomItems 	= this.bOM.bomItems;

	if ( bomItems != null && bomItemsIds != null ) {

		// iterate over array of bomItems ids
		bomItems.forEach(function (obj) {
			if ( bomItemsIds.indexOf(obj._id) > -1 ) {
				// remove the BOMItem
				this.bOM.bomItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BOM
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BOM/update/' + this.bOM;

	return  this.http.post(uri_, this.bOM );
}

	//********************************************************************
	// loadHelper - internal helper to load a BOM
	//********************************************************************	
	loadHelper( id ) {
		this.getBOM(id)
			.subscribe((res : BOM) => {
				this.bOM = res;
			});
	}
}