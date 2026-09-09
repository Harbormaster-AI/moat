from django.urls import path
from analyticsOnDjango.views import SubscriberView

urlpatterns = [
    path('', SubscriberView.index, name='index'),
	path('create', SubscriberView.get, name='create'),
	path('get/<int:subscriberId>/', SubscriberView.get, name='get'),
	path('save', SubscriberView.save, name='save'),
	path('getAll', SubscriberView.getAll, name='getAll'),
	path('delete/<int:subscriberId>/', SubscriberView.delete, name='delete'),
	path('addAlerts/<int:subscriberId>/<AlertsIds>/', SubscriberView.addAlerts, name='addAlerts'),
	path('removeAlerts/<int:subscriberId>/<AlertsIds>/', SubscriberView.removeAlerts, name='removeAlerts'),
]
