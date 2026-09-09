import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.FlightHealthEventDelegate import FlightHealthEventDelegate

 #======================================================================
# 
# Encapsulates data for View FlightHealthEvent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FlightHealthEventView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the FlightHealthEvent index.")

def get(request, flightHealthEventId ):
	delegate = FlightHealthEventDelegate()
	responseData = delegate.get( flightHealthEventId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	flightHealthEvent = json.loads(request.body)
	delegate = FlightHealthEventDelegate()
	responseData = delegate.createFromJson( flightHealthEvent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	flightHealthEvent = json.loads(request.body)
	delegate = FlightHealthEventDelegate()
	responseData = delegate.save( flightHealthEvent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, flightHealthEventId ):
	delegate = FlightHealthEventDelegate()
	responseData = delegate.delete( flightHealthEventId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FlightHealthEventDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignConnectedAircraft( request, flightHealthEventId, ConnectedAircraftId ):
	delegate = FlightHealthEventDelegate()
	responseData = delegate.saveConnectedAircraft( flightHealthEventId, ConnectedAircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignConnectedAircraft( request, flightHealthEventId ):
	delegate = FlightHealthEventDelegate()
	responseData = delegate.deleteConnectedAircraft( flightHealthEventId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

