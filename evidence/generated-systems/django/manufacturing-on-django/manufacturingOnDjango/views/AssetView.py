import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.AssetDelegate import AssetDelegate

 #======================================================================
# 
# Encapsulates data for View Asset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AssetView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Asset index.")

def get(request, assetId ):
	delegate = AssetDelegate()
	responseData = delegate.get( assetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	asset = json.loads(request.body)
	delegate = AssetDelegate()
	responseData = delegate.createFromJson( asset )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	asset = json.loads(request.body)
	delegate = AssetDelegate()
	responseData = delegate.save( asset )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, assetId ):
	delegate = AssetDelegate()
	responseData = delegate.delete( assetId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AssetDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlant( request, assetId, PlantId ):
	delegate = AssetDelegate()
	responseData = delegate.savePlant( assetId, PlantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlant( request, assetId ):
	delegate = AssetDelegate()
	responseData = delegate.deletePlant( assetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkCenter( request, assetId, WorkCenterId ):
	delegate = AssetDelegate()
	responseData = delegate.saveWorkCenter( assetId, WorkCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkCenter( request, assetId ):
	delegate = AssetDelegate()
	responseData = delegate.deleteWorkCenter( assetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMaintenanceOrders( request, assetId, MaintenanceOrdersIds ):
	delegate = AssetDelegate()
	responseData = delegate.addMaintenanceOrders( assetId, MaintenanceOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMaintenanceOrders( request, assetId, MaintenanceOrdersIds ):
	delegate = AssetDelegate()
	responseData = delegate.removeMaintenanceOrders( assetId, MaintenanceOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMaintenancePlans( request, assetId, MaintenancePlansIds ):
	delegate = AssetDelegate()
	responseData = delegate.addMaintenancePlans( assetId, MaintenancePlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMaintenancePlans( request, assetId, MaintenancePlansIds ):
	delegate = AssetDelegate()
	responseData = delegate.removeMaintenancePlans( assetId, MaintenancePlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

