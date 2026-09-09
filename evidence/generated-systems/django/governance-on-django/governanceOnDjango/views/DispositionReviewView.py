import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.DispositionReviewDelegate import DispositionReviewDelegate

 #======================================================================
# 
# Encapsulates data for View DispositionReview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DispositionReviewView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DispositionReview index.")

def get(request, dispositionReviewId ):
	delegate = DispositionReviewDelegate()
	responseData = delegate.get( dispositionReviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dispositionReview = json.loads(request.body)
	delegate = DispositionReviewDelegate()
	responseData = delegate.createFromJson( dispositionReview )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dispositionReview = json.loads(request.body)
	delegate = DispositionReviewDelegate()
	responseData = delegate.save( dispositionReview )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dispositionReviewId ):
	delegate = DispositionReviewDelegate()
	responseData = delegate.delete( dispositionReviewId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DispositionReviewDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRecord( request, dispositionReviewId, RecordId ):
	delegate = DispositionReviewDelegate()
	responseData = delegate.saveRecord( dispositionReviewId, RecordId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRecord( request, dispositionReviewId ):
	delegate = DispositionReviewDelegate()
	responseData = delegate.deleteRecord( dispositionReviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRetentionSchedule( request, dispositionReviewId, RetentionScheduleId ):
	delegate = DispositionReviewDelegate()
	responseData = delegate.saveRetentionSchedule( dispositionReviewId, RetentionScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRetentionSchedule( request, dispositionReviewId ):
	delegate = DispositionReviewDelegate()
	responseData = delegate.deleteRetentionSchedule( dispositionReviewId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

