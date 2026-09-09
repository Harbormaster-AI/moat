import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.ProductionScheduleDelegate import ProductionScheduleDelegate

 #======================================================================
# 
# Encapsulates data for View ProductionSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionScheduleView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ProductionSchedule index.")

def get(request, productionScheduleId ):
	delegate = ProductionScheduleDelegate()
	responseData = delegate.get( productionScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	productionSchedule = json.loads(request.body)
	delegate = ProductionScheduleDelegate()
	responseData = delegate.createFromJson( productionSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	productionSchedule = json.loads(request.body)
	delegate = ProductionScheduleDelegate()
	responseData = delegate.save( productionSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, productionScheduleId ):
	delegate = ProductionScheduleDelegate()
	responseData = delegate.delete( productionScheduleId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProductionScheduleDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlant( request, productionScheduleId, PlantId ):
	delegate = ProductionScheduleDelegate()
	responseData = delegate.savePlant( productionScheduleId, PlantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlant( request, productionScheduleId ):
	delegate = ProductionScheduleDelegate()
	responseData = delegate.deletePlant( productionScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWorkOrders( request, productionScheduleId, WorkOrdersIds ):
	delegate = ProductionScheduleDelegate()
	responseData = delegate.addWorkOrders( productionScheduleId, WorkOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWorkOrders( request, productionScheduleId, WorkOrdersIds ):
	delegate = ProductionScheduleDelegate()
	responseData = delegate.removeWorkOrders( productionScheduleId, WorkOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

