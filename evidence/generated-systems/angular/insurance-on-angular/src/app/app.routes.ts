import { Routes } from '@angular/router';

import { CreateInsurerComponent } from './components/Insurer/create/create.component';
import { EditInsurerComponent } from './components/Insurer/edit/edit.component';
import { IndexInsurerComponent } from './components/Insurer/index/index.component';
import { CreateInsuranceProductComponent } from './components/InsuranceProduct/create/create.component';
import { EditInsuranceProductComponent } from './components/InsuranceProduct/edit/edit.component';
import { IndexInsuranceProductComponent } from './components/InsuranceProduct/index/index.component';
import { CreateCoverageDefinitionComponent } from './components/CoverageDefinition/create/create.component';
import { EditCoverageDefinitionComponent } from './components/CoverageDefinition/edit/edit.component';
import { IndexCoverageDefinitionComponent } from './components/CoverageDefinition/index/index.component';
import { CreateDistributorComponent } from './components/Distributor/create/create.component';
import { EditDistributorComponent } from './components/Distributor/edit/edit.component';
import { IndexDistributorComponent } from './components/Distributor/index/index.component';
import { CreateAgentComponent } from './components/Agent/create/create.component';
import { EditAgentComponent } from './components/Agent/edit/edit.component';
import { IndexAgentComponent } from './components/Agent/index/index.component';
import { CreateCustomerComponent } from './components/Customer/create/create.component';
import { EditCustomerComponent } from './components/Customer/edit/edit.component';
import { IndexCustomerComponent } from './components/Customer/index/index.component';
import { CreateApplicationComponent } from './components/Application/create/create.component';
import { EditApplicationComponent } from './components/Application/edit/edit.component';
import { IndexApplicationComponent } from './components/Application/index/index.component';
import { CreateQuoteComponent } from './components/Quote/create/create.component';
import { EditQuoteComponent } from './components/Quote/edit/edit.component';
import { IndexQuoteComponent } from './components/Quote/index/index.component';
import { CreateUnderwritingDecisionComponent } from './components/UnderwritingDecision/create/create.component';
import { EditUnderwritingDecisionComponent } from './components/UnderwritingDecision/edit/edit.component';
import { IndexUnderwritingDecisionComponent } from './components/UnderwritingDecision/index/index.component';
import { CreateUnderwriterComponent } from './components/Underwriter/create/create.component';
import { EditUnderwriterComponent } from './components/Underwriter/edit/edit.component';
import { IndexUnderwriterComponent } from './components/Underwriter/index/index.component';
import { CreatePolicyComponent } from './components/Policy/create/create.component';
import { EditPolicyComponent } from './components/Policy/edit/edit.component';
import { IndexPolicyComponent } from './components/Policy/index/index.component';
import { CreateEndorsementComponent } from './components/Endorsement/create/create.component';
import { EditEndorsementComponent } from './components/Endorsement/edit/edit.component';
import { IndexEndorsementComponent } from './components/Endorsement/index/index.component';
import { CreatePolicyCoverageComponent } from './components/PolicyCoverage/create/create.component';
import { EditPolicyCoverageComponent } from './components/PolicyCoverage/edit/edit.component';
import { IndexPolicyCoverageComponent } from './components/PolicyCoverage/index/index.component';
import { CreateInsuredObjectComponent } from './components/InsuredObject/create/create.component';
import { EditInsuredObjectComponent } from './components/InsuredObject/edit/edit.component';
import { IndexInsuredObjectComponent } from './components/InsuredObject/index/index.component';
import { CreateBeneficiaryComponent } from './components/Beneficiary/create/create.component';
import { EditBeneficiaryComponent } from './components/Beneficiary/edit/edit.component';
import { IndexBeneficiaryComponent } from './components/Beneficiary/index/index.component';
import { CreateBillingAccountComponent } from './components/BillingAccount/create/create.component';
import { EditBillingAccountComponent } from './components/BillingAccount/edit/edit.component';
import { IndexBillingAccountComponent } from './components/BillingAccount/index/index.component';
import { CreateInvoiceComponent } from './components/Invoice/create/create.component';
import { EditInvoiceComponent } from './components/Invoice/edit/edit.component';
import { IndexInvoiceComponent } from './components/Invoice/index/index.component';
import { CreatePaymentComponent } from './components/Payment/create/create.component';
import { EditPaymentComponent } from './components/Payment/edit/edit.component';
import { IndexPaymentComponent } from './components/Payment/index/index.component';
import { CreateClaimComponent } from './components/Claim/create/create.component';
import { EditClaimComponent } from './components/Claim/edit/edit.component';
import { IndexClaimComponent } from './components/Claim/index/index.component';
import { CreateIncidentComponent } from './components/Incident/create/create.component';
import { EditIncidentComponent } from './components/Incident/edit/edit.component';
import { IndexIncidentComponent } from './components/Incident/index/index.component';
import { CreateExposureComponent } from './components/Exposure/create/create.component';
import { EditExposureComponent } from './components/Exposure/edit/edit.component';
import { IndexExposureComponent } from './components/Exposure/index/index.component';
import { CreateAdjusterComponent } from './components/Adjuster/create/create.component';
import { EditAdjusterComponent } from './components/Adjuster/edit/edit.component';
import { IndexAdjusterComponent } from './components/Adjuster/index/index.component';
import { CreateClaimReserveComponent } from './components/ClaimReserve/create/create.component';
import { EditClaimReserveComponent } from './components/ClaimReserve/edit/edit.component';
import { IndexClaimReserveComponent } from './components/ClaimReserve/index/index.component';
import { CreateClaimPaymentComponent } from './components/ClaimPayment/create/create.component';
import { EditClaimPaymentComponent } from './components/ClaimPayment/edit/edit.component';
import { IndexClaimPaymentComponent } from './components/ClaimPayment/index/index.component';
import { CreateServiceProviderComponent } from './components/ServiceProvider/create/create.component';
import { EditServiceProviderComponent } from './components/ServiceProvider/edit/edit.component';
import { IndexServiceProviderComponent } from './components/ServiceProvider/index/index.component';
import { CreateReinsuranceAgreementComponent } from './components/ReinsuranceAgreement/create/create.component';
import { EditReinsuranceAgreementComponent } from './components/ReinsuranceAgreement/edit/edit.component';
import { IndexReinsuranceAgreementComponent } from './components/ReinsuranceAgreement/index/index.component';
import { CreateSubrogationRecoveryComponent } from './components/SubrogationRecovery/create/create.component';
import { EditSubrogationRecoveryComponent } from './components/SubrogationRecovery/edit/edit.component';
import { IndexSubrogationRecoveryComponent } from './components/SubrogationRecovery/index/index.component';
import { CreateThirdPartyComponent } from './components/ThirdParty/create/create.component';
import { EditThirdPartyComponent } from './components/ThirdParty/edit/edit.component';
import { IndexThirdPartyComponent } from './components/ThirdParty/index/index.component';
import { CreateDocumentComponent } from './components/Document/create/create.component';
import { EditDocumentComponent } from './components/Document/edit/edit.component';
import { IndexDocumentComponent } from './components/Document/index/index.component';

