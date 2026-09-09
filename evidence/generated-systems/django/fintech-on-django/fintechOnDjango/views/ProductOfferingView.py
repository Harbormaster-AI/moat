import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.ProductOfferingDelegate import ProductOfferingDelegate

 #======================================================================
# 
# Encapsulates data for View ProductOffering
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductOfferingView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ProductOffering index.")

def get(request, productOfferingId ):
	delegate = ProductOfferingDelegate()
	responseData = delegate.get( productOfferingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	productOffering = json.loads(request.body)
	delegate = ProductOfferingDelegate()
	responseData = delegate.createFromJson( productOffering )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	productOffering = json.loads(request.body)
	delegate = ProductOfferingDelegate()
	responseData = delegate.save( productOffering )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, productOfferingId ):
	delegate = ProductOfferingDelegate()
	responseData = delegate.delete( productOfferingId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProductOfferingDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInstitution( request, productOfferingId, InstitutionId ):
	delegate = ProductOfferingDelegate()
	responseData = delegate.saveInstitution( productOfferingId, InstitutionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInstitution( request, productOfferingId ):
	delegate = ProductOfferingDelegate()
	responseData = delegate.deleteInstitution( productOfferingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPricingPlans( request, productOfferingId, PricingPlansIds ):
	delegate = ProductOfferingDelegate()
	responseData = delegate.addPricingPlans( productOfferingId, PricingPlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePricingPlans( request, productOfferingId, PricingPlansIds ):
	delegate = ProductOfferingDelegate()
	responseData = delegate.removePricingPlans( productOfferingId, PricingPlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

