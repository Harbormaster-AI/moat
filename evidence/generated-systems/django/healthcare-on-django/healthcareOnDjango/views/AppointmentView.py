import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.AppointmentDelegate import AppointmentDelegate

 #======================================================================
# 
# Encapsulates data for View Appointment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AppointmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Appointment index.")

def get(request, appointmentId ):
	delegate = AppointmentDelegate()
	responseData = delegate.get( appointmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	appointment = json.loads(request.body)
	delegate = AppointmentDelegate()
	responseData = delegate.createFromJson( appointment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	appointment = json.loads(request.body)
	delegate = AppointmentDelegate()
	responseData = delegate.save( appointment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, appointmentId ):
	delegate = AppointmentDelegate()
	responseData = delegate.delete( appointmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AppointmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, appointmentId, PatientId ):
	delegate = AppointmentDelegate()
	responseData = delegate.savePatient( appointmentId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, appointmentId ):
	delegate = AppointmentDelegate()
	responseData = delegate.deletePatient( appointmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignClinician( request, appointmentId, ClinicianId ):
	delegate = AppointmentDelegate()
	responseData = delegate.saveClinician( appointmentId, ClinicianId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignClinician( request, appointmentId ):
	delegate = AppointmentDelegate()
	responseData = delegate.deleteClinician( appointmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFacility( request, appointmentId, FacilityId ):
	delegate = AppointmentDelegate()
	responseData = delegate.saveFacility( appointmentId, FacilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFacility( request, appointmentId ):
	delegate = AppointmentDelegate()
	responseData = delegate.deleteFacility( appointmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEncounter( request, appointmentId, EncounterId ):
	delegate = AppointmentDelegate()
	responseData = delegate.saveEncounter( appointmentId, EncounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEncounter( request, appointmentId ):
	delegate = AppointmentDelegate()
	responseData = delegate.deleteEncounter( appointmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

