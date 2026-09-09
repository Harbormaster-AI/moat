import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.TradeDelegate import TradeDelegate

 #======================================================================
# 
# Encapsulates data for View Trade
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TradeView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Trade index.")

def get(request, tradeId ):
	delegate = TradeDelegate()
	responseData = delegate.get( tradeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	trade = json.loads(request.body)
	delegate = TradeDelegate()
	responseData = delegate.createFromJson( trade )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	trade = json.loads(request.body)
	delegate = TradeDelegate()
	responseData = delegate.save( trade )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, tradeId ):
	delegate = TradeDelegate()
	responseData = delegate.delete( tradeId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TradeDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, tradeId, OrderId ):
	delegate = TradeDelegate()
	responseData = delegate.saveOrder( tradeId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, tradeId ):
	delegate = TradeDelegate()
	responseData = delegate.deleteOrder( tradeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSecurity( request, tradeId, SecurityId ):
	delegate = TradeDelegate()
	responseData = delegate.saveSecurity( tradeId, SecurityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSecurity( request, tradeId ):
	delegate = TradeDelegate()
	responseData = delegate.deleteSecurity( tradeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInvestmentAccount( request, tradeId, InvestmentAccountId ):
	delegate = TradeDelegate()
	responseData = delegate.saveInvestmentAccount( tradeId, InvestmentAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInvestmentAccount( request, tradeId ):
	delegate = TradeDelegate()
	responseData = delegate.deleteInvestmentAccount( tradeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

