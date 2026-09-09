from django.urls import path
from fintechOnDjango.views import SecurityView

urlpatterns = [
    path('', SecurityView.index, name='index'),
	path('create', SecurityView.get, name='create'),
	path('get/<int:securityId>/', SecurityView.get, name='get'),
	path('save', SecurityView.save, name='save'),
	path('getAll', SecurityView.getAll, name='getAll'),
	path('delete/<int:securityId>/', SecurityView.delete, name='delete'),
	path('addPositions/<int:securityId>/<PositionsIds>/', SecurityView.addPositions, name='addPositions'),
	path('removePositions/<int:securityId>/<PositionsIds>/', SecurityView.removePositions, name='removePositions'),
	path('addTrades/<int:securityId>/<TradesIds>/', SecurityView.addTrades, name='addTrades'),
	path('removeTrades/<int:securityId>/<TradesIds>/', SecurityView.removeTrades, name='removeTrades'),
	path('addOrders/<int:securityId>/<OrdersIds>/', SecurityView.addOrders, name='addOrders'),
	path('removeOrders/<int:securityId>/<OrdersIds>/', SecurityView.removeOrders, name='removeOrders'),
]
