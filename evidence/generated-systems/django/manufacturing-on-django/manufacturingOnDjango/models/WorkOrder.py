from django.db import models
from manufacturingOnDjango.models.WorkOrderStatus import WorkOrderStatus

#======================================================================
# 
# Encapsulates data for model WorkOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkOrder Declaration
#======================================================================
class WorkOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	workOrderNumber = models.CharField(max_length=200, null=True)
	plannedStart = models.CharField(max_length=64, null=True)
	plannedEnd = models.CharField(max_length=64, null=True)
	quantity = Quantity
	priority = models.IntegerField(null=True)
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	routing = models.ForeignKey('Routing', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	bom = models.ForeignKey('BOM', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	productionSchedule = models.ForeignKey('ProductionSchedule', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	salesOrder = models.ForeignKey('SalesOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in WorkOrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.workOrderNumber
		str = str + self.plannedStart
		str = str + self.plannedEnd
		str = str + self.quantity
		str = str + self.priority
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "WorkOrder";
    
	def objectType(self):
		return "WorkOrder";
