from django.urls import path
from advertisingOnDjango.views import PaymentMethodView

urlpatterns = [
    path('', PaymentMethodView.index, name='index'),
	path('create', PaymentMethodView.get, name='create'),
	path('get/<int:paymentMethodId>/', PaymentMethodView.get, name='get'),
	path('save', PaymentMethodView.save, name='save'),
	path('getAll', PaymentMethodView.getAll, name='getAll'),
	path('delete/<int:paymentMethodId>/', PaymentMethodView.delete, name='delete'),
	path('assignBillingProfile/<int:paymentMethodId>/<int:BillingProfileId>/', PaymentMethodView.assignBillingProfile, name='assignBillingProfile'),
	path('unassignBillingProfile/<int:paymentMethodId>/', PaymentMethodView.unassignBillingProfile, name='unassignBillingProfile'),
]
