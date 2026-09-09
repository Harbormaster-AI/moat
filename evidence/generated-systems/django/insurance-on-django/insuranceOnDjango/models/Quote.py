from django.db import models

#======================================================================
# 
# Encapsulates data for model Quote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Quote Declaration
#======================================================================
class Quote (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	quoteNumber = models.CharField(max_length=200, null=True)
	totalPremium = Money
	ratingDate = models.DateField(null=True)
	asBound = models.BooleanField(null=True)
	application = models.ForeignKey('Application', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	underwritingDecisions = models.ManyToManyField('UnderwritingDecision',  blank=True, related_name='+')
	policy = models.OneToOneField('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quoteNumber
		str = str + self.totalPremium
		str = str + self.ratingDate
		str = str + self.asBound
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Quote";
    
	def objectType(self):
		return "Quote";
