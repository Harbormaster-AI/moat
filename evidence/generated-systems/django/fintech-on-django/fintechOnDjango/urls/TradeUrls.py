from django.urls import path
from fintechOnDjango.views import TradeView

urlpatterns = [
    path('', TradeView.index, name='index'),
	path('create', TradeView.get, name='create'),
	path('get/<int:tradeId>/', TradeView.get, name='get'),
	path('save', TradeView.save, name='save'),
	path('getAll', TradeView.getAll, name='getAll'),
	path('delete/<int:tradeId>/', TradeView.delete, name='delete'),
	path('assignOrder/<int:tradeId>/<int:OrderId>/', TradeView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:tradeId>/', TradeView.unassignOrder, name='unassignOrder'),
	path('assignSecurity/<int:tradeId>/<int:SecurityId>/', TradeView.assignSecurity, name='assignSecurity'),
	path('unassignSecurity/<int:tradeId>/', TradeView.unassignSecurity, name='unassignSecurity'),
	path('assignInvestmentAccount/<int:tradeId>/<int:InvestmentAccountId>/', TradeView.assignInvestmentAccount, name='assignInvestmentAccount'),
	path('unassignInvestmentAccount/<int:tradeId>/', TradeView.unassignInvestmentAccount, name='unassignInvestmentAccount'),
]
