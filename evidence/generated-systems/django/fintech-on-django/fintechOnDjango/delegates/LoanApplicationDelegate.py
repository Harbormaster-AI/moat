from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.LoanApplication import LoanApplication
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.RiskAssessment import RiskAssessment
from fintechOnDjango.models.Loan import Loan
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model LoanApplication
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanApplicationDelegate Declaration
#======================================================================
class LoanApplicationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, loanApplicationId ):
		try:	
			loanApplication = LoanApplication.objects.filter(id=loanApplicationId)
			return loanApplication.first();
		except LoanApplication.DoesNotExist:
			raise ProcessingError("LoanApplication with id " + str(loanApplicationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, loanApplication):
		for model in serializers.deserialize("json", loanApplication):
			model.save()
			return model;

	def create(self, loanApplication):
		loanApplication.save()
		return loanApplication;

	def saveFromJson(self, loanApplication):
		for model in serializers.deserialize("json", loanApplication):
			model.save()
			return loanApplication;
	
	def save(self, loanApplication):
		loanApplication.save()
		return loanApplication;
	
	def delete(self, loanApplicationId ):
		errMsg = "Failed to delete LoanApplication from db using id " + str(loanApplicationId)
		
		try:
			loanApplication = LoanApplication.objects.get(id=loanApplicationId)
			loanApplication.delete()
			return True
		except LoanApplication.DoesNotExist:
			raise ProcessingError("LoanApplication with id " + str(loanApplicationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = LoanApplication.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all LoanApplication from db")
		except Exception:
			return None;
		
	def assignCustomer( self, loanApplicationId, customerId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on LoanApplication"

		try:
			# get the LoanApplication from db
			loanApplication = self.get( loanApplicationId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			loanApplication.customer = customer
			
			#save it
			loanApplication.save()

			# reload and return the appropriate version					
			return self.get( loanApplicationId );
		except LoanApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanApplication with id " + str(loanApplicationId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, loanApplicationId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on LoanApplication"

		try:
			# get the LoanApplication from db
			loanApplication = self.get( loanApplicationId ).first()	
			
			# assign to None for unassignment
			loanApplication.customer = None			

			#save it
			loanApplication.save()

			# reload and return the appropriate version					
			return self.get( loanApplicationId );
		except LoanApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanApplication with id " + str(loanApplicationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRiskAssessment( self, loanApplicationId, riskAssessmentId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.RiskAssessmentDelegate import RiskAssessmentDelegate

		errMsg = "Failed to assign element " + str(riskAssessmentId) + " for RiskAssessment on LoanApplication"

		try:
			# get the LoanApplication from db
			loanApplication = self.get( loanApplicationId ).first()	
			
			# get the RiskAssessment from db
			riskAssessment = RiskAssessmentDelegate().get(riskAssessmentId).first();
			
			# assign the RiskAssessment		
			loanApplication.riskAssessment = riskAssessment
			
			#save it
			loanApplication.save()

			# reload and return the appropriate version					
			return self.get( loanApplicationId );
		except LoanApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanApplication with id " + str(loanApplicationId) + " does not exist.")
		except RiskAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : RiskAssessment with id " + str(riskAssessmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRiskAssessment( self, loanApplicationId ):
		errMsg = "Failed to unassign element " + str(riskAssessmentId) + " for RiskAssessment on LoanApplication"

		try:
			# get the LoanApplication from db
			loanApplication = self.get( loanApplicationId ).first()	
			
			# assign to None for unassignment
			loanApplication.riskAssessment = None			

			#save it
			loanApplication.save()

			# reload and return the appropriate version					
			return self.get( loanApplicationId );
		except LoanApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanApplication with id " + str(loanApplicationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLoan( self, loanApplicationId, loanId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.LoanDelegate import LoanDelegate

		errMsg = "Failed to assign element " + str(loanId) + " for Loan on LoanApplication"

		try:
			# get the LoanApplication from db
			loanApplication = self.get( loanApplicationId ).first()	
			
			# get the Loan from db
			loan = LoanDelegate().get(loanId).first();
			
			# assign the Loan		
			loanApplication.loan = loan
			
			#save it
			loanApplication.save()

			# reload and return the appropriate version					
			return self.get( loanApplicationId );
		except LoanApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanApplication with id " + str(loanApplicationId) + " does not exist.")
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLoan( self, loanApplicationId ):
		errMsg = "Failed to unassign element " + str(loanId) + " for Loan on LoanApplication"

		try:
			# get the LoanApplication from db
			loanApplication = self.get( loanApplicationId ).first()	
			
			# assign to None for unassignment
			loanApplication.loan = None			

			#save it
			loanApplication.save()

			# reload and return the appropriate version					
			return self.get( loanApplicationId );
		except LoanApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanApplication with id " + str(loanApplicationId) + " does not exist.")
		except Exception:
			return None;
		
