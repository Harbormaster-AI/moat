import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

 #======================================================================
# 
# Encapsulates data for View DataBreach
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataBreachView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DataBreach index.")

def get(request, dataBreachId ):
	delegate = DataBreachDelegate()
	responseData = delegate.get( dataBreachId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dataBreach = json.loads(request.body)
	delegate = DataBreachDelegate()
	responseData = delegate.createFromJson( dataBreach )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dataBreach = json.loads(request.body)
	delegate = DataBreachDelegate()
	responseData = delegate.save( dataBreach )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dataBreachId ):
	delegate = DataBreachDelegate()
	responseData = delegate.delete( dataBreachId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DataBreachDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, dataBreachId, OrganizationId ):
	delegate = DataBreachDelegate()
	responseData = delegate.saveOrganization( dataBreachId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, dataBreachId ):
	delegate = DataBreachDelegate()
	responseData = delegate.deleteOrganization( dataBreachId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMatter( request, dataBreachId, MatterId ):
	delegate = DataBreachDelegate()
	responseData = delegate.saveMatter( dataBreachId, MatterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMatter( request, dataBreachId ):
	delegate = DataBreachDelegate()
	responseData = delegate.deleteMatter( dataBreachId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcessingActivities( request, dataBreachId, ProcessingActivitiesIds ):
	delegate = DataBreachDelegate()
	responseData = delegate.addProcessingActivities( dataBreachId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcessingActivities( request, dataBreachId, ProcessingActivitiesIds ):
	delegate = DataBreachDelegate()
	responseData = delegate.removeProcessingActivities( dataBreachId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataCategories( request, dataBreachId, DataCategoriesIds ):
	delegate = DataBreachDelegate()
	responseData = delegate.addDataCategories( dataBreachId, DataCategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataCategories( request, dataBreachId, DataCategoriesIds ):
	delegate = DataBreachDelegate()
	responseData = delegate.removeDataCategories( dataBreachId, DataCategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addThirdParties( request, dataBreachId, ThirdPartiesIds ):
	delegate = DataBreachDelegate()
	responseData = delegate.addThirdParties( dataBreachId, ThirdPartiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeThirdParties( request, dataBreachId, ThirdPartiesIds ):
	delegate = DataBreachDelegate()
	responseData = delegate.removeThirdParties( dataBreachId, ThirdPartiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

