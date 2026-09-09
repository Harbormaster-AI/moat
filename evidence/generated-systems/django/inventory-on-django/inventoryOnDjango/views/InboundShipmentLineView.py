import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.InboundShipmentLineDelegate import InboundShipmentLineDelegate

 #======================================================================
# 
# Encapsulates data for View InboundShipmentLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InboundShipmentLineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InboundShipmentLine index.")

def get(request, inboundShipmentLineId ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.get( inboundShipmentLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inboundShipmentLine = json.loads(request.body)
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.createFromJson( inboundShipmentLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inboundShipmentLine = json.loads(request.body)
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.save( inboundShipmentLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inboundShipmentLineId ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.delete( inboundShipmentLineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInboundShipment( request, inboundShipmentLineId, InboundShipmentId ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.saveInboundShipment( inboundShipmentLineId, InboundShipmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInboundShipment( request, inboundShipmentLineId ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.deleteInboundShipment( inboundShipmentLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, inboundShipmentLineId, SkuId ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.saveSku( inboundShipmentLineId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, inboundShipmentLineId ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.deleteSku( inboundShipmentLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLot( request, inboundShipmentLineId, LotId ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.saveLot( inboundShipmentLineId, LotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLot( request, inboundShipmentLineId ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.deleteLot( inboundShipmentLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDestinationLocation( request, inboundShipmentLineId, DestinationLocationId ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.saveDestinationLocation( inboundShipmentLineId, DestinationLocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDestinationLocation( request, inboundShipmentLineId ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.deleteDestinationLocation( inboundShipmentLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSerialNumbers( request, inboundShipmentLineId, SerialNumbersIds ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.addSerialNumbers( inboundShipmentLineId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSerialNumbers( request, inboundShipmentLineId, SerialNumbersIds ):
	delegate = InboundShipmentLineDelegate()
	responseData = delegate.removeSerialNumbers( inboundShipmentLineId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

