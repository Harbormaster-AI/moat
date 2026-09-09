from django.db import models
from hrOnDjango.models.OfferStatus import OfferStatus

#======================================================================
# 
# Encapsulates data for model Offer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Offer Declaration
#======================================================================
class Offer (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	offerNumber = models.CharField(max_length=200, null=True)
	proposedStartDate = models.DateField(null=True)
	baseSalary = Money
	signOnBonus = Money
	requisition = models.ForeignKey('JobRequisition', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	candidate = models.ForeignKey('Candidate', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	approvedBy = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	contract = models.OneToOneField('EmploymentContract', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OfferStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.offerNumber
		str = str + self.proposedStartDate
		str = str + self.baseSalary
		str = str + self.signOnBonus
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Offer";
    
	def objectType(self):
		return "Offer";
