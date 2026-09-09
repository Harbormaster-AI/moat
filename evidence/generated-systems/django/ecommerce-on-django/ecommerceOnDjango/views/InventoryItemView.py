import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

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

def assignVariant( request, inventoryItemId, VariantId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.saveVariant( inventoryItemId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.deleteVariant( inventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFulfillmentCenter( request, inventoryItemId, FulfillmentCenterId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.saveFulfillmentCenter( inventoryItemId, FulfillmentCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFulfillmentCenter( request, inventoryItemId ):
	delegate = InventoryItemDelegate()
	responseData = delegate.deleteFulfillmentCenter( inventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

