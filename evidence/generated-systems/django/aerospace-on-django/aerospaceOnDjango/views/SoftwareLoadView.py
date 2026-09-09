import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.SoftwareLoadDelegate import SoftwareLoadDelegate

 #======================================================================
# 
# Encapsulates data for View SoftwareLoad
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SoftwareLoadView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SoftwareLoad index.")

def get(request, softwareLoadId ):
	delegate = SoftwareLoadDelegate()
	responseData = delegate.get( softwareLoadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	softwareLoad = json.loads(request.body)
	delegate = SoftwareLoadDelegate()
	responseData = delegate.createFromJson( softwareLoad )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	softwareLoad = json.loads(request.body)
	delegate = SoftwareLoadDelegate()
	responseData = delegate.save( softwareLoad )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, softwareLoadId ):
	delegate = SoftwareLoadDelegate()
	responseData = delegate.delete( softwareLoadId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SoftwareLoadDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignConnectedAircraft( request, softwareLoadId, ConnectedAircraftId ):
	delegate = SoftwareLoadDelegate()
	responseData = delegate.saveConnectedAircraft( softwareLoadId, ConnectedAircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignConnectedAircraft( request, softwareLoadId ):
	delegate = SoftwareLoadDelegate()
	responseData = delegate.deleteConnectedAircraft( softwareLoadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAvionicsSuite( request, softwareLoadId, AvionicsSuiteId ):
	delegate = SoftwareLoadDelegate()
	responseData = delegate.saveAvionicsSuite( softwareLoadId, AvionicsSuiteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAvionicsSuite( request, softwareLoadId ):
	delegate = SoftwareLoadDelegate()
	responseData = delegate.deleteAvionicsSuite( softwareLoadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

