from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Collateral import Collateral
from fintechOnDjango.models.Loan import Loan
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Collateral
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CollateralDelegate Declaration
#======================================================================
class CollateralDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, collateralId ):
		try:	
			collateral = Collateral.objects.filter(id=collateralId)
			return collateral.first();
		except Collateral.DoesNotExist:
			raise ProcessingError("Collateral with id " + str(collateralId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, collateral):
		for model in serializers.deserialize("json", collateral):
			model.save()
			return model;

	def create(self, collateral):
		collateral.save()
		return collateral;

	def saveFromJson(self, collateral):
		for model in serializers.deserialize("json", collateral):
			model.save()
			return collateral;
	
	def save(self, collateral):
		collateral.save()
		return collateral;
	
	def delete(self, collateralId ):
		errMsg = "Failed to delete Collateral from db using id " + str(collateralId)
		
		try:
			collateral = Collateral.objects.get(id=collateralId)
			collateral.delete()
			return True
		except Collateral.DoesNotExist:
			raise ProcessingError("Collateral with id " + str(collateralId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Collateral.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Collateral from db")
		except Exception:
			return None;
		
	def assignLoan( self, collateralId, loanId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.LoanDelegate import LoanDelegate

		errMsg = "Failed to assign element " + str(loanId) + " for Loan on Collateral"

		try:
			# get the Collateral from db
			collateral = self.get( collateralId ).first()	
			
			# get the Loan from db
			loan = LoanDelegate().get(loanId).first();
			
			# assign the Loan		
			collateral.loan = loan
			
			#save it
			collateral.save()

			# reload and return the appropriate version					
			return self.get( collateralId );
		except Collateral.DoesNotExist:
			raise ProcessingError(errMsg + " : Collateral with id " + str(collateralId) + " does not exist.")
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLoan( self, collateralId ):
		errMsg = "Failed to unassign element " + str(loanId) + " for Loan on Collateral"

		try:
			# get the Collateral from db
			collateral = self.get( collateralId ).first()	
			
			# assign to None for unassignment
			collateral.loan = None			

			#save it
			collateral.save()

			# reload and return the appropriate version					
			return self.get( collateralId );
		except Collateral.DoesNotExist:
			raise ProcessingError(errMsg + " : Collateral with id " + str(collateralId) + " does not exist.")
		except Exception:
			return None;
		
