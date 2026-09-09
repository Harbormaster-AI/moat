import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.System_Delegate import System_Delegate

 #======================================================================
# 
# Encapsulates data for View System_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class System_View function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the System_ index.")

def get(request, system_Id ):
	delegate = System_Delegate()
	responseData = delegate.get( system_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	system_ = json.loads(request.body)
	delegate = System_Delegate()
	responseData = delegate.createFromJson( system_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	system_ = json.loads(request.body)
	delegate = System_Delegate()
	responseData = delegate.save( system_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, system_Id ):
	delegate = System_Delegate()
	responseData = delegate.delete( system_Id )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = System_Delegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcessingActivities( request, system_Id, ProcessingActivitiesIds ):
	delegate = System_Delegate()
	responseData = delegate.addProcessingActivities( system_Id, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcessingActivities( request, system_Id, ProcessingActivitiesIds ):
	delegate = System_Delegate()
	responseData = delegate.removeProcessingActivities( system_Id, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRecordsRepositories( request, system_Id, RecordsRepositoriesIds ):
	delegate = System_Delegate()
	responseData = delegate.addRecordsRepositories( system_Id, RecordsRepositoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRecordsRepositories( request, system_Id, RecordsRepositoriesIds ):
	delegate = System_Delegate()
	responseData = delegate.removeRecordsRepositories( system_Id, RecordsRepositoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

