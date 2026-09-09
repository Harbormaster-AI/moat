from django.urls import path
from fintechOnDjango.views import PayoutView

urlpatterns = [
    path('', PayoutView.index, name='index'),
	path('create', PayoutView.get, name='create'),
	path('get/<int:payoutId>/', PayoutView.get, name='get'),
	path('save', PayoutView.save, name='save'),
	path('getAll', PayoutView.getAll, name='getAll'),
	path('delete/<int:payoutId>/', PayoutView.delete, name='delete'),
	path('assignMerchant/<int:payoutId>/<int:MerchantId>/', PayoutView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:payoutId>/', PayoutView.unassignMerchant, name='unassignMerchant'),
	path('assignSettlementBatch/<int:payoutId>/<int:SettlementBatchId>/', PayoutView.assignSettlementBatch, name='assignSettlementBatch'),
	path('unassignSettlementBatch/<int:payoutId>/', PayoutView.unassignSettlementBatch, name='unassignSettlementBatch'),
	path('assignDestinationAccount/<int:payoutId>/<int:DestinationAccountId>/', PayoutView.assignDestinationAccount, name='assignDestinationAccount'),
	path('unassignDestinationAccount/<int:payoutId>/', PayoutView.unassignDestinationAccount, name='unassignDestinationAccount'),
]
