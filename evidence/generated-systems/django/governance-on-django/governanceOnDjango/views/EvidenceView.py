import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.EvidenceDelegate import EvidenceDelegate

 #======================================================================
# 
# Encapsulates data for View Evidence
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EvidenceView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Evidence index.")

def get(request, evidenceId ):
	delegate = EvidenceDelegate()
	responseData = delegate.get( evidenceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	evidence = json.loads(request.body)
	delegate = EvidenceDelegate()
	responseData = delegate.createFromJson( evidence )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	evidence = json.loads(request.body)
	delegate = EvidenceDelegate()
	responseData = delegate.save( evidence )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, evidenceId ):
	delegate = EvidenceDelegate()
	responseData = delegate.delete( evidenceId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EvidenceDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignControlTest( request, evidenceId, ControlTestId ):
	delegate = EvidenceDelegate()
	responseData = delegate.saveControlTest( evidenceId, ControlTestId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignControlTest( request, evidenceId ):
	delegate = EvidenceDelegate()
	responseData = delegate.deleteControlTest( evidenceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignControl( request, evidenceId, ControlId ):
	delegate = EvidenceDelegate()
	responseData = delegate.saveControl( evidenceId, ControlId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignControl( request, evidenceId ):
	delegate = EvidenceDelegate()
	responseData = delegate.deleteControl( evidenceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignObligation( request, evidenceId, ObligationId ):
	delegate = EvidenceDelegate()
	responseData = delegate.saveObligation( evidenceId, ObligationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignObligation( request, evidenceId ):
	delegate = EvidenceDelegate()
	responseData = delegate.deleteObligation( evidenceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkpaper( request, evidenceId, WorkpaperId ):
	delegate = EvidenceDelegate()
	responseData = delegate.saveWorkpaper( evidenceId, WorkpaperId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkpaper( request, evidenceId ):
	delegate = EvidenceDelegate()
	responseData = delegate.deleteWorkpaper( evidenceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

