import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.BackgroundCheckDelegate import BackgroundCheckDelegate

 #======================================================================
# 
# Encapsulates data for View BackgroundCheck
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BackgroundCheckView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BackgroundCheck index.")

def get(request, backgroundCheckId ):
	delegate = BackgroundCheckDelegate()
	responseData = delegate.get( backgroundCheckId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	backgroundCheck = json.loads(request.body)
	delegate = BackgroundCheckDelegate()
	responseData = delegate.createFromJson( backgroundCheck )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	backgroundCheck = json.loads(request.body)
	delegate = BackgroundCheckDelegate()
	responseData = delegate.save( backgroundCheck )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, backgroundCheckId ):
	delegate = BackgroundCheckDelegate()
	responseData = delegate.delete( backgroundCheckId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BackgroundCheckDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCandidate( request, backgroundCheckId, CandidateId ):
	delegate = BackgroundCheckDelegate()
	responseData = delegate.saveCandidate( backgroundCheckId, CandidateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCandidate( request, backgroundCheckId ):
	delegate = BackgroundCheckDelegate()
	responseData = delegate.deleteCandidate( backgroundCheckId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRequisition( request, backgroundCheckId, RequisitionId ):
	delegate = BackgroundCheckDelegate()
	responseData = delegate.saveRequisition( backgroundCheckId, RequisitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRequisition( request, backgroundCheckId ):
	delegate = BackgroundCheckDelegate()
	responseData = delegate.deleteRequisition( backgroundCheckId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignReport( request, backgroundCheckId, ReportId ):
	delegate = BackgroundCheckDelegate()
	responseData = delegate.saveReport( backgroundCheckId, ReportId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignReport( request, backgroundCheckId ):
	delegate = BackgroundCheckDelegate()
	responseData = delegate.deleteReport( backgroundCheckId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

