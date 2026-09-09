import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.FulfillmentCenterDelegate import FulfillmentCenterDelegate

 #======================================================================
# 
# Encapsulates data for View FulfillmentCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FulfillmentCenterView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the FulfillmentCenter index.")

def get(request, fulfillmentCenterId ):
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.get( fulfillmentCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	fulfillmentCenter = json.loads(request.body)
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.createFromJson( fulfillmentCenter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	fulfillmentCenter = json.loads(request.body)
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.save( fulfillmentCenter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, fulfillmentCenterId ):
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.delete( fulfillmentCenterId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, fulfillmentCenterId, MerchantId ):
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.saveMerchant( fulfillmentCenterId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, fulfillmentCenterId ):
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.deleteMerchant( fulfillmentCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventoryItems( request, fulfillmentCenterId, InventoryItemsIds ):
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.addInventoryItems( fulfillmentCenterId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventoryItems( request, fulfillmentCenterId, InventoryItemsIds ):
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.removeInventoryItems( fulfillmentCenterId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addShipments( request, fulfillmentCenterId, ShipmentsIds ):
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.addShipments( fulfillmentCenterId, ShipmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeShipments( request, fulfillmentCenterId, ShipmentsIds ):
	delegate = FulfillmentCenterDelegate()
	responseData = delegate.removeShipments( fulfillmentCenterId, ShipmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

