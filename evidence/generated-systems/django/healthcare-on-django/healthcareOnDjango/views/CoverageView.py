import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.CoverageDelegate import CoverageDelegate

 #======================================================================
# 
# Encapsulates data for View Coverage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CoverageView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Coverage index.")

def get(request, coverageId ):
	delegate = CoverageDelegate()
	responseData = delegate.get( coverageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	coverage = json.loads(request.body)
	delegate = CoverageDelegate()
	responseData = delegate.createFromJson( coverage )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	coverage = json.loads(request.body)
	delegate = CoverageDelegate()
	responseData = delegate.save( coverage )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, coverageId ):
	delegate = CoverageDelegate()
	responseData = delegate.delete( coverageId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CoverageDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, coverageId, PatientId ):
	delegate = CoverageDelegate()
	responseData = delegate.savePatient( coverageId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, coverageId ):
	delegate = CoverageDelegate()
	responseData = delegate.deletePatient( coverageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlan( request, coverageId, PlanId ):
	delegate = CoverageDelegate()
	responseData = delegate.savePlan( coverageId, PlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlan( request, coverageId ):
	delegate = CoverageDelegate()
	responseData = delegate.deletePlan( coverageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addClaims( request, coverageId, ClaimsIds ):
	delegate = CoverageDelegate()
	responseData = delegate.addClaims( coverageId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeClaims( request, coverageId, ClaimsIds ):
	delegate = CoverageDelegate()
	responseData = delegate.removeClaims( coverageId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAuthorizations( request, coverageId, AuthorizationsIds ):
	delegate = CoverageDelegate()
	responseData = delegate.addAuthorizations( coverageId, AuthorizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAuthorizations( request, coverageId, AuthorizationsIds ):
	delegate = CoverageDelegate()
	responseData = delegate.removeAuthorizations( coverageId, AuthorizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

