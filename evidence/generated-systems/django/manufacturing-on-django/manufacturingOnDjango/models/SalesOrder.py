from django.db import models
from manufacturingOnDjango.models.SalesOrderStatus import SalesOrderStatus

#======================================================================
# 
# Encapsulates data for model SalesOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesOrder Declaration
#======================================================================
class SalesOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	orderNumber = models.CharField(max_length=200, null=True)
	orderDate = models.DateField(null=True)
	totalAmount = Money
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lines = models.ManyToManyField('SalesOrderLine',  blank=True, related_name='+')
	workOrders = models.ManyToManyField('WorkOrder',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SalesOrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.orderNumber
		str = str + self.orderDate
		str = str + self.totalAmount
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "SalesOrder";
    
	def objectType(self):
		return "SalesOrder";
