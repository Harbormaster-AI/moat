import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

 #======================================================================
# 
# Encapsulates data for View RecordsRepository
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecordsRepositoryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the RecordsRepository index.")

def get(request, recordsRepositoryId ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.get( recordsRepositoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	recordsRepository = json.loads(request.body)
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.createFromJson( recordsRepository )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	recordsRepository = json.loads(request.body)
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.save( recordsRepository )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, recordsRepositoryId ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.delete( recordsRepositoryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, recordsRepositoryId, OrganizationId ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.saveOrganization( recordsRepositoryId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, recordsRepositoryId ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.deleteOrganization( recordsRepositoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRecords( request, recordsRepositoryId, RecordsIds ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.addRecords( recordsRepositoryId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRecords( request, recordsRepositoryId, RecordsIds ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.removeRecords( recordsRepositoryId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSystems( request, recordsRepositoryId, SystemsIds ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.addSystems( recordsRepositoryId, SystemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSystems( request, recordsRepositoryId, SystemsIds ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.removeSystems( recordsRepositoryId, SystemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRetentionSchedules( request, recordsRepositoryId, RetentionSchedulesIds ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.addRetentionSchedules( recordsRepositoryId, RetentionSchedulesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRetentionSchedules( request, recordsRepositoryId, RetentionSchedulesIds ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.removeRetentionSchedules( recordsRepositoryId, RetentionSchedulesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLegalHolds( request, recordsRepositoryId, LegalHoldsIds ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.addLegalHolds( recordsRepositoryId, LegalHoldsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLegalHolds( request, recordsRepositoryId, LegalHoldsIds ):
	delegate = RecordsRepositoryDelegate()
	responseData = delegate.removeLegalHolds( recordsRepositoryId, LegalHoldsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

