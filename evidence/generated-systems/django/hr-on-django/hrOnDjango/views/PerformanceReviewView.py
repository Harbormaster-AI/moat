import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.PerformanceReviewDelegate import PerformanceReviewDelegate

 #======================================================================
# 
# Encapsulates data for View PerformanceReview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceReviewView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PerformanceReview index.")

def get(request, performanceReviewId ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.get( performanceReviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	performanceReview = json.loads(request.body)
	delegate = PerformanceReviewDelegate()
	responseData = delegate.createFromJson( performanceReview )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	performanceReview = json.loads(request.body)
	delegate = PerformanceReviewDelegate()
	responseData = delegate.save( performanceReview )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, performanceReviewId ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.delete( performanceReviewId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, performanceReviewId, EmployeeId ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.saveEmployee( performanceReviewId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, performanceReviewId ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.deleteEmployee( performanceReviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignReviewer( request, performanceReviewId, ReviewerId ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.saveReviewer( performanceReviewId, ReviewerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignReviewer( request, performanceReviewId ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.deleteReviewer( performanceReviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCycle( request, performanceReviewId, CycleId ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.saveCycle( performanceReviewId, CycleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCycle( request, performanceReviewId ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.deleteCycle( performanceReviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCompetencyRatings( request, performanceReviewId, CompetencyRatingsIds ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.addCompetencyRatings( performanceReviewId, CompetencyRatingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCompetencyRatings( request, performanceReviewId, CompetencyRatingsIds ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.removeCompetencyRatings( performanceReviewId, CompetencyRatingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGoals( request, performanceReviewId, GoalsIds ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.addGoals( performanceReviewId, GoalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGoals( request, performanceReviewId, GoalsIds ):
	delegate = PerformanceReviewDelegate()
	responseData = delegate.removeGoals( performanceReviewId, GoalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

