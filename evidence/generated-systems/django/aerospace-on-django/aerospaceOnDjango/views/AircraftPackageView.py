import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AircraftPackageDelegate import AircraftPackageDelegate

 #======================================================================
# 
# Encapsulates data for View AircraftPackage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftPackageView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AircraftPackage index.")

def get(request, aircraftPackageId ):
	delegate = AircraftPackageDelegate()
	responseData = delegate.get( aircraftPackageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aircraftPackage = json.loads(request.body)
	delegate = AircraftPackageDelegate()
	responseData = delegate.createFromJson( aircraftPackage )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aircraftPackage = json.loads(request.body)
	delegate = AircraftPackageDelegate()
	responseData = delegate.save( aircraftPackage )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aircraftPackageId ):
	delegate = AircraftPackageDelegate()
	responseData = delegate.delete( aircraftPackageId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AircraftPackageDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOptions( request, aircraftPackageId, OptionsIds ):
	delegate = AircraftPackageDelegate()
	responseData = delegate.addOptions( aircraftPackageId, OptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOptions( request, aircraftPackageId, OptionsIds ):
	delegate = AircraftPackageDelegate()
	responseData = delegate.removeOptions( aircraftPackageId, OptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVariants( request, aircraftPackageId, VariantsIds ):
	delegate = AircraftPackageDelegate()
	responseData = delegate.addVariants( aircraftPackageId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVariants( request, aircraftPackageId, VariantsIds ):
	delegate = AircraftPackageDelegate()
	responseData = delegate.removeVariants( aircraftPackageId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

