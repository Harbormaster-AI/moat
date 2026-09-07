import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InspectionLot} from '../models/InspectionLot';
import {ItemService} from '../services/Item.service';
import {WorkOrderService} from '../services/WorkOrder.service';
import {GoodsReceiptService} from '../services/GoodsReceipt.service';
import {InspectionResultService} from '../services/InspectionResult.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InspectionLotService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inspectionLot : InspectionLot;

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
	// add a InspectionLot
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInspectionLot(lotNumber, quantity, sampleSize, createdOn, Item, WorkOrder, GoodsReceipt, Results, InspectionType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/InspectionLot/create';
		const obj = {
			      		lotNumber: lotNumber,
      		quantity: quantity,
      		sampleSize: sampleSize,
      		createdOn: createdOn,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		WorkOrder: WorkOrder != null && WorkOrder.length > 0 ? WorkOrder : null,
      		GoodsReceipt: GoodsReceipt != null && GoodsReceipt.length > 0 ? GoodsReceipt : null,
      		Results: Results != null && Results.length > 0 ? Results : null,
      		InspectionType: InspectionType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InspectionLot
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInspectionLot(lotNumber, quantity, sampleSize, createdOn, Item, WorkOrder, GoodsReceipt, Results, InspectionType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InspectionLot/update/' + id;
		const obj = {
				      		lotNumber: lotNumber,
      		quantity: quantity,
      		sampleSize: sampleSize,
      		createdOn: createdOn,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		WorkOrder: WorkOrder != null && WorkOrder.length > 0 ? WorkOrder : null,
      		GoodsReceipt: GoodsReceipt != null && GoodsReceipt.length > 0 ? GoodsReceipt : null,
      		Results: Results != null && Results.length > 0 ? Results : null,
      		InspectionType: InspectionType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InspectionLot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInspectionLot(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InspectionLot/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InspectionLot
	// returns the results untouched as an Observable InspectionLot
	// InspectionLot model
	// delegates via URI
	//********************************************************************
	getInspectionLot(id) : Observable<InspectionLot> {
		const uri_ = this.apiUrl + '/InspectionLot/load/' + id;

		return this.http.get<InspectionLot>(uri_);
	}
	
	//********************************************************************
	// gets all InspectionLot
	// returns the results untouched as JSON representation of an
	// Observable array of InspectionLot models
	// delegates via URI
	//********************************************************************
	getInspectionLots() : Observable<InspectionLot[]> {
		const uri_ = this.apiUrl + '/InspectionLot/';

		return this
			.http.get<InspectionLot[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Item on a InspectionLot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( inspectionLotId, _itemId ): Observable<any> {

		// get the InspectionLot from storage
		this.loadHelper( inspectionLotId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.inspectionLot.item = tmp;

	// save the InspectionLot
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a InspectionLot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( inspectionLotId ): Observable<any> {

		// get the InspectionLot from storage
		this.loadHelper( inspectionLotId );

	// assign Item to null
	this.inspectionLot.item = null;

	// save the InspectionLot
	return this.saveHelper();
}

		//********************************************************************
	// assigns a WorkOrder on a InspectionLot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkOrder( inspectionLotId, _workOrderId ): Observable<any> {

		// get the InspectionLot from storage
		this.loadHelper( inspectionLotId );

	// get the WorkOrder from storage
	var tmp 	= new WorkOrderService(this.http).getWorkOrder(_workOrderId);

	// assign the WorkOrder
	this.inspectionLot.workOrder = tmp;

	// save the InspectionLot
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkOrder on a InspectionLot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkOrder( inspectionLotId ): Observable<any> {

		// get the InspectionLot from storage
		this.loadHelper( inspectionLotId );

	// assign WorkOrder to null
	this.inspectionLot.workOrder = null;

	// save the InspectionLot
	return this.saveHelper();
}

		//********************************************************************
	// assigns a GoodsReceipt on a InspectionLot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignGoodsReceipt( inspectionLotId, _goodsReceiptId ): Observable<any> {

		// get the InspectionLot from storage
		this.loadHelper( inspectionLotId );

	// get the GoodsReceipt from storage
	var tmp 	= new GoodsReceiptService(this.http).getGoodsReceipt(_goodsReceiptId);

	// assign the GoodsReceipt
	this.inspectionLot.goodsReceipt = tmp;

	// save the InspectionLot
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a GoodsReceipt on a InspectionLot
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignGoodsReceipt( inspectionLotId ): Observable<any> {

		// get the InspectionLot from storage
		this.loadHelper( inspectionLotId );

	// assign GoodsReceipt to null
	this.inspectionLot.goodsReceipt = null;

	// save the InspectionLot
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more resultsIds as a Results
	// to a InspectionLot
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addResults( inspectionLotId, resultsIds ): Observable<any> {

		// get the InspectionLot
		this.loadHelper( inspectionLotId );

	// split on a comma with no spaces
	var idList = resultsIds.split(',')

	// iterate over array of results ids
	idList.forEach(function (id) {
		// read the InspectionResult
		var inspectionResult = new InspectionResultService(this.http).getInspectionResult(id);
		// add the InspectionResult if not already assigned
		if ( this.inspectionLot.results.indexOf(inspectionResult) == -1 )
		this.inspectionLot.results.push(inspectionResult);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more resultsIds as a Results
	// from a InspectionLot
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeResults( inspectionLotId, resultsIds ): Observable<any> {

		// get the InspectionLot
		this.loadHelper( inspectionLotId );


	// split on a comma with no spaces
	var idList 					= resultsIds.split(',');
	var results 	= this.inspectionLot.results;

	if ( results != null && resultsIds != null ) {

		// iterate over array of results ids
		results.forEach(function (obj) {
			if ( resultsIds.indexOf(obj._id) > -1 ) {
				// remove the InspectionResult
				this.inspectionLot.results.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InspectionLot
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InspectionLot/update/' + this.inspectionLot;

	return  this.http.post(uri_, this.inspectionLot );
}

	//********************************************************************
	// loadHelper - internal helper to load a InspectionLot
	//********************************************************************	
	loadHelper( id ) {
		this.getInspectionLot(id)
			.subscribe((res : InspectionLot) => {
				this.inspectionLot = res;
			});
	}
}