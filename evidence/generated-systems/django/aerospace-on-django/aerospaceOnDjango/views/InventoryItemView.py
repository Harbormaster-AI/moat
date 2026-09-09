import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

 #======================================================================
# 
# Encapsulates data for View InventoryItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InventoryItem index.")

def get(request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.get( inventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inventoryItem = json.loads(request.body)
	delegate = InventoryItemDelegate()
	responseData = delegate.createFromJson( inventoryItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inventoryItem = json.loads(request.body)
	delegate = InventoryItemDelegate()
	responseData = delegate.save( inventoryItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.delete( inventoryItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InventoryItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignComponent( request, inventoryItemId, ComponentId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.saveComponent( inventoryItemId, ComponentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignComponent( request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.deleteComponent( inventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, inventoryItemId, WarehouseId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.saveWarehouse( inventoryItemId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.deleteWarehouse( inventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

