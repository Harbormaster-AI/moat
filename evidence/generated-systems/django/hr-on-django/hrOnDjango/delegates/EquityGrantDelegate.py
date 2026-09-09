from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.EquityGrant import EquityGrant
from hrOnDjango.models.CompensationPackage import CompensationPackage
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model EquityGrant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EquityGrantDelegate Declaration
#======================================================================
class EquityGrantDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, equityGrantId ):
		try:	
			equityGrant = EquityGrant.objects.filter(id=equityGrantId)
			return equityGrant.first();
		except EquityGrant.DoesNotExist:
			raise ProcessingError("EquityGrant with id " + str(equityGrantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, equityGrant):
		for model in serializers.deserialize("json", equityGrant):
			model.save()
			return model;

	def create(self, equityGrant):
		equityGrant.save()
		return equityGrant;

	def saveFromJson(self, equityGrant):
		for model in serializers.deserialize("json", equityGrant):
			model.save()
			return equityGrant;
	
	def save(self, equityGrant):
		equityGrant.save()
		return equityGrant;
	
	def delete(self, equityGrantId ):
		errMsg = "Failed to delete EquityGrant from db using id " + str(equityGrantId)
		
		try:
			equityGrant = EquityGrant.objects.get(id=equityGrantId)
			equityGrant.delete()
			return True
		except EquityGrant.DoesNotExist:
			raise ProcessingError("EquityGrant with id " + str(equityGrantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = EquityGrant.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all EquityGrant from db")
		except Exception:
			return None;
		
	def assignCompensationPackage( self, equityGrantId, compensationPackageId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompensationPackageDelegate import CompensationPackageDelegate

		errMsg = "Failed to assign element " + str(compensationPackageId) + " for CompensationPackage on EquityGrant"

		try:
			# get the EquityGrant from db
			equityGrant = self.get( equityGrantId ).first()	
			
			# get the CompensationPackage from db
			compensationPackage = CompensationPackageDelegate().get(compensationPackageId).first();
			
			# assign the CompensationPackage		
			equityGrant.compensationPackage = compensationPackage
			
			#save it
			equityGrant.save()

			# reload and return the appropriate version					
			return self.get( equityGrantId );
		except EquityGrant.DoesNotExist:
			raise ProcessingError(errMsg + " : EquityGrant with id " + str(equityGrantId) + " does not exist.")
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCompensationPackage( self, equityGrantId ):
		errMsg = "Failed to unassign element " + str(compensationPackageId) + " for CompensationPackage on EquityGrant"

		try:
			# get the EquityGrant from db
			equityGrant = self.get( equityGrantId ).first()	
			
			# assign to None for unassignment
			equityGrant.compensationPackage = None			

			#save it
			equityGrant.save()

			# reload and return the appropriate version					
			return self.get( equityGrantId );
		except EquityGrant.DoesNotExist:
			raise ProcessingError(errMsg + " : EquityGrant with id " + str(equityGrantId) + " does not exist.")
		except Exception:
			return None;
		
