import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.PlantDelegate import PlantDelegate

 #======================================================================
# 
# Encapsulates data for View Plant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlantView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Plant index.")

def get(request, plantId ):
	delegate = PlantDelegate()
	responseData = delegate.get( plantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	plant = json.loads(request.body)
	delegate = PlantDelegate()
	responseData = delegate.createFromJson( plant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	plant = json.loads(request.body)
	delegate = PlantDelegate()
	responseData = delegate.save( plant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, plantId ):
	delegate = PlantDelegate()
	responseData = delegate.delete( plantId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PlantDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignManufacturer( request, plantId, ManufacturerId ):
	delegate = PlantDelegate()
	responseData = delegate.saveManufacturer( plantId, ManufacturerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignManufacturer( request, plantId ):
	delegate = PlantDelegate()
	responseData = delegate.deleteManufacturer( plantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProductionLines( request, plantId, ProductionLinesIds ):
	delegate = PlantDelegate()
	responseData = delegate.addProductionLines( plantId, ProductionLinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProductionLines( request, plantId, ProductionLinesIds ):
	delegate = PlantDelegate()
	responseData = delegate.removeProductionLines( plantId, ProductionLinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWarehouses( request, plantId, WarehousesIds ):
	delegate = PlantDelegate()
	responseData = delegate.addWarehouses( plantId, WarehousesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWarehouses( request, plantId, WarehousesIds ):
	delegate = PlantDelegate()
	responseData = delegate.removeWarehouses( plantId, WarehousesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

