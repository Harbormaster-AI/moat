
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSalesCampaignComponent } from './index.component';
import { SalesCampaignService } from '../../../services/SalesCampaign.service';

describe('IndexSalesCampaignComponent', () => {
  let component: IndexSalesCampaignComponent;
  let fixture: ComponentFixture<IndexSalesCampaignComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSalesCampaignComponent
      ],
      providers: [
        SalesCampaignService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSalesCampaignComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});