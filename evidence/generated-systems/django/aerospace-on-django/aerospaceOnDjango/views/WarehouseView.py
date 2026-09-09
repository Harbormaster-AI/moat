import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

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

