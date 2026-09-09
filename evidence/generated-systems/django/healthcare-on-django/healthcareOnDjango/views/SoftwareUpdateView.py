import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.SoftwareUpdateDelegate import SoftwareUpdateDelegate

 #======================================================================
# 
# Encapsulates data for View SoftwareUpdate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SoftwareUpdateView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SoftwareUpdate index.")

def get(request, softwareUpdateId ):
	delegate = SoftwareUpdateDelegate()
	responseData = delegate.get( softwareUpdateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	softwareUpdate = json.loads(request.body)
	delegate = SoftwareUpdateDelegate()
	responseData = delegate.createFromJson( softwareUpdate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	softwareUpdate = json.loads(request.body)
	delegate = SoftwareUpdateDelegate()
	responseData = delegate.save( softwareUpdate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, softwareUpdateId ):
	delegate = SoftwareUpdateDelegate()
	responseData = delegate.delete( softwareUpdateId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SoftwareUpdateDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDevice( request, softwareUpdateId, DeviceId ):
	delegate = SoftwareUpdateDelegate()
	responseData = delegate.saveDevice( softwareUpdateId, DeviceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDevice( request, softwareUpdateId ):
	delegate = SoftwareUpdateDelegate()
	responseData = delegate.deleteDevice( softwareUpdateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

