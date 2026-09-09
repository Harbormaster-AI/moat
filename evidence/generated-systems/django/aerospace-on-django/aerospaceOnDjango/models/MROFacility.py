from django.db import models

#======================================================================
# 
# Encapsulates data for model MROFacility
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MROFacility Declaration
#======================================================================
class MROFacility (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	approvalScope = models.CharField(max_length=200, null=True)
	address = Address
	appointments = models.ManyToManyField('MaintenanceAppointment',  blank=True, related_name='+')
	workOrders = models.ManyToManyField('MaintenanceWorkOrder',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.approvalScope
		str = str + self.address
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MROFacility";
    
	def objectType(self):
		return "MROFacility";
