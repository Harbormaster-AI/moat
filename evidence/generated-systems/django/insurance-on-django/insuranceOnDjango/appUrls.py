"""mainsite URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include
urlpatterns = [
    path('Insurer/', include('insuranceOnDjango.urls.InsurerUrls')),
    path('InsuranceProduct/', include('insuranceOnDjango.urls.InsuranceProductUrls')),
    path('CoverageDefinition/', include('insuranceOnDjango.urls.CoverageDefinitionUrls')),
    path('Distributor/', include('insuranceOnDjango.urls.DistributorUrls')),
    path('Agent/', include('insuranceOnDjango.urls.AgentUrls')),
    path('Customer/', include('insuranceOnDjango.urls.CustomerUrls')),
    path('Application/', include('insuranceOnDjango.urls.ApplicationUrls')),
    path('Quote/', include('insuranceOnDjango.urls.QuoteUrls')),
    path('UnderwritingDecision/', include('insuranceOnDjango.urls.UnderwritingDecisionUrls')),
    path('Underwriter/', include('insuranceOnDjango.urls.UnderwriterUrls')),
    path('Policy/', include('insuranceOnDjango.urls.PolicyUrls')),
    path('Endorsement/', include('insuranceOnDjango.urls.EndorsementUrls')),
    path('PolicyCoverage/', include('insuranceOnDjango.urls.PolicyCoverageUrls')),
    path('InsuredObject/', include('insuranceOnDjango.urls.InsuredObjectUrls')),
    path('Beneficiary/', include('insuranceOnDjango.urls.BeneficiaryUrls')),
    path('BillingAccount/', include('insuranceOnDjango.urls.BillingAccountUrls')),
    path('Invoice/', include('insuranceOnDjango.urls.InvoiceUrls')),
    path('Payment/', include('insuranceOnDjango.urls.PaymentUrls')),
    path('Claim/', include('insuranceOnDjango.urls.ClaimUrls')),
    path('Incident/', include('insuranceOnDjango.urls.IncidentUrls')),
    path('Exposure/', include('insuranceOnDjango.urls.ExposureUrls')),
    path('Adjuster/', include('insuranceOnDjango.urls.AdjusterUrls')),
    path('ClaimReserve/', include('insuranceOnDjango.urls.ClaimReserveUrls')),
    path('ClaimPayment/', include('insuranceOnDjango.urls.ClaimPaymentUrls')),
    path('ServiceProvider/', include('insuranceOnDjango.urls.ServiceProviderUrls')),
    path('ReinsuranceAgreement/', include('insuranceOnDjango.urls.ReinsuranceAgreementUrls')),
    path('SubrogationRecovery/', include('insuranceOnDjango.urls.SubrogationRecoveryUrls')),
    path('ThirdParty/', include('insuranceOnDjango.urls.ThirdPartyUrls')),
    path('Document/', include('insuranceOnDjango.urls.DocumentUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]