import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.AdmissionDelegate import AdmissionDelegate

 #======================================================================
# 
# Encapsulates data for View Admission
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdmissionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Admission index.")

def get(request, admissionId ):
	delegate = AdmissionDelegate()
	responseData = delegate.get( admissionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	admission = json.loads(request.body)
	delegate = AdmissionDelegate()
	responseData = delegate.createFromJson( admission )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	admission = json.loads(request.body)
	delegate = AdmissionDelegate()
	responseData = delegate.save( admission )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, admissionId ):
	delegate = AdmissionDelegate()
	responseData = delegate.delete( admissionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AdmissionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEncounter( request, admissionId, EncounterId ):
	delegate = AdmissionDelegate()
	responseData = delegate.saveEncounter( admissionId, EncounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEncounter( request, admissionId ):
	delegate = AdmissionDelegate()
	responseData = delegate.deleteEncounter( admissionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFacility( request, admissionId, FacilityId ):
	delegate = AdmissionDelegate()
	responseData = delegate.saveFacility( admissionId, FacilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFacility( request, admissionId ):
	delegate = AdmissionDelegate()
	responseData = delegate.deleteFacility( admissionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

