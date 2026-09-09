from django.db import models

#======================================================================
# 
# Encapsulates data for model Aircraft
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Aircraft Declaration
#======================================================================
class Aircraft (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	msn = MSN
	deliveryDate = models.DateField(null=True)
	variant = models.ForeignKey('AircraftVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	operator = models.ForeignKey('Operator', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	registration = models.OneToOneField('Registration', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	warranty = models.OneToOneField('Warranty', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	maintenanceRecords = models.ManyToManyField('MaintenanceWorkOrder',  blank=True, related_name='+')
	connectedAircraft = models.OneToOneField('ConnectedAircraft', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	cabinLayout = models.ForeignKey('CabinLayout', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.msn
		str = str + self.deliveryDate
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Aircraft";
    
	def objectType(self):
		return "Aircraft";
