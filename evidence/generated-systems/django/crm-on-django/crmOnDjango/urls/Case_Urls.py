from django.urls import path
from crmOnDjango.views import Case_View

urlpatterns = [
    path('', Case_View.index, name='index'),
	path('create', Case_View.get, name='create'),
	path('get/<int:case_Id>/', Case_View.get, name='get'),
	path('save', Case_View.save, name='save'),
	path('getAll', Case_View.getAll, name='getAll'),
	path('delete/<int:case_Id>/', Case_View.delete, name='delete'),
	path('assignOrganization/<int:case_Id>/<int:OrganizationId>/', Case_View.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:case_Id>/', Case_View.unassignOrganization, name='unassignOrganization'),
	path('assignAccount/<int:case_Id>/<int:AccountId>/', Case_View.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:case_Id>/', Case_View.unassignAccount, name='unassignAccount'),
	path('assignContact/<int:case_Id>/<int:ContactId>/', Case_View.assignContact, name='assignContact'),
	path('unassignContact/<int:case_Id>/', Case_View.unassignContact, name='unassignContact'),
	path('assignOwner/<int:case_Id>/<int:OwnerId>/', Case_View.assignOwner, name='assignOwner'),
	path('unassignOwner/<int:case_Id>/', Case_View.unassignOwner, name='unassignOwner'),
	path('assignTeam/<int:case_Id>/<int:TeamId>/', Case_View.assignTeam, name='assignTeam'),
	path('unassignTeam/<int:case_Id>/', Case_View.unassignTeam, name='unassignTeam'),
	path('addActivities/<int:case_Id>/<ActivitiesIds>/', Case_View.addActivities, name='addActivities'),
	path('removeActivities/<int:case_Id>/<ActivitiesIds>/', Case_View.removeActivities, name='removeActivities'),
	path('addCaseComments/<int:case_Id>/<CaseCommentsIds>/', Case_View.addCaseComments, name='addCaseComments'),
	path('removeCaseComments/<int:case_Id>/<CaseCommentsIds>/', Case_View.removeCaseComments, name='removeCaseComments'),
	path('addEmails/<int:case_Id>/<EmailsIds>/', Case_View.addEmails, name='addEmails'),
	path('removeEmails/<int:case_Id>/<EmailsIds>/', Case_View.removeEmails, name='removeEmails'),
	path('addRelatedOpportunities/<int:case_Id>/<RelatedOpportunitiesIds>/', Case_View.addRelatedOpportunities, name='addRelatedOpportunities'),
	path('removeRelatedOpportunities/<int:case_Id>/<RelatedOpportunitiesIds>/', Case_View.removeRelatedOpportunities, name='removeRelatedOpportunities'),
]
