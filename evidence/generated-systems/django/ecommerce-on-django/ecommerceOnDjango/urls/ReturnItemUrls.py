from django.urls import path
from ecommerceOnDjango.views import ReturnItemView

urlpatterns = [
    path('', ReturnItemView.index, name='index'),
	path('create', ReturnItemView.get, name='create'),
	path('get/<int:returnItemId>/', ReturnItemView.get, name='get'),
	path('save', ReturnItemView.save, name='save'),
	path('getAll', ReturnItemView.getAll, name='getAll'),
	path('delete/<int:returnItemId>/', ReturnItemView.delete, name='delete'),
	path('assignReturnRequest/<int:returnItemId>/<int:ReturnRequestId>/', ReturnItemView.assignReturnRequest, name='assignReturnRequest'),
	path('unassignReturnRequest/<int:returnItemId>/', ReturnItemView.unassignReturnRequest, name='unassignReturnRequest'),
	path('assignOrderLine/<int:returnItemId>/<int:OrderLineId>/', ReturnItemView.assignOrderLine, name='assignOrderLine'),
	path('unassignOrderLine/<int:returnItemId>/', ReturnItemView.unassignOrderLine, name='unassignOrderLine'),
]
