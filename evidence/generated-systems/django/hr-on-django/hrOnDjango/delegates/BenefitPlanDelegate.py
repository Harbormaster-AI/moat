from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.BenefitPlan import BenefitPlan
from hrOnDjango.models.Organization import Organization
from hrOnDjango.models.BenefitEnrollment import BenefitEnrollment
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BenefitPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BenefitPlanDelegate Declaration
#======================================================================
class BenefitPlanDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, benefitPlanId ):
		try:	
			benefitPlan = BenefitPlan.objects.filter(id=benefitPlanId)
			return benefitPlan.first();
		except BenefitPlan.DoesNotExist:
			raise ProcessingError("BenefitPlan with id " + str(benefitPlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, benefitPlan):
		for model in serializers.deserialize("json", benefitPlan):
			model.save()
			return model;

	def create(self, benefitPlan):
		benefitPlan.save()
		return benefitPlan;

	def saveFromJson(self, benefitPlan):
		for model in serializers.deserialize("json", benefitPlan):
			model.save()
			return benefitPlan;
	
	def save(self, benefitPlan):
		benefitPlan.save()
		return benefitPlan;
	
	def delete(self, benefitPlanId ):
		errMsg = "Failed to delete BenefitPlan from db using id " + str(benefitPlanId)
		
		try:
			benefitPlan = BenefitPlan.objects.get(id=benefitPlanId)
			benefitPlan.delete()
			return True
		except BenefitPlan.DoesNotExist:
			raise ProcessingError("BenefitPlan with id " + str(benefitPlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BenefitPlan.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BenefitPlan from db")
		except Exception:
			return None;
		
	def assignOrganization( self, benefitPlanId, organizationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on BenefitPlan"

		try:
			# get the BenefitPlan from db
			benefitPlan = self.get( benefitPlanId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			benefitPlan.organization = organization
			
			#save it
			benefitPlan.save()

			# reload and return the appropriate version					
			return self.get( benefitPlanId );
		except BenefitPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitPlan with id " + str(benefitPlanId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, benefitPlanId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on BenefitPlan"

		try:
			# get the BenefitPlan from db
			benefitPlan = self.get( benefitPlanId ).first()	
			
			# assign to None for unassignment
			benefitPlan.organization = None			

			#save it
			benefitPlan.save()

			# reload and return the appropriate version					
			return self.get( benefitPlanId );
		except BenefitPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitPlan with id " + str(benefitPlanId) + " does not exist.")
		except Exception:
			return None;
		
	def addEnrollments( self, benefitPlanId, enrollmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BenefitEnrollmentDelegate import BenefitEnrollmentDelegate

		errMsg = "Failed to add elements " + str(enrollmentsIds) + " for Enrollments on BenefitPlan"

		try:
			# get the BenefitPlan
			benefitPlan = self.get( benefitPlanId ).first()
				
			# split on a comma with no spaces
			idList = enrollmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BenefitEnrollment		
				benefitEnrollment = BenefitEnrollmentDelegate().get(id).first();	
				# add the BenefitEnrollment
				benefitPlan.enrollments.add(benefitEnrollment)
				
			# save it		
			benefitPlan.save()
			
			# reload and return the appropriate version
			return self.get( benefitPlanId );
		except BenefitPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitPlan with id " + str(benefitPlanId) + " does not exist.")
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEnrollments( self, benefitPlanId, enrollmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BenefitEnrollmentDelegate import BenefitEnrollmentDelegate

		errMsg = "Failed to remove elements " + str(enrollmentsIds) + " for Enrollments on BenefitPlan"

		try:
			# get the BenefitPlan
			benefitPlan = self.get( benefitPlanId ).first()
				
			# split on a comma with no spaces
			idList = enrollmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BenefitEnrollment		
				benefitEnrollment = BenefitEnrollmentDelegate().get(id).first();	
				# add the BenefitEnrollment
				benefitPlan.enrollments.remove(benefitEnrollment)
				
			# save it		
			benefitPlan.save()
			
			# reload and return the appropriate version
			return self.get( benefitPlanId );
		except BenefitPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitPlan with id " + str(benefitPlanId) + " does not exist.")
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
