import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Publisher} from '../models/Publisher';
import {InventorySourceService} from '../services/InventorySource.service';
import {DealService} from '../services/Deal.service';
import {CreativeApprovalService} from '../services/CreativeApproval.service';
import {InsertionOrderService} from '../services/InsertionOrder.service';
import {RateCardService} from '../services/RateCard.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PublisherService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	publisher : Publisher;

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
	// add a Publisher
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPublisher(name, website, InventorySources, Deals, CreativeApprovals, InsertionOrders, RateCards, PublisherType) : Observable<any> {
		const uri_ = this.apiUrl + '/Publisher/create';
		const obj = {
			      		name: name,
      		website: website,
      		InventorySources: InventorySources != null && InventorySources.length > 0 ? InventorySources : null,
      		Deals: Deals != null && Deals.length > 0 ? Deals : null,
      		CreativeApprovals: CreativeApprovals != null && CreativeApprovals.length > 0 ? CreativeApprovals : null,
      		InsertionOrders: InsertionOrders != null && InsertionOrders.length > 0 ? InsertionOrders : null,
      		RateCards: RateCards != null && RateCards.length > 0 ? RateCards : null,
			PublisherType: PublisherType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Publisher
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePublisher(name, website, InventorySources, Deals, CreativeApprovals, InsertionOrders, RateCards, PublisherType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Publisher/update/' + id;
		const obj = {
				      		name: name,
      		website: website,
      		InventorySources: InventorySources != null && InventorySources.length > 0 ? InventorySources : null,
      		Deals: Deals != null && Deals.length > 0 ? Deals : null,
      		CreativeApprovals: CreativeApprovals != null && CreativeApprovals.length > 0 ? CreativeApprovals : null,
      		InsertionOrders: InsertionOrders != null && InsertionOrders.length > 0 ? InsertionOrders : null,
      		RateCards: RateCards != null && RateCards.length > 0 ? RateCards : null,
			PublisherType: PublisherType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Publisher
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePublisher(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Publisher/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Publisher
	// returns the results untouched as an Observable Publisher
	// Publisher model
	// delegates via URI
	//********************************************************************
	getPublisher(id) : Observable<Publisher> {
		const uri_ = this.apiUrl + '/Publisher/load/' + id;

		return this.http.get<Publisher>(uri_);
	}
	
	//********************************************************************
	// gets all Publisher
	// returns the results untouched as JSON representation of an
	// Observable array of Publisher models
	// delegates via URI
	//********************************************************************
	getPublishers() : Observable<Publisher[]> {
		const uri_ = this.apiUrl + '/Publisher/';

		return this
			.http.get<Publisher[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more inventorySourcesIds as a InventorySources
	// to a Publisher
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventorySources( publisherId, inventorySourcesIds ): Observable<any> {

		// get the Publisher
		this.loadHelper( publisherId );

	// split on a comma with no spaces
	var idList = inventorySourcesIds.split(',')

	// iterate over array of inventorySources ids
	idList.forEach(function (id) {
		// read the InventorySource
		var inventorySource = new InventorySourceService(this.http).getInventorySource(id);
		// add the InventorySource if not already assigned
		if ( this.publisher.inventorySources.indexOf(inventorySource) == -1 )
		this.publisher.inventorySources.push(inventorySource);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventorySourcesIds as a InventorySources
	// from a Publisher
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventorySources( publisherId, inventorySourcesIds ): Observable<any> {

		// get the Publisher
		this.loadHelper( publisherId );


	// split on a comma with no spaces
	var idList 					= inventorySourcesIds.split(',');
	var inventorySources 	= this.publisher.inventorySources;

	if ( inventorySources != null && inventorySourcesIds != null ) {

		// iterate over array of inventorySources ids
		inventorySources.forEach(function (obj) {
			if ( inventorySourcesIds.indexOf(obj._id) > -1 ) {
				// remove the InventorySource
				this.publisher.inventorySources.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dealsIds as a Deals
	// to a Publisher
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDeals( publisherId, dealsIds ): Observable<any> {

		// get the Publisher
		this.loadHelper( publisherId );

	// split on a comma with no spaces
	var idList = dealsIds.split(',')

	// iterate over array of deals ids
	idList.forEach(function (id) {
		// read the Deal
		var deal = new DealService(this.http).getDeal(id);
		// add the Deal if not already assigned
		if ( this.publisher.deals.indexOf(deal) == -1 )
		this.publisher.deals.push(deal);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dealsIds as a Deals
	// from a Publisher
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDeals( publisherId, dealsIds ): Observable<any> {

		// get the Publisher
		this.loadHelper( publisherId );


	// split on a comma with no spaces
	var idList 					= dealsIds.split(',');
	var deals 	= this.publisher.deals;

	if ( deals != null && dealsIds != null ) {

		// iterate over array of deals ids
		deals.forEach(function (obj) {
			if ( dealsIds.indexOf(obj._id) > -1 ) {
				// remove the Deal
				this.publisher.deals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more creativeApprovalsIds as a CreativeApprovals
	// to a Publisher
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCreativeApprovals( publisherId, creativeApprovalsIds ): Observable<any> {

		// get the Publisher
		this.loadHelper( publisherId );

	// split on a comma with no spaces
	var idList = creativeApprovalsIds.split(',')

	// iterate over array of creativeApprovals ids
	idList.forEach(function (id) {
		// read the CreativeApproval
		var creativeApproval = new CreativeApprovalService(this.http).getCreativeApproval(id);
		// add the CreativeApproval if not already assigned
		if ( this.publisher.creativeApprovals.indexOf(creativeApproval) == -1 )
		this.publisher.creativeApprovals.push(creativeApproval);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more creativeApprovalsIds as a CreativeApprovals
	// from a Publisher
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCreativeApprovals( publisherId, creativeApprovalsIds ): Observable<any> {

		// get the Publisher
		this.loadHelper( publisherId );


	// split on a comma with no spaces
	var idList 					= creativeApprovalsIds.split(',');
	var creativeApprovals 	= this.publisher.creativeApprovals;

	if ( creativeApprovals != null && creativeApprovalsIds != null ) {

		// iterate over array of creativeApprovals ids
		creativeApprovals.forEach(function (obj) {
			if ( creativeApprovalsIds.indexOf(obj._id) > -1 ) {
				// remove the CreativeApproval
				this.publisher.creativeApprovals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more insertionOrdersIds as a InsertionOrders
	// to a Publisher
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInsertionOrders( publisherId, insertionOrdersIds ): Observable<any> {

		// get the Publisher
		this.loadHelper( publisherId );

	// split on a comma with no spaces
	var idList = insertionOrdersIds.split(',')

	// iterate over array of insertionOrders ids
	idList.forEach(function (id) {
		// read the InsertionOrder
		var insertionOrder = new InsertionOrderService(this.http).getInsertionOrder(id);
		// add the InsertionOrder if not already assigned
		if ( this.publisher.insertionOrders.indexOf(insertionOrder) == -1 )
		this.publisher.insertionOrders.push(insertionOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more insertionOrdersIds as a InsertionOrders
	// from a Publisher
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInsertionOrders( publisherId, insertionOrdersIds ): Observable<any> {

		// get the Publisher
		this.loadHelper( publisherId );


	// split on a comma with no spaces
	var idList 					= insertionOrdersIds.split(',');
	var insertionOrders 	= this.publisher.insertionOrders;

	if ( insertionOrders != null && insertionOrdersIds != null ) {

		// iterate over array of insertionOrders ids
		insertionOrders.forEach(function (obj) {
			if ( insertionOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the InsertionOrder
				this.publisher.insertionOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more rateCardsIds as a RateCards
	// to a Publisher
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRateCards( publisherId, rateCardsIds ): Observable<any> {

		// get the Publisher
		this.loadHelper( publisherId );

	// split on a comma with no spaces
	var idList = rateCardsIds.split(',')

	// iterate over array of rateCards ids
	idList.forEach(function (id) {
		// read the RateCard
		var rateCard = new RateCardService(this.http).getRateCard(id);
		// add the RateCard if not already assigned
		if ( this.publisher.rateCards.indexOf(rateCard) == -1 )
		this.publisher.rateCards.push(rateCard);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more rateCardsIds as a RateCards
	// from a Publisher
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRateCards( publisherId, rateCardsIds ): Observable<any> {

		// get the Publisher
		this.loadHelper( publisherId );


	// split on a comma with no spaces
	var idList 					= rateCardsIds.split(',');
	var rateCards 	= this.publisher.rateCards;

	if ( rateCards != null && rateCardsIds != null ) {

		// iterate over array of rateCards ids
		rateCards.forEach(function (obj) {
			if ( rateCardsIds.indexOf(obj._id) > -1 ) {
				// remove the RateCard
				this.publisher.rateCards.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Publisher
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Publisher/update/' + this.publisher;

	return  this.http.post(uri_, this.publisher );
}

	//********************************************************************
	// loadHelper - internal helper to load a Publisher
	//********************************************************************	
	loadHelper( id ) {
		this.getPublisher(id)
			.subscribe((res : Publisher) => {
				this.publisher = res;
			});
	}
}