import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

 #======================================================================
# 
# Encapsulates data for View DataProcessingActivity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataProcessingActivityView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DataProcessingActivity index.")

def get(request, dataProcessingActivityId ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.get( dataProcessingActivityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dataProcessingActivity = json.loads(request.body)
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.createFromJson( dataProcessingActivity )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dataProcessingActivity = json.loads(request.body)
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.save( dataProcessingActivity )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dataProcessingActivityId ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.delete( dataProcessingActivityId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, dataProcessingActivityId, OrganizationId ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.saveOrganization( dataProcessingActivityId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, dataProcessingActivityId ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.deleteOrganization( dataProcessingActivityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataCategories( request, dataProcessingActivityId, DataCategoriesIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.addDataCategories( dataProcessingActivityId, DataCategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataCategories( request, dataProcessingActivityId, DataCategoriesIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.removeDataCategories( dataProcessingActivityId, DataCategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSystems( request, dataProcessingActivityId, SystemsIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.addSystems( dataProcessingActivityId, SystemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSystems( request, dataProcessingActivityId, SystemsIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.removeSystems( dataProcessingActivityId, SystemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRecords( request, dataProcessingActivityId, RecordsIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.addRecords( dataProcessingActivityId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRecords( request, dataProcessingActivityId, RecordsIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.removeRecords( dataProcessingActivityId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPrivacyNotices( request, dataProcessingActivityId, PrivacyNoticesIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.addPrivacyNotices( dataProcessingActivityId, PrivacyNoticesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePrivacyNotices( request, dataProcessingActivityId, PrivacyNoticesIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.removePrivacyNotices( dataProcessingActivityId, PrivacyNoticesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addThirdParties( request, dataProcessingActivityId, ThirdPartiesIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.addThirdParties( dataProcessingActivityId, ThirdPartiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeThirdParties( request, dataProcessingActivityId, ThirdPartiesIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.removeThirdParties( dataProcessingActivityId, ThirdPartiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addConsents( request, dataProcessingActivityId, ConsentsIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.addConsents( dataProcessingActivityId, ConsentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeConsents( request, dataProcessingActivityId, ConsentsIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.removeConsents( dataProcessingActivityId, ConsentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataBreaches( request, dataProcessingActivityId, DataBreachesIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.addDataBreaches( dataProcessingActivityId, DataBreachesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataBreaches( request, dataProcessingActivityId, DataBreachesIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.removeDataBreaches( dataProcessingActivityId, DataBreachesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataSubjectRequests( request, dataProcessingActivityId, DataSubjectRequestsIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.addDataSubjectRequests( dataProcessingActivityId, DataSubjectRequestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataSubjectRequests( request, dataProcessingActivityId, DataSubjectRequestsIds ):
	delegate = DataProcessingActivityDelegate()
	responseData = delegate.removeDataSubjectRequests( dataProcessingActivityId, DataSubjectRequestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

