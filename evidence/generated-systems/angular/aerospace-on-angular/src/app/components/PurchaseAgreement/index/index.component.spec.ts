
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPurchaseAgreementComponent } from './index.component';
import { PurchaseAgreementService } from '../../../services/PurchaseAgreement.service';

describe('IndexPurchaseAgreementComponent', () => {
  let component: IndexPurchaseAgreementComponent;
  let fixture: ComponentFixture<IndexPurchaseAgreementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPurchaseAgreementComponent
      ],
      providers: [
        PurchaseAgreementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPurchaseAgreementComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});