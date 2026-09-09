import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.SettlementBatchDelegate import SettlementBatchDelegate

 #======================================================================
# 
# Encapsulates data for View SettlementBatch
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SettlementBatchView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SettlementBatch index.")

def get(request, settlementBatchId ):
	delegate = SettlementBatchDelegate()
	responseData = delegate.get( settlementBatchId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	settlementBatch = json.loads(request.body)
	delegate = SettlementBatchDelegate()
	responseData = delegate.createFromJson( settlementBatch )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	settlementBatch = json.loads(request.body)
	delegate = SettlementBatchDelegate()
	responseData = delegate.save( settlementBatch )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, settlementBatchId ):
	delegate = SettlementBatchDelegate()
	responseData = delegate.delete( settlementBatchId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SettlementBatchDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProcessor( request, settlementBatchId, ProcessorId ):
	delegate = SettlementBatchDelegate()
	responseData = delegate.saveProcessor( settlementBatchId, ProcessorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProcessor( request, settlementBatchId ):
	delegate = SettlementBatchDelegate()
	responseData = delegate.deleteProcessor( settlementBatchId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, settlementBatchId, MerchantId ):
	delegate = SettlementBatchDelegate()
	responseData = delegate.saveMerchant( settlementBatchId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, settlementBatchId ):
	delegate = SettlementBatchDelegate()
	responseData = delegate.deleteMerchant( settlementBatchId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayouts( request, settlementBatchId, PayoutsIds ):
	delegate = SettlementBatchDelegate()
	responseData = delegate.addPayouts( settlementBatchId, PayoutsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayouts( request, settlementBatchId, PayoutsIds ):
	delegate = SettlementBatchDelegate()
	responseData = delegate.removePayouts( settlementBatchId, PayoutsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTransactions( request, settlementBatchId, TransactionsIds ):
	delegate = SettlementBatchDelegate()
	responseData = delegate.addTransactions( settlementBatchId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTransactions( request, settlementBatchId, TransactionsIds ):
	delegate = SettlementBatchDelegate()
	responseData = delegate.removeTransactions( settlementBatchId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

