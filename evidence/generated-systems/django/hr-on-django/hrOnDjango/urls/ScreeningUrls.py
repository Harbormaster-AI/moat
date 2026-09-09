from django.urls import path
from hrOnDjango.views import ScreeningView

urlpatterns = [
    path('', ScreeningView.index, name='index'),
	path('create', ScreeningView.get, name='create'),
	path('get/<int:screeningId>/', ScreeningView.get, name='get'),
	path('save', ScreeningView.save, name='save'),
	path('getAll', ScreeningView.getAll, name='getAll'),
	path('delete/<int:screeningId>/', ScreeningView.delete, name='delete'),
	path('assignApplication/<int:screeningId>/<int:ApplicationId>/', ScreeningView.assignApplication, name='assignApplication'),
	path('unassignApplication/<int:screeningId>/', ScreeningView.unassignApplication, name='unassignApplication'),
]
