import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.RetentionScheduleDelegate import RetentionScheduleDelegate

 #======================================================================
# 
# Encapsulates data for View RetentionSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RetentionScheduleView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the RetentionSchedule index.")

def get(request, retentionScheduleId ):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.get( retentionScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	retentionSchedule = json.loads(request.body)
	delegate = RetentionScheduleDelegate()
	responseData = delegate.createFromJson( retentionSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	retentionSchedule = json.loads(request.body)
	delegate = RetentionScheduleDelegate()
	responseData = delegate.save( retentionSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, retentionScheduleId ):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.delete( retentionScheduleId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRepositories( request, retentionScheduleId, RepositoriesIds ):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.addRepositories( retentionScheduleId, RepositoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRepositories( request, retentionScheduleId, RepositoriesIds ):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.removeRepositories( retentionScheduleId, RepositoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRecords( request, retentionScheduleId, RecordsIds ):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.addRecords( retentionScheduleId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRecords( request, retentionScheduleId, RecordsIds ):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.removeRecords( retentionScheduleId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addExceptions( request, retentionScheduleId, ExceptionsIds ):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.addExceptions( retentionScheduleId, ExceptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeExceptions( request, retentionScheduleId, ExceptionsIds ):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.removeExceptions( retentionScheduleId, ExceptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDispositionReviews( request, retentionScheduleId, DispositionReviewsIds ):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.addDispositionReviews( retentionScheduleId, DispositionReviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDispositionReviews( request, retentionScheduleId, DispositionReviewsIds ):
	delegate = RetentionScheduleDelegate()
	responseData = delegate.removeDispositionReviews( retentionScheduleId, DispositionReviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

