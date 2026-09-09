import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.BonusPlanDelegate import BonusPlanDelegate

 #======================================================================
# 
# Encapsulates data for View BonusPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BonusPlanView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BonusPlan index.")

def get(request, bonusPlanId ):
	delegate = BonusPlanDelegate()
	responseData = delegate.get( bonusPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	bonusPlan = json.loads(request.body)
	delegate = BonusPlanDelegate()
	responseData = delegate.createFromJson( bonusPlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	bonusPlan = json.loads(request.body)
	delegate = BonusPlanDelegate()
	responseData = delegate.save( bonusPlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, bonusPlanId ):
	delegate = BonusPlanDelegate()
	responseData = delegate.delete( bonusPlanId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BonusPlanDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCompensationPackages( request, bonusPlanId, CompensationPackagesIds ):
	delegate = BonusPlanDelegate()
	responseData = delegate.addCompensationPackages( bonusPlanId, CompensationPackagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCompensationPackages( request, bonusPlanId, CompensationPackagesIds ):
	delegate = BonusPlanDelegate()
	responseData = delegate.removeCompensationPackages( bonusPlanId, CompensationPackagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

