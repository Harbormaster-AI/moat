from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Branch import Branch
from fintechOnDjango.models.FinancialInstitution import FinancialInstitution
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Branch
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BranchDelegate Declaration
#======================================================================
class BranchDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, branchId ):
		try:	
			branch = Branch.objects.filter(id=branchId)
			return branch.first();
		except Branch.DoesNotExist:
			raise ProcessingError("Branch with id " + str(branchId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, branch):
		for model in serializers.deserialize("json", branch):
			model.save()
			return model;

	def create(self, branch):
		branch.save()
		return branch;

	def saveFromJson(self, branch):
		for model in serializers.deserialize("json", branch):
			model.save()
			return branch;
	
	def save(self, branch):
		branch.save()
		return branch;
	
	def delete(self, branchId ):
		errMsg = "Failed to delete Branch from db using id " + str(branchId)
		
		try:
			branch = Branch.objects.get(id=branchId)
			branch.delete()
			return True
		except Branch.DoesNotExist:
			raise ProcessingError("Branch with id " + str(branchId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Branch.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Branch from db")
		except Exception:
			return None;
		
	def assignInstitution( self, branchId, institutionId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FinancialInstitutionDelegate import FinancialInstitutionDelegate

		errMsg = "Failed to assign element " + str(institutionId) + " for Institution on Branch"

		try:
			# get the Branch from db
			branch = self.get( branchId ).first()	
			
			# get the FinancialInstitution from db
			financialInstitution = FinancialInstitutionDelegate().get(institutionId).first();
			
			# assign the Institution		
			branch.institution = financialInstitution
			
			#save it
			branch.save()

			# reload and return the appropriate version					
			return self.get( branchId );
		except Branch.DoesNotExist:
			raise ProcessingError(errMsg + " : Branch with id " + str(branchId) + " does not exist.")
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(institutionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInstitution( self, branchId ):
		errMsg = "Failed to unassign element " + str(institutionId) + " for Institution on Branch"

		try:
			# get the Branch from db
			branch = self.get( branchId ).first()	
			
			# assign to None for unassignment
			branch.financialInstitution = None			

			#save it
			branch.save()

			# reload and return the appropriate version					
			return self.get( branchId );
		except Branch.DoesNotExist:
			raise ProcessingError(errMsg + " : Branch with id " + str(branchId) + " does not exist.")
		except Exception:
			return None;
		
