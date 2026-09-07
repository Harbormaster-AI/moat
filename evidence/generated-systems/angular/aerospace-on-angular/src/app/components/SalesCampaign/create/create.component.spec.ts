
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSalesCampaignComponent } from './create.component';
import { SalesCampaignService } from '../../../services/SalesCampaign.service';
import { Router } from '@angular/router';

describe('CreateSalesCampaignComponent', () => {
  let component: CreateSalesCampaignComponent;
  let fixture: ComponentFixture<CreateSalesCampaignComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSalesCampaignComponent
      ],
      providers: [
        SalesCampaignService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSalesCampaignComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});