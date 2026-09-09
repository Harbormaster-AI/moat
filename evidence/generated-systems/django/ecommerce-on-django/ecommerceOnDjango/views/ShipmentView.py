import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.ShipmentDelegate import ShipmentDelegate

 #======================================================================
# 
# Encapsulates data for View Shipment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShipmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Shipment index.")

def get(request, shipmentId ):
	delegate = ShipmentDelegate()
	responseData = delegate.get( shipmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	shipment = json.loads(request.body)
	delegate = ShipmentDelegate()
	responseData = delegate.createFromJson( shipment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	shipment = json.loads(request.body)
	delegate = ShipmentDelegate()
	responseData = delegate.save( shipment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, shipmentId ):
	delegate = ShipmentDelegate()
	responseData = delegate.delete( shipmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ShipmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, shipmentId, OrderId ):
	delegate = ShipmentDelegate()
	responseData = delegate.saveOrder( shipmentId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, shipmentId ):
	delegate = ShipmentDelegate()
	responseData = delegate.deleteOrder( shipmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFulfillmentCenter( request, shipmentId, FulfillmentCenterId ):
	delegate = ShipmentDelegate()
	responseData = delegate.saveFulfillmentCenter( shipmentId, FulfillmentCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFulfillmentCenter( request, shipmentId ):
	delegate = ShipmentDelegate()
	responseData = delegate.deleteFulfillmentCenter( shipmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addShipmentItems( request, shipmentId, ShipmentItemsIds ):
	delegate = ShipmentDelegate()
	responseData = delegate.addShipmentItems( shipmentId, ShipmentItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeShipmentItems( request, shipmentId, ShipmentItemsIds ):
	delegate = ShipmentDelegate()
	responseData = delegate.removeShipmentItems( shipmentId, ShipmentItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

