import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Organization} from '../models/Organization';
import {GovernanceBodyService} from '../services/GovernanceBody.service';
import {PolicyService} from '../services/Policy.service';
import {RiskService} from '../services/Risk.service';
import {ThirdPartyService} from '../services/ThirdParty.service';
import {RecordsRepositoryService} from '../services/RecordsRepository.service';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {ComplianceProgramService} from '../services/ComplianceProgram.service';
import {AuditProgramService} from '../services/AuditProgram.service';
import {BusinessUnitService} from '../services/BusinessUnit.service';
import {MatterService} from '../services/Matter.service';
import {DataBreachService} from '../services/DataBreach.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OrganizationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	organization : Organization;

	//********************************************************************
	// Catch all for the return value of a service call
	//********************************************************************
	result: any;

	//********************************************************************
	// sole constructor, injected with the HttpClient
	//********************************************************************
	constructor(private http: HttpClient) {
		super();
	}

		//********************************************************************
	// add a Organization
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOrganization(name, legalName, jurisdiction, industrySector, GovernanceBodies, Policies, Risks, ThirdParties, RecordsRepositories, DataProcessingActivities, CompliancePrograms, AuditPrograms, BusinessUnits, Matters, DataBreaches) : Observable<any> {
		const uri_ = this.apiUrl + '/Organization/create';
		const obj = {
			      		name: name,
      		legalName: legalName,
      		jurisdiction: jurisdiction,
      		industrySector: industrySector,
      		GovernanceBodies: GovernanceBodies != null && GovernanceBodies.length > 0 ? GovernanceBodies : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Risks: Risks != null && Risks.length > 0 ? Risks : null,
      		ThirdParties: ThirdParties != null && ThirdParties.length > 0 ? ThirdParties : null,
      		RecordsRepositories: RecordsRepositories != null && RecordsRepositories.length > 0 ? RecordsRepositories : null,
      		DataProcessingActivities: DataProcessingActivities != null && DataProcessingActivities.length > 0 ? DataProcessingActivities : null,
      		CompliancePrograms: CompliancePrograms != null && CompliancePrograms.length > 0 ? CompliancePrograms : null,
      		AuditPrograms: AuditPrograms != null && AuditPrograms.length > 0 ? AuditPrograms : null,
      		BusinessUnits: BusinessUnits != null && BusinessUnits.length > 0 ? BusinessUnits : null,
      		Matters: Matters != null && Matters.length > 0 ? Matters : null,
			DataBreaches: DataBreaches != null && DataBreaches.length > 0 ? DataBreaches : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Organization
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOrganization(name, legalName, jurisdiction, industrySector, GovernanceBodies, Policies, Risks, ThirdParties, RecordsRepositories, DataProcessingActivities, CompliancePrograms, AuditPrograms, BusinessUnits, Matters, DataBreaches, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Organization/update/' + id;
		const obj = {
				      		name: name,
      		legalName: legalName,
      		jurisdiction: jurisdiction,
      		industrySector: industrySector,
      		GovernanceBodies: GovernanceBodies != null && GovernanceBodies.length > 0 ? GovernanceBodies : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Risks: Risks != null && Risks.length > 0 ? Risks : null,
      		ThirdParties: ThirdParties != null && ThirdParties.length > 0 ? ThirdParties : null,
      		RecordsRepositories: RecordsRepositories != null && RecordsRepositories.length > 0 ? RecordsRepositories : null,
      		DataProcessingActivities: DataProcessingActivities != null && DataProcessingActivities.length > 0 ? DataProcessingActivities : null,
      		CompliancePrograms: CompliancePrograms != null && CompliancePrograms.length > 0 ? CompliancePrograms : null,
      		AuditPrograms: AuditPrograms != null && AuditPrograms.length > 0 ? AuditPrograms : null,
      		BusinessUnits: BusinessUnits != null && BusinessUnits.length > 0 ? BusinessUnits : null,
      		Matters: Matters != null && Matters.length > 0 ? Matters : null,
			DataBreaches: DataBreaches != null && DataBreaches.length > 0 ? DataBreaches : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Organization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOrganization(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Organization/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Organization
	// returns the results untouched as an Observable Organization
	// Organization model
	// delegates via URI
	//********************************************************************
	getOrganization(id) : Observable<Organization> {
		const uri_ = this.apiUrl + '/Organization/load/' + id;

		return this.http.get<Organization>(uri_);
	}
	
	//********************************************************************
	// gets all Organization
	// returns the results untouched as JSON representation of an
	// Observable array of Organization models
	// delegates via URI
	//********************************************************************
	getOrganizations() : Observable<Organization[]> {
		const uri_ = this.apiUrl + '/Organization/';

		return this
			.http.get<Organization[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more governanceBodiesIds as a GovernanceBodies
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addGovernanceBodies( organizationId, governanceBodiesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = governanceBodiesIds.split(',')

	// iterate over array of governanceBodies ids
	idList.forEach(function (id) {
		// read the GovernanceBody
		var governanceBody = new GovernanceBodyService(this.http).getGovernanceBody(id);
		// add the GovernanceBody if not already assigned
		if ( this.organization.governanceBodies.indexOf(governanceBody) == -1 )
		this.organization.governanceBodies.push(governanceBody);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more governanceBodiesIds as a GovernanceBodies
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeGovernanceBodies( organizationId, governanceBodiesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= governanceBodiesIds.split(',');
	var governanceBodies 	= this.organization.governanceBodies;

	if ( governanceBodies != null && governanceBodiesIds != null ) {

		// iterate over array of governanceBodies ids
		governanceBodies.forEach(function (obj) {
			if ( governanceBodiesIds.indexOf(obj._id) > -1 ) {
				// remove the GovernanceBody
				this.organization.governanceBodies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( organizationId, policiesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.organization.policies.indexOf(policy) == -1 )
		this.organization.policies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( organizationId, policiesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.organization.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.organization.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more risksIds as a Risks
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRisks( organizationId, risksIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = risksIds.split(',')

	// iterate over array of risks ids
	idList.forEach(function (id) {
		// read the Risk
		var risk = new RiskService(this.http).getRisk(id);
		// add the Risk if not already assigned
		if ( this.organization.risks.indexOf(risk) == -1 )
		this.organization.risks.push(risk);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more risksIds as a Risks
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRisks( organizationId, risksIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= risksIds.split(',');
	var risks 	= this.organization.risks;

	if ( risks != null && risksIds != null ) {

		// iterate over array of risks ids
		risks.forEach(function (obj) {
			if ( risksIds.indexOf(obj._id) > -1 ) {
				// remove the Risk
				this.organization.risks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more thirdPartiesIds as a ThirdParties
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addThirdParties( organizationId, thirdPartiesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = thirdPartiesIds.split(',')

	// iterate over array of thirdParties ids
	idList.forEach(function (id) {
		// read the ThirdParty
		var thirdParty = new ThirdPartyService(this.http).getThirdParty(id);
		// add the ThirdParty if not already assigned
		if ( this.organization.thirdParties.indexOf(thirdParty) == -1 )
		this.organization.thirdParties.push(thirdParty);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more thirdPartiesIds as a ThirdParties
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeThirdParties( organizationId, thirdPartiesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= thirdPartiesIds.split(',');
	var thirdParties 	= this.organization.thirdParties;

	if ( thirdParties != null && thirdPartiesIds != null ) {

		// iterate over array of thirdParties ids
		thirdParties.forEach(function (obj) {
			if ( thirdPartiesIds.indexOf(obj._id) > -1 ) {
				// remove the ThirdParty
				this.organization.thirdParties.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more recordsRepositoriesIds as a RecordsRepositories
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRecordsRepositories( organizationId, recordsRepositoriesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = recordsRepositoriesIds.split(',')

	// iterate over array of recordsRepositories ids
	idList.forEach(function (id) {
		// read the RecordsRepository
		var recordsRepository = new RecordsRepositoryService(this.http).getRecordsRepository(id);
		// add the RecordsRepository if not already assigned
		if ( this.organization.recordsRepositories.indexOf(recordsRepository) == -1 )
		this.organization.recordsRepositories.push(recordsRepository);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more recordsRepositoriesIds as a RecordsRepositories
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRecordsRepositories( organizationId, recordsRepositoriesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= recordsRepositoriesIds.split(',');
	var recordsRepositories 	= this.organization.recordsRepositories;

	if ( recordsRepositories != null && recordsRepositoriesIds != null ) {

		// iterate over array of recordsRepositories ids
		recordsRepositories.forEach(function (obj) {
			if ( recordsRepositoriesIds.indexOf(obj._id) > -1 ) {
				// remove the RecordsRepository
				this.organization.recordsRepositories.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataProcessingActivitiesIds as a DataProcessingActivities
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataProcessingActivities( organizationId, dataProcessingActivitiesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = dataProcessingActivitiesIds.split(',')

	// iterate over array of dataProcessingActivities ids
	idList.forEach(function (id) {
		// read the DataProcessingActivity
		var dataProcessingActivity = new DataProcessingActivityService(this.http).getDataProcessingActivity(id);
		// add the DataProcessingActivity if not already assigned
		if ( this.organization.dataProcessingActivities.indexOf(dataProcessingActivity) == -1 )
		this.organization.dataProcessingActivities.push(dataProcessingActivity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataProcessingActivitiesIds as a DataProcessingActivities
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataProcessingActivities( organizationId, dataProcessingActivitiesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= dataProcessingActivitiesIds.split(',');
	var dataProcessingActivities 	= this.organization.dataProcessingActivities;

	if ( dataProcessingActivities != null && dataProcessingActivitiesIds != null ) {

		// iterate over array of dataProcessingActivities ids
		dataProcessingActivities.forEach(function (obj) {
			if ( dataProcessingActivitiesIds.indexOf(obj._id) > -1 ) {
				// remove the DataProcessingActivity
				this.organization.dataProcessingActivities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more complianceProgramsIds as a CompliancePrograms
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCompliancePrograms( organizationId, complianceProgramsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = complianceProgramsIds.split(',')

	// iterate over array of compliancePrograms ids
	idList.forEach(function (id) {
		// read the ComplianceProgram
		var complianceProgram = new ComplianceProgramService(this.http).getComplianceProgram(id);
		// add the ComplianceProgram if not already assigned
		if ( this.organization.compliancePrograms.indexOf(complianceProgram) == -1 )
		this.organization.compliancePrograms.push(complianceProgram);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more complianceProgramsIds as a CompliancePrograms
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCompliancePrograms( organizationId, complianceProgramsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= complianceProgramsIds.split(',');
	var compliancePrograms 	= this.organization.compliancePrograms;

	if ( compliancePrograms != null && complianceProgramsIds != null ) {

		// iterate over array of compliancePrograms ids
		compliancePrograms.forEach(function (obj) {
			if ( complianceProgramsIds.indexOf(obj._id) > -1 ) {
				// remove the ComplianceProgram
				this.organization.compliancePrograms.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more auditProgramsIds as a AuditPrograms
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAuditPrograms( organizationId, auditProgramsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = auditProgramsIds.split(',')

	// iterate over array of auditPrograms ids
	idList.forEach(function (id) {
		// read the AuditProgram
		var auditProgram = new AuditProgramService(this.http).getAuditProgram(id);
		// add the AuditProgram if not already assigned
		if ( this.organization.auditPrograms.indexOf(auditProgram) == -1 )
		this.organization.auditPrograms.push(auditProgram);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more auditProgramsIds as a AuditPrograms
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAuditPrograms( organizationId, auditProgramsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= auditProgramsIds.split(',');
	var auditPrograms 	= this.organization.auditPrograms;

	if ( auditPrograms != null && auditProgramsIds != null ) {

		// iterate over array of auditPrograms ids
		auditPrograms.forEach(function (obj) {
			if ( auditProgramsIds.indexOf(obj._id) > -1 ) {
				// remove the AuditProgram
				this.organization.auditPrograms.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more businessUnitsIds as a BusinessUnits
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBusinessUnits( organizationId, businessUnitsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = businessUnitsIds.split(',')

	// iterate over array of businessUnits ids
	idList.forEach(function (id) {
		// read the BusinessUnit
		var businessUnit = new BusinessUnitService(this.http).getBusinessUnit(id);
		// add the BusinessUnit if not already assigned
		if ( this.organization.businessUnits.indexOf(businessUnit) == -1 )
		this.organization.businessUnits.push(businessUnit);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more businessUnitsIds as a BusinessUnits
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBusinessUnits( organizationId, businessUnitsIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= businessUnitsIds.split(',');
	var businessUnits 	= this.organization.businessUnits;

	if ( businessUnits != null && businessUnitsIds != null ) {

		// iterate over array of businessUnits ids
		businessUnits.forEach(function (obj) {
			if ( businessUnitsIds.indexOf(obj._id) > -1 ) {
				// remove the BusinessUnit
				this.organization.businessUnits.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more mattersIds as a Matters
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMatters( organizationId, mattersIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = mattersIds.split(',')

	// iterate over array of matters ids
	idList.forEach(function (id) {
		// read the Matter
		var matter = new MatterService(this.http).getMatter(id);
		// add the Matter if not already assigned
		if ( this.organization.matters.indexOf(matter) == -1 )
		this.organization.matters.push(matter);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more mattersIds as a Matters
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMatters( organizationId, mattersIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= mattersIds.split(',');
	var matters 	= this.organization.matters;

	if ( matters != null && mattersIds != null ) {

		// iterate over array of matters ids
		matters.forEach(function (obj) {
			if ( mattersIds.indexOf(obj._id) > -1 ) {
				// remove the Matter
				this.organization.matters.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataBreachesIds as a DataBreaches
	// to a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataBreaches( organizationId, dataBreachesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );

	// split on a comma with no spaces
	var idList = dataBreachesIds.split(',')

	// iterate over array of dataBreaches ids
	idList.forEach(function (id) {
		// read the DataBreach
		var dataBreach = new DataBreachService(this.http).getDataBreach(id);
		// add the DataBreach if not already assigned
		if ( this.organization.dataBreaches.indexOf(dataBreach) == -1 )
		this.organization.dataBreaches.push(dataBreach);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataBreachesIds as a DataBreaches
	// from a Organization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataBreaches( organizationId, dataBreachesIds ): Observable<any> {

		// get the Organization
		this.loadHelper( organizationId );


	// split on a comma with no spaces
	var idList 					= dataBreachesIds.split(',');
	var dataBreaches 	= this.organization.dataBreaches;

	if ( dataBreaches != null && dataBreachesIds != null ) {

		// iterate over array of dataBreaches ids
		dataBreaches.forEach(function (obj) {
			if ( dataBreachesIds.indexOf(obj._id) > -1 ) {
				// remove the DataBreach
				this.organization.dataBreaches.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Organization
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Organization/update/' + this.organization;

	return  this.http.post(uri_, this.organization );
}

	//********************************************************************
	// loadHelper - internal helper to load a Organization
	//********************************************************************	
	loadHelper( id ) {
		this.getOrganization(id)
			.subscribe((res : Organization) => {
				this.organization = res;
			});
	}
}