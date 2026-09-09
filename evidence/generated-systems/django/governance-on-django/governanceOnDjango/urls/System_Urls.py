from django.urls import path
from governanceOnDjango.views import System_View

urlpatterns = [
    path('', System_View.index, name='index'),
	path('create', System_View.get, name='create'),
	path('get/<int:system_Id>/', System_View.get, name='get'),
	path('save', System_View.save, name='save'),
	path('getAll', System_View.getAll, name='getAll'),
	path('delete/<int:system_Id>/', System_View.delete, name='delete'),
	path('addProcessingActivities/<int:system_Id>/<ProcessingActivitiesIds>/', System_View.addProcessingActivities, name='addProcessingActivities'),
	path('removeProcessingActivities/<int:system_Id>/<ProcessingActivitiesIds>/', System_View.removeProcessingActivities, name='removeProcessingActivities'),
	path('addRecordsRepositories/<int:system_Id>/<RecordsRepositoriesIds>/', System_View.addRecordsRepositories, name='addRecordsRepositories'),
	path('removeRecordsRepositories/<int:system_Id>/<RecordsRepositoriesIds>/', System_View.removeRecordsRepositories, name='removeRecordsRepositories'),
]
