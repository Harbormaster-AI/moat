import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.PurchaseAgreementDelegate import PurchaseAgreementDelegate

 #======================================================================
# 
# Encapsulates data for View PurchaseAgreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseAgreementView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PurchaseAgreement index.")

def get(request, purchaseAgreementId ):
	delegate = PurchaseAgreementDelegate()
	responseData = delegate.get( purchaseAgreementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	purchaseAgreement = json.loads(request.body)
	delegate = PurchaseAgreementDelegate()
	responseData = delegate.createFromJson( purchaseAgreement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	purchaseAgreement = json.loads(request.body)
	delegate = PurchaseAgreementDelegate()
	responseData = delegate.save( purchaseAgreement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, purchaseAgreementId ):
	delegate = PurchaseAgreementDelegate()
	responseData = delegate.delete( purchaseAgreementId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PurchaseAgreementDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAircraftOrder( request, purchaseAgreementId, AircraftOrderId ):
	delegate = PurchaseAgreementDelegate()
	responseData = delegate.saveAircraftOrder( purchaseAgreementId, AircraftOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAircraftOrder( request, purchaseAgreementId ):
	delegate = PurchaseAgreementDelegate()
	responseData = delegate.deleteAircraftOrder( purchaseAgreementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

