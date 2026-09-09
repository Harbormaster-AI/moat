from django.urls import path
from hrOnDjango.views import PerformanceCycleView

urlpatterns = [
    path('', PerformanceCycleView.index, name='index'),
	path('create', PerformanceCycleView.get, name='create'),
	path('get/<int:performanceCycleId>/', PerformanceCycleView.get, name='get'),
	path('save', PerformanceCycleView.save, name='save'),
	path('getAll', PerformanceCycleView.getAll, name='getAll'),
	path('delete/<int:performanceCycleId>/', PerformanceCycleView.delete, name='delete'),
	path('assignOrganization/<int:performanceCycleId>/<int:OrganizationId>/', PerformanceCycleView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:performanceCycleId>/', PerformanceCycleView.unassignOrganization, name='unassignOrganization'),
	path('addReviews/<int:performanceCycleId>/<ReviewsIds>/', PerformanceCycleView.addReviews, name='addReviews'),
	path('removeReviews/<int:performanceCycleId>/<ReviewsIds>/', PerformanceCycleView.removeReviews, name='removeReviews'),
	path('addGoals/<int:performanceCycleId>/<GoalsIds>/', PerformanceCycleView.addGoals, name='addGoals'),
	path('removeGoals/<int:performanceCycleId>/<GoalsIds>/', PerformanceCycleView.removeGoals, name='removeGoals'),
]
