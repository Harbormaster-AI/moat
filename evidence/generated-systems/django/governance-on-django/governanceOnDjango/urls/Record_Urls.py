from django.urls import path
from governanceOnDjango.views import Record_View

urlpatterns = [
    path('', Record_View.index, name='index'),
	path('create', Record_View.get, name='create'),
	path('get/<int:record_Id>/', Record_View.get, name='get'),
	path('save', Record_View.save, name='save'),
	path('getAll', Record_View.getAll, name='getAll'),
	path('delete/<int:record_Id>/', Record_View.delete, name='delete'),
	path('assignRepository/<int:record_Id>/<int:RepositoryId>/', Record_View.assignRepository, name='assignRepository'),
	path('unassignRepository/<int:record_Id>/', Record_View.unassignRepository, name='unassignRepository'),
	path('assignRetentionSchedule/<int:record_Id>/<int:RetentionScheduleId>/', Record_View.assignRetentionSchedule, name='assignRetentionSchedule'),
	path('unassignRetentionSchedule/<int:record_Id>/', Record_View.unassignRetentionSchedule, name='unassignRetentionSchedule'),
	path('addProcessingActivities/<int:record_Id>/<ProcessingActivitiesIds>/', Record_View.addProcessingActivities, name='addProcessingActivities'),
	path('removeProcessingActivities/<int:record_Id>/<ProcessingActivitiesIds>/', Record_View.removeProcessingActivities, name='removeProcessingActivities'),
	path('addDataCategories/<int:record_Id>/<DataCategoriesIds>/', Record_View.addDataCategories, name='addDataCategories'),
	path('removeDataCategories/<int:record_Id>/<DataCategoriesIds>/', Record_View.removeDataCategories, name='removeDataCategories'),
	path('addLegalHolds/<int:record_Id>/<LegalHoldsIds>/', Record_View.addLegalHolds, name='addLegalHolds'),
	path('removeLegalHolds/<int:record_Id>/<LegalHoldsIds>/', Record_View.removeLegalHolds, name='removeLegalHolds'),
	path('addDataSubjectRequests/<int:record_Id>/<DataSubjectRequestsIds>/', Record_View.addDataSubjectRequests, name='addDataSubjectRequests'),
	path('removeDataSubjectRequests/<int:record_Id>/<DataSubjectRequestsIds>/', Record_View.removeDataSubjectRequests, name='removeDataSubjectRequests'),
]
