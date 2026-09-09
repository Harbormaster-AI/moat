import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.InterviewDelegate import InterviewDelegate

 #======================================================================
# 
# Encapsulates data for View Interview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InterviewView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Interview index.")

def get(request, interviewId ):
	delegate = InterviewDelegate()
	responseData = delegate.get( interviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	interview = json.loads(request.body)
	delegate = InterviewDelegate()
	responseData = delegate.createFromJson( interview )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	interview = json.loads(request.body)
	delegate = InterviewDelegate()
	responseData = delegate.save( interview )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, interviewId ):
	delegate = InterviewDelegate()
	responseData = delegate.delete( interviewId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InterviewDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRequisition( request, interviewId, RequisitionId ):
	delegate = InterviewDelegate()
	responseData = delegate.saveRequisition( interviewId, RequisitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRequisition( request, interviewId ):
	delegate = InterviewDelegate()
	responseData = delegate.deleteRequisition( interviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCandidate( request, interviewId, CandidateId ):
	delegate = InterviewDelegate()
	responseData = delegate.saveCandidate( interviewId, CandidateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCandidate( request, interviewId ):
	delegate = InterviewDelegate()
	responseData = delegate.deleteCandidate( interviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInterviewers( request, interviewId, InterviewersIds ):
	delegate = InterviewDelegate()
	responseData = delegate.addInterviewers( interviewId, InterviewersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInterviewers( request, interviewId, InterviewersIds ):
	delegate = InterviewDelegate()
	responseData = delegate.removeInterviewers( interviewId, InterviewersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

