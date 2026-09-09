import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.CareTaskDelegate import CareTaskDelegate

 #======================================================================
# 
# Encapsulates data for View CareTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CareTaskView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CareTask index.")

def get(request, careTaskId ):
	delegate = CareTaskDelegate()
	responseData = delegate.get( careTaskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	careTask = json.loads(request.body)
	delegate = CareTaskDelegate()
	responseData = delegate.createFromJson( careTask )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	careTask = json.loads(request.body)
	delegate = CareTaskDelegate()
	responseData = delegate.save( careTask )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, careTaskId ):
	delegate = CareTaskDelegate()
	responseData = delegate.delete( careTaskId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CareTaskDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCarePlan( request, careTaskId, CarePlanId ):
	delegate = CareTaskDelegate()
	responseData = delegate.saveCarePlan( careTaskId, CarePlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCarePlan( request, careTaskId ):
	delegate = CareTaskDelegate()
	responseData = delegate.deleteCarePlan( careTaskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAssignedTo( request, careTaskId, AssignedToId ):
	delegate = CareTaskDelegate()
	responseData = delegate.saveAssignedTo( careTaskId, AssignedToId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAssignedTo( request, careTaskId ):
	delegate = CareTaskDelegate()
	responseData = delegate.deleteAssignedTo( careTaskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEncounter( request, careTaskId, EncounterId ):
	delegate = CareTaskDelegate()
	responseData = delegate.saveEncounter( careTaskId, EncounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEncounter( request, careTaskId ):
	delegate = CareTaskDelegate()
	responseData = delegate.deleteEncounter( careTaskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

