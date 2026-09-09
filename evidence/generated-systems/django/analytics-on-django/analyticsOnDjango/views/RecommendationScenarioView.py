import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.RecommendationScenarioDelegate import RecommendationScenarioDelegate

 #======================================================================
# 
# Encapsulates data for View RecommendationScenario
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecommendationScenarioView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the RecommendationScenario index.")

def get(request, recommendationScenarioId ):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.get( recommendationScenarioId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	recommendationScenario = json.loads(request.body)
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.createFromJson( recommendationScenario )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	recommendationScenario = json.loads(request.body)
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.save( recommendationScenario )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, recommendationScenarioId ):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.delete( recommendationScenarioId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModels( request, recommendationScenarioId, ModelsIds ):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.addModels( recommendationScenarioId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModels( request, recommendationScenarioId, ModelsIds ):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.removeModels( recommendationScenarioId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, recommendationScenarioId, DatasetsIds ):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.addDatasets( recommendationScenarioId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, recommendationScenarioId, DatasetsIds ):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.removeDatasets( recommendationScenarioId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addExperiments( request, recommendationScenarioId, ExperimentsIds ):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.addExperiments( recommendationScenarioId, ExperimentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeExperiments( request, recommendationScenarioId, ExperimentsIds ):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.removeExperiments( recommendationScenarioId, ExperimentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAlerts( request, recommendationScenarioId, AlertsIds ):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.addAlerts( recommendationScenarioId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAlerts( request, recommendationScenarioId, AlertsIds ):
	delegate = RecommendationScenarioDelegate()
	responseData = delegate.removeAlerts( recommendationScenarioId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

