from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Underwriter import Underwriter
from insuranceOnDjango.models.UnderwritingDecision import UnderwritingDecision
from insuranceOnDjango.models.Insurer import Insurer
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Underwriter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UnderwriterDelegate Declaration
#======================================================================
class UnderwriterDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, underwriterId ):
		try:	
			underwriter = Underwriter.objects.filter(id=underwriterId)
			return underwriter.first();
		except Underwriter.DoesNotExist:
			raise ProcessingError("Underwriter with id " + str(underwriterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, underwriter):
		for model in serializers.deserialize("json", underwriter):
			model.save()
			return model;

	def create(self, underwriter):
		underwriter.save()
		return underwriter;

	def saveFromJson(self, underwriter):
		for model in serializers.deserialize("json", underwriter):
			model.save()
			return underwriter;
	
	def save(self, underwriter):
		underwriter.save()
		return underwriter;
	
	def delete(self, underwriterId ):
		errMsg = "Failed to delete Underwriter from db using id " + str(underwriterId)
		
		try:
			underwriter = Underwriter.objects.get(id=underwriterId)
			underwriter.delete()
			return True
		except Underwriter.DoesNotExist:
			raise ProcessingError("Underwriter with id " + str(underwriterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Underwriter.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Underwriter from db")
		except Exception:
			return None;
		
	def assignInsurer( self, underwriterId, insurerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsurerDelegate import InsurerDelegate

		errMsg = "Failed to assign element " + str(insurerId) + " for Insurer on Underwriter"

		try:
			# get the Underwriter from db
			underwriter = self.get( underwriterId ).first()	
			
			# get the Insurer from db
			insurer = InsurerDelegate().get(insurerId).first();
			
			# assign the Insurer		
			underwriter.insurer = insurer
			
			#save it
			underwriter.save()

			# reload and return the appropriate version					
			return self.get( underwriterId );
		except Underwriter.DoesNotExist:
			raise ProcessingError(errMsg + " : Underwriter with id " + str(underwriterId) + " does not exist.")
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInsurer( self, underwriterId ):
		errMsg = "Failed to unassign element " + str(insurerId) + " for Insurer on Underwriter"

		try:
			# get the Underwriter from db
			underwriter = self.get( underwriterId ).first()	
			
			# assign to None for unassignment
			underwriter.insurer = None			

			#save it
			underwriter.save()

			# reload and return the appropriate version					
			return self.get( underwriterId );
		except Underwriter.DoesNotExist:
			raise ProcessingError(errMsg + " : Underwriter with id " + str(underwriterId) + " does not exist.")
		except Exception:
			return None;
		
	def addDecisions( self, underwriterId, decisionsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.UnderwritingDecisionDelegate import UnderwritingDecisionDelegate

		errMsg = "Failed to add elements " + str(decisionsIds) + " for Decisions on Underwriter"

		try:
			# get the Underwriter
			underwriter = self.get( underwriterId ).first()
				
			# split on a comma with no spaces
			idList = decisionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the UnderwritingDecision		
				underwritingDecision = UnderwritingDecisionDelegate().get(id).first();	
				# add the UnderwritingDecision
				underwriter.decisions.add(underwritingDecision)
				
			# save it		
			underwriter.save()
			
			# reload and return the appropriate version
			return self.get( underwriterId );
		except Underwriter.DoesNotExist:
			raise ProcessingError(errMsg + " : Underwriter with id " + str(underwriterId) + " does not exist.")
		except UnderwritingDecision.DoesNotExist:
			raise ProcessingError(errMsg + " : UnderwritingDecision does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDecisions( self, underwriterId, decisionsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.UnderwritingDecisionDelegate import UnderwritingDecisionDelegate

		errMsg = "Failed to remove elements " + str(decisionsIds) + " for Decisions on Underwriter"

		try:
			# get the Underwriter
			underwriter = self.get( underwriterId ).first()
				
			# split on a comma with no spaces
			idList = decisionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the UnderwritingDecision		
				underwritingDecision = UnderwritingDecisionDelegate().get(id).first();	
				# add the UnderwritingDecision
				underwriter.decisions.remove(underwritingDecision)
				
			# save it		
			underwriter.save()
			
			# reload and return the appropriate version
			return self.get( underwriterId );
		except Underwriter.DoesNotExist:
			raise ProcessingError(errMsg + " : Underwriter with id " + str(underwriterId) + " does not exist.")
		except UnderwritingDecision.DoesNotExist:
			raise ProcessingError(errMsg + " : UnderwritingDecision does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
