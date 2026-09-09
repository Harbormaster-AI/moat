import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

 #======================================================================
# 
# Encapsulates data for View Warehouse
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WarehouseView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Warehouse index.")

def get(request, warehouseId ):
	delegate = WarehouseDelegate()
	responseData = delegate.get( warehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	warehouse = json.loads(request.body)
	delegate = WarehouseDelegate()
	responseData = delegate.createFromJson( warehouse )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	warehouse = json.loads(request.body)
	delegate = WarehouseDelegate()
	responseData = delegate.save( warehouse )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, warehouseId ):
	delegate = WarehouseDelegate()
	responseData = delegate.delete( warehouseId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WarehouseDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addStorageLocations( request, warehouseId, StorageLocationsIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.addStorageLocations( warehouseId, StorageLocationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeStorageLocations( request, warehouseId, StorageLocationsIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.removeStorageLocations( warehouseId, StorageLocationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventoryItems( request, warehouseId, InventoryItemsIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.addInventoryItems( warehouseId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventoryItems( request, warehouseId, InventoryItemsIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.removeInventoryItems( warehouseId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInboundShipments( request, warehouseId, InboundShipmentsIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.addInboundShipments( warehouseId, InboundShipmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInboundShipments( request, warehouseId, InboundShipmentsIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.removeInboundShipments( warehouseId, InboundShipmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOutboundAllocations( request, warehouseId, OutboundAllocationsIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.addOutboundAllocations( warehouseId, OutboundAllocationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOutboundAllocations( request, warehouseId, OutboundAllocationsIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.removeOutboundAllocations( warehouseId, OutboundAllocationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOriginTransfers( request, warehouseId, OriginTransfersIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.addOriginTransfers( warehouseId, OriginTransfersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOriginTransfers( request, warehouseId, OriginTransfersIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.removeOriginTransfers( warehouseId, OriginTransfersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDestinationTransfers( request, warehouseId, DestinationTransfersIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.addDestinationTransfers( warehouseId, DestinationTransfersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDestinationTransfers( request, warehouseId, DestinationTransfersIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.removeDestinationTransfers( warehouseId, DestinationTransfersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCycleCounts( request, warehouseId, CycleCountsIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.addCycleCounts( warehouseId, CycleCountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCycleCounts( request, warehouseId, CycleCountsIds ):
	delegate = WarehouseDelegate()
	responseData = delegate.removeCycleCounts( warehouseId, CycleCountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

