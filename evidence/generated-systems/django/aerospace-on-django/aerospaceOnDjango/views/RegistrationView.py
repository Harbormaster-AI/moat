import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.RegistrationDelegate import RegistrationDelegate

 #======================================================================
# 
# Encapsulates data for View Registration
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RegistrationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Registration index.")

def get(request, registrationId ):
	delegate = RegistrationDelegate()
	responseData = delegate.get( registrationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	registration = json.loads(request.body)
	delegate = RegistrationDelegate()
	responseData = delegate.createFromJson( registration )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	registration = json.loads(request.body)
	delegate = RegistrationDelegate()
	responseData = delegate.save( registration )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, registrationId ):
	delegate = RegistrationDelegate()
	responseData = delegate.delete( registrationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RegistrationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAircraft( request, registrationId, AircraftId ):
	delegate = RegistrationDelegate()
	responseData = delegate.saveAircraft( registrationId, AircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAircraft( request, registrationId ):
	delegate = RegistrationDelegate()
	responseData = delegate.deleteAircraft( registrationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

