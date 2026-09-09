import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.AllergyDelegate import AllergyDelegate

 #======================================================================
# 
# Encapsulates data for View Allergy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AllergyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Allergy index.")

def get(request, allergyId ):
	delegate = AllergyDelegate()
	responseData = delegate.get( allergyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	allergy = json.loads(request.body)
	delegate = AllergyDelegate()
	responseData = delegate.createFromJson( allergy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	allergy = json.loads(request.body)
	delegate = AllergyDelegate()
	responseData = delegate.save( allergy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, allergyId ):
	delegate = AllergyDelegate()
	responseData = delegate.delete( allergyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AllergyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, allergyId, PatientId ):
	delegate = AllergyDelegate()
	responseData = delegate.savePatient( allergyId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, allergyId ):
	delegate = AllergyDelegate()
	responseData = delegate.deletePatient( allergyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

