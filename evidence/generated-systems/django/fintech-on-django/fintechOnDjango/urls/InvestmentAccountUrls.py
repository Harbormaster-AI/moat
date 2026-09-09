from django.urls import path
from fintechOnDjango.views import InvestmentAccountView

urlpatterns = [
    path('', InvestmentAccountView.index, name='index'),
	path('create', InvestmentAccountView.get, name='create'),
	path('get/<int:investmentAccountId>/', InvestmentAccountView.get, name='get'),
	path('save', InvestmentAccountView.save, name='save'),
	path('getAll', InvestmentAccountView.getAll, name='getAll'),
	path('delete/<int:investmentAccountId>/', InvestmentAccountView.delete, name='delete'),
	path('assignPortfolio/<int:investmentAccountId>/<int:PortfolioId>/', InvestmentAccountView.assignPortfolio, name='assignPortfolio'),
	path('unassignPortfolio/<int:investmentAccountId>/', InvestmentAccountView.unassignPortfolio, name='unassignPortfolio'),
	path('addTrades/<int:investmentAccountId>/<TradesIds>/', InvestmentAccountView.addTrades, name='addTrades'),
	path('removeTrades/<int:investmentAccountId>/<TradesIds>/', InvestmentAccountView.removeTrades, name='removeTrades'),
	path('addOrders/<int:investmentAccountId>/<OrdersIds>/', InvestmentAccountView.addOrders, name='addOrders'),
	path('removeOrders/<int:investmentAccountId>/<OrdersIds>/', InvestmentAccountView.removeOrders, name='removeOrders'),
]
