import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AircraftFamilyDelegate import AircraftFamilyDelegate

 #======================================================================
# 
# Encapsulates data for View AircraftFamily
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftFamilyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AircraftFamily index.")

def get(request, aircraftFamilyId ):
	delegate = AircraftFamilyDelegate()
	responseData = delegate.get( aircraftFamilyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aircraftFamily = json.loads(request.body)
	delegate = AircraftFamilyDelegate()
	responseData = delegate.createFromJson( aircraftFamily )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aircraftFamily = json.loads(request.body)
	delegate = AircraftFamilyDelegate()
	responseData = delegate.save( aircraftFamily )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aircraftFamilyId ):
	delegate = AircraftFamilyDelegate()
	responseData = delegate.delete( aircraftFamilyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AircraftFamilyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProgram( request, aircraftFamilyId, ProgramId ):
	delegate = AircraftFamilyDelegate()
	responseData = delegate.saveProgram( aircraftFamilyId, ProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProgram( request, aircraftFamilyId ):
	delegate = AircraftFamilyDelegate()
	responseData = delegate.deleteProgram( aircraftFamilyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAircraftModels( request, aircraftFamilyId, AircraftModelsIds ):
	delegate = AircraftFamilyDelegate()
	responseData = delegate.addAircraftModels( aircraftFamilyId, AircraftModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAircraftModels( request, aircraftFamilyId, AircraftModelsIds ):
	delegate = AircraftFamilyDelegate()
	responseData = delegate.removeAircraftModels( aircraftFamilyId, AircraftModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

