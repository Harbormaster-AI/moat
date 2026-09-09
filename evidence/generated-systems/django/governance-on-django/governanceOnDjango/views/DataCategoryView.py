import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.DataCategoryDelegate import DataCategoryDelegate

 #======================================================================
# 
# Encapsulates data for View DataCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataCategoryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DataCategory index.")

def get(request, dataCategoryId ):
	delegate = DataCategoryDelegate()
	responseData = delegate.get( dataCategoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dataCategory = json.loads(request.body)
	delegate = DataCategoryDelegate()
	responseData = delegate.createFromJson( dataCategory )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dataCategory = json.loads(request.body)
	delegate = DataCategoryDelegate()
	responseData = delegate.save( dataCategory )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dataCategoryId ):
	delegate = DataCategoryDelegate()
	responseData = delegate.delete( dataCategoryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DataCategoryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcessingActivities( request, dataCategoryId, ProcessingActivitiesIds ):
	delegate = DataCategoryDelegate()
	responseData = delegate.addProcessingActivities( dataCategoryId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcessingActivities( request, dataCategoryId, ProcessingActivitiesIds ):
	delegate = DataCategoryDelegate()
	responseData = delegate.removeProcessingActivities( dataCategoryId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRecords( request, dataCategoryId, RecordsIds ):
	delegate = DataCategoryDelegate()
	responseData = delegate.addRecords( dataCategoryId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRecords( request, dataCategoryId, RecordsIds ):
	delegate = DataCategoryDelegate()
	responseData = delegate.removeRecords( dataCategoryId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataBreaches( request, dataCategoryId, DataBreachesIds ):
	delegate = DataCategoryDelegate()
	responseData = delegate.addDataBreaches( dataCategoryId, DataBreachesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataBreaches( request, dataCategoryId, DataBreachesIds ):
	delegate = DataCategoryDelegate()
	responseData = delegate.removeDataBreaches( dataCategoryId, DataBreachesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

