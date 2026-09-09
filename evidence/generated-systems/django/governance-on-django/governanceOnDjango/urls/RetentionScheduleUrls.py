from django.urls import path
from governanceOnDjango.views import RetentionScheduleView

urlpatterns = [
    path('', RetentionScheduleView.index, name='index'),
	path('create', RetentionScheduleView.get, name='create'),
	path('get/<int:retentionScheduleId>/', RetentionScheduleView.get, name='get'),
	path('save', RetentionScheduleView.save, name='save'),
	path('getAll', RetentionScheduleView.getAll, name='getAll'),
	path('delete/<int:retentionScheduleId>/', RetentionScheduleView.delete, name='delete'),
	path('addRepositories/<int:retentionScheduleId>/<RepositoriesIds>/', RetentionScheduleView.addRepositories, name='addRepositories'),
	path('removeRepositories/<int:retentionScheduleId>/<RepositoriesIds>/', RetentionScheduleView.removeRepositories, name='removeRepositories'),
	path('addRecords/<int:retentionScheduleId>/<RecordsIds>/', RetentionScheduleView.addRecords, name='addRecords'),
	path('removeRecords/<int:retentionScheduleId>/<RecordsIds>/', RetentionScheduleView.removeRecords, name='removeRecords'),
	path('addExceptions/<int:retentionScheduleId>/<ExceptionsIds>/', RetentionScheduleView.addExceptions, name='addExceptions'),
	path('removeExceptions/<int:retentionScheduleId>/<ExceptionsIds>/', RetentionScheduleView.removeExceptions, name='removeExceptions'),
	path('addDispositionReviews/<int:retentionScheduleId>/<DispositionReviewsIds>/', RetentionScheduleView.addDispositionReviews, name='addDispositionReviews'),
	path('removeDispositionReviews/<int:retentionScheduleId>/<DispositionReviewsIds>/', RetentionScheduleView.removeDispositionReviews, name='removeDispositionReviews'),
]
