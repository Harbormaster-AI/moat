import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.AuditFindingDelegate import AuditFindingDelegate

 #======================================================================
# 
# Encapsulates data for View AuditFinding
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditFindingView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AuditFinding index.")

def get(request, auditFindingId ):
	delegate = AuditFindingDelegate()
	responseData = delegate.get( auditFindingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	auditFinding = json.loads(request.body)
	delegate = AuditFindingDelegate()
	responseData = delegate.createFromJson( auditFinding )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	auditFinding = json.loads(request.body)
	delegate = AuditFindingDelegate()
	responseData = delegate.save( auditFinding )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, auditFindingId ):
	delegate = AuditFindingDelegate()
	responseData = delegate.delete( auditFindingId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AuditFindingDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEngagement( request, auditFindingId, EngagementId ):
	delegate = AuditFindingDelegate()
	responseData = delegate.saveEngagement( auditFindingId, EngagementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEngagement( request, auditFindingId ):
	delegate = AuditFindingDelegate()
	responseData = delegate.deleteEngagement( auditFindingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkpaper( request, auditFindingId, WorkpaperId ):
	delegate = AuditFindingDelegate()
	responseData = delegate.saveWorkpaper( auditFindingId, WorkpaperId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkpaper( request, auditFindingId ):
	delegate = AuditFindingDelegate()
	responseData = delegate.deleteWorkpaper( auditFindingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCorrectiveActions( request, auditFindingId, CorrectiveActionsIds ):
	delegate = AuditFindingDelegate()
	responseData = delegate.addCorrectiveActions( auditFindingId, CorrectiveActionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCorrectiveActions( request, auditFindingId, CorrectiveActionsIds ):
	delegate = AuditFindingDelegate()
	responseData = delegate.removeCorrectiveActions( auditFindingId, CorrectiveActionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRelatedRisks( request, auditFindingId, RelatedRisksIds ):
	delegate = AuditFindingDelegate()
	responseData = delegate.addRelatedRisks( auditFindingId, RelatedRisksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRelatedRisks( request, auditFindingId, RelatedRisksIds ):
	delegate = AuditFindingDelegate()
	responseData = delegate.removeRelatedRisks( auditFindingId, RelatedRisksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRelatedControls( request, auditFindingId, RelatedControlsIds ):
	delegate = AuditFindingDelegate()
	responseData = delegate.addRelatedControls( auditFindingId, RelatedControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRelatedControls( request, auditFindingId, RelatedControlsIds ):
	delegate = AuditFindingDelegate()
	responseData = delegate.removeRelatedControls( auditFindingId, RelatedControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addIssues( request, auditFindingId, IssuesIds ):
	delegate = AuditFindingDelegate()
	responseData = delegate.addIssues( auditFindingId, IssuesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeIssues( request, auditFindingId, IssuesIds ):
	delegate = AuditFindingDelegate()
	responseData = delegate.removeIssues( auditFindingId, IssuesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

