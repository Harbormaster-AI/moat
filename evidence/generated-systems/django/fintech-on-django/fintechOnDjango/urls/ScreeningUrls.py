from django.urls import path
from fintechOnDjango.views import ScreeningView

urlpatterns = [
    path('', ScreeningView.index, name='index'),
	path('create', ScreeningView.get, name='create'),
	path('get/<int:screeningId>/', ScreeningView.get, name='get'),
	path('save', ScreeningView.save, name='save'),
	path('getAll', ScreeningView.getAll, name='getAll'),
	path('delete/<int:screeningId>/', ScreeningView.delete, name='delete'),
	path('assignKycProfile/<int:screeningId>/<int:KycProfileId>/', ScreeningView.assignKycProfile, name='assignKycProfile'),
	path('unassignKycProfile/<int:screeningId>/', ScreeningView.unassignKycProfile, name='unassignKycProfile'),
	path('addAlerts/<int:screeningId>/<AlertsIds>/', ScreeningView.addAlerts, name='addAlerts'),
	path('removeAlerts/<int:screeningId>/<AlertsIds>/', ScreeningView.removeAlerts, name='removeAlerts'),
]
