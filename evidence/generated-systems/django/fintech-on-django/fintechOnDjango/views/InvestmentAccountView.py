import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.InvestmentAccountDelegate import InvestmentAccountDelegate

 #======================================================================
# 
# Encapsulates data for View InvestmentAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvestmentAccountView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InvestmentAccount index.")

def get(request, investmentAccountId ):
	delegate = InvestmentAccountDelegate()
	responseData = delegate.get( investmentAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	investmentAccount = json.loads(request.body)
	delegate = InvestmentAccountDelegate()
	responseData = delegate.createFromJson( investmentAccount )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	investmentAccount = json.loads(request.body)
	delegate = InvestmentAccountDelegate()
	responseData = delegate.save( investmentAccount )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, investmentAccountId ):
	delegate = InvestmentAccountDelegate()
	responseData = delegate.delete( investmentAccountId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InvestmentAccountDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPortfolio( request, investmentAccountId, PortfolioId ):
	delegate = InvestmentAccountDelegate()
	responseData = delegate.savePortfolio( investmentAccountId, PortfolioId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPortfolio( request, investmentAccountId ):
	delegate = InvestmentAccountDelegate()
	responseData = delegate.deletePortfolio( investmentAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTrades( request, investmentAccountId, TradesIds ):
	delegate = InvestmentAccountDelegate()
	responseData = delegate.addTrades( investmentAccountId, TradesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTrades( request, investmentAccountId, TradesIds ):
	delegate = InvestmentAccountDelegate()
	responseData = delegate.removeTrades( investmentAccountId, TradesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, investmentAccountId, OrdersIds ):
	delegate = InvestmentAccountDelegate()
	responseData = delegate.addOrders( investmentAccountId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, investmentAccountId, OrdersIds ):
	delegate = InvestmentAccountDelegate()
	responseData = delegate.removeOrders( investmentAccountId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

