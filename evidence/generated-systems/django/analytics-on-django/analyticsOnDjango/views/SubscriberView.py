import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.SubscriberDelegate import SubscriberDelegate

 #======================================================================
# 
# Encapsulates data for View Subscriber
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubscriberView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Subscriber index.")

def get(request, subscriberId ):
	delegate = SubscriberDelegate()
	responseData = delegate.get( subscriberId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	subscriber = json.loads(request.body)
	delegate = SubscriberDelegate()
	responseData = delegate.createFromJson( subscriber )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	subscriber = json.loads(request.body)
	delegate = SubscriberDelegate()
	responseData = delegate.save( subscriber )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, subscriberId ):
	delegate = SubscriberDelegate()
	responseData = delegate.delete( subscriberId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SubscriberDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAlerts( request, subscriberId, AlertsIds ):
	delegate = SubscriberDelegate()
	responseData = delegate.addAlerts( subscriberId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAlerts( request, subscriberId, AlertsIds ):
	delegate = SubscriberDelegate()
	responseData = delegate.removeAlerts( subscriberId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

