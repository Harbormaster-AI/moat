import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.PlannedOrderDelegate import PlannedOrderDelegate

 #======================================================================
# 
# Encapsulates data for View PlannedOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlannedOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PlannedOrder index.")

def get(request, plannedOrderId ):
	delegate = PlannedOrderDelegate()
	responseData = delegate.get( plannedOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	plannedOrder = json.loads(request.body)
	delegate = PlannedOrderDelegate()
	responseData = delegate.createFromJson( plannedOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	plannedOrder = json.loads(request.body)
	delegate = PlannedOrderDelegate()
	responseData = delegate.save( plannedOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, plannedOrderId ):
	delegate = PlannedOrderDelegate()
	responseData = delegate.delete( plannedOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PlannedOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMrpRun( request, plannedOrderId, MrpRunId ):
	delegate = PlannedOrderDelegate()
	responseData = delegate.saveMrpRun( plannedOrderId, MrpRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMrpRun( request, plannedOrderId ):
	delegate = PlannedOrderDelegate()
	responseData = delegate.deleteMrpRun( plannedOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, plannedOrderId, ItemId ):
	delegate = PlannedOrderDelegate()
	responseData = delegate.saveItem( plannedOrderId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, plannedOrderId ):
	delegate = PlannedOrderDelegate()
	responseData = delegate.deleteItem( plannedOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlant( request, plannedOrderId, PlantId ):
	delegate = PlannedOrderDelegate()
	responseData = delegate.savePlant( plannedOrderId, PlantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlant( request, plannedOrderId ):
	delegate = PlannedOrderDelegate()
	responseData = delegate.deletePlant( plannedOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

