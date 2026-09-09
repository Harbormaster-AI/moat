from django.urls import path
from crmOnDjango.views import PriceBookEntryView

urlpatterns = [
    path('', PriceBookEntryView.index, name='index'),
	path('create', PriceBookEntryView.get, name='create'),
	path('get/<int:priceBookEntryId>/', PriceBookEntryView.get, name='get'),
	path('save', PriceBookEntryView.save, name='save'),
	path('getAll', PriceBookEntryView.getAll, name='getAll'),
	path('delete/<int:priceBookEntryId>/', PriceBookEntryView.delete, name='delete'),
	path('assignPriceBook/<int:priceBookEntryId>/<int:PriceBookId>/', PriceBookEntryView.assignPriceBook, name='assignPriceBook'),
	path('unassignPriceBook/<int:priceBookEntryId>/', PriceBookEntryView.unassignPriceBook, name='unassignPriceBook'),
	path('assignProduct/<int:priceBookEntryId>/<int:ProductId>/', PriceBookEntryView.assignProduct, name='assignProduct'),
	path('unassignProduct/<int:priceBookEntryId>/', PriceBookEntryView.unassignProduct, name='unassignProduct'),
]
