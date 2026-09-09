from django.urls import path
from fintechOnDjango.views import InvestmentPortfolioView

urlpatterns = [
    path('', InvestmentPortfolioView.index, name='index'),
	path('create', InvestmentPortfolioView.get, name='create'),
	path('get/<int:investmentPortfolioId>/', InvestmentPortfolioView.get, name='get'),
	path('save', InvestmentPortfolioView.save, name='save'),
	path('getAll', InvestmentPortfolioView.getAll, name='getAll'),
	path('delete/<int:investmentPortfolioId>/', InvestmentPortfolioView.delete, name='delete'),
	path('assignCustomer/<int:investmentPortfolioId>/<int:CustomerId>/', InvestmentPortfolioView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:investmentPortfolioId>/', InvestmentPortfolioView.unassignCustomer, name='unassignCustomer'),
	path('addAccounts/<int:investmentPortfolioId>/<AccountsIds>/', InvestmentPortfolioView.addAccounts, name='addAccounts'),
	path('removeAccounts/<int:investmentPortfolioId>/<AccountsIds>/', InvestmentPortfolioView.removeAccounts, name='removeAccounts'),
	path('addOrders/<int:investmentPortfolioId>/<OrdersIds>/', InvestmentPortfolioView.addOrders, name='addOrders'),
	path('removeOrders/<int:investmentPortfolioId>/<OrdersIds>/', InvestmentPortfolioView.removeOrders, name='removeOrders'),
	path('addHoldings/<int:investmentPortfolioId>/<HoldingsIds>/', InvestmentPortfolioView.addHoldings, name='addHoldings'),
	path('removeHoldings/<int:investmentPortfolioId>/<HoldingsIds>/', InvestmentPortfolioView.removeHoldings, name='removeHoldings'),
]
