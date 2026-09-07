import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Warehouse} from '../models/Warehouse';
import {StorageLocationService} from '../services/StorageLocation.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import {InboundShipmentService} from '../services/InboundShipment.service';
import {OutboundAllocationService} from '../services/OutboundAllocation.service';
import {TransferOrderService} from '../services/TransferOrder.service';
import {CycleCountService} from '../services/CycleCount.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class WarehouseService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	warehouse : Warehouse;

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
	// add a Warehouse
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addWarehouse(name, code, address, timeZone, allowsOverAllocation, StorageLocations, InventoryItems, InboundShipments, OutboundAllocations, OriginTransfers, DestinationTransfers, CycleCounts) : Observable<any> {
		const uri_ = this.apiUrl + '/Warehouse/create';
		const obj = {
			      		name: name,
      		code: code,
      		address: address,
      		timeZone: timeZone,
      		allowsOverAllocation: allowsOverAllocation,
      		StorageLocations: StorageLocations != null && StorageLocations.length > 0 ? StorageLocations : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
      		InboundShipments: InboundShipments != null && InboundShipments.length > 0 ? InboundShipments : null,
      		OutboundAllocations: OutboundAllocations != null && OutboundAllocations.length > 0 ? OutboundAllocations : null,
      		OriginTransfers: OriginTransfers != null && OriginTransfers.length > 0 ? OriginTransfers : null,
      		DestinationTransfers: DestinationTransfers != null && DestinationTransfers.length > 0 ? DestinationTransfers : null,
			CycleCounts: CycleCounts != null && CycleCounts.length > 0 ? CycleCounts : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Warehouse
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateWarehouse(name, code, address, timeZone, allowsOverAllocation, StorageLocations, InventoryItems, InboundShipments, OutboundAllocations, OriginTransfers, DestinationTransfers, CycleCounts, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Warehouse/update/' + id;
		const obj = {
				      		name: name,
      		code: code,
      		address: address,
      		timeZone: timeZone,
      		allowsOverAllocation: allowsOverAllocation,
      		StorageLocations: StorageLocations != null && StorageLocations.length > 0 ? StorageLocations : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
      		InboundShipments: InboundShipments != null && InboundShipments.length > 0 ? InboundShipments : null,
      		OutboundAllocations: OutboundAllocations != null && OutboundAllocations.length > 0 ? OutboundAllocations : null,
      		OriginTransfers: OriginTransfers != null && OriginTransfers.length > 0 ? OriginTransfers : null,
      		DestinationTransfers: DestinationTransfers != null && DestinationTransfers.length > 0 ? DestinationTransfers : null,
			CycleCounts: CycleCounts != null && CycleCounts.length > 0 ? CycleCounts : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Warehouse
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteWarehouse(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Warehouse/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Warehouse
	// returns the results untouched as an Observable Warehouse
	// Warehouse model
	// delegates via URI
	//********************************************************************
	getWarehouse(id) : Observable<Warehouse> {
		const uri_ = this.apiUrl + '/Warehouse/load/' + id;

		return this.http.get<Warehouse>(uri_);
	}
	
	//********************************************************************
	// gets all Warehouse
	// returns the results untouched as JSON representation of an
	// Observable array of Warehouse models
	// delegates via URI
	//********************************************************************
	getWarehouses() : Observable<Warehouse[]> {
		const uri_ = this.apiUrl + '/Warehouse/';

		return this
			.http.get<Warehouse[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more storageLocationsIds as a StorageLocations
	// to a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addStorageLocations( warehouseId, storageLocationsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );

	// split on a comma with no spaces
	var idList = storageLocationsIds.split(',')

	// iterate over array of storageLocations ids
	idList.forEach(function (id) {
		// read the StorageLocation
		var storageLocation = new StorageLocationService(this.http).getStorageLocation(id);
		// add the StorageLocation if not already assigned
		if ( this.warehouse.storageLocations.indexOf(storageLocation) == -1 )
		this.warehouse.storageLocations.push(storageLocation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more storageLocationsIds as a StorageLocations
	// from a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeStorageLocations( warehouseId, storageLocationsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );


	// split on a comma with no spaces
	var idList 					= storageLocationsIds.split(',');
	var storageLocations 	= this.warehouse.storageLocations;

	if ( storageLocations != null && storageLocationsIds != null ) {

		// iterate over array of storageLocations ids
		storageLocations.forEach(function (obj) {
			if ( storageLocationsIds.indexOf(obj._id) > -1 ) {
				// remove the StorageLocation
				this.warehouse.storageLocations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more inventoryItemsIds as a InventoryItems
	// to a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventoryItems( warehouseId, inventoryItemsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );

	// split on a comma with no spaces
	var idList = inventoryItemsIds.split(',')

	// iterate over array of inventoryItems ids
	idList.forEach(function (id) {
		// read the InventoryItem
		var inventoryItem = new InventoryItemService(this.http).getInventoryItem(id);
		// add the InventoryItem if not already assigned
		if ( this.warehouse.inventoryItems.indexOf(inventoryItem) == -1 )
		this.warehouse.inventoryItems.push(inventoryItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventoryItemsIds as a InventoryItems
	// from a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventoryItems( warehouseId, inventoryItemsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );


	// split on a comma with no spaces
	var idList 					= inventoryItemsIds.split(',');
	var inventoryItems 	= this.warehouse.inventoryItems;

	if ( inventoryItems != null && inventoryItemsIds != null ) {

		// iterate over array of inventoryItems ids
		inventoryItems.forEach(function (obj) {
			if ( inventoryItemsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryItem
				this.warehouse.inventoryItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more inboundShipmentsIds as a InboundShipments
	// to a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInboundShipments( warehouseId, inboundShipmentsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );

	// split on a comma with no spaces
	var idList = inboundShipmentsIds.split(',')

	// iterate over array of inboundShipments ids
	idList.forEach(function (id) {
		// read the InboundShipment
		var inboundShipment = new InboundShipmentService(this.http).getInboundShipment(id);
		// add the InboundShipment if not already assigned
		if ( this.warehouse.inboundShipments.indexOf(inboundShipment) == -1 )
		this.warehouse.inboundShipments.push(inboundShipment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inboundShipmentsIds as a InboundShipments
	// from a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInboundShipments( warehouseId, inboundShipmentsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );


	// split on a comma with no spaces
	var idList 					= inboundShipmentsIds.split(',');
	var inboundShipments 	= this.warehouse.inboundShipments;

	if ( inboundShipments != null && inboundShipmentsIds != null ) {

		// iterate over array of inboundShipments ids
		inboundShipments.forEach(function (obj) {
			if ( inboundShipmentsIds.indexOf(obj._id) > -1 ) {
				// remove the InboundShipment
				this.warehouse.inboundShipments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more outboundAllocationsIds as a OutboundAllocations
	// to a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOutboundAllocations( warehouseId, outboundAllocationsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );

	// split on a comma with no spaces
	var idList = outboundAllocationsIds.split(',')

	// iterate over array of outboundAllocations ids
	idList.forEach(function (id) {
		// read the OutboundAllocation
		var outboundAllocation = new OutboundAllocationService(this.http).getOutboundAllocation(id);
		// add the OutboundAllocation if not already assigned
		if ( this.warehouse.outboundAllocations.indexOf(outboundAllocation) == -1 )
		this.warehouse.outboundAllocations.push(outboundAllocation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more outboundAllocationsIds as a OutboundAllocations
	// from a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOutboundAllocations( warehouseId, outboundAllocationsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );


	// split on a comma with no spaces
	var idList 					= outboundAllocationsIds.split(',');
	var outboundAllocations 	= this.warehouse.outboundAllocations;

	if ( outboundAllocations != null && outboundAllocationsIds != null ) {

		// iterate over array of outboundAllocations ids
		outboundAllocations.forEach(function (obj) {
			if ( outboundAllocationsIds.indexOf(obj._id) > -1 ) {
				// remove the OutboundAllocation
				this.warehouse.outboundAllocations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more originTransfersIds as a OriginTransfers
	// to a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOriginTransfers( warehouseId, originTransfersIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );

	// split on a comma with no spaces
	var idList = originTransfersIds.split(',')

	// iterate over array of originTransfers ids
	idList.forEach(function (id) {
		// read the TransferOrder
		var transferOrder = new TransferOrderService(this.http).getTransferOrder(id);
		// add the TransferOrder if not already assigned
		if ( this.warehouse.originTransfers.indexOf(transferOrder) == -1 )
		this.warehouse.originTransfers.push(transferOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more originTransfersIds as a OriginTransfers
	// from a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOriginTransfers( warehouseId, originTransfersIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );


	// split on a comma with no spaces
	var idList 					= originTransfersIds.split(',');
	var originTransfers 	= this.warehouse.originTransfers;

	if ( originTransfers != null && originTransfersIds != null ) {

		// iterate over array of originTransfers ids
		originTransfers.forEach(function (obj) {
			if ( originTransfersIds.indexOf(obj._id) > -1 ) {
				// remove the TransferOrder
				this.warehouse.originTransfers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more destinationTransfersIds as a DestinationTransfers
	// to a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDestinationTransfers( warehouseId, destinationTransfersIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );

	// split on a comma with no spaces
	var idList = destinationTransfersIds.split(',')

	// iterate over array of destinationTransfers ids
	idList.forEach(function (id) {
		// read the TransferOrder
		var transferOrder = new TransferOrderService(this.http).getTransferOrder(id);
		// add the TransferOrder if not already assigned
		if ( this.warehouse.destinationTransfers.indexOf(transferOrder) == -1 )
		this.warehouse.destinationTransfers.push(transferOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more destinationTransfersIds as a DestinationTransfers
	// from a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDestinationTransfers( warehouseId, destinationTransfersIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );


	// split on a comma with no spaces
	var idList 					= destinationTransfersIds.split(',');
	var destinationTransfers 	= this.warehouse.destinationTransfers;

	if ( destinationTransfers != null && destinationTransfersIds != null ) {

		// iterate over array of destinationTransfers ids
		destinationTransfers.forEach(function (obj) {
			if ( destinationTransfersIds.indexOf(obj._id) > -1 ) {
				// remove the TransferOrder
				this.warehouse.destinationTransfers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more cycleCountsIds as a CycleCounts
	// to a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCycleCounts( warehouseId, cycleCountsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );

	// split on a comma with no spaces
	var idList = cycleCountsIds.split(',')

	// iterate over array of cycleCounts ids
	idList.forEach(function (id) {
		// read the CycleCount
		var cycleCount = new CycleCountService(this.http).getCycleCount(id);
		// add the CycleCount if not already assigned
		if ( this.warehouse.cycleCounts.indexOf(cycleCount) == -1 )
		this.warehouse.cycleCounts.push(cycleCount);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more cycleCountsIds as a CycleCounts
	// from a Warehouse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCycleCounts( warehouseId, cycleCountsIds ): Observable<any> {

		// get the Warehouse
		this.loadHelper( warehouseId );


	// split on a comma with no spaces
	var idList 					= cycleCountsIds.split(',');
	var cycleCounts 	= this.warehouse.cycleCounts;

	if ( cycleCounts != null && cycleCountsIds != null ) {

		// iterate over array of cycleCounts ids
		cycleCounts.forEach(function (obj) {
			if ( cycleCountsIds.indexOf(obj._id) > -1 ) {
				// remove the CycleCount
				this.warehouse.cycleCounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Warehouse
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Warehouse/update/' + this.warehouse;

	return  this.http.post(uri_, this.warehouse );
}

	//********************************************************************
	// loadHelper - internal helper to load a Warehouse
	//********************************************************************	
	loadHelper( id ) {
		this.getWarehouse(id)
			.subscribe((res : Warehouse) => {
				this.warehouse = res;
			});
	}
}