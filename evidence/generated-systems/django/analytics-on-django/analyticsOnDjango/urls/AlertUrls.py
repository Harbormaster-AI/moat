from django.urls import path
from analyticsOnDjango.views import AlertView

urlpatterns = [
    path('', AlertView.index, name='index'),
	path('create', AlertView.get, name='create'),
	path('get/<int:alertId>/', AlertView.get, name='get'),
	path('save', AlertView.save, name='save'),
	path('getAll', AlertView.getAll, name='getAll'),
	path('delete/<int:alertId>/', AlertView.delete, name='delete'),
	path('assignMetric/<int:alertId>/<int:MetricId>/', AlertView.assignMetric, name='assignMetric'),
	path('unassignMetric/<int:alertId>/', AlertView.unassignMetric, name='unassignMetric'),
	path('assignDashboard/<int:alertId>/<int:DashboardId>/', AlertView.assignDashboard, name='assignDashboard'),
	path('unassignDashboard/<int:alertId>/', AlertView.unassignDashboard, name='unassignDashboard'),
	path('assignDataset/<int:alertId>/<int:DatasetId>/', AlertView.assignDataset, name='assignDataset'),
	path('unassignDataset/<int:alertId>/', AlertView.unassignDataset, name='unassignDataset'),
	path('assignRule/<int:alertId>/<int:RuleId>/', AlertView.assignRule, name='assignRule'),
	path('unassignRule/<int:alertId>/', AlertView.unassignRule, name='unassignRule'),
	path('addAnomalies/<int:alertId>/<AnomaliesIds>/', AlertView.addAnomalies, name='addAnomalies'),
	path('removeAnomalies/<int:alertId>/<AnomaliesIds>/', AlertView.removeAnomalies, name='removeAnomalies'),
	path('addSubscribers/<int:alertId>/<SubscribersIds>/', AlertView.addSubscribers, name='addSubscribers'),
	path('removeSubscribers/<int:alertId>/<SubscribersIds>/', AlertView.removeSubscribers, name='removeSubscribers'),
]
