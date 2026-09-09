import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.TransferOrderDelegate import TransferOrderDelegate

 #======================================================================
# 
# Encapsulates data for View TransferOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransferOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TransferOrder index.")

def get(request, transferOrderId ):
	delegate = TransferOrderDelegate()
	responseData = delegate.get( transferOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	transferOrder = json.loads(request.body)
	delegate = TransferOrderDelegate()
	responseData = delegate.createFromJson( transferOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	transferOrder = json.loads(request.body)
	delegate = TransferOrderDelegate()
	responseData = delegate.save( transferOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, transferOrderId ):
	delegate = TransferOrderDelegate()
	responseData = delegate.delete( transferOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TransferOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOriginWarehouse( request, transferOrderId, OriginWarehouseId ):
	delegate = TransferOrderDelegate()
	responseData = delegate.saveOriginWarehouse( transferOrderId, OriginWarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOriginWarehouse( request, transferOrderId ):
	delegate = TransferOrderDelegate()
	responseData = delegate.deleteOriginWarehouse( transferOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDestinationWarehouse( request, transferOrderId, DestinationWarehouseId ):
	delegate = TransferOrderDelegate()
	responseData = delegate.saveDestinationWarehouse( transferOrderId, DestinationWarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDestinationWarehouse( request, transferOrderId ):
	delegate = TransferOrderDelegate()
	responseData = delegate.deleteDestinationWarehouse( transferOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLines( request, transferOrderId, LinesIds ):
	delegate = TransferOrderDelegate()
	responseData = delegate.addLines( transferOrderId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLines( request, transferOrderId, LinesIds ):
	delegate = TransferOrderDelegate()
	responseData = delegate.removeLines( transferOrderId, LinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTransactions( request, transferOrderId, TransactionsIds ):
	delegate = TransferOrderDelegate()
	responseData = delegate.addTransactions( transferOrderId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTransactions( request, transferOrderId, TransactionsIds ):
	delegate = TransferOrderDelegate()
	responseData = delegate.removeTransactions( transferOrderId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

