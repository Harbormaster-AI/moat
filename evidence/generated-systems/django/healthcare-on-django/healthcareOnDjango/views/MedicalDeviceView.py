import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.MedicalDeviceDelegate import MedicalDeviceDelegate

 #======================================================================
# 
# Encapsulates data for View MedicalDevice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicalDeviceView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MedicalDevice index.")

def get(request, medicalDeviceId ):
	delegate = MedicalDeviceDelegate()
	responseData = delegate.get( medicalDeviceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	medicalDevice = json.loads(request.body)
	delegate = MedicalDeviceDelegate()
	responseData = delegate.createFromJson( medicalDevice )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	medicalDevice = json.loads(request.body)
	delegate = MedicalDeviceDelegate()
	responseData = delegate.save( medicalDevice )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, medicalDeviceId ):
	delegate = MedicalDeviceDelegate()
	responseData = delegate.delete( medicalDeviceId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MedicalDeviceDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, medicalDeviceId, PatientId ):
	delegate = MedicalDeviceDelegate()
	responseData = delegate.savePatient( medicalDeviceId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, medicalDeviceId ):
	delegate = MedicalDeviceDelegate()
	responseData = delegate.deletePatient( medicalDeviceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addObservations( request, medicalDeviceId, ObservationsIds ):
	delegate = MedicalDeviceDelegate()
	responseData = delegate.addObservations( medicalDeviceId, ObservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeObservations( request, medicalDeviceId, ObservationsIds ):
	delegate = MedicalDeviceDelegate()
	responseData = delegate.removeObservations( medicalDeviceId, ObservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSoftwareUpdates( request, medicalDeviceId, SoftwareUpdatesIds ):
	delegate = MedicalDeviceDelegate()
	responseData = delegate.addSoftwareUpdates( medicalDeviceId, SoftwareUpdatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSoftwareUpdates( request, medicalDeviceId, SoftwareUpdatesIds ):
	delegate = MedicalDeviceDelegate()
	responseData = delegate.removeSoftwareUpdates( medicalDeviceId, SoftwareUpdatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

