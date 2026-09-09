import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.AuditWorkpaperDelegate import AuditWorkpaperDelegate

 #======================================================================
# 
# Encapsulates data for View AuditWorkpaper
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditWorkpaperView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AuditWorkpaper index.")

def get(request, auditWorkpaperId ):
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.get( auditWorkpaperId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	auditWorkpaper = json.loads(request.body)
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.createFromJson( auditWorkpaper )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	auditWorkpaper = json.loads(request.body)
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.save( auditWorkpaper )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, auditWorkpaperId ):
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.delete( auditWorkpaperId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEngagement( request, auditWorkpaperId, EngagementId ):
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.saveEngagement( auditWorkpaperId, EngagementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEngagement( request, auditWorkpaperId ):
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.deleteEngagement( auditWorkpaperId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEvidence( request, auditWorkpaperId, EvidenceIds ):
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.addEvidence( auditWorkpaperId, EvidenceIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEvidence( request, auditWorkpaperId, EvidenceIds ):
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.removeEvidence( auditWorkpaperId, EvidenceIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFindings( request, auditWorkpaperId, FindingsIds ):
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.addFindings( auditWorkpaperId, FindingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFindings( request, auditWorkpaperId, FindingsIds ):
	delegate = AuditWorkpaperDelegate()
	responseData = delegate.removeFindings( auditWorkpaperId, FindingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

