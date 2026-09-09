import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.MaintenanceOrderDelegate import MaintenanceOrderDelegate

 #======================================================================
# 
# Encapsulates data for View MaintenanceOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MaintenanceOrder index.")

def get(request, maintenanceOrderId ):
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.get( maintenanceOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	maintenanceOrder = json.loads(request.body)
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.createFromJson( maintenanceOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	maintenanceOrder = json.loads(request.body)
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.save( maintenanceOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, maintenanceOrderId ):
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.delete( maintenanceOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAsset( request, maintenanceOrderId, AssetId ):
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.saveAsset( maintenanceOrderId, AssetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAsset( request, maintenanceOrderId ):
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.deleteAsset( maintenanceOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlan( request, maintenanceOrderId, PlanId ):
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.savePlan( maintenanceOrderId, PlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlan( request, maintenanceOrderId ):
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.deletePlan( maintenanceOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkCenter( request, maintenanceOrderId, WorkCenterId ):
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.saveWorkCenter( maintenanceOrderId, WorkCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkCenter( request, maintenanceOrderId ):
	delegate = MaintenanceOrderDelegate()
	responseData = delegate.deleteWorkCenter( maintenanceOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

