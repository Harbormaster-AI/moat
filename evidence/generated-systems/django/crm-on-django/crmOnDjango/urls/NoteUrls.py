from django.urls import path
from crmOnDjango.views import NoteView

urlpatterns = [
    path('', NoteView.index, name='index'),
	path('create', NoteView.get, name='create'),
	path('get/<int:noteId>/', NoteView.get, name='get'),
	path('save', NoteView.save, name='save'),
	path('getAll', NoteView.getAll, name='getAll'),
	path('delete/<int:noteId>/', NoteView.delete, name='delete'),
	path('assignOrganization/<int:noteId>/<int:OrganizationId>/', NoteView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:noteId>/', NoteView.unassignOrganization, name='unassignOrganization'),
	path('assignOwner/<int:noteId>/<int:OwnerId>/', NoteView.assignOwner, name='assignOwner'),
	path('unassignOwner/<int:noteId>/', NoteView.unassignOwner, name='unassignOwner'),
	path('assignAccount/<int:noteId>/<int:AccountId>/', NoteView.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:noteId>/', NoteView.unassignAccount, name='unassignAccount'),
	path('assignContact/<int:noteId>/<int:ContactId>/', NoteView.assignContact, name='assignContact'),
	path('unassignContact/<int:noteId>/', NoteView.unassignContact, name='unassignContact'),
	path('assignOpportunity/<int:noteId>/<int:OpportunityId>/', NoteView.assignOpportunity, name='assignOpportunity'),
	path('unassignOpportunity/<int:noteId>/', NoteView.unassignOpportunity, name='unassignOpportunity'),
	path('assignCase/<int:noteId>/<int:CaseId>/', NoteView.assignCase, name='assignCase'),
	path('unassignCase/<int:noteId>/', NoteView.unassignCase, name='unassignCase'),
	path('assignLead/<int:noteId>/<int:LeadId>/', NoteView.assignLead, name='assignLead'),
	path('unassignLead/<int:noteId>/', NoteView.unassignLead, name='unassignLead'),
]
