from django.urls import path
from ecommerceOnDjango.views import RefundView

urlpatterns = [
    path('', RefundView.index, name='index'),
	path('create', RefundView.get, name='create'),
	path('get/<int:refundId>/', RefundView.get, name='get'),
	path('save', RefundView.save, name='save'),
	path('getAll', RefundView.getAll, name='getAll'),
	path('delete/<int:refundId>/', RefundView.delete, name='delete'),
	path('assignPayment/<int:refundId>/<int:PaymentId>/', RefundView.assignPayment, name='assignPayment'),
	path('unassignPayment/<int:refundId>/', RefundView.unassignPayment, name='unassignPayment'),
	path('assignOrder/<int:refundId>/<int:OrderId>/', RefundView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:refundId>/', RefundView.unassignOrder, name='unassignOrder'),
]
