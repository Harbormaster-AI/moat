import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

 #======================================================================
# 
# Encapsulates data for View Record_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Record_View function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Record_ index.")

def get(request, record_Id ):
	delegate = Record_Delegate()
	responseData = delegate.get( record_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	record_ = json.loads(request.body)
	delegate = Record_Delegate()
	responseData = delegate.createFromJson( record_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	record_ = json.loads(request.body)
	delegate = Record_Delegate()
	responseData = delegate.save( record_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, record_Id ):
	delegate = Record_Delegate()
	responseData = delegate.delete( record_Id )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = Record_Delegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRepository( request, record_Id, RepositoryId ):
	delegate = Record_Delegate()
	responseData = delegate.saveRepository( record_Id, RepositoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRepository( request, record_Id ):
	delegate = Record_Delegate()
	responseData = delegate.deleteRepository( record_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRetentionSchedule( request, record_Id, RetentionScheduleId ):
	delegate = Record_Delegate()
	responseData = delegate.saveRetentionSchedule( record_Id, RetentionScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRetentionSchedule( request, record_Id ):
	delegate = Record_Delegate()
	responseData = delegate.deleteRetentionSchedule( record_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcessingActivities( request, record_Id, ProcessingActivitiesIds ):
	delegate = Record_Delegate()
	responseData = delegate.addProcessingActivities( record_Id, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcessingActivities( request, record_Id, ProcessingActivitiesIds ):
	delegate = Record_Delegate()
	responseData = delegate.removeProcessingActivities( record_Id, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataCategories( request, record_Id, DataCategoriesIds ):
	delegate = Record_Delegate()
	responseData = delegate.addDataCategories( record_Id, DataCategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataCategories( request, record_Id, DataCategoriesIds ):
	delegate = Record_Delegate()
	responseData = delegate.removeDataCategories( record_Id, DataCategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLegalHolds( request, record_Id, LegalHoldsIds ):
	delegate = Record_Delegate()
	responseData = delegate.addLegalHolds( record_Id, LegalHoldsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLegalHolds( request, record_Id, LegalHoldsIds ):
	delegate = Record_Delegate()
	responseData = delegate.removeLegalHolds( record_Id, LegalHoldsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataSubjectRequests( request, record_Id, DataSubjectRequestsIds ):
	delegate = Record_Delegate()
	responseData = delegate.addDataSubjectRequests( record_Id, DataSubjectRequestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataSubjectRequests( request, record_Id, DataSubjectRequestsIds ):
	delegate = Record_Delegate()
	responseData = delegate.removeDataSubjectRequests( record_Id, DataSubjectRequestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

