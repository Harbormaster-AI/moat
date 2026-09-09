from django.urls import path
from ecommerceOnDjango.views import TaxRuleView

urlpatterns = [
    path('', TaxRuleView.index, name='index'),
	path('create', TaxRuleView.get, name='create'),
	path('get/<int:taxRuleId>/', TaxRuleView.get, name='get'),
	path('save', TaxRuleView.save, name='save'),
	path('getAll', TaxRuleView.getAll, name='getAll'),
	path('delete/<int:taxRuleId>/', TaxRuleView.delete, name='delete'),
	path('assignMerchant/<int:taxRuleId>/<int:MerchantId>/', TaxRuleView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:taxRuleId>/', TaxRuleView.unassignMerchant, name='unassignMerchant'),
	path('addChannels/<int:taxRuleId>/<ChannelsIds>/', TaxRuleView.addChannels, name='addChannels'),
	path('removeChannels/<int:taxRuleId>/<ChannelsIds>/', TaxRuleView.removeChannels, name='removeChannels'),
]
