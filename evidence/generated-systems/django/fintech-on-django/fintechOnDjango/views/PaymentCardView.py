import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.PaymentCardDelegate import PaymentCardDelegate

 #======================================================================
# 
# Encapsulates data for View PaymentCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentCardView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PaymentCard index.")

def get(request, paymentCardId ):
	delegate = PaymentCardDelegate()
	responseData = delegate.get( paymentCardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	paymentCard = json.loads(request.body)
	delegate = PaymentCardDelegate()
	responseData = delegate.createFromJson( paymentCard )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	paymentCard = json.loads(request.body)
	delegate = PaymentCardDelegate()
	responseData = delegate.save( paymentCard )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, paymentCardId ):
	delegate = PaymentCardDelegate()
	responseData = delegate.delete( paymentCardId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PaymentCardDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, paymentCardId, CustomerId ):
	delegate = PaymentCardDelegate()
	responseData = delegate.saveCustomer( paymentCardId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, paymentCardId ):
	delegate = PaymentCardDelegate()
	responseData = delegate.deleteCustomer( paymentCardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, paymentCardId, AccountId ):
	delegate = PaymentCardDelegate()
	responseData = delegate.saveAccount( paymentCardId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, paymentCardId ):
	delegate = PaymentCardDelegate()
	responseData = delegate.deleteAccount( paymentCardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTokenizations( request, paymentCardId, TokenizationsIds ):
	delegate = PaymentCardDelegate()
	responseData = delegate.addTokenizations( paymentCardId, TokenizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTokenizations( request, paymentCardId, TokenizationsIds ):
	delegate = PaymentCardDelegate()
	responseData = delegate.removeTokenizations( paymentCardId, TokenizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDisputes( request, paymentCardId, DisputesIds ):
	delegate = PaymentCardDelegate()
	responseData = delegate.addDisputes( paymentCardId, DisputesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDisputes( request, paymentCardId, DisputesIds ):
	delegate = PaymentCardDelegate()
	responseData = delegate.removeDisputes( paymentCardId, DisputesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

