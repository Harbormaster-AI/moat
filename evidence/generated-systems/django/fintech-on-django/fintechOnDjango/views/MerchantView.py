import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.MerchantDelegate import MerchantDelegate

 #======================================================================
# 
# Encapsulates data for View Merchant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MerchantView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Merchant index.")

def get(request, merchantId ):
	delegate = MerchantDelegate()
	responseData = delegate.get( merchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	merchant = json.loads(request.body)
	delegate = MerchantDelegate()
	responseData = delegate.createFromJson( merchant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	merchant = json.loads(request.body)
	delegate = MerchantDelegate()
	responseData = delegate.save( merchant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, merchantId ):
	delegate = MerchantDelegate()
	responseData = delegate.delete( merchantId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MerchantDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTerminals( request, merchantId, TerminalsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addTerminals( merchantId, TerminalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTerminals( request, merchantId, TerminalsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removeTerminals( merchantId, TerminalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPaymentContracts( request, merchantId, PaymentContractsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addPaymentContracts( merchantId, PaymentContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePaymentContracts( request, merchantId, PaymentContractsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removePaymentContracts( merchantId, PaymentContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayouts( request, merchantId, PayoutsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addPayouts( merchantId, PayoutsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayouts( request, merchantId, PayoutsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removePayouts( merchantId, PayoutsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSettlements( request, merchantId, SettlementsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addSettlements( merchantId, SettlementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSettlements( request, merchantId, SettlementsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removeSettlements( merchantId, SettlementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDisputes( request, merchantId, DisputesIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addDisputes( merchantId, DisputesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDisputes( request, merchantId, DisputesIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removeDisputes( merchantId, DisputesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInvoices( request, merchantId, InvoicesIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addInvoices( merchantId, InvoicesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInvoices( request, merchantId, InvoicesIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removeInvoices( merchantId, InvoicesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

