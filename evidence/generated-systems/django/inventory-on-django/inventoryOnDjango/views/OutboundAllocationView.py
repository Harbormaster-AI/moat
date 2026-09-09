import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.OutboundAllocationDelegate import OutboundAllocationDelegate

 #======================================================================
# 
# Encapsulates data for View OutboundAllocation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OutboundAllocationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the OutboundAllocation index.")

def get(request, outboundAllocationId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.get( outboundAllocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	outboundAllocation = json.loads(request.body)
	delegate = OutboundAllocationDelegate()
	responseData = delegate.createFromJson( outboundAllocation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	outboundAllocation = json.loads(request.body)
	delegate = OutboundAllocationDelegate()
	responseData = delegate.save( outboundAllocation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, outboundAllocationId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.delete( outboundAllocationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, outboundAllocationId, WarehouseId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.saveWarehouse( outboundAllocationId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, outboundAllocationId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.deleteWarehouse( outboundAllocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, outboundAllocationId, SkuId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.saveSku( outboundAllocationId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, outboundAllocationId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.deleteSku( outboundAllocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInventoryItem( request, outboundAllocationId, InventoryItemId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.saveInventoryItem( outboundAllocationId, InventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInventoryItem( request, outboundAllocationId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.deleteInventoryItem( outboundAllocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignReservation( request, outboundAllocationId, ReservationId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.saveReservation( outboundAllocationId, ReservationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignReservation( request, outboundAllocationId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.deleteReservation( outboundAllocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLot( request, outboundAllocationId, LotId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.saveLot( outboundAllocationId, LotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLot( request, outboundAllocationId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.deleteLot( outboundAllocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSourceLocation( request, outboundAllocationId, SourceLocationId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.saveSourceLocation( outboundAllocationId, SourceLocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSourceLocation( request, outboundAllocationId ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.deleteSourceLocation( outboundAllocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSerialNumbers( request, outboundAllocationId, SerialNumbersIds ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.addSerialNumbers( outboundAllocationId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSerialNumbers( request, outboundAllocationId, SerialNumbersIds ):
	delegate = OutboundAllocationDelegate()
	responseData = delegate.removeSerialNumbers( outboundAllocationId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

