import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.PaymentProcessorDelegate import PaymentProcessorDelegate

 #======================================================================
# 
# Encapsulates data for View PaymentProcessor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentProcessorView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PaymentProcessor index.")

def get(request, paymentProcessorId ):
	delegate = PaymentProcessorDelegate()
	responseData = delegate.get( paymentProcessorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	paymentProcessor = json.loads(request.body)
	delegate = PaymentProcessorDelegate()
	responseData = delegate.createFromJson( paymentProcessor )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	paymentProcessor = json.loads(request.body)
	delegate = PaymentProcessorDelegate()
	responseData = delegate.save( paymentProcessor )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, paymentProcessorId ):
	delegate = PaymentProcessorDelegate()
	responseData = delegate.delete( paymentProcessorId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PaymentProcessorDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInstitutions( request, paymentProcessorId, InstitutionsIds ):
	delegate = PaymentProcessorDelegate()
	responseData = delegate.addInstitutions( paymentProcessorId, InstitutionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInstitutions( request, paymentProcessorId, InstitutionsIds ):
	delegate = PaymentProcessorDelegate()
	responseData = delegate.removeInstitutions( paymentProcessorId, InstitutionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContracts( request, paymentProcessorId, ContractsIds ):
	delegate = PaymentProcessorDelegate()
	responseData = delegate.addContracts( paymentProcessorId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContracts( request, paymentProcessorId, ContractsIds ):
	delegate = PaymentProcessorDelegate()
	responseData = delegate.removeContracts( paymentProcessorId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSettlements( request, paymentProcessorId, SettlementsIds ):
	delegate = PaymentProcessorDelegate()
	responseData = delegate.addSettlements( paymentProcessorId, SettlementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSettlements( request, paymentProcessorId, SettlementsIds ):
	delegate = PaymentProcessorDelegate()
	responseData = delegate.removeSettlements( paymentProcessorId, SettlementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

