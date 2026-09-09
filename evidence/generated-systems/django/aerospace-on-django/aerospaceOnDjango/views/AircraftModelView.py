import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AircraftModelDelegate import AircraftModelDelegate

 #======================================================================
# 
# Encapsulates data for View AircraftModel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftModelView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AircraftModel index.")

def get(request, aircraftModelId ):
	delegate = AircraftModelDelegate()
	responseData = delegate.get( aircraftModelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aircraftModel = json.loads(request.body)
	delegate = AircraftModelDelegate()
	responseData = delegate.createFromJson( aircraftModel )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aircraftModel = json.loads(request.body)
	delegate = AircraftModelDelegate()
	responseData = delegate.save( aircraftModel )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aircraftModelId ):
	delegate = AircraftModelDelegate()
	responseData = delegate.delete( aircraftModelId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AircraftModelDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFamily( request, aircraftModelId, FamilyId ):
	delegate = AircraftModelDelegate()
	responseData = delegate.saveFamily( aircraftModelId, FamilyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFamily( request, aircraftModelId ):
	delegate = AircraftModelDelegate()
	responseData = delegate.deleteFamily( aircraftModelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVariants( request, aircraftModelId, VariantsIds ):
	delegate = AircraftModelDelegate()
	responseData = delegate.addVariants( aircraftModelId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVariants( request, aircraftModelId, VariantsIds ):
	delegate = AircraftModelDelegate()
	responseData = delegate.removeVariants( aircraftModelId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEngineTypes( request, aircraftModelId, EngineTypesIds ):
	delegate = AircraftModelDelegate()
	responseData = delegate.addEngineTypes( aircraftModelId, EngineTypesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEngineTypes( request, aircraftModelId, EngineTypesIds ):
	delegate = AircraftModelDelegate()
	responseData = delegate.removeEngineTypes( aircraftModelId, EngineTypesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

