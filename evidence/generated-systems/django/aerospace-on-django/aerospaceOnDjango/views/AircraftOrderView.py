import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AircraftOrderDelegate import AircraftOrderDelegate

 #======================================================================
# 
# Encapsulates data for View AircraftOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AircraftOrder index.")

def get(request, aircraftOrderId ):
	delegate = AircraftOrderDelegate()
	responseData = delegate.get( aircraftOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aircraftOrder = json.loads(request.body)
	delegate = AircraftOrderDelegate()
	responseData = delegate.createFromJson( aircraftOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aircraftOrder = json.loads(request.body)
	delegate = AircraftOrderDelegate()
	responseData = delegate.save( aircraftOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aircraftOrderId ):
	delegate = AircraftOrderDelegate()
	responseData = delegate.delete( aircraftOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AircraftOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOperator( request, aircraftOrderId, OperatorId ):
	delegate = AircraftOrderDelegate()
	responseData = delegate.saveOperator( aircraftOrderId, OperatorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOperator( request, aircraftOrderId ):
	delegate = AircraftOrderDelegate()
	responseData = delegate.deleteOperator( aircraftOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignVariant( request, aircraftOrderId, VariantId ):
	delegate = AircraftOrderDelegate()
	responseData = delegate.saveVariant( aircraftOrderId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, aircraftOrderId ):
	delegate = AircraftOrderDelegate()
	responseData = delegate.deleteVariant( aircraftOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignQuote( request, aircraftOrderId, QuoteId ):
	delegate = AircraftOrderDelegate()
	responseData = delegate.saveQuote( aircraftOrderId, QuoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignQuote( request, aircraftOrderId ):
	delegate = AircraftOrderDelegate()
	responseData = delegate.deleteQuote( aircraftOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPurchaseAgreement( request, aircraftOrderId, PurchaseAgreementId ):
	delegate = AircraftOrderDelegate()
	responseData = delegate.savePurchaseAgreement( aircraftOrderId, PurchaseAgreementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPurchaseAgreement( request, aircraftOrderId ):
	delegate = AircraftOrderDelegate()
	responseData = delegate.deletePurchaseAgreement( aircraftOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

