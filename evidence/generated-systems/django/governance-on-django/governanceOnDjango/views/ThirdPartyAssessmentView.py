import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.ThirdPartyAssessmentDelegate import ThirdPartyAssessmentDelegate

 #======================================================================
# 
# Encapsulates data for View ThirdPartyAssessment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ThirdPartyAssessmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ThirdPartyAssessment index.")

def get(request, thirdPartyAssessmentId ):
	delegate = ThirdPartyAssessmentDelegate()
	responseData = delegate.get( thirdPartyAssessmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	thirdPartyAssessment = json.loads(request.body)
	delegate = ThirdPartyAssessmentDelegate()
	responseData = delegate.createFromJson( thirdPartyAssessment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	thirdPartyAssessment = json.loads(request.body)
	delegate = ThirdPartyAssessmentDelegate()
	responseData = delegate.save( thirdPartyAssessment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, thirdPartyAssessmentId ):
	delegate = ThirdPartyAssessmentDelegate()
	responseData = delegate.delete( thirdPartyAssessmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ThirdPartyAssessmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignThirdParty( request, thirdPartyAssessmentId, ThirdPartyId ):
	delegate = ThirdPartyAssessmentDelegate()
	responseData = delegate.saveThirdParty( thirdPartyAssessmentId, ThirdPartyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignThirdParty( request, thirdPartyAssessmentId ):
	delegate = ThirdPartyAssessmentDelegate()
	responseData = delegate.deleteThirdParty( thirdPartyAssessmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addIssues( request, thirdPartyAssessmentId, IssuesIds ):
	delegate = ThirdPartyAssessmentDelegate()
	responseData = delegate.addIssues( thirdPartyAssessmentId, IssuesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeIssues( request, thirdPartyAssessmentId, IssuesIds ):
	delegate = ThirdPartyAssessmentDelegate()
	responseData = delegate.removeIssues( thirdPartyAssessmentId, IssuesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

