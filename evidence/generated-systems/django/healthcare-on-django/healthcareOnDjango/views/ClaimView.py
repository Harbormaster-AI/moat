import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.ClaimDelegate import ClaimDelegate

 #======================================================================
# 
# Encapsulates data for View Claim
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Claim index.")

def get(request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.get( claimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	claim = json.loads(request.body)
	delegate = ClaimDelegate()
	responseData = delegate.createFromJson( claim )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	claim = json.loads(request.body)
	delegate = ClaimDelegate()
	responseData = delegate.save( claim )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.delete( claimId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ClaimDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, claimId, PatientId ):
	delegate = ClaimDelegate()
	responseData = delegate.savePatient( claimId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.deletePatient( claimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCoverage( request, claimId, CoverageId ):
	delegate = ClaimDelegate()
	responseData = delegate.saveCoverage( claimId, CoverageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCoverage( request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.deleteCoverage( claimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEncounter( request, claimId, EncounterId ):
	delegate = ClaimDelegate()
	responseData = delegate.saveEncounter( claimId, EncounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEncounter( request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.deleteEncounter( claimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPayer( request, claimId, PayerId ):
	delegate = ClaimDelegate()
	responseData = delegate.savePayer( claimId, PayerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPayer( request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.deletePayer( claimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInvoices( request, claimId, InvoicesIds ):
	delegate = ClaimDelegate()
	responseData = delegate.addInvoices( claimId, InvoicesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInvoices( request, claimId, InvoicesIds ):
	delegate = ClaimDelegate()
	responseData = delegate.removeInvoices( claimId, InvoicesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

