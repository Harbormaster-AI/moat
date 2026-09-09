import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.RegulationDelegate import RegulationDelegate

 #======================================================================
# 
# Encapsulates data for View Regulation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RegulationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Regulation index.")

def get(request, regulationId ):
	delegate = RegulationDelegate()
	responseData = delegate.get( regulationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	regulation = json.loads(request.body)
	delegate = RegulationDelegate()
	responseData = delegate.createFromJson( regulation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	regulation = json.loads(request.body)
	delegate = RegulationDelegate()
	responseData = delegate.save( regulation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, regulationId ):
	delegate = RegulationDelegate()
	responseData = delegate.delete( regulationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RegulationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addObligations( request, regulationId, ObligationsIds ):
	delegate = RegulationDelegate()
	responseData = delegate.addObligations( regulationId, ObligationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeObligations( request, regulationId, ObligationsIds ):
	delegate = RegulationDelegate()
	responseData = delegate.removeObligations( regulationId, ObligationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCompliancePrograms( request, regulationId, ComplianceProgramsIds ):
	delegate = RegulationDelegate()
	responseData = delegate.addCompliancePrograms( regulationId, ComplianceProgramsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCompliancePrograms( request, regulationId, ComplianceProgramsIds ):
	delegate = RegulationDelegate()
	responseData = delegate.removeCompliancePrograms( regulationId, ComplianceProgramsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

