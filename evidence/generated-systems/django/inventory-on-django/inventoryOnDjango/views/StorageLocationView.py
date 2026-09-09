import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

 #======================================================================
# 
# Encapsulates data for View StorageLocation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StorageLocationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the StorageLocation index.")

def get(request, storageLocationId ):
	delegate = StorageLocationDelegate()
	responseData = delegate.get( storageLocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	storageLocation = json.loads(request.body)
	delegate = StorageLocationDelegate()
	responseData = delegate.createFromJson( storageLocation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	storageLocation = json.loads(request.body)
	delegate = StorageLocationDelegate()
	responseData = delegate.save( storageLocation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, storageLocationId ):
	delegate = StorageLocationDelegate()
	responseData = delegate.delete( storageLocationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = StorageLocationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, storageLocationId, WarehouseId ):
	delegate = StorageLocationDelegate()
	responseData = delegate.saveWarehouse( storageLocationId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, storageLocationId ):
	delegate = StorageLocationDelegate()
	responseData = delegate.deleteWarehouse( storageLocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignParentLocation( request, storageLocationId, ParentLocationId ):
	delegate = StorageLocationDelegate()
	responseData = delegate.saveParentLocation( storageLocationId, ParentLocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignParentLocation( request, storageLocationId ):
	delegate = StorageLocationDelegate()
	responseData = delegate.deleteParentLocation( storageLocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChildLocations( request, storageLocationId, ChildLocationsIds ):
	delegate = StorageLocationDelegate()
	responseData = delegate.addChildLocations( storageLocationId, ChildLocationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChildLocations( request, storageLocationId, ChildLocationsIds ):
	delegate = StorageLocationDelegate()
	responseData = delegate.removeChildLocations( storageLocationId, ChildLocationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventoryItems( request, storageLocationId, InventoryItemsIds ):
	delegate = StorageLocationDelegate()
	responseData = delegate.addInventoryItems( storageLocationId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventoryItems( request, storageLocationId, InventoryItemsIds ):
	delegate = StorageLocationDelegate()
	responseData = delegate.removeInventoryItems( storageLocationId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

