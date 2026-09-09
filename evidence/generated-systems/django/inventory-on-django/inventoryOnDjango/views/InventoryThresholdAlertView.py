import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.InventoryThresholdAlertDelegate import InventoryThresholdAlertDelegate

 #======================================================================
# 
# Encapsulates data for View InventoryThresholdAlert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryThresholdAlertView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InventoryThresholdAlert index.")

def get(request, inventoryThresholdAlertId ):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.get( inventoryThresholdAlertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inventoryThresholdAlert = json.loads(request.body)
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.createFromJson( inventoryThresholdAlert )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inventoryThresholdAlert = json.loads(request.body)
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.save( inventoryThresholdAlert )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inventoryThresholdAlertId ):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.delete( inventoryThresholdAlertId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, inventoryThresholdAlertId, SkuId ):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.saveSku( inventoryThresholdAlertId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, inventoryThresholdAlertId ):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.deleteSku( inventoryThresholdAlertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, inventoryThresholdAlertId, WarehouseId ):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.saveWarehouse( inventoryThresholdAlertId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, inventoryThresholdAlertId ):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.deleteWarehouse( inventoryThresholdAlertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLocation( request, inventoryThresholdAlertId, LocationId ):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.saveLocation( inventoryThresholdAlertId, LocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLocation( request, inventoryThresholdAlertId ):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.deleteLocation( inventoryThresholdAlertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRelatedPolicy( request, inventoryThresholdAlertId, RelatedPolicyId ):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.saveRelatedPolicy( inventoryThresholdAlertId, RelatedPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRelatedPolicy( request, inventoryThresholdAlertId ):
	delegate = InventoryThresholdAlertDelegate()
	responseData = delegate.deleteRelatedPolicy( inventoryThresholdAlertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

