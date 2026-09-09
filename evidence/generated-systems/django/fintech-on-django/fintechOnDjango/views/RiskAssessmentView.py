import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.RiskAssessmentDelegate import RiskAssessmentDelegate

 #======================================================================
# 
# Encapsulates data for View RiskAssessment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskAssessmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the RiskAssessment index.")

def get(request, riskAssessmentId ):
	delegate = RiskAssessmentDelegate()
	responseData = delegate.get( riskAssessmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	riskAssessment = json.loads(request.body)
	delegate = RiskAssessmentDelegate()
	responseData = delegate.createFromJson( riskAssessment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	riskAssessment = json.loads(request.body)
	delegate = RiskAssessmentDelegate()
	responseData = delegate.save( riskAssessment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, riskAssessmentId ):
	delegate = RiskAssessmentDelegate()
	responseData = delegate.delete( riskAssessmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RiskAssessmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignApplication( request, riskAssessmentId, ApplicationId ):
	delegate = RiskAssessmentDelegate()
	responseData = delegate.saveApplication( riskAssessmentId, ApplicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignApplication( request, riskAssessmentId ):
	delegate = RiskAssessmentDelegate()
	responseData = delegate.deleteApplication( riskAssessmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

