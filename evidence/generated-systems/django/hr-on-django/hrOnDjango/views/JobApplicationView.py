import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.JobApplicationDelegate import JobApplicationDelegate

 #======================================================================
# 
# Encapsulates data for View JobApplication
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobApplicationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the JobApplication index.")

def get(request, jobApplicationId ):
	delegate = JobApplicationDelegate()
	responseData = delegate.get( jobApplicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	jobApplication = json.loads(request.body)
	delegate = JobApplicationDelegate()
	responseData = delegate.createFromJson( jobApplication )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	jobApplication = json.loads(request.body)
	delegate = JobApplicationDelegate()
	responseData = delegate.save( jobApplication )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, jobApplicationId ):
	delegate = JobApplicationDelegate()
	responseData = delegate.delete( jobApplicationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = JobApplicationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCandidate( request, jobApplicationId, CandidateId ):
	delegate = JobApplicationDelegate()
	responseData = delegate.saveCandidate( jobApplicationId, CandidateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCandidate( request, jobApplicationId ):
	delegate = JobApplicationDelegate()
	responseData = delegate.deleteCandidate( jobApplicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRequisition( request, jobApplicationId, RequisitionId ):
	delegate = JobApplicationDelegate()
	responseData = delegate.saveRequisition( jobApplicationId, RequisitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRequisition( request, jobApplicationId ):
	delegate = JobApplicationDelegate()
	responseData = delegate.deleteRequisition( jobApplicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addScreenings( request, jobApplicationId, ScreeningsIds ):
	delegate = JobApplicationDelegate()
	responseData = delegate.addScreenings( jobApplicationId, ScreeningsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeScreenings( request, jobApplicationId, ScreeningsIds ):
	delegate = JobApplicationDelegate()
	responseData = delegate.removeScreenings( jobApplicationId, ScreeningsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

