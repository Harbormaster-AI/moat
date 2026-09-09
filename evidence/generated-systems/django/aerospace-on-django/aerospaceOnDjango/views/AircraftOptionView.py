import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AircraftOptionDelegate import AircraftOptionDelegate

 #======================================================================
# 
# Encapsulates data for View AircraftOption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftOptionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AircraftOption index.")

def get(request, aircraftOptionId ):
	delegate = AircraftOptionDelegate()
	responseData = delegate.get( aircraftOptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aircraftOption = json.loads(request.body)
	delegate = AircraftOptionDelegate()
	responseData = delegate.createFromJson( aircraftOption )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aircraftOption = json.loads(request.body)
	delegate = AircraftOptionDelegate()
	responseData = delegate.save( aircraftOption )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aircraftOptionId ):
	delegate = AircraftOptionDelegate()
	responseData = delegate.delete( aircraftOptionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AircraftOptionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVariants( request, aircraftOptionId, VariantsIds ):
	delegate = AircraftOptionDelegate()
	responseData = delegate.addVariants( aircraftOptionId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVariants( request, aircraftOptionId, VariantsIds ):
	delegate = AircraftOptionDelegate()
	responseData = delegate.removeVariants( aircraftOptionId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPackages( request, aircraftOptionId, PackagesIds ):
	delegate = AircraftOptionDelegate()
	responseData = delegate.addPackages( aircraftOptionId, PackagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePackages( request, aircraftOptionId, PackagesIds ):
	delegate = AircraftOptionDelegate()
	responseData = delegate.removePackages( aircraftOptionId, PackagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

