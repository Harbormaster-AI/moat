import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.TransferOrderLineDelegate import TransferOrderLineDelegate

 #======================================================================
# 
# Encapsulates data for View TransferOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransferOrderLineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TransferOrderLine index.")

def get(request, transferOrderLineId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.get( transferOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	transferOrderLine = json.loads(request.body)
	delegate = TransferOrderLineDelegate()
	responseData = delegate.createFromJson( transferOrderLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	transferOrderLine = json.loads(request.body)
	delegate = TransferOrderLineDelegate()
	responseData = delegate.save( transferOrderLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, transferOrderLineId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.delete( transferOrderLineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTransferOrder( request, transferOrderLineId, TransferOrderId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.saveTransferOrder( transferOrderLineId, TransferOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTransferOrder( request, transferOrderLineId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.deleteTransferOrder( transferOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, transferOrderLineId, SkuId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.saveSku( transferOrderLineId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, transferOrderLineId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.deleteSku( transferOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLot( request, transferOrderLineId, LotId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.saveLot( transferOrderLineId, LotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLot( request, transferOrderLineId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.deleteLot( transferOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFromLocation( request, transferOrderLineId, FromLocationId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.saveFromLocation( transferOrderLineId, FromLocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFromLocation( request, transferOrderLineId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.deleteFromLocation( transferOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignToLocation( request, transferOrderLineId, ToLocationId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.saveToLocation( transferOrderLineId, ToLocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignToLocation( request, transferOrderLineId ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.deleteToLocation( transferOrderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSerialNumbers( request, transferOrderLineId, SerialNumbersIds ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.addSerialNumbers( transferOrderLineId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSerialNumbers( request, transferOrderLineId, SerialNumbersIds ):
	delegate = TransferOrderLineDelegate()
	responseData = delegate.removeSerialNumbers( transferOrderLineId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

