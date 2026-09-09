import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.JobProfileDelegate import JobProfileDelegate

 #======================================================================
# 
# Encapsulates data for View JobProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobProfileView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the JobProfile index.")

def get(request, jobProfileId ):
	delegate = JobProfileDelegate()
	responseData = delegate.get( jobProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	jobProfile = json.loads(request.body)
	delegate = JobProfileDelegate()
	responseData = delegate.createFromJson( jobProfile )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	jobProfile = json.loads(request.body)
	delegate = JobProfileDelegate()
	responseData = delegate.save( jobProfile )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, jobProfileId ):
	delegate = JobProfileDelegate()
	responseData = delegate.delete( jobProfileId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = JobProfileDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignJobFamily( request, jobProfileId, JobFamilyId ):
	delegate = JobProfileDelegate()
	responseData = delegate.saveJobFamily( jobProfileId, JobFamilyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignJobFamily( request, jobProfileId ):
	delegate = JobProfileDelegate()
	responseData = delegate.deleteJobFamily( jobProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCompetencies( request, jobProfileId, CompetenciesIds ):
	delegate = JobProfileDelegate()
	responseData = delegate.addCompetencies( jobProfileId, CompetenciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCompetencies( request, jobProfileId, CompetenciesIds ):
	delegate = JobProfileDelegate()
	responseData = delegate.removeCompetencies( jobProfileId, CompetenciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTrainingRecommendations( request, jobProfileId, TrainingRecommendationsIds ):
	delegate = JobProfileDelegate()
	responseData = delegate.addTrainingRecommendations( jobProfileId, TrainingRecommendationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTrainingRecommendations( request, jobProfileId, TrainingRecommendationsIds ):
	delegate = JobProfileDelegate()
	responseData = delegate.removeTrainingRecommendations( jobProfileId, TrainingRecommendationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPositions( request, jobProfileId, PositionsIds ):
	delegate = JobProfileDelegate()
	responseData = delegate.addPositions( jobProfileId, PositionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePositions( request, jobProfileId, PositionsIds ):
	delegate = JobProfileDelegate()
	responseData = delegate.removePositions( jobProfileId, PositionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

