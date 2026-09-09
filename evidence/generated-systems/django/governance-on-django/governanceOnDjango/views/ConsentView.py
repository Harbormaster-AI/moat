import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.ConsentDelegate import ConsentDelegate

 #======================================================================
# 
# Encapsulates data for View Consent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConsentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Consent index.")

def get(request, consentId ):
	delegate = ConsentDelegate()
	responseData = delegate.get( consentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	consent = json.loads(request.body)
	delegate = ConsentDelegate()
	responseData = delegate.createFromJson( consent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	consent = json.loads(request.body)
	delegate = ConsentDelegate()
	responseData = delegate.save( consent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, consentId ):
	delegate = ConsentDelegate()
	responseData = delegate.delete( consentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ConsentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPrivacyNotice( request, consentId, PrivacyNoticeId ):
	delegate = ConsentDelegate()
	responseData = delegate.savePrivacyNotice( consentId, PrivacyNoticeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPrivacyNotice( request, consentId ):
	delegate = ConsentDelegate()
	responseData = delegate.deletePrivacyNotice( consentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcessingActivities( request, consentId, ProcessingActivitiesIds ):
	delegate = ConsentDelegate()
	responseData = delegate.addProcessingActivities( consentId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcessingActivities( request, consentId, ProcessingActivitiesIds ):
	delegate = ConsentDelegate()
	responseData = delegate.removeProcessingActivities( consentId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

