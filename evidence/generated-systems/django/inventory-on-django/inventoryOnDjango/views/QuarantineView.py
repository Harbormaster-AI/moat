import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.QuarantineDelegate import QuarantineDelegate

 #======================================================================
# 
# Encapsulates data for View Quarantine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuarantineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Quarantine index.")

def get(request, quarantineId ):
	delegate = QuarantineDelegate()
	responseData = delegate.get( quarantineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	quarantine = json.loads(request.body)
	delegate = QuarantineDelegate()
	responseData = delegate.createFromJson( quarantine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	quarantine = json.loads(request.body)
	delegate = QuarantineDelegate()
	responseData = delegate.save( quarantine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, quarantineId ):
	delegate = QuarantineDelegate()
	responseData = delegate.delete( quarantineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = QuarantineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, quarantineId, WarehouseId ):
	delegate = QuarantineDelegate()
	responseData = delegate.saveWarehouse( quarantineId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, quarantineId ):
	delegate = QuarantineDelegate()
	responseData = delegate.deleteWarehouse( quarantineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLot( request, quarantineId, LotId ):
	delegate = QuarantineDelegate()
	responseData = delegate.saveLot( quarantineId, LotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLot( request, quarantineId ):
	delegate = QuarantineDelegate()
	responseData = delegate.deleteLot( quarantineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addItems( request, quarantineId, ItemsIds ):
	delegate = QuarantineDelegate()
	responseData = delegate.addItems( quarantineId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeItems( request, quarantineId, ItemsIds ):
	delegate = QuarantineDelegate()
	responseData = delegate.removeItems( quarantineId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSerialNumbers( request, quarantineId, SerialNumbersIds ):
	delegate = QuarantineDelegate()
	responseData = delegate.addSerialNumbers( quarantineId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSerialNumbers( request, quarantineId, SerialNumbersIds ):
	delegate = QuarantineDelegate()
	responseData = delegate.removeSerialNumbers( quarantineId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

