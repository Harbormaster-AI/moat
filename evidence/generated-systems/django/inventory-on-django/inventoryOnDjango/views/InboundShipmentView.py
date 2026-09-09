import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.InboundShipmentDelegate import InboundShipmentDelegate

 #======================================================================
# 
# Encapsulates data for View InboundShipment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InboundShipmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InboundShipment index.")

def get(request, inboundShipmentId ):
	delegate = InboundShipmentDelegate()
	responseData = delegate.get( inboundShipmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inboundShipment = json.loads(request.body)
	delegate = InboundShipmentDelegate()
	responseData = delegate.createFromJson( inboundShipment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inboundShipment = json.loads(request.body)
	delegate = InboundShipmentDelegate()
	responseData = delegate.save( inboundShipment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inboundShipmentId ):
	delegate = InboundShipmentDelegate()
	responseData = delegate.delete( inboundShipmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InboundShipmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, inboundShipmentId, WarehouseId ):
	delegate = InboundShipmentDelegate()
	responseData = delegate.saveWarehouse( inboundShipmentId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, inboundShipmentId ):
	delegate = InboundShipmentDelegate()
	responseData = delegate.deleteWarehouse( inboundShipmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLines( request, inboundShipmentId, LinesIds ):
	delegate = InboundShipmentDelegate()
	responseData = delegate.addLines( inboundShipmentId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLines( request, inboundShipmentId, LinesIds ):
	delegate = InboundShipmentDelegate()
	responseData = delegate.removeLines( inboundShipmentId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTransactions( request, inboundShipmentId, TransactionsIds ):
	delegate = InboundShipmentDelegate()
	responseData = delegate.addTransactions( inboundShipmentId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTransactions( request, inboundShipmentId, TransactionsIds ):
	delegate = InboundShipmentDelegate()
	responseData = delegate.removeTransactions( inboundShipmentId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

