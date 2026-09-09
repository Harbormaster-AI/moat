from django.urls import path
from governanceOnDjango.views import DataSubjectRequestView

urlpatterns = [
    path('', DataSubjectRequestView.index, name='index'),
	path('create', DataSubjectRequestView.get, name='create'),
	path('get/<int:dataSubjectRequestId>/', DataSubjectRequestView.get, name='get'),
	path('save', DataSubjectRequestView.save, name='save'),
	path('getAll', DataSubjectRequestView.getAll, name='getAll'),
	path('delete/<int:dataSubjectRequestId>/', DataSubjectRequestView.delete, name='delete'),
	path('assignOrganization/<int:dataSubjectRequestId>/<int:OrganizationId>/', DataSubjectRequestView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:dataSubjectRequestId>/', DataSubjectRequestView.unassignOrganization, name='unassignOrganization'),
	path('addProcessingActivities/<int:dataSubjectRequestId>/<ProcessingActivitiesIds>/', DataSubjectRequestView.addProcessingActivities, name='addProcessingActivities'),
	path('removeProcessingActivities/<int:dataSubjectRequestId>/<ProcessingActivitiesIds>/', DataSubjectRequestView.removeProcessingActivities, name='removeProcessingActivities'),
	path('addRecords/<int:dataSubjectRequestId>/<RecordsIds>/', DataSubjectRequestView.addRecords, name='addRecords'),
	path('removeRecords/<int:dataSubjectRequestId>/<RecordsIds>/', DataSubjectRequestView.removeRecords, name='removeRecords'),
]
