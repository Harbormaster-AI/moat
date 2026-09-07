// routerConfig.ts

import { Routes } from '@angular/router';
import { CreateOrganizationComponent } from './components/Organization/create/create.component';
import { EditOrganizationComponent } from './components/Organization/edit/edit.component';
import { IndexOrganizationComponent } from './components/Organization/index/index.component';
import { CreateUserComponent } from './components/User/create/create.component';
import { EditUserComponent } from './components/User/edit/edit.component';
import { IndexUserComponent } from './components/User/index/index.component';
import { CreateTeamComponent } from './components/Team/create/create.component';
import { EditTeamComponent } from './components/Team/edit/edit.component';
import { IndexTeamComponent } from './components/Team/index/index.component';
import { CreateTerritoryComponent } from './components/Territory/create/create.component';
import { EditTerritoryComponent } from './components/Territory/edit/edit.component';
import { IndexTerritoryComponent } from './components/Territory/index/index.component';
import { CreateAccountComponent } from './components/Account/create/create.component';
import { EditAccountComponent } from './components/Account/edit/edit.component';
import { IndexAccountComponent } from './components/Account/index/index.component';
import { CreateContactComponent } from './components/Contact/create/create.component';
import { EditContactComponent } from './components/Contact/edit/edit.component';
import { IndexContactComponent } from './components/Contact/index/index.component';
import { CreateLeadComponent } from './components/Lead/create/create.component';
import { EditLeadComponent } from './components/Lead/edit/edit.component';
import { IndexLeadComponent } from './components/Lead/index/index.component';
import { CreateOpportunityComponent } from './components/Opportunity/create/create.component';
import { EditOpportunityComponent } from './components/Opportunity/edit/edit.component';
import { IndexOpportunityComponent } from './components/Opportunity/index/index.component';
import { CreateOpportunityLineItemComponent } from './components/OpportunityLineItem/create/create.component';
import { EditOpportunityLineItemComponent } from './components/OpportunityLineItem/edit/edit.component';
import { IndexOpportunityLineItemComponent } from './components/OpportunityLineItem/index/index.component';
import { CreateOpportunityStageHistoryComponent } from './components/OpportunityStageHistory/create/create.component';
import { EditOpportunityStageHistoryComponent } from './components/OpportunityStageHistory/edit/edit.component';
import { IndexOpportunityStageHistoryComponent } from './components/OpportunityStageHistory/index/index.component';
import { CreateProductComponent } from './components/Product/create/create.component';
import { EditProductComponent } from './components/Product/edit/edit.component';
import { IndexProductComponent } from './components/Product/index/index.component';
import { CreatePriceBookComponent } from './components/PriceBook/create/create.component';
import { EditPriceBookComponent } from './components/PriceBook/edit/edit.component';
import { IndexPriceBookComponent } from './components/PriceBook/index/index.component';
import { CreatePriceBookEntryComponent } from './components/PriceBookEntry/create/create.component';
import { EditPriceBookEntryComponent } from './components/PriceBookEntry/edit/edit.component';
import { IndexPriceBookEntryComponent } from './components/PriceBookEntry/index/index.component';
import { CreateQuoteComponent } from './components/Quote/create/create.component';
import { EditQuoteComponent } from './components/Quote/edit/edit.component';
import { IndexQuoteComponent } from './components/Quote/index/index.component';
import { CreateQuoteLineItemComponent } from './components/QuoteLineItem/create/create.component';
import { EditQuoteLineItemComponent } from './components/QuoteLineItem/edit/edit.component';
import { IndexQuoteLineItemComponent } from './components/QuoteLineItem/index/index.component';
import { CreateOrderComponent } from './components/Order/create/create.component';
import { EditOrderComponent } from './components/Order/edit/edit.component';
import { IndexOrderComponent } from './components/Order/index/index.component';
import { CreateOrderItemComponent } from './components/OrderItem/create/create.component';
import { EditOrderItemComponent } from './components/OrderItem/edit/edit.component';
import { IndexOrderItemComponent } from './components/OrderItem/index/index.component';
import { CreateContractComponent } from './components/Contract/create/create.component';
import { EditContractComponent } from './components/Contract/edit/edit.component';
import { IndexContractComponent } from './components/Contract/index/index.component';
import { CreateCase_Component } from './components/Case_/create/create.component';
import { EditCase_Component } from './components/Case_/edit/edit.component';
import { IndexCase_Component } from './components/Case_/index/index.component';
import { CreateActivityComponent } from './components/Activity/create/create.component';
import { EditActivityComponent } from './components/Activity/edit/edit.component';
import { IndexActivityComponent } from './components/Activity/index/index.component';
import { CreateCampaignComponent } from './components/Campaign/create/create.component';
import { EditCampaignComponent } from './components/Campaign/edit/edit.component';
import { IndexCampaignComponent } from './components/Campaign/index/index.component';
import { CreateCampaignMemberComponent } from './components/CampaignMember/create/create.component';
import { EditCampaignMemberComponent } from './components/CampaignMember/edit/edit.component';
import { IndexCampaignMemberComponent } from './components/CampaignMember/index/index.component';
import { CreateNoteComponent } from './components/Note/create/create.component';
import { EditNoteComponent } from './components/Note/edit/edit.component';
import { IndexNoteComponent } from './components/Note/index/index.component';
import { CreateEmailMessageComponent } from './components/EmailMessage/create/create.component';
import { EditEmailMessageComponent } from './components/EmailMessage/edit/edit.component';
import { IndexEmailMessageComponent } from './components/EmailMessage/index/index.component';

