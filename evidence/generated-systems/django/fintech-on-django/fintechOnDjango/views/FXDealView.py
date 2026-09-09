import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.FXDealDelegate import FXDealDelegate

 #======================================================================
# 
# Encapsulates data for View FXDeal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FXDealView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the FXDeal index.")

def get(request, fXDealId ):
	delegate = FXDealDelegate()
	responseData = delegate.get( fXDealId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	fXDeal = json.loads(request.body)
	delegate = FXDealDelegate()
	responseData = delegate.createFromJson( fXDeal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	fXDeal = json.loads(request.body)
	delegate = FXDealDelegate()
	responseData = delegate.save( fXDeal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, fXDealId ):
	delegate = FXDealDelegate()
	responseData = delegate.delete( fXDealId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FXDealDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignQuote( request, fXDealId, QuoteId ):
	delegate = FXDealDelegate()
	responseData = delegate.saveQuote( fXDealId, QuoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignQuote( request, fXDealId ):
	delegate = FXDealDelegate()
	responseData = delegate.deleteQuote( fXDealId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPaymentOrders( request, fXDealId, PaymentOrdersIds ):
	delegate = FXDealDelegate()
	responseData = delegate.addPaymentOrders( fXDealId, PaymentOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePaymentOrders( request, fXDealId, PaymentOrdersIds ):
	delegate = FXDealDelegate()
	responseData = delegate.removePaymentOrders( fXDealId, PaymentOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

