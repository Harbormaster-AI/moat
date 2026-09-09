import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.MaintenanceAppointmentDelegate import MaintenanceAppointmentDelegate

 #======================================================================
# 
# Encapsulates data for View MaintenanceAppointment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceAppointmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MaintenanceAppointment index.")

def get(request, maintenanceAppointmentId ):
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.get( maintenanceAppointmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	maintenanceAppointment = json.loads(request.body)
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.createFromJson( maintenanceAppointment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	maintenanceAppointment = json.loads(request.body)
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.save( maintenanceAppointment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, maintenanceAppointmentId ):
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.delete( maintenanceAppointmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAircraft( request, maintenanceAppointmentId, AircraftId ):
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.saveAircraft( maintenanceAppointmentId, AircraftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAircraft( request, maintenanceAppointmentId ):
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.deleteAircraft( maintenanceAppointmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMroFacility( request, maintenanceAppointmentId, MroFacilityId ):
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.saveMroFacility( maintenanceAppointmentId, MroFacilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMroFacility( request, maintenanceAppointmentId ):
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.deleteMroFacility( maintenanceAppointmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkOrder( request, maintenanceAppointmentId, WorkOrderId ):
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.saveWorkOrder( maintenanceAppointmentId, WorkOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkOrder( request, maintenanceAppointmentId ):
	delegate = MaintenanceAppointmentDelegate()
	responseData = delegate.deleteWorkOrder( maintenanceAppointmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

