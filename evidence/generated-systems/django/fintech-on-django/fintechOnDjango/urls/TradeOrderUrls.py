from django.urls import path
from fintechOnDjango.views import TradeOrderView

urlpatterns = [
    path('', TradeOrderView.index, name='index'),
	path('create', TradeOrderView.get, name='create'),
	path('get/<int:tradeOrderId>/', TradeOrderView.get, name='get'),
	path('save', TradeOrderView.save, name='save'),
	path('getAll', TradeOrderView.getAll, name='getAll'),
	path('delete/<int:tradeOrderId>/', TradeOrderView.delete, name='delete'),
	path('assignPortfolio/<int:tradeOrderId>/<int:PortfolioId>/', TradeOrderView.assignPortfolio, name='assignPortfolio'),
	path('unassignPortfolio/<int:tradeOrderId>/', TradeOrderView.unassignPortfolio, name='unassignPortfolio'),
	path('assignSecurity/<int:tradeOrderId>/<int:SecurityId>/', TradeOrderView.assignSecurity, name='assignSecurity'),
	path('unassignSecurity/<int:tradeOrderId>/', TradeOrderView.unassignSecurity, name='unassignSecurity'),
	path('addTrades/<int:tradeOrderId>/<TradesIds>/', TradeOrderView.addTrades, name='addTrades'),
	path('removeTrades/<int:tradeOrderId>/<TradesIds>/', TradeOrderView.removeTrades, name='removeTrades'),
]
