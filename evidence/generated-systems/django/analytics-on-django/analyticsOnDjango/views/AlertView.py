import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

 #======================================================================
# 
# Encapsulates data for View Alert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AlertView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Alert index.")

def get(request, alertId ):
	delegate = AlertDelegate()
	responseData = delegate.get( alertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	alert = json.loads(request.body)
	delegate = AlertDelegate()
	responseData = delegate.createFromJson( alert )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	alert = json.loads(request.body)
	delegate = AlertDelegate()
	responseData = delegate.save( alert )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, alertId ):
	delegate = AlertDelegate()
	responseData = delegate.delete( alertId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AlertDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMetric( request, alertId, MetricId ):
	delegate = AlertDelegate()
	responseData = delegate.saveMetric( alertId, MetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMetric( request, alertId ):
	delegate = AlertDelegate()
	responseData = delegate.deleteMetric( alertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDashboard( request, alertId, DashboardId ):
	delegate = AlertDelegate()
	responseData = delegate.saveDashboard( alertId, DashboardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDashboard( request, alertId ):
	delegate = AlertDelegate()
	responseData = delegate.deleteDashboard( alertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDataset( request, alertId, DatasetId ):
	delegate = AlertDelegate()
	responseData = delegate.saveDataset( alertId, DatasetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDataset( request, alertId ):
	delegate = AlertDelegate()
	responseData = delegate.deleteDataset( alertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRule( request, alertId, RuleId ):
	delegate = AlertDelegate()
	responseData = delegate.saveRule( alertId, RuleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRule( request, alertId ):
	delegate = AlertDelegate()
	responseData = delegate.deleteRule( alertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAnomalies( request, alertId, AnomaliesIds ):
	delegate = AlertDelegate()
	responseData = delegate.addAnomalies( alertId, AnomaliesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAnomalies( request, alertId, AnomaliesIds ):
	delegate = AlertDelegate()
	responseData = delegate.removeAnomalies( alertId, AnomaliesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSubscribers( request, alertId, SubscribersIds ):
	delegate = AlertDelegate()
	responseData = delegate.addSubscribers( alertId, SubscribersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSubscribers( request, alertId, SubscribersIds ):
	delegate = AlertDelegate()
	responseData = delegate.removeSubscribers( alertId, SubscribersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

