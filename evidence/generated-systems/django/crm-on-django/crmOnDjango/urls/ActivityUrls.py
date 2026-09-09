from django.urls import path
from crmOnDjango.views import ActivityView

urlpatterns = [
    path('', ActivityView.index, name='index'),
	path('create', ActivityView.get, name='create'),
	path('get/<int:activityId>/', ActivityView.get, name='get'),
	path('save', ActivityView.save, name='save'),
	path('getAll', ActivityView.getAll, name='getAll'),
	path('delete/<int:activityId>/', ActivityView.delete, name='delete'),
	path('assignOrganization/<int:activityId>/<int:OrganizationId>/', ActivityView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:activityId>/', ActivityView.unassignOrganization, name='unassignOrganization'),
	path('assignOwner/<int:activityId>/<int:OwnerId>/', ActivityView.assignOwner, name='assignOwner'),
	path('unassignOwner/<int:activityId>/', ActivityView.unassignOwner, name='unassignOwner'),
	path('assignAccount/<int:activityId>/<int:AccountId>/', ActivityView.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:activityId>/', ActivityView.unassignAccount, name='unassignAccount'),
	path('assignContact/<int:activityId>/<int:ContactId>/', ActivityView.assignContact, name='assignContact'),
	path('unassignContact/<int:activityId>/', ActivityView.unassignContact, name='unassignContact'),
	path('assignLead/<int:activityId>/<int:LeadId>/', ActivityView.assignLead, name='assignLead'),
	path('unassignLead/<int:activityId>/', ActivityView.unassignLead, name='unassignLead'),
	path('assignOpportunity/<int:activityId>/<int:OpportunityId>/', ActivityView.assignOpportunity, name='assignOpportunity'),
	path('unassignOpportunity/<int:activityId>/', ActivityView.unassignOpportunity, name='unassignOpportunity'),
	path('assignCase/<int:activityId>/<int:CaseId>/', ActivityView.assignCase, name='assignCase'),
	path('unassignCase/<int:activityId>/', ActivityView.unassignCase, name='unassignCase'),
	path('assignCampaign/<int:activityId>/<int:CampaignId>/', ActivityView.assignCampaign, name='assignCampaign'),
	path('unassignCampaign/<int:activityId>/', ActivityView.unassignCampaign, name='unassignCampaign'),
]
