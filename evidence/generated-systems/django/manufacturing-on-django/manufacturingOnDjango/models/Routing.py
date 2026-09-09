from django.db import models
from manufacturingOnDjango.models.RoutingType import RoutingType
from manufacturingOnDjango.models.RoutingStatus import RoutingStatus

#======================================================================
# 
# Encapsulates data for model Routing
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Routing Declaration
#======================================================================
class Routing (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	routingNumber = models.CharField(max_length=200, null=True)
	revision = models.CharField(max_length=200, null=True)
	effectivityStart = models.DateField(null=True)
	effectivityEnd = models.DateField(null=True)
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	operations = models.ManyToManyField('Operation',  blank=True, related_name='+')
	routingType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RoutingType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RoutingStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.routingNumber
		str = str + self.revision
		str = str + self.effectivityStart
		str = str + self.effectivityEnd
		str = str + self.routingType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Routing";
    
	def objectType(self):
		return "Routing";
