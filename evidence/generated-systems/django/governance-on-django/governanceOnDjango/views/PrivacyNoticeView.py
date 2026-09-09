import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.PrivacyNoticeDelegate import PrivacyNoticeDelegate

 #======================================================================
# 
# Encapsulates data for View PrivacyNotice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PrivacyNoticeView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PrivacyNotice index.")

def get(request, privacyNoticeId ):
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.get( privacyNoticeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	privacyNotice = json.loads(request.body)
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.createFromJson( privacyNotice )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	privacyNotice = json.loads(request.body)
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.save( privacyNotice )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, privacyNoticeId ):
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.delete( privacyNoticeId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, privacyNoticeId, OrganizationId ):
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.saveOrganization( privacyNoticeId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, privacyNoticeId ):
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.deleteOrganization( privacyNoticeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcessingActivities( request, privacyNoticeId, ProcessingActivitiesIds ):
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.addProcessingActivities( privacyNoticeId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcessingActivities( request, privacyNoticeId, ProcessingActivitiesIds ):
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.removeProcessingActivities( privacyNoticeId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addConsents( request, privacyNoticeId, ConsentsIds ):
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.addConsents( privacyNoticeId, ConsentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeConsents( request, privacyNoticeId, ConsentsIds ):
	delegate = PrivacyNoticeDelegate()
	responseData = delegate.removeConsents( privacyNoticeId, ConsentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

