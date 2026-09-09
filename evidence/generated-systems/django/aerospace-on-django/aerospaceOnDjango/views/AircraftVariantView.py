import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

 #======================================================================
# 
# Encapsulates data for View AircraftVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftVariantView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AircraftVariant index.")

def get(request, aircraftVariantId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.get( aircraftVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aircraftVariant = json.loads(request.body)
	delegate = AircraftVariantDelegate()
	responseData = delegate.createFromJson( aircraftVariant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aircraftVariant = json.loads(request.body)
	delegate = AircraftVariantDelegate()
	responseData = delegate.save( aircraftVariant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aircraftVariantId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.delete( aircraftVariantId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AircraftVariantDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignModel( request, aircraftVariantId, ModelId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.saveModel( aircraftVariantId, ModelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignModel( request, aircraftVariantId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.deleteModel( aircraftVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEngineType( request, aircraftVariantId, EngineTypeId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.saveEngineType( aircraftVariantId, EngineTypeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEngineType( request, aircraftVariantId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.deleteEngineType( aircraftVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAvionicsSuite( request, aircraftVariantId, AvionicsSuiteId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.saveAvionicsSuite( aircraftVariantId, AvionicsSuiteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAvionicsSuite( request, aircraftVariantId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.deleteAvionicsSuite( aircraftVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignApu( request, aircraftVariantId, ApuId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.saveApu( aircraftVariantId, ApuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignApu( request, aircraftVariantId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.deleteApu( aircraftVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLandingGear( request, aircraftVariantId, LandingGearId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.saveLandingGear( aircraftVariantId, LandingGearId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLandingGear( request, aircraftVariantId ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.deleteLandingGear( aircraftVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCabinLayouts( request, aircraftVariantId, CabinLayoutsIds ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.addCabinLayouts( aircraftVariantId, CabinLayoutsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCabinLayouts( request, aircraftVariantId, CabinLayoutsIds ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.removeCabinLayouts( aircraftVariantId, CabinLayoutsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOptions( request, aircraftVariantId, OptionsIds ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.addOptions( aircraftVariantId, OptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOptions( request, aircraftVariantId, OptionsIds ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.removeOptions( aircraftVariantId, OptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPackages( request, aircraftVariantId, PackagesIds ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.addPackages( aircraftVariantId, PackagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePackages( request, aircraftVariantId, PackagesIds ):
	delegate = AircraftVariantDelegate()
	responseData = delegate.removePackages( aircraftVariantId, PackagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

