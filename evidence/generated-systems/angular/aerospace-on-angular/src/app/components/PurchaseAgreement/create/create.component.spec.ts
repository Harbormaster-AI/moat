
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePurchaseAgreementComponent } from './create.component';
import { PurchaseAgreementService } from '../../../services/PurchaseAgreement.service';
import { Router } from '@angular/router';

describe('CreatePurchaseAgreementComponent', () => {
  let component: CreatePurchaseAgreementComponent;
  let fixture: ComponentFixture<CreatePurchaseAgreementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePurchaseAgreementComponent
      ],
      providers: [
        PurchaseAgreementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePurchaseAgreementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});