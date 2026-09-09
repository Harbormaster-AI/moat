from django.db import models
from fintechOnDjango.models.InstallmentStatus import InstallmentStatus

#======================================================================
# 
# Encapsulates data for model RepaymentSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RepaymentSchedule Declaration
#======================================================================
class RepaymentSchedule (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	installmentNumber = models.IntegerField(null=True)
	dueDate = models.DateField(null=True)
	amountDue = Money
	principalDue = Money
	interestDue = Money
	loan = models.ForeignKey('Loan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	payments = models.ManyToManyField('Transaction',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InstallmentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.installmentNumber
		str = str + self.dueDate
		str = str + self.amountDue
		str = str + self.principalDue
		str = str + self.interestDue
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "RepaymentSchedule";
    
	def objectType(self):
		return "RepaymentSchedule";
