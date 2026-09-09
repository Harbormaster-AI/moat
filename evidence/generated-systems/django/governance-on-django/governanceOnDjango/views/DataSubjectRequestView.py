import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.DataSubjectRequestDelegate import DataSubjectRequestDelegate

 #======================================================================
# 
# Encapsulates data for View DataSubjectRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSubjectRequestView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DataSubjectRequest index.")

def get(request, dataSubjectRequestId ):
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.get( dataSubjectRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dataSubjectRequest = json.loads(request.body)
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.createFromJson( dataSubjectRequest )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dataSubjectRequest = json.loads(request.body)
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.save( dataSubjectRequest )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dataSubjectRequestId ):
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.delete( dataSubjectRequestId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, dataSubjectRequestId, OrganizationId ):
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.saveOrganization( dataSubjectRequestId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, dataSubjectRequestId ):
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.deleteOrganization( dataSubjectRequestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcessingActivities( request, dataSubjectRequestId, ProcessingActivitiesIds ):
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.addProcessingActivities( dataSubjectRequestId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcessingActivities( request, dataSubjectRequestId, ProcessingActivitiesIds ):
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.removeProcessingActivities( dataSubjectRequestId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRecords( request, dataSubjectRequestId, RecordsIds ):
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.addRecords( dataSubjectRequestId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRecords( request, dataSubjectRequestId, RecordsIds ):
	delegate = DataSubjectRequestDelegate()
	responseData = delegate.removeRecords( dataSubjectRequestId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

