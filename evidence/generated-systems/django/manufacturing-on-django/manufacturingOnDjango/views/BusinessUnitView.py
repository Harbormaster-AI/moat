import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.BusinessUnitDelegate import BusinessUnitDelegate

 #======================================================================
# 
# Encapsulates data for View BusinessUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessUnitView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BusinessUnit index.")

def get(request, businessUnitId ):
	delegate = BusinessUnitDelegate()
	responseData = delegate.get( businessUnitId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	businessUnit = json.loads(request.body)
	delegate = BusinessUnitDelegate()
	responseData = delegate.createFromJson( businessUnit )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	businessUnit = json.loads(request.body)
	delegate = BusinessUnitDelegate()
	responseData = delegate.save( businessUnit )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, businessUnitId ):
	delegate = BusinessUnitDelegate()
	responseData = delegate.delete( businessUnitId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BusinessUnitDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEnterprise( request, businessUnitId, EnterpriseId ):
	delegate = BusinessUnitDelegate()
	responseData = delegate.saveEnterprise( businessUnitId, EnterpriseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEnterprise( request, businessUnitId ):
	delegate = BusinessUnitDelegate()
	responseData = delegate.deleteEnterprise( businessUnitId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addItems( request, businessUnitId, ItemsIds ):
	delegate = BusinessUnitDelegate()
	responseData = delegate.addItems( businessUnitId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeItems( request, businessUnitId, ItemsIds ):
	delegate = BusinessUnitDelegate()
	responseData = delegate.removeItems( businessUnitId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPlants( request, businessUnitId, PlantsIds ):
	delegate = BusinessUnitDelegate()
	responseData = delegate.addPlants( businessUnitId, PlantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePlants( request, businessUnitId, PlantsIds ):
	delegate = BusinessUnitDelegate()
	responseData = delegate.removePlants( businessUnitId, PlantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

