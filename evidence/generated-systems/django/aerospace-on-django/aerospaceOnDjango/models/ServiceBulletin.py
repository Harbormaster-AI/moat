from django.db import models
from aerospaceOnDjango.models.ServiceBulletinCategory import ServiceBulletinCategory

#======================================================================
# 
# Encapsulates data for model ServiceBulletin
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceBulletin Declaration
#======================================================================
class ServiceBulletin (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	bulletinNumber = models.CharField(max_length=200, null=True)
	workOrders = models.ManyToManyField('MaintenanceWorkOrder',  blank=True, related_name='+')
	variants = models.ManyToManyField('AircraftVariant',  blank=True, related_name='+')
	category = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ServiceBulletinCategory])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.bulletinNumber
		str = str + self.category
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ServiceBulletin";
    
	def objectType(self):
		return "ServiceBulletin";
