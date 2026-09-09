from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.GovernanceBody import GovernanceBody
from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.models.Risk import Risk
from governanceOnDjango.models.ThirdParty import ThirdParty
from governanceOnDjango.models.RecordsRepository import RecordsRepository
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.ComplianceProgram import ComplianceProgram
from governanceOnDjango.models.AuditProgram import AuditProgram
from governanceOnDjango.models.BusinessUnit import BusinessUnit
from governanceOnDjango.models.Matter import Matter
from governanceOnDjango.models.DataBreach import DataBreach
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Organization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrganizationDelegate Declaration
#======================================================================
class OrganizationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, organizationId ):
		try:	
			organization = Organization.objects.filter(id=organizationId)
			return organization.first();
		except Organization.DoesNotExist:
			raise ProcessingError("Organization with id " + str(organizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, organization):
		for model in serializers.deserialize("json", organization):
			model.save()
			return model;

	def create(self, organization):
		organization.save()
		return organization;

	def saveFromJson(self, organization):
		for model in serializers.deserialize("json", organization):
			model.save()
			return organization;
	
	def save(self, organization):
		organization.save()
		return organization;
	
	def delete(self, organizationId ):
		errMsg = "Failed to delete Organization from db using id " + str(organizationId)
		
		try:
			organization = Organization.objects.get(id=organizationId)
			organization.delete()
			return True
		except Organization.DoesNotExist:
			raise ProcessingError("Organization with id " + str(organizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Organization.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Organization from db")
		except Exception:
			return None;
		
	def addGovernanceBodies( self, organizationId, governanceBodiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.GovernanceBodyDelegate import GovernanceBodyDelegate

		errMsg = "Failed to add elements " + str(governanceBodiesIds) + " for GovernanceBodies on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = governanceBodiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the GovernanceBody		
				governanceBody = GovernanceBodyDelegate().get(id).first();	
				# add the GovernanceBody
				organization.governanceBodies.add(governanceBody)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except GovernanceBody.DoesNotExist:
			raise ProcessingError(errMsg + " : GovernanceBody does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGovernanceBodies( self, organizationId, governanceBodiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.GovernanceBodyDelegate import GovernanceBodyDelegate

		errMsg = "Failed to remove elements " + str(governanceBodiesIds) + " for GovernanceBodies on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = governanceBodiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the GovernanceBody		
				governanceBody = GovernanceBodyDelegate().get(id).first();	
				# add the GovernanceBody
				organization.governanceBodies.remove(governanceBody)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except GovernanceBody.DoesNotExist:
			raise ProcessingError(errMsg + " : GovernanceBody does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPolicies( self, organizationId, policiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				organization.policies.add(policy)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, organizationId, policiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				organization.policies.remove(policy)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRisks( self, organizationId, risksIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

		errMsg = "Failed to add elements " + str(risksIds) + " for Risks on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = risksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Risk		
				risk = RiskDelegate().get(id).first();	
				# add the Risk
				organization.risks.add(risk)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRisks( self, organizationId, risksIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

		errMsg = "Failed to remove elements " + str(risksIds) + " for Risks on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = risksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Risk		
				risk = RiskDelegate().get(id).first();	
				# add the Risk
				organization.risks.remove(risk)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addThirdParties( self, organizationId, thirdPartiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

		errMsg = "Failed to add elements " + str(thirdPartiesIds) + " for ThirdParties on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = thirdPartiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ThirdParty		
				thirdParty = ThirdPartyDelegate().get(id).first();	
				# add the ThirdParty
				organization.thirdParties.add(thirdParty)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeThirdParties( self, organizationId, thirdPartiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

		errMsg = "Failed to remove elements " + str(thirdPartiesIds) + " for ThirdParties on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = thirdPartiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ThirdParty		
				thirdParty = ThirdPartyDelegate().get(id).first();	
				# add the ThirdParty
				organization.thirdParties.remove(thirdParty)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRecordsRepositories( self, organizationId, recordsRepositoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

		errMsg = "Failed to add elements " + str(recordsRepositoriesIds) + " for RecordsRepositories on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = recordsRepositoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RecordsRepository		
				recordsRepository = RecordsRepositoryDelegate().get(id).first();	
				# add the RecordsRepository
				organization.recordsRepositories.add(recordsRepository)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRecordsRepositories( self, organizationId, recordsRepositoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RecordsRepositoryDelegate import RecordsRepositoryDelegate

		errMsg = "Failed to remove elements " + str(recordsRepositoriesIds) + " for RecordsRepositories on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = recordsRepositoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RecordsRepository		
				recordsRepository = RecordsRepositoryDelegate().get(id).first();	
				# add the RecordsRepository
				organization.recordsRepositories.remove(recordsRepository)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except RecordsRepository.DoesNotExist:
			raise ProcessingError(errMsg + " : RecordsRepository does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataProcessingActivities( self, organizationId, dataProcessingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to add elements " + str(dataProcessingActivitiesIds) + " for DataProcessingActivities on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = dataProcessingActivitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				organization.dataProcessingActivities.add(dataProcessingActivity)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataProcessingActivities( self, organizationId, dataProcessingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to remove elements " + str(dataProcessingActivitiesIds) + " for DataProcessingActivities on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = dataProcessingActivitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				organization.dataProcessingActivities.remove(dataProcessingActivity)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCompliancePrograms( self, organizationId, complianceProgramsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ComplianceProgramDelegate import ComplianceProgramDelegate

		errMsg = "Failed to add elements " + str(complianceProgramsIds) + " for CompliancePrograms on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = complianceProgramsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ComplianceProgram		
				complianceProgram = ComplianceProgramDelegate().get(id).first();	
				# add the ComplianceProgram
				organization.compliancePrograms.add(complianceProgram)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCompliancePrograms( self, organizationId, complianceProgramsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ComplianceProgramDelegate import ComplianceProgramDelegate

		errMsg = "Failed to remove elements " + str(complianceProgramsIds) + " for CompliancePrograms on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = complianceProgramsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ComplianceProgram		
				complianceProgram = ComplianceProgramDelegate().get(id).first();	
				# add the ComplianceProgram
				organization.compliancePrograms.remove(complianceProgram)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAuditPrograms( self, organizationId, auditProgramsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditProgramDelegate import AuditProgramDelegate

		errMsg = "Failed to add elements " + str(auditProgramsIds) + " for AuditPrograms on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = auditProgramsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AuditProgram		
				auditProgram = AuditProgramDelegate().get(id).first();	
				# add the AuditProgram
				organization.auditPrograms.add(auditProgram)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except AuditProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditProgram does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAuditPrograms( self, organizationId, auditProgramsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditProgramDelegate import AuditProgramDelegate

		errMsg = "Failed to remove elements " + str(auditProgramsIds) + " for AuditPrograms on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = auditProgramsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AuditProgram		
				auditProgram = AuditProgramDelegate().get(id).first();	
				# add the AuditProgram
				organization.auditPrograms.remove(auditProgram)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except AuditProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditProgram does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addBusinessUnits( self, organizationId, businessUnitsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.BusinessUnitDelegate import BusinessUnitDelegate

		errMsg = "Failed to add elements " + str(businessUnitsIds) + " for BusinessUnits on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = businessUnitsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BusinessUnit		
				businessUnit = BusinessUnitDelegate().get(id).first();	
				# add the BusinessUnit
				organization.businessUnits.add(businessUnit)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBusinessUnits( self, organizationId, businessUnitsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.BusinessUnitDelegate import BusinessUnitDelegate

		errMsg = "Failed to remove elements " + str(businessUnitsIds) + " for BusinessUnits on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = businessUnitsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BusinessUnit		
				businessUnit = BusinessUnitDelegate().get(id).first();	
				# add the BusinessUnit
				organization.businessUnits.remove(businessUnit)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMatters( self, organizationId, mattersIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.MatterDelegate import MatterDelegate

		errMsg = "Failed to add elements " + str(mattersIds) + " for Matters on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = mattersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Matter		
				matter = MatterDelegate().get(id).first();	
				# add the Matter
				organization.matters.add(matter)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMatters( self, organizationId, mattersIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.MatterDelegate import MatterDelegate

		errMsg = "Failed to remove elements " + str(mattersIds) + " for Matters on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = mattersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Matter		
				matter = MatterDelegate().get(id).first();	
				# add the Matter
				organization.matters.remove(matter)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataBreaches( self, organizationId, dataBreachesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

		errMsg = "Failed to add elements " + str(dataBreachesIds) + " for DataBreaches on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = dataBreachesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataBreach		
				dataBreach = DataBreachDelegate().get(id).first();	
				# add the DataBreach
				organization.dataBreaches.add(dataBreach)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataBreaches( self, organizationId, dataBreachesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

		errMsg = "Failed to remove elements " + str(dataBreachesIds) + " for DataBreaches on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = dataBreachesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataBreach		
				dataBreach = DataBreachDelegate().get(id).first();	
				# add the DataBreach
				organization.dataBreaches.remove(dataBreach)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
