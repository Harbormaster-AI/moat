from django.urls import path
from fintechOnDjango.views import DisputeView

urlpatterns = [
    path('', DisputeView.index, name='index'),
	path('create', DisputeView.get, name='create'),
	path('get/<int:disputeId>/', DisputeView.get, name='get'),
	path('save', DisputeView.save, name='save'),
	path('getAll', DisputeView.getAll, name='getAll'),
	path('delete/<int:disputeId>/', DisputeView.delete, name='delete'),
	path('assignTransaction/<int:disputeId>/<int:TransactionId>/', DisputeView.assignTransaction, name='assignTransaction'),
	path('unassignTransaction/<int:disputeId>/', DisputeView.unassignTransaction, name='unassignTransaction'),
	path('assignCard/<int:disputeId>/<int:CardId>/', DisputeView.assignCard, name='assignCard'),
	path('unassignCard/<int:disputeId>/', DisputeView.unassignCard, name='unassignCard'),
	path('assignMerchant/<int:disputeId>/<int:MerchantId>/', DisputeView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:disputeId>/', DisputeView.unassignMerchant, name='unassignMerchant'),
	path('addChargebacks/<int:disputeId>/<ChargebacksIds>/', DisputeView.addChargebacks, name='addChargebacks'),
	path('removeChargebacks/<int:disputeId>/<ChargebacksIds>/', DisputeView.removeChargebacks, name='removeChargebacks'),
]
