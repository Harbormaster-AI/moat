
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexClaimPaymentComponent } from './index.component';
import { ClaimPaymentService } from '../../../services/ClaimPayment.service';

describe('IndexClaimPaymentComponent', () => {
  let component: IndexClaimPaymentComponent;
  let fixture: ComponentFixture<IndexClaimPaymentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexClaimPaymentComponent
      ],
      providers: [
        ClaimPaymentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexClaimPaymentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});