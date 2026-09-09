import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.PayoutDelegate import PayoutDelegate

 #======================================================================
# 
# Encapsulates data for View Payout
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayoutView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Payout index.")

def get(request, payoutId ):
	delegate = PayoutDelegate()
	responseData = delegate.get( payoutId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	payout = json.loads(request.body)
	delegate = PayoutDelegate()
	responseData = delegate.createFromJson( payout )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	payout = json.loads(request.body)
	delegate = PayoutDelegate()
	responseData = delegate.save( payout )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, payoutId ):
	delegate = PayoutDelegate()
	responseData = delegate.delete( payoutId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PayoutDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSeller( request, payoutId, SellerId ):
	delegate = PayoutDelegate()
	responseData = delegate.saveSeller( payoutId, SellerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSeller( request, payoutId ):
	delegate = PayoutDelegate()
	responseData = delegate.deleteSeller( payoutId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, payoutId, OrdersIds ):
	delegate = PayoutDelegate()
	responseData = delegate.addOrders( payoutId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, payoutId, OrdersIds ):
	delegate = PayoutDelegate()
	responseData = delegate.removeOrders( payoutId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

