import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.AuditEngagementDelegate import AuditEngagementDelegate

 #======================================================================
# 
# Encapsulates data for View AuditEngagement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditEngagementView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AuditEngagement index.")

def get(request, auditEngagementId ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.get( auditEngagementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	auditEngagement = json.loads(request.body)
	delegate = AuditEngagementDelegate()
	responseData = delegate.createFromJson( auditEngagement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	auditEngagement = json.loads(request.body)
	delegate = AuditEngagementDelegate()
	responseData = delegate.save( auditEngagement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, auditEngagementId ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.delete( auditEngagementId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AuditEngagementDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAuditProgram( request, auditEngagementId, AuditProgramId ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.saveAuditProgram( auditEngagementId, AuditProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAuditProgram( request, auditEngagementId ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.deleteAuditProgram( auditEngagementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBusinessUnits( request, auditEngagementId, BusinessUnitsIds ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.addBusinessUnits( auditEngagementId, BusinessUnitsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBusinessUnits( request, auditEngagementId, BusinessUnitsIds ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.removeBusinessUnits( auditEngagementId, BusinessUnitsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addControlTests( request, auditEngagementId, ControlTestsIds ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.addControlTests( auditEngagementId, ControlTestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeControlTests( request, auditEngagementId, ControlTestsIds ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.removeControlTests( auditEngagementId, ControlTestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWorkpapers( request, auditEngagementId, WorkpapersIds ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.addWorkpapers( auditEngagementId, WorkpapersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWorkpapers( request, auditEngagementId, WorkpapersIds ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.removeWorkpapers( auditEngagementId, WorkpapersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFindings( request, auditEngagementId, FindingsIds ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.addFindings( auditEngagementId, FindingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFindings( request, auditEngagementId, FindingsIds ):
	delegate = AuditEngagementDelegate()
	responseData = delegate.removeFindings( auditEngagementId, FindingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

