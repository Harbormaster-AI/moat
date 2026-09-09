import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.DisputeDelegate import DisputeDelegate

 #======================================================================
# 
# Encapsulates data for View Dispute
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DisputeView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Dispute index.")

def get(request, disputeId ):
	delegate = DisputeDelegate()
	responseData = delegate.get( disputeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dispute = json.loads(request.body)
	delegate = DisputeDelegate()
	responseData = delegate.createFromJson( dispute )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dispute = json.loads(request.body)
	delegate = DisputeDelegate()
	responseData = delegate.save( dispute )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, disputeId ):
	delegate = DisputeDelegate()
	responseData = delegate.delete( disputeId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DisputeDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTransaction( request, disputeId, TransactionId ):
	delegate = DisputeDelegate()
	responseData = delegate.saveTransaction( disputeId, TransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTransaction( request, disputeId ):
	delegate = DisputeDelegate()
	responseData = delegate.deleteTransaction( disputeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCard( request, disputeId, CardId ):
	delegate = DisputeDelegate()
	responseData = delegate.saveCard( disputeId, CardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCard( request, disputeId ):
	delegate = DisputeDelegate()
	responseData = delegate.deleteCard( disputeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, disputeId, MerchantId ):
	delegate = DisputeDelegate()
	responseData = delegate.saveMerchant( disputeId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, disputeId ):
	delegate = DisputeDelegate()
	responseData = delegate.deleteMerchant( disputeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChargebacks( request, disputeId, ChargebacksIds ):
	delegate = DisputeDelegate()
	responseData = delegate.addChargebacks( disputeId, ChargebacksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChargebacks( request, disputeId, ChargebacksIds ):
	delegate = DisputeDelegate()
	responseData = delegate.removeChargebacks( disputeId, ChargebacksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

