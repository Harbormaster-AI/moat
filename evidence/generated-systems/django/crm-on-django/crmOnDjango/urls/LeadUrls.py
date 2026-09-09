from django.urls import path
from crmOnDjango.views import LeadView

urlpatterns = [
    path('', LeadView.index, name='index'),
	path('create', LeadView.get, name='create'),
	path('get/<int:leadId>/', LeadView.get, name='get'),
	path('save', LeadView.save, name='save'),
	path('getAll', LeadView.getAll, name='getAll'),
	path('delete/<int:leadId>/', LeadView.delete, name='delete'),
	path('assignOrganization/<int:leadId>/<int:OrganizationId>/', LeadView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:leadId>/', LeadView.unassignOrganization, name='unassignOrganization'),
	path('assignOwner/<int:leadId>/<int:OwnerId>/', LeadView.assignOwner, name='assignOwner'),
	path('unassignOwner/<int:leadId>/', LeadView.unassignOwner, name='unassignOwner'),
	path('assignConvertedAccount/<int:leadId>/<int:ConvertedAccountId>/', LeadView.assignConvertedAccount, name='assignConvertedAccount'),
	path('unassignConvertedAccount/<int:leadId>/', LeadView.unassignConvertedAccount, name='unassignConvertedAccount'),
	path('assignConvertedContact/<int:leadId>/<int:ConvertedContactId>/', LeadView.assignConvertedContact, name='assignConvertedContact'),
	path('unassignConvertedContact/<int:leadId>/', LeadView.unassignConvertedContact, name='unassignConvertedContact'),
	path('assignConvertedOpportunity/<int:leadId>/<int:ConvertedOpportunityId>/', LeadView.assignConvertedOpportunity, name='assignConvertedOpportunity'),
	path('unassignConvertedOpportunity/<int:leadId>/', LeadView.unassignConvertedOpportunity, name='unassignConvertedOpportunity'),
	path('addActivities/<int:leadId>/<ActivitiesIds>/', LeadView.addActivities, name='addActivities'),
	path('removeActivities/<int:leadId>/<ActivitiesIds>/', LeadView.removeActivities, name='removeActivities'),
	path('addCampaigns/<int:leadId>/<CampaignsIds>/', LeadView.addCampaigns, name='addCampaigns'),
	path('removeCampaigns/<int:leadId>/<CampaignsIds>/', LeadView.removeCampaigns, name='removeCampaigns'),
	path('addNotes/<int:leadId>/<NotesIds>/', LeadView.addNotes, name='addNotes'),
	path('removeNotes/<int:leadId>/<NotesIds>/', LeadView.removeNotes, name='removeNotes'),
	path('addEmailMessages/<int:leadId>/<EmailMessagesIds>/', LeadView.addEmailMessages, name='addEmailMessages'),
	path('removeEmailMessages/<int:leadId>/<EmailMessagesIds>/', LeadView.removeEmailMessages, name='removeEmailMessages'),
]
