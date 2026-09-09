import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.PaymentContractDelegate import PaymentContractDelegate

 #======================================================================
# 
# Encapsulates data for View PaymentContract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentContractView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PaymentContract index.")

def get(request, paymentContractId ):
	delegate = PaymentContractDelegate()
	responseData = delegate.get( paymentContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	paymentContract = json.loads(request.body)
	delegate = PaymentContractDelegate()
	responseData = delegate.createFromJson( paymentContract )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	paymentContract = json.loads(request.body)
	delegate = PaymentContractDelegate()
	responseData = delegate.save( paymentContract )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, paymentContractId ):
	delegate = PaymentContractDelegate()
	responseData = delegate.delete( paymentContractId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PaymentContractDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, paymentContractId, MerchantId ):
	delegate = PaymentContractDelegate()
	responseData = delegate.saveMerchant( paymentContractId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, paymentContractId ):
	delegate = PaymentContractDelegate()
	responseData = delegate.deleteMerchant( paymentContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAcquirer( request, paymentContractId, AcquirerId ):
	delegate = PaymentContractDelegate()
	responseData = delegate.saveAcquirer( paymentContractId, AcquirerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAcquirer( request, paymentContractId ):
	delegate = PaymentContractDelegate()
	responseData = delegate.deleteAcquirer( paymentContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

