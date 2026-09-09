from django.db import models
from manufacturingOnDjango.models.PlannedOrderType import PlannedOrderType
from manufacturingOnDjango.models.PlannedOrderStatus import PlannedOrderStatus

#======================================================================
# 
# Encapsulates data for model PlannedOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlannedOrder Declaration
#======================================================================
class PlannedOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	plannedOrderNumber = models.CharField(max_length=200, null=True)
	quantity = Quantity
	dueDate = models.DateField(null=True)
	mrpRun = models.ForeignKey('MRPRun', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	orderType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PlannedOrderType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PlannedOrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.plannedOrderNumber
		str = str + self.quantity
		str = str + self.dueDate
		str = str + self.orderType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PlannedOrder";
    
	def objectType(self):
		return "PlannedOrder";
