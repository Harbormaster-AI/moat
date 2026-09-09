import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.CabinLayoutDelegate import CabinLayoutDelegate

 #======================================================================
# 
# Encapsulates data for View CabinLayout
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CabinLayoutView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CabinLayout index.")

def get(request, cabinLayoutId ):
	delegate = CabinLayoutDelegate()
	responseData = delegate.get( cabinLayoutId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	cabinLayout = json.loads(request.body)
	delegate = CabinLayoutDelegate()
	responseData = delegate.createFromJson( cabinLayout )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	cabinLayout = json.loads(request.body)
	delegate = CabinLayoutDelegate()
	responseData = delegate.save( cabinLayout )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, cabinLayoutId ):
	delegate = CabinLayoutDelegate()
	responseData = delegate.delete( cabinLayoutId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CabinLayoutDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignVariant( request, cabinLayoutId, VariantId ):
	delegate = CabinLayoutDelegate()
	responseData = delegate.saveVariant( cabinLayoutId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, cabinLayoutId ):
	delegate = CabinLayoutDelegate()
	responseData = delegate.deleteVariant( cabinLayoutId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAircraft( request, cabinLayoutId, AircraftIds ):
	delegate = CabinLayoutDelegate()
	responseData = delegate.addAircraft( cabinLayoutId, AircraftIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAircraft( request, cabinLayoutId, AircraftIds ):
	delegate = CabinLayoutDelegate()
	responseData = delegate.removeAircraft( cabinLayoutId, AircraftIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOptions( request, cabinLayoutId, OptionsIds ):
	delegate = CabinLayoutDelegate()
	responseData = delegate.addOptions( cabinLayoutId, OptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOptions( request, cabinLayoutId, OptionsIds ):
	delegate = CabinLayoutDelegate()
	responseData = delegate.removeOptions( cabinLayoutId, OptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

