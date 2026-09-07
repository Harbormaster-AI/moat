import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatInputModule} from '@angular/material/input';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatMomentDateModule} from "@angular/material-moment-adapter";
import {NgModule} from '@angular/core';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {RouterModule} from '@angular/router';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule} from '@angular/forms';
import {ReactiveFormsModule} from '@angular/forms';
import {AppComponent} from './app.component';
import {MatMenuModule} from '@angular/material/menu';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatSidenavModule} from '@angular/material/sidenav'

import {IndexInsurerComponent} from './components/Insurer/index/index.component';
import {CreateInsurerComponent} from './components/Insurer/create/create.component';
import {EditInsurerComponent} from './components/Insurer/edit/edit.component';
import {IndexInsuranceProductComponent} from './components/InsuranceProduct/index/index.component';
import {CreateInsuranceProductComponent} from './components/InsuranceProduct/create/create.component';
import {EditInsuranceProductComponent} from './components/InsuranceProduct/edit/edit.component';
import {IndexCoverageDefinitionComponent} from './components/CoverageDefinition/index/index.component';
import {CreateCoverageDefinitionComponent} from './components/CoverageDefinition/create/create.component';
import {EditCoverageDefinitionComponent} from './components/CoverageDefinition/edit/edit.component';
import {IndexDistributorComponent} from './components/Distributor/index/index.component';
import {CreateDistributorComponent} from './components/Distributor/create/create.component';
import {EditDistributorComponent} from './components/Distributor/edit/edit.component';
import {IndexAgentComponent} from './components/Agent/index/index.component';
import {CreateAgentComponent} from './components/Agent/create/create.component';
import {EditAgentComponent} from './components/Agent/edit/edit.component';
import {IndexCustomerComponent} from './components/Customer/index/index.component';
import {CreateCustomerComponent} from './components/Customer/create/create.component';
import {EditCustomerComponent} from './components/Customer/edit/edit.component';
import {IndexApplicationComponent} from './components/Application/index/index.component';
import {CreateApplicationComponent} from './components/Application/create/create.component';
import {EditApplicationComponent} from './components/Application/edit/edit.component';
import {IndexQuoteComponent} from './components/Quote/index/index.component';
import {CreateQuoteComponent} from './components/Quote/create/create.component';
import {EditQuoteComponent} from './components/Quote/edit/edit.component';
import {IndexUnderwritingDecisionComponent} from './components/UnderwritingDecision/index/index.component';
import {CreateUnderwritingDecisionComponent} from './components/UnderwritingDecision/create/create.component';
import {EditUnderwritingDecisionComponent} from './components/UnderwritingDecision/edit/edit.component';
import {IndexUnderwriterComponent} from './components/Underwriter/index/index.component';
import {CreateUnderwriterComponent} from './components/Underwriter/create/create.component';
import {EditUnderwriterComponent} from './components/Underwriter/edit/edit.component';
import {IndexPolicyComponent} from './components/Policy/index/index.component';
import {CreatePolicyComponent} from './components/Policy/create/create.component';
import {EditPolicyComponent} from './components/Policy/edit/edit.component';
import {IndexEndorsementComponent} from './components/Endorsement/index/index.component';
import {CreateEndorsementComponent} from './components/Endorsement/create/create.component';
import {EditEndorsementComponent} from './components/Endorsement/edit/edit.component';
import {IndexPolicyCoverageComponent} from './components/PolicyCoverage/index/index.component';
import {CreatePolicyCoverageComponent} from './components/PolicyCoverage/create/create.component';
import {EditPolicyCoverageComponent} from './components/PolicyCoverage/edit/edit.component';
import {IndexInsuredObjectComponent} from './components/InsuredObject/index/index.component';
import {CreateInsuredObjectComponent} from './components/InsuredObject/create/create.component';
import {EditInsuredObjectComponent} from './components/InsuredObject/edit/edit.component';
import {IndexBeneficiaryComponent} from './components/Beneficiary/index/index.component';
import {CreateBeneficiaryComponent} from './components/Beneficiary/create/create.component';
import {EditBeneficiaryComponent} from './components/Beneficiary/edit/edit.component';
import {IndexBillingAccountComponent} from './components/BillingAccount/index/index.component';
import {CreateBillingAccountComponent} from './components/BillingAccount/create/create.component';
import {EditBillingAccountComponent} from './components/BillingAccount/edit/edit.component';
import {IndexInvoiceComponent} from './components/Invoice/index/index.component';
import {CreateInvoiceComponent} from './components/Invoice/create/create.component';
import {EditInvoiceComponent} from './components/Invoice/edit/edit.component';
import {IndexPaymentComponent} from './components/Payment/index/index.component';
import {CreatePaymentComponent} from './components/Payment/create/create.component';
import {EditPaymentComponent} from './components/Payment/edit/edit.component';
import {IndexClaimComponent} from './components/Claim/index/index.component';
import {CreateClaimComponent} from './components/Claim/create/create.component';
import {EditClaimComponent} from './components/Claim/edit/edit.component';
import {IndexIncidentComponent} from './components/Incident/index/index.component';
import {CreateIncidentComponent} from './components/Incident/create/create.component';
import {EditIncidentComponent} from './components/Incident/edit/edit.component';
import {IndexExposureComponent} from './components/Exposure/index/index.component';
import {CreateExposureComponent} from './components/Exposure/create/create.component';
import {EditExposureComponent} from './components/Exposure/edit/edit.component';
import {IndexAdjusterComponent} from './components/Adjuster/index/index.component';
import {CreateAdjusterComponent} from './components/Adjuster/create/create.component';
import {EditAdjusterComponent} from './components/Adjuster/edit/edit.component';
import {IndexClaimReserveComponent} from './components/ClaimReserve/index/index.component';
import {CreateClaimReserveComponent} from './components/ClaimReserve/create/create.component';
import {EditClaimReserveComponent} from './components/ClaimReserve/edit/edit.component';
import {IndexClaimPaymentComponent} from './components/ClaimPayment/index/index.component';
import {CreateClaimPaymentComponent} from './components/ClaimPayment/create/create.component';
import {EditClaimPaymentComponent} from './components/ClaimPayment/edit/edit.component';
import {IndexServiceProviderComponent} from './components/ServiceProvider/index/index.component';
import {CreateServiceProviderComponent} from './components/ServiceProvider/create/create.component';
import {EditServiceProviderComponent} from './components/ServiceProvider/edit/edit.component';
import {IndexReinsuranceAgreementComponent} from './components/ReinsuranceAgreement/index/index.component';
import {CreateReinsuranceAgreementComponent} from './components/ReinsuranceAgreement/create/create.component';
import {EditReinsuranceAgreementComponent} from './components/ReinsuranceAgreement/edit/edit.component';
import {IndexSubrogationRecoveryComponent} from './components/SubrogationRecovery/index/index.component';
import {CreateSubrogationRecoveryComponent} from './components/SubrogationRecovery/create/create.component';
import {EditSubrogationRecoveryComponent} from './components/SubrogationRecovery/edit/edit.component';
import {IndexThirdPartyComponent} from './components/ThirdParty/index/index.component';
import {CreateThirdPartyComponent} from './components/ThirdParty/create/create.component';
import {EditThirdPartyComponent} from './components/ThirdParty/edit/edit.component';
import {IndexDocumentComponent} from './components/Document/index/index.component';
import {CreateDocumentComponent} from './components/Document/create/create.component';
import {EditDocumentComponent} from './components/Document/edit/edit.component';

