from django.urls import path
from fintechOnDjango.views import ComplianceAlertView

urlpatterns = [
    path('', ComplianceAlertView.index, name='index'),
	path('create', ComplianceAlertView.get, name='create'),
	path('get/<int:complianceAlertId>/', ComplianceAlertView.get, name='get'),
	path('save', ComplianceAlertView.save, name='save'),
	path('getAll', ComplianceAlertView.getAll, name='getAll'),
	path('delete/<int:complianceAlertId>/', ComplianceAlertView.delete, name='delete'),
	path('assignScreening/<int:complianceAlertId>/<int:ScreeningId>/', ComplianceAlertView.assignScreening, name='assignScreening'),
	path('unassignScreening/<int:complianceAlertId>/', ComplianceAlertView.unassignScreening, name='unassignScreening'),
	path('assignTransaction/<int:complianceAlertId>/<int:TransactionId>/', ComplianceAlertView.assignTransaction, name='assignTransaction'),
	path('unassignTransaction/<int:complianceAlertId>/', ComplianceAlertView.unassignTransaction, name='unassignTransaction'),
]
