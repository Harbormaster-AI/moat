from django.urls import path
from fintechOnDjango.views import TerminalView

urlpatterns = [
    path('', TerminalView.index, name='index'),
	path('create', TerminalView.get, name='create'),
	path('get/<int:terminalId>/', TerminalView.get, name='get'),
	path('save', TerminalView.save, name='save'),
	path('getAll', TerminalView.getAll, name='getAll'),
	path('delete/<int:terminalId>/', TerminalView.delete, name='delete'),
	path('assignMerchant/<int:terminalId>/<int:MerchantId>/', TerminalView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:terminalId>/', TerminalView.unassignMerchant, name='unassignMerchant'),
]
