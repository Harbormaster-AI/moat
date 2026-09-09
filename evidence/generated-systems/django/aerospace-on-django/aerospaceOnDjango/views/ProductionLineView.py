import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.ProductionLineDelegate import ProductionLineDelegate

 #======================================================================
# 
# Encapsulates data for View ProductionLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionLineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ProductionLine index.")

def get(request, productionLineId ):
	delegate = ProductionLineDelegate()
	responseData = delegate.get( productionLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	productionLine = json.loads(request.body)
	delegate = ProductionLineDelegate()
	responseData = delegate.createFromJson( productionLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	productionLine = json.loads(request.body)
	delegate = ProductionLineDelegate()
	responseData = delegate.save( productionLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, productionLineId ):
	delegate = ProductionLineDelegate()
	responseData = delegate.delete( productionLineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProductionLineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlant( request, productionLineId, PlantId ):
	delegate = ProductionLineDelegate()
	responseData = delegate.savePlant( productionLineId, PlantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlant( request, productionLineId ):
	delegate = ProductionLineDelegate()
	responseData = delegate.deletePlant( productionLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWorkCenters( request, productionLineId, WorkCentersIds ):
	delegate = ProductionLineDelegate()
	responseData = delegate.addWorkCenters( productionLineId, WorkCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWorkCenters( request, productionLineId, WorkCentersIds ):
	delegate = ProductionLineDelegate()
	responseData = delegate.removeWorkCenters( productionLineId, WorkCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

