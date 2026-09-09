import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.AuditProgramDelegate import AuditProgramDelegate

 #======================================================================
# 
# Encapsulates data for View AuditProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditProgramView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AuditProgram index.")

def get(request, auditProgramId ):
	delegate = AuditProgramDelegate()
	responseData = delegate.get( auditProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	auditProgram = json.loads(request.body)
	delegate = AuditProgramDelegate()
	responseData = delegate.createFromJson( auditProgram )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	auditProgram = json.loads(request.body)
	delegate = AuditProgramDelegate()
	responseData = delegate.save( auditProgram )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, auditProgramId ):
	delegate = AuditProgramDelegate()
	responseData = delegate.delete( auditProgramId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AuditProgramDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, auditProgramId, OrganizationId ):
	delegate = AuditProgramDelegate()
	responseData = delegate.saveOrganization( auditProgramId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, auditProgramId ):
	delegate = AuditProgramDelegate()
	responseData = delegate.deleteOrganization( auditProgramId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEngagements( request, auditProgramId, EngagementsIds ):
	delegate = AuditProgramDelegate()
	responseData = delegate.addEngagements( auditProgramId, EngagementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEngagements( request, auditProgramId, EngagementsIds ):
	delegate = AuditProgramDelegate()
	responseData = delegate.removeEngagements( auditProgramId, EngagementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