import * as appRoutes from './routerConfig';

import {InsurerService} from './services/Insurer.service';
import {InsuranceProductService} from './services/InsuranceProduct.service';
import {CoverageDefinitionService} from './services/CoverageDefinition.service';
import {DistributorService} from './services/Distributor.service';
import {AgentService} from './services/Agent.service';
import {CustomerService} from './services/Customer.service';
import {ApplicationService} from './services/Application.service';
import {QuoteService} from './services/Quote.service';
import {UnderwritingDecisionService} from './services/UnderwritingDecision.service';
import {UnderwriterService} from './services/Underwriter.service';
import {PolicyService} from './services/Policy.service';
import {EndorsementService} from './services/Endorsement.service';
import {PolicyCoverageService} from './services/PolicyCoverage.service';
import {InsuredObjectService} from './services/InsuredObject.service';
import {BeneficiaryService} from './services/Beneficiary.service';
import {BillingAccountService} from './services/BillingAccount.service';
import {InvoiceService} from './services/Invoice.service';
import {PaymentService} from './services/Payment.service';
import {ClaimService} from './services/Claim.service';
import {IncidentService} from './services/Incident.service';
import {ExposureService} from './services/Exposure.service';
import {AdjusterService} from './services/Adjuster.service';
import {ClaimReserveService} from './services/ClaimReserve.service';
import {ClaimPaymentService} from './services/ClaimPayment.service';
import {ServiceProviderService} from './services/ServiceProvider.service';
import {ReinsuranceAgreementService} from './services/ReinsuranceAgreement.service';
import {SubrogationRecoveryService} from './services/SubrogationRecovery.service';
import {ThirdPartyService} from './services/ThirdParty.service';
import {DocumentService} from './services/Document.service';

