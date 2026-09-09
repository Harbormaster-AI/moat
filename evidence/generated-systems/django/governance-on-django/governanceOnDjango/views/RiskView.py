import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

 #======================================================================
# 
# Encapsulates data for View Risk
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Risk index.")

def get(request, riskId ):
	delegate = RiskDelegate()
	responseData = delegate.get( riskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	risk = json.loads(request.body)
	delegate = RiskDelegate()
	responseData = delegate.createFromJson( risk )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	risk = json.loads(request.body)
	delegate = RiskDelegate()
	responseData = delegate.save( risk )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, riskId ):
	delegate = RiskDelegate()
	responseData = delegate.delete( riskId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RiskDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, riskId, OrganizationId ):
	delegate = RiskDelegate()
	responseData = delegate.saveOrganization( riskId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, riskId ):
	delegate = RiskDelegate()
	responseData = delegate.deleteOrganization( riskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addControls( request, riskId, ControlsIds ):
	delegate = RiskDelegate()
	responseData = delegate.addControls( riskId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeControls( request, riskId, ControlsIds ):
	delegate = RiskDelegate()
	responseData = delegate.removeControls( riskId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAssessments( request, riskId, AssessmentsIds ):
	delegate = RiskDelegate()
	responseData = delegate.addAssessments( riskId, AssessmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAssessments( request, riskId, AssessmentsIds ):
	delegate = RiskDelegate()
	responseData = delegate.removeAssessments( riskId, AssessmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addIssues( request, riskId, IssuesIds ):
	delegate = RiskDelegate()
	responseData = delegate.addIssues( riskId, IssuesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeIssues( request, riskId, IssuesIds ):
	delegate = RiskDelegate()
	responseData = delegate.removeIssues( riskId, IssuesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFindings( request, riskId, FindingsIds ):
	delegate = RiskDelegate()
	responseData = delegate.addFindings( riskId, FindingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFindings( request, riskId, FindingsIds ):
	delegate = RiskDelegate()
	responseData = delegate.removeFindings( riskId, FindingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

