import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {QualitySpecification} from '../models/QualitySpecification';
import {ItemService} from '../services/Item.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class QualitySpecificationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	qualitySpecification : QualitySpecification;

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
	// add a QualitySpecification
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addQualitySpecification(specCode, name, version, Item) : Observable<any> {
		const uri_ = this.apiUrl + '/QualitySpecification/create';
		const obj = {
			      		specCode: specCode,
      		name: name,
      		version: version,
			Item: Item != null && Item.length > 0 ? Item : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a QualitySpecification
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateQualitySpecification(specCode, name, version, Item, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/QualitySpecification/update/' + id;
		const obj = {
				      		specCode: specCode,
      		name: name,
      		version: version,
			Item: Item != null && Item.length > 0 ? Item : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a QualitySpecification
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteQualitySpecification(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/QualitySpecification/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a QualitySpecification
	// returns the results untouched as an Observable QualitySpecification
	// QualitySpecification model
	// delegates via URI
	//********************************************************************
	getQualitySpecification(id) : Observable<QualitySpecification> {
		const uri_ = this.apiUrl + '/QualitySpecification/load/' + id;

		return this.http.get<QualitySpecification>(uri_);
	}
	
	//********************************************************************
	// gets all QualitySpecification
	// returns the results untouched as JSON representation of an
	// Observable array of QualitySpecification models
	// delegates via URI
	//********************************************************************
	getQualitySpecifications() : Observable<QualitySpecification[]> {
		const uri_ = this.apiUrl + '/QualitySpecification/';

		return this
			.http.get<QualitySpecification[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Item on a QualitySpecification
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( qualitySpecificationId, _itemId ): Observable<any> {

		// get the QualitySpecification from storage
		this.loadHelper( qualitySpecificationId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.qualitySpecification.item = tmp;

	// save the QualitySpecification
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a QualitySpecification
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( qualitySpecificationId ): Observable<any> {

		// get the QualitySpecification from storage
		this.loadHelper( qualitySpecificationId );

	// assign Item to null
	this.qualitySpecification.item = null;

	// save the QualitySpecification
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a QualitySpecification
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/QualitySpecification/update/' + this.qualitySpecification;

	return  this.http.post(uri_, this.qualitySpecification );
}

	//********************************************************************
	// loadHelper - internal helper to load a QualitySpecification
	//********************************************************************	
	loadHelper( id ) {
		this.getQualitySpecification(id)
			.subscribe((res : QualitySpecification) => {
				this.qualitySpecification = res;
			});
	}
}