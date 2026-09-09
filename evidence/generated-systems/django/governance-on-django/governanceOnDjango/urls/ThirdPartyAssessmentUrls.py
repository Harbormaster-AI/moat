from django.urls import path
from governanceOnDjango.views import ThirdPartyAssessmentView

urlpatterns = [
    path('', ThirdPartyAssessmentView.index, name='index'),
	path('create', ThirdPartyAssessmentView.get, name='create'),
	path('get/<int:thirdPartyAssessmentId>/', ThirdPartyAssessmentView.get, name='get'),
	path('save', ThirdPartyAssessmentView.save, name='save'),
	path('getAll', ThirdPartyAssessmentView.getAll, name='getAll'),
	path('delete/<int:thirdPartyAssessmentId>/', ThirdPartyAssessmentView.delete, name='delete'),
	path('assignThirdParty/<int:thirdPartyAssessmentId>/<int:ThirdPartyId>/', ThirdPartyAssessmentView.assignThirdParty, name='assignThirdParty'),
	path('unassignThirdParty/<int:thirdPartyAssessmentId>/', ThirdPartyAssessmentView.unassignThirdParty, name='unassignThirdParty'),
	path('addIssues/<int:thirdPartyAssessmentId>/<IssuesIds>/', ThirdPartyAssessmentView.addIssues, name='addIssues'),
	path('removeIssues/<int:thirdPartyAssessmentId>/<IssuesIds>/', ThirdPartyAssessmentView.removeIssues, name='removeIssues'),
]
