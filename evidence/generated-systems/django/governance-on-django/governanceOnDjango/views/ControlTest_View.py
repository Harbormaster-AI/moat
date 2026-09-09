import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.ControlTest_Delegate import ControlTest_Delegate

 #======================================================================
# 
# Encapsulates data for View ControlTest_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlTest_View function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ControlTest_ index.")

def get(request, controlTest_Id ):
	delegate = ControlTest_Delegate()
	responseData = delegate.get( controlTest_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	controlTest_ = json.loads(request.body)
	delegate = ControlTest_Delegate()
	responseData = delegate.createFromJson( controlTest_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	controlTest_ = json.loads(request.body)
	delegate = ControlTest_Delegate()
	responseData = delegate.save( controlTest_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, controlTest_Id ):
	delegate = ControlTest_Delegate()
	responseData = delegate.delete( controlTest_Id )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ControlTest_Delegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignControl( request, controlTest_Id, ControlId ):
	delegate = ControlTest_Delegate()
	responseData = delegate.saveControl( controlTest_Id, ControlId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignControl( request, controlTest_Id ):
	delegate = ControlTest_Delegate()
	responseData = delegate.deleteControl( controlTest_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEngagement( request, controlTest_Id, EngagementId ):
	delegate = ControlTest_Delegate()
	responseData = delegate.saveEngagement( controlTest_Id, EngagementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEngagement( request, controlTest_Id ):
	delegate = ControlTest_Delegate()
	responseData = delegate.deleteEngagement( controlTest_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEvidence( request, controlTest_Id, EvidenceIds ):
	delegate = ControlTest_Delegate()
	responseData = delegate.addEvidence( controlTest_Id, EvidenceIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEvidence( request, controlTest_Id, EvidenceIds ):
	delegate = ControlTest_Delegate()
	responseData = delegate.removeEvidence( controlTest_Id, EvidenceIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

