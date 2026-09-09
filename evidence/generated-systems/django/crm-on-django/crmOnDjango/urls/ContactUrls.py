from django.urls import path
from crmOnDjango.views import ContactView

urlpatterns = [
    path('', ContactView.index, name='index'),
	path('create', ContactView.get, name='create'),
	path('get/<int:contactId>/', ContactView.get, name='get'),
	path('save', ContactView.save, name='save'),
	path('getAll', ContactView.getAll, name='getAll'),
	path('delete/<int:contactId>/', ContactView.delete, name='delete'),
	path('assignOrganization/<int:contactId>/<int:OrganizationId>/', ContactView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:contactId>/', ContactView.unassignOrganization, name='unassignOrganization'),
	path('assignAccount/<int:contactId>/<int:AccountId>/', ContactView.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:contactId>/', ContactView.unassignAccount, name='unassignAccount'),
	path('assignOwner/<int:contactId>/<int:OwnerId>/', ContactView.assignOwner, name='assignOwner'),
	path('unassignOwner/<int:contactId>/', ContactView.unassignOwner, name='unassignOwner'),
	path('addActivities/<int:contactId>/<ActivitiesIds>/', ContactView.addActivities, name='addActivities'),
	path('removeActivities/<int:contactId>/<ActivitiesIds>/', ContactView.removeActivities, name='removeActivities'),
	path('addOpportunities/<int:contactId>/<OpportunitiesIds>/', ContactView.addOpportunities, name='addOpportunities'),
	path('removeOpportunities/<int:contactId>/<OpportunitiesIds>/', ContactView.removeOpportunities, name='removeOpportunities'),
	path('addCases/<int:contactId>/<CasesIds>/', ContactView.addCases, name='addCases'),
	path('removeCases/<int:contactId>/<CasesIds>/', ContactView.removeCases, name='removeCases'),
	path('addCampaigns/<int:contactId>/<CampaignsIds>/', ContactView.addCampaigns, name='addCampaigns'),
	path('removeCampaigns/<int:contactId>/<CampaignsIds>/', ContactView.removeCampaigns, name='removeCampaigns'),
	path('addNotes/<int:contactId>/<NotesIds>/', ContactView.addNotes, name='addNotes'),
	path('removeNotes/<int:contactId>/<NotesIds>/', ContactView.removeNotes, name='removeNotes'),
	path('addEmailMessages/<int:contactId>/<EmailMessagesIds>/', ContactView.addEmailMessages, name='addEmailMessages'),
	path('removeEmailMessages/<int:contactId>/<EmailMessagesIds>/', ContactView.removeEmailMessages, name='removeEmailMessages'),
]
