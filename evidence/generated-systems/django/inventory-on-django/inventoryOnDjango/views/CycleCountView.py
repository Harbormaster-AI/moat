import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.CycleCountDelegate import CycleCountDelegate

 #======================================================================
# 
# Encapsulates data for View CycleCount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CycleCountView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CycleCount index.")

def get(request, cycleCountId ):
	delegate = CycleCountDelegate()
	responseData = delegate.get( cycleCountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	cycleCount = json.loads(request.body)
	delegate = CycleCountDelegate()
	responseData = delegate.createFromJson( cycleCount )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	cycleCount = json.loads(request.body)
	delegate = CycleCountDelegate()
	responseData = delegate.save( cycleCount )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, cycleCountId ):
	delegate = CycleCountDelegate()
	responseData = delegate.delete( cycleCountId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CycleCountDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, cycleCountId, WarehouseId ):
	delegate = CycleCountDelegate()
	responseData = delegate.saveWarehouse( cycleCountId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, cycleCountId ):
	delegate = CycleCountDelegate()
	responseData = delegate.deleteWarehouse( cycleCountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLocations( request, cycleCountId, LocationsIds ):
	delegate = CycleCountDelegate()
	responseData = delegate.addLocations( cycleCountId, LocationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLocations( request, cycleCountId, LocationsIds ):
	delegate = CycleCountDelegate()
	responseData = delegate.removeLocations( cycleCountId, LocationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEntries( request, cycleCountId, EntriesIds ):
	delegate = CycleCountDelegate()
	responseData = delegate.addEntries( cycleCountId, EntriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEntries( request, cycleCountId, EntriesIds ):
	delegate = CycleCountDelegate()
	responseData = delegate.removeEntries( cycleCountId, EntriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTransactions( request, cycleCountId, TransactionsIds ):
	delegate = CycleCountDelegate()
	responseData = delegate.addTransactions( cycleCountId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTransactions( request, cycleCountId, TransactionsIds ):
	delegate = CycleCountDelegate()
	responseData = delegate.removeTransactions( cycleCountId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

