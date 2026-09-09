from django.urls import path
from governanceOnDjango.views import LegalHoldView

urlpatterns = [
    path('', LegalHoldView.index, name='index'),
	path('create', LegalHoldView.get, name='create'),
	path('get/<int:legalHoldId>/', LegalHoldView.get, name='get'),
	path('save', LegalHoldView.save, name='save'),
	path('getAll', LegalHoldView.getAll, name='getAll'),
	path('delete/<int:legalHoldId>/', LegalHoldView.delete, name='delete'),
	path('assignMatter/<int:legalHoldId>/<int:MatterId>/', LegalHoldView.assignMatter, name='assignMatter'),
	path('unassignMatter/<int:legalHoldId>/', LegalHoldView.unassignMatter, name='unassignMatter'),
	path('addRepositories/<int:legalHoldId>/<RepositoriesIds>/', LegalHoldView.addRepositories, name='addRepositories'),
	path('removeRepositories/<int:legalHoldId>/<RepositoriesIds>/', LegalHoldView.removeRepositories, name='removeRepositories'),
	path('addRecords/<int:legalHoldId>/<RecordsIds>/', LegalHoldView.addRecords, name='addRecords'),
	path('removeRecords/<int:legalHoldId>/<RecordsIds>/', LegalHoldView.removeRecords, name='removeRecords'),
]
