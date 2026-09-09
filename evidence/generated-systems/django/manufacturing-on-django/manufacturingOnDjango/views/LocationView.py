import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.LocationDelegate import LocationDelegate

 #======================================================================
# 
# Encapsulates data for View Location
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LocationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Location index.")

def get(request, locationId ):
	delegate = LocationDelegate()
	responseData = delegate.get( locationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	location = json.loads(request.body)
	delegate = LocationDelegate()
	responseData = delegate.createFromJson( location )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	location = json.loads(request.body)
	delegate = LocationDelegate()
	responseData = delegate.save( location )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, locationId ):
	delegate = LocationDelegate()
	responseData = delegate.delete( locationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LocationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, locationId, WarehouseId ):
	delegate = LocationDelegate()
	responseData = delegate.saveWarehouse( locationId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, locationId ):
	delegate = LocationDelegate()
	responseData = delegate.deleteWarehouse( locationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventoryItems( request, locationId, InventoryItemsIds ):
	delegate = LocationDelegate()
	responseData = delegate.addInventoryItems( locationId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventoryItems( request, locationId, InventoryItemsIds ):
	delegate = LocationDelegate()
	responseData = delegate.removeInventoryItems( locationId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

