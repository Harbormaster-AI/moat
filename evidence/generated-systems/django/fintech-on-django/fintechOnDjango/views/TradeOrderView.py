import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.TradeOrderDelegate import TradeOrderDelegate

 #======================================================================
# 
# Encapsulates data for View TradeOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TradeOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TradeOrder index.")

def get(request, tradeOrderId ):
	delegate = TradeOrderDelegate()
	responseData = delegate.get( tradeOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	tradeOrder = json.loads(request.body)
	delegate = TradeOrderDelegate()
	responseData = delegate.createFromJson( tradeOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	tradeOrder = json.loads(request.body)
	delegate = TradeOrderDelegate()
	responseData = delegate.save( tradeOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, tradeOrderId ):
	delegate = TradeOrderDelegate()
	responseData = delegate.delete( tradeOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TradeOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPortfolio( request, tradeOrderId, PortfolioId ):
	delegate = TradeOrderDelegate()
	responseData = delegate.savePortfolio( tradeOrderId, PortfolioId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPortfolio( request, tradeOrderId ):
	delegate = TradeOrderDelegate()
	responseData = delegate.deletePortfolio( tradeOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSecurity( request, tradeOrderId, SecurityId ):
	delegate = TradeOrderDelegate()
	responseData = delegate.saveSecurity( tradeOrderId, SecurityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSecurity( request, tradeOrderId ):
	delegate = TradeOrderDelegate()
	responseData = delegate.deleteSecurity( tradeOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTrades( request, tradeOrderId, TradesIds ):
	delegate = TradeOrderDelegate()
	responseData = delegate.addTrades( tradeOrderId, TradesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTrades( request, tradeOrderId, TradesIds ):
	delegate = TradeOrderDelegate()
	responseData = delegate.removeTrades( tradeOrderId, TradesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

