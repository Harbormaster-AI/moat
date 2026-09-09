import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.PayoutDelegate import PayoutDelegate

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

def assignMerchant( request, payoutId, MerchantId ):
	delegate = PayoutDelegate()
	responseData = delegate.saveMerchant( payoutId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, payoutId ):
	delegate = PayoutDelegate()
	responseData = delegate.deleteMerchant( payoutId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSettlementBatch( request, payoutId, SettlementBatchId ):
	delegate = PayoutDelegate()
	responseData = delegate.saveSettlementBatch( payoutId, SettlementBatchId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSettlementBatch( request, payoutId ):
	delegate = PayoutDelegate()
	responseData = delegate.deleteSettlementBatch( payoutId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDestinationAccount( request, payoutId, DestinationAccountId ):
	delegate = PayoutDelegate()
	responseData = delegate.saveDestinationAccount( payoutId, DestinationAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDestinationAccount( request, payoutId ):
	delegate = PayoutDelegate()
	responseData = delegate.deleteDestinationAccount( payoutId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

