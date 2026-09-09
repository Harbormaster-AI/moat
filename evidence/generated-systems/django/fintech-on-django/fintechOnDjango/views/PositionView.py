import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.PositionDelegate import PositionDelegate

 #======================================================================
# 
# Encapsulates data for View Position
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PositionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Position index.")

def get(request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.get( positionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	position = json.loads(request.body)
	delegate = PositionDelegate()
	responseData = delegate.createFromJson( position )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	position = json.loads(request.body)
	delegate = PositionDelegate()
	responseData = delegate.save( position )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.delete( positionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PositionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPortfolio( request, positionId, PortfolioId ):
	delegate = PositionDelegate()
	responseData = delegate.savePortfolio( positionId, PortfolioId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPortfolio( request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.deletePortfolio( positionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSecurity( request, positionId, SecurityId ):
	delegate = PositionDelegate()
	responseData = delegate.saveSecurity( positionId, SecurityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSecurity( request, positionId ):
	delegate = PositionDelegate()
	responseData = delegate.deleteSecurity( positionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

