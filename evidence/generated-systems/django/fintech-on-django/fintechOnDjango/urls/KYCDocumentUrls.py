from django.urls import path
from fintechOnDjango.views import KYCDocumentView

urlpatterns = [
    path('', KYCDocumentView.index, name='index'),
	path('create', KYCDocumentView.get, name='create'),
	path('get/<int:kYCDocumentId>/', KYCDocumentView.get, name='get'),
	path('save', KYCDocumentView.save, name='save'),
	path('getAll', KYCDocumentView.getAll, name='getAll'),
	path('delete/<int:kYCDocumentId>/', KYCDocumentView.delete, name='delete'),
	path('assignKycProfile/<int:kYCDocumentId>/<int:KycProfileId>/', KYCDocumentView.assignKycProfile, name='assignKycProfile'),
	path('unassignKycProfile/<int:kYCDocumentId>/', KYCDocumentView.unassignKycProfile, name='unassignKycProfile'),
]
