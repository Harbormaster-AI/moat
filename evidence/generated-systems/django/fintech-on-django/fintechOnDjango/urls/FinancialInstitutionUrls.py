from django.urls import path
from fintechOnDjango.views import FinancialInstitutionView

urlpatterns = [
    path('', FinancialInstitutionView.index, name='index'),
	path('create', FinancialInstitutionView.get, name='create'),
	path('get/<int:financialInstitutionId>/', FinancialInstitutionView.get, name='get'),
	path('save', FinancialInstitutionView.save, name='save'),
	path('getAll', FinancialInstitutionView.getAll, name='getAll'),
	path('delete/<int:financialInstitutionId>/', FinancialInstitutionView.delete, name='delete'),
	path('addBranches/<int:financialInstitutionId>/<BranchesIds>/', FinancialInstitutionView.addBranches, name='addBranches'),
	path('removeBranches/<int:financialInstitutionId>/<BranchesIds>/', FinancialInstitutionView.removeBranches, name='removeBranches'),
	path('addCustomers/<int:financialInstitutionId>/<CustomersIds>/', FinancialInstitutionView.addCustomers, name='addCustomers'),
	path('removeCustomers/<int:financialInstitutionId>/<CustomersIds>/', FinancialInstitutionView.removeCustomers, name='removeCustomers'),
	path('addProductOfferings/<int:financialInstitutionId>/<ProductOfferingsIds>/', FinancialInstitutionView.addProductOfferings, name='addProductOfferings'),
	path('removeProductOfferings/<int:financialInstitutionId>/<ProductOfferingsIds>/', FinancialInstitutionView.removeProductOfferings, name='removeProductOfferings'),
	path('addPaymentProcessors/<int:financialInstitutionId>/<PaymentProcessorsIds>/', FinancialInstitutionView.addPaymentProcessors, name='addPaymentProcessors'),
	path('removePaymentProcessors/<int:financialInstitutionId>/<PaymentProcessorsIds>/', FinancialInstitutionView.removePaymentProcessors, name='removePaymentProcessors'),
	path('addCompliancePolicies/<int:financialInstitutionId>/<CompliancePoliciesIds>/', FinancialInstitutionView.addCompliancePolicies, name='addCompliancePolicies'),
	path('removeCompliancePolicies/<int:financialInstitutionId>/<CompliancePoliciesIds>/', FinancialInstitutionView.removeCompliancePolicies, name='removeCompliancePolicies'),
]
