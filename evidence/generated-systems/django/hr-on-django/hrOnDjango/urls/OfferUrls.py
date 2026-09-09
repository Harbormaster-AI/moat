from django.urls import path
from hrOnDjango.views import OfferView

urlpatterns = [
    path('', OfferView.index, name='index'),
	path('create', OfferView.get, name='create'),
	path('get/<int:offerId>/', OfferView.get, name='get'),
	path('save', OfferView.save, name='save'),
	path('getAll', OfferView.getAll, name='getAll'),
	path('delete/<int:offerId>/', OfferView.delete, name='delete'),
	path('assignRequisition/<int:offerId>/<int:RequisitionId>/', OfferView.assignRequisition, name='assignRequisition'),
	path('unassignRequisition/<int:offerId>/', OfferView.unassignRequisition, name='unassignRequisition'),
	path('assignCandidate/<int:offerId>/<int:CandidateId>/', OfferView.assignCandidate, name='assignCandidate'),
	path('unassignCandidate/<int:offerId>/', OfferView.unassignCandidate, name='unassignCandidate'),
	path('assignApprovedBy/<int:offerId>/<int:ApprovedById>/', OfferView.assignApprovedBy, name='assignApprovedBy'),
	path('unassignApprovedBy/<int:offerId>/', OfferView.unassignApprovedBy, name='unassignApprovedBy'),
	path('assignContract/<int:offerId>/<int:ContractId>/', OfferView.assignContract, name='assignContract'),
	path('unassignContract/<int:offerId>/', OfferView.unassignContract, name='unassignContract'),
]
