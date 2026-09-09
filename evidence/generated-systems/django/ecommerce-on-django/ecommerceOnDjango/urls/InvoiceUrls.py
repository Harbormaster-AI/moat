from django.urls import path
from ecommerceOnDjango.views import InvoiceView

urlpatterns = [
    path('', InvoiceView.index, name='index'),
	path('create', InvoiceView.get, name='create'),
	path('get/<int:invoiceId>/', InvoiceView.get, name='get'),
	path('save', InvoiceView.save, name='save'),
	path('getAll', InvoiceView.getAll, name='getAll'),
	path('delete/<int:invoiceId>/', InvoiceView.delete, name='delete'),
	path('assignOrder/<int:invoiceId>/<int:OrderId>/', InvoiceView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:invoiceId>/', InvoiceView.unassignOrder, name='unassignOrder'),
]
