from django.urls import path
from crmOnDjango.views import EmailMessageView

urlpatterns = [
    path('', EmailMessageView.index, name='index'),
	path('create', EmailMessageView.get, name='create'),
	path('get/<int:emailMessageId>/', EmailMessageView.get, name='get'),
	path('save', EmailMessageView.save, name='save'),
	path('getAll', EmailMessageView.getAll, name='getAll'),
	path('delete/<int:emailMessageId>/', EmailMessageView.delete, name='delete'),
	path('assignOrganization/<int:emailMessageId>/<int:OrganizationId>/', EmailMessageView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:emailMessageId>/', EmailMessageView.unassignOrganization, name='unassignOrganization'),
	path('assignOwner/<int:emailMessageId>/<int:OwnerId>/', EmailMessageView.assignOwner, name='assignOwner'),
	path('unassignOwner/<int:emailMessageId>/', EmailMessageView.unassignOwner, name='unassignOwner'),
	path('assignAccount/<int:emailMessageId>/<int:AccountId>/', EmailMessageView.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:emailMessageId>/', EmailMessageView.unassignAccount, name='unassignAccount'),
	path('assignContact/<int:emailMessageId>/<int:ContactId>/', EmailMessageView.assignContact, name='assignContact'),
	path('unassignContact/<int:emailMessageId>/', EmailMessageView.unassignContact, name='unassignContact'),
	path('assignLead/<int:emailMessageId>/<int:LeadId>/', EmailMessageView.assignLead, name='assignLead'),
	path('unassignLead/<int:emailMessageId>/', EmailMessageView.unassignLead, name='unassignLead'),
	path('assignCase/<int:emailMessageId>/<int:CaseId>/', EmailMessageView.assignCase, name='assignCase'),
	path('unassignCase/<int:emailMessageId>/', EmailMessageView.unassignCase, name='unassignCase'),
	path('assignOpportunity/<int:emailMessageId>/<int:OpportunityId>/', EmailMessageView.assignOpportunity, name='assignOpportunity'),
	path('unassignOpportunity/<int:emailMessageId>/', EmailMessageView.unassignOpportunity, name='unassignOpportunity'),
	path('assignCampaign/<int:emailMessageId>/<int:CampaignId>/', EmailMessageView.assignCampaign, name='assignCampaign'),
	path('unassignCampaign/<int:emailMessageId>/', EmailMessageView.unassignCampaign, name='unassignCampaign'),
]
