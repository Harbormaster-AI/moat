import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.ScreeningDelegate import ScreeningDelegate

 #======================================================================
# 
# Encapsulates data for View Screening
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScreeningView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Screening index.")

def get(request, screeningId ):
	delegate = ScreeningDelegate()
	responseData = delegate.get( screeningId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	screening = json.loads(request.body)
	delegate = ScreeningDelegate()
	responseData = delegate.createFromJson( screening )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	screening = json.loads(request.body)
	delegate = ScreeningDelegate()
	responseData = delegate.save( screening )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, screeningId ):
	delegate = ScreeningDelegate()
	responseData = delegate.delete( screeningId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ScreeningDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignApplication( request, screeningId, ApplicationId ):
	delegate = ScreeningDelegate()
	responseData = delegate.saveApplication( screeningId, ApplicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignApplication( request, screeningId ):
	delegate = ScreeningDelegate()
	responseData = delegate.deleteApplication( screeningId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

