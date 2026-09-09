import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.ShipmentItemDelegate import ShipmentItemDelegate

 #======================================================================
# 
# Encapsulates data for View ShipmentItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShipmentItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ShipmentItem index.")

def get(request, shipmentItemId ):
	delegate = ShipmentItemDelegate()
	responseData = delegate.get( shipmentItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	shipmentItem = json.loads(request.body)
	delegate = ShipmentItemDelegate()
	responseData = delegate.createFromJson( shipmentItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	shipmentItem = json.loads(request.body)
	delegate = ShipmentItemDelegate()
	responseData = delegate.save( shipmentItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, shipmentItemId ):
	delegate = ShipmentItemDelegate()
	responseData = delegate.delete( shipmentItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ShipmentItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignShipment( request, shipmentItemId, ShipmentId ):
	delegate = ShipmentItemDelegate()
	responseData = delegate.saveShipment( shipmentItemId, ShipmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignShipment( request, shipmentItemId ):
	delegate = ShipmentItemDelegate()
	responseData = delegate.deleteShipment( shipmentItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrderLine( request, shipmentItemId, OrderLineId ):
	delegate = ShipmentItemDelegate()
	responseData = delegate.saveOrderLine( shipmentItemId, OrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrderLine( request, shipmentItemId ):
	delegate = ShipmentItemDelegate()
	responseData = delegate.deleteOrderLine( shipmentItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

