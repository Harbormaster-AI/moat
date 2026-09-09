from django.urls import path
from manufacturingOnDjango.views import CorrectiveActionView

urlpatterns = [
    path('', CorrectiveActionView.index, name='index'),
	path('create', CorrectiveActionView.get, name='create'),
	path('get/<int:correctiveActionId>/', CorrectiveActionView.get, name='get'),
	path('save', CorrectiveActionView.save, name='save'),
	path('getAll', CorrectiveActionView.getAll, name='getAll'),
	path('delete/<int:correctiveActionId>/', CorrectiveActionView.delete, name='delete'),
	path('assignNonconformance/<int:correctiveActionId>/<int:NonconformanceId>/', CorrectiveActionView.assignNonconformance, name='assignNonconformance'),
	path('unassignNonconformance/<int:correctiveActionId>/', CorrectiveActionView.unassignNonconformance, name='unassignNonconformance'),
	path('assignOwner/<int:correctiveActionId>/<int:OwnerId>/', CorrectiveActionView.assignOwner, name='assignOwner'),
	path('unassignOwner/<int:correctiveActionId>/', CorrectiveActionView.unassignOwner, name='unassignOwner'),
]
