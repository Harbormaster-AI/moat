import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.ComplianceAlertDelegate import ComplianceAlertDelegate

 #======================================================================
# 
# Encapsulates data for View ComplianceAlert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceAlertView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ComplianceAlert index.")

def get(request, complianceAlertId ):
	delegate = ComplianceAlertDelegate()
	responseData = delegate.get( complianceAlertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	complianceAlert = json.loads(request.body)
	delegate = ComplianceAlertDelegate()
	responseData = delegate.createFromJson( complianceAlert )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	complianceAlert = json.loads(request.body)
	delegate = ComplianceAlertDelegate()
	responseData = delegate.save( complianceAlert )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, complianceAlertId ):
	delegate = ComplianceAlertDelegate()
	responseData = delegate.delete( complianceAlertId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ComplianceAlertDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignScreening( request, complianceAlertId, ScreeningId ):
	delegate = ComplianceAlertDelegate()
	responseData = delegate.saveScreening( complianceAlertId, ScreeningId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignScreening( request, complianceAlertId ):
	delegate = ComplianceAlertDelegate()
	responseData = delegate.deleteScreening( complianceAlertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTransaction( request, complianceAlertId, TransactionId ):
	delegate = ComplianceAlertDelegate()
	responseData = delegate.saveTransaction( complianceAlertId, TransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTransaction( request, complianceAlertId ):
	delegate = ComplianceAlertDelegate()
	responseData = delegate.deleteTransaction( complianceAlertId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

