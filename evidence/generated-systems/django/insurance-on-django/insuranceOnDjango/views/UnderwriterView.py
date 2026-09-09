import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.UnderwriterDelegate import UnderwriterDelegate

 #======================================================================
# 
# Encapsulates data for View Underwriter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UnderwriterView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Underwriter index.")

def get(request, underwriterId ):
	delegate = UnderwriterDelegate()
	responseData = delegate.get( underwriterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	underwriter = json.loads(request.body)
	delegate = UnderwriterDelegate()
	responseData = delegate.createFromJson( underwriter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	underwriter = json.loads(request.body)
	delegate = UnderwriterDelegate()
	responseData = delegate.save( underwriter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, underwriterId ):
	delegate = UnderwriterDelegate()
	responseData = delegate.delete( underwriterId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = UnderwriterDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInsurer( request, underwriterId, InsurerId ):
	delegate = UnderwriterDelegate()
	responseData = delegate.saveInsurer( underwriterId, InsurerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInsurer( request, underwriterId ):
	delegate = UnderwriterDelegate()
	responseData = delegate.deleteInsurer( underwriterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDecisions( request, underwriterId, DecisionsIds ):
	delegate = UnderwriterDelegate()
	responseData = delegate.addDecisions( underwriterId, DecisionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDecisions( request, underwriterId, DecisionsIds ):
	delegate = UnderwriterDelegate()
	responseData = delegate.removeDecisions( underwriterId, DecisionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

