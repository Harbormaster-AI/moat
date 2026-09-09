from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.ThirdPartyAssessment import ThirdPartyAssessment
from governanceOnDjango.models.ThirdParty import ThirdParty
from governanceOnDjango.models.Issue import Issue
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ThirdPartyAssessment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ThirdPartyAssessmentDelegate Declaration
#======================================================================
class ThirdPartyAssessmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, thirdPartyAssessmentId ):
		try:	
			thirdPartyAssessment = ThirdPartyAssessment.objects.filter(id=thirdPartyAssessmentId)
			return thirdPartyAssessment.first();
		except ThirdPartyAssessment.DoesNotExist:
			raise ProcessingError("ThirdPartyAssessment with id " + str(thirdPartyAssessmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, thirdPartyAssessment):
		for model in serializers.deserialize("json", thirdPartyAssessment):
			model.save()
			return model;

	def create(self, thirdPartyAssessment):
		thirdPartyAssessment.save()
		return thirdPartyAssessment;

	def saveFromJson(self, thirdPartyAssessment):
		for model in serializers.deserialize("json", thirdPartyAssessment):
			model.save()
			return thirdPartyAssessment;
	
	def save(self, thirdPartyAssessment):
		thirdPartyAssessment.save()
		return thirdPartyAssessment;
	
	def delete(self, thirdPartyAssessmentId ):
		errMsg = "Failed to delete ThirdPartyAssessment from db using id " + str(thirdPartyAssessmentId)
		
		try:
			thirdPartyAssessment = ThirdPartyAssessment.objects.get(id=thirdPartyAssessmentId)
			thirdPartyAssessment.delete()
			return True
		except ThirdPartyAssessment.DoesNotExist:
			raise ProcessingError("ThirdPartyAssessment with id " + str(thirdPartyAssessmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ThirdPartyAssessment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ThirdPartyAssessment from db")
		except Exception:
			return None;
		
	def assignThirdParty( self, thirdPartyAssessmentId, thirdPartyId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

		errMsg = "Failed to assign element " + str(thirdPartyId) + " for ThirdParty on ThirdPartyAssessment"

		try:
			# get the ThirdPartyAssessment from db
			thirdPartyAssessment = self.get( thirdPartyAssessmentId ).first()	
			
			# get the ThirdParty from db
			thirdParty = ThirdPartyDelegate().get(thirdPartyId).first();
			
			# assign the ThirdParty		
			thirdPartyAssessment.thirdParty = thirdParty
			
			#save it
			thirdPartyAssessment.save()

			# reload and return the appropriate version					
			return self.get( thirdPartyAssessmentId );
		except ThirdPartyAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdPartyAssessment with id " + str(thirdPartyAssessmentId) + " does not exist.")
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignThirdParty( self, thirdPartyAssessmentId ):
		errMsg = "Failed to unassign element " + str(thirdPartyId) + " for ThirdParty on ThirdPartyAssessment"

		try:
			# get the ThirdPartyAssessment from db
			thirdPartyAssessment = self.get( thirdPartyAssessmentId ).first()	
			
			# assign to None for unassignment
			thirdPartyAssessment.thirdParty = None			

			#save it
			thirdPartyAssessment.save()

			# reload and return the appropriate version					
			return self.get( thirdPartyAssessmentId );
		except ThirdPartyAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdPartyAssessment with id " + str(thirdPartyAssessmentId) + " does not exist.")
		except Exception:
			return None;
		
	def addIssues( self, thirdPartyAssessmentId, issuesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

		errMsg = "Failed to add elements " + str(issuesIds) + " for Issues on ThirdPartyAssessment"

		try:
			# get the ThirdPartyAssessment
			thirdPartyAssessment = self.get( thirdPartyAssessmentId ).first()
				
			# split on a comma with no spaces
			idList = issuesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Issue		
				issue = IssueDelegate().get(id).first();	
				# add the Issue
				thirdPartyAssessment.issues.add(issue)
				
			# save it		
			thirdPartyAssessment.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyAssessmentId );
		except ThirdPartyAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdPartyAssessment with id " + str(thirdPartyAssessmentId) + " does not exist.")
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeIssues( self, thirdPartyAssessmentId, issuesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

		errMsg = "Failed to remove elements " + str(issuesIds) + " for Issues on ThirdPartyAssessment"

		try:
			# get the ThirdPartyAssessment
			thirdPartyAssessment = self.get( thirdPartyAssessmentId ).first()
				
			# split on a comma with no spaces
			idList = issuesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Issue		
				issue = IssueDelegate().get(id).first();	
				# add the Issue
				thirdPartyAssessment.issues.remove(issue)
				
			# save it		
			thirdPartyAssessment.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyAssessmentId );
		except ThirdPartyAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdPartyAssessment with id " + str(thirdPartyAssessmentId) + " does not exist.")
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
