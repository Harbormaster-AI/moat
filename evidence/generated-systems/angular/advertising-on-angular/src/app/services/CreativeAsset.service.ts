import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CreativeAsset} from '../models/CreativeAsset';
import {CreativeFileService} from '../services/CreativeFile.service';
import {CreativeApprovalService} from '../services/CreativeApproval.service';
import {CreativeVariationService} from '../services/CreativeVariation.service';
import {LineItemService} from '../services/LineItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CreativeAssetService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	creativeAsset : CreativeAsset;

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
	// add a CreativeAsset
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCreativeAsset(name, clickUrl, landingPage, width, height, durationSeconds, Files, Approvals, Variations, LineItems, CreativeType, AdFormat) : Observable<any> {
		const uri_ = this.apiUrl + '/CreativeAsset/create';
		const obj = {
			      		name: name,
      		clickUrl: clickUrl,
      		landingPage: landingPage,
      		width: width,
      		height: height,
      		durationSeconds: durationSeconds,
      		Files: Files != null && Files.length > 0 ? Files : null,
      		Approvals: Approvals != null && Approvals.length > 0 ? Approvals : null,
      		Variations: Variations != null && Variations.length > 0 ? Variations : null,
      		LineItems: LineItems != null && LineItems.length > 0 ? LineItems : null,
      		CreativeType: CreativeType,
			AdFormat: AdFormat
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CreativeAsset
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCreativeAsset(name, clickUrl, landingPage, width, height, durationSeconds, Files, Approvals, Variations, LineItems, CreativeType, AdFormat, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CreativeAsset/update/' + id;
		const obj = {
				      		name: name,
      		clickUrl: clickUrl,
      		landingPage: landingPage,
      		width: width,
      		height: height,
      		durationSeconds: durationSeconds,
      		Files: Files != null && Files.length > 0 ? Files : null,
      		Approvals: Approvals != null && Approvals.length > 0 ? Approvals : null,
      		Variations: Variations != null && Variations.length > 0 ? Variations : null,
      		LineItems: LineItems != null && LineItems.length > 0 ? LineItems : null,
      		CreativeType: CreativeType,
			AdFormat: AdFormat
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CreativeAsset
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCreativeAsset(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CreativeAsset/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CreativeAsset
	// returns the results untouched as an Observable CreativeAsset
	// CreativeAsset model
	// delegates via URI
	//********************************************************************
	getCreativeAsset(id) : Observable<CreativeAsset> {
		const uri_ = this.apiUrl + '/CreativeAsset/load/' + id;

		return this.http.get<CreativeAsset>(uri_);
	}
	
	//********************************************************************
	// gets all CreativeAsset
	// returns the results untouched as JSON representation of an
	// Observable array of CreativeAsset models
	// delegates via URI
	//********************************************************************
	getCreativeAssets() : Observable<CreativeAsset[]> {
		const uri_ = this.apiUrl + '/CreativeAsset/';

		return this
			.http.get<CreativeAsset[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more filesIds as a Files
	// to a CreativeAsset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFiles( creativeAssetId, filesIds ): Observable<any> {

		// get the CreativeAsset
		this.loadHelper( creativeAssetId );

	// split on a comma with no spaces
	var idList = filesIds.split(',')

	// iterate over array of files ids
	idList.forEach(function (id) {
		// read the CreativeFile
		var creativeFile = new CreativeFileService(this.http).getCreativeFile(id);
		// add the CreativeFile if not already assigned
		if ( this.creativeAsset.files.indexOf(creativeFile) == -1 )
		this.creativeAsset.files.push(creativeFile);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more filesIds as a Files
	// from a CreativeAsset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFiles( creativeAssetId, filesIds ): Observable<any> {

		// get the CreativeAsset
		this.loadHelper( creativeAssetId );


	// split on a comma with no spaces
	var idList 					= filesIds.split(',');
	var files 	= this.creativeAsset.files;

	if ( files != null && filesIds != null ) {

		// iterate over array of files ids
		files.forEach(function (obj) {
			if ( filesIds.indexOf(obj._id) > -1 ) {
				// remove the CreativeFile
				this.creativeAsset.files.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more approvalsIds as a Approvals
	// to a CreativeAsset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addApprovals( creativeAssetId, approvalsIds ): Observable<any> {

		// get the CreativeAsset
		this.loadHelper( creativeAssetId );

	// split on a comma with no spaces
	var idList = approvalsIds.split(',')

	// iterate over array of approvals ids
	idList.forEach(function (id) {
		// read the CreativeApproval
		var creativeApproval = new CreativeApprovalService(this.http).getCreativeApproval(id);
		// add the CreativeApproval if not already assigned
		if ( this.creativeAsset.approvals.indexOf(creativeApproval) == -1 )
		this.creativeAsset.approvals.push(creativeApproval);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more approvalsIds as a Approvals
	// from a CreativeAsset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeApprovals( creativeAssetId, approvalsIds ): Observable<any> {

		// get the CreativeAsset
		this.loadHelper( creativeAssetId );


	// split on a comma with no spaces
	var idList 					= approvalsIds.split(',');
	var approvals 	= this.creativeAsset.approvals;

	if ( approvals != null && approvalsIds != null ) {

		// iterate over array of approvals ids
		approvals.forEach(function (obj) {
			if ( approvalsIds.indexOf(obj._id) > -1 ) {
				// remove the CreativeApproval
				this.creativeAsset.approvals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more variationsIds as a Variations
	// to a CreativeAsset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVariations( creativeAssetId, variationsIds ): Observable<any> {

		// get the CreativeAsset
		this.loadHelper( creativeAssetId );

	// split on a comma with no spaces
	var idList = variationsIds.split(',')

	// iterate over array of variations ids
	idList.forEach(function (id) {
		// read the CreativeVariation
		var creativeVariation = new CreativeVariationService(this.http).getCreativeVariation(id);
		// add the CreativeVariation if not already assigned
		if ( this.creativeAsset.variations.indexOf(creativeVariation) == -1 )
		this.creativeAsset.variations.push(creativeVariation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more variationsIds as a Variations
	// from a CreativeAsset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVariations( creativeAssetId, variationsIds ): Observable<any> {

		// get the CreativeAsset
		this.loadHelper( creativeAssetId );


	// split on a comma with no spaces
	var idList 					= variationsIds.split(',');
	var variations 	= this.creativeAsset.variations;

	if ( variations != null && variationsIds != null ) {

		// iterate over array of variations ids
		variations.forEach(function (obj) {
			if ( variationsIds.indexOf(obj._id) > -1 ) {
				// remove the CreativeVariation
				this.creativeAsset.variations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more lineItemsIds as a LineItems
	// to a CreativeAsset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLineItems( creativeAssetId, lineItemsIds ): Observable<any> {

		// get the CreativeAsset
		this.loadHelper( creativeAssetId );

	// split on a comma with no spaces
	var idList = lineItemsIds.split(',')

	// iterate over array of lineItems ids
	idList.forEach(function (id) {
		// read the LineItem
		var lineItem = new LineItemService(this.http).getLineItem(id);
		// add the LineItem if not already assigned
		if ( this.creativeAsset.lineItems.indexOf(lineItem) == -1 )
		this.creativeAsset.lineItems.push(lineItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more lineItemsIds as a LineItems
	// from a CreativeAsset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLineItems( creativeAssetId, lineItemsIds ): Observable<any> {

		// get the CreativeAsset
		this.loadHelper( creativeAssetId );


	// split on a comma with no spaces
	var idList 					= lineItemsIds.split(',');
	var lineItems 	= this.creativeAsset.lineItems;

	if ( lineItems != null && lineItemsIds != null ) {

		// iterate over array of lineItems ids
		lineItems.forEach(function (obj) {
			if ( lineItemsIds.indexOf(obj._id) > -1 ) {
				// remove the LineItem
				this.creativeAsset.lineItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a CreativeAsset
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CreativeAsset/update/' + this.creativeAsset;

	return  this.http.post(uri_, this.creativeAsset );
}

	//********************************************************************
	// loadHelper - internal helper to load a CreativeAsset
	//********************************************************************	
	loadHelper( id ) {
		this.getCreativeAsset(id)
			.subscribe((res : CreativeAsset) => {
				this.creativeAsset = res;
			});
	}
}