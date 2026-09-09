import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.OnboardingTaskDelegate import OnboardingTaskDelegate

 #======================================================================
# 
# Encapsulates data for View OnboardingTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OnboardingTaskView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the OnboardingTask index.")

def get(request, onboardingTaskId ):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.get( onboardingTaskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	onboardingTask = json.loads(request.body)
	delegate = OnboardingTaskDelegate()
	responseData = delegate.createFromJson( onboardingTask )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	onboardingTask = json.loads(request.body)
	delegate = OnboardingTaskDelegate()
	responseData = delegate.save( onboardingTask )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, onboardingTaskId ):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.delete( onboardingTaskId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, onboardingTaskId, EmployeeId ):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.saveEmployee( onboardingTaskId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, onboardingTaskId ):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.deleteEmployee( onboardingTaskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAssignedTo( request, onboardingTaskId, AssignedToId ):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.saveAssignedTo( onboardingTaskId, AssignedToId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAssignedTo( request, onboardingTaskId ):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.deleteAssignedTo( onboardingTaskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRelatedOffer( request, onboardingTaskId, RelatedOfferId ):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.saveRelatedOffer( onboardingTaskId, RelatedOfferId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRelatedOffer( request, onboardingTaskId ):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.deleteRelatedOffer( onboardingTaskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDependencies( request, onboardingTaskId, DependenciesIds ):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.addDependencies( onboardingTaskId, DependenciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDependencies( request, onboardingTaskId, DependenciesIds ):
	delegate = OnboardingTaskDelegate()
	responseData = delegate.removeDependencies( onboardingTaskId, DependenciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

