from django.urls import path
from governanceOnDjango.views import DataProcessingActivityView

urlpatterns = [
    path('', DataProcessingActivityView.index, name='index'),
	path('create', DataProcessingActivityView.get, name='create'),
	path('get/<int:dataProcessingActivityId>/', DataProcessingActivityView.get, name='get'),
	path('save', DataProcessingActivityView.save, name='save'),
	path('getAll', DataProcessingActivityView.getAll, name='getAll'),
	path('delete/<int:dataProcessingActivityId>/', DataProcessingActivityView.delete, name='delete'),
	path('assignOrganization/<int:dataProcessingActivityId>/<int:OrganizationId>/', DataProcessingActivityView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:dataProcessingActivityId>/', DataProcessingActivityView.unassignOrganization, name='unassignOrganization'),
	path('addDataCategories/<int:dataProcessingActivityId>/<DataCategoriesIds>/', DataProcessingActivityView.addDataCategories, name='addDataCategories'),
	path('removeDataCategories/<int:dataProcessingActivityId>/<DataCategoriesIds>/', DataProcessingActivityView.removeDataCategories, name='removeDataCategories'),
	path('addSystems/<int:dataProcessingActivityId>/<SystemsIds>/', DataProcessingActivityView.addSystems, name='addSystems'),
	path('removeSystems/<int:dataProcessingActivityId>/<SystemsIds>/', DataProcessingActivityView.removeSystems, name='removeSystems'),
	path('addRecords/<int:dataProcessingActivityId>/<RecordsIds>/', DataProcessingActivityView.addRecords, name='addRecords'),
	path('removeRecords/<int:dataProcessingActivityId>/<RecordsIds>/', DataProcessingActivityView.removeRecords, name='removeRecords'),
	path('addPrivacyNotices/<int:dataProcessingActivityId>/<PrivacyNoticesIds>/', DataProcessingActivityView.addPrivacyNotices, name='addPrivacyNotices'),
	path('removePrivacyNotices/<int:dataProcessingActivityId>/<PrivacyNoticesIds>/', DataProcessingActivityView.removePrivacyNotices, name='removePrivacyNotices'),
	path('addThirdParties/<int:dataProcessingActivityId>/<ThirdPartiesIds>/', DataProcessingActivityView.addThirdParties, name='addThirdParties'),
	path('removeThirdParties/<int:dataProcessingActivityId>/<ThirdPartiesIds>/', DataProcessingActivityView.removeThirdParties, name='removeThirdParties'),
	path('addConsents/<int:dataProcessingActivityId>/<ConsentsIds>/', DataProcessingActivityView.addConsents, name='addConsents'),
	path('removeConsents/<int:dataProcessingActivityId>/<ConsentsIds>/', DataProcessingActivityView.removeConsents, name='removeConsents'),
	path('addDataBreaches/<int:dataProcessingActivityId>/<DataBreachesIds>/', DataProcessingActivityView.addDataBreaches, name='addDataBreaches'),
	path('removeDataBreaches/<int:dataProcessingActivityId>/<DataBreachesIds>/', DataProcessingActivityView.removeDataBreaches, name='removeDataBreaches'),
	path('addDataSubjectRequests/<int:dataProcessingActivityId>/<DataSubjectRequestsIds>/', DataProcessingActivityView.addDataSubjectRequests, name='addDataSubjectRequests'),
	path('removeDataSubjectRequests/<int:dataProcessingActivityId>/<DataSubjectRequestsIds>/', DataProcessingActivityView.removeDataSubjectRequests, name='removeDataSubjectRequests'),
]
