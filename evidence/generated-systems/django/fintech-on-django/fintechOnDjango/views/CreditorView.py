import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.CreditorDelegate import CreditorDelegate

 #======================================================================
# 
# Encapsulates data for View Creditor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreditorView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Creditor index.")

def get(request, creditorId ):
	delegate = CreditorDelegate()
	responseData = delegate.get( creditorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	creditor = json.loads(request.body)
	delegate = CreditorDelegate()
	responseData = delegate.createFromJson( creditor )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	creditor = json.loads(request.body)
	delegate = CreditorDelegate()
	responseData = delegate.save( creditor )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, creditorId ):
	delegate = CreditorDelegate()
	responseData = delegate.delete( creditorId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CreditorDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMandates( request, creditorId, MandatesIds ):
	delegate = CreditorDelegate()
	responseData = delegate.addMandates( creditorId, MandatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMandates( request, creditorId, MandatesIds ):
	delegate = CreditorDelegate()
	responseData = delegate.removeMandates( creditorId, MandatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

