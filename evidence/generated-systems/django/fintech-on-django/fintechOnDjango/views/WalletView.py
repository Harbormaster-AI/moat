import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.WalletDelegate import WalletDelegate

 #======================================================================
# 
# Encapsulates data for View Wallet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WalletView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Wallet index.")

def get(request, walletId ):
	delegate = WalletDelegate()
	responseData = delegate.get( walletId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	wallet = json.loads(request.body)
	delegate = WalletDelegate()
	responseData = delegate.createFromJson( wallet )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	wallet = json.loads(request.body)
	delegate = WalletDelegate()
	responseData = delegate.save( wallet )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, walletId ):
	delegate = WalletDelegate()
	responseData = delegate.delete( walletId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WalletDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, walletId, CustomerId ):
	delegate = WalletDelegate()
	responseData = delegate.saveCustomer( walletId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, walletId ):
	delegate = WalletDelegate()
	responseData = delegate.deleteCustomer( walletId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTransactions( request, walletId, TransactionsIds ):
	delegate = WalletDelegate()
	responseData = delegate.addTransactions( walletId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTransactions( request, walletId, TransactionsIds ):
	delegate = WalletDelegate()
	responseData = delegate.removeTransactions( walletId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

