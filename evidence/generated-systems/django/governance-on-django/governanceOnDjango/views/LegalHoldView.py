import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.LegalHoldDelegate import LegalHoldDelegate

 #======================================================================
# 
# Encapsulates data for View LegalHold
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LegalHoldView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the LegalHold index.")

def get(request, legalHoldId ):
	delegate = LegalHoldDelegate()
	responseData = delegate.get( legalHoldId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	legalHold = json.loads(request.body)
	delegate = LegalHoldDelegate()
	responseData = delegate.createFromJson( legalHold )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	legalHold = json.loads(request.body)
	delegate = LegalHoldDelegate()
	responseData = delegate.save( legalHold )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, legalHoldId ):
	delegate = LegalHoldDelegate()
	responseData = delegate.delete( legalHoldId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LegalHoldDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMatter( request, legalHoldId, MatterId ):
	delegate = LegalHoldDelegate()
	responseData = delegate.saveMatter( legalHoldId, MatterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMatter( request, legalHoldId ):
	delegate = LegalHoldDelegate()
	responseData = delegate.deleteMatter( legalHoldId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRepositories( request, legalHoldId, RepositoriesIds ):
	delegate = LegalHoldDelegate()
	responseData = delegate.addRepositories( legalHoldId, RepositoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRepositories( request, legalHoldId, RepositoriesIds ):
	delegate = LegalHoldDelegate()
	responseData = delegate.removeRepositories( legalHoldId, RepositoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRecords( request, legalHoldId, RecordsIds ):
	delegate = LegalHoldDelegate()
	responseData = delegate.addRecords( legalHoldId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRecords( request, legalHoldId, RecordsIds ):
	delegate = LegalHoldDelegate()
	responseData = delegate.removeRecords( legalHoldId, RecordsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

