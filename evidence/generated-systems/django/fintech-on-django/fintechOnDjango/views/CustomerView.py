import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

 #======================================================================
# 
# Encapsulates data for View Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Customer index.")

def get(request, customerId ):
	delegate = CustomerDelegate()
	responseData = delegate.get( customerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	customer = json.loads(request.body)
	delegate = CustomerDelegate()
	responseData = delegate.createFromJson( customer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	customer = json.loads(request.body)
	delegate = CustomerDelegate()
	responseData = delegate.save( customer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, customerId ):
	delegate = CustomerDelegate()
	responseData = delegate.delete( customerId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CustomerDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInstitution( request, customerId, InstitutionId ):
	delegate = CustomerDelegate()
	responseData = delegate.saveInstitution( customerId, InstitutionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInstitution( request, customerId ):
	delegate = CustomerDelegate()
	responseData = delegate.deleteInstitution( customerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAccounts( request, customerId, AccountsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addAccounts( customerId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAccounts( request, customerId, AccountsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeAccounts( customerId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWallets( request, customerId, WalletsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addWallets( customerId, WalletsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWallets( request, customerId, WalletsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeWallets( customerId, WalletsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCards( request, customerId, CardsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addCards( customerId, CardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCards( request, customerId, CardsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeCards( customerId, CardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addKycProfiles( request, customerId, KycProfilesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addKycProfiles( customerId, KycProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeKycProfiles( request, customerId, KycProfilesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeKycProfiles( customerId, KycProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addConsents( request, customerId, ConsentsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addConsents( customerId, ConsentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeConsents( request, customerId, ConsentsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeConsents( customerId, ConsentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAgreements( request, customerId, AgreementsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addAgreements( customerId, AgreementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAgreements( request, customerId, AgreementsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeAgreements( customerId, AgreementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLoanApplications( request, customerId, LoanApplicationsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addLoanApplications( customerId, LoanApplicationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLoanApplications( request, customerId, LoanApplicationsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeLoanApplications( customerId, LoanApplicationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLoans( request, customerId, LoansIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addLoans( customerId, LoansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLoans( request, customerId, LoansIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeLoans( customerId, LoansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPortfolios( request, customerId, PortfoliosIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addPortfolios( customerId, PortfoliosIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePortfolios( request, customerId, PortfoliosIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removePortfolios( customerId, PortfoliosIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDisputes( request, customerId, DisputesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addDisputes( customerId, DisputesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDisputes( request, customerId, DisputesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeDisputes( customerId, DisputesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

