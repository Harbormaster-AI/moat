from django.db import models
from governanceOnDjango.models.AttestationResult import AttestationResult

#======================================================================
# 
# Encapsulates data for model Attestation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Attestation Declaration
#======================================================================
class Attestation (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	statement = models.CharField(max_length=200, null=True)
	attestor = models.CharField(max_length=200, null=True)
	dateSigned = models.DateField(null=True)
	control = models.ForeignKey('Control', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	complianceProgram = models.ForeignKey('ComplianceProgram', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	result = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AttestationResult])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.statement
		str = str + self.attestor
		str = str + self.dateSigned
		str = str + self.result
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Attestation";
    
	def objectType(self):
		return "Attestation";
