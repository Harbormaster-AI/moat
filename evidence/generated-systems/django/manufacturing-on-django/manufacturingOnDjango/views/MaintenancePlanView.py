import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.MaintenancePlanDelegate import MaintenancePlanDelegate

 #======================================================================
# 
# Encapsulates data for View MaintenancePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenancePlanView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MaintenancePlan index.")

def get(request, maintenancePlanId ):
	delegate = MaintenancePlanDelegate()
	responseData = delegate.get( maintenancePlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	maintenancePlan = json.loads(request.body)
	delegate = MaintenancePlanDelegate()
	responseData = delegate.createFromJson( maintenancePlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	maintenancePlan = json.loads(request.body)
	delegate = MaintenancePlanDelegate()
	responseData = delegate.save( maintenancePlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, maintenancePlanId ):
	delegate = MaintenancePlanDelegate()
	responseData = delegate.delete( maintenancePlanId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MaintenancePlanDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAsset( request, maintenancePlanId, AssetId ):
	delegate = MaintenancePlanDelegate()
	responseData = delegate.saveAsset( maintenancePlanId, AssetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAsset( request, maintenancePlanId ):
	delegate = MaintenancePlanDelegate()
	responseData = delegate.deleteAsset( maintenancePlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMaintenanceOrders( request, maintenancePlanId, MaintenanceOrdersIds ):
	delegate = MaintenancePlanDelegate()
	responseData = delegate.addMaintenanceOrders( maintenancePlanId, MaintenanceOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMaintenanceOrders( request, maintenancePlanId, MaintenanceOrdersIds ):
	delegate = MaintenancePlanDelegate()
	responseData = delegate.removeMaintenanceOrders( maintenancePlanId, MaintenanceOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

