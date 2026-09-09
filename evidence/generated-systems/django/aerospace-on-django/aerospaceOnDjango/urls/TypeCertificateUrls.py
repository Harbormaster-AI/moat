from django.urls import path
from aerospaceOnDjango.views import TypeCertificateView

urlpatterns = [
    path('', TypeCertificateView.index, name='index'),
	path('create', TypeCertificateView.get, name='create'),
	path('get/<int:typeCertificateId>/', TypeCertificateView.get, name='get'),
	path('save', TypeCertificateView.save, name='save'),
	path('getAll', TypeCertificateView.getAll, name='getAll'),
	path('delete/<int:typeCertificateId>/', TypeCertificateView.delete, name='delete'),
	path('assignProgram/<int:typeCertificateId>/<int:ProgramId>/', TypeCertificateView.assignProgram, name='assignProgram'),
	path('unassignProgram/<int:typeCertificateId>/', TypeCertificateView.unassignProgram, name='unassignProgram'),
]
