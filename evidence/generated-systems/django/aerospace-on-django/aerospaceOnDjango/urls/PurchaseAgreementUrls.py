from django.urls import path
from aerospaceOnDjango.views import PurchaseAgreementView

urlpatterns = [
    path('', PurchaseAgreementView.index, name='index'),
	path('create', PurchaseAgreementView.get, name='create'),
	path('get/<int:purchaseAgreementId>/', PurchaseAgreementView.get, name='get'),
	path('save', PurchaseAgreementView.save, name='save'),
	path('getAll', PurchaseAgreementView.getAll, name='getAll'),
	path('delete/<int:purchaseAgreementId>/', PurchaseAgreementView.delete, name='delete'),
	path('assignAircraftOrder/<int:purchaseAgreementId>/<int:AircraftOrderId>/', PurchaseAgreementView.assignAircraftOrder, name='assignAircraftOrder'),
	path('unassignAircraftOrder/<int:purchaseAgreementId>/', PurchaseAgreementView.unassignAircraftOrder, name='unassignAircraftOrder'),
]
