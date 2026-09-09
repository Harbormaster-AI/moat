from django.urls import path
from fintechOnDjango.views import AppliedFeeView

urlpatterns = [
    path('', AppliedFeeView.index, name='index'),
	path('create', AppliedFeeView.get, name='create'),
	path('get/<int:appliedFeeId>/', AppliedFeeView.get, name='get'),
	path('save', AppliedFeeView.save, name='save'),
	path('getAll', AppliedFeeView.getAll, name='getAll'),
	path('delete/<int:appliedFeeId>/', AppliedFeeView.delete, name='delete'),
	path('assignPaymentOrder/<int:appliedFeeId>/<int:PaymentOrderId>/', AppliedFeeView.assignPaymentOrder, name='assignPaymentOrder'),
	path('unassignPaymentOrder/<int:appliedFeeId>/', AppliedFeeView.unassignPaymentOrder, name='unassignPaymentOrder'),
	path('assignTransaction/<int:appliedFeeId>/<int:TransactionId>/', AppliedFeeView.assignTransaction, name='assignTransaction'),
	path('unassignTransaction/<int:appliedFeeId>/', AppliedFeeView.unassignTransaction, name='unassignTransaction'),
]
