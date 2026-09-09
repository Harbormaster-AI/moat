import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.InsurancePayerDelegate import InsurancePayerDelegate

 #======================================================================
# 
# Encapsulates data for View InsurancePayer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurancePayerView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InsurancePayer index.")

def get(request, insurancePayerId ):
	delegate = InsurancePayerDelegate()
	responseData = delegate.get( insurancePayerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	insurancePayer = json.loads(request.body)
	delegate = InsurancePayerDelegate()
	responseData = delegate.createFromJson( insurancePayer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	insurancePayer = json.loads(request.body)
	delegate = InsurancePayerDelegate()
	responseData = delegate.save( insurancePayer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, insurancePayerId ):
	delegate = InsurancePayerDelegate()
	responseData = delegate.delete( insurancePayerId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InsurancePayerDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPlans( request, insurancePayerId, PlansIds ):
	delegate = InsurancePayerDelegate()
	responseData = delegate.addPlans( insurancePayerId, PlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePlans( request, insurancePayerId, PlansIds ):
	delegate = InsurancePayerDelegate()
	responseData = delegate.removePlans( insurancePayerId, PlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addClaims( request, insurancePayerId, ClaimsIds ):
	delegate = InsurancePayerDelegate()
	responseData = delegate.addClaims( insurancePayerId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeClaims( request, insurancePayerId, ClaimsIds ):
	delegate = InsurancePayerDelegate()
	responseData = delegate.removeClaims( insurancePayerId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

