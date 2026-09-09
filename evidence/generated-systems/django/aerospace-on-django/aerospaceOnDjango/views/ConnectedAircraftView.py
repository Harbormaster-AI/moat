import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.ConnectedAircraftDelegate import ConnectedAircraftDelegate

 #======================================================================
# 
# Encapsulates data for View ConnectedAircraft
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConnectedAircraftView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ConnectedAircraft index.")

def get(request, connectedAircraftId ):
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.get( connectedAircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	connectedAircraft = json.loads(request.body)
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.createFromJson( connectedAircraft )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	connectedAircraft = json.loads(request.body)
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.save( connectedAircraft )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, connectedAircraftId ):
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.delete( connectedAircraftId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAircraft( request, connectedAircraftId, AircraftId ):
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.saveAircraft( connectedAircraftId, AircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAircraft( request, connectedAircraftId ):
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.deleteAircraft( connectedAircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFlightHealthEvents( request, connectedAircraftId, FlightHealthEventsIds ):
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.addFlightHealthEvents( connectedAircraftId, FlightHealthEventsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFlightHealthEvents( request, connectedAircraftId, FlightHealthEventsIds ):
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.removeFlightHealthEvents( connectedAircraftId, FlightHealthEventsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSoftwareLoads( request, connectedAircraftId, SoftwareLoadsIds ):
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.addSoftwareLoads( connectedAircraftId, SoftwareLoadsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSoftwareLoads( request, connectedAircraftId, SoftwareLoadsIds ):
	delegate = ConnectedAircraftDelegate()
	responseData = delegate.removeSoftwareLoads( connectedAircraftId, SoftwareLoadsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

