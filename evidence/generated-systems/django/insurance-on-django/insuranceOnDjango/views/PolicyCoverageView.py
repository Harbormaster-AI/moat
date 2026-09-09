import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.PolicyCoverageDelegate import PolicyCoverageDelegate

 #======================================================================
# 
# Encapsulates data for View PolicyCoverage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyCoverageView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PolicyCoverage index.")

def get(request, policyCoverageId ):
	delegate = PolicyCoverageDelegate()
	responseData = delegate.get( policyCoverageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	policyCoverage = json.loads(request.body)
	delegate = PolicyCoverageDelegate()
	responseData = delegate.createFromJson( policyCoverage )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	policyCoverage = json.loads(request.body)
	delegate = PolicyCoverageDelegate()
	responseData = delegate.save( policyCoverage )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, policyCoverageId ):
	delegate = PolicyCoverageDelegate()
	responseData = delegate.delete( policyCoverageId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PolicyCoverageDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, policyCoverageId, PolicyId ):
	delegate = PolicyCoverageDelegate()
	responseData = delegate.savePolicy( policyCoverageId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, policyCoverageId ):
	delegate = PolicyCoverageDelegate()
	responseData = delegate.deletePolicy( policyCoverageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInsuredObjects( request, policyCoverageId, InsuredObjectsIds ):
	delegate = PolicyCoverageDelegate()
	responseData = delegate.addInsuredObjects( policyCoverageId, InsuredObjectsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInsuredObjects( request, policyCoverageId, InsuredObjectsIds ):
	delegate = PolicyCoverageDelegate()
	responseData = delegate.removeInsuredObjects( policyCoverageId, InsuredObjectsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

