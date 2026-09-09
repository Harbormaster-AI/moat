import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.DirectDebitMandateDelegate import DirectDebitMandateDelegate

 #======================================================================
# 
# Encapsulates data for View DirectDebitMandate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DirectDebitMandateView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DirectDebitMandate index.")

def get(request, directDebitMandateId ):
	delegate = DirectDebitMandateDelegate()
	responseData = delegate.get( directDebitMandateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	directDebitMandate = json.loads(request.body)
	delegate = DirectDebitMandateDelegate()
	responseData = delegate.createFromJson( directDebitMandate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	directDebitMandate = json.loads(request.body)
	delegate = DirectDebitMandateDelegate()
	responseData = delegate.save( directDebitMandate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, directDebitMandateId ):
	delegate = DirectDebitMandateDelegate()
	responseData = delegate.delete( directDebitMandateId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DirectDebitMandateDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, directDebitMandateId, AccountId ):
	delegate = DirectDebitMandateDelegate()
	responseData = delegate.saveAccount( directDebitMandateId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, directDebitMandateId ):
	delegate = DirectDebitMandateDelegate()
	responseData = delegate.deleteAccount( directDebitMandateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCreditor( request, directDebitMandateId, CreditorId ):
	delegate = DirectDebitMandateDelegate()
	responseData = delegate.saveCreditor( directDebitMandateId, CreditorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCreditor( request, directDebitMandateId ):
	delegate = DirectDebitMandateDelegate()
	responseData = delegate.deleteCreditor( directDebitMandateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

