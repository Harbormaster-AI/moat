import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.InsurancePlanDelegate import InsurancePlanDelegate

 #======================================================================
# 
# Encapsulates data for View InsurancePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurancePlanView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InsurancePlan index.")

def get(request, insurancePlanId ):
	delegate = InsurancePlanDelegate()
	responseData = delegate.get( insurancePlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	insurancePlan = json.loads(request.body)
	delegate = InsurancePlanDelegate()
	responseData = delegate.createFromJson( insurancePlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	insurancePlan = json.loads(request.body)
	delegate = InsurancePlanDelegate()
	responseData = delegate.save( insurancePlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, insurancePlanId ):
	delegate = InsurancePlanDelegate()
	responseData = delegate.delete( insurancePlanId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InsurancePlanDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPayer( request, insurancePlanId, PayerId ):
	delegate = InsurancePlanDelegate()
	responseData = delegate.savePayer( insurancePlanId, PayerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPayer( request, insurancePlanId ):
	delegate = InsurancePlanDelegate()
	responseData = delegate.deletePayer( insurancePlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCoverages( request, insurancePlanId, CoveragesIds ):
	delegate = InsurancePlanDelegate()
	responseData = delegate.addCoverages( insurancePlanId, CoveragesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCoverages( request, insurancePlanId, CoveragesIds ):
	delegate = InsurancePlanDelegate()
	responseData = delegate.removeCoverages( insurancePlanId, CoveragesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