@NgModule({
  declarations: [
    IndexInsurerComponent,
    CreateInsurerComponent,
    EditInsurerComponent,
    IndexInsuranceProductComponent,
    CreateInsuranceProductComponent,
    EditInsuranceProductComponent,
    IndexCoverageDefinitionComponent,
    CreateCoverageDefinitionComponent,
    EditCoverageDefinitionComponent,
    IndexDistributorComponent,
    CreateDistributorComponent,
    EditDistributorComponent,
    IndexAgentComponent,
    CreateAgentComponent,
    EditAgentComponent,
    IndexCustomerComponent,
    CreateCustomerComponent,
    EditCustomerComponent,
    IndexApplicationComponent,
    CreateApplicationComponent,
    EditApplicationComponent,
    IndexQuoteComponent,
    CreateQuoteComponent,
    EditQuoteComponent,
    IndexUnderwritingDecisionComponent,
    CreateUnderwritingDecisionComponent,
    EditUnderwritingDecisionComponent,
    IndexUnderwriterComponent,
    CreateUnderwriterComponent,
    EditUnderwriterComponent,
    IndexPolicyComponent,
    CreatePolicyComponent,
    EditPolicyComponent,
    IndexEndorsementComponent,
    CreateEndorsementComponent,
    EditEndorsementComponent,
    IndexPolicyCoverageComponent,
    CreatePolicyCoverageComponent,
    EditPolicyCoverageComponent,
    IndexInsuredObjectComponent,
    CreateInsuredObjectComponent,
    EditInsuredObjectComponent,
    IndexBeneficiaryComponent,
    CreateBeneficiaryComponent,
    EditBeneficiaryComponent,
    IndexBillingAccountComponent,
    CreateBillingAccountComponent,
    EditBillingAccountComponent,
    IndexInvoiceComponent,
    CreateInvoiceComponent,
    EditInvoiceComponent,
    IndexPaymentComponent,
    CreatePaymentComponent,
    EditPaymentComponent,
    IndexClaimComponent,
    CreateClaimComponent,
    EditClaimComponent,
    IndexIncidentComponent,
    CreateIncidentComponent,
    EditIncidentComponent,
    IndexExposureComponent,
    CreateExposureComponent,
    EditExposureComponent,
    IndexAdjusterComponent,
    CreateAdjusterComponent,
    EditAdjusterComponent,
    IndexClaimReserveComponent,
    CreateClaimReserveComponent,
    EditClaimReserveComponent,
    IndexClaimPaymentComponent,
    CreateClaimPaymentComponent,
    EditClaimPaymentComponent,
    IndexServiceProviderComponent,
    CreateServiceProviderComponent,
    EditServiceProviderComponent,
    IndexReinsuranceAgreementComponent,
    CreateReinsuranceAgreementComponent,
    EditReinsuranceAgreementComponent,
    IndexSubrogationRecoveryComponent,
    CreateSubrogationRecoveryComponent,
    EditSubrogationRecoveryComponent,
    IndexThirdPartyComponent,
    CreateThirdPartyComponent,
    EditThirdPartyComponent,
    IndexDocumentComponent,
    CreateDocumentComponent,
    EditDocumentComponent,
    AppComponent
  ],
  imports: [

    BrowserModule, 
    NgbModule,
    MatMenuModule,
    MatToolbarModule,
    MatCheckboxModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatDatepickerModule,
	MatMomentDateModule,
    BrowserAnimationsModule,
	HttpClientModule, 
    ReactiveFormsModule,
    FormsModule,
    MatSidenavModule,    
    RouterModule.forRoot(appRoutes.InsurerRoutes), 
    RouterModule.forRoot(appRoutes.InsuranceProductRoutes), 
    RouterModule.forRoot(appRoutes.CoverageDefinitionRoutes), 
    RouterModule.forRoot(appRoutes.DistributorRoutes), 
    RouterModule.forRoot(appRoutes.AgentRoutes), 
    RouterModule.forRoot(appRoutes.CustomerRoutes), 
    RouterModule.forRoot(appRoutes.ApplicationRoutes), 
    RouterModule.forRoot(appRoutes.QuoteRoutes), 
    RouterModule.forRoot(appRoutes.UnderwritingDecisionRoutes), 
    RouterModule.forRoot(appRoutes.UnderwriterRoutes), 
    RouterModule.forRoot(appRoutes.PolicyRoutes), 
    RouterModule.forRoot(appRoutes.EndorsementRoutes), 
    RouterModule.forRoot(appRoutes.PolicyCoverageRoutes), 
    RouterModule.forRoot(appRoutes.InsuredObjectRoutes), 
    RouterModule.forRoot(appRoutes.BeneficiaryRoutes), 
    RouterModule.forRoot(appRoutes.BillingAccountRoutes), 
    RouterModule.forRoot(appRoutes.InvoiceRoutes), 
    RouterModule.forRoot(appRoutes.PaymentRoutes), 
    RouterModule.forRoot(appRoutes.ClaimRoutes), 
    RouterModule.forRoot(appRoutes.IncidentRoutes), 
    RouterModule.forRoot(appRoutes.ExposureRoutes), 
    RouterModule.forRoot(appRoutes.AdjusterRoutes), 
    RouterModule.forRoot(appRoutes.ClaimReserveRoutes), 
    RouterModule.forRoot(appRoutes.ClaimPaymentRoutes), 
    RouterModule.forRoot(appRoutes.ServiceProviderRoutes), 
    RouterModule.forRoot(appRoutes.ReinsuranceAgreementRoutes), 
    RouterModule.forRoot(appRoutes.SubrogationRecoveryRoutes), 
    RouterModule.forRoot(appRoutes.ThirdPartyRoutes), 
    RouterModule.forRoot(appRoutes.DocumentRoutes), 
  ],
  providers: [InsurerService,InsuranceProductService,CoverageDefinitionService,DistributorService,AgentService,CustomerService,ApplicationService,QuoteService,UnderwritingDecisionService,UnderwriterService,PolicyService,EndorsementService,PolicyCoverageService,InsuredObjectService,BeneficiaryService,BillingAccountService,InvoiceService,PaymentService,ClaimService,IncidentService,ExposureService,AdjusterService,ClaimReserveService,ClaimPaymentService,ServiceProviderService,ReinsuranceAgreementService,SubrogationRecoveryService,ThirdPartyService,DocumentService],
  bootstrap: [AppComponent]
})
export class AppModule { }
