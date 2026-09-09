import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

 #======================================================================
# 
# Encapsulates data for View WorkCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkCenterView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the WorkCenter index.")

def get(request, workCenterId ):
	delegate = WorkCenterDelegate()
	responseData = delegate.get( workCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	workCenter = json.loads(request.body)
	delegate = WorkCenterDelegate()
	responseData = delegate.createFromJson( workCenter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	workCenter = json.loads(request.body)
	delegate = WorkCenterDelegate()
	responseData = delegate.save( workCenter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, workCenterId ):
	delegate = WorkCenterDelegate()
	responseData = delegate.delete( workCenterId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WorkCenterDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProductionLine( request, workCenterId, ProductionLineId ):
	delegate = WorkCenterDelegate()
	responseData = delegate.saveProductionLine( workCenterId, ProductionLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProductionLine( request, workCenterId ):
	delegate = WorkCenterDelegate()
	responseData = delegate.deleteProductionLine( workCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAssets( request, workCenterId, AssetsIds ):
	delegate = WorkCenterDelegate()
	responseData = delegate.addAssets( workCenterId, AssetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAssets( request, workCenterId, AssetsIds ):
	delegate = WorkCenterDelegate()
	responseData = delegate.removeAssets( workCenterId, AssetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMaintenanceOrders( request, workCenterId, MaintenanceOrdersIds ):
	delegate = WorkCenterDelegate()
	responseData = delegate.addMaintenanceOrders( workCenterId, MaintenanceOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMaintenanceOrders( request, workCenterId, MaintenanceOrdersIds ):
	delegate = WorkCenterDelegate()
	responseData = delegate.removeMaintenanceOrders( workCenterId, MaintenanceOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

