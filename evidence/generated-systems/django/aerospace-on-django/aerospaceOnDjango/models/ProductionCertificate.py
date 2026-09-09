from django.db import models

#======================================================================
# 
# Encapsulates data for model ProductionCertificate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionCertificate Declaration
#======================================================================
class ProductionCertificate (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	certificateNumber = models.CharField(max_length=200, null=True)
	authority = models.CharField(max_length=200, null=True)
	manufacturer = models.ForeignKey('AerospaceManufacturer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.certificateNumber
		str = str + self.authority
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ProductionCertificate";
    
	def objectType(self):
		return "ProductionCertificate";
