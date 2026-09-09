import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

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

def assignEnterprise( request, plantId, EnterpriseId ):
	delegate = PlantDelegate()
	responseData = delegate.saveEnterprise( plantId, EnterpriseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEnterprise( request, plantId ):
	delegate = PlantDelegate()
	responseData = delegate.deleteEnterprise( plantId )
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

def addWorkCenters( request, plantId, WorkCentersIds ):
	delegate = PlantDelegate()
	responseData = delegate.addWorkCenters( plantId, WorkCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWorkCenters( request, plantId, WorkCentersIds ):
	delegate = PlantDelegate()
	responseData = delegate.removeWorkCenters( plantId, WorkCentersIds )
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

def addAssets( request, plantId, AssetsIds ):
	delegate = PlantDelegate()
	responseData = delegate.addAssets( plantId, AssetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAssets( request, plantId, AssetsIds ):
	delegate = PlantDelegate()
	responseData = delegate.removeAssets( plantId, AssetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProductionSchedules( request, plantId, ProductionSchedulesIds ):
	delegate = PlantDelegate()
	responseData = delegate.addProductionSchedules( plantId, ProductionSchedulesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProductionSchedules( request, plantId, ProductionSchedulesIds ):
	delegate = PlantDelegate()
	responseData = delegate.removeProductionSchedules( plantId, ProductionSchedulesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

