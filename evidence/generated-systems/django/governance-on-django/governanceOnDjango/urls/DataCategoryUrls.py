from django.urls import path
from governanceOnDjango.views import DataCategoryView

urlpatterns = [
    path('', DataCategoryView.index, name='index'),
	path('create', DataCategoryView.get, name='create'),
	path('get/<int:dataCategoryId>/', DataCategoryView.get, name='get'),
	path('save', DataCategoryView.save, name='save'),
	path('getAll', DataCategoryView.getAll, name='getAll'),
	path('delete/<int:dataCategoryId>/', DataCategoryView.delete, name='delete'),
	path('addProcessingActivities/<int:dataCategoryId>/<ProcessingActivitiesIds>/', DataCategoryView.addProcessingActivities, name='addProcessingActivities'),
	path('removeProcessingActivities/<int:dataCategoryId>/<ProcessingActivitiesIds>/', DataCategoryView.removeProcessingActivities, name='removeProcessingActivities'),
	path('addRecords/<int:dataCategoryId>/<RecordsIds>/', DataCategoryView.addRecords, name='addRecords'),
	path('removeRecords/<int:dataCategoryId>/<RecordsIds>/', DataCategoryView.removeRecords, name='removeRecords'),
	path('addDataBreaches/<int:dataCategoryId>/<DataBreachesIds>/', DataCategoryView.addDataBreaches, name='addDataBreaches'),
	path('removeDataBreaches/<int:dataCategoryId>/<DataBreachesIds>/', DataCategoryView.removeDataBreaches, name='removeDataBreaches'),
]
