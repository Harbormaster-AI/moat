from django.urls import path
from aerospaceOnDjango.views import AircraftOrderView

urlpatterns = [
    path('', AircraftOrderView.index, name='index'),
	path('create', AircraftOrderView.get, name='create'),
	path('get/<int:aircraftOrderId>/', AircraftOrderView.get, name='get'),
	path('save', AircraftOrderView.save, name='save'),
	path('getAll', AircraftOrderView.getAll, name='getAll'),
	path('delete/<int:aircraftOrderId>/', AircraftOrderView.delete, name='delete'),
	path('assignOperator/<int:aircraftOrderId>/<int:OperatorId>/', AircraftOrderView.assignOperator, name='assignOperator'),
	path('unassignOperator/<int:aircraftOrderId>/', AircraftOrderView.unassignOperator, name='unassignOperator'),
	path('assignVariant/<int:aircraftOrderId>/<int:VariantId>/', AircraftOrderView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:aircraftOrderId>/', AircraftOrderView.unassignVariant, name='unassignVariant'),
	path('assignQuote/<int:aircraftOrderId>/<int:QuoteId>/', AircraftOrderView.assignQuote, name='assignQuote'),
	path('unassignQuote/<int:aircraftOrderId>/', AircraftOrderView.unassignQuote, name='unassignQuote'),
	path('assignPurchaseAgreement/<int:aircraftOrderId>/<int:PurchaseAgreementId>/', AircraftOrderView.assignPurchaseAgreement, name='assignPurchaseAgreement'),
	path('unassignPurchaseAgreement/<int:aircraftOrderId>/', AircraftOrderView.unassignPurchaseAgreement, name='unassignPurchaseAgreement'),
]
