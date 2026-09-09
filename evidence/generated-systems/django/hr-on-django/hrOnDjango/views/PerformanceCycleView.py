import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.PerformanceCycleDelegate import PerformanceCycleDelegate

 #======================================================================
# 
# Encapsulates data for View PerformanceCycle
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceCycleView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PerformanceCycle index.")

def get(request, performanceCycleId ):
	delegate = PerformanceCycleDelegate()
	responseData = delegate.get( performanceCycleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	performanceCycle = json.loads(request.body)
	delegate = PerformanceCycleDelegate()
	responseData = delegate.createFromJson( performanceCycle )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	performanceCycle = json.loads(request.body)
	delegate = PerformanceCycleDelegate()
	responseData = delegate.save( performanceCycle )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, performanceCycleId ):
	delegate = PerformanceCycleDelegate()
	responseData = delegate.delete( performanceCycleId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PerformanceCycleDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, performanceCycleId, OrganizationId ):
	delegate = PerformanceCycleDelegate()
	responseData = delegate.saveOrganization( performanceCycleId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, performanceCycleId ):
	delegate = PerformanceCycleDelegate()
	responseData = delegate.deleteOrganization( performanceCycleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReviews( request, performanceCycleId, ReviewsIds ):
	delegate = PerformanceCycleDelegate()
	responseData = delegate.addReviews( performanceCycleId, ReviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReviews( request, performanceCycleId, ReviewsIds ):
	delegate = PerformanceCycleDelegate()
	responseData = delegate.removeReviews( performanceCycleId, ReviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGoals( request, performanceCycleId, GoalsIds ):
	delegate = PerformanceCycleDelegate()
	responseData = delegate.addGoals( performanceCycleId, GoalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGoals( request, performanceCycleId, GoalsIds ):
	delegate = PerformanceCycleDelegate()
	responseData = delegate.removeGoals( performanceCycleId, GoalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

