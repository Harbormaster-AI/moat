import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.Exception_Delegate import Exception_Delegate

 #======================================================================
# 
# Encapsulates data for View Exception_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Exception_View function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Exception_ index.")

def get(request, exception_Id ):
	delegate = Exception_Delegate()
	responseData = delegate.get( exception_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	exception_ = json.loads(request.body)
	delegate = Exception_Delegate()
	responseData = delegate.createFromJson( exception_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	exception_ = json.loads(request.body)
	delegate = Exception_Delegate()
	responseData = delegate.save( exception_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, exception_Id ):
	delegate = Exception_Delegate()
	responseData = delegate.delete( exception_Id )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = Exception_Delegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRetentionSchedule( request, exception_Id, RetentionScheduleId ):
	delegate = Exception_Delegate()
	responseData = delegate.saveRetentionSchedule( exception_Id, RetentionScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRetentionSchedule( request, exception_Id ):
	delegate = Exception_Delegate()
	responseData = delegate.deleteRetentionSchedule( exception_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, exception_Id, PolicyId ):
	delegate = Exception_Delegate()
	responseData = delegate.savePolicy( exception_Id, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, exception_Id ):
	delegate = Exception_Delegate()
	responseData = delegate.deletePolicy( exception_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignControl( request, exception_Id, ControlId ):
	delegate = Exception_Delegate()
	responseData = delegate.saveControl( exception_Id, ControlId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignControl( request, exception_Id ):
	delegate = Exception_Delegate()
	responseData = delegate.deleteControl( exception_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRisk( request, exception_Id, RiskId ):
	delegate = Exception_Delegate()
	responseData = delegate.saveRisk( exception_Id, RiskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRisk( request, exception_Id ):
	delegate = Exception_Delegate()
	responseData = delegate.deleteRisk( exception_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

