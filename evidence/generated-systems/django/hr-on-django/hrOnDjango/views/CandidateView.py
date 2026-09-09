import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.CandidateDelegate import CandidateDelegate

 #======================================================================
# 
# Encapsulates data for View Candidate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CandidateView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Candidate index.")

def get(request, candidateId ):
	delegate = CandidateDelegate()
	responseData = delegate.get( candidateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	candidate = json.loads(request.body)
	delegate = CandidateDelegate()
	responseData = delegate.createFromJson( candidate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	candidate = json.loads(request.body)
	delegate = CandidateDelegate()
	responseData = delegate.save( candidate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, candidateId ):
	delegate = CandidateDelegate()
	responseData = delegate.delete( candidateId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CandidateDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addApplications( request, candidateId, ApplicationsIds ):
	delegate = CandidateDelegate()
	responseData = delegate.addApplications( candidateId, ApplicationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeApplications( request, candidateId, ApplicationsIds ):
	delegate = CandidateDelegate()
	responseData = delegate.removeApplications( candidateId, ApplicationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInterviews( request, candidateId, InterviewsIds ):
	delegate = CandidateDelegate()
	responseData = delegate.addInterviews( candidateId, InterviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInterviews( request, candidateId, InterviewsIds ):
	delegate = CandidateDelegate()
	responseData = delegate.removeInterviews( candidateId, InterviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOffers( request, candidateId, OffersIds ):
	delegate = CandidateDelegate()
	responseData = delegate.addOffers( candidateId, OffersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOffers( request, candidateId, OffersIds ):
	delegate = CandidateDelegate()
	responseData = delegate.removeOffers( candidateId, OffersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDocuments( request, candidateId, DocumentsIds ):
	delegate = CandidateDelegate()
	responseData = delegate.addDocuments( candidateId, DocumentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDocuments( request, candidateId, DocumentsIds ):
	delegate = CandidateDelegate()
	responseData = delegate.removeDocuments( candidateId, DocumentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

