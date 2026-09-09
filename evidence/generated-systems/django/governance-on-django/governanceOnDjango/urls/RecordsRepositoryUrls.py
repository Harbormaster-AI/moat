from django.urls import path
from governanceOnDjango.views import RecordsRepositoryView

urlpatterns = [
    path('', RecordsRepositoryView.index, name='index'),
	path('create', RecordsRepositoryView.get, name='create'),
	path('get/<int:recordsRepositoryId>/', RecordsRepositoryView.get, name='get'),
	path('save', RecordsRepositoryView.save, name='save'),
	path('getAll', RecordsRepositoryView.getAll, name='getAll'),
	path('delete/<int:recordsRepositoryId>/', RecordsRepositoryView.delete, name='delete'),
	path('assignOrganization/<int:recordsRepositoryId>/<int:OrganizationId>/', RecordsRepositoryView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:recordsRepositoryId>/', RecordsRepositoryView.unassignOrganization, name='unassignOrganization'),
	path('addRecords/<int:recordsRepositoryId>/<RecordsIds>/', RecordsRepositoryView.addRecords, name='addRecords'),
	path('removeRecords/<int:recordsRepositoryId>/<RecordsIds>/', RecordsRepositoryView.removeRecords, name='removeRecords'),
	path('addSystems/<int:recordsRepositoryId>/<SystemsIds>/', RecordsRepositoryView.addSystems, name='addSystems'),
	path('removeSystems/<int:recordsRepositoryId>/<SystemsIds>/', RecordsRepositoryView.removeSystems, name='removeSystems'),
	path('addRetentionSchedules/<int:recordsRepositoryId>/<RetentionSchedulesIds>/', RecordsRepositoryView.addRetentionSchedules, name='addRetentionSchedules'),
	path('removeRetentionSchedules/<int:recordsRepositoryId>/<RetentionSchedulesIds>/', RecordsRepositoryView.removeRetentionSchedules, name='removeRetentionSchedules'),
	path('addLegalHolds/<int:recordsRepositoryId>/<LegalHoldsIds>/', RecordsRepositoryView.addLegalHolds, name='addLegalHolds'),
	path('removeLegalHolds/<int:recordsRepositoryId>/<LegalHoldsIds>/', RecordsRepositoryView.removeLegalHolds, name='removeLegalHolds'),
]
