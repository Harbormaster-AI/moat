from django.urls import path
from governanceOnDjango.views import DispositionReviewView

urlpatterns = [
    path('', DispositionReviewView.index, name='index'),
	path('create', DispositionReviewView.get, name='create'),
	path('get/<int:dispositionReviewId>/', DispositionReviewView.get, name='get'),
	path('save', DispositionReviewView.save, name='save'),
	path('getAll', DispositionReviewView.getAll, name='getAll'),
	path('delete/<int:dispositionReviewId>/', DispositionReviewView.delete, name='delete'),
	path('assignRecord/<int:dispositionReviewId>/<int:RecordId>/', DispositionReviewView.assignRecord, name='assignRecord'),
	path('unassignRecord/<int:dispositionReviewId>/', DispositionReviewView.unassignRecord, name='unassignRecord'),
	path('assignRetentionSchedule/<int:dispositionReviewId>/<int:RetentionScheduleId>/', DispositionReviewView.assignRetentionSchedule, name='assignRetentionSchedule'),
	path('unassignRetentionSchedule/<int:dispositionReviewId>/', DispositionReviewView.unassignRetentionSchedule, name='unassignRetentionSchedule'),
]
