import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

 #======================================================================
# 
# Encapsulates data for View MaintenanceWorkOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceWorkOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MaintenanceWorkOrder index.")

def get(request, maintenanceWorkOrderId ):
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.get( maintenanceWorkOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	maintenanceWorkOrder = json.loads(request.body)
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.createFromJson( maintenanceWorkOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	maintenanceWorkOrder = json.loads(request.body)
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.save( maintenanceWorkOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, maintenanceWorkOrderId ):
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.delete( maintenanceWorkOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAircraft( request, maintenanceWorkOrderId, AircraftId ):
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.saveAircraft( maintenanceWorkOrderId, AircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAircraft( request, maintenanceWorkOrderId ):
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.deleteAircraft( maintenanceWorkOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAirworthinessDirective( request, maintenanceWorkOrderId, AirworthinessDirectiveId ):
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.saveAirworthinessDirective( maintenanceWorkOrderId, AirworthinessDirectiveId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAirworthinessDirective( request, maintenanceWorkOrderId ):
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.deleteAirworthinessDirective( maintenanceWorkOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignServiceBulletin( request, maintenanceWorkOrderId, ServiceBulletinId ):
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.saveServiceBulletin( maintenanceWorkOrderId, ServiceBulletinId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignServiceBulletin( request, maintenanceWorkOrderId ):
	delegate = MaintenanceWorkOrderDelegate()
	responseData = delegate.deleteServiceBulletin( maintenanceWorkOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