export const routes: Routes = [

        {
        path: 'createInsurer',
        component: CreateInsurerComponent
},
{
    path: 'editInsurer/:id',
        component: EditInsurerComponent
},
{
    path: 'indexInsurer',
        component: IndexInsurerComponent
},
    {
        path: 'createInsuranceProduct',
        component: CreateInsuranceProductComponent
},
{
    path: 'editInsuranceProduct/:id',
        component: EditInsuranceProductComponent
},
{
    path: 'indexInsuranceProduct',
        component: IndexInsuranceProductComponent
},
    {
        path: 'createCoverageDefinition',
        component: CreateCoverageDefinitionComponent
},
{
    path: 'editCoverageDefinition/:id',
        component: EditCoverageDefinitionComponent
},
{
    path: 'indexCoverageDefinition',
        component: IndexCoverageDefinitionComponent
},
    {
        path: 'createDistributor',
        component: CreateDistributorComponent
},
{
    path: 'editDistributor/:id',
        component: EditDistributorComponent
},
{
    path: 'indexDistributor',
        component: IndexDistributorComponent
},
    {
        path: 'createAgent',
        component: CreateAgentComponent
},
{
    path: 'editAgent/:id',
        component: EditAgentComponent
},
{
    path: 'indexAgent',
        component: IndexAgentComponent
},
    {
        path: 'createCustomer',
        component: CreateCustomerComponent
},
{
    path: 'editCustomer/:id',
        component: EditCustomerComponent
},
{
    path: 'indexCustomer',
        component: IndexCustomerComponent
},
    {
        path: 'createApplication',
        component: CreateApplicationComponent
},
{
    path: 'editApplication/:id',
        component: EditApplicationComponent
},
{
    path: 'indexApplication',
        component: IndexApplicationComponent
},
    {
        path: 'createQuote',
        component: CreateQuoteComponent
},
{
    path: 'editQuote/:id',
        component: EditQuoteComponent
},
{
    path: 'indexQuote',
        component: IndexQuoteComponent
},
    {
        path: 'createUnderwritingDecision',
        component: CreateUnderwritingDecisionComponent
},
{
    path: 'editUnderwritingDecision/:id',
        component: EditUnderwritingDecisionComponent
},
{
    path: 'indexUnderwritingDecision',
        component: IndexUnderwritingDecisionComponent
},
    {
        path: 'createUnderwriter',
        component: CreateUnderwriterComponent
},
{
    path: 'editUnderwriter/:id',
        component: EditUnderwriterComponent
},
{
    path: 'indexUnderwriter',
        component: IndexUnderwriterComponent
},
    {
        path: 'createPolicy',
        component: CreatePolicyComponent
},
{
    path: 'editPolicy/:id',
        component: EditPolicyComponent
},
{
    path: 'indexPolicy',
        component: IndexPolicyComponent
},
    {
        path: 'createEndorsement',
        component: CreateEndorsementComponent
},
{
    path: 'editEndorsement/:id',
        component: EditEndorsementComponent
},
{
    path: 'indexEndorsement',
        component: IndexEndorsementComponent
},
    {
        path: 'createPolicyCoverage',
        component: CreatePolicyCoverageComponent
},
{
    path: 'editPolicyCoverage/:id',
        component: EditPolicyCoverageComponent
},
{
    path: 'indexPolicyCoverage',
        component: IndexPolicyCoverageComponent
},
    {
        path: 'createInsuredObject',
        component: CreateInsuredObjectComponent
},
{
    path: 'editInsuredObject/:id',
        component: EditInsuredObjectComponent
},
{
    path: 'indexInsuredObject',
        component: IndexInsuredObjectComponent
},
    {
        path: 'createBeneficiary',
        component: CreateBeneficiaryComponent
},
{
    path: 'editBeneficiary/:id',
        component: EditBeneficiaryComponent
},
{
    path: 'indexBeneficiary',
        component: IndexBeneficiaryComponent
},
    {
        path: 'createBillingAccount',
        component: CreateBillingAccountComponent
},
{
    path: 'editBillingAccount/:id',
        component: EditBillingAccountComponent
},
{
    path: 'indexBillingAccount',
        component: IndexBillingAccountComponent
},
    {
        path: 'createInvoice',
        component: CreateInvoiceComponent
},
{
    path: 'editInvoice/:id',
        component: EditInvoiceComponent
},
{
    path: 'indexInvoice',
        component: IndexInvoiceComponent
},
    {
        path: 'createPayment',
        component: CreatePaymentComponent
},
{
    path: 'editPayment/:id',
        component: EditPaymentComponent
},
{
    path: 'indexPayment',
        component: IndexPaymentComponent
},
    {
        path: 'createClaim',
        component: CreateClaimComponent
},
{
    path: 'editClaim/:id',
        component: EditClaimComponent
},
{
    path: 'indexClaim',
        component: IndexClaimComponent
},
    {
        path: 'createIncident',
        component: CreateIncidentComponent
},
{
    path: 'editIncident/:id',
        component: EditIncidentComponent
},
{
    path: 'indexIncident',
        component: IndexIncidentComponent
},
    {
        path: 'createExposure',
        component: CreateExposureComponent
},
{
    path: 'editExposure/:id',
        component: EditExposureComponent
},
{
    path: 'indexExposure',
        component: IndexExposureComponent
},
    {
        path: 'createAdjuster',
        component: CreateAdjusterComponent
},
{
    path: 'editAdjuster/:id',
        component: EditAdjusterComponent
},
{
    path: 'indexAdjuster',
        component: IndexAdjusterComponent
},
    {
        path: 'createClaimReserve',
        component: CreateClaimReserveComponent
},
{
    path: 'editClaimReserve/:id',
        component: EditClaimReserveComponent
},
{
    path: 'indexClaimReserve',
        component: IndexClaimReserveComponent
},
    {
        path: 'createClaimPayment',
        component: CreateClaimPaymentComponent
},
{
    path: 'editClaimPayment/:id',
        component: EditClaimPaymentComponent
},
{
    path: 'indexClaimPayment',
        component: IndexClaimPaymentComponent
},
    {
        path: 'createServiceProvider',
        component: CreateServiceProviderComponent
},
{
    path: 'editServiceProvider/:id',
        component: EditServiceProviderComponent
},
{
    path: 'indexServiceProvider',
        component: IndexServiceProviderComponent
},
    {
        path: 'createReinsuranceAgreement',
        component: CreateReinsuranceAgreementComponent
},
{
    path: 'editReinsuranceAgreement/:id',
        component: EditReinsuranceAgreementComponent
},
{
    path: 'indexReinsuranceAgreement',
        component: IndexReinsuranceAgreementComponent
},
    {
        path: 'createSubrogationRecovery',
        component: CreateSubrogationRecoveryComponent
},
{
    path: 'editSubrogationRecovery/:id',
        component: EditSubrogationRecoveryComponent
},
{
    path: 'indexSubrogationRecovery',
        component: IndexSubrogationRecoveryComponent
},
    {
        path: 'createThirdParty',
        component: CreateThirdPartyComponent
},
{
    path: 'editThirdParty/:id',
        component: EditThirdPartyComponent
},
{
    path: 'indexThirdParty',
        component: IndexThirdPartyComponent
},
    {
        path: 'createDocument',
        component: CreateDocumentComponent
},
{
    path: 'editDocument/:id',
        component: EditDocumentComponent
},
{
    path: 'indexDocument',
        component: IndexDocumentComponent
}
];