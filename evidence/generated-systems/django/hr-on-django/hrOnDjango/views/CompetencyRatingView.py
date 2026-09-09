import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.CompetencyRatingDelegate import CompetencyRatingDelegate

 #======================================================================
# 
# Encapsulates data for View CompetencyRating
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompetencyRatingView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CompetencyRating index.")

def get(request, competencyRatingId ):
	delegate = CompetencyRatingDelegate()
	responseData = delegate.get( competencyRatingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	competencyRating = json.loads(request.body)
	delegate = CompetencyRatingDelegate()
	responseData = delegate.createFromJson( competencyRating )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	competencyRating = json.loads(request.body)
	delegate = CompetencyRatingDelegate()
	responseData = delegate.save( competencyRating )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, competencyRatingId ):
	delegate = CompetencyRatingDelegate()
	responseData = delegate.delete( competencyRatingId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CompetencyRatingDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignReview( request, competencyRatingId, ReviewId ):
	delegate = CompetencyRatingDelegate()
	responseData = delegate.saveReview( competencyRatingId, ReviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignReview( request, competencyRatingId ):
	delegate = CompetencyRatingDelegate()
	responseData = delegate.deleteReview( competencyRatingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCompetency( request, competencyRatingId, CompetencyId ):
	delegate = CompetencyRatingDelegate()
	responseData = delegate.saveCompetency( competencyRatingId, CompetencyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCompetency( request, competencyRatingId ):
	delegate = CompetencyRatingDelegate()
	responseData = delegate.deleteCompetency( competencyRatingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

