import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

 #======================================================================
# 
# Encapsulates data for View Transaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransactionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Transaction index.")

def get(request, transactionId ):
	delegate = TransactionDelegate()
	responseData = delegate.get( transactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	transaction = json.loads(request.body)
	delegate = TransactionDelegate()
	responseData = delegate.createFromJson( transaction )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	transaction = json.loads(request.body)
	delegate = TransactionDelegate()
	responseData = delegate.save( transaction )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, transactionId ):
	delegate = TransactionDelegate()
	responseData = delegate.delete( transactionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TransactionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, transactionId, AccountId ):
	delegate = TransactionDelegate()
	responseData = delegate.saveAccount( transactionId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, transactionId ):
	delegate = TransactionDelegate()
	responseData = delegate.deleteAccount( transactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWallet( request, transactionId, WalletId ):
	delegate = TransactionDelegate()
	responseData = delegate.saveWallet( transactionId, WalletId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWallet( request, transactionId ):
	delegate = TransactionDelegate()
	responseData = delegate.deleteWallet( transactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPaymentOrder( request, transactionId, PaymentOrderId ):
	delegate = TransactionDelegate()
	responseData = delegate.savePaymentOrder( transactionId, PaymentOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPaymentOrder( request, transactionId ):
	delegate = TransactionDelegate()
	responseData = delegate.deletePaymentOrder( transactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, transactionId, MerchantId ):
	delegate = TransactionDelegate()
	responseData = delegate.saveMerchant( transactionId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, transactionId ):
	delegate = TransactionDelegate()
	responseData = delegate.deleteMerchant( transactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCard( request, transactionId, CardId ):
	delegate = TransactionDelegate()
	responseData = delegate.saveCard( transactionId, CardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCard( request, transactionId ):
	delegate = TransactionDelegate()
	responseData = delegate.deleteCard( transactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRelatedTransactions( request, transactionId, RelatedTransactionsIds ):
	delegate = TransactionDelegate()
	responseData = delegate.addRelatedTransactions( transactionId, RelatedTransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRelatedTransactions( request, transactionId, RelatedTransactionsIds ):
	delegate = TransactionDelegate()
	responseData = delegate.removeRelatedTransactions( transactionId, RelatedTransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAlerts( request, transactionId, AlertsIds ):
	delegate = TransactionDelegate()
	responseData = delegate.addAlerts( transactionId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAlerts( request, transactionId, AlertsIds ):
	delegate = TransactionDelegate()
	responseData = delegate.removeAlerts( transactionId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

