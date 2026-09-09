from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.CompensationPackage import CompensationPackage
from hrOnDjango.models.EmploymentContract import EmploymentContract
from hrOnDjango.models.SalaryComponent import SalaryComponent
from hrOnDjango.models.BonusPlan import BonusPlan
from hrOnDjango.models.EquityGrant import EquityGrant
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CompensationPackage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompensationPackageDelegate Declaration
#======================================================================
class CompensationPackageDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, compensationPackageId ):
		try:	
			compensationPackage = CompensationPackage.objects.filter(id=compensationPackageId)
			return compensationPackage.first();
		except CompensationPackage.DoesNotExist:
			raise ProcessingError("CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, compensationPackage):
		for model in serializers.deserialize("json", compensationPackage):
			model.save()
			return model;

	def create(self, compensationPackage):
		compensationPackage.save()
		return compensationPackage;

	def saveFromJson(self, compensationPackage):
		for model in serializers.deserialize("json", compensationPackage):
			model.save()
			return compensationPackage;
	
	def save(self, compensationPackage):
		compensationPackage.save()
		return compensationPackage;
	
	def delete(self, compensationPackageId ):
		errMsg = "Failed to delete CompensationPackage from db using id " + str(compensationPackageId)
		
		try:
			compensationPackage = CompensationPackage.objects.get(id=compensationPackageId)
			compensationPackage.delete()
			return True
		except CompensationPackage.DoesNotExist:
			raise ProcessingError("CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CompensationPackage.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CompensationPackage from db")
		except Exception:
			return None;
		
	def assignContract( self, compensationPackageId, contractId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentContractDelegate import EmploymentContractDelegate

		errMsg = "Failed to assign element " + str(contractId) + " for Contract on CompensationPackage"

		try:
			# get the CompensationPackage from db
			compensationPackage = self.get( compensationPackageId ).first()	
			
			# get the EmploymentContract from db
			employmentContract = EmploymentContractDelegate().get(contractId).first();
			
			# assign the Contract		
			compensationPackage.contract = employmentContract
			
			#save it
			compensationPackage.save()

			# reload and return the appropriate version					
			return self.get( compensationPackageId );
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(contractId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignContract( self, compensationPackageId ):
		errMsg = "Failed to unassign element " + str(contractId) + " for Contract on CompensationPackage"

		try:
			# get the CompensationPackage from db
			compensationPackage = self.get( compensationPackageId ).first()	
			
			# assign to None for unassignment
			compensationPackage.employmentContract = None			

			#save it
			compensationPackage.save()

			# reload and return the appropriate version					
			return self.get( compensationPackageId );
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except Exception:
			return None;
		
	def addSalaryComponents( self, compensationPackageId, salaryComponentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.SalaryComponentDelegate import SalaryComponentDelegate

		errMsg = "Failed to add elements " + str(salaryComponentsIds) + " for SalaryComponents on CompensationPackage"

		try:
			# get the CompensationPackage
			compensationPackage = self.get( compensationPackageId ).first()
				
			# split on a comma with no spaces
			idList = salaryComponentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SalaryComponent		
				salaryComponent = SalaryComponentDelegate().get(id).first();	
				# add the SalaryComponent
				compensationPackage.salaryComponents.add(salaryComponent)
				
			# save it		
			compensationPackage.save()
			
			# reload and return the appropriate version
			return self.get( compensationPackageId );
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except SalaryComponent.DoesNotExist:
			raise ProcessingError(errMsg + " : SalaryComponent does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSalaryComponents( self, compensationPackageId, salaryComponentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.SalaryComponentDelegate import SalaryComponentDelegate

		errMsg = "Failed to remove elements " + str(salaryComponentsIds) + " for SalaryComponents on CompensationPackage"

		try:
			# get the CompensationPackage
			compensationPackage = self.get( compensationPackageId ).first()
				
			# split on a comma with no spaces
			idList = salaryComponentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SalaryComponent		
				salaryComponent = SalaryComponentDelegate().get(id).first();	
				# add the SalaryComponent
				compensationPackage.salaryComponents.remove(salaryComponent)
				
			# save it		
			compensationPackage.save()
			
			# reload and return the appropriate version
			return self.get( compensationPackageId );
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except SalaryComponent.DoesNotExist:
			raise ProcessingError(errMsg + " : SalaryComponent does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addBonusPlans( self, compensationPackageId, bonusPlansIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BonusPlanDelegate import BonusPlanDelegate

		errMsg = "Failed to add elements " + str(bonusPlansIds) + " for BonusPlans on CompensationPackage"

		try:
			# get the CompensationPackage
			compensationPackage = self.get( compensationPackageId ).first()
				
			# split on a comma with no spaces
			idList = bonusPlansIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BonusPlan		
				bonusPlan = BonusPlanDelegate().get(id).first();	
				# add the BonusPlan
				compensationPackage.bonusPlans.add(bonusPlan)
				
			# save it		
			compensationPackage.save()
			
			# reload and return the appropriate version
			return self.get( compensationPackageId );
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except BonusPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BonusPlan does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBonusPlans( self, compensationPackageId, bonusPlansIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BonusPlanDelegate import BonusPlanDelegate

		errMsg = "Failed to remove elements " + str(bonusPlansIds) + " for BonusPlans on CompensationPackage"

		try:
			# get the CompensationPackage
			compensationPackage = self.get( compensationPackageId ).first()
				
			# split on a comma with no spaces
			idList = bonusPlansIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BonusPlan		
				bonusPlan = BonusPlanDelegate().get(id).first();	
				# add the BonusPlan
				compensationPackage.bonusPlans.remove(bonusPlan)
				
			# save it		
			compensationPackage.save()
			
			# reload and return the appropriate version
			return self.get( compensationPackageId );
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except BonusPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BonusPlan does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEquityGrants( self, compensationPackageId, equityGrantsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EquityGrantDelegate import EquityGrantDelegate

		errMsg = "Failed to add elements " + str(equityGrantsIds) + " for EquityGrants on CompensationPackage"

		try:
			# get the CompensationPackage
			compensationPackage = self.get( compensationPackageId ).first()
				
			# split on a comma with no spaces
			idList = equityGrantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EquityGrant		
				equityGrant = EquityGrantDelegate().get(id).first();	
				# add the EquityGrant
				compensationPackage.equityGrants.add(equityGrant)
				
			# save it		
			compensationPackage.save()
			
			# reload and return the appropriate version
			return self.get( compensationPackageId );
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except EquityGrant.DoesNotExist:
			raise ProcessingError(errMsg + " : EquityGrant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEquityGrants( self, compensationPackageId, equityGrantsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EquityGrantDelegate import EquityGrantDelegate

		errMsg = "Failed to remove elements " + str(equityGrantsIds) + " for EquityGrants on CompensationPackage"

		try:
			# get the CompensationPackage
			compensationPackage = self.get( compensationPackageId ).first()
				
			# split on a comma with no spaces
			idList = equityGrantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EquityGrant		
				equityGrant = EquityGrantDelegate().get(id).first();	
				# add the EquityGrant
				compensationPackage.equityGrants.remove(equityGrant)
				
			# save it		
			compensationPackage.save()
			
			# reload and return the appropriate version
			return self.get( compensationPackageId );
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except EquityGrant.DoesNotExist:
			raise ProcessingError(errMsg + " : EquityGrant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