export const OrganizationRoutes: Routes = [
  { path: 'createOrganization',
    component: CreateOrganizationComponent
  },
  {
    path: 'editOrganization/:id',
    component: EditOrganizationComponent
  },
  { path: 'indexOrganization',
    component: IndexOrganizationComponent
  }
];
export const UserRoutes: Routes = [
  { path: 'createUser',
    component: CreateUserComponent
  },
  {
    path: 'editUser/:id',
    component: EditUserComponent
  },
  { path: 'indexUser',
    component: IndexUserComponent
  }
];
export const TeamRoutes: Routes = [
  { path: 'createTeam',
    component: CreateTeamComponent
  },
  {
    path: 'editTeam/:id',
    component: EditTeamComponent
  },
  { path: 'indexTeam',
    component: IndexTeamComponent
  }
];
export const TerritoryRoutes: Routes = [
  { path: 'createTerritory',
    component: CreateTerritoryComponent
  },
  {
    path: 'editTerritory/:id',
    component: EditTerritoryComponent
  },
  { path: 'indexTerritory',
    component: IndexTerritoryComponent
  }
];
export const AccountRoutes: Routes = [
  { path: 'createAccount',
    component: CreateAccountComponent
  },
  {
    path: 'editAccount/:id',
    component: EditAccountComponent
  },
  { path: 'indexAccount',
    component: IndexAccountComponent
  }
];
export const ContactRoutes: Routes = [
  { path: 'createContact',
    component: CreateContactComponent
  },
  {
    path: 'editContact/:id',
    component: EditContactComponent
  },
  { path: 'indexContact',
    component: IndexContactComponent
  }
];
export const LeadRoutes: Routes = [
  { path: 'createLead',
    component: CreateLeadComponent
  },
  {
    path: 'editLead/:id',
    component: EditLeadComponent
  },
  { path: 'indexLead',
    component: IndexLeadComponent
  }
];
export const OpportunityRoutes: Routes = [
  { path: 'createOpportunity',
    component: CreateOpportunityComponent
  },
  {
    path: 'editOpportunity/:id',
    component: EditOpportunityComponent
  },
  { path: 'indexOpportunity',
    component: IndexOpportunityComponent
  }
];
export const OpportunityLineItemRoutes: Routes = [
  { path: 'createOpportunityLineItem',
    component: CreateOpportunityLineItemComponent
  },
  {
    path: 'editOpportunityLineItem/:id',
    component: EditOpportunityLineItemComponent
  },
  { path: 'indexOpportunityLineItem',
    component: IndexOpportunityLineItemComponent
  }
];
export const OpportunityStageHistoryRoutes: Routes = [
  { path: 'createOpportunityStageHistory',
    component: CreateOpportunityStageHistoryComponent
  },
  {
    path: 'editOpportunityStageHistory/:id',
    component: EditOpportunityStageHistoryComponent
  },
  { path: 'indexOpportunityStageHistory',
    component: IndexOpportunityStageHistoryComponent
  }
];
export const ProductRoutes: Routes = [
  { path: 'createProduct',
    component: CreateProductComponent
  },
  {
    path: 'editProduct/:id',
    component: EditProductComponent
  },
  { path: 'indexProduct',
    component: IndexProductComponent
  }
];
export const PriceBookRoutes: Routes = [
  { path: 'createPriceBook',
    component: CreatePriceBookComponent
  },
  {
    path: 'editPriceBook/:id',
    component: EditPriceBookComponent
  },
  { path: 'indexPriceBook',
    component: IndexPriceBookComponent
  }
];
export const PriceBookEntryRoutes: Routes = [
  { path: 'createPriceBookEntry',
    component: CreatePriceBookEntryComponent
  },
  {
    path: 'editPriceBookEntry/:id',
    component: EditPriceBookEntryComponent
  },
  { path: 'indexPriceBookEntry',
    component: IndexPriceBookEntryComponent
  }
];
export const QuoteRoutes: Routes = [
  { path: 'createQuote',
    component: CreateQuoteComponent
  },
  {
    path: 'editQuote/:id',
    component: EditQuoteComponent
  },
  { path: 'indexQuote',
    component: IndexQuoteComponent
  }
];
export const QuoteLineItemRoutes: Routes = [
  { path: 'createQuoteLineItem',
    component: CreateQuoteLineItemComponent
  },
  {
    path: 'editQuoteLineItem/:id',
    component: EditQuoteLineItemComponent
  },
  { path: 'indexQuoteLineItem',
    component: IndexQuoteLineItemComponent
  }
];
export const OrderRoutes: Routes = [
  { path: 'createOrder',
    component: CreateOrderComponent
  },
  {
    path: 'editOrder/:id',
    component: EditOrderComponent
  },
  { path: 'indexOrder',
    component: IndexOrderComponent
  }
];
export const OrderItemRoutes: Routes = [
  { path: 'createOrderItem',
    component: CreateOrderItemComponent
  },
  {
    path: 'editOrderItem/:id',
    component: EditOrderItemComponent
  },
  { path: 'indexOrderItem',
    component: IndexOrderItemComponent
  }
];
export const ContractRoutes: Routes = [
  { path: 'createContract',
    component: CreateContractComponent
  },
  {
    path: 'editContract/:id',
    component: EditContractComponent
  },
  { path: 'indexContract',
    component: IndexContractComponent
  }
];
export const Case_Routes: Routes = [
  { path: 'createCase_',
    component: CreateCase_Component
  },
  {
    path: 'editCase_/:id',
    component: EditCase_Component
  },
  { path: 'indexCase_',
    component: IndexCase_Component
  }
];
export const ActivityRoutes: Routes = [
  { path: 'createActivity',
    component: CreateActivityComponent
  },
  {
    path: 'editActivity/:id',
    component: EditActivityComponent
  },
  { path: 'indexActivity',
    component: IndexActivityComponent
  }
];
export const CampaignRoutes: Routes = [
  { path: 'createCampaign',
    component: CreateCampaignComponent
  },
  {
    path: 'editCampaign/:id',
    component: EditCampaignComponent
  },
  { path: 'indexCampaign',
    component: IndexCampaignComponent
  }
];
export const CampaignMemberRoutes: Routes = [
  { path: 'createCampaignMember',
    component: CreateCampaignMemberComponent
  },
  {
    path: 'editCampaignMember/:id',
    component: EditCampaignMemberComponent
  },
  { path: 'indexCampaignMember',
    component: IndexCampaignMemberComponent
  }
];
export const NoteRoutes: Routes = [
  { path: 'createNote',
    component: CreateNoteComponent
  },
  {
    path: 'editNote/:id',
    component: EditNoteComponent
  },
  { path: 'indexNote',
    component: IndexNoteComponent
  }
];
export const EmailMessageRoutes: Routes = [
  { path: 'createEmailMessage',
    component: CreateEmailMessageComponent
  },
  {
    path: 'editEmailMessage/:id',
    component: EditEmailMessageComponent
  },
  { path: 'indexEmailMessage',
    component: IndexEmailMessageComponent
  }
];
