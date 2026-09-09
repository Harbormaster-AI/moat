import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.ObservationDelegate import ObservationDelegate

 #======================================================================
# 
# Encapsulates data for View Observation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ObservationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Observation index.")

def get(request, observationId ):
	delegate = ObservationDelegate()
	responseData = delegate.get( observationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	observation = json.loads(request.body)
	delegate = ObservationDelegate()
	responseData = delegate.createFromJson( observation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	observation = json.loads(request.body)
	delegate = ObservationDelegate()
	responseData = delegate.save( observation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, observationId ):
	delegate = ObservationDelegate()
	responseData = delegate.delete( observationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ObservationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEncounter( request, observationId, EncounterId ):
	delegate = ObservationDelegate()
	responseData = delegate.saveEncounter( observationId, EncounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEncounter( request, observationId ):
	delegate = ObservationDelegate()
	responseData = delegate.deleteEncounter( observationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, observationId, PatientId ):
	delegate = ObservationDelegate()
	responseData = delegate.savePatient( observationId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, observationId ):
	delegate = ObservationDelegate()
	responseData = delegate.deletePatient( observationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDevice( request, observationId, DeviceId ):
	delegate = ObservationDelegate()
	responseData = delegate.saveDevice( observationId, DeviceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDevice( request, observationId ):
	delegate = ObservationDelegate()
	responseData = delegate.deleteDevice( observationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLabResult( request, observationId, LabResultId ):
	delegate = ObservationDelegate()
	responseData = delegate.saveLabResult( observationId, LabResultId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLabResult( request, observationId ):
	delegate = ObservationDelegate()
	responseData = delegate.deleteLabResult( observationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

