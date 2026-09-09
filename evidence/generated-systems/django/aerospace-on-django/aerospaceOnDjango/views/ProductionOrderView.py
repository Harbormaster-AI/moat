import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.ProductionOrderDelegate import ProductionOrderDelegate

 #======================================================================
# 
# Encapsulates data for View ProductionOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ProductionOrder index.")

def get(request, productionOrderId ):
	delegate = ProductionOrderDelegate()
	responseData = delegate.get( productionOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	productionOrder = json.loads(request.body)
	delegate = ProductionOrderDelegate()
	responseData = delegate.createFromJson( productionOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	productionOrder = json.loads(request.body)
	delegate = ProductionOrderDelegate()
	responseData = delegate.save( productionOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, productionOrderId ):
	delegate = ProductionOrderDelegate()
	responseData = delegate.delete( productionOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProductionOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignVariant( request, productionOrderId, VariantId ):
	delegate = ProductionOrderDelegate()
	responseData = delegate.saveVariant( productionOrderId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, productionOrderId ):
	delegate = ProductionOrderDelegate()
	responseData = delegate.deleteVariant( productionOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlant( request, productionOrderId, PlantId ):
	delegate = ProductionOrderDelegate()
	responseData = delegate.savePlant( productionOrderId, PlantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlant( request, productionOrderId ):
	delegate = ProductionOrderDelegate()
	responseData = delegate.deletePlant( productionOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAircraftOrder( request, productionOrderId, AircraftOrderId ):
	delegate = ProductionOrderDelegate()
	responseData = delegate.saveAircraftOrder( productionOrderId, AircraftOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAircraftOrder( request, productionOrderId ):
	delegate = ProductionOrderDelegate()
	responseData = delegate.deleteAircraftOrder( productionOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

