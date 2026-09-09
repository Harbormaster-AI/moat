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
    path('Organization/', include('crmOnDjango.urls.OrganizationUrls')),
    path('User/', include('crmOnDjango.urls.UserUrls')),
    path('Team/', include('crmOnDjango.urls.TeamUrls')),
    path('Territory/', include('crmOnDjango.urls.TerritoryUrls')),
    path('Account/', include('crmOnDjango.urls.AccountUrls')),
    path('Contact/', include('crmOnDjango.urls.ContactUrls')),
    path('Lead/', include('crmOnDjango.urls.LeadUrls')),
    path('Opportunity/', include('crmOnDjango.urls.OpportunityUrls')),
    path('OpportunityLineItem/', include('crmOnDjango.urls.OpportunityLineItemUrls')),
    path('OpportunityStageHistory/', include('crmOnDjango.urls.OpportunityStageHistoryUrls')),
    path('Product/', include('crmOnDjango.urls.ProductUrls')),
    path('PriceBook/', include('crmOnDjango.urls.PriceBookUrls')),
    path('PriceBookEntry/', include('crmOnDjango.urls.PriceBookEntryUrls')),
    path('Quote/', include('crmOnDjango.urls.QuoteUrls')),
    path('QuoteLineItem/', include('crmOnDjango.urls.QuoteLineItemUrls')),
    path('Order/', include('crmOnDjango.urls.OrderUrls')),
    path('OrderItem/', include('crmOnDjango.urls.OrderItemUrls')),
    path('Contract/', include('crmOnDjango.urls.ContractUrls')),
    path('Case_/', include('crmOnDjango.urls.Case_Urls')),
    path('Activity/', include('crmOnDjango.urls.ActivityUrls')),
    path('Campaign/', include('crmOnDjango.urls.CampaignUrls')),
    path('CampaignMember/', include('crmOnDjango.urls.CampaignMemberUrls')),
    path('Note/', include('crmOnDjango.urls.NoteUrls')),
    path('EmailMessage/', include('crmOnDjango.urls.EmailMessageUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]