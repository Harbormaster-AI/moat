import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

 #======================================================================
# 
# Encapsulates data for View Aircraft
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Aircraft index.")

def get(request, aircraftId ):
	delegate = AircraftDelegate()
	responseData = delegate.get( aircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aircraft = json.loads(request.body)
	delegate = AircraftDelegate()
	responseData = delegate.createFromJson( aircraft )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aircraft = json.loads(request.body)
	delegate = AircraftDelegate()
	responseData = delegate.save( aircraft )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aircraftId ):
	delegate = AircraftDelegate()
	responseData = delegate.delete( aircraftId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AircraftDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignVariant( request, aircraftId, VariantId ):
	delegate = AircraftDelegate()
	responseData = delegate.saveVariant( aircraftId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, aircraftId ):
	delegate = AircraftDelegate()
	responseData = delegate.deleteVariant( aircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOperator( request, aircraftId, OperatorId ):
	delegate = AircraftDelegate()
	responseData = delegate.saveOperator( aircraftId, OperatorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOperator( request, aircraftId ):
	delegate = AircraftDelegate()
	responseData = delegate.deleteOperator( aircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRegistration( request, aircraftId, RegistrationId ):
	delegate = AircraftDelegate()
	responseData = delegate.saveRegistration( aircraftId, RegistrationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRegistration( request, aircraftId ):
	delegate = AircraftDelegate()
	responseData = delegate.deleteRegistration( aircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarranty( request, aircraftId, WarrantyId ):
	delegate = AircraftDelegate()
	responseData = delegate.saveWarranty( aircraftId, WarrantyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarranty( request, aircraftId ):
	delegate = AircraftDelegate()
	responseData = delegate.deleteWarranty( aircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignConnectedAircraft( request, aircraftId, ConnectedAircraftId ):
	delegate = AircraftDelegate()
	responseData = delegate.saveConnectedAircraft( aircraftId, ConnectedAircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignConnectedAircraft( request, aircraftId ):
	delegate = AircraftDelegate()
	responseData = delegate.deleteConnectedAircraft( aircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCabinLayout( request, aircraftId, CabinLayoutId ):
	delegate = AircraftDelegate()
	responseData = delegate.saveCabinLayout( aircraftId, CabinLayoutId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCabinLayout( request, aircraftId ):
	delegate = AircraftDelegate()
	responseData = delegate.deleteCabinLayout( aircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMaintenanceRecords( request, aircraftId, MaintenanceRecordsIds ):
	delegate = AircraftDelegate()
	responseData = delegate.addMaintenanceRecords( aircraftId, MaintenanceRecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMaintenanceRecords( request, aircraftId, MaintenanceRecordsIds ):
	delegate = AircraftDelegate()
	responseData = delegate.removeMaintenanceRecords( aircraftId, MaintenanceRecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

