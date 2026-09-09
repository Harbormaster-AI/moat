import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.RunParameterDelegate import RunParameterDelegate

 #======================================================================
# 
# Encapsulates data for View RunParameter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RunParameterView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the RunParameter index.")

def get(request, runParameterId ):
	delegate = RunParameterDelegate()
	responseData = delegate.get( runParameterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	runParameter = json.loads(request.body)
	delegate = RunParameterDelegate()
	responseData = delegate.createFromJson( runParameter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	runParameter = json.loads(request.body)
	delegate = RunParameterDelegate()
	responseData = delegate.save( runParameter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, runParameterId ):
	delegate = RunParameterDelegate()
	responseData = delegate.delete( runParameterId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RunParameterDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTrainingRun( request, runParameterId, TrainingRunId ):
	delegate = RunParameterDelegate()
	responseData = delegate.saveTrainingRun( runParameterId, TrainingRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTrainingRun( request, runParameterId ):
	delegate = RunParameterDelegate()
	responseData = delegate.deleteTrainingRun( runParameterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

