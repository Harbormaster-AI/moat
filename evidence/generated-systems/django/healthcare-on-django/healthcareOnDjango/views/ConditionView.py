import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.ConditionDelegate import ConditionDelegate

 #======================================================================
# 
# Encapsulates data for View Condition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConditionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Condition index.")

def get(request, conditionId ):
	delegate = ConditionDelegate()
	responseData = delegate.get( conditionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	condition = json.loads(request.body)
	delegate = ConditionDelegate()
	responseData = delegate.createFromJson( condition )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	condition = json.loads(request.body)
	delegate = ConditionDelegate()
	responseData = delegate.save( condition )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, conditionId ):
	delegate = ConditionDelegate()
	responseData = delegate.delete( conditionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ConditionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, conditionId, PatientId ):
	delegate = ConditionDelegate()
	responseData = delegate.savePatient( conditionId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, conditionId ):
	delegate = ConditionDelegate()
	responseData = delegate.deletePatient( conditionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

