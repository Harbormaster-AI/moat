from django.db import models
from manufacturingOnDjango.models.SupplierTier import SupplierTier
from manufacturingOnDjango.models.PaymentTerms import PaymentTerms

#======================================================================
# 
# Encapsulates data for model Supplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Supplier Declaration
#======================================================================
class Supplier (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	supplierCode = models.CharField(max_length=200, null=True)
	address = Address
	enterprises = models.ManyToManyField('Enterprise',  blank=True, related_name='+')
	items = models.ManyToManyField('Item',  blank=True, related_name='+')
	purchaseOrders = models.ManyToManyField('PurchaseOrder',  blank=True, related_name='+')
	supplierTier = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SupplierTier])
	paymentTerms = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentTerms])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.supplierCode
		str = str + self.address
		str = str + self.supplierTier
		str = str + self.paymentTerms
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Supplier";
    
	def objectType(self):
		return "Supplier";
