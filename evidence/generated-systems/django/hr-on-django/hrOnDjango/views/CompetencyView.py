import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.CompetencyDelegate import CompetencyDelegate

 #======================================================================
# 
# Encapsulates data for View Competency
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompetencyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Competency index.")

def get(request, competencyId ):
	delegate = CompetencyDelegate()
	responseData = delegate.get( competencyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	competency = json.loads(request.body)
	delegate = CompetencyDelegate()
	responseData = delegate.createFromJson( competency )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	competency = json.loads(request.body)
	delegate = CompetencyDelegate()
	responseData = delegate.save( competency )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, competencyId ):
	delegate = CompetencyDelegate()
	responseData = delegate.delete( competencyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CompetencyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addJobProfiles( request, competencyId, JobProfilesIds ):
	delegate = CompetencyDelegate()
	responseData = delegate.addJobProfiles( competencyId, JobProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeJobProfiles( request, competencyId, JobProfilesIds ):
	delegate = CompetencyDelegate()
	responseData = delegate.removeJobProfiles( competencyId, JobProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCompetencyRatings( request, competencyId, CompetencyRatingsIds ):
	delegate = CompetencyDelegate()
	responseData = delegate.addCompetencyRatings( competencyId, CompetencyRatingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCompetencyRatings( request, competencyId, CompetencyRatingsIds ):
	delegate = CompetencyDelegate()
	responseData = delegate.removeCompetencyRatings( competencyId, CompetencyRatingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

