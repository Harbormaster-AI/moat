import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.InvestmentPortfolioDelegate import InvestmentPortfolioDelegate

 #======================================================================
# 
# Encapsulates data for View InvestmentPortfolio
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvestmentPortfolioView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InvestmentPortfolio index.")

def get(request, investmentPortfolioId ):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.get( investmentPortfolioId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	investmentPortfolio = json.loads(request.body)
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.createFromJson( investmentPortfolio )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	investmentPortfolio = json.loads(request.body)
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.save( investmentPortfolio )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, investmentPortfolioId ):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.delete( investmentPortfolioId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, investmentPortfolioId, CustomerId ):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.saveCustomer( investmentPortfolioId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, investmentPortfolioId ):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.deleteCustomer( investmentPortfolioId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAccounts( request, investmentPortfolioId, AccountsIds ):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.addAccounts( investmentPortfolioId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAccounts( request, investmentPortfolioId, AccountsIds ):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.removeAccounts( investmentPortfolioId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, investmentPortfolioId, OrdersIds ):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.addOrders( investmentPortfolioId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, investmentPortfolioId, OrdersIds ):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.removeOrders( investmentPortfolioId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addHoldings( request, investmentPortfolioId, HoldingsIds ):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.addHoldings( investmentPortfolioId, HoldingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeHoldings( request, investmentPortfolioId, HoldingsIds ):
	delegate = InvestmentPortfolioDelegate()
	responseData = delegate.removeHoldings( investmentPortfolioId, HoldingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

