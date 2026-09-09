import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.PaymentOrderDelegate import PaymentOrderDelegate

 #======================================================================
# 
# Encapsulates data for View PaymentOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PaymentOrder index.")

def get(request, paymentOrderId ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.get( paymentOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	paymentOrder = json.loads(request.body)
	delegate = PaymentOrderDelegate()
	responseData = delegate.createFromJson( paymentOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	paymentOrder = json.loads(request.body)
	delegate = PaymentOrderDelegate()
	responseData = delegate.save( paymentOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, paymentOrderId ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.delete( paymentOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PaymentOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSourceAccount( request, paymentOrderId, SourceAccountId ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.saveSourceAccount( paymentOrderId, SourceAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSourceAccount( request, paymentOrderId ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.deleteSourceAccount( paymentOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDestinationAccount( request, paymentOrderId, DestinationAccountId ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.saveDestinationAccount( paymentOrderId, DestinationAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDestinationAccount( request, paymentOrderId ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.deleteDestinationAccount( paymentOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBeneficiary( request, paymentOrderId, BeneficiaryId ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.saveBeneficiary( paymentOrderId, BeneficiaryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBeneficiary( request, paymentOrderId ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.deleteBeneficiary( paymentOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFxDeal( request, paymentOrderId, FxDealId ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.saveFxDeal( paymentOrderId, FxDealId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFxDeal( request, paymentOrderId ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.deleteFxDeal( paymentOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTransactions( request, paymentOrderId, TransactionsIds ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.addTransactions( paymentOrderId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTransactions( request, paymentOrderId, TransactionsIds ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.removeTransactions( paymentOrderId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFees( request, paymentOrderId, FeesIds ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.addFees( paymentOrderId, FeesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFees( request, paymentOrderId, FeesIds ):
	delegate = PaymentOrderDelegate()
	responseData = delegate.removeFees( paymentOrderId, FeesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

